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

type Note struct {

	BaseType string `json:"@baseType,omitempty"`

	SchemaLocation string `json:"@schemaLocation,omitempty"`

	Type_ string `json:"@type,omitempty"`

	Author string `json:"author,omitempty"`

	Date time.Time `json:"date,omitempty"`

	Id string `json:"id,omitempty"`

	Text string `json:"text,omitempty"`
}