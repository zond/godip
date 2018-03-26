package variants

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
	vrt "github.com/zond/godip/variants/common"
)

var temp_directory = "test_output_maps"
var svg = "http://www.w3.org/2000/svg"

func init() {
	os.Mkdir(temp_directory, 0755)
	for _, variant := range OrderedVariants {
		os.Mkdir(filepath.Join(temp_directory, variant.Name), 0755)
	}
}

func createTempPath(variant string, name string) string {
	return filepath.Join(temp_directory, variant, name)
}

func createFile(variant string, name string, bin_data []byte) {
	path := createTempPath(variant, name)
	ioutil.WriteFile(path, bin_data, 0644)
}

func openFile(variant string, name string) *os.File {
	file, err := os.OpenFile(createTempPath(variant, name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func provincesContain(provinces []dip.Province, needle string) bool {
	for _, province := range provinces {
		if needle == string(province) {
			return true
		}
	}
	return false
}

func variantContainsSC(variant vrt.Variant, needle string) bool {
	for _, nation := range variant.Nations {
		if provincesContain(variant.Graph().SCs(nation), needle) {
			return true
		}
	}
	return provincesContain(variant.Graph().SCs(cla.Neutral), needle)
}

func variantContainsProvince(variant vrt.Variant, needle string) bool {
	return provincesContain(variant.Graph().Provinces(), needle)
}

func findAttr(attrs []xml.Attr, name string) *xml.Attr {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return &attr
		}
	}
	return nil
}

func setAttr(attrs []xml.Attr, name string, value string) []xml.Attr {
	for i, attr := range attrs {
		if attr.Name.Local == name {
			attrs[i].Value = value
			return attrs
		}
	}
	return append(attrs, xml.Attr{xml.Name{"", name}, value})
}

// Return a slice with the first found occurance of an attribute removed.
// If no matching attribute is found then the whole slice is returned.
func removeAttr(attrs []xml.Attr, name string) []xml.Attr {
	for i, attr := range attrs {
		if attr.Name.Local == name {
			copy(attrs[i:], attrs[i+1:])
			attrs[len(attrs)-1] = xml.Attr{xml.Name{"", ""}, ""}
			return attrs[:len(attrs)-1]
		}
	}
	return attrs
}

type vector struct {
	dx float64
	dy float64
}

func (v vector) length() float64 {
	return math.Sqrt(v.dx*v.dx + v.dy*v.dy)
}

func (a vector) add(b vector) vector {
	return vector{a.dx + b.dx, a.dy + b.dy}
}

func (v vector) mul(m float64) vector {
	return vector{v.dx * m, v.dy * m}
}

func (v vector) div(d float64) vector {
	return vector{v.dx / d, v.dy / d}
}

func (v vector) dir() vector {
	return v.div(v.length())
}

func (v vector) orth() vector {
	return vector{-v.dy, v.dx}.dir()
}

type coordinates struct {
	x float64
	y float64
}

func (c coordinates) add(v vector) coordinates {
	return coordinates{c.x + v.dx, c.y + v.dy}
}

func (c coordinates) sub(v vector) coordinates {
	return coordinates{c.x - v.dx, c.y - v.dy}
}

func (a coordinates) to(b coordinates) vector {
	return vector{b.x - a.x, b.y - a.y}
}

func toStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func addArrow(encoder *xml.Encoder, startProvince dip.Province, endProvince dip.Province, provinceCenters map[string]coordinates) {
	start := provinceCenters[string(startProvince)]
	end := provinceCenters[string(endProvince)]
	middle := coordinates{(start.x + end.x) / 2, (start.y + end.y) / 2}

	boundF := 3.0
	headF1 := boundF * 3
	headF2 := boundF * 6
	spacer := boundF * 2
	part1 := start.to(middle)
	part2 := middle.to(end)
	start0 := start.add(part1.dir().mul(spacer)).add(part1.orth().mul(boundF))
	start1 := start.add(part1.dir().mul(spacer)).sub(part1.orth().mul(boundF))
	sumOrth := part1.orth().add(part2.orth())
	avgOrth := sumOrth.div(sumOrth.length())
	control0 := middle.add(avgOrth.mul(boundF))
	control1 := middle.sub(avgOrth.mul(boundF))
	end0 := end.sub(part2.dir().mul(spacer + headF2)).add(part2.orth().mul(boundF))
	end1 := end.sub(part2.dir().mul(spacer + headF2)).sub(part2.orth().mul(boundF))
	end3 := end.sub(part2.dir().mul(spacer))
	head0 := end0.add(part2.orth().mul(headF1))
	head1 := end1.sub(part2.orth().mul(headF1))

	path := xml.StartElement{xml.Name{svg, "path"}, []xml.Attr{}}
	path.Attr = setAttr(path.Attr, "style", "fill:#ff0000;stroke:#000000;stroke-width:0.5;stroke-miterlimit:4;stroke-opacity:1.0;fill-opacity:0.7;")
	var d = "M " + toStr(start0.x) + "," + toStr(start0.y)
	d += " C " + toStr(control0.x) + "," + toStr(control0.y) + "," + toStr(control0.x) + "," + toStr(control0.y) + "," + toStr(end0.x) + "," + toStr(end0.y)
	d += " L " + toStr(head0.x) + "," + toStr(head0.y)
	d += " L " + toStr(end3.x) + "," + toStr(end3.y)
	d += " L " + toStr(head1.x) + "," + toStr(head1.y)
	d += " L " + toStr(end1.x) + "," + toStr(end1.y)
	d += " C " + toStr(control1.x) + "," + toStr(control1.y) + "," + toStr(control1.x) + "," + toStr(control1.y) + "," + toStr(start1.x) + "," + toStr(start1.y)
	d += " z"
	path.Attr = setAttr(path.Attr, "d", d)
	encoder.EncodeToken(path)
	encoder.EncodeToken(xml.EndElement{xml.Name{svg, "path"}})
}

// Create svg files which can be inspected manually to check the binary map data is correct.
func TestDrawMaps(t *testing.T) {
	if os.Getenv("DRAW_MAPS") != "true" {
		fmt.Println("Skipping test to draw debug maps. Please use the environment variable DRAW_MAPS=true to enable.")
		return
	}
	for _, variant := range OrderedVariants {
		// Output what the empty map looks like
		b, err := variant.SVGMap()
		if err != nil {
			panic(err)
		}
		createFile(variant.Name, "empty.svg", b)

		// Fill each SC red and output it
		xmlFile := bytes.NewReader(b)
		decoder := xml.NewDecoder(xmlFile)
		encoder := xml.NewEncoder(openFile(variant.Name, "supply_centers.svg"))
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
						startElement.Attr = removeAttr(startElement.Attr, "style")
					} else if variantContainsSC(variant, id) {
						// Colour each supply center province red.
						var styleAttr = findAttr(startElement.Attr, "style")
						if styleAttr != nil {
							var style = styleAttr.Value
							var re = regexp.MustCompile(`fill:[^;]+`)
							newStyle := re.ReplaceAllString(style, `fill:#ff0000`)
							styleAttr.Value = newStyle
							startElement.Attr = setAttr(startElement.Attr, "style", newStyle)
						}
					}
				}
				// Remove the duplicate xmlns attribute from the root element.
				// See https://github.com/golang/go/issues/13400 for the ongoing issues with this.
				if startElement.Name.Local == "svg" {
					startElement.Attr = removeAttr(startElement.Attr, "xmlns")
				}
			}
			encoder.EncodeToken(token)
		}
		encoder.Flush()

		// Mark each province selectable and output it
		xmlFile = bytes.NewReader(b)
		decoder = xml.NewDecoder(xmlFile)
		encoder = xml.NewEncoder(openFile(variant.Name, "selectable_provinces.svg"))
		highlights := []xml.StartElement{}
		for {
			outputHighlights := false
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch startElement := token.(type) {
			case xml.StartElement:
				idAttr := findAttr(startElement.Attr, "id")
				if idAttr != nil {
					id := idAttr.Value
					// Ensure the provinces layer is visible.
					if id == "provinces" {
						startElement.Attr = removeAttr(startElement.Attr, "style")
					} else if variantContainsProvince(variant, id) {
						// Duplicate each province region.
						styleAttr := findAttr(startElement.Attr, "style")
						if styleAttr != nil {
							style := styleAttr.Value
							re := regexp.MustCompile(`fill:[^;]+`)
							newStyle := re.ReplaceAllString(style, `fill:#ffffff`)
							styleAttr.Value = newStyle
							startElement.Attr = setAttr(startElement.Attr, "style", newStyle)
						}
						highlights = append(highlights, startElement)
					} else if id == "highlights" {
						outputHighlights = true
					}
				}
				// Remove the duplicate xmlns attribute from the root element.
				// See https://github.com/golang/go/issues/13400 for the ongoing issues with this.
				if startElement.Name.Local == "svg" {
					startElement.Attr = removeAttr(startElement.Attr, "xmlns")
				}
			}
			encoder.EncodeToken(token)
			if outputHighlights {
				for _, highlight := range highlights {
					id := findAttr(highlight.Attr, "id").Value
					highlight.Attr = setAttr(highlight.Attr, "id", id+"_highlight")
					highlight.Attr = setAttr(highlight.Attr, "style", "fill:url(#stripes)")
					highlight.Attr = setAttr(highlight.Attr, "fill-opacity", "1")
					highlight.Attr = setAttr(highlight.Attr, "stroke", "none")
					encoder.EncodeToken(highlight)
					encoder.EncodeToken(xml.EndElement{highlight.Name})
				}
			}
		}
		encoder.Flush()

		// Find all the types of edge that exist.
		edgeTypes := make(map[dip.Flag]bool)
		for _, start := range variant.Graph().Provinces() {
			for _, flags := range variant.Graph().Edges(start) {
				for flag, b := range flags {
					if b {
						edgeTypes[flag] = true
					}
				}
			}
		}

		// Draw arrows between connected provinces
		provinceCenters := make(map[string]coordinates)
		for edgeType := range edgeTypes {
			xmlFile = bytes.NewReader(b)
			decoder = xml.NewDecoder(xmlFile)
			encoder = xml.NewEncoder(openFile(variant.Name, "orders_"+string(edgeType)+".svg"))
			for {
				outputOrders := false
				token, _ := decoder.Token()
				if token == nil {
					break
				}
				switch startElement := token.(type) {
				case xml.StartElement:
					var idAttr = findAttr(startElement.Attr, "id")
					if idAttr != nil {
						id := idAttr.Value
						// Find the location of each province.
						if strings.HasSuffix(id, "Center") {
							province := strings.Replace(id, "Center", "", 1)
							if variantContainsProvince(variant, province) {
								// Extract the coordinates from the d attribute.
								var dAttr = findAttr(startElement.Attr, "d")
								if dAttr != nil {
									var d = dAttr.Value
									var re = regexp.MustCompile(`^m\s+([\d-.]+),([\d-.]+)\s+`)
									result := re.FindStringSubmatch(d)
									x, errX := strconv.ParseFloat(result[1], 64)
									y, errY := strconv.ParseFloat(result[2], 64)
									if errX != nil || errY != nil {
										panic("Error extracting center of province")
									}
									provinceCenters[province] = coordinates{x, y}
								}
							}
						} else if id == "orders" {
							outputOrders = true
						}
					}
					// Remove the duplicate xmlns attribute from the root element.
					// See https://github.com/golang/go/issues/13400 for the ongoing issues with this.
					if startElement.Name.Local == "svg" {
						startElement.Attr = removeAttr(startElement.Attr, "xmlns")
					}
				}
				encoder.EncodeToken(token)
				if outputOrders {
					// TODO Draw arrows for all orders here.
					for _, start := range variant.Graph().Provinces() {
						for end, flags := range variant.Graph().Edges(start) {
							if flags[edgeType] {
								addArrow(encoder, start, end, provinceCenters)
							}
						}
					}
				}
			}
			encoder.Flush()
		}

		// Find all the types of province that exist.
		provinceTypes := make(map[dip.Flag]bool)
		for _, province := range variant.Graph().Provinces() {
			for flag, b := range variant.Graph().Flags(province) {
				if b {
					provinceTypes[flag] = true
				}
			}
		}

		// Draw each type of unit in each type of province
		for provinceType, _ := range provinceTypes {
			for _, unitType := range variant.UnitTypes {
				xmlFile = bytes.NewReader(b)
				decoder = xml.NewDecoder(xmlFile)
				encoder = xml.NewEncoder(openFile(variant.Name, "units_"+string(unitType)+"_"+string(provinceType)+".svg"))
				for {
					outputUnits := false
					token, _ := decoder.Token()
					if token == nil {
						break
					}
					switch startElement := token.(type) {
					case xml.StartElement:
						idAttr := findAttr(startElement.Attr, "id")
						if idAttr != nil {
							id := idAttr.Value
							// Ensure the provinces layer is visible.
							if id == "provinces" {
								startElement.Attr = removeAttr(startElement.Attr, "style")
							} else if variantContainsProvince(variant, id) {
								// Duplicate each province region.
								styleAttr := findAttr(startElement.Attr, "style")
								if styleAttr != nil {
									style := styleAttr.Value
									re := regexp.MustCompile(`fill:[^;]+`)
									newStyle := re.ReplaceAllString(style, `fill:#ffffff`)
									styleAttr.Value = newStyle
									startElement.Attr = setAttr(startElement.Attr, "style", newStyle)
								}
								highlights = append(highlights, startElement)
							} else if id == "units" {
								outputUnits = true
							}
						}
						// Remove the duplicate xmlns attribute from the root element.
						// See https://github.com/golang/go/issues/13400 for the ongoing issues with this.
						if startElement.Name.Local == "svg" {
							startElement.Attr = removeAttr(startElement.Attr, "xmlns")
						}
					}
					encoder.EncodeToken(token)
					if outputUnits {
						unit, err := variant.SVGUnits[unitType]()
						if err != nil {
							panic(err)
						}
						for province, coordinates := range provinceCenters {
							if variant.Graph().Flags(dip.Province(province))[provinceType] {
								unitFile := bytes.NewReader(unit)
								unitDecoder := xml.NewDecoder(unitFile)
								for {
									unitToken, _ := unitDecoder.Token()
									if unitToken == nil {
										break
									}
									switch startElement := unitToken.(type) {
									case xml.StartElement:
										idAttr := findAttr(startElement.Attr, "id")
										if idAttr != nil {
											id := idAttr.Value
											if id == "hull" || id == "body" || id == "shadow" {
												offset := vector{0.0, 0.0}
												if id == "hull" {
													offset = vector{-65.0, -15.0}
												} else if id == "body" {
													offset = vector{-40.0, -5.0}
												} else if id == "shadow" {
													offset = vector{-40.0, -5.0}
												}
												xStr := toStr(coordinates.x + offset.dx)
												yStr := toStr(coordinates.y + offset.dy)
												startElement.Attr = setAttr(startElement.Attr, "transform", "translate("+xStr+","+yStr+")")
												unitToken = startElement
											}
										}
										// Remove the duplicate xmlns attribute from the root element.
										// See https://github.com/golang/go/issues/13400 for the ongoing issues with this.
										if startElement.Name.Local == "svg" {
											startElement.Attr = removeAttr(startElement.Attr, "xmlns")
											// Remove the bounds from the embedded svg.
											startElement.Attr = removeAttr(startElement.Attr, "width")
											startElement.Attr = removeAttr(startElement.Attr, "height")
										}
									}
									encoder.EncodeToken(unitToken)
								}
							}
						}
					}
				}
				encoder.Flush()
			}
		}

		// Highlight home centers for each nation
		for _, nation := range variant.Nations {
			xmlFile := bytes.NewReader(b)
			decoder := xml.NewDecoder(xmlFile)
			encoder := xml.NewEncoder(openFile(variant.Name, "home_centers_"+string(nation)+".svg"))
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
							startElement.Attr = removeAttr(startElement.Attr, "style")
						} else if provincesContain(variant.Graph().SCs(nation), id) {
							// Colour each home center province red.
							var styleAttr = findAttr(startElement.Attr, "style")
							if styleAttr != nil {
								var style = styleAttr.Value
								var re = regexp.MustCompile(`fill:[^;]+`)
								newStyle := re.ReplaceAllString(style, `fill:#ff0000`)
								styleAttr.Value = newStyle
								startElement.Attr = setAttr(startElement.Attr, "style", newStyle)
							}
						}
					}
					// Remove the duplicate xmlns attribute from the root element.
					// See https://github.com/golang/go/issues/13400 for the ongoing issues with this.
					if startElement.Name.Local == "svg" {
						startElement.Attr = removeAttr(startElement.Attr, "xmlns")
					}
				}
				encoder.EncodeToken(token)
			}
			encoder.Flush()
		}
	}
}
