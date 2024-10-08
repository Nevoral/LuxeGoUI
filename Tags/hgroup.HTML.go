// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func HgroupHtml(tags []any) *HgroupTagHtml {
	node := &HgroupTagHtml{
		tag: &tag{
			name:                "hgroup",
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

type HgroupTagHtml struct {
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
func (h *HgroupTagHtml) CustomAttribute(attributes ...*Attribute) *HgroupTagHtml {
	h.registerAttributes(attributes...)
	return h
}

/*
Children - Method for nesting tags into parent tag
*/
func (h *HgroupTagHtml) Children(tags ...any) *HgroupTagHtml {
	return h.supportedChildrenCheck(tags)
}

func (h *HgroupTagHtml) supportedChildrenCheck(tags []any) *HgroupTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			h.registerChildren(TextContentSvg(val).getTag())
		case content:
			h.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				h.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return h
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
func (h *HgroupTagHtml) Onabort(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onabort", value)
	return h
}

/*
Onautocomplete -
*/
func (h *HgroupTagHtml) Onautocomplete(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onautocomplete", value)
	return h
}

/*
Onautocompleteerror -
*/
func (h *HgroupTagHtml) Onautocompleteerror(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onautocompleteerror", value)
	return h
}

/*
Onblur -
*/
func (h *HgroupTagHtml) Onblur(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onblur", value)
	return h
}

/*
Oncancel -
*/
func (h *HgroupTagHtml) Oncancel(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("oncancel", value)
	return h
}

/*
Oncanplay -
*/
func (h *HgroupTagHtml) Oncanplay(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("oncanplay", value)
	return h
}

/*
Oncanplaythrough -
*/
func (h *HgroupTagHtml) Oncanplaythrough(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("oncanplaythrough", value)
	return h
}

/*
Onchange -
*/
func (h *HgroupTagHtml) Onchange(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onchange", value)
	return h
}

/*
Onclick -
*/
func (h *HgroupTagHtml) Onclick(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onclick", value)
	return h
}

/*
Onclose -
*/
func (h *HgroupTagHtml) Onclose(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onclose", value)
	return h
}

/*
Oncontextmenu -
*/
func (h *HgroupTagHtml) Oncontextmenu(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("oncontextmenu", value)
	return h
}

/*
Oncuechange -
*/
func (h *HgroupTagHtml) Oncuechange(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("oncuechange", value)
	return h
}

/*
Ondblclick -
*/
func (h *HgroupTagHtml) Ondblclick(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondblclick", value)
	return h
}

/*
Ondrag -
*/
func (h *HgroupTagHtml) Ondrag(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondrag", value)
	return h
}

/*
Ondragend -
*/
func (h *HgroupTagHtml) Ondragend(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondragend", value)
	return h
}

/*
Ondragenter -
*/
func (h *HgroupTagHtml) Ondragenter(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondragenter", value)
	return h
}

/*
Ondragleave -
*/
func (h *HgroupTagHtml) Ondragleave(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondragleave", value)
	return h
}

/*
Ondragover -
*/
func (h *HgroupTagHtml) Ondragover(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondragover", value)
	return h
}

/*
Ondragstart -
*/
func (h *HgroupTagHtml) Ondragstart(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondragstart", value)
	return h
}

/*
Ondrop -
*/
func (h *HgroupTagHtml) Ondrop(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondrop", value)
	return h
}

/*
Ondurationchange -
*/
func (h *HgroupTagHtml) Ondurationchange(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ondurationchange", value)
	return h
}

/*
Onemptied -
*/
func (h *HgroupTagHtml) Onemptied(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onemptied", value)
	return h
}

/*
Onended -
*/
func (h *HgroupTagHtml) Onended(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onended", value)
	return h
}

/*
Onfocus -
*/
func (h *HgroupTagHtml) Onfocus(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onfocus", value)
	return h
}

/*
Oninput -
*/
func (h *HgroupTagHtml) Oninput(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("oninput", value)
	return h
}

/*
Oninvalid -
*/
func (h *HgroupTagHtml) Oninvalid(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("oninvalid", value)
	return h
}

/*
Onkeydown -
*/
func (h *HgroupTagHtml) Onkeydown(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onkeydown", value)
	return h
}

/*
Onkeypress -
*/
func (h *HgroupTagHtml) Onkeypress(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onkeypress", value)
	return h
}

/*
Onkeyup -
*/
func (h *HgroupTagHtml) Onkeyup(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onkeyup", value)
	return h
}

/*
Onloadeddata -
*/
func (h *HgroupTagHtml) Onloadeddata(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onloadeddata", value)
	return h
}

/*
Onloadedmetadata -
*/
func (h *HgroupTagHtml) Onloadedmetadata(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onloadedmetadata", value)
	return h
}

/*
Onloadstart -
*/
func (h *HgroupTagHtml) Onloadstart(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onloadstart", value)
	return h
}

/*
Onmousedown -
*/
func (h *HgroupTagHtml) Onmousedown(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmousedown", value)
	return h
}

/*
Onmouseenter -
*/
func (h *HgroupTagHtml) Onmouseenter(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmouseenter", value)
	return h
}

/*
Onmouseleave -
*/
func (h *HgroupTagHtml) Onmouseleave(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmouseleave", value)
	return h
}

/*
Onmousemove -
*/
func (h *HgroupTagHtml) Onmousemove(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmousemove", value)
	return h
}

/*
Onmouseout -
*/
func (h *HgroupTagHtml) Onmouseout(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmouseout", value)
	return h
}

/*
Onmouseover -
*/
func (h *HgroupTagHtml) Onmouseover(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmouseover", value)
	return h
}

/*
Onmouseup -
*/
func (h *HgroupTagHtml) Onmouseup(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmouseup", value)
	return h
}

/*
Onmousewheel -
*/
func (h *HgroupTagHtml) Onmousewheel(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmousewheel", value)
	return h
}

/*
Onpause -
*/
func (h *HgroupTagHtml) Onpause(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onpause", value)
	return h
}

/*
Onplay -
*/
func (h *HgroupTagHtml) Onplay(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onplay", value)
	return h
}

/*
Onplaying -
*/
func (h *HgroupTagHtml) Onplaying(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onplaying", value)
	return h
}

/*
Onprogress -
*/
func (h *HgroupTagHtml) Onprogress(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onprogress", value)
	return h
}

/*
Onratechange -
*/
func (h *HgroupTagHtml) Onratechange(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onratechange", value)
	return h
}

/*
Onreset -
*/
func (h *HgroupTagHtml) Onreset(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onreset", value)
	return h
}

/*
Onscroll -
*/
func (h *HgroupTagHtml) Onscroll(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onscroll", value)
	return h
}

/*
Onseeked -
*/
func (h *HgroupTagHtml) Onseeked(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onseeked", value)
	return h
}

/*
Onseeking -
*/
func (h *HgroupTagHtml) Onseeking(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onseeking", value)
	return h
}

/*
Onselect -
*/
func (h *HgroupTagHtml) Onselect(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onselect", value)
	return h
}

/*
Onshow -
*/
func (h *HgroupTagHtml) Onshow(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onshow", value)
	return h
}

/*
Onsort -
*/
func (h *HgroupTagHtml) Onsort(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onsort", value)
	return h
}

/*
Onstalled -
*/
func (h *HgroupTagHtml) Onstalled(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onstalled", value)
	return h
}

/*
Onsubmit -
*/
func (h *HgroupTagHtml) Onsubmit(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onsubmit", value)
	return h
}

/*
Onsuspend -
*/
func (h *HgroupTagHtml) Onsuspend(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onsuspend", value)
	return h
}

/*
Ontimeupdate -
*/
func (h *HgroupTagHtml) Ontimeupdate(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ontimeupdate", value)
	return h
}

/*
Ontoggle -
*/
func (h *HgroupTagHtml) Ontoggle(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ontoggle", value)
	return h
}

/*
Onvolumechange -
*/
func (h *HgroupTagHtml) Onvolumechange(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onvolumechange", value)
	return h
}

/*
Onwaiting -
*/
func (h *HgroupTagHtml) Onwaiting(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onwaiting", value)
	return h
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (h *HgroupTagHtml) Onafterprint(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onafterprint", value)
	return h
}

/*
Onbeforeprint -
*/
func (h *HgroupTagHtml) Onbeforeprint(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onbeforeprint", value)
	return h
}

/*
Onbeforeunload -
*/
func (h *HgroupTagHtml) Onbeforeunload(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onbeforeunload", value)
	return h
}

/*
Onerror -
*/
func (h *HgroupTagHtml) Onerror(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onerror", value)
	return h
}

/*
Onhashchange -
*/
func (h *HgroupTagHtml) Onhashchange(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onhashchange", value)
	return h
}

/*
Onload -
*/
func (h *HgroupTagHtml) Onload(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onload", value)
	return h
}

/*
Onmessage -
*/
func (h *HgroupTagHtml) Onmessage(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onmessage", value)
	return h
}

/*
Onoffline -
*/
func (h *HgroupTagHtml) Onoffline(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onoffline", value)
	return h
}

/*
Ononline -
*/
func (h *HgroupTagHtml) Ononline(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("ononline", value)
	return h
}

/*
Onpagehide -
*/
func (h *HgroupTagHtml) Onpagehide(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onpagehide", value)
	return h
}

/*
Onpageshow -
*/
func (h *HgroupTagHtml) Onpageshow(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onpageshow", value)
	return h
}

/*
Onpopstate -
*/
func (h *HgroupTagHtml) Onpopstate(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onpopstate", value)
	return h
}

/*
Onresize -
*/
func (h *HgroupTagHtml) Onresize(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onresize", value)
	return h
}

/*
Onstorage -
*/
func (h *HgroupTagHtml) Onstorage(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onstorage", value)
	return h
}

/*
Onunload -
*/
func (h *HgroupTagHtml) Onunload(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("onunload", value)
	return h
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (h *HgroupTagHtml) AccessKey(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("accessKey", value)
	return h
}

/*
Aria -
*/
func (h *HgroupTagHtml) Aria(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria", value)
	return h
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (h *HgroupTagHtml) Autocapitalize(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("autocapitalize", value)
	return h
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (h *HgroupTagHtml) Autofocus(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("autofocus", value)
	return h
}

/*
Class -
*/
func (h *HgroupTagHtml) Class(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("class", value)
	return h
}

/*
Contenteditable -
*/
func (h *HgroupTagHtml) Contenteditable(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("contenteditable", value)
	return h
}

/*
Data -
*/
func (h *HgroupTagHtml) Data(name, value string) *HgroupTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	h.registerAttribute(dataName, value)
	return h
}

/*
Dir -
*/
func (h *HgroupTagHtml) Dir(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("dir", value)
	return h
}

/*
Draggable -
*/
func (h *HgroupTagHtml) Draggable(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("draggable", value)
	return h
}

/*
EnterKeyHint -
*/
func (h *HgroupTagHtml) EnterKeyHint(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("enterKeyHint", value)
	return h
}

/*
ExportParts -
*/
func (h *HgroupTagHtml) ExportParts(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("exportParts", value)
	return h
}

/*
Hidden -
*/
func (h *HgroupTagHtml) Hidden(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("hidden", value)
	return h
}

/*
Id -
*/
func (h *HgroupTagHtml) Id(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("id", value)
	return h
}

/*
Inert -
*/
func (h *HgroupTagHtml) Inert(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("inert", value)
	return h
}

/*
InputMode -
*/
func (h *HgroupTagHtml) InputMode(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("inputMode", value)
	return h
}

/*
Is -
*/
func (h *HgroupTagHtml) Is(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("is", value)
	return h
}

/*
ItemId -
*/
func (h *HgroupTagHtml) ItemId(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("itemId", value)
	return h
}

/*
ItemProp -
*/
func (h *HgroupTagHtml) ItemProp(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("itemProp", value)
	return h
}

/*
ItemRef -
*/
func (h *HgroupTagHtml) ItemRef(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("itemRef", value)
	return h
}

/*
ItemScope -
*/
func (h *HgroupTagHtml) ItemScope(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("itemScope", value)
	return h
}

/*
ItemType -
*/
func (h *HgroupTagHtml) ItemType(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("itemType", value)
	return h
}

/*
Lang -
*/
func (h *HgroupTagHtml) Lang(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("lang", value)
	return h
}

/*
Nonce -
*/
func (h *HgroupTagHtml) Nonce(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("nonce", value)
	return h
}

/*
Part -
*/
func (h *HgroupTagHtml) Part(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("part", value)
	return h
}

/*
Popover -
*/
func (h *HgroupTagHtml) Popover() *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("popover", "")
	return h
}

/*
Role -
*/
func (h *HgroupTagHtml) Role(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("role", value)
	return h
}

/*
Slot -
*/
func (h *HgroupTagHtml) Slot(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("slot", value)
	return h
}

/*
Spellcheck -
*/
func (h *HgroupTagHtml) Spellcheck(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("spellcheck", value)
	return h
}

/*
Style -
*/
func (h *HgroupTagHtml) Style(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("style", value)
	return h
}

/*
Tabindex -
*/
func (h *HgroupTagHtml) Tabindex(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("tabindex", value)
	return h
}

/*
Title -
*/
func (h *HgroupTagHtml) Title(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("title", value)
	return h
}

/*
Translate -
*/
func (h *HgroupTagHtml) Translate(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("translate", value)
	return h
}

/*
VirtualKeyBoardPolicy -
*/
func (h *HgroupTagHtml) VirtualKeyBoardPolicy(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("virtualKeyBoardPolicy", value)
	return h
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (h *HgroupTagHtml) AriaAtomic(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-atomic", value)
	return h
}

/*
AriaBusy -
*/
func (h *HgroupTagHtml) AriaBusy(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-busy", value)
	return h
}

/*
AriaControls -
*/
func (h *HgroupTagHtml) AriaControls(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-controls", value)
	return h
}

/*
AriaCurrent -
*/
func (h *HgroupTagHtml) AriaCurrent(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-current", value)
	return h
}

/*
AriaDescribedby -
*/
func (h *HgroupTagHtml) AriaDescribedby(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-describedby", value)
	return h
}

/*
AriaDescription -
*/
func (h *HgroupTagHtml) AriaDescription(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-description", value)
	return h
}

/*
AriaDetails -
*/
func (h *HgroupTagHtml) AriaDetails(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-details", value)
	return h
}

/*
AriaDisabled -
*/
func (h *HgroupTagHtml) AriaDisabled(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-disabled", value)
	return h
}

/*
AriaDropeffect -
*/
func (h *HgroupTagHtml) AriaDropeffect(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-dropeffect", value)
	return h
}

/*
AriaErrormessage -
*/
func (h *HgroupTagHtml) AriaErrormessage(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-errormessage", value)
	return h
}

/*
AriaFlowto -
*/
func (h *HgroupTagHtml) AriaFlowto(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-flowto", value)
	return h
}

/*
AriaGrabbed -
*/
func (h *HgroupTagHtml) AriaGrabbed(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-grabbed", value)
	return h
}

/*
AriaHaspopup -
*/
func (h *HgroupTagHtml) AriaHaspopup(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-haspopup", value)
	return h
}

/*
AriaHidden -
*/
func (h *HgroupTagHtml) AriaHidden(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-hidden", value)
	return h
}

/*
AriaInvalid -
*/
func (h *HgroupTagHtml) AriaInvalid(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-invalid", value)
	return h
}

/*
AriaKeyshortcuts -
*/
func (h *HgroupTagHtml) AriaKeyshortcuts(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-keyshortcuts", value)
	return h
}

/*
AriaLabel -
*/
func (h *HgroupTagHtml) AriaLabel(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-label", value)
	return h
}

/*
AriaLabelledby -
*/
func (h *HgroupTagHtml) AriaLabelledby(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-labelledby", value)
	return h
}

/*
AriaLive -
*/
func (h *HgroupTagHtml) AriaLive(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-live", value)
	return h
}

/*
AriaOwns -
*/
func (h *HgroupTagHtml) AriaOwns(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-owns", value)
	return h
}

/*
AriaRelevant -
*/
func (h *HgroupTagHtml) AriaRelevant(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-relevant", value)
	return h
}

/*
AriaRoledescription -
*/
func (h *HgroupTagHtml) AriaRoledescription(value string) *HgroupTagHtml {
	if h.attributes == nil {
		h.attributes = []*Attribute{}
	}
	h.registerAttribute("aria-roledescription", value)
	return h
}
