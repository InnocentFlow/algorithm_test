package seq

import (
	"fmt"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LongestCommonSubseqOutput(text1 string, text2 string) string {
	r1 := []rune(text1)
	r2 := []rune(text2)
	if len(r1) > len(r2) {
		r1, r2 = r2, r1
	}

	h := len(r1)
	w := len(r2)
	dp := make([][]int, h+1)
	for idx := range dp {
		dp[idx] = make([]int, w+1)
	}

	for x, char1 := range r1 {
		for y, char2 := range r2 {
			if char2 == char1 {
				dp[x+1][y+1] = dp[x][y] + 1
			} else {
				dp[x+1][y+1] = max(dp[x][y+1], dp[x+1][y])
			}
		}
	}

	resChar := make([]string, dp[h][w])
	currentL := dp[len(r1)][len(r2)]

	for i := h; i > 0 && currentL > 0; {
		for j := w; j > 0 && currentL > 0; {
			if dp[i-1][j] == currentL {
				i--
			} else if dp[i][j-1] == currentL {
				j--
			} else {
				resChar[currentL-1] = string(r1[i-1])
				currentL--
				i--
				j--
			}
		}
	}
	return strings.Join(resChar, "")
}

func LongestCommonSubseq(text1 string, text2 string) int {
	r1 := []rune(text1)
	r2 := []rune(text2)
	if len(r1) > len(r2) {
		r1, r2 = r2, r1
	}
	fmt.Printf("%s\n%s\n", string(r1), string(r2))
	dp := make([][]int, len(r1)+1)
	for idx := range dp {
		dp[idx] = make([]int, len(r2)+1)
	}

	for x, char1 := range r1 {
		for y, char2 := range r2 {
			if char2 == char1 {
				dp[x+1][y+1] = dp[x][y] + 1
			} else {
				dp[x+1][y+1] = max(dp[x][y+1], dp[x+1][y])
			}
		}
		fmt.Printf("\n%v\n", dp[x+1])
	}
	return dp[len(r1)][len(r2)]
}
