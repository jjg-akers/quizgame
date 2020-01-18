package main

import "flag"

func main() {
	// first add a flag
	// presents problems, accepts input, checks answers
	// get the file name - use Flage package
	//		flag name, filename default val, help text
	// Flags are great when people only have the binaries to run
	// With flags defined you can run binaries with -h to print out the help text
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the formate of question, answer")

	// tell flag to parse
	flag.Parse()
	_ = csvFilename
}
