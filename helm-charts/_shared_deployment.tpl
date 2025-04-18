apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "name" . }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app: {{ include "name" . }}
  template:
    metadata:
      labels:
        app: {{ include "name" . }}
    spec:
      containers:
        - name: {{ include "name" . }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          ports:
            - containerPort: {{ .Values.service.targetPort }}