package eventsubscriptions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryAttributeMapping interface {
}

func unmarshalDeliveryAttributeMappingImplementation(input []byte) (DeliveryAttributeMapping, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryAttributeMapping into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Dynamic") {
		var out DynamicDeliveryAttributeMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DynamicDeliveryAttributeMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Static") {
		var out StaticDeliveryAttributeMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StaticDeliveryAttributeMapping: %+v", err)
		}
		return out, nil
	}

	type RawDeliveryAttributeMappingImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawDeliveryAttributeMappingImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
