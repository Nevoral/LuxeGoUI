// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func MaskSvg(tags []any) *MaskTagSvg {
	node := &MaskTagSvg{
		tag: &tag{
			name:                "mask",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.FullTagType,
			namespace:           spec.SVG,
			children:            nil,
			textContent:         "",
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
	return node.supportedChildrenCheck(tags)
}

type MaskTagSvg struct {
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
func (m *MaskTagSvg) CustomAttribute(attributes ...*Attribute) *MaskTagSvg {
	m.registerAttributes(attributes...)
	return m
}

/*
Children - Method for nesting tags into parent tag
*/
func (m *MaskTagSvg) Children(tags ...any) *MaskTagSvg {
	return m.supportedChildrenCheck(tags)
}

func (m *MaskTagSvg) supportedChildrenCheck(tags []any) *MaskTagSvg {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			m.registerChildren(TextContentSvg(val).getTag())
		case content:
			m.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				m.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return m
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Height -
*/
func (m *MaskTagSvg) Height(value string) *MaskTagSvg {
	if m.attributes == nil {
		m.attributes = []*Attribute{}
	}
	m.registerAttribute("height", value)
	return m
}

/*
MaskContentUnits -
*/
func (m *MaskTagSvg) MaskContentUnits(value string) *MaskTagSvg {
	if m.attributes == nil {
		m.attributes = []*Attribute{}
	}
	m.registerAttribute("maskContentUnits", value)
	return m
}

/*
MaskUnits -
*/
func (m *MaskTagSvg) MaskUnits(value string) *MaskTagSvg {
	if m.attributes == nil {
		m.attributes = []*Attribute{}
	}
	m.registerAttribute("maskUnits", value)
	return m
}

/*
Width -
*/
func (m *MaskTagSvg) Width(value string) *MaskTagSvg {
	if m.attributes == nil {
		m.attributes = []*Attribute{}
	}
	m.registerAttribute("width", value)
	return m
}

/*
X -
*/
func (m *MaskTagSvg) X(value string) *MaskTagSvg {
	if m.attributes == nil {
		m.attributes = []*Attribute{}
	}
	m.registerAttribute("x", value)
	return m
}

/*
Y -
*/
func (m *MaskTagSvg) Y(value string) *MaskTagSvg {
	if m.attributes == nil {
		m.attributes = []*Attribute{}
	}
	m.registerAttribute("y", value)
	return m
}
