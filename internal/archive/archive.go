package archive

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseUrl = "https://archiveofourown.org"

type Work struct {
	Title string
	Work  []string
}

func GetWork(workToFetch string) (Work, error) {
	var work Work
	workRequest, err := http.Get(fmt.Sprintf("%s/works/%s", baseUrl, workToFetch))
	if err != nil {
		return work, err
	}
	defer workRequest.Body.Close()
	if workRequest.StatusCode != http.StatusOK {
		return work, fmt.Errorf("error retrieving work: status code %d",
			workRequest.StatusCode)
	}
	workDoc, err := goquery.NewDocumentFromReader(workRequest.Body)
	if err != nil {
		return work, err
	}
	workDoc.Find("div.userstuff p").Each(func(i int, s *goquery.Selection) {
		work.Work = append(work.Work, strings.TrimSpace(s.Text()))
	})
	if len(work.Work) == 0 {
		return work, errors.New("work not found in page")
	}
	work.Title = strings.TrimSpace(workDoc.Find("h2.title").First().Text())
	return work, nil
}
