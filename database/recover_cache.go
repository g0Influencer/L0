package database

import (
	"L0/cache"
	"L0/models"
	"log"
)

func Recover(){
	Connect()
	cache.Init()

	query_order:= `select * from orders`
	query_delivery:= `select name,phone,zip,city,address,region,email from delivery where order_uid = $1`
	query_payment:= `select transaction,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee from payment where order_uid = $1`
	query_item:= `select chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status from item where order_uid = $1`


	order_rows,err:=db.Query(query_order)
	if err!= nil{
		log.Println(err)
	}

	defer order_rows.Close()

	for order_rows.Next(){
		order:=models.Order{}

		order_rows.Scan(&order.OrderUID,&order.TrackNumber,&order.Entry,&order.Locale,
			&order.InternalSignature,&order.CustomerID,&order.DeliveryService,
			&order.ShardKey,&order.SmID,&order.DateCreated,&order.OofShard)


		delivery:=models.Delivery{}

		delivery_rows,err:= db.Query(query_delivery, order.OrderUID)
		if err!= nil{
			log.Println(err)
		}
		defer delivery_rows.Close()

		for delivery_rows.Next() {
			delivery_rows.Scan(&delivery.Name, &delivery.Phone, &delivery.Zip,
				&delivery.City, &delivery.Address, &delivery.Region,&delivery.Email)
		}
		order.Delivery = delivery

		payment:=models.Payment{}
		payment_rows,err:= db.Query(query_payment, order.OrderUID)
		if err!= nil{
			log.Println(err)
		}
		defer payment_rows.Close()

		for payment_rows.Next() {
			payment_rows.Scan(&payment.Transaction, &payment.RequestID, &payment.Currency,
				&payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank,
				&payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
		}
		order.Payment = payment

		items:=make([]models.Item,0)

		item_rows,err:=db.Query(query_item, order.OrderUID)
		if err!= nil{
			log.Println(err)
		}
		defer item_rows.Close()

		for item_rows.Next() {
			i:=models.Item{}
			item_rows.Scan(&i.ChrtID, &i.TrackNumber, &i.Price, &i.Rid, &i.Name,
				&i.Sale, &i.Size, &i.TotalPrice, &i.NmID, &i.Brand, &i.Status)
			items = append(items,i)
		}
		order.Items = items
		cache.Set(order)

	}


}
