// Code generated by "stringer -linecomment -type ReturnAttr"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ReturnAttrInReg-0]
	_ = x[ReturnAttrNoAlias-1]
	_ = x[ReturnAttrNonNull-2]
	_ = x[ReturnAttrSignExt-3]
	_ = x[ReturnAttrZeroExt-4]
}

const _ReturnAttr_name = "inregnoaliasnonnullsignextzeroext"

var _ReturnAttr_index = [...]uint8{0, 5, 12, 19, 26, 33}

func (i ReturnAttr) String() string {
	if i >= ReturnAttr(len(_ReturnAttr_index)-1) {
		return "ReturnAttr(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ReturnAttr_name[_ReturnAttr_index[i]:_ReturnAttr_index[i+1]]
}
