/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role.
type ProductSpecification struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// Complements the description of an element (for instance a product) through video, pictures...
	Attachment []AttachmentRef `json:"attachment,omitempty"`
	// The manufacturer or trademark of the specification
	Brand string `json:"brand,omitempty"`
	// A type of ProductSpecification that belongs to a grouping of ProductSpecifications made available to the market. It inherits of all attributes of ProductSpecification.
	BundledProductSpecification []BundledProductSpecification `json:"bundledProductSpecification,omitempty"`

	Code string `json:"code,omitempty"`

	CreatedBy string `json:"createdBy,omitempty"`

	CreatedOn time.Time `json:"createdOn,omitempty"`
	// A narrative that explains in detail what the product specification is
	Description string `json:"description,omitempty"`

	ExternalID string `json:"externalID,omitempty"`
	// Reference of the product specification
	Href string `json:"href,omitempty"`
	// Unique identifier of the product specification
	Id string `json:"id,omitempty"`
	// isBundle determines whether a productSpecification represents a single productSpecification (false), or a bundle of productSpecification (true).
	IsBundle bool `json:"isBundle,omitempty"`
	// Date and time of the last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// Name of the product specification
	Name string `json:"name,omitempty"`
	// An identification number assigned to uniquely identity the specification
	ProductNumber string `json:"productNumber,omitempty"`
	// A characteristic quality or distinctive feature of a ProductSpecification.  The characteristic can be take on a discrete value, such as color, can take on a range of values, (for example, sensitivity of 100-240 mV), or can be derived from a formula (for example, usage time (hrs) = 30 - talk time *3). Certain characteristics, such as color, may be configured during the ordering or some other process.
	ProductSpecCharacteristic []SpecCharacteristic `json:"productSpecCharacteristic,omitempty"`
	// A migration, substitution, dependency or exclusivity relationship between/among product specifications.
	ProductSpecificationRelationship []ProductSpecificationRelationship `json:"productSpecificationRelationship,omitempty"`
	// A related party defines party or party role linked to a specific entity.
	RelatedParty []RelatedParty `json:"relatedParty,omitempty"`

	ResourceSpecification []SpecificationRef `json:"resourceSpecification,omitempty"`

	ServiceSpecification []SpecificationRef `json:"serviceSpecification,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// Product specification version
	Version string `json:"version,omitempty"`
}
