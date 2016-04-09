package main

import "fmt"
import "bufio"
import "os"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el texto: ")
	text, _ := reader.ReadString('\n')
	s := []byte(text)
	for i:=0;i<len(s);i++{
		fmt.Println(string(s[0:i]))
	}
	for i:=len(s)-2;i>0;i--{
		fmt.Println(string(s[0:i]))
	}
}
