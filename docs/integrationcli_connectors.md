## integrationcli connectors

Manage connections for Integration Connectors

### Synopsis

Manage connections for Integration Connectors

### Options

```
  -h, --help          help for connectors
  -p, --proj string   Integration GCP Project name
  -r, --reg string    Integration region name
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
  -t, --token string     Google OAuth Token
      --verbose          Enable verbose output from integrationcli
```

### SEE ALSO

* [integrationcli](integrationcli.md)	 - Utility to work with GCP App Integration & Connectors
* [integrationcli connectors create](integrationcli_connectors_create.md)	 - Create a new connection
* [integrationcli connectors custom](integrationcli_connectors_custom.md)	 - Manage custom connections for Integration Connectors
* [integrationcli connectors delete](integrationcli_connectors_delete.md)	 - Delete a connection
* [integrationcli connectors eventsubs](integrationcli_connectors_eventsubs.md)	 - Manage Event Subscriptions for an Integration Connector
* [integrationcli connectors export](integrationcli_connectors_export.md)	 - Export connections in a region to a folder
* [integrationcli connectors get](integrationcli_connectors_get.md)	 - Get connection details
* [integrationcli connectors iam](integrationcli_connectors_iam.md)	 - Manage IAM permissions for the connection
* [integrationcli connectors import](integrationcli_connectors_import.md)	 - Import connections to a region from a folder
* [integrationcli connectors list](integrationcli_connectors_list.md)	 - List all connections in the region
* [integrationcli connectors managedzones](integrationcli_connectors_managedzones.md)	 - Manage DNS Peering with Integration Connectors
* [integrationcli connectors nodecount](integrationcli_connectors_nodecount.md)	 - Manage connection node count
* [integrationcli connectors operations](integrationcli_connectors_operations.md)	 - View connection operations
* [integrationcli connectors repair](integrationcli_connectors_repair.md)	 - Tries to repair eventing related event subscriptions
* [integrationcli connectors update](integrationcli_connectors_update.md)	 - Update an existing connection

###### Auto generated by spf13/cobra on 1-Apr-2025
