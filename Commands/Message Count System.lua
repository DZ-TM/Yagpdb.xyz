{{/*
	Regex: .*
	Commands:
		-topmessages
		-msgcount
*/}}

{{$x:=dbIncr .User.ID "msgCount" 1}}{{$list:=""}}
{{if reFind `\A-m(?:essages?|sg)?c(?:ount)?` .Message.Content}}
	 You have typed {{$x}} messages.
{{end}}
{{if reFind `\A-t(?:op)?m(?:essages?|sg)?` .Message.Content}}
	{{range dbTopEntries "msgCount" 10 0}}
		{{- $list =print $list .User.Username " - " .Value "\n" -}}
	{{end}}
	{{sendMessage nil (cembed "title" "Message Count Leaderboard" "description" $list)}}
{{end}}
