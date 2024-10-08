// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func AudioHtml(tags []any) *AudioTagHtml {
	node := &AudioTagHtml{
		tag: &tag{
			name:                "audio",
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

type AudioTagHtml struct {
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
func (a *AudioTagHtml) CustomAttribute(attributes ...*Attribute) *AudioTagHtml {
	a.registerAttributes(attributes...)
	return a
}

/*
Children - Method for nesting tags into parent tag
*/
func (a *AudioTagHtml) Children(tags ...any) *AudioTagHtml {
	return a.supportedChildrenCheck(tags)
}

func (a *AudioTagHtml) supportedChildrenCheck(tags []any) *AudioTagHtml {
	for _, tagObj := range tags {
		switch val := tagObj.(type) {
		case string:
			a.registerChildren(TextContentSvg(val).getTag())
		case content:
			a.registerChildren(val.getTag())
		case []content:
			for _, child := range val {
				a.registerChildren(child.getTag())
			}
		default:
			panic(fmt.Sprintf("Unsupported content type. %v", val))
		}
	}
	return a
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Autoplay - A Boolean attribute: if specified, the audio will automatically begin playback as soon as it can do so, without waiting for the entire audio file to finish downloading.
A Boolean attribute: if specified, the audio will automatically begin playback as soon as it can do so, without waiting for the entire audio file to finish downloading.

	Note: Sites that automatically play audio (or videos with an audio track) can be an unpleasant experience for users, so should be avoided when possible. If you must offer autoplay functionality, you should make it opt-in (requiring a user to specifically enable it). However, this can be useful when creating media elements whose source will be set at a later time, under user control. See our autoplay guide for additional information about how to properly use autoplay.
*/
func (a *AudioTagHtml) Autoplay() *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("autoplay", "")
	return a
}

/*
Controls - If this attribute is present, the browser will offer controls to allow the user to control audio playback, including volume, seeking, and pause/resume playback.
If this attribute is present, the browser will offer controls to allow the user to control audio playback, including volume, seeking, and pause/resume playback.
*/
func (a *AudioTagHtml) Controls() *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("controls", "")
	return a
}

/*
Controlslist - The controlslist attribute, when specified, helps the browser select what controls to show for the audio element whenever the browser shows its own set of controls (that is, when the controls attribute is specified).
The controlslist attribute, when specified, helps the browser select what controls to show for the audio element whenever the browser shows its own set of controls (that is, when the controls attribute is specified).

	The allowed values are nodownload, nofullscreen and noremoteplayback.
*/
func (a *AudioTagHtml) Controlslist(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("controlslist", value)
	return a
}

/*
Crossorigin - This enumerated attribute indicates whether to use CORS to fetch the related audio file. CORS-enabled resources can be reused in the <canvas> element without being tainted. The allowed values are:
This enumerated attribute indicates whether to use CORS to fetch the related audio file. CORS-enabled resources can be reused in the <canvas> element without being tainted. The allowed values are:

	anonymous

	    Sends a cross-origin request without a credential. In other words, it sends the Origin: HTTP header without a cookie, X.509 certificate, or performing HTTP Basic authentication. If the server does not give credentials to the origin site (by not setting the Access-Control-Allow-Origin: HTTP header), the resource will be tainted, and its usage restricted.
	use-credentials

	    Sends a cross-origin request with a credential. In other words, it sends the Origin: HTTP header with a cookie, a certificate, or performing HTTP Basic authentication. If the server does not give credentials to the origin site (through Access-Control-Allow-Credentials: HTTP header), the resource will be tainted and its usage restricted.

	When not present, the resource is fetched without a CORS request (i.e. without sending the Origin: HTTP header), preventing its non-tainted use in <canvas> elements. If invalid, it is handled as if the enumerated keyword anonymous was used. See CORS settings attributes for additional information.
*/
func (a *AudioTagHtml) Crossorigin(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("crossorigin", value)
	return a
}

/*
Disableremoteplayback - A Boolean attribute used to disable the capability of remote playback in devices that are attached using wired (HDMI, DVI, etc.) and wireless technologies (Miracast, Chromecast, DLNA, AirPlay, etc.). See this proposed specification for more information.
A Boolean attribute used to disable the capability of remote playback in devices that are attached using wired (HDMI, DVI, etc.) and wireless technologies (Miracast, Chromecast, DLNA, AirPlay, etc.). See this proposed specification for more information.

	In Safari, you can use x-webkit-airplay="deny" as a fallback.
*/
func (a *AudioTagHtml) Disableremoteplayback(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("disableremoteplayback", value)
	return a
}

/*
Loop - A Boolean attribute: if specified, the audio player will automatically seek back to the start upon reaching the end of the audio.
A Boolean attribute: if specified, the audio player will automatically seek back to the start upon reaching the end of the audio.
*/
func (a *AudioTagHtml) Loop() *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("loop", "")
	return a
}

/*
Muted - A Boolean attribute that indicates whether the audio will be initially silenced. Its default value is false.
A Boolean attribute that indicates whether the audio will be initially silenced. Its default value is false.
*/
func (a *AudioTagHtml) Muted() *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("muted", "")
	return a
}

