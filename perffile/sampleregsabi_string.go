// Code generated by "stringer -type=SampleRegsABI"; DO NOT EDIT.

package perffile

import "strconv"

const _SampleRegsABI_name = "SampleRegsABINoneSampleRegsABI32SampleRegsABI64"

var _SampleRegsABI_index = [...]uint8{0, 17, 32, 47}

func (i SampleRegsABI) String() string {
	if i >= SampleRegsABI(len(_SampleRegsABI_index)-1) {
		return "SampleRegsABI(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SampleRegsABI_name[_SampleRegsABI_index[i]:_SampleRegsABI_index[i+1]]
}
