{{/*
	Made by DZ#6669 (438789314101379072)

	Trigger Type: RegEx
	Trigger: \A!d\sbump(\s+|\z)
*/}}

{{/* configuration area */}}
{{$voicechannelid := 740556133038555246}}
{{$bumpchannelid := 740547487655526400}}

{{/* do not edit below (unless you know what you're doing c: ) */}}
{{if .ExecData}}
	{{editChannelName $voicechannelid "Bump Now!"}}
{{else}}
	{{if and (eq .Channel.ID $bumpchannelid) (not (dbGet 0 "bump"))}}
		{{dbSetExpire 0 "bump" 1 86400}}{{editChannelName $voicechannelid (print "Next Bump in " (((dbGet 0 "bump").ExpiresAt.Sub currentTime).Round .TimeSecond))}}{{execAdmin "giverep" .User}}{{execCC .CCID nil 86400 "data"}}
	{{end}}
{{end}}
