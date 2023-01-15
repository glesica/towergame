# Tower Game

## Development

Run the game:

```shell
go run ./app
```

## Tower Control Language

Below is an incomplete description of language features and conventions
we can consider implementing.

### Execution Model

A virtual machine will execute the program running on each tower. A certain
number of instructions will be executed, based on the configuration of the
tower's "computer" and possibly other game-world factors. Then, the state of the
virtual machine after the instructions have been executed will be converted by
the game engine to tower updates.

For example, there might be a global variable or register within the VM that
determines whether the tower should rotate left, right, or not at all. The state
of this variable after execution will be read externally and used to decide how
to update the tower in-game.

