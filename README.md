# csvbytes
Low allocation CSV Parser which returns [][]byte

This is a modified version of the Golang standard library csv parser.

[Golang standard csv parser](https://golang.org/pkg/encoding/csv/#Reader.Read)

The difference is that this implementation returns a slice of bytes to reduce allocations

    func (r *Reader) Read() (record [][]byte, err error) 

Please handle the byte slice with care as its contents will be destroyed by the next read.


## running tests
    
    $ go test ./... -bench=. -v
    
## results
    
    === RUN   TestParser
    --- PASS: TestParser (0.00s)
    
    BenchmarkParser-8   	  500000	      2869 ns/op    4640 B/op	      12 allocs/op
    PASS
    ok  	github.com/telecoda/csvbytes	1.474s

