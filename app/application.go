package app

import (
	"bufio"
	"fmt"
	"github.com/joycezemitchell/quiz_master/datasource/mysql/schema"
	"os"
)

const (
	WELCOME_MESSAGE = "Welcome to Quiz Master!"
)

// Run Application ...
func Run() {
	// Create schema
	schema.CreateSchema()

	fmt.Println(WELCOME_MESSAGE, "\n\n")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = RunCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

