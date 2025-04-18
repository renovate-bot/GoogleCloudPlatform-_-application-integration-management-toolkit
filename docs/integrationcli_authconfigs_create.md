## integrationcli authconfigs create

Create an authconfig

### Synopsis

Create an authconfig

```
integrationcli authconfigs create [flags]
```

### Examples

```
Create a new user name auth config: integrationcli authconfigs create -f samples/ac_username.json
Create a new OIDC auth config: integrationcli authconfigs create -f samples/ac_oidc.json
Create a new auth token auth config: integrationcli authconfigs create -f samples/ac_authtoken.json
Create a new auth config from Cloud KMS Encrypted files: integrationcli authconfigs create -e samples/b64encoded_ac.txt -k locations/$region/keyRings/$key/cryptoKeys/$cryptokey
```

### Options

```
  -e, --encrypted-file string     Base64 encoded, Cloud KMS encrypted Auth Config JSON file path
  -k, --encryption-keyid string   Cloud KMS key for decrypting Auth Config; Format = locations/*keyRings/*/cryptoKeys/*
  -f, --file string               Auth Config JSON file path
  -h, --help                      help for create
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --api api          Sets the control plane API. Must be one of prod, staging or autopush; default is prod
      --default-token    Use Google default application credentials access token
      --disable-check    Disable check for newer versions
      --metadata-token   Metadata OAuth2 access token
      --no-output        Disable printing all statements to stdout
      --print-output     Control printing of info log statements (default true)
  -p, --proj string      Integration GCP Project name
  -r, --reg string       Integration region name
  -t, --token string     Google OAuth Token
      --verbose          Enable verbose output from integrationcli
```

### SEE ALSO

* [integrationcli authconfigs](integrationcli_authconfigs.md)	 - Manage integration auth configurations

###### Auto generated by spf13/cobra on 1-Apr-2025
