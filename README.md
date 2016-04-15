KMP
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/trigram/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/kmp?status.svg)](https://godoc.org/github.com/kkdai/kmp)  [![Build Status](https://travis-ci.org/kkdai/kmp.svg?branch=master)](https://travis-ci.org/kkdai/kmp)


This is a KMP(Knuth–Morris–Pratt algorithm) implement and related string function `Strstr` and `Strchr` which using `KMP` feature inside.

In computer science, the Knuth–Morris–Pratt string searching algorithm (or KMP algorithm) searches for occurrences of a "word" W within a main "text string" S by employing the observation that when a mismatch occurs, the word itself embodies sufficient information to determine where the next match could begin, thus bypassing re-examination of previously matched characters.
 

A chinese version How KMP work note is [here](http://www.evanlin.com/about-kmp/) 
 
Install
---------------
`go get github.com/kkdai/kmp`


Usage
---------------

```go
package main

import (
	"fmt"
	
	"github.com/kkdai/kmp"
)

func main() {
	
	//Using KMP
	fmt.Println(KMP("co", "cocacola"))
	//0, 4
	
	//Using Strstr
	fmt.Println(Strstr("cocacola", "co"))
	//0
	
	//Using Strchr
	fmt.Println(Strstr("cocacola", "co"))
	//4
}	
```


Inspired
---------------

- [Knuth-Morris-Pratt algorithm](http://www-igm.univ-mlv.fr/~lecroq/string/node8.html)
- [理解KMP算法](http://zhangbuhuai.com/2015/07/06/KMP/)
- [String Matching: 各種演算法介紹](http://www.csie.ntnu.edu.tw/~u91029/StringMatching.html)
- [youtube: Tutorial: The Knuth-Morris-Pratt (KMP) String Matching Algorithm](https://www.youtube.com/watch?v=2ogqPWJSftE)
- [youtube: Knuth–Morris–Pratt(KMP) Pattern Matching(Substring search)](https://www.youtube.com/watch?v=GTJr8OvyEVQ)
- [Regular Expression Matching with a Trigram Index](https://swtch.com/~rsc/regexp/regexp4.html)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

