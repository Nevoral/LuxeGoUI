// Package tags do not edit, this file was autogenerated.
package tags

import spec "github.com/Nevoral/LuxeGoUI/Specification"

func TspanSvg() *TspanTagSvg {
	return &TspanTagSvg{
		tag: &tag{
			name:                "tspan",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.SelfClosingType,
			namespace:           spec.SVG,
			children:            nil,
			textContent:         "",
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

type TspanTagSvg struct {
	*tag
}

/*
************************************************************************************************************************
*-------------------------------------------------- Extension Method --------------------------------------------------*
************************************************************************************************************************
 */

/*
CustomAttribute - This is method for adding custom attribute, that are not in HTML5 specification.
*/
func (t *TspanTagSvg) CustomAttribute(attributes ...*Attribute) *TspanTagSvg {
	t.registerAttributes(attributes...)
	return t
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Dx -
*/
func (t *TspanTagSvg) Dx(value string) *TspanTagSvg {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("dx", value)
	return t
}

/*
Dy -
*/
func (t *TspanTagSvg) Dy(value string) *TspanTagSvg {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("dy", value)
	return t
}

/*
LengthAdjust -
*/
func (t *TspanTagSvg) LengthAdjust(value string) *TspanTagSvg {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("lengthAdjust", value)
	return t
}

/*
Rotate -
*/
func (t *TspanTagSvg) Rotate(value string) *TspanTagSvg {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("rotate", value)
	return t
}

/*
TextLength -
*/
func (t *TspanTagSvg) TextLength(value string) *TspanTagSvg {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("textLength", value)
	return t
}

/*
X -
*/
func (t *TspanTagSvg) X(value string) *TspanTagSvg {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("x", value)
	return t
}

/*
Y -
*/
func (t *TspanTagSvg) Y(value string) *TspanTagSvg {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("y", value)
	return t
}
