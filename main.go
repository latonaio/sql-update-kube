package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sql-update-kube/config"
	"sql-update-kube/database"
	"sql-update-kube/database/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/latonaio/golang-logging-library/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func getBodyHeader(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func main() {
	l := logger.NewLogger()

	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), c.RMQ.QueueTo())
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
		sessionId := getBodyHeader(data)
		rmq.AddSendTemp(map[string]interface{}{"runtime_session_id": sessionId})
		// l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": sessionId})
		msg.Success()
		str, err := json.Marshal(data["message"])
		if err != nil {
			l.Error(err.Error())
			rmq.Send(c.RMQ.QueueTo()[0], data)
			continue
		}
		row, ok := data["function"]
		if !ok {
			rmq.Send(c.RMQ.QueueTo()[0], data)
			continue
		}
		f, ok := row.(string)
		if !ok {
			rmq.Send(c.RMQ.QueueTo()[0], data)
			continue
		}
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
			rmq.Send(c.RMQ.QueueTo()[0], data)
			continue
		}
		rmq.Send(c.RMQ.QueueTo()[0], map[string]interface{}{"status": "create new row"})
	}
}
