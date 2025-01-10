# `cinder`命令

``` bash
cinder [flage...] <subcommand> ...

Command-line interface to the OpenStack Cinder API.

Positional arguments:
  <subcommand>
    absolute-limits     Lists absolute limits for a user.
    api-version         Display the server API version information. (Supported
                        by API versions 3.0 - 3.latest)
    attachment-complete
                        Complete an attachment for a cinder volume. (Supported
                        by API versions 3.44 - 3.latest)
    attachment-create   Create an attachment for a cinder volume. (Supported
                        by API versions 3.27 - 3.latest)
    attachment-delete   Delete an attachment for a cinder volume. (Supported
                        by API versions 3.27 - 3.latest)
    attachment-list     Lists all attachments. (Supported by API versions 3.27
                        - 3.latest)
    attachment-show     Show detailed information for attachment. (Supported
                        by API versions 3.27 - 3.latest)
    attachment-update   Update an attachment for a cinder volume. (Supported
                        by API versions 3.27 - 3.latest)
    availability-zone-list
                        Lists all availability zones.
    backup-abort        Aborts one or more backups.
    backup-create       Creates a volume backup.
    backup-delete       Removes one or more backups.
    backup-export       Export backup metadata record.
    backup-import       Import backup metadata record.
    backup-list         Lists all backups.
    backup-reset-state  Explicitly updates the backup state.
    backup-restore      Restores a backup.
    backup-show         Shows backup details.
    backup-update       Updates a backup. (Supported by API versions 3.9 -
                        3.latest)
    calc-backup-hmac    Calculate backup hmac.
    calc-snapshot-hmac  Calculate snapshot hmac.
    calc-volume-hmac    Calculate volume hmac.
    cgsnapshot-create   Creates a cgsnapshot.
    cgsnapshot-delete   Removes one or more cgsnapshots.
    cgsnapshot-list     Lists all cgsnapshots.
    cgsnapshot-show     Shows cgsnapshot details.
    clean-reserved-time
                        Clean volume reserved time.
    clone-image-metadata
                        Clone image metadata.
    cluster-disable     Disables clustered services. (Supported by API
                        versions 3.7 - 3.latest)
    cluster-enable      Enables clustered services. (Supported by API versions
                        3.7 - 3.latest)
    cluster-list        Lists clustered services with optional filtering.
                        (Supported by API versions 3.7 - 3.latest)
    cluster-show        Show detailed information on a clustered service.
                        (Supported by API versions 3.7 - 3.latest)
    config-metadata     Sets or unsets or updates config metadata.
    consisgroup-create  Creates a consistency group.
    consisgroup-create-from-src
                        Creates a consistency group from a cgsnapshot or a
                        source CG.
    consisgroup-delete  Removes one or more consistency groups.
    consisgroup-list    Lists all consistency groups.
    consisgroup-show    Shows details of a consistency group.
    consisgroup-update  Updates a consistency group.
    create              Creates a volume.
    credentials         Shows user credentials returned from auth.
    delete              Removes one or more volumes.
    delete-cancel       Cancel one or more deletions.
    delete-confirm      Confirm one or more deletions.
    delete-volume-hmac  Delete volume hmac.
    encryption-type-create
                        Creates encryption type for a volume type. Admin only.
    encryption-type-delete
                        Deletes encryption type for a volume type. Admin only.
    encryption-type-list
                        Shows encryption type details for volume types. Admin
                        only.
    encryption-type-show
                        Shows encryption type details for a volume type. Admin
                        only.
    encryption-type-update
                        Update encryption type information for a volume type
                        (Admin Only).
    endpoints           Discovers endpoints registered by authentication
                        service.
    extend              Attempts to extend size of an existing volume.
    extra-specs-list    Lists current volume types and extra specs.
    failover-host       Failover a replicating cinder-volume host.
    force-delete        Attempts force-delete of volume, regardless of state.
    force-delete-system-reserved
                        Force delete certain system_reserved volume(s).
    freeze-host         Freeze and disable the specified cinder-volume host.
    get-capabilities    Show backend volume stats and properties. Admin only.
    get-cinder-version
    get-delete-confirm-resource
                        Show a delete confirm resource by verify_id.
    get-pools           Show pool information for backends. Admin only.
    get-qos             Shows volume qos.
    get-reserved-time   Get volume reserved time.
    get-system-reserved-volumes
                        Get system reserved (user deleted) volumes with
                        reserved time.
    get-volume-snapshots-used
                        Shows volume qos.
    get-volumes-with-reserved-time
                        Get volumes with reserved time.
    group-create        Creates a group. (Supported by API versions 3.13 -
                        3.latest)
    group-create-from-src
                        Creates a group from a group snapshot or a source
                        group. (Supported by API versions 3.14 - 3.latest)
    group-delete        Removes one or more groups. (Supported by API versions
                        3.13 - 3.latest)
    group-disable-replication
                        Disables replication for group. (Supported by API
                        versions 3.38 - 3.latest)
    group-enable-replication
                        Enables replication for group. (Supported by API
                        versions 3.38 - 3.latest)
    group-failover-replication
                        Fails over replication for group. (Supported by API
                        versions 3.38 - 3.latest)
    group-list          Lists all groups. (Supported by API versions 3.13 -
                        3.latest)
    group-list-replication-targets
                        Lists replication targets for group. (Supported by API
                        versions 3.38 - 3.latest)
    group-show          Shows details of a group. (Supported by API versions
                        3.13 - 3.latest)
    group-snapshot-create
                        Creates a group snapshot. (Supported by API versions
                        3.14 - 3.latest)
    group-snapshot-delete
                        Removes one or more group snapshots. (Supported by API
                        versions 3.14 - 3.latest)
    group-snapshot-list
                        Lists all group snapshots.
    group-snapshot-show
                        Shows group snapshot details. (Supported by API
                        versions 3.14 - 3.latest)
    group-specs-list    Lists current group types and specs. (Supported by API
                        versions 3.11 - 3.latest)
    group-type-create   Creates a group type. (Supported by API versions 3.11
                        - 3.latest)
    group-type-default  List the default group type. (Supported by API
                        versions 3.11 - 3.latest)
    group-type-delete   Deletes group type or types. (Supported by API
                        versions 3.11 - 3.latest)
    group-type-key      Sets or unsets group_spec for a group type. (Supported
                        by API versions 3.11 - 3.latest)
    group-type-list     Lists available 'group types'. (Admin only will see
                        private types) (Supported by API versions 3.11 -
                        3.latest)
    group-type-show     Show group type details. (Supported by API versions
                        3.11 - 3.latest)
    group-type-update   Updates group type name, description, and/or
                        is_public. (Supported by API versions 3.11 - 3.latest)
    group-update        Updates a group. (Supported by API versions 3.13 -
                        3.latest)
    image-metadata      Sets or deletes volume image metadata.
    image-metadata-show
                        Shows volume image metadata.
    interrupt-io        Interrupt or recover IO operations on volumes.
    list                Lists all volumes.
    list-delete-confirm-resources
                        Get delete confirm resources.
    list-filters        List enabled filters. (Supported by API versions 3.33
                        - 3.latest)
    live-retype         Changes the volume type for a volume.
    manage              Manage an existing volume.
    manageable-list     Lists all manageable volumes. (Supported by API
                        versions 3.8 - 3.latest)
    message-delete      Removes one or more messages. (Supported by API
                        versions 3.3 - 3.latest)
    message-list        Lists all messages. (Supported by API versions 3.3 -
                        3.latest)
    message-show        Shows message details. (Supported by API versions 3.3
                        - 3.latest)
    metadata            Sets or deletes volume metadata.
    metadata-show       Shows volume metadata.
    metadata-update-all
                        Updates volume metadata.
    migrate             Migrates volume to a new host.
    qos-associate       Associates qos specs with specified volume type.
    qos-create          Creates a qos specs.
    qos-delete          Deletes a specified qos specs.
    qos-disassociate    Disassociates qos specs from specified volume type.
    qos-disassociate-all
                        Disassociates qos specs from all its associations.
    qos-get-association
                        Lists all associations for specified qos specs.
    qos-key             Sets or unsets specifications for a qos spec.
    qos-list            Lists qos specs.
    qos-show            Shows qos specs details.
    quota-class-show    Lists quotas for a quota class.
    quota-class-update  Updates quotas for a quota class.
    quota-defaults      Lists default quotas for a tenant.
    quota-delete        Delete the quotas for a tenant.
    quota-show          Lists quotas for a tenant.
    quota-update        Updates quotas for a tenant.
    quota-usage         Lists quota usage for a tenant.
    rate-limits         Lists rate limits for a user.
    readonly-mode-update
                        Updates volume read-only access-mode flag.
    rename              Renames a volume.
    replication-promote
                        Promote a secondary volume to primary for a
                        relationship.
    replication-reenable
                        Sync the secondary volume with primary for a
                        relationship.
    reset-state         Explicitly updates the entity state in the Cinder
                        database.
    retype              Changes the volume type for a volume.
    revert-to-snapshot  Revert a volume to the specified snapshot. (Supported
                        by API versions 3.40 - 3.latest)
    service-disable     Disables the service.
    service-enable      Enables the service.
    service-get-log     (Supported by API versions 3.32 - 3.latest)
    service-list        Lists all services. Filter by host and service binary.
                        (Supported by API versions 3.0 - 3.latest)
    service-set-log     (Supported by API versions 3.32 - 3.latest)
    set-bootable        Update bootable status of a volume.
    set-qos             Set qos for volume.
    set-reserved-time   Set volume reserved time.
    show                Shows volume details.
    snapshot-cg-create  Creates a cgsnapshot.
    snapshot-cg-delete  Removes one or more cgsnapshots.
    snapshot-cg-remove  Removes one or more snapshots from a consistency
                        snapshot group.
    snapshot-cg-restore
                        Restores a single cgsnapshot.
    snapshot-cg-show    Shows cgsnapshot details.
    snapshot-create     Creates a snapshot.
    snapshot-delete     Removes one or more snapshots.
    snapshot-list       Lists all snapshots.
    snapshot-manage     Manage an existing snapshot.
    snapshot-manageable-list
                        Lists all manageable snapshots. (Supported by API
                        versions 3.8 - 3.latest)
    snapshot-metadata   Sets or deletes snapshot metadata.
    snapshot-metadata-show
                        Shows snapshot metadata.
    snapshot-metadata-update-all
                        Updates snapshot metadata.
    snapshot-rename     Renames a snapshot.
    snapshot-reset-state
                        Explicitly updates the snapshot state.
    snapshot-show       Shows snapshot details.
    snapshot-unmanage   Stop managing a snapshot.
    summary             Get volumes summary. (Supported by API versions 3.12 -
                        3.latest)
    system-restore      Restore a system_reserved volume.
    thaw-host           Thaw and enable the specified cinder-volume host.
    transfer-accept     Accepts a volume transfer.
    transfer-create     Creates a volume transfer.
    transfer-delete     Undoes a transfer.
    transfer-list       Lists all transfers.
    transfer-show       Shows transfer details.
    type-access-add     Adds volume type access for the given project.
    type-access-list    Print access information about the given volume type.
    type-access-remove  Removes volume type access for the given project.
    type-create         Creates a volume type.
    type-default        List the default volume type.
    type-delete         Deletes volume type or types.
    type-key            Sets or unsets extra_spec for a volume type.
    type-list           Lists available 'volume types'.
    type-show           Show volume type details.
    type-update         Updates volume type name, description, and/or
                        is_public.
    unmanage            Stop managing a volume.
    update-volume-type  Changes the volume type for a volume to diff qos in
                        same pool.
    upload-to-image     Uploads volume to Image Service as an image.
    veri-backup-hmac    Verify snapshot hmac.
    veri-snapshot-hmac  Verify snapshot hmac.
    veri-volume-hmac    Verify volume hmac.
    version-list        List all API versions. (Supported by API versions 3.0
                        - 3.latest)
    volume-lock         Lock a volume, so the volume cannot be deleted until
                        it's unlocked.
    volume-unlock       UnLock a volume, so the volume can be deleted.
    work-cleanup        Request cleanup of services with optional filtering.
                        (Supported by API versions 3.24 - 3.latest)
    bash-completion     Prints arguments for bash_completion.
    help                Shows help about this program or one of its
                        subcommands.
    list-extensions     Lists all available os-api extensions.

Optional arguments:
  --version             show program's version number and exit
  -d, --debug           Shows debugging output.
  --service-type <service-type>
                        Service type. For most actions, default is volume.
  --service-name <service-name>
                        Service name. Default=env[CINDER_SERVICE_NAME].
  --volume-service-name <volume-service-name>
                        Volume service name.
                        Default=env[CINDER_VOLUME_SERVICE_NAME].
  --os-endpoint-type <os-endpoint-type>
                        Endpoint type, which is publicURL or internalURL.
                        Default=env[OS_ENDPOINT_TYPE] or nova
                        env[CINDER_ENDPOINT_TYPE] or publicURL.
  --endpoint-type <endpoint-type>
                        DEPRECATED! Use --os-endpoint-type.
  --os-volume-api-version <volume-api-ver>
                        Block Storage API version. Accepts X, X.Y (where X is
                        major and Y is minor
                        part).Default=env[OS_VOLUME_API_VERSION].
  --bypass-url <bypass-url>
                        DEPRECATED! Use os_endpoint. Use this API endpoint
                        instead of the Service Catalog. Defaults to
                        env[CINDERCLIENT_BYPASS_URL].
  --os-endpoint <os-endpoint>
                        Use this API endpoint instead of the Service Catalog.
                        Defaults to env[CINDER_ENDPOINT].
  --retries <retries>   Number of retries.
  --profile HMAC_KEY    HMAC key to use for encrypting context data for
                        performance profiling of operation. This key needs to
                        match the one configured on the cinder api server.
                        Without key the profiling will not be triggered even
                        if osprofiler is enabled on server side. Defaults to
                        env[OS_PROFILE].
  --os-auth-type <name>, --os-auth-plugin <name>
                        Authentication type to use
  --os-auth-strategy <auth-strategy>
                        Authentication strategy (Env: OS_AUTH_STRATEGY,
                        default keystone). For now, any other value will
                        disable the authentication.
  --os-auth-system <os-auth-system>
                        DEPRECATED! Use --os-auth-type. Defaults to
                        env[OS_AUTH_SYSTEM].
  --os-token <token>    Defaults to env[OS_TOKEN].
  --os-url <url>        Defaults to env[OS_URL].
  --os-delete-token <delete-token>
                        Delete token for admin role, default=(Env:
                        OS_DELETE_TOKEN)

API Connection Options:
  Options controlling the HTTP API Connections

  --insecure            Explicitly allow client to perform "insecure" TLS
                        (https) requests. The server's certificate will not be
                        verified against any certificate authorities. This
                        option should be used with caution.
  --os-cacert <ca-certificate>
                        Specify a CA bundle file to use in verifying a TLS
                        (https) server certificate. Defaults to
                        env[OS_CACERT].
  --os-cert <certificate>
                        Defaults to env[OS_CERT].
  --os-key <key>        Defaults to env[OS_KEY].
  --timeout <seconds>   Set request timeout (in seconds).

Authentication Options:
  Options specific to the password plugin.

  --os-auth-url OS_AUTH_URL
                        Authentication URL
  --os-system-scope OS_SYSTEM_SCOPE
                        Scope for system operations
  --os-domain-id OS_DOMAIN_ID
                        Domain ID to scope to
  --os-domain-name OS_DOMAIN_NAME
                        Domain name to scope to
  --os-project-id OS_PROJECT_ID, --os-tenant-id OS_PROJECT_ID
                        Project ID to scope to
  --os-project-name OS_PROJECT_NAME, --os-tenant-name OS_PROJECT_NAME
                        Project name to scope to
  --os-project-domain-id OS_PROJECT_DOMAIN_ID
                        Domain ID containing project
  --os-project-domain-name OS_PROJECT_DOMAIN_NAME
                        Domain name containing project
  --os-trust-id OS_TRUST_ID
                        Trust ID
  --os-default-domain-id OS_DEFAULT_DOMAIN_ID
                        Optional domain ID to use with v3 and v2 parameters.
                        It will be used for both the user and project domain
                        in v3 and ignored in v2 authentication.
  --os-default-domain-name OS_DEFAULT_DOMAIN_NAME
                        Optional domain name to use with v3 API and v2
                        parameters. It will be used for both the user and
                        project domain in v3 and ignored in v2 authentication.
  --os-user-id OS_USER_ID
                        User id
  --os-username OS_USERNAME, --os-user-name OS_USERNAME
                        Username
  --os-user-domain-id OS_USER_DOMAIN_ID
                        User's domain id
  --os-user-domain-name OS_USER_DOMAIN_NAME
                        User's domain name
  --os-password OS_PASSWORD
                        User's password

Run "cinder help SUBCOMMAND" for help on a subcommand.
```

