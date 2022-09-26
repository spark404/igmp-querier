Querier
===


Very simple commandline tool to send a IGMP membership query. Mainly usefull together 
with a packetcapture tool like wireshark or tcpdump. 

Building
---

This version works on both darwin and linux. It's OS specific as the bind to interface method differs
per OS.

```
go build -o igmp-querier ./cmd/igmp-querier
```


Usage
---

```
igmp-querier <2|3> <interface>
```

Example:
```
# ./igmp-querier 3 eth0
```


