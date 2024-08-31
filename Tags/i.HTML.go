// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func IHtml(tags []any) *ITagHtml {
	node := &ITagHtml{
		tag: &tag{
			name:                "i",
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

type ITagHtml struct {
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
func (i *ITagHtml) CustomAttribute(attributes ...*Attribute) *ITagHtml {
	i.registerAttributes(attributes...)
	return i
}

/*
Children - Method for nesting tags into parent tag
*/
func (i *ITagHtml) Children(tags ...any) *ITagHtml {
	return i.supportedChildrenCheck(tags)
}

func (i *ITagHtml) supportedChildrenCheck(tags []any) *ITagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			i.registerChildren(TextContentSvg(val).getTag())
		case content:
			i.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				i.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return i
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
func (i *ITagHtml) AccessKey(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("accessKey", value)
	return i
}

/*
Aria -
*/
func (i *ITagHtml) Aria(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria", value)
	return i
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (i *ITagHtml) Autocapitalize(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("autocapitalize", value)
	return i
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (i *ITagHtml) Autofocus(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("autofocus", value)
	return i
}

/*
Class -
*/
func (i *ITagHtml) Class(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("class", value)
	return i
}

/*
Contenteditable -
*/
func (i *ITagHtml) Contenteditable(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("contenteditable", value)
	return i
}

/*
Data -
*/
func (i *ITagHtml) Data(name, value string) *ITagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	i.registerAttribute(dataName, value)
	return i
}

/*
Dir -
*/
func (i *ITagHtml) Dir(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("dir", value)
	return i
}

/*
Draggable -
*/
func (i *ITagHtml) Draggable(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("draggable", value)
	return i
}

/*
EnterKeyHint -
*/
func (i *ITagHtml) EnterKeyHint(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("enterKeyHint", value)
	return i
}

/*
ExportParts -
*/
func (i *ITagHtml) ExportParts(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("exportParts", value)
	return i
}

/*
Hidden -
*/
func (i *ITagHtml) Hidden(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("hidden", value)
	return i
}

/*
Id -
*/
func (i *ITagHtml) Id(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("id", value)
	return i
}

/*
Inert -
*/
func (i *ITagHtml) Inert(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inert", value)
	return i
}

/*
InputMode -
*/
func (i *ITagHtml) InputMode(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inputMode", value)
	return i
}

/*
Is -
*/
func (i *ITagHtml) Is(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("is", value)
	return i
}

/*
ItemId -
*/
func (i *ITagHtml) ItemId(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemId", value)
	return i
}

/*
ItemProp -
*/
func (i *ITagHtml) ItemProp(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemProp", value)
	return i
}

/*
ItemRef -
*/
func (i *ITagHtml) ItemRef(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemRef", value)
	return i
}

/*
ItemScope -
*/
func (i *ITagHtml) ItemScope(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemScope", value)
	return i
}

/*
ItemType -
*/
func (i *ITagHtml) ItemType(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemType", value)
	return i
}

/*
Lang -
*/
func (i *ITagHtml) Lang(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("lang", value)
	return i
}

/*
Nonce -
*/
func (i *ITagHtml) Nonce(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("nonce", value)
	return i
}

/*
Part -
*/
func (i *ITagHtml) Part(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("part", value)
	return i
}

/*
Popover -
*/
func (i *ITagHtml) Popover() *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("popover", "")
	return i
}

/*
Role -
*/
func (i *ITagHtml) Role(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("role", value)
	return i
}

/*
Slot -
*/
func (i *ITagHtml) Slot(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("slot", value)
	return i
}

/*
Spellcheck -
*/
func (i *ITagHtml) Spellcheck(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("spellcheck", value)
	return i
}

/*
Style -
*/
func (i *ITagHtml) Style(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("style", value)
	return i
}

/*
Tabindex -
*/
func (i *ITagHtml) Tabindex(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("tabindex", value)
	return i
}

/*
Title -
*/
func (i *ITagHtml) Title(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("title", value)
	return i
}

/*
Translate -
*/
func (i *ITagHtml) Translate(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("translate", value)
	return i
}

/*
VirtualKeyBoardPolicy -
*/
func (i *ITagHtml) VirtualKeyBoardPolicy(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("virtualKeyBoardPolicy", value)
	return i
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (i *ITagHtml) AriaAtomic(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-atomic", value)
	return i
}

/*
AriaBusy -
*/
func (i *ITagHtml) AriaBusy(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-busy", value)
	return i
}

/*
AriaControls -
*/
func (i *ITagHtml) AriaControls(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-controls", value)
	return i
}

/*
AriaCurrent -
*/
func (i *ITagHtml) AriaCurrent(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-current", value)
	return i
}

/*
AriaDescribedby -
*/
func (i *ITagHtml) AriaDescribedby(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-describedby", value)
	return i
}

/*
AriaDescription -
*/
func (i *ITagHtml) AriaDescription(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-description", value)
	return i
}

/*
AriaDetails -
*/
func (i *ITagHtml) AriaDetails(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-details", value)
	return i
}

/*
AriaDisabled -
*/
func (i *ITagHtml) AriaDisabled(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-disabled", value)
	return i
}

/*
AriaDropeffect -
*/
func (i *ITagHtml) AriaDropeffect(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-dropeffect", value)
	return i
}

/*
AriaErrormessage -
*/
func (i *ITagHtml) AriaErrormessage(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-errormessage", value)
	return i
}

/*
AriaFlowto -
*/
func (i *ITagHtml) AriaFlowto(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-flowto", value)
	return i
}

/*
AriaGrabbed -
*/
func (i *ITagHtml) AriaGrabbed(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-grabbed", value)
	return i
}

/*
AriaHaspopup -
*/
func (i *ITagHtml) AriaHaspopup(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-haspopup", value)
	return i
}

/*
AriaHidden -
*/
func (i *ITagHtml) AriaHidden(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-hidden", value)
	return i
}

/*
AriaInvalid -
*/
func (i *ITagHtml) AriaInvalid(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-invalid", value)
	return i
}

/*
AriaKeyshortcuts -
*/
func (i *ITagHtml) AriaKeyshortcuts(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-keyshortcuts", value)
	return i
}

/*
AriaLabel -
*/
func (i *ITagHtml) AriaLabel(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-label", value)
	return i
}

/*
AriaLabelledby -
*/
func (i *ITagHtml) AriaLabelledby(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-labelledby", value)
	return i
}

/*
AriaLive -
*/
func (i *ITagHtml) AriaLive(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-live", value)
	return i
}

/*
AriaOwns -
*/
func (i *ITagHtml) AriaOwns(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-owns", value)
	return i
}

/*
AriaRelevant -
*/
func (i *ITagHtml) AriaRelevant(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-relevant", value)
	return i
}

/*
AriaRoledescription -
*/
func (i *ITagHtml) AriaRoledescription(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-roledescription", value)
	return i
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (i *ITagHtml) Onabort(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onabort", value)
	return i
}

/*
Onautocomplete -
*/
func (i *ITagHtml) Onautocomplete(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocomplete", value)
	return i
}

/*
Onautocompleteerror -
*/
func (i *ITagHtml) Onautocompleteerror(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocompleteerror", value)
	return i
}

/*
Onblur -
*/
func (i *ITagHtml) Onblur(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onblur", value)
	return i
}

/*
Oncancel -
*/
func (i *ITagHtml) Oncancel(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncancel", value)
	return i
}

/*
Oncanplay -
*/
func (i *ITagHtml) Oncanplay(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplay", value)
	return i
}

/*
Oncanplaythrough -
*/
func (i *ITagHtml) Oncanplaythrough(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplaythrough", value)
	return i
}

/*
Onchange -
*/
func (i *ITagHtml) Onchange(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onchange", value)
	return i
}

/*
Onclick -
*/
func (i *ITagHtml) Onclick(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclick", value)
	return i
}

/*
Onclose -
*/
func (i *ITagHtml) Onclose(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclose", value)
	return i
}

/*
Oncontextmenu -
*/
func (i *ITagHtml) Oncontextmenu(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncontextmenu", value)
	return i
}

/*
Oncuechange -
*/
func (i *ITagHtml) Oncuechange(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncuechange", value)
	return i
}

/*
Ondblclick -
*/
func (i *ITagHtml) Ondblclick(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondblclick", value)
	return i
}

/*
Ondrag -
*/
func (i *ITagHtml) Ondrag(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrag", value)
	return i
}

/*
Ondragend -
*/
func (i *ITagHtml) Ondragend(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragend", value)
	return i
}

/*
Ondragenter -
*/
func (i *ITagHtml) Ondragenter(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragenter", value)
	return i
}

/*
Ondragleave -
*/
func (i *ITagHtml) Ondragleave(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragleave", value)
	return i
}

/*
Ondragover -
*/
func (i *ITagHtml) Ondragover(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragover", value)
	return i
}

/*
Ondragstart -
*/
func (i *ITagHtml) Ondragstart(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragstart", value)
	return i
}

/*
Ondrop -
*/
func (i *ITagHtml) Ondrop(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrop", value)
	return i
}

/*
Ondurationchange -
*/
func (i *ITagHtml) Ondurationchange(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondurationchange", value)
	return i
}

/*
Onemptied -
*/
func (i *ITagHtml) Onemptied(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onemptied", value)
	return i
}

/*
Onended -
*/
func (i *ITagHtml) Onended(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onended", value)
	return i
}

/*
Onfocus -
*/
func (i *ITagHtml) Onfocus(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onfocus", value)
	return i
}

/*
Oninput -
*/
func (i *ITagHtml) Oninput(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninput", value)
	return i
}

/*
Oninvalid -
*/
func (i *ITagHtml) Oninvalid(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninvalid", value)
	return i
}

/*
Onkeydown -
*/
func (i *ITagHtml) Onkeydown(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeydown", value)
	return i
}

/*
Onkeypress -
*/
func (i *ITagHtml) Onkeypress(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeypress", value)
	return i
}

/*
Onkeyup -
*/
func (i *ITagHtml) Onkeyup(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeyup", value)
	return i
}

/*
Onloadeddata -
*/
func (i *ITagHtml) Onloadeddata(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadeddata", value)
	return i
}

/*
Onloadedmetadata -
*/
func (i *ITagHtml) Onloadedmetadata(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadedmetadata", value)
	return i
}

/*
Onloadstart -
*/
func (i *ITagHtml) Onloadstart(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadstart", value)
	return i
}

/*
Onmousedown -
*/
func (i *ITagHtml) Onmousedown(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousedown", value)
	return i
}

/*
Onmouseenter -
*/
func (i *ITagHtml) Onmouseenter(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseenter", value)
	return i
}

/*
Onmouseleave -
*/
func (i *ITagHtml) Onmouseleave(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseleave", value)
	return i
}

/*
Onmousemove -
*/
func (i *ITagHtml) Onmousemove(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousemove", value)
	return i
}

/*
Onmouseout -
*/
func (i *ITagHtml) Onmouseout(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseout", value)
	return i
}

/*
Onmouseover -
*/
func (i *ITagHtml) Onmouseover(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseover", value)
	return i
}

/*
Onmouseup -
*/
func (i *ITagHtml) Onmouseup(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseup", value)
	return i
}

/*
Onmousewheel -
*/
func (i *ITagHtml) Onmousewheel(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousewheel", value)
	return i
}

/*
Onpause -
*/
func (i *ITagHtml) Onpause(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpause", value)
	return i
}

/*
Onplay -
*/
func (i *ITagHtml) Onplay(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplay", value)
	return i
}

/*
Onplaying -
*/
func (i *ITagHtml) Onplaying(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplaying", value)
	return i
}

/*
Onprogress -
*/
func (i *ITagHtml) Onprogress(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onprogress", value)
	return i
}

/*
Onratechange -
*/
func (i *ITagHtml) Onratechange(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onratechange", value)
	return i
}

/*
Onreset -
*/
func (i *ITagHtml) Onreset(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onreset", value)
	return i
}

/*
Onscroll -
*/
func (i *ITagHtml) Onscroll(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onscroll", value)
	return i
}

/*
Onseeked -
*/
func (i *ITagHtml) Onseeked(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeked", value)
	return i
}

/*
Onseeking -
*/
func (i *ITagHtml) Onseeking(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeking", value)
	return i
}

/*
Onselect -
*/
func (i *ITagHtml) Onselect(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onselect", value)
	return i
}

/*
Onshow -
*/
func (i *ITagHtml) Onshow(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onshow", value)
	return i
}

/*
Onsort -
*/
func (i *ITagHtml) Onsort(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsort", value)
	return i
}

/*
Onstalled -
*/
func (i *ITagHtml) Onstalled(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstalled", value)
	return i
}

/*
Onsubmit -
*/
func (i *ITagHtml) Onsubmit(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsubmit", value)
	return i
}

/*
Onsuspend -
*/
func (i *ITagHtml) Onsuspend(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsuspend", value)
	return i
}

/*
Ontimeupdate -
*/
func (i *ITagHtml) Ontimeupdate(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontimeupdate", value)
	return i
}

/*
Ontoggle -
*/
func (i *ITagHtml) Ontoggle(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontoggle", value)
	return i
}

/*
Onvolumechange -
*/
func (i *ITagHtml) Onvolumechange(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onvolumechange", value)
	return i
}

/*
Onwaiting -
*/
func (i *ITagHtml) Onwaiting(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onwaiting", value)
	return i
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (i *ITagHtml) Onafterprint(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onafterprint", value)
	return i
}

/*
Onbeforeprint -
*/
func (i *ITagHtml) Onbeforeprint(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeprint", value)
	return i
}

/*
Onbeforeunload -
*/
func (i *ITagHtml) Onbeforeunload(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeunload", value)
	return i
}

/*
Onerror -
*/
func (i *ITagHtml) Onerror(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onerror", value)
	return i
}

/*
Onhashchange -
*/
func (i *ITagHtml) Onhashchange(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onhashchange", value)
	return i
}

/*
Onload -
*/
func (i *ITagHtml) Onload(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onload", value)
	return i
}

/*
Onmessage -
*/
func (i *ITagHtml) Onmessage(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmessage", value)
	return i
}

/*
Onoffline -
*/
func (i *ITagHtml) Onoffline(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onoffline", value)
	return i
}

/*
Ononline -
*/
func (i *ITagHtml) Ononline(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ononline", value)
	return i
}

/*
Onpagehide -
*/
func (i *ITagHtml) Onpagehide(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpagehide", value)
	return i
}

/*
Onpageshow -
*/
func (i *ITagHtml) Onpageshow(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpageshow", value)
	return i
}

/*
Onpopstate -
*/
func (i *ITagHtml) Onpopstate(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpopstate", value)
	return i
}

/*
Onresize -
*/
func (i *ITagHtml) Onresize(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onresize", value)
	return i
}

/*
Onstorage -
*/
func (i *ITagHtml) Onstorage(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstorage", value)
	return i
}

/*
Onunload -
*/
func (i *ITagHtml) Onunload(value string) *ITagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onunload", value)
	return i
}
