// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresqlfirewallrule

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
)

type SqlFirewallRuleManager interface {
	CreateOrUpdateSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string, startIP string, endIP string) (result bool, err error)
	DeleteSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (err error)
	GetSQLFirewallRule(ctx context.Context, resourceGroupName string, serverName string, ruleName string) (result sql.FirewallRule, err error)
	GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
}