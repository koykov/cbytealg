# CbyteAlg

Alloc-free replacements for packages [bytes](https://golang.org/pkg/bytes) and
[strings](https://golang.org/pkg/strings) based on [cbyte](https://github.com/koykov/cbyte).

Also package provides BatchReplaceStr, - alloc-free replacement for [strings.Replacer](https://golang.org/pkg/strings/#Replacer).
Byte slice batch replacer available under name BatchReplace.

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

## BatchReplace

In fact it isn't a replacement of [strings.Replacer](https://golang.org/pkg/strings/#Replacer) since vanilla replacer
made for concurrent use, whereas BatchReplacer made to reduce allocations for big lists of replacements.

Usage example:

```go
originalStr := "this IS a string that contains {tag0}, {tag1}, tag2 and #s"
expectStr := "this WAS a string that contains 'very long substring', 1234567890, 154.195628217573 and etc..."

// Use pool instead of direct using of NewBatchReplace() or NewBatchReplaceStr().
// Pool may help you to get zero allocations on long distance and under high load.
r := cbytealg.BatchStrPool.Get(originalStr)
defer cbytealg.BatchStrPool.Put(r)
res := r.Replace("IS", "WAS").
    Replace("{tag0}", "'very long substring'").
    ReplaceInt("{tag1}", int64(1234567890)).
    ReplaceFloat("tag2", float64(154.195628217573)).
    Replace("#s", "etc...").
    Commit()
fmt.Println(res == expectStr) // true
```

## Benchmarks

```
BenchmarkBatchReplaceStr_Replace-8          2000000       830 ns/op       0 B/op       0 allocs/op
BenchmarkBatchReplaceStrNative_Replace-8    2000000       922 ns/op     544 B/op      10 allocs/op
BenchmarkBatchReplace_Replace-8             2000000       805 ns/op       0 B/op       0 allocs/op
BenchmarkBatchReplaceNative_Replace-8       2000000       775 ns/op     304 B/op       6 allocs/op
BenchmarkTrim-8                            20000000      80.6 ns/op       0 B/op       0 allocs/op
BenchmarkTrim_Native-8                     10000000       141 ns/op      48 B/op       2 allocs/op
BenchmarkSplit-8                            5000000       241 ns/op       0 B/op       0 allocs/op
BenchmarkSplit_Native-8                    10000000       149 ns/op      80 B/op       1 allocs/op
BenchmarkJoin-8                            10000000       216 ns/op       0 B/op       0 allocs/op
BenchmarkJoin_Native-8                     20000000      66.8 ns/op      16 B/op       1 allocs/op
BenchmarkReplace-8                         10000000       230 ns/op       0 B/op       0 allocs/op
BenchmarkReplace_Native-8                  20000000       101 ns/op      32 B/op       1 allocs/op
BenchmarkRepeat-8                             20000     60513 ns/op       0 B/op       0 allocs/op
BenchmarkRepeat_Native-8                      10000    135109 ns/op  655364 B/op       1 allocs/op
BenchmarkTrimStr-8                         20000000      90.5 ns/op       0 B/op       0 allocs/op
BenchmarkTrimStr_Native-8                  10000000       151 ns/op      48 B/op       2 allocs/op
BenchmarkSplitStr-8                         5000000       244 ns/op       0 B/op       0 allocs/op
BenchmarkSplitStr_Native-8                 10000000       140 ns/op      48 B/op       1 allocs/op
BenchmarkJoinStr-8                         10000000       219 ns/op       0 B/op       0 allocs/op
BenchmarkJoinStr_Native-8                  20000000      72.4 ns/op      16 B/op       1 allocs/op
BenchmarkReplaceStr-8                       5000000       252 ns/op       0 B/op       0 allocs/op
BenchmarkReplaceStr_Native-8               10000000       131 ns/op      64 B/op       2 allocs/op
BenchmarkRepeatStr-8                          20000     56727 ns/op       0 B/op       0 allocs/op
BenchmarkRepeatStr_Native-8                   10000    136147 ns/op  655363 B/op       1 allocs/op
```

As you see, cbytealg produces zero allocations during work, but usually works a bit longer than vanilla versions.

Well it's a acceptable cost to reduce GC pressure.
