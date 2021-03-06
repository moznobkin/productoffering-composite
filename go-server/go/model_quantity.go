/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An amount in a given unit
type Quantity struct {
	// Numeric value in a given unit
	Amount float32 `json:"amount,omitempty"`
	// Unit
	Units string `json:"units,omitempty"`
}
