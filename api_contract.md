# HomePage

API endpoint: `/`

Method: any

Response: HTML control dashboard.

# Sensors

## Aggregated state 
Endpoint: `/control/currentState`

Method: any (TODO: leave only `GET`)

Returns:
```json
{
    "TempInside": 27.9,
    "TempOutside": 0,
    "TempHeater": 77.8,
    "TempReverse": 92.3,
    "TempEntryRoom": -1,
    "TempWaterBoiler": 0,
    "TempRecuperator": 0,
    "PumpState": true,
    "HeaterState": true,
    "Timestamp": 1616250448,
    "index": 0
}
```

## Granular sensor data
---
### Heater pump

#### Current state

Endpoint: `/control/pump`

Method: any (TODO: leave only `GET`)

Possible values: [`On`, `Off`, `Auto`]

Returns raw enum value:
```
Auto
```
#### Configuration
Endpoint: `/control/pump`

Query params: `?state=On`

Method: any (TODO: leave only `POST`)

Possible values: [`On`, `Off`, `Auto`]

---
### Electric heater

#### Current state

Endpoint: `/control/heat`

Method: any (TODO: leave only `GET`)

Possible values: [`On`, `Off`, `Auto`]

Returns raw enum value:
```
Auto
```
#### Configuration

Endpoint: `/control/pump`

Query params: `?state=On`

Method: any (TODO: leave only `POST`)

Possible values: [`On`, `Off`, `Auto`]

---
#### Config
Endpoint: `/control/config`

Method: any (TODO: leave only `GET`)

Returns JSON:
```
{
    "PumpState": "Auto",
    "HeaterState": "On",
    "DesiredTemp": "22"
}
```
---
#### Sensor state subscription
Endpoint: `/relays`

Method: WS

Returns:
```
{"Type":"desiredTempChanged","Key":"state","Value":"22"}
{"Type":"pumpStateChanged","Key":"state","Value":"Auto"}
{"Type":"heatStateChanged","Key":"state","Value":"On"}
{"Type":"heatStateChanged","Key":"state","Value":"Auto"}
{"Type":"heatStateChanged","Key":"state","Value":"On"}
```