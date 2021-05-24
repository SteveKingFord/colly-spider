package cnki

import (
	"encoding/json"
)

type QueryJson struct {
	Platform  string `json:"Platform"`
	DBCode    string `json:"DBCode"`
	KuaKuCode string `json:"KuaKuCode"`
	QNode     QNode  `json:"QNode"`
}

type QNode struct {
	QGroup []QGroupItem `json:"QGroup"`
}

type QGroupItem struct {
	Key        string      `json:"Key"`
	Title      string      `json:"Title"`
	Logic      int         `json:"Logic"`
	Items      []Item      `json:"Items"`
	ChildItems []ChildItem `json:"ChildItems"`
}

type Item struct {
	Key         string `json:"Key"`
	Title       string `json:"Title"`
	Logic       int    `json:"Logic"`
	Name        string `json:"Name"`
	Value       string `json:"Value"`
	Value2      string `json:"Value2"`
	Operate     string `json:"Operate"`
	BlurType    string `json:"BlurType"`
	ExtendType  int    `json:"ExtendType"`
	ExtendValue string `json:"ExtendValue"`
}

type ChildItem struct {
	Key   string `json:"Key"`
	Title string `json:"Title"`
	Logic int    `json:"Logic"`
	Items []Item `json:"Items"`
	ChildItems []interface{}
}

func NewQueryJson() string{
	q:= &QueryJson{
		Platform:  "",
		DBCode:    "SCDB",
		KuaKuCode: "CJFQ,CDMD,CIPD,CCND,CISD,SNAD,BDZK,CCVD,CJFN,CCJD",
		QNode: QNode{
			QGroup: []QGroupItem{
				{
					Key:   "Subject",
					Title: "",
					Logic: 1,
					Items: []Item{
						{
							Title:    "主题",
							Name:     "SU",
							Value:    "脂肪肝",
							Operate:  "%=",
							BlurType: "",
						},
					},
					ChildItems: []ChildItem{},
				},
				{
					Key:   "SCDBGroup",
					Title: "",
					Logic: 1,
					Items: []Item{},
					ChildItems: []ChildItem{
						{
							Key:   "2",
							Title: "",
							Logic: 1,
							Items: []Item{
								{
									Key:         "脂肪肝",
									Title:       "脂肪肝",
									Logic:       2,
									Name:        "主要主题",
									Operate:     "",
									Value:       "脂肪肝",
									ExtendType:  0,
									ExtendValue: "",
									Value2:      "",
									BlurType:    "",
								},
							},
							ChildItems:[]interface{}{},
						},
						{
							Key:   "3",
							Title: "",
							Logic: 1,
							Items: []Item{
								{
									Key:         "2021",
									Title:       "2021",
									Logic:       2,
									Name:        "年",
									Operate:     "",
									Value:       "2021",
									ExtendType:  0,
									ExtendValue: "",
									Value2:      "",
									BlurType:    "",
								},
								{
									Key:         "2020",
									Title:       "2020",
									Logic:       2,
									Name:        "年",
									Operate:     "",
									Value:       "20210",
									ExtendType:  0,
									ExtendValue: "",
									Value2:      "",
									BlurType:    "",
								},
								{
									Key:         "2019",
									Title:       "2019",
									Logic:       2,
									Name:        "年",
									Operate:     "",
									Value:       "2019",
									ExtendType:  0,
									ExtendValue: "",
									Value2:      "",
									BlurType:    "",
								},
								{
									Key:         "2018",
									Title:       "2018",
									Logic:       2,
									Name:        "年",
									Operate:     "",
									Value:       "2018",
									ExtendType:  0,
									ExtendValue: "",
									Value2:      "",
									BlurType:    "",
								},
								{
									Key:         "2017",
									Title:       "2017",
									Logic:       2,
									Name:        "年",
									Operate:     "",
									Value:       "2017",
									ExtendType:  0,
									ExtendValue: "",
									Value2:      "",
									BlurType:    "",
								},
								{
									Key:         "2016",
									Title:       "2016",
									Logic:       2,
									Name:        "年",
									Operate:     "",
									Value:       "2016",
									ExtendType:  0,
									ExtendValue: "",
									Value2:      "",
									BlurType:    "",
								},
							},
							ChildItems:[]interface{}{},
						},
					},
				},
			},
		},
	}

	jsonData,_ := json.Marshal(&q)

	return string(jsonData)

}


