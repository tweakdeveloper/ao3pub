package archive

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseUrl = "https://archiveofourown.org"

func GetWorkText(work string) ([]string, error) {
	workRequest, err := http.Get(fmt.Sprintf("%s/works/%s", baseUrl, work))
	if err != nil {
		return nil, err
	}
	defer workRequest.Body.Close()
	if workRequest.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error retrieving work: status code %d",
			workRequest.StatusCode)
	}
	workDoc, err := goquery.NewDocumentFromReader(workRequest.Body)
	if err != nil {
		return nil, err
	}
	var paragraphs []string
	workDoc.Find("div.userstuff p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, strings.TrimSpace(s.Text()))
	})
	if len(paragraphs) == 0 {
		return nil, errors.New("work not found in page")
	}
	return paragraphs, nil
}
