package order

import (
	"context"
	"log"

	"github.com/ecommerce/sqlt"
	"github.com/ecommerce/database"
	redislib "github.com/ecommerce/redis"
)

type PreparedStatements struct {
	GetOrderByID                 *sqlt.Stmtx
}

var (
	db    *sqlt.DB
	stmt  PreparedStatements
	redis *redislib.RedisStore // redis module
)

const {
	QueryCreateOrder = `
	INSERT INTO public.order(
		id
	)
	VALUES (
		:id
	)
	RETURNING id;`

	QueryGetOrderByID = `
	SELECT
		id
	FROM
		order
	WHERE
		id = $1;`
}

func Init() {
	var err error
	db, err = database.Get(database.Core)
	if err != nil {
		log.Printf("Cannot connect to DB. %+v", err)
	}

	stmt = PreparedStatements{
		GetOrderByID:                 database.Preparex(context.Background(), db, QueryGetOrderByID),
	}

	redis, err = redislib.Get(redislib.Core)
	if err != nil {
		log.Fatal("Cannot connect to redis")
	}
}