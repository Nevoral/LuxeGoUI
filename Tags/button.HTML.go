// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func ButtonHtml(tags []any) *ButtonTagHtml {
	node := &ButtonTagHtml{
		tag: &tag{
			name:                "button",
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

type ButtonTagHtml struct {
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
func (b *ButtonTagHtml) CustomAttribute(attributes ...*Attribute) *ButtonTagHtml {
	b.registerAttributes(attributes...)
	return b
}

/*
Children - Method for nesting tags into parent tag
*/
func (b *ButtonTagHtml) Children(tags ...any) *ButtonTagHtml {
	return b.supportedChildrenCheck(tags)
}

func (b *ButtonTagHtml) supportedChildrenCheck(tags []any) *ButtonTagHtml {
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

/*
Autocomplete -
*/
func (b *ButtonTagHtml) Autocomplete(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("autocomplete", value)
	return b
}

/*
Disabled -
*/
func (b *ButtonTagHtml) Disabled() *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("disabled", "")
	return b
}

/*
Form -
*/
func (b *ButtonTagHtml) Form(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("form", value)
	return b
}

/*
Formaction -
*/
func (b *ButtonTagHtml) Formaction(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("formaction", value)
	return b
}

/*
Formenctype -
*/
func (b *ButtonTagHtml) Formenctype(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("formenctype", value)
	return b
}

/*
Formmethod -
*/
func (b *ButtonTagHtml) Formmethod(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("formmethod", value)
	return b
}

/*
Formnovalidate -
*/
func (b *ButtonTagHtml) Formnovalidate(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("formnovalidate", value)
	return b
}

/*
Formtarget -
*/
func (b *ButtonTagHtml) Formtarget(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("formtarget", value)
	return b
}

/*
Name -
*/
func (b *ButtonTagHtml) Name(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("name", value)
	return b
}

/*
Popovertarget -
*/
func (b *ButtonTagHtml) Popovertarget(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("popovertarget", value)
	return b
}

/*
Popovertargetaction -
*/
func (b *ButtonTagHtml) Popovertargetaction(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("popovertargetaction", value)
	return b
}

/*
Type - Specifies the type of an <input> element.
Specifies the type of an <input> element.
*/
func (b *ButtonTagHtml) Type(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("type", value)
	return b
}

/*
Value - Specifies the value of an <input> element.
Specifies the value of an <input> element.
*/
func (b *ButtonTagHtml) Value(value string) *ButtonTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("value", value)
	return b
}
