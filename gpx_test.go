package gpx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gpx "github.com/sudhanshuraheja/go-garmin-gpx"
)

func Test_WrongFileParser(t *testing.T) {
	file := "./samples/DoesNotExist.gpx"
	_, err := gpx.ParseFile(file)
	assert.Contains(t, err.Error(), "no such file or directory")
}

func Test_InvalidXMLParser(t *testing.T) {
	file := "./samples/error.gpx"
	_, err := gpx.ParseFile(file)
	assert.Contains(t, err.Error(), "invalid syntax")
}

func Test_MapboxParser(t *testing.T) {
	file := "./samples/mapbox.gpx"
	g, err := gpx.ParseFile(file)
	assert.Nil(t, err)

	assert.Equal(t, "connect.garmin.com", g.Metadata.Links[0].URL)
	assert.Equal(t, "Garmin Connect", g.Metadata.Links[0].Text)
	assert.Equal(t, "2012-10-24T23:22:51.000Z", g.Metadata.Timestamp)

	assert.Equal(t, "Untitled", g.Tracks[0].Name)

	point0 := g.Tracks[0].TrackSegments[0].TrackPoint[0]
	assert.Equal(t, gpx.Longitude(-77.02016168273985), point0.Longitude)
	assert.Equal(t, gpx.Latitude(38.92747367732227), point0.Latitude)
	assert.Equal(t, 25.600000381469727, point0.Elevation)
	assert.Equal(t, "2012-10-24T23:29:40.000Z", point0.Timestamp)
	assert.Equal(t, gpx.BeatsPerMinute(130), point0.Extensions.TrackPointExtensions.HeartRate)

	point1 := g.Tracks[0].TrackSegments[0].TrackPoint[1]
	assert.Equal(t, gpx.Longitude(-77.02014584094286), point1.Longitude)
	assert.Equal(t, gpx.Latitude(38.927609380334616), point1.Latitude)
	assert.Equal(t, 35.599998474121094, point1.Elevation)
	assert.Equal(t, "2012-10-24T23:30:00.000Z", point1.Timestamp)
	assert.Equal(t, gpx.BeatsPerMinute(134), point1.Extensions.TrackPointExtensions.HeartRate)
}

func Test_SpecParser(t *testing.T) {
	file := "./samples/spec.gpx"
	g, err := gpx.ParseFile(file)
	assert.Nil(t, err)

	assert.Equal(t, "name", g.Metadata.Name)
	assert.Equal(t, "description", g.Metadata.Description)
	assert.Equal(t, "authorName", g.Metadata.Author.Name)
	assert.Equal(t, "emailID", g.Metadata.Author.Email.ID)
	assert.Equal(t, "emailDomain", g.Metadata.Author.Email.Domain)
	assert.Equal(t, "http://url.tld", g.Metadata.Author.Link.URL)
	assert.Equal(t, "someText", g.Metadata.Author.Link.Text)
	assert.Equal(t, "linkType", g.Metadata.Author.Link.Type)
	assert.Equal(t, "copyrightAuthor", g.Metadata.Copyright.Author)
	assert.Equal(t, 2019, g.Metadata.Copyright.Year)
	assert.Equal(t, "http://url.tld", g.Metadata.Copyright.License)
	assert.Equal(t, "2018-02-26T22:58:34Z", g.Metadata.Timestamp)
	assert.Equal(t, "keywords", g.Metadata.Keywords)
	assert.Equal(t, gpx.Latitude(-90.0), g.Metadata.Bounds.MinimumLatitude)
	assert.Equal(t, gpx.Longitude(-180.0), g.Metadata.Bounds.MinimumLongitude)
	assert.Equal(t, gpx.Latitude(90.0), g.Metadata.Bounds.MaximumLatitude)
	assert.Equal(t, gpx.Longitude(180.0), g.Metadata.Bounds.MaximumLongitude)

	assert.Equal(t, gpx.Latitude(90.0), g.Waypoints[0].Latitude)
	assert.Equal(t, gpx.Longitude(180.0), g.Waypoints[0].Longitude)
	assert.Equal(t, 12.0, g.Waypoints[0].Elevation)
	assert.Equal(t, gpx.Degrees(360.0), g.Waypoints[0].MagneticVariation)
	assert.Equal(t, 0.0, g.Waypoints[0].GeoIDHeight)
	assert.Equal(t, gpx.Fix("3d"), g.Waypoints[0].Fix)
	assert.Equal(t, 1, g.Waypoints[0].Sat)
	assert.Equal(t, 0.0, g.Waypoints[0].HorizontalDilutionOfPrecision)
	assert.Equal(t, 0.0, g.Waypoints[0].HorizontalDilutionOfPrecision)
	assert.Equal(t, 0.0, g.Waypoints[0].PositionDilutionOfPrecision)
	assert.Equal(t, 0.0, g.Waypoints[0].AgeOfGpsData)
	assert.Equal(t, gpx.DGPSStation(1023), g.Waypoints[0].DifferentialGPSID)

	assert.Equal(t, "routeName", g.Routes[0].Name)
	assert.Equal(t, 10.0, g.Routes[0].RoutePoints[0].Elevation)
	assert.Equal(t, "trackName", g.Tracks[0].Name)
	assert.Equal(t, "trackType", g.Tracks[0].Type)
	assert.Equal(t, "pointName", g.Tracks[0].TrackSegments[0].TrackPoint[0].Name)
}

