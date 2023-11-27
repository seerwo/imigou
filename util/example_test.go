package util

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_name(t *testing.T) {
	text  := `<request>
	<kjCustomerOutstore>
		<tradeMode>业务模式（9610=跨境直邮，1210=跨境保税，CC=快件直邮）</tradeMode>
		<saleOrderCode>订单号（必填）</saleOrderCode>
		<flowno>支付订单号</flowno>
		<orderTime>下单时间（必填：2021-01-18 14:21:21）</orderTime>
		<payTime>支付时间（必填：2021-01-18 14:21:21）</payTime>
		<dealerShopName>店铺名称（必填）</dealerShopName>
		<postFee>运费</postFee>
		<insuranceFee>保价费</insuranceFee>		
		<tariffAmount>关税（必填，taxType=N时填0）</tariffAmount>
		<addedValueTaxAmount>增值税（必填，taxType=N时填0）</addedValueTaxAmount>
		<consumptionDutyAmount>消费税（必填，taxType=N时填0）</consumptionDutyAmount>
		<taxAmount>合计税费（必填，taxType=N时填0）</taxAmount>
		<taxType>是否价税分离（Y=单价为税前价，N=单价为含税价）</taxType>
		<disAmount>优惠金额</disAmount>
		<buyerPayment>实付金额（必填）</buyerPayment>
		<grossWeight>商品毛重（必填）</grossWeight>
		<buyerAccount>买家账号（必填）</buyerAccount>
		<buyerIdNumber>订购人身份证号（必填）</buyerIdNumber>
		<buyerName>订购人姓名（必填）</buyerName>
		<paymentMethodName>支付方式（必填，CC可不填））</paymentMethodName>
		<paymentNo>支付单号（必填，与支付机构交互的流水号，CC可不填）</paymentNo>
		<receiverName>收件人（必填）</receiverName>
		<receiverPhone>联系电话（必填）</receiverPhone>
		<province>省（必填）</province>
		<city>市（必填）</city>
		<district>区（必填）</district>
		<receiverAddress>收件人地址（必填，不含省、市、区）</receiverAddress>
		<remark>备注信息</remark>
		<promotionList>
			<promotion>
				<promotionAmount>优惠金额</promotionAmount>
				<promotionRemark>优惠信息说明</promotionRemark>
			</promotion>
		</promotionList>
		<detailList>
			<detail>
				<productNumberCode>商品货号（必填）</productNumberCode>
				<productNumberName>商品名称（必填）</productNumberName>
				<outstoreNumber>数量（必填）</outstoreNumber>
				<price>单价（必填，人民币，taxType=N时填税后价格，taxType=Y时填税前价格）</price>
			</detail>
		</detailList>
	</kjCustomerOutstore>
</request>`


	encodeText := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Printf("Encoded text: %s\n", encodeText)

	resource := encodeText + "T2020" + "123456"
	md5Text,_ := CalculateSign(resource, "", "")
	fmt.Printf("Encoded text: %s\n", md5Text)
	return
}

//func ExampleTest(t *testing.T){
//
//	assert.Equal(t, Day0204(), 0)
//	return
//}