/*
Preload - This enumerated attribute is intended to provide a hint to the browser about what the author thinks will lead to the best user experience. It may have one of the following values:
This enumerated attribute is intended to provide a hint to the browser about what the author thinks will lead to the best user experience. It may have one of the following values:

	    none: Indicates that the audio should not be preloaded.
	    metadata: Indicates that only audio metadata (e.g. length) is fetched.
	    auto: Indicates that the whole audio file can be downloaded, even if the user is not expected to use it.
	    empty string: A synonym of the auto value.

	The default value is different for each browser. The spec advises it to be set to metadata.

	Note:

	    The autoplay attribute has precedence over preload. If autoplay is specified, the browser would obviously need to start downloading the audio for playback.
	    The browser is not forced by the specification to follow the value of this attribute; it is a mere hint.
*/
func (a *AudioTagHtml) Preload(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("preload", value)
	return a
}

/*
Src - The URL of the audio to embed. This is subject to HTTP access controls. This is optional; you may instead use the <source> element within the audio block to specify the audio to embed.
The URL of the audio to embed. This is subject to HTTP access controls. This is optional; you may instead use the <source> element within the audio block to specify the audio to embed.
*/
func (a *AudioTagHtml) Src(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("src", value)
	return a
}

/*
************************************************************************************************************************
*------------------------------------------------- Global Attributes --------------------------------------------------*
************************************************************************************************************************
 */

/*
AccessKey -
*/
func (a *AudioTagHtml) AccessKey(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("accessKey", value)
	return a
}

/*
Aria -
*/
func (a *AudioTagHtml) Aria(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria", value)
	return a
}

/*
Autocapitalize - Controls whether and how text input is automatically capitalized.
Controls whether and how text input is automatically capitalized.
*/
func (a *AudioTagHtml) Autocapitalize(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("autocapitalize", value)
	return a
}

/*
Autofocus - Specifies that an element should automatically get focus when the page loads.
Specifies that an element should automatically get focus when the page loads.
*/
func (a *AudioTagHtml) Autofocus(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("autofocus", value)
	return a
}

/*
Class -
*/
func (a *AudioTagHtml) Class(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("class", value)
	return a
}

/*
Contenteditable -
*/
func (a *AudioTagHtml) Contenteditable(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("contenteditable", value)
	return a
}

/*
Data -
*/
func (a *AudioTagHtml) Data(name, value string) *AudioTagHtml {
	var dataName string
	if name == "" {
		dataName = "data"
	} else {
		dataName = fmt.Sprintf("data-%s", name)
	}
	a.registerAttribute(dataName, value)
	return a
}

