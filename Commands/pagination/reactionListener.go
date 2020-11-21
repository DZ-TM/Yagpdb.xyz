{{/*
	Made by DZ#6669 (438789314101379072)
	Trigger Type: Reaction
	Trigger: Added Only
 
	NOTE: Read README.md for instructions on configuration.
*/}}
 
{{/* configuration area */}}
{{$delete:=cslice "Image"}}
{{$pages:=cslice
	(sdict
		"title" "Title A"
		"description" "Description A"
		"image" (sdict "url" "https://m.hindustantimes.com/rf/image_size_960x540/HT/p2/2018/05/16/Pictures/_1571873a-58de-11e8-b431-73159b4b09e2.jpg")
	)
	(sdict
		"title" "Title B"
		"description" "Description B"
		"image" (sdict "url" "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/funny-dog-captions-1563456605.jpg?crop=0.747xw:1.00xh;0.0459xw,0&resize=480:*")
	)
	(sdict
		"title" "Title C"
		"description" "Description C"
		"image" (sdict "url" "https://github.com/jigsawpieces/dog-api-images/blob/master/doberman/doberman.jpg?raw=true")
	)
}}
 
{{/* do not edit below */}}
{{if and .ReactionAdded .Message.Embeds}}
	{{$embed:=index .Message.Embeds 0|structToSdict}}
	{{range $k, $v:=$embed}}
		{{- if eq (kindOf $v true) "struct"}}
			{{- $embed.Set $k (structToSdict $v)}}
		{{- end -}}
	{{end}}
	{{with $embed.Author}}
		{{.Set "icon_url" .IconURL}}
	{{end}}
	{{with $embed.Footer}}
		{{.Set "icon_url" .IconURL}}
	{{end}}
	{{range $delete}}
		{{- $embed.Del . -}}
	{{end}}
	{{range $i,$_:=$pages}}
		{{- $_.Set "description" (print (or $_.description "") "[\u200b](" $i ")")}}
		{{- $pages.Set $i $_ -}}
	{{end}}
	{{if $pageNum:=reFindAllSubmatches `\[\p{Cf}\]\((\d+)\)` (index .Message.Embeds 0).Description}}
		{{$pageNum:=index $pageNum 0 1|toInt}}
		{{deleteMessageReaction nil .Message.ID .User.ID "⬅️" "➡️"}}
		{{$pageNum:=add $pageNum (or (and (eq .Reaction.Emoji.Name "⬅️") (gt $pageNum 0) -1) (and (eq .Reaction.Emoji.Name "➡️") (lt $pageNum (add -1 (len $pages))) 1) 0)}}
		{{with index $pages $pageNum}}
			{{range $k, $v:=.}}
				{{- $embed.Set $k $v -}}
			{{end}}
			{{editMessage nil $.Message.ID (cembed $embed)}}
		{{end}}
	{{end}}
{{end}}
