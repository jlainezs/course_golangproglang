# Modules

## go mod

It is used to configure the workspace. It accepts a few commands, detailed below:

```
download    download modules to local cache
edit        edit go.mod from tools or scripts
graph       print module requirement graph
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed
```

Allows to manage modules.

After running ``go mod init <my_module>``, a ``go.mod`` file is created in the directory. This file will contain the module dependencies (wich are modules) which we can manage with ``go get`` commands.

The ``go.mod`` file contains the module name, the go version and a list of modules required by the module (if any). For each module it shows the repository and the module version:

```
module mymodule

go 1.20

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
```

## go get

Add module dependencies.

```
go get <repository_url@version>
```
Version is a version tag in the repository. Also we can use a commit hash. 

## Module versioning

Modules can be referenced using a commit refrence or a version tag. The recommended way to go is use version tags,
which are defined using ```git tag``` commands.

Versioning introduces the use of semantic versioning. Given a version numver like X.Y.Z :

- X is the major version. Notes that major changes are introduced. Backard compatibility can be broken.
- Y is the minor version. The version introduces minor changes. Application is backward compatible.
- Z is the patch version. The version introdudes bg fixes. Application is backward compatible.

## Module vs Package

A module can contain several packages.

## Resources

- [How to use Go modules](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)
- [Practical Go lessons, modules](https://www.practical-go-lessons.com/chap-17-go-modules)
- [Semantic versioning](https://semver.org)
