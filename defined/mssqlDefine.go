package defined

import "time"

type GoodsList struct {
	Goods_site        int64     `db:"goods_code"`
	Goods_code        int64     `db:"goods_code"`
	Goods_bktop       int64     `db:"goods_bktop"`
	Goods_bkmid       int64     `db:"goods_bkmid"`
	Goods_bkbot       int64     `db:"goods_bkbot"`
	Goods_name        string    `db:"goods_name"`
	Goods_bcode       string    `db:"goods_bcode"`
	Goods_bprice      float64   `db:"goods_bprice"`
	Goods_sprice      int64     `db:"goods_sprice"`
	Goods_stype       int64     `db:"goods_stype"`
	Goods_stock       int64     `db:"goods_stock"`
	Goods_stocktype   int64     `db:"goods_stocktype"`
	Goods_lastdate    time.Time `db:"goods_lastdate"`
	Goods_sspec       string    `db:"goods_sspec"`
	Goods_mileage     int64     `db:"goods_mileage"`
	Goods_tax         int64     `db:"goods_tax"`
	Goods_status      int64     `db:"goods_status"`
	Goods_rdate       time.Time `db:"goods_rdate"`
	Goods_edate       time.Time `db:"goods_edate"`
	Goods_sup         int64     `db:"goods_sup"`
	Goods_man         int64     `db:"goods_man"`
	Goods_remark      string    `db:"goods_remark"`
	Goods_keyword     string    `db:"goods_keyword"`
	Goods_emp         int64     `db:"goods_emp"`
	Goods_abc         int64     `db:"goods_abc"`
	Goods_pricedate   time.Time `db:"goods_pricedate"`
	Goods_minstock    int64     `db:"goods_minstock"`
	Goods_minpur      int64     `db:"goods_minpur"`
	Goods_dayqty      float64   `db:"goods_dayqty"`
	Goods_bguaranty   int64     `db:"goods_bguaranty"`
	Goods_sguaranty   int64     `db:"goods_sguaranty"`
	Goods_netqty      string    `db:"goods_netqty"`
	Goods_unitqty     string    `db:"goods_unitqty"`
	Goods_type        int64     `db:"goods_type"`
	Goods_fee         float64   `db:"goods_fee"`
	Goods_stprice     float64   `db:"goods_stprice"`
	Goods_dc          float64   `db:"goods_dc"`
	Goods_pricepolicy int64     `db:"goods_pricepolicy"`
	Goods_stockdate   time.Time `db:"goods_stockdate"`
	Goods_brand       int64     `db:"goods_brand"`
	Goods_storetype   int64     `db:"goods_storetype"`
	NS_VF             int64     `db:"NS_VF"`
	Goods_margin      float64   `db:"goods_margin"`
	Goods_bpricedate  time.Time `db:"goods_bpricedate"`
	Goods_purdate     time.Time `db:"goods_purdate"`
	Send_gb           string    `db:"send_gb"`
	Goods_wsprice1    int64     `db:"goods_wsprice1"`
	Goods_wsprice2    int64     `db:"goods_wsprice2"`
	Goods_wsprice3    int64     `db:"goods_wsprice3"`
	Goods_wsprice4    int64     `db:"goods_wsprice4"`
	Goods_wsprice5    int64     `db:"goods_wsprice5"`
	Goods_agelimit    int64     `db:"goods_agelimit"`
	Goods_origin      string    `db:"goods_origin"`
	Gl_date_str       time.Time `db:"gl_date_str"`
}

type Attr struct {
	Mart_id  string `db:"mart_id"`
	Bcode    string `db:"bcode"`
	Price    int64  `db:"price"`
	Tax_free int64 `db:"tax_free"`
}


// goods_bcode
// goods_sprice
// goods_tax
// gl_date

// bcode : item.goods_bcode.trim(),
// price : item.goods_sprice,
// tax_free : item.goods_tax,
// gl_date : item.gl_date