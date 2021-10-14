package main

import (
	"fmt"
	"io"
)

func printSlice(w io.Writer, name string, values []int) {
	fmt.Fprintf(w, "%s = [", name)
	fmt.Fprint(w, values[0])
	for i := 1; i < len(values); i++ {
		fmt.Fprint(w, ", ")
		fmt.Fprint(w, values[i])
	}
	fmt.Fprintf(w, "]\n")
}

func printDot(w io.Writer, parents []int, inflows []int, subInflows []int, cutIx int) {
	fmt.Fprintln(w, "digraph sewer {")
	fmt.Fprintln(w, `  rankdir="BT"`)
	fmt.Fprintln(w, `  node [color=mediumseagreen, fillcolor=palegreen, style=filled, penwidth=3, fontname="Arial"]`)
	fmt.Fprintln(w, "  edge [penwidth=2]")
	for i := 0; i < len(parents); i++ {
		fmt.Fprintf(w, "  %d [label=\"%d/%d\\n(%d total)\"]\n", i, i, inflows[i], subInflows[i])
		// fmt.Fprintf(w, "  %d [label=\"%d/%d\"]\n", i, i, inflows[i])
	}
	for trg, src := range parents {
		if src < 0 {
			continue
		}
		fmt.Fprintf(w, "  %d -> %d", trg, src)
		if trg == cutIx {
			fmt.Fprint(w, " [color=red]")
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w, "}")
}
