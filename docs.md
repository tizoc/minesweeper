# Minesweeper API

## Game state representation

The game state is represented by a json object with the following fields:

- `board` is an array of integers representing each cell in the board. Values >= 0 are for uncovered cells, and the value represents the amount of nearby mines. A value of -1 represents a bomb, a value of -127 represents a covered cell and a value of -126 represents a flag.
- `height` is the height of the board.
- `width` is the width of the board.
- `mines` is the total amount of mines in the board.
- `status` is one of `"NotStarted"`, `"Started"`, `"Won"` or `"Lost`".
- `startedAt` is the time at which the game started.
- `endedAt` is the time at which the game was won or lost.

The size of `board` is equal to `height * width`.

Example:

```json
{
    "board": [
        0, 0, 1, -1, 1, 0, /* .... */ -127
    ],
    "height": 9,
    "width": 9,
    "mines": 10,
    "status": "Lost",
    "startedAt": "2019-11-29T21:59:43.020639Z",
    "endedAt": "2019-11-29T21:59:56.694715Z"
}
```

## Endpoints

## POST /minesweeper/games

Content-Type: application/json
Input: a json object with the following fields:

- `height`, an integer > 0 that defines the height of the board (defaults to 9 if not provided)
- `width`, an integer > 0 that defines the width of the board (defaults to 9 if not provided)
- `mines`, an integer > 0 that defines the amount of mines to be placed (defaults to 10 if not provided)

Output:

A json object with a single key `"gameID"`, containing the ID value of the created game. Reply code 201.

## GET /minesweeper/games/{game-id}

Output:

A representation of the game state as described above it it exists. If not, a 404 reply is sent instead.

## POST /minesweeper/games/{game-id}/plays

Content-Type: application/json
Input: a json object with the following fields:

- `location`, a reference to one of the cells in the board (>= 0 integer)
- `action`, one of `"flag"` (toggles a flag in the given location) or `"uncover"` (uncovers the given location)

Output:

An update representation of the game state as described above it it exists with reply code 201 (play created). If not, a 404 reply is sent instead.
