# go-junior

## Lesson materials


### Tips for printing using fmt:

#### General:
- %v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names

- %#v	a Go-syntax representation of the value
	(floating-point infinities and NaNs print as ±Inf and NaN)

- %T	a Go-syntax representation of the type of the value

- %%	a literal percent sign; consumes no value

- %s	the uninterpreted bytes of the string or slice

#### Boolean:
- %t	the word true or false