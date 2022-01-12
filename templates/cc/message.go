package cc

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}
	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
	{
		pgv::ValidationMsg inner_err;
		if ({{ hasAccessor .}} && !pgv::Validator<{{ ctype $f.Type }}>::CheckMessage({{ accessor . }}, &inner_err)) {
			{{ errCause . "inner_err" "下级参数有误" }}
		}
	}
	{{ end }}
`

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if (!{{ hasAccessor . }}) {
			{{ err . "value is required" }}
		}
	{{ end }}
`
