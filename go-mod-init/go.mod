module day1

go 1.17

require (
	router v0.0.0
	runRouter v0.0.0
	test v0.0.0
	getRouter v0.0.0
)
replace (
	router => ./router
	runRouter => ./runRouter
	test => ./test
	getRouter => ./router/getRouter

)