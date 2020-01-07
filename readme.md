```
BenchmarkBatchReplaceStr_Replace-8          1000000      1322 ns/op       0 B/op       0 allocs/op
BenchmarkBatchReplaceStrNative_Replace-8    1000000      1041 ns/op     544 B/op      10 allocs/op
BenchmarkBatchReplace_Replace-8             1000000      1257 ns/op       0 B/op       0 allocs/op
BenchmarkBatchReplaceNative_Replace-8       2000000       765 ns/op     304 B/op       6 allocs/op
BenchmarkTrim-8                            20000000      78.3 ns/op       0 B/op       0 allocs/op
BenchmarkTrim_Native-8                     10000000       129 ns/op      48 B/op       2 allocs/op
BenchmarkSplit-8                           10000000       232 ns/op       0 B/op       0 allocs/op
BenchmarkSplit_Native-8                    10000000       140 ns/op      80 B/op       1 allocs/op
BenchmarkJoin-8                            10000000       220 ns/op       0 B/op       0 allocs/op
BenchmarkJoin_Native-8                     20000000      74.8 ns/op      16 B/op       1 allocs/op
BenchmarkReplace-8                          5000000       256 ns/op       0 B/op       0 allocs/op
BenchmarkReplace_Native-8                  20000000      98.1 ns/op      32 B/op       1 allocs/op
BenchmarkTrimStr-8                         20000000      86.6 ns/op       0 B/op       0 allocs/op
BenchmarkTrimStr_Native-8                  10000000       145 ns/op      48 B/op       2 allocs/op
BenchmarkSplitStr-8                        10000000       242 ns/op       0 B/op       0 allocs/op
BenchmarkSplitStr_Native-8                 10000000       137 ns/op      48 B/op       1 allocs/op
BenchmarkJoinStr-8                         10000000       219 ns/op       0 B/op       0 allocs/op
BenchmarkJoinStr_Native-8                  20000000      71.2 ns/op      16 B/op       1 allocs/op
BenchmarkReplaceStr-8                       5000000       265 ns/op       0 B/op       0 allocs/op
BenchmarkReplaceStr_Native-8               10000000       130 ns/op      64 B/op       2 allocs/op
```
