package main

import (
	"github.com/k0kubun/pp"
	"github.com/kevincobain2000/go-gibberish/gibberish"
)

func main() {

	raw := `It’s hard to understate how shocked and happy I am that Andor exists. I have been banging a drum for years that Disney has been putting out programming that often appropriates the aesthetic of social change and revolution while advancing pretty regressive narratives (see my take on Black Panther and She-Hulk as examples).

	Yet with Andor, we have a show that is saying something explicit about the need for direct action in fighting fascism without pulling any punches. It is an earnest text that covers a lot of ground, and like every commenter with half a brain, the fact that the Disney corporation greenlit it is shocking to me. We get a show that depicts fascism as it actually is, and that is sadly too rare in pop culture.`

	j := gibberish.NewGibberish()
	r := j.Detect(raw)

	pp.Println(r)
}
