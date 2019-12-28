# Simple 2 Router and Switch

## Introduction

* Use `Cisco Packet Tracer` to model connectivity between `2 ISR4321 routers` and a `3650-24PS switch`.

---

## Instructions

1. Start `Packet Tracer`!

2. Add 2 `ISR4321 routers`.

3. Add a `3650-24PS switch`.

    1. Add a `AC power supply module` to the switch to power it on.

4. Add 2 `straight through cables` to connect the `routers` to `ports` on the `switch`.

5. Enter `Router0`. 

    1. Enter `no` to skip initial setup.

    2. Enter `enable` to enter `privileged mode`.

    3. Enter `configure terminal` to enter `configuration mode`. NB: Or `conf t<tab>`.

    4. Enter `host Router0` to name the router.

        * `?` show the options.
    
    5. Enter `interface gigabitEthernet 0/0/0` to select the Fast Gigabit ethernet 0 interface.

    6. Enter `no shutdown` to bring up the interface.

    7. Enter `ip address 10.1.1.1 255.255.255.0` to assign the interface and IP address of `10.1.1.1` on the subnet `10.1.1.0`.

    8. Enter `end` to stop configuring the interface.

    9. Enter `ping 10.1.1.1` to ping Router0 from Router0.

    10. Enter `copy running-config startup-config` to save the config. NB: Ur, use `wr`.

6. Enter `Router1`

    1. Perform a similar set of operations but assign the address `10.1.1.2`.

        ```bash
        Would you like to enter the initial configuration dialog? [yes/no]: no

        Press RETURN to get started!

        Router>enable
        Router#configure terminal
        Enter configuration commands, one per line.  End with CNTL/Z.
        Router(config)#host Router1
        Router1(config)#interface gigabitEthernet 0/0/0
        Router1(config-if)#no shutdown 
        %LINK-5-CHANGED: Interface GigabitEthernet0/0/0, changed state to up
        %LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet0/0/0, changed state to up
        Router1(config-if)#ip address 10.1.1.2 255.255.255.0
        Router1(config-if)#end
        Router1#
        %SYS-5-CONFIG_I: Configured from console by console
        Router1#copy running-config startup-config 
        Destination filename [startup-config]? 
        Building configuration...
        [OK]
        Router1#        
        ```

    2. Enter `ping 10.1.1.2` to ping Router1 from Router1.

    3. Enter `ping 10.1.1.1` to ping Router0 from Router1.

7. Save the workspace and close `Packet Tracer`.

8. Open `Packet Tracer` and load the workspace. 

    1. Wait for things to start running.

    2. Open a shell on Router0.

    3. Enter `enable` to enter `privileged mode`.

    4. Enter `show run` to see the configuration.

    5. Enter `ping 10.1.1.2` to check we can still ping Router1.

---

## References

* [Packet Tracer Tutorial](https://www.youtube.com/watch?v=fnQB6cN3UWo)