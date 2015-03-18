package bit

import "math"

const NotFound = math.MaxUint64

// ScanRight scans from the highest to the lowest bit
// if x == 0 it will return NotFound
func ScanRight(x uint64) uint64 {
	if x == 0 {
		return NotFound
	}
	n := uint64(0)

	for x > scanTableMask {
		x >>= scanTableBits
		n += scanTableBits
	}
	return scanRightTable[x] + n
}

func slowScanRight(x uint64) uint64 {
	for k := 63; k >= 0; k -= 1 {
		if x&(1<<byte(k)) != 0 {
			return uint64(k)
		}
	}
	return NotFound
}

// ScanLeft scans from the lowest to the highest bit
// if x == 0 it will return NotFound
func ScanLeft(x uint64) uint64 {
	if x == 0 {
		return NotFound
	}

	n := uint64(0)
	for x&scanTableMask == 0 {
		n += scanTableBits
		x >>= scanTableBits
	}
	return scanLeftTable[x&scanTableMask] + n
}

func slowScanLeft(x uint64) uint64 {
	for k := byte(0); k < 64; k += 1 {
		if x&(1<<k) != 0 {
			return uint64(k)
		}
	}
	return NotFound
}
