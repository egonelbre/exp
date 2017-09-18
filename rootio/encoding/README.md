Results:

```
goos: windows
goarch: amd64
pkg: github.com/egonelbre/exp/rootio/encoding

// Naming
//   B / L = big endian / little endian
//   A / S = (at, data) / (data[at:])
//   F / R = forward vs reverse
//   U = no bounds checks

binary.BigEndian-8              2000000               654 ns/op
BAR-8                           1000000              1303 ns/op
BAF-8                           1000000              1297 ns/op
LAF-8                           1000000              1298 ns/op
LAR-8                           1000000              1302 ns/op
NAD-8                           2000000               652 ns/op
NADU-8                          5000000               340 ns/op
BSR-8                           2000000               654 ns/op
BSF-8                           1000000              1618 ns/op
binary.LittleEndian-8           2000000               664 ns/op
LSF-8                           1000000              1633 ns/op
LSR-8                           2000000               653 ns/op
NSD-8                           1000000              1337 ns/op
NSDU-8                          1000000              1334 ns/op
```