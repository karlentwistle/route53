package route53

import (
	"encoding/xml"
	"net/http"
)

type CreateHostedZoneRequest struct {
	Name            string
	CallerReference string
	Comment         string `xml:"HostedZoneConfig>Comment"`
	Xmlns           string `xml:"xmlns,attr"`
}

func (hz *CreateHostedZoneRequest) XML() (s string, err error) {
	hz.Xmlns = `https://route53.amazonaws.com/doc/2012-12-12/`
	byteXML, err := xml.MarshalIndent(hz, "", `  `)
	if err != nil {
		return "", err
	}
	s = xml.Header + string(byteXML)
	return
}

func (hz *CreateHostedZoneRequest) Create(a AccessIdentifiers) (req *http.Response, err error) {
	postData, err := hz.XML()
	if err != nil {
		return nil, err
	}
	req, err = post(postURL, postData, a.headers())
	return
}
