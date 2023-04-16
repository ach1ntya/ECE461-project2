package main

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"
)

func code_review_metric(personal_token string, owner string, repo string) float64 {

	client := graphql.NewClient("https://api.github.com/graphql")
	req := graphql.NewRequest(fmt.Sprintf(`query {
      repository(owner: "%s", name: "%s") {
        pullRequests(last: 100) {
          nodes {
            reviews(first: 1) {
              totalCount
            }
          }
        }
      }
    }`, owner, repo))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", personal_token))

	ctx := context.Background()
	var respData struct {
		Repository struct {
			PullRequests struct {
				Nodes []struct {
					Reviews struct {
						TotalCount int `json:"totalCount"`
					} `json:"reviews"`
				} `json:"nodes"`
			} `json:"pullRequests"`
		} `json:"repository"`
	}
	if err := client.Run(ctx, req, &respData); err != nil {
		panic(err)
	}

	reviewedPullRequests := 0
	for _, node := range respData.Repository.PullRequests.Nodes {
		if node.Reviews.TotalCount > 0 {
			reviewedPullRequests++
		}
	}

	var score float64
	var totalPullRequests = len(respData.Repository.PullRequests.Nodes)
	if totalPullRequests > 0 {
		score = float64(reviewedPullRequests) / float64(totalPullRequests)
	} else {
		score = 0
	}

	return score
}
