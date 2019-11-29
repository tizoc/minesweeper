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

**TODO**

## POST /minesweeper/games/{game-id}/plays

**TODO**
