# Redis-Mini

| Command | Syntax | Example | Output |
|---------|--------|---------|--------|
| SET | SET key value | `SET myKey "Hello"` | "OK"|
| GET | GET key | `GET myKey` | "Hello"|
| DEL | DEL key[...key] | `DEL myKey` | 1 |
| INCR | INCR key | `INCR myLikes` | n+1 |
| MGET | MGET key[...key] | `GET myKey myLikes invalidKey` | "Hello"</br>1</br>nil |
| INCRBY | INCRBY key number | `INCRBY myLikes -1` | n-1 |
| INCRBYFLOAT | INCRBYFLOAT key float | `INCRBYFLOAT myLikes 0.01` | n+0.01 |
<!-- | KEYS | KEYS patter* | `KEYS my*` | | -->
<!-- | EXPIRE | EXPIRE key seconds | `EXPIRE myKey 10` | 1 |  -->
<!-- | EXISTS | EXISTS key[...key] | `EXISTS myKey myLikes` | 1, 0|  -->
