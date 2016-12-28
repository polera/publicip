#publicip

This package returns the public facing IP address of the calling client (a la https://icanhazip.com, but from Go!)

Author
==
James Polera <james@uncryptic.com>

Dependencies
==
publicip uses Glide for dependency management.  After cloning this package, run:
```bash
glide up
```

Credits
==
This package was inspired by both:

[public-ip (nodejs)](https://github.com/sindresorhus/public-ip/blob/master/index.js)

[OpenDNS::MyIP (Perl)](https://metacpan.org/pod/OpenDNS::MyIP)

Example
==
```go
package main

import (
  "fmt"
  "github.com/polera/publicip"
)

func main() {

  myIpAddr, err := publicip.GetIP()
  if err != nil {
    fmt.Printf("Error getting IP address: %s\n", err)
  } else {
    fmt.Printf("Public IP address is: %s", myIpAddr)
  }

}

```


