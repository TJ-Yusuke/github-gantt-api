package presentation

import (
	"encoding/json"
	"fmt"
	"github-gantt-api/src/usecase"
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
	issues, err := ic.useCase.GetIssues(projectId)
	if err != nil {
		return nil, fmt.Errorf("could not get Issues because of '%v'", err)
	}
	res, err := json.Marshal(issues)
	if err != nil {
		return nil, fmt.Errorf("convert json error '%v'", err)
	}
	return res, nil
}
