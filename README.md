jcdecauxssb
===========

Go library for the JcDecaux self-service bikes api

Install
-------

	go get https://github.com/nodvos/go-jcdecaux.ssb/jcdecauxssb/

Usage
-----

Create a new client this way:

	client := JcDecaux.NewClient(APIKEY,nil)

With the client set, you can query the api this way:

	// Get the list of contracts:
	contracts := client.Contracts()
	// Get all the station of a specific contract (in this case Paris)
	parisStations := client.Stations("Paris")
	// Retrieve my specific station
	myStation := parisStations[42]
	// Update it
	myStation = client.Station("Paris",myStation.Number)

The full JcDecaux API is documented at https://developer.jcdecaux.com/#/opendata/vls?page=dynamic.
