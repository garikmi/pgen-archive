package main

import (
	"github.com/garikmi/pgen/cmd/internal/helpers"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GenerateWord generates a password made up of words
func (app *application) generateWord() string {
	const sConsonants string = "BCDFGHJKLMNPRSTVWYZBlBrChClCrDrFlFrGhGlGnGrKnPhPlPrQuScShSkSlSmSnSpStThTrWhWrSchScrShmShrSquStrThr"
	const vowels string = "AeAiAoAuEaEeEiEuIaIeIoOaOeOiOoOuUeUiAEIOU"
	const eConsonants string = "BCDFGLMNPRSTXZBtChCkCtFtGhGnLbLdLfLkLlLmLnLpLtMbMnMpNkNgNtPhPtRbRcRdRfRgRkRlRmRnRpRtRvRzShSkSpSsStZzLchLshLthRchRshRstRthSchTch"

	var result string

	// build regex string
	re, err := regexp.Compile(`[A-Z][a-z]*`)
	if err != nil {
		log.Fatal(err)
	}

	// regex parse parts
	s := re.FindAllString(sConsonants, -1)
	v := re.FindAllString(vowels, -1)
	e := re.FindAllString(eConsonants, -1)

	// random generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// generate find password
	for i := 1; i <= 3; i++ {
		if helpers.IsTrue() {
			result += strings.ToUpper(generateVowel(&s, &v, &e, r1))
		} else {
			result += generateVowel(&s, &v, &e, r1)
		}
		if helpers.IsTrue() {
			result += strconv.Itoa(r1.Intn(9))
		}
		if i < 3 {
			result += "-"
		}
	}

	return result
}

// generateVowel generates a vowel with prefix and suffix
func generateVowel(s, v, e *[]string, rand *rand.Rand) string {
	var result string

	for i := 1; i <= rand.Intn(2)+1; i++ {
		if helpers.IsTrue() {
			result += (*s)[rand.Intn(len(*s))]
		}
		result += (*v)[rand.Intn(len(*v))]
		if helpers.IsTrue() {
			result += (*e)[rand.Intn(len(*e))]
		}
	}

	return strings.ToLower(result)
}

// GenerateRandom generates a password made up of random characters
func (app *application) generateRandom() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var result string

	// random generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 1; i <= 12; i++ {
		result += string(letters[r1.Intn(len(letters))])
	}

	return result
}
