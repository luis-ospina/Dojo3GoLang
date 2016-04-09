package main

import "fmt"
import "bufio"
import "os"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el texto a repetir: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
