// Package tags do not edit, this file was autogenerated.
package tags

import spec "github.com/Nevoral/LuxeGoUI/Specification"

func WbrHtml() *WbrTagHtml {
	return &WbrTagHtml{
		tag: &tag{
			name:                "wbr",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.SelfClosingType,
			namespace:           spec.HTML,
			children:            nil,
			textContent:         "",
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

type WbrTagHtml struct {
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
func (w *WbrTagHtml) CustomAttribute(attributes ...*Attribute) *WbrTagHtml {
	w.registerAttributes(attributes...)
	return w
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */
