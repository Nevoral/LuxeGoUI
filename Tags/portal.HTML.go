// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func PortalHtml(tags []any) *PortalTagHtml {
	node := &PortalTagHtml{
		tag: &tag{
			name:                "portal",
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

type PortalTagHtml struct {
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
func (p *PortalTagHtml) CustomAttribute(attributes ...*Attribute) *PortalTagHtml {
	p.registerAttributes(attributes...)
	return p
}

/*
Children - Method for nesting tags into parent tag
*/
func (p *PortalTagHtml) Children(tags ...any) *PortalTagHtml {
	return p.supportedChildrenCheck(tags)
}

func (p *PortalTagHtml) supportedChildrenCheck(tags []any) *PortalTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			p.registerChildren(TextContentSvg(val).getTag())
		case content:
			p.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				p.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return p
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Referrerpolicy -
*/
func (p *PortalTagHtml) Referrerpolicy(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("referrerpolicy", value)
	return p
}

/*
Src - Specifies the URL of an image.
Specifies the URL of an image.
*/
func (p *PortalTagHtml) Src(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("src", value)
	return p
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (p *PortalTagHtml) AriaAtomic(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-atomic", value)
	return p
}

/*
AriaBusy -
*/
func (p *PortalTagHtml) AriaBusy(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-busy", value)
	return p
}

/*
AriaControls -
*/
func (p *PortalTagHtml) AriaControls(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-controls", value)
	return p
}

/*
AriaCurrent -
*/
func (p *PortalTagHtml) AriaCurrent(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-current", value)
	return p
}

/*
AriaDescribedby -
*/
func (p *PortalTagHtml) AriaDescribedby(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-describedby", value)
	return p
}

/*
AriaDescription -
*/
func (p *PortalTagHtml) AriaDescription(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-description", value)
	return p
}

/*
AriaDetails -
*/
func (p *PortalTagHtml) AriaDetails(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-details", value)
	return p
}

/*
AriaDisabled -
*/
func (p *PortalTagHtml) AriaDisabled(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-disabled", value)
	return p
}

/*
AriaDropeffect -
*/
func (p *PortalTagHtml) AriaDropeffect(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-dropeffect", value)
	return p
}

/*
AriaErrormessage -
*/
func (p *PortalTagHtml) AriaErrormessage(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-errormessage", value)
	return p
}

/*
AriaFlowto -
*/
func (p *PortalTagHtml) AriaFlowto(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-flowto", value)
	return p
}

/*
AriaGrabbed -
*/
func (p *PortalTagHtml) AriaGrabbed(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-grabbed", value)
	return p
}

/*
AriaHaspopup -
*/
func (p *PortalTagHtml) AriaHaspopup(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-haspopup", value)
	return p
}

/*
AriaHidden -
*/
func (p *PortalTagHtml) AriaHidden(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-hidden", value)
	return p
}

/*
AriaInvalid -
*/
func (p *PortalTagHtml) AriaInvalid(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-invalid", value)
	return p
}

/*
AriaKeyshortcuts -
*/
func (p *PortalTagHtml) AriaKeyshortcuts(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-keyshortcuts", value)
	return p
}

/*
AriaLabel -
*/
func (p *PortalTagHtml) AriaLabel(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-label", value)
	return p
}

/*
AriaLabelledby -
*/
func (p *PortalTagHtml) AriaLabelledby(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-labelledby", value)
	return p
}

/*
AriaLive -
*/
func (p *PortalTagHtml) AriaLive(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-live", value)
	return p
}

/*
AriaOwns -
*/
func (p *PortalTagHtml) AriaOwns(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-owns", value)
	return p
}

/*
AriaRelevant -
*/
func (p *PortalTagHtml) AriaRelevant(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-relevant", value)
	return p
}

/*
AriaRoledescription -
*/
func (p *PortalTagHtml) AriaRoledescription(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria-roledescription", value)
	return p
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (p *PortalTagHtml) Onabort(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onabort", value)
	return p
}

/*
Onautocomplete -
*/
func (p *PortalTagHtml) Onautocomplete(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onautocomplete", value)
	return p
}

/*
Onautocompleteerror -
*/
func (p *PortalTagHtml) Onautocompleteerror(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onautocompleteerror", value)
	return p
}

/*
Onblur -
*/
func (p *PortalTagHtml) Onblur(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onblur", value)
	return p
}

/*
Oncancel -
*/
func (p *PortalTagHtml) Oncancel(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("oncancel", value)
	return p
}

/*
Oncanplay -
*/
func (p *PortalTagHtml) Oncanplay(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("oncanplay", value)
	return p
}

/*
Oncanplaythrough -
*/
func (p *PortalTagHtml) Oncanplaythrough(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("oncanplaythrough", value)
	return p
}

/*
Onchange -
*/
func (p *PortalTagHtml) Onchange(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onchange", value)
	return p
}

/*
Onclick -
*/
func (p *PortalTagHtml) Onclick(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onclick", value)
	return p
}

/*
Onclose -
*/
func (p *PortalTagHtml) Onclose(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onclose", value)
	return p
}

/*
Oncontextmenu -
*/
func (p *PortalTagHtml) Oncontextmenu(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("oncontextmenu", value)
	return p
}

/*
Oncuechange -
*/
func (p *PortalTagHtml) Oncuechange(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("oncuechange", value)
	return p
}

/*
Ondblclick -
*/
func (p *PortalTagHtml) Ondblclick(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondblclick", value)
	return p
}

/*
Ondrag -
*/
func (p *PortalTagHtml) Ondrag(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondrag", value)
	return p
}

/*
Ondragend -
*/
func (p *PortalTagHtml) Ondragend(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondragend", value)
	return p
}

/*
Ondragenter -
*/
func (p *PortalTagHtml) Ondragenter(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondragenter", value)
	return p
}

/*
Ondragleave -
*/
func (p *PortalTagHtml) Ondragleave(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondragleave", value)
	return p
}

/*
Ondragover -
*/
func (p *PortalTagHtml) Ondragover(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondragover", value)
	return p
}

/*
Ondragstart -
*/
func (p *PortalTagHtml) Ondragstart(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondragstart", value)
	return p
}

/*
Ondrop -
*/
func (p *PortalTagHtml) Ondrop(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondrop", value)
	return p
}

/*
Ondurationchange -
*/
func (p *PortalTagHtml) Ondurationchange(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ondurationchange", value)
	return p
}

/*
Onemptied -
*/
func (p *PortalTagHtml) Onemptied(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onemptied", value)
	return p
}

/*
Onended -
*/
func (p *PortalTagHtml) Onended(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onended", value)
	return p
}

/*
Onfocus -
*/
func (p *PortalTagHtml) Onfocus(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onfocus", value)
	return p
}

/*
Oninput -
*/
func (p *PortalTagHtml) Oninput(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("oninput", value)
	return p
}

/*
Oninvalid -
*/
func (p *PortalTagHtml) Oninvalid(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("oninvalid", value)
	return p
}

/*
Onkeydown -
*/
func (p *PortalTagHtml) Onkeydown(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onkeydown", value)
	return p
}

/*
Onkeypress -
*/
func (p *PortalTagHtml) Onkeypress(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onkeypress", value)
	return p
}

/*
Onkeyup -
*/
func (p *PortalTagHtml) Onkeyup(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onkeyup", value)
	return p
}

/*
Onloadeddata -
*/
func (p *PortalTagHtml) Onloadeddata(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onloadeddata", value)
	return p
}

/*
Onloadedmetadata -
*/
func (p *PortalTagHtml) Onloadedmetadata(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onloadedmetadata", value)
	return p
}

/*
Onloadstart -
*/
func (p *PortalTagHtml) Onloadstart(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onloadstart", value)
	return p
}

/*
Onmousedown -
*/
func (p *PortalTagHtml) Onmousedown(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmousedown", value)
	return p
}

/*
Onmouseenter -
*/
func (p *PortalTagHtml) Onmouseenter(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmouseenter", value)
	return p
}

/*
Onmouseleave -
*/
func (p *PortalTagHtml) Onmouseleave(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmouseleave", value)
	return p
}

/*
Onmousemove -
*/
func (p *PortalTagHtml) Onmousemove(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmousemove", value)
	return p
}

/*
Onmouseout -
*/
func (p *PortalTagHtml) Onmouseout(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmouseout", value)
	return p
}

/*
Onmouseover -
*/
func (p *PortalTagHtml) Onmouseover(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmouseover", value)
	return p
}

/*
Onmouseup -
*/
func (p *PortalTagHtml) Onmouseup(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmouseup", value)
	return p
}

/*
Onmousewheel -
*/
func (p *PortalTagHtml) Onmousewheel(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmousewheel", value)
	return p
}

/*
Onpause -
*/
func (p *PortalTagHtml) Onpause(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onpause", value)
	return p
}

/*
Onplay -
*/
func (p *PortalTagHtml) Onplay(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onplay", value)
	return p
}

/*
Onplaying -
*/
func (p *PortalTagHtml) Onplaying(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onplaying", value)
	return p
}

/*
Onprogress -
*/
func (p *PortalTagHtml) Onprogress(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onprogress", value)
	return p
}

/*
Onratechange -
*/
func (p *PortalTagHtml) Onratechange(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onratechange", value)
	return p
}

/*
Onreset -
*/
func (p *PortalTagHtml) Onreset(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onreset", value)
	return p
}

/*
Onscroll -
*/
func (p *PortalTagHtml) Onscroll(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onscroll", value)
	return p
}

/*
Onseeked -
*/
func (p *PortalTagHtml) Onseeked(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onseeked", value)
	return p
}

/*
Onseeking -
*/
func (p *PortalTagHtml) Onseeking(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onseeking", value)
	return p
}

/*
Onselect -
*/
func (p *PortalTagHtml) Onselect(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onselect", value)
	return p
}

/*
Onshow -
*/
func (p *PortalTagHtml) Onshow(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onshow", value)
	return p
}

/*
Onsort -
*/
func (p *PortalTagHtml) Onsort(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onsort", value)
	return p
}

/*
Onstalled -
*/
func (p *PortalTagHtml) Onstalled(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onstalled", value)
	return p
}

/*
Onsubmit -
*/
func (p *PortalTagHtml) Onsubmit(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onsubmit", value)
	return p
}

/*
Onsuspend -
*/
func (p *PortalTagHtml) Onsuspend(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onsuspend", value)
	return p
}

/*
Ontimeupdate -
*/
func (p *PortalTagHtml) Ontimeupdate(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ontimeupdate", value)
	return p
}

/*
Ontoggle -
*/
func (p *PortalTagHtml) Ontoggle(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ontoggle", value)
	return p
}

/*
Onvolumechange -
*/
func (p *PortalTagHtml) Onvolumechange(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onvolumechange", value)
	return p
}

/*
Onwaiting -
*/
func (p *PortalTagHtml) Onwaiting(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onwaiting", value)
	return p
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (p *PortalTagHtml) Onafterprint(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onafterprint", value)
	return p
}

/*
Onbeforeprint -
*/
func (p *PortalTagHtml) Onbeforeprint(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onbeforeprint", value)
	return p
}

/*
Onbeforeunload -
*/
func (p *PortalTagHtml) Onbeforeunload(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onbeforeunload", value)
	return p
}

/*
Onerror -
*/
func (p *PortalTagHtml) Onerror(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onerror", value)
	return p
}

/*
Onhashchange -
*/
func (p *PortalTagHtml) Onhashchange(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onhashchange", value)
	return p
}

/*
Onload -
*/
func (p *PortalTagHtml) Onload(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onload", value)
	return p
}

/*
Onmessage -
*/
func (p *PortalTagHtml) Onmessage(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onmessage", value)
	return p
}

/*
Onoffline -
*/
func (p *PortalTagHtml) Onoffline(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onoffline", value)
	return p
}

/*
Ononline -
*/
func (p *PortalTagHtml) Ononline(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("ononline", value)
	return p
}

/*
Onpagehide -
*/
func (p *PortalTagHtml) Onpagehide(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onpagehide", value)
	return p
}

/*
Onpageshow -
*/
func (p *PortalTagHtml) Onpageshow(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onpageshow", value)
	return p
}

/*
Onpopstate -
*/
func (p *PortalTagHtml) Onpopstate(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onpopstate", value)
	return p
}

/*
Onresize -
*/
func (p *PortalTagHtml) Onresize(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onresize", value)
	return p
}

/*
Onstorage -
*/
func (p *PortalTagHtml) Onstorage(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onstorage", value)
	return p
}

/*
Onunload -
*/
func (p *PortalTagHtml) Onunload(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("onunload", value)
	return p
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (p *PortalTagHtml) AccessKey(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("accessKey", value)
	return p
}

/*
Aria -
*/
func (p *PortalTagHtml) Aria(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("aria", value)
	return p
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (p *PortalTagHtml) Autocapitalize(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("autocapitalize", value)
	return p
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (p *PortalTagHtml) Autofocus(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("autofocus", value)
	return p
}

/*
Class -
*/
func (p *PortalTagHtml) Class(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("class", value)
	return p
}

/*
Contenteditable -
*/
func (p *PortalTagHtml) Contenteditable(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("contenteditable", value)
	return p
}

/*
Data -
*/
func (p *PortalTagHtml) Data(name, value string) *PortalTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	p.registerAttribute(dataName, value)
	return p
}

/*
Dir -
*/
func (p *PortalTagHtml) Dir(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("dir", value)
	return p
}

/*
Draggable -
*/
func (p *PortalTagHtml) Draggable(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("draggable", value)
	return p
}

/*
EnterKeyHint -
*/
func (p *PortalTagHtml) EnterKeyHint(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("enterKeyHint", value)
	return p
}

/*
ExportParts -
*/
func (p *PortalTagHtml) ExportParts(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("exportParts", value)
	return p
}

/*
Hidden -
*/
func (p *PortalTagHtml) Hidden(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("hidden", value)
	return p
}

/*
Id -
*/
func (p *PortalTagHtml) Id(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("id", value)
	return p
}

/*
Inert -
*/
func (p *PortalTagHtml) Inert(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("inert", value)
	return p
}

/*
InputMode -
*/
func (p *PortalTagHtml) InputMode(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("inputMode", value)
	return p
}

/*
Is -
*/
func (p *PortalTagHtml) Is(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("is", value)
	return p
}

/*
ItemId -
*/
func (p *PortalTagHtml) ItemId(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("itemId", value)
	return p
}

/*
ItemProp -
*/
func (p *PortalTagHtml) ItemProp(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("itemProp", value)
	return p
}

/*
ItemRef -
*/
func (p *PortalTagHtml) ItemRef(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("itemRef", value)
	return p
}

/*
ItemScope -
*/
func (p *PortalTagHtml) ItemScope(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("itemScope", value)
	return p
}

/*
ItemType -
*/
func (p *PortalTagHtml) ItemType(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("itemType", value)
	return p
}

/*
Lang -
*/
func (p *PortalTagHtml) Lang(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("lang", value)
	return p
}

/*
Nonce -
*/
func (p *PortalTagHtml) Nonce(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("nonce", value)
	return p
}

/*
Part -
*/
func (p *PortalTagHtml) Part(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("part", value)
	return p
}

/*
Popover -
*/
func (p *PortalTagHtml) Popover() *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("popover", "")
	return p
}

/*
Role -
*/
func (p *PortalTagHtml) Role(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("role", value)
	return p
}

/*
Slot -
*/
func (p *PortalTagHtml) Slot(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("slot", value)
	return p
}

/*
Spellcheck -
*/
func (p *PortalTagHtml) Spellcheck(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("spellcheck", value)
	return p
}

/*
Style -
*/
func (p *PortalTagHtml) Style(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("style", value)
	return p
}

/*
Tabindex -
*/
func (p *PortalTagHtml) Tabindex(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("tabindex", value)
	return p
}

/*
Title -
*/
func (p *PortalTagHtml) Title(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("title", value)
	return p
}

/*
Translate -
*/
func (p *PortalTagHtml) Translate(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("translate", value)
	return p
}

/*
VirtualKeyBoardPolicy -
*/
func (p *PortalTagHtml) VirtualKeyBoardPolicy(value string) *PortalTagHtml {
	if p.attributes == nil {
		p.attributes = []*Attribute{}
	}
	p.registerAttribute("virtualKeyBoardPolicy", value)
	return p
}
