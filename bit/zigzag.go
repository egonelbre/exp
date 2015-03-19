package bit

// ZEncode implements zig-zag encoding
// The sign will be encoded as the least significant bit in the result
// ZEncode(x) < abs(x)*2
func ZEncode(x int64) uint64 {
	ux := uint64(x)
	return (ux << 1) ^ (ux >> 63)
}

// ZDecode is the inverse operation of ZEncode
// ZDecode(ZEncode(x)) == x
func ZDecode(ux uint64) int64 {
	return int64((ux >> 1) ^ (ux << 63))
}
