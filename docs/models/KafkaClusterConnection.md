# KafkaClusterConnection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The datacenter to connect your instance to. | |
|**LanId** | **string** | The numeric LAN ID to connect your instance to. | |
|**Cidr** | **string** | The IP and subnet for your instance. Note the following unavailable IP range: 10.244.0.0/11  | |
|**BrokerAddresses** | Pointer to **[]string** | IP addresses and subnet of cluster brokers. Note the following unavailable IP range: 10.224.0.0/11  | [optional] |

## Methods

### NewKafkaClusterConnection

`func NewKafkaClusterConnection(datacenterId string, lanId string, cidr string, ) *KafkaClusterConnection`

NewKafkaClusterConnection instantiates a new KafkaClusterConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaClusterConnectionWithDefaults

`func NewKafkaClusterConnectionWithDefaults() *KafkaClusterConnection`

NewKafkaClusterConnectionWithDefaults instantiates a new KafkaClusterConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *KafkaClusterConnection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *KafkaClusterConnection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *KafkaClusterConnection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *KafkaClusterConnection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *KafkaClusterConnection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *KafkaClusterConnection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetCidr

`func (o *KafkaClusterConnection) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *KafkaClusterConnection) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *KafkaClusterConnection) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetBrokerAddresses

`func (o *KafkaClusterConnection) GetBrokerAddresses() []string`

GetBrokerAddresses returns the BrokerAddresses field if non-nil, zero value otherwise.

### GetBrokerAddressesOk

`func (o *KafkaClusterConnection) GetBrokerAddressesOk() (*[]string, bool)`

GetBrokerAddressesOk returns a tuple with the BrokerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAddresses

`func (o *KafkaClusterConnection) SetBrokerAddresses(v []string)`

SetBrokerAddresses sets BrokerAddresses field to given value.

### HasBrokerAddresses

`func (o *KafkaClusterConnection) HasBrokerAddresses() bool`

HasBrokerAddresses returns a boolean if a field has been set.


