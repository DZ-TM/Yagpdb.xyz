{{/*

MADE BY: DZ (438789314101379072)
EDITTED BY: mangowhite (209208777365389312)
SUGGESTED TRIGGER: howgay

DESCRIPTION: This CC gives the user an amount of Gayness that will attach to them like a parasite that will never go away. But that's not all... it will also display all the most gays in the server in a Top. 

*/}}

{{/*VARIABLES*/}}

{{$howGay:=randInt 0 101}}
{{$po:=0}}
{{$color:=randInt 16777216}}
{{$user:= ""}}{{with .Message.Mentions}}{{$user = userArg (index . 0)}}{{end}}
{{$embed:= sdict "color" $color "title" "How Gay"}}


{{/*CODE*/}}
{{if $user}}
	{{range .Guild.Roles}}
		{{if and (in (getMember $user.ID).Roles .ID) .Color (lt $po .Position)}}
			{{$po = .Position}}
			{{$color = .Color}}
		{{end}}
	{{end}}
	{{if not ($db:=dbGet $user.ID "howGay")}}
		{{dbSet $user.ID "howGay" $howGay}}
	{{else if $db}}
		{{$list:= ""}}
		{{$num:= 0}}
		{{range (dbTopEntries "howGay" 10 0)}}
			{{$num = add $num 1}}
			{{$list = print $list $num " • " .User.Username " - " .Value "%\n"}}
		{{end}}
		{{$embed.Set "fields" (cslice (sdict "name" "The Most Gays" "value" $list))}}
		{{$embed.Set "description" (print (or $user.Username (getMember $user.ID).Nick) " is " $db.Value "% gay!  🏳️‍🌈")}}
	{{else if (ge $db 100)}}
		{{$embed.Set "image" (sdict "url" "https://cdn.discordapp.com/attachments/832547182858862602/872972692469461023/unknown.png")}}
	{{end}}
	{{sendMessage nil (complexMessage 
			"content" $user.Mention 
			"embed" (cembed $embed)
		)}}
{{else}}
	❌ You need to input a valid user mention!
{{end}}
