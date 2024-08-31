// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func DfnHtml(tags []any) *DfnTagHtml {
	node := &DfnTagHtml{
		tag: &tag{
			name:                "dfn",
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

type DfnTagHtml struct {
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
func (d *DfnTagHtml) CustomAttribute(attributes ...*Attribute) *DfnTagHtml {
	d.registerAttributes(attributes...)
	return d
}

/*
Children - Method for nesting tags into parent tag
*/
func (d *DfnTagHtml) Children(tags ...any) *DfnTagHtml {
	return d.supportedChildrenCheck(tags)
}

func (d *DfnTagHtml) supportedChildrenCheck(tags []any) *DfnTagHtml {
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
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (d *DfnTagHtml) AriaAtomic(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-atomic", value)
	return d
}

/*
AriaBusy -
*/
func (d *DfnTagHtml) AriaBusy(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-busy", value)
	return d
}

/*
AriaControls -
*/
func (d *DfnTagHtml) AriaControls(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-controls", value)
	return d
}

/*
AriaCurrent -
*/
func (d *DfnTagHtml) AriaCurrent(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-current", value)
	return d
}

/*
AriaDescribedby -
*/
func (d *DfnTagHtml) AriaDescribedby(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-describedby", value)
	return d
}

/*
AriaDescription -
*/
func (d *DfnTagHtml) AriaDescription(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-description", value)
	return d
}

/*
AriaDetails -
*/
func (d *DfnTagHtml) AriaDetails(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-details", value)
	return d
}

/*
AriaDisabled -
*/
func (d *DfnTagHtml) AriaDisabled(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-disabled", value)
	return d
}

/*
AriaDropeffect -
*/
func (d *DfnTagHtml) AriaDropeffect(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-dropeffect", value)
	return d
}

/*
AriaErrormessage -
*/
func (d *DfnTagHtml) AriaErrormessage(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-errormessage", value)
	return d
}

/*
AriaFlowto -
*/
func (d *DfnTagHtml) AriaFlowto(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-flowto", value)
	return d
}

/*
AriaGrabbed -
*/
func (d *DfnTagHtml) AriaGrabbed(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-grabbed", value)
	return d
}

/*
AriaHaspopup -
*/
func (d *DfnTagHtml) AriaHaspopup(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-haspopup", value)
	return d
}

/*
AriaHidden -
*/
func (d *DfnTagHtml) AriaHidden(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-hidden", value)
	return d
}

/*
AriaInvalid -
*/
func (d *DfnTagHtml) AriaInvalid(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-invalid", value)
	return d
}

/*
AriaKeyshortcuts -
*/
func (d *DfnTagHtml) AriaKeyshortcuts(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-keyshortcuts", value)
	return d
}

/*
AriaLabel -
*/
func (d *DfnTagHtml) AriaLabel(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-label", value)
	return d
}

/*
AriaLabelledby -
*/
func (d *DfnTagHtml) AriaLabelledby(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-labelledby", value)
	return d
}

/*
AriaLive -
*/
func (d *DfnTagHtml) AriaLive(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-live", value)
	return d
}

/*
AriaOwns -
*/
func (d *DfnTagHtml) AriaOwns(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-owns", value)
	return d
}

/*
AriaRelevant -
*/
func (d *DfnTagHtml) AriaRelevant(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-relevant", value)
	return d
}

/*
AriaRoledescription -
*/
func (d *DfnTagHtml) AriaRoledescription(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("aria-roledescription", value)
	return d
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (d *DfnTagHtml) Onabort(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onabort", value)
	return d
}

/*
Onautocomplete -
*/
func (d *DfnTagHtml) Onautocomplete(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onautocomplete", value)
	return d
}

/*
Onautocompleteerror -
*/
func (d *DfnTagHtml) Onautocompleteerror(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onautocompleteerror", value)
	return d
}

/*
Onblur -
*/
func (d *DfnTagHtml) Onblur(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onblur", value)
	return d
}

/*
Oncancel -
*/
func (d *DfnTagHtml) Oncancel(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncancel", value)
	return d
}

/*
Oncanplay -
*/
func (d *DfnTagHtml) Oncanplay(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncanplay", value)
	return d
}

/*
Oncanplaythrough -
*/
func (d *DfnTagHtml) Oncanplaythrough(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncanplaythrough", value)
	return d
}

/*
Onchange -
*/
func (d *DfnTagHtml) Onchange(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onchange", value)
	return d
}

/*
Onclick -
*/
func (d *DfnTagHtml) Onclick(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onclick", value)
	return d
}

/*
Onclose -
*/
func (d *DfnTagHtml) Onclose(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onclose", value)
	return d
}

/*
Oncontextmenu -
*/
func (d *DfnTagHtml) Oncontextmenu(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncontextmenu", value)
	return d
}

/*
Oncuechange -
*/
func (d *DfnTagHtml) Oncuechange(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oncuechange", value)
	return d
}

/*
Ondblclick -
*/
func (d *DfnTagHtml) Ondblclick(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondblclick", value)
	return d
}

/*
Ondrag -
*/
func (d *DfnTagHtml) Ondrag(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondrag", value)
	return d
}

/*
Ondragend -
*/
func (d *DfnTagHtml) Ondragend(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragend", value)
	return d
}

/*
Ondragenter -
*/
func (d *DfnTagHtml) Ondragenter(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragenter", value)
	return d
}

/*
Ondragleave -
*/
func (d *DfnTagHtml) Ondragleave(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragleave", value)
	return d
}

/*
Ondragover -
*/
func (d *DfnTagHtml) Ondragover(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragover", value)
	return d
}

/*
Ondragstart -
*/
func (d *DfnTagHtml) Ondragstart(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondragstart", value)
	return d
}

/*
Ondrop -
*/
func (d *DfnTagHtml) Ondrop(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondrop", value)
	return d
}

/*
Ondurationchange -
*/
func (d *DfnTagHtml) Ondurationchange(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ondurationchange", value)
	return d
}

/*
Onemptied -
*/
func (d *DfnTagHtml) Onemptied(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onemptied", value)
	return d
}

/*
Onended -
*/
func (d *DfnTagHtml) Onended(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onended", value)
	return d
}

/*
Onfocus -
*/
func (d *DfnTagHtml) Onfocus(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onfocus", value)
	return d
}

/*
Oninput -
*/
func (d *DfnTagHtml) Oninput(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oninput", value)
	return d
}

/*
Oninvalid -
*/
func (d *DfnTagHtml) Oninvalid(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("oninvalid", value)
	return d
}

/*
Onkeydown -
*/
func (d *DfnTagHtml) Onkeydown(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onkeydown", value)
	return d
}

/*
Onkeypress -
*/
func (d *DfnTagHtml) Onkeypress(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onkeypress", value)
	return d
}

/*
Onkeyup -
*/
func (d *DfnTagHtml) Onkeyup(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onkeyup", value)
	return d
}

/*
Onloadeddata -
*/
func (d *DfnTagHtml) Onloadeddata(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onloadeddata", value)
	return d
}

/*
Onloadedmetadata -
*/
func (d *DfnTagHtml) Onloadedmetadata(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onloadedmetadata", value)
	return d
}

/*
Onloadstart -
*/
func (d *DfnTagHtml) Onloadstart(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onloadstart", value)
	return d
}

/*
Onmousedown -
*/
func (d *DfnTagHtml) Onmousedown(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmousedown", value)
	return d
}

/*
Onmouseenter -
*/
func (d *DfnTagHtml) Onmouseenter(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseenter", value)
	return d
}

/*
Onmouseleave -
*/
func (d *DfnTagHtml) Onmouseleave(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseleave", value)
	return d
}

/*
Onmousemove -
*/
func (d *DfnTagHtml) Onmousemove(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmousemove", value)
	return d
}

/*
Onmouseout -
*/
func (d *DfnTagHtml) Onmouseout(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseout", value)
	return d
}

/*
Onmouseover -
*/
func (d *DfnTagHtml) Onmouseover(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseover", value)
	return d
}

/*
Onmouseup -
*/
func (d *DfnTagHtml) Onmouseup(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmouseup", value)
	return d
}

/*
Onmousewheel -
*/
func (d *DfnTagHtml) Onmousewheel(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmousewheel", value)
	return d
}

/*
Onpause -
*/
func (d *DfnTagHtml) Onpause(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpause", value)
	return d
}

/*
Onplay -
*/
func (d *DfnTagHtml) Onplay(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onplay", value)
	return d
}

/*
Onplaying -
*/
func (d *DfnTagHtml) Onplaying(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onplaying", value)
	return d
}

/*
Onprogress -
*/
func (d *DfnTagHtml) Onprogress(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onprogress", value)
	return d
}

/*
Onratechange -
*/
func (d *DfnTagHtml) Onratechange(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onratechange", value)
	return d
}

/*
Onreset -
*/
func (d *DfnTagHtml) Onreset(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onreset", value)
	return d
}

/*
Onscroll -
*/
func (d *DfnTagHtml) Onscroll(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onscroll", value)
	return d
}

/*
Onseeked -
*/
func (d *DfnTagHtml) Onseeked(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onseeked", value)
	return d
}

/*
Onseeking -
*/
func (d *DfnTagHtml) Onseeking(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onseeking", value)
	return d
}

/*
Onselect -
*/
func (d *DfnTagHtml) Onselect(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onselect", value)
	return d
}

/*
Onshow -
*/
func (d *DfnTagHtml) Onshow(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onshow", value)
	return d
}

/*
Onsort -
*/
func (d *DfnTagHtml) Onsort(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onsort", value)
	return d
}

/*
Onstalled -
*/
func (d *DfnTagHtml) Onstalled(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onstalled", value)
	return d
}

/*
Onsubmit -
*/
func (d *DfnTagHtml) Onsubmit(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onsubmit", value)
	return d
}

/*
Onsuspend -
*/
func (d *DfnTagHtml) Onsuspend(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onsuspend", value)
	return d
}

/*
Ontimeupdate -
*/
func (d *DfnTagHtml) Ontimeupdate(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ontimeupdate", value)
	return d
}

/*
Ontoggle -
*/
func (d *DfnTagHtml) Ontoggle(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ontoggle", value)
	return d
}

/*
Onvolumechange -
*/
func (d *DfnTagHtml) Onvolumechange(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onvolumechange", value)
	return d
}

/*
Onwaiting -
*/
func (d *DfnTagHtml) Onwaiting(value string) *DfnTagHtml {
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
func (d *DfnTagHtml) Onafterprint(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onafterprint", value)
	return d
}

/*
Onbeforeprint -
*/
func (d *DfnTagHtml) Onbeforeprint(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onbeforeprint", value)
	return d
}

/*
Onbeforeunload -
*/
func (d *DfnTagHtml) Onbeforeunload(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onbeforeunload", value)
	return d
}

/*
Onerror -
*/
func (d *DfnTagHtml) Onerror(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onerror", value)
	return d
}

/*
Onhashchange -
*/
func (d *DfnTagHtml) Onhashchange(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onhashchange", value)
	return d
}

/*
Onload -
*/
func (d *DfnTagHtml) Onload(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onload", value)
	return d
}

/*
Onmessage -
*/
func (d *DfnTagHtml) Onmessage(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onmessage", value)
	return d
}

/*
Onoffline -
*/
func (d *DfnTagHtml) Onoffline(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onoffline", value)
	return d
}

/*
Ononline -
*/
func (d *DfnTagHtml) Ononline(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("ononline", value)
	return d
}

/*
Onpagehide -
*/
func (d *DfnTagHtml) Onpagehide(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpagehide", value)
	return d
}

/*
Onpageshow -
*/
func (d *DfnTagHtml) Onpageshow(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpageshow", value)
	return d
}

/*
Onpopstate -
*/
func (d *DfnTagHtml) Onpopstate(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onpopstate", value)
	return d
}

/*
Onresize -
*/
func (d *DfnTagHtml) Onresize(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onresize", value)
	return d
}

/*
Onstorage -
*/
func (d *DfnTagHtml) Onstorage(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("onstorage", value)
	return d
}

/*
Onunload -
*/
func (d *DfnTagHtml) Onunload(value string) *DfnTagHtml {
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
func (d *DfnTagHtml) AccessKey(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("accessKey", value)
	return d
}

/*
Aria -
*/
func (d *DfnTagHtml) Aria(value string) *DfnTagHtml {
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
func (d *DfnTagHtml) Autocapitalize(value string) *DfnTagHtml {
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
func (d *DfnTagHtml) Autofocus(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("autofocus", value)
	return d
}

/*
Class -
*/
func (d *DfnTagHtml) Class(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("class", value)
	return d
}

/*
Contenteditable -
*/
func (d *DfnTagHtml) Contenteditable(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("contenteditable", value)
	return d
}

/*
Data -
*/
func (d *DfnTagHtml) Data(name, value string) *DfnTagHtml {
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
func (d *DfnTagHtml) Dir(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("dir", value)
	return d
}

/*
Draggable -
*/
func (d *DfnTagHtml) Draggable(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("draggable", value)
	return d
}

/*
EnterKeyHint -
*/
func (d *DfnTagHtml) EnterKeyHint(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("enterKeyHint", value)
	return d
}

/*
ExportParts -
*/
func (d *DfnTagHtml) ExportParts(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("exportParts", value)
	return d
}

/*
Hidden -
*/
func (d *DfnTagHtml) Hidden(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("hidden", value)
	return d
}

/*
Id -
*/
func (d *DfnTagHtml) Id(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("id", value)
	return d
}

/*
Inert -
*/
func (d *DfnTagHtml) Inert(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("inert", value)
	return d
}

/*
InputMode -
*/
func (d *DfnTagHtml) InputMode(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("inputMode", value)
	return d
}

/*
Is -
*/
func (d *DfnTagHtml) Is(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("is", value)
	return d
}

/*
ItemId -
*/
func (d *DfnTagHtml) ItemId(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemId", value)
	return d
}

/*
ItemProp -
*/
func (d *DfnTagHtml) ItemProp(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemProp", value)
	return d
}

/*
ItemRef -
*/
func (d *DfnTagHtml) ItemRef(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemRef", value)
	return d
}

/*
ItemScope -
*/
func (d *DfnTagHtml) ItemScope(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemScope", value)
	return d
}

/*
ItemType -
*/
func (d *DfnTagHtml) ItemType(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("itemType", value)
	return d
}

/*
Lang -
*/
func (d *DfnTagHtml) Lang(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("lang", value)
	return d
}

/*
Nonce -
*/
func (d *DfnTagHtml) Nonce(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("nonce", value)
	return d
}

/*
Part -
*/
func (d *DfnTagHtml) Part(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("part", value)
	return d
}

/*
Popover -
*/
func (d *DfnTagHtml) Popover() *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("popover", "")
	return d
}

/*
Role -
*/
func (d *DfnTagHtml) Role(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("role", value)
	return d
}

/*
Slot -
*/
func (d *DfnTagHtml) Slot(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("slot", value)
	return d
}

/*
Spellcheck -
*/
func (d *DfnTagHtml) Spellcheck(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("spellcheck", value)
	return d
}

/*
Style -
*/
func (d *DfnTagHtml) Style(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("style", value)
	return d
}

/*
Tabindex -
*/
func (d *DfnTagHtml) Tabindex(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("tabindex", value)
	return d
}

/*
Title -
*/
func (d *DfnTagHtml) Title(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("title", value)
	return d
}

/*
Translate -
*/
func (d *DfnTagHtml) Translate(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("translate", value)
	return d
}

/*
VirtualKeyBoardPolicy -
*/
func (d *DfnTagHtml) VirtualKeyBoardPolicy(value string) *DfnTagHtml {
	if d.attributes == nil {
		d.attributes = []*Attribute{}
	}
	d.registerAttribute("virtualKeyBoardPolicy", value)
	return d
}
