package doc

import (
	"strings"
	"text/template"

	"github.com/tweakdeveloper/ao3pub/internal/archive"
)

const workTemplate = `\documentclass{article}
\usepackage[english]{babel}
\usepackage{csquotes}
\MakeOuterQuote{"}
\usepackage{setspace} \doublespacing
\usepackage[margin=1in]{geometry}
\usepackage[pdftitle={ {{.Title}} }]{hyperref}
\usepackage{fontspec}
\newfontfamily\DejaSans{DejaVu Sans}
\usepackage{xeCJK}
\IfFontExistsTF{Noto Serif JP}{\setCJKmainfont{Noto Serif JP}}{\setCJKmainfont{Noto Serif CJK JP}}
\setlength{\parindent}{0.25in}
\begin{document}
{{range .Work}}{{range .}}{{if .Italicized}}\textit{ {{- .Text -}} }{{else if .Bold}}\textbf{ {{- .Text -}} }{{else}} {{.Text}} {{end}}{{end}}

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
