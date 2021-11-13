package githubutils

import (
	"fmt"
	"testing"
)

func TestDummyFun(t *testing.T) {
	fmt.Println("started")
	ghD := GHDetails{
		hostName: "api.github.com",
		baseUrl:  "repos",
		authorId: "ashwahegde",
		repoName: "my_coursera_certficates",
		treeSha:  "master",
	}

	fmt.Println("hey")
	if dummyFun() != "hey" {
		t.Error("invalid answer")
	}

	// aTree := GetChildren(ghD.generateUrl())
	// for _, aNode := range aTree {
	// 	fmt.Println(aNode)
	// }
	// Breadth First Traversal
	q := []string{}
	q = append(q, ghD.generateUrl())
	for len(q) > 0 {
		aTree := GetChildren(q[0])
		q = q[1:]
		for _, aNode := range aTree {
			fmt.Println(aNode)
			if aNode.RType == "tree" {
				q = append(q, aNode.Url)
			}
		}
	}
}
