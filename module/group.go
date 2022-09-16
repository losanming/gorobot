package module

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
