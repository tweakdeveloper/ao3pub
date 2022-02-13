package doc

import (
	"strings"
	"text/template"

	"github.com/tweakdeveloper/ao3pub2/internal/archive"
)

const workTemplate = `\documentclass{article}
\usepackage{setspace} \doublespacing
\usepackage[margin=1in]{geometry}
\usepackage[pdftitle={ {{.Title}} }]{hyperref}
\begin{document}
{{range .Work}}\paragraph{}
{{.}}
{{end}}\end{document}
`

func GetTemplateFromWork(work archive.Work) (string, error) {
	workTemplate, err := template.New("work").Parse(workTemplate)
	if err != nil {
		return "", err
	}
	var workBuilder strings.Builder
	err = workTemplate.Execute(&workBuilder, work)
	if err != nil {
		return "", err
	}
	return workBuilder.String(), nil
}
