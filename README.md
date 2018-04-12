# rnv api
This project contains a golang implementation for integrating the *rnv Start.Info API*. __To use the api, you need an [access token](https://opendata.rnv-online.de/startinfo-api)!__

## How to get this package
Use go flow to get this module:
```
go get github.com/Rhein-Neckar-Verkehr/golang-rnv-api
```

### Set Access Token
You can set the token as environment variable:
```
export rnv_api_token=yourToken // Unix example
```
Or use a config file:
* Copy the *tokenConfExample.json* file,
* add your *RNV_API_TOKEN* in the *token* field and
* rename this file to *tokenConf.json*.

### Proxy settings
This module uses the environment variable ```https_proxy``` for making request through a proxy. __NTLM-Authentication is not supported__

## How to use this package
Run tests within the package:
```
go test
```
The duration of a test run takes approximately 1.5 minutes.


Full Documentation is in the ```doc/``` folder in [docco](github.com/dhconnelly/docgo) style.
This module contains examples for each interface of the *rnv Start.Info API* (see ```doc/apiExample.html```).
For the correct use of query parameter for each interface, see the testable examples in ```doc/apiParamsExample_test.html```.
