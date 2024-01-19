// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-connector-sdk/tree/main/cmd/paramgen

package source

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (Config) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"caCert": {
			Default:     "",
			Description: "caCert is the Kafka broker's certificate.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"clientCert": {
			Default:     "",
			Description: "clientCert is the Kafka client's certificate.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"clientID": {
			Default:     "conduit-connector-kafka",
			Description: "clientID is a unique identifier for client connections established by this connector.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"clientKey": {
			Default:     "",
			Description: "clientKey is the Kafka client's private key.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"groupID": {
			Default:     "",
			Description: "groupID defines the consumer group id.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"insecureSkipVerify": {
			Default:     "",
			Description: "insecureSkipVerify defines whether to validate the broker's certificate chain and host name. If 'true', accepts any certificate presented by the server and any host name in that certificate.",
			Type:        sdk.ParameterTypeBool,
			Validations: []sdk.Validation{},
		},
		"readFromBeginning": {
			Default:     "",
			Description: "readFromBeginning determines from whence the consumer group should begin consuming when it finds a partition without a committed offset. If this options is set to true it will start with the first message in that partition.",
			Type:        sdk.ParameterTypeBool,
			Validations: []sdk.Validation{},
		},
		"saslMechanism": {
			Default:     "",
			Description: "saslMechanism configures the connector to use SASL authentication. If empty, no authentication will be performed.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"PLAIN", "SCRAM-SHA-256", "SCRAM-SHA-512"}},
			},
		},
		"saslPassword": {
			Default:     "",
			Description: "saslPassword sets up the password used with SASL authentication.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"saslUsername": {
			Default:     "",
			Description: "saslUsername sets up the username used with SASL authentication.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"servers": {
			Default:     "",
			Description: "servers is a list of Kafka bootstrap servers, which will be used to discover all the servers in a cluster.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"tls.enabled": {
			Default:     "",
			Description: "tls.enabled defines whether TLS is needed to communicate with the Kafka cluster.",
			Type:        sdk.ParameterTypeBool,
			Validations: []sdk.Validation{},
		},
		"topic": {
			Default:     "",
			Description: "topic is the Kafka topic.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
	}
}
