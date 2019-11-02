# DNS - Domain Name Service

## Introduction

1. The `Domain Name Service` assigns hierarchical `names` to `ip addresses`.

2. The `/etc/hosts` file on a machine can usually be used to provide a default mapping.

3. The `/etc/resolv.conf` file on a machine is used to:

    1.  The `default search domains` to resolve `Full Qualified Domain Names` from partial name.
    
    2. Configure 4 upstream DNS `nameservers`.

4. The `/etc/nsswitch.conf` __Name Service Switch__ provides a variety of sources for common configuration databases and name resolution mechanisms.

    1. This includes DNS.



---

## References

* [DNS - Wikipedia](https://en.wikipedia.org/wiki/Domain_Name_System)

* [/etc/hosts - Wikipedia](https://en.wikipedia.org/wiki/Hosts_(file))

* [/etc/resolv.conf - Wikipedia](https://en.wikipedia.org/wiki/Resolv.conf)

* []()

* [FQDN - Wikipedia](https://en.wikipedia.org/wiki/Fully_qualified_domain_name)