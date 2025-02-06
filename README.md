# Redis-Mini

| Command | Syntax | Example | Output |
|---------|--------|---------|--------|
| SET | SET key value | `SET myKey "Hello"` | "OK"| 
| GET | GET key | `GET myKey` | "Hello"| 
| INCR | INCR key | `INCR myLikes` | n+1| 
| MGET | MGET key[...key] | `GET myKey myLikes invalidKey` | "Hello"</br>1</br>nil|
| KEYS | KEYS patter* | `KEYS my*` | |
| EXPIRE | EXPIRE key seconds | `EXPIRE myKey 10` | 1 | 
| DEL | DEL key[...key] | `DEL myKey` | 1 | 
| EXISTS | EXISTS key[...key] | `EXISTS myKey myLikes` | 1, 0| 
