```
BenchmarkBasic-4                   	    5000	   2937549 ns/op	 196.08 MB/s
BenchmarkInstruct-4                	    5000	   3224819 ns/op	 178.61 MB/s
BenchmarkOrdone-4                  	    5000	   2876686 ns/op	 200.23 MB/s
BenchmarkOnearr-4                  	    5000	   2663603 ns/op	 216.25 MB/s
BenchmarkOnearrRev-4               	   10000	   2227535 ns/op	 258.58 MB/s
BenchmarkOnearrRevSpecialize-4     	   10000	   1969871 ns/op	 292.40 MB/s
BenchmarkPremul-4                  	    5000	   3438574 ns/op	 167.51 MB/s
BenchmarkOnearrPremul-4            	    5000	   2774488 ns/op	 207.61 MB/s
BenchmarkSpecialize-4              	    5000	   2568444 ns/op	 224.26 MB/s
BenchmarkUnroll-4                  	   10000	   2059076 ns/op	 279.74 MB/s
BenchmarkUnrollReverse-4           	   10000	   2104658 ns/op	 273.68 MB/s
BenchmarkUnrollInverse-4           	   10000	   1968053 ns/op	 292.67 MB/s
BenchmarkUnrollInReverse-4         	   10000	   2002197 ns/op	 287.68 MB/s
BenchmarkUnrollInReverseSwitch-4   	    5000	   2521456 ns/op	 228.44 MB/s
BenchmarkUnrollInReverseBool-4     	   10000	   1922431 ns/op	 299.62 MB/s
```

Full table:
```
go1.10beta2
windows/amd64
i7-2820QM CPU @ 2.30GHz

name                                 3x20x40x24    24x20x40x3  24x20x40x30  5x50x100x150  150x100x50x5  100x100x100x100  Average
BenchmarkUnrollInReverseHardcode     N.A           N.A         N.A          N.A           N.A           282.48           282.48
BenchmarkUnrollInReverse             238.79        221.01      242.10       242.79        236.88        250.09           238.61
BenchmarkUnrollInverse               238.18        208.09      226.56       258.80        213.70        253.52           233.14
BenchmarkUnrollInReverseAdvance      221.97        227.76      201.82       253.53        238.25        246.09           231.57
BenchmarkUnrollReverse               231.67        209.97      220.50       235.31        221.87        238.73           226.34
BenchmarkOnearrRevSpecializeAdvance  228.30        228.07      240.76       241.28        162.22        231.55           222.03
BenchmarkUnroll                      215.60        211.02      186.14       238.95        225.02        240.08           219.47
BenchmarkOnearrRevAdvance            233.46        212.61      219.58       234.32        166.89        231.00           216.31
BenchmarkOnearrRev                   229.44        198.01      222.65       240.77        149.54        237.23           212.94
BenchmarkOnearrRevSpecialize         210.98        217.43      231.30       208.71        195.63        202.00           211.01
BenchmarkUnrollInReverseBool         213.17        209.72      235.69       198.79        207.41        196.46           210.21
BenchmarkUnrollInReverseSwitch       179.38        175.98      178.33       201.74        187.58        199.93           187.16
BenchmarkOnearr                      183.30        165.18      187.66       203.35        178.98        199.00           186.25
BenchmarkSpecialize                  187.10        175.33      187.67       190.53        178.58        194.64           185.64
BenchmarkOnearrPremul                179.19        166.17      182.14       192.40        179.46        192.58           181.99
BenchmarkBasic                       159.80        149.56      162.46       170.60        158.04        173.00           162.24
BenchmarkOrdone                      169.01        148.75      162.51       174.08        142.52        173.45           161.72
BenchmarkInstruct                    150.63        135.21      149.15       158.70        147.27        157.67           149.77
BenchmarkPremul                      134.78        130.76      137.08       143.73        136.75        143.47           137.76
```