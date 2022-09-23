package commands

import (
	"errors"
	"fmt"
	"github.com/joycezemitchell/quiz_master/domain/question"
	"github.com/joycezemitchell/quiz_master/services"
	"strconv"
)

func GetQuestionsByID(arrCommandStr []string) error {

	// Check parameter
	if len(arrCommandStr) != 2 {
		return errors.New(fmt.Sprintf("%s required for 1 argument", QUESTION))
	}

	i, _ := strconv.Atoi(arrCommandStr[1])
	q2, err := services.QuestionService.GetQuestionByID(question.Question{
		Id: i,
	})

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Q:%v", q2.Question))
	fmt.Println(fmt.Sprintf("A:%v", q2.Answer))
	return nil
}
