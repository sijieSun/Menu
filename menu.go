package main

import "fmt"

func main(){
	fmt.Print("MENU:\n>>1.Command\n>>2.Quit\n")
	fmt.Println("Please input your command")
	for{
		var cmd string
		fmt.Println(">>")
		fmt.Println(&cmd)
		if cmd == "Command" {
			fmt.Println("This is a Command!")
		} else if	cmd == "Quit" {
			fmt.Println("Bye!")
		} else {
			fmt.Println("Error! Please input 'Command' or 'Quit'!")
		}
	}
}