package repositories

import (
	"database/sql"
	"fmt"
	"go_iris/common"
	"go_iris/datamodels"
	"strconv"
)

// 定义接口与方法
type IProduct interface {
	Conn() error
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

//[duck typing] 实现IProduct的接口
type ProductManage struct {
	table  string
	dbConn *sql.DB
}

// 检测数据库连接
func (p ProductManage) Conn() (err error) {
	if p.dbConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return
		}
		p.dbConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

// 插入一条数据
func (p ProductManage) Insert(product *datamodels.Product) (productId int64, err error) {
	// 1. 判断连接是否存在
	if err = p.Conn(); err != nil {
		return
	}

	// 2. 拼sql
	stmt, err := p.dbConn.Prepare(`INSERT product SET productName=?, productNum=?, productImage=?, productUrl=?`)
	if err != nil {
		return
	}

	// 3. exec
	result, err := stmt.Exec(
		product.ProductName,
		product.ProductNum,
		product.ProductImage,
		product.ProductUrl)
	if err != nil {
		return
	}

	return result.LastInsertId()
}

// 删除一条数据
func (p ProductManage) Delete(id int64) bool {
	// 1. 判断连接是否存在
	err := p.Conn()
	if err != nil {
		return false
	}

	stmt, err := p.dbConn.Prepare(`delete from product where id=?`)
	if err != nil {
		return false
	}

	_, err = stmt.Exec(strconv.FormatInt(id, 10))
	if err != nil {
		return false
	}
	return true
}

// 更新一条数据
func (p ProductManage) Update(product *datamodels.Product) (err error) {
	if err = p.Conn(); err != nil {
		return
	}

	if product.Id <= 0 {
		return fmt.Errorf("id is invaild")
	}

	mySql := `UPDATE product set 
productName=?, productNum=?, productImage=?, productUrl=? 
where id=` + strconv.FormatInt(product.Id, 10)
	stmt, err := p.dbConn.Prepare(mySql)
	if err != nil {
		return
	}

	_, err = stmt.Exec(
		product.ProductName,
		product.ProductNum,
		product.ProductImage,
		product.ProductUrl)
	if err != nil {
		return
	}

	return nil
}

// 获取一条数据
func (p ProductManage) SelectByKey(id int64) (product *datamodels.Product, err error) {
	product = &datamodels.Product{}

	if err = p.Conn(); err != nil {
		return product, err
	}

	row, err := p.dbConn.Query(`SELECT * FROM product WHERE id=` + strconv.FormatInt(id, 10))
	defer row.Close()
	if err != nil {
		return product, err
	}

	result := common.GetResultRow(row)
	if len(result) == 0 {
		return product, nil
	}

	common.DataToStructByTagSql(result, product)
	return product, nil
}

// 获取所有数据
func (p ProductManage) SelectAll() (products []*datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}

	rows, err := p.dbConn.Query(`select * FROM product`)
	if err != nil {
		return
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		products = append(products, product)
	}
	return
}

// new instance
func NewProductManage(table string, mysql *sql.DB) IProduct {
	return &ProductManage{
		table:  table,
		dbConn: mysql,
	}
}
