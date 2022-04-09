package database

import "L0/models"

func Insert(order models.Order){
	query_order := `INSERT INTO orders(uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	query_delivery := `INSERT INTO delivery(name, phone, zip, city, address, region, email, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	query_payment := `INSERT INTO payment(transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	query_items := `INSERT INTO item(chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`


	db.QueryRow(query_order, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	db.QueryRow(query_delivery, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email, order.OrderUID)
	db.QueryRow(query_payment, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.OrderUID)
	for _, item := range order.Items {
		db.QueryRow(query_items, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status, order.OrderUID)
	}

}

