// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func CaptionHtml(tags []any) *CaptionTagHtml {
	node := &CaptionTagHtml{
		tag: &tag{
			name:                "caption",
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

type CaptionTagHtml struct {
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
func (c *CaptionTagHtml) CustomAttribute(attributes ...*Attribute) *CaptionTagHtml {
	c.registerAttributes(attributes...)
	return c
}

/*
Children - Method for nesting tags into parent tag
*/
func (c *CaptionTagHtml) Children(tags ...any) *CaptionTagHtml {
	return c.supportedChildrenCheck(tags)
}

func (c *CaptionTagHtml) supportedChildrenCheck(tags []any) *CaptionTagHtml {
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
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (c *CaptionTagHtml) Onabort(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onabort", value)
	return c
}

/*
Onautocomplete -
*/
func (c *CaptionTagHtml) Onautocomplete(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onautocomplete", value)
	return c
}

/*
Onautocompleteerror -
*/
func (c *CaptionTagHtml) Onautocompleteerror(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onautocompleteerror", value)
	return c
}

/*
Onblur -
*/
func (c *CaptionTagHtml) Onblur(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onblur", value)
	return c
}

/*
Oncancel -
*/
func (c *CaptionTagHtml) Oncancel(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncancel", value)
	return c
}

/*
Oncanplay -
*/
func (c *CaptionTagHtml) Oncanplay(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncanplay", value)
	return c
}

/*
Oncanplaythrough -
*/
func (c *CaptionTagHtml) Oncanplaythrough(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncanplaythrough", value)
	return c
}

/*
Onchange -
*/
func (c *CaptionTagHtml) Onchange(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onchange", value)
	return c
}

/*
Onclick -
*/
func (c *CaptionTagHtml) Onclick(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onclick", value)
	return c
}

/*
Onclose -
*/
func (c *CaptionTagHtml) Onclose(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onclose", value)
	return c
}

/*
Oncontextmenu -
*/
func (c *CaptionTagHtml) Oncontextmenu(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncontextmenu", value)
	return c
}

/*
Oncuechange -
*/
func (c *CaptionTagHtml) Oncuechange(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oncuechange", value)
	return c
}

/*
Ondblclick -
*/
func (c *CaptionTagHtml) Ondblclick(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondblclick", value)
	return c
}

/*
Ondrag -
*/
func (c *CaptionTagHtml) Ondrag(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondrag", value)
	return c
}

/*
Ondragend -
*/
func (c *CaptionTagHtml) Ondragend(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragend", value)
	return c
}

/*
Ondragenter -
*/
func (c *CaptionTagHtml) Ondragenter(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragenter", value)
	return c
}

/*
Ondragleave -
*/
func (c *CaptionTagHtml) Ondragleave(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragleave", value)
	return c
}

/*
Ondragover -
*/
func (c *CaptionTagHtml) Ondragover(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragover", value)
	return c
}

/*
Ondragstart -
*/
func (c *CaptionTagHtml) Ondragstart(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondragstart", value)
	return c
}

/*
Ondrop -
*/
func (c *CaptionTagHtml) Ondrop(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondrop", value)
	return c
}

/*
Ondurationchange -
*/
func (c *CaptionTagHtml) Ondurationchange(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ondurationchange", value)
	return c
}

/*
Onemptied -
*/
func (c *CaptionTagHtml) Onemptied(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onemptied", value)
	return c
}

/*
Onended -
*/
func (c *CaptionTagHtml) Onended(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onended", value)
	return c
}

/*
Onfocus -
*/
func (c *CaptionTagHtml) Onfocus(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onfocus", value)
	return c
}

/*
Oninput -
*/
func (c *CaptionTagHtml) Oninput(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oninput", value)
	return c
}

/*
Oninvalid -
*/
func (c *CaptionTagHtml) Oninvalid(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("oninvalid", value)
	return c
}

/*
Onkeydown -
*/
func (c *CaptionTagHtml) Onkeydown(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeydown", value)
	return c
}

/*
Onkeypress -
*/
func (c *CaptionTagHtml) Onkeypress(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeypress", value)
	return c
}

/*
Onkeyup -
*/
func (c *CaptionTagHtml) Onkeyup(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onkeyup", value)
	return c
}

/*
Onloadeddata -
*/
func (c *CaptionTagHtml) Onloadeddata(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadeddata", value)
	return c
}

/*
Onloadedmetadata -
*/
func (c *CaptionTagHtml) Onloadedmetadata(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadedmetadata", value)
	return c
}

/*
Onloadstart -
*/
func (c *CaptionTagHtml) Onloadstart(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onloadstart", value)
	return c
}

/*
Onmousedown -
*/
func (c *CaptionTagHtml) Onmousedown(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousedown", value)
	return c
}

/*
Onmouseenter -
*/
func (c *CaptionTagHtml) Onmouseenter(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseenter", value)
	return c
}

/*
Onmouseleave -
*/
func (c *CaptionTagHtml) Onmouseleave(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseleave", value)
	return c
}

/*
Onmousemove -
*/
func (c *CaptionTagHtml) Onmousemove(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousemove", value)
	return c
}

/*
Onmouseout -
*/
func (c *CaptionTagHtml) Onmouseout(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseout", value)
	return c
}

/*
Onmouseover -
*/
func (c *CaptionTagHtml) Onmouseover(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseover", value)
	return c
}

/*
Onmouseup -
*/
func (c *CaptionTagHtml) Onmouseup(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmouseup", value)
	return c
}

/*
Onmousewheel -
*/
func (c *CaptionTagHtml) Onmousewheel(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmousewheel", value)
	return c
}

/*
Onpause -
*/
func (c *CaptionTagHtml) Onpause(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpause", value)
	return c
}

/*
Onplay -
*/
func (c *CaptionTagHtml) Onplay(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onplay", value)
	return c
}

/*
Onplaying -
*/
func (c *CaptionTagHtml) Onplaying(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onplaying", value)
	return c
}

/*
Onprogress -
*/
func (c *CaptionTagHtml) Onprogress(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onprogress", value)
	return c
}

/*
Onratechange -
*/
func (c *CaptionTagHtml) Onratechange(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onratechange", value)
	return c
}

/*
Onreset -
*/
func (c *CaptionTagHtml) Onreset(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onreset", value)
	return c
}

/*
Onscroll -
*/
func (c *CaptionTagHtml) Onscroll(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onscroll", value)
	return c
}

/*
Onseeked -
*/
func (c *CaptionTagHtml) Onseeked(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onseeked", value)
	return c
}

/*
Onseeking -
*/
func (c *CaptionTagHtml) Onseeking(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onseeking", value)
	return c
}

/*
Onselect -
*/
func (c *CaptionTagHtml) Onselect(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onselect", value)
	return c
}

/*
Onshow -
*/
func (c *CaptionTagHtml) Onshow(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onshow", value)
	return c
}

/*
Onsort -
*/
func (c *CaptionTagHtml) Onsort(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsort", value)
	return c
}

/*
Onstalled -
*/
func (c *CaptionTagHtml) Onstalled(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onstalled", value)
	return c
}

/*
Onsubmit -
*/
func (c *CaptionTagHtml) Onsubmit(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsubmit", value)
	return c
}

/*
Onsuspend -
*/
func (c *CaptionTagHtml) Onsuspend(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onsuspend", value)
	return c
}

/*
Ontimeupdate -
*/
func (c *CaptionTagHtml) Ontimeupdate(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ontimeupdate", value)
	return c
}

/*
Ontoggle -
*/
func (c *CaptionTagHtml) Ontoggle(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ontoggle", value)
	return c
}

/*
Onvolumechange -
*/
func (c *CaptionTagHtml) Onvolumechange(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onvolumechange", value)
	return c
}

/*
Onwaiting -
*/
func (c *CaptionTagHtml) Onwaiting(value string) *CaptionTagHtml {
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
func (c *CaptionTagHtml) Onafterprint(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onafterprint", value)
	return c
}

/*
Onbeforeprint -
*/
func (c *CaptionTagHtml) Onbeforeprint(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onbeforeprint", value)
	return c
}

/*
Onbeforeunload -
*/
func (c *CaptionTagHtml) Onbeforeunload(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onbeforeunload", value)
	return c
}

/*
Onerror -
*/
func (c *CaptionTagHtml) Onerror(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onerror", value)
	return c
}

/*
Onhashchange -
*/
func (c *CaptionTagHtml) Onhashchange(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onhashchange", value)
	return c
}

/*
Onload -
*/
func (c *CaptionTagHtml) Onload(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onload", value)
	return c
}

/*
Onmessage -
*/
func (c *CaptionTagHtml) Onmessage(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onmessage", value)
	return c
}

/*
Onoffline -
*/
func (c *CaptionTagHtml) Onoffline(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onoffline", value)
	return c
}

/*
Ononline -
*/
func (c *CaptionTagHtml) Ononline(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("ononline", value)
	return c
}

/*
Onpagehide -
*/
func (c *CaptionTagHtml) Onpagehide(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpagehide", value)
	return c
}

/*
Onpageshow -
*/
func (c *CaptionTagHtml) Onpageshow(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpageshow", value)
	return c
}

/*
Onpopstate -
*/
func (c *CaptionTagHtml) Onpopstate(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onpopstate", value)
	return c
}

/*
Onresize -
*/
func (c *CaptionTagHtml) Onresize(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onresize", value)
	return c
}

/*
Onstorage -
*/
func (c *CaptionTagHtml) Onstorage(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onstorage", value)
	return c
}

/*
Onunload -
*/
func (c *CaptionTagHtml) Onunload(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("onunload", value)
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
func (c *CaptionTagHtml) AccessKey(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("accessKey", value)
	return c
}

/*
Aria -
*/
func (c *CaptionTagHtml) Aria(value string) *CaptionTagHtml {
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
func (c *CaptionTagHtml) Autocapitalize(value string) *CaptionTagHtml {
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
func (c *CaptionTagHtml) Autofocus(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("autofocus", value)
	return c
}

/*
Class -
*/
func (c *CaptionTagHtml) Class(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("class", value)
	return c
}

/*
Contenteditable -
*/
func (c *CaptionTagHtml) Contenteditable(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("contenteditable", value)
	return c
}

/*
Data -
*/
func (c *CaptionTagHtml) Data(name, value string) *CaptionTagHtml {
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
func (c *CaptionTagHtml) Dir(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("dir", value)
	return c
}

/*
Draggable -
*/
func (c *CaptionTagHtml) Draggable(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("draggable", value)
	return c
}

/*
EnterKeyHint -
*/
func (c *CaptionTagHtml) EnterKeyHint(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("enterKeyHint", value)
	return c
}

/*
ExportParts -
*/
func (c *CaptionTagHtml) ExportParts(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("exportParts", value)
	return c
}

/*
Hidden -
*/
func (c *CaptionTagHtml) Hidden(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("hidden", value)
	return c
}

/*
Id -
*/
func (c *CaptionTagHtml) Id(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("id", value)
	return c
}

/*
Inert -
*/
func (c *CaptionTagHtml) Inert(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("inert", value)
	return c
}

/*
InputMode -
*/
func (c *CaptionTagHtml) InputMode(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("inputMode", value)
	return c
}

/*
Is -
*/
func (c *CaptionTagHtml) Is(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("is", value)
	return c
}

/*
ItemId -
*/
func (c *CaptionTagHtml) ItemId(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemId", value)
	return c
}

/*
ItemProp -
*/
func (c *CaptionTagHtml) ItemProp(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemProp", value)
	return c
}

/*
ItemRef -
*/
func (c *CaptionTagHtml) ItemRef(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemRef", value)
	return c
}

/*
ItemScope -
*/
func (c *CaptionTagHtml) ItemScope(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemScope", value)
	return c
}

/*
ItemType -
*/
func (c *CaptionTagHtml) ItemType(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("itemType", value)
	return c
}

/*
Lang -
*/
func (c *CaptionTagHtml) Lang(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("lang", value)
	return c
}

/*
Nonce -
*/
func (c *CaptionTagHtml) Nonce(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("nonce", value)
	return c
}

/*
Part -
*/
func (c *CaptionTagHtml) Part(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("part", value)
	return c
}

/*
Popover -
*/
func (c *CaptionTagHtml) Popover() *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("popover", "")
	return c
}

/*
Role -
*/
func (c *CaptionTagHtml) Role(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("role", value)
	return c
}

/*
Slot -
*/
func (c *CaptionTagHtml) Slot(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("slot", value)
	return c
}

/*
Spellcheck -
*/
func (c *CaptionTagHtml) Spellcheck(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("spellcheck", value)
	return c
}

/*
Style -
*/
func (c *CaptionTagHtml) Style(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("style", value)
	return c
}

/*
Tabindex -
*/
func (c *CaptionTagHtml) Tabindex(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("tabindex", value)
	return c
}

/*
Title -
*/
func (c *CaptionTagHtml) Title(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("title", value)
	return c
}

/*
Translate -
*/
func (c *CaptionTagHtml) Translate(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("translate", value)
	return c
}

/*
VirtualKeyBoardPolicy -
*/
func (c *CaptionTagHtml) VirtualKeyBoardPolicy(value string) *CaptionTagHtml {
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
func (c *CaptionTagHtml) AriaAtomic(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-atomic", value)
	return c
}

/*
AriaBusy -
*/
func (c *CaptionTagHtml) AriaBusy(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-busy", value)
	return c
}

/*
AriaControls -
*/
func (c *CaptionTagHtml) AriaControls(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-controls", value)
	return c
}

/*
AriaCurrent -
*/
func (c *CaptionTagHtml) AriaCurrent(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-current", value)
	return c
}

/*
AriaDescribedby -
*/
func (c *CaptionTagHtml) AriaDescribedby(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-describedby", value)
	return c
}

/*
AriaDescription -
*/
func (c *CaptionTagHtml) AriaDescription(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-description", value)
	return c
}

/*
AriaDetails -
*/
func (c *CaptionTagHtml) AriaDetails(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-details", value)
	return c
}

/*
AriaDisabled -
*/
func (c *CaptionTagHtml) AriaDisabled(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-disabled", value)
	return c
}

/*
AriaDropeffect -
*/
func (c *CaptionTagHtml) AriaDropeffect(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-dropeffect", value)
	return c
}

/*
AriaErrormessage -
*/
func (c *CaptionTagHtml) AriaErrormessage(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-errormessage", value)
	return c
}

/*
AriaFlowto -
*/
func (c *CaptionTagHtml) AriaFlowto(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-flowto", value)
	return c
}

/*
AriaGrabbed -
*/
func (c *CaptionTagHtml) AriaGrabbed(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-grabbed", value)
	return c
}

/*
AriaHaspopup -
*/
func (c *CaptionTagHtml) AriaHaspopup(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-haspopup", value)
	return c
}

/*
AriaHidden -
*/
func (c *CaptionTagHtml) AriaHidden(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-hidden", value)
	return c
}

/*
AriaInvalid -
*/
func (c *CaptionTagHtml) AriaInvalid(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-invalid", value)
	return c
}

/*
AriaKeyshortcuts -
*/
func (c *CaptionTagHtml) AriaKeyshortcuts(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-keyshortcuts", value)
	return c
}

/*
AriaLabel -
*/
func (c *CaptionTagHtml) AriaLabel(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-label", value)
	return c
}

/*
AriaLabelledby -
*/
func (c *CaptionTagHtml) AriaLabelledby(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-labelledby", value)
	return c
}

/*
AriaLive -
*/
func (c *CaptionTagHtml) AriaLive(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-live", value)
	return c
}

/*
AriaOwns -
*/
func (c *CaptionTagHtml) AriaOwns(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-owns", value)
	return c
}

/*
AriaRelevant -
*/
func (c *CaptionTagHtml) AriaRelevant(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-relevant", value)
	return c
}

/*
AriaRoledescription -
*/
func (c *CaptionTagHtml) AriaRoledescription(value string) *CaptionTagHtml {
	if c.attributes == nil {
		c.attributes = []*Attribute{}
	}
	c.registerAttribute("aria-roledescription", value)
	return c
}
