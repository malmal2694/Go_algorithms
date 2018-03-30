package longestSubsequence

import "fmt"

func ComputeLcsTable(x, y string) [][]int {
	var lenLongSubseq [][]int // length_longest_subsequence
	var longestDigit int

	lenLongSubseq = make([][]int, len(x)+1)
	for i := 1; i <= len(x); i++ {
		lenLongSubseq[i] = append(lenLongSubseq[i], 0) //Append 0 to lenLongSubseq[i, 0]
	}
	for j := 0; j <= len(y); j++ {
		lenLongSubseq[0] = append(lenLongSubseq[0], 0) //Append 0 to lenLongSubseq[0, j]
	}

	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				lenLongSubseq[i] = append(lenLongSubseq[i], (lenLongSubseq[i-1][j-1])+1) //Set lenLongSubseq[i,j] to [i - 1, j - 1] + 1
			} else {
				longestDigit = lenLongSubseq[i][j-1]      //Set lengestDigit to lenLongSubseq[i, j - 1]
				if lenLongSubseq[i-1][j] > longestDigit { // If lenLongSubseq[i - 1, j] bigger than lenLongSubseq[i, j - 1] than set longestDigit to lenLongSubseq[i - 1, j]
					longestDigit = lenLongSubseq[i-1][j]
				}
				lenLongSubseq[i] = append(lenLongSubseq[i], longestDigit)

			}
		}
	}
	return lenLongSubseq
}
func AssembleLcs(x, y string, l [][]int, i, j int) string {
	ii := fmt.Sprintf("0%s", x) //Create ii variable and add "0" character and x variable
	jj := fmt.Sprintf("0%s", y) //Create jj variable and add "0" character and y variable

	if l[i][j] == 0 {
		return "" //If l[i, j] equals 0, then return the empty string
	} else if ii[i] == jj[j] {
		//a := fmt.Sprintf("%s%c", AssembleLcs(x, y, l, i-1, j-1), ii[i]) // x[i])
		fmt.Printf("hello %s   %c\n", a, ii[i])
		return a
	} else if l[i][j-1] > l[i-1][j] {
		return AssembleLcs(x, y, l, i, j-1)
	} else {
		return AssembleLcs(x, y, l, i-1, j)
	}
}
