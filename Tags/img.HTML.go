// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func ImgHtml() *ImgTagHtml {
	return &ImgTagHtml{
		tag: &tag{
			name:                "img",
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

type ImgTagHtml struct {
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
func (i *ImgTagHtml) CustomAttribute(attributes ...*Attribute) *ImgTagHtml {
	i.registerAttributes(attributes...)
	return i
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Alt -
*/
func (i *ImgTagHtml) Alt(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("alt", value)
	return i
}

/*
Crossorigin -
*/
func (i *ImgTagHtml) Crossorigin(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("crossorigin", value)
	return i
}

/*
Decoding -
*/
func (i *ImgTagHtml) Decoding(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("decoding", value)
	return i
}

/*
Elementtiming -
*/
func (i *ImgTagHtml) Elementtiming(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("elementtiming", value)
	return i
}

/*
Fetchpriority -
*/
func (i *ImgTagHtml) Fetchpriority(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("fetchpriority", value)
	return i
}

/*
Height -
*/
func (i *ImgTagHtml) Height(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("height", value)
	return i
}

/*
Ismap -
*/
func (i *ImgTagHtml) Ismap() *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ismap", "")
	return i
}

/*
Loading -
*/
func (i *ImgTagHtml) Loading(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("loading", value)
	return i
}

/*
Referrerpolicy -
*/
func (i *ImgTagHtml) Referrerpolicy(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("referrerpolicy", value)
	return i
}

/*
Sizes -
*/
func (i *ImgTagHtml) Sizes(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("sizes", value)
	return i
}

/*
Src - Specifies the URL of an image.
Specifies the URL of an image.
*/
func (i *ImgTagHtml) Src(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("src", value)
	return i
}

/*
Srcset -
*/
func (i *ImgTagHtml) Srcset(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("srcset", value)
	return i
}

/*
Usemap -
*/
func (i *ImgTagHtml) Usemap(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("usemap", value)
	return i
}

/*
Width -
*/
func (i *ImgTagHtml) Width(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("width", value)
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
func (i *ImgTagHtml) Onabort(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onabort", value)
	return i
}

/*
Onautocomplete -
*/
func (i *ImgTagHtml) Onautocomplete(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocomplete", value)
	return i
}

/*
Onautocompleteerror -
*/
func (i *ImgTagHtml) Onautocompleteerror(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocompleteerror", value)
	return i
}

/*
Onblur -
*/
func (i *ImgTagHtml) Onblur(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onblur", value)
	return i
}

/*
Oncancel -
*/
func (i *ImgTagHtml) Oncancel(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncancel", value)
	return i
}

/*
Oncanplay -
*/
func (i *ImgTagHtml) Oncanplay(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplay", value)
	return i
}

/*
Oncanplaythrough -
*/
func (i *ImgTagHtml) Oncanplaythrough(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplaythrough", value)
	return i
}

/*
Onchange -
*/
func (i *ImgTagHtml) Onchange(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onchange", value)
	return i
}

/*
Onclick -
*/
func (i *ImgTagHtml) Onclick(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclick", value)
	return i
}

/*
Onclose -
*/
func (i *ImgTagHtml) Onclose(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclose", value)
	return i
}

/*
Oncontextmenu -
*/
func (i *ImgTagHtml) Oncontextmenu(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncontextmenu", value)
	return i
}

/*
Oncuechange -
*/
func (i *ImgTagHtml) Oncuechange(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncuechange", value)
	return i
}

/*
Ondblclick -
*/
func (i *ImgTagHtml) Ondblclick(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondblclick", value)
	return i
}

/*
Ondrag -
*/
func (i *ImgTagHtml) Ondrag(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrag", value)
	return i
}

/*
Ondragend -
*/
func (i *ImgTagHtml) Ondragend(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragend", value)
	return i
}

/*
Ondragenter -
*/
func (i *ImgTagHtml) Ondragenter(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragenter", value)
	return i
}

/*
Ondragleave -
*/
func (i *ImgTagHtml) Ondragleave(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragleave", value)
	return i
}

/*
Ondragover -
*/
func (i *ImgTagHtml) Ondragover(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragover", value)
	return i
}

/*
Ondragstart -
*/
func (i *ImgTagHtml) Ondragstart(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragstart", value)
	return i
}

/*
Ondrop -
*/
func (i *ImgTagHtml) Ondrop(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrop", value)
	return i
}

/*
Ondurationchange -
*/
func (i *ImgTagHtml) Ondurationchange(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondurationchange", value)
	return i
}

/*
Onemptied -
*/
func (i *ImgTagHtml) Onemptied(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onemptied", value)
	return i
}

/*
Onended -
*/
func (i *ImgTagHtml) Onended(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onended", value)
	return i
}

/*
Onfocus -
*/
func (i *ImgTagHtml) Onfocus(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onfocus", value)
	return i
}

/*
Oninput -
*/
func (i *ImgTagHtml) Oninput(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninput", value)
	return i
}

/*
Oninvalid -
*/
func (i *ImgTagHtml) Oninvalid(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninvalid", value)
	return i
}

/*
Onkeydown -
*/
func (i *ImgTagHtml) Onkeydown(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeydown", value)
	return i
}

/*
Onkeypress -
*/
func (i *ImgTagHtml) Onkeypress(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeypress", value)
	return i
}

/*
Onkeyup -
*/
func (i *ImgTagHtml) Onkeyup(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeyup", value)
	return i
}

/*
Onloadeddata -
*/
func (i *ImgTagHtml) Onloadeddata(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadeddata", value)
	return i
}

/*
Onloadedmetadata -
*/
func (i *ImgTagHtml) Onloadedmetadata(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadedmetadata", value)
	return i
}

/*
Onloadstart -
*/
func (i *ImgTagHtml) Onloadstart(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadstart", value)
	return i
}

/*
Onmousedown -
*/
func (i *ImgTagHtml) Onmousedown(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousedown", value)
	return i
}

/*
Onmouseenter -
*/
func (i *ImgTagHtml) Onmouseenter(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseenter", value)
	return i
}

/*
Onmouseleave -
*/
func (i *ImgTagHtml) Onmouseleave(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseleave", value)
	return i
}

/*
Onmousemove -
*/
func (i *ImgTagHtml) Onmousemove(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousemove", value)
	return i
}

/*
Onmouseout -
*/
func (i *ImgTagHtml) Onmouseout(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseout", value)
	return i
}

/*
Onmouseover -
*/
func (i *ImgTagHtml) Onmouseover(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseover", value)
	return i
}

/*
Onmouseup -
*/
func (i *ImgTagHtml) Onmouseup(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseup", value)
	return i
}

/*
Onmousewheel -
*/
func (i *ImgTagHtml) Onmousewheel(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousewheel", value)
	return i
}

/*
Onpause -
*/
func (i *ImgTagHtml) Onpause(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpause", value)
	return i
}

/*
Onplay -
*/
func (i *ImgTagHtml) Onplay(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplay", value)
	return i
}

/*
Onplaying -
*/
func (i *ImgTagHtml) Onplaying(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplaying", value)
	return i
}

/*
Onprogress -
*/
func (i *ImgTagHtml) Onprogress(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onprogress", value)
	return i
}

/*
Onratechange -
*/
func (i *ImgTagHtml) Onratechange(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onratechange", value)
	return i
}

/*
Onreset -
*/
func (i *ImgTagHtml) Onreset(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onreset", value)
	return i
}

/*
Onscroll -
*/
func (i *ImgTagHtml) Onscroll(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onscroll", value)
	return i
}

/*
Onseeked -
*/
func (i *ImgTagHtml) Onseeked(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeked", value)
	return i
}

/*
Onseeking -
*/
func (i *ImgTagHtml) Onseeking(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeking", value)
	return i
}

/*
Onselect -
*/
func (i *ImgTagHtml) Onselect(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onselect", value)
	return i
}

/*
Onshow -
*/
func (i *ImgTagHtml) Onshow(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onshow", value)
	return i
}

/*
Onsort -
*/
func (i *ImgTagHtml) Onsort(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsort", value)
	return i
}

/*
Onstalled -
*/
func (i *ImgTagHtml) Onstalled(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstalled", value)
	return i
}

/*
Onsubmit -
*/
func (i *ImgTagHtml) Onsubmit(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsubmit", value)
	return i
}

/*
Onsuspend -
*/
func (i *ImgTagHtml) Onsuspend(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsuspend", value)
	return i
}

/*
Ontimeupdate -
*/
func (i *ImgTagHtml) Ontimeupdate(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontimeupdate", value)
	return i
}

/*
Ontoggle -
*/
func (i *ImgTagHtml) Ontoggle(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontoggle", value)
	return i
}

/*
Onvolumechange -
*/
func (i *ImgTagHtml) Onvolumechange(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onvolumechange", value)
	return i
}

/*
Onwaiting -
*/
func (i *ImgTagHtml) Onwaiting(value string) *ImgTagHtml {
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
func (i *ImgTagHtml) Onafterprint(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onafterprint", value)
	return i
}

/*
Onbeforeprint -
*/
func (i *ImgTagHtml) Onbeforeprint(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeprint", value)
	return i
}

/*
Onbeforeunload -
*/
func (i *ImgTagHtml) Onbeforeunload(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeunload", value)
	return i
}

/*
Onerror -
*/
func (i *ImgTagHtml) Onerror(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onerror", value)
	return i
}

/*
Onhashchange -
*/
func (i *ImgTagHtml) Onhashchange(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onhashchange", value)
	return i
}

/*
Onload -
*/
func (i *ImgTagHtml) Onload(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onload", value)
	return i
}

/*
Onmessage -
*/
func (i *ImgTagHtml) Onmessage(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmessage", value)
	return i
}

/*
Onoffline -
*/
func (i *ImgTagHtml) Onoffline(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onoffline", value)
	return i
}

/*
Ononline -
*/
func (i *ImgTagHtml) Ononline(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ononline", value)
	return i
}

/*
Onpagehide -
*/
func (i *ImgTagHtml) Onpagehide(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpagehide", value)
	return i
}

/*
Onpageshow -
*/
func (i *ImgTagHtml) Onpageshow(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpageshow", value)
	return i
}

/*
Onpopstate -
*/
func (i *ImgTagHtml) Onpopstate(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpopstate", value)
	return i
}

/*
Onresize -
*/
func (i *ImgTagHtml) Onresize(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onresize", value)
	return i
}

/*
Onstorage -
*/
func (i *ImgTagHtml) Onstorage(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstorage", value)
	return i
}

/*
Onunload -
*/
func (i *ImgTagHtml) Onunload(value string) *ImgTagHtml {
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
func (i *ImgTagHtml) AccessKey(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("accessKey", value)
	return i
}

/*
Aria -
*/
func (i *ImgTagHtml) Aria(value string) *ImgTagHtml {
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
func (i *ImgTagHtml) Autocapitalize(value string) *ImgTagHtml {
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
func (i *ImgTagHtml) Autofocus(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("autofocus", value)
	return i
}

/*
Class -
*/
func (i *ImgTagHtml) Class(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("class", value)
	return i
}

/*
Contenteditable -
*/
func (i *ImgTagHtml) Contenteditable(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("contenteditable", value)
	return i
}

/*
Data -
*/
func (i *ImgTagHtml) Data(name, value string) *ImgTagHtml {
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
func (i *ImgTagHtml) Dir(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("dir", value)
	return i
}

/*
Draggable -
*/
func (i *ImgTagHtml) Draggable(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("draggable", value)
	return i
}

/*
EnterKeyHint -
*/
func (i *ImgTagHtml) EnterKeyHint(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("enterKeyHint", value)
	return i
}

/*
ExportParts -
*/
func (i *ImgTagHtml) ExportParts(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("exportParts", value)
	return i
}

/*
Hidden -
*/
func (i *ImgTagHtml) Hidden(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("hidden", value)
	return i
}

/*
Id -
*/
func (i *ImgTagHtml) Id(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("id", value)
	return i
}

/*
Inert -
*/
func (i *ImgTagHtml) Inert(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inert", value)
	return i
}

/*
InputMode -
*/
func (i *ImgTagHtml) InputMode(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inputMode", value)
	return i
}

/*
Is -
*/
func (i *ImgTagHtml) Is(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("is", value)
	return i
}

/*
ItemId -
*/
func (i *ImgTagHtml) ItemId(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemId", value)
	return i
}

/*
ItemProp -
*/
func (i *ImgTagHtml) ItemProp(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemProp", value)
	return i
}

/*
ItemRef -
*/
func (i *ImgTagHtml) ItemRef(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemRef", value)
	return i
}

/*
ItemScope -
*/
func (i *ImgTagHtml) ItemScope(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemScope", value)
	return i
}

/*
ItemType -
*/
func (i *ImgTagHtml) ItemType(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemType", value)
	return i
}

/*
Lang -
*/
func (i *ImgTagHtml) Lang(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("lang", value)
	return i
}

/*
Nonce -
*/
func (i *ImgTagHtml) Nonce(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("nonce", value)
	return i
}

/*
Part -
*/
func (i *ImgTagHtml) Part(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("part", value)
	return i
}

/*
Popover -
*/
func (i *ImgTagHtml) Popover() *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("popover", "")
	return i
}

/*
Role -
*/
func (i *ImgTagHtml) Role(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("role", value)
	return i
}

/*
Slot -
*/
func (i *ImgTagHtml) Slot(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("slot", value)
	return i
}

/*
Spellcheck -
*/
func (i *ImgTagHtml) Spellcheck(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("spellcheck", value)
	return i
}

/*
Style -
*/
func (i *ImgTagHtml) Style(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("style", value)
	return i
}

/*
Tabindex -
*/
func (i *ImgTagHtml) Tabindex(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("tabindex", value)
	return i
}

/*
Title -
*/
func (i *ImgTagHtml) Title(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("title", value)
	return i
}

/*
Translate -
*/
func (i *ImgTagHtml) Translate(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("translate", value)
	return i
}

/*
VirtualKeyBoardPolicy -
*/
func (i *ImgTagHtml) VirtualKeyBoardPolicy(value string) *ImgTagHtml {
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
func (i *ImgTagHtml) AriaAtomic(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-atomic", value)
	return i
}

/*
AriaBusy -
*/
func (i *ImgTagHtml) AriaBusy(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-busy", value)
	return i
}

/*
AriaControls -
*/
func (i *ImgTagHtml) AriaControls(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-controls", value)
	return i
}

/*
AriaCurrent -
*/
func (i *ImgTagHtml) AriaCurrent(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-current", value)
	return i
}

/*
AriaDescribedby -
*/
func (i *ImgTagHtml) AriaDescribedby(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-describedby", value)
	return i
}

/*
AriaDescription -
*/
func (i *ImgTagHtml) AriaDescription(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-description", value)
	return i
}

/*
AriaDetails -
*/
func (i *ImgTagHtml) AriaDetails(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-details", value)
	return i
}

/*
AriaDisabled -
*/
func (i *ImgTagHtml) AriaDisabled(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-disabled", value)
	return i
}

/*
AriaDropeffect -
*/
func (i *ImgTagHtml) AriaDropeffect(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-dropeffect", value)
	return i
}

/*
AriaErrormessage -
*/
func (i *ImgTagHtml) AriaErrormessage(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-errormessage", value)
	return i
}

/*
AriaFlowto -
*/
func (i *ImgTagHtml) AriaFlowto(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-flowto", value)
	return i
}

/*
AriaGrabbed -
*/
func (i *ImgTagHtml) AriaGrabbed(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-grabbed", value)
	return i
}

/*
AriaHaspopup -
*/
func (i *ImgTagHtml) AriaHaspopup(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-haspopup", value)
	return i
}

/*
AriaHidden -
*/
func (i *ImgTagHtml) AriaHidden(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-hidden", value)
	return i
}

/*
AriaInvalid -
*/
func (i *ImgTagHtml) AriaInvalid(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-invalid", value)
	return i
}

/*
AriaKeyshortcuts -
*/
func (i *ImgTagHtml) AriaKeyshortcuts(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-keyshortcuts", value)
	return i
}

/*
AriaLabel -
*/
func (i *ImgTagHtml) AriaLabel(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-label", value)
	return i
}

/*
AriaLabelledby -
*/
func (i *ImgTagHtml) AriaLabelledby(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-labelledby", value)
	return i
}

/*
AriaLive -
*/
func (i *ImgTagHtml) AriaLive(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-live", value)
	return i
}

/*
AriaOwns -
*/
func (i *ImgTagHtml) AriaOwns(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-owns", value)
	return i
}

/*
AriaRelevant -
*/
func (i *ImgTagHtml) AriaRelevant(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-relevant", value)
	return i
}

/*
AriaRoledescription -
*/
func (i *ImgTagHtml) AriaRoledescription(value string) *ImgTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-roledescription", value)
	return i
}
