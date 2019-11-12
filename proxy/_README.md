# Network Proxies

## Introduction

* In `computer networks`, a `proxy` is a `computer system` or `application` that acts as an `intermediary` for `requests` and `responses` between `clients` and  other `machines`.

* `Proxies` are not all the same. Some are `half proxies`, some are `full proxies`; some are `forward` and some are `reverse`.


* `Proxy` Architecture

    ```
    [Client] <---> [Forward Proxy] <---> [Public Internet] <---> [Reverse Proxy] <---> [Resource Servers]
    ```
---

## Half-Proxy - One network stack connection.

1. `Half proxies` only maintain one connection in the network stack.

2. The term `half proxy` can denote two types of usage:

    1. `Direct Server Return` - Requests are proxied, but the responses are not, they are sent directly to the client. e.g. `streaming protocols`.

    2. `Delayed Binding` - The initial first few requests of a communication are examined by the proxy. e.g. to route a `TLS handshake`.

        1. `Initial Connection` is proxied.

        2. `Subsequent connections` are not proxied.

3. `Half proxies` are not a protocol endpoint and can only forward packet traffic.

    1. This typically restricts them to `Layer 4` port switching, routing, NAT-ing, etc.

4. Almost all `half-proxies` fall into the category of `reverse proxies`.

---

## Full-Proxy - Two network stack connections.

1. `Full proxies` maintain two connection in the network stack.

    1. `Inbound Connection` - Between the client (via public internet) and the proxy itself.
    
    2. `Outbound Connection` - Between the proxy itself and the destination servers.

2. `Full proxies` are an actual protocol endpoint, it must fully implement the protocols as both `a client` and `a server` (a packet-based design does not).

3. `Full proxies` have unique `inbound` and `outbound` connections; each can have its own TCP connection behavior.

    1. `Buffering` 
    
    2. `Retransmits`
    
    3. `TCP options`.

4. `Full proxies` can look at `incoming requests` and `outbound responses` and can manipulate both if the solution allows it. 

    1. `SSL termination`

    2. `Load balancing`

    3. `Connection optimisation`

    4. `Protocol Gateway`.

---

## Forward Proxy

1. `Forward proxies` usually sit between the `client` and `public internet`. `Bridge`. 

2. `Forward proxies` handle __outbound__ requests from clients.

3. `Forward proxies` can perform several services:

    1. `Content filtering`
    
    2. `Caching`

    3. `Authentication` and `Authorization`.

4. `Forward proxies` are generally HTTP (Web) proxies.

---

## Reverse-Proxy

1. `Reverse proxies` usually sit between `public internet` and `resource servers`.

2. `Reverse proxies` handle __inbound__ requests to servers.

3. `Reverse proxies` can perform several services:

    1. `Load balancing`.

    2. `SSL termination`.

4. `Reverse proxies` handle a wide range of protocols.

---

## Tunneling Proxy

1. A proxy server that passes unmodified requests and responses is usually called a `gateway` or sometimes a `tunneling proxy`.

---

## Open proxies (Forwarding Proxies)

1. An `open proxy` is a `forwarding proxy server` that is accessible by any Internet user.

    1. `Anonymous Proxy` - Conceal IP address.

        1. This server reveals іts identity as а server, but, does not disclose the initial IP address.
        
        2. Though this server cаn be discovered easily, іt cаn be beneficial for some users as іt hides the IP address.

    2. `Transparent Proxy` - Cache server content.

        1. This server reveals іts identity as а server, аnd with the support of HTTP headers the initial IP address cаn be viewed. 
        
        2. The main benefit of using this sort of server іs іts ability to cache the website and other resources.

2. An `anonymous open proxy` allows users to `conceal their IP address` while browsing the Web or using other Internet services. 

    * There are varying degrees of anonymity and methods of 'tricking' the client into revealing itself regardless of the proxy being used.

3. Most `VPNs` work through an open proxy.

---

## References

* [Proxy - Wikipedia](https://en.wikipedia.org/wiki/Proxy_server)

    * [Open Proxy - Wikipedia](https://en.wikipedia.org/wiki/Open_proxy)

    * [Reverse Proxy - Wikipedia](https://en.wikipedia.org/wiki/Reverse_proxy)

* [What is a Proxy?](https://devcentral.f5.com/s/articles/what-is-a-proxy-25839)

* [Concise Guide to Proxies](https://devcentral.f5.com/s/articles/the-concise-guide-to-proxies)

* [Full vs Half Proxies](https://devcentral.f5.com/s/articles/three-things-your-proxy-cant-do-unless-its-a-full-proxy)
