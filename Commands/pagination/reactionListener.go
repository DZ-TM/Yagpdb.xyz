{{/*
	Made by DZ#6669 (438789314101379072)
	Trigger Type: Reaction
	Trigger: Added Only

	NOTE: Read README.md for instructions on configuration.
*/}}

{{/* configuration area */}}
{{$isWrap:=true}}{{/* if this is set to true, this makes it so if the pagination is on the final page, it moves to the first page assuming the 'right_arrow' reaction is clicked or if it is on the first page and the 'left_arrow' reaction is clicked, it moves to the final page */}}
{{$delete:=cslice "image"}}{{/* this means that if there was, for example an image field under the first embed and not in another stage, it will not retain the image / it deletes it per run */}}
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
}}{{/* embeds within a slice where the embeds are set as sdicts, the pages are in order of the slice e.g. the first element in the slice is the first page */}}

{{/* do not edit below */}}
{{if and .ReactionAdded .Message.Embeds (eq .Message.Author.ID 204255221017214977)}}
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
		{{- $embed.Del (lower .|title) -}}
	{{end}}
	{{range $i,$_:=$pages}}
		{{- (index $pages $i).Set "description" (print (or .description "") "[](" $i ")")}}
	{{end}}
	{{if $pageNum:=reFindAllSubmatches `\[\]\((\d+)\)` (index .Message.Embeds 0).Description}}
		{{$pageNum:=len $pages|mod (index $pageNum 0 1)|toInt}}
		{{deleteMessageReaction nil .Message.ID .User.ID "⬅️" "➡️"}}
		{{if eq .Reaction.Emoji.Name "⬅️"}}
			{{if $isWrap}}
				{{if eq $pageNum 0}}
					{{$pageNum =len $pages|add -1}}
				{{else}}
					{{$pageNum =sub $pageNum 1}}
				{{end}}
			{{else if gt $pageNum 0}}
				{{$pageNum =sub $pageNum 1}}
			{{end}}
		{{else if eq .Reaction.Emoji.Name "➡️"}}
			{{if or $isWrap (lt $pageNum (len $pages|add -1))}}
				{{$pageNum =add $pageNum 1}}
			{{end}}
		{{end}}
		{{$pageNum =len $pages|mod $pageNum|toInt}}
		{{range $k, $v:=index $pages $pageNum}}
			{{- $embed.Set $k $v -}}
		{{end}}
		{{editMessage nil $.Message.ID (cembed $embed)}}
	{{end}}
{{end}}
