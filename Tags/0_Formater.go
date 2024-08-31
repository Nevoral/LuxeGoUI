package tags

import (
	"bytes"
	"context"
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
	"golang.org/x/net/html"
	"slices"
	"strings"
)

var voidTag = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "source", "track", "wbr"}

type HtmlFormater struct {
	RawHtml          []byte
	DocType          *DOCTYPETagHtml
	HtmlAst          *tag
	InitNestingLevel int
	RenderFormat     RenderFormat
	Output           bytes.Buffer
	OutputFunc       func(output []byte) error // for terminal output file output...
}

func (h *HtmlFormater) GenerateFormatedOutput() error {
	doc, err := html.Parse(bytes.NewReader(h.RawHtml))
	if err != nil {
		return err
	}

	if err = h.scannerHtml(doc, h.HtmlAst); err != nil {
		return err
	}
	h.HtmlAst.updateNestingLevel(h.InitNestingLevel)
	h.HtmlAst.updateParent(nil)
	if !IsSameFormat(h.RenderFormat, DefaultHtml) {
		h.DocType.updateRenderFormat(h.RenderFormat)
		h.HtmlAst.updateRenderFormat(h.RenderFormat)
	}
	if err = h.DocType.render(context.Background(), &h.Output); err != nil {
		return err
	}
	if err = h.HtmlAst.render(context.Background(), &h.Output); err != nil {
		return err
	}
	return h.OutputFunc(h.Output.Bytes())
}

func (h *HtmlFormater) scannerHtml(n *html.Node, parentTag *tag) error {
	switch n.Type {
	case html.DocumentNode:
		return h.scannerHtml(n.FirstChild, parentTag)
	case html.RawNode:
		return h.scannerHtml(n.FirstChild, parentTag)
	case html.TextNode:
		textContent := strings.TrimSpace(n.Data)
		if textContent == "" {
			return nil
		}
		if parentTag == nil {
			h.HtmlAst = &tag{
				name:                "TextContent",
				attributes:          nil,
				supportedAttributes: nil,
				tagType:             spec.TextContentType,
				namespace:           h.DocType.getNamespace(),
				children:            nil,
				textContent:         textContent,
				parent:              nil,
				nestingLevel:        0,
				renderFormat:        DefaultHtml,
			}
		} else {
			parentTag.registerChildren(&tag{
				name:                "TextContent",
				attributes:          nil,
				supportedAttributes: nil,
				tagType:             spec.TextContentType,
				namespace:           h.DocType.getNamespace(),
				children:            nil,
				textContent:         textContent,
				parent:              parentTag,
				nestingLevel:        0,
				renderFormat:        DefaultHtml,
			})
		}
		return nil
	case html.CommentNode:
		comment := strings.TrimSpace(n.Data)
		if comment == "" {
			return nil
		}
		if parentTag == nil {
			h.HtmlAst = &tag{
				name:                "!--",
				attributes:          nil,
				supportedAttributes: nil,
				tagType:             spec.CommentType,
				namespace:           h.DocType.getNamespace(),
				children:            nil,
				textContent:         comment,
				parent:              parentTag,
				nestingLevel:        0,
				renderFormat:        DefaultHtml,
			}
		} else {
			parentTag.registerChildren(&tag{
				name:                "!--",
				attributes:          nil,
				supportedAttributes: nil,
				tagType:             spec.CommentType,
				namespace:           h.DocType.getNamespace(),
				children:            nil,
				textContent:         comment,
				parent:              parentTag,
				nestingLevel:        0,
				renderFormat:        DefaultHtml,
			})
		}
		return nil
	case html.DoctypeNode:
		h.DocType = DOCTYPEHtml()
		if strings.ToLower(n.Data) == "svg" {
			h.DocType.Svg()
		} else if strings.ToLower(n.Data) == "math" {
			h.DocType.Math()
		} else {
			h.DocType.Html()
		}
		for _, attr := range n.Attr {
			if attr.Key == "public" {
				h.DocType.PUBLIC(attr.Val)
				continue
			}
			h.DocType.SYSTEM(attr.Val)
		}
		c := n.NextSibling
		if c != nil {
			err := h.scannerHtml(c, nil)
			if err != nil {
				return err
			}
		}
		return nil
	case html.ElementNode:
		var (
			node          tag
			attributeList []*Attribute
			namespace     spec.Namespace
		)

		if n.Namespace == "svg" {
			namespace = spec.SVG
		} else if n.Namespace == "math" {
			namespace = spec.MATH
		} else if n.Namespace == "" && h.DocType.getNamespace() == spec.XHTML {
			namespace = spec.XHTML
		} else {
			namespace = spec.HTML
		}

		if n.Attr == nil || len(n.Attr) == 0 {
			attributeList = nil
		}
		for _, attr := range n.Attr {
			if attr.Val == "" {
				attributeList = append(attributeList, &Attribute{
					Name:          attr.Key,
					Value:         "",
					Boolean:       true,
					attributeType: unknown,
				})
			} else {
				attributeList = append(attributeList, &Attribute{
					Name:          attr.Key,
					Value:         attr.Val,
					Boolean:       false,
					attributeType: unknown,
				})
			}
		}
		if slices.Contains(voidTag, n.Data) {
			node = tag{
				name:                n.Data,
				attributes:          attributeList,
				supportedAttributes: nil,
				tagType:             spec.SelfClosingType,
				namespace:           namespace,
				children:            nil,
				textContent:         "",
				parent:              parentTag,
				nestingLevel:        0,
				renderFormat:        DefaultHtml,
			}
		} else {
			node = tag{
				name:                n.Data,
				attributes:          attributeList,
				supportedAttributes: nil,
				tagType:             spec.FullTagType,
				namespace:           namespace,
				children:            nil,
				textContent:         "",
				parent:              parentTag,
				nestingLevel:        0,
				renderFormat:        DefaultHtml,
			}
		}
		if parentTag == nil {
			h.HtmlAst = &node
		} else {
			parentTag.registerChildren(&node)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := h.scannerHtml(c, &node)
			if err != nil {
				return err
			}
		}
	default:
		panic(fmt.Sprintf("unhandled default case %d", n.Type))
	}
	return nil
}
