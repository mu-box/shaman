[![shaman logo](http://microbox.rocks/assets/readme-headers/shaman.png)](http://microbox.cloud/open-source#shaman)
[![Build Status](https://github.com/mu-box/shaman/actions/workflows/ci.yaml/badge.svg)](https://github.com/mu-box/shaman/actions)

# Shaman

Small, lightweight, api-driven dns server.

## Routes:

| Route | Description | Payload | Output |
| --- | --- | --- | --- |
| **POST** /records | Adds the domain and full record | json domain object | json domain object |
| **PUT** /records | Update all domains and records (replaces all) | json array of domain objects | json array of domain objects |
| **GET** /records | Returns a list of domains we have records for | nil | string array of domains |
| **PUT** /records/{domain} | Update domain's records (replaces all) | json domain object | json domain object |
| **GET** /records/{domain} | Returns the records for that domain | nil | json domain object |
| **DELETE** /records/{domain} | Delete a domain | nil | success message |

## Usage Example:

#### add domain
```sh
$ curl -k -H "X-AUTH-TOKEN: secret" https://localhost:1632/records -d \
       '{"domain":"microbox.cloud","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"}]}'
# {"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"}]}
```

#### list domains
```sh
$ curl -k -H "X-AUTH-TOKEN: secret" https://localhost:1632/records
# ["microbox.cloud"]
```
or add `?full=true` for the full records
```sh
$ curl -k -H "X-AUTH-TOKEN: secret" https://localhost:1632/records?full=true
# [{"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"}]}]
```

#### update domains
```sh
$ curl -k -H "X-AUTH-TOKEN: secret" https://localhost:1632/records -d \
       '[{"domain":"microbox.cloud","records":[{"address":"127.0.0.1"}]}]' \
       -X PUT
# [{"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.1"}]}]
```

#### update domain
```sh
$ curl -k -H "X-AUTH-TOKEN: secret" https://localhost:1632/records/microbox.cloud -d \
       '{"domain":"microbox.cloud","records":[{"address":"127.0.0.2"}]}' \
       -X PUT
# {"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"}]}
```

#### delete domain
```sh
$ curl -k -H "X-AUTH-TOKEN: secret" https://localhost:1632/records/microbox.cloud \
       -X DELETE
# {"msg":"success"}
```

#### get domain
```sh
$ curl -k -H "X-AUTH-TOKEN: secret" https://localhost:1632/records/microbox.cloud
# {"err":"failed to find record for domain - 'microbox.cloud'"}
```

[![oss logo](http://microbox.rocks/assets/open-src/microbox-open-src.png)](http://microbox.cloud/open-source)
