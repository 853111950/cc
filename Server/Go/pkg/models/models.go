package models

import (
	"time"

	"ApsaraLive/pkg/alicloud/live"
)

type RoomInfo struct {
	// 直播Id
	ID string `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	// 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"index:idx_create_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at"`
	// 直播开始时间
	StartedAt time.Time `json:"started_at"`
	// 直播结束时间
	StoppedAt time.Time `json:"stopped_at"`

	// 直播标题
	Title string `json:"title"`
	// 直播公告
	Notice string `json:"notice"`
	// 直播封面
	CoverUrl string `json:"cover_url"`

	// 主播Id
	AnchorId string `json:"anchor_id"`

	// 主播Nick
	AnchorNick string `json:"anchor_nick"`

	// 扩展字段
	Extends string `json:"extends" gorm:"size:65535"`

	// 直播状态，0-准备中，1-已开始，2-已结束
	Status int `json:"status"`

	// 直播模式 0-普通直播, 1-连麦直播，2-PK直播
	Mode int `json:"mode"`
	// 群组Id
	ChatId string `json:"chat_id"`

	// 连麦Id
	MeetingId string `json:"meeting_id"`

	// 点播Id
	VodId string `json:"vod_id"`

	// 连麦成员信息（json序列化）
	MeetingInfo string `json:"meeting_info" gorm:"size:65535"`

	// 直播间统计
	Metrics *Metrics `json:"metrics,omitempty" gorm:"-"`

	// 直播转录制，点播信息
	VodInfo *VodInfo `json:"vod_info,omitempty" gorm:"-"`

	// 用户状态
	UserStatus *UserStatus `json:"user_status,omitempty" gorm:"-"`

	// 推流相关地址信息，动态生成
	PushInfo *PushLiveInfo `json:"push_url_info,omitempty" gorm:"-"`
	// 拉流相关地址信息，动态生成
	PullInfo *live.PullLiveInfo `json:"pull_url_info,omitempty" gorm:"-"`

	// 连麦PK 等信息， 动态生成
	LinkInfo *LinkInfo `json:"link_info,omitempty" gorm:"-"`
}

// 前缀字段（artc://）+固定字段（live.aliyun.com）+拉流标识位（play）+roomid（房间ID）+SdkIAppID（连麦应用ID）+UserID（连麦观众ID）+timestamp（有效时长时间戳）+token
// 播放前缀+播流域名+AppName（live） + StreamID（由连麦应用ID_房间ID_主播ID_camera组成）+auth_key

type Metrics struct {
	OnlineCount uint64 `json:"online_count" gorm:"-"`
	LikeCount   uint64 `json:"like_count" gorm:"-"`
	Pv          uint64 `json:"pv" gorm:"-"`
	Uv          uint64 `json:"uv" gorm:"-"`
}

type UserStatus struct {
	Mute       bool     `json:"mute" gorm:"-"`
	MuteSource []string `json:"mute_source" gorm:"-"`
}

type LinkInfo struct {
	// 推流地址
	RtcPushUrl string `json:"rtc_push_url"`
	// 拉流地址
	RtcPullUrl string `json:"rtc_pull_url"`
	// 普通观众CDN拉流地址
	CdnPullInfo *live.PullLiveInfo `json:"cdn_pull_info"`
}

type PushLiveInfo struct {
	RtmpUrl string `json:"rtmp_url"`
	RtsUrl  string `json:"rts_url"`
	SrtUrl  string `json:"srt_url"`
}

type Status struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

// LivePushStreamEvent 推流事件
// 见：https://help.aliyun.com/document_detail/199365.html?spm=5176.13499635.help.dexternal.35d92699jvVrc7#section-mxt-vfh-b6s
type LivePushStreamEvent struct {
	Action  string `form:"action" json:"action" binding:"required"`
	App     string `form:"app" json:"app" binding:"required"`
	Ip      string `form:"ip" json:"ip"`
	Id      string `form:"id" json:"id" binding:"required"`
	AppName string `form:"appname" json:"appname" binding:"required"`
	Time    int32  `form:"time" json:"time"`
	Height  string `form:"height" json:"height"`
	Width   string `form:"width" json:"width"`
	// 签名
	LiveSignature string `json:"liveSignature"`
	// 签名所需要的time
	LiveTimestamp string `json:"liveTimestamp"`
}

type JumpUrlRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	LiveId   string `json:"live_id" binding:"required"`
	UserName string `json:"user_name"`
}

type JumpUrlResponse struct {
	LiveJumpUrl string `json:"live_jump_url"`
}

type AuthTokenRequest struct {
	UserId    string `json:"user_id" binding:"required"`
	LiveId    string `json:"live_id" binding:"required"`
	UserName  string `json:"user_name"`
	AppServer string `json:"app_server"`
	Token     string `json:"token"`
}

type AuthTokenResponse struct {
	LoginToken string `json:"login_token"`
}

const (
	// PUBLIC 表示推流
	PUBLIC = "publish"

	// PUBLIC_DONE 表示断流
	PUBLIC_DONE = "publish_done"
)

const (
	// LiveStatusPrepare 0 准备中
	LiveStatusPrepare = iota
	// LiveStatusOn 1 已开始
	LiveStatusOn
	// LiveStatusOff 2 已结束
	LiveStatusOff
)

const (
	// LiveModeNormal 0 常规模式
	LiveModeNormal = iota
	// LiveModeLink 1 连麦模式
	LiveModeLink
	// LiveModePk 2 pk模式
	LiveModePk
)