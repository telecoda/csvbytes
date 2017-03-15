package csvbytes

import (
	"io"
	"strings"
	"testing"
)

var csvString = `header1,header2,header3
row1,value1.1,value1.2
row2,value2.1,value2.2`

func TestParser(t *testing.T) {

	expectedRows := 2

	altSeparator := ','
	var headerFields []string
	var fields [][]byte
	var err error
	input := strings.NewReader(csvString)

	r := NewReader(input)
	r.LazyQuotes = true
	r.Comma = altSeparator
	// read header row
	if fields, err = r.Read(); err != nil {
		if err == io.EOF {
			t.Fatal("CSV file is empty")
		} else {
			t.Fatal(err)
		}
	} else {
		// convert header names
		headerFields = make([]string, len(fields))
		for n, _ := range fields {
			headerFields[n] = string(fields[n])
		}
	}
	// read rows
	n := 0
	for {
		if fields, err := r.Read(); err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("Errors on line: %d %s", n, err)
		} else {
			if len(fields) != len(headerFields) {
				t.Fatalf("Expected %d columns but received %d\n", len(headerFields), len(fields))
			}
		}
		n++
	}

	if n != expectedRows {
		t.Fatalf("Expected %d lines but received %d\n", expectedRows, n)
	}

}

func BenchmarkParser(b *testing.B) {

	expectedRows := 2

	altSeparator := ','
	var headerFields []string
	var fields [][]byte
	var err error

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		input := strings.NewReader(csvString)
		r := NewReader(input)
		r.LazyQuotes = true
		r.Comma = altSeparator
		// read header row
		if fields, err = r.Read(); err != nil {
			if err == io.EOF {
				b.Fatal("CSV file is empty")
			} else {
				b.Fatal(err)
			}
		} else {
			// convert header names
			headerFields = make([]string, len(fields))
			for n, _ := range fields {
				headerFields[n] = string(fields[n])
			}
		}
		// read rows
		n := 0
		for {
			if fields, err := r.Read(); err != nil {
				if err == io.EOF {
					break
				}
				b.Fatalf("Errors on line: %d %s", n, err)
			} else {
				if len(fields) != len(headerFields) {
					b.Fatalf("Expected %d columns but received %d\n", len(headerFields), len(fields))
				}
			}
			n++
		}

		if n != expectedRows {
			b.Fatalf("Expected %d lines but received %d\n", expectedRows, n)
		}

	}

}
