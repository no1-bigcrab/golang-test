package config

type (
	configs struct {
		PUBLIC_PATH     string
		TEMPLATE_PATH   string
		CSS_PATH        string
		JS_PATH         string
		CONTROllER_PATH string
		MODELS_PATH     string
		MIDDLEWARE_PATH string
		RESPONSIVE_PATH string
		CONFIG_PATH     string
		UNTILS_PATH     string
	}
)

//ReturnURLS export urls
func returnURLS() configs {

	var urlPatterns configs

	//path public
	urlPatterns.PUBLIC_PATH = "/public/"
	urlPatterns.TEMPLATE_PATH = urlPatterns.PUBLIC_PATH + "/views/"
	urlPatterns.CSS_PATH = urlPatterns.PUBLIC_PATH + "/css/"
	urlPatterns.JS_PATH = urlPatterns.PUBLIC_PATH + "/js/"

	//path controller
	urlPatterns.CONTROllER_PATH = "/controller/"
	//path model
	urlPatterns.MODELS_PATH = "/models/"
	//path middlewares
	urlPatterns.MIDDLEWARE_PATH = "/middlewares/"

	//path responsive
	urlPatterns.RESPONSIVE_PATH = "/responses/"

	//path config
	urlPatterns.CONFIG_PATH = "/config/"

	//path utils
	urlPatterns.UNTILS_PATH = "/utils/"

	return urlPatterns
}
