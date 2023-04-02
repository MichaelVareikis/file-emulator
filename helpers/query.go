package helper

var Query string = `
{
	"criteria": {
		"expressions": [
			{
				"column": "type",
				"component": "OperatingSystem",
				"operator": "is",
				"qualifier": "OSX"
			},
			{
				"column": "filewave_client_name",
				"component": "Client",
				"operator": "is_not",
				"qualifier": null
			}
		],
		"logic": "all"
	},
	"favorite": true,
	"fields": [
		{
			"column": "serial_number",
			"component": "Client"
		},
		{
			"column": "filewave_client_name",
			"component": "Client"
		},
		{
			"column": "last_check_in",
			"component": "Client"
		},
		{
			"column": "name",
			"component": "OperatingSystem"
		},
		{
			"column": "version",
			"component": "OperatingSystem"
		},
		{
			"column": "device_product_name",
			"component": "Client"
		},
		{
			"column": "ethernet_mac_address",
			"component": "CustomFields"
		},
		{
			"column": "wireless_mac_address",
			"component": "CustomFields"
		}
	],
	"group": 2,
	"main_component": "OperatingSystem",
	"name": "Axonius file",
	"id": 151,
	"version": 3
}`
