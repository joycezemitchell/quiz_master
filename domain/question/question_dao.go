package question

import (
	"context"
	"errors"
	"fmt"
	"github.com/joycezemitchell/quiz_master/datasource/mysql/dataq"
	"log"
)

const (
	queryAddQuestion     = "INSERT INTO questions(id, question, answer) VALUES(%v, '%v', '%v')"
	queryGetQuestionByID = "SELECT id, question, answer FROM questions WHERE id=?"
	queryGetAllQuestions = "SELECT id, question, answer FROM questions"
	queryDeleteQuestion  = "DELETE FROM questions WHERE id=%v"
	queryUpdateQuestion  = "UPDATE questions SET question='%v', answer='%v' WHERE id=%v"
)

// AddQuestion ...
func (question *Question) AddQuestion(input Question) error {

	ctx := context.Background()
	tx, err := dataq.Client.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	// Add question
	_, err = tx.ExecContext(ctx, fmt.Sprintf(
		queryAddQuestion,
		input.Id,
		input.Question,
		input.Answer,
	))
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

// GetQuestionByID ...
func (question *Question) GetQuestionByID() (Question, error) {

	stmt, err := dataq.Client.Prepare(queryGetQuestionByID)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(question.Id)
	if getErr := result.Scan(&question.Id, &question.Question, &question.Answer); getErr != nil {
		log.Println(getErr)
		return Question{}, errors.New("error when tying to find admin")
	}

	return Question{
		Id:       question.Id,
		Question: question.Question,
		Answer:   question.Answer,
	}, nil

}

// GetAllQuestions
func (question *Question) GetAllQuestions() ([]Question, error) {

	qs := []Question{}

	stmt, err := dataq.Client.Prepare(queryGetAllQuestions)
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	results, err := stmt.Query()

	if err != nil {
		log.Println(err)
	}

	for results.Next() {
		q2 := Question{}
		results.Scan(&q2.Id, &q2.Question, &q2.Answer)
		qs = append(qs, q2)
	}

	return qs, nil

}

// DeleleteQuestion ...
func (question *Question) DeleleteQuestion() error {

	ctx := context.Background()
	tx, err := dataq.Client.BeginTx(ctx, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	// Delete question
	_, err = tx.ExecContext(ctx, fmt.Sprintf(queryDeleteQuestion, question.Id))

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

// UpdateQuestion ...
func (question *Question) UpdateQuestion(input Question) error {

	ctx := context.Background()
	tx, err := dataq.Client.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// Update question
	_, err = tx.ExecContext(ctx, fmt.Sprintf(
		queryUpdateQuestion,
		input.Question,
		input.Answer,
		input.Id,
	))

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

