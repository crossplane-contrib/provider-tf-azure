/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1active "github.com/ulucinar/provider-tf-azure/apis/active/v1alpha1"
	v1alpha1advanced "github.com/ulucinar/provider-tf-azure/apis/advanced/v1alpha1"
	v1alpha1analysis "github.com/ulucinar/provider-tf-azure/apis/analysis/v1alpha1"
	v1alpha1api "github.com/ulucinar/provider-tf-azure/apis/api/v1alpha1"
	v1alpha1app "github.com/ulucinar/provider-tf-azure/apis/app/v1alpha1"
	v1alpha1application "github.com/ulucinar/provider-tf-azure/apis/application/v1alpha1"
	v1alpha1attestation "github.com/ulucinar/provider-tf-azure/apis/attestation/v1alpha1"
	v1alpha1automation "github.com/ulucinar/provider-tf-azure/apis/automation/v1alpha1"
	v1alpha1availability "github.com/ulucinar/provider-tf-azure/apis/availability/v1alpha1"
	v1alpha1backup "github.com/ulucinar/provider-tf-azure/apis/backup/v1alpha1"
	v1alpha1bastion "github.com/ulucinar/provider-tf-azure/apis/bastion/v1alpha1"
	v1alpha1batch "github.com/ulucinar/provider-tf-azure/apis/batch/v1alpha1"
	v1alpha1blueprint "github.com/ulucinar/provider-tf-azure/apis/blueprint/v1alpha1"
	v1alpha1bot "github.com/ulucinar/provider-tf-azure/apis/bot/v1alpha1"
	v1alpha1cdn "github.com/ulucinar/provider-tf-azure/apis/cdn/v1alpha1"
	v1alpha1cognitive "github.com/ulucinar/provider-tf-azure/apis/cognitive/v1alpha1"
	v1alpha1communication "github.com/ulucinar/provider-tf-azure/apis/communication/v1alpha1"
	v1alpha1consumption "github.com/ulucinar/provider-tf-azure/apis/consumption/v1alpha1"
	v1alpha1container "github.com/ulucinar/provider-tf-azure/apis/container/v1alpha1"
	v1alpha1cosmosdb "github.com/ulucinar/provider-tf-azure/apis/cosmosdb/v1alpha1"
	v1alpha1cost "github.com/ulucinar/provider-tf-azure/apis/cost/v1alpha1"
	v1alpha1custom "github.com/ulucinar/provider-tf-azure/apis/custom/v1alpha1"
	v1alpha1dashboard "github.com/ulucinar/provider-tf-azure/apis/dashboard/v1alpha1"
	v1alpha1data "github.com/ulucinar/provider-tf-azure/apis/data/v1alpha1"
	v1alpha1database "github.com/ulucinar/provider-tf-azure/apis/database/v1alpha1"
	v1alpha1databox "github.com/ulucinar/provider-tf-azure/apis/databox/v1alpha1"
	v1alpha1databricks "github.com/ulucinar/provider-tf-azure/apis/databricks/v1alpha1"
	v1alpha1dedicated "github.com/ulucinar/provider-tf-azure/apis/dedicated/v1alpha1"
	v1alpha1dev "github.com/ulucinar/provider-tf-azure/apis/dev/v1alpha1"
	v1alpha1devspace "github.com/ulucinar/provider-tf-azure/apis/devspace/v1alpha1"
	v1alpha1digital "github.com/ulucinar/provider-tf-azure/apis/digital/v1alpha1"
	v1alpha1disk "github.com/ulucinar/provider-tf-azure/apis/disk/v1alpha1"
	v1alpha1dns "github.com/ulucinar/provider-tf-azure/apis/dns/v1alpha1"
	v1alpha1eventgrid "github.com/ulucinar/provider-tf-azure/apis/eventgrid/v1alpha1"
	v1alpha1eventhub "github.com/ulucinar/provider-tf-azure/apis/eventhub/v1alpha1"
	v1alpha1express "github.com/ulucinar/provider-tf-azure/apis/express/v1alpha1"
	v1alpha1firewall "github.com/ulucinar/provider-tf-azure/apis/firewall/v1alpha1"
	v1alpha1frontdoor "github.com/ulucinar/provider-tf-azure/apis/frontdoor/v1alpha1"
	v1alpha1function "github.com/ulucinar/provider-tf-azure/apis/function/v1alpha1"
	v1alpha1hdinsight "github.com/ulucinar/provider-tf-azure/apis/hdinsight/v1alpha1"
	v1alpha1healthbot "github.com/ulucinar/provider-tf-azure/apis/healthbot/v1alpha1"
	v1alpha1healthcare "github.com/ulucinar/provider-tf-azure/apis/healthcare/v1alpha1"
	v1alpha1hpc "github.com/ulucinar/provider-tf-azure/apis/hpc/v1alpha1"
	v1alpha1image "github.com/ulucinar/provider-tf-azure/apis/image/v1alpha1"
	v1alpha1integration "github.com/ulucinar/provider-tf-azure/apis/integration/v1alpha1"
	v1alpha1iot "github.com/ulucinar/provider-tf-azure/apis/iot/v1alpha1"
	v1alpha1iotcentral "github.com/ulucinar/provider-tf-azure/apis/iotcentral/v1alpha1"
	v1alpha1iothub "github.com/ulucinar/provider-tf-azure/apis/iothub/v1alpha1"
	v1alpha1ip "github.com/ulucinar/provider-tf-azure/apis/ip/v1alpha1"
	v1alpha1key "github.com/ulucinar/provider-tf-azure/apis/key/v1alpha1"
	v1alpha1kubernetes "github.com/ulucinar/provider-tf-azure/apis/kubernetes/v1alpha1"
	v1alpha1kusto "github.com/ulucinar/provider-tf-azure/apis/kusto/v1alpha1"
	v1alpha1lb "github.com/ulucinar/provider-tf-azure/apis/lb/v1alpha1"
	v1alpha1lighthouse "github.com/ulucinar/provider-tf-azure/apis/lighthouse/v1alpha1"
	v1alpha1linux "github.com/ulucinar/provider-tf-azure/apis/linux/v1alpha1"
	v1alpha1local "github.com/ulucinar/provider-tf-azure/apis/local/v1alpha1"
	v1alpha1log "github.com/ulucinar/provider-tf-azure/apis/log/v1alpha1"
	v1alpha1logic "github.com/ulucinar/provider-tf-azure/apis/logic/v1alpha1"
	v1alpha1machine "github.com/ulucinar/provider-tf-azure/apis/machine/v1alpha1"
	v1alpha1maintenance "github.com/ulucinar/provider-tf-azure/apis/maintenance/v1alpha1"
	v1alpha1managed "github.com/ulucinar/provider-tf-azure/apis/managed/v1alpha1"
	v1alpha1management "github.com/ulucinar/provider-tf-azure/apis/management/v1alpha1"
	v1alpha1maps "github.com/ulucinar/provider-tf-azure/apis/maps/v1alpha1"
	v1alpha1mariadb "github.com/ulucinar/provider-tf-azure/apis/mariadb/v1alpha1"
	v1alpha1marketplace "github.com/ulucinar/provider-tf-azure/apis/marketplace/v1alpha1"
	v1alpha1media "github.com/ulucinar/provider-tf-azure/apis/media/v1alpha1"
	v1alpha1monitor "github.com/ulucinar/provider-tf-azure/apis/monitor/v1alpha1"
	v1alpha1mssql "github.com/ulucinar/provider-tf-azure/apis/mssql/v1alpha1"
	v1alpha1mysql "github.com/ulucinar/provider-tf-azure/apis/mysql/v1alpha1"
	v1alpha1nat "github.com/ulucinar/provider-tf-azure/apis/nat/v1alpha1"
	v1alpha1netapp "github.com/ulucinar/provider-tf-azure/apis/netapp/v1alpha1"
	v1alpha1network "github.com/ulucinar/provider-tf-azure/apis/network/v1alpha1"
	v1alpha1notification "github.com/ulucinar/provider-tf-azure/apis/notification/v1alpha1"
	v1alpha1orchestrated "github.com/ulucinar/provider-tf-azure/apis/orchestrated/v1alpha1"
	v1alpha1packet "github.com/ulucinar/provider-tf-azure/apis/packet/v1alpha1"
	v1alpha1point "github.com/ulucinar/provider-tf-azure/apis/point/v1alpha1"
	v1alpha1policy "github.com/ulucinar/provider-tf-azure/apis/policy/v1alpha1"
	v1alpha1portal "github.com/ulucinar/provider-tf-azure/apis/portal/v1alpha1"
	v1alpha1postgresql "github.com/ulucinar/provider-tf-azure/apis/postgresql/v1alpha1"
	v1alpha1powerbi "github.com/ulucinar/provider-tf-azure/apis/powerbi/v1alpha1"
	v1alpha1private "github.com/ulucinar/provider-tf-azure/apis/private/v1alpha1"
	v1alpha1proximity "github.com/ulucinar/provider-tf-azure/apis/proximity/v1alpha1"
	v1alpha1public "github.com/ulucinar/provider-tf-azure/apis/public/v1alpha1"
	v1alpha1purview "github.com/ulucinar/provider-tf-azure/apis/purview/v1alpha1"
	v1alpha1recovery "github.com/ulucinar/provider-tf-azure/apis/recovery/v1alpha1"
	v1alpha1redis "github.com/ulucinar/provider-tf-azure/apis/redis/v1alpha1"
	v1alpha1relay "github.com/ulucinar/provider-tf-azure/apis/relay/v1alpha1"
	v1alpha1resource "github.com/ulucinar/provider-tf-azure/apis/resource/v1alpha1"
	v1alpha1role "github.com/ulucinar/provider-tf-azure/apis/role/v1alpha1"
	v1alpha1route "github.com/ulucinar/provider-tf-azure/apis/route/v1alpha1"
	v1alpha1search "github.com/ulucinar/provider-tf-azure/apis/search/v1alpha1"
	v1alpha1security "github.com/ulucinar/provider-tf-azure/apis/security/v1alpha1"
	v1alpha1sentinel "github.com/ulucinar/provider-tf-azure/apis/sentinel/v1alpha1"
	v1alpha1service "github.com/ulucinar/provider-tf-azure/apis/service/v1alpha1"
	v1alpha1servicebus "github.com/ulucinar/provider-tf-azure/apis/servicebus/v1alpha1"
	v1alpha1shared "github.com/ulucinar/provider-tf-azure/apis/shared/v1alpha1"
	v1alpha1signalr "github.com/ulucinar/provider-tf-azure/apis/signalr/v1alpha1"
	v1alpha1site "github.com/ulucinar/provider-tf-azure/apis/site/v1alpha1"
	v1alpha1snapshot "github.com/ulucinar/provider-tf-azure/apis/snapshot/v1alpha1"
	v1alpha1spatial "github.com/ulucinar/provider-tf-azure/apis/spatial/v1alpha1"
	v1alpha1spring "github.com/ulucinar/provider-tf-azure/apis/spring/v1alpha1"
	v1alpha1sql "github.com/ulucinar/provider-tf-azure/apis/sql/v1alpha1"
	v1alpha1ssh "github.com/ulucinar/provider-tf-azure/apis/ssh/v1alpha1"
	v1alpha1stack "github.com/ulucinar/provider-tf-azure/apis/stack/v1alpha1"
	v1alpha1static "github.com/ulucinar/provider-tf-azure/apis/static/v1alpha1"
	v1alpha1storage "github.com/ulucinar/provider-tf-azure/apis/storage/v1alpha1"
	v1alpha1stream "github.com/ulucinar/provider-tf-azure/apis/stream/v1alpha1"
	v1alpha1subnet "github.com/ulucinar/provider-tf-azure/apis/subnet/v1alpha1"
	v1alpha1subscription "github.com/ulucinar/provider-tf-azure/apis/subscription/v1alpha1"
	v1alpha1synapse "github.com/ulucinar/provider-tf-azure/apis/synapse/v1alpha1"
	v1alpha1template "github.com/ulucinar/provider-tf-azure/apis/template/v1alpha1"
	v1alpha1tenant "github.com/ulucinar/provider-tf-azure/apis/tenant/v1alpha1"
	v1alpha1traffic "github.com/ulucinar/provider-tf-azure/apis/traffic/v1alpha1"
	v1alpha1user "github.com/ulucinar/provider-tf-azure/apis/user/v1alpha1"
	v1alpha1 "github.com/ulucinar/provider-tf-azure/apis/v1alpha1"
	v1alpha1video "github.com/ulucinar/provider-tf-azure/apis/video/v1alpha1"
	v1alpha1virtual "github.com/ulucinar/provider-tf-azure/apis/virtual/v1alpha1"
	v1alpha1vmware "github.com/ulucinar/provider-tf-azure/apis/vmware/v1alpha1"
	v1alpha1vpn "github.com/ulucinar/provider-tf-azure/apis/vpn/v1alpha1"
	v1alpha1web "github.com/ulucinar/provider-tf-azure/apis/web/v1alpha1"
	v1alpha1windows "github.com/ulucinar/provider-tf-azure/apis/windows/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1active.SchemeBuilder.AddToScheme,
		v1alpha1advanced.SchemeBuilder.AddToScheme,
		v1alpha1analysis.SchemeBuilder.AddToScheme,
		v1alpha1api.SchemeBuilder.AddToScheme,
		v1alpha1app.SchemeBuilder.AddToScheme,
		v1alpha1application.SchemeBuilder.AddToScheme,
		v1alpha1attestation.SchemeBuilder.AddToScheme,
		v1alpha1automation.SchemeBuilder.AddToScheme,
		v1alpha1availability.SchemeBuilder.AddToScheme,
		v1alpha1backup.SchemeBuilder.AddToScheme,
		v1alpha1bastion.SchemeBuilder.AddToScheme,
		v1alpha1batch.SchemeBuilder.AddToScheme,
		v1alpha1blueprint.SchemeBuilder.AddToScheme,
		v1alpha1bot.SchemeBuilder.AddToScheme,
		v1alpha1cdn.SchemeBuilder.AddToScheme,
		v1alpha1cognitive.SchemeBuilder.AddToScheme,
		v1alpha1communication.SchemeBuilder.AddToScheme,
		v1alpha1consumption.SchemeBuilder.AddToScheme,
		v1alpha1container.SchemeBuilder.AddToScheme,
		v1alpha1cosmosdb.SchemeBuilder.AddToScheme,
		v1alpha1cost.SchemeBuilder.AddToScheme,
		v1alpha1custom.SchemeBuilder.AddToScheme,
		v1alpha1dashboard.SchemeBuilder.AddToScheme,
		v1alpha1data.SchemeBuilder.AddToScheme,
		v1alpha1database.SchemeBuilder.AddToScheme,
		v1alpha1databox.SchemeBuilder.AddToScheme,
		v1alpha1databricks.SchemeBuilder.AddToScheme,
		v1alpha1dedicated.SchemeBuilder.AddToScheme,
		v1alpha1dev.SchemeBuilder.AddToScheme,
		v1alpha1devspace.SchemeBuilder.AddToScheme,
		v1alpha1digital.SchemeBuilder.AddToScheme,
		v1alpha1disk.SchemeBuilder.AddToScheme,
		v1alpha1dns.SchemeBuilder.AddToScheme,
		v1alpha1eventgrid.SchemeBuilder.AddToScheme,
		v1alpha1eventhub.SchemeBuilder.AddToScheme,
		v1alpha1express.SchemeBuilder.AddToScheme,
		v1alpha1firewall.SchemeBuilder.AddToScheme,
		v1alpha1frontdoor.SchemeBuilder.AddToScheme,
		v1alpha1function.SchemeBuilder.AddToScheme,
		v1alpha1hdinsight.SchemeBuilder.AddToScheme,
		v1alpha1healthbot.SchemeBuilder.AddToScheme,
		v1alpha1healthcare.SchemeBuilder.AddToScheme,
		v1alpha1hpc.SchemeBuilder.AddToScheme,
		v1alpha1image.SchemeBuilder.AddToScheme,
		v1alpha1integration.SchemeBuilder.AddToScheme,
		v1alpha1iot.SchemeBuilder.AddToScheme,
		v1alpha1iotcentral.SchemeBuilder.AddToScheme,
		v1alpha1iothub.SchemeBuilder.AddToScheme,
		v1alpha1ip.SchemeBuilder.AddToScheme,
		v1alpha1key.SchemeBuilder.AddToScheme,
		v1alpha1kubernetes.SchemeBuilder.AddToScheme,
		v1alpha1kusto.SchemeBuilder.AddToScheme,
		v1alpha1lb.SchemeBuilder.AddToScheme,
		v1alpha1lighthouse.SchemeBuilder.AddToScheme,
		v1alpha1linux.SchemeBuilder.AddToScheme,
		v1alpha1local.SchemeBuilder.AddToScheme,
		v1alpha1log.SchemeBuilder.AddToScheme,
		v1alpha1logic.SchemeBuilder.AddToScheme,
		v1alpha1machine.SchemeBuilder.AddToScheme,
		v1alpha1maintenance.SchemeBuilder.AddToScheme,
		v1alpha1managed.SchemeBuilder.AddToScheme,
		v1alpha1management.SchemeBuilder.AddToScheme,
		v1alpha1maps.SchemeBuilder.AddToScheme,
		v1alpha1mariadb.SchemeBuilder.AddToScheme,
		v1alpha1marketplace.SchemeBuilder.AddToScheme,
		v1alpha1media.SchemeBuilder.AddToScheme,
		v1alpha1monitor.SchemeBuilder.AddToScheme,
		v1alpha1mssql.SchemeBuilder.AddToScheme,
		v1alpha1mysql.SchemeBuilder.AddToScheme,
		v1alpha1nat.SchemeBuilder.AddToScheme,
		v1alpha1netapp.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1notification.SchemeBuilder.AddToScheme,
		v1alpha1orchestrated.SchemeBuilder.AddToScheme,
		v1alpha1packet.SchemeBuilder.AddToScheme,
		v1alpha1point.SchemeBuilder.AddToScheme,
		v1alpha1policy.SchemeBuilder.AddToScheme,
		v1alpha1portal.SchemeBuilder.AddToScheme,
		v1alpha1postgresql.SchemeBuilder.AddToScheme,
		v1alpha1powerbi.SchemeBuilder.AddToScheme,
		v1alpha1private.SchemeBuilder.AddToScheme,
		v1alpha1proximity.SchemeBuilder.AddToScheme,
		v1alpha1public.SchemeBuilder.AddToScheme,
		v1alpha1purview.SchemeBuilder.AddToScheme,
		v1alpha1recovery.SchemeBuilder.AddToScheme,
		v1alpha1redis.SchemeBuilder.AddToScheme,
		v1alpha1relay.SchemeBuilder.AddToScheme,
		v1alpha1resource.SchemeBuilder.AddToScheme,
		v1alpha1role.SchemeBuilder.AddToScheme,
		v1alpha1route.SchemeBuilder.AddToScheme,
		v1alpha1search.SchemeBuilder.AddToScheme,
		v1alpha1security.SchemeBuilder.AddToScheme,
		v1alpha1sentinel.SchemeBuilder.AddToScheme,
		v1alpha1service.SchemeBuilder.AddToScheme,
		v1alpha1servicebus.SchemeBuilder.AddToScheme,
		v1alpha1shared.SchemeBuilder.AddToScheme,
		v1alpha1signalr.SchemeBuilder.AddToScheme,
		v1alpha1site.SchemeBuilder.AddToScheme,
		v1alpha1snapshot.SchemeBuilder.AddToScheme,
		v1alpha1spatial.SchemeBuilder.AddToScheme,
		v1alpha1spring.SchemeBuilder.AddToScheme,
		v1alpha1sql.SchemeBuilder.AddToScheme,
		v1alpha1ssh.SchemeBuilder.AddToScheme,
		v1alpha1stack.SchemeBuilder.AddToScheme,
		v1alpha1static.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1stream.SchemeBuilder.AddToScheme,
		v1alpha1subnet.SchemeBuilder.AddToScheme,
		v1alpha1subscription.SchemeBuilder.AddToScheme,
		v1alpha1synapse.SchemeBuilder.AddToScheme,
		v1alpha1template.SchemeBuilder.AddToScheme,
		v1alpha1tenant.SchemeBuilder.AddToScheme,
		v1alpha1traffic.SchemeBuilder.AddToScheme,
		v1alpha1user.SchemeBuilder.AddToScheme,
		v1alpha1video.SchemeBuilder.AddToScheme,
		v1alpha1virtual.SchemeBuilder.AddToScheme,
		v1alpha1vmware.SchemeBuilder.AddToScheme,
		v1alpha1vpn.SchemeBuilder.AddToScheme,
		v1alpha1web.SchemeBuilder.AddToScheme,
		v1alpha1windows.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
