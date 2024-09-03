package main

import "github.com/sirupsen/logrus"
import "github.com/fatih/color"

func main() {
	logrus.Println("Hello, Hexlet!")
	color.Cyan("Prints text in cyan.")
	red := color.New(color.FgRed).PrintfFunc()
	red("Warning \n")
	red("Fire\n")
}
