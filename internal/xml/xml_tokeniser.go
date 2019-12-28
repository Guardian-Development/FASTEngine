package xml

import (
	"encoding/xml"
	"errors"
	"io"
)

// Tag provides a typed version of an XML document
type Tag struct {
	Type       string
	Attributes map[string]string
	NestedTags []Tag
}

// LoadTagsFrom takes an XML decoder and reads the XML document into an Tag type
func LoadTagsFrom(decoder *xml.Decoder) (Tag, error) {
	rootTag := Tag{}
	err := populateTag(decoder, &rootTag)
	return rootTag, err
}

func populateTag(decoder *xml.Decoder, parentTag *Tag) error {
	for {
		token, err := decoder.Token()

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			err := processNewElement(decoder, parentTag, t)
			if err != nil {
				return err
			}
		case xml.EndElement:
			return nil
		case xml.CharData:
		case xml.Comment:
		case xml.ProcInst:
		default:
			return errors.New("Unable to parse templates XML, there was an unexpected element type in the file")
		}
	}
}

func processNewElement(decoder *xml.Decoder, parentTag *Tag, element xml.StartElement) error {
	// if we have already processed this element, this StartElement is a sub element of the parentTag
	if parentTag.Type != "" {
		childTag := Tag{
			Type:       element.Name.Local,
			Attributes: parseAttributes(element.Attr),
		}
		err := populateTag(decoder, &childTag)

		if err != nil {
			return err
		}

		parentTag.NestedTags = append(parentTag.NestedTags, childTag)
	} else {
		parentTag.Type = element.Name.Local
		parentTag.Attributes = parseAttributes(element.Attr)
	}

	return nil
}

func parseAttributes(attributes []xml.Attr) map[string]string {
	xmlAttributes := make(map[string]string)
	for _, attribute := range attributes {
		xmlAttributes[attribute.Name.Local] = attribute.Value
	}

	return xmlAttributes
}
