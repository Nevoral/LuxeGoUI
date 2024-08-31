// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func TheadHtml(tags []any) *TheadTagHtml {
	node := &TheadTagHtml{
		tag: &tag{
			name:                "thead",
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

type TheadTagHtml struct {
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
func (t *TheadTagHtml) CustomAttribute(attributes ...*Attribute) *TheadTagHtml {
	t.registerAttributes(attributes...)
	return t
}

/*
Children - Method for nesting tags into parent tag
*/
func (t *TheadTagHtml) Children(tags ...any) *TheadTagHtml {
	return t.supportedChildrenCheck(tags)
}

func (t *TheadTagHtml) supportedChildrenCheck(tags []any) *TheadTagHtml {
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
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (t *TheadTagHtml) Onabort(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onabort", value)
	return t
}

/*
Onautocomplete -
*/
func (t *TheadTagHtml) Onautocomplete(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onautocomplete", value)
	return t
}

/*
Onautocompleteerror -
*/
func (t *TheadTagHtml) Onautocompleteerror(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onautocompleteerror", value)
	return t
}

/*
Onblur -
*/
func (t *TheadTagHtml) Onblur(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onblur", value)
	return t
}

/*
Oncancel -
*/
func (t *TheadTagHtml) Oncancel(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncancel", value)
	return t
}

/*
Oncanplay -
*/
func (t *TheadTagHtml) Oncanplay(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncanplay", value)
	return t
}

/*
Oncanplaythrough -
*/
func (t *TheadTagHtml) Oncanplaythrough(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncanplaythrough", value)
	return t
}

/*
Onchange -
*/
func (t *TheadTagHtml) Onchange(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onchange", value)
	return t
}

/*
Onclick -
*/
func (t *TheadTagHtml) Onclick(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onclick", value)
	return t
}

/*
Onclose -
*/
func (t *TheadTagHtml) Onclose(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onclose", value)
	return t
}

/*
Oncontextmenu -
*/
func (t *TheadTagHtml) Oncontextmenu(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncontextmenu", value)
	return t
}

/*
Oncuechange -
*/
func (t *TheadTagHtml) Oncuechange(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncuechange", value)
	return t
}

/*
Ondblclick -
*/
func (t *TheadTagHtml) Ondblclick(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondblclick", value)
	return t
}

/*
Ondrag -
*/
func (t *TheadTagHtml) Ondrag(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondrag", value)
	return t
}

/*
Ondragend -
*/
func (t *TheadTagHtml) Ondragend(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragend", value)
	return t
}

/*
Ondragenter -
*/
func (t *TheadTagHtml) Ondragenter(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragenter", value)
	return t
}

/*
Ondragleave -
*/
func (t *TheadTagHtml) Ondragleave(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragleave", value)
	return t
}

/*
Ondragover -
*/
func (t *TheadTagHtml) Ondragover(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragover", value)
	return t
}

/*
Ondragstart -
*/
func (t *TheadTagHtml) Ondragstart(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragstart", value)
	return t
}

/*
Ondrop -
*/
func (t *TheadTagHtml) Ondrop(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondrop", value)
	return t
}

/*
Ondurationchange -
*/
func (t *TheadTagHtml) Ondurationchange(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondurationchange", value)
	return t
}

/*
Onemptied -
*/
func (t *TheadTagHtml) Onemptied(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onemptied", value)
	return t
}

/*
Onended -
*/
func (t *TheadTagHtml) Onended(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onended", value)
	return t
}

/*
Onfocus -
*/
func (t *TheadTagHtml) Onfocus(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onfocus", value)
	return t
}

/*
Oninput -
*/
func (t *TheadTagHtml) Oninput(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oninput", value)
	return t
}

/*
Oninvalid -
*/
func (t *TheadTagHtml) Oninvalid(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oninvalid", value)
	return t
}

/*
Onkeydown -
*/
func (t *TheadTagHtml) Onkeydown(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeydown", value)
	return t
}

/*
Onkeypress -
*/
func (t *TheadTagHtml) Onkeypress(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeypress", value)
	return t
}

/*
Onkeyup -
*/
func (t *TheadTagHtml) Onkeyup(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeyup", value)
	return t
}

/*
Onloadeddata -
*/
func (t *TheadTagHtml) Onloadeddata(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadeddata", value)
	return t
}

/*
Onloadedmetadata -
*/
func (t *TheadTagHtml) Onloadedmetadata(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadedmetadata", value)
	return t
}

/*
Onloadstart -
*/
func (t *TheadTagHtml) Onloadstart(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadstart", value)
	return t
}

/*
Onmousedown -
*/
func (t *TheadTagHtml) Onmousedown(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousedown", value)
	return t
}

/*
Onmouseenter -
*/
func (t *TheadTagHtml) Onmouseenter(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseenter", value)
	return t
}

/*
Onmouseleave -
*/
func (t *TheadTagHtml) Onmouseleave(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseleave", value)
	return t
}

/*
Onmousemove -
*/
func (t *TheadTagHtml) Onmousemove(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousemove", value)
	return t
}

/*
Onmouseout -
*/
func (t *TheadTagHtml) Onmouseout(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseout", value)
	return t
}

/*
Onmouseover -
*/
func (t *TheadTagHtml) Onmouseover(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseover", value)
	return t
}

/*
Onmouseup -
*/
func (t *TheadTagHtml) Onmouseup(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseup", value)
	return t
}

/*
Onmousewheel -
*/
func (t *TheadTagHtml) Onmousewheel(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousewheel", value)
	return t
}

/*
Onpause -
*/
func (t *TheadTagHtml) Onpause(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpause", value)
	return t
}

/*
Onplay -
*/
func (t *TheadTagHtml) Onplay(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onplay", value)
	return t
}

/*
Onplaying -
*/
func (t *TheadTagHtml) Onplaying(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onplaying", value)
	return t
}

/*
Onprogress -
*/
func (t *TheadTagHtml) Onprogress(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onprogress", value)
	return t
}

/*
Onratechange -
*/
func (t *TheadTagHtml) Onratechange(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onratechange", value)
	return t
}

/*
Onreset -
*/
func (t *TheadTagHtml) Onreset(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onreset", value)
	return t
}

/*
Onscroll -
*/
func (t *TheadTagHtml) Onscroll(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onscroll", value)
	return t
}

/*
Onseeked -
*/
func (t *TheadTagHtml) Onseeked(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onseeked", value)
	return t
}

/*
Onseeking -
*/
func (t *TheadTagHtml) Onseeking(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onseeking", value)
	return t
}

/*
Onselect -
*/
func (t *TheadTagHtml) Onselect(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onselect", value)
	return t
}

/*
Onshow -
*/
func (t *TheadTagHtml) Onshow(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onshow", value)
	return t
}

/*
Onsort -
*/
func (t *TheadTagHtml) Onsort(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsort", value)
	return t
}

/*
Onstalled -
*/
func (t *TheadTagHtml) Onstalled(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onstalled", value)
	return t
}

/*
Onsubmit -
*/
func (t *TheadTagHtml) Onsubmit(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsubmit", value)
	return t
}

/*
Onsuspend -
*/
func (t *TheadTagHtml) Onsuspend(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsuspend", value)
	return t
}

/*
Ontimeupdate -
*/
func (t *TheadTagHtml) Ontimeupdate(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ontimeupdate", value)
	return t
}

/*
Ontoggle -
*/
func (t *TheadTagHtml) Ontoggle(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ontoggle", value)
	return t
}

/*
Onvolumechange -
*/
func (t *TheadTagHtml) Onvolumechange(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onvolumechange", value)
	return t
}

/*
Onwaiting -
*/
func (t *TheadTagHtml) Onwaiting(value string) *TheadTagHtml {
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
func (t *TheadTagHtml) Onafterprint(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onafterprint", value)
	return t
}

/*
Onbeforeprint -
*/
func (t *TheadTagHtml) Onbeforeprint(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onbeforeprint", value)
	return t
}

/*
Onbeforeunload -
*/
func (t *TheadTagHtml) Onbeforeunload(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onbeforeunload", value)
	return t
}

/*
Onerror -
*/
func (t *TheadTagHtml) Onerror(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onerror", value)
	return t
}

/*
Onhashchange -
*/
func (t *TheadTagHtml) Onhashchange(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onhashchange", value)
	return t
}

/*
Onload -
*/
func (t *TheadTagHtml) Onload(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onload", value)
	return t
}

/*
Onmessage -
*/
func (t *TheadTagHtml) Onmessage(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmessage", value)
	return t
}

/*
Onoffline -
*/
func (t *TheadTagHtml) Onoffline(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onoffline", value)
	return t
}

/*
Ononline -
*/
func (t *TheadTagHtml) Ononline(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ononline", value)
	return t
}

/*
Onpagehide -
*/
func (t *TheadTagHtml) Onpagehide(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpagehide", value)
	return t
}

/*
Onpageshow -
*/
func (t *TheadTagHtml) Onpageshow(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpageshow", value)
	return t
}

/*
Onpopstate -
*/
func (t *TheadTagHtml) Onpopstate(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpopstate", value)
	return t
}

/*
Onresize -
*/
func (t *TheadTagHtml) Onresize(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onresize", value)
	return t
}

/*
Onstorage -
*/
func (t *TheadTagHtml) Onstorage(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onstorage", value)
	return t
}

/*
Onunload -
*/
func (t *TheadTagHtml) Onunload(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onunload", value)
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
func (t *TheadTagHtml) AccessKey(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("accessKey", value)
	return t
}

/*
Aria -
*/
func (t *TheadTagHtml) Aria(value string) *TheadTagHtml {
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
func (t *TheadTagHtml) Autocapitalize(value string) *TheadTagHtml {
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
func (t *TheadTagHtml) Autofocus(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("autofocus", value)
	return t
}

/*
Class -
*/
func (t *TheadTagHtml) Class(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("class", value)
	return t
}

/*
Contenteditable -
*/
func (t *TheadTagHtml) Contenteditable(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("contenteditable", value)
	return t
}

/*
Data -
*/
func (t *TheadTagHtml) Data(name, value string) *TheadTagHtml {
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
func (t *TheadTagHtml) Dir(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("dir", value)
	return t
}

/*
Draggable -
*/
func (t *TheadTagHtml) Draggable(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("draggable", value)
	return t
}

/*
EnterKeyHint -
*/
func (t *TheadTagHtml) EnterKeyHint(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("enterKeyHint", value)
	return t
}

/*
ExportParts -
*/
func (t *TheadTagHtml) ExportParts(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("exportParts", value)
	return t
}

/*
Hidden -
*/
func (t *TheadTagHtml) Hidden(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("hidden", value)
	return t
}

/*
Id -
*/
func (t *TheadTagHtml) Id(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("id", value)
	return t
}

/*
Inert -
*/
func (t *TheadTagHtml) Inert(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("inert", value)
	return t
}

/*
InputMode -
*/
func (t *TheadTagHtml) InputMode(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("inputMode", value)
	return t
}

/*
Is -
*/
func (t *TheadTagHtml) Is(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("is", value)
	return t
}

/*
ItemId -
*/
func (t *TheadTagHtml) ItemId(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemId", value)
	return t
}

/*
ItemProp -
*/
func (t *TheadTagHtml) ItemProp(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemProp", value)
	return t
}

/*
ItemRef -
*/
func (t *TheadTagHtml) ItemRef(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemRef", value)
	return t
}

/*
ItemScope -
*/
func (t *TheadTagHtml) ItemScope(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemScope", value)
	return t
}

/*
ItemType -
*/
func (t *TheadTagHtml) ItemType(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemType", value)
	return t
}

/*
Lang -
*/
func (t *TheadTagHtml) Lang(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("lang", value)
	return t
}

/*
Nonce -
*/
func (t *TheadTagHtml) Nonce(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("nonce", value)
	return t
}

/*
Part -
*/
func (t *TheadTagHtml) Part(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("part", value)
	return t
}

/*
Popover -
*/
func (t *TheadTagHtml) Popover() *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("popover", "")
	return t
}

/*
Role -
*/
func (t *TheadTagHtml) Role(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("role", value)
	return t
}

/*
Slot -
*/
func (t *TheadTagHtml) Slot(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("slot", value)
	return t
}

/*
Spellcheck -
*/
func (t *TheadTagHtml) Spellcheck(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("spellcheck", value)
	return t
}

/*
Style -
*/
func (t *TheadTagHtml) Style(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("style", value)
	return t
}

/*
Tabindex -
*/
func (t *TheadTagHtml) Tabindex(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("tabindex", value)
	return t
}

/*
Title -
*/
func (t *TheadTagHtml) Title(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("title", value)
	return t
}

/*
Translate -
*/
func (t *TheadTagHtml) Translate(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("translate", value)
	return t
}

/*
VirtualKeyBoardPolicy -
*/
func (t *TheadTagHtml) VirtualKeyBoardPolicy(value string) *TheadTagHtml {
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
func (t *TheadTagHtml) AriaAtomic(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-atomic", value)
	return t
}

/*
AriaBusy -
*/
func (t *TheadTagHtml) AriaBusy(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-busy", value)
	return t
}

/*
AriaControls -
*/
func (t *TheadTagHtml) AriaControls(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-controls", value)
	return t
}

/*
AriaCurrent -
*/
func (t *TheadTagHtml) AriaCurrent(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-current", value)
	return t
}

/*
AriaDescribedby -
*/
func (t *TheadTagHtml) AriaDescribedby(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-describedby", value)
	return t
}

/*
AriaDescription -
*/
func (t *TheadTagHtml) AriaDescription(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-description", value)
	return t
}

/*
AriaDetails -
*/
func (t *TheadTagHtml) AriaDetails(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-details", value)
	return t
}

/*
AriaDisabled -
*/
func (t *TheadTagHtml) AriaDisabled(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-disabled", value)
	return t
}

/*
AriaDropeffect -
*/
func (t *TheadTagHtml) AriaDropeffect(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-dropeffect", value)
	return t
}

/*
AriaErrormessage -
*/
func (t *TheadTagHtml) AriaErrormessage(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-errormessage", value)
	return t
}

/*
AriaFlowto -
*/
func (t *TheadTagHtml) AriaFlowto(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-flowto", value)
	return t
}

/*
AriaGrabbed -
*/
func (t *TheadTagHtml) AriaGrabbed(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-grabbed", value)
	return t
}

/*
AriaHaspopup -
*/
func (t *TheadTagHtml) AriaHaspopup(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-haspopup", value)
	return t
}

/*
AriaHidden -
*/
func (t *TheadTagHtml) AriaHidden(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-hidden", value)
	return t
}

/*
AriaInvalid -
*/
func (t *TheadTagHtml) AriaInvalid(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-invalid", value)
	return t
}

/*
AriaKeyshortcuts -
*/
func (t *TheadTagHtml) AriaKeyshortcuts(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-keyshortcuts", value)
	return t
}

/*
AriaLabel -
*/
func (t *TheadTagHtml) AriaLabel(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-label", value)
	return t
}

/*
AriaLabelledby -
*/
func (t *TheadTagHtml) AriaLabelledby(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-labelledby", value)
	return t
}

/*
AriaLive -
*/
func (t *TheadTagHtml) AriaLive(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-live", value)
	return t
}

/*
AriaOwns -
*/
func (t *TheadTagHtml) AriaOwns(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-owns", value)
	return t
}

/*
AriaRelevant -
*/
func (t *TheadTagHtml) AriaRelevant(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-relevant", value)
	return t
}

/*
AriaRoledescription -
*/
func (t *TheadTagHtml) AriaRoledescription(value string) *TheadTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-roledescription", value)
	return t
}
