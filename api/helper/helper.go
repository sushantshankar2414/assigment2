package helpers

import (
	"os"
	"strings"
)
func EnforceHttp(url string)string{
	if url[:4] != "http"{
		return "http/" + url
	}
	return url
}

func DomainError(url string) bool{
	if url == os.Getenv("DOMAIN"){
		 return false
	}
	generatedUrl := strings.Replace(url,"http://","",1)
	generatedUrl = strings.Replace(generatedUrl,"https://","",1)
	generatedUrl = strings.Replace(generatedUrl,"www.","",1)
	generatedUrl = strings.Split(generatedUrl,"/")[0]
	if generatedUrl == os.Getenv("DOMAIN"){
		return false
	}
	return true
}