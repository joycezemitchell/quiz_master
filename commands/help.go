package commands

import "fmt"

func Help() error {

	fmt.Println("Command | Description")
	fmt.Println("create_question <no> <question> <answer> | Creates a question")
	fmt.Println("update_question <no> <question> <answer> | Updates a question")
	fmt.Println("delete_question <no> | Deletes a question")
	fmt.Println("question <no> | Shows a question")
	fmt.Println("questions | Shows question list")

	return nil
}