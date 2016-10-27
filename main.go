package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jmcvetta/napping"
	"github.com/libgit2/git2go"
)

func main() {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	urlStr := "http://codereview.infotecs.ru:8080/services/json/v1"
	//repoPath := flag.String("repo", dir, "path to the git repo")
	url := flag.String("url", urlStr, "path to the git repo")
	userName := flag.String("user", "user", "user name for ccollab")
	password := flag.String("passwd", "", "password for ccollab")
	//reviewType := flag.String("type", "new", "review type new|last|reviewId")
	//numCommits := flag.Int("num", 1, "number of commits, which will be added to review")

	ccollabCreateSession(*url, *userName, *password)

}

func getCommitMessage(repoPath string) string {
	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	headCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
		panic(err)
	}
	return headCommit.Message()
}

func ccollabCreateSession(url string, userName string, passwd string) napping.Session {
	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "myvalue")
	s.Header = h

	// var jsonStr = []byte(`
	// [
	//        {"command" : "Examples.checkLoggedIn"},
	//        {"command" : "SessionService.getLoginTicket",
	//                "args":{"login":"jsmith","password":"qwerty12345"}},
	//        {"command" : "Examples.checkLoggedIn"}
	// ]`)

	var jsonStr = []byte(`{"command" : "Examples.checkLoggedIn"}`)
	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		log.Panic(err)
	}
	log.Print(data)

	resp, err := s.Post(url, &data, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response Status:", resp.Status())
	fmt.Println("response Headers:", resp.HttpResponse().Header)
	fmt.Println("response Body:", resp.RawText())

	return s
}
