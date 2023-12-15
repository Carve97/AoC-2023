package main

import (
	"bufio"
	"os"
)

func Solve15() {
	var res int64
	var cur int64
	f, _ := os.Open("./data/15")
	reader = bufio.NewReader(f)
	for {
		b, err := reader.ReadByte()
		if err != nil || b == ',' {
			res += cur
			cur = 0
			if err != nil {
				break
			}
			continue
		}

		cur += int64(b)
		cur = (cur * 17) % 256
	}
	print(res)
}

func Solve15_2() {
	var hash int
	var label []byte
	type lens struct {
		pos   int
		focal int
	}
	boxes := make([]map[string]*lens, 256)

	getBox := func(hash int) map[string]*lens {
		if boxes[hash] == nil {
			boxes[hash] = make(map[string]*lens)
		}
		return boxes[hash]
	}

	addToBox := func(hash int, label string, focalByte byte) {
		box := getBox(hash)
		focal := int(focalByte - '0')
		if lens, in := box[label]; in {
			lens.focal = focal
			return
		}
		box[label] = &lens{pos: len(box) + 1, focal: focal}
	}
	removeFromBox := func(hash int, label string) {
		box := getBox(hash)
		if val, in := box[label]; in {
			delete(box, label)

			// shift the lens that were behind this one
			for _, v := range box {
				if v.pos > val.pos {
					v.pos--
				}
			}
		}
	}

	f, _ := os.Open("./data/15")
	reader = bufio.NewReader(f)
	for {
		b, err := reader.ReadByte()
		if b == '-' || b == '=' {
			if b == '=' {
				focal, _ := reader.ReadByte()
				addToBox(hash, string(label), focal)
			}
			if b == '-' {
				removeFromBox(hash, string(label))
			}

			// next will be either ',' or EOF
			_, err = reader.ReadByte()
			if err != nil {
				break
			}

			label = []byte{}
			hash = 0
			continue
		}

		label = append(label, b)
		hash += int(b)
		hash = (hash * 17) % 256
	}

	var res int64
	for i, box := range boxes {
		for _, lens := range box {
			res += int64((i + 1) * lens.pos * lens.focal)
		}
	}
	print(res)
}
