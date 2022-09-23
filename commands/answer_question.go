package commands

import (
	"errors"
	"fmt"
	"github.com/joycezemitchell/quiz_master/domain/question"
	"github.com/joycezemitchell/quiz_master/services"
	"github.com/joycezemitchell/quiz_master/util"
	"strconv"
	"strings"
)

func AnswerQuestion(arrCommandStr []string) error {

	// Check parameter
	if len(arrCommandStr) != 3 {
		return errors.New(fmt.Sprintf("%s required for 2 arguments", QUESTION))
	}

	i, _ := strconv.Atoi(arrCommandStr[1])
	answer := strings.Trim(arrCommandStr[2], "\"")

	q2, err := services.QuestionService.GetQuestionByID(question.Question{
		Id: i,
	})

	if err != nil {
		return err
	}

	checkAnswer(q2.Answer, answer)

	return nil
}

func checkAnswer(correctAnswer string, userAnswer string) error {

	ca := strings.ToLower(correctAnswer)
	ua := strings.ToLower(userAnswer)

	// Check if answer is int
	intCa, err := strconv.Atoi(ca)
	if err == nil {
		ca = util.Convert1to1000(intCa)
	}

	intUa, err := strconv.Atoi(ua)
	if err == nil {
		ua = util.Convert1to1000(intUa)
	}

	if ca != ua {
		fmt.Println("INCORRECT!")
		return errors.New("incorrect")
	}

	fmt.Println("CORRECT!")
	return nil
}
