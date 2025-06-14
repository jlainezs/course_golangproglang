# Note on executing

McAfee on Windows 11 reports a false positive when executing

`> go run main.go`

from the command line.

To solve this issue, run a build then launch the application.

```
> go build main.go
> ./main.exe
```
