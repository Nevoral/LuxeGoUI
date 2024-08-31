// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func InsHtml(tags []any) *InsTagHtml {
	node := &InsTagHtml{
		tag: &tag{
			name:                "ins",
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

type InsTagHtml struct {
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
func (i *InsTagHtml) CustomAttribute(attributes ...*Attribute) *InsTagHtml {
	i.registerAttributes(attributes...)
	return i
}

/*
Children - Method for nesting tags into parent tag
*/
func (i *InsTagHtml) Children(tags ...any) *InsTagHtml {
	return i.supportedChildrenCheck(tags)
}

func (i *InsTagHtml) supportedChildrenCheck(tags []any) *InsTagHtml {
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
Cite -
*/
func (i *InsTagHtml) Cite(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("cite", value)
	return i
}

/*
Datetime -
*/
func (i *InsTagHtml) Datetime(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("datetime", value)
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
func (i *InsTagHtml) Onabort(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onabort", value)
	return i
}

/*
Onautocomplete -
*/
func (i *InsTagHtml) Onautocomplete(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocomplete", value)
	return i
}

/*
Onautocompleteerror -
*/
func (i *InsTagHtml) Onautocompleteerror(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocompleteerror", value)
	return i
}

/*
Onblur -
*/
func (i *InsTagHtml) Onblur(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onblur", value)
	return i
}

/*
Oncancel -
*/
func (i *InsTagHtml) Oncancel(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncancel", value)
	return i
}

/*
Oncanplay -
*/
func (i *InsTagHtml) Oncanplay(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplay", value)
	return i
}

/*
Oncanplaythrough -
*/
func (i *InsTagHtml) Oncanplaythrough(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplaythrough", value)
	return i
}

/*
Onchange -
*/
func (i *InsTagHtml) Onchange(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onchange", value)
	return i
}

/*
Onclick -
*/
func (i *InsTagHtml) Onclick(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclick", value)
	return i
}

/*
Onclose -
*/
func (i *InsTagHtml) Onclose(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclose", value)
	return i
}

/*
Oncontextmenu -
*/
func (i *InsTagHtml) Oncontextmenu(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncontextmenu", value)
	return i
}

/*
Oncuechange -
*/
func (i *InsTagHtml) Oncuechange(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncuechange", value)
	return i
}

/*
Ondblclick -
*/
func (i *InsTagHtml) Ondblclick(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondblclick", value)
	return i
}

/*
Ondrag -
*/
func (i *InsTagHtml) Ondrag(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrag", value)
	return i
}

/*
Ondragend -
*/
func (i *InsTagHtml) Ondragend(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragend", value)
	return i
}

/*
Ondragenter -
*/
func (i *InsTagHtml) Ondragenter(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragenter", value)
	return i
}

/*
Ondragleave -
*/
func (i *InsTagHtml) Ondragleave(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragleave", value)
	return i
}

/*
Ondragover -
*/
func (i *InsTagHtml) Ondragover(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragover", value)
	return i
}

/*
Ondragstart -
*/
func (i *InsTagHtml) Ondragstart(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragstart", value)
	return i
}

/*
Ondrop -
*/
func (i *InsTagHtml) Ondrop(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrop", value)
	return i
}

/*
Ondurationchange -
*/
func (i *InsTagHtml) Ondurationchange(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondurationchange", value)
	return i
}

/*
Onemptied -
*/
func (i *InsTagHtml) Onemptied(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onemptied", value)
	return i
}

/*
Onended -
*/
func (i *InsTagHtml) Onended(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onended", value)
	return i
}

/*
Onfocus -
*/
func (i *InsTagHtml) Onfocus(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onfocus", value)
	return i
}

/*
Oninput -
*/
func (i *InsTagHtml) Oninput(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninput", value)
	return i
}

/*
Oninvalid -
*/
func (i *InsTagHtml) Oninvalid(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninvalid", value)
	return i
}

/*
Onkeydown -
*/
func (i *InsTagHtml) Onkeydown(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeydown", value)
	return i
}

/*
Onkeypress -
*/
func (i *InsTagHtml) Onkeypress(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeypress", value)
	return i
}

/*
Onkeyup -
*/
func (i *InsTagHtml) Onkeyup(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeyup", value)
	return i
}

/*
Onloadeddata -
*/
func (i *InsTagHtml) Onloadeddata(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadeddata", value)
	return i
}

/*
Onloadedmetadata -
*/
func (i *InsTagHtml) Onloadedmetadata(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadedmetadata", value)
	return i
}

/*
Onloadstart -
*/
func (i *InsTagHtml) Onloadstart(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadstart", value)
	return i
}

/*
Onmousedown -
*/
func (i *InsTagHtml) Onmousedown(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousedown", value)
	return i
}

/*
Onmouseenter -
*/
func (i *InsTagHtml) Onmouseenter(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseenter", value)
	return i
}

/*
Onmouseleave -
*/
func (i *InsTagHtml) Onmouseleave(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseleave", value)
	return i
}

/*
Onmousemove -
*/
func (i *InsTagHtml) Onmousemove(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousemove", value)
	return i
}

/*
Onmouseout -
*/
func (i *InsTagHtml) Onmouseout(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseout", value)
	return i
}

/*
Onmouseover -
*/
func (i *InsTagHtml) Onmouseover(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseover", value)
	return i
}

/*
Onmouseup -
*/
func (i *InsTagHtml) Onmouseup(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseup", value)
	return i
}

/*
Onmousewheel -
*/
func (i *InsTagHtml) Onmousewheel(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousewheel", value)
	return i
}

/*
Onpause -
*/
func (i *InsTagHtml) Onpause(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpause", value)
	return i
}

/*
Onplay -
*/
func (i *InsTagHtml) Onplay(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplay", value)
	return i
}

/*
Onplaying -
*/
func (i *InsTagHtml) Onplaying(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplaying", value)
	return i
}

/*
Onprogress -
*/
func (i *InsTagHtml) Onprogress(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onprogress", value)
	return i
}

/*
Onratechange -
*/
func (i *InsTagHtml) Onratechange(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onratechange", value)
	return i
}

/*
Onreset -
*/
func (i *InsTagHtml) Onreset(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onreset", value)
	return i
}

/*
Onscroll -
*/
func (i *InsTagHtml) Onscroll(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onscroll", value)
	return i
}

/*
Onseeked -
*/
func (i *InsTagHtml) Onseeked(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeked", value)
	return i
}

/*
Onseeking -
*/
func (i *InsTagHtml) Onseeking(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeking", value)
	return i
}

/*
Onselect -
*/
func (i *InsTagHtml) Onselect(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onselect", value)
	return i
}

/*
Onshow -
*/
func (i *InsTagHtml) Onshow(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onshow", value)
	return i
}

/*
Onsort -
*/
func (i *InsTagHtml) Onsort(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsort", value)
	return i
}

/*
Onstalled -
*/
func (i *InsTagHtml) Onstalled(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstalled", value)
	return i
}

/*
Onsubmit -
*/
func (i *InsTagHtml) Onsubmit(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsubmit", value)
	return i
}

/*
Onsuspend -
*/
func (i *InsTagHtml) Onsuspend(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsuspend", value)
	return i
}

/*
Ontimeupdate -
*/
func (i *InsTagHtml) Ontimeupdate(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontimeupdate", value)
	return i
}

/*
Ontoggle -
*/
func (i *InsTagHtml) Ontoggle(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontoggle", value)
	return i
}

/*
Onvolumechange -
*/
func (i *InsTagHtml) Onvolumechange(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onvolumechange", value)
	return i
}

/*
Onwaiting -
*/
func (i *InsTagHtml) Onwaiting(value string) *InsTagHtml {
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
func (i *InsTagHtml) Onafterprint(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onafterprint", value)
	return i
}

/*
Onbeforeprint -
*/
func (i *InsTagHtml) Onbeforeprint(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeprint", value)
	return i
}

/*
Onbeforeunload -
*/
func (i *InsTagHtml) Onbeforeunload(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeunload", value)
	return i
}

/*
Onerror -
*/
func (i *InsTagHtml) Onerror(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onerror", value)
	return i
}

/*
Onhashchange -
*/
func (i *InsTagHtml) Onhashchange(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onhashchange", value)
	return i
}

/*
Onload -
*/
func (i *InsTagHtml) Onload(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onload", value)
	return i
}

/*
Onmessage -
*/
func (i *InsTagHtml) Onmessage(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmessage", value)
	return i
}

/*
Onoffline -
*/
func (i *InsTagHtml) Onoffline(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onoffline", value)
	return i
}

/*
Ononline -
*/
func (i *InsTagHtml) Ononline(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ononline", value)
	return i
}

/*
Onpagehide -
*/
func (i *InsTagHtml) Onpagehide(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpagehide", value)
	return i
}

/*
Onpageshow -
*/
func (i *InsTagHtml) Onpageshow(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpageshow", value)
	return i
}

/*
Onpopstate -
*/
func (i *InsTagHtml) Onpopstate(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpopstate", value)
	return i
}

/*
Onresize -
*/
func (i *InsTagHtml) Onresize(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onresize", value)
	return i
}

/*
Onstorage -
*/
func (i *InsTagHtml) Onstorage(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstorage", value)
	return i
}

/*
Onunload -
*/
func (i *InsTagHtml) Onunload(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onunload", value)
	return i
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (i *InsTagHtml) AccessKey(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("accessKey", value)
	return i
}

/*
Aria -
*/
func (i *InsTagHtml) Aria(value string) *InsTagHtml {
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
func (i *InsTagHtml) Autocapitalize(value string) *InsTagHtml {
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
func (i *InsTagHtml) Autofocus(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("autofocus", value)
	return i
}

/*
Class -
*/
func (i *InsTagHtml) Class(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("class", value)
	return i
}

/*
Contenteditable -
*/
func (i *InsTagHtml) Contenteditable(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("contenteditable", value)
	return i
}

/*
Data -
*/
func (i *InsTagHtml) Data(name, value string) *InsTagHtml {
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
func (i *InsTagHtml) Dir(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("dir", value)
	return i
}

/*
Draggable -
*/
func (i *InsTagHtml) Draggable(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("draggable", value)
	return i
}

/*
EnterKeyHint -
*/
func (i *InsTagHtml) EnterKeyHint(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("enterKeyHint", value)
	return i
}

/*
ExportParts -
*/
func (i *InsTagHtml) ExportParts(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("exportParts", value)
	return i
}

/*
Hidden -
*/
func (i *InsTagHtml) Hidden(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("hidden", value)
	return i
}

/*
Id -
*/
func (i *InsTagHtml) Id(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("id", value)
	return i
}

/*
Inert -
*/
func (i *InsTagHtml) Inert(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inert", value)
	return i
}

/*
InputMode -
*/
func (i *InsTagHtml) InputMode(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inputMode", value)
	return i
}

/*
Is -
*/
func (i *InsTagHtml) Is(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("is", value)
	return i
}

/*
ItemId -
*/
func (i *InsTagHtml) ItemId(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemId", value)
	return i
}

/*
ItemProp -
*/
func (i *InsTagHtml) ItemProp(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemProp", value)
	return i
}

/*
ItemRef -
*/
func (i *InsTagHtml) ItemRef(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemRef", value)
	return i
}

/*
ItemScope -
*/
func (i *InsTagHtml) ItemScope(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemScope", value)
	return i
}

/*
ItemType -
*/
func (i *InsTagHtml) ItemType(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemType", value)
	return i
}

/*
Lang -
*/
func (i *InsTagHtml) Lang(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("lang", value)
	return i
}

/*
Nonce -
*/
func (i *InsTagHtml) Nonce(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("nonce", value)
	return i
}

/*
Part -
*/
func (i *InsTagHtml) Part(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("part", value)
	return i
}

/*
Popover -
*/
func (i *InsTagHtml) Popover() *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("popover", "")
	return i
}

/*
Role -
*/
func (i *InsTagHtml) Role(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("role", value)
	return i
}

/*
Slot -
*/
func (i *InsTagHtml) Slot(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("slot", value)
	return i
}

/*
Spellcheck -
*/
func (i *InsTagHtml) Spellcheck(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("spellcheck", value)
	return i
}

/*
Style -
*/
func (i *InsTagHtml) Style(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("style", value)
	return i
}

/*
Tabindex -
*/
func (i *InsTagHtml) Tabindex(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("tabindex", value)
	return i
}

/*
Title -
*/
func (i *InsTagHtml) Title(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("title", value)
	return i
}

/*
Translate -
*/
func (i *InsTagHtml) Translate(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("translate", value)
	return i
}

/*
VirtualKeyBoardPolicy -
*/
func (i *InsTagHtml) VirtualKeyBoardPolicy(value string) *InsTagHtml {
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
func (i *InsTagHtml) AriaAtomic(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-atomic", value)
	return i
}

/*
AriaBusy -
*/
func (i *InsTagHtml) AriaBusy(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-busy", value)
	return i
}

/*
AriaControls -
*/
func (i *InsTagHtml) AriaControls(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-controls", value)
	return i
}

/*
AriaCurrent -
*/
func (i *InsTagHtml) AriaCurrent(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-current", value)
	return i
}

/*
AriaDescribedby -
*/
func (i *InsTagHtml) AriaDescribedby(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-describedby", value)
	return i
}

/*
AriaDescription -
*/
func (i *InsTagHtml) AriaDescription(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-description", value)
	return i
}

/*
AriaDetails -
*/
func (i *InsTagHtml) AriaDetails(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-details", value)
	return i
}

/*
AriaDisabled -
*/
func (i *InsTagHtml) AriaDisabled(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-disabled", value)
	return i
}

/*
AriaDropeffect -
*/
func (i *InsTagHtml) AriaDropeffect(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-dropeffect", value)
	return i
}

/*
AriaErrormessage -
*/
func (i *InsTagHtml) AriaErrormessage(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-errormessage", value)
	return i
}

/*
AriaFlowto -
*/
func (i *InsTagHtml) AriaFlowto(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-flowto", value)
	return i
}

/*
AriaGrabbed -
*/
func (i *InsTagHtml) AriaGrabbed(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-grabbed", value)
	return i
}

/*
AriaHaspopup -
*/
func (i *InsTagHtml) AriaHaspopup(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-haspopup", value)
	return i
}

/*
AriaHidden -
*/
func (i *InsTagHtml) AriaHidden(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-hidden", value)
	return i
}

/*
AriaInvalid -
*/
func (i *InsTagHtml) AriaInvalid(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-invalid", value)
	return i
}

/*
AriaKeyshortcuts -
*/
func (i *InsTagHtml) AriaKeyshortcuts(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-keyshortcuts", value)
	return i
}

/*
AriaLabel -
*/
func (i *InsTagHtml) AriaLabel(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-label", value)
	return i
}

/*
AriaLabelledby -
*/
func (i *InsTagHtml) AriaLabelledby(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-labelledby", value)
	return i
}

/*
AriaLive -
*/
func (i *InsTagHtml) AriaLive(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-live", value)
	return i
}

/*
AriaOwns -
*/
func (i *InsTagHtml) AriaOwns(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-owns", value)
	return i
}

/*
AriaRelevant -
*/
func (i *InsTagHtml) AriaRelevant(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-relevant", value)
	return i
}

/*
AriaRoledescription -
*/
func (i *InsTagHtml) AriaRoledescription(value string) *InsTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-roledescription", value)
	return i
}
