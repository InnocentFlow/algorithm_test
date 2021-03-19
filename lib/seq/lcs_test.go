package seq

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubseq(t *testing.T) {
	src := "abcde"
	dst := "agdedbcd"
	//src := "acdgde"
	//dst := "agdedbcd"
	cs := LongestCommonSubseqOutput(src, dst)
	fmt.Printf("longest common sequence is %s", cs)
}
