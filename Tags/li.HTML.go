// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func LiHtml(tags []any) *LiTagHtml {
	node := &LiTagHtml{
		tag: &tag{
			name:                "li",
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

type LiTagHtml struct {
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
func (l *LiTagHtml) CustomAttribute(attributes ...*Attribute) *LiTagHtml {
	l.registerAttributes(attributes...)
	return l
}

/*
Children - Method for nesting tags into parent tag
*/
func (l *LiTagHtml) Children(tags ...any) *LiTagHtml {
	return l.supportedChildrenCheck(tags)
}

func (l *LiTagHtml) supportedChildrenCheck(tags []any) *LiTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			l.registerChildren(TextContentSvg(val).getTag())
		case content:
			l.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				l.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return l
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Value - Specifies the value of an <input> element.
Specifies the value of an <input> element.
*/
func (l *LiTagHtml) Value(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("value", value)
	return l
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (l *LiTagHtml) Onabort(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onabort", value)
	return l
}

/*
Onautocomplete -
*/
func (l *LiTagHtml) Onautocomplete(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onautocomplete", value)
	return l
}

/*
Onautocompleteerror -
*/
func (l *LiTagHtml) Onautocompleteerror(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onautocompleteerror", value)
	return l
}

/*
Onblur -
*/
func (l *LiTagHtml) Onblur(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onblur", value)
	return l
}

/*
Oncancel -
*/
func (l *LiTagHtml) Oncancel(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("oncancel", value)
	return l
}

/*
Oncanplay -
*/
func (l *LiTagHtml) Oncanplay(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("oncanplay", value)
	return l
}

/*
Oncanplaythrough -
*/
func (l *LiTagHtml) Oncanplaythrough(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("oncanplaythrough", value)
	return l
}

/*
Onchange -
*/
func (l *LiTagHtml) Onchange(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onchange", value)
	return l
}

/*
Onclick -
*/
func (l *LiTagHtml) Onclick(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onclick", value)
	return l
}

/*
Onclose -
*/
func (l *LiTagHtml) Onclose(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onclose", value)
	return l
}

/*
Oncontextmenu -
*/
func (l *LiTagHtml) Oncontextmenu(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("oncontextmenu", value)
	return l
}

/*
Oncuechange -
*/
func (l *LiTagHtml) Oncuechange(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("oncuechange", value)
	return l
}

/*
Ondblclick -
*/
func (l *LiTagHtml) Ondblclick(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondblclick", value)
	return l
}

/*
Ondrag -
*/
func (l *LiTagHtml) Ondrag(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondrag", value)
	return l
}

/*
Ondragend -
*/
func (l *LiTagHtml) Ondragend(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondragend", value)
	return l
}

/*
Ondragenter -
*/
func (l *LiTagHtml) Ondragenter(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondragenter", value)
	return l
}

/*
Ondragleave -
*/
func (l *LiTagHtml) Ondragleave(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondragleave", value)
	return l
}

/*
Ondragover -
*/
func (l *LiTagHtml) Ondragover(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondragover", value)
	return l
}

/*
Ondragstart -
*/
func (l *LiTagHtml) Ondragstart(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondragstart", value)
	return l
}

/*
Ondrop -
*/
func (l *LiTagHtml) Ondrop(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondrop", value)
	return l
}

/*
Ondurationchange -
*/
func (l *LiTagHtml) Ondurationchange(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ondurationchange", value)
	return l
}

/*
Onemptied -
*/
func (l *LiTagHtml) Onemptied(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onemptied", value)
	return l
}

/*
Onended -
*/
func (l *LiTagHtml) Onended(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onended", value)
	return l
}

/*
Onfocus -
*/
func (l *LiTagHtml) Onfocus(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onfocus", value)
	return l
}

/*
Oninput -
*/
func (l *LiTagHtml) Oninput(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("oninput", value)
	return l
}

/*
Oninvalid -
*/
func (l *LiTagHtml) Oninvalid(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("oninvalid", value)
	return l
}

/*
Onkeydown -
*/
func (l *LiTagHtml) Onkeydown(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onkeydown", value)
	return l
}

/*
Onkeypress -
*/
func (l *LiTagHtml) Onkeypress(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onkeypress", value)
	return l
}

/*
Onkeyup -
*/
func (l *LiTagHtml) Onkeyup(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onkeyup", value)
	return l
}

/*
Onloadeddata -
*/
func (l *LiTagHtml) Onloadeddata(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onloadeddata", value)
	return l
}

/*
Onloadedmetadata -
*/
func (l *LiTagHtml) Onloadedmetadata(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onloadedmetadata", value)
	return l
}

/*
Onloadstart -
*/
func (l *LiTagHtml) Onloadstart(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onloadstart", value)
	return l
}

/*
Onmousedown -
*/
func (l *LiTagHtml) Onmousedown(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmousedown", value)
	return l
}

/*
Onmouseenter -
*/
func (l *LiTagHtml) Onmouseenter(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmouseenter", value)
	return l
}

/*
Onmouseleave -
*/
func (l *LiTagHtml) Onmouseleave(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmouseleave", value)
	return l
}

/*
Onmousemove -
*/
func (l *LiTagHtml) Onmousemove(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmousemove", value)
	return l
}

/*
Onmouseout -
*/
func (l *LiTagHtml) Onmouseout(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmouseout", value)
	return l
}

/*
Onmouseover -
*/
func (l *LiTagHtml) Onmouseover(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmouseover", value)
	return l
}

/*
Onmouseup -
*/
func (l *LiTagHtml) Onmouseup(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmouseup", value)
	return l
}

/*
Onmousewheel -
*/
func (l *LiTagHtml) Onmousewheel(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmousewheel", value)
	return l
}

/*
Onpause -
*/
func (l *LiTagHtml) Onpause(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onpause", value)
	return l
}

/*
Onplay -
*/
func (l *LiTagHtml) Onplay(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onplay", value)
	return l
}

/*
Onplaying -
*/
func (l *LiTagHtml) Onplaying(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onplaying", value)
	return l
}

/*
Onprogress -
*/
func (l *LiTagHtml) Onprogress(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onprogress", value)
	return l
}

/*
Onratechange -
*/
func (l *LiTagHtml) Onratechange(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onratechange", value)
	return l
}

/*
Onreset -
*/
func (l *LiTagHtml) Onreset(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onreset", value)
	return l
}

/*
Onscroll -
*/
func (l *LiTagHtml) Onscroll(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onscroll", value)
	return l
}

/*
Onseeked -
*/
func (l *LiTagHtml) Onseeked(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onseeked", value)
	return l
}

/*
Onseeking -
*/
func (l *LiTagHtml) Onseeking(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onseeking", value)
	return l
}

/*
Onselect -
*/
func (l *LiTagHtml) Onselect(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onselect", value)
	return l
}

/*
Onshow -
*/
func (l *LiTagHtml) Onshow(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onshow", value)
	return l
}

/*
Onsort -
*/
func (l *LiTagHtml) Onsort(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onsort", value)
	return l
}

/*
Onstalled -
*/
func (l *LiTagHtml) Onstalled(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onstalled", value)
	return l
}

/*
Onsubmit -
*/
func (l *LiTagHtml) Onsubmit(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onsubmit", value)
	return l
}

/*
Onsuspend -
*/
func (l *LiTagHtml) Onsuspend(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onsuspend", value)
	return l
}

/*
Ontimeupdate -
*/
func (l *LiTagHtml) Ontimeupdate(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ontimeupdate", value)
	return l
}

/*
Ontoggle -
*/
func (l *LiTagHtml) Ontoggle(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ontoggle", value)
	return l
}

/*
Onvolumechange -
*/
func (l *LiTagHtml) Onvolumechange(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onvolumechange", value)
	return l
}

/*
Onwaiting -
*/
func (l *LiTagHtml) Onwaiting(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onwaiting", value)
	return l
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (l *LiTagHtml) Onafterprint(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onafterprint", value)
	return l
}

/*
Onbeforeprint -
*/
func (l *LiTagHtml) Onbeforeprint(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onbeforeprint", value)
	return l
}

/*
Onbeforeunload -
*/
func (l *LiTagHtml) Onbeforeunload(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onbeforeunload", value)
	return l
}

/*
Onerror -
*/
func (l *LiTagHtml) Onerror(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onerror", value)
	return l
}

/*
Onhashchange -
*/
func (l *LiTagHtml) Onhashchange(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onhashchange", value)
	return l
}

/*
Onload -
*/
func (l *LiTagHtml) Onload(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onload", value)
	return l
}

/*
Onmessage -
*/
func (l *LiTagHtml) Onmessage(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onmessage", value)
	return l
}

/*
Onoffline -
*/
func (l *LiTagHtml) Onoffline(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onoffline", value)
	return l
}

/*
Ononline -
*/
func (l *LiTagHtml) Ononline(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("ononline", value)
	return l
}

/*
Onpagehide -
*/
func (l *LiTagHtml) Onpagehide(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onpagehide", value)
	return l
}

/*
Onpageshow -
*/
func (l *LiTagHtml) Onpageshow(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onpageshow", value)
	return l
}

/*
Onpopstate -
*/
func (l *LiTagHtml) Onpopstate(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onpopstate", value)
	return l
}

/*
Onresize -
*/
func (l *LiTagHtml) Onresize(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onresize", value)
	return l
}

/*
Onstorage -
*/
func (l *LiTagHtml) Onstorage(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onstorage", value)
	return l
}

/*
Onunload -
*/
func (l *LiTagHtml) Onunload(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("onunload", value)
	return l
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (l *LiTagHtml) AccessKey(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("accessKey", value)
	return l
}

/*
Aria -
*/
func (l *LiTagHtml) Aria(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria", value)
	return l
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (l *LiTagHtml) Autocapitalize(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("autocapitalize", value)
	return l
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (l *LiTagHtml) Autofocus(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("autofocus", value)
	return l
}

/*
Class -
*/
func (l *LiTagHtml) Class(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("class", value)
	return l
}

/*
Contenteditable -
*/
func (l *LiTagHtml) Contenteditable(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("contenteditable", value)
	return l
}

/*
Data -
*/
func (l *LiTagHtml) Data(name, value string) *LiTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	l.registerAttribute(dataName, value)
	return l
}

/*
Dir -
*/
func (l *LiTagHtml) Dir(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("dir", value)
	return l
}

/*
Draggable -
*/
func (l *LiTagHtml) Draggable(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("draggable", value)
	return l
}

/*
EnterKeyHint -
*/
func (l *LiTagHtml) EnterKeyHint(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("enterKeyHint", value)
	return l
}

/*
ExportParts -
*/
func (l *LiTagHtml) ExportParts(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("exportParts", value)
	return l
}

/*
Hidden -
*/
func (l *LiTagHtml) Hidden(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("hidden", value)
	return l
}

/*
Id -
*/
func (l *LiTagHtml) Id(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("id", value)
	return l
}

/*
Inert -
*/
func (l *LiTagHtml) Inert(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("inert", value)
	return l
}

/*
InputMode -
*/
func (l *LiTagHtml) InputMode(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("inputMode", value)
	return l
}

/*
Is -
*/
func (l *LiTagHtml) Is(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("is", value)
	return l
}

/*
ItemId -
*/
func (l *LiTagHtml) ItemId(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("itemId", value)
	return l
}

/*
ItemProp -
*/
func (l *LiTagHtml) ItemProp(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("itemProp", value)
	return l
}

/*
ItemRef -
*/
func (l *LiTagHtml) ItemRef(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("itemRef", value)
	return l
}

/*
ItemScope -
*/
func (l *LiTagHtml) ItemScope(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("itemScope", value)
	return l
}

/*
ItemType -
*/
func (l *LiTagHtml) ItemType(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("itemType", value)
	return l
}

/*
Lang -
*/
func (l *LiTagHtml) Lang(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("lang", value)
	return l
}

/*
Nonce -
*/
func (l *LiTagHtml) Nonce(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("nonce", value)
	return l
}

/*
Part -
*/
func (l *LiTagHtml) Part(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("part", value)
	return l
}

/*
Popover -
*/
func (l *LiTagHtml) Popover() *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("popover", "")
	return l
}

/*
Role -
*/
func (l *LiTagHtml) Role(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("role", value)
	return l
}

/*
Slot -
*/
func (l *LiTagHtml) Slot(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("slot", value)
	return l
}

/*
Spellcheck -
*/
func (l *LiTagHtml) Spellcheck(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("spellcheck", value)
	return l
}

/*
Style -
*/
func (l *LiTagHtml) Style(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("style", value)
	return l
}

/*
Tabindex -
*/
func (l *LiTagHtml) Tabindex(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("tabindex", value)
	return l
}

/*
Title -
*/
func (l *LiTagHtml) Title(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("title", value)
	return l
}

/*
Translate -
*/
func (l *LiTagHtml) Translate(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("translate", value)
	return l
}

/*
VirtualKeyBoardPolicy -
*/
func (l *LiTagHtml) VirtualKeyBoardPolicy(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("virtualKeyBoardPolicy", value)
	return l
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (l *LiTagHtml) AriaAtomic(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-atomic", value)
	return l
}

/*
AriaBusy -
*/
func (l *LiTagHtml) AriaBusy(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-busy", value)
	return l
}

/*
AriaControls -
*/
func (l *LiTagHtml) AriaControls(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-controls", value)
	return l
}

/*
AriaCurrent -
*/
func (l *LiTagHtml) AriaCurrent(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-current", value)
	return l
}

/*
AriaDescribedby -
*/
func (l *LiTagHtml) AriaDescribedby(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-describedby", value)
	return l
}

/*
AriaDescription -
*/
func (l *LiTagHtml) AriaDescription(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-description", value)
	return l
}

/*
AriaDetails -
*/
func (l *LiTagHtml) AriaDetails(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-details", value)
	return l
}

/*
AriaDisabled -
*/
func (l *LiTagHtml) AriaDisabled(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-disabled", value)
	return l
}

/*
AriaDropeffect -
*/
func (l *LiTagHtml) AriaDropeffect(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-dropeffect", value)
	return l
}

/*
AriaErrormessage -
*/
func (l *LiTagHtml) AriaErrormessage(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-errormessage", value)
	return l
}

/*
AriaFlowto -
*/
func (l *LiTagHtml) AriaFlowto(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-flowto", value)
	return l
}

/*
AriaGrabbed -
*/
func (l *LiTagHtml) AriaGrabbed(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-grabbed", value)
	return l
}

/*
AriaHaspopup -
*/
func (l *LiTagHtml) AriaHaspopup(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-haspopup", value)
	return l
}

/*
AriaHidden -
*/
func (l *LiTagHtml) AriaHidden(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-hidden", value)
	return l
}

/*
AriaInvalid -
*/
func (l *LiTagHtml) AriaInvalid(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-invalid", value)
	return l
}

/*
AriaKeyshortcuts -
*/
func (l *LiTagHtml) AriaKeyshortcuts(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-keyshortcuts", value)
	return l
}

/*
AriaLabel -
*/
func (l *LiTagHtml) AriaLabel(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-label", value)
	return l
}

/*
AriaLabelledby -
*/
func (l *LiTagHtml) AriaLabelledby(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-labelledby", value)
	return l
}

/*
AriaLive -
*/
func (l *LiTagHtml) AriaLive(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-live", value)
	return l
}

/*
AriaOwns -
*/
func (l *LiTagHtml) AriaOwns(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-owns", value)
	return l
}

/*
AriaRelevant -
*/
func (l *LiTagHtml) AriaRelevant(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-relevant", value)
	return l
}

/*
AriaRoledescription -
*/
func (l *LiTagHtml) AriaRoledescription(value string) *LiTagHtml {
	if l.attributes == nil {
		l.attributes = []*Attribute{}
	}
	l.registerAttribute("aria-roledescription", value)
	return l
}
