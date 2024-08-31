// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func SvgHtml(tags []any) *SvgTagHtml {
	node := &SvgTagHtml{
		tag: &tag{
			name:                "svg",
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

type SvgTagHtml struct {
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
func (s *SvgTagHtml) CustomAttribute(attributes ...*Attribute) *SvgTagHtml {
	s.registerAttributes(attributes...)
	return s
}

/*
Children - Method for nesting tags into parent tag
*/
func (s *SvgTagHtml) Children(tags ...any) *SvgTagHtml {
	return s.supportedChildrenCheck(tags)
}

func (s *SvgTagHtml) supportedChildrenCheck(tags []any) *SvgTagHtml {
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
BaseProfile -
*/
func (s *SvgTagHtml) BaseProfile(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("baseProfile", value)
	return s
}

/*
Fill -
*/
func (s *SvgTagHtml) Fill(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("fill", value)
	return s
}

/*
Height -
*/
func (s *SvgTagHtml) Height(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("height", value)
	return s
}

/*
PreserveAspectRatio -
*/
func (s *SvgTagHtml) PreserveAspectRatio(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("preserveAspectRatio", value)
	return s
}

/*
Version -
*/
func (s *SvgTagHtml) Version(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("version", value)
	return s
}

/*
ViewBox -
*/
func (s *SvgTagHtml) ViewBox(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("viewBox", value)
	return s
}

/*
Width -
*/
func (s *SvgTagHtml) Width(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("width", value)
	return s
}

/*
X -
*/
func (s *SvgTagHtml) X(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("x", value)
	return s
}

/*
XmlBase -
*/
func (s *SvgTagHtml) XmlBase(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("xml:base", value)
	return s
}

/*
XmlLang -
*/
func (s *SvgTagHtml) XmlLang(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("xml:lang", value)
	return s
}

/*
XmlSpace -
*/
func (s *SvgTagHtml) XmlSpace(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("xml:space", value)
	return s
}

/*
Xmlns -
*/
func (s *SvgTagHtml) Xmlns(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("xmlns", value)
	return s
}

/*
XmlnsXlink -
*/
func (s *SvgTagHtml) XmlnsXlink(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("xmlns:xlink", value)
	return s
}

/*
Y -
*/
func (s *SvgTagHtml) Y(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("y", value)
	return s
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (s *SvgTagHtml) AccessKey(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("accessKey", value)
	return s
}

/*
Aria -
*/
func (s *SvgTagHtml) Aria(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria", value)
	return s
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (s *SvgTagHtml) Autocapitalize(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("autocapitalize", value)
	return s
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (s *SvgTagHtml) Autofocus(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("autofocus", value)
	return s
}

/*
Class -
*/
func (s *SvgTagHtml) Class(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("class", value)
	return s
}

/*
Contenteditable -
*/
func (s *SvgTagHtml) Contenteditable(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("contenteditable", value)
	return s
}

/*
Data -
*/
func (s *SvgTagHtml) Data(name, value string) *SvgTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	s.registerAttribute(dataName, value)
	return s
}

/*
Dir -
*/
func (s *SvgTagHtml) Dir(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("dir", value)
	return s
}

/*
Draggable -
*/
func (s *SvgTagHtml) Draggable(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("draggable", value)
	return s
}

/*
EnterKeyHint -
*/
func (s *SvgTagHtml) EnterKeyHint(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("enterKeyHint", value)
	return s
}

/*
ExportParts -
*/
func (s *SvgTagHtml) ExportParts(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("exportParts", value)
	return s
}

/*
Hidden -
*/
func (s *SvgTagHtml) Hidden(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("hidden", value)
	return s
}

/*
Id -
*/
func (s *SvgTagHtml) Id(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("id", value)
	return s
}

/*
Inert -
*/
func (s *SvgTagHtml) Inert(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inert", value)
	return s
}

/*
InputMode -
*/
func (s *SvgTagHtml) InputMode(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inputMode", value)
	return s
}

/*
Is -
*/
func (s *SvgTagHtml) Is(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("is", value)
	return s
}

/*
ItemId -
*/
func (s *SvgTagHtml) ItemId(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemId", value)
	return s
}

/*
ItemProp -
*/
func (s *SvgTagHtml) ItemProp(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemProp", value)
	return s
}

/*
ItemRef -
*/
func (s *SvgTagHtml) ItemRef(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemRef", value)
	return s
}

/*
ItemScope -
*/
func (s *SvgTagHtml) ItemScope(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemScope", value)
	return s
}

/*
ItemType -
*/
func (s *SvgTagHtml) ItemType(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemType", value)
	return s
}

/*
Lang -
*/
func (s *SvgTagHtml) Lang(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("lang", value)
	return s
}

/*
Nonce -
*/
func (s *SvgTagHtml) Nonce(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("nonce", value)
	return s
}

/*
Part -
*/
func (s *SvgTagHtml) Part(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("part", value)
	return s
}

/*
Popover -
*/
func (s *SvgTagHtml) Popover() *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("popover", "")
	return s
}

/*
Role -
*/
func (s *SvgTagHtml) Role(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("role", value)
	return s
}

/*
Slot -
*/
func (s *SvgTagHtml) Slot(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("slot", value)
	return s
}

/*
Spellcheck -
*/
func (s *SvgTagHtml) Spellcheck(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("spellcheck", value)
	return s
}

/*
Style -
*/
func (s *SvgTagHtml) Style(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("style", value)
	return s
}

/*
Tabindex -
*/
func (s *SvgTagHtml) Tabindex(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("tabindex", value)
	return s
}

/*
Title -
*/
func (s *SvgTagHtml) Title(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("title", value)
	return s
}

/*
Translate -
*/
func (s *SvgTagHtml) Translate(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("translate", value)
	return s
}

/*
VirtualKeyBoardPolicy -
*/
func (s *SvgTagHtml) VirtualKeyBoardPolicy(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("virtualKeyBoardPolicy", value)
	return s
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (s *SvgTagHtml) AriaAtomic(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-atomic", value)
	return s
}

/*
AriaBusy -
*/
func (s *SvgTagHtml) AriaBusy(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-busy", value)
	return s
}

/*
AriaControls -
*/
func (s *SvgTagHtml) AriaControls(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-controls", value)
	return s
}

/*
AriaCurrent -
*/
func (s *SvgTagHtml) AriaCurrent(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-current", value)
	return s
}

/*
AriaDescribedby -
*/
func (s *SvgTagHtml) AriaDescribedby(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-describedby", value)
	return s
}

/*
AriaDescription -
*/
func (s *SvgTagHtml) AriaDescription(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-description", value)
	return s
}

/*
AriaDetails -
*/
func (s *SvgTagHtml) AriaDetails(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-details", value)
	return s
}

/*
AriaDisabled -
*/
func (s *SvgTagHtml) AriaDisabled(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-disabled", value)
	return s
}

/*
AriaDropeffect -
*/
func (s *SvgTagHtml) AriaDropeffect(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-dropeffect", value)
	return s
}

/*
AriaErrormessage -
*/
func (s *SvgTagHtml) AriaErrormessage(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-errormessage", value)
	return s
}

/*
AriaFlowto -
*/
func (s *SvgTagHtml) AriaFlowto(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-flowto", value)
	return s
}

/*
AriaGrabbed -
*/
func (s *SvgTagHtml) AriaGrabbed(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-grabbed", value)
	return s
}

/*
AriaHaspopup -
*/
func (s *SvgTagHtml) AriaHaspopup(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-haspopup", value)
	return s
}

/*
AriaHidden -
*/
func (s *SvgTagHtml) AriaHidden(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-hidden", value)
	return s
}

/*
AriaInvalid -
*/
func (s *SvgTagHtml) AriaInvalid(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-invalid", value)
	return s
}

/*
AriaKeyshortcuts -
*/
func (s *SvgTagHtml) AriaKeyshortcuts(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-keyshortcuts", value)
	return s
}

/*
AriaLabel -
*/
func (s *SvgTagHtml) AriaLabel(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-label", value)
	return s
}

/*
AriaLabelledby -
*/
func (s *SvgTagHtml) AriaLabelledby(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-labelledby", value)
	return s
}

/*
AriaLive -
*/
func (s *SvgTagHtml) AriaLive(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-live", value)
	return s
}

/*
AriaOwns -
*/
func (s *SvgTagHtml) AriaOwns(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-owns", value)
	return s
}

/*
AriaRelevant -
*/
func (s *SvgTagHtml) AriaRelevant(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-relevant", value)
	return s
}

/*
AriaRoledescription -
*/
func (s *SvgTagHtml) AriaRoledescription(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-roledescription", value)
	return s
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (s *SvgTagHtml) Onabort(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onabort", value)
	return s
}

/*
Onautocomplete -
*/
func (s *SvgTagHtml) Onautocomplete(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocomplete", value)
	return s
}

/*
Onautocompleteerror -
*/
func (s *SvgTagHtml) Onautocompleteerror(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocompleteerror", value)
	return s
}

/*
Onblur -
*/
func (s *SvgTagHtml) Onblur(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onblur", value)
	return s
}

/*
Oncancel -
*/
func (s *SvgTagHtml) Oncancel(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncancel", value)
	return s
}

/*
Oncanplay -
*/
func (s *SvgTagHtml) Oncanplay(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplay", value)
	return s
}

/*
Oncanplaythrough -
*/
func (s *SvgTagHtml) Oncanplaythrough(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplaythrough", value)
	return s
}

/*
Onchange -
*/
func (s *SvgTagHtml) Onchange(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onchange", value)
	return s
}

/*
Onclick -
*/
func (s *SvgTagHtml) Onclick(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclick", value)
	return s
}

/*
Onclose -
*/
func (s *SvgTagHtml) Onclose(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclose", value)
	return s
}

/*
Oncontextmenu -
*/
func (s *SvgTagHtml) Oncontextmenu(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncontextmenu", value)
	return s
}

/*
Oncuechange -
*/
func (s *SvgTagHtml) Oncuechange(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncuechange", value)
	return s
}

/*
Ondblclick -
*/
func (s *SvgTagHtml) Ondblclick(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondblclick", value)
	return s
}

/*
Ondrag -
*/
func (s *SvgTagHtml) Ondrag(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrag", value)
	return s
}

/*
Ondragend -
*/
func (s *SvgTagHtml) Ondragend(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragend", value)
	return s
}

/*
Ondragenter -
*/
func (s *SvgTagHtml) Ondragenter(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragenter", value)
	return s
}

/*
Ondragleave -
*/
func (s *SvgTagHtml) Ondragleave(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragleave", value)
	return s
}

/*
Ondragover -
*/
func (s *SvgTagHtml) Ondragover(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragover", value)
	return s
}

/*
Ondragstart -
*/
func (s *SvgTagHtml) Ondragstart(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragstart", value)
	return s
}

/*
Ondrop -
*/
func (s *SvgTagHtml) Ondrop(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrop", value)
	return s
}

/*
Ondurationchange -
*/
func (s *SvgTagHtml) Ondurationchange(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondurationchange", value)
	return s
}

/*
Onemptied -
*/
func (s *SvgTagHtml) Onemptied(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onemptied", value)
	return s
}

/*
Onended -
*/
func (s *SvgTagHtml) Onended(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onended", value)
	return s
}

/*
Onfocus -
*/
func (s *SvgTagHtml) Onfocus(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onfocus", value)
	return s
}

/*
Oninput -
*/
func (s *SvgTagHtml) Oninput(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninput", value)
	return s
}

/*
Oninvalid -
*/
func (s *SvgTagHtml) Oninvalid(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninvalid", value)
	return s
}

/*
Onkeydown -
*/
func (s *SvgTagHtml) Onkeydown(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeydown", value)
	return s
}

/*
Onkeypress -
*/
func (s *SvgTagHtml) Onkeypress(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeypress", value)
	return s
}

/*
Onkeyup -
*/
func (s *SvgTagHtml) Onkeyup(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeyup", value)
	return s
}

/*
Onloadeddata -
*/
func (s *SvgTagHtml) Onloadeddata(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadeddata", value)
	return s
}

/*
Onloadedmetadata -
*/
func (s *SvgTagHtml) Onloadedmetadata(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadedmetadata", value)
	return s
}

/*
Onloadstart -
*/
func (s *SvgTagHtml) Onloadstart(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadstart", value)
	return s
}

/*
Onmousedown -
*/
func (s *SvgTagHtml) Onmousedown(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousedown", value)
	return s
}

/*
Onmouseenter -
*/
func (s *SvgTagHtml) Onmouseenter(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseenter", value)
	return s
}

/*
Onmouseleave -
*/
func (s *SvgTagHtml) Onmouseleave(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseleave", value)
	return s
}

/*
Onmousemove -
*/
func (s *SvgTagHtml) Onmousemove(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousemove", value)
	return s
}

/*
Onmouseout -
*/
func (s *SvgTagHtml) Onmouseout(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseout", value)
	return s
}

/*
Onmouseover -
*/
func (s *SvgTagHtml) Onmouseover(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseover", value)
	return s
}

/*
Onmouseup -
*/
func (s *SvgTagHtml) Onmouseup(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseup", value)
	return s
}

/*
Onmousewheel -
*/
func (s *SvgTagHtml) Onmousewheel(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousewheel", value)
	return s
}

/*
Onpause -
*/
func (s *SvgTagHtml) Onpause(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpause", value)
	return s
}

/*
Onplay -
*/
func (s *SvgTagHtml) Onplay(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplay", value)
	return s
}

/*
Onplaying -
*/
func (s *SvgTagHtml) Onplaying(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplaying", value)
	return s
}

/*
Onprogress -
*/
func (s *SvgTagHtml) Onprogress(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onprogress", value)
	return s
}

/*
Onratechange -
*/
func (s *SvgTagHtml) Onratechange(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onratechange", value)
	return s
}

/*
Onreset -
*/
func (s *SvgTagHtml) Onreset(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onreset", value)
	return s
}

/*
Onscroll -
*/
func (s *SvgTagHtml) Onscroll(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onscroll", value)
	return s
}

/*
Onseeked -
*/
func (s *SvgTagHtml) Onseeked(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeked", value)
	return s
}

/*
Onseeking -
*/
func (s *SvgTagHtml) Onseeking(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeking", value)
	return s
}

/*
Onselect -
*/
func (s *SvgTagHtml) Onselect(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onselect", value)
	return s
}

/*
Onshow -
*/
func (s *SvgTagHtml) Onshow(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onshow", value)
	return s
}

/*
Onsort -
*/
func (s *SvgTagHtml) Onsort(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsort", value)
	return s
}

/*
Onstalled -
*/
func (s *SvgTagHtml) Onstalled(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstalled", value)
	return s
}

/*
Onsubmit -
*/
func (s *SvgTagHtml) Onsubmit(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsubmit", value)
	return s
}

/*
Onsuspend -
*/
func (s *SvgTagHtml) Onsuspend(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsuspend", value)
	return s
}

/*
Ontimeupdate -
*/
func (s *SvgTagHtml) Ontimeupdate(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontimeupdate", value)
	return s
}

/*
Ontoggle -
*/
func (s *SvgTagHtml) Ontoggle(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontoggle", value)
	return s
}

/*
Onvolumechange -
*/
func (s *SvgTagHtml) Onvolumechange(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onvolumechange", value)
	return s
}

/*
Onwaiting -
*/
func (s *SvgTagHtml) Onwaiting(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onwaiting", value)
	return s
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (s *SvgTagHtml) Onafterprint(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onafterprint", value)
	return s
}

/*
Onbeforeprint -
*/
func (s *SvgTagHtml) Onbeforeprint(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeprint", value)
	return s
}

/*
Onbeforeunload -
*/
func (s *SvgTagHtml) Onbeforeunload(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeunload", value)
	return s
}

/*
Onerror -
*/
func (s *SvgTagHtml) Onerror(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onerror", value)
	return s
}

/*
Onhashchange -
*/
func (s *SvgTagHtml) Onhashchange(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onhashchange", value)
	return s
}

/*
Onload -
*/
func (s *SvgTagHtml) Onload(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onload", value)
	return s
}

/*
Onmessage -
*/
func (s *SvgTagHtml) Onmessage(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmessage", value)
	return s
}

/*
Onoffline -
*/
func (s *SvgTagHtml) Onoffline(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onoffline", value)
	return s
}

/*
Ononline -
*/
func (s *SvgTagHtml) Ononline(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ononline", value)
	return s
}

/*
Onpagehide -
*/
func (s *SvgTagHtml) Onpagehide(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpagehide", value)
	return s
}

/*
Onpageshow -
*/
func (s *SvgTagHtml) Onpageshow(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpageshow", value)
	return s
}

/*
Onpopstate -
*/
func (s *SvgTagHtml) Onpopstate(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpopstate", value)
	return s
}

/*
Onresize -
*/
func (s *SvgTagHtml) Onresize(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onresize", value)
	return s
}

/*
Onstorage -
*/
func (s *SvgTagHtml) Onstorage(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstorage", value)
	return s
}

/*
Onunload -
*/
func (s *SvgTagHtml) Onunload(value string) *SvgTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onunload", value)
	return s
}
