package main

/*
 * Terminology
 * ===========
 *
 * The sudoku is made up 81 cells organised in 9 rows and 9 cols
 * These are referred to as the board or b for short.
 * A 3x3 cells group is called a box. There are 9 boxes.
 *
 * For strategies that can use row/col/box interchangably these can
 * be accessed as 27 groups i.e. 9 rows, 9 cols, and 9 boxes.
 */

import (
	"fmt"
	"strings"
)

type Cell struct {
	row       int
	col       int
	box       int
	possibles []string
}

// --- Methods of Cell ---
func (c Cell) solved() bool {
	return len(c.possibles) == 1
}

func (cell Cell) inRow(row int) bool {
	return cell.row == row
}

func (cell Cell) inCol(col int) bool {
	return cell.col == col
}

func (cell Cell) inBox(box int) bool {
	return cell.box == box
}

func (cell *Cell) solve(value string) {
	cell.possibles = []string{value}
}

func (cell *Cell) hasPossible(s string) bool {
	if cell.solved() {
		return false
	}
	for _, possible := range cell.possibles {
		if possible == s {
			return true
		}
	}
	return false
}

func (cell *Cell) removePossible(value string) bool {
	if cell.solved() {
		return false
	}
	newPossibles := []string{}
	result := false
	for _, possible := range cell.possibles {
		if possible == value {
			result = true
		} else {
			newPossibles = append(newPossibles, possible)
		}
	}
	if len(newPossibles) != len(cell.possibles) {
		// fmt.Printf("Setting cell %v %v from %v to %v\n", cell.row, cell.col, cell.possibles, newPossibles)
		cell.possibles = newPossibles
	}
	return result
}

func (cell *Cell) removePossibles(values []string) bool {
	result := false
	for _, value := range values {
		if cell.removePossible(value) {
			result = true
		}
	}
	return result
}

func (cell *Cell) hasAtLeastOneOf(possibles []string) bool {
	for _, possible := range possibles {
		if cell.hasPossible(possible) {
			return true
		}
	}
	return false
}

func (cell *Cell) hasAllOf(possibles []string) bool {
	if cell.solved() {
		return false
	}
	for _, possible := range possibles {
		if !cell.hasPossible(possible) {
			return false
		}
	}
	return true
}

// -----------------------

func removeFromCells(cells []*Cell, values []string) bool {
	found := false
	for _, cell := range cells {
		if cell.removePossibles(values) {
			found = true
		}
	}
	return found
}

func filterInclude(cells []*Cell, include func(*Cell) bool) []*Cell {
	result := []*Cell{}
	for _, c := range cells {
		if include(c) {
			result = append(result, c)
		}
	}
	return result
}

func filterExclude(cells []*Cell, include func(*Cell) bool) []*Cell {
	result := []*Cell{}
	for _, c := range cells {
		if !include(c) {
			result = append(result, c)
		}
	}
	return result
}

func filterHasPossible(cells []*Cell, s string) []*Cell {
	result := []*Cell{}
	for _, c := range cells {
		if c.hasPossible(s) {
			result = append(result, c)
		}
	}
	return result
}