func NewQuery() string {
	q:= &QueryJson{
		Platform:  "",
		DBCode:    "SCDB",
		KuaKuCode: "CJFQ,CDMD,CIPD,CCND,CISD,SNAD,BDZK,CCVD,CJFN,CCJD",
		QNode: QNode{
			QGroup: []QGroupItem{
				{
					Key:   "Subject",
					Title: "",
					Logic: 1,
					Items: []Item{
						{
							Title:    "主题",
							Name:     "SU",
							Value:    "脂肪肝合并慢乙肝",
							Operate:  "%=",
							BlurType: "",
						},
						{
							Title:    "主题",
							Name:     "SU",
							Value:    "脂肪肝合并慢乙肝",
							Operate:  "%=",
							BlurType: "ddx",
						},
					},
					ChildItems: []ChildItem{},
				},
			},
		},
	}

	jsonData,_ := json.Marshal(&q)

	return string(jsonData)
}

var HeadMap = map[string]string{
	"Accept":           "text/html, */*; q=0.01",
	"Accept-Encoding":  "gzip, deflate, br",
	"Accept-Language":  "zh-CN,zh;q=0.9,zh-TW;q=0.8",
	"Connection":       "keep-alive",
	"Content-Length":   "7851",
	"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
	"Cookie":           " Ecp_ClientId=3210521172703839997; ASP.NET_SessionId=1wrlotpv1sehk03tqv1glkqp; SID_kns8=123119; SID_kns_new=kns123109; cnkiUserKey=c7c89b56-5950-1e8d-85ee-b702e7e750fe; CurrSortField=%e5%8f%91%e8%a1%a8%e6%97%b6%e9%97%b4%2f(%e5%8f%91%e8%a1%a8%e6%97%b6%e9%97%b4%2c%27TIME%27); CurrSortFieldType=desc; Ecp_ClientIp=59.57.195.38; SID_kcms=124109; SID_kxreader_new=011124; c_m_LinID=LinID=WEEvREcwSlJHSldTTEYzVnBFbFF4cDdyWllTUXBFbW5Fek5idG1OUGRoVT0=$9A4hF_YAuvQ5obgVAqNKPCYcEjKensW4IQMovwHtwkF4VYPoHbKxJw!!&ot=05/28/2021 17:29:19; LID=WEEvREcwSlJHSldTTEYzVnBFbFF4cDdyWllTUXBFbW5Fek5idG1OUGRoVT0=$9A4hF_YAuvQ5obgVAqNKPCYcEjKensW4IQMovwHtwkF4VYPoHbKxJw!!; c_m_expire=2021-05-28 17:29:19; Ecp_session=1; Ecp_LoginStuts={\"IsAutoLogin\":false,\"UserName\":\"1311369756@qq.com\",\"ShowName\":\"1311369756%40qq.com\",\"UserType\":\"jf\",\"BUserName\":\"\",\"BShowName\":\"\",\"BUserType\":\"\",\"r\":\"Jyzucl\"}; _pk_ref=%5B%22%22%2C%22%22%2C1621828634%2C%22https%3A%2F%2Fwww.google.com%2F%22%5D; _pk_ses=*; Ecp_loginuserjf=1311369756@qq.com; knsLeftGroupSelectItem=null9%3B3%3B4%3B; _pk_id=1ed33360-c5fb-4f01-bfd1-3b7ca1106460.1621589234.2.1621828962.1621828634.",
	"Host":             "kns.cnki.net",
	"Origin":           "https://kns.cnki.net",
	"Referer":          "https://kns.cnki.net/kns8/defaultresult/index",
	"sec-ch-ua":        "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"",
	"sec-ch-ua-mobile": "?0",
	"Sec-Fetch-Dest":   "empty",
	"Sec-Fetch-Mode":   "cors",
	"Sec-Fetch-Site":   "same-origin",
	"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
	"X-Requested-With": "XMLHttpRequest",
}

