package urls

type (
	urls struct {
		UrlApi_PATH     string
		PUBLIC_PATH     string
		TEMPLATE_PATH   string
		CSS_PATH        string
		JS_PATH         string
		CONTROllER_PATH string
		MODELS_PATH     string
		MIDDLEWARE_PATH string
		RESPONSIVE_PATH string
		URLS_PATH       string
		UNTILS_PATH     string
	}
)

//pathUrl export urls
func PathUrl() urls {

	var UrlPatterns urls
	UrlPatterns.UrlApi_PATH = "./api/"
	//path public
	UrlPatterns.PUBLIC_PATH = UrlPatterns.UrlApi_PATH + "public/"
	UrlPatterns.TEMPLATE_PATH = UrlPatterns.PUBLIC_PATH + "views/"
	UrlPatterns.CSS_PATH = UrlPatterns.PUBLIC_PATH + "css/"
	UrlPatterns.JS_PATH = UrlPatterns.PUBLIC_PATH + "js/"

	//path controller
	UrlPatterns.CONTROllER_PATH = "/controller/"
	//path model
	UrlPatterns.MODELS_PATH = "/models/"
	//path middlewares
	UrlPatterns.MIDDLEWARE_PATH = "/middlewares/"

	//path responsive
	UrlPatterns.RESPONSIVE_PATH = "/responses/"

	//path config
	UrlPatterns.URLS_PATH = "/urls/"

	//path utils
	UrlPatterns.UNTILS_PATH = "/utils/"

	return UrlPatterns
}
