# backend-engineer-app-josh
## Prerequisites
* Go is installed (1.17 or higher - see go-sqlite3 supported Go versions [here](https://github.com/mattn/go-sqlite3/blob/v1.14.15/.github/workflows/go.yaml)).
* Evironment variable `CGO_ENABLED` is set to `1`.

## Setup & Run
1. Clone this repo.
2. Set `PORT` environment variable (default port is 8080).
3. Run: `go build .`
4. Run: `go run .`
5. Open a new terminal and run: `curl localhost:$PORT/employees`.