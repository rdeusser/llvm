// Code generated by "string2enum -linecomment -type FPred /home/u/Desktop/go/src/github.com/llir/llvm/ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

const _FPred_name = "falseoeqogeogtoleoltoneordtrueuequgeugtuleultuneuno"

var _FPred_index = [...]uint8{0, 5, 8, 11, 14, 17, 20, 23, 26, 30, 33, 36, 39, 42, 45, 48, 51}

func FPredFromString(s string) enum.FPred {
	if len(s) == 0 {
		return 0
	}
	for i := range _FPred_index[:len(_FPred_index)-1] {
		if s == _FPred_name[_FPred_index[i]:_FPred_index[i+1]] {
			return enum.FPred(i)
		}
	}
	panic(fmt.Errorf("unable to locate FPred enum corresponding to %q", s))
}