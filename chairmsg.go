package main

import (
	"errors"
	"fmt"
	"math"
)

type ChairMessage map[string]interface{}

var SAMPLE_MESSAGE = map[string]interface{}{
	"seqno": 0,
	"seqmax": 255,
	"timestamp_type": "rel",
	"timestamp": 123,
	"temperature": 1,
	"humidity": 1,
	"occupancy": false,
	"bottom_heat": 0,
	"bottom_fan": 0,
	"back_heat": 0,
	"back_fan": 0,
}

func NewChairMessage() ChairMessage {
	return make(map[string]interface{})
}

func (cm ChairMessage) SanityCheck() error {
	var seqmax uint64
	var val interface{}
	var ok bool
	
	/* Check required fields. */
	if val, ok = cm["seqmax"]; ok {
		if seqmax, ok = val.(uint64); !ok {
			return errors.New(fmt.Sprintf("required field seqmax is not a uint64: %v", val))
		}
	} else {
		return errors.New("missing required field seqmax")
	}
	
	if _, ok = cm["seqno"]; ok {
		if seqno, ok := val.(uint64); ok {
			if seqno >= seqmax {
				return errors.New(fmt.Sprintf("seqno  is not less than seqmax: %v >= %v", seqno, seqmax))
			}
		} else {
			return errors.New(fmt.Sprintf("required field seqno is not a uint64: %v", val))
		}
	} else {
		return errors.New("missing required field seqno")
	}
	
	if val, ok = cm["timestamp"]; ok {
		if _, ok := val.(int64); !ok {
			return errors.New(fmt.Sprintf("required field timestamp is not an int64: %v", val))
		}
	} else {
		return errors.New("missing required field timestamp")
	}
	
	if val, ok = cm["timestamp_type"]; ok {
		if tt, ok := val.(string); ok {
			if tt != "rel" && tt != "abs" {
				return errors.New(fmt.Sprintf("required field timestamp_type is invalid: %v", tt))
			}
		} else {
			return errors.New(fmt.Sprintf("required field timestamp_type is not a string: %v", val))
		}
	} else {
		return errors.New("missing required field timestamp_type")
	}
	
	/* Check optional fields. */
	if val, ok = cm["temperature"]; ok {
		if temp, ok := val.(float64); ok {
			if math.IsNaN(temp) || math.IsInf(temp, 0) || temp < 0 || temp >= 100 {
				return errors.New(fmt.Sprintf("optional field temperature is invalid: %v", temp))
			}
		} else {
			return errors.New(fmt.Sprintf("optional field temperature is not a float64: %v", val))
		}
	}
	
	if val, ok = cm["humidity"]; ok {
		if hum, ok := val.(float64); ok {
			if math.IsNaN(hum) || math.IsInf(hum, 0) || hum < 0 || hum >= 100 {
				return errors.New(fmt.Sprintf("optional field humidity is invalid: %v", hum))
			}
		} else {
			return errors.New(fmt.Sprintf("optional field humidity is not a float64: %v", val))
		}
	}
	
	if val, ok = cm["occupancy"]; ok {
		if _, ok := val.(bool); !ok {
			return errors.New(fmt.Sprintf("optional field occupancy is not a bool: %v", val))
		}
	}
	
	if val, ok = cm["bottom_heat"]; ok {
		if bottomh, ok := val.(uint8); ok {
			if bottomh > 100 {
				return errors.New(fmt.Sprintf("optional field bottom_heat is invalid: %v", bottomh))
			}
		} else {
			return errors.New(fmt.Sprintf("optional field bottom_heat is not a uint8: %v", val))
		}
	}
	
	if val, ok = cm["bottom_fan"]; ok {
		if bottomf, ok := val.(uint8); ok {
			if bottomf > 100 {
				return errors.New(fmt.Sprintf("optional field bottom_fan is invalid: %v", bottomf))
			}
		} else {
			return errors.New(fmt.Sprintf("optional field bottom_fan is not a uint8: %v", val))
		}
	}
	
	if val, ok = cm["back_heat"]; ok {
		if backh, ok := val.(uint8); ok {
			if backh > 100 {
				return errors.New(fmt.Sprintf("optional field back_heat is invalid: %v", backh))
			}
		} else {
			return errors.New(fmt.Sprintf("optional field back_heat is not a uint8: %v", val))
		}
	}
	
	if val, ok = cm["back_fan"]; ok {
		if backf, ok := val.(uint8); ok {
			if backf > 100 {
				return errors.New(fmt.Sprintf("optional field back_fan is invalid: %v", backf))
			}
		} else {
			return errors.New(fmt.Sprintf("optional field back_fan is not a uint8: %v", val))
		}
	}
	
	/* Check extra fields. */
	for key := range cm {
		if _, ok = SAMPLE_MESSAGE[key]; !ok {
			return errors.New(fmt.Sprintf("received extra key: %v", key))
		}
	}
	
	return nil
}

func (cm ChairMessage) SeqMax() uint64 {
	return cm["seqmax"].(uint64)
}

func (cm ChairMessage) SeqNo() uint64 {
	return cm["seqno"].(uint64)
}

func (cm ChairMessage) TimestampType() string {
	return cm["timestamp_type"].(string)
}

func (cm ChairMessage) Timestamp() int64 {
	return cm["timestamp"].(int64)
}

func (cm ChairMessage) Temperature() (temp float64, ok bool) {
	val, ok := cm["temperature"]
	if ok {
		temp = val.(float64)
	}
	return
}

func (cm ChairMessage) Humidity() (hum float64, ok bool) {
	val, ok := cm["humidity"]
	if ok {
		hum = val.(float64)
	}
	return
}

func (cm ChairMessage) Occupancy() (occ bool, ok bool) {
	val, ok := cm["occupancy"]
	if ok {
		occ = val.(bool)
	}
	return
}

func (cm ChairMessage) BottomHeat() (bottomh uint8, ok bool) {
	val, ok := cm["bottom_heat"]
	if ok {
		bottomh = val.(uint8)
	}
	return
}

func (cm ChairMessage) BottomFan() (bottomf uint8, ok bool) {
	val, ok := cm["bottom_fan"]
	if ok {
		bottomf = val.(uint8)
	}
	return
}

func (cm ChairMessage) BackHeat() (backh uint8, ok bool) {
	val, ok := cm["back_heater"]
	if ok {
		backh = val.(uint8)
	}
	return
}

func (cm ChairMessage) BackFan() (backf uint8, ok bool) {
	val, ok := cm["back_fan"]
	if ok {
		backf = val.(uint8)
	}
	return
}

