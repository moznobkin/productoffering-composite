/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Attachment reference. An attachment complements the description of an element (for instance a product) through video, pictures
type AttachmentRef struct {
	// A string. The (immediate) base class type of this REST resource
	BaseType string `json:"@baseType,omitempty"`
	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty"`
	// A string class type of target specification
	Type_ string `json:"@type,omitempty"`
	// A narrative text describing the content of the attachment
	Description string `json:"description,omitempty"`
	// URL serving as reference for the attachment resource
	Href string `json:"href"`
	// Unique-Identifier for this attachment
	Id string `json:"id"`
	// Link to the attachment media/content
	Url string `json:"url,omitempty"`
}
