# TODO

### IN PROGRESS

- [x] `set-door-passcodes` (cf. https://github.com/uhppoted/uhppoted/issues/40)
- [x] Replace Event pointer in GetStatusResponse with zero value (cf. https://github.com/uhppoted/uhppote-core/issues/18)
- [x] Fix [old event published on each card swipe](https://github.com/uhppoted/uhppoted-mqtt/issues/15)
- [x] Rework events handling (cf. https://github.com/uhppoted/uhppoted-mqtt/issues/16)
      - [x] Publish only received events to _events_ topic
      - [x] Publish all events to _events/feed_ topic
            - [x] Move IsDevNull to lib.os
            - [x] Seperate _live_ and _feed_ topics
            - [x] Make EventMap private to _listen_
            - [x] Retrieve events on swipe event
                  - [x] Retrieve events
                  - [x] Task queue
                  - [x] Limit retrieve task queue size
                  - [x] Rate limit task queue
            - [x] Merge to _main_
            - [x] Cleanup uhppoted-lib listen/events
      - [x] README
      - [x] CHANGELOG

- [ ] Remove startup warnings for missing encryption/signing/etc files if auth is not enabled.
- [ ] Clean up Paho logging
- [ ] MQTT v5

## TODO

- [ ] [Sparkplug B](https://github.com/eclipse-sparkplug/sparkplug)
- [ ] [MQTT Dash](https://iot.stackexchange.com/questions/6561/generic-mobile-applications-for-smart-home-devices)
- [ ] [UCANs](https://ucan.xyz/)
- [ ] (optionally) Generate uhppoted.conf if it doesn't exist
- [ ] Make reconnect time configurable
- [ ] Relook at encoding reply content - maybe json.RawMessage can preserve the field order
- [ ] Replace values passed in Context with initialised struct
- [ ] last-will-and-testament (?)
- [ ] publish add/delete card, etc to event stream
- [ ] MQTT v5.0
- [ ] [JSON-RPC](https://en.wikipedia.org/wiki/JSON-RPC) (?)
- [ ] Add to CLI
- [ ] Non-ephemeral key transport:  https://tools.ietf.org/html/rfc5990#appendix-A
- [ ] user:open/get permissions require matching card number 
- [ ] [AEAD](http://alexander.holbreich.org/message-authentication)
- [ ] Support for multiple brokers
- [ ] NACL/tweetnacl
- [ ] Report system events for e.g. listen bound/not bound

### Documentation

- [ ] TeX protocol description
- [ ] godoc
- [ ] build documentation
- [ ] install documentation
- [ ] user manuals
- [ ] man/info page

### Other

1.  github project page
2.  Integration tests
3.  Verify fields in listen events/status replies against SDK:
    - battery status can be (at least) 0x00, 0x01 and 0x04
4.  EventLogger 
    - MacOS: use [system logging](https://developer.apple.com/documentation/os/logging)
    - Windows: event logging
5.  Update file watchers to fsnotify when that is merged into the standard library (1.4 ?)
    - https://github.com/golang/go/issues/4068
6. [Teserakt E2E encryption](https://teserakt.io)
7. [Fernet encryption](https://asecuritysite.com/encryption/fernet)
8. [IoT standards](https://iot.stackexchange.com/questions/5363/mqtt-json-format-for-process-automation-industry)
9. [StackExchange: MQTT security tests](https://iot.stackexchange.com/questions/452/what-simple-security-tests-can-i-perform-on-my-mqtt-network)
10. [VerneMQ](https://vernemq.com)
11. [SparkplugB](https://cogentdatahub.com/connect/mqtt/sparkplug-b)

## NOTES

1. [os_arch.go](https://gist.github.com/camabeh/a02e6846e00251e1820c784516c0318f)
