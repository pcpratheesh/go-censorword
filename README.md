[![Create Release](https://github.com/pcpratheesh/go-censorword/actions/workflows/release.yml/badge.svg)](https://github.com/pcpratheesh/go-censorword/actions/workflows/release.yml)

# go-censorword
go-censorword is a lightweight and easy-to-use tool that allows you to detect and filter out profanity words from your text-based content. Whether you're building a social media platform, a chat app, or just want to keep your comments section clean, this package can help.

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
The go-censorword package uses a censorWord [here](censor/censor.go) list to check profanities, but also provides an option for you to override this list with your own contents. You can create a list of bad words that are not included in the original blacklist by using the **customCensorList** method.

```go
CustomCensorList([]string{}) 
```

## How to use
```go
// this would initialize the detector object.
var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")

// override the existing list of bad words with your own
detector.CustomCensorList([]string{
    "bad", "word","one",
})

// censor the word
actualString := "with having some bad words"
filterString, err := detector.CensorWord(actualString)
if err != nil {
    panic(err)
}

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
- All words having repeated characters more than twice 
