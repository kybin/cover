# cover

cover opens html report for go code coverage.

It essentially wraps these two commands.

```
go test -coverprofile=/tmp/coverage.out
go tool cover -html=/tmp/coverage.out
```

