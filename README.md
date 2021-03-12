## APISIX api
> Debug version!!! Do not use this for produce env!!!

### Target
[x] Register route.     
[x] Register route by expired.     
[x] Register service.     

### Usage
```
package main

import (
    "github.com/eavesmy/gear"
    "github.com/eavesmy/apisix"
)

func main(){
    app := gear.New()
    
    svc := &apisix.Svc{
        Name: "test",
        Port:  "8000",
        XAPIKEY: "apisix api key.",
    }
    svc.RegisterService()
    svc.RegisterRouter("/*",50) // router and ttl
    svc.RegisterRouter("/*") //  just router 


    app.Listen(":8000")
}


```
