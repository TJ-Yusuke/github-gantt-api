package repository

import (
	"github-gantt-api/src/domain/entity"
	"time"
)

type IssueRepository interface {
	GetIssues() (*entity.Issue, error)
	SetStartDate(date time.Time, issue *entity.Issue) error
	SetDueDate(date time.Time, issue *entity.Issue) error
	SetProgress(progress uint8, issue *entity.Issue) error
}
