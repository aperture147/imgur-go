package imgur_go

type CommonResponse struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

type ImageResponse struct {
	Data ImageData `json:"data"`
	CommonResponse
}

type AlbumResponse struct {
	Data AlbumData `json:"data"`
	CommonResponse
}

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
