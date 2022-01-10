package main

import (
	"context"
	"encoding/json"
	"sql-update-kube/config"
	"sql-update-kube/database"
	"sql-update-kube/database/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/latonaio/golang-logging-library/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	l := logger.NewLogger()

	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), nil)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()
	ctx := context.Background()

	for msg := range iter {
		data := msg.Data()
		str, _ := json.Marshal(msg.Data()["message"])
		f := data["function"].(string)
		msg.Success()

		switch f {
		case "XXXXXXXXXXXXXXXXXXX":
			pms := &[]models.XXXXXXXXXXXXXXXXXXXDatum{}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
					}
				}
			}
		case "YYYYYYYYYYYYYYYYYYY":
			sd := &models.YYYYYYYYYYYYYYYYYYYDatum{}
			json.Unmarshal(str, sd)
			err = sd.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = sd.Update(ctx, db, boil.Infer())
				if err != nil {
					l.Error(err)
				}
			}
		default:
			l.Error("catch unknown function %s", f)
		}
	}
}
