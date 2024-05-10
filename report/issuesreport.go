package report

import (
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var Report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
