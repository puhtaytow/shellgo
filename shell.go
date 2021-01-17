package shellgo

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
)

func STDINReader() []rune {
	var output []rune
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return output
}

func Parse(input string) string {
	var buffer bytes.Buffer
	r, _ := regexp.Compile("\t[0-9a-fA-F]+")
	matched := r.FindAllString(input, -1)

	for _, e := range matched {
		buffer.WriteString(strings.Trim(e, "\t"))
	}

	return buffer.String()
}