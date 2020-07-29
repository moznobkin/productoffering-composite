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

// Product Offering Qualification
type ProductOfferingQualification struct {

	BaseType string `json:"@baseType,omitempty"`

	SchemaLocation string `json:"@schemaLocation,omitempty"`

	Type_ string `json:"@type,omitempty"`

	Category *CategoryRef `json:"category,omitempty"`

	Channel *ChannelRef `json:"channel,omitempty"`

	Context *Context `json:"context,omitempty"`
	// Description of the productOfferingQualifcation
	Description string `json:"description,omitempty"`

	EffectiveQualificationDate time.Time `json:"effectiveQualificationDate,omitempty"`

	ExpectedPOQCompletionDate time.Time `json:"expectedPOQCompletionDate,omitempty"`

	ExpirationDate time.Time `json:"expirationDate,omitempty"`
	// Hyperlink to access the productOfferingQualification
	Href string `json:"href,omitempty"`
	// Unique identifier of the productOfferingQualification resource
	Id string `json:"id,omitempty"`

	InstantSyncQualification bool `json:"instantSyncQualification,omitempty"`

	Note []Note `json:"note,omitempty"`

	Place []PlaceRef `json:"place,omitempty"`

	ProductOfferingQualificationDate time.Time `json:"productOfferingQualificationDate,omitempty"`
	// A list of Product Offering Qualification Item Container for the Product Offering Qualification
	ProductOfferingQualificationItemContainer []PoqItemContainer `json:"productOfferingQualificationItemContainer,omitempty"`
	// An indicator which when the value is true means that alternative solutions should be provided
	ProvideAlternative bool `json:"provideAlternative,omitempty"`

	ProvideOnlyAvailable bool `json:"provideOnlyAvailable,omitempty"`
	// An indicator which when the value is true means that unavailability reason are expected for non available product offering
	ProvideUnavailabilityReason bool `json:"provideUnavailabilityReason,omitempty"`

	QualificationResult string `json:"qualificationResult,omitempty"`

	RelatedParty []RelatedParty `json:"relatedParty,omitempty"`

	RequestedPOQCompletionDate time.Time `json:"requestedPOQCompletionDate,omitempty"`

	State string `json:"state,omitempty"`
}
