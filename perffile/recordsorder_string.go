// Code generated by "stringer -type=RecordsOrder"; DO NOT EDIT.

package perffile

import "strconv"

const _RecordsOrder_name = "RecordsFileOrderRecordsCausalOrderRecordsTimeOrder"

var _RecordsOrder_index = [...]uint8{0, 16, 34, 50}

func (i RecordsOrder) String() string {
	if i < 0 || i >= RecordsOrder(len(_RecordsOrder_index)-1) {
		return "RecordsOrder(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RecordsOrder_name[_RecordsOrder_index[i]:_RecordsOrder_index[i+1]]
}
