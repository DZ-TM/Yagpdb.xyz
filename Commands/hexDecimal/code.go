{{/*
	Made by DZ#5559 (438789314101379072)

	Trigger Type: RegEx
	Trigger: \A(-|<@!?204255221017214977>)\s*(h(?:ex(?:adecimal)?)?|d(?:ecimal)?)(\s+|\z)
*/}}
{{/* configuration area */}}
{{$fail:=sdict "color" 14565697}}
{{$success:=sdict "color" 586350}}

{{/* do not edit below */}}
{{$decimal:=0}}{{$dict:=dict}}
{{define "change"}}
	{{range $k,$v:=.change}}
		{{- if $v}}
			{{- $.embed.Set $k $v}}
		{{- end -}}
	{{end}}
	{{sendMessage nil (cembed .embed)}}
{{end}}
{{if .CmdArgs}}
	{{if in .Cmd "h"}}
		{{if $data:=index .CmdArgs 0|reFind `\d+`}}
			{{if and (ge ($data =toInt $data) 0) (le $data 16777216)}}
				{{$success.Set "color" $data}}
			{{end}}
			{{template "change" (sdict
				"embed" $success
				"change" (sdict
					"title" "Hex Result:"
					"description" (printf "%X" $data)
					"footer" (sdict
						"icon_url" (.User.AvatarURL "1024")
						"text" (print "Decimal: " $data)
					)
				)
			)}}
		{{else}}
			{{template "change" (sdict "embed" $fail "change" (sdict "title" "Incorrect Usage of Command" "description" "The command was used incorrectly.\n```\n-hex <int>\n```"))}}
		{{end}}
	{{else}}
		{{if $data:=index .CmdArgs 0|reFind `0*[A-Fa-f\d]{1,6}`|upper}}
			{{range len ($split:=split $data "")|seq 0}}
				{{- $decimal =add $decimal (mult (toInt ((sdict "0" 0 "1" 1 "2" 2 "3" 3 "4" 4 "5" 5 "6" 6 "7" 7 "8" 8 "9" 9 "A" 10 "B" 11 "C" 12 "D" 13 "E" 14 "F" 15).Get (sub (len $split) 1 .|index $split|str))) (pow 16 .)) -}}
			{{end}}
			{{template "change" (sdict
				"embed" $success
				"change" (sdict
					"title" "Decimal Result:"
					"description" (str $decimal)
					"footer" (sdict
						"icon_url" (.User.AvatarURL "1024")
						"text" (print "Hex: " (reReplace `^0*` $data ""))
					)
					"color" $decimal
				)
			)}}
		{{else}}
			{{template "change" (sdict "embed" $fail "change" (sdict "title" "Incorrect Usage of Command" "description" "The command was used incorrectly.\n```\n-decimal <hex>\n```"))}}
		{{end}}
	{{end}}
{{else}}
	{{template "change" (sdict "embed" $fail "change" (sdict "title" "Insufficient Args" "description" "Not enough arguments were provided.\n```\n-hex <int>\n-decimal <hex>\n```"))}}
{{end}}
