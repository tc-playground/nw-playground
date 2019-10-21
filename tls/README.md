# mTLS in Go

## Introduction

1. `TLS` (Transport Layer Security) provides the necessary encryption for applications when communicating over a network. 

    1. __Privacy__ - Communication is private.

    2. __Integrity__ - Identification of participants.

2. `TLS` requires a `CA` (Certificate Authority) to issue a `X.509` digital certificate to a `service`.

3. `TLS` requires `clients` of the `service` use the issued `X.509` digital certificate to validate the `service` with the `CA`.

4. `TLS` relies on `PKI` (Public Key Infrastructure).

5. The TLS protocol comprises two layers: 

    1. The `TLS record` protocol.
    
    2. The `TLS handshake` protocol. 


> TLS is the precursor to SSL.

---

## Features

1. __Secure Private Connection__

    1. `Symmetric Cryptography` is used to encrypted the transmitted data.

    2. A `session` is established where the participant negotiated and agree upon an `encryption algorithm` and a `shared secret`.

2. __Trusted__

    1. The `identity` of the communicating parties can be authenticated using `public-key cryptography`.

3. __Reliable__

    1. The `connection` is reliable because each message transmitted includes a `message integrity check` using a `message authentication code`.

---

## Related Protocols

* `HTTPS` (Hypertext Transfer Protocol Secure) is an extension of HTTP that leverages TLS for security.

* `mTLS` or `Mutual TLS` is a process where secure transmission of data is achieved by both the `server` and the `client` by completing the TLS process in both directions.

---

## References

* [TLS](https://en.wikipedia.org/wiki/Transport_Layer_Security) - Wikipedia.

* [x509](https://en.wikipedia.org/wiki/X.509) - Wikipedia

* __Acunetix TLS Tutorial__

    1. [TLS Introduction](https://www.acunetix.com/blog/articles/tls-security-what-is-tls-ssl-part-1/)

    2. [TLS (SSL) History](https://www.acunetix.com/blog/articles/history-of-tls-ssl-part-2/)

    3. [TLS Terminology](https://www.acunetix.com/blog/articles/tls-ssl-terminology-basics-part-3/)

    4. [TLS Certificates](https://www.acunetix.com/blog/articles/tls-ssl-certificates-part-4/)

    5. [TLS Algorithm](https://www.acunetix.com/blog/articles/establishing-tls-ssl-connection-part-5/)

    6. [TLS Vulnerabilities](https://www.acunetix.com/blog/articles/tls-vulnerabilities-attacks-final-part/)

