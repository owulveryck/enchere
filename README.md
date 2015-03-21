# About

This is a personnal test to write a code in go.
This is a "lab" where I test the functionnality I learn.

# The code
What it does:

Connect to the website, get the json in a seperate goroutine, and pass it to a goroutine that writes the data back in a database

# Howto use it
    go get github.com/owulveryck/enchere
    go run cmd/enchere.go -user=username -password=PASSWORD -baseurl=the.url.of.the.site
