# Minesweeper API

## Game state representation

**TODO**: document state object retuned by all the endpoints defined below

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
