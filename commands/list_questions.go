package commands

import (
	"fmt"
	"github.com/joycezemitchell/quiz_master/services"
)

func GetAllQuestions() error {

	questions, err := services.QuestionService.GetAllQuestions()

	if err != nil {
		return err
	}

	fmt.Println("No | Question | Answer")
	for _, q := range questions {
		fmt.Println(fmt.Sprintf("%v, %v, %v", q.Id, q.Question, q.Answer))
	}

	return nil
}
