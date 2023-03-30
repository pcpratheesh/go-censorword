[![Create Release](https://github.com/pcpratheesh/go-censorword/actions/workflows/release.yml/badge.svg)](https://github.com/pcpratheesh/go-censorword/actions/workflows/release.yml)

# go-censorword
go-censorword is a lightweight and easy-to-use tool that allows you to detect and filter out profanity words from your text-based content. Whether you're building a social media platform, a chat app, or just want to keep your comments section clean, this package can help.

## Installation
```sh
    go get -u github.com/pcpratheesh/go-censorword
```

for previous version

```sh
    go get -u github.com/pcpratheesh/go-censorword@v1.1.0
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
var detector = gocensorword.NewDetector(
	gocensorword.WithCensorReplaceChar("*"),
    
    // override the existing list of bad words with your own
    gocensorword.WithCustomCensorList([]string{
        "bad", "word","one",
    }),
)

// censor the word
actualString := "with having some bad words"
filterString, err := detector.CensorWord(actualString)
if err != nil {
    panic(err)
}

```

## Option Methods
- *WithCensorReplaceChar(string)* : This method can be used to replace the filtered word characters with asterisks (*), dashes (-) or custom characters, like the pound sign (#) or at sign (@).
- *WithCustomCensorList([]string)* : The list of your own profanity words
- *WithSanitizeSpecialCharacters(bool)*: To sanitize the special characters in the word
- *WithKeepPrefixChar(bool)*: If you want to Kept the prefix Character (eg : F****)
- *WithKeepSuffixChar(bool)*: If you want to Kept the suffix Character (eg : ****K)

```go
detector := NewDetector(
    gocensorword.WithCensorReplaceChar("*"),
)

resultString, err := detector.CensorWord(inputString)

if err != nil {
    panic(err)
}
```


In the future, we should implement the following points
- Support for other language profanities
- All words having repeated characters more than twice 


## Contributing
Contributions to the Profanity Filter package are welcome and encouraged! If you find a bug or have a feature request, please open an issue on GitHub. If you'd like to contribute code, please fork the repository, make your changes, and submit a pull request.


## License
The Profanity Filter package is licensed under the MIT License. See the LICENSE file for more information.
