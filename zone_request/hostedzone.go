package hostedzone

import (
	"encoding/xml"
	"net/http"
	"route53"
)

const postURL = `https://route53.amazonaws.com/2012-12-12/hostedzone`

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

func (hz *CreateHostedZoneRequest) Create() (req *http.Response, err error) {
	postData, err := hz.XML()
	if err != nil {
		return nil, err
	}
	req, err = route53.RemotePost(postURL, postData)
	return
}
