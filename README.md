# GO-GPX

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

## Samples

You can find some samples of GPX files in the `/samples` folder