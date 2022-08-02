```
❯ go test -bench . -benchmem -benchtime 5s  --cpu 2,32
goos: linux
goarch: amd64
pkg: github.com/raja/benchmark-id
cpu: AMD Ryzen 9 5950X 16-Core Processor
BenchmarkPoolUUID-2                               	100000000	        59.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkPoolUUID-32                              	99271846	        58.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkXIDNew-2                                 	131183766	        45.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkXIDNew-32                                	133419596	        43.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkULIDMake-2                               	78722870	        73.83 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDMake-32                              	77999992	        76.45 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDNew/WithMonotonicEntropy-2           	83727796	        70.91 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDNew/WithMonotonicEntropy-32          	83093264	        70.80 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDNew/WithCryptoMonotonicEntropy-2     	71932791	        82.21 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDNew/WithCryptoMonotonicEntropy-32    	72203370	        82.77 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDNew/WithCryptoThreadSafe-2           	70527390	        83.59 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDNew/WithCryptoThreadSafe-32          	70614345	        83.51 ns/op	      16 B/op	       1 allocs/op
BenchmarkPoolUUIDString-2                         	62119964	        96.45 ns/op	      48 B/op	       1 allocs/op
BenchmarkPoolUUIDString-32                        	62149693	        97.10 ns/op	      48 B/op	       1 allocs/op
BenchmarkULIDString-2                             	58883454	       100.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDString-32                            	59165128	        99.21 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDCryptoString-2                       	56789270	       105.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkULIDCryptoString-32                      	56335917	       105.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkParallelPoolUUIDString-2                 	59811583	       108.6 ns/op	      48 B/op	       1 allocs/op
BenchmarkParallelPoolUUIDString-32                	46424649	       128.3 ns/op	      48 B/op	       1 allocs/op
BenchmarkParallelXIDString-2                      	82259113	        73.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkParallelXIDString-32                     	430154944	        13.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkParallelULIDMakeString-2                 	63526720	       125.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkParallelULIDMakeString-32                	20212610	       297.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkParallelULIDCryptoThreadSafeString-2     	48970002	       126.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkParallelULIDCryptoThreadSafeString-32    	36091648	       166.9 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/raja/benchmark-id	174.801s
```
