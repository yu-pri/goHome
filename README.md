Private home automation on raspberry PI

TODO:
+ show Pump and Heater state (On/Off)
+ rewrite logic on/off/auto pump and heater

+ scheduling to manage from command line

- handle gobot interruption (set relay to initial state)
- save configuration state between reboots
- remove older data from history (older that 24hrs - leave each 60th record)
- replace relay to more reliable
- gas/fog control (2 points)

- notification sending (email/sms)
 - pump off, but heater temperature is higher (sms)
 - too cold notification (sms)
 - system about freezing (sms)
 - fog/gas (sms)


Far planes
- boiler control
- power consumption monitor
- door sensors
- light control
