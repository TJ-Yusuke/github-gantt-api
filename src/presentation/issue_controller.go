package presentation

import (
	"fmt"
	"github-gantt-api/src/domain/entity"
	"github-gantt-api/src/usecase"
	"sort"
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
	// 期日の最新順に並び替える
	sort.Slice(issues, func(i, j int) bool {
		return issues[i].DueDate.Before(issues[j].DueDate)
	})
	return issues, nil
}
