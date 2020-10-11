{{/*
	Made by DZ#6669 (438789314101379072)

	Trigger Type: RegEx
	Trigger: \A!d\sbump(?:\s+|\z)
*/}}

{{/* configuration area */}}
{{$voiceChannelID := 740556133038555246}}
{{$thanksMessage := "Thanks for bumping us!"}}{{/* thanks message if the user successfully bumps the server */}}
{{$bumpChannel := 764595403219664898}}{{/* channel to send the message telling users it is possible to bump again */}}
{{$bumpPing := "<@&724944927330533407>"}}{{/* role to ping when it is possible to bump once again */}}
{{$bumpMessage := "Please bump us, your support is appreciated!"}}{{/* message to be sent when it is possible for the server to be bumped once again */}}

{{/* do not edit below (unless you know what you're doing c: ) */}}
{{if .ExecData}}
	{{editChannelName $voiceChannelID "Bump Now!"}}
	{{sendMessage $bumpChannel (complexMessage "content" $bumpPing "embed" (cembed "title" "Bump" "description" $bumpMessage))}}
{{else}}
	{{if not (dbGet 0 "bump")}}
		{{dbSetExpire 0 "bump" 1 7200}}
		{{editChannelName $voiceChannelID (print "Next Bump in " (((dbGet 0 "bump).ExpiresAt.Sub currentTime).Round .TimeSecond))}}
		{{execCC .CCID nil 7200 "data"}}
		{{sendMessage nil (cembed "title" "Bump!" "description" $thanksMessage)}}
	{{end}}
{{end}}
