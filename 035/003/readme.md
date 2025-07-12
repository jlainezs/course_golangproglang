# Hands-on exercise #3

Start with this code. Get the code ready to BET on (add benchmarks, examples, tests). Write a
table test. Add documentation to the code. Run the following in this order:

- tests
- benchmarks
- coverage
- coverage shown in web browser
- examples shown in documentation in a web browser

## Solution

```shell
003> go test ./...
?       exercise3       [no test files]
ok      exercise3/mymath        0.011s
```

```shell
mymath> go test -bench . 
goos: windows
goarch: amd64
pkg: exercise3/mymath
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkCenteredAvg-24         100000000               10.09 ns/op
PASS
ok      exercise3/mymath        1.033s
```

```shell
mymath> go test -cover
PASS
coverage: 100.0% of statements
ok      exercise3/mymath        0.013s
```

```shell
mymath> go test -coverprofile c.out
PASS
coverage: 100.0% of statements
ok      exercise3/mymath        0.013s
mymath> go tool cover -html c.out
```

```shell
mymath> go test -coverprofile c.out
PASS
coverage: 100.0% of statements
ok      exercise3/mymath        0.013s
mymath> go tool cover -html c.out
```

