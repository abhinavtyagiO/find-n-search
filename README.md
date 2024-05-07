
# About

A simple GO program that gets the JSON data from the comics API `https://xkcd.com/%d/info.0.json` and stores it in a file and get the data that matches with the given search terms.

`get-comics/main.go` takes in a file name as an argument and stores the JSON data read from the API.

```
$ go run  ./get-comics xkcd.json
Read 2927 comics
```


`find-comics/main.go` reads some command line arguments as the search terms and find comics whose title or transcript matches all the words.

```
$ go run  ./find-comics xkcd.json sleep run 
https://xkcd.com/124/ 5/7/2006 Blogofractal
https://xkcd.com/574/ 27/4/2009 Swine Flu
https://xkcd.com/658/ 4/11/2009 Orbitals
https://xkcd.com/1023/ 29/2/2012 Late-Night PBS
Found 4 comics
```

