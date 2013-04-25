package route53

import (

)

// var resourceRecordSets = RecordSetsRequest{
//   ChangeBatch: ChangeBatch{
//     Comment: "optional comment about the changes in this change batch request",
//     Changes: []Change{
//       {
//         Action: "CREATE",
//         ResourceRecordSet: ResourceRecordSet{
//           Name: "DNS domain name",
//           Type: "DNS record type",
//           TTL:  300,
//           ResourceRecords: []ResourceRecord{
//             {
//               Value: "applicable value for the record type",
//             },
//           },
//           HealthCheckId: "optional ID of a Route 53 health check",
//         },
//       },
//     },
//   },
// }

// func TestCreateResourceRecordSetsXML(t *testing.T) {
//   responseXML, err := createResourceRecordSetsXML(resourceRecordSets)
//   if err != nil {
//     t.Fatal("Error:", err)
//   }

//   if string(responseXML) != changeResourceRecordSets {
//     t.Fatal("returned XML is incorrectly formatted", responseXML)
//   }
// }

// // TODO Add a test for remote post
// func TestRemotePost(t *testing.T) {

// }
