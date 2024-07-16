# KafkaClusterConnectionAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**BrokerAddresses** | Pointer to **[]string** | IP addresses and subnet of cluster brokers. Note the following unavailable IP range: 10.224.0.0/11  | [optional] |

## Methods

### NewKafkaClusterConnectionAllOf

`func NewKafkaClusterConnectionAllOf() *KafkaClusterConnectionAllOf`

NewKafkaClusterConnectionAllOf instantiates a new KafkaClusterConnectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaClusterConnectionAllOfWithDefaults

`func NewKafkaClusterConnectionAllOfWithDefaults() *KafkaClusterConnectionAllOf`

NewKafkaClusterConnectionAllOfWithDefaults instantiates a new KafkaClusterConnectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerAddresses

`func (o *KafkaClusterConnectionAllOf) GetBrokerAddresses() []string`

GetBrokerAddresses returns the BrokerAddresses field if non-nil, zero value otherwise.

### GetBrokerAddressesOk

`func (o *KafkaClusterConnectionAllOf) GetBrokerAddressesOk() (*[]string, bool)`

GetBrokerAddressesOk returns a tuple with the BrokerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAddresses

`func (o *KafkaClusterConnectionAllOf) SetBrokerAddresses(v []string)`

SetBrokerAddresses sets BrokerAddresses field to given value.

### HasBrokerAddresses

`func (o *KafkaClusterConnectionAllOf) HasBrokerAddresses() bool`

HasBrokerAddresses returns a boolean if a field has been set.


