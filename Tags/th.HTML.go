// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func ThHtml(tags []any) *ThTagHtml {
	node := &ThTagHtml{
		tag: &tag{
			name:                "th",
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

type ThTagHtml struct {
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
func (t *ThTagHtml) CustomAttribute(attributes ...*Attribute) *ThTagHtml {
	t.registerAttributes(attributes...)
	return t
}

/*
Children - Method for nesting tags into parent tag
*/
func (t *ThTagHtml) Children(tags ...any) *ThTagHtml {
	return t.supportedChildrenCheck(tags)
}

func (t *ThTagHtml) supportedChildrenCheck(tags []any) *ThTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			t.registerChildren(TextContentSvg(val).getTag())
		case content:
			t.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				t.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return t
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Abbr -
*/
func (t *ThTagHtml) Abbr(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("abbr", value)
	return t
}

/*
Colspan -
*/
func (t *ThTagHtml) Colspan(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("colspan", value)
	return t
}

/*
Headers -
*/
func (t *ThTagHtml) Headers(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("headers", value)
	return t
}

/*
Rowspan - Specify how many rows a cell should span in a table.
Specify how many rows a cell should span in a table.
*/
func (t *ThTagHtml) Rowspan(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("rowspan", value)
	return t
}

/*
Scope -
*/
func (t *ThTagHtml) Scope(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("scope", value)
	return t
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (t *ThTagHtml) AccessKey(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("accessKey", value)
	return t
}

/*
Aria -
*/
func (t *ThTagHtml) Aria(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria", value)
	return t
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (t *ThTagHtml) Autocapitalize(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("autocapitalize", value)
	return t
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (t *ThTagHtml) Autofocus(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("autofocus", value)
	return t
}

/*
Class -
*/
func (t *ThTagHtml) Class(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("class", value)
	return t
}

/*
Contenteditable -
*/
func (t *ThTagHtml) Contenteditable(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("contenteditable", value)
	return t
}

/*
Data -
*/
func (t *ThTagHtml) Data(name, value string) *ThTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	t.registerAttribute(dataName, value)
	return t
}

/*
Dir -
*/
func (t *ThTagHtml) Dir(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("dir", value)
	return t
}

/*
Draggable -
*/
func (t *ThTagHtml) Draggable(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("draggable", value)
	return t
}

/*
EnterKeyHint -
*/
func (t *ThTagHtml) EnterKeyHint(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("enterKeyHint", value)
	return t
}

/*
ExportParts -
*/
func (t *ThTagHtml) ExportParts(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("exportParts", value)
	return t
}

/*
Hidden -
*/
func (t *ThTagHtml) Hidden(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("hidden", value)
	return t
}

/*
Id -
*/
func (t *ThTagHtml) Id(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("id", value)
	return t
}

/*
Inert -
*/
func (t *ThTagHtml) Inert(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("inert", value)
	return t
}

/*
InputMode -
*/
func (t *ThTagHtml) InputMode(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("inputMode", value)
	return t
}

/*
Is -
*/
func (t *ThTagHtml) Is(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("is", value)
	return t
}

/*
ItemId -
*/
func (t *ThTagHtml) ItemId(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemId", value)
	return t
}

/*
ItemProp -
*/
func (t *ThTagHtml) ItemProp(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemProp", value)
	return t
}

/*
ItemRef -
*/
func (t *ThTagHtml) ItemRef(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemRef", value)
	return t
}

/*
ItemScope -
*/
func (t *ThTagHtml) ItemScope(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemScope", value)
	return t
}

/*
ItemType -
*/
func (t *ThTagHtml) ItemType(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemType", value)
	return t
}

/*
Lang -
*/
func (t *ThTagHtml) Lang(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("lang", value)
	return t
}

/*
Nonce -
*/
func (t *ThTagHtml) Nonce(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("nonce", value)
	return t
}

/*
Part -
*/
func (t *ThTagHtml) Part(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("part", value)
	return t
}

/*
Popover -
*/
func (t *ThTagHtml) Popover() *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("popover", "")
	return t
}

/*
Role -
*/
func (t *ThTagHtml) Role(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("role", value)
	return t
}

/*
Slot -
*/
func (t *ThTagHtml) Slot(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("slot", value)
	return t
}

/*
Spellcheck -
*/
func (t *ThTagHtml) Spellcheck(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("spellcheck", value)
	return t
}

/*
Style -
*/
func (t *ThTagHtml) Style(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("style", value)
	return t
}

/*
Tabindex -
*/
func (t *ThTagHtml) Tabindex(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("tabindex", value)
	return t
}

/*
Title -
*/
func (t *ThTagHtml) Title(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("title", value)
	return t
}

/*
Translate -
*/
func (t *ThTagHtml) Translate(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("translate", value)
	return t
}

/*
VirtualKeyBoardPolicy -
*/
func (t *ThTagHtml) VirtualKeyBoardPolicy(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("virtualKeyBoardPolicy", value)
	return t
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (t *ThTagHtml) AriaAtomic(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-atomic", value)
	return t
}

/*
AriaBusy -
*/
func (t *ThTagHtml) AriaBusy(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-busy", value)
	return t
}

/*
AriaControls -
*/
func (t *ThTagHtml) AriaControls(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-controls", value)
	return t
}

/*
AriaCurrent -
*/
func (t *ThTagHtml) AriaCurrent(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-current", value)
	return t
}

/*
AriaDescribedby -
*/
func (t *ThTagHtml) AriaDescribedby(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-describedby", value)
	return t
}

/*
AriaDescription -
*/
func (t *ThTagHtml) AriaDescription(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-description", value)
	return t
}

/*
AriaDetails -
*/
func (t *ThTagHtml) AriaDetails(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-details", value)
	return t
}

/*
AriaDisabled -
*/
func (t *ThTagHtml) AriaDisabled(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-disabled", value)
	return t
}

/*
AriaDropeffect -
*/
func (t *ThTagHtml) AriaDropeffect(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-dropeffect", value)
	return t
}

/*
AriaErrormessage -
*/
func (t *ThTagHtml) AriaErrormessage(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-errormessage", value)
	return t
}

/*
AriaFlowto -
*/
func (t *ThTagHtml) AriaFlowto(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-flowto", value)
	return t
}

/*
AriaGrabbed -
*/
func (t *ThTagHtml) AriaGrabbed(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-grabbed", value)
	return t
}

/*
AriaHaspopup -
*/
func (t *ThTagHtml) AriaHaspopup(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-haspopup", value)
	return t
}

/*
AriaHidden -
*/
func (t *ThTagHtml) AriaHidden(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-hidden", value)
	return t
}

/*
AriaInvalid -
*/
func (t *ThTagHtml) AriaInvalid(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-invalid", value)
	return t
}

/*
AriaKeyshortcuts -
*/
func (t *ThTagHtml) AriaKeyshortcuts(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-keyshortcuts", value)
	return t
}

/*
AriaLabel -
*/
func (t *ThTagHtml) AriaLabel(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-label", value)
	return t
}

/*
AriaLabelledby -
*/
func (t *ThTagHtml) AriaLabelledby(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-labelledby", value)
	return t
}

/*
AriaLive -
*/
func (t *ThTagHtml) AriaLive(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-live", value)
	return t
}

/*
AriaOwns -
*/
func (t *ThTagHtml) AriaOwns(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-owns", value)
	return t
}

/*
AriaRelevant -
*/
func (t *ThTagHtml) AriaRelevant(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-relevant", value)
	return t
}

/*
AriaRoledescription -
*/
func (t *ThTagHtml) AriaRoledescription(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-roledescription", value)
	return t
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (t *ThTagHtml) Onabort(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onabort", value)
	return t
}

/*
Onautocomplete -
*/
func (t *ThTagHtml) Onautocomplete(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onautocomplete", value)
	return t
}

/*
Onautocompleteerror -
*/
func (t *ThTagHtml) Onautocompleteerror(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onautocompleteerror", value)
	return t
}

/*
Onblur -
*/
func (t *ThTagHtml) Onblur(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onblur", value)
	return t
}

/*
Oncancel -
*/
func (t *ThTagHtml) Oncancel(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncancel", value)
	return t
}

/*
Oncanplay -
*/
func (t *ThTagHtml) Oncanplay(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncanplay", value)
	return t
}

/*
Oncanplaythrough -
*/
func (t *ThTagHtml) Oncanplaythrough(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncanplaythrough", value)
	return t
}

/*
Onchange -
*/
func (t *ThTagHtml) Onchange(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onchange", value)
	return t
}

/*
Onclick -
*/
func (t *ThTagHtml) Onclick(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onclick", value)
	return t
}

/*
Onclose -
*/
func (t *ThTagHtml) Onclose(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onclose", value)
	return t
}

/*
Oncontextmenu -
*/
func (t *ThTagHtml) Oncontextmenu(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncontextmenu", value)
	return t
}

/*
Oncuechange -
*/
func (t *ThTagHtml) Oncuechange(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncuechange", value)
	return t
}

/*
Ondblclick -
*/
func (t *ThTagHtml) Ondblclick(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondblclick", value)
	return t
}

/*
Ondrag -
*/
func (t *ThTagHtml) Ondrag(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondrag", value)
	return t
}

/*
Ondragend -
*/
func (t *ThTagHtml) Ondragend(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragend", value)
	return t
}

/*
Ondragenter -
*/
func (t *ThTagHtml) Ondragenter(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragenter", value)
	return t
}

/*
Ondragleave -
*/
func (t *ThTagHtml) Ondragleave(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragleave", value)
	return t
}

/*
Ondragover -
*/
func (t *ThTagHtml) Ondragover(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragover", value)
	return t
}

/*
Ondragstart -
*/
func (t *ThTagHtml) Ondragstart(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragstart", value)
	return t
}

/*
Ondrop -
*/
func (t *ThTagHtml) Ondrop(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondrop", value)
	return t
}

/*
Ondurationchange -
*/
func (t *ThTagHtml) Ondurationchange(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondurationchange", value)
	return t
}

/*
Onemptied -
*/
func (t *ThTagHtml) Onemptied(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onemptied", value)
	return t
}

/*
Onended -
*/
func (t *ThTagHtml) Onended(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onended", value)
	return t
}

/*
Onfocus -
*/
func (t *ThTagHtml) Onfocus(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onfocus", value)
	return t
}

/*
Oninput -
*/
func (t *ThTagHtml) Oninput(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oninput", value)
	return t
}

/*
Oninvalid -
*/
func (t *ThTagHtml) Oninvalid(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oninvalid", value)
	return t
}

/*
Onkeydown -
*/
func (t *ThTagHtml) Onkeydown(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeydown", value)
	return t
}

/*
Onkeypress -
*/
func (t *ThTagHtml) Onkeypress(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeypress", value)
	return t
}

/*
Onkeyup -
*/
func (t *ThTagHtml) Onkeyup(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeyup", value)
	return t
}

/*
Onloadeddata -
*/
func (t *ThTagHtml) Onloadeddata(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadeddata", value)
	return t
}

/*
Onloadedmetadata -
*/
func (t *ThTagHtml) Onloadedmetadata(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadedmetadata", value)
	return t
}

/*
Onloadstart -
*/
func (t *ThTagHtml) Onloadstart(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadstart", value)
	return t
}

/*
Onmousedown -
*/
func (t *ThTagHtml) Onmousedown(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousedown", value)
	return t
}

/*
Onmouseenter -
*/
func (t *ThTagHtml) Onmouseenter(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseenter", value)
	return t
}

/*
Onmouseleave -
*/
func (t *ThTagHtml) Onmouseleave(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseleave", value)
	return t
}

/*
Onmousemove -
*/
func (t *ThTagHtml) Onmousemove(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousemove", value)
	return t
}

/*
Onmouseout -
*/
func (t *ThTagHtml) Onmouseout(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseout", value)
	return t
}

/*
Onmouseover -
*/
func (t *ThTagHtml) Onmouseover(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseover", value)
	return t
}

/*
Onmouseup -
*/
func (t *ThTagHtml) Onmouseup(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseup", value)
	return t
}

/*
Onmousewheel -
*/
func (t *ThTagHtml) Onmousewheel(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousewheel", value)
	return t
}

/*
Onpause -
*/
func (t *ThTagHtml) Onpause(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpause", value)
	return t
}

/*
Onplay -
*/
func (t *ThTagHtml) Onplay(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onplay", value)
	return t
}

/*
Onplaying -
*/
func (t *ThTagHtml) Onplaying(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onplaying", value)
	return t
}

/*
Onprogress -
*/
func (t *ThTagHtml) Onprogress(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onprogress", value)
	return t
}

/*
Onratechange -
*/
func (t *ThTagHtml) Onratechange(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onratechange", value)
	return t
}

/*
Onreset -
*/
func (t *ThTagHtml) Onreset(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onreset", value)
	return t
}

/*
Onscroll -
*/
func (t *ThTagHtml) Onscroll(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onscroll", value)
	return t
}

/*
Onseeked -
*/
func (t *ThTagHtml) Onseeked(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onseeked", value)
	return t
}

/*
Onseeking -
*/
func (t *ThTagHtml) Onseeking(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onseeking", value)
	return t
}

/*
Onselect -
*/
func (t *ThTagHtml) Onselect(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onselect", value)
	return t
}

/*
Onshow -
*/
func (t *ThTagHtml) Onshow(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onshow", value)
	return t
}

/*
Onsort -
*/
func (t *ThTagHtml) Onsort(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsort", value)
	return t
}

/*
Onstalled -
*/
func (t *ThTagHtml) Onstalled(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onstalled", value)
	return t
}

/*
Onsubmit -
*/
func (t *ThTagHtml) Onsubmit(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsubmit", value)
	return t
}

/*
Onsuspend -
*/
func (t *ThTagHtml) Onsuspend(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsuspend", value)
	return t
}

/*
Ontimeupdate -
*/
func (t *ThTagHtml) Ontimeupdate(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ontimeupdate", value)
	return t
}

/*
Ontoggle -
*/
func (t *ThTagHtml) Ontoggle(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ontoggle", value)
	return t
}

/*
Onvolumechange -
*/
func (t *ThTagHtml) Onvolumechange(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onvolumechange", value)
	return t
}

/*
Onwaiting -
*/
func (t *ThTagHtml) Onwaiting(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onwaiting", value)
	return t
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (t *ThTagHtml) Onafterprint(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onafterprint", value)
	return t
}

/*
Onbeforeprint -
*/
func (t *ThTagHtml) Onbeforeprint(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onbeforeprint", value)
	return t
}

/*
Onbeforeunload -
*/
func (t *ThTagHtml) Onbeforeunload(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onbeforeunload", value)
	return t
}

/*
Onerror -
*/
func (t *ThTagHtml) Onerror(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onerror", value)
	return t
}

/*
Onhashchange -
*/
func (t *ThTagHtml) Onhashchange(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onhashchange", value)
	return t
}

/*
Onload -
*/
func (t *ThTagHtml) Onload(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onload", value)
	return t
}

/*
Onmessage -
*/
func (t *ThTagHtml) Onmessage(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmessage", value)
	return t
}

/*
Onoffline -
*/
func (t *ThTagHtml) Onoffline(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onoffline", value)
	return t
}

/*
Ononline -
*/
func (t *ThTagHtml) Ononline(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ononline", value)
	return t
}

/*
Onpagehide -
*/
func (t *ThTagHtml) Onpagehide(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpagehide", value)
	return t
}

/*
Onpageshow -
*/
func (t *ThTagHtml) Onpageshow(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpageshow", value)
	return t
}

/*
Onpopstate -
*/
func (t *ThTagHtml) Onpopstate(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpopstate", value)
	return t
}

/*
Onresize -
*/
func (t *ThTagHtml) Onresize(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onresize", value)
	return t
}

/*
Onstorage -
*/
func (t *ThTagHtml) Onstorage(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onstorage", value)
	return t
}

/*
Onunload -
*/
func (t *ThTagHtml) Onunload(value string) *ThTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onunload", value)
	return t
}
