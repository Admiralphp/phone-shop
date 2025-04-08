apiVersion: v1
kind: Service
metadata:
  name: {{ include "name" . }}
spec:
  selector:
    app: {{ include "name" . }}
  ports:
    - protocol: TCP
      port: {{ .Values.service.port }}
      targetPort: {{ .Values.service.targetPort }}
  type: ClusterIP