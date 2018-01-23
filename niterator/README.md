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