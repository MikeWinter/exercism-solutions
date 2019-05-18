package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(spec string) (*Matrix, error) {
	m := &Matrix{}
	for i, row := range strings.Split(spec, "\n") {
		if len(strings.TrimSpace(row)) == 0 {
			return nil, fmt.Errorf("matrix: empty row %d", i)
		}
		els, err := elements(row)
		if err != nil {
			return nil, err
		}
		if i != 0 && len(els) != len((*m)[0]) {
			return nil, fmt.Errorf("matrix: uneven row length (expected %d, found %d)", len((*m)[0]), len(row))
		}
		*m = append(*m, els)
	}
	return m, nil
}

func elements(spec string) ([]int, error) {
	var row []int
	for _, el := range strings.Fields(spec) {
		v, err := strconv.Atoi(el)
		if err != nil {
			return nil, err
		}
		row = append(row, v)
	}
	return row, nil
}

func (m *Matrix) Rows() [][]int {
	var rows [][]int
	if m != nil {
		for i := range *m {
			row := make([]int, len((*m)[i]))
			copy(row, (*m)[i])
			rows = append(rows, row)
		}
	}
	return rows
}

func (m *Matrix) Cols() [][]int {
	var cols [][]int
	if m != nil {
		for i, n := 0, len((*m)[0]); i < n; i++ {
			col := make([]int, len(*m))
			for j, row := range *m {
				col[j] = row[i]
			}
			cols = append(cols, col)
		}
	}
	return cols
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) || col < 0 || col >= len((*m)[0]) {
		return false
	}
	(*m)[row][col] = val
	return true
}
