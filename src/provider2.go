package main

import "fmt"

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"example_server": resourceServer(),
		},
	}
}
