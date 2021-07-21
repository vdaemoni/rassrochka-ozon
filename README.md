# In-memory DB for Ozon Fintech internship

## To run:

Server:
```bash
go run srcs/tcp-server.go
```
Client:
```bash
go run srcs/tcp-client.go
```

## Usage:

client:
```bash

# set a key in memory db
set <key> <value>

# get a key from the memory db
get <key>

# delete a key from the memory db
del <key>

# disconnect
exit
```
