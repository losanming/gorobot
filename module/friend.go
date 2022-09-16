package module

import (
	"encoding/json"
	"mytest/cqhttpServer/utils"
)

type FriendInfo struct {
	UserId   int64  `json:"user_id"`  //QQ号
	NickName string `json:"nickname"` //昵称
	Remark   string `json:"remark"`   //备注名
}

func GetFriendList() (friend_list []FriendInfo, err error) {
	url := utils.GetBaseUrl("get_friend_list")
	resp, err := utils.SendRequest(url, nil, nil, "GET")
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &friend_list)
	if err != nil {
		return
	}
	return friend_list, err
}
