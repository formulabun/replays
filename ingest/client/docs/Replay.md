# Replay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**GameMap** | Pointer to **float32** |  | [optional] 
**Time** | Pointer to **float32** |  | [optional] 
**Lap** | Pointer to **float32** |  | [optional] 
**PlayerEntries** | Pointer to [**[]PlayerEntry**](PlayerEntry.md) |  | [optional] 

## Methods

### NewReplay

`func NewReplay() *Replay`

NewReplay instantiates a new Replay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplayWithDefaults

`func NewReplayWithDefaults() *Replay`

NewReplayWithDefaults instantiates a new Replay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Replay) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Replay) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Replay) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Replay) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetGameMap

`func (o *Replay) GetGameMap() float32`

GetGameMap returns the GameMap field if non-nil, zero value otherwise.

### GetGameMapOk

`func (o *Replay) GetGameMapOk() (*float32, bool)`

GetGameMapOk returns a tuple with the GameMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMap

`func (o *Replay) SetGameMap(v float32)`

SetGameMap sets GameMap field to given value.

### HasGameMap

`func (o *Replay) HasGameMap() bool`

HasGameMap returns a boolean if a field has been set.

### GetTime

`func (o *Replay) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Replay) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Replay) SetTime(v float32)`

SetTime sets Time field to given value.

### HasTime

`func (o *Replay) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetLap

`func (o *Replay) GetLap() float32`

GetLap returns the Lap field if non-nil, zero value otherwise.

### GetLapOk

`func (o *Replay) GetLapOk() (*float32, bool)`

GetLapOk returns a tuple with the Lap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLap

`func (o *Replay) SetLap(v float32)`

SetLap sets Lap field to given value.

### HasLap

`func (o *Replay) HasLap() bool`

HasLap returns a boolean if a field has been set.

### GetPlayerEntries

`func (o *Replay) GetPlayerEntries() []PlayerEntry`

GetPlayerEntries returns the PlayerEntries field if non-nil, zero value otherwise.

### GetPlayerEntriesOk

`func (o *Replay) GetPlayerEntriesOk() (*[]PlayerEntry, bool)`

GetPlayerEntriesOk returns a tuple with the PlayerEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerEntries

`func (o *Replay) SetPlayerEntries(v []PlayerEntry)`

SetPlayerEntries sets PlayerEntries field to given value.

### HasPlayerEntries

`func (o *Replay) HasPlayerEntries() bool`

HasPlayerEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


