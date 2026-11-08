package textutil

import "strings"

func Reverse(s string) (string) {
	convertedByteSlice := make([]byte, len(s))

	for i := range(s) {
		convertedByteSlice[i] = s[i]
	}

	l := 0
	r := len(convertedByteSlice) - 1
	
	for l <= r {
		t := convertedByteSlice[l]
		convertedByteSlice[l] = convertedByteSlice[r]
		convertedByteSlice[r] = t

		l += 1;
		r -= 1
	}

	reversedString := string(convertedByteSlice)
	return reversedString
}

func Shout(s string) (string) {
	return addBang(strings.ToUpper(s))
}

func addBang(s string) (string) {
	return s + "!"
}