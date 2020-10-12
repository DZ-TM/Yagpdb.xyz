### README Requires Serious Editing
> These custom commands are meant to be used alongside bumpbot (DISBOARD#2760).

## 10m Interval

- Edits channel name every 10m to show when you can bump your server.

- Configuration Values: set $channelid to the ID of the channel you want it to edit. It is recommended to use a voice channel.
  > Example: {{$channelid := 606353968079044611}}

## Bump Command

- Gives rep when a user uses the bump command at the correct time.

- Edits the channel name to a timer starting with 2h0m0s to start the interval command.

- Edits the interval channel timer to "Bump Now!" when the bump is ready. 

- Cofiguration Values: $voiceChannelID, $thanksMessage, $bumpChannel, $bumpPing, $bumpMessage. Explanation and examples below.
  > **$voiceChannelID**: The channel that the bot will edit the name to the timer or to "Bump Now!"
  > > {{$voiceChannelID := 606354015776669714}}

  > **$thanksMessage**: A simple thank you message when a user bumps the server successfully.
  > > {{$thanksMessage := "Thanks for bumping the server, it's much apprieciated! Have a hug from us as a thanks!"}}
  
  > **$bumpChannel**: The channel that the message will be sent to when the next bump is available.
  > > {{$bumpChannel := 606353949913645067}}
  
  > **$bumpPing**: The role that will be pinged along with the message when the bump is available. Please only change the numbers contained in the `<@& >` that's in the code.
  > > {{$bumpPing := "<@&586003189233614848>"}}
  
  > **$bumpMessage**: The message that will be sent when the bump in available.
  > > {{$bumpMessage := "The bump command is ready! Bump it up peeps!"}}
  
## Disboard Mess
- we are not allowed to give __roles as rewards__ for bumps.

- Disboard does not have the power to forbid that but will remove you from their service if you do so.

- *STOP DISBOARD FROM FORCING US TO NOT REWARD USERS FOR THEIR ACTIONS*

**#BlameDisboard with this mess here**

<h1 align="center"><img src="https://i.imgur.com/WUL0zgc.png"></img></h1>
