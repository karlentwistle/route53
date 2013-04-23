package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const createHostedZoneRequest = `<?xml version="1.0" encoding="UTF-8"?>

<CreateHostedZoneRequest xmlns="https://route53.amazonaws.com/doc/2012-12-12/">
   <Name>DNS domain name</Name>
   <CallerReference>unique description</CallerReference>
   <HostedZoneConfig>
      <Comment>optional comment</Comment>
   </HostedZoneConfig>
</CreateHostedZoneRequest>`

type emptyHandler struct{}

func (h *emptyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Test-Header", "Testing")
	w.Header().Set("Date", time.Now().Format(time.RFC1123))
	io.WriteString(w, "hello, world!\n")
}

var accessIdentifiers = AccessIdentifiers{AccessKey: "foo", SecretKey: "bar"}
var hostedZoneConfig  = HostedZoneConfig{Comment: "optional comment"}
var zoneRequest = ZoneRequest{
  Name: "DNS domain name", 
  CallerReference: "unique description", 
  HostedZoneConfig: hostedZoneConfig,
}
var hostedZone = HostedZone{
  AccessIdentifiers: accessIdentifiers, 
  HostedZoneRequest: zoneRequest, 
  Endpoint: "",
}

func TestSignature(t *testing.T) {
	response := signature("kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct", time.Unix(0, 0))
	if response != "wP04ISYdJymhU5Ix2tGl9kFU71ccwx2Nd1QEUtsONVI=" {
		t.Fatal("incorrect signature encoding", response)
	}
}

func TestRemoteHeaders(t *testing.T) {
	handler := &emptyHandler{}
	server := httptest.NewServer(handler)
	headers, err := remoteHeaders(server.URL)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if headers.Get("Test-Header") == "Test-Header" {
		t.Fatal("could not read headers", headers)
	}
}

func TestRemoteTime(t *testing.T) {
	handler := &emptyHandler{}
	server := httptest.NewServer(handler)
	remoteTime, err := remoteTime(server.URL)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if remoteTime != time.Now().Format(time.RFC1123) {
		t.Fatal("remote time did not return", remoteTime)
	}
}

func TestCreateHostedZoneXML(t *testing.T) {
	responseXML, err := createHostedZoneXML(zoneRequest)
  if err != nil {
    t.Fatal("Error:", err)
  }

	if string(responseXML) != createHostedZoneRequest {
		t.Fatal("returned XML is incorrectly formatted", responseXML)
	}
}

func TestRemotePost(t *testing.T) {
  
}
