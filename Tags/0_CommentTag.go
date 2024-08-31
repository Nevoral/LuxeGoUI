package tags

import (
	spec "github.com/Nevoral/LuxeGoUI/Specification"
	"strings"
)

func CommentHtml(comment string) *CommentTagHtml {
	return &CommentTagHtml{
		tag: &tag{
			name:                "!--",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.CommentType,
			namespace:           spec.HTML,
			children:            nil,
			textContent:         strings.TrimSpace(comment),
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

func CommentSvg(comment string) *CommentTagHtml {
	return &CommentTagHtml{
		tag: &tag{
			name:                "!--",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.CommentType,
			namespace:           spec.SVG,
			children:            nil,
			textContent:         strings.TrimSpace(comment),
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

func CommentMath(comment string) *CommentTagHtml {
	return &CommentTagHtml{
		tag: &tag{
			name:                "!--",
			attributes:          nil,
			supportedAttributes: nil,
			tagType:             spec.CommentType,
			namespace:           spec.MATH,
			children:            nil,
			textContent:         strings.TrimSpace(comment),
			parent:              nil,
			nestingLevel:        0,
			renderFormat:        DefaultHtml,
		},
	}
}

type CommentTagHtml struct {
	*tag
}
