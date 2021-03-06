/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A type of ProductSpecification that belongs to a grouping of ProductSpecifications made available to the market. It inherits of all attributes of ProductSpecification.
type BundledProductSpecification struct {
	// Reference of the product specification
	Href string `json:"href,omitempty"`
	// Unique identifier of the product specification
	Id string `json:"id,omitempty"`
	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// Name of the product specification
	Name string `json:"name,omitempty"`
}
