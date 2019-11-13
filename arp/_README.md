# ARP (Address Resolution Protocol)

## Introduction

* `ARP` stands for `Address Resolution Protocol`, which is used to find the `Media Access Control (MAC)` address of a network neighbour for a given IPv4 Address.

* `ARP` is used to discover the `link address` associated with an `ip address`.

* `ARP` can also be used to `announce` an IP/MAC mapping.

* `Neighbor Discovery Protocol (NDP)` is the equivalent protocol for IPv6 networks. 

---

## Usage

* Two machine on the same (Ethernet) LAN.

* The IP of a target machine is known (maybe via DNS), but, the source machine does not have the destination machines MAC address in it's arp cache.

* An ARP request is broadcast to all machines on the LAN with the IP of target destination machine, and it's own source IP and MAC addresses.

* The target machine responds with it's MAC and IP address to the broadcaster, which may add the mapping to it's arp cache.

---

## Commands

1. `man arp` - Manual page.

2. `arp` - Show ARP table.

3. `arp -d [address]` - Delete a ARP table entry.

4. `arp -s [address] [hw_addr]` - Ass a new ARP table entry.

---

## References

* [Address Resolution Protocol](https://en.wikipedia.org/wiki/Address_Resolution_Protocol) - Wikipedia

* __ `cat /etc/net/arp`

* `man arp`