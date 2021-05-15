package presentation

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github-gantt-api/src/domain/repository"
	"github-gantt-api/src/usecase"
	"time"
)

type IssueController struct {
	useCase    usecase.IssueUseCase
	repository repository.IssueRepository
}

func NewIssueController(useCase usecase.IssueUseCase, repository repository.IssueRepository) *IssueController {
	ic := new(IssueController)
	ic.useCase = useCase
	ic.repository = repository
	return ic
}

func (ic *IssueController) GetIssues(projectId uint16) (response []byte, error error) {
	issues, err := ic.useCase.GetIssuesList(projectId)
	if err != nil {
		return nil, fmt.Errorf("could not get Issues because of '%v'", err)
	}
	res, err := json.Marshal(issues)
	if err != nil {
		return nil, fmt.Errorf("convert json error '%v'", err)
	}
	return res, nil
}
func (ic *IssueController) SetStartDate(date []byte, issueIdByte []byte) error {
	issueId := binary.BigEndian.Uint16(issueIdByte)
	issue, err := ic.repository.GetIssue(issueId)
	if err != nil {
		return fmt.Errorf("failed to fetch issue %v", err)
	}
	layout := "Mon Jan 2 15:04:05 MST 2006"
	startDate, err := time.Parse(layout, string(date))
	if err != nil {
		return fmt.Errorf("failed to convert string to Time.time '%v'", err)
	}
	err2 := ic.useCase.SetStartDate(startDate, issue)
	if err2 != nil {
		return fmt.Errorf("failed to setStartDate '%v'", err2)
	}
	return nil
}

func (ic *IssueController) SetDueDate(date []byte, issueIdByte []byte) error {
	issueId := binary.BigEndian.Uint16(issueIdByte)
	issue, err := ic.repository.GetIssue(issueId)
	if err != nil {
		return fmt.Errorf("failed to fetch issue %v", err)
	}
	layout := "Mon Jan 2 15:04:05 MST 2006"
	dueDate, err := time.Parse(layout, string(date))
	if err != nil {
		return fmt.Errorf("failed to convert string to Time.time '%v'", err)
	}
	err2 := ic.useCase.SetDueDate(dueDate, issue)
	if err2 != nil {
		return fmt.Errorf("failed to SetDueDate '%v'", err2)
	}
	return nil
}

func (ic *IssueController) SetProgress(progress uint8, issueIdByte []byte) error {
	issueId := binary.BigEndian.Uint16(issueIdByte)
	issue, err := ic.repository.GetIssue(issueId)
	if err != nil {
		return fmt.Errorf("failed to fetch issue %v", err)
	}
	err2 := ic.useCase.SetProgress(progress, issue)
	if err2 != nil {
		return fmt.Errorf("failed to SetProgress '%v'", err2)
	}
	return nil
}
