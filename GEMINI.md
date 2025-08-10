# GEMINI.md
This file provides guidance to GEMINI when working with code in this repository.

## Commands

### Build
To build the application:
```sh
go build
```

### Run
To run the application:
```sh
go run main.go
```
The server will start on port 80 by default, but it can be configured via the `PORT` environment variable.

### Test
To run all tests:
```sh
go test ./...
```

To run a single test file:
```sh
go test [path to test file]
```
For example:
```sh
go test model/Reversi_test.go
```

## Code Architecture

This is a web application written in Go that implements the game of Reversi.

### Backend
The backend is written in Go and is located in the root directory and the `model` directory.
- `main.go`: This is the entry point of the application. It starts a web server and handles requests.
- `model/`: This directory contains the core logic of the Reversi game.
    - `Reversi.go`: Contains the main game logic.
    - `ReversiPlay.go`: Handles the game play flow.
    - `ReversiSetting.go`: Manages the game settings.
    - The other files in this directory define data structures and interfaces for the game.

The server exposes a single API endpoint `/FrontController` that accepts commands to interact with the game. The game state is stored in a server-side session.

### Frontend
The frontend is located in the `static/` directory and consists of HTML, CSS, and JavaScript.
- `index.html`: The main page of the application.
- `js/func.js`: Contains the JavaScript logic for interacting with the backend.
- `css/style.css`: Contains the custom styles for the application.
