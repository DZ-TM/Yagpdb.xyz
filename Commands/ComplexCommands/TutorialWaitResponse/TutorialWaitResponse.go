{{/*
	Made by DZ#6669 (438789314101379072)
	Trigger Type: RegEx
	Trigger: .*
*/}}

{{/* configuration area */}}
{{$trigger:=`-trigger`}} {{/* the trigger is set to -trigger, basic regex knowledge is required to edit this */}}
{{$timer:=60}} {{/* 60s timer, if timer ends with no response, the command is automatically disabled */}}

{{/* DO NOT EDIT BELOW (unless you want to learn how it works or know what you're doing c: ) */}}
{{$databaseValue:=toInt (dbGet .User.ID "waitResponse").Value}}
{{$dontChangeStage:=0}}
{{$changeStage:=0}}
{{$colour:=123456}}
{{$position:=0}}

{{/* sets colour to role colour */}}
{{range .Guild.Roles}}
	{{- if and (in $.Member.Roles .ID) .Color (lt $position .Position)}}
		{{- $position =.Position}}{{$colour =.Color}}
	{{- end -}}
{{end}}

{{/* sets default value for $embed */}}
{{$embed:=sdict "author" (sdict "icon_url" (.User.AvatarURL "256") "name" (print "User: " .User.String)) "footer" (sdict "text" "Type cancel to cancel the tutorial.") "color" $colour}}

{{/* checks if it was executed by scheduleUniqueCC */}}
{{if .ExecData}}

	{{/* sets timed out embed message */}}
	{{$embed.Set "title" "Tutorial Timed Out!"}}
	{{$embed.Set "description" (print "You have not entered a response after " $timer " seconds. As such, the tutorial has been cancelled.")}}

	{{/* sends message */}}
	{{sendMessage nil (cembed $embed)}}

