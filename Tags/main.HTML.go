// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func MainHtml(tags []any) *MainTagHtml {
	node := &MainTagHtml{
		tag: &tag{
			name:                "main",
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

type MainTagHtml struct {
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
func (m *MainTagHtml) CustomAttribute(attributes ...*Attribute) *MainTagHtml {
	m.registerAttributes(attributes...)
	return m
}

/*
Children - Method for nesting tags into parent tag
*/
func (m *MainTagHtml) Children(tags ...any) *MainTagHtml {
	return m.supportedChildrenCheck(tags)
}

func (m *MainTagHtml) supportedChildrenCheck(tags []any) *MainTagHtml {
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
