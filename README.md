# Weeks

Calculate my age in weeks.

Inspired by Four Thousand Weeks by Oliver Burkeman.

Visit [my blog](https://nathany.com/4000) for an overview and my learnings from this exercise.

## Run

Build and run:

```console
go run weeks.go
```

Build then run:

```console
go build
./weeks
```

## Tests

Run the tests and include % test coverage:

```console
go test -cover 
```

Open test coverage in a web browser:

```console
go test -coverprofile=c.out
go tool cover -html="c.out"
```