## create  

```bash
cinder create --name janus_a --image-id 939d8d9f-9c40-4369-8de6-4b701be38799 50 
939d8d9f-9c40-4369-8de6-4b701be38799
```

```bash
usage: cinder create [flags] [size]

Creates a volume.

Positional arguments:
  <size>                Size of volume, in GiBs. (Required unless snapshot-id
                        /source-volid is specified).

Optional arguments:
  --consisgroup-id <consistencygroup-id>
                        ID of a consistency group where the new volume belongs
                        to. Default=None.
  --snapshot-id <snapshot-id>
                        Creates volume from snapshot ID. Default=None.
  --source-volid <source-volid>
                        Creates volume from volume ID. Default=None.
  --source-replica <source-replica>
                        Creates volume from replicated volume ID.
                        Default=None.
  --image-id <image-id>
                        Creates volume from image ID. Default=None.
  --image <image>       Creates a volume from image (ID or name).
                        Default=None.
  --name <name>         Volume name. Default=None.
  --description <description>
                        Volume description. Default=None.
  --volume-type <volume-type>
                        Volume type. Default=None.
  --availability-zone <availability-zone>
                        Availability zone for volume. Default=None.
  --metadata [<key=value> [<key=value> ...]]
                        Metadata key and value pairs. Default=None.
  --hint <key=value>    Scheduler hint, like in nova.
  --allow-multiattach   Allow volume to be attached more than once.
                        Default=False
  --poll                Wait for volume creation until it completes.
  --count <number>      Create <number> volumes (limited by quota).
  --id <id>             Creates a volume with specific id. Default=None.
```

