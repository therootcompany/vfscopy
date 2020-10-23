# vfscopy

Copy a Virtual FileSystem, such as
[http.FileSystem](https://golang.org/pkg/net/http/#FileSystem),
recursively to a native file system destination.

```go
httpfs := http.Dir("/tmp/public/")
vfs := vfscopy.NewVFS(httpfs)

if err := Copy(vfs, ".", "/tmp/dst/"); nil != err {
    fmt.Fprintf(os.Stderr, "couldn't copy vfs: %v\n", err)
}
```

## Test

```bash
go generate ./...
go test ./...
```
