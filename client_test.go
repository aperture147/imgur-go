package imgur_go

import (
	"os"
	"testing"
)

func TestClient_PostImageUrl(t *testing.T) {
	client := NewClient(os.Getenv("IMGUR_CLIENT_ID"), os.Getenv("IMGUR_TOKEN"))
	info, err := client.PostImageUrl(ImageInfo{
		Title:       "Test",
		Description: "Test Image",
		Url:         "https://i.picsum.photos/id/887/200/200.jpg?hmac=yOynpt597y5pLfJ5SsRVVKZiT5MXElbhtgUYeRzu3S4",
	})

	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("PostImageUrl succeeded")
	t.Log(info.Url)
}
