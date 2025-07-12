# Hands-on exercise #1
Start with [this code](https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/01/starting-code). Get the code ready to BET on the code (add benchmarks, examples, tests).
Run the following in this order:
- tests
- benchmarks
- coverage
- coverage shown in web browser
- examples shown in documentation in a web browser

## Log
### Test

```shell
> go test ./...
?       exercise1       [no test files]
ok      exercise1/dog   0.583s
> 
```
### Benchmark
```shell
001> cd dog                                                                                                                                                                                                                                                                                             
dog> go test -bench .
goos: windows
goarch: amd64
pkg: exercise1/dog
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkYears-24               1000000000               0.2394 ns/op
BenchmarkYearsTwo-24            385373780                3.115 ns/op
PASS
ok      exercise1/dog   2.298s
```

## Coverage
To run coverage, on Win11 with McAfee installed, you should disable realtime scanning
as it raises a false positive when launching the coverage tool and the command crashes.

```shell
001> go test -cover ./...
        exercise1               coverage: 0.0% of statements
ok      exercise1/dog   0.013s  coverage: 100.0% of statements
```

> Note: in real life expect coverage to be in the [70,80( interval.


### Coverage in browser

```shell
dog> go test -coverprofile c.out
PASS
coverage: 100.0% of statements
ok      exercise1/dog   0.014s
dog> go tool cover -html c.out  
```
>__NOTE__: do not use -html=c.out syntax as it is said in the course.

## Examples in documentation

```shell
001> godoc -http=:8080 
using module mode; GOMOD=D:\Projectes\Cursos\GolangProgrammingLanguage\GoProgLangExercises\035\001\go.mod
go: no module dependencies to download
```

Now browse http://localhost:8080, and look for package exercise1/dog
