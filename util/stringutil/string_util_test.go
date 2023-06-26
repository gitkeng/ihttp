package stringutil_test

import (
	"testing"

	"github.com/gitkeng/ihttp/util/stringutil"
)

var dataTest = `
{
   "order": {
       "voucher_platform": 0,
       "voucher": 0,
       "order_number": 418536614269569,
       "voucher_seller": 0,
       "created_at": "2021-07-29 11:02:19 +0700",
       "voucher_code": "",
       "gift_option": false,
       "customer_last_name": "",
       "updated_at": "2021-07-30 10:09:05 +0700",
       "promised_shipping_times": "",
       "price": "159.00",
       "national_registration_number": "",
       "payment_method": "COD",
       "customer_first_name": "ม****************ง",
       "shipping_fee": 30,
       "items_count": 2,
       "delivery_info": "",
       "statuses": [
           "canceled"
       ],
       "address_billing": {
           "country": "Thailand",
           "address3": "น*****************i",
           "address2": "",
           "city": "บางใหญ่/ Bang Yai",
           "address1": "1****************************************************น",
           "phone2": "",
           "last_name": "",
           "phone": "66********69",
           "customer_email": "",
           "post_code": "11140",
           "address5": "1***0",
           "address4": "บ***************i",
           "first_name": "มงคลเดข  อินต๊ะวัง"
       },
       "extra_attributes": "{\"TaxInvoiceRequested\":false}",
       "order_id": 418536614269569,
       "gift_message": "",
       "remarks": "",
       "address_shipping": {
           "country": "Thailand",
           "address3": "ก********************k",
           "address2": "",
           "city": "บางซื่อ/ Bang Sue",
           "address1": "1******************6",
           "phone2": "",
           "last_name": "",
           "phone": "66********53",
           "customer_email": "",
           "post_code": "10800",
           "address5": "1***0",
           "address4": "บ***************e",
           "first_name": "ศรีสุดา  กองกาญจนานันท์"
       }
   },
   "order_item": {
       "data": [
           {
               "paid_price": 129,
               "product_main_image": "https://th-live-02.slatic.net/p/f33c3c68054b19c8121f65b9f706dc68.jpg",
               "tax_amount": 0,
               "voucher_platform": 0,
               "reason": "เปลี่ยนใจ",
               "product_detail_url": "https://www.lazada.co.th/products/i2293323634-s7724256761.html?urlFlag=true\u0026mp=1",
               "promised_shipping_time": "",
               "purchase_order_id": "",
               "voucher_seller": 0,
               "shipping_type": "Dropshipping",
               "created_at": "2021-07-29 11:02:19 +0700",
               "voucher_code": "",
               "package_id": "",
               "variation": "",
               "updated_at": "2021-07-30 10:09:05 +0700",
               "purchase_order_number": "",
               "currency": "THB",
               "shipping_provider_type": "standard",
               "sku": "2293323634-1617677348413-0",
               "invoice_number": "INV21072905",
               "cancel_return_initiator": "buyer-cancel",
               "shop_sku": "2293323634_TH-7724256761",
               "is_digital": 0,
               "item_price": 129,
               "shipping_service_cost": 0,
               "tracking_code_pre": "",
               "tracking_code": "",
               "shipping_amount": 23.57,
               "order_item_id": 418536614369569,
               "reason_detail": "",
               "shop_id": "น.ส. สุพรรษา  กระจายความรู้",
               "return_status": "",
               "name": "ทิชชู่เปียก Lucky6ห่อห่อล่ะ50แผ่น ทิชชู่ ทิชชู่เปียก ผ้าเปียก ทิชชู่เปียก lucky ทิชชู่เปียกลัคกี้ ผ้าเปียก กระดาษทิชชู่",
               "shipment_provider": "",
               "voucher_amount": 0,
               "digital_delivery_info": "mongkoldet19@gmail.com",
               "extra_attributes": "",
               "order_id": 418536614269569,
               "status": "canceled"
           },
           {
               "paid_price": 30,
               "product_main_image": "https://th-live.slatic.net/p/1c18f76c0826976f53c2cec8be7edefb.jpg",
               "tax_amount": 0,
               "voucher_platform": 0,
               "reason": "เปลี่ยนใจ",
               "product_detail_url": "https://www.lazada.co.th/products/i2634985241-s9457728132.html?urlFlag=true\u0026mp=1",
               "promised_shipping_time": "",
               "purchase_order_id": "",
               "voucher_seller": 0,
               "shipping_type": "Dropshipping",
               "created_at": "2021-07-29 11:02:19 +0700",
               "voucher_code": "",
               "package_id": "",
               "variation": "",
               "updated_at": "2021-07-30 10:09:05 +0700",
               "purchase_order_number": "",
               "currency": "THB",
               "shipping_provider_type": "standard",
               "sku": "S001",
               "invoice_number": "INV21072905",
               "cancel_return_initiator": "buyer-cancel",
               "shop_sku": "2634985241_TH-9457728132",
               "is_digital": 0,
               "item_price": 30,
               "shipping_service_cost": 0,
               "tracking_code_pre": "",
               "tracking_code": "",
               "shipping_amount": 6.43,
               "order_item_id": 418536614469569,
               "reason_detail": "",
               "shop_id": "น.ส. สุพรรษา  กระจายความรู้",
               "return_status": "",
               "name": "ทิชชู่เปียก Lucky2ห่อห่อล่ะ20แผ่น ทิชชู่เปียก ทิชชู่ ขายส่ง พกพา ทิชชู่เปียก lucky ทิชชู่เปียกลัคกี้ ทิชชู่ กระดาษทิชชู่",
               "shipment_provider": "",
               "voucher_amount": 0,
               "digital_delivery_info": "mongkoldet19@gmail.com",
               "extra_attributes": "",
               "order_id": 418536614269569,
               "status": "canceled"
           }
       ],
       "code": "0",
       "request_id": "210134d716283843692568123"
   }
}
`

func TestCompressAndDecompress(t *testing.T) {
	compressData, err := stringutil.Compress(dataTest, false)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("compress data:\n%s", compressData)
	t.Logf("before compress size: %d", len(dataTest))
	t.Logf("after compress size: %d", len(compressData))

	decompressData, err := stringutil.DeCompress(compressData, false)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("decompress data:\n%s", decompressData)
	t.Logf("before decompress size: %d", len(compressData))
	t.Logf("after decompress size: %d", len(decompressData))

}

func TestToSnake(t *testing.T) {
	datas := []string{
		"HelloWorld",
		"Hello World",
		"Hello_World",
		"HaDmz 123",
	}
	for _, data := range datas {
		t.Logf("snake case: %s", stringutil.ToSnake(data))
	}
}