func Test_StLouisZooParser(t *testing.T) {
	file := "./samples/StLouisZoo.gpx"
	g, err := gpx.ParseFile(file)
	assert.Nil(t, err)

	assert.Equal(t, "St Louis Zoo sample", g.Metadata.Name)
	assert.Equal(t, "This self guided,GPS enabled tour of the world famous St. Louis Zoo, has 85 points of interest. Narratives in english,explaining each exhibit and provides guidance to zoo facilities.This audio tour guide can enhance your next visit.", g.Metadata.Description)
	assert.Equal(t, "wizardone, using GeoTours", g.Metadata.Author.Name)
	assert.Equal(t, "http://www.geovative.com/view?t=GEIF", g.Metadata.Author.Link.URL)
	assert.Equal(t, "St Louis Zoo sample", g.Metadata.Author.Link.Text)
	assert.Equal(t, "2008-02-26T19:49:13", g.Metadata.Timestamp)
	assert.Contains(t, g.Metadata.Keywords, "Audio tour guide")

	assert.Equal(t, gpx.Latitude(38.63473), g.Waypoints[0].Latitude)
	assert.Equal(t, gpx.Longitude(-90.29408), g.Waypoints[0].Longitude)
	assert.Equal(t, "Asian Elephant", g.Waypoints[0].Name)
	assert.Equal(t, "Data/Location3152-1", g.Waypoints[0].Links[0].URL)
	assert.Equal(t, "Waypoint", g.Waypoints[0].Symbol)
	assert.Equal(t, gpx.Metres(15.24), g.Waypoints[0].Extensions.WayPointExtensions.Proximity)
	assert.Equal(t, gpx.DisplayMode("SymbolAndName"), g.Waypoints[0].Extensions.WayPointExtensions.DisplayMode)

	assert.Equal(t, gpx.Latitude(38.63368), g.Waypoints[1].Latitude)
	assert.Equal(t, gpx.Longitude(-90.28679), g.Waypoints[1].Longitude)
	assert.Equal(t, "Bactrian Camel", g.Waypoints[1].Name)
	assert.Equal(t, "Data/Location3152-2", g.Waypoints[1].Links[0].URL)
	assert.Equal(t, "Waypoint", g.Waypoints[1].Symbol)
	assert.Equal(t, gpx.Metres(15.24), g.Waypoints[0].Extensions.WayPointExtensions.Proximity)
	assert.Equal(t, gpx.DisplayMode("SymbolAndName"), g.Waypoints[0].Extensions.WayPointExtensions.DisplayMode)
}

func Test_StravaSampleParser(t *testing.T) {
	file := "./samples/strava-1427712053.gpx"
	g, err := gpx.ParseFile(file)
	assert.Nil(t, err)

	assert.Equal(t, "2018-02-26T22:58:34Z", g.Metadata.Timestamp)

	assert.Equal(t, "Really bad GPS", g.Tracks[0].Name)
	assert.Equal(t, 16.6, g.Tracks[0].TrackSegments[0].TrackPoint[0].Elevation)
	assert.Equal(t, gpx.DegreesCelcius(28), g.Tracks[0].TrackSegments[0].TrackPoint[0].Extensions.TrackPointExtensions.Temperature)
	assert.Equal(t, gpx.BeatsPerMinute(95), g.Tracks[0].TrackSegments[0].TrackPoint[0].Extensions.TrackPointExtensions.HeartRate)
	assert.Equal(t, gpx.RevolutionsPerMinute(0), g.Tracks[0].TrackSegments[0].TrackPoint[0].Extensions.TrackPointExtensions.Cadence)
	assert.Equal(t, 16.4, g.Tracks[0].TrackSegments[0].TrackPoint[1].Elevation)
	assert.Equal(t, gpx.DegreesCelcius(28), g.Tracks[0].TrackSegments[0].TrackPoint[1].Extensions.TrackPointExtensions.Temperature)
	assert.Equal(t, gpx.BeatsPerMinute(95), g.Tracks[0].TrackSegments[0].TrackPoint[1].Extensions.TrackPointExtensions.HeartRate)
	assert.Equal(t, gpx.RevolutionsPerMinute(0), g.Tracks[0].TrackSegments[0].TrackPoint[1].Extensions.TrackPointExtensions.Cadence)
}

func Test_WikipediaParser(t *testing.T) {
	file := "./samples/wikipedia-sample.gpx"
	g, err := gpx.ParseFile(file)
	assert.Nil(t, err)

	assert.Equal(t, "2009-10-17T22:58:43Z", g.Metadata.Timestamp)
	assert.Equal(t, "http://www.garmin.com", g.Metadata.Links[0].URL)
	assert.Equal(t, "Garmin International", g.Metadata.Links[0].Text)

	assert.Equal(t, "Example GPX Document", g.Tracks[0].Name)
	assert.Equal(t, gpx.Latitude(47.644548), g.Tracks[0].TrackSegments[0].TrackPoint[0].Latitude)
	assert.Equal(t, 4.46, g.Tracks[0].TrackSegments[0].TrackPoint[0].Elevation)
	assert.Equal(t, gpx.Latitude(47.644548), g.Tracks[0].TrackSegments[0].TrackPoint[1].Latitude)
	assert.Equal(t, 4.94, g.Tracks[0].TrackSegments[0].TrackPoint[1].Elevation)
	assert.Equal(t, gpx.Latitude(47.644548), g.Tracks[0].TrackSegments[0].TrackPoint[2].Latitude)
	assert.Equal(t, 6.87, g.Tracks[0].TrackSegments[0].TrackPoint[2].Elevation)
}