var QueryData = map[string]string{
	"IsSearch":          "true",
	"queryJson":         NewQuery(),
	"SearchSql":         "0645419CC2F0B23BC604FFC82ADF67C692F6944D0202D2DC6656A6C7AC1EDC809A155BF2DE12E2F98C4BD55B8DC0937B674DE0CD2B86BB3266824FC750BED96C00B6FFF8B53C732897B1E623A882A32F0BF5CF51B6EAB5042BB9034B07B74C7567AD9C4111D92A13DEE9BBC8802C0D282F61171EB1794B796728CBF9FEE654FB6C550FE2154E7B199FF2B0A2F14AFB850735D5697534FBB43620CC6627D49477153E142292C900AECEE9968C0C2D01B114E7D3A5C688C197F2D6353CEEA237DECD555EDDFE4641352CE4DE3CD5DB4ADF5126EB9AAB81EEB6A9C7048B3ABC508E2511FFECD5B809118E6FE24E065BDF56F8D679AD87AEC534278173EA90C7BD2F782DF316E0689F3F9CF7130A3D164C0B9F67E0BCFD40B11A893E059B385C139C93381B0E49D9B341E9F873F36CD7419DB9DCFAEE1BC089BBDD3F69BD336BCE0E0CB84C933EE6C249556E6F2832A7DBC426AC9B3FF7C49275C5C9A40BAEE1916D88FE3B89F46F3855840137991445DDDE36419A95CC34A3A67996D1066FC72587ECC3178EF9FE478D64E053B67FE75AF13E01C467EF6BB65CB74CBC667876B06414BBE808D9BFDE5C7B37665C52B788CE77B25BBF742357D34924C53618AAB0F1F923AD03D5C0AAD60B43CDF34DE9F2642F6432E1AE8588C458E5A4F89EA592507F2D75D75E5CFC0C5EC385E23CDDE479774AB9AA0C848545A87A13C88243D0F316BDD0D2D277F4156124D62AF5D1752EE9DF1AA61E73BD1851CA0CCBB0E643BBC465A5A252D0379F50BBCDD9057F7D9B5963318588BCA97EB74575AB118421EE8725A4E43DD27C107D64B1941DBCCB6E5126BC57C7E14BD5142219D1F0BCC312D76C5E166382E7A14D0D8212827A25ACD9BD083214B358CE6CF5DAD6A25B1B6D6513211C44E00F79B2862723CAC5828518BDDC3D1C4A4450501000B35D9173CD222F55DADD4BC7BB18D9AB47CE6299C5EC3F7B7BD795B65A1DB63746D7856CA667C6D0C28309D4176A54A771C8303BA11FC70D5FE03FE808B35150681CE1449249A6513EB94BE9DBE1002F0C41B2DC471CF1792524EDFEBC6BA3838E316B9FCF39E2F6DC7F08C6AF7C05DE1C44FD30BA2E7423256156E8E3F388E4069C0C2BB459C523134D5AF0CD59A6508A08A3F57F4BD33020E8A966326745D125A0599DE6710B72EC02E3A4C914207F9652A46A952710D24D86D42ED92C2C366727E97A13A4F495FCB3B07537A94EAA1BF1A29EB467F2F53042A8A74A590B5636DEF39E4E8A1E43617743BB24907B87DBD84C61D0A6F2475AEB9E541B0C4B446801E4FD8ACBAED1E7552A7735301749E63599FF0910F191473383338EE3F0A2CDC5CADB7DCC7A46065DC7C652185ECD579930A9F6E5B6F63ABA7283868365C8DC3E75D73A1B22FF62D7CF02FBB0B889D3F1B608E16B29829170DE0CE3B755F1E1DAABE2117B7D52BA44F2D615A436C2CCAE062EB78384486DBD8A3D510021BFC733621A0E51C8D44C892DDE003CD64B91ADCF8DEC5083E07CE303EB2C227636F3B0E69E45F866BCBCEED86318ECF8F3549C70C8DC88EC01D9722EFC2E7B1B295C85CE0A2BC0E30BCC6524B1911DFEE2008357E11A4C29F27C42831635581BD326FAE2AB9FEC74DE31A3B24DDBBC88E5AB175864FD51198762AE146D3286556F3F48D51F3B7F0153BCC6B7EB6D7B5A8F10B9134628CD3C6BE979DC26FAA476536DEB0B698E7A6F0A0CE2C8AFE58CC1AFBF821234E5D21F9B47B023257C14DD78FE096DFF136165BF9C2BCD5125A2393A3971592CAAEB0DC2A36E7EC40EF057080E5A29C48E4CCE2ACF786042AD4005433491CCC7D55A1F15575AB41D530C1F7FD19F28475C9C9A4548FE7C3E4298D460D882648C9A5F07172482123C39907BD9A885F16D29EACCFD03D27F87E9563E99A6C2C7E798E603561E97AE3CDBB185A7970E655C6016908654B7FCC8EDF36A5E2AC94ACBD29BD2A83C20C87C19CADB29BD5F61690A845FB6EB947DA23FECFFDD4F9139492EE149614DB54850E02374EA602B42C726A65DFDA81C54B2E027880B49CCE2D16607AA87C08EE0986FED8540FAF2761C44C208DA20D659C86777FD26498A3CACBC1C4270BAF96A001072AA7F7AEC171D1D3E514F27D957B885EF217A32C65565D765794A6CE877D442CBFA4562554C6CA73A4510A401BFFFD4923AA560C414750FDAD7420DEBEBF232B8E26A50308E29B8A238724B2D14AAA156E4145EB3A151A145A3012102CFC8CA4B46C54B2652F1D29CEA9F905CE25DF61012DD58ADE9EA3B6ED24E713229222EB39FA58DD73A4B7674304AB8B2784B6F627826F530C6B1F381A634E0FC0560BAAAA10275B8A0936E0CC385B1377764459D5695DA7DBB061574363AA79D8EDC1B322D0879DF89D6E77ACD922F22E6CAD8340A581D5403E8840817CD0AA07CE8AD202FB76691C4813A8231BDEEEF75CAC34BA184CA6CDE69D3F142B18F31BA50471BA6EDA17A0BE6BFF5F45D74982F859639775110016E95D0FDBA72FE7F352D132AC80A8C926AF008D15022D4CC35268E721919B89E6AF590DDAA397692058787D721B88E5D464876B867E84593A1F4F336606B026BC4DA7AFC7C985208BE09C4097D2315813C353A6E37C13A4CF459A97D78D90C882B1D3B29701560CFCE5B2EDE98D7B6BFD015115C4CF71F4FB3BD65282B3F46C349C2A214319C734C7EC3D8BCEB3F014228511F7CAB7F3BAAE56364BA0862FAFB5FD2976C6F922357B42D01899DE2CC26E9752BD0BAA6675BF07FF4E07A509BFBBC2B954CC82040C171E5097116C378D7C4E3A79B8B3274785CD855AA53AA6664D2E1F989DA386A99708EA13B18977E553E8EE5962AD27FDA3CCFFAFED3296FD7EC753B2C13BB976C89785191CBAD8AD3B1144F33844FADA1FFA87FE05DB449A89B001896C964A73D1DA51829B2A89FEB9E4882C3705AA2CBF33FAC61A8B7C23A65F1ACF482956914F11EA49111306EA30D8D6D20530BD06907CCDBA0AACF0526D227597DB70A53B36B70E185AF6B4BE69CDE00AF660F2F29A6685C0001E2CC2203C48A0F833931D5FEDA149066C0B0584616FD5CC0782BE9B9330F94CC31DDE6AFF20CE4E746A4D1CFD8166EF60DC60061738F",
	"PageName":          "DefaultResult",
	"HandlerId":         "67",
	"DBCode":            "SCDB",
	"KuaKuCodes":        "CJFQ,CDMD,CIPD,CCND,CISD,SNAD,BDZK,CCVD,CJFN,CCJD",
	"CurPage":           "1",
	"RecordsCntPerPage": "20",
	"CurDisplayMode":    "listmode",
	"CurrSortField":     "%e5%8f%91%e8%a1%a8%e6%97%b6%e9%97%b4%2f(%e5%8f%91%e8%a1%a8%e6%97%b6%e9%97%b4%2c%27TIME%27)",
	"CurrSortFieldType": "desc",
	"IsSortSearch":      "false",
	"IsSentenceSearch":  "false",
	"Subject":           "",
}
