package main

import ( 
  "route53"
)

func main() {
  var accessIdentifiers = route53.AccessIdentifiers{
    AccessKey: "YourAWSAccessKey",
    SecretKey: "YourAWSSecretKey",
  }
  fmt.Println(accessIdentifiers.Zones())
}
