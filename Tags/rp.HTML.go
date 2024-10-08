// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func RpHtml(tags []any) *RpTagHtml {
	node := &RpTagHtml{
		tag: &tag{
			name:                "rp",
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

type RpTagHtml struct {
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
func (r *RpTagHtml) CustomAttribute(attributes ...*Attribute) *RpTagHtml {
	r.registerAttributes(attributes...)
	return r
}

/*
Children - Method for nesting tags into parent tag
*/
func (r *RpTagHtml) Children(tags ...any) *RpTagHtml {
	return r.supportedChildrenCheck(tags)
}

func (r *RpTagHtml) supportedChildrenCheck(tags []any) *RpTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			r.registerChildren(TextContentSvg(val).getTag())
		case content:
			r.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				r.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return r
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
func (r *RpTagHtml) AccessKey(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("accessKey", value)
	return r
}

/*
Aria -
*/
func (r *RpTagHtml) Aria(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria", value)
	return r
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (r *RpTagHtml) Autocapitalize(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("autocapitalize", value)
	return r
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (r *RpTagHtml) Autofocus(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("autofocus", value)
	return r
}

/*
Class -
*/
func (r *RpTagHtml) Class(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("class", value)
	return r
}

/*
Contenteditable -
*/
func (r *RpTagHtml) Contenteditable(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("contenteditable", value)
	return r
}

/*
Data -
*/
func (r *RpTagHtml) Data(name, value string) *RpTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	r.registerAttribute(dataName, value)
	return r
}

/*
Dir -
*/
func (r *RpTagHtml) Dir(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("dir", value)
	return r
}

/*
Draggable -
*/
func (r *RpTagHtml) Draggable(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("draggable", value)
	return r
}

/*
EnterKeyHint -
*/
func (r *RpTagHtml) EnterKeyHint(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("enterKeyHint", value)
	return r
}

/*
ExportParts -
*/
func (r *RpTagHtml) ExportParts(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("exportParts", value)
	return r
}

/*
Hidden -
*/
func (r *RpTagHtml) Hidden(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("hidden", value)
	return r
}

/*
Id -
*/
func (r *RpTagHtml) Id(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("id", value)
	return r
}

/*
Inert -
*/
func (r *RpTagHtml) Inert(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("inert", value)
	return r
}

/*
InputMode -
*/
func (r *RpTagHtml) InputMode(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("inputMode", value)
	return r
}

/*
Is -
*/
func (r *RpTagHtml) Is(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("is", value)
	return r
}

/*
ItemId -
*/
func (r *RpTagHtml) ItemId(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("itemId", value)
	return r
}

/*
ItemProp -
*/
func (r *RpTagHtml) ItemProp(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("itemProp", value)
	return r
}

/*
ItemRef -
*/
func (r *RpTagHtml) ItemRef(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("itemRef", value)
	return r
}

/*
ItemScope -
*/
func (r *RpTagHtml) ItemScope(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("itemScope", value)
	return r
}

/*
ItemType -
*/
func (r *RpTagHtml) ItemType(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("itemType", value)
	return r
}

/*
Lang -
*/
func (r *RpTagHtml) Lang(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("lang", value)
	return r
}

/*
Nonce -
*/
func (r *RpTagHtml) Nonce(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("nonce", value)
	return r
}

/*
Part -
*/
func (r *RpTagHtml) Part(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("part", value)
	return r
}

/*
Popover -
*/
func (r *RpTagHtml) Popover() *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("popover", "")
	return r
}

/*
Role -
*/
func (r *RpTagHtml) Role(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("role", value)
	return r
}

/*
Slot -
*/
func (r *RpTagHtml) Slot(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("slot", value)
	return r
}

/*
Spellcheck -
*/
func (r *RpTagHtml) Spellcheck(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("spellcheck", value)
	return r
}

/*
Style -
*/
func (r *RpTagHtml) Style(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("style", value)
	return r
}

/*
Tabindex -
*/
func (r *RpTagHtml) Tabindex(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("tabindex", value)
	return r
}

/*
Title -
*/
func (r *RpTagHtml) Title(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("title", value)
	return r
}

/*
Translate -
*/
func (r *RpTagHtml) Translate(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("translate", value)
	return r
}

/*
VirtualKeyBoardPolicy -
*/
func (r *RpTagHtml) VirtualKeyBoardPolicy(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("virtualKeyBoardPolicy", value)
	return r
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (r *RpTagHtml) AriaAtomic(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-atomic", value)
	return r
}

/*
AriaBusy -
*/
func (r *RpTagHtml) AriaBusy(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-busy", value)
	return r
}

/*
AriaControls -
*/
func (r *RpTagHtml) AriaControls(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-controls", value)
	return r
}

/*
AriaCurrent -
*/
func (r *RpTagHtml) AriaCurrent(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-current", value)
	return r
}

/*
AriaDescribedby -
*/
func (r *RpTagHtml) AriaDescribedby(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-describedby", value)
	return r
}

/*
AriaDescription -
*/
func (r *RpTagHtml) AriaDescription(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-description", value)
	return r
}

/*
AriaDetails -
*/
func (r *RpTagHtml) AriaDetails(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-details", value)
	return r
}

/*
AriaDisabled -
*/
func (r *RpTagHtml) AriaDisabled(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-disabled", value)
	return r
}

/*
AriaDropeffect -
*/
func (r *RpTagHtml) AriaDropeffect(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-dropeffect", value)
	return r
}

/*
AriaErrormessage -
*/
func (r *RpTagHtml) AriaErrormessage(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-errormessage", value)
	return r
}

/*
AriaFlowto -
*/
func (r *RpTagHtml) AriaFlowto(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-flowto", value)
	return r
}

/*
AriaGrabbed -
*/
func (r *RpTagHtml) AriaGrabbed(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-grabbed", value)
	return r
}

/*
AriaHaspopup -
*/
func (r *RpTagHtml) AriaHaspopup(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-haspopup", value)
	return r
}

/*
AriaHidden -
*/
func (r *RpTagHtml) AriaHidden(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-hidden", value)
	return r
}

/*
AriaInvalid -
*/
func (r *RpTagHtml) AriaInvalid(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-invalid", value)
	return r
}

/*
AriaKeyshortcuts -
*/
func (r *RpTagHtml) AriaKeyshortcuts(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-keyshortcuts", value)
	return r
}

/*
AriaLabel -
*/
func (r *RpTagHtml) AriaLabel(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-label", value)
	return r
}

/*
AriaLabelledby -
*/
func (r *RpTagHtml) AriaLabelledby(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-labelledby", value)
	return r
}

/*
AriaLive -
*/
func (r *RpTagHtml) AriaLive(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-live", value)
	return r
}

/*
AriaOwns -
*/
func (r *RpTagHtml) AriaOwns(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-owns", value)
	return r
}

/*
AriaRelevant -
*/
func (r *RpTagHtml) AriaRelevant(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-relevant", value)
	return r
}

/*
AriaRoledescription -
*/
func (r *RpTagHtml) AriaRoledescription(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("aria-roledescription", value)
	return r
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (r *RpTagHtml) Onabort(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onabort", value)
	return r
}

/*
Onautocomplete -
*/
func (r *RpTagHtml) Onautocomplete(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onautocomplete", value)
	return r
}

/*
Onautocompleteerror -
*/
func (r *RpTagHtml) Onautocompleteerror(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onautocompleteerror", value)
	return r
}

/*
Onblur -
*/
func (r *RpTagHtml) Onblur(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onblur", value)
	return r
}

/*
Oncancel -
*/
func (r *RpTagHtml) Oncancel(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("oncancel", value)
	return r
}

/*
Oncanplay -
*/
func (r *RpTagHtml) Oncanplay(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("oncanplay", value)
	return r
}

/*
Oncanplaythrough -
*/
func (r *RpTagHtml) Oncanplaythrough(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("oncanplaythrough", value)
	return r
}

/*
Onchange -
*/
func (r *RpTagHtml) Onchange(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onchange", value)
	return r
}

/*
Onclick -
*/
func (r *RpTagHtml) Onclick(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onclick", value)
	return r
}

/*
Onclose -
*/
func (r *RpTagHtml) Onclose(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onclose", value)
	return r
}

/*
Oncontextmenu -
*/
func (r *RpTagHtml) Oncontextmenu(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("oncontextmenu", value)
	return r
}

/*
Oncuechange -
*/
func (r *RpTagHtml) Oncuechange(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("oncuechange", value)
	return r
}

/*
Ondblclick -
*/
func (r *RpTagHtml) Ondblclick(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondblclick", value)
	return r
}

/*
Ondrag -
*/
func (r *RpTagHtml) Ondrag(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondrag", value)
	return r
}

/*
Ondragend -
*/
func (r *RpTagHtml) Ondragend(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondragend", value)
	return r
}

/*
Ondragenter -
*/
func (r *RpTagHtml) Ondragenter(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondragenter", value)
	return r
}

/*
Ondragleave -
*/
func (r *RpTagHtml) Ondragleave(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondragleave", value)
	return r
}

/*
Ondragover -
*/
func (r *RpTagHtml) Ondragover(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondragover", value)
	return r
}

/*
Ondragstart -
*/
func (r *RpTagHtml) Ondragstart(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondragstart", value)
	return r
}

/*
Ondrop -
*/
func (r *RpTagHtml) Ondrop(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondrop", value)
	return r
}

/*
Ondurationchange -
*/
func (r *RpTagHtml) Ondurationchange(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ondurationchange", value)
	return r
}

/*
Onemptied -
*/
func (r *RpTagHtml) Onemptied(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onemptied", value)
	return r
}

/*
Onended -
*/
func (r *RpTagHtml) Onended(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onended", value)
	return r
}

/*
Onfocus -
*/
func (r *RpTagHtml) Onfocus(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onfocus", value)
	return r
}

/*
Oninput -
*/
func (r *RpTagHtml) Oninput(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("oninput", value)
	return r
}

/*
Oninvalid -
*/
func (r *RpTagHtml) Oninvalid(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("oninvalid", value)
	return r
}

/*
Onkeydown -
*/
func (r *RpTagHtml) Onkeydown(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onkeydown", value)
	return r
}

/*
Onkeypress -
*/
func (r *RpTagHtml) Onkeypress(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onkeypress", value)
	return r
}

/*
Onkeyup -
*/
func (r *RpTagHtml) Onkeyup(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onkeyup", value)
	return r
}

/*
Onloadeddata -
*/
func (r *RpTagHtml) Onloadeddata(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onloadeddata", value)
	return r
}

/*
Onloadedmetadata -
*/
func (r *RpTagHtml) Onloadedmetadata(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onloadedmetadata", value)
	return r
}

/*
Onloadstart -
*/
func (r *RpTagHtml) Onloadstart(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onloadstart", value)
	return r
}

/*
Onmousedown -
*/
func (r *RpTagHtml) Onmousedown(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmousedown", value)
	return r
}

/*
Onmouseenter -
*/
func (r *RpTagHtml) Onmouseenter(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmouseenter", value)
	return r
}

/*
Onmouseleave -
*/
func (r *RpTagHtml) Onmouseleave(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmouseleave", value)
	return r
}

/*
Onmousemove -
*/
func (r *RpTagHtml) Onmousemove(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmousemove", value)
	return r
}

/*
Onmouseout -
*/
func (r *RpTagHtml) Onmouseout(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmouseout", value)
	return r
}

/*
Onmouseover -
*/
func (r *RpTagHtml) Onmouseover(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmouseover", value)
	return r
}

/*
Onmouseup -
*/
func (r *RpTagHtml) Onmouseup(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmouseup", value)
	return r
}

/*
Onmousewheel -
*/
func (r *RpTagHtml) Onmousewheel(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmousewheel", value)
	return r
}

/*
Onpause -
*/
func (r *RpTagHtml) Onpause(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onpause", value)
	return r
}

/*
Onplay -
*/
func (r *RpTagHtml) Onplay(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onplay", value)
	return r
}

/*
Onplaying -
*/
func (r *RpTagHtml) Onplaying(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onplaying", value)
	return r
}

/*
Onprogress -
*/
func (r *RpTagHtml) Onprogress(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onprogress", value)
	return r
}

/*
Onratechange -
*/
func (r *RpTagHtml) Onratechange(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onratechange", value)
	return r
}

/*
Onreset -
*/
func (r *RpTagHtml) Onreset(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onreset", value)
	return r
}

/*
Onscroll -
*/
func (r *RpTagHtml) Onscroll(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onscroll", value)
	return r
}

/*
Onseeked -
*/
func (r *RpTagHtml) Onseeked(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onseeked", value)
	return r
}

/*
Onseeking -
*/
func (r *RpTagHtml) Onseeking(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onseeking", value)
	return r
}

/*
Onselect -
*/
func (r *RpTagHtml) Onselect(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onselect", value)
	return r
}

/*
Onshow -
*/
func (r *RpTagHtml) Onshow(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onshow", value)
	return r
}

/*
Onsort -
*/
func (r *RpTagHtml) Onsort(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onsort", value)
	return r
}

/*
Onstalled -
*/
func (r *RpTagHtml) Onstalled(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onstalled", value)
	return r
}

/*
Onsubmit -
*/
func (r *RpTagHtml) Onsubmit(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onsubmit", value)
	return r
}

/*
Onsuspend -
*/
func (r *RpTagHtml) Onsuspend(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onsuspend", value)
	return r
}

/*
Ontimeupdate -
*/
func (r *RpTagHtml) Ontimeupdate(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ontimeupdate", value)
	return r
}

/*
Ontoggle -
*/
func (r *RpTagHtml) Ontoggle(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ontoggle", value)
	return r
}

/*
Onvolumechange -
*/
func (r *RpTagHtml) Onvolumechange(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onvolumechange", value)
	return r
}

/*
Onwaiting -
*/
func (r *RpTagHtml) Onwaiting(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onwaiting", value)
	return r
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (r *RpTagHtml) Onafterprint(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onafterprint", value)
	return r
}

/*
Onbeforeprint -
*/
func (r *RpTagHtml) Onbeforeprint(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onbeforeprint", value)
	return r
}

/*
Onbeforeunload -
*/
func (r *RpTagHtml) Onbeforeunload(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onbeforeunload", value)
	return r
}

/*
Onerror -
*/
func (r *RpTagHtml) Onerror(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onerror", value)
	return r
}

/*
Onhashchange -
*/
func (r *RpTagHtml) Onhashchange(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onhashchange", value)
	return r
}

/*
Onload -
*/
func (r *RpTagHtml) Onload(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onload", value)
	return r
}

/*
Onmessage -
*/
func (r *RpTagHtml) Onmessage(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onmessage", value)
	return r
}

/*
Onoffline -
*/
func (r *RpTagHtml) Onoffline(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onoffline", value)
	return r
}

/*
Ononline -
*/
func (r *RpTagHtml) Ononline(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("ononline", value)
	return r
}

/*
Onpagehide -
*/
func (r *RpTagHtml) Onpagehide(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onpagehide", value)
	return r
}

/*
Onpageshow -
*/
func (r *RpTagHtml) Onpageshow(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onpageshow", value)
	return r
}

/*
Onpopstate -
*/
func (r *RpTagHtml) Onpopstate(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onpopstate", value)
	return r
}

/*
Onresize -
*/
func (r *RpTagHtml) Onresize(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onresize", value)
	return r
}

/*
Onstorage -
*/
func (r *RpTagHtml) Onstorage(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onstorage", value)
	return r
}

/*
Onunload -
*/
func (r *RpTagHtml) Onunload(value string) *RpTagHtml {
	if r.attributes == nil {
		r.attributes = []*Attribute{}
	}
	r.registerAttribute("onunload", value)
	return r
}
