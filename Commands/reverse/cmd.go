{{/*
	Made by DZ#6669 (438789314101379072)
	Trigger Type: Command
	Trigger: username
*/}}
{{/* configuration area */}}
{{$fail:=sdict "color" 14565697}}
{{$success:=sdict "color" 586350}}

{{/* do not edit below */}}
{{$result:=""}}
{{define "send"}}
	{{range $k,$v:=.change}}
		{{- $.embed.Set $k $v -}}
	{{end}}
	{{sendMessage nil (cembed .embed)}}
{{end}}
{{if .StrippedMsg}}
	{{range $i,$_:=($split:=split .StrippedMsg "")}}
		{{- $result =print $result (sub (len $split) 1 $i|index $split) -}}
	{{end}}
	{{template "send" (sdict
		"embed" $success
		"change" (sdict
			"title" "Reversed Text:"
			"description" $result
		)
	)}}
{{else}}
	{{template "send" (sdict
		"embed" $fail
		"change" (sdict
			"title" "No Text Entered"
			"description" "Please enter some arguments to reverse.\n```\n-reverse [Text]\n```"
		)
	)}}
{{end}}
