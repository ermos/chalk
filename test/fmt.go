package main

import "github.com/ermos/chalk"

func main() {
	chalk.Println(chalk.Blue(chalk.Bluef("%s is not allowed", "toto")))
}