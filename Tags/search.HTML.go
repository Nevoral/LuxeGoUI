// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func SearchHtml(tags []any) *SearchTagHtml {
	node := &SearchTagHtml{
		tag: &tag{
			name:                "search",
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

type SearchTagHtml struct {
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
func (s *SearchTagHtml) CustomAttribute(attributes ...*Attribute) *SearchTagHtml {
	s.registerAttributes(attributes...)
	return s
}

/*
Children - Method for nesting tags into parent tag
*/
func (s *SearchTagHtml) Children(tags ...any) *SearchTagHtml {
	return s.supportedChildrenCheck(tags)
}

func (s *SearchTagHtml) supportedChildrenCheck(tags []any) *SearchTagHtml {
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
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (s *SearchTagHtml) AriaAtomic(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-atomic", value)
	return s
}

/*
AriaBusy -
*/
func (s *SearchTagHtml) AriaBusy(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-busy", value)
	return s
}

/*
AriaControls -
*/
func (s *SearchTagHtml) AriaControls(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-controls", value)
	return s
}

/*
AriaCurrent -
*/
func (s *SearchTagHtml) AriaCurrent(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-current", value)
	return s
}

/*
AriaDescribedby -
*/
func (s *SearchTagHtml) AriaDescribedby(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-describedby", value)
	return s
}

/*
AriaDescription -
*/
func (s *SearchTagHtml) AriaDescription(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-description", value)
	return s
}

/*
AriaDetails -
*/
func (s *SearchTagHtml) AriaDetails(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-details", value)
	return s
}

/*
AriaDisabled -
*/
func (s *SearchTagHtml) AriaDisabled(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-disabled", value)
	return s
}

/*
AriaDropeffect -
*/
func (s *SearchTagHtml) AriaDropeffect(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-dropeffect", value)
	return s
}

/*
AriaErrormessage -
*/
func (s *SearchTagHtml) AriaErrormessage(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-errormessage", value)
	return s
}

/*
AriaFlowto -
*/
func (s *SearchTagHtml) AriaFlowto(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-flowto", value)
	return s
}

/*
AriaGrabbed -
*/
func (s *SearchTagHtml) AriaGrabbed(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-grabbed", value)
	return s
}

/*
AriaHaspopup -
*/
func (s *SearchTagHtml) AriaHaspopup(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-haspopup", value)
	return s
}

/*
AriaHidden -
*/
func (s *SearchTagHtml) AriaHidden(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-hidden", value)
	return s
}

/*
AriaInvalid -
*/
func (s *SearchTagHtml) AriaInvalid(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-invalid", value)
	return s
}

/*
AriaKeyshortcuts -
*/
func (s *SearchTagHtml) AriaKeyshortcuts(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-keyshortcuts", value)
	return s
}

/*
AriaLabel -
*/
func (s *SearchTagHtml) AriaLabel(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-label", value)
	return s
}

/*
AriaLabelledby -
*/
func (s *SearchTagHtml) AriaLabelledby(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-labelledby", value)
	return s
}

/*
AriaLive -
*/
func (s *SearchTagHtml) AriaLive(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-live", value)
	return s
}

/*
AriaOwns -
*/
func (s *SearchTagHtml) AriaOwns(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-owns", value)
	return s
}

/*
AriaRelevant -
*/
func (s *SearchTagHtml) AriaRelevant(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-relevant", value)
	return s
}

/*
AriaRoledescription -
*/
func (s *SearchTagHtml) AriaRoledescription(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("aria-roledescription", value)
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
func (s *SearchTagHtml) Onabort(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onabort", value)
	return s
}

/*
Onautocomplete -
*/
func (s *SearchTagHtml) Onautocomplete(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocomplete", value)
	return s
}

/*
Onautocompleteerror -
*/
func (s *SearchTagHtml) Onautocompleteerror(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onautocompleteerror", value)
	return s
}

/*
Onblur -
*/
func (s *SearchTagHtml) Onblur(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onblur", value)
	return s
}

/*
Oncancel -
*/
func (s *SearchTagHtml) Oncancel(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncancel", value)
	return s
}

/*
Oncanplay -
*/
func (s *SearchTagHtml) Oncanplay(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplay", value)
	return s
}

/*
Oncanplaythrough -
*/
func (s *SearchTagHtml) Oncanplaythrough(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncanplaythrough", value)
	return s
}

/*
Onchange -
*/
func (s *SearchTagHtml) Onchange(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onchange", value)
	return s
}

/*
Onclick -
*/
func (s *SearchTagHtml) Onclick(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclick", value)
	return s
}

/*
Onclose -
*/
func (s *SearchTagHtml) Onclose(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onclose", value)
	return s
}

/*
Oncontextmenu -
*/
func (s *SearchTagHtml) Oncontextmenu(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncontextmenu", value)
	return s
}

/*
Oncuechange -
*/
func (s *SearchTagHtml) Oncuechange(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oncuechange", value)
	return s
}

/*
Ondblclick -
*/
func (s *SearchTagHtml) Ondblclick(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondblclick", value)
	return s
}

/*
Ondrag -
*/
func (s *SearchTagHtml) Ondrag(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrag", value)
	return s
}

/*
Ondragend -
*/
func (s *SearchTagHtml) Ondragend(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragend", value)
	return s
}

/*
Ondragenter -
*/
func (s *SearchTagHtml) Ondragenter(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragenter", value)
	return s
}

/*
Ondragleave -
*/
func (s *SearchTagHtml) Ondragleave(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragleave", value)
	return s
}

/*
Ondragover -
*/
func (s *SearchTagHtml) Ondragover(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragover", value)
	return s
}

/*
Ondragstart -
*/
func (s *SearchTagHtml) Ondragstart(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondragstart", value)
	return s
}

/*
Ondrop -
*/
func (s *SearchTagHtml) Ondrop(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondrop", value)
	return s
}

/*
Ondurationchange -
*/
func (s *SearchTagHtml) Ondurationchange(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ondurationchange", value)
	return s
}

/*
Onemptied -
*/
func (s *SearchTagHtml) Onemptied(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onemptied", value)
	return s
}

/*
Onended -
*/
func (s *SearchTagHtml) Onended(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onended", value)
	return s
}

/*
Onfocus -
*/
func (s *SearchTagHtml) Onfocus(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onfocus", value)
	return s
}

/*
Oninput -
*/
func (s *SearchTagHtml) Oninput(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninput", value)
	return s
}

/*
Oninvalid -
*/
func (s *SearchTagHtml) Oninvalid(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("oninvalid", value)
	return s
}

/*
Onkeydown -
*/
func (s *SearchTagHtml) Onkeydown(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeydown", value)
	return s
}

/*
Onkeypress -
*/
func (s *SearchTagHtml) Onkeypress(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeypress", value)
	return s
}

/*
Onkeyup -
*/
func (s *SearchTagHtml) Onkeyup(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onkeyup", value)
	return s
}

/*
Onloadeddata -
*/
func (s *SearchTagHtml) Onloadeddata(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadeddata", value)
	return s
}

/*
Onloadedmetadata -
*/
func (s *SearchTagHtml) Onloadedmetadata(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadedmetadata", value)
	return s
}

/*
Onloadstart -
*/
func (s *SearchTagHtml) Onloadstart(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onloadstart", value)
	return s
}

/*
Onmousedown -
*/
func (s *SearchTagHtml) Onmousedown(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousedown", value)
	return s
}

/*
Onmouseenter -
*/
func (s *SearchTagHtml) Onmouseenter(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseenter", value)
	return s
}

/*
Onmouseleave -
*/
func (s *SearchTagHtml) Onmouseleave(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseleave", value)
	return s
}

/*
Onmousemove -
*/
func (s *SearchTagHtml) Onmousemove(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousemove", value)
	return s
}

/*
Onmouseout -
*/
func (s *SearchTagHtml) Onmouseout(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseout", value)
	return s
}

/*
Onmouseover -
*/
func (s *SearchTagHtml) Onmouseover(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseover", value)
	return s
}

/*
Onmouseup -
*/
func (s *SearchTagHtml) Onmouseup(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmouseup", value)
	return s
}

/*
Onmousewheel -
*/
func (s *SearchTagHtml) Onmousewheel(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmousewheel", value)
	return s
}

/*
Onpause -
*/
func (s *SearchTagHtml) Onpause(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpause", value)
	return s
}

/*
Onplay -
*/
func (s *SearchTagHtml) Onplay(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplay", value)
	return s
}

/*
Onplaying -
*/
func (s *SearchTagHtml) Onplaying(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onplaying", value)
	return s
}

/*
Onprogress -
*/
func (s *SearchTagHtml) Onprogress(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onprogress", value)
	return s
}

/*
Onratechange -
*/
func (s *SearchTagHtml) Onratechange(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onratechange", value)
	return s
}

/*
Onreset -
*/
func (s *SearchTagHtml) Onreset(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onreset", value)
	return s
}

/*
Onscroll -
*/
func (s *SearchTagHtml) Onscroll(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onscroll", value)
	return s
}

/*
Onseeked -
*/
func (s *SearchTagHtml) Onseeked(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeked", value)
	return s
}

/*
Onseeking -
*/
func (s *SearchTagHtml) Onseeking(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onseeking", value)
	return s
}

/*
Onselect -
*/
func (s *SearchTagHtml) Onselect(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onselect", value)
	return s
}

/*
Onshow -
*/
func (s *SearchTagHtml) Onshow(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onshow", value)
	return s
}

/*
Onsort -
*/
func (s *SearchTagHtml) Onsort(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsort", value)
	return s
}

/*
Onstalled -
*/
func (s *SearchTagHtml) Onstalled(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstalled", value)
	return s
}

/*
Onsubmit -
*/
func (s *SearchTagHtml) Onsubmit(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsubmit", value)
	return s
}

/*
Onsuspend -
*/
func (s *SearchTagHtml) Onsuspend(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onsuspend", value)
	return s
}

/*
Ontimeupdate -
*/
func (s *SearchTagHtml) Ontimeupdate(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontimeupdate", value)
	return s
}

/*
Ontoggle -
*/
func (s *SearchTagHtml) Ontoggle(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ontoggle", value)
	return s
}

/*
Onvolumechange -
*/
func (s *SearchTagHtml) Onvolumechange(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onvolumechange", value)
	return s
}

/*
Onwaiting -
*/
func (s *SearchTagHtml) Onwaiting(value string) *SearchTagHtml {
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
func (s *SearchTagHtml) Onafterprint(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onafterprint", value)
	return s
}

/*
Onbeforeprint -
*/
func (s *SearchTagHtml) Onbeforeprint(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeprint", value)
	return s
}

/*
Onbeforeunload -
*/
func (s *SearchTagHtml) Onbeforeunload(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onbeforeunload", value)
	return s
}

/*
Onerror -
*/
func (s *SearchTagHtml) Onerror(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onerror", value)
	return s
}

/*
Onhashchange -
*/
func (s *SearchTagHtml) Onhashchange(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onhashchange", value)
	return s
}

/*
Onload -
*/
func (s *SearchTagHtml) Onload(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onload", value)
	return s
}

/*
Onmessage -
*/
func (s *SearchTagHtml) Onmessage(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onmessage", value)
	return s
}

/*
Onoffline -
*/
func (s *SearchTagHtml) Onoffline(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onoffline", value)
	return s
}

/*
Ononline -
*/
func (s *SearchTagHtml) Ononline(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("ononline", value)
	return s
}

/*
Onpagehide -
*/
func (s *SearchTagHtml) Onpagehide(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpagehide", value)
	return s
}

/*
Onpageshow -
*/
func (s *SearchTagHtml) Onpageshow(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpageshow", value)
	return s
}

/*
Onpopstate -
*/
func (s *SearchTagHtml) Onpopstate(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onpopstate", value)
	return s
}

/*
Onresize -
*/
func (s *SearchTagHtml) Onresize(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onresize", value)
	return s
}

/*
Onstorage -
*/
func (s *SearchTagHtml) Onstorage(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("onstorage", value)
	return s
}

/*
Onunload -
*/
func (s *SearchTagHtml) Onunload(value string) *SearchTagHtml {
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
func (s *SearchTagHtml) AccessKey(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("accessKey", value)
	return s
}

/*
Aria -
*/
func (s *SearchTagHtml) Aria(value string) *SearchTagHtml {
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
func (s *SearchTagHtml) Autocapitalize(value string) *SearchTagHtml {
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
func (s *SearchTagHtml) Autofocus(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("autofocus", value)
	return s
}

/*
Class -
*/
func (s *SearchTagHtml) Class(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("class", value)
	return s
}

/*
Contenteditable -
*/
func (s *SearchTagHtml) Contenteditable(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("contenteditable", value)
	return s
}

/*
Data -
*/
func (s *SearchTagHtml) Data(name, value string) *SearchTagHtml {
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
func (s *SearchTagHtml) Dir(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("dir", value)
	return s
}

/*
Draggable -
*/
func (s *SearchTagHtml) Draggable(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("draggable", value)
	return s
}

/*
EnterKeyHint -
*/
func (s *SearchTagHtml) EnterKeyHint(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("enterKeyHint", value)
	return s
}

/*
ExportParts -
*/
func (s *SearchTagHtml) ExportParts(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("exportParts", value)
	return s
}

/*
Hidden -
*/
func (s *SearchTagHtml) Hidden(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("hidden", value)
	return s
}

/*
Id -
*/
func (s *SearchTagHtml) Id(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("id", value)
	return s
}

/*
Inert -
*/
func (s *SearchTagHtml) Inert(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inert", value)
	return s
}

/*
InputMode -
*/
func (s *SearchTagHtml) InputMode(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("inputMode", value)
	return s
}

/*
Is -
*/
func (s *SearchTagHtml) Is(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("is", value)
	return s
}

/*
ItemId -
*/
func (s *SearchTagHtml) ItemId(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemId", value)
	return s
}

/*
ItemProp -
*/
func (s *SearchTagHtml) ItemProp(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemProp", value)
	return s
}

/*
ItemRef -
*/
func (s *SearchTagHtml) ItemRef(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemRef", value)
	return s
}

/*
ItemScope -
*/
func (s *SearchTagHtml) ItemScope(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemScope", value)
	return s
}

/*
ItemType -
*/
func (s *SearchTagHtml) ItemType(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("itemType", value)
	return s
}

/*
Lang -
*/
func (s *SearchTagHtml) Lang(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("lang", value)
	return s
}

/*
Nonce -
*/
func (s *SearchTagHtml) Nonce(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("nonce", value)
	return s
}

/*
Part -
*/
func (s *SearchTagHtml) Part(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("part", value)
	return s
}

/*
Popover -
*/
func (s *SearchTagHtml) Popover() *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("popover", "")
	return s
}

/*
Role -
*/
func (s *SearchTagHtml) Role(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("role", value)
	return s
}

/*
Slot -
*/
func (s *SearchTagHtml) Slot(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("slot", value)
	return s
}

/*
Spellcheck -
*/
func (s *SearchTagHtml) Spellcheck(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("spellcheck", value)
	return s
}

/*
Style -
*/
func (s *SearchTagHtml) Style(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("style", value)
	return s
}

/*
Tabindex -
*/
func (s *SearchTagHtml) Tabindex(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("tabindex", value)
	return s
}

/*
Title -
*/
func (s *SearchTagHtml) Title(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("title", value)
	return s
}

/*
Translate -
*/
func (s *SearchTagHtml) Translate(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("translate", value)
	return s
}

/*
VirtualKeyBoardPolicy -
*/
func (s *SearchTagHtml) VirtualKeyBoardPolicy(value string) *SearchTagHtml {
	if s.attributes == nil {
		s.attributes = []*Attribute{}
	}
	s.registerAttribute("virtualKeyBoardPolicy", value)
	return s
}
