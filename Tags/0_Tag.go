package tags

import (
	"context"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
	"io"
	"slices"
)

type content interface {
	getTag() *tag
	getName() string
	getAttributes() []*Attribute
	getSupportedAttributes() map[string][]string
	getTagType() spec.TagType
	getNamespace() spec.Namespace
	getChildren() []*tag
	getTextContent() string
	getParent() *tag
	getNestingLevel() int
	getRenderFormat() RenderFormat

	setNamespace(namespace spec.Namespace)
	setTextContent(content string)

	registerAttribute(name, value string)
	registerAttributes(attributes ...*Attribute)
	registerChildren(children ...*tag)
	registerParent(parent *tag)

	updateParent(parent *tag)
	updateNestingLevel(offset int)
	updateRenderFormat(fn RenderFormat)

	render(ctx context.Context, writer io.Writer) error
}

type tag struct {
	name                string
	attributes          []*Attribute
	supportedAttributes map[string][]string // This map contains name of supported group as key and list of attributes that are in this group
	tagType             spec.TagType
	namespace           spec.Namespace
	children            []*tag
	textContent         string
	parent              *tag
	nestingLevel        int
	renderFormat        RenderFormat
}

func (t *tag) getTag() *tag {
	return t
}

func (t *tag) getName() string {
	return t.name
}

func (t *tag) getAttributes() []*Attribute {
	return t.attributes
}

func (t *tag) getSupportedAttributes() map[string][]string {
	return t.supportedAttributes
}

func (t *tag) getTagType() spec.TagType {
	return t.tagType
}

func (t *tag) getNamespace() spec.Namespace {
	return t.namespace
}

func (t *tag) getChildren() []*tag {
	return t.children
}

func (t *tag) getTextContent() string {
	return t.textContent
}

func (t *tag) getParent() *tag {
	return t.parent
}

func (t *tag) getNestingLevel() int {
	return t.nestingLevel
}

func (t *tag) getRenderFormat() RenderFormat {
	return t.renderFormat
}

func (t *tag) setNamespace(namespace spec.Namespace) {
	t.namespace = namespace
}

func (t *tag) setTextContent(content string) {
	t.textContent = content
}

func (t *tag) registerAttribute(name, value string) {
	if t.attributes == nil {
		t.attributes = make([]*Attribute, 0)
	}
	if value == "" {
		t.attributes = append(t.attributes, &Attribute{
			Name:          name,
			Value:         value,
			Boolean:       true,
			attributeType: html5Compliant,
		})
		return
	}

	t.attributes = append(t.attributes, &Attribute{
		Name:          name,
		Value:         value,
		Boolean:       false,
		attributeType: html5Compliant,
	})
}

func (t *tag) registerAttributes(attributes ...*Attribute) {
	if t.attributes == nil {
		t.attributes = make([]*Attribute, 0)
	}
	for _, attribute := range attributes {
		t.attributes = append(t.attributes, &Attribute{
			Name:          attribute.Name,
			Value:         attribute.Value,
			Boolean:       attribute.Boolean,
			attributeType: custom,
		})
	}
}

func (t *tag) registerChildren(children ...*tag) {
	if t.children == nil {
		t.children = make([]*tag, 0)
	}
	t.children = append(t.children, children...)
	for _, child := range children {
		child.registerParent(t)
	}
}

func (t *tag) registerParent(parent *tag) {
	t.parent = parent
}

func (t *tag) updateParent(parent *tag) {
	t.parent = parent
	for _, child := range t.children {
		child.updateParent(t)
	}
}

func (t *tag) updateNestingLevel(offset int) {
	t.nestingLevel = offset
	for _, child := range t.children {
		child.updateNestingLevel(offset + 1)
	}
}

func (t *tag) updateRenderFormat(fn RenderFormat) {
	t.renderFormat = fn
	for _, child := range t.children {
		child.updateRenderFormat(t.renderFormat)
	}
}

func (t *tag) render(ctx context.Context, writer io.Writer) error {
	return t.renderFormat(ctx, writer, t)
}

func (t *tag) isSupported(name string) bool {
	if t.supportedAttributes == nil || len(t.supportedAttributes) == 0 {
		return false
	}

	for _, attributes := range t.supportedAttributes {
		if slices.Contains(attributes, name) {
			return true
		}
	}
	return false
}

func (t *tag) getCategoryName(name string) string {
	for category, attributes := range t.supportedAttributes {
		if slices.Contains(attributes, name) {
			return category
		}
	}
	panic("Error: this should not happened, because this method is called after was checked that this attribute is Supported by this tag.")
}
