// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func StyleHtml(tags []any) *StyleTagHtml {
	node := &StyleTagHtml{
		tag: &tag{
			name:                "style",
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

type StyleTagHtml struct {
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
func (s *StyleTagHtml) CustomAttribute(attributes ...*Attribute) *StyleTagHtml {
	s.registerAttributes(attributes...)
	return s
}

/*
Children - Method for nesting tags into parent tag
*/
func (s *StyleTagHtml) Children(tags ...any) *StyleTagHtml {
	return s.supportedChildrenCheck(tags)
}

func (s *StyleTagHtml) supportedChildrenCheck(tags []any) *StyleTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			s.registerChildren(TextContentSvg(val).getTag())
		case content:
			s.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				s.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return s
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Blocking -
*/
func (s *StyleTagHtml) Blocking(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("blocking", value)
	return s
}

/*
Media -
*/
func (s *StyleTagHtml) Media(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("media", value)
	return s
}

/*
Type - Specifies the type of an <input> element.
Specifies the type of an <input> element.
*/
func (s *StyleTagHtml) Type(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("type", value)
	return s
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (s *StyleTagHtml) Onabort(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onabort", value)
	return s
}

/*
Onautocomplete -
*/
func (s *StyleTagHtml) Onautocomplete(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocomplete", value)
	return s
}

/*
Onautocompleteerror -
*/
func (s *StyleTagHtml) Onautocompleteerror(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocompleteerror", value)
	return s
}

/*
Onblur -
*/
func (s *StyleTagHtml) Onblur(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onblur", value)
	return s
}

/*
Oncancel -
*/
func (s *StyleTagHtml) Oncancel(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncancel", value)
	return s
}

/*
Oncanplay -
*/
func (s *StyleTagHtml) Oncanplay(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplay", value)
	return s
}

/*
Oncanplaythrough -
*/
func (s *StyleTagHtml) Oncanplaythrough(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplaythrough", value)
	return s
}

/*
Onchange -
*/
func (s *StyleTagHtml) Onchange(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onchange", value)
	return s
}

/*
Onclick -
*/
func (s *StyleTagHtml) Onclick(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclick", value)
	return s
}

/*
Onclose -
*/
func (s *StyleTagHtml) Onclose(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclose", value)
	return s
}

/*
Oncontextmenu -
*/
func (s *StyleTagHtml) Oncontextmenu(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncontextmenu", value)
	return s
}

/*
Oncuechange -
*/
func (s *StyleTagHtml) Oncuechange(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncuechange", value)
	return s
}

/*
Ondblclick -
*/
func (s *StyleTagHtml) Ondblclick(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondblclick", value)
	return s
}

/*
Ondrag -
*/
func (s *StyleTagHtml) Ondrag(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrag", value)
	return s
}

/*
Ondragend -
*/
func (s *StyleTagHtml) Ondragend(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragend", value)
	return s
}

/*
Ondragenter -
*/
func (s *StyleTagHtml) Ondragenter(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragenter", value)
	return s
}

/*
Ondragleave -
*/
func (s *StyleTagHtml) Ondragleave(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragleave", value)
	return s
}

/*
Ondragover -
*/
func (s *StyleTagHtml) Ondragover(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragover", value)
	return s
}

/*
Ondragstart -
*/
func (s *StyleTagHtml) Ondragstart(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragstart", value)
	return s
}

/*
Ondrop -
*/
func (s *StyleTagHtml) Ondrop(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrop", value)
	return s
}

/*
Ondurationchange -
*/
func (s *StyleTagHtml) Ondurationchange(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondurationchange", value)
	return s
}

/*
Onemptied -
*/
func (s *StyleTagHtml) Onemptied(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onemptied", value)
	return s
}

/*
Onended -
*/
func (s *StyleTagHtml) Onended(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onended", value)
	return s
}

/*
Onfocus -
*/
func (s *StyleTagHtml) Onfocus(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onfocus", value)
	return s
}

/*
Oninput -
*/
func (s *StyleTagHtml) Oninput(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninput", value)
	return s
}

/*
Oninvalid -
*/
func (s *StyleTagHtml) Oninvalid(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninvalid", value)
	return s
}

/*
Onkeydown -
*/
func (s *StyleTagHtml) Onkeydown(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeydown", value)
	return s
}

/*
Onkeypress -
*/
func (s *StyleTagHtml) Onkeypress(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeypress", value)
	return s
}

/*
Onkeyup -
*/
func (s *StyleTagHtml) Onkeyup(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeyup", value)
	return s
}

/*
Onloadeddata -
*/
func (s *StyleTagHtml) Onloadeddata(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadeddata", value)
	return s
}

/*
Onloadedmetadata -
*/
func (s *StyleTagHtml) Onloadedmetadata(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadedmetadata", value)
	return s
}

/*
Onloadstart -
*/
func (s *StyleTagHtml) Onloadstart(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadstart", value)
	return s
}

/*
Onmousedown -
*/
func (s *StyleTagHtml) Onmousedown(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousedown", value)
	return s
}

/*
Onmouseenter -
*/
func (s *StyleTagHtml) Onmouseenter(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseenter", value)
	return s
}

/*
Onmouseleave -
*/
func (s *StyleTagHtml) Onmouseleave(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseleave", value)
	return s
}

/*
Onmousemove -
*/
func (s *StyleTagHtml) Onmousemove(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousemove", value)
	return s
}

/*
Onmouseout -
*/
func (s *StyleTagHtml) Onmouseout(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseout", value)
	return s
}

/*
Onmouseover -
*/
func (s *StyleTagHtml) Onmouseover(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseover", value)
	return s
}

/*
Onmouseup -
*/
func (s *StyleTagHtml) Onmouseup(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseup", value)
	return s
}

/*
Onmousewheel -
*/
func (s *StyleTagHtml) Onmousewheel(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousewheel", value)
	return s
}

/*
Onpause -
*/
func (s *StyleTagHtml) Onpause(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpause", value)
	return s
}

/*
Onplay -
*/
func (s *StyleTagHtml) Onplay(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplay", value)
	return s
}

/*
Onplaying -
*/
func (s *StyleTagHtml) Onplaying(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplaying", value)
	return s
}

/*
Onprogress -
*/
func (s *StyleTagHtml) Onprogress(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onprogress", value)
	return s
}

/*
Onratechange -
*/
func (s *StyleTagHtml) Onratechange(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onratechange", value)
	return s
}

/*
Onreset -
*/
func (s *StyleTagHtml) Onreset(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onreset", value)
	return s
}

/*
Onscroll -
*/
func (s *StyleTagHtml) Onscroll(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onscroll", value)
	return s
}

/*
Onseeked -
*/
func (s *StyleTagHtml) Onseeked(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeked", value)
	return s
}

/*
Onseeking -
*/
func (s *StyleTagHtml) Onseeking(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeking", value)
	return s
}

/*
Onselect -
*/
func (s *StyleTagHtml) Onselect(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onselect", value)
	return s
}

/*
Onshow -
*/
func (s *StyleTagHtml) Onshow(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onshow", value)
	return s
}

/*
Onsort -
*/
func (s *StyleTagHtml) Onsort(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsort", value)
	return s
}

/*
Onstalled -
*/
func (s *StyleTagHtml) Onstalled(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstalled", value)
	return s
}

/*
Onsubmit -
*/
func (s *StyleTagHtml) Onsubmit(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsubmit", value)
	return s
}

/*
Onsuspend -
*/
func (s *StyleTagHtml) Onsuspend(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsuspend", value)
	return s
}

/*
Ontimeupdate -
*/
func (s *StyleTagHtml) Ontimeupdate(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontimeupdate", value)
	return s
}

/*
Ontoggle -
*/
func (s *StyleTagHtml) Ontoggle(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontoggle", value)
	return s
}

/*
Onvolumechange -
*/
func (s *StyleTagHtml) Onvolumechange(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onvolumechange", value)
	return s
}

/*
Onwaiting -
*/
func (s *StyleTagHtml) Onwaiting(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onwaiting", value)
	return s
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (s *StyleTagHtml) Onafterprint(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onafterprint", value)
	return s
}

/*
Onbeforeprint -
*/
func (s *StyleTagHtml) Onbeforeprint(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeprint", value)
	return s
}

/*
Onbeforeunload -
*/
func (s *StyleTagHtml) Onbeforeunload(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeunload", value)
	return s
}

/*
Onerror -
*/
func (s *StyleTagHtml) Onerror(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onerror", value)
	return s
}

/*
Onhashchange -
*/
func (s *StyleTagHtml) Onhashchange(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onhashchange", value)
	return s
}

/*
Onload -
*/
func (s *StyleTagHtml) Onload(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onload", value)
	return s
}

/*
Onmessage -
*/
func (s *StyleTagHtml) Onmessage(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmessage", value)
	return s
}

/*
Onoffline -
*/
func (s *StyleTagHtml) Onoffline(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onoffline", value)
	return s
}

/*
Ononline -
*/
func (s *StyleTagHtml) Ononline(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ononline", value)
	return s
}

/*
Onpagehide -
*/
func (s *StyleTagHtml) Onpagehide(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpagehide", value)
	return s
}

/*
Onpageshow -
*/
func (s *StyleTagHtml) Onpageshow(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpageshow", value)
	return s
}

/*
Onpopstate -
*/
func (s *StyleTagHtml) Onpopstate(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpopstate", value)
	return s
}

/*
Onresize -
*/
func (s *StyleTagHtml) Onresize(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onresize", value)
	return s
}

/*
Onstorage -
*/
func (s *StyleTagHtml) Onstorage(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstorage", value)
	return s
}

/*
Onunload -
*/
func (s *StyleTagHtml) Onunload(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onunload", value)
	return s
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (s *StyleTagHtml) AccessKey(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("accessKey", value)
	return s
}

/*
Aria -
*/
func (s *StyleTagHtml) Aria(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria", value)
	return s
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (s *StyleTagHtml) Autocapitalize(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("autocapitalize", value)
	return s
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (s *StyleTagHtml) Autofocus(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("autofocus", value)
	return s
}

/*
Class -
*/
func (s *StyleTagHtml) Class(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("class", value)
	return s
}

/*
Contenteditable -
*/
func (s *StyleTagHtml) Contenteditable(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("contenteditable", value)
	return s
}

/*
Data -
*/
func (s *StyleTagHtml) Data(name, value string) *StyleTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	s.registerAttribute(dataName, value)
	return s
}

/*
Dir -
*/
func (s *StyleTagHtml) Dir(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("dir", value)
	return s
}

/*
Draggable -
*/
func (s *StyleTagHtml) Draggable(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("draggable", value)
	return s
}

/*
EnterKeyHint -
*/
func (s *StyleTagHtml) EnterKeyHint(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("enterKeyHint", value)
	return s
}

/*
ExportParts -
*/
func (s *StyleTagHtml) ExportParts(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("exportParts", value)
	return s
}

/*
Hidden -
*/
func (s *StyleTagHtml) Hidden(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("hidden", value)
	return s
}

/*
Id -
*/
func (s *StyleTagHtml) Id(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("id", value)
	return s
}

/*
Inert -
*/
func (s *StyleTagHtml) Inert(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inert", value)
	return s
}

/*
InputMode -
*/
func (s *StyleTagHtml) InputMode(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inputMode", value)
	return s
}

/*
Is -
*/
func (s *StyleTagHtml) Is(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("is", value)
	return s
}

/*
ItemId -
*/
func (s *StyleTagHtml) ItemId(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemId", value)
	return s
}

/*
ItemProp -
*/
func (s *StyleTagHtml) ItemProp(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemProp", value)
	return s
}

/*
ItemRef -
*/
func (s *StyleTagHtml) ItemRef(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemRef", value)
	return s
}

/*
ItemScope -
*/
func (s *StyleTagHtml) ItemScope(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemScope", value)
	return s
}

/*
ItemType -
*/
func (s *StyleTagHtml) ItemType(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemType", value)
	return s
}

/*
Lang -
*/
func (s *StyleTagHtml) Lang(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("lang", value)
	return s
}

/*
Nonce -
*/
func (s *StyleTagHtml) Nonce(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("nonce", value)
	return s
}

/*
Part -
*/
func (s *StyleTagHtml) Part(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("part", value)
	return s
}

/*
Popover -
*/
func (s *StyleTagHtml) Popover() *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("popover", "")
	return s
}

/*
Role -
*/
func (s *StyleTagHtml) Role(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("role", value)
	return s
}

/*
Slot -
*/
func (s *StyleTagHtml) Slot(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("slot", value)
	return s
}

/*
Spellcheck -
*/
func (s *StyleTagHtml) Spellcheck(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("spellcheck", value)
	return s
}

/*
Style -
*/
func (s *StyleTagHtml) Style(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("style", value)
	return s
}

/*
Tabindex -
*/
func (s *StyleTagHtml) Tabindex(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("tabindex", value)
	return s
}

/*
Title -
*/
func (s *StyleTagHtml) Title(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("title", value)
	return s
}

/*
Translate -
*/
func (s *StyleTagHtml) Translate(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("translate", value)
	return s
}

/*
VirtualKeyBoardPolicy -
*/
func (s *StyleTagHtml) VirtualKeyBoardPolicy(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("virtualKeyBoardPolicy", value)
	return s
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (s *StyleTagHtml) AriaAtomic(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-atomic", value)
	return s
}

/*
AriaBusy -
*/
func (s *StyleTagHtml) AriaBusy(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-busy", value)
	return s
}

/*
AriaControls -
*/
func (s *StyleTagHtml) AriaControls(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-controls", value)
	return s
}

/*
AriaCurrent -
*/
func (s *StyleTagHtml) AriaCurrent(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-current", value)
	return s
}

/*
AriaDescribedby -
*/
func (s *StyleTagHtml) AriaDescribedby(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-describedby", value)
	return s
}

/*
AriaDescription -
*/
func (s *StyleTagHtml) AriaDescription(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-description", value)
	return s
}

/*
AriaDetails -
*/
func (s *StyleTagHtml) AriaDetails(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-details", value)
	return s
}

/*
AriaDisabled -
*/
func (s *StyleTagHtml) AriaDisabled(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-disabled", value)
	return s
}

/*
AriaDropeffect -
*/
func (s *StyleTagHtml) AriaDropeffect(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-dropeffect", value)
	return s
}

/*
AriaErrormessage -
*/
func (s *StyleTagHtml) AriaErrormessage(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-errormessage", value)
	return s
}

/*
AriaFlowto -
*/
func (s *StyleTagHtml) AriaFlowto(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-flowto", value)
	return s
}

/*
AriaGrabbed -
*/
func (s *StyleTagHtml) AriaGrabbed(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-grabbed", value)
	return s
}

/*
AriaHaspopup -
*/
func (s *StyleTagHtml) AriaHaspopup(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-haspopup", value)
	return s
}

/*
AriaHidden -
*/
func (s *StyleTagHtml) AriaHidden(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-hidden", value)
	return s
}

/*
AriaInvalid -
*/
func (s *StyleTagHtml) AriaInvalid(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-invalid", value)
	return s
}

/*
AriaKeyshortcuts -
*/
func (s *StyleTagHtml) AriaKeyshortcuts(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-keyshortcuts", value)
	return s
}

/*
AriaLabel -
*/
func (s *StyleTagHtml) AriaLabel(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-label", value)
	return s
}

/*
AriaLabelledby -
*/
func (s *StyleTagHtml) AriaLabelledby(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-labelledby", value)
	return s
}

/*
AriaLive -
*/
func (s *StyleTagHtml) AriaLive(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-live", value)
	return s
}

/*
AriaOwns -
*/
func (s *StyleTagHtml) AriaOwns(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-owns", value)
	return s
}

/*
AriaRelevant -
*/
func (s *StyleTagHtml) AriaRelevant(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-relevant", value)
	return s
}

/*
AriaRoledescription -
*/
func (s *StyleTagHtml) AriaRoledescription(value string) *StyleTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-roledescription", value)
	return s
}
