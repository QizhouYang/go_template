package message

import (
	"go_template/pkg/constant"
	"go_template/pkg/util/message/client"
)

type MessageClient interface {
	SendMessage(vars map[string]interface{}) error
}

func NewMessageClient(vars map[string]interface{}) (MessageClient, error) {
	if vars["type"] == constant.Email {
		return client.NewEmailClient(vars)
	}
	if vars["type"] == constant.DingTalk {
		return client.NewDingTalkClient(vars)
	}
	if vars["type"] == constant.WorkWeiXin {
		return client.NewWorkWeixinClient(vars)
	}
	return nil, nil
}
