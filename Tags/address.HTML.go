// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func AddressHtml(tags []any) *AddressTagHtml {
	node := &AddressTagHtml{
		tag: &tag{
			name:                "address",
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

type AddressTagHtml struct {
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
func (a *AddressTagHtml) CustomAttribute(attributes ...*Attribute) *AddressTagHtml {
	a.registerAttributes(attributes...)
	return a
}

/*
Children - Method for nesting tags into parent tag
*/
func (a *AddressTagHtml) Children(tags ...any) *AddressTagHtml {
	return a.supportedChildrenCheck(tags)
}

func (a *AddressTagHtml) supportedChildrenCheck(tags []any) *AddressTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			a.registerChildren(TextContentSvg(val).getTag())
		case content:
			a.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				a.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return a
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */
