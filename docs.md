# Minesweeper API

## Game state representation

**TODO**: document state object retuned by all the endpoints defined below

## Endpoints

## POST /minesweeper/games

Content-Type: application/json
Input: a json object with the following fields:

- height, an integer > 0 that defines the height of the board (defaults to 9 if not provided)
- width, an integer > 0 that defines the width of the board (defaults to 9 if not provided)
- bombs, an integer > 0 that defines the amount of bombs to be placed (defaults to 10 if not provided)

Output:

A json object with the game state as described above and deply code 201.


## GET /minesweeper/games/{game-id}

**TODO**

## POST /minesweeper/games/{game-id}/plays

**TODO**
