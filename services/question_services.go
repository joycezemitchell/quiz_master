package services

import (
	"github.com/joycezemitchell/quiz_master/domain/question"
)

var (
	QuestionService questionServiceInterface = &questionService{}
)

type questionService struct{}

type questionServiceInterface interface {
	AddQuestion(request question.Question) error
	GetAllQuestions() ([]question.Question, error)
	GetQuestionByID(request question.Question) (question.Question, error)
	DeleleteQuestion(request question.Question) error
	UpdateQuestion(request question.Question) error
}

// AddQuestion ...
func (p *questionService) AddQuestion(request question.Question) error {

	daoQuestion := &question.Question{}

	// Add question
	daoQuestion.AddQuestion(request)

	return nil
}

// GetQuestionByID ...
func (p *questionService) GetQuestionByID(request question.Question) (question.Question, error) {

	daoQuestion := &question.Question{
		Id: request.Id,
	}

	// Get question by ID
	q2, err := daoQuestion.GetQuestionByID()

	if err != nil {
		return question.Question{}, err
	}

	return q2, nil
}

// GetAllQuestions ...
func (p *questionService) GetAllQuestions() ([]question.Question, error) {

	daoQuestion := &question.Question{}

	// Get all questions
	q2, err := daoQuestion.GetAllQuestions()

	if err != nil {
		return nil, err
	}

	return q2, nil
}

// DeleleteQuestion ...
func (p *questionService) DeleleteQuestion(request question.Question) error {
	daoQuestion := &question.Question{
		Id: request.Id,
	}

	// Delete question
	err := daoQuestion.DeleleteQuestion()

	if err != nil {
		return err
	}

	return nil
}

// UpdateQuestion ...
func (p *questionService) UpdateQuestion(request question.Question) error {

	daoQuestion := &question.Question{}

	// Update question
	daoQuestion.UpdateQuestion(request)

	return nil
}
