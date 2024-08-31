// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func MenuHtml(tags []any) *MenuTagHtml {
	node := &MenuTagHtml{
		tag: &tag{
			name:                "menu",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.FullTagType,
			namespace:           spec.HTML,
			children:            nil,
			textContent:         "",
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
	return node.supportedChildrenCheck(tags)
}

type MenuTagHtml struct {
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
func (m *MenuTagHtml) CustomAttribute(attributes ...*Attribute) *MenuTagHtml {
	m.registerAttributes(attributes...)
	return m
}

/*
Children - Method for nesting tags into parent tag
*/
func (m *MenuTagHtml) Children(tags ...any) *MenuTagHtml {
	return m.supportedChildrenCheck(tags)
}

func (m *MenuTagHtml) supportedChildrenCheck(tags []any) *MenuTagHtml {
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
