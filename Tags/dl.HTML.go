// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func DlHtml(tags []any) *DlTagHtml {
	node := &DlTagHtml{
		tag: &tag{
			name:                "dl",
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

type DlTagHtml struct {
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
func (d *DlTagHtml) CustomAttribute(attributes ...*Attribute) *DlTagHtml {
	d.registerAttributes(attributes...)
	return d
}

/*
Children - Method for nesting tags into parent tag
*/
func (d *DlTagHtml) Children(tags ...any) *DlTagHtml {
	return d.supportedChildrenCheck(tags)
}

func (d *DlTagHtml) supportedChildrenCheck(tags []any) *DlTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			d.registerChildren(TextContentSvg(val).getTag())
		case content:
			d.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				d.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return d
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
func (d *DlTagHtml) Onabort(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onabort", value)
	return d
}

/*
Onautocomplete -
*/
func (d *DlTagHtml) Onautocomplete(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onautocomplete", value)
	return d
}

/*
Onautocompleteerror -
*/
func (d *DlTagHtml) Onautocompleteerror(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onautocompleteerror", value)
	return d
}

/*
Onblur -
*/
func (d *DlTagHtml) Onblur(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onblur", value)
	return d
}

/*
Oncancel -
*/
func (d *DlTagHtml) Oncancel(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncancel", value)
	return d
}

/*
Oncanplay -
*/
func (d *DlTagHtml) Oncanplay(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncanplay", value)
	return d
}

/*
Oncanplaythrough -
*/
func (d *DlTagHtml) Oncanplaythrough(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncanplaythrough", value)
	return d
}

/*
Onchange -
*/
func (d *DlTagHtml) Onchange(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onchange", value)
	return d
}

/*
Onclick -
*/
func (d *DlTagHtml) Onclick(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onclick", value)
	return d
}

/*
Onclose -
*/
func (d *DlTagHtml) Onclose(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onclose", value)
	return d
}

/*
Oncontextmenu -
*/
func (d *DlTagHtml) Oncontextmenu(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncontextmenu", value)
	return d
}

/*
Oncuechange -
*/
func (d *DlTagHtml) Oncuechange(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncuechange", value)
	return d
}

/*
Ondblclick -
*/
func (d *DlTagHtml) Ondblclick(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondblclick", value)
	return d
}

/*
Ondrag -
*/
func (d *DlTagHtml) Ondrag(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondrag", value)
	return d
}

/*
Ondragend -
*/
func (d *DlTagHtml) Ondragend(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragend", value)
	return d
}

/*
Ondragenter -
*/
func (d *DlTagHtml) Ondragenter(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragenter", value)
	return d
}

/*
Ondragleave -
*/
func (d *DlTagHtml) Ondragleave(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragleave", value)
	return d
}

/*
Ondragover -
*/
func (d *DlTagHtml) Ondragover(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragover", value)
	return d
}

/*
Ondragstart -
*/
func (d *DlTagHtml) Ondragstart(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragstart", value)
	return d
}

/*
Ondrop -
*/
func (d *DlTagHtml) Ondrop(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondrop", value)
	return d
}

/*
Ondurationchange -
*/
func (d *DlTagHtml) Ondurationchange(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondurationchange", value)
	return d
}

/*
Onemptied -
*/
func (d *DlTagHtml) Onemptied(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onemptied", value)
	return d
}

/*
Onended -
*/
func (d *DlTagHtml) Onended(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onended", value)
	return d
}

/*
Onfocus -
*/
func (d *DlTagHtml) Onfocus(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onfocus", value)
	return d
}

/*
Oninput -
*/
func (d *DlTagHtml) Oninput(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oninput", value)
	return d
}

/*
Oninvalid -
*/
func (d *DlTagHtml) Oninvalid(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oninvalid", value)
	return d
}

/*
Onkeydown -
*/
func (d *DlTagHtml) Onkeydown(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onkeydown", value)
	return d
}

/*
Onkeypress -
*/
func (d *DlTagHtml) Onkeypress(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onkeypress", value)
	return d
}

/*
Onkeyup -
*/
func (d *DlTagHtml) Onkeyup(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onkeyup", value)
	return d
}

/*
Onloadeddata -
*/
func (d *DlTagHtml) Onloadeddata(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onloadeddata", value)
	return d
}

/*
Onloadedmetadata -
*/
func (d *DlTagHtml) Onloadedmetadata(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onloadedmetadata", value)
	return d
}

/*
Onloadstart -
*/
func (d *DlTagHtml) Onloadstart(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onloadstart", value)
	return d
}

/*
Onmousedown -
*/
func (d *DlTagHtml) Onmousedown(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmousedown", value)
	return d
}

/*
Onmouseenter -
*/
func (d *DlTagHtml) Onmouseenter(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseenter", value)
	return d
}

/*
Onmouseleave -
*/
func (d *DlTagHtml) Onmouseleave(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseleave", value)
	return d
}

/*
Onmousemove -
*/
func (d *DlTagHtml) Onmousemove(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmousemove", value)
	return d
}

/*
Onmouseout -
*/
func (d *DlTagHtml) Onmouseout(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseout", value)
	return d
}

/*
Onmouseover -
*/
func (d *DlTagHtml) Onmouseover(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseover", value)
	return d
}

/*
Onmouseup -
*/
func (d *DlTagHtml) Onmouseup(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseup", value)
	return d
}

/*
Onmousewheel -
*/
func (d *DlTagHtml) Onmousewheel(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmousewheel", value)
	return d
}

/*
Onpause -
*/
func (d *DlTagHtml) Onpause(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpause", value)
	return d
}

/*
Onplay -
*/
func (d *DlTagHtml) Onplay(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onplay", value)
	return d
}

/*
Onplaying -
*/
func (d *DlTagHtml) Onplaying(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onplaying", value)
	return d
}

/*
Onprogress -
*/
func (d *DlTagHtml) Onprogress(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onprogress", value)
	return d
}

/*
Onratechange -
*/
func (d *DlTagHtml) Onratechange(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onratechange", value)
	return d
}

/*
Onreset -
*/
func (d *DlTagHtml) Onreset(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onreset", value)
	return d
}

/*
Onscroll -
*/
func (d *DlTagHtml) Onscroll(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onscroll", value)
	return d
}

/*
Onseeked -
*/
func (d *DlTagHtml) Onseeked(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onseeked", value)
	return d
}

/*
Onseeking -
*/
func (d *DlTagHtml) Onseeking(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onseeking", value)
	return d
}

/*
Onselect -
*/
func (d *DlTagHtml) Onselect(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onselect", value)
	return d
}

/*
Onshow -
*/
func (d *DlTagHtml) Onshow(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onshow", value)
	return d
}

/*
Onsort -
*/
func (d *DlTagHtml) Onsort(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onsort", value)
	return d
}

/*
Onstalled -
*/
func (d *DlTagHtml) Onstalled(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onstalled", value)
	return d
}

/*
Onsubmit -
*/
func (d *DlTagHtml) Onsubmit(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onsubmit", value)
	return d
}

/*
Onsuspend -
*/
func (d *DlTagHtml) Onsuspend(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onsuspend", value)
	return d
}

/*
Ontimeupdate -
*/
func (d *DlTagHtml) Ontimeupdate(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ontimeupdate", value)
	return d
}

/*
Ontoggle -
*/
func (d *DlTagHtml) Ontoggle(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ontoggle", value)
	return d
}

/*
Onvolumechange -
*/
func (d *DlTagHtml) Onvolumechange(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onvolumechange", value)
	return d
}

/*
Onwaiting -
*/
func (d *DlTagHtml) Onwaiting(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onwaiting", value)
	return d
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (d *DlTagHtml) Onafterprint(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onafterprint", value)
	return d
}

/*
Onbeforeprint -
*/
func (d *DlTagHtml) Onbeforeprint(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onbeforeprint", value)
	return d
}

/*
Onbeforeunload -
*/
func (d *DlTagHtml) Onbeforeunload(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onbeforeunload", value)
	return d
}

/*
Onerror -
*/
func (d *DlTagHtml) Onerror(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onerror", value)
	return d
}

/*
Onhashchange -
*/
func (d *DlTagHtml) Onhashchange(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onhashchange", value)
	return d
}

/*
Onload -
*/
func (d *DlTagHtml) Onload(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onload", value)
	return d
}

/*
Onmessage -
*/
func (d *DlTagHtml) Onmessage(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmessage", value)
	return d
}

/*
Onoffline -
*/
func (d *DlTagHtml) Onoffline(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onoffline", value)
	return d
}

/*
Ononline -
*/
func (d *DlTagHtml) Ononline(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ononline", value)
	return d
}

/*
Onpagehide -
*/
func (d *DlTagHtml) Onpagehide(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpagehide", value)
	return d
}

/*
Onpageshow -
*/
func (d *DlTagHtml) Onpageshow(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpageshow", value)
	return d
}

/*
Onpopstate -
*/
func (d *DlTagHtml) Onpopstate(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpopstate", value)
	return d
}

/*
Onresize -
*/
func (d *DlTagHtml) Onresize(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onresize", value)
	return d
}

/*
Onstorage -
*/
func (d *DlTagHtml) Onstorage(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onstorage", value)
	return d
}

/*
Onunload -
*/
func (d *DlTagHtml) Onunload(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onunload", value)
	return d
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (d *DlTagHtml) AccessKey(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("accessKey", value)
	return d
}

/*
Aria -
*/
func (d *DlTagHtml) Aria(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria", value)
	return d
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (d *DlTagHtml) Autocapitalize(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("autocapitalize", value)
	return d
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (d *DlTagHtml) Autofocus(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("autofocus", value)
	return d
}

/*
Class -
*/
func (d *DlTagHtml) Class(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("class", value)
	return d
}

/*
Contenteditable -
*/
func (d *DlTagHtml) Contenteditable(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("contenteditable", value)
	return d
}

/*
Data -
*/
func (d *DlTagHtml) Data(name, value string) *DlTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	d.registerAttribute(dataName, value)
	return d
}

/*
Dir -
*/
func (d *DlTagHtml) Dir(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("dir", value)
	return d
}

/*
Draggable -
*/
func (d *DlTagHtml) Draggable(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("draggable", value)
	return d
}

/*
EnterKeyHint -
*/
func (d *DlTagHtml) EnterKeyHint(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("enterKeyHint", value)
	return d
}

/*
ExportParts -
*/
func (d *DlTagHtml) ExportParts(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("exportParts", value)
	return d
}

/*
Hidden -
*/
func (d *DlTagHtml) Hidden(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("hidden", value)
	return d
}

/*
Id -
*/
func (d *DlTagHtml) Id(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("id", value)
	return d
}

/*
Inert -
*/
func (d *DlTagHtml) Inert(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("inert", value)
	return d
}

/*
InputMode -
*/
func (d *DlTagHtml) InputMode(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("inputMode", value)
	return d
}

/*
Is -
*/
func (d *DlTagHtml) Is(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("is", value)
	return d
}

/*
ItemId -
*/
func (d *DlTagHtml) ItemId(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemId", value)
	return d
}

/*
ItemProp -
*/
func (d *DlTagHtml) ItemProp(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemProp", value)
	return d
}

/*
ItemRef -
*/
func (d *DlTagHtml) ItemRef(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemRef", value)
	return d
}

/*
ItemScope -
*/
func (d *DlTagHtml) ItemScope(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemScope", value)
	return d
}

/*
ItemType -
*/
func (d *DlTagHtml) ItemType(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemType", value)
	return d
}

/*
Lang -
*/
func (d *DlTagHtml) Lang(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("lang", value)
	return d
}

/*
Nonce -
*/
func (d *DlTagHtml) Nonce(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("nonce", value)
	return d
}

/*
Part -
*/
func (d *DlTagHtml) Part(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("part", value)
	return d
}

/*
Popover -
*/
func (d *DlTagHtml) Popover() *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("popover", "")
	return d
}

/*
Role -
*/
func (d *DlTagHtml) Role(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("role", value)
	return d
}

/*
Slot -
*/
func (d *DlTagHtml) Slot(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("slot", value)
	return d
}

/*
Spellcheck -
*/
func (d *DlTagHtml) Spellcheck(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("spellcheck", value)
	return d
}

/*
Style -
*/
func (d *DlTagHtml) Style(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("style", value)
	return d
}

/*
Tabindex -
*/
func (d *DlTagHtml) Tabindex(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("tabindex", value)
	return d
}

/*
Title -
*/
func (d *DlTagHtml) Title(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("title", value)
	return d
}

/*
Translate -
*/
func (d *DlTagHtml) Translate(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("translate", value)
	return d
}

/*
VirtualKeyBoardPolicy -
*/
func (d *DlTagHtml) VirtualKeyBoardPolicy(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("virtualKeyBoardPolicy", value)
	return d
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (d *DlTagHtml) AriaAtomic(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-atomic", value)
	return d
}

/*
AriaBusy -
*/
func (d *DlTagHtml) AriaBusy(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-busy", value)
	return d
}

/*
AriaControls -
*/
func (d *DlTagHtml) AriaControls(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-controls", value)
	return d
}

/*
AriaCurrent -
*/
func (d *DlTagHtml) AriaCurrent(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-current", value)
	return d
}

/*
AriaDescribedby -
*/
func (d *DlTagHtml) AriaDescribedby(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-describedby", value)
	return d
}

/*
AriaDescription -
*/
func (d *DlTagHtml) AriaDescription(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-description", value)
	return d
}

/*
AriaDetails -
*/
func (d *DlTagHtml) AriaDetails(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-details", value)
	return d
}

/*
AriaDisabled -
*/
func (d *DlTagHtml) AriaDisabled(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-disabled", value)
	return d
}

/*
AriaDropeffect -
*/
func (d *DlTagHtml) AriaDropeffect(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-dropeffect", value)
	return d
}

/*
AriaErrormessage -
*/
func (d *DlTagHtml) AriaErrormessage(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-errormessage", value)
	return d
}

/*
AriaFlowto -
*/
func (d *DlTagHtml) AriaFlowto(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-flowto", value)
	return d
}

/*
AriaGrabbed -
*/
func (d *DlTagHtml) AriaGrabbed(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-grabbed", value)
	return d
}

/*
AriaHaspopup -
*/
func (d *DlTagHtml) AriaHaspopup(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-haspopup", value)
	return d
}

/*
AriaHidden -
*/
func (d *DlTagHtml) AriaHidden(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-hidden", value)
	return d
}

/*
AriaInvalid -
*/
func (d *DlTagHtml) AriaInvalid(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-invalid", value)
	return d
}

/*
AriaKeyshortcuts -
*/
func (d *DlTagHtml) AriaKeyshortcuts(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-keyshortcuts", value)
	return d
}

/*
AriaLabel -
*/
func (d *DlTagHtml) AriaLabel(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-label", value)
	return d
}

/*
AriaLabelledby -
*/
func (d *DlTagHtml) AriaLabelledby(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-labelledby", value)
	return d
}

/*
AriaLive -
*/
func (d *DlTagHtml) AriaLive(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-live", value)
	return d
}

/*
AriaOwns -
*/
func (d *DlTagHtml) AriaOwns(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-owns", value)
	return d
}

/*
AriaRelevant -
*/
func (d *DlTagHtml) AriaRelevant(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-relevant", value)
	return d
}

/*
AriaRoledescription -
*/
func (d *DlTagHtml) AriaRoledescription(value string) *DlTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-roledescription", value)
	return d
}
