package global

// Freelancer Type
const (
	FreelancerType_Independent = iota // 0
	FreelancerType_Agency             // 1
)

// Last Activity
const (
	LastActive_Within2Weeks = iota // 0
	LastActive_Within1Month        // 1
	LastActive_Within2Months
)

// Availability
const (
	Available_True = iota // 0
	Available_False
)

// Available Hours
const (
	Available_Morethan30hrs = iota // 0
	Available_Lessthan30hrs
	Available_AsNeeded
)

// Profile Visibility
const (
	Visibility_Public = iota // 0
	Visibility_OnlyLogin
	Visibility_Private
)

// Skill Experience Level
const (
	ExperienceLevel_ENTRY = iota // 0
	ExperienceLevel_INTERMEDIATE
	ExperienceLevel_EXPERT
)

// Hours Rate
const (
	HoursRate_Lessthan50Yuan = iota // 0
	HoursRate_50_100Yuan
	HoursRate_100_150Yuan
	HoursRate_150_200Yua
	HoursRate_200_300Yua
	HoursRate_Morethan300Yua
)

// Hours Billed
const (
	HoursBilled_Atleast1Hour = iota // 0
	HoursBilled_Atleast100Hour
	HoursBilled_Morethan1000Hour
)

// Job Success
const (
	JobSuccess_80 = iota // 0
	JobSuccess_90
)

// Contact Sign Enum
const (
	ContactSign_Null = iota // 0
	ContactSign_Client
	ContactSign_Freelancer
	ContactSign_Dealed
)

// Account billing status
const (
	BillingStatus_Create = iota // 0
	BillingStatus_Deal
	BillingStatus_Cancel
)

// Account billing CostType
const (
	BillingCostType_PlatformFee = iota // 0 平台服务费 +
	BillingCostType_Recharge           // 甲方充值到平台甲方账户 +
	BillingCostType_RechargeFee        // 甲方充值手续费 -
	BillingCostType_A_Transfer         // 甲方账户向乙方账户转账，扣除平台服务费 -
	BillingCostType_B_Transfer         // 乙方账户接受甲方转账 +
	BillingCostType_Cash               // 甲,乙方提现 -
	BillingCostType_CashFee            // 提现支付手续费 -
)

// Account billing payment type
const (
	PayType_BankRemittance = 0 // 银行汇款 Fee : 0%
	PayType_WechatPay      = 7 // 微信支付 Fee : 0.7%
)

// 平台服务费率：10%，同一商家如果交易额超过10000元，费率为5%
const (
	PlatformFeeRate_10 = 10
	PlatformFeeRate_5  = 5
)

// Currency Enum
// HKD港币、KRW韩元、CHF瑞郎、SGD新加坡元、MYR马来西亚币、IDR印尼、NZD新西兰、VND越南、THB泰铢、PHP菲律宾
const (
	CURRENCY_CNY = "CNY" // 人民币
	CURRENCY_USD = "USD" // 美元
	CURRENCY_EUR = "EUR" // 欧元
	CURRENCY_GBP = "GBP" // 英镑
	CURRENCY_JPY = "JPY" // 日元
	CURRENCY_CAD = "CAD" // 加元
	CURRENCY_AUD = "AUD" // 澳元
)
