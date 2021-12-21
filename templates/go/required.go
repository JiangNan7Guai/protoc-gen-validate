package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if m.{{ .Field.Name.UpperCamelCase }} == nil {
			err := {{ err . "必填项!" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
