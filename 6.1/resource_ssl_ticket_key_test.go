// Copyright (C) 2018-2022, Pulse Secure, LLC.
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package main

/*
 * This test covers the following cases:
 *   - Creation and deletion of a vtm_ssl_ticket_key object with minimal configuration
 */

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	vtm "github.com/pulse-vadc/go-vtm/6.1"
)

func TestResourceSslTicketKey(t *testing.T) {
	objName := acctest.RandomWithPrefix("TestSslTicketKey")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslTicketKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: getBasicSslTicketKeyConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslTicketKeyExists,
				),
			},
		},
	})
}

func testAccCheckSslTicketKeyExists(s *terraform.State) error {
	for _, tfResource := range s.RootModule().Resources {
		if tfResource.Type != "vtm_ssl_ticket_key" {
			continue
		}
		objectName := tfResource.Primary.Attributes["name"]
		tm := testAccProvider.Meta().(*vtm.VirtualTrafficManager)
		if _, err := tm.GetSslTicketKey(objectName); err != nil {
			return fmt.Errorf("SslTicketKey %s does not exist: %#v", objectName, err)
		}
	}

	return nil
}

func testAccCheckSslTicketKeyDestroy(s *terraform.State) error {
	for _, tfResource := range s.RootModule().Resources {
		if tfResource.Type != "vtm_ssl_ticket_key" {
			continue
		}
		objectName := tfResource.Primary.Attributes["name"]
		tm := testAccProvider.Meta().(*vtm.VirtualTrafficManager)
		if _, err := tm.GetSslTicketKey(objectName); err == nil {
			return fmt.Errorf("SslTicketKey %s still exists", objectName)
		}
	}

	return nil
}

func getBasicSslTicketKeyConfig(name string) string {
	return fmt.Sprintf(`
        resource "vtm_ssl_ticket_key" "test_vtm_ssl_ticket_key" {
			name = "%s"
			identifier = "01234567890123456789012345678901"
			key = "abc123"
			validity_end = 10
			validity_start = 10

        }`,
		name,
	)
}
