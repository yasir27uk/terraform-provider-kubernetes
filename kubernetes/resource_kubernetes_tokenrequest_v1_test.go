// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetes

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	api "k8s.io/api/core/v1"
)

func TestAccKubernetesTokenRequestV1_basic(t *testing.T) {
	var conf api.ServiceAccount
	resourceName := "kubernetes_service_account_v1.tokentest"
	tokenName := "kubernetes_token_request_v1.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		IDRefreshName:     resourceName,
		IDRefreshIgnore:   []string{"metadata.0.resource_version"},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckKubernetesServiceAccountV1Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKubernetesTokenRequestV1Config_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckKubernetesServiceAccountV1Exists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "metadata.0.name", "tokentest"),
					resource.TestCheckResourceAttr(tokenName, "metadata.0.name", "tokentest"),
					resource.TestCheckResourceAttr(tokenName, "spec.0.audiences.0", "api"),
					resource.TestCheckResourceAttr(tokenName, "spec.0.audiences.1", "vault"),
					resource.TestCheckResourceAttr(tokenName, "spec.0.audiences.2", "factors"),
					resource.TestCheckResourceAttrSet(tokenName, "token"),
				),
			},
		},
	})
}

func testAccKubernetesTokenRequestV1Config_basic() string {
	return fmt.Sprintf(`resource "kubernetes_service_account_v1" "tokentest" {
  metadata {
    name = "tokentest"
  }
}

resource "kubernetes_token_request_v1" "test" {
  metadata {
    name = kubernetes_service_account_v1.tokentest.metadata.0.name
  }
  spec {
    audiences = [
      "api",
      "vault",
      "factors"
    ]
  }
}


`)
}
