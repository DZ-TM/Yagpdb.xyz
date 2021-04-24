## Note
`reSplit` has now been added to YAGPDB.xyz, so this snippet of code is no longer needed. Taken from [Golang docs](https://golang.org/pkg/regexp/#example_Regexp_Split):
```
func (re *Regexp) Split(s string, n int) []string
```
> Split slices s into substrings separated by the expression and returns a slice of the substrings between those expression matches.<br><br>
The slice returned by this method consists of all the substrings of s not contained in the slice returned by FindAllString. When called on an expression that contains no metacharacters, it is equivalent to strings.SplitN.
### How does it work?
###### Syntax of reSplit
```
reSplit RegEx string max
```
The regex must be any valid RegEx in Golang flavour of RegEx.

Max must be an integer, if it is -1 or equivalent, it returns all results, otherwise the result is limited (in the amount of elements it can have in the slice) by the max.

reSplit returns a string slice.
