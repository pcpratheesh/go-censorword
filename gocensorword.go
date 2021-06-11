package gocensorword

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/pcpratheesh/go-censorword/censor"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	Transformer transform.Transformer
)

type CensorWordDetection struct {
	CensorList                []string
	CensorReplaceChar         string
	KeepPrefixChar            bool
	KeepSuffixChar            bool
	SanitizeSpecialCharacters bool
	TextNormalization         bool
}

// this will create a new CensorWordDetection object
func NewDetector() *CensorWordDetection {
	return &CensorWordDetection{
		CensorList:                censor.CensorWordsList,
		CensorReplaceChar:         censor.CensorChar,
		KeepPrefixChar:            false,
		KeepSuffixChar:            false,
		SanitizeSpecialCharacters: true,
		TextNormalization:         true,
	}
}

// change the default censor list
// can provide own censor words list
func (censor *CensorWordDetection) CustomCensorList(list []string) *CensorWordDetection {
	censor.CensorList = list
	return censor
}

// change the censorReplaceCharacter
func (censor *CensorWordDetection) SetCensorReplaceChar(char string) *CensorWordDetection {
	censor.CensorReplaceChar = char
	return censor
}

// sanitize special characters
func (censor *CensorWordDetection) WithSanitizeSpecialCharacters(status bool) *CensorWordDetection {
	censor.SanitizeSpecialCharacters = status
	return censor
}

// sanitize text normalization
func (censor *CensorWordDetection) WithTextNormalization(status bool) *CensorWordDetection {
	censor.TextNormalization = status
	return censor
}

// sanitize text normalization
func (censor *CensorWordDetection) normalizeText(word string) string {
	if Transformer == nil {
		Transformer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	}
	word, _, _ = transform.String(Transformer, word)
	return word
}

// remove special characters from string
func (censor *CensorWordDetection) SanitizeCharacter(str string) string {
	str = strings.ToLower(str)
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	str = re.ReplaceAllString(str, " ")

	return str
}

// Censor Word
func (censor *CensorWordDetection) CensorWord(word string) (string, error) {

	// sanitize with text normalization
	word = censor.normalizeText(word)

	if censor.SanitizeSpecialCharacters {
		word = censor.SanitizeCharacter(word)
	}

	//sort on descending
	sort.Strings(censor.CensorList)
	sort.Slice(censor.CensorList, func(i, j int) bool {
		return len(censor.CensorList[i]) > len(censor.CensorList[j])
	})

	// check the list is empty
	if ok := len(censor.CensorList) > 0; !ok {
		return "", fmt.Errorf("found empty censor word list")
	}
	// convert str into a slice
	for _, forbiddenWord := range censor.CensorList {

		// should replace incase sensitive
		pattern := regexp.MustCompile(fmt.Sprintf(`(?i)%s`, forbiddenWord))
		var replacePattern, prefix, suffix string
		wordLength := len(forbiddenWord)

		if censor.KeepPrefixChar {
			prefix = string(forbiddenWord[0])
			wordLength--
		}
		if censor.KeepSuffixChar {
			suffix = string(forbiddenWord[len(forbiddenWord)-1])
			wordLength--
		}

		replacePattern = fmt.Sprintf(
			"%s%s%s", prefix, strings.Repeat(censor.CensorReplaceChar, wordLength), suffix,
		)
		word = pattern.ReplaceAllString(word, replacePattern)
	}
	// join the string
	return word, nil
}
