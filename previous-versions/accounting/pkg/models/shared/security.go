// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	AuthHeader string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}