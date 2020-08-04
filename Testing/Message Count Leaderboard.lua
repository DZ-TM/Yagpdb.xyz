{{$list:=""}}
{{range dbTopEntries "msgCount" 10 0}}
	{{- $list =print $list .User.Username " - " .Value "\n" -}}
{{end}}{{$list}}
