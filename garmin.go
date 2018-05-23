package gpx

import "encoding/xml"

// This files defines Garmin extensions to be used with the GPX 1.1 schema
// https://www8.garmin.com/xmlschemas/GpxExtensions/v3/GpxExtensionsv3.xsd
// https://www8.garmin.com/xmlschemas/WaypointExtensionv1.xsd
// https://www8.garmin.com/xmlschemas/TrackPointExtensionv1.xsd

// WayPointExtensions extend GPX by adding your own elements from another schema
type WayPointExtensions struct {
	XMLName            xml.Name           `xml:"extensions"`
	WayPointExtensions *WayPointExtension `xml:"WaypointExtension,omitempty"`
}

// WayPointExtension add data fields available in Garmin GDB waypoints that cannot be represented in waypoints in GPX 1.1 instances
type WayPointExtension struct {
	XMLName     xml.Name       `xml:"WaypointExtension"`
	Proximity   Metres         `xml:"Proximity,omitempty"`
	Temperature DegreesCelcius `xml:"Temperature,omitempty"`
	Depth       Metres         `xml:"Depth,omitempty"`
	DisplayMode DisplayMode    `xml:"DisplayMode,omitempty"`
	Categories  Categories     `xml:"Categories,omitempty"`
	Address     Address        `xml:"Address,omitempty"`
	PhoneNumber []PhoneNumber  `xml:"PhoneNumber,omitempty"`
	Samples     int            `xml:"Samples,omitempty"`
	Expiration  string         `xml:"Expiration,omitempty"`
}

// RouteExtensions extend GPX by adding your own elements from another schema
type RouteExtensions struct {
	XMLName         xml.Name        `xml:"extensions"`
	RouteExtensions *RouteExtension `xml:"RouteExtension,omitempty"`
}

// RouteExtension tracks temperature, heart rate and cadence specific to garmin devices
type RouteExtension struct {
	XMLName      xml.Name         `xml:"RouteExtension"`
	IsAutoNamed  bool             `xml:"IsAutoNamed,omitempty"`
	DisplayColor DisplayColor     `xml:"DisplayColor,omitempty"`
	Extensions   GarminExtensions `xml:"Extensions,omitempty"`
}

// RoutePointExtensions extend GPX by adding your own elements from another schema
type RoutePointExtensions struct {
	XMLName              xml.Name             `xml:"extensions"`
	RoutePointExtensions *RoutePointExtension `xml:"RoutePointExtension,omitempty"`
}

// RoutePointExtension tracks temperature, heart rate and cadence specific to garmin devices
type RoutePointExtension struct {
	XMLName        xml.Name         `xml:"RoutePointExtension"`
	Subclass       SubClass         `xml:"Subclass,omitempty"`
	AutoRoutePoint AutoRoutePoint   `xml:"rpt,omitempty"`
	Extensions     GarminExtensions `xml:"Extensions,omitempty"`
}

// TrackExtensions extend GPX by adding your own elements from another schema
type TrackExtensions struct {
	XMLName         xml.Name        `xml:"extensions"`
	TrackExtensions *TrackExtension `xml:"TrackExtension,omitempty"`
}

// TrackExtension tracks temperature, heart rate and cadence specific to garmin devices
type TrackExtension struct {
	XMLName      xml.Name         `xml:"TrackExtension"`
	DisplayColor DisplayColor     `xml:"DisplayColor,omitempty"`
	Extensions   GarminExtensions `xml:"Extensions,omitempty"`
}

// TrackPointExtensions extend GPX by adding your own elements from another schema
type TrackPointExtensions struct {
	XMLName              xml.Name             `xml:"extensions"`
	TrackPointExtensions *TrackPointExtension `xml:"TrackPointExtension,omitempty"`
}

// TrackPointExtension tracks temperature, heart rate and cadence specific to garmin devices
// From https://www8.garmin.com/xmlschemas/TrackPointExtensionv1.xsd
type TrackPointExtension struct {
	XMLName      xml.Name             `xml:"TrackPointExtension"`
	Temperature  DegreesCelcius       `xml:"atemp,omitempty"`
	WTemperature DegreesCelcius       `xml:"wtemp,omitempty"`
	Depth        Metres               `xml:"depth,omitempty"`
	HeartRate    BeatsPerMinute       `xml:"hr,omitempty"`
	Cadence      RevolutionsPerMinute `xml:"cad,omitempty"`
	Extensions   GarminExtensionsV1   `xml:"extensions,omitempty"`
}

