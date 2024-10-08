// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func TfootHtml(tags []any) *TfootTagHtml {
	node := &TfootTagHtml{
		tag: &tag{
			name:                "tfoot",
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

type TfootTagHtml struct {
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
func (t *TfootTagHtml) CustomAttribute(attributes ...*Attribute) *TfootTagHtml {
	t.registerAttributes(attributes...)
	return t
}

/*
Children - Method for nesting tags into parent tag
*/
func (t *TfootTagHtml) Children(tags ...any) *TfootTagHtml {
	return t.supportedChildrenCheck(tags)
}

func (t *TfootTagHtml) supportedChildrenCheck(tags []any) *TfootTagHtml {
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
func (t *TfootTagHtml) Onabort(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onabort", value)
	return t
}

/*
Onautocomplete -
*/
func (t *TfootTagHtml) Onautocomplete(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onautocomplete", value)
	return t
}

/*
Onautocompleteerror -
*/
func (t *TfootTagHtml) Onautocompleteerror(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onautocompleteerror", value)
	return t
}

/*
Onblur -
*/
func (t *TfootTagHtml) Onblur(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onblur", value)
	return t
}

/*
Oncancel -
*/
func (t *TfootTagHtml) Oncancel(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncancel", value)
	return t
}

/*
Oncanplay -
*/
func (t *TfootTagHtml) Oncanplay(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncanplay", value)
	return t
}

/*
Oncanplaythrough -
*/
func (t *TfootTagHtml) Oncanplaythrough(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncanplaythrough", value)
	return t
}

/*
Onchange -
*/
func (t *TfootTagHtml) Onchange(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onchange", value)
	return t
}

/*
Onclick -
*/
func (t *TfootTagHtml) Onclick(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onclick", value)
	return t
}

/*
Onclose -
*/
func (t *TfootTagHtml) Onclose(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onclose", value)
	return t
}

/*
Oncontextmenu -
*/
func (t *TfootTagHtml) Oncontextmenu(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncontextmenu", value)
	return t
}

/*
Oncuechange -
*/
func (t *TfootTagHtml) Oncuechange(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oncuechange", value)
	return t
}

/*
Ondblclick -
*/
func (t *TfootTagHtml) Ondblclick(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondblclick", value)
	return t
}

/*
Ondrag -
*/
func (t *TfootTagHtml) Ondrag(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondrag", value)
	return t
}

/*
Ondragend -
*/
func (t *TfootTagHtml) Ondragend(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragend", value)
	return t
}

/*
Ondragenter -
*/
func (t *TfootTagHtml) Ondragenter(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragenter", value)
	return t
}

/*
Ondragleave -
*/
func (t *TfootTagHtml) Ondragleave(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragleave", value)
	return t
}

/*
Ondragover -
*/
func (t *TfootTagHtml) Ondragover(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragover", value)
	return t
}

/*
Ondragstart -
*/
func (t *TfootTagHtml) Ondragstart(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondragstart", value)
	return t
}

/*
Ondrop -
*/
func (t *TfootTagHtml) Ondrop(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondrop", value)
	return t
}

/*
Ondurationchange -
*/
func (t *TfootTagHtml) Ondurationchange(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ondurationchange", value)
	return t
}

/*
Onemptied -
*/
func (t *TfootTagHtml) Onemptied(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onemptied", value)
	return t
}

/*
Onended -
*/
func (t *TfootTagHtml) Onended(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onended", value)
	return t
}

/*
Onfocus -
*/
func (t *TfootTagHtml) Onfocus(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onfocus", value)
	return t
}

/*
Oninput -
*/
func (t *TfootTagHtml) Oninput(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oninput", value)
	return t
}

/*
Oninvalid -
*/
func (t *TfootTagHtml) Oninvalid(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("oninvalid", value)
	return t
}

/*
Onkeydown -
*/
func (t *TfootTagHtml) Onkeydown(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeydown", value)
	return t
}

/*
Onkeypress -
*/
func (t *TfootTagHtml) Onkeypress(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeypress", value)
	return t
}

/*
Onkeyup -
*/
func (t *TfootTagHtml) Onkeyup(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onkeyup", value)
	return t
}

/*
Onloadeddata -
*/
func (t *TfootTagHtml) Onloadeddata(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadeddata", value)
	return t
}

/*
Onloadedmetadata -
*/
func (t *TfootTagHtml) Onloadedmetadata(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadedmetadata", value)
	return t
}

/*
Onloadstart -
*/
func (t *TfootTagHtml) Onloadstart(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onloadstart", value)
	return t
}

/*
Onmousedown -
*/
func (t *TfootTagHtml) Onmousedown(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousedown", value)
	return t
}

/*
Onmouseenter -
*/
func (t *TfootTagHtml) Onmouseenter(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseenter", value)
	return t
}

/*
Onmouseleave -
*/
func (t *TfootTagHtml) Onmouseleave(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseleave", value)
	return t
}

/*
Onmousemove -
*/
func (t *TfootTagHtml) Onmousemove(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousemove", value)
	return t
}

/*
Onmouseout -
*/
func (t *TfootTagHtml) Onmouseout(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseout", value)
	return t
}

/*
Onmouseover -
*/
func (t *TfootTagHtml) Onmouseover(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseover", value)
	return t
}

/*
Onmouseup -
*/
func (t *TfootTagHtml) Onmouseup(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmouseup", value)
	return t
}

/*
Onmousewheel -
*/
func (t *TfootTagHtml) Onmousewheel(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmousewheel", value)
	return t
}

/*
Onpause -
*/
func (t *TfootTagHtml) Onpause(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpause", value)
	return t
}

/*
Onplay -
*/
func (t *TfootTagHtml) Onplay(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onplay", value)
	return t
}

/*
Onplaying -
*/
func (t *TfootTagHtml) Onplaying(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onplaying", value)
	return t
}

/*
Onprogress -
*/
func (t *TfootTagHtml) Onprogress(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onprogress", value)
	return t
}

/*
Onratechange -
*/
func (t *TfootTagHtml) Onratechange(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onratechange", value)
	return t
}

/*
Onreset -
*/
func (t *TfootTagHtml) Onreset(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onreset", value)
	return t
}

/*
Onscroll -
*/
func (t *TfootTagHtml) Onscroll(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onscroll", value)
	return t
}

/*
Onseeked -
*/
func (t *TfootTagHtml) Onseeked(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onseeked", value)
	return t
}

/*
Onseeking -
*/
func (t *TfootTagHtml) Onseeking(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onseeking", value)
	return t
}

/*
Onselect -
*/
func (t *TfootTagHtml) Onselect(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onselect", value)
	return t
}

/*
Onshow -
*/
func (t *TfootTagHtml) Onshow(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onshow", value)
	return t
}

/*
Onsort -
*/
func (t *TfootTagHtml) Onsort(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsort", value)
	return t
}

/*
Onstalled -
*/
func (t *TfootTagHtml) Onstalled(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onstalled", value)
	return t
}

/*
Onsubmit -
*/
func (t *TfootTagHtml) Onsubmit(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsubmit", value)
	return t
}

/*
Onsuspend -
*/
func (t *TfootTagHtml) Onsuspend(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onsuspend", value)
	return t
}

/*
Ontimeupdate -
*/
func (t *TfootTagHtml) Ontimeupdate(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ontimeupdate", value)
	return t
}

/*
Ontoggle -
*/
func (t *TfootTagHtml) Ontoggle(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ontoggle", value)
	return t
}

/*
Onvolumechange -
*/
func (t *TfootTagHtml) Onvolumechange(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onvolumechange", value)
	return t
}

/*
Onwaiting -
*/
func (t *TfootTagHtml) Onwaiting(value string) *TfootTagHtml {
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
func (t *TfootTagHtml) Onafterprint(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onafterprint", value)
	return t
}

/*
Onbeforeprint -
*/
func (t *TfootTagHtml) Onbeforeprint(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onbeforeprint", value)
	return t
}

/*
Onbeforeunload -
*/
func (t *TfootTagHtml) Onbeforeunload(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onbeforeunload", value)
	return t
}

/*
Onerror -
*/
func (t *TfootTagHtml) Onerror(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onerror", value)
	return t
}

/*
Onhashchange -
*/
func (t *TfootTagHtml) Onhashchange(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onhashchange", value)
	return t
}

/*
Onload -
*/
func (t *TfootTagHtml) Onload(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onload", value)
	return t
}

/*
Onmessage -
*/
func (t *TfootTagHtml) Onmessage(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onmessage", value)
	return t
}

/*
Onoffline -
*/
func (t *TfootTagHtml) Onoffline(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onoffline", value)
	return t
}

/*
Ononline -
*/
func (t *TfootTagHtml) Ononline(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("ononline", value)
	return t
}

/*
Onpagehide -
*/
func (t *TfootTagHtml) Onpagehide(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpagehide", value)
	return t
}

/*
Onpageshow -
*/
func (t *TfootTagHtml) Onpageshow(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpageshow", value)
	return t
}

/*
Onpopstate -
*/
func (t *TfootTagHtml) Onpopstate(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onpopstate", value)
	return t
}

/*
Onresize -
*/
func (t *TfootTagHtml) Onresize(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onresize", value)
	return t
}

/*
Onstorage -
*/
func (t *TfootTagHtml) Onstorage(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("onstorage", value)
	return t
}

/*
Onunload -
*/
func (t *TfootTagHtml) Onunload(value string) *TfootTagHtml {
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
func (t *TfootTagHtml) AccessKey(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("accessKey", value)
	return t
}

/*
Aria -
*/
func (t *TfootTagHtml) Aria(value string) *TfootTagHtml {
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
func (t *TfootTagHtml) Autocapitalize(value string) *TfootTagHtml {
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
func (t *TfootTagHtml) Autofocus(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("autofocus", value)
	return t
}

/*
Class -
*/
func (t *TfootTagHtml) Class(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("class", value)
	return t
}

/*
Contenteditable -
*/
func (t *TfootTagHtml) Contenteditable(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("contenteditable", value)
	return t
}

/*
Data -
*/
func (t *TfootTagHtml) Data(name, value string) *TfootTagHtml {
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
func (t *TfootTagHtml) Dir(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("dir", value)
	return t
}

/*
Draggable -
*/
func (t *TfootTagHtml) Draggable(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("draggable", value)
	return t
}

/*
EnterKeyHint -
*/
func (t *TfootTagHtml) EnterKeyHint(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("enterKeyHint", value)
	return t
}

/*
ExportParts -
*/
func (t *TfootTagHtml) ExportParts(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("exportParts", value)
	return t
}

/*
Hidden -
*/
func (t *TfootTagHtml) Hidden(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("hidden", value)
	return t
}

/*
Id -
*/
func (t *TfootTagHtml) Id(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("id", value)
	return t
}

/*
Inert -
*/
func (t *TfootTagHtml) Inert(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("inert", value)
	return t
}

/*
InputMode -
*/
func (t *TfootTagHtml) InputMode(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("inputMode", value)
	return t
}

/*
Is -
*/
func (t *TfootTagHtml) Is(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("is", value)
	return t
}

/*
ItemId -
*/
func (t *TfootTagHtml) ItemId(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemId", value)
	return t
}

/*
ItemProp -
*/
func (t *TfootTagHtml) ItemProp(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemProp", value)
	return t
}

/*
ItemRef -
*/
func (t *TfootTagHtml) ItemRef(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemRef", value)
	return t
}

/*
ItemScope -
*/
func (t *TfootTagHtml) ItemScope(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemScope", value)
	return t
}

/*
ItemType -
*/
func (t *TfootTagHtml) ItemType(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("itemType", value)
	return t
}

/*
Lang -
*/
func (t *TfootTagHtml) Lang(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("lang", value)
	return t
}

/*
Nonce -
*/
func (t *TfootTagHtml) Nonce(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("nonce", value)
	return t
}

/*
Part -
*/
func (t *TfootTagHtml) Part(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("part", value)
	return t
}

/*
Popover -
*/
func (t *TfootTagHtml) Popover() *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("popover", "")
	return t
}

/*
Role -
*/
func (t *TfootTagHtml) Role(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("role", value)
	return t
}

/*
Slot -
*/
func (t *TfootTagHtml) Slot(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("slot", value)
	return t
}

/*
Spellcheck -
*/
func (t *TfootTagHtml) Spellcheck(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("spellcheck", value)
	return t
}

/*
Style -
*/
func (t *TfootTagHtml) Style(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("style", value)
	return t
}

/*
Tabindex -
*/
func (t *TfootTagHtml) Tabindex(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("tabindex", value)
	return t
}

/*
Title -
*/
func (t *TfootTagHtml) Title(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("title", value)
	return t
}

/*
Translate -
*/
func (t *TfootTagHtml) Translate(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("translate", value)
	return t
}

/*
VirtualKeyBoardPolicy -
*/
func (t *TfootTagHtml) VirtualKeyBoardPolicy(value string) *TfootTagHtml {
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
func (t *TfootTagHtml) AriaAtomic(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-atomic", value)
	return t
}

/*
AriaBusy -
*/
func (t *TfootTagHtml) AriaBusy(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-busy", value)
	return t
}

/*
AriaControls -
*/
func (t *TfootTagHtml) AriaControls(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-controls", value)
	return t
}

/*
AriaCurrent -
*/
func (t *TfootTagHtml) AriaCurrent(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-current", value)
	return t
}

/*
AriaDescribedby -
*/
func (t *TfootTagHtml) AriaDescribedby(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-describedby", value)
	return t
}

/*
AriaDescription -
*/
func (t *TfootTagHtml) AriaDescription(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-description", value)
	return t
}

/*
AriaDetails -
*/
func (t *TfootTagHtml) AriaDetails(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-details", value)
	return t
}

/*
AriaDisabled -
*/
func (t *TfootTagHtml) AriaDisabled(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-disabled", value)
	return t
}

/*
AriaDropeffect -
*/
func (t *TfootTagHtml) AriaDropeffect(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-dropeffect", value)
	return t
}

/*
AriaErrormessage -
*/
func (t *TfootTagHtml) AriaErrormessage(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-errormessage", value)
	return t
}

/*
AriaFlowto -
*/
func (t *TfootTagHtml) AriaFlowto(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-flowto", value)
	return t
}

/*
AriaGrabbed -
*/
func (t *TfootTagHtml) AriaGrabbed(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-grabbed", value)
	return t
}

/*
AriaHaspopup -
*/
func (t *TfootTagHtml) AriaHaspopup(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-haspopup", value)
	return t
}

/*
AriaHidden -
*/
func (t *TfootTagHtml) AriaHidden(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-hidden", value)
	return t
}

/*
AriaInvalid -
*/
func (t *TfootTagHtml) AriaInvalid(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-invalid", value)
	return t
}

/*
AriaKeyshortcuts -
*/
func (t *TfootTagHtml) AriaKeyshortcuts(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-keyshortcuts", value)
	return t
}

/*
AriaLabel -
*/
func (t *TfootTagHtml) AriaLabel(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-label", value)
	return t
}

/*
AriaLabelledby -
*/
func (t *TfootTagHtml) AriaLabelledby(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-labelledby", value)
	return t
}

/*
AriaLive -
*/
func (t *TfootTagHtml) AriaLive(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-live", value)
	return t
}

/*
AriaOwns -
*/
func (t *TfootTagHtml) AriaOwns(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-owns", value)
	return t
}

/*
AriaRelevant -
*/
func (t *TfootTagHtml) AriaRelevant(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-relevant", value)
	return t
}

/*
AriaRoledescription -
*/
func (t *TfootTagHtml) AriaRoledescription(value string) *TfootTagHtml {
	if t.attributes == nil {
		t.attributes = []*Attribute{}
	}
	t.registerAttribute("aria-roledescription", value)
	return t
}
