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
		Url:         "https://cdn.discordapp.com/attachments/725214545307762740/727192857131352206/wp_ss_20161110_0001.png",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log("image URL:", info.Url)
}
