// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func TimeHtml(tags []any) *TimeTagHtml {
	node := &TimeTagHtml{
		tag: &tag{
			name:                "time",
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

type TimeTagHtml struct {
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
func (t *TimeTagHtml) CustomAttribute(attributes ...*Attribute) *TimeTagHtml {
	t.registerAttributes(attributes...)
	return t
}

/*
Children - Method for nesting tags into parent tag
*/
func (t *TimeTagHtml) Children(tags ...any) *TimeTagHtml {
	return t.supportedChildrenCheck(tags)
}

func (t *TimeTagHtml) supportedChildrenCheck(tags []any) *TimeTagHtml {
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

/*
Datetime -
*/
func (t *TimeTagHtml) Datetime(value string) *TimeTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("datetime", value)
	return t
}
