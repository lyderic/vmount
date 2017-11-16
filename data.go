package main

import (
	"encoding/xml"
)

const version = "0.0.2"

type Configuration struct {
	FavoritesPath string
}

var config Configuration

type VeraCrypt struct {
	XMLName   xml.Name  `xml:"VeraCrypt"`
	Favorites Favorites `xml:"favorites"`
}

type Favorites struct {
	XMLName xml.Name `xml:"favorites"`
	Volumes []Volume `xml:"volume"`
}

type Volume struct {
	XMLName         xml.Name `xml:"volume"`
	Path            string   `xml:",chardata"`
	ShortPath       string
	Mountpoint      string `xml:"mountpoint,attr"`
	ShortMountpoint string
	Readonly        int `xml:"readonly,attr"`
	Slotnumber      int `xml:"slotnumber,attr"`
	System          int `xml:"system,attr"`
}

const template = `<?xml version="1.0" encoding="utf-8"?>
<VeraCrypt>
  <favorites>
    <volume mountpoint="/path/to/mountpoint" readonly="0" slotnumber="1" system="0">/path/to/volume</volume>
  </favorites>
</VeraCrypt>
`
