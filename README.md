# wan.do in go
This is a little play to create an url shortener.

This is the desired list of requirements:

- [ ] Publish a port that listen to requests
- [ ] Have a path to redirect requests giving a 302 status code given an id matching an id, requests 
that doesn't match an id will get a 404 (maybe a pretty 404 status page?) 
- [ ] Have a path to receive info and save the data giving a 201 if the information is valid.
- [ ] Connect to a RDBMS
- [ ] Connect to a NoSQL