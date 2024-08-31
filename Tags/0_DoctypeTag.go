package tags

import (
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func DOCTYPEHtml() *DOCTYPETagHtml {
	return &DOCTYPETagHtml{
		tag: &tag{
			name:                "!DOCTYPE",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.DoctypeType,
			namespace:           spec.HTML,
			children:            nil,
			textContent:         "",
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

type DOCTYPETagHtml struct {
	*tag
}

func (d *DOCTYPETagHtml) Html() *DOCTYPETagHtml {
	d.namespace = spec.HTML
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("html", "")
	return d
}

func (d *DOCTYPETagHtml) Svg() *DOCTYPETagHtml {
	d.namespace = spec.SVG
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("svg", "")
	return d
}

func (d *DOCTYPETagHtml) Math() *DOCTYPETagHtml {
	d.namespace = spec.MATH
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("math", "")
	return d
}

func (d *DOCTYPETagHtml) PUBLIC(value string) *DOCTYPETagHtml {
	if d.namespace == spec.HTML {
		d.namespace = spec.XHTML
	}
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("PUBLIC", value)
	return d
}

func (d *DOCTYPETagHtml) SYSTEM(value string) *DOCTYPETagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("SYSTEM", value)
	return d
}
