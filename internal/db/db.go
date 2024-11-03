// package db

// import (
// 	"context"
// 	"database/sql"
// 	"time"
// )

// // create + configure new connection to postgreSQL
// func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {

// 	// create a new db conncetion pool, it technically prepares the database object for later use
// 	db, err := sql.Open("postgres", addr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.SetMaxOpenConns(maxOpenConns)
// 	db.SetMaxIdleConns(maxIdleConns)

// 	// converting the string into a time
// 	duration, err := time.ParseDuration(maxIdleTime)
// 	if err != nil {
// 		return nil, err
// 	}

// 	db.SetConnMaxIdleTime(duration)

// 	// if it takes more than 5 seconds to connect we will have a time out
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// by defering we ensure the resources are released once the timeout expries / there is a new connection
// 	defer cancel()

// 	// testing the connection by pinging within the 5 second timeout we set up above
// 	if err = db.PingContext(ctx); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
