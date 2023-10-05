# User Management

## Create User

Create user [USER ID].

``` bash
cfs-cli user create [USER ID] [flags]
```

### Flags

| Name       | Description                                                                     |
|------------|---------------------------------------------------------------------------------|
| access-key | Specify the access key for the user to use the object storage function.         |
| secret-key | Specify the secret key for the user to use the object storage function.         |
| password   | Specify the user password.                                                      |
| user-type  | Specify the user type, optional values are normal or admin (default is normal). |
| y, yes     | Skip all questions and set the answer to "yes".                                 |

## Delete User

Delete user [USER ID].

``` bash
cfs-cli user delete [USER ID] [flags]
```

### Flags

| Name   | Description                                     |
|--------|-------------------------------------------------|
| y, yes | Skip all questions and set the answer to "yes". |

## Show User

Get information of user [USER ID].

```bash
cfs-cli user info [USER ID]
```

## List Users

Get a list of all current users.

```bash
cfs-cli user list
```

## Update User Permission

Update the permission [PERM] of user [USER ID] for volume [VOLUME].

[PERM] can be "READONLY/RO", "READWRITE/RW", or "NONE".

```bash
cfs-cli user perm [USER ID] [VOLUME] [PERM]
```

## Update User Information

Update the information of user [USER ID].

```bash
cfs-cli user update [USER ID] [flags]
```

### Flags

| Name       | Description                                                 |
|------------|-------------------------------------------------------------|
| access-key | The updated access key value.                               |
| secret-key | The updated secret key value.                               |
| user-type  | The updated user type, optional values are normal or admin. |
| y, yes     | Skip all questions and set the answer to "yes".             |
