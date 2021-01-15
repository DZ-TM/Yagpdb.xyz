{{/*
	Regex: .*
	Commands:
		-topmessages
		-msgcount
*/}}

{{/* DONT TOUCH UNLESS YOU KNOW WHAT UR DOING c: */}}
{{/* define variables */}}
{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png") }}
{{ $prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (exec "Prefix")) 0 1 }}
{{ $x: = dbIncr .User.ID "msgCount" 1 }}{{ $list:="" }}
{{ $embed := cembed "thumbnail" (sdict "url" $avatar) "title" "Message Counter" "description" (print .User.Mention " have sent " $x  " message(s) ") "color" 0xFF0000 }}
{{ if reFind (print `\A\` $prefix `m(?:essages?|sg)?c(?:ount)?`) .Message.Content }} {{/* msgcount command */}}
	{{ sendMessage nil $embed }}
{{ end }}	
{{ if reFind (print `\A\` $prefix `t(?:op)?m(?:essages?|sg)?`) .Message.Content }} {{/* topmessages command */}}
{{/* better leaderboard, made by joe, modified by blox c: */}}
{{ $page := 1 }} {{/* Default page to start at */}}
{{ $skip := mult (sub $page 1) 10 }} {{/* Amount of entries to skip */}}
{{ $users := dbTopEntries "msgCount" 10 $skip }} {{/* Retrieve the relevant DB entries with the parameters provided */}}
{{ if not (len $users) }}
	There were no users on that page! {{/* If there were no users, return */}}
{{ else }}
	{{ $rank := $skip }} {{/* Instantiate rank variable with value of $skip */}}
	{{ $display := "" }} {{/* The description for the leaderboard description */}}
	{{- range $users -}}
		{{ $coins := toInt .Value }} {{/* Messages for this user entry */}}
		{{ $rank = add $rank 1 }} {{/* Increment rank variable */}}
		{{ $display = printf "%s\n• [%s](https://yagpdb.xyz) :: %d Messages"
			$display .User.String $coins
		}} {{/* Format this line */}}
	{{- end -}}
	{{ $id := sendMessageRetID nil (cembed
		"title" (print "❯ Messages leaderboard")
		"thumbnail" (sdict "url" "https://i.imgur.com/mJ7zu6k.png")
		"color" 14232643
		"description" $display
		"footer" (sdict "text" (joinStr "" "Page " $page))
	) }} {{/* Construct and send the embed */}}
	{{ addMessageReactions nil $id "⬅" "➡" }} {{/* Add reactions for pagination */}}
{{ end }}
