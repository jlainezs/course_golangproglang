# Hands-on exercise #3

- Using goroutines, create an incrementer program
  - have a variable to hold the incrementer value
  - launch a bunch of goroutines
    - each goroutine should
      - read the incrementer value
        - store it in a new variable
      - yield the processor with runtime.Gosched()
      - increment the new variable
      - write the value in the new variable back to the incrementer
        variable
- use waitgroups to wait for all of your goroutines to finish
- the above will create a race condition.
- Prove that it is a race condition by using the -race flag
- if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

# Enable CGO

## Powershell

First, install gcc C compiler, otherwise this message will appear:
```
# runtime/cgo
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
```

Second, enable CGO by setting an environment variable:
```PowerShell
> $env:CGO_ENABLED = 1
```
