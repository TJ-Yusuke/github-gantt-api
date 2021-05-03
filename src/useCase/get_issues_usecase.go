package useCase

import (
	"fmt"
	"github-gantt-api/src/domain/entity"
	"github-gantt-api/src/domain/repository"
)

type GetIssuesUseCase struct {
	repository repository.IssueRepository
}

func NewGetIssuesUseCase(repository repository.IssueRepository) *GetIssuesUseCase {
	giu := new(GetIssuesUseCase)
	giu.repository = repository
	return giu
}

func (giu *GetIssuesUseCase) GetIssues(projectId uint16) ([]*entity.Issue, error) {
	issues, err := giu.repository.GetIssues(projectId)
	if err != nil {
		return nil, fmt.Errorf("could not get issues via github api of '%v'", err)
	}
	return issues, nil
}
