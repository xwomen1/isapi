package hikivision

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"os"
)

type DownloadRequest struct {
	XMLName      xml.Name `xml:"downloadRequest"`
	PlaybackURI  string   `xml:"playbackURI"`
	StartTime    string   `xml:"-"`
	EndTime      string   `xml:"-"`
	MetadataList []MetadataDescriptor
}

// MetadataDescriptor struct đại diện cho mô tả metadata
type MetadataDescriptor struct {
	MetadataDescriptor string `xml:"metadataDescriptor"`
}

// DownloadVideoResult struct đại diện cho kết quả tải video
type DownloadVideoResult struct {
	XMLName      xml.Name `xml:"downloadVideo"`
	XMLVersion   string   `xml:"version,attr"`
	XMLNamespace string   `xml:"xmlns,attr"`
	Response     bool     `xml:"responseStatus"`
	StatusStrg   string   `xml:"responseStatusStrg"`
}

func (c *Client) PostDownloadVideo(data *DownloadRequest) (resp *ResponseStatus, respData *DownloadVideoResult, err error) {
	// Construct the API endpoint URL
	path := "/ISAPI/ContentMgmt/download"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, nil, err
	}
	bodyBytes, err := c.PostXML(u, data)
	if err != nil {
		fmt.Println("======> Post XML Unsuccessful: ", err)
		return nil, nil, err
	}
	err = os.MkdirAll("videos", 0755)
	if err != nil {
		fmt.Println("======> Creating Directory Unsuccessful: ", err)
		return nil, nil, err
	}

	file, err := os.Create("videos/response.mp4")
	if err != nil {
		fmt.Println("======> Creating File Unsuccessful: ", err)
		return nil, nil, err
	}
	defer file.Close()

	_, err = file.Write(bodyBytes)
	if err != nil {
		fmt.Println("======> Writing File Unsuccessful: ", err)
		return nil, nil, err
	}

	fmt.Println("=======> Post data success: ", respData)
	return resp, respData, nil
}
