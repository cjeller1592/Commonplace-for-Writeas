package main

import (
	"fmt"
	"github.com/advancedlogic/GoOse"
	"go.code.as/writeas.v2"
)

func main() {

	// Log in as a Write.as User

	c := writeas.NewClient()

	// Put in your Write.as username and password to log in and get an auth token
	u, err := c.LogIn("username", "password")
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		access_token := u.AccessToken
		c.SetToken(access_token)
	}

	// Put in the url of the article/blog post you want to put on Write.as
	// Use GoOse to extract the title and content from the the article/blog post
	// Play around with it - some urls work better than others with output

	fmt.Print("Enter url here: ")
	var input string
	fmt.Scanln(&input)
	url := fmt.Sprintf(input)

	g := goose.New()
	article, err := g.ExtractFromURL(url)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		title := article.Title
		content := article.CleanedText

		// Feed the title and content into a new Write.as post
		// The url of the new post will show in the cmd line
		// 

		p, err := c.CreatePost(&writeas.PostParams{
			Title:   title,
			Content: content + "\n\n [Link to Source](" + url + ")",
			Collection: "commonplace", // Put in the collection (blog) name you want to publish the post in
		})
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			fmt.Println("Check out the article: https://write.as/" + p.ID)
		}

	}

}
