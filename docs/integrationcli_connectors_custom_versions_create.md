## integrationcli connectors custom versions create

Create a new custom connection version

### Synopsis

Create a new customer connection version in a region

```
integrationcli connectors custom versions create [flags]
```

### Examples

```
Create a custom connection version: integrationcli connectors custom versions create --id $version -n $name -f samples/custom-connection.json --sa=connectors --default-token
```

### Options

```
  -f, --file string   Custom Connection Version details JSON file path
  -h, --help          help for create
      --id string     Identifier assigned to the custom connection version
  -n, --name string   Connection name
      --sa string     Service Account name for the connection; do not include @<project-id>.iam.gserviceaccount.com
      --sp string     Service Account Project for the connection. Default is the connection's project id
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

* [integrationcli connectors custom versions](integrationcli_connectors_custom_versions.md)	 - Manage custom connection versions for Integration Connectors

###### Auto generated by spf13/cobra on 1-Apr-2025
