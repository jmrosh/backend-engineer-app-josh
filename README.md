# backend-engineer-app-josh

## Required
* Go is installed.
* Go environment variable `CGO_ENABLED` is set to `1` (default).
  * To check, run: `go env | grep CGO_ENABLED`.

## Recommended
* Go Version 1.17 or higher - see go-sqlite3 supported Go versions [here](https://github.com/mattn/go-sqlite3/blob/v1.14.15/.github/workflows/go.yaml).
* These instructions were intended for mac/linux. If on windows, run your windows equivalent commands.

## Setup & Run
1. Open a new terminal.
2. Clone this repo.
3. Change directory to this repo.
   * e.g. `cd backend-engineer-app-josh`
4. Set `PORT` environment variable (default port is 8080).
     * e.g. `export PORT=8080` or change default port directly in `.env`.
5. Build and run.
   * Run: `go build .`
   * Run: `go run .`
6. Open a new terminal to send a request.
   * Set `PORT` environment variable again to match value set above (in the terminal that's running the server).
   * Run: `curl localhost:$PORT/employees`.

## Changing Port
* To test different port values:
  1. Stop the server with ctrl+c
  2. Run: `export PORT=<port>`
     * The PORT env var set in the terminal will override the default value in `.env`
  3. Update the PORT env var in the terminal which you send the request (curl command).
  4. Restart the server by running `go run .`
* To check the current port the server is configured with:
  1. Stop the server with ctrl+c
  2. Run: `env | grep PORT`
