
Example Usage

Create A New Zone

var hostedZoneConfig = HostedZoneConfig{
  Comment: "optional comment",
}

var zoneRequest = ZoneRequest{
  Name: "foo.bar.com", 
  CallerReference: "random tracking string", 
  HostedZoneConfig: hostedZoneConfig,
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
