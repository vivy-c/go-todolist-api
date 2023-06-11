package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (req *ToDoCreateReq) Validate() error {
	return validation.ValidateStruct(req,
		// Validasi field password
		validation.Field(&req.Todo,
			validation.Required.Error("Activities is required"),
			// validation.Length(6, 20).Error("Password harus memiliki panjang antara 6 hingga 20 karakter"),
		),
	)
}
