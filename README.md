# backend-engineer-app-josh
## Recommended
* Go is installed (1.17 or higher - see go-sqlite3 supported Go versions [here](https://github.com/mattn/go-sqlite3/blob/v1.14.15/.github/workflows/go.yaml)).
* Evironment variable `CGO_ENABLED` is set to `1`.

## Setup & Run
1. Clone this repo.
2. Open a terminal to start the server.
   * Set `PORT` environment variable (default port is 8080).
    * e.g. `export PORT=8080` or change default port directly in `.env`.
   * Run: `go build .`
   * Run: `go run .`
3. Open a new terminal to send a request.
   * Set `PORT` environment variable again to match value set above (in the terminal that's running the server).
   * Run: `curl localhost:$PORT/employees`.

## Ports
* To test different port values:
  1. Stop the server with ctrl+c
  2. Run: `export PORT=<port>`
    * The PORT env var set in the terminal will override the default value in `.env`
  3. Restart the server by running `go run .`
* To check the current port the server is configured with:
  1. Stop the server with ctrl+c
  2. Run: `env | grep PORT`
