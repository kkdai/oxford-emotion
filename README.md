Project Oxford Emotion API for Golang
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/oxford-emotion/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/oxford-face?status.svg)](https://godoc.org/github.com/kkdai/oxford-emotion)  [![Build Status](https://travis-ci.org/kkdai/oxford-emotion.svg)](https://travis-ci.org/kkdai/oxford-emotion)
 
![](https://www.projectoxford.ai/images/bright/emotion/EmotionAPI.png)
 
This package is [Project Oxford](https://www.projectoxford.ai/) [Face API](https://www.projectoxford.ai/emotion) in Golang

What is Project Oxford Emotion API
---------------

[Project Oxford](https://www.projectoxford.ai/) is a web services from Microsoft. It contains following services. (refer from this [page](https://www.projectoxford.ai/emotion).)

Installation
---------------
```
go get github.com/kkdai/oxford-emotion
```

How to use it
---------------

Sign-up for Microsoft Translator API (see [here](http://blogs.msdn.com/b/translation/p/gettingstarted1.aspx) for more detail) and get your developer credentials. Use the client ID and secret to instantiate a translator as shown below.

```go
package main

import (
	"fmt"
	"os"

	. "github.com/kkdai/oxford-face"
)

func main() {
	key := os.Getenv("MSFT_KEY")
	if key == "" {
		fmt.Println("Please export your key to environment first, `export MSFT_KEY=12234`")
		return
	}
	f := NewFace(key)

	//Detect
	ret, err := f.DetectFile(nil, "./1.jpg")
	fmt.Println("ret:", ret, " err=", err)
}
```

Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`

Inspired
---------------

- [Project Oxford: Emotion API](https://www.projectoxford.ai/emotion)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

