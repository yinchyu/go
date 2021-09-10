module runRouter

go 1.17



require (
	router v0.0.0
	//getRouter v0.0.0
)
replace (
	getRouter => ../router/getRouter
	router => ../router
)