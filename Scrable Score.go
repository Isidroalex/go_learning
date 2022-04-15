package scrabble

import "strings"

//Score counting score of the word for Scramble's rules
func Score(word string) int {

	alphabet := make(map[byte]int)
	word = strings.ToLower(word)
	wordByte := []byte(word)
	score := 0
	alphabet['a'], alphabet['e'], alphabet['i'], alphabet['o'], alphabet['u'], alphabet['l'], alphabet['n'], alphabet['r'], alphabet['s'], alphabet['t'] = 1, 1, 1, 1, 1, 1, 1, 1, 1, 1
	alphabet['d'], alphabet['g'] = 2, 2
	alphabet['b'], alphabet['c'], alphabet['m'], alphabet['p'] = 3, 3, 3, 3
	alphabet['f'], alphabet['h'], alphabet['v'], alphabet['w'], alphabet['y'] = 4, 4, 4, 4, 4
	alphabet['k'] = 5
	alphabet['j'], alphabet['x'] = 8, 8
	alphabet['q'], alphabet['z'] = 10, 10

	for i := 0; i < len(word); i++ {
		score = score + alphabet[wordByte[i]]
	}
	return score

}
