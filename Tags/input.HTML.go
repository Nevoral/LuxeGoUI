// Package tags do not edit, this file was autogenerated.
package tags

import (
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
)

func InputHtml() *InputTagHtml {
	return &InputTagHtml{
		tag: &tag{
			name:                "input",
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

type InputTagHtml struct {
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
func (i *InputTagHtml) CustomAttribute(attributes ...*Attribute) *InputTagHtml {
	i.registerAttributes(attributes...)
	return i
}

/*
************************************************************************************************************************
*------------------------------------------------ Specific Attributes -------------------------------------------------*
************************************************************************************************************************
 */

/*
Accept -
*/
func (i *InputTagHtml) Accept(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("accept", value)
	return i
}

/*
Alt -
*/
func (i *InputTagHtml) Alt(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("alt", value)
	return i
}

/*
Autocomplete -
*/
func (i *InputTagHtml) Autocomplete(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("autocomplete", value)
	return i
}

/*
Autocorrect -
*/
func (i *InputTagHtml) Autocorrect(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("autocorrect", value)
	return i
}

/*
Capture -
*/
func (i *InputTagHtml) Capture(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("capture", value)
	return i
}

/*
Checked -
*/
func (i *InputTagHtml) Checked() *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("checked", "")
	return i
}

/*
Dirname -
*/
func (i *InputTagHtml) Dirname(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("dirname", value)
	return i
}

/*
Disabled -
*/
func (i *InputTagHtml) Disabled() *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("disabled", "")
	return i
}

/*
Form -
*/
func (i *InputTagHtml) Form(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("form", value)
	return i
}

/*
Formaction -
*/
func (i *InputTagHtml) Formaction(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("formaction", value)
	return i
}

/*
Formenctype -
*/
func (i *InputTagHtml) Formenctype(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("formenctype", value)
	return i
}

/*
Formmethod -
*/
func (i *InputTagHtml) Formmethod(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("formmethod", value)
	return i
}

/*
Formnovalidate -
*/
func (i *InputTagHtml) Formnovalidate(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("formnovalidate", value)
	return i
}

/*
Formtarget -
*/
func (i *InputTagHtml) Formtarget(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("formtarget", value)
	return i
}

/*
Height -
*/
func (i *InputTagHtml) Height(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("height", value)
	return i
}

/*
Incremental -
*/
func (i *InputTagHtml) Incremental(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("incremental", value)
	return i
}

/*
Inputmode -
*/
func (i *InputTagHtml) Inputmode(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inputmode", value)
	return i
}

/*
List -
*/
func (i *InputTagHtml) List(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("list", value)
	return i
}

/*
Max -
*/
func (i *InputTagHtml) Max(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("max", value)
	return i
}

/*
Maxlength -
*/
func (i *InputTagHtml) Maxlength(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("maxlength", value)
	return i
}

/*
Min -
*/
func (i *InputTagHtml) Min(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("min", value)
	return i
}

/*
Minlength -
*/
func (i *InputTagHtml) Minlength(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("minlength", value)
	return i
}

/*
Multiple -
*/
func (i *InputTagHtml) Multiple() *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("multiple", "")
	return i
}

/*
Name -
*/
func (i *InputTagHtml) Name(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("name", value)
	return i
}

/*
Orient -
*/
func (i *InputTagHtml) Orient(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("orient", value)
	return i
}

/*
Pattern -
*/
func (i *InputTagHtml) Pattern(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("pattern", value)
	return i
}

/*
Placeholder -
*/
func (i *InputTagHtml) Placeholder(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("placeholder", value)
	return i
}

/*
Popovertarget -
*/
func (i *InputTagHtml) Popovertarget(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("popovertarget", value)
	return i
}

/*
Popovertargetaction -
*/
func (i *InputTagHtml) Popovertargetaction(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("popovertargetaction", value)
	return i
}

/*
Readonly -
*/
func (i *InputTagHtml) Readonly() *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("readonly", "")
	return i
}

/*
Required -
*/
func (i *InputTagHtml) Required() *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("required", "")
	return i
}

