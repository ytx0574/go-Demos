package main

import (
	"C"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"go-Demos/LearnGoProject/demo27_Sirial/model"
)

/*
JSON序列化
1. 序列化时, 优先使用字段的json tag定义, 没有则使用字段本身定义, 反序列化同样如此
2. 反序列化时, 可反序列化为任何数据结构, 只要它里面的字段能对应上. 对应上的直接填充值, 结构中没有的字段, 则填充为默认值
3. tag的标准写法为`key:"value"`, 内部实现是一个字符串, 而字符串的解析是基于本格式来实现的
*/
func main() {

	var model1_ = model.JSONModel1{
		Id : 11,
		Name : "大哥",
		Gender : "男",
		Age :12,
		Phone : "996",
		Email : "gmail.com",
	}

	//转JSON
	model1JSONBytes, err := json.Marshal(model1_)
	if err != nil {
		fmt.Printf("model1_转JSON失败 err:%v\n", err)
	}
	fmt.Printf("model1_转jSON String: %v\n", string(model1JSONBytes))

	//转JSON
	var model4_ model.JSONModel4 = model.JSONModel4{
		Id : 11111,
		Name : "大哥1111",
		Gender : "男111",
		Age :12111,
		Phone : "996111",
		Email : "gmail.com111",
	}
	model1JSONBytes4, err := json.Marshal(model4_)
	if err != nil {
		fmt.Printf("model4_转JSON失败 err:%v\n", err)
	}
	fmt.Printf("mode4_转jSON String: %v\n", string(model1JSONBytes4))


	//JSON 反序列化
	var model1 model.JSONModel1 = model.JSONModel1{}
	err = json.Unmarshal(model1JSONBytes, &model1)
	if err != nil {
		fmt.Printf("JSON转model1失败 err:%v\n", err)
	}
	fmt.Printf("model1 = %v \n", model1)

	var model2 model.JSONModel2 = model.JSONModel2{}
	err = json.Unmarshal(model1JSONBytes, &model2)
	if err != nil {
		fmt.Printf("JSON转model3失0败 err:%v\n", err)
	}
	fmt.Printf("model3 = %v \n", model2)

	var model3 model.JSONModel3 = model.JSONModel3{}
	err = json.Unmarshal(model1JSONBytes, &model3)
	if err != nil {
		fmt.Printf("JSON转model3失0败 err:%v\n", err)
	}
	fmt.Printf("model3 = %v \n", model3)



	var model11 model.JSONModel1 = model.JSONModel1{}
	err = json.Unmarshal(model1JSONBytes4, &model11)
	if err != nil {
		fmt.Printf("JSON转model11失败 err:%v\n", err)
	}
	fmt.Printf("model11 = %v \n", model11)

	var model33 model.JSONModel3 = model.JSONModel3{}
	err = json.Unmarshal(model1JSONBytes, &model33)
	if err != nil {
		fmt.Printf("JSON转model33失0败 err:%v\n", err)
	}
	fmt.Printf("model33 = %v \n", model3)


	//默认的两种base64加密变量
	//base64.RawStdEncoding
	//base64.RawURLEncoding = base64.NewEncoding("")

	var strBytes = []byte("j3u32hfJOJO发无法joj解耦覅佛")
	//自己创建加密因子
	var base64Encoding *base64.Encoding = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	var base64EncodingString = base64Encoding.EncodeToString(strBytes)
	fmt.Printf("base64String: %v\n", base64EncodingString)

	strbase64DencodingBytes, err := base64Encoding.DecodeString(base64EncodingString)
	if err != nil {
		 fmt.Printf("base64 dencoding err:%v\n", err)
	}
	fmt.Printf("dencoding base64String:%v\n", string(strbase64DencodingBytes))


	var base64EncodingBytes = make([]byte, base64Encoding.EncodedLen(len(strBytes)))
	base64Encoding.Encode(base64EncodingBytes, strBytes)
	fmt.Printf("base64String1:%v\n", string(base64EncodingBytes))

	var strDencodingBytes1 = make([]byte, base64Encoding.DecodedLen(len(base64EncodingBytes)))
	//返回解密长度和错误信息. 解密的长度不等于解密切片的长度
	n, err := base64Encoding.Decode(strDencodingBytes1, base64EncodingBytes)
	if err != nil {
		fmt.Printf("base64 dencoding err:%v\n", err)
	}
	fmt.Printf("dencoding base64string1:%v, n:%d, decoding len:%d, 原始串长度:%d\n", string(strDencodingBytes1), n, len(strDencodingBytes1), len(strBytes))


	var hexString = hex.EncodeToString([]byte("解耦323lji我"))
	fmt.Printf("ecoding hex string:%v\n", hexString)

	dencodeHexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Printf("dencoding ")
	}
	fmt.Printf("dencoding hex string:%v\n", string(dencodeHexBytes))

	 structs := structs.New(model4_)
	 fmt.Println(structs.Values(), structs.Names(), structs.Fields())

	 for _, v := range structs.Fields() {
	 	fmt.Println(v.Name(), v.Kind(), v.Value(), v.Tag("other_tag"))
	 }
}
