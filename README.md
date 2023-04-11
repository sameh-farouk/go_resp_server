# GO_RESP_SERVER
[RESP](https://redis.io/docs/reference/protocol-spec/) server with two custom commands.

This is a learning by doing project, built in few hours while reading how resp protocol works.

Disclaimer:
Don't use this in production.

 ## Compile & Run Server
run the command below to compile the server:

```bash
go build ./cmd/server/
```

Next, run the server as:

```bash
./server
```

The above command should run the TCP server and we can make tcp connections to it.

## Testing the TCP Server
The server include handlers for two custom commands, `HI` and `TESTURL`.

To test the resp server, you can use redis-cli using the command:

```bash
redis-cli -p 3333
```

Try `HI` command, the server should respond with a Hi back

```bash
"Hi back!""
```

also you can use `TESTURL` command to check if url is reachable.
it should response with 1 if the url is reachable, otherwise 0:

```bash
redis-cli -p 3333
127.0.0.1:3333> TESTURL threefold.io
(integer) 1
127.0.0.1:3333> TESTURL gle.co
(integer) 0
```

if redis-cli not available, you can use netcat

```bash
{ echo -e '*1\r\n$2\r\nHI\r\n'; sleep 1; } | netcat localhost 3333
```

output

```bash
$8
Hi back!
```

{ echo -e '*2\r\n$7\r\nTESTURL\r\n$5\r\ng.com'; sleep 3; } | netcat localhost 3333

output

```bash
:0
```

