package route53

import (
  
)

// // changeResourceRecordSets
// type RecordSet struct {
//   AccessIdentifiers AccessIdentifiers
//   RecordSetsRequest RecordSetsRequest
//   Endpoint          string
// }

// type RecordSetsRequest struct {
//   XMLName     xml.Name
//   ChangeBatch ChangeBatch
// }

// type ChangeBatch struct {
//   Comment string
//   Changes []Change `xml:"Changes>Change"`
// }

// type Change struct {
//   Action            string
//   ResourceRecordSet ResourceRecordSet
// }

// type ResourceRecordSet struct {
//   Name            string
//   Type            string
//   TTL             int
//   ResourceRecords []ResourceRecord `xml:"ResourceRecords>ResourceRecord"`
//   HealthCheckId   string
// }

// type ResourceRecord struct {
//   Value string
// }

// func createResourceRecordSetsXML(resourceRecordSetsRequest RecordSetsRequest) (response string, err error) {
//   resourceRecordSetsRequest.XMLName = xml.Name{
//     Space: endpoint + `/doc/` + api + `/`,
//     Local: "ChangeResourceRecordSetsRequest",
//   }
//   byteXML, err := xml.MarshalIndent(resourceRecordSetsRequest, "", `   `)
//   if err != nil {
//     return "", err
//   }
//   response = xml.Header + string(byteXML)
//   return
// }
