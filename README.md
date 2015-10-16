# go-oauth-samples

In those samples, we try to have few external dependencies beside the one from the standard libraries. 
Those code examples are not fully secure because they do not manage the state parameter required to implement OAuth 2 RFC 6749 10.12 CSRF Protection. 
Indeed, you need to setup a (short-lived) state cookie and adds tons of code to handle this process. 
Check out the library of Dalton Hubble [here](https://github.com/dghubble/gologin). I think he did a great job implementing those login processes.
