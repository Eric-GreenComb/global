package constant

/**
section for config
*/
// Gearman
const (
	EtcdKeyGearmanAddr = "/banerwai/gearman/conn"
)

// Redis
const (
	EtcdKeyRedisAddr = "/banerwai/redis/conn"
)

// for set key
const (
	EtcdKeyTplEmail   = "/banerwai/tpl/email/"
	EtcdKeyTplSms     = "/banerwai/tpl/sms/"
	EtcdKeyTplWeb     = "/banerwai/tpl/web/"
	EtcdKeyTplContact = "/banerwai/tpl/contact/"

	EtcdKeyJSONCategory = "/banerwai/json/category/"
)

/**
section for service reg/discovery
*/
// micros/query
const (
	EtcdKeyMicrosQueryCategory = "banerwai/micros/query/category/addr" // port: 39010
	EtcdKeyMicrosQueryRender   = "banerwai/micros/query/render/addr"   // port: 39030

	EtcdKeyMicrosQueryToken = "banerwai/micros/query/token/addr" // port: 39040

	EtcdKeyMicrosQueryAuth   = "banerwai/micros/query/auth/addr"   // port: 39020
	EtcdKeyMicrosQueryUser   = "banerwai/micros/query/user/addr"   // port: 39060
	EtcdKeyMicrosQueryResume = "banerwai/micros/query/resume/addr" // port: 39070

	EtcdKeyMicrosQueryProfile     = "banerwai/micros/query/profile/addr"     // port: 39050
	EtcdKeyMicrosQueryWorkhistory = "banerwai/micros/query/workhistory/addr" // port: 39080

	EtcdKeyMicrosQueryContact = "banerwai/micros/query/contact/addr" // port: 39090
	EtcdKeyMicrosQueryAccount = "banerwai/micros/query/account/addr" // port: 39100
)

// micros/command
const (
	EtcdKeyMicrosCommandCategory = "banerwai/micros/command/category/addr" // port: 36010
	EtcdKeyMicrosCommandRender   = "banerwai/micros/command/render/addr"   // port: 36030

	EtcdKeyMicrosCommandToken = "banerwai/micros/command/token/addr" // port: 36040

	EtcdKeyMicrosCommandAuth   = "banerwai/micros/command/auth/addr"   // port: 36020
	EtcdKeyMicrosCommandUser   = "banerwai/micros/command/user/addr"   // port: 36060
	EtcdKeyMicrosCommandResume = "banerwai/micros/command/resume/addr" // port: 36070

	EtcdKeyMicrosCommandProfile     = "banerwai/micros/command/profile/addr"     // port: 36050
	EtcdKeyMicrosCommandWorkhistory = "banerwai/micros/command/workhistory/addr" // port: 36080

	EtcdKeyMicrosCommandContact = "banerwai/micros/command/contact/addr" // port: 36090
	EtcdKeyMicrosCommandAccount = "banerwai/micros/command/account/addr" // port: 36100
)
