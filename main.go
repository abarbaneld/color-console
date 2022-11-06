package main

import (
	"fmt"
	"github.com/abarbaneld/color-console/color"
)


func main() {
    fmt.Println(color.White + "This is White" + color.Reset)
    fmt.Println(color.Red + "This is Red" + color.Reset)
    fmt.Println(color.Green + "This is Green" + color.Reset)
    fmt.Println(color.Yellow + "This is Yellow" + color.Reset)
    fmt.Println(color.Blue + "This is Blue" + color.Reset)
    fmt.Println(color.Purple + "This is Purple" + color.Reset)
    fmt.Println(color.Cyan + "This is Cyan" + color.Reset)
    fmt.Println(color.Gray + "This is Gray" + color.Reset)
	fmt.Println("----------------------")
	color.Println(color.White, "This is White")
    color.Println(color.Red, "This is Red")
    color.Println(color.Green, "This is Green")
    color.Println(color.Yellow, "This is Yellow")
    color.Println(color.Blue, "This is Blue")
    color.Println(color.Purple,"This is Purple")
    color.Println(color.Cyan,"This is Cyan")
    color.Println(color.Gray,"This is Gray")	
	
}