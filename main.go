package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tusupov/gostringbfs/list"
)

func main() {

	intputFilePath := flag.String("i", "github.com/tusupov/gostringbfs/input_big.txt", "Input file path")
	findFilePath := flag.String("f", "github.com/tusupov/gostringbfs/find_big.txt", "Find file path")
	flag.Parse()

	inFile, err := os.Open(*intputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	findFile, err := os.Open(*findFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer findFile.Close()

	in := bufio.NewScanner(inFile)
	find := bufio.NewScanner(findFile)

	substrMap := list.New()

	// Read substring list
	for in.Scan() {
		line := in.Text()
		listLine := strings.Split(line, ",")
		if len(listLine) != 2 || len(listLine[0]) == 0 || len(listLine[1]) == 0 {
			continue
		}
		substrMap.Put(listLine[1], listLine[0])
	}

	// Read find text
	for find.Scan() {
		needText := find.Text()
		if len(needText) == 0 {
			continue
		}
		if result, ok := substrMap.FindText(needText); !ok {
			fmt.Println("-1")
		} else {
			fmt.Println(strings.Join(result, ","))
		}
	}

}
