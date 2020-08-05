{{/*
	Made by DZ#6669 (438789314101379072)

	Trigger Type: RegEx
	Trigger: .*
*/}}

{{/* configuration area */}}
{{$channelid := 740547487655526400}}

{{/* do not edit below (unless you know what you're doing c: ) */}}
{{$db:=dbGet .Channel.ID "stickymessage"}}
{{if eq .Channel.ID $channelid}}
	{{if $db}}
		{{deleteMessage nil (toInt $db.Value) 0}}
	{{end}}
	{{if not (reFind `\A!d\sbump(\s+|\z)` .Message.Content)}}
		{{deleteTrigger 0}}
	{{end}}
	{{$id := sendMessageRetID nil (cembed "title" "Bump This Server!" "description" "Please bump this server by typing `!d bump`" "color" 123456)}}{{dbSet .Channel.ID "stickymessage" (str $id)}}
{{end}}
