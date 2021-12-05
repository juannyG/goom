package product

import (
	"database/sql"
)

type Product struct {
        ExternalProductId string
        AutoshipEnabled   bool
        Live              bool
}

const query = `
SELECT p.external_product_id, p.autoship_enabled, p.live
   FROM product_product AS p, merchant_merchant
   WHERE merchant_merchant.id=p.merchant_id
   AND merchant_merchant.public_id = ?
   AND p.external_product_id = ?
`

func GetProduct(db *sql.DB, merchantPublicId string, externalProductId string) (Product, error) {
	var p Product
	stmt, err := db.Prepare(query)
	if err != nil {
		return p, err
	}
	defer stmt.Close()

	err = stmt.
		QueryRow(merchantPublicId, externalProductId).
		Scan(&p.ExternalProductId, &p.AutoshipEnabled, &p.Live)
	return p, err
}
