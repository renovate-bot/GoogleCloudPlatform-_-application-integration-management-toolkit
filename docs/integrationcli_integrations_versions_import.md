## integrationcli integrations versions import

Import all versions of an Integration flows to a region

### Synopsis

Import all versions of an Integration flows to a region

```
integrationcli integrations versions import [flags]
```

### Options

```
  -f, --folder string   Folder to import Integration flows
  -h, --help            help for import
  -n, --name string     Name of the Integration flow
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

* [integrationcli integrations versions](integrationcli_integrations_versions.md)	 - Manage integrations flow versions

###### Auto generated by spf13/cobra on 1-Apr-2025
