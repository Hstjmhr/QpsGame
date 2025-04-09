package response

import (
	"hall/model/request"
	common "msqp"
)

type UpdateUserAddressRes struct {
	common.Result
	UpdateUserData request.UpdateUserAddressReq `json:"updateUserData"`
}
