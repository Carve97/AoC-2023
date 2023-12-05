package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var lineReader = bufio.NewScanner(bufio.NewReader(os.Stdin))
var writer = bufio.NewWriter(os.Stdout)

func read(vals ...any) bool {
	_, err := fmt.Fscan(reader, vals...)
	return err == nil
}

func readLine(val *string) bool {
	res := lineReader.Scan()
	*val = lineReader.Text()
	return res
}

func print(vals ...any) {
	fmt.Fprintln(writer, vals...)
}

func main() {
	defer writer.Flush()
	lineReader.Split(bufio.ScanLines)

	Solve5()
}
