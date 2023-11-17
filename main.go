package main

import (
    "flag"
    "fmt"
    "github.com/fatih/color"
)

func main() {
    // Define command line arguments
    textPtr := flag.String("text", "", "Text to print")
    colorPtr := flag.String("color", "red", "Color of the text")

    // Parse command line arguments
    flag.Parse()

    // Choose color
    var printFunc func(a ...interface{}) string
    switch *colorPtr {
    case "red":
        printFunc = color.New(color.FgRed).SprintFunc()
    case "green":
        printFunc = color.New(color.FgGreen).SprintFunc()
    case "blue":
        printFunc = color.New(color.FgBlue).SprintFunc()
    default:
        fmt.Println("Invalid color option. Choose from 'red', 'green', 'blue'.")
        return
    }

    // Print colored text
    fmt.Println(printFunc(*textPtr))
}
