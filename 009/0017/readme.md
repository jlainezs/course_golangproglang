# Build for specific platforms

- using previous code, build for windows
  - GOOS=windows go build
- build for mac
  - GOOS=darwin go build
- build for linux
  - GOOS=linux go build

Under powershell define the environment variables as it follows:

`$Env:GOOS="linux";$Env:GOARCH="amd64";go build main.go`