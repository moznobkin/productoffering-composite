# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/MyTeam889/Life/0.0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubscription**](ServiceAPIApi.md#DeleteSubscription) | **Delete** /API/v1/subscriptions/{serviceid} | Delete subscription
[**GetOffers**](ServiceAPIApi.md#GetOffers) | **Get** /API/v1/offerslist | returns all published offers
[**GetSubscriptions**](ServiceAPIApi.md#GetSubscriptions) | **Get** /API/v1/subscriptionslist | Retrieves subscriptions
[**PostSubscription**](ServiceAPIApi.md#PostSubscription) | **Post** /API/v1/subscriptions/{serviceid} | Create subscription

# **DeleteSubscription**
> OkResponse DeleteSubscription(ctx, serviceid, optional)
Delete subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceid** | **string**| VAS-P subscription identifier | 
 **optional** | ***ServiceAPIApiDeleteSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAPIApiDeleteSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubscriptionInfo**](SubscriptionInfo.md)|  | 

### Return type

[**OkResponse**](OkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOffers**
> OffersListResponse GetOffers(ctx, optional)
returns all published offers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceAPIApiGetOffersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAPIApiGetOffersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msisdn** | **optional.String**| msisdn | 
 **category** | **optional.String**| Category name. If not specified - all | 

### Return type

[**OffersListResponse**](OffersListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptions**
> Subscription GetSubscriptions(ctx, optional)
Retrieves subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceAPIApiGetSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAPIApiGetSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msisdn** | **optional.String**| msisdn | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSubscription**
> OkResponse PostSubscription(ctx, serviceid, optional)
Create subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceid** | **string**| VAS-P subscription identifier | 
 **optional** | ***ServiceAPIApiPostSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAPIApiPostSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubscriptionInfo**](SubscriptionInfo.md)|  | 

### Return type

[**OkResponse**](OkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

