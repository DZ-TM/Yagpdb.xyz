{{$howGay:= randInt 0 101}} 
{{if eq .User.ID 495426188517572629}}
{{$howGay = 100}}
{{end}}
{{$po := 0}} {{$co := randInt 16777216}}
{{$user := .User}} {{$member := .Member}}
{{if .Message.Mentions}}
	{{$user = index .Message.Mentions 0}} {{$member = getMember $user}}
{{end}}
{{range .Guild.Roles}}
	{{- if and (in $member.Roles .ID) .Color (lt $po .Position)}}
		{{- $po =.Position}} {{$co =.Color}}
	{{- end -}}
{{end}}
{{sendMessage nil (complexMessage "content" $user.Mention "embed" (cembed "color" $co "title" "gay r8 machine" "description" (print $user.Username " is " $howGay "% gay 🏳️‍🌈")))}}
