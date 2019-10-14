# mTLS in Go

## Introduction

1. `TLS` (Transport Layer Security) provides the necessary encryption for applications when communicating over a network. 

2. `TLS` requires a `CA` (Certificate Authority) to issue a `X.509` digital certificate to a `service`.

3. `TLS` requires `clients` of the `service` use the issued `X.509` digital certificate to validate the `service` with the `CA`.

4. `TLS` relies on `PKI` (Public Key Infrastructure).

---

## Related Protocols

* `HTTPS` (Hypertext Transfer Protocol Secure) is an extension of HTTP that leverages TLS for security.

* `mTLS` or `Mutual TLS` is a process where secure transmission of data is achieved by both the `server` and the `client` by completing the TLS process in both directions.

---

## References

* [TLS](https://en.wikipedia.org/wiki/Transport_Layer_Security) - Wikipedia.

* [x509](https://en.wikipedia.org/wiki/X.509) - Wikipedia
