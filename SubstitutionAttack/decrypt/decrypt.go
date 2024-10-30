package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type letter struct {
	Key   string
	Value float32
}
type substitution struct {
	letter string
	sub    string
}
type bigrams struct {
	bigram string
	freq   float32
}
type trigrams struct {
	trigram string
	freq    float32
}
type values struct {
	letter             string
	substitutionValues map[string]float32
}

var sortedFrequencies []letter
var sortedKey []substitution
var sortedBigrams []bigrams
var sortedTrigrams []trigrams
var possibleSubs []values

func readCryptotext() string {
	cryptotext, error := os.ReadFile("cryptotext.txt")
	if error != nil {
		panic(error)
	}

	return string(cryptotext)
}

func findFrequencies(alphabet []string, cryptotext string) {

	frequencies := make(map[string]float32)
	length := len(cryptotext)
	for _, letter := range alphabet {
		frequency := strings.Count(cryptotext, string(letter))

		var freqFloat float32 = float32(frequency) / float32(length)

		frequencies[string(letter)] = freqFloat
	}

	for key, value := range frequencies {
		sortedFrequencies = append(sortedFrequencies, letter{key, value})
	}

	sort.Slice(sortedFrequencies, func(i, j int) bool {
		return sortedFrequencies[i].Value > sortedFrequencies[j].Value
	})


}
func addPossibleValues1() {
	//adauga prima valoare posibila pentru fiecare litera dupa calculul frecventei
	var topFreq = []string{"E", "T", "A", "O", "I", "N", "S", "H", "R", "D", "L", "U", "C", "M", "F", "W", "G", "Y", "P", "B", "V", "K", "X", "J", "Q", "Z"}

	for i := 0; i < 4; i++ {
		letter := sortedFrequencies[i].Key
		freq := sortedFrequencies[i].Value
		possibleSubs = append(possibleSubs, values{string(letter), make(map[string]float32)})
		possibleSubs[len(possibleSubs)-1].substitutionValues[topFreq[i]] = float32(freq)
	}
	letter := sortedFrequencies[len(sortedFrequencies)-1].Key
	freq := sortedFrequencies[len(sortedFrequencies)-1].Value
	possibleSubs = append(possibleSubs, values{string(letter), make(map[string]float32)})
	possibleSubs[len(possibleSubs)-1].substitutionValues[topFreq[len(topFreq)-1]] = float32(freq)

}
func addPossibleValues2() {

	//adauga a doua valoare posibila pentru fiecare litera dupa calcul frecventelor digramelor
	var topBigrams = []string{"TH", "HE", "IN", "ER", "AN", "ER", "ND", "ON", "EN", "AT", "OU", "ED", "HA", "TO", "OR", "IT", "IS", "HI", "ES", "NG"}
	for i, bigram := range sortedBigrams {
		if i < len(topBigrams) {
			for indx, x := range possibleSubs {
				if x.letter == string(bigram.bigram[0]) {
					currentFreq := possibleSubs[indx].substitutionValues[string(topBigrams[i][0])]
					if currentFreq < float32(bigram.freq) {
						possibleSubs[indx].substitutionValues[string(topBigrams[i][0])] = float32(bigram.freq)
					}
				}
				if x.letter == string(bigram.bigram[1]) {
					currentFreq := possibleSubs[indx].substitutionValues[string(topBigrams[i][1])]
					if currentFreq < float32(bigram.freq) {
						possibleSubs[indx].substitutionValues[string(topBigrams[i][1])] = float32(bigram.freq)
					}
				}
			}
		} else {
			break
		}
	}

}
func addPossibleValues3() {
	//adauga a treia valoare posibila pentru fiecare litera dupa calcul frecventelor trigramelor
	var topTrigrams = []string{"THE", "AND", "ING", "HER", "HAT", "HIS", "THA", "ERE", "FOR", "ENT", "ION", "TER", "WAS", "YOU", "ITH", "VER", "ALL", "WIT", "THI", "TIO"}
	for i := 0; i < 5; i++ {
		trigram := sortedTrigrams[i]
		for j, letter := range trigram.trigram {
			exists := false
			for _, element := range possibleSubs {
				if element.letter == string(letter) {
					exists = true
				}
			}
			if !exists {
				possibleSubs = append(possibleSubs, values{string(letter), make(map[string]float32)})

				possibleSubs[len(possibleSubs)-1].substitutionValues[string(topTrigrams[i][j])] = float32(trigram.freq)
			}
		}
	}

}

func containsBigram(stringsArray []bigrams, element string) bool {
	for _, x := range stringsArray {
		if x.bigram == element {
			return true
		}
	}
	return false
}
func containsTrigram(stringsArray []trigrams, element string) bool {
	for _, x := range stringsArray {
		if x.trigram == element {
			return true
		}
	}
	return false
}

func bigramsFreq(cryptotext string) {

	for i := 0; i < len(cryptotext)-2; i++ {
		frequency := strings.Count(cryptotext, cryptotext[i:i+2])
		var freqFloat float32 = float32(frequency) / float32(len(cryptotext)-1)

		if !containsBigram(sortedBigrams, cryptotext[i:i+2]) {
			sortedBigrams = append(sortedBigrams, bigrams{cryptotext[i : i+2], freqFloat})
		}
	}
	sort.Slice(sortedBigrams, func(i, j int) bool {
		return sortedBigrams[i].freq > sortedBigrams[j].freq
	})


}
func trigramsFreq(cryptotext string) {

	for i := 0; i < len(cryptotext)-3; i++ {
		frequency := strings.Count(cryptotext, cryptotext[i:i+3])
		var freqFloat float32 = float32(frequency) / float32(len(cryptotext)-2)
		if !containsTrigram(sortedTrigrams, cryptotext[i:i+3]) {
			sortedTrigrams = append(sortedTrigrams, trigrams{cryptotext[i : i+3], freqFloat})
		}

	}
	sort.Slice(sortedTrigrams, func(i, j int) bool {
		return sortedTrigrams[i].freq > sortedTrigrams[j].freq
	})
	
}
func findKey() {
	key := make(map[string]string)

	for _, element := range possibleSubs {
		var maximumP float32 = float32(math.Inf(-1))
		var maxLetter string = ""
		// fmt.Printf("%s\n", element.letter)
		for associatedLetter, probability := range element.substitutionValues {
			if probability > maximumP {
				maximumP = probability
				maxLetter = associatedLetter
			}
		}
		key[string(element.letter)] = maxLetter

	}

	for key, value := range key {
		sortedKey = append(sortedKey, substitution{key, value})
	}

	sort.Slice(sortedKey, func(i, j int) bool {
		return sortedKey[i].letter < sortedKey[j].letter
	})

	for _, substitution := range sortedKey {
		fmt.Printf("%s: %s\n", substitution.letter, substitution.sub)
	}

}

func main() {

	var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	cryptotext := readCryptotext()
	findFrequencies(alphabet, cryptotext)
	addPossibleValues1()

	trigramsFreq(cryptotext)
	addPossibleValues3()

	bigramsFreq(cryptotext)
	addPossibleValues2()

	findKey()

}
