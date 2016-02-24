# oohhc
A simple set of management servers and an admin client for metadata that uses the group store for persistent storage.


## oohhc-acctd
This is the account web service implemented in gRPC.  It can either be run on the command line or using systemctl

command line argument | environment variable | default | description
--------------------- | -------------------- | ------- | -----------
tls | OOHHC_ACCT_TLS | true | Connection uses TLS if true, else plain TCP
cert_file | OOHHC_ACCT_CERT_FILE | /etc/oort/server.crt | The TLS cert file
key_file |  OOHHC_ACCT_KEY_FILE | /etc/oort/server.key | The TLS key file
port | OOHHC_ACCT_PORT | 8449 | The acct server port
oortgrouphost | OOHHC_OORT_GROUP_HOST | 127.0.0.1:6380 | host:port to use when connecting to oort group store
skipverify | OOHHC_ACCT_SKIP_VERIFY | true | don't verify cert
superkey | OOHHC_ACCT_SKIP_VERIFY | 123456789 | Super User key used for authentication


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