var b [81]Cell
var rows = [][]*Cell{
	{&b[0], &b[1], &b[2], &b[3], &b[4], &b[5], &b[6], &b[7], &b[8]},
	{&b[9], &b[10], &b[11], &b[12], &b[13], &b[14], &b[15], &b[16], &b[17]},
	{&b[18], &b[19], &b[20], &b[21], &b[22], &b[23], &b[24], &b[25], &b[26]},
	{&b[27], &b[28], &b[29], &b[30], &b[31], &b[32], &b[33], &b[34], &b[35]},
	{&b[36], &b[37], &b[38], &b[39], &b[40], &b[41], &b[42], &b[43], &b[44]},
	{&b[45], &b[46], &b[47], &b[48], &b[49], &b[50], &b[51], &b[52], &b[53]},
	{&b[54], &b[55], &b[56], &b[57], &b[58], &b[59], &b[60], &b[61], &b[62]},
	{&b[63], &b[64], &b[65], &b[66], &b[67], &b[68], &b[69], &b[70], &b[71]},
	{&b[72], &b[73], &b[74], &b[75], &b[76], &b[77], &b[78], &b[79], &b[80]},
}
var cols = [][]*Cell{
	{&b[0], &b[9], &b[18], &b[27], &b[36], &b[45], &b[54], &b[63], &b[72]},
	{&b[1], &b[10], &b[19], &b[28], &b[37], &b[46], &b[55], &b[64], &b[73]},
	{&b[2], &b[11], &b[20], &b[29], &b[38], &b[47], &b[56], &b[65], &b[74]},
	{&b[3], &b[12], &b[21], &b[30], &b[39], &b[48], &b[57], &b[66], &b[75]},
	{&b[4], &b[13], &b[22], &b[31], &b[40], &b[49], &b[58], &b[67], &b[76]},
	{&b[5], &b[14], &b[23], &b[32], &b[41], &b[50], &b[59], &b[68], &b[77]},
	{&b[6], &b[15], &b[24], &b[33], &b[42], &b[51], &b[60], &b[69], &b[78]},
	{&b[7], &b[16], &b[25], &b[34], &b[43], &b[52], &b[61], &b[70], &b[79]},
	{&b[8], &b[17], &b[26], &b[35], &b[44], &b[53], &b[62], &b[71], &b[80]},
}
var boxes = [][]*Cell{
	{&b[0], &b[1], &b[2], &b[9], &b[10], &b[11], &b[18], &b[19], &b[20]},
	{&b[3], &b[4], &b[5], &b[12], &b[13], &b[14], &b[21], &b[22], &b[23]},
	{&b[6], &b[7], &b[8], &b[15], &b[16], &b[17], &b[24], &b[25], &b[26]},
	{&b[27], &b[28], &b[29], &b[36], &b[37], &b[38], &b[45], &b[46], &b[47]},
	{&b[30], &b[31], &b[32], &b[39], &b[40], &b[41], &b[48], &b[49], &b[50]},
	{&b[33], &b[34], &b[35], &b[42], &b[43], &b[44], &b[51], &b[52], &b[53]},
	{&b[54], &b[55], &b[56], &b[63], &b[64], &b[65], &b[72], &b[73], &b[74]},
	{&b[57], &b[58], &b[59], &b[66], &b[67], &b[68], &b[75], &b[76], &b[77]},
	{&b[60], &b[61], &b[62], &b[69], &b[70], &b[71], &b[78], &b[79], &b[80]},
}
var blocks = [][]*Cell{}
var numbers []string = strings.Split("123456789", "")
var combinations = [][]string{}

func init() {
	for i, _ := range b {
		row := i / 9
		col := i % 9
		box := (row / 3 * 3) + col/3
		b[i] = Cell{possibles: numbers, row: row, col: col, box: box}
	}
	for _, row := range rows {
		blocks = append(blocks, row)
	}
	for _, col := range cols {
		blocks = append(blocks, col)
	}
	for _, box := range boxes {
		blocks = append(blocks, box)
	}
	combinations = makeCombinations(numbers, 2)
}

// Found on topcoder
// Imagine all numbers from 0 to 2^len-1
// The bit patterns of these numbers are the combinations
func makeCombinations(elems []string, min int) [][]string {
	result := [][]string{}
	n := len(elems)
	for num := 0; num < (1 << uint(n)); num++ {
		combination := []string{}
		for ndx := 0; ndx < n; ndx++ {
			// (is the bit "on" in this number?)
			if num&(1<<uint(ndx)) != 0 {
				// (then add it to the combination)
				combination = append(combination, elems[ndx])
			}
		}
		if len(combination) >= min {
			result = append(result, combination)
		}
	}
	return result
}

