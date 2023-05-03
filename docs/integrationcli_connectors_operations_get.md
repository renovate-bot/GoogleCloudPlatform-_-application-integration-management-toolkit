## integrationcli connectors operations get

Get operation details

### Synopsis

Get operation details from a connection created in a region

```
integrationcli connectors operations get [flags]
```

### Options

```
  -h, --help          help for get
  -n, --name string   The name of the operation
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --api api          Sets the control plane API. Must be one of prod, staging or autopush; default is prod
      --disable-check    Disable check for newer versions
      --no-output        Disable printing all statements to stdout
      --print-output     Control printing of info log statements (default true)
  -p, --proj string      Integration GCP Project name
  -r, --reg string       Integration region name
  -t, --token string     Google OAuth Token
      --verbose          Enable verbose output from integrationcli
```

### SEE ALSO

* [integrationcli connectors operations](integrationcli_connectors_operations.md)	 - View connection operations

###### Auto generated by spf13/cobra on 29-Apr-2023