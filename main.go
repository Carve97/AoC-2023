package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var lineReader = bufio.NewScanner(bufio.NewReader(os.Stdin))
var writer = bufio.NewWriter(os.Stdout)

func read(vals ...any) bool {
	_, err := fmt.Fscan(reader, vals...)
	return err == nil
}

func readFromReader(r io.Reader, vals ...any) bool {
	_, err := fmt.Fscan(r, vals...)
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

func lastElem[T any](slice []T) T {
	var t T
	if len(slice) == 0 {
		return t
	}
	return slice[len(slice)-1]
}

func main() {
	defer writer.Flush()
	lineReader.Split(bufio.ScanLines)

	Solve14_2()
}
