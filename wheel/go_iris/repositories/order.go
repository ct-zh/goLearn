package repositories

import (
	"database/sql"
	"go_iris/common"
	"go_iris/datamodels"
	"strconv"
)

type IOrder interface {
	Conn() error
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderManager struct {
	table  string
	dbConn *sql.DB
}

func (o OrderManager) SelectAllWithInfo() (OrderMap map[int]map[string]string, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "select `order`.id, product.productName, `order`.orderStatus from `order` left join product on `order`.productId=product.id"
	rows, errRows := o.dbConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}

	return common.GetResultRows(rows), nil
}

func (o OrderManager) Conn() error {
	if o.dbConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.dbConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

func (o OrderManager) Insert(order *datamodels.Order) (productId int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}

	sql := "INSERT `order` SET userId=?,productId=?,orderStatus=?"

	stmt, errStmt := o.dbConn.Prepare(sql)
	if errStmt != nil {
		return productId, errStmt
	}

	result, errResult := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if errResult != nil {
		return productId, errResult
	}
	return result.LastInsertId()
}

func (o OrderManager) Delete(orderId int64) (isOk bool) {
	if err := o.Conn(); err != nil {
		return
	}
	sql := "delete from " + o.table + " where id =?"
	stmt, errStmt := o.dbConn.Prepare(sql)
	if errStmt != nil {
		return
	}
	_, err := stmt.Exec(orderId)
	if err != nil {
		return
	}
	return true
}

func (o OrderManager) Update(order *datamodels.Order) error {
	if errConn := o.Conn(); errConn != nil {
		return errConn
	}

	sql := "Update " + o.table + " set userId=?,productId=?,orderStatus=? Where id=" + strconv.FormatInt(order.ID, 10)
	stmt, errStmt := o.dbConn.Prepare(sql)
	if errStmt != nil {
		return errStmt
	}
	_, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	return err
}

func (o OrderManager) SelectByKey(orderId int64) (order *datamodels.Order, err error) {
	if errConn := o.Conn(); errConn != nil {
		return &datamodels.Order{}, errConn
	}

	sql := "Select * From " + o.table + " where ID=" + strconv.FormatInt(orderId, 10)
	row, errRow := o.dbConn.Query(sql)
	if errRow != nil {
		return &datamodels.Order{}, errRow
	}

	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Order{}, err
	}

	order = &datamodels.Order{}
	common.DataToStructByTagSql(result, order)
	return
}

func (o OrderManager) SelectAll() (orderArray []*datamodels.Order, err error) {
	if errConn := o.Conn(); errConn != nil {
		return nil, errConn
	}
	sql := "Select * from `" + o.table + "`"

	rows, errRows := o.dbConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, err
	}

	for _, v := range result {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orderArray = append(orderArray, order)
	}
	return
}

func NewOrderManager(t string, mysql *sql.DB) IOrder {
	return &OrderManager{
		table:  t,
		dbConn: mysql,
	}
}
