# TopicLogRetention

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**RetentionTime** | **int32** | The time in milliseconds that a message is retained in the topic log. Messages older than the retention time are deleted. If value is &#x60;0&#x60;, messages are retained indefinitely unless other retention is set.  | [default to 0]|
|**SegmentBytes** | **int32** | The maximum size in bytes that the topic log can grow to. When the log reaches this size, the oldest messages are deleted. If value is &#x60;0&#x60;, messages are retained indefinitely unless other retention is set.  | [default to 0]|

## Methods

### NewTopicLogRetention

`func NewTopicLogRetention(retentionTime int32, segmentBytes int32, ) *TopicLogRetention`

NewTopicLogRetention instantiates a new TopicLogRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicLogRetentionWithDefaults

`func NewTopicLogRetentionWithDefaults() *TopicLogRetention`

NewTopicLogRetentionWithDefaults instantiates a new TopicLogRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionTime

`func (o *TopicLogRetention) GetRetentionTime() int32`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *TopicLogRetention) GetRetentionTimeOk() (*int32, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *TopicLogRetention) SetRetentionTime(v int32)`

SetRetentionTime sets RetentionTime field to given value.


### GetSegmentBytes

`func (o *TopicLogRetention) GetSegmentBytes() int32`

GetSegmentBytes returns the SegmentBytes field if non-nil, zero value otherwise.

### GetSegmentBytesOk

`func (o *TopicLogRetention) GetSegmentBytesOk() (*int32, bool)`

GetSegmentBytesOk returns a tuple with the SegmentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentBytes

`func (o *TopicLogRetention) SetSegmentBytes(v int32)`

SetSegmentBytes sets SegmentBytes field to given value.



