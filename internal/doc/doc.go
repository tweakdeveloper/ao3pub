package doc

import (
	"strings"
	"text/template"
)

const workTemplate = `\documentclass{article}
\usepackage{setspace} \doublespacing
\usepackage[margin=1in]{geometry}
\begin{document}
{{range .}}\paragraph{}
{{.}}
{{end}}\end{document}
`

func GetTemplateFromWork(work []string) (string, error) {
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
