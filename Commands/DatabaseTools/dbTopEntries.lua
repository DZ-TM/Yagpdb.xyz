{{/*
	Made by DZ#6669 (438789314101379072)

	Trigger Type: Optional
 	Trigger: Optional
*/}}

{{/* configuration area */}}
{{$delay:=10}} {{/* 10s delay */}}
{{$limit:=5}} {{/* times to loop over, must be a positive number */}}

{{/* do not edit below (unless you know what you're doing c: ) */}}
{{$msg:=""}}{{$x:=10}}{{$y:=0}}
{{if and (lt .ExecData.x (mult $limit 10)) .ExecData.x}}
	{{$x =add .ExecData.x 10}}{{$y =add .ExecData.y 10}}
{{end}}
{{range dbTopEntries "%" $x $y}}
	{{- $msg =print $msg .Key ":" .Value "\n" -}}
{{end}}
{{execCC .CCID nil 5 (sdict "x" $x "y" $y)}}{{sendMessage nil $msg}}