## list

```bash
cinder list [flag...]

Optional arguments:
  --all-tenants [<0|1>]
                        Shows details for all tenants. Admin only.
  --name <name>         Filters results by a name. Default=None. This option
                        is deprecated and will be removed in newer release.
                        Please use '--filters' option which is introduced
                        since 3.33 instead.
  --status <status>     Filters results by a status. Default=None. This option
                        is deprecated and will be removed in newer release.
                        Please use '--filters' option which is introduced
                        since 3.33 instead.
  --bootable [<True|true|False|false>]
                        Filters results by bootable status. Default=None. This
                        option is deprecated and will be removed in newer
                        release. Please use '--filters' option which is
                        introduced since 3.33 instead.
  --migration_status <migration_status>
                        Filters results by a migration status. Default=None.
                        Admin only. This option is deprecated and will be
                        removed in newer release. Please use '--filters'
                        option which is introduced since 3.33 instead.
  --metadata [<key=value> [<key=value> ...]]
                        Filters results by a metadata key and value pair.
                        Default=None. This option is deprecated and will be
                        removed in newer release. Please use '--filters'
                        option which is introduced since 3.33 instead.
  --marker <marker>     Begin returning volumes that appear later in the
                        volume list than that represented by this volume id.
                        Default=None.
  --limit <limit>       Maximum number of volumes to return. Default=None.
  --fields <fields>     Comma-separated list of fields to display. Use the
                        show command to see which fields are available.
                        Unavailable/non-existent fields will be ignored.
                        Default=None.
  --sort <key>[:<direction>]
                        Comma-separated list of sort keys and directions in
                        the form of <key>[:<asc|desc>]. Valid keys: id,
                        status, size, availability_zone, name, bootable,
                        created_at, reference. Default=None.
  --tenant [<tenant>]   Display information from single tenant (Admin only).
```

