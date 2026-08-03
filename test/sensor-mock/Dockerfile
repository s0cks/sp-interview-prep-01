ARG PYTHON_VERSION=3.11-slim
FROM python:${PYTHON_VERSION}
ENV SENSOR_ID="sensor-0"
ENV SENSOR_SERVICE_HOST="http://localhost:8080"
ENV SENSOR_PATTERN="steady"

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD [ "python", "sensor.py" ]
