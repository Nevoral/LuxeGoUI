// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func QHtml(tags []any) *QTagHtml {
	node := &QTagHtml{
		tag: &tag{
			name:                "q",
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

type QTagHtml struct {
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
func (q *QTagHtml) CustomAttribute(attributes ...*Attribute) *QTagHtml {
	q.registerAttributes(attributes...)
	return q
}

/*
Children - Method for nesting tags into parent tag
*/
func (q *QTagHtml) Children(tags ...any) *QTagHtml {
	return q.supportedChildrenCheck(tags)
}

func (q *QTagHtml) supportedChildrenCheck(tags []any) *QTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			q.registerChildren(TextContentSvg(val).getTag())
		case content:
			q.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				q.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return q
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Cite -
*/
func (q *QTagHtml) Cite(value string) *QTagHtml {
	if q.attributes == nil {
		q.attributes = []*Attribute{}
	}
	q.registerAttribute("cite", value)
	return q
}
