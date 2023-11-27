package order

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/seerwo/imigou/util"
	"github.com/seerwo/imigou/wms/context"
)

const(
	ADD_KJ_CUSTOMER_OUTSTORE = "AddKjCustomerOutstore"
	CANCEL_KJ_CUSTOMER_OUTSTORE = "CancelKjCustomerOutstore"
	KJ_CUSTOMER_OUTSTORE_RETURN = "KjCustomerOutstoreReturn"
	KJ_CUSTOMER_OUTSTORE_STATUSRESPONSE = "KjCustomerOutstoreStatusResponse"
	SEARCH_KJ_CUSTOMER_OUTSTORE = "SearchKjCustomerOutstore"
)

type AddKjCustomerOutstore struct {
	Request struct {
		TradeMode string `json:"tradeMode"`                         //业务模式（9610=跨境直邮，1210=跨境保税，CC=快件直邮）
		SaleOrderCode string `json:"saleOrderCode"`                 //订单号（必填）
		Flowno string `json:"flowno"`                               //支付订单号
		OrderTime string `json:"orderTime"`                         //下单时间（必填：2021-01-18 14:21:21）
		PayTime string `json:"payTime"`                             //支付时间（必填：2021-01-18 14:21:21）
		DealerShopName string `json:"dealerShopName"`               //店铺名称（必填）
		PostFee string `json:"postFee"`                             //运费
		InsuranceFee string `json:"insuranceFee"`                   //保价费
		TariffAmount string `json:"tariffAmount"`                   //关税（必填，taxType=N时填0）
		AddedValueTaxAmount string `json:"addedValueTaxAmount"`     //增值税（必填，taxType=N时填0）
		ConsumptionDutyAmount string `json:"consumptionDutyAmount"` //消费税（必填，taxType=N时填0）
		TaxAmount string `json:"taxAmount"`                         //合计税费（必填，taxType=N时填0）
		TaxType string `json:"taxType"`                             //是否价税分离（Y=单价为税前价，N=单价为含税价）
		DisAmount string `json:"disAmount"`                         //优惠金额
		buyerPayment string `json:"buyerPayment"`                   //实付金额（必填）
		GrossWeight string `json:"grossWeight"`                     //商品毛重（必填）
		BuyerAccount string `json:"buyerAccount"`                   //买家账号（必填）
		BuyerIdNumber string `json:"buyerIdNumber"`                 //订购人身份证号（必填）
		BuyerName string `json:"buyerName"`                         //订购人姓名（必填）
		PaymentMethodName string `json:"paymentMethodName"`         //支付方式（必填，CC可不填））
		PaymentNo string `json:"paymentNo"`                         //支付单号（必填，与支付机构交互的流水号，CC可不填）
		ReceiverName string `json:"receiverName"`                   //收件人（必填）
		ReceiverPhone string `json:"receiverPhone"`                 //联系电话（必填）
		Province string `json:"province"`                           //省（必填）
		City string `json:"city"`                                   //市（必填）
		District string `json:"district"`                           //区（必填）
		ReceiverAddress string `json:"receiverAddress"`             //收件人地址（必填，不含省、市、区）
		Remark string `json:"remark"`                               //备注信息
		PromotionList []Promotion `json:"promotionList"`
		DetailList []Detail `json:"detailList"`
	} `json:"request"`
}

type Response struct {
	util.CommonError
}

type Promotion struct {
	PromotionAmount string `json:"promotionAmount"` //优惠金额
	PromotionRemark string `json:"promotionRemark"` //优惠信息说明
}

type Detail struct {
	ProductNumberCode string `json:"productNumberCode"` //商品货号（必填）
	ProductNumberName string `json:"productNumberName"` //商品名称（必填）
	OutstoreNumber string `json:"outstoreNumber"` //数量（必填）
	Price string `json:"price"` //单价（必填，人民币，taxType=N时填税后价格，taxType=Y时填税前价格）
}

type CancelKjCustomerOutstore struct {
	SaleOrderCode string `json:"saleOrderCode"` //订单号
}

type KjCustomerOutstoreReturn struct {
	SaleOrderCode string `json:"saleOrderCode"` //订单号
	LogisticsCode string `json:"logisticsCode"` //快递公司
	LogisticsNumber string `json:"logisticsNumber"` //运单号
	MarkPlace1 string `json:"markPlace1"` //大头笔或集包地
	MarkPlace2 string `json:"markPlace2"` //大头笔或三段码
	OutstoreStatus string `json:"outstoreStatus"` //发货单状态: 10-预校验不通过，21-已报海关，22-海关单证放行，23-海关单证审核不过，30-已发仓库，31-库存不足，32-仓库已打包，90-包裹出库
	OutstoreStatusMsg string `json:"outstoreStatusMsg"` //发货单状态信息（10,23,31时必填）
	OccurTime string `json:"occurTime"` //状态更新时间(2021-05-05 12:00:00)
}

type KjCustomerOutstoreStatusResponse struct {
	Flag string `json:"flag"` //success|failure，报文是否接收成功
	Code string `json:"code"` //0=正常，非0=系统异常，需要确认message信息
	Message string `json:"message"` //异常信息，code=非0时必填
	KjCustomerOutstoreReturn KjCustomerOutstoreReturn
}

type SearchKjCustomerOutstore struct {
	SaleOrderCode string `json:"saleOrderCode"` //订单号
}

//Order struct
type Order struct {
	*context.Context
}

//NewOrder instance
func NewOrder(context *context.Context) *Order {
	order := new(Order)
	order.Context = context
	return order
}

//AddKjCustomerOutstore
func (order *Order) GetAddKjCustomerOutstore(req AddKjCustomerOutstore) (res Response, err error){
	jsonBytes, err := xml.Marshal(req)
	if err != nil {
		return
	}

	accessParam := ""
	if accessParam, err = order.GetAccessParam(string(jsonBytes)); err != nil {
		return
	}
	//uri := fmt.Sprintf("%s%s", util.WMS_WEB_URL + CREATE_DN_URL , accessParam)

	var response []byte
	if response, err = util.NewHTTPPost(util.WMS_WEB_URL, string(accessParam)); err != nil {
		return
	}

	if err = json.Unmarshal(response,&res); err != nil {
		return
	}
	if res.Code != 0 && res.Flag == "false" {
		err = fmt.Errorf("GetAddKjCustomerOutstore Error , errcode=%d , errmsg=%s", res.Code, res.Message)
		return
	}
	return
}
//cancelDN