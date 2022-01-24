# Golang 연습 2021-12-31~
## **Golang 문법**   

> ## 2022-01-24정리   
> echo를 사용하여 web server 동작시키기 
> ```go
>     package main
>     
>     import (
>     "strings"
>     
>     "github.com/Pigonhair/go-learning/scrapper"
>     "github.com/labstack/echo"
>     )
>     
>     func handleHome(c echo.Context) error {
>         return c.File("home.html")
>     }
>
>     func handleScrape(c echo.Context) error {
>         term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
>         return nil
>     }      
>     
>     func main() {
>         e := echo.New()
>         e.GET("/", handleHome)
>         e.POST("/scrape", handleScrape)
>         e.Logger.Fatal(e.Start(":1323"))
>     }
>   ```
> home.html
> ```html
>    <!DOCTYPE html>
>    <html lang="en">
>  
>     <head>
>         <meta charset="UTF-8">
>         <meta http-equiv="X-UA-Compatible" content="IE=edge">
>         <meta name="viewport" content="width=device-width, initial-scale=1.0">
>         <title>Go Jobs</title>
>     </head>
>
>     <body>
>         <h1>Go Jobs</h1>
>         <h3>Indeed.com scrapper</h3>
>         <form method="POST" action="/scrape">
>             <input placeholder="what job do u want" name="term" />
>             <button>Search</button>
>         </form>
>     </body>
>
>     </html>
>   ```
