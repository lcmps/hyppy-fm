package models

// ImgHelper is a helper structure
type ImgHelper []struct {
	Size string `xml:"size,attr"`
	Url  string `xml:",chardata"`
}

// CollageParams default structure
type CollageParams struct {
	Username string `json:"username"`
	Period   string `json:"period"`
	Size     int    `json:"size"`
}

type UriParams struct {
	Username string `form:"username"`
	Period   string `form:"period"`
	Size     int    `form:"size"`
}

// InternalError default error structure
type InternalError struct {
	Reason string
}

// ResponseClg default response for clg/ endpoint
type ResponseClg struct {
	Path string
}
