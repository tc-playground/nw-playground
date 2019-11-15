# Tun / Tap Virtual Network Interfaces

## Introduction

1. `TUN` and `TAP` are `virtual network` `kernel interfaces`.

    1. `TUN` devices work at `Layer 3` and handles `IP Packets`.

    2. `TAP` devices work at `Layer 2` and handles `Ethernet Frames`.

2. `TUN`/`TAP` devices are purely software and not backed by physical devices.

    1. A `user-space` program can `attach` to a `TUN`/`TAP` device. 
    
        * This allows the operating system to `forward` packets to it.

    2. A `user-space` program can also `send` packets to a `TUN`/`TAP` device. 
    
        * In this case the `TUN`/`TAP` device delivers (or "injects") these packets to the operating-system network stack. 
        
        * This emulates their reception from an external source.

3. `TUN` is short for `Network Tunnel` and can be used to implement `routing` functionality, such as `VCNs` and `overlay networks`.

4. `TAP` is short for `Network Tap` and can be used to implement `bridging` functionality, such as creating a `network bridge`.

5. __TUN and TAP in the Network Stack__

    ```
    +---+--------------------+                             +---+--------------------+
    | 7 | Application Layer  |                             | 7 | Application Layer  |
    +---+--------------------+                             +---+--------------------+
    | 6 | Presentation Layer |                             | 6 | Presentation Layer |
    +---+--------------------+                             +---+--------------------+
    | 5 | Session Layer      |                             | 5 | Session Layer      |
    +---+--------------------+                             +---+--------------------+
    | 4 | Transport Layer    |                             | 4 | Transport Layer    |
    +---+--------------------+                             +---+--------------------+
    | 3 | Network Layer      | <=   TUN (IP Packets)    => | 3 | Network Layer      |
    +---+--------------------+                             +---+--------------------+
    | 2 | Data Link Layer    | <= TAP (Ethernet Frames) => | 2 | Data Link Layer    |
    +---+--------------------+                             +---+--------------------+
    | 1 | Physical Layer     |                             | 1 | Physical Layer     |
    +---+--------------------+                             +---+--------------------+
    ```

---

## Applications

1. Virtual private networks

2. Virtual-machine networking

3. Network simulation

4. NAT (Network Address Translation)

---

## References

* __tuntap__

    * [tuntap Tutorial](https://backreference.org/2010/03/26/tuntap-interface-tutorial)

    * [tuntap - Wikipedia](https://en.wikipedia.org/wiki/TUN/TAP)

    * [tuntap - Linux Kernel Docs](https://github.com/torvalds/linux/blob/master/Documentation/networking/tuntap.txt)

* __Applications__

    * [MAC VLANs and Virtual Ethernets](http://www.pocketnix.org/posts/Linux%20Networking:%20MAC%20VLANs%20and%20Virtual%20Ethernets)

    * [vtun - sourceforge](http://vtun.sourceforge.net/tun/index.html)

    * [ns3 - Network Simulator](https://www.nsnam.org/)