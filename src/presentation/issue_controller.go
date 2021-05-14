package presentation

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github-gantt-api/src/usecase"
	"time"
)

type IssueController struct {
	useCase usecase.IssueUseCase
}

func NewIssueController(useCase usecase.IssueUseCase) *IssueController {
	ic := new(IssueController)
	ic.useCase = useCase
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
func (ic *IssueController) SetStartDate(date []byte, oldIssueIdByte []byte) error {
	oldIssueId := binary.BigEndian.Uint16(oldIssueIdByte)
	oldIssue, err := ic.useCase.GetIssue(oldIssueId)
	if err != nil {
		return fmt.Errorf("failed to fetch issue %v", err)
	}
	layout := "Mon Jan 2 15:04:05 MST 2006"
	startDate, err := time.Parse(layout, string(date))
	if err != nil {
		return fmt.Errorf("failed to convert string to Time.time '%v'", err)
	}
	err2 := ic.useCase.SetStartDate(startDate, oldIssue)
	if err2 != nil {
		return fmt.Errorf("failed to setStartDate '%v'", err2)
	}
	return nil
}
