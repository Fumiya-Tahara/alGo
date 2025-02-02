package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() string {
	sc.Scan()

	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	d := scan()

	switch d {
	case "N":
		fmt.Println("S")
	case "S":
		fmt.Println("N")
	case "E":
		fmt.Println("W")
	case "W":
		fmt.Println("E")
	case "NE":
		fmt.Println("SW")
	case "SW":
		fmt.Println("NE")
	case "NW":
		fmt.Println("SE")
	case "SE":
		fmt.Println("NW")
	}

}
