package usecase

import (
	"fmt"
	"github-gantt-api/src/domain/entity"
	"github-gantt-api/src/domain/repository"
	"time"
)

type IssueUseCase struct {
	repository repository.IssueRepository
}

func NewIssueUseCase(repository repository.IssueRepository) *IssueUseCase {
	iu := new(IssueUseCase)
	iu.repository = repository
	return iu
}

func (iu *IssueUseCase) GetIssues(projectId uint16) ([]*entity.Issue, error) {
	issues, err := iu.repository.GetIssues(projectId)
	if err != nil {
		return nil, fmt.Errorf("could not get issues via github api of '%v'", err)
	}
	return issues, nil
}

func (iu *IssueUseCase) SetStartDate(date time.Time, oldIssue *entity.Issue) {
	newIssue := oldIssue.SetStartDate(date)
	err := iu.repository.UpdateIssue(newIssue)
	assertError("could not set startDate", err)
	return
}

func (iu *IssueUseCase) SetDueDate(date time.Time, oldIssue *entity.Issue) {
	newIssue, err := oldIssue.SetDueDate(date)
	assertError("could not set dueDate", err)
	updateErr := iu.repository.UpdateIssue(newIssue)
	assertError("could not set dueDate", updateErr)
	return
}

func (iu *IssueUseCase) SetLabel(label uint16, oldIssue *entity.Issue) {
	newIssue := oldIssue.SetLabel(label)
	err := iu.repository.UpdateIssue(newIssue)
	if err != nil {
		assertError("could not set Label", err)
	}
	return
}

func (iu *IssueUseCase) SetProgress(progress uint8, oldIssue *entity.Issue) {
	newIssue, err := oldIssue.SetProgress(progress)
	assertError("could not set progress", err)
	updateErr := iu.repository.UpdateIssue(newIssue)
	assertError("could not set progress", updateErr)
	return
}

func (iu *IssueUseCase) SetAssignee(assignee string, oldIssue *entity.Issue) {
	newIssue := oldIssue.SetAssignee(assignee)
	err := iu.repository.UpdateIssue(newIssue)
	assertError("could not set assignee", err)
	return
}

func assertError(s string, err error) {
	if err != nil {
		fmt.Errorf("%v because of '%v'", s, err)
	}
}
