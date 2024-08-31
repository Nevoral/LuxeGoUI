// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func HtmlHtml(tags []any) *HtmlTagHtml {
	node := &HtmlTagHtml{
		tag: &tag{
			name:                "html",
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

type HtmlTagHtml struct {
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
func (h *HtmlTagHtml) CustomAttribute(attributes ...*Attribute) *HtmlTagHtml {
	h.registerAttributes(attributes...)
	return h
}

/*
Children - Method for nesting tags into parent tag
*/
func (h *HtmlTagHtml) Children(tags ...any) *HtmlTagHtml {
	return h.supportedChildrenCheck(tags)
}

func (h *HtmlTagHtml) supportedChildrenCheck(tags []any) *HtmlTagHtml {
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

/*
Xmlns -
*/
func (h *HtmlTagHtml) Xmlns(value string) *HtmlTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("xmlns", value)
	return h
}
