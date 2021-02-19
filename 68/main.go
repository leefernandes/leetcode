package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	// fill rows greedily, pad remainder with whitespaces

	count := map[int]int{0: 0}
	row := 0
	rows := [][]string{{}}
	for i := range words {
		word := words[i]
		wordLen := len(word)

		if v := count[row] + wordLen; v > maxWidth {
			// inc row, and move onto the next
			row++
			rows = append(rows, []string{})
		}

		if v := count[row]; v == 0 {
			count[row] += wordLen
		} else {
			count[row] += wordLen + 1
		}

		rows[row] = append(rows[row], word)
	}

	output := []string{}

	rowsLen := len(rows)
	for i := 0; i < rowsLen; i++ {
		row := rows[i]
		rowCharLen := count[i]
		if d := maxWidth - rowCharLen; d > 0 {
			nGaps := len(row) - 1
			if i == rowsLen-1 {
				padded := fmt.Sprintf("%-*s", maxWidth-rowCharLen, "")
				rows[i][nGaps] = row[nGaps] + padded
			} else {
				if nGaps > 0 {
					k := 0
					for j := 0; j < d; j++ {
						if k == nGaps {
							k = 0
						}

						row[k] += " "

						k++
					}

				} else {
					padded := fmt.Sprintf("%-*s", maxWidth, row[0])
					rows[i][0] = padded
				}
			}
		}

		output = append(output, strings.Join(row, " "))
	}

	return output
}
