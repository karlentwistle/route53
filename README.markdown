# Route53 #

## A simple Google Go package for interacting with Route 53. ##

This was created to scratch my own itch and I use it in conjunction with my Google Go program https://github.com/karlentwistle/routemaster. This is not feature complete and is missing HealthChecks

Example Usage

Create A New Zone

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

Change Resource Record Sets (You can set the action to CREATE or DELETE)

    package main

    import ( 
      "route53"
    )
    
    func main() {
      
      var resourceRecordSets = ChangeResourceRecordSetsRequest{
        Comment: "optional comment",
        Changes: []Change{
          {
            Action:        "CREATE",
            Name:          "DNS domain name",
            Type:          "DNS record type",
            TTL:           300,
            Value:         "applicable value for the record type",
            HealthCheckId: "optional ID of a Route 53 health check",
          },
        },
      }
      var accessIdentifiers = route53.AccessIdentifiers{
          AccessKey: "YourAWSAccessKey",
          SecretKey: "YourAWSSecretKey",
      }
      request, error := resourceRecordSets.Create(accessIdentifiers)
        
    }
