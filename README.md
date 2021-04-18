# furry-broccoli

NOTE: name suggested by github's "new repository" -- I sort of like it :-)

This is a coding exercise for an interview (company name not disclosed).  It
uses the MBTA v3 API to retrieve train schedule information and allow the user
to select route, stop, and direction to get info on the next train leaving the
stop in the given direction.

It was developed using Go 1.16.3, but should work with any version >= 1.14.

If a key for the MBTA API is provided using the environment variable
`MBTA_API_KEY`, it will pass it in each call to the API.

To run, ensure that a relatively new version of Go is installed (I use
[gimme](https://github.com/travis-ci/gimme).)

```
$ eval $(gimme 1.16.3)
$ go build
$ MBTA_API_KEY=myapikey ./furry-brocolli
```

Use the cursor keys to select route, stop, and direction.  The next predicted
departure time will be displayed.  If no departures are predicted with that
info, that will be indicated.

Possible improvements:
- add a way to escape the prompt system (only Ctrl-Z and kill xxx work
  currently).
- use templates and http server to produce as a web service rather than
  terminal