// TrackPointExtension tracks temperature, heart rate and cadence specific to garmin devices
// From https://www8.garmin.com/xmlschemas/GpxExtensions/v3/GpxExtensionsv3.xsd
// type TrackPointExtension struct {
// 	XMLName     xml.Name       `xml:"TrackPointExtension"`
// 	Temperature DegreesCelcius `xml:"Temperature,omitempty"`
// 	Depth       Metres         `xml:"Depth,omitempty"`
// 	Extensions  Extensions     `xml:"Extensions,omitempty"`
// }

// GarminExtensions handles extensions in garmin extensions
type GarminExtensions struct {
	XMLName xml.Name `xml:"Extensions"`
}

// GarminExtensionsV1 handles extensions in garmin extensions
type GarminExtensionsV1 struct {
	XMLName xml.Name `xml:"extensions"`
}

// Categories contains a list of categories that a waypoint has been assigned
type Categories struct {
	XMLName  xml.Name `xml:"Categories"`
	Category []string `xml:"Category,omitempty"`
}

// Address is the address, duh
type Address struct {
	XMLName       xml.Name          `xml:"Address"`
	StreetAddress []string          `xml:"StreetAddress,omitempty"`
	City          string            `xml:"City,omitempty"`
	State         string            `xml:"State,omitempty"`
	Country       string            `xml:"Country,omitempty"`
	PostalCode    string            `xml:"PostalCode,omitempty"`
	Extensions    *GarminExtensions `xml:"Extensions,omitempty"`
}

// PhoneNumber saves the phone number and type
type PhoneNumber struct {
	XMLName  xml.Name `xml:"PhoneNumber"`
	Category string   `xml:"Category,attr,omitempty"`
	Number   string   `xml:"Number,omitempty"`
}

// DisplayMode contains a string that specifies how a waypoint should be displayed on a map.
type DisplayMode string

const (
	// SymbolOnly shows only the symbol
	SymbolOnly DisplayMode = "SymbolOnly"
	// SymbolAndName shows both symbol and name
	SymbolAndName DisplayMode = "SymbolAndName"
	// SymbolAndDescription shows both symbol and description
	SymbolAndDescription DisplayMode = "SymbolAndDescription"
)

// DisplayColor contains which color to show in
type DisplayColor string

const (
	// Black color
	Black DisplayColor = "Black"
	// DarkRed color
	DarkRed DisplayColor = "DarkRed"
	// DarkGreen color
	DarkGreen DisplayColor = "DarkGreen"
	// DarkYellow color
	DarkYellow DisplayColor = "DarkYellow"
	// DarkBlue color
	DarkBlue DisplayColor = "DarkBlue"
	// DarkMagenta color
	DarkMagenta DisplayColor = "DarkMagenta"
	// DarkCyan color
	DarkCyan DisplayColor = "DarkCyan"
	// LightGray color
	LightGray DisplayColor = "LightGray"
	// DarkGray color
	DarkGray DisplayColor = "DarkGray"
	// Red color
	Red DisplayColor = "Red"
	// Green color
	Green DisplayColor = "Green"
	// Yellow color
	Yellow DisplayColor = "Yellow"
	// Blue color
	Blue DisplayColor = "Blue"
	// Magenta color
	Magenta DisplayColor = "Magenta"
	// Cyan color
	Cyan DisplayColor = "Cyan"
	// White color
	White DisplayColor = "White"
	// Transparent color
	Transparent DisplayColor = "Transparent"
)

// AutoRoutePoint (not sure what this does)
type AutoRoutePoint struct {
	XMLName   xml.Name  `xml:"rpt"`
	Latitude  Latitude  `xml:"lat,attr,omitempty"`
	Longitude Longitude `xml:"lon,attr,omitempty"`
	SubClass  SubClass  `xml:"Subclass,omitempty"`
}

// SubClass (not sure what this does)
type SubClass string

// Metres is used to measure length
type Metres float64

// DegreesCelcius is used to measure degree celcius
type DegreesCelcius float64

// BeatsPerMinute is used to measure frequency
type BeatsPerMinute int

// RevolutionsPerMinute is used to measure cadence
type RevolutionsPerMinute int
