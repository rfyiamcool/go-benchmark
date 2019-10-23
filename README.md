# go-benchmark

golang lib benchmark


## Other Benchmark test

[batch notify channel](batch_notify_channel/README.md)


## Benchmark

hash

```
go test -bench=. -run=none -benchmem hash/hash_test.go
goos: darwin
goarch: amd64

BenchmarkFnv-4         	100000000	        22.1 ns/op	       4 B/op	       1 allocs/op
BenchmarkCrc32-4       	50000000	        26.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmur32-4    	20000000	        65.0 ns/op	      80 B/op	       1 allocs/op
BenchmarkFnvString-4   	200000000	         6.45 ns/op	       0 B/op	       0 allocs/op
```

string plus

```
BenchmarkStringsJoin-4      	20000000	        76.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkAppendOperator-4   	 5000000	       368 ns/op	    1104 B/op	       5 allocs/op
BenchmarkHardCoding-4       	20000000	        65.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkPrintf-4           	 5000000	       358 ns/op	     160 B/op	       7 allocs/op
BenchmarkByteArray-4        	10000000	       133 ns/op	     176 B/op	       4 allocs/op
BenchmarkCapByteArray-4     	30000000	        52.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkBytesBuffer-4      	10000000	       134 ns/op	     128 B/op	       2 allocs/op
BenchmarkCapBytesBuffer-4   	20000000	        73.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkStringBuffer-4     	20000000	        95.9 ns/op	     128 B/op	       2 allocs/op
BenchmarkStringBuilder-4    	20000000	       116 ns/op	     112 B/op	       3 allocs/op
```

interface assert

```
Benchmark_NormalSwitch-4      	2000000000	         1.61 ns/op	       0 B/op	       0 allocs/op
Benchmark_InterfaceSwitch-4   	200000000	         9.94 ns/op	       0 B/op	       0 allocs/op
Benchmark_TypeSwitch-4        	100000000	        18.7 ns/op	       0 B/op	       0 allocs/op
```

defer

```
BenchmarkDefer-4            	30000000	        46.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeferLoop3-4       	20000000	       114 ns/op	       0 B/op	       0 allocs/op
BenchmarkLastCallLoop3-4    	200000000	         6.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeferLoop5-4       	10000000	       191 ns/op	       0 B/op	       0 allocs/op
BenchmarkLastCallLoop5-4    	200000000	         8.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeferLoop10-4      	 5000000	       370 ns/op	       0 B/op	       0 allocs/op
BenchmarkLastCallLoop10-4   	100000000	        15.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeferLoop20-4      	 2000000	       837 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeferNo-4          	2000000000	         1.79 ns/op	       0 B/op	       0 allocs/op
```

reflect select

```
BenchmarkReflectSelect-4           10000            294436 ns/op           99134 B/op       1081 allocs/op
BenchmarkGoSelect-4                  100          15649996 ns/op            4125 B/op         36 allocs/op
```
