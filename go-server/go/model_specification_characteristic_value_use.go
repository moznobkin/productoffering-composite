/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A use of the SpecificationCharacteristicValue by a ProductOffering to which additional properties (attributes) apply or override the properties of similar properties contained in SpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding specification. The available characteristic values for a SpecificationCharacteristic in a specification can be modified at the ProductOffering level. For example, a characteristic 'Color' might have values White, Blue, Green, and Red. But, the list of values can be restricted to e.g. White and Blue in an associated product offering. It should be noted that the list of values in 'SpecificationCharacteristicValueUse' is a strict subset of the list of values as defined in the corresponding specification characteristics.
type SpecificationCharacteristicValueUse struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	//  hyperlink reference to the schema describing this resource.
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// A narrative that explains in detail what the productSpecificationCharacteristic is
	Description string `json:"description,omitempty"`
	// The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality.
	MaxCardinality int32 `json:"maxCardinality,omitempty"`
	// The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality.
	MinCardinality int32 `json:"minCardinality,omitempty"`
	// Name of the associated productSpecificationCharacteristic
	Name string `json:"name"`

	ReferenceId string `json:"referenceId,omitempty"`
	// A number or text that can be assigned to a SpecificationCharacteristic.
	SpecCharacteristicValue []SpecificationCharacteristicValue `json:"specCharacteristicValue"`

	Specification *SpecificationRef `json:"specification,omitempty"`
	// When sub-classing, this defines the sub-class specification type
	SpecificationType string `json:"specificationType,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// A kind of value that the characteristic can take on, such as numeric, text and so forth
	ValueType string `json:"valueType,omitempty"`
}