## delete

```bash
cinder delete [flags] <volume> [<volume> ...]	# 删除volume

Positional arguments:
  <volume>        Name or ID of volume or volumes to delete.

Optional arguments:
  --cascade       Remove any snapshots along with volume. Default=False.
  --remove-in-db  Remove deleted record in db at the same time. Default=False.
```

## rename

```bash
cinder rename [flag] <volume> [<name>]

Positional arguments:
  <volume>              Name or ID of volume to rename.
  <name>                New name for volume.

Optional arguments:
  --description <description>
                        Volume description. Default=None.
```

## attachment-list

```bash
cinder --os-volume-api-version 3.50  attachment-list   --status

cinder --os-volume-api-version 3.50  attachment-list  --volume-id 0e07009a-5066-4c9e-aa8a-ca0e631b6329
```

```bash
usage: cinder attachment-list [--all-tenants [<0|1>]]
                              [--volume-id <volume-id>] [--status <status>]
                              [--marker <marker>] [--limit <limit>]
                              [--sort <key>[:<direction>]]
                              [--tenant [<tenant>]]
                              [--filters [<key=value> [<key=value> ...]]]

Lists all attachments.

Optional arguments:
  --all-tenants [<0|1>]
                        Shows details for all tenants. Admin only.
  --volume-id <volume-id>
                        Filters results by a volume ID. Default=None. This
                        option is deprecated and will be removed in newer
                        release. Please use '--filters' option which is
                        introduced since 3.33 instead.
  --status <status>     Filters results by a status. Default=None. This option
                        is deprecated and will be removed in newer release.
                        Please use '--filters' option which is introduced
                        since 3.33 instead.
  --marker <marker>     Begin returning attachments that appear later in
                        attachment list than that represented by this id.
                        Default=None.
  --limit <limit>       Maximum number of attachments to return. Default=None.
  --sort <key>[:<direction>]
                        Comma-separated list of sort keys and directions in
                        the form of <key>[:<asc|desc>]. Valid keys: id,
                        status, size, availability_zone, name, bootable,
                        created_at, reference. Default=None.
  --tenant [<tenant>]   Display information from single tenant (Admin only).
  --filters [<key=value> [<key=value> ...]]
                        Filter key and value pairs. Please use 'cinder list-
                        filters' to check enabled filters from server. Use
                        'key~=value' for inexact filtering if the key
                        supports. Default=None. (Supported by API version 3.33
                        and later)
```



