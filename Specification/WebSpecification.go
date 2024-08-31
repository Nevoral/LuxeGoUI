package specification

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"slices"
	"strings"
	"text/template"
)

type Namespace int

const (
	HTML Namespace = iota
	SVG
	MATH
	XHTML
)

func (n Namespace) String() string {
	switch n {
	case HTML:
		return "HTML"
	case SVG:
		return "SVG"
	case MATH:
		return "MATH"
	case XHTML:
		return "XHTML"
	default:
		return "Unknown"
	}
}

type TagType int

const (
	DoctypeType TagType = iota
	SelfClosingType
	CommentType
	TextContentType
	FullTagType
	DocumentType
	ComponentType
)

func (t TagType) String() string {
	switch t {
	case DoctypeType:
		return "DoctypeType"
	case SelfClosingType:
		return "SelfClosingType"
	case CommentType:
		return "CommentType"
	case TextContentType:
		return "TextContentType"
	case FullTagType:
		return "FullTagType"
	case DocumentType:
		return "DocumentType"
	case ComponentType:
		return "ComponentType"
	default:
		return "Unknown"
	}
}

type Comment string

// BuildComment - this method will create comment for attribute or tag.
// starts after "[TagName/AttributeName] - %s", BuildComment
func (c Comment) BuildComment() string {
	parts := strings.Split(strings.TrimSpace(string(c)), "\n")

	var result string
	for index, val := range parts {
		if index == 0 {
			result += fmt.Sprintf("%s", val)
		}
		result += fmt.Sprintf("\n%s", val)
	}
	return result
}

type WebSpecification map[Namespace]*Config

func (w *WebSpecification) GetConfig(tagType Namespace) *Config {
	conf := *w
	return conf[tagType]
}

// CreateTomlFile creates a TOML file from the provided config struct
func (w *WebSpecification) CreateTomlFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	encoder := toml.NewEncoder(file)
	if err = encoder.Encode(w); err != nil {
		return fmt.Errorf("could not encode config to TOML: %v", err)
	}

	return nil
}

// LoadFromTomlFile loads the TOML data from the specified file into the WebConfig struct.
func (w *WebSpecification) LoadFromTomlFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("could not close file:", err)
		}
	}(file)

	if _, err = toml.NewDecoder(file).Decode(w); err != nil {
		return fmt.Errorf("could not decode TOML: %v", err)
	}

	return nil
}

// TODO: not ready later
// GenerateGoFile generates a Go file with the configuration data.
func (w *WebSpecification) GenerateGoFile(filename string) error {
	fn := template.FuncMap{"toLower": strings.ToLower,
		"firstCharLower": func(s string) string {
			return strings.ToLower(s[0:1])
		},
		"capFirstChar": func(s string) string {
			return strings.ToUpper(s[0:1]) + s[1:]
		},
	}
	tmpl, err := template.New("goTemplate").Funcs(fn).ParseFiles("./internal/templates/ConfigFile.gohtml")
	if err != nil {
		return fmt.Errorf("could not parse template: %v", err)
	}

	var pageBuf bytes.Buffer
	if err = tmpl.ExecuteTemplate(&pageBuf, "goWebConfigFile", w); err != nil {
		fmt.Println("Error executing header block:", err)
		return err
	}

	// fmt.Println(buf.String())
	if err = os.WriteFile(filename, pageBuf.Bytes(), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

type AttributeCategories string

const (
	SpecificAttributes                 AttributeCategories = "Specific Attributes"
	GlobalAttributes                                       = "Global Attributes"
	AriaAttributes                                         = "Aria Attributes"
	DocumentActions                                        = "Document Actions"
	WindowActions                                          = "Window Actions"
	AnimationAdditionAttributes                            = "Animation Addition Attributes"
	AnimationEventAttributes                               = "Animation Event Attributes"
	AnimationTargetElementAttributes                       = "Animation Target Element Attributes"
	AnimationAttributeTargetAttributes                     = "Animation Attribute Target Attributes"
	AnimationTimingAttributes                              = "Animation Timing Attributes"
	AnimationValueAttributes                               = "Animation Value Attributes"
	ConditionalProcessingAttributes                        = "Conditional Processing Attributes"
	CoreAttributes                                         = "Core Attributes"
	PresentationAttributes                                 = "Presentation Attributes"
	FilterPrimitiveAttributes                              = "Filter Primitive Attributes"
	TransferFunctionAttributes                             = "Transfer Function Attributes"
	GlobalEventAttributes                                  = "Global Event Attributes"
	DocumentElementEventAttributes                         = "Document Element Event Attributes"
)

func (a AttributeCategories) String() string {
	return string(a)
}

type Config struct {
	Tags                 []*TagConfig
	AttributesCategories map[AttributeCategories][]*AttributeConfig
}

func (c *Config) GetAttributeDefaultValue(name string, category AttributeCategories) string {
	atrIndex := slices.IndexFunc(c.AttributesCategories[category], func(e *AttributeConfig) bool {
		return e.Name == name
	})
	return c.AttributesCategories[category][atrIndex].InitialValue
}

func (c *Config) GetAttributeBoolean(name string, category AttributeCategories) bool {
	atrIndex := slices.IndexFunc(c.AttributesCategories[category], func(e *AttributeConfig) bool {
		return e.Name == name
	})
	return c.AttributesCategories[category][atrIndex].Boolean
}

func (c *Config) GetTagConfig(name string) (*TagConfig, error) {
	tagIndex := slices.IndexFunc(c.Tags, func(e *TagConfig) bool {
		return strings.ToLower(e.Name) == name
	})
	if tagIndex < 0 {
		msg := fmt.Errorf("Error: in specification isn't any tag called %s.", name)
		err := logError(msg)
		if err != nil {
			return nil, err
		}
		return nil, msg
	}
	return c.Tags[tagIndex], nil
}

func (c *Config) CheckValueValidity(name, value string, category AttributeCategories) bool {
	atrIndex := slices.IndexFunc(c.AttributesCategories[category], func(e *AttributeConfig) bool {
		return e.Name == name
	})
	if len(c.AttributesCategories[category][atrIndex].SupportedValues) == 0 {
		return true
	}
	if _, ok := c.AttributesCategories[category][atrIndex].SupportedValues[value]; !ok {
		return false
	}
	return true
}

type TagConfig struct {
	Name                       string
	TagType                    TagType
	Comment                    Comment
	DocumentationURL           string
	AttributesCategorySupports map[AttributeCategories][]string
	SpecificAttributes         []*AttributeConfig
	SupportedChildrenTags      []string
}

type AttributeConfig struct {
	Name             string
	Boolean          bool
	Comment          Comment
	DocumentationURL string
	InitialValue     string             // if exist value will be there if not it will be empty string ""
	SupportedValues  map[string]Comment // if map has 0 len its
}
