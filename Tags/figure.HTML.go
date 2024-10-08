// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func FigureHtml(tags []any) *FigureTagHtml {
	node := &FigureTagHtml{
		tag: &tag{
			name:                "figure",
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

type FigureTagHtml struct {
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
func (f *FigureTagHtml) CustomAttribute(attributes ...*Attribute) *FigureTagHtml {
	f.registerAttributes(attributes...)
	return f
}

/*
Children - Method for nesting tags into parent tag
*/
func (f *FigureTagHtml) Children(tags ...any) *FigureTagHtml {
	return f.supportedChildrenCheck(tags)
}

func (f *FigureTagHtml) supportedChildrenCheck(tags []any) *FigureTagHtml {
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
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (f *FigureTagHtml) AccessKey(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("accessKey", value)
	return f
}

/*
Aria -
*/
func (f *FigureTagHtml) Aria(value string) *FigureTagHtml {
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
func (f *FigureTagHtml) Autocapitalize(value string) *FigureTagHtml {
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
func (f *FigureTagHtml) Autofocus(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("autofocus", value)
	return f
}

/*
Class -
*/
func (f *FigureTagHtml) Class(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("class", value)
	return f
}

/*
Contenteditable -
*/
func (f *FigureTagHtml) Contenteditable(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("contenteditable", value)
	return f
}

/*
Data -
*/
func (f *FigureTagHtml) Data(name, value string) *FigureTagHtml {
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
func (f *FigureTagHtml) Dir(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("dir", value)
	return f
}

/*
Draggable -
*/
func (f *FigureTagHtml) Draggable(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("draggable", value)
	return f
}

/*
EnterKeyHint -
*/
func (f *FigureTagHtml) EnterKeyHint(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("enterKeyHint", value)
	return f
}

/*
ExportParts -
*/
func (f *FigureTagHtml) ExportParts(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("exportParts", value)
	return f
}

/*
Hidden -
*/
func (f *FigureTagHtml) Hidden(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("hidden", value)
	return f
}

/*
Id -
*/
func (f *FigureTagHtml) Id(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("id", value)
	return f
}

/*
Inert -
*/
func (f *FigureTagHtml) Inert(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("inert", value)
	return f
}

/*
InputMode -
*/
func (f *FigureTagHtml) InputMode(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("inputMode", value)
	return f
}

/*
Is -
*/
func (f *FigureTagHtml) Is(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("is", value)
	return f
}

/*
ItemId -
*/
func (f *FigureTagHtml) ItemId(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemId", value)
	return f
}

/*
ItemProp -
*/
func (f *FigureTagHtml) ItemProp(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemProp", value)
	return f
}

/*
ItemRef -
*/
func (f *FigureTagHtml) ItemRef(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemRef", value)
	return f
}

/*
ItemScope -
*/
func (f *FigureTagHtml) ItemScope(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemScope", value)
	return f
}

/*
ItemType -
*/
func (f *FigureTagHtml) ItemType(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("itemType", value)
	return f
}

/*
Lang -
*/
func (f *FigureTagHtml) Lang(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("lang", value)
	return f
}

/*
Nonce -
*/
func (f *FigureTagHtml) Nonce(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("nonce", value)
	return f
}

/*
Part -
*/
func (f *FigureTagHtml) Part(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("part", value)
	return f
}

/*
Popover -
*/
func (f *FigureTagHtml) Popover() *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("popover", "")
	return f
}

/*
Role -
*/
func (f *FigureTagHtml) Role(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("role", value)
	return f
}

/*
Slot -
*/
func (f *FigureTagHtml) Slot(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("slot", value)
	return f
}

/*
Spellcheck -
*/
func (f *FigureTagHtml) Spellcheck(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("spellcheck", value)
	return f
}

/*
Style -
*/
func (f *FigureTagHtml) Style(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("style", value)
	return f
}

/*
Tabindex -
*/
func (f *FigureTagHtml) Tabindex(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("tabindex", value)
	return f
}

/*
Title -
*/
func (f *FigureTagHtml) Title(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("title", value)
	return f
}

/*
Translate -
*/
func (f *FigureTagHtml) Translate(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("translate", value)
	return f
}

/*
VirtualKeyBoardPolicy -
*/
func (f *FigureTagHtml) VirtualKeyBoardPolicy(value string) *FigureTagHtml {
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
func (f *FigureTagHtml) AriaAtomic(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-atomic", value)
	return f
}

/*
AriaBusy -
*/
func (f *FigureTagHtml) AriaBusy(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-busy", value)
	return f
}

/*
AriaControls -
*/
func (f *FigureTagHtml) AriaControls(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-controls", value)
	return f
}

/*
AriaCurrent -
*/
func (f *FigureTagHtml) AriaCurrent(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-current", value)
	return f
}

/*
AriaDescribedby -
*/
func (f *FigureTagHtml) AriaDescribedby(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-describedby", value)
	return f
}

/*
AriaDescription -
*/
func (f *FigureTagHtml) AriaDescription(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-description", value)
	return f
}

/*
AriaDetails -
*/
func (f *FigureTagHtml) AriaDetails(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-details", value)
	return f
}

/*
AriaDisabled -
*/
func (f *FigureTagHtml) AriaDisabled(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-disabled", value)
	return f
}

/*
AriaDropeffect -
*/
func (f *FigureTagHtml) AriaDropeffect(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-dropeffect", value)
	return f
}

/*
AriaErrormessage -
*/
func (f *FigureTagHtml) AriaErrormessage(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-errormessage", value)
	return f
}

/*
AriaFlowto -
*/
func (f *FigureTagHtml) AriaFlowto(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-flowto", value)
	return f
}

/*
AriaGrabbed -
*/
func (f *FigureTagHtml) AriaGrabbed(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-grabbed", value)
	return f
}

/*
AriaHaspopup -
*/
func (f *FigureTagHtml) AriaHaspopup(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-haspopup", value)
	return f
}

/*
AriaHidden -
*/
func (f *FigureTagHtml) AriaHidden(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-hidden", value)
	return f
}

/*
AriaInvalid -
*/
func (f *FigureTagHtml) AriaInvalid(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-invalid", value)
	return f
}

/*
AriaKeyshortcuts -
*/
func (f *FigureTagHtml) AriaKeyshortcuts(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-keyshortcuts", value)
	return f
}

/*
AriaLabel -
*/
func (f *FigureTagHtml) AriaLabel(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-label", value)
	return f
}

/*
AriaLabelledby -
*/
func (f *FigureTagHtml) AriaLabelledby(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-labelledby", value)
	return f
}

/*
AriaLive -
*/
func (f *FigureTagHtml) AriaLive(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-live", value)
	return f
}

/*
AriaOwns -
*/
func (f *FigureTagHtml) AriaOwns(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-owns", value)
	return f
}

/*
AriaRelevant -
*/
func (f *FigureTagHtml) AriaRelevant(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("aria-relevant", value)
	return f
}

/*
AriaRoledescription -
*/
func (f *FigureTagHtml) AriaRoledescription(value string) *FigureTagHtml {
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
func (f *FigureTagHtml) Onabort(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onabort", value)
	return f
}

/*
Onautocomplete -
*/
func (f *FigureTagHtml) Onautocomplete(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onautocomplete", value)
	return f
}

/*
Onautocompleteerror -
*/
func (f *FigureTagHtml) Onautocompleteerror(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onautocompleteerror", value)
	return f
}

/*
Onblur -
*/
func (f *FigureTagHtml) Onblur(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onblur", value)
	return f
}

/*
Oncancel -
*/
func (f *FigureTagHtml) Oncancel(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncancel", value)
	return f
}

/*
Oncanplay -
*/
func (f *FigureTagHtml) Oncanplay(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncanplay", value)
	return f
}

/*
Oncanplaythrough -
*/
func (f *FigureTagHtml) Oncanplaythrough(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncanplaythrough", value)
	return f
}

/*
Onchange -
*/
func (f *FigureTagHtml) Onchange(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onchange", value)
	return f
}

/*
Onclick -
*/
func (f *FigureTagHtml) Onclick(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onclick", value)
	return f
}

/*
Onclose -
*/
func (f *FigureTagHtml) Onclose(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onclose", value)
	return f
}

/*
Oncontextmenu -
*/
func (f *FigureTagHtml) Oncontextmenu(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncontextmenu", value)
	return f
}

/*
Oncuechange -
*/
func (f *FigureTagHtml) Oncuechange(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oncuechange", value)
	return f
}

/*
Ondblclick -
*/
func (f *FigureTagHtml) Ondblclick(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondblclick", value)
	return f
}

/*
Ondrag -
*/
func (f *FigureTagHtml) Ondrag(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondrag", value)
	return f
}

/*
Ondragend -
*/
func (f *FigureTagHtml) Ondragend(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragend", value)
	return f
}

/*
Ondragenter -
*/
func (f *FigureTagHtml) Ondragenter(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragenter", value)
	return f
}

/*
Ondragleave -
*/
func (f *FigureTagHtml) Ondragleave(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragleave", value)
	return f
}

/*
Ondragover -
*/
func (f *FigureTagHtml) Ondragover(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragover", value)
	return f
}

/*
Ondragstart -
*/
func (f *FigureTagHtml) Ondragstart(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondragstart", value)
	return f
}

/*
Ondrop -
*/
func (f *FigureTagHtml) Ondrop(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondrop", value)
	return f
}

/*
Ondurationchange -
*/
func (f *FigureTagHtml) Ondurationchange(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ondurationchange", value)
	return f
}

/*
Onemptied -
*/
func (f *FigureTagHtml) Onemptied(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onemptied", value)
	return f
}

/*
Onended -
*/
func (f *FigureTagHtml) Onended(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onended", value)
	return f
}

/*
Onfocus -
*/
func (f *FigureTagHtml) Onfocus(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onfocus", value)
	return f
}

/*
Oninput -
*/
func (f *FigureTagHtml) Oninput(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oninput", value)
	return f
}

/*
Oninvalid -
*/
func (f *FigureTagHtml) Oninvalid(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("oninvalid", value)
	return f
}

/*
Onkeydown -
*/
func (f *FigureTagHtml) Onkeydown(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onkeydown", value)
	return f
}

/*
Onkeypress -
*/
func (f *FigureTagHtml) Onkeypress(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onkeypress", value)
	return f
}

/*
Onkeyup -
*/
func (f *FigureTagHtml) Onkeyup(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onkeyup", value)
	return f
}

/*
Onloadeddata -
*/
func (f *FigureTagHtml) Onloadeddata(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onloadeddata", value)
	return f
}

/*
Onloadedmetadata -
*/
func (f *FigureTagHtml) Onloadedmetadata(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onloadedmetadata", value)
	return f
}

/*
Onloadstart -
*/
func (f *FigureTagHtml) Onloadstart(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onloadstart", value)
	return f
}

/*
Onmousedown -
*/
func (f *FigureTagHtml) Onmousedown(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmousedown", value)
	return f
}

/*
Onmouseenter -
*/
func (f *FigureTagHtml) Onmouseenter(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseenter", value)
	return f
}

/*
Onmouseleave -
*/
func (f *FigureTagHtml) Onmouseleave(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseleave", value)
	return f
}

/*
Onmousemove -
*/
func (f *FigureTagHtml) Onmousemove(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmousemove", value)
	return f
}

/*
Onmouseout -
*/
func (f *FigureTagHtml) Onmouseout(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseout", value)
	return f
}

/*
Onmouseover -
*/
func (f *FigureTagHtml) Onmouseover(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseover", value)
	return f
}

/*
Onmouseup -
*/
func (f *FigureTagHtml) Onmouseup(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmouseup", value)
	return f
}

/*
Onmousewheel -
*/
func (f *FigureTagHtml) Onmousewheel(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmousewheel", value)
	return f
}

/*
Onpause -
*/
func (f *FigureTagHtml) Onpause(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpause", value)
	return f
}

/*
Onplay -
*/
func (f *FigureTagHtml) Onplay(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onplay", value)
	return f
}

/*
Onplaying -
*/
func (f *FigureTagHtml) Onplaying(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onplaying", value)
	return f
}

/*
Onprogress -
*/
func (f *FigureTagHtml) Onprogress(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onprogress", value)
	return f
}

/*
Onratechange -
*/
func (f *FigureTagHtml) Onratechange(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onratechange", value)
	return f
}

/*
Onreset -
*/
func (f *FigureTagHtml) Onreset(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onreset", value)
	return f
}

/*
Onscroll -
*/
func (f *FigureTagHtml) Onscroll(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onscroll", value)
	return f
}

/*
Onseeked -
*/
func (f *FigureTagHtml) Onseeked(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onseeked", value)
	return f
}

/*
Onseeking -
*/
func (f *FigureTagHtml) Onseeking(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onseeking", value)
	return f
}

/*
Onselect -
*/
func (f *FigureTagHtml) Onselect(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onselect", value)
	return f
}

/*
Onshow -
*/
func (f *FigureTagHtml) Onshow(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onshow", value)
	return f
}

/*
Onsort -
*/
func (f *FigureTagHtml) Onsort(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onsort", value)
	return f
}

/*
Onstalled -
*/
func (f *FigureTagHtml) Onstalled(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onstalled", value)
	return f
}

/*
Onsubmit -
*/
func (f *FigureTagHtml) Onsubmit(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onsubmit", value)
	return f
}

/*
Onsuspend -
*/
func (f *FigureTagHtml) Onsuspend(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onsuspend", value)
	return f
}

/*
Ontimeupdate -
*/
func (f *FigureTagHtml) Ontimeupdate(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ontimeupdate", value)
	return f
}

/*
Ontoggle -
*/
func (f *FigureTagHtml) Ontoggle(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ontoggle", value)
	return f
}

/*
Onvolumechange -
*/
func (f *FigureTagHtml) Onvolumechange(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onvolumechange", value)
	return f
}

/*
Onwaiting -
*/
func (f *FigureTagHtml) Onwaiting(value string) *FigureTagHtml {
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
func (f *FigureTagHtml) Onafterprint(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onafterprint", value)
	return f
}

/*
Onbeforeprint -
*/
func (f *FigureTagHtml) Onbeforeprint(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onbeforeprint", value)
	return f
}

/*
Onbeforeunload -
*/
func (f *FigureTagHtml) Onbeforeunload(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onbeforeunload", value)
	return f
}

/*
Onerror -
*/
func (f *FigureTagHtml) Onerror(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onerror", value)
	return f
}

/*
Onhashchange -
*/
func (f *FigureTagHtml) Onhashchange(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onhashchange", value)
	return f
}

/*
Onload -
*/
func (f *FigureTagHtml) Onload(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onload", value)
	return f
}

/*
Onmessage -
*/
func (f *FigureTagHtml) Onmessage(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onmessage", value)
	return f
}

/*
Onoffline -
*/
func (f *FigureTagHtml) Onoffline(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onoffline", value)
	return f
}

/*
Ononline -
*/
func (f *FigureTagHtml) Ononline(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("ononline", value)
	return f
}

/*
Onpagehide -
*/
func (f *FigureTagHtml) Onpagehide(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpagehide", value)
	return f
}

/*
Onpageshow -
*/
func (f *FigureTagHtml) Onpageshow(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpageshow", value)
	return f
}

/*
Onpopstate -
*/
func (f *FigureTagHtml) Onpopstate(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onpopstate", value)
	return f
}

/*
Onresize -
*/
func (f *FigureTagHtml) Onresize(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onresize", value)
	return f
}

/*
Onstorage -
*/
func (f *FigureTagHtml) Onstorage(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onstorage", value)
	return f
}

/*
Onunload -
*/
func (f *FigureTagHtml) Onunload(value string) *FigureTagHtml {
	if f.attributes == nil {
		f.attributes = []*Attribute{}
	}
	f.registerAttribute("onunload", value)
	return f
}
