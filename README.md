# go-censorword

go-censorword is a lightweight library for detecting profanities in Go string.


## Installation
```
    go get -u github.com/pcpratheesh/go-censorword
```
## Usage
```go
import (
	"github.com/pcpratheesh/go-censorword"
)
```

## In working
go-censorword package uses a censorWord [here](censor/censor.go) list to check the profanities. 
Als we have provided an option to override this list of contents by using
```go
CustomCensorList() 
```
You can provide your own list to search and replace the profanities

## Example
```go
var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")
detector.CustomCensorList([]string{
    "bad", "word","one",
})
```
## Example
```go
detector := NewDetector().SetCensorReplaceChar("*")

resultString, err := detector.CensorWord(inputString)

if err != nil {
    panic(err)
}
```


In the future, we should implement the following points
- Support for other language profanities
- All words having repeated characters more than twice (eg : fuck -> fuuuuuck)
- Should remove word match conditions (eg :-> fucker, fucking..etc)