# Weeks

Calculate my age in weeks.

> "Missing out on something -- indeed, on almost everything -- is basically guaranteed. Which isn’t actually a problem anyway, it turns out, because 'missing out' is what makes our choices meaningful in the first place."
> -- Four Thousand Weeks, Oliver Burkeman

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

## Learnings

### Time zone abbreviations

Time zone abbreviations such as `CST` are ambiguous.[^1] It's no trouble to display them, but parsing them doesn't work quite right. Oddly, the Go standard library makes an attempt at it, instead of rejecting layouts containing MST.

```go
badTime, err := time.Parse("2006-01-02 3:04 PM (MST)", "1977-04-05 11:58 AM (PST)")
if err != nil {
    panic(err)
}
fmt.Println("bad", badTime)

location, err := time.LoadLocation("America/Vancouver")
if err != nil {
    panic(err)
}
goodTime, err := time.ParseInLocation("2006-01-02 3:04 PM", "1977-04-05 11:58 AM", location)
if err != nil {
    panic(err)
}
fmt.Println("good", goodTime)
```

This code[^2] produces the following result:

```
bad 1977-04-05 11:58:00 +0000 PST
good 1977-04-05 11:58:00 -0800 PST
```

Too bad the time zone offset is completely wrong! You can imagine how displaying the time with a format string would hide the offset and make it seem like everything was working (yes, that happened). Except the duration calculations were off.

At least it's well documented:

> "If the zone abbreviation is unknown, Parse records... the given zone abbreviation and a zero offset."[^3]

When I first implemented this, I didn't read the documentation carefully enough. My bad. After tracking down the bug, I landed on a GitHub Issue.

> "It is not a goal that time.Time.Format and time.Parse be exact reverses of each other."[^4]

Switching to `time.ParseInLocation` solved all the things.

[^1]: [Wikipedia](https://en.wikipedia.org/wiki/List_of_time_zone_abbreviations)
[^2]: [Go Playground](https://go.dev/play/p/8gQYa00Yv2o)) for time zone parsing
[^3]: [time.Parse documentation](https://pkg.go.dev/time#Parse)
[^4]: [GitHub Issue](https://github.com/golang/go/issues/24071)

### DivMod

My initial implementation of splitting a duration into weeks, days, hours was rather cumbersome. I hit upon the idea of taking the remainder to simplify the code.

The [math.Remainder](https://pkg.go.dev/math@go1.22.0#Remainder) doesn't produce the result I was looking for. Don't be mislead by the name. Use [math.Mod](https://pkg.go.dev/math@go1.22.0#Mod) instead, it's just what I needed.[^5]

```go
fmt.Println(math.Remainder(14.0, 3.0)) // -1
fmt.Println(math.Mod(14.0, 3.0))       // 2
```

Modulus using `%` works for integers but not floating point numbers.

```
// fmt.Println(14.0 % 3.0)
// invalid operation: operator % not defined...
```

The Go standard library has a DivMod function for `big.Int`, but not for floating point. So I wrote a DivMod function and it cleaned up my code nicely.

[^5]: [Go Playground](https://go.dev/play/p/5TdKuPK7hVl) for math.Remainder

