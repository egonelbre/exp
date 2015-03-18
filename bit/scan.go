package bit

const NotFound = 0xFF

// ScanRight scans from the highest to the lowest bit
// if x == 0 it will return NotFound
func ScanRight(x uint64) uint {
	if x == 0 {
		return NotFound
	}
	n := uint(0)

	for x > scanTableMask {
		x >>= scanTableBits
		n += scanTableBits
	}
	return uint(scanRightTable[x]) + n
}

func slowScanRight(x uint64) uint {
	for k := 63; k >= 0; k -= 1 {
		if x&(1<<byte(k)) != 0 {
			return uint(k)
		}
	}
	return NotFound
}

// ScanLeft scans from the lowest to the highest bit
// if x == 0 it will return NotFound
func ScanLeft(x uint64) uint {
	if x == 0 {
		return NotFound
	}

	n := uint(0)
	for x&scanTableMask == 0 {
		n += scanTableBits
		x >>= scanTableBits
	}
	return uint(scanLeftTable[x&scanTableMask]) + n
}

func slowScanLeft(x uint64) uint {
	for k := byte(0); k < 64; k += 1 {
		if x&(1<<k) != 0 {
			return uint(k)
		}
	}
	return NotFound
}
