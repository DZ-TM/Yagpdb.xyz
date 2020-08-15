{{/*
	Made by DZ#6669 (438789314101379072)

	Trigger Type: Interval
	Time: 10m
*/}}

{{/* configuration area */}}
{{$channelid := 740556133038555246}}{{/* ID of the channel you want it to edit the name of, recommended to use a voice channel */}}

{{/* do not edit below (unless you know what you're doing c: ) */}}
{{$db:=dbGet 0 "bump"}}
{{if $db}}
	{{editChannelName $channelid (print "Next Bump in " (($db.ExpiresAt.Sub currentTime).Round .TimeSecond))}}
{{else}}
	{{editChannelName $channelid "Bump Now!"}}
{{end}}
