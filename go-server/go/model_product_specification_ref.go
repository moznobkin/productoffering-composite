/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Product specification reference: A Specification is a detailed description of a tangible or intangible object made available externally in the form of a Offering to customers or other parties playing a party role.
type ProductSpecificationRef struct {
	// Reference of the product specification.
	Href string `json:"href,omitempty"`
	// Unique identifier of the product specification.
	Id string `json:"id"`
	// Name of the product specification.
	Name string `json:"name,omitempty"`

	ResourceSpecification []ResourceSpecification `json:"resourceSpecification,omitempty"`

	ServiceSpecification []ServiceSpecification `json:"serviceSpecification,omitempty"`
	// Product specification version.
	Version string `json:"version,omitempty"`
}