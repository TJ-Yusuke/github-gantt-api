package usecase

import (
	"fmt"
	"github-gantt-api/src/domain/entity"
	"github-gantt-api/src/domain/repository"
	"sort"
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

func (iu *IssueUseCase) GetIssue(issueId uint16) (*entity.Issue, error) {
	issue, err := iu.repository.GetIssue(issueId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch issue ID:%v, because of '%v'", issueId, err)
	}
	return issue, nil
}

func (iu *IssueUseCase) GetIssuesList(projectId uint16) ([]*entity.Issue, error) {
	issues, err := iu.repository.GetIssuesList(projectId)
	if err != nil {
		return nil, fmt.Errorf("could not get issues via github api of '%v'", err)
	}
	// 期日の最新順に並び替える
	sort.Slice(issues, func(i, j int) bool {
		return issues[i].DueDate.Before(issues[j].DueDate)
	})
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

func assertError(s string, err error) error {
	if err != nil {
		return fmt.Errorf("%v because of '%v'", s, err)
	}
	return nil
}
