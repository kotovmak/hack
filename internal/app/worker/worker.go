package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"hack/internal/app/config"
	"hack/internal/app/model"
	"hack/internal/app/store"
	"hack/internal/app/websocket"
	"log"
	"os"
	"os/exec"
	"time"

	"golang.org/x/sync/errgroup"
)

type Worker struct {
	cfg      config.Config
	store    store.Store
	recordCh chan model.File
	ws       *websocket.WS
}

func New(cfg config.Config, s store.Store, ws *websocket.WS) *Worker {
	return &Worker{
		cfg:      cfg,
		store:    s,
		recordCh: make(chan model.File),
		ws:       ws,
	}
}

func (w *Worker) Init(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	for i := 0; i < w.cfg.NumOfWorkers; i++ {
		g.Go(func() error {
			var err error
			for ch := range w.recordCh {
				err = w.handle(ctx, ch)
			}
			return err
		})
	}
	return nil
}

func (w *Worker) handle(ctx context.Context, f model.File) error {
	// TODO Открыть файл и запустить скрипт обработки модели
	log.Printf("[WORKER] Start python job")
	c := exec.Command(
		"/usr/local/bin/python3",
		"/app/services/predict-loyal-city/main.py",
		"-t",
		"/app/"+f.Name,
		"-c",
		"/app/services/predict-loyal-city/data/cities.csv",
		"-p",
		"/app/prediction_debit.json",
		"-pt",
		"debit",
	)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	f.SendAt = time.Now()
	err := w.store.File().Update(ctx, f)
	if err != nil {
		return err
	}

	w.ws.Clients.Range(func(key, value interface{}) bool {
		result := model.PredictResult{
			Status:  f.Status,
			Message: "",
			Data:    []model.RegionPredict{},
		}

		response, err := json.Marshal(result)
		if err != nil {
			log.Printf("key %s, error %s", key, err.Error())
			return false
		}

		value.(*websocket.Client).WriteMessage(string(response))
		return false
	})

	log.Printf("[WORKER] End python job")
	return nil
}

func (w *Worker) check(ctx context.Context, f model.File) error {
	// TODO Открыть файл результата и записать в базу
	localFile, err := os.ReadFile("prediction_debit.json")
	if err != nil {
		return nil
	}

	a := &model.PredictResult{}
	err = json.Unmarshal(localFile, a)
	if err != nil {
		log.Printf("file parsing err %s", err)
		return nil
	}

	if a.Status == "SUCCESS" || a.Status == "INVALID" {
		log.Printf("change status to %s", a.Status)
		f.Status = a.Status

		// update file table
		f.ReceivedAt = time.Now()
		if len(a.Message) > 0 {
			f.StatusMessage = a.Message
		}
		err := w.store.File().Update(ctx, f)
		if err != nil {
			return err
		}

		// update predict table
		err = w.store.Region().PredictListUpdate(ctx, a.Data)
		if err != nil {
			log.Printf("update table %s", err)
			return err
		}

		log.Printf("Start WS send")
		w.ws.Clients.Range(func(key, value interface{}) bool {
			log.Printf("send to client %s", a.Status)
			result := model.PredictResult{
				Status:  a.Status,
				Message: a.Message,
				Data:    a.Data,
			}

			response, err := json.Marshal(result)
			if err != nil {
				log.Printf("key %s, error %s", key, err.Error())
				return false
			}

			value.(*websocket.Client).WriteMessage(string(response))
			return false
		})
		log.Printf("End WS send")

	}
	return nil
}

func (w *Worker) Add(o model.File) {
	w.recordCh <- o
	log.Println("send File to chan")
}

func (w *Worker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			file, err := w.store.File().GetLast(ctx)
			if err != nil && err != sql.ErrNoRows {
				log.Println(err)
				return err
			}
			if err == nil && file.Status == "PROCESSED" {
				go w.check(ctx, file)
			}
			time.Sleep(3 * time.Second)
		}
	}
}
