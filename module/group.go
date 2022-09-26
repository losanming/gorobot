package module

import (
	"bytes"
	"encoding/json"
	"example.com/m/global"
	"example.com/m/utils"
	"fmt"
)

type GroupInfo struct {
	GroupId         int64  `json:"group_id"`          //群号
	GroupName       string `json:"group_name"`        //群名称
	GroupMemo       string `json:"group_memo"`        //群备注
	GroupCreateTime uint32 `json:"group_create_time"` //群创建时间
	GroupLevel      uint32 `json:"group_level"`       //群等级
	MemberCount     int32  `json:"member_count"`      //成员数
	MaxMemberCount  int32  `json:"max_member_count"`  //最大成员数
}
type GroupList struct {
	GroupLists []GroupInfo `json:"data"`
}

type GroupMemberInfo struct {
	GroupId         int64  `json:"group_id"`          //群号
	UserId          int64  `json:"user_id"`           //QQ号
	NickName        string `json:"nickname"`          //昵称
	Card            string `json:"card"`              //群名片/备注
	Sex             string `json:"sex"`               //性别，male或female 或 unknown
	Age             int32  `json:"age"`               //年龄
	Area            string `json:"area"`              //地区
	JoinTime        int32  `json:"join_time"`         //加群时间戳
	LastSentTime    int32  `json:"last_sent_time"`    //最后发言时间戳
	Level           string `json:"level"`             //成员等级
	Role            string `json:"role"`              //角色, owner 或 admin 或 member
	Unfriendly      bool   `json:"unfriendly"`        //是否不良记录成员
	Title           string `json:"title"`             //专属头衔
	TitleExpireTime int64  `json:"title_expire_time"` //专属头衔过期时间戳
	CardChangeable  bool   `json:"card_changeable"`   //是否允许修改群名片
	ShutUpTimestamp int64  `json:"shut_up_timestamp"` //禁言到期时间
}

type GroupMemberList struct {
	GroupMemberInfos []GroupMemberInfo `json:"data"`
}
type GroupMemberInfoList struct {
	GroupId       int64
	GroupInfoList GroupMemberList
}
type SendGroupMsg struct {
	GroupId    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

func GetGroupList() (group_list GroupList, err error) {
	url := global.HOSTPORT + "get_group_list"
	resp, err := utils.SendRequest(url, nil, nil, "GET")
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &group_list)
	if err != nil {
		return
	}
	return group_list, err
}
func GetGroupMemberInfoByGroupId(id int64) (temp GroupMemberInfoList, err error) {
	var group_member_info_list GroupMemberList
	url := global.HOSTPORT + fmt.Sprintf("get_group_member_list?group_id=%v&no_cache=false", id)
	resp, err := utils.SendRequest(url, nil, nil, "GET")
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &group_member_info_list)
	if err != nil {
		return
	}
	temp.GroupId = id
	temp.GroupInfoList = group_member_info_list
	return temp, err
}

func SendGroupMsgByGroupId(group_id int64, msg string) {

}

func SendMsgById(group_id int64, msg string) (err error) {
	var send SendGroupMsg
	send.GroupId = group_id
	send.Message = msg
	send.AutoEscape = false
	data, _ := json.Marshal(send)

	url := global.HOSTPORT + fmt.Sprintf("send_group_msg")
	_, err = utils.SendRequest(url, bytes.NewBuffer(data), nil, "POST")
	if err != nil {
		return err
	}
	return err
}
