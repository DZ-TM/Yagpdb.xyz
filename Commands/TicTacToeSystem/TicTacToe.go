{{/*
	Made by DZ#6669 (438789314101379072)
	Trigger Type: RegEx
	Trigger: \A(?:\-|<@!?204255221017214977>)\s*(?:tictactoe|ttt)(?:\s+|\z)
*/}}
 
{{/* configuration area */}}
{{$delay:=3600}}{{/*time in seconds of which, if no response is received, ends the game*/}}
{{$emotes:=cslice "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣" "6️⃣" "7️⃣" "8️⃣" "9️⃣"}}{{/*two things to keep in mind 1) use 9 DIFFERENT emojis and 2) of emojis being used, use on of three formats by discord 1. for animated emotes <a:example:123456789> 2. for custom emotes <:example:123456789> 3. for normal emojis, use their literal emotes or their unicodes*/}}
{{$embed:=sdict "color" 123456}}{{/*basic layout for embed, other details will need to be edited in the actual code below, learn how to edit it at */}}
 
{{/* do not edit below */}}
{{$emoteData:=cslice}}{{$desc:=""}}
{{define "change"}}
	{{range $k,$v:=.change}}
		{{- $.embed.Set $k $v -}}
	{{end}}
	{{if not .correct}}
		{{sendMessage nil (cembed .embed)}}
	{{end}}
{{end}}
{{range $i,$_:=$emotes}}
	{{- $emoteData =$emoteData.Append (reFind `(?:[\w~]{2,32}:\d{17,19}|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(?:\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` $_)}}
	{{- if mod (add $i 1) 3}}
		{{- $desc =print $desc $_ "\t"}}
	{{- else}}
		{{- $desc =print $desc $_ "\n"}}
	{{- end -}}
{{end}}
{{if .CmdArgs}}
	{{with index .CmdArgs 0|reFind `\d{17,}`|toInt|getMember}}
		{{if or .User.Bot (eq .User.ID $.User.ID)}}
			{{template "change" (sdict "embed" $embed "change" (sdict "title" "Invalid Input" "description" "Please input a valid user. This user must be in the server and not be a bot."))}}
		{{else}}
			{{template "change" (sdict "embed" $embed "change" (sdict "footer" (sdict "icon_url" ($.User.AvatarURL "256") "text" (print (or $.Member.Nick $.User.Username) "’s turn!")) "title" "TicTacToe" "description" $desc) "correct" 1)}}
			{{$retID:=sendMessageRetID nil (cembed $embed)}}
			{{dbSetExpire $retID "ttt" (sdict "x" (userArg $.User.ID) "o" .User "turn" "x" "steps" (cslice "empty" "empty" "empty" "empty" "empty" "empty" "empty" "empty" "empty")) $delay}}
			{{range $emoteData}}
				{{- addMessageReactions nil $retID . -}}
			{{end}}
		{{end}}
	{{else}}
	 	 {{template "change" (sdict "embed" $embed "change" (sdict "title" "Invalid Input" "description" "Please use the command correctly, the correct syntax is `-tictactoe @user`."))}}
	{{end}}
{{else}}
	 {{template "change" (sdict "embed" $embed "change" (sdict "title" "Insufficient Arguments" "description" "You have not entered any arguments. The correct command syntax is `-tictactoe @user`."))}}
{{end}}
