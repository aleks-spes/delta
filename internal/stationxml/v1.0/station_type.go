package stationxml

/*
 *  WARNING: CODE GENERATED AUTOMATICALLY.
 *
 *  Use "go generate" to update these files.
 *
 *  THIS FILE SHOULD NOT BE EDITED BY HAND.
 *
 */

type StationType struct {
	BaseNodeType

	Latitude LatitudeType `xml:"Latitude"`

	Longitude LongitudeType `xml:"Longitude"`

	Elevation DistanceType `xml:"Elevation"`

	Site SiteType `xml:"Site"`

	Vault string `xml:"Vault,omitempty"`

	Geology string `xml:"Geology,omitempty"`

	Equipment []EquipmentType `xml:"Equipment,omitempty"`

	Operator []Operator `xml:"Operator,omitempty"`

	CreationDate DateTime `xml:"CreationDate"`

	TerminationDate DateTime `xml:"TerminationDate,omitempty"`

	TotalNumberChannels *CounterType `xml:"TotalNumberChannels,omitempty"`

	SelectedNumberChannels *CounterType `xml:"SelectedNumberChannels,omitempty"`

	ExternalReference []ExternalReferenceType `xml:"ExternalReference,omitempty"`

	Channel []ChannelType `xml:"Channel,omitempty"`
}
