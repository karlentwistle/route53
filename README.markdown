# Route53 #

## A simple Google Go package for interacting with Route 53. ##

This was created to scratch my own itch and I use it in conjunction with my Google Go program https://github.com/karlentwistle/routemaster. This is not feature complete and is missing HealthChecks

Example Usage

Create A New Zone


    var zoneRequest = ZoneRequest{
      Name: "foo.bar.com", 
      CallerReference: "random tracking string", 
      HostedZoneConfig: HostedZoneConfig{
        Comment: "optional comment",
      },
    }

    var accessIdentifiers = AccessIdentifiers{
      AccessKey: "022QF06E7MXBSH9DHM02", 
      SecretKey: "kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct",
    }

    var hostedZone = HostedZone{
      AccessIdentifiers: accessIdentifiers, 
      HostedZoneRequest: zoneRequest, 
    }

    test, _ := hostedZone.CreateHostedZone()
    fmt.Println("Error")
    bodyBytes, _ := ioutil.ReadAll(test.Body) 
    bodyString := string(bodyBytes) 
    fmt.Println(bodyString)
