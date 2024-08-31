// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func StrongHtml(tags []any) *StrongTagHtml {
	node := &StrongTagHtml{
		tag: &tag{
			name:                "strong",
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

type StrongTagHtml struct {
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
func (s *StrongTagHtml) CustomAttribute(attributes ...*Attribute) *StrongTagHtml {
	s.registerAttributes(attributes...)
	return s
}

/*
Children - Method for nesting tags into parent tag
*/
func (s *StrongTagHtml) Children(tags ...any) *StrongTagHtml {
	return s.supportedChildrenCheck(tags)
}

func (s *StrongTagHtml) supportedChildrenCheck(tags []any) *StrongTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			s.registerChildren(TextContentSvg(val).getTag())
		case content:
			s.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				s.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return s
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */
