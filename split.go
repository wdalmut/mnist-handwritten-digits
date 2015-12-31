package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	var input string
	var out string
	var tt string

	flag.StringVar(&input, "in", "", "The mnist source image file")
	flag.StringVar(&out, "out", ".", "The output folder")
	flag.StringVar(&tt, "type", "img", "The file type (img, label)")

	flag.Parse()

	data, err := ioutil.ReadFile(input)

	if err != nil {
		log.Fatalf("Unable to open file: %s", input)
	}

	switch tt {
	case "img":
	case "image":
		splitImages(data, out)
		break
	case "label":
	case "txt":
		splitText(data, out)
		break
	}

	os.Exit(0)
}

func splitText(data []byte, out string) {
	magicNumber := extractUint32(data, 0)
	numberOfItems := extractUint32(data, 4)

	fmt.Printf("Magic number %d, no %d\n", magicNumber, numberOfItems)

	for i := 0; i < numberOfItems; i++ {
		el := data[8+i]

		ioutil.WriteFile(fmt.Sprintf("%s/%d.txt", out, i), []byte(strconv.Itoa(int(el))), 0755)
	}
}

func splitImages(data []byte, out string) {
	magicNumber := extractUint32(data, 0)
	numberOfItems := extractUint32(data, 4)
	h := extractUint32(data, 8)
	w := extractUint32(data, 12)

	fmt.Printf("Magic number %d, no %d, %dx%d\n", magicNumber, numberOfItems, h, w)

	for i := 0; i < numberOfItems; i++ {
		start := 16 + i*h*w
		stop := start + h*w

		ioutil.WriteFile(fmt.Sprintf("%s/%d.data", out, i), data[start:stop], 0755)
	}
}

func extractUint32(data []byte, pos int) int {
	var magicNumber uint32

	magicNumber = uint32(data[pos+3])
	magicNumber += ((uint32(data[pos+2])) << 8)
	magicNumber += ((uint32(data[pos+1])) << 16)
	magicNumber += ((uint32(data[pos+0])) << 24)

	return int(magicNumber)
}
