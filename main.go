package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sql-update-kube/config"
	"sql-update-kube/database/models"

	database "github.com/latonaio/golang-mysql-network-connector"

	_ "github.com/go-sql-driver/mysql"
	logger "github.com/latonaio/golang-logging-library-for-data-platform"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
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
	defer db.Close()

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
		l := logger.NewLogger()
		l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
		data := msg.Data()
		str, _ := json.Marshal(msg.Data()["message"])
		msg.Success()
		f := data["function"].(string)

		switch f {
		case "OrdersHeader":
			datum := models.DataPlatformOrdersHeaderDatum{}
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
			err = json.Unmarshal(str, &datum)
			if err != nil {
				l.Error(err)
				continue
			}

			err = datum.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = datum.Update(ctx, db, boil.Blacklist())
				if err != nil {
					l.Error(err)
					msg.Respond(map[string]interface{}{"result": "failed"})
					continue
				}
			}

		case "OrdersHeaderPartner":
			datum := models.DataPlatformOrdersHeaderPartnerDatum{}
			jmap := map[string]interface{}{}
			json.Unmarshal(str, &jmap)
			for k, v := range jmap {
				if v == "" {
					delete(jmap, k)
				}
			}
			str, err = json.Marshal(jmap)
			if err != nil {
				l.Error(err.Error())
				continue
			}
			err = json.Unmarshal(str, &datum)
			if err != nil {
				l.Error(err)
				continue
			}

			err = datum.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = datum.Update(ctx, db, boil.Blacklist())
				if err != nil {
					l.Error(err)
					msg.Respond(map[string]interface{}{"result": "failed"})
					continue
				}
			}

		case "OrdersHeaderPartnerPlant":
			datum := models.DataPlatformOrdersHeaderPartnerPlantDatum{}
			jmap := map[string]interface{}{}
			json.Unmarshal(str, &jmap)
			for k, v := range jmap {
				if v == "" {
					delete(jmap, k)
				}
			}
			str, err = json.Marshal(jmap)
			if err != nil {
				l.Error(err.Error())
				continue
			}
			err = json.Unmarshal(str, &datum)
			if err != nil {
				l.Error(err)
				continue
			}

			err = datum.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = datum.Update(ctx, db, boil.Greylist())
				if err != nil {
					l.Error(err)
					msg.Respond(map[string]interface{}{"result": "failed"})
					continue
				}
			}

			// case "OrdersHeaderPartnerContact":
			// 	dpohp := models.DataPlatformOrdersHeaderPartnerContactDatum{}
			// 	jmap := map[string]interface{}{}
			// 	json.Unmarshal(str, &jmap)
			// 	for k, v := range jmap {
			// 		if v == "" {
			// 			delete(jmap, k)
			// 		}
			// 	}
			// 	str, err = json.Marshal(jmap)
			// 	if err != nil {
			// 		l.Error(err.Error())
			// 		continue
			// 	}
			// 	err = json.Unmarshal(str, &dpohp)
			// 	if err != nil {
			// 		l.Error(err)
			// 		continue
			// 	}
			// 	fmt.Printf("%s\n", str)

			// 	err = dpohp.Insert(ctx, db, boil.Infer())
			// 	if err != nil {
			// 		l.Info("insert failed: %+v ; try update", err)
			// 		_, err = dpohp.Update(ctx, db, boil.Greylist())
			// 		if err != nil {
			// 			l.Error(err)
			// 			continue
			// 		}
			// 	}
			// 	if err != nil {
			// 		l.Error(err)
			// 		continue
			// 	}

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

		case "DeliveryDocumentHeader":
			datum := models.DataPlatformDeliveryDocumentHeaderDatum{}
			jmap := map[string]interface{}{}
			json.Unmarshal(str, &jmap)
			for k, v := range jmap {
				if v == nil || v == "" {
					delete(jmap, k)
				}
			}
			str, err = json.Marshal(jmap)
			if err != nil {
				l.Error(err.Error())
				continue
			}
			err = json.Unmarshal(str, &datum)
			if err != nil {
				l.Error(err)
				continue
			}

			err = datum.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = datum.Update(ctx, db, boil.Blacklist())
				if err != nil {
					l.Error(err)
					msg.Respond(map[string]interface{}{"result": "failed"})
					continue
				}
			}

		case "PartnerFunctionPartnerFunction":
			datum := models.DataPlatformPartnerFunctionPartnerFunctionDatum{}
			jmap := map[string]interface{}{}
			json.Unmarshal(str, &jmap)
			for k, v := range jmap {
				if v == nil || v == "" {
					delete(jmap, k)
				}
			}
			str, err = json.Marshal(jmap)
			if err != nil {
				l.Error(err.Error())
				continue
			}
			err = json.Unmarshal(str, &datum)
			if err != nil {
				l.Error(err)
				continue
			}

			err = datum.Insert(ctx, db, boil.Infer())
			if err != nil {
				l.Info("insert failed: %+v ; try update", err)
				_, err = datum.Update(ctx, db, boil.Blacklist())
				if err != nil {
					l.Error(err)
					msg.Respond(map[string]interface{}{"result": "failed"})
					continue
				}
			}

		case "ProductMasterGeneral":
			datum := &[]models.SapProductMasterGeneralDatum{}
			str, err := json.Marshal(data["Orders"])
			if err != nil {
				l.Error(err.Error())
				continue
			}
			json.Unmarshal(str, datum)
			for _, pm := range *datum {
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
		msg.Respond(map[string]interface{}{"result": "success"})
		l.Info(data)
	}
}
func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}
