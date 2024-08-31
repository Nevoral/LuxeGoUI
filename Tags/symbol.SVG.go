// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func SymbolSvg(tags []any) *SymbolTagSvg {
	node := &SymbolTagSvg{
		tag: &tag{
			name:                "symbol",
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

type SymbolTagSvg struct {
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
func (s *SymbolTagSvg) CustomAttribute(attributes ...*Attribute) *SymbolTagSvg {
	s.registerAttributes(attributes...)
	return s
}

/*
Children - Method for nesting tags into parent tag
*/
func (s *SymbolTagSvg) Children(tags ...any) *SymbolTagSvg {
	return s.supportedChildrenCheck(tags)
}

func (s *SymbolTagSvg) supportedChildrenCheck(tags []any) *SymbolTagSvg {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			s.registerChildren(TextContentSvg(val).getTag())
		case content:
			s.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				s.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return s
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
PreserveAspectRatio -
*/
func (s *SymbolTagSvg) PreserveAspectRatio(value string) *SymbolTagSvg {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("preserveAspectRatio", value)
	return s
}

/*
RefX -
*/
func (s *SymbolTagSvg) RefX(value string) *SymbolTagSvg {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("refX", value)
	return s
}

/*
RefY -
*/
func (s *SymbolTagSvg) RefY(value string) *SymbolTagSvg {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("refY", value)
	return s
}

/*
ViewBox -
*/
func (s *SymbolTagSvg) ViewBox(value string) *SymbolTagSvg {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("viewBox", value)
	return s
}
