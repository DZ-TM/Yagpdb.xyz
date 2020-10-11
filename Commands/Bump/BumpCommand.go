{{/*
	Made by DZ#6669 (438789314101379072)

	Trigger Type: RegEx
	Trigger: \A!d\sbump(?:\s+|\z)
*/}}

{{/* configuration area */}}
{{$voicechannelid := 740556133038555246}}
{{$message := "Thanks for bumping us!"}}

{{/* do not edit below (unless you know what you're doing c: ) */}}
{{if .ExecData}}
	{{editChannelName $voicechannelid "Bump Now!"}}
{{else}}
	{{if not (dbGet 0 "bump")}}
		{{dbSetExpire 0 "bump" 1 7200}}
		{{editChannelName $voicechannelid (print "Next Bump in " (((dbGet 0 "bump").ExpiresAt.Sub currentTime).Round .TimeSecond))}}
		{{execCC .CCID nil 7200 "data"}}
		{{sendMessage nil (cembed "title" "Bump!" "description" $message)}}
	{{end}}
{{end}}
