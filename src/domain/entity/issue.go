package entity

import (
	"errors"
	"time"
)

type Issue struct {
	Id        uint16
	Title     string
	Url       string
	RepoId    uint16
	LabelId   uint16
	StartDate time.Time
	DueDate   time.Time
	Progress  uint8
	Assignee  string
}

func NewIssue(id uint16, title string, url string, repoId uint16) *Issue {
	issue := new(Issue)
	issue.Id = id
	issue.Title = title
	issue.Url = url
	issue.RepoId = repoId
	return issue
}

func (issue *Issue) SetLabel(labelId uint16) *Issue {
	issue.LabelId = labelId
	return issue
}

func (issue *Issue) SetStartDate(date time.Time) *Issue {
	issue.StartDate = date
	return issue
}

// SetDueDate 期日は開始日よりも前に設定することはできない
func (issue *Issue) SetDueDate(date time.Time) (*Issue, error) {
	if issue.StartDate.Before(date) || issue.StartDate.Equal(date) {
		issue.DueDate = date
		return issue, nil
	} else {
		return nil, errors.New("cannot set DueDate before StartDate\n")
	}
}

// SetProgress 進捗は常に0~100の間の整数
func (issue *Issue) SetProgress(progress uint8) (*Issue, error) {
	if progress <= 100 {
		issue.Progress = progress
		return issue, nil
	} else {
		return nil, errors.New("Progress must be always in a range 0 to 100\n")
	}
}

func (issue *Issue) SetAssignee(assignee string) *Issue {
	issue.Assignee = assignee
	return issue
}
