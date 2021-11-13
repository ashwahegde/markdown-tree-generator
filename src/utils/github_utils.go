package githubutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

// ghAPIBaseUrl := "https://api.github.com/repos/"
// "ashwahegde/my_coursera_certficates/git/trees/master?recursive=1"
type GHDetails struct {
	hostName string
	baseUrl  string
	authorId string
	repoName string
	treeSha  string
}

func (ghD GHDetails) generateUrl() string {
	tUrl := url.URL{
		Scheme: "https",
		Host:   ghD.hostName,
		Path: path.Join(
			ghD.baseUrl,
			ghD.authorId,
			ghD.repoName,
			"git/trees",
			ghD.treeSha,
		),
	}
	return tUrl.String()
}

type GHTree struct {
	Path  string `json:"path"`
	RType string `json:"type"`
	Sha   string `json:"sha"`
	Url   string `json:"url"`
}

type GHResponse struct {
	Sha  string   `json:"sha"`
	Tree []GHTree `json:"tree"`
}

func GetChildren(readyUrl string) []GHTree {
	// if treeSha == "" {
	// 	treeSha = "master"
	// }
	resp, err := http.Get(readyUrl)

	if err != nil {
		log.Fatal("Invalid repsonse from GitHub API.", err)
		os.Exit(1)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	respGH := GHResponse{}
	// var respGH = new(GHResponse)
	jsonErr := json.Unmarshal([]byte(body), &respGH)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return respGH.Tree
}

func dummyFun() string {
	return "hey"
}
