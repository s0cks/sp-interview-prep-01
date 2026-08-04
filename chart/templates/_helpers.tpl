{{- define "sensors.api.selectorLabels" -}}
app.kubernetes.io/name: {{ .Release.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}-api
{{- end -}}

{{- define "sensors.api.labels" -}}
{{- include "sensors.api.selectorLabels" . }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end -}}

{{- define "sensors.web.selectorLabels" -}}
app.kubernetes.io/name: {{ .Release.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}-web
{{- end -}}

{{- define "sensors.web.labels" -}}
{{- include "sensors.web.selectorLabels" . }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end -}}

{{- define "sensors.mockSensor.selectorLabels" -}}
app.kubernetes.io/name: {{ .Release.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}-mock-sensor
{{- end -}}

{{- define "sensors.mockSensor.labels" -}}
{{- include "sensors.mockSensor.selectorLabels" . }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end -}}

{{- define "sensors.api.imageFullName" -}}
{{- if .Values.api.image -}}
{{ .Values.api.image.registry | default "ghcr.io/s0cks" }}/{{ .Values.api.image.name | default "sp-interview-prep-01-sensor-service" }}:{{ .Values.api.image.version | default .Chart.AppVersion }}
{{- else -}}
ghcr.io/s0cks/sp-interview-prep-01-sensor-service:{{ .Chart.AppVersion }}
{{- end -}}
{{- end -}}

{{- define "sensors.web.imageFullName" -}}
{{- if .Values.web.image -}}
{{ .Values.web.image.registry | default "ghcr.io/s0cks" }}/{{ .Values.web.image.name | default "sp-interview-prep-01-sensor-web" }}:{{ .Values.web.image.version | default .Chart.AppVersion }}
{{- else -}}
ghcr.io/s0cks/sp-interview-prep-01-sensor-web:{{ .Chart.AppVersion }}
{{- end -}}
{{- end -}}

{{- define "sensors.mockSensor.imageFullName" -}}
{{- if .Values.mockSensor.image -}}
{{ .Values.mockSensor.image.registry | default "ghcr.io/s0cks" }}/{{ .Values.mockSensor.image.name | default "sp-interview-prep-01-sensor-mock" }}:{{ .Values.mockSensor.image.version | default .Chart.AppVersion }}
{{- else -}}
ghcr.io/s0cks/sp-interview-prep-01-sensor-web:{{ .Chart.AppVersion }}
{{- end -}}
{{- end -}}
