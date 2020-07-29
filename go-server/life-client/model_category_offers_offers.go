/*
 * Product aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger_life_client

type CategoryOffersOffers struct {
	// unique identifier
	Id int64 `json:"id,omitempty"`
	// offer name
	Name string `json:"name,omitempty"`
	// short description
	Shortdescription string `json:"shortdescription,omitempty"`
	// full description
	Detaileddescription string `json:"detaileddescription,omitempty"`
	// reference to the offer image in the list
	Miniature string `json:"miniature,omitempty"`
	// reference to the offer image in the full description
	Header string `json:"header,omitempty"`
	Offertype string `json:"offertype,omitempty"`
	// reference to service, required if offertype = url
	Serviceurl string `json:"serviceurl,omitempty"`
	// the text of the button
	Buttontext string `json:"buttontext,omitempty"`
	// offer priority, bigger number means higher priority
	Priority int32 `json:"priority,omitempty"`
	// array of labels data which will be used for offer
	Labels []CategoryOffersLabels `json:"labels,omitempty"`
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}
