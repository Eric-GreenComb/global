package global

/**
section for config
*/
// Gearman
const (
	ETCD_KEY_GEARMAN_ADDR = "/banerwai/gearman/conn"
)

// Redis
const (
	ETCD_KEY_REDIS_ADDR = "/banerwai/redis/conn"
)

// for set key
const (
	ETCD_KEY_TPL_EMAIL = "/banerwai/tpl/email/"
	ETCD_KEY_TPL_SMS   = "/banerwai/tpl/sms/"
	ETCD_KEY_TPL_WEB   = "/banerwai/tpl/web/"
)

/**
section for service reg/discovery
*/
// micros/query
const (
	ETCD_KEY_MICROS_QUERY_CATEGORY = "banerwai/micros/query/category/addr" // port: 39010
	ETCD_KEY_MICROS_QUERY_RENDER   = "banerwai/micros/query/render/addr"   // port: 39030

	ETCD_KEY_MICROS_QUERY_TOKEN = "banerwai/micros/query/token/addr" // port: 39040

	ETCD_KEY_MICROS_QUERY_AUTH   = "banerwai/micros/query/auth/addr"   // port: 39020
	ETCD_KEY_MICROS_QUERY_USER   = "banerwai/micros/query/user/addr"   // port: 39060
	ETCD_KEY_MICROS_QUERY_RESUME = "banerwai/micros/query/resume/addr" // port: 39070

	ETCD_KEY_MICROS_QUERY_PROFILE     = "banerwai/micros/query/profile/addr"     // port: 39050
	ETCD_KEY_MICROS_QUERY_WORKHISTORY = "banerwai/micros/query/workhistory/addr" // port: 39080
)

// command/query
const (
	ETCD_KEY_MICROS_COMMAND_CATEGORY = "banerwai/micros/command/category/addr" // port: 36010
	ETCD_KEY_MICROS_COMMAND_RENDER   = "banerwai/micros/command/render/addr"   // port: 36030

	ETCD_KEY_MICROS_COMMAND_TOKEN = "banerwai/micros/command/token/addr" // port: 36040

	ETCD_KEY_MICROS_COMMAND_AUTH   = "banerwai/micros/command/auth/addr"   // port: 36020
	ETCD_KEY_MICROS_COMMAND_USER   = "banerwai/micros/command/user/addr"   // port: 36060
	ETCD_KEY_MICROS_COMMAND_RESUME = "banerwai/micros/command/resume/addr" // port: 36070

	ETCD_KEY_MICROS_COMMAND_PROFILE     = "banerwai/micros/command/profile/addr"     // port: 36050
	ETCD_KEY_MICROS_COMMAND_WORKHISTORY = "banerwai/micros/command/workhistory/addr" // port: 36080
)
