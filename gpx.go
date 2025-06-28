package gpx

import (
	"io"
	"os"
	"strings"

	xml "github.com/Zauberstuhl/go-xml"
)

// ParseFile takes a file and parses it
func ParseFile(fileName string) (*GPX, error) {
	g := GPX{}

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return &g, err
	}

	err = Parse(bytes, &g)
	if err != nil {
		return &g, err
	}
	return &g, nil
}

// Parse bytes of xml
func Parse(bytes []byte, g *GPX) error {
	err := xml.Unmarshal(bytes, g)
	if err != nil {
		return err
	}
	return nil
}

// Write GPX file
func Write(g *GPX, fileName string) error {
	output, err := xml.MarshalIndent(g, "", "    ")
	if err != nil {
		return err
	}

	fileData := append([]byte(xml.Header), output...)
	var path string

	if strings.HasSuffix(fileName, "gpx") {
		path = fileName
	} else {
		path = fileName + ".gpx"
	}

	err = os.WriteFile(path, fileData, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Comments from http://www.topografix.com/GPX/1/1/
// Extensions from the following
// https://www8.garmin.com/xmlschemas/GpxExtensions/v3/GpxExtensionsv3.xsd
// https://www8.garmin.com/xmlschemas/WaypointExtensionv1.xsd
// https://www8.garmin.com/xmlschemas/TrackPointExtensionv1.xsd

// GPX is the root element
type GPX struct {
	XMLName   xml.Name   `xml:"gpx"`
	Version   string     `xml:"version,attr"`
	Creator   string     `xml:"creator,attr"`
	Metadata  Metadata   `xml:"metadata,omitempty"`
	Waypoints []WayPoint `xml:"wpt,omitempty"`
	Routes    []Route    `xml:"rte,omitempty"`
	Tracks    []Track    `xml:"trk"`
}

// Metadata has information about the GPX file
type Metadata struct {
	XMLName     xml.Name    `xml:"metadata"`
	Name        string      `xml:"name,omitempty"`
	Description string      `xml:"desc,omitempty"`
	Author      *Person     `xml:"author,omitempty"`
	Copyright   *Copyright  `xml:"copyright,omitempty"`
	Links       []Link      `xml:"link,omitempty"`
	Timestamp   string      `xml:"time,omitempty"`
	Keywords    string      `xml:"keywords,omitempty"`
	Bounds      *Bounds     `xml:"bounds,omitempty"`
	Extensions  *Extensions `xml:"extensions,omitempty"`
}

// WayPoint is a point of interest, or named feature on a map.
type WayPoint struct {
	XMLName                       xml.Name           `xml:"wpt"`
	Latitude                      Latitude           `xml:"lat,attr"`
	Longitude                     Longitude          `xml:"lon,attr"`
	Elevation                     float64            `xml:"ele,omitempty"`
	Timestamp                     string             `xml:"time,omitempty"`
	MagneticVariation             Degrees            `xml:"magvar,omitempty"`
	GeoIDHeight                   float64            `xml:"geoidheight,omitempty"`
	Name                          string             `xml:"name,omitempty"`
	Comment                       string             `xml:"cmt,omitempty"`
	Description                   string             `xml:"desc,omitempty"`
	Source                        string             `xml:"src,omitempty"`
	Links                         []Link             `xml:"link,omitempty"`
	Symbol                        string             `xml:"sym,omitempty"`
	Type                          string             `xml:"type,omitempty"`
	Fix                           Fix                `xml:"fix,omitempty"`
	Sat                           int                `xml:"sat,omitempty"`
	HorizontalDilutionOfPrecision float64            `xml:"hdop,omitempty"`
	VerticalDilutionOfPrecision   float64            `xml:"vdop,omitempty"`
	PositionDilutionOfPrecision   float64            `xml:"pdop,omitempty"`
	AgeOfGpsData                  float64            `xml:"ageofgpsdata,omitempty"`
	DifferentialGPSID             DGPSStation        `xml:"dgpsid,omitempty"`
	Extensions                    WayPointExtensions `xml:"extensions,omitempty"`
}

// Route is an ordered list of Waypoints representing a series of points leading to a destination.
type Route struct {
	XMLName     xml.Name        `xml:"rte"`
	Name        string          `xml:"name,omitempty"`
	Comment     string          `xml:"cmt,omitempty"`
	Description string          `xml:"desc,omitempty"`
	Source      string          `xml:"src,omitempty"`
	Links       []Link          `xml:"link"`
	Number      int             `xml:"number,omitempty"`
	Type        string          `xml:"type,omitempty"`
	Extensions  RouteExtensions `xml:"extensions,omitempty"`
	RoutePoints []RoutePoint    `xml:"rtept"`
}

// RoutePoint is a point of interest, or named feature on a map.
type RoutePoint struct {
	XMLName                       xml.Name             `xml:"rtept"`
	Latitude                      Latitude             `xml:"lat,attr"`
	Longitude                     Longitude            `xml:"lon,attr"`
	Elevation                     float64              `xml:"ele,omitempty"`
	Timestamp                     string               `xml:"time,omitempty"`
	MagneticVariation             Degrees              `xml:"magvar,omitempty"`
	GeoIDHeight                   float64              `xml:"geoidheight,omitempty"`
	Name                          string               `xml:"name,omitempty"`
	Comment                       string               `xml:"cmt,omitempty"`
	Description                   string               `xml:"desc,omitempty"`
	Source                        string               `xml:"src,omitempty"`
	Links                         []Link               `xml:"link"`
	Symbol                        string               `xml:"sym,omitempty"`
	Type                          string               `xml:"type,omitempty"`
	Fix                           Fix                  `xml:"fix,omitempty"`
	Sat                           int                  `xml:"sat,omitempty"`
	HorizontalDilutionOfPrecision float64              `xml:"hdop,omitempty"`
	VerticalDilutionOfPrecision   float64              `xml:"vdop,omitempty"`
	PositionDilutionOfPrecision   float64              `xml:"pdop,omitempty"`
	AgeOfGpsData                  float64              `xml:"ageofgpsdata,omitempty"`
	DifferentialGPSID             DGPSStation          `xml:"dgpsid,omitempty"`
	Extensions                    RoutePointExtensions `xml:"extensions,omitempty"`
}

// Track represents a track - an ordered list of points describing a path
type Track struct {
	XMLName       xml.Name         `xml:"trk"`
	Name          string           `xml:"name,omitempty"`
	Comment       string           `xml:"cmt,omitempty"`
	Description   string           `xml:"desc,omitempty"`
	Source        string           `xml:"src,omitempty"`
	Links         []Link           `xml:"link"`
	Number        int              `xml:"number,omitempty"`
	Type          string           `xml:"type,omitempty"`
	Extensions    *TrackExtensions `xml:"extensions,omitempty"`
	TrackSegments []TrackSegment   `xml:"trkseg"`
}

// TrackSegment has a list of continious span of TrackPoints
type TrackSegment struct {
	XMLName    xml.Name     `xml:"trkseg"`
	TrackPoint []TrackPoint `xml:"trkpt"`
	Extensions *Extensions  `xml:"extensions,omitempty"`
}

// TrackPoint is a point of interest, or named feature on a map.
type TrackPoint struct {
	XMLName                       xml.Name              `xml:"trkpt"`
	Latitude                      Latitude              `xml:"lat,attr"`
	Longitude                     Longitude             `xml:"lon,attr"`
	Elevation                     float64               `xml:"ele,omitempty"`
	Timestamp                     string                `xml:"time,omitempty"`
	MagneticVariation             Degrees               `xml:"magvar,omitempty"`
	GeoIDHeight                   float64               `xml:"geoidheight,omitempty"`
	Name                          string                `xml:"name,omitempty"`
	Comment                       string                `xml:"cmt,omitempty"`
	Description                   string                `xml:"desc,omitempty"`
	Source                        string                `xml:"src,omitempty"`
	Links                         []Link                `xml:"link"`
	Symbol                        string                `xml:"sym,omitempty"`
	Type                          string                `xml:"type,omitempty"`
	Fix                           Fix                   `xml:"fix,omitempty"`
	Sat                           int                   `xml:"sat,omitempty"`
	HorizontalDilutionOfPrecision float64               `xml:"hdop,omitempty"`
	VerticalDilutionOfPrecision   float64               `xml:"vdop,omitempty"`
	PositionDilutionOfPrecision   float64               `xml:"pdop,omitempty"`
	AgeOfGpsData                  float64               `xml:"ageofgpsdata,omitempty"`
	DifferentialGPSID             DGPSStation           `xml:"dgpsid,omitempty"`
	Extensions                    *TrackPointExtensions `xml:"extensions,omitempty"`
}

// Copyright has information about holder and license
type Copyright struct {
	XMLName xml.Name `xml:"copyright"`
	Author  string   `xml:"author,attr"`
	Year    int      `xml:"year,omitempty"`
	License string   `xml:"license,omitempty"`
}

// Link is for an external resource with additional information.
type Link struct {
	XMLName xml.Name `xml:"link"`
	URL     string   `xml:"href,attr,omitempty"`
	Text    string   `xml:"text,omitempty"`
	Type    string   `xml:"type,omitempty"`
}

// Email address which is broken into two parts (id and domain)
type Email struct {
	XMLName xml.Name `xml:"email"`
	ID      string   `xml:"id,attr,omitempty"`
	Domain  string   `xml:"domain,attr,omitempty"`
}

// Person is a person or an organisation
type Person struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name,omitempty"`
	Email   Email    `xml:"email,omitempty"`
	Link    Link     `xml:"link,omitempty"`
}

// Point with optional elevation and time
type Point struct {
	XMLName   xml.Name  `xml:"pt"`
	Latitude  Latitude  `xml:"lat,attr"`
	Longitude Longitude `xml:"lon,attr"`
	Elevation float64   `xml:"ele,omitempty"`
	Timestamp string    `xml:"time,omitempty"`
}

// PointSegment is a sequence of Points
type PointSegment struct {
	XMLName xml.Name `xml:"ptseg"`
	Points  []Point  `xml:"pt"`
}

// Bounds are two latitude longitude pairs defining the extent of an element.
type Bounds struct {
	XMLName          xml.Name  `xml:"bounds"`
	MinimumLatitude  Latitude  `xml:"minlat,attr"`
	MaximumLatitude  Latitude  `xml:"maxlat,attr"`
	MinimumLongitude Longitude `xml:"minlon,attr"`
	MaximumLongitude Longitude `xml:"maxlon,attr"`
}

// Latitude is the latitude of the point. Decimal degrees, WGS84 datum. The value varies between -90.0 to 90.0
type Latitude float64

// Longitude is the longitude of the point. Decimal degrees, WGS84 datum. The value varies between -180.0 and 180.0
type Longitude float64

// Degrees is used for bearing, heading, course. Units are decimal degrees, true (not magnetic). The value varies between 0.0 and 360.0
type Degrees float64

// Fix represents type of GPS fix
type Fix string

const (
	// None means we didn't get a fix
	None Fix = "none"
	// TwoDimensional fix
	TwoDimensional Fix = "2d"
	// ThreeDimensional fix
	ThreeDimensional Fix = "3d"
	// DGPS means a digital GPS fix
	DGPS Fix = "dgps"
	// PPS means that a military signal was used
	PPS Fix = "pps"
)

// DGPSStation represents a differential GPS station and varies between 0 to 1023
type DGPSStation int

// Extensions extend GPX by adding your own elements from another schema
type Extensions struct {
	XMLName xml.Name `xml:"extensions"`
}

type Decoder struct {
	decoder *xml.Decoder
}

func NewDecoder(r io.Reader) Decoder {
	return Decoder{
		decoder: xml.NewDecoder(r),
	}
}

func (dec *Decoder) Decode(v *GPX) error {
	return dec.decoder.Decode(v)
}

type Encoder struct {
	encoder *xml.Encoder
}

func NewEncoder(w io.Writer) Encoder {
	return Encoder{
		encoder: xml.NewEncoder(w),
	}
}

func (enc *Encoder) Encode(v *GPX) error {
	return enc.encoder.Encode(v)
}
