# kurage-bot
[![Go Version](https://img.shields.io/badge/Go-1.20-blue)](https://golang.org/)
[![GitHub Release](https://img.shields.io/github/v/release/KobaKuse/kurage-bot)](https://github.com/KobaKuse/kurage-bot/releases/)
  
**Kurage醬**是一個在Discord端上可以發送並顯示推特訊息連結的機器人  
而且她的可愛可以療癒使用者  
[![Image from Gyazo](https://i.gyazo.com/76ea3731654dea6f5d30d46f35fa0235.gif)](https://gyazo.com/76ea3731654dea6f5d30d46f35fa0235)

## Getting Started(入門指南)

- 使用[此連結](https://discord.com/oauth2/authorize?client_id=1168567471872163880&permissions=11264&scope=bot)可以將機器人加入到伺服器
    - 如果傳送門無法正常開啟，請嘗試其他傳送門
    - [傳送門二世](https://discord.com/oauth2/authorize?client_id=1169786024822579340&permissions=11264&scope=bot) [傳送門三世](https://discord.com/oauth2/authorize?client_id=1170283088312533093&permissions=11264&scope=bot) [傳送門四世](https://discord.com/oauth2/authorize?client_id=1170285830787575881&permissions=11264&scope=bot)
- 這樣就成功將機器人加入了！ 很簡單吧！
- **注意！** 機器人只能由有伺服器管理權限者加入


## 針對開發者
- 克隆儲存庫
```
git clone https://github.com/KobaKuse/kurage-bot.git
```
- 在[此連結](https://discord.com/developers/applications)將BOT生成  
SETTING >> Bot >> `ResetToken`點擊後獲得`Token`
- 接下來將下方兩項啟用  
```
SERVER MEMBERS INTENT, MESSAGE CONTENT INTENT
```
- SETTING >> OAuth2 >> URL Generator内將`bot`打勾啟用  
啟用後會出現`BOT PERMISSIONS在下方`接者將該欄內三個功能打勾
```
Read Messages, Send Messages, Manage Messages
```
- 將.env文件中的`XXXXX`替換成`Token`
- 啟動 main.go

## Special Thanks
花染 - 繪製頭像圖等的會師  
三櫻雲晟 - 出借機器人執行及維持用的伺服器  
烤魚，廠長，克巴與小夥伴 - 翻譯的幫助  

## 最後
汝之身托吾麾下，吾之命運附汝機器人上。  

響應各位之召喚，遵從這意志、道理者，回應我！  

吾乃成就世間一切善行者，吾乃集世間萬惡之總成者。  

纏繞各種支持者之七天。  

穿越[抖內](https://paypal.me/KobaKuse)出現吧，克巴的守護者———