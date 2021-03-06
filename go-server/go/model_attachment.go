/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Complements the description of an element (for instance a product) through video, pictures...
type Attachment struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// Attachment type such as video, picture
	AttachmentType string `json:"attachmentType,omitempty"`
	// The actual contents of the attachment object, if embedded, encoded as base64
	Content string `json:"content,omitempty"`
	// A narrative text describing the content of the attachment
	Description string `json:"description,omitempty"`
	// URI for this Attachment
	Href string `json:"href,omitempty"`
	// Unique identifier for this particular attachment
	Id string `json:"id"`
	// Attachment mime type such as extension file for video, picture and document
	MimeType string `json:"mimeType,omitempty"`
	// The name of the attachment
	Name string `json:"name,omitempty"`

	Size *Quantity `json:"size,omitempty"`
	// Uniform Resource Locator, is a web page address (a subset of URI)
	Url string `json:"url,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
}
