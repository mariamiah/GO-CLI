package main

import (
	"os"
	"log"
	"strings"
	"fmt"
	"bufio"
	"flag"
)

func main(){
	path := flag.String("path", "mrm.log", "The path to the log that should be analysed")
	level := flag.String("level", "WARNING", "log level to search for")
	flag.Parse()
	f, err := os.Open(*path)
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level){
			fmt.Println(s)
		}
	}
}