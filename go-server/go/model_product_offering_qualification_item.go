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

// Representing a Product Offering Qualification Item
type ProductOfferingQualificationItem struct {

	BaseType string `json:"@baseType,omitempty"`

	SchemaLocation string `json:"@schemaLocation,omitempty"`

	Type_ string `json:"@type,omitempty"`
	// action for the Product Offering Qualification Item
	Action string `json:"action,omitempty"`

	Category *CategoryRef `json:"category,omitempty"`

	ExpectedActivationDate time.Time `json:"expectedActivationDate,omitempty"`
	// Unique identifier of the Product Offering Qualification Item
	Id string `json:"id,omitempty"`

	Note []Note `json:"note,omitempty"`

	ProductOffering *ProductOfferingRef `json:"productOffering,omitempty"`

	QualificationItemResult string `json:"qualificationItemResult,omitempty"`

	State string `json:"state,omitempty"`
}