package isogram

import "strings"

func IsIsogram(word string) bool {
	var result bool
	wordByte := []byte(strings.ToLower(word))
	letterAmount := make(map[byte]int)
	for i := range wordByte {
		letterAmount[wordByte[i]] += 1
		if letterAmount[wordByte[i]] > 1 {
			if wordByte[i] == byte(' ') || wordByte[i] == byte('-') {
				continue
			}
			result = false
			return result
		}
	}
	return true

}
