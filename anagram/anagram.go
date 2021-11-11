package anagram

import (
	"fmt"
	"reflect"
	"regexp"
)

func IsAnagram(str1 string, str2 string) bool {
	reg, _ := regexp.Compile("[^a-zA-z0-9]+")
	newStr1 := reg.ReplaceAllString(str1, "")
	newStr2 := reg.ReplaceAllString(str2, "")
	var hash1, hash2 = make(map[string]int), make(map[string]int)
	for _, s := range newStr1 {
		if v, ok := hash1[string(s)]; !ok {
			hash1[string(s)] = 1
		} else {
			hash1[string(s)] += v
		}
	}
	for _, s := range newStr2 {
		if v, ok := hash2[string(s)]; !ok {
			hash2[string(s)] = 1
		} else {
			hash2[string(s)] += v
		}
	}

	fmt.Println("Hash1", hash1, "hash2:", hash2)
	return reflect.DeepEqual(hash1, hash2)
}

func IsAnagramsimple(s string, t string) bool {
	alphabets := make([]int, 26)

	for i := range s {
		alphabets[s[i]-'a']++
	}

	for i := range t {
		alphabets[t[i]-'a']--
	}

	for i := range alphabets {
		if alphabets[i] != 0 {
			return false
		}
	}

	return true

}
