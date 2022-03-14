package dbutils

const ArrivalTable = `
	CREATE TABLE IF NOT EXISTS arrival (
		id SERIAL PRIMARY KEY,
		stop_code VARCHAR(40) NOT NULL,
		bus_line VARCHAR(40) NOT NULL,
		time TIME NOT NULL
	)
`
