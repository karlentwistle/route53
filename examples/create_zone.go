package main

import ( 
  "route53"
)

func main() {

    var zoneRequest = route53.CreateHostedZoneRequest{
        Name: "foo.bar.org", 
        CallerReference: "random tracking string2", 
        Comment: "optional comment",
    }
    var accessIdentifiers = route53.AccessIdentifiers{
        AccessKey: "YourAWSAccessKey",
        SecretKey: "YourAWSSecretKey",
    }
    request, error := zoneRequest.Create(accessIdentifiers)
    
}
