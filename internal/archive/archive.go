package archive

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"unicode"

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
	for i := 0; i < len(work.Work); i++ {
		work.Work[i] = fixUnbalancedQuotes(work.Work[i])
		work.Work[i] = fixUnicodeChars(work.Work[i])
	}
	work.Title = strings.TrimSpace(workDoc.Find("h2.title").First().Text())
	return work, nil
}

func fixUnbalancedQuotes(text string) string {
	unbalanced := false
	for i := 0; i < len(text); i++ {
		if text[i] == '"' {
			unbalanced = !unbalanced
		}
	}
	if unbalanced {
		return text + "\""
	} else {
		return text
	}
}

func fixUnicodeChars(text string) string {
	var textBuilder strings.Builder
	for _, char := range text {
		if unicode.IsSymbol(char) {
			textBuilder.WriteString(fmt.Sprintf(`{\DejaSans %c}`, char))
		} else {
			textBuilder.WriteRune(char)
		}
	}
	return textBuilder.String()
}
