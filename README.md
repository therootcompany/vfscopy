# vfscopy

Copy a Virtual FileSystem, such as
[http.FileSystem](https://golang.org/pkg/net/http/#FileSystem),
recursively to a native file system destination.

Works with any file system that implements http.FileSystem,
such as `vfsgen`, `fileb0x`, `gobindata`.

```go
httpfs := http.Dir("/tmp/public/")
vfs := vfscopy.NewVFS(httpfs)

if err := Copy(vfs, ".", "/tmp/dst/"); nil != err {
    fmt.Fprintf(os.Stderr, "couldn't copy vfs: %v\n", err)
}
```

## Test

```bash
# Generate the test virtual file system
go generate ./...

# Run the tests
go test ./...
```

# License

The MIT License (MIT)

We used the recursive native file system copy implementation at
https://github.com/otiai10/copy as a starting point and added
virtual file system support.
