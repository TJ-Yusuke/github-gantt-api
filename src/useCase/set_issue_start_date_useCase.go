package useCase

import (
	"fmt"
	"github-gantt-api/src/domain/entity"
	"github-gantt-api/src/domain/repository"
	"time"
)

type SetIssueStartDateUseCase struct {
	repository repository.IssueRepository
}

func NewSetIssueStartDateUseCase(repository repository.IssueRepository) *SetIssueStartDateUseCase {
	su := new(SetIssueStartDateUseCase)
	su.repository = repository
	return su
}

func (su *SetIssueStartDateUseCase) SetStartDate(date time.Time, issue *entity.Issue) {
	issue.SetStartDate(date)
	err := su.repository.UpdateIssue(issue)
	if err != nil {
		fmt.Errorf("could not set startDate because of '%v'", err)
	}
	return
}
