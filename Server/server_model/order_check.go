package serverModel

import "errors"

//需要了解index标签是什么意思
type AddOrderReq struct {
	UserName string `gorm:"index:idx_order_username" json:"user_name"`
	Amount float64 `json:"amount"`
	FileUrl string `json:"file_url"`
}

//错误检测
func (r *AddOrderReq)  IsValid() error{
	switch  {
	case r.Amount == 0:
		return errors.New("amount is not zero")
	case r.UserName == "":
		return errors.New("user_name is not found")
	default:
		return nil
	}
}

type Order struct {
	ID uint `grom:"primarykey"`
	OrderNo string `gorm:"index:idx_order_no" json:"order_no"`
	UserName string `gorm:"index:idx_order_username" json:"user_name"`
	Amount float64 `json:"amount"`
	FileUrl string `json:"file_url"`
}

func (r *Order) IsValid() error  {
	switch  {
	case r.OrderNo == "":
		return errors.New("order_no is not found")
	case r.Amount == 0:
		return errors.New("amount is not zero")
	default:
		return nil
	}

}



func (r *Order) IsDeleteValid() error  {
	switch  {
	case r.OrderNo == "":
		return errors.New("order_no is not found")
	default:
		return nil
	}
}
