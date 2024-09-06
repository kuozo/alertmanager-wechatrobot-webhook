package transformer

import (
	"bytes"
	"fmt"

	"github.com/k8stech/alertmanager-wechatrobot-webhook/model"
)

// 新增一个函数来获取告警颜色
func getAlertColor(severity string) string {
	switch severity {
	case "critical":
		return "warning"
	case "warning":
		return "comment"
	case "info":
		return "info"
	default:
		return "comment"
	}
}

// 限制微信 markdown 消息的最大长度
const maxMarkdownLength = 4096

// 分片发送的最大条数
const maxAlertsPerMessage = 5

// TransformToMarkdown transform alertmanager notification to wechat markdown message
func TransformToMarkdown(notification model.Notification, grafanaURL string, alertDomain string) (markdowns []*model.WeChatMarkdown, robotURL string, err error) {

	status := notification.Status
	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer
	alertCount := 0

	// 函数用于检查 buffer 长度，并在超长时分片
	flushBuffer := func() {
		if buffer.Len() > 0 {
			markdown := &model.WeChatMarkdown{
				MsgType: "markdown",
				Markdown: &model.Markdown{
					Content: buffer.String(),
				},
			}
			markdowns = append(markdowns, markdown)
			buffer.Reset() // 清空 buffer 以便下一次使用
		}
	}

	// 开始构建消息
	buffer.WriteString(fmt.Sprintf("### 当前状态:%s \n", status))

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		// 动态获取 namespace, container, alertDomain
		namespace := labels["namespace"]
		pod := labels["pod"]
		severity := labels["severity"]
		alertColor := getAlertColor(severity)

		// 增加告警内容到 buffer
		buffer.WriteString(fmt.Sprintf("\n# 告警: <font color='%s'>%s</font>\n", alertColor, annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("\n>【级别】 %s\n", severity))
		buffer.WriteString(fmt.Sprintf("\n>【类型】 %s\n", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>【主机】 %s\n", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("\n>【内容】 %s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n>【当前状态】 %s \n", status))
		buffer.WriteString(fmt.Sprintf("\n>【触发时间】 %s\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
		buffer.WriteString(fmt.Sprintf("\n[跳转Grafana看板](https://%s?orgId=1&var-NameSpace=%s&var-Pod=%s)", grafanaURL, namespace, pod))
		buffer.WriteString(fmt.Sprintf("\n[告警规则详情](http://%s/alerts?search=)", alertDomain))
		buffer.WriteString(fmt.Sprintf("\n[日志详情](https://aws-au-loki-grafana.vnnox.com/d/o6-BGgnnk/loki-kubernetes-logs?orgId=1&from=now-1h&to=now&var-query=&var-namespace=au&var-stream=All&var-container=vnnox-middle-oauth)"))

		// 检查当前长度是否超过最大长度
		if buffer.Len() >= maxMarkdownLength || alertCount >= maxAlertsPerMessage {
			flushBuffer()
			alertCount = 0 // 重置告警计数
		}
		alertCount++
	}

	// 最后一次刷新
	flushBuffer()

	return
}
