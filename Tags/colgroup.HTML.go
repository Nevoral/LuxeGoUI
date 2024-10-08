// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func ColgroupHtml(tags []any) *ColgroupTagHtml {
	node := &ColgroupTagHtml{
		tag: &tag{
			name:                "colgroup",
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

type ColgroupTagHtml struct {
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
func (c *ColgroupTagHtml) CustomAttribute(attributes ...*Attribute) *ColgroupTagHtml {
	c.registerAttributes(attributes...)
	return c
}

/*
Children - Method for nesting tags into parent tag
*/
func (c *ColgroupTagHtml) Children(tags ...any) *ColgroupTagHtml {
	return c.supportedChildrenCheck(tags)
}

func (c *ColgroupTagHtml) supportedChildrenCheck(tags []any) *ColgroupTagHtml {
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
Span -
*/
func (c *ColgroupTagHtml) Span(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("span", value)
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
func (c *ColgroupTagHtml) AccessKey(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("accessKey", value)
	return c
}

/*
Aria -
*/
func (c *ColgroupTagHtml) Aria(value string) *ColgroupTagHtml {
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
func (c *ColgroupTagHtml) Autocapitalize(value string) *ColgroupTagHtml {
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
func (c *ColgroupTagHtml) Autofocus(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("autofocus", value)
	return c
}

/*
Class -
*/
func (c *ColgroupTagHtml) Class(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("class", value)
	return c
}

/*
Contenteditable -
*/
func (c *ColgroupTagHtml) Contenteditable(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("contenteditable", value)
	return c
}

/*
Data -
*/
func (c *ColgroupTagHtml) Data(name, value string) *ColgroupTagHtml {
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
func (c *ColgroupTagHtml) Dir(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("dir", value)
	return c
}

/*
Draggable -
*/
func (c *ColgroupTagHtml) Draggable(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("draggable", value)
	return c
}

/*
EnterKeyHint -
*/
func (c *ColgroupTagHtml) EnterKeyHint(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("enterKeyHint", value)
	return c
}

/*
ExportParts -
*/
func (c *ColgroupTagHtml) ExportParts(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("exportParts", value)
	return c
}

/*
Hidden -
*/
func (c *ColgroupTagHtml) Hidden(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("hidden", value)
	return c
}

/*
Id -
*/
func (c *ColgroupTagHtml) Id(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("id", value)
	return c
}

/*
Inert -
*/
func (c *ColgroupTagHtml) Inert(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("inert", value)
	return c
}

/*
InputMode -
*/
func (c *ColgroupTagHtml) InputMode(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("inputMode", value)
	return c
}

/*
Is -
*/
func (c *ColgroupTagHtml) Is(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("is", value)
	return c
}

/*
ItemId -
*/
func (c *ColgroupTagHtml) ItemId(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemId", value)
	return c
}

/*
ItemProp -
*/
func (c *ColgroupTagHtml) ItemProp(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemProp", value)
	return c
}

/*
ItemRef -
*/
func (c *ColgroupTagHtml) ItemRef(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemRef", value)
	return c
}

/*
ItemScope -
*/
func (c *ColgroupTagHtml) ItemScope(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemScope", value)
	return c
}

/*
ItemType -
*/
func (c *ColgroupTagHtml) ItemType(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemType", value)
	return c
}

/*
Lang -
*/
func (c *ColgroupTagHtml) Lang(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("lang", value)
	return c
}

/*
Nonce -
*/
func (c *ColgroupTagHtml) Nonce(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("nonce", value)
	return c
}

/*
Part -
*/
func (c *ColgroupTagHtml) Part(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("part", value)
	return c
}

/*
Popover -
*/
func (c *ColgroupTagHtml) Popover() *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("popover", "")
	return c
}

/*
Role -
*/
func (c *ColgroupTagHtml) Role(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("role", value)
	return c
}

/*
Slot -
*/
func (c *ColgroupTagHtml) Slot(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("slot", value)
	return c
}

/*
Spellcheck -
*/
func (c *ColgroupTagHtml) Spellcheck(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("spellcheck", value)
	return c
}

/*
Style -
*/
func (c *ColgroupTagHtml) Style(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("style", value)
	return c
}

/*
Tabindex -
*/
func (c *ColgroupTagHtml) Tabindex(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("tabindex", value)
	return c
}

/*
Title -
*/
func (c *ColgroupTagHtml) Title(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("title", value)
	return c
}

/*
Translate -
*/
func (c *ColgroupTagHtml) Translate(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("translate", value)
	return c
}

/*
VirtualKeyBoardPolicy -
*/
func (c *ColgroupTagHtml) VirtualKeyBoardPolicy(value string) *ColgroupTagHtml {
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
func (c *ColgroupTagHtml) AriaAtomic(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-atomic", value)
	return c
}

/*
AriaBusy -
*/
func (c *ColgroupTagHtml) AriaBusy(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-busy", value)
	return c
}

/*
AriaControls -
*/
func (c *ColgroupTagHtml) AriaControls(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-controls", value)
	return c
}

/*
AriaCurrent -
*/
func (c *ColgroupTagHtml) AriaCurrent(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-current", value)
	return c
}

/*
AriaDescribedby -
*/
func (c *ColgroupTagHtml) AriaDescribedby(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-describedby", value)
	return c
}

/*
AriaDescription -
*/
func (c *ColgroupTagHtml) AriaDescription(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-description", value)
	return c
}

/*
AriaDetails -
*/
func (c *ColgroupTagHtml) AriaDetails(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-details", value)
	return c
}

/*
AriaDisabled -
*/
func (c *ColgroupTagHtml) AriaDisabled(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-disabled", value)
	return c
}

/*
AriaDropeffect -
*/
func (c *ColgroupTagHtml) AriaDropeffect(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-dropeffect", value)
	return c
}

/*
AriaErrormessage -
*/
func (c *ColgroupTagHtml) AriaErrormessage(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-errormessage", value)
	return c
}

/*
AriaFlowto -
*/
func (c *ColgroupTagHtml) AriaFlowto(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-flowto", value)
	return c
}

/*
AriaGrabbed -
*/
func (c *ColgroupTagHtml) AriaGrabbed(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-grabbed", value)
	return c
}

/*
AriaHaspopup -
*/
func (c *ColgroupTagHtml) AriaHaspopup(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-haspopup", value)
	return c
}

/*
AriaHidden -
*/
func (c *ColgroupTagHtml) AriaHidden(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-hidden", value)
	return c
}

/*
AriaInvalid -
*/
func (c *ColgroupTagHtml) AriaInvalid(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-invalid", value)
	return c
}

/*
AriaKeyshortcuts -
*/
func (c *ColgroupTagHtml) AriaKeyshortcuts(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-keyshortcuts", value)
	return c
}

/*
AriaLabel -
*/
func (c *ColgroupTagHtml) AriaLabel(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-label", value)
	return c
}

/*
AriaLabelledby -
*/
func (c *ColgroupTagHtml) AriaLabelledby(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-labelledby", value)
	return c
}

/*
AriaLive -
*/
func (c *ColgroupTagHtml) AriaLive(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-live", value)
	return c
}

/*
AriaOwns -
*/
func (c *ColgroupTagHtml) AriaOwns(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-owns", value)
	return c
}

/*
AriaRelevant -
*/
func (c *ColgroupTagHtml) AriaRelevant(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-relevant", value)
	return c
}

/*
AriaRoledescription -
*/
func (c *ColgroupTagHtml) AriaRoledescription(value string) *ColgroupTagHtml {
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
func (c *ColgroupTagHtml) Onabort(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onabort", value)
	return c
}

/*
Onautocomplete -
*/
func (c *ColgroupTagHtml) Onautocomplete(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onautocomplete", value)
	return c
}

/*
Onautocompleteerror -
*/
func (c *ColgroupTagHtml) Onautocompleteerror(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onautocompleteerror", value)
	return c
}

/*
Onblur -
*/
func (c *ColgroupTagHtml) Onblur(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onblur", value)
	return c
}

/*
Oncancel -
*/
func (c *ColgroupTagHtml) Oncancel(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncancel", value)
	return c
}

/*
Oncanplay -
*/
func (c *ColgroupTagHtml) Oncanplay(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncanplay", value)
	return c
}

/*
Oncanplaythrough -
*/
func (c *ColgroupTagHtml) Oncanplaythrough(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncanplaythrough", value)
	return c
}

/*
Onchange -
*/
func (c *ColgroupTagHtml) Onchange(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onchange", value)
	return c
}

/*
Onclick -
*/
func (c *ColgroupTagHtml) Onclick(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onclick", value)
	return c
}

/*
Onclose -
*/
func (c *ColgroupTagHtml) Onclose(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onclose", value)
	return c
}

/*
Oncontextmenu -
*/
func (c *ColgroupTagHtml) Oncontextmenu(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncontextmenu", value)
	return c
}

/*
Oncuechange -
*/
func (c *ColgroupTagHtml) Oncuechange(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncuechange", value)
	return c
}

/*
Ondblclick -
*/
func (c *ColgroupTagHtml) Ondblclick(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondblclick", value)
	return c
}

/*
Ondrag -
*/
func (c *ColgroupTagHtml) Ondrag(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondrag", value)
	return c
}

/*
Ondragend -
*/
func (c *ColgroupTagHtml) Ondragend(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragend", value)
	return c
}

/*
Ondragenter -
*/
func (c *ColgroupTagHtml) Ondragenter(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragenter", value)
	return c
}

/*
Ondragleave -
*/
func (c *ColgroupTagHtml) Ondragleave(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragleave", value)
	return c
}

/*
Ondragover -
*/
func (c *ColgroupTagHtml) Ondragover(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragover", value)
	return c
}

/*
Ondragstart -
*/
func (c *ColgroupTagHtml) Ondragstart(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragstart", value)
	return c
}

/*
Ondrop -
*/
func (c *ColgroupTagHtml) Ondrop(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondrop", value)
	return c
}

/*
Ondurationchange -
*/
func (c *ColgroupTagHtml) Ondurationchange(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondurationchange", value)
	return c
}

/*
Onemptied -
*/
func (c *ColgroupTagHtml) Onemptied(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onemptied", value)
	return c
}

/*
Onended -
*/
func (c *ColgroupTagHtml) Onended(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onended", value)
	return c
}

/*
Onfocus -
*/
func (c *ColgroupTagHtml) Onfocus(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onfocus", value)
	return c
}

/*
Oninput -
*/
func (c *ColgroupTagHtml) Oninput(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oninput", value)
	return c
}

/*
Oninvalid -
*/
func (c *ColgroupTagHtml) Oninvalid(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oninvalid", value)
	return c
}

/*
Onkeydown -
*/
func (c *ColgroupTagHtml) Onkeydown(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeydown", value)
	return c
}

/*
Onkeypress -
*/
func (c *ColgroupTagHtml) Onkeypress(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeypress", value)
	return c
}

/*
Onkeyup -
*/
func (c *ColgroupTagHtml) Onkeyup(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeyup", value)
	return c
}

/*
Onloadeddata -
*/
func (c *ColgroupTagHtml) Onloadeddata(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadeddata", value)
	return c
}

/*
Onloadedmetadata -
*/
func (c *ColgroupTagHtml) Onloadedmetadata(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadedmetadata", value)
	return c
}

/*
Onloadstart -
*/
func (c *ColgroupTagHtml) Onloadstart(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadstart", value)
	return c
}

/*
Onmousedown -
*/
func (c *ColgroupTagHtml) Onmousedown(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousedown", value)
	return c
}

/*
Onmouseenter -
*/
func (c *ColgroupTagHtml) Onmouseenter(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseenter", value)
	return c
}

/*
Onmouseleave -
*/
func (c *ColgroupTagHtml) Onmouseleave(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseleave", value)
	return c
}

/*
Onmousemove -
*/
func (c *ColgroupTagHtml) Onmousemove(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousemove", value)
	return c
}

/*
Onmouseout -
*/
func (c *ColgroupTagHtml) Onmouseout(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseout", value)
	return c
}

/*
Onmouseover -
*/
func (c *ColgroupTagHtml) Onmouseover(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseover", value)
	return c
}

/*
Onmouseup -
*/
func (c *ColgroupTagHtml) Onmouseup(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseup", value)
	return c
}

/*
Onmousewheel -
*/
func (c *ColgroupTagHtml) Onmousewheel(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousewheel", value)
	return c
}

/*
Onpause -
*/
func (c *ColgroupTagHtml) Onpause(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpause", value)
	return c
}

/*
Onplay -
*/
func (c *ColgroupTagHtml) Onplay(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onplay", value)
	return c
}

/*
Onplaying -
*/
func (c *ColgroupTagHtml) Onplaying(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onplaying", value)
	return c
}

/*
Onprogress -
*/
func (c *ColgroupTagHtml) Onprogress(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onprogress", value)
	return c
}

/*
Onratechange -
*/
func (c *ColgroupTagHtml) Onratechange(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onratechange", value)
	return c
}

/*
Onreset -
*/
func (c *ColgroupTagHtml) Onreset(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onreset", value)
	return c
}

/*
Onscroll -
*/
func (c *ColgroupTagHtml) Onscroll(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onscroll", value)
	return c
}

/*
Onseeked -
*/
func (c *ColgroupTagHtml) Onseeked(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onseeked", value)
	return c
}

/*
Onseeking -
*/
func (c *ColgroupTagHtml) Onseeking(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onseeking", value)
	return c
}

/*
Onselect -
*/
func (c *ColgroupTagHtml) Onselect(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onselect", value)
	return c
}

/*
Onshow -
*/
func (c *ColgroupTagHtml) Onshow(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onshow", value)
	return c
}

/*
Onsort -
*/
func (c *ColgroupTagHtml) Onsort(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsort", value)
	return c
}

/*
Onstalled -
*/
func (c *ColgroupTagHtml) Onstalled(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onstalled", value)
	return c
}

/*
Onsubmit -
*/
func (c *ColgroupTagHtml) Onsubmit(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsubmit", value)
	return c
}

/*
Onsuspend -
*/
func (c *ColgroupTagHtml) Onsuspend(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsuspend", value)
	return c
}

/*
Ontimeupdate -
*/
func (c *ColgroupTagHtml) Ontimeupdate(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ontimeupdate", value)
	return c
}

/*
Ontoggle -
*/
func (c *ColgroupTagHtml) Ontoggle(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ontoggle", value)
	return c
}

/*
Onvolumechange -
*/
func (c *ColgroupTagHtml) Onvolumechange(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onvolumechange", value)
	return c
}

/*
Onwaiting -
*/
func (c *ColgroupTagHtml) Onwaiting(value string) *ColgroupTagHtml {
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
func (c *ColgroupTagHtml) Onafterprint(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onafterprint", value)
	return c
}

/*
Onbeforeprint -
*/
func (c *ColgroupTagHtml) Onbeforeprint(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onbeforeprint", value)
	return c
}

/*
Onbeforeunload -
*/
func (c *ColgroupTagHtml) Onbeforeunload(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onbeforeunload", value)
	return c
}

/*
Onerror -
*/
func (c *ColgroupTagHtml) Onerror(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onerror", value)
	return c
}

/*
Onhashchange -
*/
func (c *ColgroupTagHtml) Onhashchange(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onhashchange", value)
	return c
}

/*
Onload -
*/
func (c *ColgroupTagHtml) Onload(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onload", value)
	return c
}

/*
Onmessage -
*/
func (c *ColgroupTagHtml) Onmessage(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmessage", value)
	return c
}

/*
Onoffline -
*/
func (c *ColgroupTagHtml) Onoffline(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onoffline", value)
	return c
}

/*
Ononline -
*/
func (c *ColgroupTagHtml) Ononline(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ononline", value)
	return c
}

/*
Onpagehide -
*/
func (c *ColgroupTagHtml) Onpagehide(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpagehide", value)
	return c
}

/*
Onpageshow -
*/
func (c *ColgroupTagHtml) Onpageshow(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpageshow", value)
	return c
}

/*
Onpopstate -
*/
func (c *ColgroupTagHtml) Onpopstate(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpopstate", value)
	return c
}

/*
Onresize -
*/
func (c *ColgroupTagHtml) Onresize(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onresize", value)
	return c
}

/*
Onstorage -
*/
func (c *ColgroupTagHtml) Onstorage(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onstorage", value)
	return c
}

/*
Onunload -
*/
func (c *ColgroupTagHtml) Onunload(value string) *ColgroupTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onunload", value)
	return c
}
