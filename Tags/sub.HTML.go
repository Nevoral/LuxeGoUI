// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func SubHtml(tags []any) *SubTagHtml {
	node := &SubTagHtml{
		tag: &tag{
			name:                "sub",
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

type SubTagHtml struct {
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
func (s *SubTagHtml) CustomAttribute(attributes ...*Attribute) *SubTagHtml {
	s.registerAttributes(attributes...)
	return s
}

/*
Children - Method for nesting tags into parent tag
*/
func (s *SubTagHtml) Children(tags ...any) *SubTagHtml {
	return s.supportedChildrenCheck(tags)
}

func (s *SubTagHtml) supportedChildrenCheck(tags []any) *SubTagHtml {
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
func (s *SubTagHtml) AccessKey(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("accessKey", value)
	return s
}

/*
Aria -
*/
func (s *SubTagHtml) Aria(value string) *SubTagHtml {
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
func (s *SubTagHtml) Autocapitalize(value string) *SubTagHtml {
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
func (s *SubTagHtml) Autofocus(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("autofocus", value)
	return s
}

/*
Class -
*/
func (s *SubTagHtml) Class(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("class", value)
	return s
}

/*
Contenteditable -
*/
func (s *SubTagHtml) Contenteditable(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("contenteditable", value)
	return s
}

/*
Data -
*/
func (s *SubTagHtml) Data(name, value string) *SubTagHtml {
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
func (s *SubTagHtml) Dir(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("dir", value)
	return s
}

/*
Draggable -
*/
func (s *SubTagHtml) Draggable(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("draggable", value)
	return s
}

/*
EnterKeyHint -
*/
func (s *SubTagHtml) EnterKeyHint(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("enterKeyHint", value)
	return s
}

/*
ExportParts -
*/
func (s *SubTagHtml) ExportParts(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("exportParts", value)
	return s
}

/*
Hidden -
*/
func (s *SubTagHtml) Hidden(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("hidden", value)
	return s
}

/*
Id -
*/
func (s *SubTagHtml) Id(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("id", value)
	return s
}

/*
Inert -
*/
func (s *SubTagHtml) Inert(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inert", value)
	return s
}

/*
InputMode -
*/
func (s *SubTagHtml) InputMode(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inputMode", value)
	return s
}

/*
Is -
*/
func (s *SubTagHtml) Is(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("is", value)
	return s
}

/*
ItemId -
*/
func (s *SubTagHtml) ItemId(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemId", value)
	return s
}

/*
ItemProp -
*/
func (s *SubTagHtml) ItemProp(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemProp", value)
	return s
}

/*
ItemRef -
*/
func (s *SubTagHtml) ItemRef(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemRef", value)
	return s
}

/*
ItemScope -
*/
func (s *SubTagHtml) ItemScope(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemScope", value)
	return s
}

/*
ItemType -
*/
func (s *SubTagHtml) ItemType(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemType", value)
	return s
}

/*
Lang -
*/
func (s *SubTagHtml) Lang(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("lang", value)
	return s
}

/*
Nonce -
*/
func (s *SubTagHtml) Nonce(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("nonce", value)
	return s
}

/*
Part -
*/
func (s *SubTagHtml) Part(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("part", value)
	return s
}

/*
Popover -
*/
func (s *SubTagHtml) Popover() *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("popover", "")
	return s
}

/*
Role -
*/
func (s *SubTagHtml) Role(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("role", value)
	return s
}

/*
Slot -
*/
func (s *SubTagHtml) Slot(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("slot", value)
	return s
}

/*
Spellcheck -
*/
func (s *SubTagHtml) Spellcheck(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("spellcheck", value)
	return s
}

/*
Style -
*/
func (s *SubTagHtml) Style(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("style", value)
	return s
}

/*
Tabindex -
*/
func (s *SubTagHtml) Tabindex(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("tabindex", value)
	return s
}

/*
Title -
*/
func (s *SubTagHtml) Title(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("title", value)
	return s
}

/*
Translate -
*/
func (s *SubTagHtml) Translate(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("translate", value)
	return s
}

/*
VirtualKeyBoardPolicy -
*/
func (s *SubTagHtml) VirtualKeyBoardPolicy(value string) *SubTagHtml {
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
func (s *SubTagHtml) AriaAtomic(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-atomic", value)
	return s
}

/*
AriaBusy -
*/
func (s *SubTagHtml) AriaBusy(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-busy", value)
	return s
}

/*
AriaControls -
*/
func (s *SubTagHtml) AriaControls(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-controls", value)
	return s
}

/*
AriaCurrent -
*/
func (s *SubTagHtml) AriaCurrent(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-current", value)
	return s
}

/*
AriaDescribedby -
*/
func (s *SubTagHtml) AriaDescribedby(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-describedby", value)
	return s
}

/*
AriaDescription -
*/
func (s *SubTagHtml) AriaDescription(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-description", value)
	return s
}

/*
AriaDetails -
*/
func (s *SubTagHtml) AriaDetails(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-details", value)
	return s
}

/*
AriaDisabled -
*/
func (s *SubTagHtml) AriaDisabled(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-disabled", value)
	return s
}

/*
AriaDropeffect -
*/
func (s *SubTagHtml) AriaDropeffect(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-dropeffect", value)
	return s
}

/*
AriaErrormessage -
*/
func (s *SubTagHtml) AriaErrormessage(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-errormessage", value)
	return s
}

/*
AriaFlowto -
*/
func (s *SubTagHtml) AriaFlowto(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-flowto", value)
	return s
}

/*
AriaGrabbed -
*/
func (s *SubTagHtml) AriaGrabbed(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-grabbed", value)
	return s
}

/*
AriaHaspopup -
*/
func (s *SubTagHtml) AriaHaspopup(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-haspopup", value)
	return s
}

/*
AriaHidden -
*/
func (s *SubTagHtml) AriaHidden(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-hidden", value)
	return s
}

/*
AriaInvalid -
*/
func (s *SubTagHtml) AriaInvalid(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-invalid", value)
	return s
}

/*
AriaKeyshortcuts -
*/
func (s *SubTagHtml) AriaKeyshortcuts(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-keyshortcuts", value)
	return s
}

/*
AriaLabel -
*/
func (s *SubTagHtml) AriaLabel(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-label", value)
	return s
}

/*
AriaLabelledby -
*/
func (s *SubTagHtml) AriaLabelledby(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-labelledby", value)
	return s
}

/*
AriaLive -
*/
func (s *SubTagHtml) AriaLive(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-live", value)
	return s
}

/*
AriaOwns -
*/
func (s *SubTagHtml) AriaOwns(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-owns", value)
	return s
}

/*
AriaRelevant -
*/
func (s *SubTagHtml) AriaRelevant(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-relevant", value)
	return s
}

/*
AriaRoledescription -
*/
func (s *SubTagHtml) AriaRoledescription(value string) *SubTagHtml {
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
func (s *SubTagHtml) Onabort(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onabort", value)
	return s
}

/*
Onautocomplete -
*/
func (s *SubTagHtml) Onautocomplete(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocomplete", value)
	return s
}

/*
Onautocompleteerror -
*/
func (s *SubTagHtml) Onautocompleteerror(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocompleteerror", value)
	return s
}

/*
Onblur -
*/
func (s *SubTagHtml) Onblur(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onblur", value)
	return s
}

/*
Oncancel -
*/
func (s *SubTagHtml) Oncancel(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncancel", value)
	return s
}

/*
Oncanplay -
*/
func (s *SubTagHtml) Oncanplay(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplay", value)
	return s
}

/*
Oncanplaythrough -
*/
func (s *SubTagHtml) Oncanplaythrough(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplaythrough", value)
	return s
}

/*
Onchange -
*/
func (s *SubTagHtml) Onchange(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onchange", value)
	return s
}

/*
Onclick -
*/
func (s *SubTagHtml) Onclick(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclick", value)
	return s
}

/*
Onclose -
*/
func (s *SubTagHtml) Onclose(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclose", value)
	return s
}

/*
Oncontextmenu -
*/
func (s *SubTagHtml) Oncontextmenu(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncontextmenu", value)
	return s
}

/*
Oncuechange -
*/
func (s *SubTagHtml) Oncuechange(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncuechange", value)
	return s
}

/*
Ondblclick -
*/
func (s *SubTagHtml) Ondblclick(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondblclick", value)
	return s
}

/*
Ondrag -
*/
func (s *SubTagHtml) Ondrag(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrag", value)
	return s
}

/*
Ondragend -
*/
func (s *SubTagHtml) Ondragend(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragend", value)
	return s
}

/*
Ondragenter -
*/
func (s *SubTagHtml) Ondragenter(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragenter", value)
	return s
}

/*
Ondragleave -
*/
func (s *SubTagHtml) Ondragleave(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragleave", value)
	return s
}

/*
Ondragover -
*/
func (s *SubTagHtml) Ondragover(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragover", value)
	return s
}

/*
Ondragstart -
*/
func (s *SubTagHtml) Ondragstart(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragstart", value)
	return s
}

/*
Ondrop -
*/
func (s *SubTagHtml) Ondrop(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrop", value)
	return s
}

/*
Ondurationchange -
*/
func (s *SubTagHtml) Ondurationchange(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondurationchange", value)
	return s
}

/*
Onemptied -
*/
func (s *SubTagHtml) Onemptied(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onemptied", value)
	return s
}

/*
Onended -
*/
func (s *SubTagHtml) Onended(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onended", value)
	return s
}

/*
Onfocus -
*/
func (s *SubTagHtml) Onfocus(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onfocus", value)
	return s
}

/*
Oninput -
*/
func (s *SubTagHtml) Oninput(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninput", value)
	return s
}

/*
Oninvalid -
*/
func (s *SubTagHtml) Oninvalid(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninvalid", value)
	return s
}

/*
Onkeydown -
*/
func (s *SubTagHtml) Onkeydown(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeydown", value)
	return s
}

/*
Onkeypress -
*/
func (s *SubTagHtml) Onkeypress(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeypress", value)
	return s
}

/*
Onkeyup -
*/
func (s *SubTagHtml) Onkeyup(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeyup", value)
	return s
}

/*
Onloadeddata -
*/
func (s *SubTagHtml) Onloadeddata(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadeddata", value)
	return s
}

/*
Onloadedmetadata -
*/
func (s *SubTagHtml) Onloadedmetadata(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadedmetadata", value)
	return s
}

/*
Onloadstart -
*/
func (s *SubTagHtml) Onloadstart(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadstart", value)
	return s
}

/*
Onmousedown -
*/
func (s *SubTagHtml) Onmousedown(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousedown", value)
	return s
}

/*
Onmouseenter -
*/
func (s *SubTagHtml) Onmouseenter(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseenter", value)
	return s
}

/*
Onmouseleave -
*/
func (s *SubTagHtml) Onmouseleave(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseleave", value)
	return s
}

/*
Onmousemove -
*/
func (s *SubTagHtml) Onmousemove(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousemove", value)
	return s
}

/*
Onmouseout -
*/
func (s *SubTagHtml) Onmouseout(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseout", value)
	return s
}

/*
Onmouseover -
*/
func (s *SubTagHtml) Onmouseover(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseover", value)
	return s
}

/*
Onmouseup -
*/
func (s *SubTagHtml) Onmouseup(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseup", value)
	return s
}

/*
Onmousewheel -
*/
func (s *SubTagHtml) Onmousewheel(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousewheel", value)
	return s
}

/*
Onpause -
*/
func (s *SubTagHtml) Onpause(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpause", value)
	return s
}

/*
Onplay -
*/
func (s *SubTagHtml) Onplay(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplay", value)
	return s
}

/*
Onplaying -
*/
func (s *SubTagHtml) Onplaying(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplaying", value)
	return s
}

/*
Onprogress -
*/
func (s *SubTagHtml) Onprogress(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onprogress", value)
	return s
}

/*
Onratechange -
*/
func (s *SubTagHtml) Onratechange(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onratechange", value)
	return s
}

/*
Onreset -
*/
func (s *SubTagHtml) Onreset(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onreset", value)
	return s
}

/*
Onscroll -
*/
func (s *SubTagHtml) Onscroll(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onscroll", value)
	return s
}

/*
Onseeked -
*/
func (s *SubTagHtml) Onseeked(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeked", value)
	return s
}

/*
Onseeking -
*/
func (s *SubTagHtml) Onseeking(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeking", value)
	return s
}

/*
Onselect -
*/
func (s *SubTagHtml) Onselect(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onselect", value)
	return s
}

/*
Onshow -
*/
func (s *SubTagHtml) Onshow(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onshow", value)
	return s
}

/*
Onsort -
*/
func (s *SubTagHtml) Onsort(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsort", value)
	return s
}

/*
Onstalled -
*/
func (s *SubTagHtml) Onstalled(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstalled", value)
	return s
}

/*
Onsubmit -
*/
func (s *SubTagHtml) Onsubmit(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsubmit", value)
	return s
}

/*
Onsuspend -
*/
func (s *SubTagHtml) Onsuspend(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsuspend", value)
	return s
}

/*
Ontimeupdate -
*/
func (s *SubTagHtml) Ontimeupdate(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontimeupdate", value)
	return s
}

/*
Ontoggle -
*/
func (s *SubTagHtml) Ontoggle(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontoggle", value)
	return s
}

/*
Onvolumechange -
*/
func (s *SubTagHtml) Onvolumechange(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onvolumechange", value)
	return s
}

/*
Onwaiting -
*/
func (s *SubTagHtml) Onwaiting(value string) *SubTagHtml {
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
func (s *SubTagHtml) Onafterprint(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onafterprint", value)
	return s
}

/*
Onbeforeprint -
*/
func (s *SubTagHtml) Onbeforeprint(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeprint", value)
	return s
}

/*
Onbeforeunload -
*/
func (s *SubTagHtml) Onbeforeunload(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeunload", value)
	return s
}

/*
Onerror -
*/
func (s *SubTagHtml) Onerror(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onerror", value)
	return s
}

/*
Onhashchange -
*/
func (s *SubTagHtml) Onhashchange(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onhashchange", value)
	return s
}

/*
Onload -
*/
func (s *SubTagHtml) Onload(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onload", value)
	return s
}

/*
Onmessage -
*/
func (s *SubTagHtml) Onmessage(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmessage", value)
	return s
}

/*
Onoffline -
*/
func (s *SubTagHtml) Onoffline(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onoffline", value)
	return s
}

/*
Ononline -
*/
func (s *SubTagHtml) Ononline(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ononline", value)
	return s
}

/*
Onpagehide -
*/
func (s *SubTagHtml) Onpagehide(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpagehide", value)
	return s
}

/*
Onpageshow -
*/
func (s *SubTagHtml) Onpageshow(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpageshow", value)
	return s
}

/*
Onpopstate -
*/
func (s *SubTagHtml) Onpopstate(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpopstate", value)
	return s
}

/*
Onresize -
*/
func (s *SubTagHtml) Onresize(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onresize", value)
	return s
}

/*
Onstorage -
*/
func (s *SubTagHtml) Onstorage(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstorage", value)
	return s
}

/*
Onunload -
*/
func (s *SubTagHtml) Onunload(value string) *SubTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onunload", value)
	return s
}
