package table

import (
	"io"
	"strings"
)

type Table struct {
	line [][]string
}

func New() *Table {
	return &Table{}
}

func (t *Table) Add(line []string) {
	t.line = append(t.line, line)
}

func (t *Table) Render(out io.Writer) error {
	if len(t.line) == 0 {
		return nil
	}
	max := make([]int, len(t.line[0]))
	for _, line := range t.line {
		for i, field := range line {
			if l := len(field); max[i] < l {
				max[i] = l
			}
		}
	}
	for i, _ := range max {
		if i == len(max)-1 {
			max[i] = 0
		} else {
			max[i] += 1
		}
	}
	LF := []byte("\n")
	for _, line := range t.line {
		for i, field := range line {
			_, err := out.Write([]byte(pad(field, max[i])))
			if err != nil {
				return err
			}
		}
		_, err := out.Write(LF)
		if err != nil {
			return err
		}
	}
	return nil
}

func pad(str string, size int) string {
	if len(str) >= size {
		return str
	}
	return str + strings.Repeat(" ", size-len(str))
}
