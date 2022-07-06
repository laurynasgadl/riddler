# Riddler
Riddler is a Golang package that provides functions to generate random data.

## Usage
Use `go get` to install the latest version of Riddler:
`go get -u github.com/laurynasgadl/riddler@latest`

### Strings
```go
import (
	"github.com/laurynasgadl/riddler"
	"github.com/laurynasgadl/riddler/strings"
)

str1 := riddler.String()
str2 := riddler.String(strings.WithLength(10))
str3 := riddler.String(
    strings.WithLength(32),
    strings.WithCharset(strings.AlphaNumCharset),
)
```
String generation can be adjusted using options `WithLength` (how long should the generated string be) and `WithCharset` (what characters should be used for the generation).

## Changelog
All notable changes between versions are tracked [here](https://github.com/laurynasgadl/riddler/blob/master/CHANGELOG.md).
