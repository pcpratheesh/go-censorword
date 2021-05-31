package gocensorword_test

import (
	"fmt"
	"testing"

	gocensorword "github.com/pcpratheesh/go-censorword"
	"github.com/stretchr/testify/require"
)

func TestBadWord(t *testing.T) {
	var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")

	word := "bitch"
	resultString, err := detector.CensorWord(word)
	if err != nil {
		panic(err)
	}

	require.Equal(t, resultString, "*****")
}

func TestWithCustomList(t *testing.T) {
	var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")

	word := "bad ass"
	detector.CustomCensorList([]string{
		"ass", "bitch",
	})
	resultString, err := detector.CensorWord(word)
	if err != nil {
		panic(err)
	}

	require.Equal(t, resultString, "bad ***")
}

func TestBadWordFirstLetterKept(t *testing.T) {
	var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")

	word := "bitch"
	detector.KeepPrefixChar = true
	resultString, err := detector.CensorWord(word)
	if err != nil {
		panic(err)
	}

	require.Equal(t, resultString, "b****")
}

func TestBadWordFirstAndLastLetterKept(t *testing.T) {
	var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")

	word := "bitch"
	detector.KeepPrefixChar = true
	detector.KeepSuffixChar = true
	resultString, err := detector.CensorWord(word)
	if err != nil {
		panic(err)
	}

	require.Equal(t, resultString, "b***h")
}

func TestBadWordEmptyList(t *testing.T) {
	var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")
	detector.CustomCensorList([]string{})
	word := "bitch"
	_, err := detector.CensorWord(word)
	require.NotNil(t, err)
}

func TestBadFullLength(t *testing.T) {
	var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")
	detector.KeepPrefixChar = true
	detector.KeepSuffixChar = true
	word := "fuck post content asshole suck sucker"
	resultString, _ := detector.CensorWord(word)
	require.Equal(t, resultString, "f**k post content a*****e s**k s****r")
}

func TestBadWithCustomReplacePattern(t *testing.T) {
	var detector = gocensorword.NewDetector().SetCensorReplaceChar("*")
	detector.KeepPrefixChar = true
	detector.KeepSuffixChar = true
	detector.ReplaceCheckPattern = `\b%s\b`
	word := "pass ass fucker sucker"
	resultString, _ := detector.CensorWord(word)
	fmt.Println("resulr----", resultString)
	require.Equal(t, resultString, "pass a*s f****r s****r")
}
