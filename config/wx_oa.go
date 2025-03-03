package config

import (
	"os"
	"strings"
)

const (
	Wx_Token_key           = "WX_TOKEN"
	Wx_App_Id_key          = "WX_APP_ID"
	Wx_App_Secret_key      = "WX_APP_SECRET"
	Wx_Subscribe_Reply_key = "WX_SUBSCRIBE_REPLY"
	Wx_Help_Reply_key      = "WX_HELP_REPLY"

	Wx_Event_Key_Chat_Gpt_key   = "AI_CHAT_GPT"
	Wx_Event_Key_Chat_Spark_key = "AI_CHAT_SPARK"
	Wx_Event_Key_Chat_Qwen_key  = "AI_CHAT_QWEN"

	Wx_Command_Help      = "/help"
	Wx_Command_Gpt       = "/gpt"
	Wx_Command_Spark     = "/spark"
	Wx_Command_Qwen      = "/qwen"
	Wx_Command_Gemini    = "/gemini"
	Wx_Command_Prompt    = "/prompt"
	Wx_Command_RmPrompt  = "/cpt"
	Wx_Command_GetPrompt = "/getpt"
	Wx_Command_SetModel  = "/setmodel"
	Wx_Command_GetModel  = "/getmodel"
	Wx_Command_Clear     = "/clear"

	Wx_Todo_Add  = "/ta"
	Wx_Todo_Del  = "/td"
	Wx_Todo_List = "/tl"

	Wx_Coin = "/cb"
)

var (
	Wx_Commands = []string{Wx_Command_Help, Wx_Command_Gpt, Wx_Command_Spark, Wx_Command_Qwen, Wx_Command_Gemini}
)

func GetWxToken() string {
	return os.Getenv(Wx_Token_key)
}
func GetWxAppId() string {
	return os.Getenv(Wx_App_Id_key)
}
func GetWxAppSecret() string {
	return os.Getenv(Wx_App_Secret_key)
}
func GetWxSubscribeReply() string {
	subscribeMsg := os.Getenv(Wx_Subscribe_Reply_key)
	return strings.ReplaceAll(subscribeMsg, "\\n", "\n")
}
func GetWxHelpReply() string {
	helpMsg := os.Getenv(Wx_Help_Reply_key)
	if helpMsg == "" {
		helpMsg = "欢迎关注 青松awa ~\n我是你的AI助理，有问题可以问我哦ヾ(≧▽≦*)o~\n如果需要自定义我的身份请输入\n/prompt ~\n示例：你想自定义我的身份为一个医生，请输入\n'/prompt 你是一名医生。'\n\nヾ(•ω•`)o"
	}
	return strings.ReplaceAll(helpMsg, "\\n", "\n")
}
func GetWxEventKeyChatGpt() string {
	return os.Getenv(Wx_Event_Key_Chat_Gpt_key)
}
func GetWxEventKeyChatSpark() string {
	return os.Getenv(Wx_Event_Key_Chat_Spark_key)
}
func GetWxEventKeyChatQwen() string {
	return os.Getenv(Wx_Event_Key_Chat_Qwen_key)
}

func GetBotWelcomeReply(botType string) string {
	switch botType {
	case Bot_Type_Gpt:
		return GetGptWelcomeReply()
	case Bot_Type_Gemini:
		return GetGeminiWelcomeReply()
	case Bot_Type_Spark:
		return GetSparkWelcomeReply()
	case Bot_Type_Qwen:
		return GetQwenWelcomeReply()
	}

	return botType
}