func parse(s string) {
	if len(s) != 81 {
		fmt.Printf("!!! Parse expected length 81 but got %v\n", len(s))
	}
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c != "." {
			b[i].possibles = []string{c}
		}
	}
}

func boardSolved() bool {
	solved := true
	for _, cell := range b {
		solved = solved && cell.solved()
	}
	return solved
}

func printb() {
	fmt.Println()
	for i, _ := range b {
		if b[i].solved() {
			fmt.Printf("%-10s", "    "+b[i].possibles[0])
		} else {
			fmt.Printf("%-10s", strings.Join(b[i].possibles, ""))
		}

		if i > 0 {
			j := i + 1
			if j%9 == 0 {
				fmt.Println()
				if j == 27 || j == 54 {
					fmt.Println("==============================|==============================|==============================")
				}
			} else if j%3 == 0 {
				fmt.Print("|")
			}
		}
	}
	fmt.Println()
}

func nameOfBlock(block int) string {
	if block < 9 {
		return fmt.Sprintf("Row %v", block+1)
	} else if block < 18 {
		return fmt.Sprintf("Col %v", block-9+1)
	} else {
		return fmt.Sprintf("Box %v", block-18+1)
	}
}

func signatureOfBlock(i int) string {
	block := blocks[i]
	result := nameOfBlock(i) + ":"
	for _, cell := range block {
		result += strings.Join(cell.possibles, "") + "|"
	}
	return result
}

func removeSolved() {
	fmt.Println("=== Remove Solved")
	found := true
	for found {
		found = false
		for i, block := range blocks {
			oldSignature := signatureOfBlock(i)
			solved := []string{}
			for _, cell := range block {
				if cell.solved() {
					solved = append(solved, cell.possibles[0])
				}
			}
			if removeFromCells(block, solved) {
				found = true
			}
			newSignature := signatureOfBlock(i)
			if newSignature != oldSignature {
				// fmt.Printf("%s -> %s\n", oldSignature, newSignature)
			}
		}
	}
}

/*
 * If a cell is the only one to contain a possible then it is the solution
 */
func singles() bool {
	fmt.Println("=== Singles")
	for index, cells := range blocks {
		for _, possible := range numbers {
			matches := filterHasPossible(cells, possible)
			if len(matches) == 1 {
				fmt.Printf("Single %s in %v\n", possible, nameOfBlock(index))
				matches[0].solve(possible)
				return true
			}
		}
	}
	return false
}

/*
 * If two or more cells are the only ones to contain a combo then the combo can be removed from other cells
 */
func nakeds() bool {
	fmt.Println("=== Nakeds")
	for _, combo := range combinations {
		// Need to match combo is [1,2,3] and value is [1] or [2,3] etc.
		comboFlavours := makeCombinations(combo, 1)
		hasCombo := func(cell *Cell) bool {
			possibles := strings.Join(cell.possibles, "")
			for _, c := range comboFlavours {
				if possibles == strings.Join(c, "") {
					return true
				}
			}
			return false
		}
		for index, block := range blocks {
			matches := filterInclude(block, hasCombo)
			if len(matches) == len(combo) {
				inMatches := func(cell *Cell) bool {
					for _, match := range matches {
						if cell == match {
							return true
						}
					}
					return false
				}
				others := filterExclude(block, inMatches)
				found := removeFromCells(others, combo)
				if found {
					fmt.Printf("Naked %v found in %s\n", combo, nameOfBlock(index))
					return true
				}
			}
		}
	}
	return false
}

/*
 * If two or more cells are the only ones to contain a combo then any other possibles can be removed from those cells
 */
func hiddens() bool {
	fmt.Println("=== Hiddens")
	for _, combo := range combinations {
		hasCombo := func(cell *Cell) bool {
			return cell.hasAllOf(combo)
		}
		for index, block := range blocks {
			matches := filterInclude(block, hasCombo)
			if len(matches) == len(combo) {
				for _, match := range matches {
					match.possibles = combo
				}
				fmt.Printf("Hidden %v found in %s\n", combo, nameOfBlock(index))
				return true
			}
		}
	}
	return false
}

