package repository

import (
	"github-gantt-api/src/domain/entity"
)

type IssueRepository interface {
	// 開始日と期日が登録されているものしかとってこない
	GetIssues(projectId uint16) ([]*entity.Issue, error)
	UpdateIssue(issue *entity.Issue) error
}
