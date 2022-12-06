package product

const (
	readProduct = `SELECT id, name, price, size, quantity WHERE id = ? AND status = 1`

	readProductList = `SELECT id, name, price, size, quantity WHERE status = 1`
)
