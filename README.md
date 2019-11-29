# minesweeper-API

To build:

    make build

To run tests:

    make test:

To run:

    make run

## Notes

### What is supported

- Revealing a cell that does not contain nearby mines produces cascade reveal effect on nearby cells.
- Covered cells can be flagged.
- Time is tracked.
- Detects game finish (either won or lost).
- Multiple games supported.
- Configurable board size and amount of mines.

### What is pending

- Example client.
- This is basic-REST, no HATEOAS or even close.
- Error responses don't contain enough information.
- Some validation is missing when creating the game.
- Not parallel-safe, a possible solution would be to communicate each game instance through a channel to enforce plays to be sequential.
- No out-or-process persistence. Game state could be marshalled and stored in plain files, an external database user, or even the insterlal state encoded and encrypted and included as part of the HTTP replies to make everything stateless server-side (state representation is compact enough for this to be a viable option).