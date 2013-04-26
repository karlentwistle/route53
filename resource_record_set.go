package route53

import (
	"encoding/xml"
	"net/http"
  "fmt"
)

type ChangeResourceRecordSetsRequest struct {
	ZoneID  string   `xml:"-"`
	Comment string   `xml:"ChangeBatch>Comment"`
	Changes []Change `xml:"ChangeBatch>Changes>Change"`
	Xmlns   string   `xml:"xmlns,attr"`
}

type Change struct {
	Action        string
	Name          string `xml:"ResourceRecordSet>Name"`
	Type          string `xml:"ResourceRecordSet>Type"`
	TTL           int    `xml:"ResourceRecordSet>TTL"`
	Value         string `xml:"ResourceRecordSet>ResourceRecords>ResourceRecord>Value"`
}

func (c *ChangeResourceRecordSetsRequest) XML() (s string, err error) {
	c.Xmlns = `https://route53.amazonaws.com/doc/2012-12-12/`
	byteXML, err := xml.MarshalIndent(c, "", `   `)
	if err != nil {
		return "", err
	}
	s = xml.Header + string(byteXML)
	return
}

func (c *ChangeResourceRecordSetsRequest) Create(a AccessIdentifiers) (req *http.Response, err error) {
	postData, err := c.XML()
	if err != nil {
		return nil, err
	}
	url := postURL + `/` + c.ZoneID + `/rrset`
  fmt.Println(url)
	req, err = RemotePost(url, postData, a)
	return
}
