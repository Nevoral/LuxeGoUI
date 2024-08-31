// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func BdiHtml(tags []any) *BdiTagHtml {
	node := &BdiTagHtml{
		tag: &tag{
			name:                "bdi",
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

type BdiTagHtml struct {
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
func (b *BdiTagHtml) CustomAttribute(attributes ...*Attribute) *BdiTagHtml {
	b.registerAttributes(attributes...)
	return b
}

/*
Children - Method for nesting tags into parent tag
*/
func (b *BdiTagHtml) Children(tags ...any) *BdiTagHtml {
	return b.supportedChildrenCheck(tags)
}

func (b *BdiTagHtml) supportedChildrenCheck(tags []any) *BdiTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			b.registerChildren(TextContentSvg(val).getTag())
		case content:
			b.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				b.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return b
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */
