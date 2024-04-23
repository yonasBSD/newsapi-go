package main

import (
    "fmt"
    "net/http"
    "time"
    "context"
    "os"
    "strings"
    "errors"

    "github.com/jellydator/newsapi-go"
)

func main() {

  if len(os.Args) < 2 {
    err := errors.New("Missing search query.")
    fmt.Println(err)
    os.Exit(1)
  }

  client := newsapi.NewClient("29fe881bf1d54b86b4e5003910f3f470", newsapi.WithHTTPClient(&http.Client{
    Timeout: 5 * time.Second,
  }))

  articles, _, err := client.TopHeadlines(context.Background(), newsapi.TopHeadlinesParams{
    Query: strings.Join(os.Args[1:], " "),
  })

  if err != nil {
    // handle error
    panic(err)
  }

  fmt.Println("# News")
  for _, v:= range articles {
    fmt.Printf("## %s\n\n[Link](%s)\n\n%s\n%s\n\n", v.Title, v.URL, v.Description, v.Content)
  }
}
