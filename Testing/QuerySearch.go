{{/* -bannedwords search/query <query> */}}

{{$bannedWords:=(cslice).AppendSlice (dbGet 0 "bannedWords").Value}}{{$queryList:=""}}{{$list:=""}}{{$a:=.StrippedMsg}}
{{if $a}}
	{{range $bannedWords}}
		{{- if reFind (print `(?i)(\b|\s*\/)` (reReplace "[^\\w]" $a "") `(/\S*|\S*)`) .}}
			{{- $queryList =joinStr ", " $queryList .}}
		{{- end -}}
	{{end}}
	{{if $queryList}}
		Your search has provided the following word(s): {{$queryList}}.
	{{else}}
		No word(s) containing `{{$a}}` could be found.
	{{end}}
{{else}}
	 Nothing to query was provided.
{{end}}
