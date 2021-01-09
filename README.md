# Press

Compression Algorithm

Support FileIO

- [x] Tar/UnTar
- [x] Gzip/UnGzip
- [x] Tgz/UnTgz

# Install

Use the alias "press". To use Press in your Go code:

```go
import "github.com/AkvicorEdwards/press"
```

To install Press in your $GOPATH:

```shell script
go get "github.com/AkvicorEdwards/press"
```

# API

```go
Tar(source []string, targetPath, targetFilename string)
TarToFileIO(sources []string, target *os.File)
TarToWriter(sources []string, target *tar.Writer) (err error)
UnTar(tarball, target string)
UnTarFromFileIO(tarball *os.File, target string)
UnTarFromReader(reader *tar.Reader, target string)
```

```go
Gzip(source, targetPath, targetFilename string)
GzipToFileIO(reader, target *os.File)
UnGzip(source, target string)
UnGzipFromFileIO(reader, target *os.File)
```

```go
Tgz(source []string, targetPath, targetFilename string)
TgzToFileIO(sources []string, target *os.File)
UnTgz(tarball, target string)
UnTgzFromFileIO(tarball *os.File, target string)
```
