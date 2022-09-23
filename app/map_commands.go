package app

import (
	"fmt"
	"github.com/joycezemitchell/quiz_master/commands"
	"os"
	"strings"
)

// RunCommand ...
func RunCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	quoted := false

	arrCommandStr := strings.FieldsFunc(commandStr, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return !quoted && r == ' '
	})

	var err error

	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case commands.HELP:
		err = commands.Help()
	case commands.CREATE_QUESTIONS:
		err = commands.CreateQuestion(arrCommandStr)
	case commands.UPDATE_QUESTION:
		err = commands.UpdateQuestion(arrCommandStr)
	case commands.QUESTION:
		err = commands.GetQuestionsByID(arrCommandStr)
	case commands.QUESTIONS:
		err = commands.GetAllQuestions()
	case commands.DELETE_QUESTION:
		err = commands.DeleleteQuestion(arrCommandStr)
	case commands.ANSWER_QUESTION:
		err = commands.AnswerQuestion(arrCommandStr)
	default:
		fmt.Println("Bad command or filename")
	}

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
