package repository_impl

import (
	"context"
	"fmt"
	"github-gantt-api/src/domain/entity"
	"github.com/google/go-github/v35/github"
	"strconv"
	"strings"
)

type IssueRepository struct {
	client *github.Client
}

func NewIssueRepository(client *github.Client) *IssueRepository {
	issueRepository := new(IssueRepository)
	issueRepository.client = client
	return issueRepository
}

func (ir *IssueRepository) GetIssuesList(ctx context.Context, projectNumber uint16, org string) ([]*entity.Issue, error) {
	projects, _, err := ir.client.Organizations.ListProjects(ctx, org, &github.ProjectListOptions{State: "all"})
	if err != nil {
		return nil, fmt.Errorf("could not get ListProjects '%v'", err)
	}
	var issuesList []*entity.Issue
	// Organizationのprojectから対応するprojectを取り出す
	for _, project := range projects {
		if *project.Number == int(projectNumber) {
			columns, _, err := ir.client.Projects.ListProjectColumns(ctx, *project.ID, &github.ListOptions{})
			if err != nil {
				return nil, fmt.Errorf("could not get ListProjectColums '%v'", err)
			}
			// 取り出したprojectのcolumnをルームさせてcardを取り出す
			for _, column := range columns {
				cards, _, err := ir.client.Projects.ListProjectCards(ctx, *column.ID, &github.ProjectCardListOptions{})
				if err != nil {
					return nil, fmt.Errorf("could not get ListProjectCards '%v'", err)
				}
				// 取り出したcardをループさせてissueを取得する
				for _, card := range cards {
					if card.ContentURL != nil {
						values := strings.Split(*card.ContentURL, "/")
						issueNumberString := values[len(values)-1]
						issueNumber, _ := strconv.Atoi(issueNumberString)
						repo := values[len(values)-3]
						owner := values[len(values)-4]
						issueStruct, _, err := ir.client.Issues.Get(ctx, owner, repo, issueNumber)
						if err != nil {
							return nil, fmt.Errorf("could not get Issue '%v'", err)
						}
						// 取得したissueをドメインに合わせてインスタンスを作成する
						issue := entity.NewIssue(uint16(*issueStruct.ID), *issueStruct.Title, *issueStruct.HTMLURL, uint16(*issueStruct.Repository.ID))
						issuesList = append(issuesList, issue)
					}
				}
			}
		}
	}
	fmt.Println("success getting IssuesList")
	return issuesList, nil
}
