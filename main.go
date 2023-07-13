package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sql-update-kube/config"
	"sql-update-kube/database/models"

	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"

	_ "github.com/go-sql-driver/mysql"
	logger "github.com/latonaio/golang-logging-library-for-data-platform"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	l := logger.NewLogger()

	l.Info("start sql-update-kube 1")

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
		err := callProcess(ctx, db, msg)
		if err != nil {
			msg.Respond(map[string]interface{}{"result": "failed"})
			continue
		}
		msg.Respond(map[string]interface{}{"result": "success"})
	}
}

func callProcess(ctx context.Context, db *database.Mysql, msg rabbitmq.RabbitmqMessage) (err error) {
	l := logger.NewLogger()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
	data := msg.Data()
	str, _ := json.Marshal(msg.Data()["message"])
	msg.Success()
	f := data["function"].(string)

	switch f {
	//case "OrdersHeader":
	//	datum := models.DataPlatformOrdersHeaderDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "OrdersItem":
	//	datum := models.DataPlatformOrdersItemDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Whitelist(keys...))
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "OrdersPartner":
	//	datum := models.DataPlatformOrdersPartnerDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "OrdersAddress":
	//	datum := models.DataPlatformOrdersAddressDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "OrdersItemPricingElement":
	//	datum := models.DataPlatformOrdersItemPricingElementDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "OrdersItemScheduleLine":
	//	datum := models.DataPlatformOrdersItemScheduleLineDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "DeliveryDocumentHeader":
	//	datum := models.DataPlatformDeliveryDocumentHeaderDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "DeliveryDocumentItem":
	//	datum := models.DataPlatformDeliveryDocumentItemDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "DeliveryDocumentPartner":
	//	datum := models.DataPlatformDeliveryDocumentPartnerDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "DeliveryDocumentAddress":
	//	datum := models.DataPlatformDeliveryDocumentAddressDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "InvoiceDocumentHeader":
	//	datum := models.DataPlatformInvoiceDocumentHeaderDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "InvoiceDocumentItem":
	//	datum := models.DataPlatformInvoiceDocumentItemDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "InvoiceDocumentItemPricingElement":
	//	datum := models.DataPlatformInvoiceDocumentItemPricingElementDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "InvoiceDocumentPartner":
	//	datum := models.DataPlatformInvoiceDocumentPartnerDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "InvoiceDocumentAddress":
	//	datum := models.DataPlatformInvoiceDocumentAddressDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	case "PartnerFunctionPartnerFunction":
		datum := models.DataPlatformPartnerFunctionPartnerFunctionDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMaterDoc":
		datum := models.DataPlatformProductMasterGeneralDocDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMaster":
		datum := models.DataPlatformProductMasterGeneralDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}
		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "BarcodeGenerate":
		datum := models.DataPlatformProductMasterGeneralDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}
		err = datum.Update(ctx, db, boil.Whitelist(keys...))
		if err != nil {
			l.Error(err)
			return err
		}

	case "PaymentRequisitionHeader":
		datum := models.DataPlatformPaymentRequisitionHeaderDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "Address":
		datum := models.DataPlatformAddressAddressDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	//case "ProductionOrderHeader":
	//	datum := models.DataPlatformProductionOrderHeaderDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "ProductionOrderItem":
	//	datum := models.DataPlatformProductionOrderItemDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "ProductionOrderItemComponent":
	//	datum := models.DataPlatformProductionOrderItemComponentDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}
	//
	//case "ProductionOrderItemComponentStockConfirmation":
	//	datum := models.DataPlatformProductionOrderItemComponentStockConfDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "ProductionOrderItemComponentCosting":
	//	datum := models.DataPlatformProductionOrderItemComponentCostingDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	//case "ProductionOrderItemOperations":
	//	datum := models.DataPlatformProductionOrderItemOperationsDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	case "BillOfMaterialHeader":
		datum := models.DataPlatformBillOfMaterialHeaderDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "OperationsHeader":
		datum := models.DataPlatformOperationsHeaderDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "PriceMasterPriceMaster":
		datum := models.DataPlatformPriceMasterPriceMasterDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "BillOfMaterialItem":
		datum := models.DataPlatformBillOfMaterialItemDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	//case "OperationsItem":
	//	datum := models.DataPlatformOperationsItemDatum{}
	//	jmap := map[string]interface{}{}
	//	keys := make([]string, 0, len(jmap))
	//	json.Unmarshal(str, &jmap)
	//	for k, v := range jmap {
	//		if v == nil || v == "" {
	//			delete(jmap, k)
	//		} else {
	//			keys = append(keys, k)
	//		}
	//	}
	//	str, err = json.Marshal(jmap)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//	err = json.Unmarshal(str, &datum)
	//	if err != nil {
	//		l.Error(err)
	//		return err
	//	}
	//
	//	err = datum.Insert(ctx, db, boil.Infer())
	//	if err != nil {
	//		l.Info("insert failed: %+v ; try update", err)
	//		err = datum.Update(ctx, db, boil.Whitelist(keys...))
	//		if err != nil {
	//			l.Error(err)
	//			return err
	//		}
	//	}

	case "ProductMasterGeneral":
		datum := models.DataPlatformProductMasterGeneralDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterBusinessPartner":
		datum := models.DataPlatformProductMasterBusinessPartnerDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterAllergen":
		datum := models.DataPlatformProductMasterAllergenDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterCalories":
		datum := models.DataPlatformProductMasterCaloriesDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterNutritionalInfo":
		datum := models.DataPlatformProductMasterNutritionalInfoDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterProductDescription":
		datum := models.DataPlatformProductMasterProductDescriptionDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterProductDescByBP":
		datum := models.DataPlatformProductMasterProductDescByBPDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterBPPlant":
		datum := models.DataPlatformProductMasterBPPlantDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterAccounting":
		datum := models.DataPlatformProductMasterAccountingDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterMRPArea":
		datum := models.DataPlatformProductMasterMRPAreaDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterQuality":
		datum := models.DataPlatformProductMasterQualityDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterStorageLocation":
		datum := models.DataPlatformProductMasterStorageLocationDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterStorageBin":
		datum := models.DataPlatformProductMasterStorageBinDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterWorkScheduling":
		datum := models.DataPlatformProductMasterWorkSchedulingDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductMasterTax":
		datum := models.DataPlatformProductMasterTaxDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "BusinessPartnerGeneral":
		datum := models.DataPlatformBusinessPartnerGeneralDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "BusinessPartnerAccounting":
		datum := models.DataPlatformBusinessPartnerAccountingDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "BusinessPartnerGeneralFinInst":
		datum := models.DataPlatformBusinessPartnerGeneralFinInstDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "BusinessPartnerRole":
		datum := models.DataPlatformBusinessPartnerRoleDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStock":
		datum := models.DataPlatformProductStockProductStockDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStockAvailability":
		datum := models.DataPlatformProductStockProductStockAvailabilityDatum{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStockByBatch":
		datum := models.DataPlatformProductStockProductStockByBTCH{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStockAvailabilityByBatch":
		datum := models.DataPlatformProductStockProductStockAvailByBTCH{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStockByStorageBin":
		datum := models.DataPlatformProductStockProductStockBySTRGBin{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStockAvailabilityByStorageBin":
		datum := models.DataPlatformProductStockProductStockAvailBySTRGBin{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStockByStorageBinByBatch":
		datum := models.DataPlatformProductStockProductStockBySTRGBinBTCH{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	case "ProductStockAvailabilityByStorageBinByBatch":
		datum := models.DataPlatformProductStockProductStockAvailBySTRGBinBTCH{}
		jmap := map[string]interface{}{}
		keys := make([]string, 0, len(jmap))
		json.Unmarshal(str, &jmap)
		for k, v := range jmap {
			if v == nil || v == "" {
				delete(jmap, k)
			} else {
				keys = append(keys, k)
			}
		}
		str, err = json.Marshal(jmap)
		if err != nil {
			l.Error(err)
			return err
		}
		err = json.Unmarshal(str, &datum)
		if err != nil {
			l.Error(err)
			return err
		}

		err = datum.Insert(ctx, db, boil.Infer())
		if err != nil {
			l.Info("insert failed: %+v ; try update", err)
			err = datum.Update(ctx, db, boil.Whitelist(keys...))
			if err != nil {
				l.Error(err)
				return err
			}
		}

	// case "PaymentRequisitionItem":
	// 	datum := models.DataPlatformPaymentRequisitionItemDatum{}
	// 	jmap := map[string]interface{}{}
	// 	json.Unmarshal(str, &jmap)
	// 	for k, v := range jmap {
	// 		if v == nil || v == "" {
	// 			delete(jmap, k)
	// 		}
	// 	}
	// 	str, err = json.Marshal(jmap)
	// 	if err != nil {
	// 		l.Error(err)
	// 		return err
	// 	}
	// 	err = json.Unmarshal(str, &datum)
	// 	if err != nil {
	// 		l.Error(err)
	// 		return err
	// 	}

	// 	err = datum.Insert(ctx, db, boil.Infer())
	// 	if err != nil {
	// 		l.Info("insert failed: %+v ; try update", err)
	// 		err = datum.Update(ctx, db, boil.Blacklist())
	// 		if err != nil {
	// 			l.Error(err)
	// 			return err
	// 		}
	// 	}

	default:
		err = xerrors.Errorf("unknown function name %s", f)
		l.Error(err)
		return err
	}
	l.Info(data)

	return nil
}

func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}
