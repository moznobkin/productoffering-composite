# CategoryOffersOffers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | unique identifier | [optional] [default to null]
**Name** | **string** | offer name | [optional] [default to null]
**Shortdescription** | **string** | short description | [optional] [default to null]
**Detaileddescription** | **string** | full description | [optional] [default to null]
**Miniature** | **string** | reference to the offer image in the list | [optional] [default to null]
**Header** | **string** | reference to the offer image in the full description | [optional] [default to null]
**Offertype** | **string** |  | [optional] [default to null]
**Serviceurl** | **string** | reference to service, required if offertype &#x3D; url | [optional] [default to null]
**Buttontext** | **string** | the text of the button | [optional] [default to null]
**Priority** | **int32** | offer priority, bigger number means higher priority | [optional] [default to null]
**Labels** | [**[]CategoryOffersLabels**](CategoryOffers_labels.md) | array of labels data which will be used for offer | [optional] [default to null]
**Subscriptions** | [**[]Subscription**](Subscription.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

