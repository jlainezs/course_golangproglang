# Section 34

About testing

## Test launch error

When launching tests you may face this error:

```shell
> go test
# main.test
$WORK\b001\_testmain.go:14:8: could not import main (cannot import "main")
FAIL    main [build failed]
```

To solve it, I've changed the go.mod file content from

```
module main

go 1.24.3
```

to

```
module tst_Demo

go 1.24.3
```

Then launch tests again and now it works.

``` shell
> go test
PASS
ok      tst     0.540s
```

For details about what's going on [see this page](https://appliedgo.net/testmain/)
