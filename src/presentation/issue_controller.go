package presentation

import (
	"fmt"
	"github-gantt-api/src/domain/entity"
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

func (ic *IssueController) GetIssues(projectId uint16) ([]*entity.Issue, error) {
	issues, err := ic.useCase.GetIssues(projectId)
	if err != nil {
		return nil, fmt.Errorf("could not get Issues because of '%v'", err)
	}
	return issues, nil
}
