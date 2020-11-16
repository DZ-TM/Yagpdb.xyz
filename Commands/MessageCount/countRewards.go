{{/*
    Trigger Type: Regex
    Trigger     : .*
*/}}

{{$rewards := dict 50 ROLE-ID 75 ROLE-ID 100 ROLE-ID 125 ROLE-ID 150 ROLE-ID}}
{{/* The rewards count goes from 50 messages and upwards, you can change/add to this if you want */}}

{{if ($db := dbGet .User.ID "msgCount")}}
{{/* Checks if they have a message count */}}

    {{range $k, $v := $rewards -}}
    {{- /* Ranges the KEY:VALUE pairs of the rewards */}}

        {{- if and (ge (toInt $db.Value) $k) (not (hasRoleID $v)) -}}
        {{- /* Checks if the user has a message count >= each KEY of the rewards AND that they do not already have that role reward */}}

            {{- addRoleID $v -}}
            {{/* Adds the role(s) corresponding to the VALUE of the rewards if true */}}

			      {{- print "Congrats! You messaged **" $db.Value "** times and earned the <@&" $v "> role!\n" -}}
			      {{/* Sends the congrats message */}}

        {{- end -}}

    {{- end}}
    
{{end}}
