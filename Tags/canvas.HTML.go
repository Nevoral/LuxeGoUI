// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func CanvasHtml(tags []any) *CanvasTagHtml {
	node := &CanvasTagHtml{
		tag: &tag{
			name:                "canvas",
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

type CanvasTagHtml struct {
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
func (c *CanvasTagHtml) CustomAttribute(attributes ...*Attribute) *CanvasTagHtml {
	c.registerAttributes(attributes...)
	return c
}

/*
Children - Method for nesting tags into parent tag
*/
func (c *CanvasTagHtml) Children(tags ...any) *CanvasTagHtml {
	return c.supportedChildrenCheck(tags)
}

func (c *CanvasTagHtml) supportedChildrenCheck(tags []any) *CanvasTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			c.registerChildren(TextContentSvg(val).getTag())
		case content:
			c.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				c.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return c
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Height -
*/
func (c *CanvasTagHtml) Height(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("height", value)
	return c
}

/*
Width -
*/
func (c *CanvasTagHtml) Width(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("width", value)
	return c
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (c *CanvasTagHtml) AccessKey(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("accessKey", value)
	return c
}

/*
Aria -
*/
func (c *CanvasTagHtml) Aria(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria", value)
	return c
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (c *CanvasTagHtml) Autocapitalize(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("autocapitalize", value)
	return c
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (c *CanvasTagHtml) Autofocus(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("autofocus", value)
	return c
}

/*
Class -
*/
func (c *CanvasTagHtml) Class(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("class", value)
	return c
}

/*
Contenteditable -
*/
func (c *CanvasTagHtml) Contenteditable(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("contenteditable", value)
	return c
}

/*
Data -
*/
func (c *CanvasTagHtml) Data(name, value string) *CanvasTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	c.registerAttribute(dataName, value)
	return c
}

/*
Dir -
*/
func (c *CanvasTagHtml) Dir(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("dir", value)
	return c
}

/*
Draggable -
*/
func (c *CanvasTagHtml) Draggable(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("draggable", value)
	return c
}

/*
EnterKeyHint -
*/
func (c *CanvasTagHtml) EnterKeyHint(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("enterKeyHint", value)
	return c
}

/*
ExportParts -
*/
func (c *CanvasTagHtml) ExportParts(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("exportParts", value)
	return c
}

/*
Hidden -
*/
func (c *CanvasTagHtml) Hidden(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("hidden", value)
	return c
}

/*
Id -
*/
func (c *CanvasTagHtml) Id(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("id", value)
	return c
}

/*
Inert -
*/
func (c *CanvasTagHtml) Inert(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("inert", value)
	return c
}

/*
InputMode -
*/
func (c *CanvasTagHtml) InputMode(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("inputMode", value)
	return c
}

/*
Is -
*/
func (c *CanvasTagHtml) Is(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("is", value)
	return c
}

/*
ItemId -
*/
func (c *CanvasTagHtml) ItemId(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemId", value)
	return c
}

/*
ItemProp -
*/
func (c *CanvasTagHtml) ItemProp(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemProp", value)
	return c
}

/*
ItemRef -
*/
func (c *CanvasTagHtml) ItemRef(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemRef", value)
	return c
}

/*
ItemScope -
*/
func (c *CanvasTagHtml) ItemScope(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemScope", value)
	return c
}

/*
ItemType -
*/
func (c *CanvasTagHtml) ItemType(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemType", value)
	return c
}

/*
Lang -
*/
func (c *CanvasTagHtml) Lang(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("lang", value)
	return c
}

/*
Nonce -
*/
func (c *CanvasTagHtml) Nonce(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("nonce", value)
	return c
}

/*
Part -
*/
func (c *CanvasTagHtml) Part(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("part", value)
	return c
}

/*
Popover -
*/
func (c *CanvasTagHtml) Popover() *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("popover", "")
	return c
}

/*
Role -
*/
func (c *CanvasTagHtml) Role(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("role", value)
	return c
}

/*
Slot -
*/
func (c *CanvasTagHtml) Slot(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("slot", value)
	return c
}

/*
Spellcheck -
*/
func (c *CanvasTagHtml) Spellcheck(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("spellcheck", value)
	return c
}

/*
Style -
*/
func (c *CanvasTagHtml) Style(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("style", value)
	return c
}

/*
Tabindex -
*/
func (c *CanvasTagHtml) Tabindex(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("tabindex", value)
	return c
}

/*
Title -
*/
func (c *CanvasTagHtml) Title(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("title", value)
	return c
}

/*
Translate -
*/
func (c *CanvasTagHtml) Translate(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("translate", value)
	return c
}

/*
VirtualKeyBoardPolicy -
*/
func (c *CanvasTagHtml) VirtualKeyBoardPolicy(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("virtualKeyBoardPolicy", value)
	return c
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (c *CanvasTagHtml) AriaAtomic(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-atomic", value)
	return c
}

/*
AriaBusy -
*/
func (c *CanvasTagHtml) AriaBusy(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-busy", value)
	return c
}

/*
AriaControls -
*/
func (c *CanvasTagHtml) AriaControls(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-controls", value)
	return c
}

/*
AriaCurrent -
*/
func (c *CanvasTagHtml) AriaCurrent(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-current", value)
	return c
}

/*
AriaDescribedby -
*/
func (c *CanvasTagHtml) AriaDescribedby(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-describedby", value)
	return c
}

/*
AriaDescription -
*/
func (c *CanvasTagHtml) AriaDescription(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-description", value)
	return c
}

/*
AriaDetails -
*/
func (c *CanvasTagHtml) AriaDetails(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-details", value)
	return c
}

/*
AriaDisabled -
*/
func (c *CanvasTagHtml) AriaDisabled(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-disabled", value)
	return c
}

/*
AriaDropeffect -
*/
func (c *CanvasTagHtml) AriaDropeffect(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-dropeffect", value)
	return c
}

/*
AriaErrormessage -
*/
func (c *CanvasTagHtml) AriaErrormessage(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-errormessage", value)
	return c
}

/*
AriaFlowto -
*/
func (c *CanvasTagHtml) AriaFlowto(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-flowto", value)
	return c
}

/*
AriaGrabbed -
*/
func (c *CanvasTagHtml) AriaGrabbed(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-grabbed", value)
	return c
}

/*
AriaHaspopup -
*/
func (c *CanvasTagHtml) AriaHaspopup(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-haspopup", value)
	return c
}

/*
AriaHidden -
*/
func (c *CanvasTagHtml) AriaHidden(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-hidden", value)
	return c
}

/*
AriaInvalid -
*/
func (c *CanvasTagHtml) AriaInvalid(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-invalid", value)
	return c
}

/*
AriaKeyshortcuts -
*/
func (c *CanvasTagHtml) AriaKeyshortcuts(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-keyshortcuts", value)
	return c
}

/*
AriaLabel -
*/
func (c *CanvasTagHtml) AriaLabel(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-label", value)
	return c
}

/*
AriaLabelledby -
*/
func (c *CanvasTagHtml) AriaLabelledby(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-labelledby", value)
	return c
}

/*
AriaLive -
*/
func (c *CanvasTagHtml) AriaLive(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-live", value)
	return c
}

/*
AriaOwns -
*/
func (c *CanvasTagHtml) AriaOwns(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-owns", value)
	return c
}

/*
AriaRelevant -
*/
func (c *CanvasTagHtml) AriaRelevant(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-relevant", value)
	return c
}

/*
AriaRoledescription -
*/
func (c *CanvasTagHtml) AriaRoledescription(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-roledescription", value)
	return c
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (c *CanvasTagHtml) Onabort(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onabort", value)
	return c
}

/*
Onautocomplete -
*/
func (c *CanvasTagHtml) Onautocomplete(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onautocomplete", value)
	return c
}

/*
Onautocompleteerror -
*/
func (c *CanvasTagHtml) Onautocompleteerror(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onautocompleteerror", value)
	return c
}

/*
Onblur -
*/
func (c *CanvasTagHtml) Onblur(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onblur", value)
	return c
}

/*
Oncancel -
*/
func (c *CanvasTagHtml) Oncancel(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncancel", value)
	return c
}

/*
Oncanplay -
*/
func (c *CanvasTagHtml) Oncanplay(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncanplay", value)
	return c
}

/*
Oncanplaythrough -
*/
func (c *CanvasTagHtml) Oncanplaythrough(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncanplaythrough", value)
	return c
}

/*
Onchange -
*/
func (c *CanvasTagHtml) Onchange(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onchange", value)
	return c
}

/*
Onclick -
*/
func (c *CanvasTagHtml) Onclick(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onclick", value)
	return c
}

/*
Onclose -
*/
func (c *CanvasTagHtml) Onclose(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onclose", value)
	return c
}

/*
Oncontextmenu -
*/
func (c *CanvasTagHtml) Oncontextmenu(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncontextmenu", value)
	return c
}

/*
Oncuechange -
*/
func (c *CanvasTagHtml) Oncuechange(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncuechange", value)
	return c
}

/*
Ondblclick -
*/
func (c *CanvasTagHtml) Ondblclick(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondblclick", value)
	return c
}

/*
Ondrag -
*/
func (c *CanvasTagHtml) Ondrag(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondrag", value)
	return c
}

/*
Ondragend -
*/
func (c *CanvasTagHtml) Ondragend(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragend", value)
	return c
}

/*
Ondragenter -
*/
func (c *CanvasTagHtml) Ondragenter(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragenter", value)
	return c
}

/*
Ondragleave -
*/
func (c *CanvasTagHtml) Ondragleave(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragleave", value)
	return c
}

/*
Ondragover -
*/
func (c *CanvasTagHtml) Ondragover(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragover", value)
	return c
}

/*
Ondragstart -
*/
func (c *CanvasTagHtml) Ondragstart(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragstart", value)
	return c
}

/*
Ondrop -
*/
func (c *CanvasTagHtml) Ondrop(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondrop", value)
	return c
}

/*
Ondurationchange -
*/
func (c *CanvasTagHtml) Ondurationchange(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondurationchange", value)
	return c
}

/*
Onemptied -
*/
func (c *CanvasTagHtml) Onemptied(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onemptied", value)
	return c
}

/*
Onended -
*/
func (c *CanvasTagHtml) Onended(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onended", value)
	return c
}

/*
Onfocus -
*/
func (c *CanvasTagHtml) Onfocus(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onfocus", value)
	return c
}

/*
Oninput -
*/
func (c *CanvasTagHtml) Oninput(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oninput", value)
	return c
}

/*
Oninvalid -
*/
func (c *CanvasTagHtml) Oninvalid(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oninvalid", value)
	return c
}

/*
Onkeydown -
*/
func (c *CanvasTagHtml) Onkeydown(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeydown", value)
	return c
}

/*
Onkeypress -
*/
func (c *CanvasTagHtml) Onkeypress(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeypress", value)
	return c
}

/*
Onkeyup -
*/
func (c *CanvasTagHtml) Onkeyup(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeyup", value)
	return c
}

/*
Onloadeddata -
*/
func (c *CanvasTagHtml) Onloadeddata(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadeddata", value)
	return c
}

/*
Onloadedmetadata -
*/
func (c *CanvasTagHtml) Onloadedmetadata(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadedmetadata", value)
	return c
}

/*
Onloadstart -
*/
func (c *CanvasTagHtml) Onloadstart(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadstart", value)
	return c
}

/*
Onmousedown -
*/
func (c *CanvasTagHtml) Onmousedown(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousedown", value)
	return c
}

/*
Onmouseenter -
*/
func (c *CanvasTagHtml) Onmouseenter(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseenter", value)
	return c
}

/*
Onmouseleave -
*/
func (c *CanvasTagHtml) Onmouseleave(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseleave", value)
	return c
}

/*
Onmousemove -
*/
func (c *CanvasTagHtml) Onmousemove(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousemove", value)
	return c
}

/*
Onmouseout -
*/
func (c *CanvasTagHtml) Onmouseout(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseout", value)
	return c
}

/*
Onmouseover -
*/
func (c *CanvasTagHtml) Onmouseover(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseover", value)
	return c
}

/*
Onmouseup -
*/
func (c *CanvasTagHtml) Onmouseup(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseup", value)
	return c
}

/*
Onmousewheel -
*/
func (c *CanvasTagHtml) Onmousewheel(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousewheel", value)
	return c
}

/*
Onpause -
*/
func (c *CanvasTagHtml) Onpause(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpause", value)
	return c
}

/*
Onplay -
*/
func (c *CanvasTagHtml) Onplay(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onplay", value)
	return c
}

/*
Onplaying -
*/
func (c *CanvasTagHtml) Onplaying(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onplaying", value)
	return c
}

/*
Onprogress -
*/
func (c *CanvasTagHtml) Onprogress(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onprogress", value)
	return c
}

/*
Onratechange -
*/
func (c *CanvasTagHtml) Onratechange(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onratechange", value)
	return c
}

/*
Onreset -
*/
func (c *CanvasTagHtml) Onreset(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onreset", value)
	return c
}

/*
Onscroll -
*/
func (c *CanvasTagHtml) Onscroll(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onscroll", value)
	return c
}

/*
Onseeked -
*/
func (c *CanvasTagHtml) Onseeked(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onseeked", value)
	return c
}

/*
Onseeking -
*/
func (c *CanvasTagHtml) Onseeking(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onseeking", value)
	return c
}

/*
Onselect -
*/
func (c *CanvasTagHtml) Onselect(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onselect", value)
	return c
}

/*
Onshow -
*/
func (c *CanvasTagHtml) Onshow(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onshow", value)
	return c
}

/*
Onsort -
*/
func (c *CanvasTagHtml) Onsort(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsort", value)
	return c
}

/*
Onstalled -
*/
func (c *CanvasTagHtml) Onstalled(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onstalled", value)
	return c
}

/*
Onsubmit -
*/
func (c *CanvasTagHtml) Onsubmit(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsubmit", value)
	return c
}

/*
Onsuspend -
*/
func (c *CanvasTagHtml) Onsuspend(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsuspend", value)
	return c
}

/*
Ontimeupdate -
*/
func (c *CanvasTagHtml) Ontimeupdate(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ontimeupdate", value)
	return c
}

/*
Ontoggle -
*/
func (c *CanvasTagHtml) Ontoggle(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ontoggle", value)
	return c
}

/*
Onvolumechange -
*/
func (c *CanvasTagHtml) Onvolumechange(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onvolumechange", value)
	return c
}

/*
Onwaiting -
*/
func (c *CanvasTagHtml) Onwaiting(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onwaiting", value)
	return c
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (c *CanvasTagHtml) Onafterprint(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onafterprint", value)
	return c
}

/*
Onbeforeprint -
*/
func (c *CanvasTagHtml) Onbeforeprint(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onbeforeprint", value)
	return c
}

/*
Onbeforeunload -
*/
func (c *CanvasTagHtml) Onbeforeunload(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onbeforeunload", value)
	return c
}

/*
Onerror -
*/
func (c *CanvasTagHtml) Onerror(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onerror", value)
	return c
}

/*
Onhashchange -
*/
func (c *CanvasTagHtml) Onhashchange(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onhashchange", value)
	return c
}

/*
Onload -
*/
func (c *CanvasTagHtml) Onload(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onload", value)
	return c
}

/*
Onmessage -
*/
func (c *CanvasTagHtml) Onmessage(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmessage", value)
	return c
}

/*
Onoffline -
*/
func (c *CanvasTagHtml) Onoffline(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onoffline", value)
	return c
}

/*
Ononline -
*/
func (c *CanvasTagHtml) Ononline(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ononline", value)
	return c
}

/*
Onpagehide -
*/
func (c *CanvasTagHtml) Onpagehide(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpagehide", value)
	return c
}

/*
Onpageshow -
*/
func (c *CanvasTagHtml) Onpageshow(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpageshow", value)
	return c
}

/*
Onpopstate -
*/
func (c *CanvasTagHtml) Onpopstate(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpopstate", value)
	return c
}

/*
Onresize -
*/
func (c *CanvasTagHtml) Onresize(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onresize", value)
	return c
}

/*
Onstorage -
*/
func (c *CanvasTagHtml) Onstorage(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onstorage", value)
	return c
}

/*
Onunload -
*/
func (c *CanvasTagHtml) Onunload(value string) *CanvasTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onunload", value)
	return c
}
