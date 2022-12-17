# Tower Game

## Development

Run the game:

```shell
go run ./app
```

## Tower Control Language

Below is an incomplete description of language features and conventions
we can consider implementing.

### Asynchronous Operations

There are a lot of asynchronous operations that will occur because the
language is, effectively, controlling a robot that exists in real time.
We need a way to pause program execution until the tower has completed
a given operation (or reached a given state).

For example, we are likely
to end up with a way to aim a tower in a particular direction. A tower
will require some time before it is finally aimed in the correct
direction.

We could handle this sort of thing by making asynchronous
operations return the amount of time they require and giving the user
a `sleep` command, or by introducing the concept of futures with an
`await` instruction. A third option would be to provide an easy way to
poll the tower, or environment, state.

`sleep <double>` - pause program execution for the given number of
seconds. This might be valuable to have either way. It is simple, but
also imprecise and doesn't allow for operations whose duration can't
be predicted ahead of time.

`await <Future>` - pause program execution until the given `Future` has
completed. This is probably the simplest from a player standpoint, but
less flexible than the option below, and more difficult to implement.

`waitFor <bool>` - pause program execution until the given boolean
expression evaluates to `true`. One challenge here is that it would be
pretty easy to write a condition that never terminates (like equality
testing two doubles).
