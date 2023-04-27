# Socket

Socket contains benchmarks for different network sockets. Each benchmark iteration re-uses the same socket, so initial connection overhead on streamed data is negated

## Example Results

### Mac

```
goos: darwin
goarch: arm64
pkg: socket/serve
BenchmarkNetworks/unix-8                  317290              3806 ns/op              16 B/op          1 allocs/op
BenchmarkNetworks/unixgram-8              179508              6331 ns/op             144 B/op          2 allocs/op
BenchmarkNetworks/tcp-8                   133105              8911 ns/op              16 B/op          1 allocs/op
BenchmarkNetworks/udp-8                    90218             11716 ns/op              80 B/op          3 allocs/op
```

### Linux VM
```
goos: linux
goarch: amd64
pkg: socket/serve
cpu: Intel(R) Core(TM) i9-10885H CPU @ 2.40GHz
BenchmarkNetworks/unix-4                  104856             13194 ns/op              16 B/op          1 allocs/op
BenchmarkNetworks/unixgram-4               76328             13566 ns/op              16 B/op          1 allocs/op
BenchmarkNetworks/udp-4                    23764             53888 ns/op              80 B/op          3 allocs/op
BenchmarkNetworks/tcp-4                     9110            121897 ns/op              16 B/op          1 allocs/op
```