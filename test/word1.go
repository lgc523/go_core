package test

import "unicode"

// IsPalindrome1  判断一个字符串是否是会问字符串
// 忽略大小写，一级非字母字符
func IsPalindrome1(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
