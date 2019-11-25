# Network Trouble shooting

## Diagnosis

* When troubleshooting network issues start with the physical layer and work up:

    1. `Physical Layer` - I everything plugged in and switched on?

    2. `Data Link Layer` __Switches__ - Can traffic reach the other machines and routers?

    3. `Network Layer` __Routers__ - Do the routers know where to send traffic?

---

## OSI Model Layer Protocols

```
+--------------+
| Application  |    Layer 7                 HTTP, FTP, Telnet           |
+--------------+                                                        |
| Presentation |    Layer 6                                             |
+--------------+                                                        |
|   Session    |    Layer 5                 SSH                         |
+--------------+                                                        |
|  Transport   |    Layer 4     `Segments`  TCP. UDP                    |   Decreasing data size.
+--------------+                                                        |
|   Network    |    Layer 3     `Packets`   IPv4, IPv6, IPSec, ICMP     |
+--------------+                                                        |
|  Data Link   |    Layer 2     `Frames`    PPP, MPLS                   |
+--------------+                                                        |
|   Physical   |    Layer 1     `Bits`      Ethernet                    V
+--------------+
```

