### andiDB
Key-Value Pair database written in Golang with 0 dependency, fast, supports HTTP client, async background save.

#### data types
* string
* array

#### to support types
* complex objects
* nested objects/arrays
* more


### client
example, build client first

```
// root(andiDB server)
go build -v
```

```
// andiDB http client
// cd client/
go build -v

```
### commands


#### set

```bash
$ ./client set server1 10.0.1.1
{"status":200,"response":"ok"}
```

#### get
```bash
$ ./client get server1
{"status":200,"response":"10.0.1.1"}
```

#### lpush
```bash
$ ./client lpush servers 10.0.1.1
$ ./client lpush servers 10.0.1.5
$ ./client lpush servers 10.0.1.6
$ ./client lpush servers 10.0.1.4
$ ./client lpush servers 10.0.1.7
```

#### lrange
```bash
./client lrange servers 0 1
{"status":200,"response":["10.0.1.1"]}

./client lrange servers 0 3
{"status":200,"response":["10.0.1.1","10.0.1.5","10.0.1.6"]}

./client lrange servers 0 -1
{"status":200,"response":["10.0.1.1","10.0.1.5","10.0.1.6","10.0.1.4","10.0.1.7"]}
```
