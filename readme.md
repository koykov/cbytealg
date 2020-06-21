# CbyteAlg

Alloc-free replacements for packages [bytes](https://golang.org/pkg/bytes) and
[strings](https://golang.org/pkg/strings) based on [cbyte](https://github.com/koykov/cbyte).

Note, this is an experimental repo.

## Index

Currently supported only replacements for most popular function that produces
allocations:
* Trim()
* TrimLeft()
* TrimRight()
* Split()
* SplitN()
* Join()
* Replace()

All function below is a replacements for corresponding function from [bytes](https://golang.org/pkg/bytes) package.
String versions are available under the same names + suffix "Str", e.g. TrimStr().

## Usage

Use functions of this package the same as vanilla versions, but keep in mind that results are produced by
[cbyte](https://github.com/koykov/cbyte) and need to be released manually.

Example:
```go
chunks := []string{"foo", "bar", "string"}
s := cbytealg.JoinStr(chunks, " ")
fmt.Println(s) // "foo bar string"
// Don't forget to call release functions after work.
// You will caught memory leak otherwise.
cbyte.ReleaseStr(s)
``` 

The only exceptions is a trim functions. They return non-cbyte strings and bytes and they can't be released manually. 

## Benchmarks

```
BenchmarkSplit-8               	 5000000	       257 ns/op	       0 B/op	       0 allocs/op
BenchmarkSplit_Native-8        	10000000	       140 ns/op	      80 B/op	       1 allocs/op
BenchmarkJoin-8                	10000000	       206 ns/op	       0 B/op	       0 allocs/op
BenchmarkJoin_Native-8         	20000000	        64.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkReplace-8             	10000000	       231 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace_Native-8      	20000000	        94.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkTrimStr_Native-8      	10000000	       136 ns/op	      48 B/op	       2 allocs/op
BenchmarkSplitStr-8            	 5000000	       243 ns/op	       0 B/op	       0 allocs/op
BenchmarkSplitStr_Native-8     	10000000	       132 ns/op	      48 B/op	       1 allocs/op
BenchmarkAppendSplitStr-8      	30000000	        56.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkJoinStr-8             	10000000	       221 ns/op	       0 B/op	       0 allocs/op
BenchmarkJoinStr_Native-8      	20000000	        69.7 ns/op	      16 B/op	       1 allocs/op
BenchmarkReplaceStr-8          	 5000000	       243 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplaceStr_Native-8   	10000000	       130 ns/op	      64 B/op	       2 allocs/op
```

As you see, cbytealg produces zero allocations during work, but usually works a bit longer than vanilla versions.

Well it's a acceptable cost to reduce GC pressure.
