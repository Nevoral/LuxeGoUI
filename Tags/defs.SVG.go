// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func DefsSvg(tags []any) *DefsTagSvg {
	node := &DefsTagSvg{
		tag: &tag{
			name:                "defs",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.FullTagType,
			namespace:           spec.SVG,
			children:            nil,
			textContent:         "",
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
	return node.supportedChildrenCheck(tags)
}

type DefsTagSvg struct {
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
func (d *DefsTagSvg) CustomAttribute(attributes ...*Attribute) *DefsTagSvg {
	d.registerAttributes(attributes...)
	return d
}

/*
Children - Method for nesting tags into parent tag
*/
func (d *DefsTagSvg) Children(tags ...any) *DefsTagSvg {
	return d.supportedChildrenCheck(tags)
}

func (d *DefsTagSvg) supportedChildrenCheck(tags []any) *DefsTagSvg {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			d.registerChildren(TextContentSvg(val).getTag())
		case content:
			d.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				d.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return d
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */
