package tags

import (
	"bytes"
	"context"
	"fmt"
	spec "github.com/Nevoral/LuxeGoUI/Specification"
	"io"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

type RenderFormat func(ctx context.Context, writer io.Writer, t *tag) error

var (
	EfficientHtml RenderFormat = efficientHtml
	DefaultHtml   RenderFormat = defaultHtml
	LuxeGoFunc    RenderFormat = luxeGoFunc
	LuxeGoTempl   RenderFormat = luxeGoTempl
)

// IsSameFormat - Compare two RenderFormat functions using reflect
func IsSameFormat(f1, f2 RenderFormat) bool {
	return reflect.ValueOf(f1).Pointer() == reflect.ValueOf(f2).Pointer()
}

func efficientHtml(ctx context.Context, writer io.Writer, t *tag) (err error) {
	switch t.getTagType() {
	case spec.DocumentType:
		var page bytes.Buffer
		for _, child := range t.getChildren() {
			err = child.render(ctx, &page)
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(writer, "%s", page.String())
	case spec.ComponentType:
		var component bytes.Buffer
		for _, child := range t.getChildren() {
			err = child.render(ctx, &component)
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(writer, "%s", component.String())
	case spec.DoctypeType:
		attr := t.getAttributes()
		if len(attr) > 1 {
			switch attr[1].Name {
			case "SYSTEM":
				_, err = fmt.Fprintf(writer, "<!DOCTYPE %s SYSTEM \"%s\">", attr[0].Name, attr[1].Value)
			default:
				_, err = fmt.Fprintf(writer, "<!DOCTYPE %s PUBLIC \"%s\" \"%s\">", attr[0].Name, attr[1].Value, attr[2].Value)
			}
			return err
		}
		if len(attr) == 1 {
			_, err = fmt.Fprintf(writer, "<!DOCTYPE %s>", attr[0].Name)
			return err
		}
		_, err = fmt.Fprintf(writer, "<!DOCTYPE html>")
	case spec.CommentType:
		re := regexp.MustCompile(`\s+`)
		_, err = fmt.Fprintf(writer, "<!--%s-->", re.ReplaceAllString(t.getTextContent(), " "))
	case spec.TextContentType:
		re := regexp.MustCompile(`\s+`)
		_, err = fmt.Fprintf(writer, "%s", re.ReplaceAllString(t.getTextContent(), " "))
	case spec.SelfClosingType:
		if t.getNamespace() == spec.HTML {
			_, err = fmt.Fprintf(writer, "<%s%s>", t.getName(), efficientAttributes(t.getAttributes()))
			return err
		}
		_, err = fmt.Fprintf(writer, "<%s%s />", t.getName(), efficientAttributes(t.getAttributes()))
	default:
		_, err = fmt.Fprintf(writer, "<%s%s>", t.getName(), efficientAttributes(t.getAttributes()))
		if err != nil {
			return err
		}
		var result bytes.Buffer
		for _, child := range t.getChildren() {
			err = child.render(ctx, &result)
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(writer, "%s</%s>", result.String(), t.getName())
	}
	return err
}

func efficientAttributes(attr []*Attribute) string {
	if attr == nil {
		return ""
	}
	var result string
	for _, val := range attr {
		result += val.renderHtmlAttribute()
	}
	return result
}

func defaultHtml(ctx context.Context, writer io.Writer, t *tag) (err error) {
	indent := strings.Repeat(" ", t.getNestingLevel()*4)
	switch t.getTagType() {
	case spec.DocumentType:
		var page bytes.Buffer
		for _, child := range t.getChildren() {
			err = child.render(ctx, &page)
			if err != nil {
				return err
			}
			page.WriteString("\n" + indent)
		}
		_, err = fmt.Fprintf(writer, "%s", page.String())
	case spec.ComponentType:
		var component bytes.Buffer
		for _, child := range t.getChildren() {
			err = child.render(ctx, &component)
			if err != nil {
				return err
			}
			component.WriteString("\n" + indent)
		}
		_, err = fmt.Fprintf(writer, "%s", component.String())
	case spec.DoctypeType:
		attr := t.getAttributes()
		if len(attr) > 1 {
			switch attr[1].Name {
			case "SYSTEM":
				_, err = fmt.Fprintf(writer, "%s<!DOCTYPE %s SYSTEM \"%s\">\n", indent, attr[0].Name, attr[1].Value)
			default:
				_, err = fmt.Fprintf(writer, "%s<!DOCTYPE %s PUBLIC \"%s\" \"%s\">\n", indent, attr[0].Name, attr[1].Value, attr[2].Value)
			}
			return err
		}
		if len(attr) == 1 {
			_, err = fmt.Fprintf(writer, "%s<!DOCTYPE %s>\n", indent, attr[0].Name)
			return err
		}
		_, err = fmt.Fprintf(writer, "%s<!DOCTYPE html>\n", indent)
	case spec.CommentType:
		commentLine := strings.Split(strings.TrimSpace(t.getTextContent()), "\n")
		if len(commentLine) > 1 {
			comment := fmt.Sprintf("%s<!--", indent)
			for _, line := range commentLine {
				comment += fmt.Sprintf("\n%s%s", indent, strings.TrimSpace(line))
			}
			_, err = fmt.Fprintf(writer, "%s\n%s-->", comment, indent)
			return err
		}
		_, err = fmt.Fprintf(writer, "%s<!--%s-->", indent, t.getTextContent())
	case spec.TextContentType:
		textLine := strings.Split(strings.TrimSpace(t.getTextContent()), "\n")
		if len(textLine) > 1 {
			text := ""
			for _, line := range textLine {
				text += fmt.Sprintf("\n%s%s", indent, strings.TrimSpace(line))
			}
			_, err = fmt.Fprintf(writer, "%s", text[1:])
			return err
		}
		_, err = fmt.Fprintf(writer, "%s%s", indent, t.getTextContent())
	case spec.SelfClosingType:
		if t.getNamespace() == spec.HTML {
			_, err = fmt.Fprintf(writer, "%s<%s%s>", indent, t.getName(), defaultAttributes(t.getAttributes(), t.getNestingLevel()))
			return err
		}
		_, err = fmt.Fprintf(writer, "%s<%s%s />", indent, t.getName(), defaultAttributes(t.getAttributes(), t.getNestingLevel()))
	default:
		openTag := fmt.Sprintf("%s<%s%s>", indent, t.getName(), defaultAttributes(t.getAttributes(), t.getNestingLevel()))
		_, err = fmt.Fprintf(writer, openTag)

		var result bytes.Buffer
		for _, child := range t.getChildren() {
			if len(t.getChildren()) == 1 && child.getName() == "TextContent" {
				if len(openTag)+len(child.getTextContent()) < 110 {
					err = child.render(ctx, &result)
					_, err = fmt.Fprintf(writer, "%s</%s>", strings.TrimSpace(result.String()), t.getName())
					return err
				}
			}
			result.WriteString("\n")
			err = child.render(ctx, &result)
			if err != nil {
				return err
			}
		}
		if result.Len() == 0 {
			_, err = fmt.Fprintf(writer, "</%s>", t.getName())
			return err
		}
		_, err = fmt.Fprintf(writer, "%s\n%s</%s>", result.String(), indent, t.getName())
	}
	return err
}

func defaultAttributes(attr []*Attribute, offset int) string {
	if attr == nil {
		return ""
	}

	var (
		result      string
		line        string
		lineCounter = 0
	)
	indent := strings.Repeat(" ", (offset+1)*4)

	for _, val := range attr {
		attrStr := val.renderHtmlAttribute()
		if len(line)+len(attr)+offset*4 >= 110 {
			if len(line) < len(attr) {
				result += attrStr + "\n" + indent
				lineCounter++
				continue
			}
			result += line + "\n" + indent
			line = attrStr
			lineCounter++
			continue
		}
		line += attrStr
	}
	if lineCounter > 0 {
		result += line + "\n" + strings.Repeat(" ", offset*4)
	} else {
		result += line
	}
	return result
}

func luxeGoFunc(ctx context.Context, writer io.Writer, t *tag) (err error) {
	indent := strings.Repeat("\t", t.getNestingLevel())
	switch t.getTagType() {
	case spec.DocumentType:
		var page bytes.Buffer
		for _, child := range t.getChildren() {
			err = child.render(ctx, &page)
			if err != nil {
				return err
			}
			page.WriteString("\n" + indent)
		}
		_, err = fmt.Fprintf(writer, "%s", page.String())
	case spec.ComponentType:
		var component bytes.Buffer
		for _, child := range t.getChildren() {
			err = child.render(ctx, &component)
			if err != nil {
				return err
			}
			component.WriteString("\n" + indent)
		}
		_, err = fmt.Fprintf(writer, "%s", component.String())
	case spec.DoctypeType:
		if len(t.getAttributes()) > 0 {
			_, err = fmt.Fprintf(writer, "%slx.DOCTYPE()%s,", indent, luxeGoAttributes(t.getAttributes(), t.getNestingLevel()))
			return err
		}
		_, err = fmt.Fprintf(writer, "%slx.DOCTYPE().Html(),", indent)
	case spec.CommentType:
		commentLine := strings.Split(strings.TrimSpace(t.getTextContent()), "\n")
		if len(commentLine) > 1 {
			_, err = fmt.Fprintf(writer, "%slx.Comment(`", indent)
			comment := ""
			for _, line := range commentLine {
				comment += fmt.Sprintf("\n%s%s", indent, strings.TrimSpace(line))
			}
			_, err = fmt.Fprintf(writer, "%s`),", strings.TrimSpace(comment))
			return err
		}
		_, err = fmt.Fprintf(writer, "%slx.Comment(\"%s\"),", indent, t.getTextContent())
	case spec.TextContentType:
		textLine := strings.Split(strings.TrimSpace(t.getTextContent()), "\n")
		if len(textLine) > 1 {
			text := ""
			for _, line := range textLine {
				text += fmt.Sprintf("\n%s%s", indent, strings.TrimSpace(line))
			}
			_, err = fmt.Fprintf(writer, "`%s`,", text[1:])
			return err
		}
		_, err = fmt.Fprintf(writer, "%s\"%s\"", indent, t.getTextContent())
	case spec.SelfClosingType:
		_, err = fmt.Fprintf(writer, "%slx.%s()%s,", indent, capitalizeFirst(t.getName()), luxeGoAttributes(t.getAttributes(), t.getNestingLevel()))
	default:
		attr := luxeGoAttributes(t.getAttributes(), t.getNestingLevel())
		openTag := fmt.Sprintf("%slx.%s{}", indent, capitalizeFirst(t.getName()))
		_, err = fmt.Fprintf(writer, openTag)

		var result bytes.Buffer
		for _, child := range t.getChildren() {
			if len(t.getChildren()) == 1 && child.getName() == "TextContent" {
				if len(openTag)+len(child.getTextContent()) < 110 {
					err = child.render(ctx, &result)
					_, err = fmt.Fprintf(writer, "%s.Children(%s),", attr, strings.TrimSpace(result.String()))
					return err
				}
			}
			result.WriteString("\n")
			err = child.render(ctx, &result)
			if err != nil {
				return err
			}
		}
		if result.Len() == 0 {
			_, err = fmt.Fprintf(writer, "%s,", attr)
			return err
		}
		_, err = fmt.Fprintf(writer, "%s.Children(%s),", attr, result.String())
	}
	return err
}

func luxeGoAttributes(attr []*Attribute, offset int) string {
	if attr == nil {
		return ""
	}

	var (
		result      string
		line        string
		lineCounter = 0
	)
	indent := strings.Repeat("\t", offset+1)

	for _, val := range attr {
		attrStr := val.renderLuxeGoAttribute()
		if len(line)+len(attr)+offset*4 >= 110 {
			if len(line) < len(attr) {
				result += attrStr + ".\n" + indent
				lineCounter++
				continue
			}
			result += line + "\n" + indent
			line = attrStr
			lineCounter++
			continue
		}
		line += attrStr
	}
	if lineCounter > 0 {
		result += line + "\n" + strings.Repeat("\t", offset)
	} else {
		result += line
	}
	return result
}

// TODO: not even started
func luxeGoTempl(ctx context.Context, writer io.Writer, t *tag) (err error) {
	//	indent := strings.Repeat("\t", t.getNestingLevel())
	//	switch t.getTagType() {
	//	case DocumentType:
	//		var page bytes.Buffer
	//		for _, child := range t.getChildren() {
	//			err = child.render(ctx, &page)
	//			if err != nil {
	//				return err
	//			}
	//			page.WriteString("\n" + indent)
	//		}
	//		_, err = fmt.Fprintf(writer, "%s", page.String())
	//	case ComponentType:
	//		var component bytes.Buffer
	//		for _, child := range t.getChildren() {
	//			err = child.render(ctx, &component)
	//			if err != nil {
	//				return err
	//			}
	//			component.WriteString("\n" + indent)
	//		}
	//		_, err = fmt.Fprintf(writer, "%s", component.String())
	//	case DoctypeType:
	//		attr := t.getAttributes()
	//		if len(attr) > 1 {
	//			switch attr[1].Name {
	//			case "SYSTEM":
	//				_, err = fmt.Fprintf(writer, "%s<!DOCTYPE %s SYSTEM \"%s\">", indent, attr[0].Name, attr[1].Value)
	//			default:
	//				_, err = fmt.Fprintf(writer, "%s<!DOCTYPE %s PUBLIC \"%s\" \"%s\">", indent, attr[0].Name, attr[1].Value, attr[2].Value)
	//			}
	//			return err
	//		}
	//		if len(attr) == 1 {
	//			_, err = fmt.Fprintf(writer, "%s<!DOCTYPE %s>", indent, attr[0].Name)
	//			return err
	//		}
	//		_, err = fmt.Fprintf(writer, "%s<!DOCTYPE html>", indent)
	//	case CommentType:
	//		commentLine := strings.Split(strings.TrimSpace(t.getTextContent()), "\n")
	//		if len(commentLine) > 1 {
	//			comment := fmt.Sprintf("%s<!--", indent)
	//			for _, line := range commentLine {
	//				comment += fmt.Sprintf("\n%s%s", indent, strings.TrimSpace(line))
	//			}
	//			_, err = fmt.Fprintf(writer, "%s\n%s-->", comment, indent)
	//			return err
	//		}
	//		_, err = fmt.Fprintf(writer, "%s<!--%s-->", indent, t.getTextContent())
	//	case TextContentType:
	//		textLine := strings.Split(strings.TrimSpace(t.getTextContent()), "\n")
	//		if len(textLine) > 1 {
	//			text := ""
	//			for _, line := range textLine {
	//				text += fmt.Sprintf("\n%s%s", indent, strings.TrimSpace(line))
	//			}
	//			_, err = fmt.Fprintf(writer, "%s", text[1:])
	//			return err
	//		}
	//		_, err = fmt.Fprintf(writer, "%s%s", indent, t.getTextContent())
	//	case SelfClosingType:
	//		if t.getNamespace() == spec.HTML {
	//			_, err = fmt.Fprintf(writer, "%s<%s%s>", indent, t.getName(), defaultAttributes(t.getAttributes(), t.getNestingLevel()))
	//			return err
	//		}
	//		_, err = fmt.Fprintf(writer, "%s<%s%s />", indent, t.getName(), defaultAttributes(t.getAttributes(), t.getNestingLevel()))
	//	default:
	//		openTag := fmt.Sprintf("%s<%s%s>", indent, t.getName(), defaultAttributes(t.getAttributes(), t.getNestingLevel()))
	//		_, err = fmt.Fprintf(writer, openTag)
	//
	//		var result bytes.Buffer
	//		for _, child := range t.getChildren() {
	//			if len(t.getChildren()) == 1 && child.getName() == "TextContent" {
	//				if len(openTag)+len(child.getTextContent()) < 110 {
	//					err = child.render(ctx, &result)
	//					_, err = fmt.Fprintf(writer, "%s</%s>", strings.TrimSpace(result.String()), t.getName())
	//					return err
	//				}
	//			}
	//			result.WriteString("\n")
	//			err = child.render(ctx, &result)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//		if result.Len() == 0 {
	//			_, err = fmt.Fprintf(writer, "</%s>", t.getName())
	//			return err
	//		}
	//		_, err = fmt.Fprintf(writer, "%s\n%s</%s>", result.String(), indent, t.getName())
	//	}
	return err
}

func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func removeAndCamelCase(atrName string, chars ...string) string {
	name := atrName
	for _, char := range chars {
		if strings.Contains(atrName, char) {
			index := strings.Index(atrName, char)
			name = strings.Replace(name, char, "", -1)
			name = name[:index] + strings.ToUpper(string(name[index])) + name[index+1:]
		}
	}
	return name
}
