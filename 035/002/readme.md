# Hands-on exercise #2
Start with [this code](https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/02/01-code-starting). Get the code ready to BET on (add benchmarks, examples, tests) however
do not write an example for the func that returns a map; and only write a test for it as an extra
challenge. Add documentation to the code. Run the following in this order:

- tests
- benchmarks
- coverage
- coverage shown in web browser
- examples shown in documentation in a web browser

## Solution
### Basic exercise requirement

```shell
002> go test ./...
?       exercise2       [no test files]
?       exercise2/quote [no test files]
ok      exercise2/word  0.632s
```

```shell
word> go test -bench . 
goos: windows
goarch: amd64
pkg: exercise2/word
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkCount-24       32298088                36.31 ns/op
PASS
ok      exercise2/word  1.754s
```

```shell
word> go test -cover . 
ok      exercise2/word  0.230s  coverage: 28.6% of statements
```

```shell
word> go test -coverprofile c.out
PASS
coverage: 28.6% of statements
ok      exercise2/word  0.214s
word> go tool cover -html c.out
word> 
```
```shell
002> godoc -http=:8080
using module mode; GOMOD=D:\Projectes\Cursos\GolangProgrammingLanguage\GoProgLangExercises\035\002\go.mod
go: no module dependencies to download
```
### Advanced exercise requirement

```shell
002> go test ./...
?       exercise2       [no test files]
?       exercise2/quote [no test files]
ok      exercise2/word  0.445s
```

```shell
word> go test -bench .
goos: windows
goarch: amd64
pkg: exercise2/word
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkCount-24               31770741                36.62 ns/op
BenchmarkUseCount-24             9454125               128.8 ns/op
PASS
ok      exercise2/word  3.033s
```

```shell
word> go test -cover .
ok      exercise2/word  0.012s  coverage: 100.0% of statements
```

```shell
word> go test -coverprofile c.out
PASS
coverage: 100.0% of statements
ok      exercise2/word  0.032s
word> go tool cover -html c.out
```

```shell
word> go test -coverprofile c.out
PASS
coverage: 100.0% of statements
ok      exercise2/word  0.032s
word> go tool cover -html c.out
```

```shell
002> godoc -http=:8080
using module mode; GOMOD=D:\Projectes\Cursos\GolangProgrammingLanguage\GoProgLangExercises\035\002\go.mod
go: no module dependencies to download
```
