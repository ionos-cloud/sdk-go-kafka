# Topic

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the Kafka cluster topic. Must be 63 characters or less and must begin and end with an alphanumeric character (&#x60;[a-z0-9A-Z]&#x60;) with dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), dots (&#x60;.&#x60;), and alphanumerics between.  | |
|**ReplicationFactor** | **int32** | The number of replicas of the topic. The replication factor determines how many copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number of brokers in the Kafka cluster.  | |
|**NumberOfPartitions** | **int32** | The number of partitions of the topic. Partitions allow for parallel processing of messages. The partition count must be greater than or equal to the replication factor.  | |
|**LogRetention** | [**TopicLogRetention**](TopicLogRetention.md) |  | |

## Methods

### NewTopic

`func NewTopic(name string, replicationFactor int32, numberOfPartitions int32, logRetention TopicLogRetention, ) *Topic`

NewTopic instantiates a new Topic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicWithDefaults

`func NewTopicWithDefaults() *Topic`

NewTopicWithDefaults instantiates a new Topic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Topic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Topic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Topic) SetName(v string)`

SetName sets Name field to given value.


### GetReplicationFactor

`func (o *Topic) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *Topic) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *Topic) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.


### GetNumberOfPartitions

`func (o *Topic) GetNumberOfPartitions() int32`

GetNumberOfPartitions returns the NumberOfPartitions field if non-nil, zero value otherwise.

### GetNumberOfPartitionsOk

`func (o *Topic) GetNumberOfPartitionsOk() (*int32, bool)`

GetNumberOfPartitionsOk returns a tuple with the NumberOfPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPartitions

`func (o *Topic) SetNumberOfPartitions(v int32)`

SetNumberOfPartitions sets NumberOfPartitions field to given value.


### GetLogRetention

`func (o *Topic) GetLogRetention() TopicLogRetention`

GetLogRetention returns the LogRetention field if non-nil, zero value otherwise.

### GetLogRetentionOk

`func (o *Topic) GetLogRetentionOk() (*TopicLogRetention, bool)`

GetLogRetentionOk returns a tuple with the LogRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetention

`func (o *Topic) SetLogRetention(v TopicLogRetention)`

SetLogRetention sets LogRetention field to given value.



