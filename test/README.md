# sensor-mock

This script simulates a sensor sending data to the REST service using Python + [mimesis](https://mimesis.name/master/).

## Dependencies

- [python](https://www.python.org/) --- v3.14.6
- The following pip requirements:

  - mimesis --- 21.0.0
  - requests --- 2.34.2

## Configuring

You can configure this script using the following environment variables:

| Environment Variable |     Example    | Description                                                                                                        |
|---------------------:|:--------------:|:-------------------------------------------------------------------------------------------------------------------|
|  SENSOR_SERVICE_HOST | localhost:8080 | The base url for the recipient service.                                                                            |
|            SENSOR_ID |    sensor-0    | The name for the sensor                                                                                            |
|       SENSOR_PATTERN |     steady     | The pattern for the sensor to follow.<br/>Valid values are:<ul> <li>steady</li> <li>wave</li> <li>spike</li> </ul> |
|     SENSOR_INTERVALS |       10       | The interval for the mock sensor to use                                                                            |
|         SENSOR_DELAY |       1.0      | The delay between intervals for the mock sensor                                                                    |

### Patterns

The mock sensor supports the following patterns:

|   Name | Description                        |
|-------:|:-----------------------------------|
|   wave | Represents a wave of data          |
|  spike | Represents a massive spike of data |
| steady | Represents a steady stream of data |
