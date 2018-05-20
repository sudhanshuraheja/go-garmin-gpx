package gpx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gpx "github.com/sudhanshuraheja/go-gpx"
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

	assert.Equal(t, "Untitled", g.Tracks.Name)

	assert.Equal(t, gpx.Longitude(-77.02016168273985), g.Tracks.TrackSegments[0].TrackPoint[0].Longitude)
	assert.Equal(t, gpx.Latitude(38.92747367732227), g.Tracks.TrackSegments[0].TrackPoint[0].Latitude)
	assert.Equal(t, 25.600000381469727, g.Tracks.TrackSegments[0].TrackPoint[0].Elevation)
	assert.Equal(t, "2012-10-24T23:29:40.000Z", g.Tracks.TrackSegments[0].TrackPoint[0].Timestamp)

	assert.Equal(t, gpx.Longitude(-77.02014584094286), g.Tracks.TrackSegments[0].TrackPoint[1].Longitude)
	assert.Equal(t, gpx.Latitude(38.927609380334616), g.Tracks.TrackSegments[0].TrackPoint[1].Latitude)
	assert.Equal(t, 35.599998474121094, g.Tracks.TrackSegments[0].TrackPoint[1].Elevation)
	assert.Equal(t, "2012-10-24T23:30:00.000Z", g.Tracks.TrackSegments[0].TrackPoint[1].Timestamp)

	// <trk>
	// <trkseg>
	//   <trkpt lon="-77.02016168273985" lat="38.92747367732227">
	//     <extensions>
	//       <gpxtpx:TrackPointExtension>
	//         <gpxtpx:hr>130</gpxtpx:hr>
	//       </gpxtpx:TrackPointExtension>
	//     </extensions>
	//   </trkpt>
	//   <trkpt lon="-77.02014584094286" lat="38.927609380334616">
	//     <extensions>
	//       <gpxtpx:TrackPointExtension>
	//         <gpxtpx:hr>134</gpxtpx:hr>
	//       </gpxtpx:TrackPointExtension>
	//     </extensions>
	//   </trkpt>

}

func Test_SpecParser(t *testing.T) {
	// file := "./samples/spec.gpx"
}

func Test_StLouisZooParser(t *testing.T) {
	// file := "./samples/StLouisZoo.gpx"
}

func Test_StravaSampleParser(t *testing.T) {
	// file := "./samples/strava-1427712053.gpx"
}

func Test_WikipediaParser(t *testing.T) {
	// file := "./samples/wikipedia-sample.gpx"
}
