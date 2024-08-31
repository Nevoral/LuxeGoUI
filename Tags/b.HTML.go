// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func BHtml(tags []any) *BTagHtml {
	node := &BTagHtml{
		tag: &tag{
			name:                "b",
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

type BTagHtml struct {
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
func (b *BTagHtml) CustomAttribute(attributes ...*Attribute) *BTagHtml {
	b.registerAttributes(attributes...)
	return b
}

/*
Children - Method for nesting tags into parent tag
*/
func (b *BTagHtml) Children(tags ...any) *BTagHtml {
	return b.supportedChildrenCheck(tags)
}

func (b *BTagHtml) supportedChildrenCheck(tags []any) *BTagHtml {
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
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (b *BTagHtml) AccessKey(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("accessKey", value)
	return b
}

/*
Aria -
*/
func (b *BTagHtml) Aria(value string) *BTagHtml {
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
func (b *BTagHtml) Autocapitalize(value string) *BTagHtml {
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
func (b *BTagHtml) Autofocus(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("autofocus", value)
	return b
}

/*
Class -
*/
func (b *BTagHtml) Class(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("class", value)
	return b
}

/*
Contenteditable -
*/
func (b *BTagHtml) Contenteditable(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("contenteditable", value)
	return b
}

/*
Data -
*/
func (b *BTagHtml) Data(name, value string) *BTagHtml {
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
func (b *BTagHtml) Dir(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("dir", value)
	return b
}

/*
Draggable -
*/
func (b *BTagHtml) Draggable(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("draggable", value)
	return b
}

/*
EnterKeyHint -
*/
func (b *BTagHtml) EnterKeyHint(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("enterKeyHint", value)
	return b
}

/*
ExportParts -
*/
func (b *BTagHtml) ExportParts(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("exportParts", value)
	return b
}

/*
Hidden -
*/
func (b *BTagHtml) Hidden(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("hidden", value)
	return b
}

/*
Id -
*/
func (b *BTagHtml) Id(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("id", value)
	return b
}

/*
Inert -
*/
func (b *BTagHtml) Inert(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("inert", value)
	return b
}

/*
InputMode -
*/
func (b *BTagHtml) InputMode(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("inputMode", value)
	return b
}

/*
Is -
*/
func (b *BTagHtml) Is(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("is", value)
	return b
}

/*
ItemId -
*/
func (b *BTagHtml) ItemId(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemId", value)
	return b
}

/*
ItemProp -
*/
func (b *BTagHtml) ItemProp(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemProp", value)
	return b
}

/*
ItemRef -
*/
func (b *BTagHtml) ItemRef(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemRef", value)
	return b
}

/*
ItemScope -
*/
func (b *BTagHtml) ItemScope(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemScope", value)
	return b
}

/*
ItemType -
*/
func (b *BTagHtml) ItemType(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemType", value)
	return b
}

/*
Lang -
*/
func (b *BTagHtml) Lang(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("lang", value)
	return b
}

/*
Nonce -
*/
func (b *BTagHtml) Nonce(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("nonce", value)
	return b
}

/*
Part -
*/
func (b *BTagHtml) Part(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("part", value)
	return b
}

/*
Popover -
*/
func (b *BTagHtml) Popover() *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("popover", "")
	return b
}

/*
Role -
*/
func (b *BTagHtml) Role(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("role", value)
	return b
}

/*
Slot -
*/
func (b *BTagHtml) Slot(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("slot", value)
	return b
}

/*
Spellcheck -
*/
func (b *BTagHtml) Spellcheck(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("spellcheck", value)
	return b
}

/*
Style -
*/
func (b *BTagHtml) Style(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("style", value)
	return b
}

/*
Tabindex -
*/
func (b *BTagHtml) Tabindex(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("tabindex", value)
	return b
}

/*
Title -
*/
func (b *BTagHtml) Title(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("title", value)
	return b
}

/*
Translate -
*/
func (b *BTagHtml) Translate(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("translate", value)
	return b
}

/*
VirtualKeyBoardPolicy -
*/
func (b *BTagHtml) VirtualKeyBoardPolicy(value string) *BTagHtml {
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
func (b *BTagHtml) AriaAtomic(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-atomic", value)
	return b
}

/*
AriaBusy -
*/
func (b *BTagHtml) AriaBusy(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-busy", value)
	return b
}

/*
AriaControls -
*/
func (b *BTagHtml) AriaControls(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-controls", value)
	return b
}

/*
AriaCurrent -
*/
func (b *BTagHtml) AriaCurrent(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-current", value)
	return b
}

/*
AriaDescribedby -
*/
func (b *BTagHtml) AriaDescribedby(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-describedby", value)
	return b
}

/*
AriaDescription -
*/
func (b *BTagHtml) AriaDescription(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-description", value)
	return b
}

/*
AriaDetails -
*/
func (b *BTagHtml) AriaDetails(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-details", value)
	return b
}

/*
AriaDisabled -
*/
func (b *BTagHtml) AriaDisabled(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-disabled", value)
	return b
}

/*
AriaDropeffect -
*/
func (b *BTagHtml) AriaDropeffect(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-dropeffect", value)
	return b
}

/*
AriaErrormessage -
*/
func (b *BTagHtml) AriaErrormessage(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-errormessage", value)
	return b
}

/*
AriaFlowto -
*/
func (b *BTagHtml) AriaFlowto(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-flowto", value)
	return b
}

/*
AriaGrabbed -
*/
func (b *BTagHtml) AriaGrabbed(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-grabbed", value)
	return b
}

/*
AriaHaspopup -
*/
func (b *BTagHtml) AriaHaspopup(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-haspopup", value)
	return b
}

/*
AriaHidden -
*/
func (b *BTagHtml) AriaHidden(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-hidden", value)
	return b
}

/*
AriaInvalid -
*/
func (b *BTagHtml) AriaInvalid(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-invalid", value)
	return b
}

/*
AriaKeyshortcuts -
*/
func (b *BTagHtml) AriaKeyshortcuts(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-keyshortcuts", value)
	return b
}

/*
AriaLabel -
*/
func (b *BTagHtml) AriaLabel(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-label", value)
	return b
}

/*
AriaLabelledby -
*/
func (b *BTagHtml) AriaLabelledby(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-labelledby", value)
	return b
}

/*
AriaLive -
*/
func (b *BTagHtml) AriaLive(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-live", value)
	return b
}

/*
AriaOwns -
*/
func (b *BTagHtml) AriaOwns(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-owns", value)
	return b
}

/*
AriaRelevant -
*/
func (b *BTagHtml) AriaRelevant(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-relevant", value)
	return b
}

/*
AriaRoledescription -
*/
func (b *BTagHtml) AriaRoledescription(value string) *BTagHtml {
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
func (b *BTagHtml) Onabort(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onabort", value)
	return b
}

/*
Onautocomplete -
*/
func (b *BTagHtml) Onautocomplete(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onautocomplete", value)
	return b
}

/*
Onautocompleteerror -
*/
func (b *BTagHtml) Onautocompleteerror(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onautocompleteerror", value)
	return b
}

/*
Onblur -
*/
func (b *BTagHtml) Onblur(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onblur", value)
	return b
}

/*
Oncancel -
*/
func (b *BTagHtml) Oncancel(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncancel", value)
	return b
}

/*
Oncanplay -
*/
func (b *BTagHtml) Oncanplay(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncanplay", value)
	return b
}

/*
Oncanplaythrough -
*/
func (b *BTagHtml) Oncanplaythrough(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncanplaythrough", value)
	return b
}

/*
Onchange -
*/
func (b *BTagHtml) Onchange(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onchange", value)
	return b
}

/*
Onclick -
*/
func (b *BTagHtml) Onclick(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onclick", value)
	return b
}

/*
Onclose -
*/
func (b *BTagHtml) Onclose(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onclose", value)
	return b
}

/*
Oncontextmenu -
*/
func (b *BTagHtml) Oncontextmenu(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncontextmenu", value)
	return b
}

/*
Oncuechange -
*/
func (b *BTagHtml) Oncuechange(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncuechange", value)
	return b
}

/*
Ondblclick -
*/
func (b *BTagHtml) Ondblclick(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondblclick", value)
	return b
}

/*
Ondrag -
*/
func (b *BTagHtml) Ondrag(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondrag", value)
	return b
}

/*
Ondragend -
*/
func (b *BTagHtml) Ondragend(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragend", value)
	return b
}

/*
Ondragenter -
*/
func (b *BTagHtml) Ondragenter(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragenter", value)
	return b
}

/*
Ondragleave -
*/
func (b *BTagHtml) Ondragleave(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragleave", value)
	return b
}

/*
Ondragover -
*/
func (b *BTagHtml) Ondragover(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragover", value)
	return b
}

/*
Ondragstart -
*/
func (b *BTagHtml) Ondragstart(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragstart", value)
	return b
}

/*
Ondrop -
*/
func (b *BTagHtml) Ondrop(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondrop", value)
	return b
}

/*
Ondurationchange -
*/
func (b *BTagHtml) Ondurationchange(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondurationchange", value)
	return b
}

/*
Onemptied -
*/
func (b *BTagHtml) Onemptied(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onemptied", value)
	return b
}

/*
Onended -
*/
func (b *BTagHtml) Onended(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onended", value)
	return b
}

/*
Onfocus -
*/
func (b *BTagHtml) Onfocus(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onfocus", value)
	return b
}

/*
Oninput -
*/
func (b *BTagHtml) Oninput(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oninput", value)
	return b
}

/*
Oninvalid -
*/
func (b *BTagHtml) Oninvalid(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oninvalid", value)
	return b
}

/*
Onkeydown -
*/
func (b *BTagHtml) Onkeydown(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeydown", value)
	return b
}

/*
Onkeypress -
*/
func (b *BTagHtml) Onkeypress(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeypress", value)
	return b
}

/*
Onkeyup -
*/
func (b *BTagHtml) Onkeyup(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeyup", value)
	return b
}

/*
Onloadeddata -
*/
func (b *BTagHtml) Onloadeddata(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadeddata", value)
	return b
}

/*
Onloadedmetadata -
*/
func (b *BTagHtml) Onloadedmetadata(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadedmetadata", value)
	return b
}

/*
Onloadstart -
*/
func (b *BTagHtml) Onloadstart(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadstart", value)
	return b
}

/*
Onmousedown -
*/
func (b *BTagHtml) Onmousedown(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousedown", value)
	return b
}

/*
Onmouseenter -
*/
func (b *BTagHtml) Onmouseenter(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseenter", value)
	return b
}

/*
Onmouseleave -
*/
func (b *BTagHtml) Onmouseleave(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseleave", value)
	return b
}

/*
Onmousemove -
*/
func (b *BTagHtml) Onmousemove(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousemove", value)
	return b
}

/*
Onmouseout -
*/
func (b *BTagHtml) Onmouseout(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseout", value)
	return b
}

/*
Onmouseover -
*/
func (b *BTagHtml) Onmouseover(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseover", value)
	return b
}

/*
Onmouseup -
*/
func (b *BTagHtml) Onmouseup(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseup", value)
	return b
}

/*
Onmousewheel -
*/
func (b *BTagHtml) Onmousewheel(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousewheel", value)
	return b
}

/*
Onpause -
*/
func (b *BTagHtml) Onpause(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpause", value)
	return b
}

/*
Onplay -
*/
func (b *BTagHtml) Onplay(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onplay", value)
	return b
}

/*
Onplaying -
*/
func (b *BTagHtml) Onplaying(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onplaying", value)
	return b
}

/*
Onprogress -
*/
func (b *BTagHtml) Onprogress(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onprogress", value)
	return b
}

/*
Onratechange -
*/
func (b *BTagHtml) Onratechange(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onratechange", value)
	return b
}

/*
Onreset -
*/
func (b *BTagHtml) Onreset(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onreset", value)
	return b
}

/*
Onscroll -
*/
func (b *BTagHtml) Onscroll(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onscroll", value)
	return b
}

/*
Onseeked -
*/
func (b *BTagHtml) Onseeked(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onseeked", value)
	return b
}

/*
Onseeking -
*/
func (b *BTagHtml) Onseeking(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onseeking", value)
	return b
}

/*
Onselect -
*/
func (b *BTagHtml) Onselect(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onselect", value)
	return b
}

/*
Onshow -
*/
func (b *BTagHtml) Onshow(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onshow", value)
	return b
}

/*
Onsort -
*/
func (b *BTagHtml) Onsort(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsort", value)
	return b
}

/*
Onstalled -
*/
func (b *BTagHtml) Onstalled(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onstalled", value)
	return b
}

/*
Onsubmit -
*/
func (b *BTagHtml) Onsubmit(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsubmit", value)
	return b
}

/*
Onsuspend -
*/
func (b *BTagHtml) Onsuspend(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsuspend", value)
	return b
}

/*
Ontimeupdate -
*/
func (b *BTagHtml) Ontimeupdate(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ontimeupdate", value)
	return b
}

/*
Ontoggle -
*/
func (b *BTagHtml) Ontoggle(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ontoggle", value)
	return b
}

/*
Onvolumechange -
*/
func (b *BTagHtml) Onvolumechange(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onvolumechange", value)
	return b
}

/*
Onwaiting -
*/
func (b *BTagHtml) Onwaiting(value string) *BTagHtml {
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
func (b *BTagHtml) Onafterprint(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onafterprint", value)
	return b
}

/*
Onbeforeprint -
*/
func (b *BTagHtml) Onbeforeprint(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onbeforeprint", value)
	return b
}

/*
Onbeforeunload -
*/
func (b *BTagHtml) Onbeforeunload(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onbeforeunload", value)
	return b
}

/*
Onerror -
*/
func (b *BTagHtml) Onerror(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onerror", value)
	return b
}

/*
Onhashchange -
*/
func (b *BTagHtml) Onhashchange(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onhashchange", value)
	return b
}

/*
Onload -
*/
func (b *BTagHtml) Onload(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onload", value)
	return b
}

/*
Onmessage -
*/
func (b *BTagHtml) Onmessage(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmessage", value)
	return b
}

/*
Onoffline -
*/
func (b *BTagHtml) Onoffline(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onoffline", value)
	return b
}

/*
Ononline -
*/
func (b *BTagHtml) Ononline(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ononline", value)
	return b
}

/*
Onpagehide -
*/
func (b *BTagHtml) Onpagehide(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpagehide", value)
	return b
}

/*
Onpageshow -
*/
func (b *BTagHtml) Onpageshow(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpageshow", value)
	return b
}

/*
Onpopstate -
*/
func (b *BTagHtml) Onpopstate(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpopstate", value)
	return b
}

/*
Onresize -
*/
func (b *BTagHtml) Onresize(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onresize", value)
	return b
}

/*
Onstorage -
*/
func (b *BTagHtml) Onstorage(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onstorage", value)
	return b
}

/*
Onunload -
*/
func (b *BTagHtml) Onunload(value string) *BTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onunload", value)
	return b
}
