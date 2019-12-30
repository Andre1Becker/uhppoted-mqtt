package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"uhppoted"
)

func (m *MQTTD) getDevices(impl *uhppoted.UHPPOTED, ctx context.Context, msg MQTT.Message) {
	operation := "get-devices"
	rq := uhppoted.GetDevicesRequest{}

	response, status, err := impl.GetDevices(ctx, rq)
	if err != nil {
		m.OnError(ctx, operation, "Error retrieving list of devices", status, err)
		return
	}

	if response != nil {
		reply := struct {
			MetaInfo *metainfo `json:"meta-info,omitempty"`
			uhppoted.GetDevicesResponse
		}{
			MetaInfo:           getMetaInfo(ctx, operation),
			GetDevicesResponse: *response,
		}

		m.Reply(ctx, reply)
	}
}

func (m *MQTTD) getDevice(impl *uhppoted.UHPPOTED, ctx context.Context, msg MQTT.Message) {
	operation := "get-device"
	body := struct {
		DeviceID *uint32 `json:"device-id"`
	}{}

	if err := json.Unmarshal(msg.Payload(), &body); err != nil {
		m.OnError(ctx, operation, "Cannot parse request", uhppoted.StatusBadRequest, err)
		return
	}

	if body.DeviceID == nil || *body.DeviceID == 0 {
		m.OnError(ctx, operation, "Missing/invalid device ID", uhppoted.StatusBadRequest, fmt.Errorf("Missing/invalid device ID '%s'", string(msg.Payload())))
		return
	}

	rq := uhppoted.GetDeviceRequest{
		DeviceID: *body.DeviceID,
	}

	response, status, err := impl.GetDevice(ctx, rq)
	if err != nil {
		m.OnError(ctx, operation, "Error retrieving device", status, err)
		return
	}

	if response != nil {
		reply := struct {
			MetaInfo *metainfo `json:"meta-info,omitempty"`
			uhppoted.GetDeviceResponse
		}{
			MetaInfo:          getMetaInfo(ctx, operation),
			GetDeviceResponse: *response,
		}

		m.Reply(ctx, reply)
	}
}
