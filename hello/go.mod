module vimday.com/hello

go 1.14

require (
	rsc.io/quote v1.5.2
	vimday.com/greetings v0.0.0-00010101000000-000000000000
)

replace vimday.com/greetings => ../greetings
