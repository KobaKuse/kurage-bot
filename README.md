# kurage-bot
Kurage醬是一個在Discord端上可以發送並顯示推特訊息連結的機器人  
以及她的可愛可以療癒使用者

## Getting Started(入門指南)

- [此鏈接](https://discord.com/oauth2/authorize?client_id=1168567471872163880&permissions=11264&scope=bot)使用後可以將機器人加入到伺服器
- 這樣就成功將機器人加入了！ 很簡單吧！
- **注意** 機器人只能由有伺服器管理權限者加入

## 針對開發者
- 克隆儲存庫
```
git clone https://github.com/KobaKuse/kurage-bot.git
```
- 在[此鏈接](https://discord.com/developers/applications)將BOT生成  
SETTING >> Bot >> ResetToken 點擊後獲得 Token
- 接下來將下方兩項開啟  
```
SERVER MEMBERS INTENT, MESSAGE CONTENT INTENT
```
- SETTING >> OAuth2 >> URL Generator内將`bot`打勾開啟  
開啟後會出現`BOT PERMISSIONS在下方`接者將該欄內三個功能開啟
```
Read Messages, Send Messages, Manage Messages
```
- 將.env文件中的XXXXXX替換成Token
- 運行 main.go

## Special Thanks
花染 - 繪製頭像圖等的會師  
Felix - 出借bot執行及維持用的伺服器  
烤魚 - 翻譯 README.md 翻譯成中文

## 最後
