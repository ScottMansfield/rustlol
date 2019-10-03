package main

import (
	"fmt"
	"os/exec"
  "strconv"
	"testing"
)

func genBench(filename string) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			exec.Command("rustc", filename).Run()
		}
	}
}

func benchRange(b *testing.B, start, end, inc int) {
	for i := start; i < end; i += inc {
		b.Run(strconv.Itoa(i), genBench(fmt.Sprintf("print_%v.rs", i)))
	}
}

func BenchmarkOnes(b *testing.B) {
  benchRange(b, 1, 10, 1)
}

func BenchmarkTens(b *testing.B) {
  benchRange(b, 10, 100, 10)
}

func BenchmarkHundreds(b *testing.B) {
  benchRange(b, 100, 1000, 100)
}

func BenchmarkThousands(b *testing.B) {
  benchRange(b, 1000, 10000, 1000)
}

func BenchmarkTenThousands(b *testing.B) {
  benchRange(b, 10000, 100000, 10000)
}

func BenchmarkHundredThousands(b *testing.B) {
  benchRange(b, 100000, 1000000, 100000)
}

func BenchmarkMillion(b *testing.B) {
  benchRange(b, 1000000, 1000001, 1)
}
