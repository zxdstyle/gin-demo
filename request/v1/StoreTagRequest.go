package v1

type StoreTagRequest struct {
	Name string	`form:"name" json:"name" valid:"Required" attr:"标签名"`
	Status int	`form:"status" json:"status" valid:"Range(0,1)" attr:"状态"`
}
