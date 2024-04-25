package qbo

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type qboConfig struct {
	ClientId *string `cty:"clientId"`
    ClientSecret *string `cty:"clientSecret"`
    RealmId *string `cty:"realmId"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"clientId": {
		Type: schema.TypeString,
	},
	"clientSecret": {
		Type: schema.TypeString,
	},
	"realmId": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &qboConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) qboConfig {
	if connection == nil || connection.Config == nil {
		return qboConfig{}
	}
	config, _ := connection.Config.(qboConfig)
	return config
}