package command
import (
	// "time"
	// "fmt"
	// "example.com/main/command/tool"
	discord "github.com/bwmarrin/discordgo"
)
type Help struct {}

func AddHelp() Command{
	return Help{}
}

func (self Help) Handle(bot *discord.Session,msg *discord.MessageCreate) error{
	// embed := tool.NewEmbed().SetTitle("指令列表:").SetDescription(
	// 	"[>>點此加入官方支援區](https://discord.gg/6G3Gjr8yFF)" + "\n" + "[>>點此邀請機器人!](https://discord.com/api/oauth2/authorize?client_id=784044530931466250&permissions=8&scope=bot)",
	// ).SetAuthor(
	// 	"嗨~ 我是 PRØ-TECTER",
	// 	"https://cdn.discordapp.com/attachments/796395974792708097/796397747385270363/000_1.jpeg",
	// ).AddField(
	// 	"***.purge <數量>***  | 清除訊息",
	// 	"`使用權限` : __`管理員權限`__",
	// ).AddField(
	// 	"***.bl <+/-/add/del> <用戶>***  | 將用戶 加入/移除 黑名單",
	// 	"> [<用戶> 以是tag或是id或是名子]\n> 黑名單內的用戶加入伺服器 會被再次ban除\n`使用權限` : __`伺服器擁有者`__",
	// ).AddField(
	// 	"***.p***  | 開啟/關閉 頻道保護",
	// 	"> [預設為關閉 第一次輸入會開啟 再次輸入會關閉]\n> 防止頻道被刪除 刪除後會立即回復 (權限等都會相同，但內容會被清除)\n`使用權限` : __`伺服器擁有者`__",
	// ).AddField(
	// 	"***.pall <on/off>***  | 開啟/關閉 全頻道保護",
	// 	"> [開啟所有文字頻道的保護]\n> 名單內的成員不會偵測刷頻行為\n`使用權限` : __`伺服器擁有者`__",
	// ).AddField(
	// 	"***.igg***  | 開啟/關閉 忽略頻道",
	// 	"> [預設為關閉 第一次輸入會開啟 再次輸入會關閉]\n> 不想一直打指令就用它吧\n`使用權限` : __`伺服器擁有者`__",
	// ).AddField(
	// 	"***.wl <+/-/add/del> <用戶>***  | 將用戶 加入/移除 白名單",
	// 	"> [使用方式與黑名單相同]\n> 名單內的成員不會偵測刷頻行為\n`使用權限` : __`伺服器擁有者`__",
	// ).SetFooter("目前版本 : PRØ-TECTER NEO v1")
	// final_embed := embed.MessageEmbed
	// _,err := bot.ChannelMessageSendEmbed(msg.ChannelID, final_embed)
	// if err != nil{
	// 	return err
	// }
	bot.ChannelMessageSend(msg.ChannelID,"neo!help")
	return nil
}

func (self Help) GetCommandName() string{
	return "help"
}

func (self Help) ToArguments(args []string) error{
	return nil
}