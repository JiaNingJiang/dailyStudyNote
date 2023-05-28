package hashTable

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	file, err := os.Open("word-list-large.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	var wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	var elemNum = uint64(len(wordlist))
	var errRate = 0.0001

	bloomFilter := NewBloomFilter(elemNum, errRate)

	fmt.Println("SampleSize: ", bloomFilter.SampleSize)
	fmt.Println("BloomSize: ", bloomFilter.BloomSize)
	fmt.Println("HashFuncNum: ", bloomFilter.HashFuncNum)
	fmt.Println("ErrRate: ", bloomFilter.ErrRate)

	for _, word := range wordlist {
		bloomFilter.Add([]byte(word))
	}

	var testcases = []struct {
		Elem string
		Want bool
	}{
		{"hello", true},
		{"zoo", false},
		{"word", true},
		{"alibaba", false},
	}

	for _, oneCase := range testcases {
		got := bloomFilter.IsContain([]byte(oneCase.Elem))
		if got != oneCase.Want {
			if got {
				t.Error("should not contain", oneCase.Elem)
			} else {
				t.Error("should contain", oneCase.Elem)
			}
			t.Error("got != want")
			return
		}
	}
}
