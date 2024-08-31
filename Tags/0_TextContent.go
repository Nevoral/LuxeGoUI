package tags

import (
	spec "github.com/Nevoral/LuxeGoUI/Specification"
	"strings"
)

func TextContentHtml(content string) *TextContentNode {
	return &TextContentNode{
		tag: &tag{
			name:                "TextContent",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.TextContentType,
			namespace:           spec.HTML,
			children:            nil,
			textContent:         strings.TrimSpace(content),
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

func TextContentSvg(content string) *TextContentNode {
	return &TextContentNode{
		tag: &tag{
			name:                "TextContent",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.TextContentType,
			namespace:           spec.SVG,
			children:            nil,
			textContent:         strings.TrimSpace(content),
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

func TextContentMath(content string) *TextContentNode {
	return &TextContentNode{
		tag: &tag{
			name:                "TextContent",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.TextContentType,
			namespace:           spec.MATH,
			children:            nil,
			textContent:         strings.TrimSpace(content),
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

type TextContentNode struct {
	*tag
}
