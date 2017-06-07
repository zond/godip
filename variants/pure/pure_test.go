package pure

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	cla "github.com/zond/godip/variants/classical/common"
	dip "github.com/zond/godip/common"
	vrt "github.com/zond/godip/variants/common"
)

var temp_directory = "test_output_maps"

func init() {
	_ = os.Mkdir(temp_directory, 0755)
}

func createTempPath(name string) string {
	return filepath.Join(temp_directory, name)
}

func createFile(name string, bin_data []byte) {
	path := createTempPath(name)
	ioutil.WriteFile(path, bin_data, 0644)
}

func openFile(name string) *os.File {
	file, err := os.OpenFile(createTempPath(name), os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func provincesContain(provinces []dip.Province, needle string) bool {
	for _, sc := range provinces {
		if needle == string(sc) {
			return true
		}
	}
	return false
}

func variantContains(variant vrt.Variant, needle string) bool {
	for _, nation := range variant.Nations {
		if provincesContain(variant.Graph().SCs(nation), needle) {
			return true
		}
	}
	return provincesContain(variant.Graph().SCs(cla.Neutral), needle)
}

func findAttr(attrs []xml.Attr, name string) *xml.Attr {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return &attr
		}
	}
	return nil
}

func setAttr(attrs []xml.Attr, name string, value string) {
	for i, attr := range attrs {
		if attr.Name.Local == name {
			attrs[i].Value = value
			return
		}
	}
	attrs = append(attrs, xml.Attr{xml.Name{"http://www.w3.org/2000/svg", name}, value})
}

func removeAttr(attrs []xml.Attr, name string) {
	for i, attr := range attrs {
		if attr.Name.Local == name {
			copy(attrs[i:], attrs[i+1:])
			attrs[len(attrs)-1] = xml.Attr{xml.Name{"",""}, ""}
			attrs = attrs[:len(attrs)-1]
			return
		}
	}
}

// Create svg files which can be inspected manually to check the binary map data is correct.
func TestDrawMaps(t *testing.T) {
	variant := PureVariant
	
	// Output what the empty map looks like
	b, err := variant.SVGMap()
	if err != nil {
		panic(err)
	}
	createFile("empty.svg", b)
	
	// Fill each SC red and output it
	xmlFile := bytes.NewReader(b)
	decoder := xml.NewDecoder(xmlFile)
	encoder := xml.NewEncoder(openFile("supply_centers.svg"))
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch startElement := token.(type) {
			case xml.StartElement:
			var idAttr = findAttr(startElement.Attr, "id")
			if idAttr != nil {
				id := idAttr.Value
				// Ensure the provinces layer is visible.
				if id == "provinces" {
					removeAttr(startElement.Attr, "style")
				} else if variantContains(variant, id) {
					// Colour each supply center province red.
					var styleAttr = findAttr(startElement.Attr, "style")
					if styleAttr != nil {
						var style = styleAttr.Value
						var re = regexp.MustCompile(`fill:[^;]+`)
						newStyle := re.ReplaceAllString(style, `fill:#ff0000`)
						styleAttr.Value = newStyle
						setAttr(startElement.Attr, "style", newStyle)
					}
				}
			}
			// Remove the duplicate xmlns attribute from the root element.
			// See https://github.com/golang/go/issues/13400 for the ongoing issues with this.
			if startElement.Name.Local == "svg" {
				removeAttr(startElement.Attr, "xmlns")
			}
		}
		encoder.EncodeToken(token)
	}
	encoder.Flush()
}