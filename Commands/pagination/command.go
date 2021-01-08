{{/*
	Made by DZ#6669 (438789314101379072)
	Trigger Type: Command
	Trigger of your choice.
 
	NOTE: Read README.md for instructions on configuration.
*/}}

{{/* configuration area */}}
{{$embed := sdict
	"title" "Title A"
	"description" "Description A"
	"image" (sdict "url" "https://m.hindustantimes.com/rf/image_size_960x540/HT/p2/2018/05/16/Pictures/_1571873a-58de-11e8-b431-73159b4b09e2.jpg")
	"color" 123456
 }}{{/* first element of the slice in the reactionListener */}}

{{/* do not edit below */}}
{{$embed.Set "description" (print (or $embed.description "") "[](0)")}}
{{$mID:=sendMessageRetID nil (cembed $embed)}}
{{addMessageReactions nil $mID "⬅️" "➡️"}}
