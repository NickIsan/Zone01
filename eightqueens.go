// package piscine

// import "github.com/01-edu/z01"

// // help function
// func absolute(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

// positions := make([]int, 8)

// func isSafe(positions []int, column int, row int) bool {
// 	for i := 0; i < column; i++ {
// 		if positions[i] == row {
// 		return false
// 		}
	
// 		if abs(positions[i]-row) == column-i {
// 			return false
// 		}
	
// 	}

// 	return true
// }

// func printSolution(positions {}int) {
// 	for i := 0; i < 8; i++ {
// 		z01.PrintRune(rune(positions[i] + '0'))
// 	}
// 	z01.PrintRune('\n')
// }

// func solveQueens(positions []int, column int) {
// 	if column == 8 {
// 		printSolution(positions)
// 		return
// 	}

// 	for row := 1; row <= 8; row++ {
// 		if isSafe(positions, column, row) {
// 			positions[column] = row
// 			sollveQueens(positions, column+1)
// 		}
// 	}
// }	

// func EightQueens() {
// 	positions := make([]int, 8)
// 	solveQueens(positions, 0)
// }



// func solve(column int, positions []int)

// // already in ascending order
// for row := 1; row <= 8; row++ {

// }

// // if safe
// positions[column] = row
// solve(column+1, positions)
// // this is the recursion

// // stop the recursion
// if column == 8 {
// 	printSolution(positions)
// 	return
// }



// // Same diagonal when line difference = column difference
// if row-positions[i] == column-i
