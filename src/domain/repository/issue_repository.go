package repository

import (
	"github-gantt-api/src/domain/entity"
)

type IssueRepository interface {
	GetIssue(issueId uint16) (*entity.Issue, error)
	// GetIssuesList 開始日と期日が登録されているものしかとってこない
	GetIssuesList(projectNumber uint16, org string) ([]*entity.Issue, error)
	UpdateIssue(issue *entity.Issue) error
}
