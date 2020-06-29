package imgur_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type Client struct {
	Id    string
	Token string
	*http.Client
}

const ApiBase = "https://api.imgur.com/3/"

type ImageInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"link"`
	Name        string `json:"name,omitempty"`
	DeleteHash  string `json:"deletehash"`
}

type AlbumInfo struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Url         string      `json:"link"`
	Images      []ImageData `json:"images"`
	DeleteHash  string      `json:"deletehash"`
}

func NewClient(id, token string) *Client {
	return &Client{
		Id:     strings.TrimSpace(id),
		Token:  strings.TrimSpace(token),
		Client: &http.Client{},
	}
}

func (c *Client) PostImage(image Image) (*ImageInfo, error) {
	return c.PostForm("/upload", map[string]io.Reader{
		"image":       image,
		"title":       strings.NewReader(image.Info.Title),
		"description": strings.NewReader(image.Info.Description)})
}

func (c *Client) PostImageUrl(info ImageInfo) (*ImageInfo, error) {
	return c.PostForm("/upload", map[string]io.Reader{
		"image":       strings.NewReader(info.Url),
		"title":       strings.NewReader(info.Title),
		"description": strings.NewReader(info.Description)})
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	} else {
		req.Header.Add("Authorization", "Client-ID "+c.Id)
	}

	return c.Client.Do(req)
}

func (c *Client) PostForm(path string, form map[string]io.Reader) (*ImageInfo, error) {
	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)
	var err error
	for field, data := range form {
		// Prevent some error related to nil data
		if data == nil {
			return nil, errors.New(fmt.Sprint("field ", field, " doesn't contain any data"))
		}

		var fw io.Writer

		if rc, ok := data.(NamedReader); ok {
			if rc.Name() != "" {
				if fw, err = w.CreateFormFile(field, rc.Name()); err != nil {
					return nil, err
				}
			} else {
				if fw, err = w.CreateFormField(field); err != nil {
					return nil, err
				}
			}
			err = rc.Close()
			if err != nil {
				return nil, err
			}
		} else {
			if fw, err = w.CreateFormField(field); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, data); err != nil {
			return nil, err
		}
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, ApiBase+path, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", w.FormDataContentType())
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		v := new(ServerError)
		_ = json.Unmarshal(b, v)
		return nil, v
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	v := &ImageResponse{}
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return &v.Data.ImageInfo, nil
}
