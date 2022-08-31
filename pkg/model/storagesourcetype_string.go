// Code generated by "stringer -type=StorageSourceType --trimprefix=StorageSource"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[storageSourceUnknown-0]
	_ = x[StorageSourceIPFS-1]
	_ = x[StorageSourceURLDownload-2]
	_ = x[StorageSourceFilecoinUnsealed-3]
	_ = x[StorageSourceFilecoin-4]
	_ = x[storageSourceDone-5]
}

const _StorageSourceType_name = "storageSourceUnknownIPFSURLDownloadFilecoinUnsealedFilecoinstorageSourceDone"

var _StorageSourceType_index = [...]uint8{0, 20, 24, 35, 51, 59, 76}

func (i StorageSourceType) String() string {
	if i < 0 || i >= StorageSourceType(len(_StorageSourceType_index)-1) {
		return "StorageSourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StorageSourceType_name[_StorageSourceType_index[i]:_StorageSourceType_index[i+1]]
}
