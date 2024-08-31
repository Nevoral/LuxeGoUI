// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func TfootHtml(tags []any) *TfootTagHtml {
	node := &TfootTagHtml{
		tag: &tag{
			name:                "tfoot",
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

type TfootTagHtml struct {
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
func (t *TfootTagHtml) CustomAttribute(attributes ...*Attribute) *TfootTagHtml {
	t.registerAttributes(attributes...)
	return t
}

/*
Children - Method for nesting tags into parent tag
*/
func (t *TfootTagHtml) Children(tags ...any) *TfootTagHtml {
	return t.supportedChildrenCheck(tags)
}

func (t *TfootTagHtml) supportedChildrenCheck(tags []any) *TfootTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			t.registerChildren(TextContentSvg(val).getTag())
		case content:
			t.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				t.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return t
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */
