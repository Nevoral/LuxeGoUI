// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func FieldsetHtml(tags []any) *FieldsetTagHtml {
	node := &FieldsetTagHtml{
		tag: &tag{
			name:                "fieldset",
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

type FieldsetTagHtml struct {
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
func (f *FieldsetTagHtml) CustomAttribute(attributes ...*Attribute) *FieldsetTagHtml {
	f.registerAttributes(attributes...)
	return f
}

/*
Children - Method for nesting tags into parent tag
*/
func (f *FieldsetTagHtml) Children(tags ...any) *FieldsetTagHtml {
	return f.supportedChildrenCheck(tags)
}

func (f *FieldsetTagHtml) supportedChildrenCheck(tags []any) *FieldsetTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			f.registerChildren(TextContentSvg(val).getTag())
		case content:
			f.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				f.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return f
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Disabled -
*/
func (f *FieldsetTagHtml) Disabled() *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("disabled", "")
	return f
}

/*
Form -
*/
func (f *FieldsetTagHtml) Form(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("form", value)
	return f
}

/*
Name -
*/
func (f *FieldsetTagHtml) Name(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("name", value)
	return f
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (f *FieldsetTagHtml) AccessKey(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("accessKey", value)
	return f
}

/*
Aria -
*/
func (f *FieldsetTagHtml) Aria(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria", value)
	return f
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (f *FieldsetTagHtml) Autocapitalize(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("autocapitalize", value)
	return f
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (f *FieldsetTagHtml) Autofocus(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("autofocus", value)
	return f
}

/*
Class -
*/
func (f *FieldsetTagHtml) Class(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("class", value)
	return f
}

/*
Contenteditable -
*/
func (f *FieldsetTagHtml) Contenteditable(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("contenteditable", value)
	return f
}

/*
Data -
*/
func (f *FieldsetTagHtml) Data(name, value string) *FieldsetTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	f.registerAttribute(dataName, value)
	return f
}

/*
Dir -
*/
func (f *FieldsetTagHtml) Dir(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("dir", value)
	return f
}

/*
Draggable -
*/
func (f *FieldsetTagHtml) Draggable(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("draggable", value)
	return f
}

/*
EnterKeyHint -
*/
func (f *FieldsetTagHtml) EnterKeyHint(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("enterKeyHint", value)
	return f
}

/*
ExportParts -
*/
func (f *FieldsetTagHtml) ExportParts(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("exportParts", value)
	return f
}

/*
Hidden -
*/
func (f *FieldsetTagHtml) Hidden(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("hidden", value)
	return f
}

/*
Id -
*/
func (f *FieldsetTagHtml) Id(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("id", value)
	return f
}

/*
Inert -
*/
func (f *FieldsetTagHtml) Inert(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("inert", value)
	return f
}

/*
InputMode -
*/
func (f *FieldsetTagHtml) InputMode(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("inputMode", value)
	return f
}

/*
Is -
*/
func (f *FieldsetTagHtml) Is(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("is", value)
	return f
}

/*
ItemId -
*/
func (f *FieldsetTagHtml) ItemId(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemId", value)
	return f
}

/*
ItemProp -
*/
func (f *FieldsetTagHtml) ItemProp(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemProp", value)
	return f
}

/*
ItemRef -
*/
func (f *FieldsetTagHtml) ItemRef(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemRef", value)
	return f
}

/*
ItemScope -
*/
func (f *FieldsetTagHtml) ItemScope(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemScope", value)
	return f
}

/*
ItemType -
*/
func (f *FieldsetTagHtml) ItemType(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemType", value)
	return f
}

/*
Lang -
*/
func (f *FieldsetTagHtml) Lang(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("lang", value)
	return f
}

/*
Nonce -
*/
func (f *FieldsetTagHtml) Nonce(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("nonce", value)
	return f
}

/*
Part -
*/
func (f *FieldsetTagHtml) Part(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("part", value)
	return f
}

/*
Popover -
*/
func (f *FieldsetTagHtml) Popover() *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("popover", "")
	return f
}

/*
Role -
*/
func (f *FieldsetTagHtml) Role(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("role", value)
	return f
}

/*
Slot -
*/
func (f *FieldsetTagHtml) Slot(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("slot", value)
	return f
}

/*
Spellcheck -
*/
func (f *FieldsetTagHtml) Spellcheck(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("spellcheck", value)
	return f
}

/*
Style -
*/
func (f *FieldsetTagHtml) Style(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("style", value)
	return f
}

/*
Tabindex -
*/
func (f *FieldsetTagHtml) Tabindex(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("tabindex", value)
	return f
}

/*
Title -
*/
func (f *FieldsetTagHtml) Title(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("title", value)
	return f
}

/*
Translate -
*/
func (f *FieldsetTagHtml) Translate(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("translate", value)
	return f
}

/*
VirtualKeyBoardPolicy -
*/
func (f *FieldsetTagHtml) VirtualKeyBoardPolicy(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("virtualKeyBoardPolicy", value)
	return f
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (f *FieldsetTagHtml) AriaAtomic(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-atomic", value)
	return f
}

/*
AriaBusy -
*/
func (f *FieldsetTagHtml) AriaBusy(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-busy", value)
	return f
}

/*
AriaControls -
*/
func (f *FieldsetTagHtml) AriaControls(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-controls", value)
	return f
}

/*
AriaCurrent -
*/
func (f *FieldsetTagHtml) AriaCurrent(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-current", value)
	return f
}

/*
AriaDescribedby -
*/
func (f *FieldsetTagHtml) AriaDescribedby(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-describedby", value)
	return f
}

/*
AriaDescription -
*/
func (f *FieldsetTagHtml) AriaDescription(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-description", value)
	return f
}

/*
AriaDetails -
*/
func (f *FieldsetTagHtml) AriaDetails(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-details", value)
	return f
}

/*
AriaDisabled -
*/
func (f *FieldsetTagHtml) AriaDisabled(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-disabled", value)
	return f
}

/*
AriaDropeffect -
*/
func (f *FieldsetTagHtml) AriaDropeffect(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-dropeffect", value)
	return f
}

/*
AriaErrormessage -
*/
func (f *FieldsetTagHtml) AriaErrormessage(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-errormessage", value)
	return f
}

/*
AriaFlowto -
*/
func (f *FieldsetTagHtml) AriaFlowto(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-flowto", value)
	return f
}

/*
AriaGrabbed -
*/
func (f *FieldsetTagHtml) AriaGrabbed(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-grabbed", value)
	return f
}

/*
AriaHaspopup -
*/
func (f *FieldsetTagHtml) AriaHaspopup(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-haspopup", value)
	return f
}

/*
AriaHidden -
*/
func (f *FieldsetTagHtml) AriaHidden(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-hidden", value)
	return f
}

/*
AriaInvalid -
*/
func (f *FieldsetTagHtml) AriaInvalid(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-invalid", value)
	return f
}

/*
AriaKeyshortcuts -
*/
func (f *FieldsetTagHtml) AriaKeyshortcuts(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-keyshortcuts", value)
	return f
}

/*
AriaLabel -
*/
func (f *FieldsetTagHtml) AriaLabel(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-label", value)
	return f
}

/*
AriaLabelledby -
*/
func (f *FieldsetTagHtml) AriaLabelledby(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-labelledby", value)
	return f
}

/*
AriaLive -
*/
func (f *FieldsetTagHtml) AriaLive(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-live", value)
	return f
}

/*
AriaOwns -
*/
func (f *FieldsetTagHtml) AriaOwns(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-owns", value)
	return f
}

/*
AriaRelevant -
*/
func (f *FieldsetTagHtml) AriaRelevant(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-relevant", value)
	return f
}

/*
AriaRoledescription -
*/
func (f *FieldsetTagHtml) AriaRoledescription(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-roledescription", value)
	return f
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (f *FieldsetTagHtml) Onabort(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onabort", value)
	return f
}

/*
Onautocomplete -
*/
func (f *FieldsetTagHtml) Onautocomplete(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onautocomplete", value)
	return f
}

/*
Onautocompleteerror -
*/
func (f *FieldsetTagHtml) Onautocompleteerror(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onautocompleteerror", value)
	return f
}

/*
Onblur -
*/
func (f *FieldsetTagHtml) Onblur(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onblur", value)
	return f
}

/*
Oncancel -
*/
func (f *FieldsetTagHtml) Oncancel(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncancel", value)
	return f
}

/*
Oncanplay -
*/
func (f *FieldsetTagHtml) Oncanplay(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncanplay", value)
	return f
}

/*
Oncanplaythrough -
*/
func (f *FieldsetTagHtml) Oncanplaythrough(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncanplaythrough", value)
	return f
}

/*
Onchange -
*/
func (f *FieldsetTagHtml) Onchange(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onchange", value)
	return f
}

/*
Onclick -
*/
func (f *FieldsetTagHtml) Onclick(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onclick", value)
	return f
}

/*
Onclose -
*/
func (f *FieldsetTagHtml) Onclose(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onclose", value)
	return f
}

/*
Oncontextmenu -
*/
func (f *FieldsetTagHtml) Oncontextmenu(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncontextmenu", value)
	return f
}

/*
Oncuechange -
*/
func (f *FieldsetTagHtml) Oncuechange(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncuechange", value)
	return f
}

/*
Ondblclick -
*/
func (f *FieldsetTagHtml) Ondblclick(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondblclick", value)
	return f
}

/*
Ondrag -
*/
func (f *FieldsetTagHtml) Ondrag(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondrag", value)
	return f
}

/*
Ondragend -
*/
func (f *FieldsetTagHtml) Ondragend(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragend", value)
	return f
}

/*
Ondragenter -
*/
func (f *FieldsetTagHtml) Ondragenter(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragenter", value)
	return f
}

/*
Ondragleave -
*/
func (f *FieldsetTagHtml) Ondragleave(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragleave", value)
	return f
}

/*
Ondragover -
*/
func (f *FieldsetTagHtml) Ondragover(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragover", value)
	return f
}

/*
Ondragstart -
*/
func (f *FieldsetTagHtml) Ondragstart(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragstart", value)
	return f
}

/*
Ondrop -
*/
func (f *FieldsetTagHtml) Ondrop(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondrop", value)
	return f
}

/*
Ondurationchange -
*/
func (f *FieldsetTagHtml) Ondurationchange(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondurationchange", value)
	return f
}

/*
Onemptied -
*/
func (f *FieldsetTagHtml) Onemptied(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onemptied", value)
	return f
}

/*
Onended -
*/
func (f *FieldsetTagHtml) Onended(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onended", value)
	return f
}

/*
Onfocus -
*/
func (f *FieldsetTagHtml) Onfocus(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onfocus", value)
	return f
}

/*
Oninput -
*/
func (f *FieldsetTagHtml) Oninput(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oninput", value)
	return f
}

/*
Oninvalid -
*/
func (f *FieldsetTagHtml) Oninvalid(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oninvalid", value)
	return f
}

/*
Onkeydown -
*/
func (f *FieldsetTagHtml) Onkeydown(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onkeydown", value)
	return f
}

/*
Onkeypress -
*/
func (f *FieldsetTagHtml) Onkeypress(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onkeypress", value)
	return f
}

/*
Onkeyup -
*/
func (f *FieldsetTagHtml) Onkeyup(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onkeyup", value)
	return f
}

/*
Onloadeddata -
*/
func (f *FieldsetTagHtml) Onloadeddata(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onloadeddata", value)
	return f
}

/*
Onloadedmetadata -
*/
func (f *FieldsetTagHtml) Onloadedmetadata(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onloadedmetadata", value)
	return f
}

/*
Onloadstart -
*/
func (f *FieldsetTagHtml) Onloadstart(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onloadstart", value)
	return f
}

/*
Onmousedown -
*/
func (f *FieldsetTagHtml) Onmousedown(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmousedown", value)
	return f
}

/*
Onmouseenter -
*/
func (f *FieldsetTagHtml) Onmouseenter(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseenter", value)
	return f
}

/*
Onmouseleave -
*/
func (f *FieldsetTagHtml) Onmouseleave(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseleave", value)
	return f
}

/*
Onmousemove -
*/
func (f *FieldsetTagHtml) Onmousemove(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmousemove", value)
	return f
}

/*
Onmouseout -
*/
func (f *FieldsetTagHtml) Onmouseout(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseout", value)
	return f
}

/*
Onmouseover -
*/
func (f *FieldsetTagHtml) Onmouseover(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseover", value)
	return f
}

/*
Onmouseup -
*/
func (f *FieldsetTagHtml) Onmouseup(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseup", value)
	return f
}

/*
Onmousewheel -
*/
func (f *FieldsetTagHtml) Onmousewheel(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmousewheel", value)
	return f
}

/*
Onpause -
*/
func (f *FieldsetTagHtml) Onpause(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpause", value)
	return f
}

/*
Onplay -
*/
func (f *FieldsetTagHtml) Onplay(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onplay", value)
	return f
}

/*
Onplaying -
*/
func (f *FieldsetTagHtml) Onplaying(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onplaying", value)
	return f
}

/*
Onprogress -
*/
func (f *FieldsetTagHtml) Onprogress(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onprogress", value)
	return f
}

/*
Onratechange -
*/
func (f *FieldsetTagHtml) Onratechange(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onratechange", value)
	return f
}

/*
Onreset -
*/
func (f *FieldsetTagHtml) Onreset(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onreset", value)
	return f
}

/*
Onscroll -
*/
func (f *FieldsetTagHtml) Onscroll(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onscroll", value)
	return f
}

/*
Onseeked -
*/
func (f *FieldsetTagHtml) Onseeked(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onseeked", value)
	return f
}

/*
Onseeking -
*/
func (f *FieldsetTagHtml) Onseeking(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onseeking", value)
	return f
}

/*
Onselect -
*/
func (f *FieldsetTagHtml) Onselect(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onselect", value)
	return f
}

/*
Onshow -
*/
func (f *FieldsetTagHtml) Onshow(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onshow", value)
	return f
}

/*
Onsort -
*/
func (f *FieldsetTagHtml) Onsort(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onsort", value)
	return f
}

/*
Onstalled -
*/
func (f *FieldsetTagHtml) Onstalled(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onstalled", value)
	return f
}

/*
Onsubmit -
*/
func (f *FieldsetTagHtml) Onsubmit(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onsubmit", value)
	return f
}

/*
Onsuspend -
*/
func (f *FieldsetTagHtml) Onsuspend(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onsuspend", value)
	return f
}

/*
Ontimeupdate -
*/
func (f *FieldsetTagHtml) Ontimeupdate(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ontimeupdate", value)
	return f
}

/*
Ontoggle -
*/
func (f *FieldsetTagHtml) Ontoggle(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ontoggle", value)
	return f
}

/*
Onvolumechange -
*/
func (f *FieldsetTagHtml) Onvolumechange(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onvolumechange", value)
	return f
}

/*
Onwaiting -
*/
func (f *FieldsetTagHtml) Onwaiting(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onwaiting", value)
	return f
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (f *FieldsetTagHtml) Onafterprint(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onafterprint", value)
	return f
}

/*
Onbeforeprint -
*/
func (f *FieldsetTagHtml) Onbeforeprint(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onbeforeprint", value)
	return f
}

/*
Onbeforeunload -
*/
func (f *FieldsetTagHtml) Onbeforeunload(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onbeforeunload", value)
	return f
}

/*
Onerror -
*/
func (f *FieldsetTagHtml) Onerror(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onerror", value)
	return f
}

/*
Onhashchange -
*/
func (f *FieldsetTagHtml) Onhashchange(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onhashchange", value)
	return f
}

/*
Onload -
*/
func (f *FieldsetTagHtml) Onload(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onload", value)
	return f
}

/*
Onmessage -
*/
func (f *FieldsetTagHtml) Onmessage(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmessage", value)
	return f
}

/*
Onoffline -
*/
func (f *FieldsetTagHtml) Onoffline(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onoffline", value)
	return f
}

/*
Ononline -
*/
func (f *FieldsetTagHtml) Ononline(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ononline", value)
	return f
}

/*
Onpagehide -
*/
func (f *FieldsetTagHtml) Onpagehide(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpagehide", value)
	return f
}

/*
Onpageshow -
*/
func (f *FieldsetTagHtml) Onpageshow(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpageshow", value)
	return f
}

/*
Onpopstate -
*/
func (f *FieldsetTagHtml) Onpopstate(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpopstate", value)
	return f
}

/*
Onresize -
*/
func (f *FieldsetTagHtml) Onresize(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onresize", value)
	return f
}

/*
Onstorage -
*/
func (f *FieldsetTagHtml) Onstorage(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onstorage", value)
	return f
}

/*
Onunload -
*/
func (f *FieldsetTagHtml) Onunload(value string) *FieldsetTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onunload", value)
	return f
}
