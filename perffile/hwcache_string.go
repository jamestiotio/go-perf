// Code generated by "stringer -type=HWCache"; DO NOT EDIT.

package perffile

import "strconv"

const _HWCache_name = "HWCacheL1DHWCacheL1IHWCacheLLHWCacheDTLBHWCacheITLBHWCacheBPUHWCacheNode"

var _HWCache_index = [...]uint8{0, 10, 20, 29, 40, 51, 61, 72}

func (i HWCache) String() string {
	if i >= HWCache(len(_HWCache_index)-1) {
		return "HWCache(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _HWCache_name[_HWCache_index[i]:_HWCache_index[i+1]]
}
