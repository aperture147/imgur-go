package imgur_go

type ImageData struct {
	ImageInfo
	Id         string `json:"id"`
	Timestamp  int64  `json:"datetime"`
	Type       string `json:"type"`
	Animated   bool   `json:"animated"`
	DeleteHash string `json:"deletehash,omitempty"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

type AlbumData struct {
	AlbumInfo
	Id          string `json:"id"`
	Timestamp   int64  `json:"datetime"`
	ImagesCount int    `json:"images_count"`
}

type ResponseStatus struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

type ImageResponse struct {
	Data ImageData `json:"data"`
	ResponseStatus
}

type AlbumResponse struct {
	Data AlbumData `json:"data"`
	ResponseStatus
}

type ErrorResponse struct {
	Data struct {
		Error interface{} `json:"error"`
	} `json:"data"`
	ResponseStatus
}
