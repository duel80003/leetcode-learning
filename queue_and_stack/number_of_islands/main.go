package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				helper(grid, i, j)
				ret++
			}
		}
	}
	return ret
}

func helper(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	fmt.Printf("i: %v, j %v, grid %v \n", i, j, string(grid[i]))

	if grid[i][j] == '1' {
		grid[i][j] = '0'
		helper(grid, i-1, j)
		helper(grid, i, j-1)
		helper(grid, i+1, j)
		helper(grid, i, j+1)
	}
}
func main() {
	// a1 := [][]byte{}
	// row1 := []byte{'1', '1', '1', '1', '0'}
	// row2 := []byte{'1', '1', '0', '1', '0'}
	// row3 := []byte{'1', '1', '0', '0', '0'}
	// row4 := []byte{'0', '0', '0', '0', '0'}
	// a1 = append(a1, row1)
	// a1 = append(a1, row2)
	// a1 = append(a1, row3)
	// a1 = append(a1, row4)
	// fmt.Println(numIslands(a1))

	a2 := [][]byte{}
	row21 := []byte{'0', '0', '0', '0'}
	row22 := []byte{'1', '1', '0', '1'}
	row23 := []byte{'0', '0', '0', '1'}
	a2 = append(a2, row21)
	a2 = append(a2, row22)
	a2 = append(a2, row23)
	fmt.Println(numIslands(a2))
	fmt.Println(a2)
}
