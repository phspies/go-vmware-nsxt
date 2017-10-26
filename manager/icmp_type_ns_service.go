/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type IcmpTypeNsServiceEntry struct {

	ResourceType string `json:"resource_type"`

	// ICMP message code
	IcmpCode int64 `json:"icmp_code,omitempty"`

	// ICMP message type
	IcmpType int64 `json:"icmp_type,omitempty"`

	Protocol string `json:"protocol"`
}

type IcmpTypeNsService struct {

        NsService

        NsserviceElement IcmpTypeNsServiceEntry `json:"nsservice_element"`
}
