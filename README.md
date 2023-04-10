[![codecov](https://codecov.io/gh/kevincobain2000/go-gibberish/branch/master/graph/badge.svg)](https://codecov.io/gh/kevincobain2000/go-gibberish)


<p align="center">
  <a href="https://github.com/kevincobain2000/go-gibberish">
    <img alt="go-gibberish" src="logo.png" width="360">
  </a>
</p>

<h3 align="center">Gibberish Text Detector in Golang.</h3>

<p align="center">
  Detect meaningless sentences in English language.
  <br>
  How to detect if a text has no meaning?
  <br>
  This package uses a dictionary of words to detect if a text is gibberish or not.
</p>

**Languages:** English

**Robust** 1:1 dictionary lookup with heuristics.

**Handles** Punctuation, numbers, and etc.

**Performance:** Blazing fast.



## DEMO


https://go-gibberish.vercel.app/?q=Lorem%20ipsum%20dolor%20sit%20amet,%20consectetur%20adipiscing%20elit.%20Suspendisse%20urna%20turpis,%20porta%20eu%20consectetur%20viverra,%20mattis%20in%20dui.%20Curabitur%20feugiat,%20odio%20in%20viverra%20pulvinar,%20lacus%20dui%20vestibulum%20eros,%20ac%20congue%20elit%20metus%20scelerisque%20tellus.%20Nulla%20faucibus%20eros%20sed%20hendrerit%20faucibus.%20Fusce%20vehicula%20sapien%20lacus,%20at%20venenatis%20metus%20euismod%20in.%20Pellentesque%20urna%20ligula,%20pharetra%20vitae%20tincidunt%20non,%20placerat%20eu%20tellus.%20Vivamus%20tincidunt%20dolor%20ac%20turpis%20tincidunt,%20et%20aliquam%20enim%20sollicitudin.%20Fusce%20imperdiet%20neque%20in%20euismod%20sollicitudin.%20Interdum%20et%20malesuada%20fames%20ac%20ante%20ipsum%20primis%20in%20faucibus.%20Suspendisse%20potenti.%20Donec%20vulputate%20vitae%20mauris%20vel%20placerat.%20Mauris%20mauris%20enim,%20finibus%20quis%20magna%20sed,%20pretium%20gravida%20nisl.



## Installation

```sh
go install github.com/kevincobain2000/go-gibberish
```

## USAGE

```go

package main

import (
	"github.com/kevincobain2000/go-gibberish/gibberish"
)

func main() {

	raw := `Yqxdl vyq wklv qrwe vkhu hqfubswlrq lv qrw yhuvlrq vwdwhphqw iru frpplwphqw lq wkh lqwhuqdwlrq dqg lqwhuqdwlrq ri wkh frqwdl`
	j := gibberish.NewGibberish()
	r := j.Detect(raw)

	fmt.Println(r)
        // outputs
        // &gibberish.Gibberish{
        //   ConfidenceThreshhold: 0.750000,
        //     IsGibberish:          false,
        //     Confidence:           0.992366,
        // }    
}

```


### CHANGE LOG

- v1.0 - Initial release includes detection of gibberish text in English language


### ROADMAP

- [ ] Add support for other languages
- [ ] Add support for bigrams for better detection
- [ ] Add support for repeating words
