// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func EmbedHtml() *EmbedTagHtml {
	return &EmbedTagHtml{
		tag: &tag{
			name:                "embed",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.SelfClosingType,
			namespace:           spec.HTML,
			children:            nil,
			textContent:         "",
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

type EmbedTagHtml struct {
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
func (e *EmbedTagHtml) CustomAttribute(attributes ...*Attribute) *EmbedTagHtml {
	e.registerAttributes(attributes...)
	return e
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Height -
*/
func (e *EmbedTagHtml) Height(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("height", value)
	return e
}

/*
Src - Specifies the URL of an image.
Specifies the URL of an image.
*/
func (e *EmbedTagHtml) Src(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("src", value)
	return e
}

/*
Type - Specifies the type of an <input> element.
Specifies the type of an <input> element.
*/
func (e *EmbedTagHtml) Type(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("type", value)
	return e
}

/*
Width -
*/
func (e *EmbedTagHtml) Width(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("width", value)
	return e
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (e *EmbedTagHtml) Onabort(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onabort", value)
	return e
}

/*
Onautocomplete -
*/
func (e *EmbedTagHtml) Onautocomplete(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onautocomplete", value)
	return e
}

/*
Onautocompleteerror -
*/
func (e *EmbedTagHtml) Onautocompleteerror(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onautocompleteerror", value)
	return e
}

/*
Onblur -
*/
func (e *EmbedTagHtml) Onblur(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onblur", value)
	return e
}

/*
Oncancel -
*/
func (e *EmbedTagHtml) Oncancel(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("oncancel", value)
	return e
}

/*
Oncanplay -
*/
func (e *EmbedTagHtml) Oncanplay(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("oncanplay", value)
	return e
}

/*
Oncanplaythrough -
*/
func (e *EmbedTagHtml) Oncanplaythrough(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("oncanplaythrough", value)
	return e
}

/*
Onchange -
*/
func (e *EmbedTagHtml) Onchange(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onchange", value)
	return e
}

/*
Onclick -
*/
func (e *EmbedTagHtml) Onclick(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onclick", value)
	return e
}

/*
Onclose -
*/
func (e *EmbedTagHtml) Onclose(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onclose", value)
	return e
}

/*
Oncontextmenu -
*/
func (e *EmbedTagHtml) Oncontextmenu(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("oncontextmenu", value)
	return e
}

/*
Oncuechange -
*/
func (e *EmbedTagHtml) Oncuechange(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("oncuechange", value)
	return e
}

/*
Ondblclick -
*/
func (e *EmbedTagHtml) Ondblclick(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondblclick", value)
	return e
}

/*
Ondrag -
*/
func (e *EmbedTagHtml) Ondrag(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondrag", value)
	return e
}

/*
Ondragend -
*/
func (e *EmbedTagHtml) Ondragend(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondragend", value)
	return e
}

/*
Ondragenter -
*/
func (e *EmbedTagHtml) Ondragenter(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondragenter", value)
	return e
}

/*
Ondragleave -
*/
func (e *EmbedTagHtml) Ondragleave(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondragleave", value)
	return e
}

/*
Ondragover -
*/
func (e *EmbedTagHtml) Ondragover(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondragover", value)
	return e
}

/*
Ondragstart -
*/
func (e *EmbedTagHtml) Ondragstart(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondragstart", value)
	return e
}

/*
Ondrop -
*/
func (e *EmbedTagHtml) Ondrop(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondrop", value)
	return e
}

/*
Ondurationchange -
*/
func (e *EmbedTagHtml) Ondurationchange(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ondurationchange", value)
	return e
}

/*
Onemptied -
*/
func (e *EmbedTagHtml) Onemptied(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onemptied", value)
	return e
}

/*
Onended -
*/
func (e *EmbedTagHtml) Onended(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onended", value)
	return e
}

/*
Onfocus -
*/
func (e *EmbedTagHtml) Onfocus(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onfocus", value)
	return e
}

/*
Oninput -
*/
func (e *EmbedTagHtml) Oninput(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("oninput", value)
	return e
}

/*
Oninvalid -
*/
func (e *EmbedTagHtml) Oninvalid(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("oninvalid", value)
	return e
}

/*
Onkeydown -
*/
func (e *EmbedTagHtml) Onkeydown(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onkeydown", value)
	return e
}

/*
Onkeypress -
*/
func (e *EmbedTagHtml) Onkeypress(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onkeypress", value)
	return e
}

/*
Onkeyup -
*/
func (e *EmbedTagHtml) Onkeyup(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onkeyup", value)
	return e
}

/*
Onloadeddata -
*/
func (e *EmbedTagHtml) Onloadeddata(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onloadeddata", value)
	return e
}

/*
Onloadedmetadata -
*/
func (e *EmbedTagHtml) Onloadedmetadata(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onloadedmetadata", value)
	return e
}

/*
Onloadstart -
*/
func (e *EmbedTagHtml) Onloadstart(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onloadstart", value)
	return e
}

/*
Onmousedown -
*/
func (e *EmbedTagHtml) Onmousedown(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmousedown", value)
	return e
}

/*
Onmouseenter -
*/
func (e *EmbedTagHtml) Onmouseenter(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmouseenter", value)
	return e
}

/*
Onmouseleave -
*/
func (e *EmbedTagHtml) Onmouseleave(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmouseleave", value)
	return e
}

/*
Onmousemove -
*/
func (e *EmbedTagHtml) Onmousemove(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmousemove", value)
	return e
}

/*
Onmouseout -
*/
func (e *EmbedTagHtml) Onmouseout(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmouseout", value)
	return e
}

/*
Onmouseover -
*/
func (e *EmbedTagHtml) Onmouseover(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmouseover", value)
	return e
}

/*
Onmouseup -
*/
func (e *EmbedTagHtml) Onmouseup(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmouseup", value)
	return e
}

/*
Onmousewheel -
*/
func (e *EmbedTagHtml) Onmousewheel(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmousewheel", value)
	return e
}

/*
Onpause -
*/
func (e *EmbedTagHtml) Onpause(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onpause", value)
	return e
}

/*
Onplay -
*/
func (e *EmbedTagHtml) Onplay(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onplay", value)
	return e
}

/*
Onplaying -
*/
func (e *EmbedTagHtml) Onplaying(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onplaying", value)
	return e
}

/*
Onprogress -
*/
func (e *EmbedTagHtml) Onprogress(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onprogress", value)
	return e
}

/*
Onratechange -
*/
func (e *EmbedTagHtml) Onratechange(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onratechange", value)
	return e
}

/*
Onreset -
*/
func (e *EmbedTagHtml) Onreset(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onreset", value)
	return e
}

/*
Onscroll -
*/
func (e *EmbedTagHtml) Onscroll(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onscroll", value)
	return e
}

/*
Onseeked -
*/
func (e *EmbedTagHtml) Onseeked(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onseeked", value)
	return e
}

/*
Onseeking -
*/
func (e *EmbedTagHtml) Onseeking(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onseeking", value)
	return e
}

/*
Onselect -
*/
func (e *EmbedTagHtml) Onselect(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onselect", value)
	return e
}

/*
Onshow -
*/
func (e *EmbedTagHtml) Onshow(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onshow", value)
	return e
}

/*
Onsort -
*/
func (e *EmbedTagHtml) Onsort(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onsort", value)
	return e
}

/*
Onstalled -
*/
func (e *EmbedTagHtml) Onstalled(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onstalled", value)
	return e
}

/*
Onsubmit -
*/
func (e *EmbedTagHtml) Onsubmit(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onsubmit", value)
	return e
}

/*
Onsuspend -
*/
func (e *EmbedTagHtml) Onsuspend(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onsuspend", value)
	return e
}

/*
Ontimeupdate -
*/
func (e *EmbedTagHtml) Ontimeupdate(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ontimeupdate", value)
	return e
}

/*
Ontoggle -
*/
func (e *EmbedTagHtml) Ontoggle(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ontoggle", value)
	return e
}

/*
Onvolumechange -
*/
func (e *EmbedTagHtml) Onvolumechange(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onvolumechange", value)
	return e
}

/*
Onwaiting -
*/
func (e *EmbedTagHtml) Onwaiting(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onwaiting", value)
	return e
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (e *EmbedTagHtml) Onafterprint(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onafterprint", value)
	return e
}

/*
Onbeforeprint -
*/
func (e *EmbedTagHtml) Onbeforeprint(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onbeforeprint", value)
	return e
}

/*
Onbeforeunload -
*/
func (e *EmbedTagHtml) Onbeforeunload(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onbeforeunload", value)
	return e
}

/*
Onerror -
*/
func (e *EmbedTagHtml) Onerror(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onerror", value)
	return e
}

/*
Onhashchange -
*/
func (e *EmbedTagHtml) Onhashchange(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onhashchange", value)
	return e
}

/*
Onload -
*/
func (e *EmbedTagHtml) Onload(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onload", value)
	return e
}

/*
Onmessage -
*/
func (e *EmbedTagHtml) Onmessage(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onmessage", value)
	return e
}

/*
Onoffline -
*/
func (e *EmbedTagHtml) Onoffline(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onoffline", value)
	return e
}

/*
Ononline -
*/
func (e *EmbedTagHtml) Ononline(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("ononline", value)
	return e
}

/*
Onpagehide -
*/
func (e *EmbedTagHtml) Onpagehide(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onpagehide", value)
	return e
}

/*
Onpageshow -
*/
func (e *EmbedTagHtml) Onpageshow(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onpageshow", value)
	return e
}

/*
Onpopstate -
*/
func (e *EmbedTagHtml) Onpopstate(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onpopstate", value)
	return e
}

/*
Onresize -
*/
func (e *EmbedTagHtml) Onresize(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onresize", value)
	return e
}

/*
Onstorage -
*/
func (e *EmbedTagHtml) Onstorage(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onstorage", value)
	return e
}

/*
Onunload -
*/
func (e *EmbedTagHtml) Onunload(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("onunload", value)
	return e
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (e *EmbedTagHtml) AccessKey(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("accessKey", value)
	return e
}

/*
Aria -
*/
func (e *EmbedTagHtml) Aria(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria", value)
	return e
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (e *EmbedTagHtml) Autocapitalize(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("autocapitalize", value)
	return e
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (e *EmbedTagHtml) Autofocus(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("autofocus", value)
	return e
}

/*
Class -
*/
func (e *EmbedTagHtml) Class(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("class", value)
	return e
}

/*
Contenteditable -
*/
func (e *EmbedTagHtml) Contenteditable(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("contenteditable", value)
	return e
}

/*
Data -
*/
func (e *EmbedTagHtml) Data(name, value string) *EmbedTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	e.registerAttribute(dataName, value)
	return e
}

/*
Dir -
*/
func (e *EmbedTagHtml) Dir(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("dir", value)
	return e
}

/*
Draggable -
*/
func (e *EmbedTagHtml) Draggable(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("draggable", value)
	return e
}

/*
EnterKeyHint -
*/
func (e *EmbedTagHtml) EnterKeyHint(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("enterKeyHint", value)
	return e
}

/*
ExportParts -
*/
func (e *EmbedTagHtml) ExportParts(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("exportParts", value)
	return e
}

/*
Hidden -
*/
func (e *EmbedTagHtml) Hidden(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("hidden", value)
	return e
}

/*
Id -
*/
func (e *EmbedTagHtml) Id(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("id", value)
	return e
}

/*
Inert -
*/
func (e *EmbedTagHtml) Inert(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("inert", value)
	return e
}

/*
InputMode -
*/
func (e *EmbedTagHtml) InputMode(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("inputMode", value)
	return e
}

/*
Is -
*/
func (e *EmbedTagHtml) Is(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("is", value)
	return e
}

/*
ItemId -
*/
func (e *EmbedTagHtml) ItemId(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("itemId", value)
	return e
}

/*
ItemProp -
*/
func (e *EmbedTagHtml) ItemProp(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("itemProp", value)
	return e
}

/*
ItemRef -
*/
func (e *EmbedTagHtml) ItemRef(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("itemRef", value)
	return e
}

/*
ItemScope -
*/
func (e *EmbedTagHtml) ItemScope(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("itemScope", value)
	return e
}

/*
ItemType -
*/
func (e *EmbedTagHtml) ItemType(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("itemType", value)
	return e
}

/*
Lang -
*/
func (e *EmbedTagHtml) Lang(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("lang", value)
	return e
}

/*
Nonce -
*/
func (e *EmbedTagHtml) Nonce(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("nonce", value)
	return e
}

/*
Part -
*/
func (e *EmbedTagHtml) Part(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("part", value)
	return e
}

/*
Popover -
*/
func (e *EmbedTagHtml) Popover() *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("popover", "")
	return e
}

/*
Role -
*/
func (e *EmbedTagHtml) Role(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("role", value)
	return e
}

/*
Slot -
*/
func (e *EmbedTagHtml) Slot(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("slot", value)
	return e
}

/*
Spellcheck -
*/
func (e *EmbedTagHtml) Spellcheck(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("spellcheck", value)
	return e
}

/*
Style -
*/
func (e *EmbedTagHtml) Style(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("style", value)
	return e
}

/*
Tabindex -
*/
func (e *EmbedTagHtml) Tabindex(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("tabindex", value)
	return e
}

/*
Title -
*/
func (e *EmbedTagHtml) Title(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("title", value)
	return e
}

/*
Translate -
*/
func (e *EmbedTagHtml) Translate(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("translate", value)
	return e
}

/*
VirtualKeyBoardPolicy -
*/
func (e *EmbedTagHtml) VirtualKeyBoardPolicy(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("virtualKeyBoardPolicy", value)
	return e
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (e *EmbedTagHtml) AriaAtomic(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-atomic", value)
	return e
}

/*
AriaBusy -
*/
func (e *EmbedTagHtml) AriaBusy(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-busy", value)
	return e
}

/*
AriaControls -
*/
func (e *EmbedTagHtml) AriaControls(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-controls", value)
	return e
}

/*
AriaCurrent -
*/
func (e *EmbedTagHtml) AriaCurrent(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-current", value)
	return e
}

/*
AriaDescribedby -
*/
func (e *EmbedTagHtml) AriaDescribedby(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-describedby", value)
	return e
}

/*
AriaDescription -
*/
func (e *EmbedTagHtml) AriaDescription(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-description", value)
	return e
}

/*
AriaDetails -
*/
func (e *EmbedTagHtml) AriaDetails(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-details", value)
	return e
}

/*
AriaDisabled -
*/
func (e *EmbedTagHtml) AriaDisabled(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-disabled", value)
	return e
}

/*
AriaDropeffect -
*/
func (e *EmbedTagHtml) AriaDropeffect(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-dropeffect", value)
	return e
}

/*
AriaErrormessage -
*/
func (e *EmbedTagHtml) AriaErrormessage(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-errormessage", value)
	return e
}

/*
AriaFlowto -
*/
func (e *EmbedTagHtml) AriaFlowto(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-flowto", value)
	return e
}

/*
AriaGrabbed -
*/
func (e *EmbedTagHtml) AriaGrabbed(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-grabbed", value)
	return e
}

/*
AriaHaspopup -
*/
func (e *EmbedTagHtml) AriaHaspopup(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-haspopup", value)
	return e
}

/*
AriaHidden -
*/
func (e *EmbedTagHtml) AriaHidden(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-hidden", value)
	return e
}

/*
AriaInvalid -
*/
func (e *EmbedTagHtml) AriaInvalid(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-invalid", value)
	return e
}

/*
AriaKeyshortcuts -
*/
func (e *EmbedTagHtml) AriaKeyshortcuts(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-keyshortcuts", value)
	return e
}

/*
AriaLabel -
*/
func (e *EmbedTagHtml) AriaLabel(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-label", value)
	return e
}

/*
AriaLabelledby -
*/
func (e *EmbedTagHtml) AriaLabelledby(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-labelledby", value)
	return e
}

/*
AriaLive -
*/
func (e *EmbedTagHtml) AriaLive(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-live", value)
	return e
}

/*
AriaOwns -
*/
func (e *EmbedTagHtml) AriaOwns(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-owns", value)
	return e
}

/*
AriaRelevant -
*/
func (e *EmbedTagHtml) AriaRelevant(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-relevant", value)
	return e
}

/*
AriaRoledescription -
*/
func (e *EmbedTagHtml) AriaRoledescription(value string) *EmbedTagHtml {
	if e.attributes == nil {
		e.attributes = []*Attribute{}
	}
	e.registerAttribute("aria-roledescription", value)
	return e
}
