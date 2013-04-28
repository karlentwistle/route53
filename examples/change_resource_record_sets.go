    package main

    import ( 
      "route53"
    )
    
    func main() {
      
      var resourceRecordSets = route53.ChangeResourceRecordSetsRequest{
        Comment: "optional comment",
        Changes: []Change{
          {
            Action:        "CREATE",
            Name:          "DNS domain name",
            Type:          "DNS record type",
            TTL:           300,
            Value:         "applicable value for the record type",
          },
        },
      }
      var accessIdentifiers = route53.AccessIdentifiers{
          AccessKey: "YourAWSAccessKey",
          SecretKey: "YourAWSSecretKey",
      }
      request, error := resourceRecordSets.Create(accessIdentifiers)
        
    }