/*
Dir -
*/
func (a *AudioTagHtml) Dir(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("dir", value)
	return a
}

/*
Draggable -
*/
func (a *AudioTagHtml) Draggable(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("draggable", value)
	return a
}

/*
EnterKeyHint -
*/
func (a *AudioTagHtml) EnterKeyHint(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("enterKeyHint", value)
	return a
}

/*
ExportParts -
*/
func (a *AudioTagHtml) ExportParts(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("exportParts", value)
	return a
}

/*
Hidden -
*/
func (a *AudioTagHtml) Hidden(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("hidden", value)
	return a
}

/*
Id -
*/
func (a *AudioTagHtml) Id(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("id", value)
	return a
}

/*
Inert -
*/
func (a *AudioTagHtml) Inert(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("inert", value)
	return a
}

/*
InputMode -
*/
func (a *AudioTagHtml) InputMode(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("inputMode", value)
	return a
}

/*
Is -
*/
func (a *AudioTagHtml) Is(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("is", value)
	return a
}

/*
ItemId -
*/
func (a *AudioTagHtml) ItemId(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("itemId", value)
	return a
}

/*
ItemProp -
*/
func (a *AudioTagHtml) ItemProp(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("itemProp", value)
	return a
}

/*
ItemRef -
*/
func (a *AudioTagHtml) ItemRef(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("itemRef", value)
	return a
}

/*
ItemScope -
*/
func (a *AudioTagHtml) ItemScope(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("itemScope", value)
	return a
}

/*
ItemType -
*/
func (a *AudioTagHtml) ItemType(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("itemType", value)
	return a
}

/*
Lang -
*/
func (a *AudioTagHtml) Lang(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("lang", value)
	return a
}

/*
Nonce -
*/
func (a *AudioTagHtml) Nonce(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("nonce", value)
	return a
}

/*
Part -
*/
func (a *AudioTagHtml) Part(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("part", value)
	return a
}

/*
Popover -
*/
func (a *AudioTagHtml) Popover() *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("popover", "")
	return a
}

/*
Role -
*/
func (a *AudioTagHtml) Role(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("role", value)
	return a
}

/*
Slot -
*/
func (a *AudioTagHtml) Slot(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("slot", value)
	return a
}

/*
Spellcheck -
*/
func (a *AudioTagHtml) Spellcheck(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("spellcheck", value)
	return a
}

/*
Style -
*/
func (a *AudioTagHtml) Style(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("style", value)
	return a
}

/*
Tabindex -
*/
func (a *AudioTagHtml) Tabindex(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("tabindex", value)
	return a
}

/*
Title -
*/
func (a *AudioTagHtml) Title(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("title", value)
	return a
}

/*
Translate -
*/
func (a *AudioTagHtml) Translate(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("translate", value)
	return a
}

/*
VirtualKeyBoardPolicy -
*/
func (a *AudioTagHtml) VirtualKeyBoardPolicy(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("virtualKeyBoardPolicy", value)
	return a
}

/*
************************************************************************************************************************
*-------------------------------------------------- Aria Attributes ---------------------------------------------------*
************************************************************************************************************************
 */

/*
AriaAtomic -
*/
func (a *AudioTagHtml) AriaAtomic(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-atomic", value)
	return a
}

/*
AriaBusy -
*/
func (a *AudioTagHtml) AriaBusy(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-busy", value)
	return a
}

/*
AriaControls -
*/
func (a *AudioTagHtml) AriaControls(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-controls", value)
	return a
}

/*
AriaCurrent -
*/
func (a *AudioTagHtml) AriaCurrent(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-current", value)
	return a
}

/*
AriaDescribedby -
*/
func (a *AudioTagHtml) AriaDescribedby(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-describedby", value)
	return a
}

/*
AriaDescription -
*/
func (a *AudioTagHtml) AriaDescription(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-description", value)
	return a
}

