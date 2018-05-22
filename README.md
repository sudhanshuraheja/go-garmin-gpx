# GO-GARMIN-GPX

GPX parsing library for the Go language which also support all Garmin extensions
#gpx #gpx-library #go #golang

## What is GPX

GPX is an XML schema defined as a common data format for software applications. You can find more about it on [Wikipedia](https://en.wikipedia.org/wiki/GPS_Exchange_Format). The schema itself can be found [here](http://www.topografix.com/GPX/1/1/).

## Give me more details about GPX

As of v1.1 (released on Aug 9, 2004), each GPX file has the following:

- wptType (WayPoint) which lists each individual waypoint
- rteType (Route) is a list of points leading to a waypoint which suggests where a person should or might go
- trkType (Track) is a list of points which explains the path that the person took
- extensions to handle data which isn't part of the main spec

## What is the Garmin Extension Spec

The [Garmin extension spec](https://www8.garmin.com/xmlschemas/GpxExtensions/v3/GpxExtensionsv3.xsd) details how to handle the data for Waypoint, Route, RoutePoint, Track and TrackPoint Extensions

## What all does this library cover

It can handle the following

- Metadata
- Bounds
- WayPoints
- Routes
- TrackSegments
- TrackPoints
- PointSegments
- Point
- Latitude
- Longitude
- Degrees
- Fix
- DGPSStation
- Copyright
- Person
- Link
- Email

It can handle the following from the Garmin Extension spec

- Waypoint Extensions
- Route Extensions
- RoutePoint Extensions
- Track Extensions
- TrackPoint Extensions

## Getting Started

You can install it into your project using

```bash
dep ensure -add github.com/sudhanshuraheja/go-garmin-gpx
```

If you would like to run the automated tests for the complete package, run this

```bash
make coverage
open ./coverage.html
```

We use the default golang coding conventions. Run the following to test for those

```bash
make fmt
make vet
make lint
```

## Usage

```go
file := "./samples/mapbox.gpx"
g, err := gpx.ParseFile(file)

if err != nil {
    return err
}

fmt.Println(g.Metadata.Timestamp)
```

## Samples

You can find some samples of GPX files in the `/samples` folder

## Roadmap

- Avoid reading the complete file into memory

## Contributing

Please read [CONTRIBUTING.md](https://github.com/sudhanshuraheja/go-garmin-gpx/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](https://github.com/sudhanshuraheja/go-garmin-gpx/blob/master/LICENSE) file for details