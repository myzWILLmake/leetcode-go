package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=535 lang=golang
 *
 * [535] Encode and Decode TinyURL
 */

// @lc code=start
type Codec struct {
	idx     int
	baseUrl string
	enData  map[string]int
	deData  []string
}

func Constructor() Codec {
	c := Codec{}
	c.baseUrl = "http://tinyurl.com/"
	c.enData = make(map[string]int)
	c.deData = make([]string, 0)
	return c
}

func (this *Codec) getShortUrl(n int) string {
	s := strconv.Itoa(n)
	return this.baseUrl + s
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if x, ok := this.enData[longUrl]; ok {
		return this.getShortUrl(x)
	} else {
		idx := len(this.deData)
		this.enData[longUrl] = idx
		this.deData = append(this.deData, longUrl)
		return this.getShortUrl(idx)
	}
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	if strings.Index(shortUrl, this.baseUrl) != 0 {
		return ""
	}

	code, ok := strconv.Atoi(shortUrl[19:])
	if ok != nil {
		return ""
	}

	return this.deData[code]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

// @lc code=end

func main() {
	obj := Constructor()
	url := obj.encode("https://leetcode.com/problems/design-tinyurl")
	ans := obj.decode(url)
	fmt.Println(url)
	fmt.Println(ans)
}