/*
AriaDetails -
*/
func (a *AudioTagHtml) AriaDetails(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-details", value)
	return a
}

/*
AriaDisabled -
*/
func (a *AudioTagHtml) AriaDisabled(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-disabled", value)
	return a
}

/*
AriaDropeffect -
*/
func (a *AudioTagHtml) AriaDropeffect(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-dropeffect", value)
	return a
}

/*
AriaErrormessage -
*/
func (a *AudioTagHtml) AriaErrormessage(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-errormessage", value)
	return a
}

/*
AriaFlowto -
*/
func (a *AudioTagHtml) AriaFlowto(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-flowto", value)
	return a
}

/*
AriaGrabbed -
*/
func (a *AudioTagHtml) AriaGrabbed(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-grabbed", value)
	return a
}

/*
AriaHaspopup -
*/
func (a *AudioTagHtml) AriaHaspopup(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-haspopup", value)
	return a
}

/*
AriaHidden -
*/
func (a *AudioTagHtml) AriaHidden(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-hidden", value)
	return a
}

/*
AriaInvalid -
*/
func (a *AudioTagHtml) AriaInvalid(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-invalid", value)
	return a
}

/*
AriaKeyshortcuts -
*/
func (a *AudioTagHtml) AriaKeyshortcuts(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-keyshortcuts", value)
	return a
}

/*
AriaLabel -
*/
func (a *AudioTagHtml) AriaLabel(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-label", value)
	return a
}

/*
AriaLabelledby -
*/
func (a *AudioTagHtml) AriaLabelledby(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-labelledby", value)
	return a
}

/*
AriaLive -
*/
func (a *AudioTagHtml) AriaLive(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-live", value)
	return a
}

/*
AriaOwns -
*/
func (a *AudioTagHtml) AriaOwns(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-owns", value)
	return a
}

/*
AriaRelevant -
*/
func (a *AudioTagHtml) AriaRelevant(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-relevant", value)
	return a
}

/*
AriaRoledescription -
*/
func (a *AudioTagHtml) AriaRoledescription(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("aria-roledescription", value)
	return a
}

/*
************************************************************************************************************************
*-------------------------------------------- Document Action Attributes ----------------------------------------------*
************************************************************************************************************************
 */

/*
Onabort -
*/
func (a *AudioTagHtml) Onabort(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onabort", value)
	return a
}

/*
Onautocomplete -
*/
func (a *AudioTagHtml) Onautocomplete(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onautocomplete", value)
	return a
}

/*
Onautocompleteerror -
*/
func (a *AudioTagHtml) Onautocompleteerror(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onautocompleteerror", value)
	return a
}

/*
Onblur -
*/
func (a *AudioTagHtml) Onblur(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onblur", value)
	return a
}

/*
Oncancel -
*/
func (a *AudioTagHtml) Oncancel(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("oncancel", value)
	return a
}

/*
Oncanplay -
*/
func (a *AudioTagHtml) Oncanplay(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("oncanplay", value)
	return a
}

/*
Oncanplaythrough -
*/
func (a *AudioTagHtml) Oncanplaythrough(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("oncanplaythrough", value)
	return a
}

/*
Onchange -
*/
func (a *AudioTagHtml) Onchange(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onchange", value)
	return a
}

/*
Onclick -
*/
func (a *AudioTagHtml) Onclick(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onclick", value)
	return a
}

/*
Onclose -
*/
func (a *AudioTagHtml) Onclose(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onclose", value)
	return a
}

/*
Oncontextmenu -
*/
func (a *AudioTagHtml) Oncontextmenu(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("oncontextmenu", value)
	return a
}

/*
Oncuechange -
*/
func (a *AudioTagHtml) Oncuechange(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("oncuechange", value)
	return a
}

/*
Ondblclick -
*/
func (a *AudioTagHtml) Ondblclick(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondblclick", value)
	return a
}

