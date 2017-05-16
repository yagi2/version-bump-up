# Android Version Bump Tool[WIP]
## Description
Management Android versionName and versionCode of app/build.gradle.  
It is not yet supported when it is replaced by a variable or managed by another file. (Issue #1)
  
## Feature
- auto increment versionCode. e.g. versionCode 5 to versionCode 6
- auto increment {major|minor|patch} versionName from each subcommand.

## Usage  
* major  
    - `$ version-bump-up major`
* minor  
    - `$ version-bump-up minor`
* patch  
    - `$ version-bump-up patch`
* set  
    - `$ version-bump-up set [versionName] [versionCode]`  
        - e.g. `$ version-bump-up set 1.0.1 2`
* check
    - `$ version-bump-up check`

## Install

To install, use `go get`:

```bash
$ go get -d github.com/yagi2/version-bump-up
```

## Contribution

1. Fork ([https://github.com/yagi2/version-bump-up/fork](https://github.com/yagi2/version-bump-up/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[yagi2](https://github.com/yagi2)

## todo
- add test code.
- support replaced by a variable and managed by another file.