## reset-state

```bash
cinder reset-state [flag...] <entity> [<entity> ...]	不会改变实际状态

Positional arguments:
  <entity>              Name or ID of entity to update.

Optional arguments:
  --type <type>         Type of entity to update. Available resources are:
                        'volume', 'snapshot', 'backup', 'group' (since 3.20)
                        and 'group-snapshot' (since 3.19), Default=volume.
                        
  --state <available|error>       
  --attach-status <attach-status>
                        This is only used for a volume entity. The attach
                        status to assign to the volume in the database, with
                        no regard to the actual status. Valid values are
                        "attached" and "detached". Default=None, that means
                        the status is unchanged.
  --reset-migration-status
                        This is only used for a volume entity. Clears the
                        migration status of the volume in the DataBase that
                        indicates the volume is source or destination of
                        volume migration, with no regard to the actual status
```

## backup-show

```bash
cinder backup-show <backup>

Positional arguments:
  <backup>  Name or ID of backup.
```

## backup state

```bash
usage: cinder backup-reset-state [--state <state>] <backup> [<backup> ...]

Explicitly updates the backup state.

Positional arguments:
  <backup>         Name or ID of the backup to modify.

Optional arguments:
  --state <state>  The state to assign to the backup. Valid values are
                   "available", "error". Default=available.

```