/*
Ondrag -
*/
func (a *AudioTagHtml) Ondrag(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondrag", value)
	return a
}

/*
Ondragend -
*/
func (a *AudioTagHtml) Ondragend(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondragend", value)
	return a
}

/*
Ondragenter -
*/
func (a *AudioTagHtml) Ondragenter(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondragenter", value)
	return a
}

/*
Ondragleave -
*/
func (a *AudioTagHtml) Ondragleave(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondragleave", value)
	return a
}

/*
Ondragover -
*/
func (a *AudioTagHtml) Ondragover(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondragover", value)
	return a
}

/*
Ondragstart -
*/
func (a *AudioTagHtml) Ondragstart(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondragstart", value)
	return a
}

/*
Ondrop -
*/
func (a *AudioTagHtml) Ondrop(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondrop", value)
	return a
}

/*
Ondurationchange -
*/
func (a *AudioTagHtml) Ondurationchange(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ondurationchange", value)
	return a
}

/*
Onemptied -
*/
func (a *AudioTagHtml) Onemptied(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onemptied", value)
	return a
}

/*
Onended -
*/
func (a *AudioTagHtml) Onended(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onended", value)
	return a
}

/*
Onfocus -
*/
func (a *AudioTagHtml) Onfocus(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onfocus", value)
	return a
}

/*
Oninput -
*/
func (a *AudioTagHtml) Oninput(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("oninput", value)
	return a
}

/*
Oninvalid -
*/
func (a *AudioTagHtml) Oninvalid(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("oninvalid", value)
	return a
}

/*
Onkeydown -
*/
func (a *AudioTagHtml) Onkeydown(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onkeydown", value)
	return a
}

/*
Onkeypress -
*/
func (a *AudioTagHtml) Onkeypress(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onkeypress", value)
	return a
}

/*
Onkeyup -
*/
func (a *AudioTagHtml) Onkeyup(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onkeyup", value)
	return a
}

/*
Onloadeddata -
*/
func (a *AudioTagHtml) Onloadeddata(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onloadeddata", value)
	return a
}

/*
Onloadedmetadata -
*/
func (a *AudioTagHtml) Onloadedmetadata(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onloadedmetadata", value)
	return a
}

/*
Onloadstart -
*/
func (a *AudioTagHtml) Onloadstart(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onloadstart", value)
	return a
}

/*
Onmousedown -
*/
func (a *AudioTagHtml) Onmousedown(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmousedown", value)
	return a
}

/*
Onmouseenter -
*/
func (a *AudioTagHtml) Onmouseenter(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmouseenter", value)
	return a
}

/*
Onmouseleave -
*/
func (a *AudioTagHtml) Onmouseleave(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmouseleave", value)
	return a
}

/*
Onmousemove -
*/
func (a *AudioTagHtml) Onmousemove(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmousemove", value)
	return a
}

/*
Onmouseout -
*/
func (a *AudioTagHtml) Onmouseout(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmouseout", value)
	return a
}

/*
Onmouseover -
*/
func (a *AudioTagHtml) Onmouseover(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmouseover", value)
	return a
}

/*
Onmouseup -
*/
func (a *AudioTagHtml) Onmouseup(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmouseup", value)
	return a
}

/*
Onmousewheel -
*/
func (a *AudioTagHtml) Onmousewheel(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmousewheel", value)
	return a
}

/*
Onpause -
*/
func (a *AudioTagHtml) Onpause(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onpause", value)
	return a
}

/*
Onplay -
*/
func (a *AudioTagHtml) Onplay(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onplay", value)
	return a
}

/*
Onplaying -
*/
func (a *AudioTagHtml) Onplaying(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onplaying", value)
	return a
}

/*
Onprogress -
*/
func (a *AudioTagHtml) Onprogress(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onprogress", value)
	return a
}

