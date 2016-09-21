package constant

// Freelancer Type
const (
	FreelancerTypeIndependent = iota // 0
	FreelancerTypeAgency             // 1
)

// Last Activity
const (
	LastActiveWithin2Weeks = iota // 0
	LastActiveWithin1Month        // 1
	LastActiveWithin2Months
)

// Availability
const (
	AvailableTrue = iota // 0
	AvailableFalse
)

// Available Hours
const (
	AvailableMorethan30hrs = iota // 0
	AvailableLessthan30hrs
	AvailableAsNeeded
)

// Profile Visibility
const (
	VisibilityPublic = iota // 0
	VisibilityOnlyLogin
	VisibilityPrivate
)

// Skill Experience Level
const (
	ExperienceLevelEntry = iota // 0
	ExperienceLevelIntermediate
	ExperienceLevelExpert
)

// Hours Rate
const (
	HoursRateLessthan50 = iota // 0
	HoursRate50to100
	HoursRate100to150
	HoursRate150to200
	HoursRate200to300
	HoursRateMorethan300
)

// Hours Billed
const (
	HoursBilledAtleast1Hour = iota // 0
	HoursBilledAtleast100Hour
	HoursBilledMorethan1000Hour
)

// Job Success
const (
	JobSuccess80 = iota // 0
	JobSuccess90
)

// Contact Sign Enum
const (
	ContactSignNull = iota // 0
	ContactSignClient
	ContactSignFreelancer
	ContactSignDealed
)

// Account billing status
const (
	BillingStatusCreate = iota // 0
	BillingStatusDeal
	BillingStatusCancel
)

// Account billing CostType
const (
	BillingCostTypePlatformFee = iota // 0 平台服务费 +
	BillingCostTypeRecharge           // 甲方充值到平台甲方账户 +
	BillingCostTypeRechargeFee        // 甲方充值手续费 -
	BillingCostTypeATransfer          // 甲方账户向乙方账户转账，扣除平台服务费 -
	BillingCostTypeBTransfer          // 乙方账户接受甲方转账 +
	BillingCostTypeCash               // 甲,乙方提现 -
	BillingCostTypeCashFee            // 提现支付手续费 -
)

// Account billing payment type
const (
	PayTypeBankRemittance = 0 // 银行汇款 Fee : 0%
	PayTypeWechatPay      = 7 // 微信支付 Fee : 0.7%
)

// 平台服务费率：10%，同一商家如果交易额超过10000元，费率为5%
const (
	PlatformFeeRate10 = 10
	PlatformFeeRate5  = 5
)

// Currency Enum
// HKD港币、KRW韩元、CHF瑞郎、SGD新加坡元、MYR马来西亚币、IDR印尼、NZD新西兰、VND越南、THB泰铢、PHP菲律宾
const (
	CurrencyCNY = "CNY" // 人民币
	CurrencyUSD = "USD" // 美元
	CurrencyEUR = "EUR" // 欧元
	CurrencyGBP = "GBP" // 英镑
	CurrencyJPY = "JPY" // 日元
	CurrencyCAD = "CAD" // 加元
	CurrencyAUD = "AUD" // 澳元
)
