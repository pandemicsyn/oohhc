# oohhc
A simple set of management servers and an admin client for metadata that uses the group store for persistent storage.


## oohhc-acctd
The account data will be stored in the group store using the following format:

```
   Key:               /acct
   ChildKey:          [uuid]
   Value:   { 
               "id": [uuid],                 # same value as the ChildKey
               "name": "[string]",           
               "token": [uuid],     
               "status": "[string]",         # active, deleted, suspended
               "createdate": [timestamp],
               "deletedate": [timestamp]     # default 0
        }
```

### Arguments
This is the account web service implemented in gRPC.  It can either be run on the command line or using systemctl

command line argument | environment variable | default | description
--------------------- | -------------------- | ------- | -----------
tls  | OOHHC_ACCT_TLS | true |  Connection uses TLS if true, else plain TCP
cert_file | OOHHC_ACCT_CERT_FILE | /etc/oort/server.crt | The TLS cert file
key_file | OOHHC_ACCT_KEY_FILE | /etc/oort/server.key | The TLS key file
port | OOHHC_ACCT_PORT | 8449 | The acctd server port
oortgrouphost | OOHHC_ACCT_OORT_GROUP_HOST | 127.0.0.1:6380 | host:port to use when connecting to oort group
skipverify | OOHHC_ACCT_SKIP_VERIFY | true | don't verify cert
superkey | OOHHC_ACCT_SUPERUSER_KEY | 123456789 | Super User key used for authentication
mutualtlsGS | OOHHC_ACCT_GS_MUTUALTLS | true | Turn on MutualTLS for Group Store
insecureSkipVerifyGS | OOHHC_ACCT_GS_SKIP_VERIFY | false | Don't verify cert for Group Store
certfileGS | OOHHC_ACCT_GS_CERT_FILE | /etc/oort/client.crt | The client TLS cert file for the Group Store
keyFileGS | OOHHC_ACCT_GS_KEY_FILE | /etc/oort/client.key | The client TLS key file for the Group Store
cafileGS |	OOHHC_ACCT_GS_CA_FILE | /etc/oort/ca.pem | The client CA file


### Deploy files for oohhc-acctd to use with systemd
```
cp -av $GOPATH/src/github.com/letterj/oohhc/packaging/root/usr/share/oohhc/systemd/oohhc-acctd.service /lib/systemd/system
touch /etc/default/oohhc-acctd
```

---

## oohhc-filesysd
The file system management data will be stored in the group store using the following format:

filesystem
```
   Key:               /acct/[uuid]/fs
   ChildKey:          [uuid]
   Value:   { 
               "id": [uuid],                 # same value as the ChildKey
               "name": "[string]",           
               "sizeinbytes": [int64],     
               "status": "[string]",         # active, deleted, suspended
               "createdate": [timestamp],
               "deletedate": [timestamp]     # default 0
        }
```

ipaddress
```
   Key:               /acct/[uuid]/fs/[uuid]
   ChildKey:          [uuid]
   Value:   { 
               "id": [uuid],                 # same value as the ChildKey
               "addr": "[string]",           
               "status": "[string]",         # active, deleted, suspended
               "createdate": [timestamp],
               "deletedate": [timestamp]     # default 0
        }
```

### Argumemnts
This is the file system management web service implemented in gRPC.  It can either be run on the command line or using systemctl

command line argument | environment variable | default | description
--------------------- | -------------------- | ------- | -----------
tls  | OOHHC_FSYS_TLS | true |  Connection uses TLS if true, else plain TCP
cert_file | OOHHC_FSYS_CERT_FILE | /etc/oort/server.crt | The TLS cert file
key_file | OOHHC_FSYS_KEY_FILE | /etc/oort/server.key | The TLS key file
port | OOHHC_FSYS_PORT | 8449 | The acctd server port
oortgrouphost | OOHHC_FSYS_OORT_GROUP_HOST | 127.0.0.1:6380 | host:port to use when connecting to oort group
skipverify | OOHHC_FSYS_SKIP_VERIFY | true | don't verify cert
mutualtlsGS | OOHHC_FSYS_GS_MUTUALTLS | true | Turn on MutualTLS for Group Store
insecureSkipVerifyGS | OOHHC_FSYS_GS_SKIP_VERIFY | false | Don't verify cert for Group Store
certfileGS | OOHHC_FSYS_GS_CERT_FILE | /etc/oort/client.crt | The client TLS cert file for the Group Store
keyFileGS | OOHHC_FSYS_GS_KEY_FILE | /etc/oort/client.key | The client TLS key file for the Group Store
cafileGS |	OOHHC_FSYS_GS_CA_FILE | /etc/oort/ca.pem | The client CA file

### Deploy files for oohhc-filesysd to use with systemd
```
cp -av $GOPATH/src/github.com/letterj/oohhc/packaging/root/usr/share/oohhc/systemd/oohhc-filesysd.service /lib/systemd/system
touch /etc/default/oohhc-filesysd
```

---

## oohhc-cli
This is the command line client that works with the account web service oohhc-acctd

```
NAME:
   oohhc-cli - Client used to manage accounts for FSAAAS
USAGE:
   oohhc-cli [global options] command [command options] [arguments...]
VERSION:
   0.0.1
COMMANDS:
   create, c	create a new account
   list, l	list all accounts
   get, g	details on a specific account
   delete, d	mark an account deleted
   update, u	update the information on an account
        OPTIONS:
          --name, -N 		New name for an account.
          --token, -T 		New token for the account
          --status, -S 	New status for an account.
   help, h	Shows a list of commands or help for one command
GLOBAL OPTIONS:
   --key, -k 				Access key for oohhc-acctd [$OOHHC_ACCESS_KEY]
   --server, -s "127.0.0.1:8449"	Address of the oohhc-acctd server [$OOHHC_SERVER_ADDRESS]
   --help, -h				show help
   --version, -v			print the version
```