## backup-delete

```bash
cinder backup-delete [--force] <backup> [<backup> ...]

Removes one or more backups.

Positional arguments:
  <backup>  Name or ID of backup(s) to delete.

Optional arguments:
  --force   Allows deleting backup of a volume when its status is other than
            "available" or "error". Default=False.
```

## backup-reset-state

```bash
cinder backup-reset-state [flag] <backup>... 

Explicitly updates the backup state.

Positional arguments:
  <backup>         Name or ID of the backup to modify.

Optional arguments:
  --state <state>  The state to assign to the backup. Valid values are
                   "available", "error". Default=available.

```

## backup-restore

```bash
cinder backup-restore [--volume <volume>] [--force-host]
                             [--name <name>]
                             <backup>

Restores a backup.

Positional arguments:
  <backup>           Name or ID of backup to restore.

Optional arguments:
  --volume <volume>  Name or ID of existing volume to which to restore. This
                     is mutually exclusive with --name and takes priority.
                     Default=None.
  --force-host       Use volume host first. Default=False.
  --name <name>      Use the name for new volume creation to restore. This is
                     mutually exclusive with --volume (or the deprecated
                     --volume-id) and --volume (or --volume-id) takes
                     priority. Default=None.
```



## attachment-create 

```bash
cinder --os-volume-api-version 3.50 attachment-create janus ae5cebfc-630d-43db-bf34-8847afbc926c
```

