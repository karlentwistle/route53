package route53

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


const changeResourceRecordSets = `<?xml version="1.0" encoding="UTF-8"?>

<ChangeResourceRecordSetsRequest xmlns="https://route53.amazonaws.com/doc/2012-12-12/">
   <ChangeBatch>
      <Comment>optional comment about the changes in this change batch request</Comment>
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
               <HealthCheckId>optional ID of a Route 53 health check</HealthCheckId>
            </ResourceRecordSet>
         </Change>
      </Changes>
   </ChangeBatch>
</ChangeResourceRecordSetsRequest>`

var accessIdentifiers = AccessIdentifiers{AccessKey: "foo", SecretKey: "bar"}

type emptyHandler struct{}

func (h *emptyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Test-Header", "Testing")
	w.Header().Set("Date", time.Now().Format(time.RFC1123))
	io.WriteString(w, "hello, world!\n")
}

func TestSignature(t *testing.T) {
	response := accessIdentifiers.CreateSignature()
	if response != "TlXhIyeN+etmLEuVWrz4i+deyqGhDs5P/9wLYq6IyHE=" {
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

func TestAwsRequestHeaders(t *testing.T) {
	awsHeaders := accessIdentifiers.CreateHeaders()

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

var resourceRecordSets = RecordSetsRequest{
	ChangeBatch: ChangeBatch{
		Comment: "optional comment about the changes in this change batch request",
		Changes: []Change{
			{
				Action: "CREATE",
				ResourceRecordSet: ResourceRecordSet{
					Name: "DNS domain name",
					Type: "DNS record type",
					TTL:  300,
					ResourceRecords: []ResourceRecord{
						{
							Value: "applicable value for the record type",
						},
					},
					HealthCheckId: "optional ID of a Route 53 health check",
				},
			},
		},
	},
}

func TestCreateResourceRecordSetsXML(t *testing.T) {
	responseXML, err := createResourceRecordSetsXML(resourceRecordSets)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if string(responseXML) != changeResourceRecordSets {
		t.Fatal("returned XML is incorrectly formatted", responseXML)
	}
}

// TODO Add a test for remote post
func TestRemotePost(t *testing.T) {

}
