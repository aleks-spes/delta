package stationxml

/*
 *  WARNING: CODE GENERATED AUTOMATICALLY.
 *
 *  Use "go generate" to update these files.
 *
 *  THIS FILE SHOULD NOT BE EDITED BY HAND.
 *
 */

type CoefficientsType struct {
	BaseFilterType

	CfTransferFunctionType CfTransferFunctionType `xml:"CfTransferFunctionType"`

	Numerator []Numerator `xml:"Numerator,omitempty"`

	Denominator []Denominator `xml:"Denominator,omitempty"`
}
