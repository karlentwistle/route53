package route53

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var t time.Time
var accessIdentifiers = AccessIdentifiers{AccessKey: "foo", SecretKey: "bar", time: t.Add(2)}

type webHandler struct {
	location string
}

func (h *webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Test-Header", "Testing")
	w.Header().Set("Date", time.Now().Format(time.RFC1123))
	if h.location == "" {
		io.WriteString(w, "hello, world!\n")
	} else {
		content, _ := ioutil.ReadFile(h.location)
		io.WriteString(w, string(content))
	}
}

func TestSignature(t *testing.T) {
	response := accessIdentifiers.signature()
	if response != "TlXhIyeN+etmLEuVWrz4i+deyqGhDs5P/9wLYq6IyHE=" {
		t.Fatal("incorrect signature encoding", response)
	}
}

func TestRemoteHeaders(t *testing.T) {
	handler := &webHandler{}
	server := httptest.NewServer(handler)
	headers, err := getHeaders(server.URL)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if headers.Get("Test-Header") == "Test-Header" {
		t.Fatal("could not read headers", headers)
	}
}

func TestRemoteTime(t *testing.T) {
	handler := &webHandler{}
	server := httptest.NewServer(handler)
	remoteTime, err := remoteTime(server.URL)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if remoteTime != time.Now().Format(time.RFC1123) {
		t.Fatal("remote time did not return", remoteTime)
	}
}

func TestGetBody(t *testing.T) {
	handler := &webHandler{}
	server := httptest.NewServer(handler)
	resp, err := getBody(server.URL, nil)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if string(resp) != "hello, world!\n" {
		t.Fatal("remote body did not return correctly", resp)
	}
}

func TestAwsRequestHeaders(t *testing.T) {
	awsHeaders := accessIdentifiers.headers()

	if awsHeaders.Get("Date") != accessIdentifiers.time.UTC().Format(time.ANSIC) {
		t.Fatal("incorrect Date in headers", awsHeaders)
	}

	if awsHeaders.Get("Content-Type") != "text/xml; charset=UTF-8" {
		t.Fatal("incorrect Content-Type in headers", awsHeaders)
	}

	if awsHeaders.Get("X-Amzn-Authorization") != ("AWS3-HTTPS AWSAccessKeyId=foo,Algorithm=HmacSHA256,Signature=TlXhIyeN+etmLEuVWrz4i+deyqGhDs5P/9wLYq6IyHE=") {
		t.Fatal("incorrect X-Amzn-Authorization in headers", awsHeaders)
	}

}

////////////// hosted_zone ////////////// 

const createHostedZoneRequest = `<?xml version="1.0" encoding="UTF-8"?>

<CreateHostedZoneRequest xmlns="https://route53.amazonaws.com/doc/2012-12-12/">
  <Name>DNS domain name</Name>
  <CallerReference>unique description</CallerReference>
  <HostedZoneConfig>
    <Comment>optional comment</Comment>
  </HostedZoneConfig>
</CreateHostedZoneRequest>`

var zoneRequest = CreateHostedZoneRequest{
	Name:            "DNS domain name",
	CallerReference: "unique description",
	Comment:         "optional comment",
}

func TestCreateHostedZoneXML(t *testing.T) {
	responseXML, err := zoneRequest.XML()
	if err != nil {
		t.Fatal("Error:", err)
	}

	if responseXML != createHostedZoneRequest {
		t.Fatal("returned XML is incorrectly formatted", responseXML, createHostedZoneRequest)
	}
}

////////////// resource_record ////////////// 

const changeResourceRecordSets = `<?xml version="1.0" encoding="UTF-8"?>

<ChangeResourceRecordSetsRequest xmlns="https://route53.amazonaws.com/doc/2012-12-12/">
   <ChangeBatch>
      <Comment>optional comment</Comment>
      <Changes>
         <Change>
            <Action>CREATE</Action>
            <ResourceRecordSet>
               <Name>DNS domain name</Name>
               <Type>DNS record type</Type>
               <TTL>300</TTL>
               <ResourceRecords>
                  <ResourceRecord>
                     <Value>applicable value for the record type</Value>
                  </ResourceRecord>
               </ResourceRecords>
            </ResourceRecordSet>
         </Change>
      </Changes>
   </ChangeBatch>
</ChangeResourceRecordSetsRequest>`

var resourceRecordSets = ChangeResourceRecordSetsRequest{
	Comment: "optional comment",
	Changes: []Change{
		{
			Action: "CREATE",
			Name:   "DNS domain name",
			Type:   "DNS record type",
			TTL:    300,
			Value:  "applicable value for the record type",
		},
	},
}

func TestCreateResourceRecordSetsXML(t *testing.T) {
	responseXML, err := resourceRecordSets.XML()
	if err != nil {
		t.Fatal("Error:", err)
	}

	if string(responseXML) != changeResourceRecordSets {
		t.Fatal("returned XML is incorrectly formatted", responseXML)
	}
}

////////////// read remote zones ////////////// 

func TestGenerateZones(t *testing.T) {
	response_xml, err := ioutil.ReadFile("spec/fixtures/list_hosted_zones_response.xml")

	if err != nil {
		t.Fatal("Error:", err)
	}

	hostedZones := generateZones(response_xml)

	if hostedZones.HostedZone[0].Id != "/hostedzone/foo" {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}

	if hostedZones.HostedZone[0].Name != "foo.who.com." {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}

	if hostedZones.HostedZone[0].CallerReference != "F2FCD646" {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}

	if hostedZones.HostedZone[0].RecordSetCount != 4 {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}

	if hostedZones.HostedZone[1].Id != "/hostedzone/mho" {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}

	if hostedZones.HostedZone[1].Name != "mho.woo.com." {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}

	if hostedZones.HostedZone[1].CallerReference != "96BA065A" {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}

	if hostedZones.HostedZone[1].RecordSetCount != 6 {
		t.Fatal("XML Unmarshal incorrectly", hostedZones)
	}
}

func TestAccessIdentifiersZones(t *testing.T) {
	handler := &webHandler{
		location: "spec/fixtures/list_hosted_zones_response.xml",
	}
	server := httptest.NewServer(handler)
	accessIdentifiers := accessIdentifiers
	accessIdentifiers.endpoint = server.URL
	hostedZones := accessIdentifiers.Zones()

	if len(hostedZones.HostedZone) != 2 {
		t.Fatal("Error reading remote hostedZones", hostedZones)
	}

}
