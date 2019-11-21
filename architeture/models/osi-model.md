# OSI Model

## Introductions

## OSI Layers

1. The `OSI Model` has __7 layers__.

```
+--------------+
| Application  |    Layer 7                 L7 Header                   |
+--------------+                                                        |
| Presentation |    Layer 6                 L6 Header                   |
+--------------+                                                        |
|   Session    |    Layer 5                 L5 Header                   |
+--------------+                                                        |
|  Transport   |    Layer 4     `Segments`  L4 Header                   |   Decreasing data size.
+--------------+                                                        |
|   Network    |    Layer 3     `Packets`   L3 Header                   |
+--------------+                                                        |
|  Data Link   |    Layer 2     `Frames`    L2 Header + L2 Trailer      |
+--------------+                                                        |
|   Physical   |    Layer 1     `Bits`                                  V
+--------------+
```

2. `Outbound data`  moves down the stack.

    1. At each layer as data moves __down__ the stack it is __split__ into smaller data chunks.

    2. At each layer as data moves __down__ the stack `headers` are added to each data chunk.

        * `Layer 2 frames` also have a `trailer` added.

3. `Inbound data`  moves up the stack.

    1. At each layer as data moves __up__ the stack `headers` are removed from each data chunk.

        * `Layer 2 frames` also have a `trailer` remove.

    2. At each layer as data moves __up__ the stack it is __aggregated__ into larger data chunks.

4. `Header Types` are __layer specific__ - `L2 Header`, `L3 Header`, etc.

5. A `Protocol Data Unit (PDU)` is combination of `data`  and `header type` - `L7 PDU`, `L6 PDU`, etc.

---

## Layer 7 - Application

> __Unit Type__ - None (Data is Data)

1. The `application` requiring two-way network communication.

2. Ensure the remote partner is available.

3. Ensure `data integrity`, `privacy`, `recovery features` are agreed.

---

## Layer 6 - Presentations

> __Unit Type__ - None (Data is Data)

1. Handles the data's application type. e.g. pdf.

2. Handles `encryption` and `decryption` of data.

---

## Layer 5 - Session

> __Unit Type__ - None (Data is Data)

1. The `manager` of the overall `data transfer process`.

2. Handles `establishing`, `maintaining`, and `terminating` sessions.

---

## Layer 4 - Transport

> __Unit Type__ - `Segment`

1. Establish a `logical end to end connection` between two systems.

    1. `Transmission Control Protocol (TCP)`

    2. `Universal Datagram Protocol (UDP)` 

---

## Layer 3 - Network

> __Unit Type__ - `Packet`

> __Device__ - `Routers`

1. Where `Routers` and `routing` occurs. A router's job is to answer two basic questions:

    1. What valid `paths` exist from here to our `destination`?

    2. What is the `best path` to use?

---

## Layer 2 - Data Link Layer

> __Unit Type__ - `Frame`

> __Device__ - `Switches`

> `Ethernet` is the most common implementation of `Layer 2`.

1. Where `Switches` occur. A switch's job is so get frames to a particular `physical address`:

    1. `Media Access Control (MAC) Address` - These are also known asL

        1. `Layer 2 Address` / `L2 Address`.

        2. `Hardware Address`.

        3. `Burned-in Address (BIA)`.

        4. `Physical Address` - NB: These do not run a the `physical layer`.

---

## Layer 1 - Physical

> __Unit Type__ - `Bit`

1. At the end of the day, all data is transmitted as `Bits` down the wire.







