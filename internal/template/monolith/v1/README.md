# {{.BinaryName}}
{{.BinaryName}} is a monolith service which contains both, a backend server
providing a REST-API and a frontend application. The frontend reside in its own
independent [module](webapp/README.md).

## running
The simplest way to build and launch is using the make file.

```go
make run
```

## building
To just build for a release, you invoke:

```
make release
./.build/release/{{.BinaryName}}
```