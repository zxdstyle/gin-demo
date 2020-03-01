package v1

type UpdateTagRequest struct {
	Name string	`form:"name" json:"name" valid:"MaxSize(255)" attr:"标签名"`
	Status int	`form:"status" json:"status" valid:"Range(0,1)" attr:"状态"`
}
