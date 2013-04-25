package hostedzone

import (
  "testing"
)

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
  Comment: "optional comment",
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
