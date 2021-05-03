package repository

import (
	"github-gantt-api/src/domain/entity"
)

type IssueRepository interface {
	GetIssues(projectId uint16) ([]*entity.Issue, error)
	UpdateIssue(issue *entity.Issue) error
}
