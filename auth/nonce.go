package auth

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/uhppoted/uhppoted-lib/kvs"
	"github.com/uhppoted/uhppoted-mqtt/log"
)

type Nonce struct {
	ignore bool
	mqttd  struct {
		*kvs.KeyValueStore
		filepath string
	}
	counters struct {
		*kvs.KeyValueStore
		filepath string
	}
}

func NewNonce(verify bool, server, clients string) (*Nonce, error) {
	var err error

	var f = func(value string) (interface{}, error) {
		return strconv.ParseUint(value, 10, 64)
	}

	nonce := Nonce{
		ignore: !verify,
		mqttd: struct {
			*kvs.KeyValueStore
			filepath string
		}{
			kvs.NewKeyValueStore("nonce:mqttd", f),
			server,
		},
		counters: struct {
			*kvs.KeyValueStore
			filepath string
		}{
			kvs.NewKeyValueStore("nonce:clients", f),
			clients,
		},
	}

	if err = nonce.mqttd.LoadFromFile(server); err != nil {
		log.Warnf(LOG_TAG, "%v", err)
	}

	if err = nonce.counters.LoadFromFile(clients); err != nil {
		log.Warnf(LOG_TAG, "%v", err)
	}

	return &nonce, nil
}

func (n *Nonce) Validate(clientID *string, nonce *uint64) error {
	if !n.ignore || (clientID != nil && nonce != nil) {
		if clientID == nil {
			return errors.New("missing client-id")
		}

		if nonce == nil {
			return errors.New("missing nonce missing")
		}

		c, ok := n.counters.Get(*clientID)
		if !ok {
			c = uint64(0)
		}

		if *nonce <= c.(uint64) {
			return fmt.Errorf("nonce reused: %s, %d", *clientID, *nonce)
		}

		n.counters.Store(*clientID, *nonce, n.counters.filepath)
	}

	return nil
}

func (n *Nonce) Next() uint64 {
	c, ok := n.mqttd.Get("mqttd")
	if !ok {
		c = uint64(0)
	}

	nonce := c.(uint64) + 1

	n.mqttd.Store("mqttd", nonce, n.mqttd.filepath)

	return nonce
}
