package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k8stech/alertmanager-wechatrobot-webhook/model"
	"github.com/k8stech/alertmanager-wechatrobot-webhook/transformer"
)

// Send send markdown message to dingtalk
func Send(notification model.Notification, defaultRobot string, grafanaURL string, alertDomain string) (err error) {
	notificationJSON, _ := json.Marshal(notification)
	fmt.Println("Nova notification JSON:", string(notificationJSON))
	markdown, robotURL, err := transformer.TransformToMarkdown(notification, grafanaURL, alertDomain)
	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	var wechatRobotURL string

	if robotURL != "" {
		wechatRobotURL = robotURL
	} else {
		wechatRobotURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + defaultRobot
	}

	req, err := http.NewRequest(
		"POST",
		wechatRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
