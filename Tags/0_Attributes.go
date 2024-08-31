package tags

import (
	"fmt"
	"strings"
)

type attributeType int

const (
	html5Compliant attributeType = iota
	custom
	unknown
)

type Attribute struct {
	Name          string
	Value         string
	Boolean       bool
	attributeType attributeType
}

func (a *Attribute) renderHtmlAttribute() string {
	if a.Boolean {
		return fmt.Sprintf(" %s", a.Name)
	} else {
		//if a.Value == "" {
		//	return fmt.Sprintf(" %s=\"%s\"", a.Name, strings.ReplaceAll(a.DefaultValue, "%", "%%"))
		//}
		return fmt.Sprintf(" %s=\"%s\"", a.Name, strings.ReplaceAll(a.Value, "%", "%%"))
	}
}

func (a *Attribute) renderLuxeGoAttribute() string {
	name := removeAndCamelCase(a.Name, ":", "-")
	name = capitalizeFirst(name)
	if a.Boolean {
		return fmt.Sprintf(".%s()", name)
	} else {
		return fmt.Sprintf(".%s(\"%s\")", name, strings.ReplaceAll(a.Value, "%", "%%"))
	}
}
