// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func SpanHtml(tags []any) *SpanTagHtml {
	node := &SpanTagHtml{
		tag: &tag{
			name:                "span",
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

type SpanTagHtml struct {
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
func (s *SpanTagHtml) CustomAttribute(attributes ...*Attribute) *SpanTagHtml {
	s.registerAttributes(attributes...)
	return s
}

/*
Children - Method for nesting tags into parent tag
*/
func (s *SpanTagHtml) Children(tags ...any) *SpanTagHtml {
	return s.supportedChildrenCheck(tags)
}

func (s *SpanTagHtml) supportedChildrenCheck(tags []any) *SpanTagHtml {
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
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (s *SpanTagHtml) AccessKey(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("accessKey", value)
	return s
}

/*
Aria -
*/
func (s *SpanTagHtml) Aria(value string) *SpanTagHtml {
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
func (s *SpanTagHtml) Autocapitalize(value string) *SpanTagHtml {
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
func (s *SpanTagHtml) Autofocus(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("autofocus", value)
	return s
}

/*
Class -
*/
func (s *SpanTagHtml) Class(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("class", value)
	return s
}

/*
Contenteditable -
*/
func (s *SpanTagHtml) Contenteditable(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("contenteditable", value)
	return s
}

/*
Data -
*/
func (s *SpanTagHtml) Data(name, value string) *SpanTagHtml {
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
func (s *SpanTagHtml) Dir(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("dir", value)
	return s
}

/*
Draggable -
*/
func (s *SpanTagHtml) Draggable(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("draggable", value)
	return s
}

/*
EnterKeyHint -
*/
func (s *SpanTagHtml) EnterKeyHint(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("enterKeyHint", value)
	return s
}

/*
ExportParts -
*/
func (s *SpanTagHtml) ExportParts(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("exportParts", value)
	return s
}

/*
Hidden -
*/
func (s *SpanTagHtml) Hidden(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("hidden", value)
	return s
}

/*
Id -
*/
func (s *SpanTagHtml) Id(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("id", value)
	return s
}

/*
Inert -
*/
func (s *SpanTagHtml) Inert(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inert", value)
	return s
}

/*
InputMode -
*/
func (s *SpanTagHtml) InputMode(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inputMode", value)
	return s
}

/*
Is -
*/
func (s *SpanTagHtml) Is(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("is", value)
	return s
}

/*
ItemId -
*/
func (s *SpanTagHtml) ItemId(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemId", value)
	return s
}

/*
ItemProp -
*/
func (s *SpanTagHtml) ItemProp(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemProp", value)
	return s
}

/*
ItemRef -
*/
func (s *SpanTagHtml) ItemRef(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemRef", value)
	return s
}

/*
ItemScope -
*/
func (s *SpanTagHtml) ItemScope(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemScope", value)
	return s
}

/*
ItemType -
*/
func (s *SpanTagHtml) ItemType(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemType", value)
	return s
}

/*
Lang -
*/
func (s *SpanTagHtml) Lang(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("lang", value)
	return s
}

/*
Nonce -
*/
func (s *SpanTagHtml) Nonce(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("nonce", value)
	return s
}

/*
Part -
*/
func (s *SpanTagHtml) Part(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("part", value)
	return s
}

/*
Popover -
*/
func (s *SpanTagHtml) Popover() *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("popover", "")
	return s
}

/*
Role -
*/
func (s *SpanTagHtml) Role(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("role", value)
	return s
}

/*
Slot -
*/
func (s *SpanTagHtml) Slot(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("slot", value)
	return s
}

/*
Spellcheck -
*/
func (s *SpanTagHtml) Spellcheck(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("spellcheck", value)
	return s
}

/*
Style -
*/
func (s *SpanTagHtml) Style(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("style", value)
	return s
}

/*
Tabindex -
*/
func (s *SpanTagHtml) Tabindex(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("tabindex", value)
	return s
}

/*
Title -
*/
func (s *SpanTagHtml) Title(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("title", value)
	return s
}

/*
Translate -
*/
func (s *SpanTagHtml) Translate(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("translate", value)
	return s
}

/*
VirtualKeyBoardPolicy -
*/
func (s *SpanTagHtml) VirtualKeyBoardPolicy(value string) *SpanTagHtml {
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
func (s *SpanTagHtml) AriaAtomic(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-atomic", value)
	return s
}

/*
AriaBusy -
*/
func (s *SpanTagHtml) AriaBusy(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-busy", value)
	return s
}

/*
AriaControls -
*/
func (s *SpanTagHtml) AriaControls(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-controls", value)
	return s
}

/*
AriaCurrent -
*/
func (s *SpanTagHtml) AriaCurrent(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-current", value)
	return s
}

/*
AriaDescribedby -
*/
func (s *SpanTagHtml) AriaDescribedby(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-describedby", value)
	return s
}

/*
AriaDescription -
*/
func (s *SpanTagHtml) AriaDescription(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-description", value)
	return s
}

/*
AriaDetails -
*/
func (s *SpanTagHtml) AriaDetails(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-details", value)
	return s
}

/*
AriaDisabled -
*/
func (s *SpanTagHtml) AriaDisabled(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-disabled", value)
	return s
}

/*
AriaDropeffect -
*/
func (s *SpanTagHtml) AriaDropeffect(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-dropeffect", value)
	return s
}

/*
AriaErrormessage -
*/
func (s *SpanTagHtml) AriaErrormessage(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-errormessage", value)
	return s
}

/*
AriaFlowto -
*/
func (s *SpanTagHtml) AriaFlowto(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-flowto", value)
	return s
}

/*
AriaGrabbed -
*/
func (s *SpanTagHtml) AriaGrabbed(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-grabbed", value)
	return s
}

/*
AriaHaspopup -
*/
func (s *SpanTagHtml) AriaHaspopup(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-haspopup", value)
	return s
}

/*
AriaHidden -
*/
func (s *SpanTagHtml) AriaHidden(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-hidden", value)
	return s
}

/*
AriaInvalid -
*/
func (s *SpanTagHtml) AriaInvalid(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-invalid", value)
	return s
}

/*
AriaKeyshortcuts -
*/
func (s *SpanTagHtml) AriaKeyshortcuts(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-keyshortcuts", value)
	return s
}

/*
AriaLabel -
*/
func (s *SpanTagHtml) AriaLabel(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-label", value)
	return s
}

/*
AriaLabelledby -
*/
func (s *SpanTagHtml) AriaLabelledby(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-labelledby", value)
	return s
}

/*
AriaLive -
*/
func (s *SpanTagHtml) AriaLive(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-live", value)
	return s
}

/*
AriaOwns -
*/
func (s *SpanTagHtml) AriaOwns(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-owns", value)
	return s
}

/*
AriaRelevant -
*/
func (s *SpanTagHtml) AriaRelevant(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-relevant", value)
	return s
}

/*
AriaRoledescription -
*/
func (s *SpanTagHtml) AriaRoledescription(value string) *SpanTagHtml {
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
func (s *SpanTagHtml) Onabort(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onabort", value)
	return s
}

/*
Onautocomplete -
*/
func (s *SpanTagHtml) Onautocomplete(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocomplete", value)
	return s
}

/*
Onautocompleteerror -
*/
func (s *SpanTagHtml) Onautocompleteerror(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocompleteerror", value)
	return s
}

/*
Onblur -
*/
func (s *SpanTagHtml) Onblur(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onblur", value)
	return s
}

/*
Oncancel -
*/
func (s *SpanTagHtml) Oncancel(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncancel", value)
	return s
}

/*
Oncanplay -
*/
func (s *SpanTagHtml) Oncanplay(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplay", value)
	return s
}

/*
Oncanplaythrough -
*/
func (s *SpanTagHtml) Oncanplaythrough(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplaythrough", value)
	return s
}

/*
Onchange -
*/
func (s *SpanTagHtml) Onchange(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onchange", value)
	return s
}

/*
Onclick -
*/
func (s *SpanTagHtml) Onclick(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclick", value)
	return s
}

/*
Onclose -
*/
func (s *SpanTagHtml) Onclose(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclose", value)
	return s
}

/*
Oncontextmenu -
*/
func (s *SpanTagHtml) Oncontextmenu(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncontextmenu", value)
	return s
}

/*
Oncuechange -
*/
func (s *SpanTagHtml) Oncuechange(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncuechange", value)
	return s
}

/*
Ondblclick -
*/
func (s *SpanTagHtml) Ondblclick(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondblclick", value)
	return s
}

/*
Ondrag -
*/
func (s *SpanTagHtml) Ondrag(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrag", value)
	return s
}

/*
Ondragend -
*/
func (s *SpanTagHtml) Ondragend(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragend", value)
	return s
}

/*
Ondragenter -
*/
func (s *SpanTagHtml) Ondragenter(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragenter", value)
	return s
}

/*
Ondragleave -
*/
func (s *SpanTagHtml) Ondragleave(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragleave", value)
	return s
}

/*
Ondragover -
*/
func (s *SpanTagHtml) Ondragover(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragover", value)
	return s
}

/*
Ondragstart -
*/
func (s *SpanTagHtml) Ondragstart(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragstart", value)
	return s
}

/*
Ondrop -
*/
func (s *SpanTagHtml) Ondrop(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrop", value)
	return s
}

/*
Ondurationchange -
*/
func (s *SpanTagHtml) Ondurationchange(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondurationchange", value)
	return s
}

/*
Onemptied -
*/
func (s *SpanTagHtml) Onemptied(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onemptied", value)
	return s
}

/*
Onended -
*/
func (s *SpanTagHtml) Onended(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onended", value)
	return s
}

/*
Onfocus -
*/
func (s *SpanTagHtml) Onfocus(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onfocus", value)
	return s
}

/*
Oninput -
*/
func (s *SpanTagHtml) Oninput(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninput", value)
	return s
}

/*
Oninvalid -
*/
func (s *SpanTagHtml) Oninvalid(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninvalid", value)
	return s
}

/*
Onkeydown -
*/
func (s *SpanTagHtml) Onkeydown(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeydown", value)
	return s
}

/*
Onkeypress -
*/
func (s *SpanTagHtml) Onkeypress(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeypress", value)
	return s
}

/*
Onkeyup -
*/
func (s *SpanTagHtml) Onkeyup(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeyup", value)
	return s
}

/*
Onloadeddata -
*/
func (s *SpanTagHtml) Onloadeddata(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadeddata", value)
	return s
}

/*
Onloadedmetadata -
*/
func (s *SpanTagHtml) Onloadedmetadata(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadedmetadata", value)
	return s
}

/*
Onloadstart -
*/
func (s *SpanTagHtml) Onloadstart(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadstart", value)
	return s
}

/*
Onmousedown -
*/
func (s *SpanTagHtml) Onmousedown(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousedown", value)
	return s
}

/*
Onmouseenter -
*/
func (s *SpanTagHtml) Onmouseenter(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseenter", value)
	return s
}

/*
Onmouseleave -
*/
func (s *SpanTagHtml) Onmouseleave(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseleave", value)
	return s
}

/*
Onmousemove -
*/
func (s *SpanTagHtml) Onmousemove(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousemove", value)
	return s
}

/*
Onmouseout -
*/
func (s *SpanTagHtml) Onmouseout(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseout", value)
	return s
}

/*
Onmouseover -
*/
func (s *SpanTagHtml) Onmouseover(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseover", value)
	return s
}

/*
Onmouseup -
*/
func (s *SpanTagHtml) Onmouseup(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseup", value)
	return s
}

/*
Onmousewheel -
*/
func (s *SpanTagHtml) Onmousewheel(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousewheel", value)
	return s
}

/*
Onpause -
*/
func (s *SpanTagHtml) Onpause(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpause", value)
	return s
}

/*
Onplay -
*/
func (s *SpanTagHtml) Onplay(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplay", value)
	return s
}

/*
Onplaying -
*/
func (s *SpanTagHtml) Onplaying(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplaying", value)
	return s
}

/*
Onprogress -
*/
func (s *SpanTagHtml) Onprogress(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onprogress", value)
	return s
}

/*
Onratechange -
*/
func (s *SpanTagHtml) Onratechange(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onratechange", value)
	return s
}

/*
Onreset -
*/
func (s *SpanTagHtml) Onreset(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onreset", value)
	return s
}

/*
Onscroll -
*/
func (s *SpanTagHtml) Onscroll(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onscroll", value)
	return s
}

/*
Onseeked -
*/
func (s *SpanTagHtml) Onseeked(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeked", value)
	return s
}

/*
Onseeking -
*/
func (s *SpanTagHtml) Onseeking(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeking", value)
	return s
}

/*
Onselect -
*/
func (s *SpanTagHtml) Onselect(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onselect", value)
	return s
}

/*
Onshow -
*/
func (s *SpanTagHtml) Onshow(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onshow", value)
	return s
}

/*
Onsort -
*/
func (s *SpanTagHtml) Onsort(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsort", value)
	return s
}

/*
Onstalled -
*/
func (s *SpanTagHtml) Onstalled(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstalled", value)
	return s
}

/*
Onsubmit -
*/
func (s *SpanTagHtml) Onsubmit(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsubmit", value)
	return s
}

/*
Onsuspend -
*/
func (s *SpanTagHtml) Onsuspend(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsuspend", value)
	return s
}

/*
Ontimeupdate -
*/
func (s *SpanTagHtml) Ontimeupdate(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontimeupdate", value)
	return s
}

/*
Ontoggle -
*/
func (s *SpanTagHtml) Ontoggle(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontoggle", value)
	return s
}

/*
Onvolumechange -
*/
func (s *SpanTagHtml) Onvolumechange(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onvolumechange", value)
	return s
}

/*
Onwaiting -
*/
func (s *SpanTagHtml) Onwaiting(value string) *SpanTagHtml {
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
func (s *SpanTagHtml) Onafterprint(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onafterprint", value)
	return s
}

/*
Onbeforeprint -
*/
func (s *SpanTagHtml) Onbeforeprint(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeprint", value)
	return s
}

/*
Onbeforeunload -
*/
func (s *SpanTagHtml) Onbeforeunload(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeunload", value)
	return s
}

/*
Onerror -
*/
func (s *SpanTagHtml) Onerror(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onerror", value)
	return s
}

/*
Onhashchange -
*/
func (s *SpanTagHtml) Onhashchange(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onhashchange", value)
	return s
}

/*
Onload -
*/
func (s *SpanTagHtml) Onload(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onload", value)
	return s
}

/*
Onmessage -
*/
func (s *SpanTagHtml) Onmessage(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmessage", value)
	return s
}

/*
Onoffline -
*/
func (s *SpanTagHtml) Onoffline(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onoffline", value)
	return s
}

/*
Ononline -
*/
func (s *SpanTagHtml) Ononline(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ononline", value)
	return s
}

/*
Onpagehide -
*/
func (s *SpanTagHtml) Onpagehide(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpagehide", value)
	return s
}

/*
Onpageshow -
*/
func (s *SpanTagHtml) Onpageshow(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpageshow", value)
	return s
}

/*
Onpopstate -
*/
func (s *SpanTagHtml) Onpopstate(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpopstate", value)
	return s
}

/*
Onresize -
*/
func (s *SpanTagHtml) Onresize(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onresize", value)
	return s
}

/*
Onstorage -
*/
func (s *SpanTagHtml) Onstorage(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstorage", value)
	return s
}

/*
Onunload -
*/
func (s *SpanTagHtml) Onunload(value string) *SpanTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onunload", value)
	return s
}
