package request

type UpdateUserAddressReq struct {
	Address  string `json:"address"`
	Location string `json:"location"`
}
