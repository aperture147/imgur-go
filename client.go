package imgur_go

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

type Client struct {
	Id    string
	Token string
	*http.Client
}

const ImgurApiBase = "https://api.imgur.com/3"

func NewClient(id, token string) *Client {
	return &Client{
		Id:     id,
		Token:  token,
		Client: &http.Client{},
	}
}

func (c *Client) PostImage(image *Image) error {
	return c.PostForm("/upload", map[string]io.Reader{"image": image})
}

func (c *Client) PostImageUrl(name, url string) error {
	return c.PostForm("/upload", map[string]io.Reader{"image": strings.NewReader(url)})
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Client-ID "+c.Id)
	req.Header.Add("Authorization", "Bearer "+c.Token)

	return c.Client.Do(req)
}

func (c *Client) PostForm(path string, form map[string]io.Reader) error {
	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)
	defer w.Close()

	var err error
	for field, data := range form {
		if data == nil {
			return errors.New(fmt.Sprint("field ", field, " doesn't contain any data"))
		}
		var fw io.Writer

		if rc, ok := data.(NamedReadCloser); ok {
			if rc.Name() != "" {
				if fw, err = w.CreateFormFile(field, rc.Name()); err != nil {
					return err
				}
			} else {
				if fw, err = w.CreateFormField(field); err != nil {
					return err
				}
			}

			rc.Close()
		} else {
			if fw, err = w.CreateFormField(field); err != nil {
				return err
			}
		}
		if _, err = io.Copy(fw, data); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(http.MethodPost, ImgurApiBase+path, &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Length", "0")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	log.Printf(res.Status)
	return nil
}
