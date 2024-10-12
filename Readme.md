# GoLang Tutorial

This project is a set of code which is used to understand golang and it's working while doing the golang course from freecodecamp and boot.dev.

## Run test case

### To run the root test cases

```
go test
```

### To run all the test cases in sub-folders

```
go test ./...
```

### To run specific folder level test case

```
go test ./folder-name
```

To run the tests in verbose mode, add the `-v` to the command.

### To run specific test cases, use `-run [TestCaseName]`

```
go test ./concurrency -v -run TestRangeInChannel
```
