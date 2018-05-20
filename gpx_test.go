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

	point0 := g.Tracks.TrackSegments[0].TrackPoint[0]
	assert.Equal(t, gpx.Longitude(-77.02016168273985), point0.Longitude)
	assert.Equal(t, gpx.Latitude(38.92747367732227), point0.Latitude)
	assert.Equal(t, 25.600000381469727, point0.Elevation)
	assert.Equal(t, "2012-10-24T23:29:40.000Z", point0.Timestamp)
	assert.Equal(t, 130, point0.Extensions.TrackPointExtensions.HeartRate)

	point1 := g.Tracks.TrackSegments[0].TrackPoint[1]
	assert.Equal(t, gpx.Longitude(-77.02014584094286), point1.Longitude)
	assert.Equal(t, gpx.Latitude(38.927609380334616), point1.Latitude)
	assert.Equal(t, 35.599998474121094, point1.Elevation)
	assert.Equal(t, "2012-10-24T23:30:00.000Z", point1.Timestamp)
	assert.Equal(t, 134, point1.Extensions.TrackPointExtensions.HeartRate)
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

	// <copyright author="copyrightAuthor">
	// 	<year>2019</year>
	// 	<license>http://url.tld</license>
	// </copyright>
	// <link href="http://url.tld">
	// 	<text>someText</text>
	// 	<type>linkType</type>
	// </link>
	// <time>2018-02-26T22:58:34Z</time>
	// <keywords>keywords</keywords>
	// <bounds minlat="-90.0" minlon="-180.0" maxlat="90.0" minlon="180.0"></bounds>
	// <extensions></extensions>

	// <wpt lat="90.0" lon="180.0">
	//     <ele>12.0</ele>
	//     <time>2018-02-26T22:58:34Z</time>
	//     <magvar>360.0</magvar>
	//     <geoidheight>0.0</geoidheight>
	//     <name>pointName</name>
	//     <cmt>pointComment</cmt>
	//     <desc>pointDescription</desc>
	//     <src>pointSource</src>
	//     <link href="http://url.tld">
	//         <text>someText</text>
	//         <type>linkType</type>
	//     </link>
	//     <sym>pointSymbol</sym>
	//     <type>pointType</type>
	//     <fix>3d</fix>
	//     <sat>1</sat>
	//     <hdop>0.0</hdop>
	//     <vdop>0.0</vdop>
	//     <pdop>0.0</pdop>
	//     <ageofdgpsdata>0.0</ageofdgpsdata>
	//     <dgpsid>1023</dgpsid>
	//     <extensions></extensions>
	// </wpt>

	// <rte>
	//     <name>routeName</name>
	//     <cmt>routeComment</cmt>
	//     <desc>routeDescription</desc>
	//     <src>routeSrc</src>
	//     <link href="http://url.tld">
	//         <text>someText</text>
	//         <type>linkType</type>
	//     </link>
	//     <number>1</number>
	//     <type>routeType</type>
	//     <extensions></extensions>
	//     <rtept lat="90.0" lon="180.0">
	//         <ele>10.0</ele>
	//         <time>2018-02-26T22:58:34Z</time>
	//         <magvar>360.0</magvar>
	//         <geoidheight>0.0</geoidheight>
	//         <name>pointName</name>
	//         <cmt>pointComment</cmt>
	//         <desc>pointDescription</desc>
	//         <src>pointSource</src>
	//         <link href="http://url.tld">
	//             <text>someText</text>
	//             <type>linkType</type>
	//         </link>
	//         <sym>pointSymbol</sym>
	//         <type>pointType</type>
	//         <fix>3d</fix>
	//         <sat>1</sat>
	//         <hdop>0.0</hdop>
	//         <vdop>0.0</vdop>
	//         <pdop>0.0</pdop>
	//         <ageofdgpsdata>0.0</ageofdgpsdata>
	//         <dgpsid>1023</dgpsid>
	//         <extensions></extensions>
	//     </rtept>
	// </rte>

	// <trk>
	//     <name>trackName</name>
	//     <cmt>trackComment</cmt>
	//     <desc>trackDescription</desc>
	//     <src>trackSource</src>
	//     <link href="http://url.tld">
	//         <text>someText</text>
	//         <type>linkType</type>
	//     </link>
	//     <number>1</number>
	//     <type>trackType</type>
	//     <extensions></extensions>
	//     <trkseg>
	//         <trkpt lat="90.0" lon="180.0">
	//             <ele>10.0</ele>
	//             <time>2018-02-26T22:58:34Z</time>
	//             <magvar>360.0</magvar>
	//             <geoidheight>0.0</geoidheight>
	//             <name>pointName</name>
	//             <cmt>pointComment</cmt>
	//             <desc>pointDescription</desc>
	//             <src>pointSource</src>
	//             <link href="http://url.tld">
	//                 <text>someText</text>
	//                 <type>linkType</type>
	//             </link>
	//             <sym>pointSymbol</sym>
	//             <type>pointType</type>
	//             <fix>3d</fix>
	//             <sat>1</sat>
	//             <hdop>0.0</hdop>
	//             <vdop>0.0</vdop>
	//             <pdop>0.0</pdop>
	//             <ageofdgpsdata>0.0</ageofdgpsdata>
	//             <dgpsid>1023</dgpsid>
	//             <extensions></extensions>
	//         </trkpt>
	//         <extensions></extensions>
	//     </trkseg>
	// </trk>

}

