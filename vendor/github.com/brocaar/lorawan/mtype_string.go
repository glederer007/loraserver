// Code generated by "stringer -type=MType"; DO NOT EDIT.

package lorawan

import "fmt"

const _MType_name = "JoinRequestJoinAcceptUnconfirmedDataUpUnconfirmedDataDownConfirmedDataUpConfirmedDataDownRFUProprietary"

var _MType_index = [...]uint8{0, 11, 21, 38, 57, 72, 89, 92, 103}

func (i MType) String() string {
	if i >= MType(len(_MType_index)-1) {
		return fmt.Sprintf("MType(%d)", i)
	}
	return _MType_name[_MType_index[i]:_MType_index[i+1]]
}