/*
Onratechange -
*/
func (a *AudioTagHtml) Onratechange(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onratechange", value)
	return a
}

/*
Onreset -
*/
func (a *AudioTagHtml) Onreset(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onreset", value)
	return a
}

/*
Onscroll -
*/
func (a *AudioTagHtml) Onscroll(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onscroll", value)
	return a
}

/*
Onseeked -
*/
func (a *AudioTagHtml) Onseeked(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onseeked", value)
	return a
}

/*
Onseeking -
*/
func (a *AudioTagHtml) Onseeking(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onseeking", value)
	return a
}

/*
Onselect -
*/
func (a *AudioTagHtml) Onselect(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onselect", value)
	return a
}

/*
Onshow -
*/
func (a *AudioTagHtml) Onshow(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onshow", value)
	return a
}

/*
Onsort -
*/
func (a *AudioTagHtml) Onsort(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onsort", value)
	return a
}

/*
Onstalled -
*/
func (a *AudioTagHtml) Onstalled(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onstalled", value)
	return a
}

/*
Onsubmit -
*/
func (a *AudioTagHtml) Onsubmit(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onsubmit", value)
	return a
}

/*
Onsuspend -
*/
func (a *AudioTagHtml) Onsuspend(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onsuspend", value)
	return a
}

/*
Ontimeupdate -
*/
func (a *AudioTagHtml) Ontimeupdate(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ontimeupdate", value)
	return a
}

/*
Ontoggle -
*/
func (a *AudioTagHtml) Ontoggle(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ontoggle", value)
	return a
}

/*
Onvolumechange -
*/
func (a *AudioTagHtml) Onvolumechange(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onvolumechange", value)
	return a
}

/*
Onwaiting -
*/
func (a *AudioTagHtml) Onwaiting(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onwaiting", value)
	return a
}

/*
************************************************************************************************************************
*--------------------------------------------- Window Action Attributes -----------------------------------------------*
************************************************************************************************************************
 */

/*
Onafterprint -
*/
func (a *AudioTagHtml) Onafterprint(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onafterprint", value)
	return a
}

/*
Onbeforeprint -
*/
func (a *AudioTagHtml) Onbeforeprint(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onbeforeprint", value)
	return a
}

/*
Onbeforeunload -
*/
func (a *AudioTagHtml) Onbeforeunload(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onbeforeunload", value)
	return a
}

/*
Onerror -
*/
func (a *AudioTagHtml) Onerror(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onerror", value)
	return a
}

/*
Onhashchange -
*/
func (a *AudioTagHtml) Onhashchange(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onhashchange", value)
	return a
}

/*
Onload -
*/
func (a *AudioTagHtml) Onload(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onload", value)
	return a
}

/*
Onmessage -
*/
func (a *AudioTagHtml) Onmessage(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onmessage", value)
	return a
}

/*
Onoffline -
*/
func (a *AudioTagHtml) Onoffline(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onoffline", value)
	return a
}

/*
Ononline -
*/
func (a *AudioTagHtml) Ononline(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("ononline", value)
	return a
}

/*
Onpagehide -
*/
func (a *AudioTagHtml) Onpagehide(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onpagehide", value)
	return a
}

/*
Onpageshow -
*/
func (a *AudioTagHtml) Onpageshow(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onpageshow", value)
	return a
}

/*
Onpopstate -
*/
func (a *AudioTagHtml) Onpopstate(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onpopstate", value)
	return a
}

/*
Onresize -
*/
func (a *AudioTagHtml) Onresize(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onresize", value)
	return a
}

/*
Onstorage -
*/
func (a *AudioTagHtml) Onstorage(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onstorage", value)
	return a
}

/*
Onunload -
*/
func (a *AudioTagHtml) Onunload(value string) *AudioTagHtml {
	if a.attributes == nil {
		a.attributes = []*Attribute{}
	}
	a.registerAttribute("onunload", value)
	return a
}
