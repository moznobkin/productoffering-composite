/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Representing a Product Offering Qualification Item Container
type PoqItemContainer struct {

	BaseType string `json:"@baseType,omitempty"`

	SchemaLocation string `json:"@schemaLocation,omitempty"`

	Type_ string `json:"@type,omitempty"`

	Category *CategoryRef `json:"category,omitempty"`
	// Unique identifier of the Product Offering Qualification Item Container
	Id string `json:"id,omitempty"`
	// Process Context of the Product Offering Qualification Item Container
	ProcessContext string `json:"processContext,omitempty"`
	// A list of Product Offering Qualification Item for the Product Offering Qualification 
	ProductOfferingQualificationItem []ProductOfferingQualificationItem `json:"productOfferingQualificationItem,omitempty"`
}
