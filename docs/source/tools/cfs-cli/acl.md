# Access Control List Management (ACL)

## Add an IP to the Blacklist

Add the specified IP to the volume's ACL blacklist, preventing that IP from accessing or manipulating the volume.

```bash
cfs-cli acl add [volName] [ipAddress]
```

## Check if an IP is on the Blacklist

Check if a specified IP has been added to the volume's ACL blacklist.

```bash
cfs-cli acl check [volName] [ipAddress]
```

## Remove an IP from the Blacklist

Remove the specified IP from the volume's ACL blacklist, allowing that IP to access or manipulate the volume again.

```bash
cfs-cli acl del [volName] [ipAddress]
```

## List All IPs on the Blacklist

View all the IP addresses that are on the volume's ACL blacklist.

```bash
cfs-cli acl list [volName]
```
