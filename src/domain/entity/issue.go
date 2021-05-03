package entity

import (
	"errors"
	"time"
)

type Issue struct {
	Id        uint16
	Title     string
	Url       string
	Label     Label
	StartDate time.Time
	DueDate   time.Time
	Progress  uint8
	Assignee  string
}

func NewIssue(id uint16, title string, url string) *Issue {
	issue := new(Issue)
	issue.Id = id
	issue.Title = title
	issue.Url = url
	return issue
}

func (issue *Issue) SetLabel(label Label) *Issue {
	issue.Label = label
	return issue
}

func (issue *Issue) SetStartDate(date time.Time) *Issue {
	issue.StartDate = date
	return issue
}

// 期日は開始日よりも前に設定することはできない
func (issue *Issue) SetDueDate(date time.Time) (*Issue, error) {
	if issue.StartDate.Before(date) || issue.StartDate.Equal(date) {
		issue.DueDate = date
		return issue, nil
	} else {
		return nil, errors.New("cannot set DueDate before StartDate\n")
	}
}

// 進捗は常に0~100の間の整数
func (issue *Issue) SetProgress(progress uint8) (*Issue, error) {
	if progress <= 100 {
		issue.Progress = progress
		return issue, nil
	} else {
		return nil, errors.New("Progress is always in a range 0 to 100\n")
	}
}

func (issue *Issue) SetAssignee(assignee string) *Issue {
	issue.Assignee = assignee
	return issue
}

type Label struct {
	Id          uint16
	Url         string
	Name        string
	Description string
	Color       string
}
