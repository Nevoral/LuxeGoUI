// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func BaseHtml() *BaseTagHtml {
	return &BaseTagHtml{
		tag: &tag{
			name:                "base",
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

type BaseTagHtml struct {
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
func (b *BaseTagHtml) CustomAttribute(attributes ...*Attribute) *BaseTagHtml {
	b.registerAttributes(attributes...)
	return b
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Href -
*/
func (b *BaseTagHtml) Href(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("href", value)
	return b
}

/*
Target -
*/
func (b *BaseTagHtml) Target(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("target", value)
	return b
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (b *BaseTagHtml) AccessKey(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("accessKey", value)
	return b
}

/*
Aria -
*/
func (b *BaseTagHtml) Aria(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria", value)
	return b
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (b *BaseTagHtml) Autocapitalize(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("autocapitalize", value)
	return b
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (b *BaseTagHtml) Autofocus(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("autofocus", value)
	return b
}

/*
Class -
*/
func (b *BaseTagHtml) Class(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("class", value)
	return b
}

/*
Contenteditable -
*/
func (b *BaseTagHtml) Contenteditable(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("contenteditable", value)
	return b
}

/*
Data -
*/
func (b *BaseTagHtml) Data(name, value string) *BaseTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	b.registerAttribute(dataName, value)
	return b
}

/*
Dir -
*/
func (b *BaseTagHtml) Dir(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("dir", value)
	return b
}

/*
Draggable -
*/
func (b *BaseTagHtml) Draggable(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("draggable", value)
	return b
}

/*
EnterKeyHint -
*/
func (b *BaseTagHtml) EnterKeyHint(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("enterKeyHint", value)
	return b
}

/*
ExportParts -
*/
func (b *BaseTagHtml) ExportParts(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("exportParts", value)
	return b
}

/*
Hidden -
*/
func (b *BaseTagHtml) Hidden(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("hidden", value)
	return b
}

/*
Id -
*/
func (b *BaseTagHtml) Id(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("id", value)
	return b
}

/*
Inert -
*/
func (b *BaseTagHtml) Inert(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("inert", value)
	return b
}

/*
InputMode -
*/
func (b *BaseTagHtml) InputMode(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("inputMode", value)
	return b
}

/*
Is -
*/
func (b *BaseTagHtml) Is(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("is", value)
	return b
}

/*
ItemId -
*/
func (b *BaseTagHtml) ItemId(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemId", value)
	return b
}

/*
ItemProp -
*/
func (b *BaseTagHtml) ItemProp(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemProp", value)
	return b
}

/*
ItemRef -
*/
func (b *BaseTagHtml) ItemRef(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemRef", value)
	return b
}

/*
ItemScope -
*/
func (b *BaseTagHtml) ItemScope(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemScope", value)
	return b
}

/*
ItemType -
*/
func (b *BaseTagHtml) ItemType(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemType", value)
	return b
}

/*
Lang -
*/
func (b *BaseTagHtml) Lang(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("lang", value)
	return b
}

/*
Nonce -
*/
func (b *BaseTagHtml) Nonce(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("nonce", value)
	return b
}

/*
Part -
*/
func (b *BaseTagHtml) Part(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("part", value)
	return b
}

/*
Popover -
*/
func (b *BaseTagHtml) Popover() *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("popover", "")
	return b
}

/*
Role -
*/
func (b *BaseTagHtml) Role(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("role", value)
	return b
}

/*
Slot -
*/
func (b *BaseTagHtml) Slot(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("slot", value)
	return b
}

/*
Spellcheck -
*/
func (b *BaseTagHtml) Spellcheck(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("spellcheck", value)
	return b
}

/*
Style -
*/
func (b *BaseTagHtml) Style(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("style", value)
	return b
}

/*
Tabindex -
*/
func (b *BaseTagHtml) Tabindex(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("tabindex", value)
	return b
}

/*
Title -
*/
func (b *BaseTagHtml) Title(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("title", value)
	return b
}

/*
Translate -
*/
func (b *BaseTagHtml) Translate(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("translate", value)
	return b
}

/*
VirtualKeyBoardPolicy -
*/
func (b *BaseTagHtml) VirtualKeyBoardPolicy(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("virtualKeyBoardPolicy", value)
	return b
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (b *BaseTagHtml) AriaAtomic(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-atomic", value)
	return b
}

/*
AriaBusy -
*/
func (b *BaseTagHtml) AriaBusy(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-busy", value)
	return b
}

/*
AriaControls -
*/
func (b *BaseTagHtml) AriaControls(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-controls", value)
	return b
}

/*
AriaCurrent -
*/
func (b *BaseTagHtml) AriaCurrent(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-current", value)
	return b
}

/*
AriaDescribedby -
*/
func (b *BaseTagHtml) AriaDescribedby(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-describedby", value)
	return b
}

/*
AriaDescription -
*/
func (b *BaseTagHtml) AriaDescription(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-description", value)
	return b
}

/*
AriaDetails -
*/
func (b *BaseTagHtml) AriaDetails(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-details", value)
	return b
}

/*
AriaDisabled -
*/
func (b *BaseTagHtml) AriaDisabled(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-disabled", value)
	return b
}

/*
AriaDropeffect -
*/
func (b *BaseTagHtml) AriaDropeffect(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-dropeffect", value)
	return b
}

/*
AriaErrormessage -
*/
func (b *BaseTagHtml) AriaErrormessage(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-errormessage", value)
	return b
}

/*
AriaFlowto -
*/
func (b *BaseTagHtml) AriaFlowto(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-flowto", value)
	return b
}

/*
AriaGrabbed -
*/
func (b *BaseTagHtml) AriaGrabbed(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-grabbed", value)
	return b
}

/*
AriaHaspopup -
*/
func (b *BaseTagHtml) AriaHaspopup(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-haspopup", value)
	return b
}

/*
AriaHidden -
*/
func (b *BaseTagHtml) AriaHidden(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-hidden", value)
	return b
}

/*
AriaInvalid -
*/
func (b *BaseTagHtml) AriaInvalid(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-invalid", value)
	return b
}

/*
AriaKeyshortcuts -
*/
func (b *BaseTagHtml) AriaKeyshortcuts(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-keyshortcuts", value)
	return b
}

/*
AriaLabel -
*/
func (b *BaseTagHtml) AriaLabel(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-label", value)
	return b
}

/*
AriaLabelledby -
*/
func (b *BaseTagHtml) AriaLabelledby(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-labelledby", value)
	return b
}

/*
AriaLive -
*/
func (b *BaseTagHtml) AriaLive(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-live", value)
	return b
}

/*
AriaOwns -
*/
func (b *BaseTagHtml) AriaOwns(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-owns", value)
	return b
}

/*
AriaRelevant -
*/
func (b *BaseTagHtml) AriaRelevant(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-relevant", value)
	return b
}

/*
AriaRoledescription -
*/
func (b *BaseTagHtml) AriaRoledescription(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-roledescription", value)
	return b
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (b *BaseTagHtml) Onabort(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onabort", value)
	return b
}

/*
Onautocomplete -
*/
func (b *BaseTagHtml) Onautocomplete(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onautocomplete", value)
	return b
}

/*
Onautocompleteerror -
*/
func (b *BaseTagHtml) Onautocompleteerror(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onautocompleteerror", value)
	return b
}

/*
Onblur -
*/
func (b *BaseTagHtml) Onblur(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onblur", value)
	return b
}

/*
Oncancel -
*/
func (b *BaseTagHtml) Oncancel(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncancel", value)
	return b
}

/*
Oncanplay -
*/
func (b *BaseTagHtml) Oncanplay(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncanplay", value)
	return b
}

/*
Oncanplaythrough -
*/
func (b *BaseTagHtml) Oncanplaythrough(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncanplaythrough", value)
	return b
}

/*
Onchange -
*/
func (b *BaseTagHtml) Onchange(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onchange", value)
	return b
}

/*
Onclick -
*/
func (b *BaseTagHtml) Onclick(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onclick", value)
	return b
}

/*
Onclose -
*/
func (b *BaseTagHtml) Onclose(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onclose", value)
	return b
}

/*
Oncontextmenu -
*/
func (b *BaseTagHtml) Oncontextmenu(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncontextmenu", value)
	return b
}

/*
Oncuechange -
*/
func (b *BaseTagHtml) Oncuechange(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncuechange", value)
	return b
}

/*
Ondblclick -
*/
func (b *BaseTagHtml) Ondblclick(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondblclick", value)
	return b
}

/*
Ondrag -
*/
func (b *BaseTagHtml) Ondrag(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondrag", value)
	return b
}

/*
Ondragend -
*/
func (b *BaseTagHtml) Ondragend(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragend", value)
	return b
}

/*
Ondragenter -
*/
func (b *BaseTagHtml) Ondragenter(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragenter", value)
	return b
}

/*
Ondragleave -
*/
func (b *BaseTagHtml) Ondragleave(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragleave", value)
	return b
}

/*
Ondragover -
*/
func (b *BaseTagHtml) Ondragover(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragover", value)
	return b
}

/*
Ondragstart -
*/
func (b *BaseTagHtml) Ondragstart(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragstart", value)
	return b
}

/*
Ondrop -
*/
func (b *BaseTagHtml) Ondrop(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondrop", value)
	return b
}

/*
Ondurationchange -
*/
func (b *BaseTagHtml) Ondurationchange(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondurationchange", value)
	return b
}

/*
Onemptied -
*/
func (b *BaseTagHtml) Onemptied(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onemptied", value)
	return b
}

/*
Onended -
*/
func (b *BaseTagHtml) Onended(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onended", value)
	return b
}

/*
Onfocus -
*/
func (b *BaseTagHtml) Onfocus(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onfocus", value)
	return b
}

/*
Oninput -
*/
func (b *BaseTagHtml) Oninput(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oninput", value)
	return b
}

/*
Oninvalid -
*/
func (b *BaseTagHtml) Oninvalid(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oninvalid", value)
	return b
}

/*
Onkeydown -
*/
func (b *BaseTagHtml) Onkeydown(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeydown", value)
	return b
}

/*
Onkeypress -
*/
func (b *BaseTagHtml) Onkeypress(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeypress", value)
	return b
}

/*
Onkeyup -
*/
func (b *BaseTagHtml) Onkeyup(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeyup", value)
	return b
}

/*
Onloadeddata -
*/
func (b *BaseTagHtml) Onloadeddata(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadeddata", value)
	return b
}

/*
Onloadedmetadata -
*/
func (b *BaseTagHtml) Onloadedmetadata(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadedmetadata", value)
	return b
}

/*
Onloadstart -
*/
func (b *BaseTagHtml) Onloadstart(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadstart", value)
	return b
}

/*
Onmousedown -
*/
func (b *BaseTagHtml) Onmousedown(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousedown", value)
	return b
}

/*
Onmouseenter -
*/
func (b *BaseTagHtml) Onmouseenter(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseenter", value)
	return b
}

/*
Onmouseleave -
*/
func (b *BaseTagHtml) Onmouseleave(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseleave", value)
	return b
}

/*
Onmousemove -
*/
func (b *BaseTagHtml) Onmousemove(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousemove", value)
	return b
}

/*
Onmouseout -
*/
func (b *BaseTagHtml) Onmouseout(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseout", value)
	return b
}

/*
Onmouseover -
*/
func (b *BaseTagHtml) Onmouseover(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseover", value)
	return b
}

/*
Onmouseup -
*/
func (b *BaseTagHtml) Onmouseup(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseup", value)
	return b
}

/*
Onmousewheel -
*/
func (b *BaseTagHtml) Onmousewheel(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousewheel", value)
	return b
}

/*
Onpause -
*/
func (b *BaseTagHtml) Onpause(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpause", value)
	return b
}

/*
Onplay -
*/
func (b *BaseTagHtml) Onplay(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onplay", value)
	return b
}

/*
Onplaying -
*/
func (b *BaseTagHtml) Onplaying(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onplaying", value)
	return b
}

/*
Onprogress -
*/
func (b *BaseTagHtml) Onprogress(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onprogress", value)
	return b
}

/*
Onratechange -
*/
func (b *BaseTagHtml) Onratechange(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onratechange", value)
	return b
}

/*
Onreset -
*/
func (b *BaseTagHtml) Onreset(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onreset", value)
	return b
}

/*
Onscroll -
*/
func (b *BaseTagHtml) Onscroll(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onscroll", value)
	return b
}

/*
Onseeked -
*/
func (b *BaseTagHtml) Onseeked(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onseeked", value)
	return b
}

/*
Onseeking -
*/
func (b *BaseTagHtml) Onseeking(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onseeking", value)
	return b
}

/*
Onselect -
*/
func (b *BaseTagHtml) Onselect(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onselect", value)
	return b
}

/*
Onshow -
*/
func (b *BaseTagHtml) Onshow(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onshow", value)
	return b
}

/*
Onsort -
*/
func (b *BaseTagHtml) Onsort(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsort", value)
	return b
}

/*
Onstalled -
*/
func (b *BaseTagHtml) Onstalled(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onstalled", value)
	return b
}

/*
Onsubmit -
*/
func (b *BaseTagHtml) Onsubmit(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsubmit", value)
	return b
}

/*
Onsuspend -
*/
func (b *BaseTagHtml) Onsuspend(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsuspend", value)
	return b
}

/*
Ontimeupdate -
*/
func (b *BaseTagHtml) Ontimeupdate(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ontimeupdate", value)
	return b
}

/*
Ontoggle -
*/
func (b *BaseTagHtml) Ontoggle(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ontoggle", value)
	return b
}

/*
Onvolumechange -
*/
func (b *BaseTagHtml) Onvolumechange(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onvolumechange", value)
	return b
}

/*
Onwaiting -
*/
func (b *BaseTagHtml) Onwaiting(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onwaiting", value)
	return b
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (b *BaseTagHtml) Onafterprint(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onafterprint", value)
	return b
}

/*
Onbeforeprint -
*/
func (b *BaseTagHtml) Onbeforeprint(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onbeforeprint", value)
	return b
}

/*
Onbeforeunload -
*/
func (b *BaseTagHtml) Onbeforeunload(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onbeforeunload", value)
	return b
}

/*
Onerror -
*/
func (b *BaseTagHtml) Onerror(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onerror", value)
	return b
}

/*
Onhashchange -
*/
func (b *BaseTagHtml) Onhashchange(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onhashchange", value)
	return b
}

/*
Onload -
*/
func (b *BaseTagHtml) Onload(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onload", value)
	return b
}

/*
Onmessage -
*/
func (b *BaseTagHtml) Onmessage(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmessage", value)
	return b
}

/*
Onoffline -
*/
func (b *BaseTagHtml) Onoffline(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onoffline", value)
	return b
}

/*
Ononline -
*/
func (b *BaseTagHtml) Ononline(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ononline", value)
	return b
}

/*
Onpagehide -
*/
func (b *BaseTagHtml) Onpagehide(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpagehide", value)
	return b
}

/*
Onpageshow -
*/
func (b *BaseTagHtml) Onpageshow(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpageshow", value)
	return b
}

/*
Onpopstate -
*/
func (b *BaseTagHtml) Onpopstate(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpopstate", value)
	return b
}

/*
Onresize -
*/
func (b *BaseTagHtml) Onresize(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onresize", value)
	return b
}

/*
Onstorage -
*/
func (b *BaseTagHtml) Onstorage(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onstorage", value)
	return b
}

/*
Onunload -
*/
func (b *BaseTagHtml) Onunload(value string) *BaseTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onunload", value)
	return b
}