```bash
usage: cinder attachment-create [flag...] <volume> <server_id>

Create an attachment for a cinder volume.

Positional arguments:
  <volume>              Name or ID of volume or volumes to attach.
  <server_id>           ID of server attaching to.

Optional arguments:
  --connect <connect>   Make an active connection using provided connector
                        info (True or False).
  --initiator <initiator>
                        iqn of the initiator attaching to. Default=None.
  --ip <ip>             ip of the system attaching to. Default=None.
  --host <host>         Name of the host attaching to. Default=None.
  --platform <platform>
                        Platform type. Default=x86_64.
  --ostype <ostype>     OS type. Default=linux2.
  --multipath <multipath>
                        Use multipath. Default=False.
  --mountpoint <mountpoint>
                        Mountpoint volume will be attached at. Default=None.
```

## attachment-complete

```bash
cinder --os-volume-api-version 3.50 attachment-complete 1342ecd8-547d-4580-a8e8-57cff8c51d19
```

```bash
cinder attachment-complete <attachment> [<attachment> ...]

Complete an attachment for a cinder volume.

Positional arguments:
  <attachment>  ID of attachment or attachments to delete.
```

## attachment-delete

```bash
cinder --os-volume-api-version 3.50 attachment-delete
```

```bash
 cinder attachment-delete <attachment> [<attachment> ...]

Delete an attachment for a cinder volume.

Positional arguments:
  <attachment>  ID of attachment or attachments to delete.
```

## attachment-show

```bash
cinder --os-volume-api-version 3.50 attachment-show 
```

## attachment

### attachment-create 

```bash
cinder --os-volume-api-version 3.50 attachment-create janus ae5cebfc-630d-43db-bf34-8847afbc926c
```

