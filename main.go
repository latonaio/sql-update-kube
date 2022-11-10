package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sql-update-kube/config"
	"sql-update-kube/database"
	"sql-update-kube/database/models"

	_ "github.com/go-sql-driver/mysql"
	logger "github.com/latonaio/golang-logging-library-for-data-platform"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	fmt.Printf("THIS VERSION IS 2\n")
	l := logger.NewLogger()

	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), "", c.RMQ.QueueTo(), -1)
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
		// str, err := json.Marshal(data["message"])
		// if err != nil {
		// 	l.Error(err.Error())
		// 	continue
		// }
		f := data["service_label"].(string)
		msg.Success()

		switch f {
		case "ORDERS":
			dpoh := models.DataPlatformOrdersHeaderDatum{}
			str, err := json.Marshal(data["Orders"])
			jmap := map[string]interface{}{}
			json.Unmarshal(str, &jmap)
			for k, v := range jmap {
				if v == nil || v == "" {
					delete(jmap, k)
				}
			}
			jmap["TransactionCurrency"] = "JPY"
			str, err = json.Marshal(jmap)
			if err != nil {
				l.Error(err.Error())
				continue
			}
			err = json.Unmarshal(str, &dpoh)
			if err != nil {
				l.Error(err)
				continue
			}
			fmt.Printf("%s\n", str)

			err = dpoh.Upsert(ctx, db, boil.Infer(), boil.Infer())
			if err != nil {
				l.Error(err)
				continue
			}

		// case "OrdersHeaderPDF":
		// 	pms := &[]models.DataPlatformOrdersHeaderPDFDatum{}
		// 	json.Unmarshal(str, pms)
		// 	for _, pm := range *pms {
		// 		err = pm.Insert(ctx, db, boil.Infer())
		// 		if err != nil {
		// 			l.Info("insert failed: %+v ; try update", err)
		// 			_, err = pm.Update(ctx, db, boil.Infer())
		// 			if err != nil {
		// 				l.Error(err)
		// 				continue
		// 			}
		// 		}
		// 	}
		case "ProductMasterGeneral":
			pms := &[]models.SapProductMasterGeneralDatum{}
			str, err := json.Marshal(data["Orders"])
			if err != nil {
				l.Error(err.Error())
				continue
			}
			json.Unmarshal(str, pms)
			for _, pm := range *pms {
				err = pm.Insert(ctx, db, boil.Infer())
				if err != nil {
					l.Info("insert failed: %+v ; try update", err)
					_, err = pm.Update(ctx, db, boil.Infer())
					if err != nil {
						l.Error(err)
						continue
					}
				}
			}
		default:
			l.Error("catch unknown function %s", f)
		}
		l.Info("insert done: %v", data)
	}
}
