package commands

import (
	"errors"
	"fmt"
	"github.com/joycezemitchell/quiz_master/domain/question"
	"github.com/joycezemitchell/quiz_master/services"
	"strconv"
)

func DeleleteQuestion(arrCommandStr []string) error {

	// Check parameter
	if len(arrCommandStr) != 2 {
		return errors.New(fmt.Sprintf("%s required for 1 argument", QUESTION))
	}

	i, _ := strconv.Atoi(arrCommandStr[1])

	qf, _ := services.QuestionService.GetQuestionByID(question.Question{
		Id: i,
	})

	// Check if exist
	if qf == (question.Question{}) {
		return errors.New(fmt.Sprintf("Question no %v not found", i))
	}

	err := services.QuestionService.DeleleteQuestion(question.Question{
		Id: i,
	})

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Question no %v was deleted!", i))

	return nil
}
