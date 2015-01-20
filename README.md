# gore – helper for error with details

### Usage

```go
err := gore.New("hello")
gore.Append("Some context message")
gore.Appendf("Another formatted context massage %s", "for define this error")

gerr := err.(*Err)
log.Print(gerr.Caller.ShortFileName())
```

See for examples `gore_test.go`

### Install

```shell
go get github.com/kavkaz/gore
```
