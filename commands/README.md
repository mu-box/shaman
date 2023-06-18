[![shaman logo](http://microbox.rocks/assets/readme-headers/shaman.png)](http://microbox.cloud/open-source#shaman)
[![Build Status](https://github.com/mu-box/shaman/actions/workflows/ci.yaml/badge.svg)](https://github.com/mu-box/shaman/actions)

# Shaman

Small, lightweight, api-driven dns server.

## CLI Commands:

```
shaman - api driven dns server

Usage:
  shaman [flags]
  shaman [command]

Available Commands:
  add         Add a domain to shaman
  delete      Remove a domain from shaman
  list        List all domains in shaman
  get         Get records for a domain
  update      Update records for a domain
  reset       Reset all domains in shaman

Flags:
  -C, --api-crt string            Path to SSL crt for API access
  -k, --api-key string            Path to SSL key for API access
  -p, --api-key-password string   Password for SSL key
  -H, --api-listen string         Listen address for the API (ip:port) (default "127.0.0.1:1632")
  -c, --config-file string        Configuration file to load
  -O, --dns-listen string         Listen address for DNS requests (ip:port) (default "127.0.0.1:53")
  -d, --domain string             Parent domain for requests (default ".")
  -i, --insecure                  Disable tls key checking (client) and listen on http (api). Also disables auth-token
  -2, --l2-connect string         Connection string for the l2 cache (default "scribble:///var/db/shaman")
  -l, --log-level string          Log level to output [fatal|error|info|debug|trace] (default "INFO")
  -s, --server                    Run in server mode
  -t, --token string              Token for API Access (default "secret")
  -T, --ttl int                   Default TTL for DNS records (default 60)
  -v, --version                   Print version info and exit

Use "shaman [command] --help" for more information about a command.
```

## Server Usage Example:
```
$ shaman --server
```
or
```
$ shaman -c config.json
```

>config.json
>```json
{
  "api-crt": "",
  "api-key": "",
  "api-key-password": "",
  "api-listen": "127.0.0.1:1632",
  "token": "secret",
  "insecure": false,
  "l2-connect": "scribble:///var/db/shaman",
  "ttl": 60,
  "domain": ".",
  "dns-listen": "127.0.0.1:53",
  "log-level": "info",
  "server": true
}
```

## Client Usage Example:

#### add records

```sh
$ shaman -i add -d microbox.cloud -A 127.0.0.1
# {"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.1"}]}

$ shaman -i add -j '{"domain":"microbox.cloud","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"}]}'
# {"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"},{"ttl":60,"class":"IN","type":"A","address":"127.0.0.1"}]}
```

#### delete record

```sh
$ shaman -i delete -d microbox.cloud
# {"msg":"success"}
```

#### update record

```sh
$ shaman -i update -d microbox.cloud -A 127.0.0.2
# {"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"}]}
```

#### get record

```sh
$ shaman -i get -d microbox.cloud
# {"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.2"}]}
```

#### reset records

```sh
$ shaman -i reset -j '[{"domain":"microbox.cloud", "records":[{"address":"127.0.0.5"}]}]'
# [{"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.5"}]}]
```

#### list records

```sh
$ shaman -i list
# ["microbox.cloud"]

$ shaman -i list -f
# [{"domain":"microbox.cloud.","records":[{"ttl":60,"class":"IN","type":"A","address":"127.0.0.5"}]}]
```

[![oss logo](http://microbox.rocks/assets/open-src/microbox-open-src.png)](http://microbox.cloud/open-source)
