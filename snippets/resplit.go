{{/*
    This is a snippet of code to mimic a reSplit function with the syntax "reSplit String RegEx".
    The code to create the function has already been PRed, but due to the feature freeze it is yet to be added to the bot.

    Note:
        To make it work, you need to input both 1) the string and 2) the regular expression.
        This is done through a sdict so we can retrieve the output later.
        Look at $sdict for an example of input to the template we make.

    How it works:
        It gets a unique character that is not within the string itself.
        It replaces all matches of the string by that unique character.
        It splits the string by that character.

    Made by DZ#6669 (438789314101379072)
*/}}

{{define "reSplit"}}
    {{$runeStr := toRune .string}}
    {{$specialChar := ""}}
    {{$cslice := seq 0 (len $runeStr|add 1)}}{{/* int slice of incrementing numbers from 0 to the length of the string, this means there will always be at least one unique character / rune which is not in the string itself */}}
    {{range $cslice}}
        {{- if not (in $runeStr .)}}
            {{- $specialChar = printf "%c" .}}{{/* we use printf to get the symbol representation of the character, rather than its rune so we can use it later on */}}
        {{- end -}}
    {{end}}
    {{.Set "result" (split (reReplace .regex .string $specialChar) $specialChar)}}
{{end}}
{{$sdict := sdict "string" .StrippedMsg "regex" `\n+`}}
{{template "reSplit" $sdict}}
{{json $sdict.result}}{{/* yay output */}}
