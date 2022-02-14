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

type TextFragment struct {
	Italicized bool
	Bold       bool
	Text       string
}

type Paragraph []TextFragment

type Work struct {
	Title string
	Work  []Paragraph
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
		var para Paragraph
		s.Children().Each(func(j int, child *goquery.Selection) {
			var fragment TextFragment
			if child.Is("em") {
				fragment.Italicized = true
			} else if child.Is("b") {
				fragment.Bold = true
			}
			fragment.Text = strings.TrimSpace(child.Text())
			para = append(para, fragment)
		})
		work.Work = append(work.Work, para)
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

func fixUnbalancedQuotes(text Paragraph) Paragraph {
	unbalanced := false
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i].Text); j++ {
			if text[i].Text[j] == '"' {
				unbalanced = !unbalanced
			}
		}
	}
	if unbalanced {
		lastFragment := len(text) - 1
		text[lastFragment].Text = text[lastFragment].Text + "\""
	}
	return text
}

func fixUnicodeChars(text Paragraph) Paragraph {
	for i := 0; i < len(text); i++ {
		var textBuilder strings.Builder
		for _, char := range text[i].Text {
			if unicode.IsSymbol(char) {
				textBuilder.WriteString(fmt.Sprintf(`{\DejaSans %c}`, char))
			} else {
				textBuilder.WriteRune(char)
			}
		}
		text[i].Text = textBuilder.String()
	}
	return text
}
