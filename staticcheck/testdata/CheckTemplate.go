package pkg

import (
	th "html/template"
	tt "text/template"
)

const tmpl1 = `{{.Name}} {{.LastName}`
const tmpl2 = `{{fn}}`

func fn() {
	_, _ = tt.New("").Parse(tmpl1) // MATCH /template/
	_, _ = tt.New("").Parse(tmpl2)
	t1 := tt.New("")
	_, _ = t1.Parse(tmpl1)
	_, _ = th.New("").Parse(tmpl1) // MATCH /template/
	_, _ = th.New("").Parse(tmpl2)
	t2 := th.New("")
	_, _ = t2.Parse(tmpl1)
	_, _ = tt.New("").Delims("[[", "]]").Parse("{{abc-}}")
}
