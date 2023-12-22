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

type keyVal[k comparable, v any] struct {
	key k
	val v
}

type intHeap[v any] struct {
	slice []keyVal[int, v]
}

func (h *intHeap[v]) Len() int {
	return len(h.slice)
}
func (h *intHeap[v]) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}
func (h *intHeap[v]) Less(i, j int) bool {
	return h.slice[i].key < h.slice[j].key
}
func (h *intHeap[v]) Push(value any) {
	val := value.(keyVal[int, v])
	h.slice = append(h.slice, val)
}
func (h *intHeap[v]) Pop() any {
	val := h.slice[len(h.slice)-1]
	h.slice = h.slice[:len(h.slice)-1]
	return val
}

func main() {
	defer writer.Flush()
	lineReader.Split(bufio.ScanLines)

	Solve19_2()
}
