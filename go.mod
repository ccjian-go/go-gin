module go-gin

go 1.14

require (
	github.com/gin-gonic/gin v1.5.0
	my/apis v0.0.0 // indirect
	my/database v0.0.0
	my/logger v0.0.0 // indirect
	my/models v0.0.0 // indirect
	my/route v0.0.0
)

replace (
	my/apis => ./apis
	my/database => ./database
	my/logger => ./logger
	my/models => ./models
	my/route => ./route
)