/*
 * If 2 or 3 cells in a box have a possible only in the same row/col then it can be removed from the rest of that row/col
 */
func pointingPairs() bool {
	fmt.Println("=== Pointing Pairs")
	for i, box := range boxes {
		cellInBox := func(cell *Cell) bool {
			return cell.box == i
		}
		for _, number := range numbers {
			matches := filterHasPossible(box, number)
			scanRow := false
			scanCol := false
			if len(matches) == 2 {
				if matches[0].row == matches[1].row {
					scanRow = true
				} else if matches[0].col == matches[1].col {
					scanCol = true
				}
			} else if len(matches) == 3 {
				if matches[0].row == matches[1].row && matches[0].row == matches[2].row {
					scanRow = true
				} else if matches[0].col == matches[1].col && matches[0].col == matches[2].col {
					scanCol = true
				}
			}

			if scanRow {
				row := matches[0].row
				others := filterExclude(rows[row], cellInBox)
				if removeFromCells(others, []string{number}) {
					fmt.Printf("Pointing pair: %v in box %v row %v\n", number, i+1, row+1)
					return true
				}
			}

			if scanCol {
				col := matches[0].col
				others := filterExclude(cols[col], cellInBox)
				if removeFromCells(others, []string{number}) {
					fmt.Printf("Pointing pair: %v in box %v col %v\n", number, i+1, col+1)
					return true
				}
			}

		}
	}
	return false
}

/*
 * If 2 or 3 cells in a row/col have a possible only in the same box then it can be removed from the rest of that box
 */
func boxLineReduction() bool {
	fmt.Println("=== Box Line Reduction")
	for i, row := range rows {
		cellInRow := func(cell *Cell) bool {
			return cell.row == i
		}
		for _, number := range numbers {
			matches := filterHasPossible(row, number)
			if len(matches) == 0 {
				continue
			}
			box := matches[0].box
			if (len(matches) == 2 && matches[1].box == box) ||
				(len(matches) == 3 && matches[1].box == box && matches[2].box == box) {
				others := filterExclude(boxes[box], cellInRow)
				if removeFromCells(others, []string{number}) {
					fmt.Printf("Box Line Reduction: %v in box %v row %v\n", number, box+1, i+1)
					return true
				}

			}
		}
	}

	for i, col := range cols {
		cellInCol := func(cell *Cell) bool {
			return cell.col == i
		}
		for _, number := range numbers {
			matches := filterHasPossible(col, number)
			if len(matches) == 0 {
				continue
			}
			box := matches[0].box
			if (len(matches) == 2 && matches[1].box == box) ||
				(len(matches) == 3 && matches[1].box == box && matches[2].box == box) {
				others := filterExclude(boxes[box], cellInCol)
				if removeFromCells(others, []string{number}) {
					fmt.Printf("Box Line Reduction: %v in box %v col %v\n", number, box+1, i+1)
					return true
				}

			}
		}
	}
	return false
}

func main() {
	strategies := []func() bool{
		singles,
		nakeds,
		hiddens,
		pointingPairs,
		boxLineReduction,
	}
	parse("5286...4913649..257942.563....1..2....78263....25.9.6.24.3..9768.97.2413.7.9.4582")
	printb()
	removeSolved()
	printb()
	if boardSolved() {
		fmt.Println("Done !!!")
		return
	}
	for {
		found := false
		for _, strategy := range strategies {
			if strategy() {
				found = true
				if boardSolved() {
					fmt.Println("Done !!!")
					return
				}
			}
			if found {
				printb()
				removeSolved()
				printb()
				if boardSolved() {
					fmt.Println("Done !!!")
					return
				}
				continue
			}
		}

		if !found {
			fmt.Println("Beats me !!!")
			return
		}
	}
}
