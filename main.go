package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("Apache_2k.log")
	if err != nil {
		fmt.Println("ERROR NO FILE")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		globalPos := 0
		text := scanner.Text()
		fmt.Printf("Line is: %s\n", text)
		dateStart := strings.IndexByte(text[globalPos:], '[')
		dateEnd := strings.IndexByte(text[dateStart:], ']')
		date := text[dateStart+1 : dateEnd]

		fmt.Printf("Date %s\n", date)
		globalPos = dateEnd + dateStart + globalPos + 1

		sevStart := strings.IndexByte(text[globalPos:], '[')

		sevEnd := strings.IndexByte(text[globalPos:], ']')

		severity := text[sevStart+globalPos+1 : sevEnd+globalPos]

		fmt.Printf("Severity %s\n", severity)
		globalPos = sevEnd + sevStart + 1 + globalPos

		actStart := strings.IndexByte(text[globalPos:], '[')
		if actStart == -1 {
			actStart := globalPos
			actEnd := strings.IndexByte(text[actStart:], ' ')
			actor := text[actStart : actEnd+globalPos]
			msg := text[actEnd+globalPos+1:]
			fmt.Printf("Actor %s\n", actor)
			fmt.Printf("Msg %s\n\n", msg)
		} else {
			actEnd := strings.IndexByte(text[globalPos:], ']')
			actor := text[actStart+globalPos+1 : actEnd+globalPos]
			msg := text[actEnd+1+globalPos:]
			fmt.Printf("Actor %s\n", actor)
			fmt.Printf("Msg %s\n\n", msg)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
