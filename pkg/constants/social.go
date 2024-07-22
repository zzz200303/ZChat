package constants

// HandlerResult 处理结果 1. 未处理，2. 处理， 3. 拒绝
type HandlerResult int

const (
	NoHandlerResult     HandlerResult = iota + 1 // 未处理
	PassHandlerResult                            // 通过
	RefuseHandlerResult                          // 拒绝
	CancelHandlerResult
)

// redis存储格式
const (
	LIKE_KEY          = "assets_like_relation"
	OFFLINE_MESSAGE   = "offline_message"
	DEFAULT_VALUE     = "1"
	DEFAULT_NOT_VALUE = "0"
)

// 编号前准
const (
	TALK_PREFIX    = "T"
	COMMENT_PREFIX = "C"
	LIKE_PREFIX    = "L"
)

// GroupRoleLevel 群等级 1. 创建者，2. 管理者，3. 普通
type GroupRoleLevel int

const (
	CreatorGroupRoleLevel GroupRoleLevel = iota + 1 // 为什么会 从1开始？
	ManagerGroupRoleLevel
	AtLargeGroupRoleLevel
)

// GroupJoinSource 进群申请的方式： 1. 邀请， 2. 申请
type GroupJoinSource int

const (
	InviteGroupJoinSource GroupJoinSource = iota + 1
	PutInGroupJoinSource
)
