{{/*
	This is the pagination code for the message leaderboard
  
	Recommended trigger: added reactions only
	Trigger type: reaction
*/}}
{{/* dont touch */}}
{{ $action := .Reaction.Emoji.Name }} {{/* The action being ran */}}
{{ $validEmojis := cslice "➡" "⬅" }} {{/* Valid emojis */}}
{{ $isValid := false }} {{/* Whether this is actually a valid embed / leaderboard embed */}}
{{ $page := 1 }} {{/* The current page */}}
{{ with and (eq .ReactionMessage.Author.ID 204255221017214977) .ReactionMessage.Embeds }} {{/* Checks for validity */}}
	{{ $embed := index . 0 }} {{/* The first embed */}}
	{{ if and (eq $embed.Title (print "❯ Messages leaderboard")) $embed.Footer }} {{/* More checks */}}
		{{ $page = reFind `\d+` $embed.Footer.Text }} {{/* We presume that this is valid, and get the page num */}}
		{{ $isValid = true }} {{/* Yay, it is valid */}}
		{{ deleteMessageReaction nil $.ReactionMessage.ID $.User.ID $action }}
	{{ end }}
{{ end }}
{{ if and (in $validEmojis $action) $isValid $page }} {{/* Even more checks for validity... */}}
	{{ if eq $action "➡" }} {{ $page = add $page 1 }} {{/* Update page according to emoji */}}
	{{ else if eq $action "⬅"}} {{ $page = sub $page 1 }} {{ end }}
	{{ if ge $page 1 }} {{/* Otherwise, dbTopEntries throws error due to negative skip */}}
		{{ $skip := mult (sub $page 1) 10 }} {{/* Get skip */}}
		{{ $users := dbTopEntries "msgCount" 10 $skip }} {{/* Fetch entries */}}
		{{ if (len $users) }} {{/* If there are users on this page, proceed */}}
			{{ $rank := $skip }}
			{{ $display := "" }} {{/* Display for leaderboard embed */}}
			{{- range $users -}} {{/* Loop over users and format */}}
				{{ $coins := toInt .Value }} {{/* The user XP */}}
				{{ $rank = add $rank 1 }} {{/* The user rank */}}
				{ $display = printf "%s\n• [%s](https://yagpdb.xyz) :: %d Messages"
            $display .User.String $coins
        }} {{/* Format display */}}
			{{- end -}}
			{{ editMessage nil .ReactionMessage.ID (cembed
				"title" "❯ Messages leaderboard"
				"thumbnail" (sdict "url" "https://i.imgur.com/mJ7zu6k.png")
				"color" 14232643
				"description" $display
				"footer" (sdict "text" (joinStr "" "Page " $page))
						) }} {{/* Edit embed */}}
		{{ end }}
	{{ end }}
{{ end }}
