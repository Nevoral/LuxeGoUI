// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func BrHtml() *BrTagHtml {
	return &BrTagHtml{
		tag: &tag{
			name:                "br",
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

type BrTagHtml struct {
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
func (b *BrTagHtml) CustomAttribute(attributes ...*Attribute) *BrTagHtml {
	b.registerAttributes(attributes...)
	return b
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
func (b *BrTagHtml) AriaAtomic(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-atomic", value)
	return b
}

/*
AriaBusy -
*/
func (b *BrTagHtml) AriaBusy(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-busy", value)
	return b
}

/*
AriaControls -
*/
func (b *BrTagHtml) AriaControls(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-controls", value)
	return b
}

/*
AriaCurrent -
*/
func (b *BrTagHtml) AriaCurrent(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-current", value)
	return b
}

/*
AriaDescribedby -
*/
func (b *BrTagHtml) AriaDescribedby(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-describedby", value)
	return b
}

/*
AriaDescription -
*/
func (b *BrTagHtml) AriaDescription(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-description", value)
	return b
}

/*
AriaDetails -
*/
func (b *BrTagHtml) AriaDetails(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-details", value)
	return b
}

/*
AriaDisabled -
*/
func (b *BrTagHtml) AriaDisabled(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-disabled", value)
	return b
}

/*
AriaDropeffect -
*/
func (b *BrTagHtml) AriaDropeffect(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-dropeffect", value)
	return b
}

/*
AriaErrormessage -
*/
func (b *BrTagHtml) AriaErrormessage(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-errormessage", value)
	return b
}

/*
AriaFlowto -
*/
func (b *BrTagHtml) AriaFlowto(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-flowto", value)
	return b
}

/*
AriaGrabbed -
*/
func (b *BrTagHtml) AriaGrabbed(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-grabbed", value)
	return b
}

/*
AriaHaspopup -
*/
func (b *BrTagHtml) AriaHaspopup(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-haspopup", value)
	return b
}

/*
AriaHidden -
*/
func (b *BrTagHtml) AriaHidden(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-hidden", value)
	return b
}

/*
AriaInvalid -
*/
func (b *BrTagHtml) AriaInvalid(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-invalid", value)
	return b
}

/*
AriaKeyshortcuts -
*/
func (b *BrTagHtml) AriaKeyshortcuts(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-keyshortcuts", value)
	return b
}

/*
AriaLabel -
*/
func (b *BrTagHtml) AriaLabel(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-label", value)
	return b
}

/*
AriaLabelledby -
*/
func (b *BrTagHtml) AriaLabelledby(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-labelledby", value)
	return b
}

/*
AriaLive -
*/
func (b *BrTagHtml) AriaLive(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-live", value)
	return b
}

/*
AriaOwns -
*/
func (b *BrTagHtml) AriaOwns(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-owns", value)
	return b
}

/*
AriaRelevant -
*/
func (b *BrTagHtml) AriaRelevant(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-relevant", value)
	return b
}

/*
AriaRoledescription -
*/
func (b *BrTagHtml) AriaRoledescription(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria-roledescription", value)
	return b
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (b *BrTagHtml) Onabort(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onabort", value)
	return b
}

/*
Onautocomplete -
*/
func (b *BrTagHtml) Onautocomplete(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onautocomplete", value)
	return b
}

/*
Onautocompleteerror -
*/
func (b *BrTagHtml) Onautocompleteerror(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onautocompleteerror", value)
	return b
}

/*
Onblur -
*/
func (b *BrTagHtml) Onblur(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onblur", value)
	return b
}

/*
Oncancel -
*/
func (b *BrTagHtml) Oncancel(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncancel", value)
	return b
}

/*
Oncanplay -
*/
func (b *BrTagHtml) Oncanplay(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncanplay", value)
	return b
}

/*
Oncanplaythrough -
*/
func (b *BrTagHtml) Oncanplaythrough(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncanplaythrough", value)
	return b
}

/*
Onchange -
*/
func (b *BrTagHtml) Onchange(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onchange", value)
	return b
}

/*
Onclick -
*/
func (b *BrTagHtml) Onclick(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onclick", value)
	return b
}

/*
Onclose -
*/
func (b *BrTagHtml) Onclose(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onclose", value)
	return b
}

/*
Oncontextmenu -
*/
func (b *BrTagHtml) Oncontextmenu(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncontextmenu", value)
	return b
}

/*
Oncuechange -
*/
func (b *BrTagHtml) Oncuechange(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oncuechange", value)
	return b
}

/*
Ondblclick -
*/
func (b *BrTagHtml) Ondblclick(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondblclick", value)
	return b
}

/*
Ondrag -
*/
func (b *BrTagHtml) Ondrag(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondrag", value)
	return b
}

/*
Ondragend -
*/
func (b *BrTagHtml) Ondragend(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragend", value)
	return b
}

/*
Ondragenter -
*/
func (b *BrTagHtml) Ondragenter(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragenter", value)
	return b
}

/*
Ondragleave -
*/
func (b *BrTagHtml) Ondragleave(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragleave", value)
	return b
}

/*
Ondragover -
*/
func (b *BrTagHtml) Ondragover(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragover", value)
	return b
}

/*
Ondragstart -
*/
func (b *BrTagHtml) Ondragstart(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondragstart", value)
	return b
}

/*
Ondrop -
*/
func (b *BrTagHtml) Ondrop(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondrop", value)
	return b
}

/*
Ondurationchange -
*/
func (b *BrTagHtml) Ondurationchange(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ondurationchange", value)
	return b
}

/*
Onemptied -
*/
func (b *BrTagHtml) Onemptied(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onemptied", value)
	return b
}

/*
Onended -
*/
func (b *BrTagHtml) Onended(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onended", value)
	return b
}

/*
Onfocus -
*/
func (b *BrTagHtml) Onfocus(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onfocus", value)
	return b
}

/*
Oninput -
*/
func (b *BrTagHtml) Oninput(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oninput", value)
	return b
}

/*
Oninvalid -
*/
func (b *BrTagHtml) Oninvalid(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("oninvalid", value)
	return b
}

/*
Onkeydown -
*/
func (b *BrTagHtml) Onkeydown(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeydown", value)
	return b
}

/*
Onkeypress -
*/
func (b *BrTagHtml) Onkeypress(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeypress", value)
	return b
}

/*
Onkeyup -
*/
func (b *BrTagHtml) Onkeyup(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onkeyup", value)
	return b
}

/*
Onloadeddata -
*/
func (b *BrTagHtml) Onloadeddata(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadeddata", value)
	return b
}

/*
Onloadedmetadata -
*/
func (b *BrTagHtml) Onloadedmetadata(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadedmetadata", value)
	return b
}

/*
Onloadstart -
*/
func (b *BrTagHtml) Onloadstart(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onloadstart", value)
	return b
}

/*
Onmousedown -
*/
func (b *BrTagHtml) Onmousedown(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousedown", value)
	return b
}

/*
Onmouseenter -
*/
func (b *BrTagHtml) Onmouseenter(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseenter", value)
	return b
}

/*
Onmouseleave -
*/
func (b *BrTagHtml) Onmouseleave(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseleave", value)
	return b
}

/*
Onmousemove -
*/
func (b *BrTagHtml) Onmousemove(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousemove", value)
	return b
}

/*
Onmouseout -
*/
func (b *BrTagHtml) Onmouseout(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseout", value)
	return b
}

/*
Onmouseover -
*/
func (b *BrTagHtml) Onmouseover(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseover", value)
	return b
}

/*
Onmouseup -
*/
func (b *BrTagHtml) Onmouseup(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmouseup", value)
	return b
}

/*
Onmousewheel -
*/
func (b *BrTagHtml) Onmousewheel(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmousewheel", value)
	return b
}

/*
Onpause -
*/
func (b *BrTagHtml) Onpause(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpause", value)
	return b
}

/*
Onplay -
*/
func (b *BrTagHtml) Onplay(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onplay", value)
	return b
}

/*
Onplaying -
*/
func (b *BrTagHtml) Onplaying(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onplaying", value)
	return b
}

/*
Onprogress -
*/
func (b *BrTagHtml) Onprogress(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onprogress", value)
	return b
}

/*
Onratechange -
*/
func (b *BrTagHtml) Onratechange(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onratechange", value)
	return b
}

/*
Onreset -
*/
func (b *BrTagHtml) Onreset(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onreset", value)
	return b
}

/*
Onscroll -
*/
func (b *BrTagHtml) Onscroll(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onscroll", value)
	return b
}

/*
Onseeked -
*/
func (b *BrTagHtml) Onseeked(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onseeked", value)
	return b
}

/*
Onseeking -
*/
func (b *BrTagHtml) Onseeking(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onseeking", value)
	return b
}

/*
Onselect -
*/
func (b *BrTagHtml) Onselect(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onselect", value)
	return b
}

/*
Onshow -
*/
func (b *BrTagHtml) Onshow(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onshow", value)
	return b
}

/*
Onsort -
*/
func (b *BrTagHtml) Onsort(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsort", value)
	return b
}

/*
Onstalled -
*/
func (b *BrTagHtml) Onstalled(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onstalled", value)
	return b
}

/*
Onsubmit -
*/
func (b *BrTagHtml) Onsubmit(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsubmit", value)
	return b
}

/*
Onsuspend -
*/
func (b *BrTagHtml) Onsuspend(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onsuspend", value)
	return b
}

/*
Ontimeupdate -
*/
func (b *BrTagHtml) Ontimeupdate(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ontimeupdate", value)
	return b
}

/*
Ontoggle -
*/
func (b *BrTagHtml) Ontoggle(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ontoggle", value)
	return b
}

/*
Onvolumechange -
*/
func (b *BrTagHtml) Onvolumechange(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onvolumechange", value)
	return b
}

/*
Onwaiting -
*/
func (b *BrTagHtml) Onwaiting(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onwaiting", value)
	return b
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (b *BrTagHtml) Onafterprint(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onafterprint", value)
	return b
}

/*
Onbeforeprint -
*/
func (b *BrTagHtml) Onbeforeprint(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onbeforeprint", value)
	return b
}

/*
Onbeforeunload -
*/
func (b *BrTagHtml) Onbeforeunload(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onbeforeunload", value)
	return b
}

/*
Onerror -
*/
func (b *BrTagHtml) Onerror(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onerror", value)
	return b
}

/*
Onhashchange -
*/
func (b *BrTagHtml) Onhashchange(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onhashchange", value)
	return b
}

/*
Onload -
*/
func (b *BrTagHtml) Onload(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onload", value)
	return b
}

/*
Onmessage -
*/
func (b *BrTagHtml) Onmessage(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onmessage", value)
	return b
}

/*
Onoffline -
*/
func (b *BrTagHtml) Onoffline(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onoffline", value)
	return b
}

/*
Ononline -
*/
func (b *BrTagHtml) Ononline(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("ononline", value)
	return b
}

/*
Onpagehide -
*/
func (b *BrTagHtml) Onpagehide(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpagehide", value)
	return b
}

/*
Onpageshow -
*/
func (b *BrTagHtml) Onpageshow(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpageshow", value)
	return b
}

/*
Onpopstate -
*/
func (b *BrTagHtml) Onpopstate(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onpopstate", value)
	return b
}

/*
Onresize -
*/
func (b *BrTagHtml) Onresize(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onresize", value)
	return b
}

/*
Onstorage -
*/
func (b *BrTagHtml) Onstorage(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onstorage", value)
	return b
}

/*
Onunload -
*/
func (b *BrTagHtml) Onunload(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("onunload", value)
	return b
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (b *BrTagHtml) AccessKey(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("accessKey", value)
	return b
}

/*
Aria -
*/
func (b *BrTagHtml) Aria(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("aria", value)
	return b
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (b *BrTagHtml) Autocapitalize(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("autocapitalize", value)
	return b
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (b *BrTagHtml) Autofocus(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("autofocus", value)
	return b
}

/*
Class -
*/
func (b *BrTagHtml) Class(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("class", value)
	return b
}

/*
Contenteditable -
*/
func (b *BrTagHtml) Contenteditable(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("contenteditable", value)
	return b
}

/*
Data -
*/
func (b *BrTagHtml) Data(name, value string) *BrTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	b.registerAttribute(dataName, value)
	return b
}

/*
Dir -
*/
func (b *BrTagHtml) Dir(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("dir", value)
	return b
}

/*
Draggable -
*/
func (b *BrTagHtml) Draggable(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("draggable", value)
	return b
}

/*
EnterKeyHint -
*/
func (b *BrTagHtml) EnterKeyHint(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("enterKeyHint", value)
	return b
}

/*
ExportParts -
*/
func (b *BrTagHtml) ExportParts(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("exportParts", value)
	return b
}

/*
Hidden -
*/
func (b *BrTagHtml) Hidden(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("hidden", value)
	return b
}

/*
Id -
*/
func (b *BrTagHtml) Id(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("id", value)
	return b
}

/*
Inert -
*/
func (b *BrTagHtml) Inert(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("inert", value)
	return b
}

/*
InputMode -
*/
func (b *BrTagHtml) InputMode(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("inputMode", value)
	return b
}

/*
Is -
*/
func (b *BrTagHtml) Is(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("is", value)
	return b
}

/*
ItemId -
*/
func (b *BrTagHtml) ItemId(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemId", value)
	return b
}

/*
ItemProp -
*/
func (b *BrTagHtml) ItemProp(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemProp", value)
	return b
}

/*
ItemRef -
*/
func (b *BrTagHtml) ItemRef(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemRef", value)
	return b
}

/*
ItemScope -
*/
func (b *BrTagHtml) ItemScope(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemScope", value)
	return b
}

/*
ItemType -
*/
func (b *BrTagHtml) ItemType(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("itemType", value)
	return b
}

/*
Lang -
*/
func (b *BrTagHtml) Lang(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("lang", value)
	return b
}

/*
Nonce -
*/
func (b *BrTagHtml) Nonce(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("nonce", value)
	return b
}

/*
Part -
*/
func (b *BrTagHtml) Part(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("part", value)
	return b
}

/*
Popover -
*/
func (b *BrTagHtml) Popover() *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("popover", "")
	return b
}

/*
Role -
*/
func (b *BrTagHtml) Role(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("role", value)
	return b
}

/*
Slot -
*/
func (b *BrTagHtml) Slot(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("slot", value)
	return b
}

/*
Spellcheck -
*/
func (b *BrTagHtml) Spellcheck(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("spellcheck", value)
	return b
}

/*
Style -
*/
func (b *BrTagHtml) Style(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("style", value)
	return b
}

/*
Tabindex -
*/
func (b *BrTagHtml) Tabindex(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("tabindex", value)
	return b
}

/*
Title -
*/
func (b *BrTagHtml) Title(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("title", value)
	return b
}

/*
Translate -
*/
func (b *BrTagHtml) Translate(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("translate", value)
	return b
}

/*
VirtualKeyBoardPolicy -
*/
func (b *BrTagHtml) VirtualKeyBoardPolicy(value string) *BrTagHtml {
	if b.attributes == nil {
		b.attributes = []*Attribute{}
	}
	b.registerAttribute("virtualKeyBoardPolicy", value)
	return b
}
