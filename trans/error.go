package trans

type ECODE_T int32

const (
	SUCCESS ECODE_T = 0

	ERROR_INNER   ECODE_T = 500
	ERROR_NOFOUND ECODE_T = 404
	ERROR_PARAM   ECODE_T = 420

	// 数据库操作错误
	ERROR_DB_NOT_FOUND  ECODE_T = 300001
	ERROR_DB_INNER      ECODE_T = 300002
	ERROR_DB_NOT_RECORD ECODE_T = 300003

	//
	ERROR_NOLOGIN          ECODE_T = 400001
	ERROR_OWNER            ECODE_T = 400002
	ERROR_RESOURCE_INVALID ECODE_T = 400003

	//
	ERROR_OP_EXPIRED ECODE_T = 500001
)

var ERRORMSG map[ECODE_T]string = map[ECODE_T]string{
	ERROR_INNER:            "内部错误",
	ERROR_PARAM:            "参数错误",
	ERROR_NOLOGIN:          "未登录",
	ERROR_OWNER:            "非法操作，请确认您是此资源所有人",
	ERROR_RESOURCE_INVALID: "资源不存在",

	ERROR_DB_INNER:      "数据操作失败",
	ERROR_DB_NOT_RECORD: "记录不存在",
	ERROR_OP_EXPIRED:    "操作过期",
}
