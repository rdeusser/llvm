// Code generated by "stringer -linecomment -type Tail"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TailNone-0]
	_ = x[TailMustTail-1]
	_ = x[TailNoTail-2]
	_ = x[TailTail-3]
}

const _Tail_name = "nonemusttailnotailtail"

var _Tail_index = [...]uint8{0, 4, 12, 18, 22}

func (i Tail) String() string {
	if i >= Tail(len(_Tail_index)-1) {
		return "Tail(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Tail_name[_Tail_index[i]:_Tail_index[i+1]]
}
