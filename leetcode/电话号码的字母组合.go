package main

func letterCombinations(digits string) []string {
	a := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	for i := 0; i < len(digits); i++ {
		a[digits[i]]
	}
}
