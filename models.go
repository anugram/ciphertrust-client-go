package ciphertrust

type User struct {
	UserID   types.String `tfsdk:"user_id"`
	Name     types.String `tfsdk:"username"`
	UserName types.String `tfsdk:"nickname"`
	Nickname types.String `tfsdk:"email"`
	Email    types.String `tfsdk:"name"`
}
