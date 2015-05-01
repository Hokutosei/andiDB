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
./bin/andidb-server
```

```
// andiDB http client
./bin/andidb-http-client

```
### commands


#### set

```bash
$ ./bin/andidb-http-client set server1 10.0.1.1
{"status":200,"response":"ok"}
```

#### get
```bash
$ ./bin/andidb-http-client get server1
{"status":200,"response":"10.0.1.1"}
```

#### lpush
```bash
$ ./bin/andidb-http-client lpush servers 10.0.1.1
$ ./bin/andidb-http-client lpush servers 10.0.1.5
$ ./bin/andidb-http-client lpush servers 10.0.1.6
$ ./bin/andidb-http-client lpush servers 10.0.1.4
$ ./bin/andidb-http-client lpush servers 10.0.1.7
```

#### lrange
```bash
./bin/andidb-http-client lrange servers 0 1
{"status":200,"response":["10.0.1.1"]}

./bin/andidb-http-client lrange servers 0 3
{"status":200,"response":["10.0.1.1","10.0.1.5","10.0.1.6"]}

./bin/andidb-http-client lrange servers 0 -1
{"status":200,"response":["10.0.1.1","10.0.1.5","10.0.1.6","10.0.1.4","10.0.1.7"]}
```

### keys

regex string should be passed on the values for key search

```bash
./bin/andidb-http-client keys *
```
