{{/*
	 Made by DZ#6669 (438789314101379072)
	 Inspired by Maventine#0001 (480561240813600768)

	 Trigger Type: RegEx
	 Trigger: \A(?:-\s?|<@!?204255221017214977>\s*)hug(?:gies?|gy|ged)?(?:\s+|\z)
*/}}

{{/* do not edit below (unless you know what you're doing c: ) */}}
{{$desc:=""}}{{$title:=""}}{{$user:=""}}{{$footer:=print .User.Username " has hugged " (dbIncr .User.ID "gaveHug" 1) " times!"}}{{$color:=123456}}{{$position:=0}}{{$hugImages:=cslice "https://media.discordapp.net/attachments/725361612441125003/741324135606780045/image0.jpg" "https://cdn.discordapp.com/attachments/725361612441125003/741324402641338468/image0.gif" "https://media.discordapp.net/attachments/725361612441125003/741324493993148436/image0.jpg" "https://media.discordapp.net/attachments/725361612441125003/741324706459942912/image0.gif" "https://media.discordapp.net/attachments/725361612441125003/741324922776715294/image0.gif" "https://cdn.discordapp.com/attachments/725361612441125003/741325176964382842/image0.gif"}}
{{range .Guild.Roles}}
	{{- if and (in $.Member.Roles .ID) .Color (lt $position .Position)}}
		{{- $position =.Position}}{{$color =.Color}}
	{{- end -}}
{{end}}
{{if .CmdArgs}}
	{{$user =userArg (index .CmdArgs 0)}}
	{{if $user}}
		{{if gt (len .CmdArgs) 1}}
			{{$desc =joinStr " " (slice .CmdArgs 1)}}
		{{end}}
		{{if ne $user.ID .User.ID}}
			{{$title =print .User.Username " gave " $user.Username " a huggie!"}}{{$footer =print $footer "\n" $user.Username " got hugged " (dbIncr $user.ID "gotHugged" 1) " times!"}}
		{{else}}
			{{$title =print .User.Username " gave themselves a huggie!"}}
		{{end}}
	{{else}}
		{{$title =print .User.Username " gave themselves a huggie!"}}{{$desc =.StrippedMsg}}
	{{end}}
{{else}}
	{{$title =print .User.Username " gave themselves a huggie!"}}
{{end}}
{{sendMessage nil (cembed "title" $title "description" $desc "image" (sdict "url" (index (shuffle $hugImages) 0)) "color" $color "footer" (sdict "text" $footer))}}
