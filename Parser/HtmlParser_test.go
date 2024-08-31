package parser

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Test cases for ParseBody function
func TestParseBody(t *testing.T) {
	rawHTML := `<!DOCTYPE html><html><head><meta charset="utf-8"><title>Test Page</title></head><body><h1>Hello, Go!</h1><input type="text" name="text1"><div class="nevim asdfj" popover>Hello</div><p>This is an example paragraph.</p></body></html>`

	reader := strings.NewReader(rawHTML)
	htmlStructure := ParseBody(reader)
	var w strings.Builder
	err := htmlStructure.Build(context.Background(), &w)
	result := w.String()

	// // Define expected result
	// var expected = html.WebContent(
	// 	html.DOCTYPE().Html(),
	// 	html.Html(
	// 		html.Head(
	// 			html.Meta().Charset("utf-8"),
	// 			html.Title("Test Page"),
	// 		),
	// 		html.Body(
	// 			html.H1("Hello, Go!"),
	// 			html.Input().Type("text").name("text1"),
	// 			html.Div("Hello").Class("nevim asdfj").Popover(),
	// 			html.P("This is an example paragraph."),
	// 		),
	// 	),
	// )

	// var b strings.Builder
	// err = expected.Build(context.Background(), &b)
	if err != nil {
		return
	}
	// exp := fmt.Sprintf("%s\n", b.String())

	// Populate expected result according to your logic
	if result != rawHTML {
		CompareStrings(result, rawHTML)
		t.Fatalf("Structures are not equal")
	}
}

// Todo: Comment tag doesn't work
func TestNestedDiv(t *testing.T) {
	rawHTML := `<div class="nevim asdfj" popover><!--Hello -->Hello<div class="nevim asdfj" popover>dalsi kontent<input class="afdsfa"></div></div>`

	reader := strings.NewReader(rawHTML)
	htmlStructure := ParseBody(reader)
	var w strings.Builder

	err := htmlStructure.Build(context.Background(), &w)

	result := w.String()

	// Define expected result
	// var expected = html.Div(
	// 	html.Comment("Hello"),
	// 	"Hello",
	// 	html.Div(
	// 		"dalsi kontent",
	// 		html.Input().Class("afdsfa"),
	// 	).Class("nevim asdfj").Popover(),
	// ).Class("nevim asdfj").Popover()

	// var b strings.Builder
	// err = expected.Build(context.Background(), &b)
	if err != nil {
		return
	}
	// exp := fmt.Sprintf("%s\n", b.String())

	// Populate expected result according to your logic
	if result != rawHTML {
		CompareStrings(result, rawHTML)
		t.Fatalf("Structures are not equal")
	}
}

func TestHtmlFile(t *testing.T) {
	// Read HTML content from file
	htmlFile := "./TestingAssets/index.html" // Ensure this path is valid
	rawHTML, err := os.ReadFile(htmlFile)
	if err != nil {
		t.Fatalf("Failed to read HTML file: %v", err)
	}

	rawHTMLStr := string(rawHTML)
	reader := strings.NewReader(rawHTMLStr)
	htmlStructure := ParseBody(reader)
	var w strings.Builder
	err = htmlStructure.Build(context.Background(), &w)
	if err != nil {
		t.Fatalf("Build function failed: %v", err)
	}

	// Build expected output format
	result := w.String()

	// Compare result with expected output
	if result != rawHTMLStr {
		CompareStrings(result, rawHTMLStr)
		t.Fatalf("Structures are not equal")
	}
}

// CompareStrings highlights the differences between two strings and prints the results in the terminal
func CompareStrings(s1, s2 string) {
	len1, len2 := len(s1), len(s2)
	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	var sb1, sb2 strings.Builder

	for i := 0; i < maxLen; i++ {
		var char1, char2 byte
		if i < len1 {
			char1 = s1[i]
		} else {
			char1 = ' ' // Add space if second string is longer
		}
		if i < len2 {
			char2 = s2[i]
		} else {
			char2 = ' ' // Add space if first string is longer
		}

		if char1 == char2 {
			sb1.WriteByte(char1)
			sb2.WriteByte(char2)
		} else {
			sb1.WriteString(fmt.Sprintf("\033[31m%c\033[0m", char1)) // Print in red if different
			sb2.WriteString(fmt.Sprintf("\033[32m%c\033[0m", char2)) // Print in green if different
		}
	}

	fmt.Println("String 1:", sb1.String())
	fmt.Println("String 2:", sb2.String())
}