```bash
usage: cinder attachment-create [flag...] <volume> <server_id>

Create an attachment for a cinder volume.

Positional arguments:
  <volume>              Name or ID of volume or volumes to attach.
  <server_id>           ID of server attaching to.

Optional arguments:
  --connect <connect>   Make an active connection using provided connector
                        info (True or False).
  --initiator <initiator>
                        iqn of the initiator attaching to. Default=None.
  --ip <ip>             ip of the system attaching to. Default=None.
  --host <host>         Name of the host attaching to. Default=None.
  --platform <platform>
                        Platform type. Default=x86_64.
  --ostype <ostype>     OS type. Default=linux2.
  --multipath <multipath>
                        Use multipath. Default=False.
  --mountpoint <mountpoint>
                        Mountpoint volume will be attached at. Default=None.
```

### attachment-complete

```bash
cinder --os-volume-api-version 3.50 attachment-complete 1342ecd8-547d-4580-a8e8-57cff8c51d19
```

```bash
cinder attachment-complete <attachment> [<attachment> ...]

Complete an attachment for a cinder volume.

Positional arguments:
  <attachment>  ID of attachment or attachments to delete.
```



## backup

### backup-create 

```bash
cinder backup-create --name janus_ee --snapshot-id  cb8d259a-54e3-4593-83ce-6252f9a8d20a janus
```

```bash
cinder backup-create [flag...] <volume>

Positional arguments:
  <volume>              Name or ID of volume to backup.

Optional arguments:
  --backup-id <backup-id>
                        Specified uuid4 string as backup id. Default=None.
  --container <container>
                        Backup container name. Default=None.
  --backend <backend>   Backup backend name. ceph or swift.
  --name <name>         Backup name. Default=None.
  --description <description>
                        Backup description. Default=None.
  --incremental         Incremental backup. Default=False.
  --force               Allows or disallows backup of a volume when the volume
                        is attached to an instance. If set to True, backs up
                        the volume whether its status is "available" or "in-
                        use". The backup of an "in-use" volume means your data
                        is crash consistent. Default=False.
  --force-host          Use volume host first. Default=False.
  --snapshot-id <snapshot-id>
                        ID of snapshot to backup. Default=None.
```

### backup-list

```bash
cinder backup-list --volume_id a840bf60-0fc1-495f-939a-2b048962a75a
```

```bash
cinder backup-list [flag...]

Optional arguments:
  --all-tenants [<all_tenants>]
                        Shows details for all tenants. Admin only.
  --name <name>         Filters results by a name. Default=None. This option
                        is deprecated and will be removed in newer release.
                        Please use '--filters' option which is introduced
                        since 3.33 instead.
  --status <status>     Filters results by a status. Default=None. This option
                        is deprecated and will be removed in newer release.
                        Please use '--filters' option which is introduced
                        since 3.33 instead.
  --volume-id <volume-id>
                        Filters results by a volume ID. Default=None. This
                        option is deprecated and will be removed in newer
                        release. Please use '--filters' option which is
                        introduced since 3.33 instead.
  --marker <marker>     Begin returning backups that appear later in the
                        backup list than that represented by this id.
                        Default=None.
  --limit <limit>       Maximum number of backups to return. Default=None.
  --sort <key>[:<direction>]
                        Comma-separated list of sort keys and directions in
                        the form of <key>[:<asc|desc>]. Valid keys: id,
                        status, size, availability_zone, name, bootable,
                        created_at, reference. Default=None.
```

### backup-show

```bash
cinder backup-show 3a140bf5-9516-4a8d-a0f8-ad576cd2df1b
```

```bash
cinder backup-show <backup>

Positional arguments:
  <backup>  Name or ID of backup.
```

### backup-delete

```bash
cinder backup-delete [--force] <backup> [<backup> ...]

Removes one or more backups.

Positional arguments:
  <backup>  Name or ID of backup(s) to delete.

Optional arguments:
  --force   Allows deleting backup of a volume when its status is other than
            "available" or "error". Default=False.
```

## 
