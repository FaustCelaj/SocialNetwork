// package main

// import (
// 	"log"

// 	"github.com/faustcelaj/social_project/internal/db"
// 	"github.com/faustcelaj/social_project/internal/env"
// 	"github.com/faustcelaj/social_project/internal/store"
// )

// func main() {
// 	cfg := config{
// 		addr: env.GetString("ADDR", ":8080"),
// 		db: dbConfig{
// 			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
// 			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
// 			maxIdleConns: env.GetInt("DB_MAX_IDEL_CONNS", 30),
// 			maxIdleTime:  env.GetString("DB_MAX_IDEL_TIME", "15min"),
// 		},
// 	}

// 	db, err := db.New(
// 		cfg.db.addr,
// 		cfg.db.maxOpenConns,
// 		cfg.db.maxIdleConns,
// 		cfg.db.maxIdleTime,
// 	)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	defer db.Close()
// 	log.Println("database connection pool established")

// 	store := store.NewStorage(db)

// 	app := &application{
// 		config: cfg,
// 		store:  store,
// 	}

// 	mux := app.mount()

// 	log.Fatal(app.run(mux))
// }

package main

import (
	"log"

	"github.com/faustcelaj/social_project/internal/env"
	"github.com/faustcelaj/social_project/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
