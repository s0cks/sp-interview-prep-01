import os
import requests
import time
import math
import random
from mimesis import Numeric

num_gen = Numeric()


class SensorServiceClient:
    def __init__(self, url, sensor_id):
        self._url = url
        self._sensor_id = sensor_id

    def send(self, val):
        url = f"{self._url}/sensors/{self._sensor_id}"
        print(f"posting to: {url}")
        return requests.post(
            url,
            json={
                "data": val,
            },
        )


def steady_pattern(step, base=22.0):
    """Generates a stable baseline with slight random jitter."""

    jitter = num_gen.float_number(start=-0.5, end=0.5)
    return round(base + jitter, 2)


def spike_pattern(step, base=22.0):
    """Generates mostly steady data with a massive random spike every 10 steps."""

    if step % 10 == 0:
        return round(base + num_gen.float_number(start=15.0, end=30.0), 2)
    return steady_pattern(step, base)


def wave_pattern(step, base=22.0):
    """Generates a smooth, cyclic sine-wave pattern with mild noise."""

    sine_value = math.sin(step * 0.2) * 5.0  # Fluctuates up/down by 5 units
    jitter = num_gen.float_number(start=-0.2, end=0.2)
    return round(base + sine_value + jitter, 2)


def run_sensor_simulator(url, sensor_id, pattern_type="steady", intervals=5, delay=1.0):
    """
    Runs the simulation loop.
    pattern_type can be: "steady", "spike", or "wave"
    """

    patterns = {"steady": steady_pattern, "spike": spike_pattern, "wave": wave_pattern}
    selected_pattern = patterns.get(pattern_type, steady_pattern)
    print(f"Starting simulation using [{pattern_type}] pattern sending to {url}...\n")

    for step in range(1, intervals + 1):
        sensor_value = selected_pattern(step)
        payload = {"data": sensor_value}

        try:
            response = requests.post(
                f"{url}/sensors/{sensor_id}", json=payload, timeout=2.0
            )

            print(
                f"[{step}/{intervals}] sent: {payload} | response: {response.status_code}"
            )
        except requests.exceptions.RequestException as e:
            print(f"[{step}/{intervals}] network error: {e}")

        time.sleep(delay)


if __name__ == "__main__":
    service_url = os.getenv("SENSOR_SERVICE_HOST", "http://localhost:8080")
    sensor_id = os.getenv("SENSOR_ID", "sensor-0")
    sensor_pattern = os.getenv("SENSOR_PATTERN", "steady")
    sensor_intervals = int(os.getenv("SENSOR_INTERVALS", 10))
    sensor_delay = float(os.getenv("SENSOR_DELAY", 1.0))

    run_sensor_simulator(
        url=service_url,
        sensor_id=sensor_id,
        pattern_type=sensor_pattern,
        intervals=sensor_intervals,
        delay=sensor_delay,
    )
