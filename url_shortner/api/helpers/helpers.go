
func EnforceHttp()string{
	if url[:4] != "http"{
		return "http/" + url
	}
	return url
}

func DomainError(url string) bool{
	if url == os.Getenv("DOMAIN"){
		return false
	}
	generatedUrl := string.Replace(url,"http://","",1)
	generatedUrl := string.Replace(generatedUrl,"https://","",1)
	generatedUrl := string.Replace(generatedUrl,"www.","",1)
	generatedUrl := string.Split(generatedUrl,"/")[0]
	if generatedUrl == os.Getenv("DOMAIN"){
		return false
	}
	return true
}