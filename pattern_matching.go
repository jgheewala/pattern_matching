package pattern_matching

import (
	"strings"
)

var empty struct{}

/*
*
* '.' Matches any single character. '*' Matches zero or more of the preceding element.
* The matching should cover the entire input string (not partial).
* Some examples:
* isMatch("aa","a") → false
* isMatch("aa","aa") → true
* isMatch("aaa","aa") → false
* isMatch("aa", "a*") → true
* isMatch("aa", ".*") → true
* isMatch("ab", ".*") → true
* isMatch("aab", "c*a*b") → true
 */

type StringInfo struct {
	DontCare bool
	Count    int
}

func IsRegularExpressionMatch(s, p string) bool {
	sl := len(s) + 1
	pl := len(p) + 1
	rv := make([][]bool, sl)
	for idx := 0; idx < sl; idx++ {
		rv[idx] = make([]bool, pl)
		//fmt.Println(idx, len(rv[idx]))
	}

	rv[0][0] = true
	//fmt.Println(sl, pl, sl+pl, len(rv[0]))

	for idx := 1; idx < len(rv[0]); idx++ {
		if string(p[idx-1]) == "*" {
			rv[0][idx] = rv[0][idx-2]
		}
	}

	for i := 1; i < sl; i++ {
		for j := 1; j < len(rv[0]); j++ {
			//fmt.Printf("%d %d", i, j)
			if string(p[j-1]) == "." || p[j-1] == s[i-1] {
				rv[i][j] = rv[i-1][j-1]
			} else if string(p[j-1]) == "*" {
				rv[i][j] = rv[i][j-2]
				if string(p[j-2]) == "." || p[j-2] == s[i-1] {
					rv[i][j] = rv[i][j] || rv[i-1][j]
				}
			} else {
				rv[i][j] = false
			}
		}
		//fmt.Println()
	}

	return rv[sl-1][pl-1]
}

/*
* "Given a pattern and a string input - find if the string follows the same pattern and return 0 or 1.
 */
func IsWordPatternII290LeetCode(pattern, input string) bool {
	chMap := make(map[byte]string)
	iMap := make(map[string]struct{})
	return isWordMatch(input, 0, pattern, 0, chMap, iMap)
}

func isWordMatch(str string, i int, pat string, j int, chMap map[byte]string, iMap map[string]struct{}) bool {
	//fmt.Println(i, j)
	if i == len(str) && j == len(pat) {
		return true
	}

	if i >= len(str) || j >= len(pat) {
		return false
	}

	ch := pat[j]
	s, ok := chMap[ch]

	if ok {
		if !strings.HasPrefix(str[i:], s) {
			return false
		}
		return isWordMatch(str, i+len(s), pat, j+1, chMap, iMap)
	} else {
		for k := i; k < len(str); k++ {
			m := str[i : k+1]
			if _, ok := iMap[m]; ok {
				continue
			}
			//fmt.Println(k, ch, s, m)
			chMap[ch] = m
			iMap[m] = empty
			//fmt.Println(chMap)
			//fmt.Println(iMap)
			if isWordMatch(str, k+1, pat, j+1, chMap, iMap) {
				return true
			}

			delete(chMap, ch)
			delete(iMap, m)
			//fmt.Println(chMap)
			//fmt.Println(iMap)
		}
	}

	return false
}

func IsWordPattern290LeetCode(pattern, input string) bool {
	if len(pattern) == 0 || len(input) == 0 {
		return false
	}

	str := strings.Split(input, " ")

	if len(pattern) != len(str) {
		//fmt.Println("Len of pattern != len of str", len(pattern), len(str))
		return false
	}

	chMap := make(map[byte]string)
	iMap := make(map[string]struct{})

	for idx := 0; idx < len(pattern); idx++ {
		ch := pattern[idx]
		s, ok := chMap[ch]
		//fmt.Println(ch, s)
		if ok {
			if s != str[idx] {
				return false
			}
		} else {
			_, ok := iMap[str[idx]]
			if ok {
				return false
			}
			chMap[ch] = str[idx]
			iMap[str[idx]] = empty
		}
		//fmt.Println(chMap, iMap)
	}
	return true
}
