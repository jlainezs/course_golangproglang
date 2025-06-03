- create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your program should do the following
  - create a random int between 0 and 250
  - store the value of that int in a variable with the identifier of x
    - func Intn(n int) int
      - rand.Intn()
  - print out the name and value of the variable
  - use an if statement to do the following
    - if the value is between 0 and 100
      - print between 0 and 100
    - if the value is between 101 and 200
      - print between 101 and 200
    - if the value is between 201 and 250
      - print between 201 and 250
  - re: inclusive, exclusive – does rand.Intn()
    - include zero in its output?
    - include 250 in its output?
    - show this in code using the numbers 0 - 3
  - hint:
    - &&
  

When executing the program, the antivirus may block it. 
- Sign the exe, before executing it
- Pause the antivirus
- Add this exe file to the antivirus exclusions list

The antivirus blocks the execution if we use `println` instead of `fmt.Println`

# Sign exe

On Windows, install make with chocolatery `choco install make` and sign the executable [See here](https://stackoverflow.com/questions/51717409/is-there-any-way-to-sign-the-windows-executables-generated-by-the-go-compiler)

```
# golang makefile based on https://golangdocs.com/makefiles-golang
BINARY_NAME=mygoproject.exe
 
build:
    go build -o ${BINARY_NAME} main.go

    # This runs signtool with a cert in your profile store instead of a *.pfx file, 
    # to avoid needing to store a password in the makefile or environment variable: 
    # https://stackoverflow.com/questions/26998439/signtool-with-certificate-stored-in-local-computer
    signtool sign /sm /s My /n <certificateSubjectName> /t http://timestamp.digicert.com ${BINARY_NAME}
 
run:
    go build -o ${BINARY_NAME} main.go
    ./${BINARY_NAME}
 
clean:
    go clean
    rm ${BINARY_NAME}
```

Run `make build` on a terminal.