/*
Results -
*/
func (i *InputTagHtml) Results(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("results", value)
	return i
}

/*
Size -
*/
func (i *InputTagHtml) Size(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("size", value)
	return i
}

/*
Src - Specifies the URL of an image.
Specifies the URL of an image.
*/
func (i *InputTagHtml) Src(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("src", value)
	return i
}

/*
Step -
*/
func (i *InputTagHtml) Step(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("step", value)
	return i
}

/*
Type - Specifies the type of an <input> element.
Specifies the type of an <input> element.
*/
func (i *InputTagHtml) Type(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("type", value)
	return i
}

/*
Value - Specifies the value of an <input> element.
Specifies the value of an <input> element.
*/
func (i *InputTagHtml) Value(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("value", value)
	return i
}

/*
Webkitdirectory -
*/
func (i *InputTagHtml) Webkitdirectory(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("webkitdirectory", value)
	return i
}

/*
Width -
*/
func (i *InputTagHtml) Width(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("width", value)
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
func (i *InputTagHtml) AccessKey(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("accessKey", value)
	return i
}

/*
Aria -
*/
func (i *InputTagHtml) Aria(value string) *InputTagHtml {
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
func (i *InputTagHtml) Autocapitalize(value string) *InputTagHtml {
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
func (i *InputTagHtml) Autofocus(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("autofocus", value)
	return i
}

/*
Class -
*/
func (i *InputTagHtml) Class(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("class", value)
	return i
}

/*
Contenteditable -
*/
func (i *InputTagHtml) Contenteditable(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("contenteditable", value)
	return i
}

/*
Data -
*/
func (i *InputTagHtml) Data(name, value string) *InputTagHtml {
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
func (i *InputTagHtml) Dir(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("dir", value)
	return i
}

/*
Draggable -
*/
func (i *InputTagHtml) Draggable(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("draggable", value)
	return i
}

/*
EnterKeyHint -
*/
func (i *InputTagHtml) EnterKeyHint(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("enterKeyHint", value)
	return i
}

/*
ExportParts -
*/
func (i *InputTagHtml) ExportParts(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("exportParts", value)
	return i
}

/*
Hidden -
*/
func (i *InputTagHtml) Hidden(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("hidden", value)
	return i
}

/*
Id -
*/
func (i *InputTagHtml) Id(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("id", value)
	return i
}

/*
Inert -
*/
func (i *InputTagHtml) Inert(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inert", value)
	return i
}

/*
InputMode -
*/
func (i *InputTagHtml) InputMode(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("inputMode", value)
	return i
}

/*
Is -
*/
func (i *InputTagHtml) Is(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("is", value)
	return i
}

/*
ItemId -
*/
func (i *InputTagHtml) ItemId(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemId", value)
	return i
}

/*
ItemProp -
*/
func (i *InputTagHtml) ItemProp(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemProp", value)
	return i
}

/*
ItemRef -
*/
func (i *InputTagHtml) ItemRef(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemRef", value)
	return i
}

/*
ItemScope -
*/
func (i *InputTagHtml) ItemScope(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemScope", value)
	return i
}

/*
ItemType -
*/
func (i *InputTagHtml) ItemType(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("itemType", value)
	return i
}

/*
Lang -
*/
func (i *InputTagHtml) Lang(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("lang", value)
	return i
}

/*
Nonce -
*/
func (i *InputTagHtml) Nonce(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("nonce", value)
	return i
}

/*
Part -
*/
func (i *InputTagHtml) Part(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("part", value)
	return i
}

/*
Popover -
*/
func (i *InputTagHtml) Popover() *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("popover", "")
	return i
}

/*
Role -
*/
func (i *InputTagHtml) Role(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("role", value)
	return i
}

/*
Slot -
*/
func (i *InputTagHtml) Slot(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("slot", value)
	return i
}

/*
Spellcheck -
*/
func (i *InputTagHtml) Spellcheck(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("spellcheck", value)
	return i
}

/*
Style -
*/
func (i *InputTagHtml) Style(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("style", value)
	return i
}

/*
Tabindex -
*/
func (i *InputTagHtml) Tabindex(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("tabindex", value)
	return i
}

/*
Title -
*/
func (i *InputTagHtml) Title(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("title", value)
	return i
}

/*
Translate -
*/
func (i *InputTagHtml) Translate(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("translate", value)
	return i
}

/*
VirtualKeyBoardPolicy -
*/
func (i *InputTagHtml) VirtualKeyBoardPolicy(value string) *InputTagHtml {
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
func (i *InputTagHtml) AriaAtomic(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-atomic", value)
	return i
}

/*
AriaBusy -
*/
func (i *InputTagHtml) AriaBusy(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-busy", value)
	return i
}

/*
AriaControls -
*/
func (i *InputTagHtml) AriaControls(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-controls", value)
	return i
}

/*
AriaCurrent -
*/
func (i *InputTagHtml) AriaCurrent(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-current", value)
	return i
}

/*
AriaDescribedby -
*/
func (i *InputTagHtml) AriaDescribedby(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-describedby", value)
	return i
}

/*
AriaDescription -
*/
func (i *InputTagHtml) AriaDescription(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-description", value)
	return i
}

/*
AriaDetails -
*/
func (i *InputTagHtml) AriaDetails(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-details", value)
	return i
}

/*
AriaDisabled -
*/
func (i *InputTagHtml) AriaDisabled(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-disabled", value)
	return i
}

/*
AriaDropeffect -
*/
func (i *InputTagHtml) AriaDropeffect(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-dropeffect", value)
	return i
}

/*
AriaErrormessage -
*/
func (i *InputTagHtml) AriaErrormessage(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-errormessage", value)
	return i
}

/*
AriaFlowto -
*/
func (i *InputTagHtml) AriaFlowto(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-flowto", value)
	return i
}

/*
AriaGrabbed -
*/
func (i *InputTagHtml) AriaGrabbed(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-grabbed", value)
	return i
}

/*
AriaHaspopup -
*/
func (i *InputTagHtml) AriaHaspopup(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-haspopup", value)
	return i
}

/*
AriaHidden -
*/
func (i *InputTagHtml) AriaHidden(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-hidden", value)
	return i
}

/*
AriaInvalid -
*/
func (i *InputTagHtml) AriaInvalid(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-invalid", value)
	return i
}

/*
AriaKeyshortcuts -
*/
func (i *InputTagHtml) AriaKeyshortcuts(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-keyshortcuts", value)
	return i
}

/*
AriaLabel -
*/
func (i *InputTagHtml) AriaLabel(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-label", value)
	return i
}

/*
AriaLabelledby -
*/
func (i *InputTagHtml) AriaLabelledby(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-labelledby", value)
	return i
}

/*
AriaLive -
*/
func (i *InputTagHtml) AriaLive(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-live", value)
	return i
}

/*
AriaOwns -
*/
func (i *InputTagHtml) AriaOwns(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-owns", value)
	return i
}

/*
AriaRelevant -
*/
func (i *InputTagHtml) AriaRelevant(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-relevant", value)
	return i
}

/*
AriaRoledescription -
*/
func (i *InputTagHtml) AriaRoledescription(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("aria-roledescription", value)
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
func (i *InputTagHtml) Onabort(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onabort", value)
	return i
}

/*
Onautocomplete -
*/
func (i *InputTagHtml) Onautocomplete(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocomplete", value)
	return i
}

/*
Onautocompleteerror -
*/
func (i *InputTagHtml) Onautocompleteerror(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onautocompleteerror", value)
	return i
}

/*
Onblur -
*/
func (i *InputTagHtml) Onblur(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onblur", value)
	return i
}

/*
Oncancel -
*/
func (i *InputTagHtml) Oncancel(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncancel", value)
	return i
}

/*
Oncanplay -
*/
func (i *InputTagHtml) Oncanplay(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplay", value)
	return i
}

/*
Oncanplaythrough -
*/
func (i *InputTagHtml) Oncanplaythrough(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncanplaythrough", value)
	return i
}

/*
Onchange -
*/
func (i *InputTagHtml) Onchange(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onchange", value)
	return i
}

/*
Onclick -
*/
func (i *InputTagHtml) Onclick(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclick", value)
	return i
}

/*
Onclose -
*/
func (i *InputTagHtml) Onclose(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onclose", value)
	return i
}

/*
Oncontextmenu -
*/
func (i *InputTagHtml) Oncontextmenu(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncontextmenu", value)
	return i
}

/*
Oncuechange -
*/
func (i *InputTagHtml) Oncuechange(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oncuechange", value)
	return i
}

/*
Ondblclick -
*/
func (i *InputTagHtml) Ondblclick(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondblclick", value)
	return i
}

/*
Ondrag -
*/
func (i *InputTagHtml) Ondrag(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrag", value)
	return i
}

/*
Ondragend -
*/
func (i *InputTagHtml) Ondragend(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragend", value)
	return i
}

/*
Ondragenter -
*/
func (i *InputTagHtml) Ondragenter(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragenter", value)
	return i
}

/*
Ondragleave -
*/
func (i *InputTagHtml) Ondragleave(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragleave", value)
	return i
}

/*
Ondragover -
*/
func (i *InputTagHtml) Ondragover(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragover", value)
	return i
}

/*
Ondragstart -
*/
func (i *InputTagHtml) Ondragstart(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondragstart", value)
	return i
}

/*
Ondrop -
*/
func (i *InputTagHtml) Ondrop(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondrop", value)
	return i
}

/*
Ondurationchange -
*/
func (i *InputTagHtml) Ondurationchange(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ondurationchange", value)
	return i
}

/*
Onemptied -
*/
func (i *InputTagHtml) Onemptied(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onemptied", value)
	return i
}

/*
Onended -
*/
func (i *InputTagHtml) Onended(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onended", value)
	return i
}

/*
Onfocus -
*/
func (i *InputTagHtml) Onfocus(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onfocus", value)
	return i
}

/*
Oninput -
*/
func (i *InputTagHtml) Oninput(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninput", value)
	return i
}

/*
Oninvalid -
*/
func (i *InputTagHtml) Oninvalid(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("oninvalid", value)
	return i
}

/*
Onkeydown -
*/
func (i *InputTagHtml) Onkeydown(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeydown", value)
	return i
}

/*
Onkeypress -
*/
func (i *InputTagHtml) Onkeypress(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeypress", value)
	return i
}

/*
Onkeyup -
*/
func (i *InputTagHtml) Onkeyup(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onkeyup", value)
	return i
}

/*
Onloadeddata -
*/
func (i *InputTagHtml) Onloadeddata(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadeddata", value)
	return i
}

/*
Onloadedmetadata -
*/
func (i *InputTagHtml) Onloadedmetadata(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadedmetadata", value)
	return i
}

/*
Onloadstart -
*/
func (i *InputTagHtml) Onloadstart(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onloadstart", value)
	return i
}

/*
Onmousedown -
*/
func (i *InputTagHtml) Onmousedown(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousedown", value)
	return i
}

/*
Onmouseenter -
*/
func (i *InputTagHtml) Onmouseenter(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseenter", value)
	return i
}

/*
Onmouseleave -
*/
func (i *InputTagHtml) Onmouseleave(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseleave", value)
	return i
}

/*
Onmousemove -
*/
func (i *InputTagHtml) Onmousemove(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousemove", value)
	return i
}

/*
Onmouseout -
*/
func (i *InputTagHtml) Onmouseout(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseout", value)
	return i
}

/*
Onmouseover -
*/
func (i *InputTagHtml) Onmouseover(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseover", value)
	return i
}

/*
Onmouseup -
*/
func (i *InputTagHtml) Onmouseup(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmouseup", value)
	return i
}

/*
Onmousewheel -
*/
func (i *InputTagHtml) Onmousewheel(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmousewheel", value)
	return i
}

/*
Onpause -
*/
func (i *InputTagHtml) Onpause(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpause", value)
	return i
}

/*
Onplay -
*/
func (i *InputTagHtml) Onplay(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplay", value)
	return i
}

/*
Onplaying -
*/
func (i *InputTagHtml) Onplaying(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onplaying", value)
	return i
}

/*
Onprogress -
*/
func (i *InputTagHtml) Onprogress(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onprogress", value)
	return i
}

/*
Onratechange -
*/
func (i *InputTagHtml) Onratechange(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onratechange", value)
	return i
}

/*
Onreset -
*/
func (i *InputTagHtml) Onreset(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onreset", value)
	return i
}

/*
Onscroll -
*/
func (i *InputTagHtml) Onscroll(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onscroll", value)
	return i
}

/*
Onseeked -
*/
func (i *InputTagHtml) Onseeked(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeked", value)
	return i
}

/*
Onseeking -
*/
func (i *InputTagHtml) Onseeking(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onseeking", value)
	return i
}

/*
Onselect -
*/
func (i *InputTagHtml) Onselect(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onselect", value)
	return i
}

/*
Onshow -
*/
func (i *InputTagHtml) Onshow(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onshow", value)
	return i
}

/*
Onsort -
*/
func (i *InputTagHtml) Onsort(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsort", value)
	return i
}

/*
Onstalled -
*/
func (i *InputTagHtml) Onstalled(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstalled", value)
	return i
}

/*
Onsubmit -
*/
func (i *InputTagHtml) Onsubmit(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsubmit", value)
	return i
}

/*
Onsuspend -
*/
func (i *InputTagHtml) Onsuspend(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onsuspend", value)
	return i
}

/*
Ontimeupdate -
*/
func (i *InputTagHtml) Ontimeupdate(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontimeupdate", value)
	return i
}

/*
Ontoggle -
*/
func (i *InputTagHtml) Ontoggle(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ontoggle", value)
	return i
}

/*
Onvolumechange -
*/
func (i *InputTagHtml) Onvolumechange(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onvolumechange", value)
	return i
}

/*
Onwaiting -
*/
func (i *InputTagHtml) Onwaiting(value string) *InputTagHtml {
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
func (i *InputTagHtml) Onafterprint(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onafterprint", value)
	return i
}

/*
Onbeforeprint -
*/
func (i *InputTagHtml) Onbeforeprint(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeprint", value)
	return i
}

/*
Onbeforeunload -
*/
func (i *InputTagHtml) Onbeforeunload(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onbeforeunload", value)
	return i
}

/*
Onerror -
*/
func (i *InputTagHtml) Onerror(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onerror", value)
	return i
}

/*
Onhashchange -
*/
func (i *InputTagHtml) Onhashchange(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onhashchange", value)
	return i
}

/*
Onload -
*/
func (i *InputTagHtml) Onload(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onload", value)
	return i
}

/*
Onmessage -
*/
func (i *InputTagHtml) Onmessage(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onmessage", value)
	return i
}

/*
Onoffline -
*/
func (i *InputTagHtml) Onoffline(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onoffline", value)
	return i
}

/*
Ononline -
*/
func (i *InputTagHtml) Ononline(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("ononline", value)
	return i
}

/*
Onpagehide -
*/
func (i *InputTagHtml) Onpagehide(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpagehide", value)
	return i
}

/*
Onpageshow -
*/
func (i *InputTagHtml) Onpageshow(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpageshow", value)
	return i
}

/*
Onpopstate -
*/
func (i *InputTagHtml) Onpopstate(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onpopstate", value)
	return i
}

/*
Onresize -
*/
func (i *InputTagHtml) Onresize(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onresize", value)
	return i
}

/*
Onstorage -
*/
func (i *InputTagHtml) Onstorage(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onstorage", value)
	return i
}

/*
Onunload -
*/
func (i *InputTagHtml) Onunload(value string) *InputTagHtml {
	if i.attributes == nil {
		i.attributes = []*Attribute{}
	}
	i.registerAttribute("onunload", value)
	return i
}