{{/* checks if it was not executed by scheduleUniqueCC */}}
{{else}}

	{{/* if no database */}}
	{{if not $databaseValue}}

		{{/* checks if message matches regex to begin tutorial */}}
		{{if reFind (print `\A(?i)` $trigger `(\s+|\z)`) .Message.Content}}

			{{/* sets embed for tutorial */}}
			{{$embed.Set "title" "Activating Tutorial"}}
			{{$embed.Set "description" "Please enter a **positive** number **below 100**."}}
			{{sendMessage nil (cembed $embed)}}

			{{/* sets $changeStage to true for usage later and replaces delay for "cancelled" with $timer */}}
			{{$changeStage =1}}
			{{scheduleUniqueCC .CCID nil $timer (print "cancelled " .User.ID) 1}}
		{{end}}

	{{/* if $databaseValue has a value */}}
	{{else}}

		{{/* checks if $databaseValue is set to the first stage */}}
		{{if eq $databaseValue 1}}

			{{/* checks if the entire message content is a number */}}
			{{with toInt .Message.Content}}

				{{/* checks if the number is a positive number below 100 */}}
				{{if and (gt . 0) (le . 100)}}

					{{/* sets tutorial to seconds stage */}}
					{{$embed.Set "title" "Stage 2"}}
					{{$embed.Set "description" (print "Thank you for your input of `" . "`\n\nPlease enter a sentence with **more than 3 words and more than 25 characters**.")}}

					{{/* sets $changeStage to true for usage later */}}
					{{$changeStage =1}}

				{{/* if user inputted incorrect details */}}
				{{else}}

					{{/* gives error message for incorrect input for stage 1 */}}
					{{$embed.Set "title" "Incorrect Input"}}
					{{$embed.Set "description" "Please enter a **positive** number **below 100**."}}

					{{/* sets $changeStage to false for usage later */}}
					{{$changeStage =0}}
				{{end}}

			{{/* if not number inputted */}}
			{{else}}

				{{/* gives error message for incorrect input for stage 1 */}}
				{{$embed.Set "title" "Incorrect Input"}}
				{{$embed.Set "description" "Please enter a **positive** number **below 100**."}}


				{{/* sets $changeStage to false for usage later */}}
				{{$changeStage = 0}}
			{{end}}

			{{/* replaces delay for "cancelled" with $timer */}}
			{{scheduleUniqueCC .CCID nil $timer (print "cancelled " .User.ID) 1}}

		{{/* checks if $databaseValue is set to the second stage */}}
		{{else if eq $databaseValue 2}}

			{{/* checks if there is content */}}
			{{if .Args}}

				{{/* checks if there are at least 3 words and 25 character */}}
				{{if and (ge (len .Args) 3) (ge (len (toRune .Message.Content)) 25)}}

					{{/* sets tutorial to third stage */}}
					{{$embed.Set "title" "Stage 3"}}
					{{$embed.Set "description" (print "Thank you for your input of `" .Message.Content "`\n\nPlease enter `finished` to complete your tutorial.")}}

					{{/* sets $changeStage to true for usage later */}}
					{{$changeStage =1}}

				{{/* if user inputted incorrect data */}}
				{{else}}

					{{/* sets tutorial to third stage */}}
					{{$embed.Set "title" "Incorrect Input"}}
					{{$embed.Set "description" "Please enter at least 3 words and 25 characters."}}

					{{/* sets $changeStage to false for usage later */}}
					{{$changeStage =0}}
				{{end}}

			{{/* if nothing inputted */}}
			{{else}}

				{{/* sets tutorial to third stage */}}
				{{$embed.Set "title" "Incorrect Input"}}
				{{$embed.Set "description" "Please enter at least 3 words and 25 characters."}}

				{{/* sets $changeStage to false for usage later */}}
				{{$changeStage =0}}
			{{end}}

			{{/* replaces delay for "cancelled" with $timer */}}
			{{scheduleUniqueCC .CCID nil $timer (print "cancelled " .User.ID) 1}}

		{{/* checks if $databaseValue is set to the third stage */}}
		{{else if eq $databaseValue 3}}

			{{/* checks if message content is equal to "finished" */}}
			{{if eq (lower .Message.Content) "finished"}}

				{{/* sets tutorial to third stage */}}
				{{$embed.Set "title" "Tutorial Finished!"}}
				{{$embed.Set "description" (print "Thank you for your input of `" .Message.Content "`\n\nYou have officially finished your tutorial.")}}


				{{/* deletes database for "waitResponse" */}}
				{{dbDel .User.ID "waitResponse"}}

				{{/* sets $changeStage to false for later usage and cancels the execution of "cancelled" */}}
				{{$changeStage =0}}
				{{cancelScheduledUniqueCC .CCID (print "cancelled " .User.ID)}}

			{{/* if user inputted incorrect data */}}
			{{else}}

				{{/* error message for last stage */}}
				{{$embed.Set "title" "Incorrect Input"}}
				{{$embed.Set "description" "Please enter `finished` to finish your tutorial."}}

				{{/* sets $changeStage to false for usage later and replaces delay for "cancelled" with $timer */}}
				{{$changeStage =0}}
				{{scheduleUniqueCC .CCID nil $timer (print "cancelled " .User.ID) 1}}
			{{end}}
		{{end}}

		{{/* checks if user inputted cancel */}}
		{{if eq (lower .Message.Content) "cancel"}}

			{{/* sets embed to "tutorial was cancelled message" */}}
			{{$embed.Set "title" "Tutorial was Cancelled"}}
			{{$embed.Set "description" (print (or .Member.Nick .User.Username) " has decided to cancel the tutorial.")}}

			{{/* deletes database for "waitResponse" */}}
			{{dbDel .User.ID "waitResponse"}}

			{{/* sets $changeStage to false for later usage and cancels the execution for "cancelled" */}}
			{{$changeStage =0}}
			{{cancelScheduledUniqueCC .CCID (print "cancelled " .User.ID)}}
		{{end}}
	{{end}}
	{{if and (or $databaseValue (dbGet .User.ID "waitResponse"))}}
	{{/* sends message if database has value, used to make it not spam chat */}}
	{{sendMessage nil (cembed $embed)}}
        {{end}}
{{end}}

{{/* used to change stage to next stage, the reason we use dbSetExpire instead of dbIncr is because dbIncr would still have the same expiration date as the old dbSetExpire, we use dbSetExpire to replace that expiration date */}}
{{if $changeStage}}
	{{dbSetExpire .User.ID "waitResponse" (str (add $databaseValue 1)) $timer}}
{{end}}
