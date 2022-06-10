# hack
Moscow City Hack 2022

## Installation locally
### Prerequisites

- go 1.17
- golang-migrate cli (install on unix/macOs: brew install golang-migrate)

1. Clone this repository

2. Install dependencies
    ```bash
    $ go mod download
    $ migrate -path migrations -database "postgresql://user:password@database:5432/hack?sslmode=disable" up
    $ cp .env.example .env
    ```

## Run app local

1. Start server locally
    ```bash
    $ make
    ```
2. Open in browser 
    ```
    http://51.250.44.134/status
    ```
