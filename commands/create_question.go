package commands

import (
	"errors"
	"fmt"
	q "github.com/joycezemitchell/quiz_master/domain/question"
	"github.com/joycezemitchell/quiz_master/services"
	"log"
	"strconv"
	"strings"
)

func CreateQuestion(arrCommandStr []string) error {

	// Check parameter
	if len(arrCommandStr) != 4 {
		return errors.New(fmt.Sprintf("%s required for 3 arguments", CREATE_QUESTIONS))
	}

	id := arrCommandStr[1]
	question := strings.Trim(arrCommandStr[2], "\"")
	answer := strings.Trim(arrCommandStr[3], "\"")

	i, _ := strconv.Atoi(id)

	request := &q.Question{
		Id:       i,
		Question: question,
		Answer:   answer,
	}

	q2,_ := request.GetQuestionByID()

	// Check if question already exist
	if q2 != (q.Question{}) {
		return errors.New(fmt.Sprintf("Question no %v already existed", i))
	}

	err := services.QuestionService.AddQuestion(*request)

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(fmt.Sprintf("Question no %v was created", id))
	fmt.Println(fmt.Sprintf("Q:%v", question))
	fmt.Println(fmt.Sprintf("A:%v", answer))

	return nil
}


