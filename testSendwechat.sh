curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
        "version": "4",
        "groupKey": "{}:{alertname=\"node节点cpu使用率过高\"}",
        "status": "firing",
        "Receiver": "webhook",
        "GroupLabels": {
                "alertname": "node节点cpu使用率过高"
        },
        "CommonLabels": {
                "alertname": "node节点cpu使用率过高",
                "container": "kube-rbac-proxy",
                "endpoint": "https",
                "instance": "172.32.2.191",
                "job": "node-exporter",
                "namespace": "monitoring",
                "pod": "node-exporter-8xpz6",
                "prometheus": "monitoring/k8s",
                "service": "node-exporter",
                "severity": "warning"
        },
        "CommonAnnotations": {
                "description": "集群名称:储能-ems-cn  node名称:172.32.2.191  cpu使用率超过85%,当前值:4%",
                "summary": "node节点cpu使用率过高"
        },
        "ExternalURL": "http://alertmanager-main-0:9093",
        "Alerts": [
                {
                        "labels": {
                                "alertname": "node节点cpu使用率过高",
                                "container": "kube-rbac-proxy",
                                "endpoint": "https",
                                "instance": "172.32.2.191",
                                "job": "node-exporter",
                                "namespace": "monitoring",
                                "pod": "node-exporter-8xpz6",
                                "prometheus": "monitoring/k8s",
                                "service": "node-exporter",
                                "severity": "warning"
                        },
                        "Annotations": {
                                "description": "集群名称:储能-ems-cn  node名称:172.32.2.191  cpu使用率超过85%,当前值:4%",
                                "summary": "node节点cpu使用率过高"
                        },
                        "startsAt": "2024-09-10T06:47:48.741Z",
                        "endsAt": "0001-01-01T00:00:00Z"
                }
        ]
}
' http://10.40.61.135:8999/webhook\?key\=77d13f