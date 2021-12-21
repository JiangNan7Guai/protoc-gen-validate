package goshared

const numTpl = `
	{{ if .Rules.GetIgnoreEmpty }}
		if {{ accessor . }} != 0 {
	{{ end }}

	{{ template "const" . }}
	{{ template "ltgt" . }}
	{{ template "in" . }}
	{{ template "required" . }}

	{{ if .Rules.GetIgnoreEmpty }}
		}
	{{ end }}

`
