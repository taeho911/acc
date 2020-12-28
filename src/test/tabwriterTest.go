package main

import (
	"fmt"
	"text/tabwriter"
	"os"
)

type person struct {
	index uint
	name string
	job string
	age uint
}

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	
	kim := person{
		index: 1,
		name: "Kim Taeho",
		job: "Application developer",
		age: 28,
	}

	fmt.Println("kim:", kim)

	fmt.Fprintf(w, "INDEX\tNAME\tJOB\tAGE\n")
	fmt.Fprintf(w, "%d\t%s\t%s\t%d\n", kim.index, kim.name, kim.job, kim.age)
	w.Flush()
}