# mTLS in Go

## Introduction

* `TLS` (Transport Layer Security) provides the necessary encryption for applications when communicating over a network. 

* `HTTPS` (Hypertext Transfer Protocol Secure) is an extension of HTTP that leverages TLS for security.

* `mTLS` or `Mutual TLS` is a process where secure transmission of data is achieved by both the `server` and the `client`.

---

## Tutorial

### Insecure Cleint/Server

1. Create an `insecure_server` using the Go `ListenAndServe` method. Invoke with the `insecure_client`:

    1. In one terminal run the server: `make run-insecure-server`.

    2. In another terminal run the client: `make run-insecure-client`.

        * "Hello World should be returned.

### Secure Client/Server

1. Generate a `PEM` format certificate for `localhost`: `make create-certs`

    ```bash
    openssl req -newkey rsa:2048 \
        -new -nodes -x509 \
        -days 3650 \
        -out cert.pem \
        -keyout key.pem \
        -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"
    ```

2. Create a `secure_server` the server by using the Go `ListenAndServeTLS` method. In the `secure-client` change the protocol to `https`: 

    ```go
    log.Fatal(http.ListenAndServeTLS(":8443", "./pki/cert.pem", "./pki/key.pem", nil))
    ```

    > NB: A browser will show it as a security risk as we created a `self-signed certificate`: `https://localhost:8443/hello`

    ```bash
    2019/10/13 18:01:32 http2: server: error reading preface from client 127.0.0.1:56136: remote error: tls: bad certificate
    ```

    1. In one terminal run the server: `make run-insecure-server`.

        ```
        2019/10/13 18:09:27 http: TLS handshake error from 127.0.0.1:56368: remote error: tls: bad certificate
        ```

    2. In another terminal run the client: `make run-insecure-client`.

        ```
        2019/10/13 18:11:55 Get https://localhost:8443/hello: x509: certificate signed by unknown authority
        ```

## Mutually Secured Client/Server

1. Create am `mtls-server` and `mtls-client` that:

    1. Read in the `cert.pem` / `key.pem` pair.

    2. Create a Go `x509.NewCertPool` and register the `cert.pem`.

    3. In the `mtls-server`, create the `TLS Config` with the CA pool and enable Client certificate validation.

    4. In the `mtls-client`, create the `http.Transport` with the client certificate configured.

    * In one terminal run the server: `make run-mtls-server`.

    * In another terminal run the client: `make run-mtls-client`.

        * "Hello World should be returned.



---

## References

* [Guide to mTLS in Go ](https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go)