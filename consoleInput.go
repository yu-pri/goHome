package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"goHome/home"
	"log"
	"os"
	"strings"
)

func processInput(x *home.HData) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("#: ")
		text, _ := reader.ReadString('\n')

		if err := json.NewDecoder(strings.NewReader(text)).Decode(x); err != nil {
			log.Println(err)
			continue
		}
		reportHData(x)
		//fmt.Println(x)
	}
}