func Test_StLouisZooParser(t *testing.T) {
	// file := "./samples/StLouisZoo.gpx"

	//     <metadata>
	//         <name>St Louis Zoo sample</name>
	//         <desc>This self guided,GPS enabled tour of the world famous St. Louis Zoo, has 85 points of interest. Narratives in english,explaining each exhibit and provides guidance to zoo facilities.This audio tour guide can enhance your next visit.</desc>
	//         <author>
	//             <name>wizardone, using GeoTours</name>
	//             <link href="http://www.geovative.com/view?t=GEIF">
	//                 <text>St Louis Zoo sample</text>
	//             </link>
	//         </author>
	//         <link href="http://www.geovative.com/view?t=GEIF">
	//             <text>St Louis Zoo sample</text>
	//         </link>
	//         <time>2008-02-26T19:49:13</time>
	//         <keywords>Audio tour guide
	// St.Louis Mo.
	// Zoo
	// Forest Park
	// Animals</keywords>
	//     </metadata>
	//     <wpt lat="38.63473" lon="-90.29408">
	//         <name>Asian Elephant</name>
	//         <desc>elephant</desc>
	//         <link href="Data/Location3152-1" />
	//         <sym>Waypoint</sym>
	//         <extensions>
	//             <gpxx:WaypointExtension xmlns:gpxx="http://www.garmin.com/xmlschemas/GpxExtensions/v3" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.garmin.com/xmlschemas/GpxExtensions/v3 http://www.garmin.com/xmlschemas/GpxExtensions/v3/GpxExtensionsv3.xsd">
	//                 <gpxx:Proximity>15.24</gpxx:Proximity>
	//                 <gpxx:DisplayMode>SymbolAndName</gpxx:DisplayMode>
	//             </gpxx:WaypointExtension>
	//         </extensions>
	//     </wpt>
	//     <wpt lat="38.63368" lon="-90.28679">
	//         <name>Bactrian Camel</name>
	//         <desc>camel</desc>
	//         <link href="Data/Location3152-2" />
	//         <sym>Waypoint</sym>
	//         <extensions>
	//             <gpxx:WaypointExtension xmlns:gpxx="http://www.garmin.com/xmlschemas/GpxExtensions/v3" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.garmin.com/xmlschemas/GpxExtensions/v3 http://www.garmin.com/xmlschemas/GpxExtensions/v3/GpxExtensionsv3.xsd">
	//                 <gpxx:Proximity>15.24</gpxx:Proximity>
	//                 <gpxx:DisplayMode>SymbolAndName</gpxx:DisplayMode>
	//             </gpxx:WaypointExtension>
	//         </extensions>
	//     </wpt>

}

func Test_StravaSampleParser(t *testing.T) {
	file := "./samples/strava-1427712053.gpx"
	g, err := gpx.ParseFile(file)
	assert.Nil(t, err)

	assert.Equal(t, 16.6, g.Tracks.TrackSegments[0].TrackPoint[0].Elevation)
	assert.Equal(t, 95, g.Tracks.TrackSegments[0].TrackPoint[0].Extensions.TrackPointExtensions.HeartRate)

	//    <metadata>
	// 	<time>2018-02-26T22:58:34Z</time>
	//    </metadata>
	//    <trk>
	// 	<name>Really bad GPS</name>
	// 	<trkseg>
	// 	 <trkpt lat="1.2793450" lon="103.8432030">
	// 	  <ele>16.6</ele>
	// 	  <time>2018-02-26T22:58:34Z</time>
	// 	  <extensions>
	// 	   <gpxtpx:TrackPointExtension>
	// 		<gpxtpx:atemp>28</gpxtpx:atemp>
	// 		<gpxtpx:hr>95</gpxtpx:hr>
	// 		<gpxtpx:cad>0</gpxtpx:cad>
	// 	   </gpxtpx:TrackPointExtension>
	// 	  </extensions>
	// 	 </trkpt>
	// 	 <trkpt lat="1.2793580" lon="103.8432120">
	// 	  <ele>16.4</ele>
	// 	  <time>2018-02-26T22:58:35Z</time>
	// 	  <extensions>
	// 	   <gpxtpx:TrackPointExtension>
	// 		<gpxtpx:atemp>28</gpxtpx:atemp>
	// 		<gpxtpx:hr>95</gpxtpx:hr>
	// 		<gpxtpx:cad>0</gpxtpx:cad>
	// 	   </gpxtpx:TrackPointExtension>
	// 	  </extensions>
	// 	 </trkpt>

}

func Test_WikipediaParser(t *testing.T) {
	// file := "./samples/wikipedia-sample.gpx"

	// 	<metadata>
	//     <link href="http://www.garmin.com">
	//       <text>Garmin International</text>
	//     </link>
	//     <time>2009-10-17T22:58:43Z</time>
	//   </metadata>
	//   <trk>
	//     <name>Example GPX Document</name>
	//     <trkseg>
	//       <trkpt lat="47.644548" lon="-122.326897">
	//         <ele>4.46</ele>
	//         <time>2009-10-17T18:37:26Z</time>
	//       </trkpt>
	//       <trkpt lat="47.644548" lon="-122.326897">
	//         <ele>4.94</ele>
	//         <time>2009-10-17T18:37:31Z</time>
	//       </trkpt>
	//       <trkpt lat="47.644548" lon="-122.326897">
	//         <ele>6.87</ele>
	//         <time>2009-10-17T18:37:34Z</time>
	//       </trkpt>
	//     </trkseg>
	//   </trk>
}
