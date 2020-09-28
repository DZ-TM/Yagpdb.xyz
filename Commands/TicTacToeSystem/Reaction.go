{{/*
	Made by DZ#6669 (438789314101379072)
	Trigger Type: Reactions
	Trigger: Added Reactions Only
*/}}
 
{{/* configuration area */}}
{{$delay:=3600}}{{/*time in seconds of which, if no response is received, ends the game*/}}
{{$emotes:=cslice "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣" "6️⃣" "7️⃣" "8️⃣" "9️⃣"}}{{/*two things to keep in mind 1) use 9 DIFFERENT emojis and 2) of emojis being used, use on of three formats by discord 1. for animated emotes <a:example:123456789> 2. for custom emotes <:example:123456789> 3. for normal emojis, use their literal emotes or their unicodes*/}}
{{$players:=cslice "❌" "⭕️"}}{{/* emoji for players, follows same format as $emotes */}}

{{/* do not edit below */}}
{{$emoIndex:=dict "x" (index $players 0) "o" (index $players 1)}}{{$emoteCheck:=cslice}}{{$emoteData:=cslice}}{{$desc:=""}}{{$fURL:=""}}{{$fText:=""}}{{$win:=0}}
{{range $i,$_:=$emotes}}
	{{- $emoIndex.Set $i (reFind `(?:[\w~]{2,32}|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(?:\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` $_)}}
	{{- $emoteCheck =$emoteCheck.Append (print (reFind `(?:[\w~]{2,32}|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(?:\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` $_) ":" (or (reReplace `>\z` (reFind `\d{17,19}>\z` $_) "") "empty"))}}
	{{- $emoteData =$emoteData.Append (reFind `(?:[\w~]{2,32}:\d{17,19}|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(?:\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` $_) -}}
{{end}}
{{if and .ReactionAdded (in $emoteCheck (print .Reaction.Emoji.Name ":" (or .Reaction.Emoji.ID "empty")))}}
	{{if $data:=dbGet .Message.ID "ttt"}}
		{{$data =sdict $data.Value}}
		{{$embed:=index .Message.Embeds 0|structToSdict}}
		{{$data.Set "steps" ((cslice).AppendSlice $data.steps)}}
		{{if eq $.User.ID ($data.Get $data.turn).ID}}
			{{range $i,$_:=$data.steps}}
				{{- if not $i}}
					{{- $data.Set "steps" cslice}}
				{{- end}}
				{{- if and ($emoIndex.Get $i|eq $.Reaction.Emoji.Name) (eq $_ "empty")}}
					{{- $data.Set "steps" ($data.steps.Append $data.turn)}}
					{{- $data.Set "turn" ((sdict "x" "o" "o" "x").Get $data.turn)}}
				{{- else}}
					{{- $data.Set "steps" ($data.steps.Append $_)}}
				{{- end -}}
			{{end}}
			{{range $k,$v:=$embed}}
				{{- if eq (kindOf $v true) "struct"}}
					{{- $embed.Set $k (structToSdict $v)}}
				{{- end -}}
			{{end}}
			{{range cslice (cslice 0 1 2) (cslice 3 4 5) (cslice 6 7 8) (cslice 0 3 6) (cslice 1 4 7) (cslice 2 5 8) (cslice 0 4 8) (cslice 2 4 6)}}
				{{- if and (eq (index . 0|index $data.steps) (index . 1|index $data.steps)) (eq (index . 1|index $data.steps) (index . 2|index $data.steps)) (eq (index . 0|index $data.steps) (index . 2|index $data.steps)) (eq (index . 0|index $data.steps) "x" "o")}}
					{{- $win =1}}
				{{- end -}}
			{{end}}
			{{if $win}}
				{{$fText =print (or .Member.Nick .User.Username)  " won!"}}
				{{dbDel .Message.ID "ttt"}}
				{{$fURL =.User.AvatarURL "256"}}
			{{else if not (in $data.steps "empty")}}
				{{$fText ="Game resulted in a draw!"}}
				{{dbDel .Message.ID "ttt"}}
				{{$fURL =print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".png"}}
			{{else}}
				{{dbSetExpire .Message.ID "ttt" $data $delay}}
				{{$fText =print (or (getMember ($data.Get $data.turn).ID).Nick ($data.Get $data.turn).Username) "'s turn"}}
				{{$fURL =(userArg ($data.Get $data.turn).ID).AvatarURL "256"}}
			{{end}}
			{{with $embed.Footer}}
				{{.Set "icon_url" $fURL}}
				{{.Set "text" $fText}}
			{{end}}
			{{with $embed.Author}}
				{{.Set "icon_url" .IconURL}}
			{{end}}
			{{range $i,$_:=$data.steps}}
				{{- $item:=""}}
				{{- if eq $_ "empty"}}
					{{- $item =index $emotes $i}}
				{{- else}}
					{{- $item =$emoIndex.Get $_}}
				{{- end -}}
				{{- if mod (add $i 1) 3}}
					{{- $desc =print $desc $item "\t"}}
				{{- else}}
					{{- $desc =print $desc $item "\n"}}
				{{- end -}}
			{{end}}
			{{$embed.Set "description" $desc}}
			{{editMessage nil .Message.ID (cembed $embed)}}
		{{end}}
	{{end}}
{{end}}
