// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func H1Html(tags []any) *H1TagHtml {
	node := &H1TagHtml{
		tag: &tag{
			name:                "h1",
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

type H1TagHtml struct {
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
func (h *H1TagHtml) CustomAttribute(attributes ...*Attribute) *H1TagHtml {
	h.registerAttributes(attributes...)
	return h
}

/*
Children - Method for nesting tags into parent tag
*/
func (h *H1TagHtml) Children(tags ...any) *H1TagHtml {
	return h.supportedChildrenCheck(tags)
}

func (h *H1TagHtml) supportedChildrenCheck(tags []any) *H1TagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			h.registerChildren(TextContentSvg(val).getTag())
		case content:
			h.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				h.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return h
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */
