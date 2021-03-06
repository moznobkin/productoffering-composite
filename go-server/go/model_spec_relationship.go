/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SpecRelationship struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// Reference of the target serviceSpecification
	Href string `json:"href,omitempty"`
	// Unique identifier of the target serviceSpecification
	Id string `json:"id,omitempty"`
	// The name given to the target service specification instance
	Name string `json:"name,omitempty"`
	// Type of relationship such as migration, substitution, dependency, exclusivity
	RelationshipType string `json:"relationshipType,omitempty"`
	// The association role for this service specification
	Role string `json:"role,omitempty"`
	// The period for which the ser
	ValidFor string `json:"validFor,omitempty"`
}
