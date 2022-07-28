package longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func Test_longestSubstring(t *testing.T) {
	fmt.Println("abcabcbb", longestSubstring("abceabcbb"))
	fmt.Println("bbbbb", longestSubstring("bbbbb"))
	fmt.Println("pwwkew", longestSubstring("pwwkewwabc"))
}
