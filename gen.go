package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	genSet(1, 10, 1)
	genSet(10, 100, 10)
	genSet(100, 1000, 100)
	genSet(1000, 10000, 1000)
	genSet(10000, 100000, 10000)
	genSet(100000, 1000000, 100000)
	genSet(1000000, 1000001, 1)
}

func genSet(start, end, inc int) {
	for i := start; i < end; i += inc {
		genFile(i)
	}
}

func genFile(n int) {
	filename := fmt.Sprintf("print_%v.rs", n)
	fmt.Println(filename)

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	w.WriteString("fn main() {\n")
	w.WriteString("    use std::io::Write;\n")
	w.WriteString("    static HELLO_WORLD: &[u8] = b\"Hello, world!\\n\";\n")
	w.WriteString("    let stdout = std::io::stdout();\n")
	w.WriteString("    let mut out = stdout.lock();\n")

	for i := 0; i < n; i++ {
		w.WriteString("    out.write_all(HELLO_WORLD).ok();\n")
	}

	w.WriteString("}\n")

	w.Flush()
	f.Close()
}
