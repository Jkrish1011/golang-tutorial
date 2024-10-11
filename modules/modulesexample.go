package modules

/*
A Go repository typically contains only one module, located at the root of the repository.
a file named go.mod at the root of a project declares the module. It contains:

1/ The module path.
2/ The version of Go lnaguage the project requires.
3/ Optionally, any external package dependencies the project has.

Note:
There is no common repository site for downloading go packages like npm.js is for javascript
GoLang hosts and pulls the other projects/packages from github

Packages
--------

All functions first character to be capitalized to be exportable.

example:

func Reverse(s string) string {
...
}

if it was reverse, it is not accessible outside this package.
Lowercase names aren't exported for external use.

Using Local Package

NOTE: This is for local testing and not for production. In production, push the package to github and access it from github.
inside go.mod file

```
replace github.com/username/package_name v0.0.0 => ../package_name (mention relative path from the current project's root)

require(
	github.com/username/package_name v0.0.0
)
```
*/
