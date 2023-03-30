package gocensorword

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/pcpratheesh/go-censorword/censor"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	Transformer transform.Transformer
)

type Options func(*CensorWordDetection)
type CensorWordDetection struct {
	CensorList                []string
	CensorReplaceChar         string
	KeepPrefixChar            bool
	KeepSuffixChar            bool
	SanitizeSpecialCharacters bool
	TextNormalization         bool
	ReplaceCheckPattern       string
}

// this will create a new CensorWordDetection object
func NewDetector(options ...Options) *CensorWordDetection {
	detector := &CensorWordDetection{
		CensorList:                censor.CensorWordsList,
		CensorReplaceChar:         censor.CensorChar,
		KeepPrefixChar:            false,
		KeepSuffixChar:            false,
		SanitizeSpecialCharacters: true,
		TextNormalization:         true,
		ReplaceCheckPattern:       "(?i)%s",
	}

	// add / update new options
	for _, opt := range options {
		opt(detector)
	}

	return detector
}

// WithCustomCensorList
// change the default censor list
// can provide own censor words list
func WithCustomCensorList(list []string) Options {
	return func(detector *CensorWordDetection) {
		detector.CensorList = list
	}
}

// WithCensorReplaceChar
func WithCensorReplaceChar(char string) Options {
	return func(detector *CensorWordDetection) {
		detector.CensorReplaceChar = char
	}
}

// WithSanitizeSpecialCharacters
func WithSanitizeSpecialCharacters(status bool) Options {
	return func(detector *CensorWordDetection) {
		detector.SanitizeSpecialCharacters = status
	}
}

// WithTextNormalization
func WithTextNormalization(status bool) Options {
	return func(detector *CensorWordDetection) {
		detector.TextNormalization = status
	}
}

// WithKeepPrefixChar
func WithKeepPrefixChar() Options {
	return func(detector *CensorWordDetection) {
		detector.KeepPrefixChar = true
	}
}

// WithKeepPrefixChar
func WithKeepSuffixChar() Options {
	return func(detector *CensorWordDetection) {
		detector.KeepSuffixChar = true
	}
}

// WithReplaceCheckPattern
func WithReplaceCheckPattern(pattern string) Options {
	return func(detector *CensorWordDetection) {
		detector.ReplaceCheckPattern = pattern
	}
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
	if censor.TextNormalization {
		word = censor.normalizeText(word)
	}

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
	for _, fword := range censor.CensorList {

		forbiddenWord := fword
		forbiddenWord = strings.ToValidUTF8(forbiddenWord, "")
		if !utf8.ValidString(forbiddenWord) {
			continue
		}

		// should replace incase sensitive
		patterFormat := fmt.Sprintf(censor.ReplaceCheckPattern, forbiddenWord)
		pattern := regexp.MustCompile(patterFormat)
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
