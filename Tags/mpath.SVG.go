// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func MpathSvg(tags []any) *MpathTagSvg {
	node := &MpathTagSvg{
		tag: &tag{
			name:                "mpath",
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

type MpathTagSvg struct {
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
func (m *MpathTagSvg) CustomAttribute(attributes ...*Attribute) *MpathTagSvg {
	m.registerAttributes(attributes...)
	return m
}

/*
Children - Method for nesting tags into parent tag
*/
func (m *MpathTagSvg) Children(tags ...any) *MpathTagSvg {
	return m.supportedChildrenCheck(tags)
}

func (m *MpathTagSvg) supportedChildrenCheck(tags []any) *MpathTagSvg {
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
Href -
*/
func (m *MpathTagSvg) Href(value string) *MpathTagSvg {
	if m.attributes == nil {
		m.attributes = []*Attribute{}
	}
	m.registerAttribute("href", value)
	return m
}
