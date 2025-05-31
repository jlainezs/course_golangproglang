# GOPATH

- check the $GOPATH/bin
  - echo $GOPATH
  - go env
- Navigate to $GOPATH/bin and show its contents
- install a program 
  - check it is in $GOPATH/bin


Check $GOPATH:
In powershell,

```
echo $Env:GOPATH
```

Mavigate to $GOPATH/bin

```
cd $Env:GOPATH/bin
ls
```
Install program:

```
go install .
```

On the other terminal, ls shows the applications installed.

Open a third terminal and execute

```
>appwrkspace
Hello, world!
>
```