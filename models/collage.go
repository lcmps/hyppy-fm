package models

// ImgHelper is a helper structure
type ImgHelper []struct {
	Size string `xml:"size,attr"`
	Url  string `xml:",chardata"`
}
