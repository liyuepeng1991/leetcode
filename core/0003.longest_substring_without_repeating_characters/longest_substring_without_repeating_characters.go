package longest_substring_without_repeating_characters

// 题目描述：
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 。

// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 1
// 2
// 3

// 示例 2:
// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 1
// 2
// 3

// 示例 3:
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func longestSubstring(s string) string {
	//有256个字符
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}
	left0 := 0
	right0 := 0
	left := 0
	right := 0
	maxlen := 0
	for i := 0; i < len(s); i++ {
		right = i + 1
		if location[s[i]] >= left { //出现了相同字符
			left = location[s[i]] + 1
		} else if i+1-left > maxlen {
			maxlen = i + 1 - left
			left0 = left
			right0 = right
		}
		location[s[i]] = i
	}
	return s[left0:right0]
}
