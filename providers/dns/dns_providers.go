package dns

import (
	"fmt"

	"github.com/casbin/lego/v4/challenge"
	"github.com/casbin/lego/v4/challenge/dns01"
	"github.com/casbin/lego/v4/providers/dns/acmedns"
	"github.com/casbin/lego/v4/providers/dns/alidns"
	"github.com/casbin/lego/v4/providers/dns/allinkl"
	"github.com/casbin/lego/v4/providers/dns/arvancloud"
	"github.com/casbin/lego/v4/providers/dns/auroradns"
	"github.com/casbin/lego/v4/providers/dns/autodns"
	"github.com/casbin/lego/v4/providers/dns/azure"
	"github.com/casbin/lego/v4/providers/dns/bindman"
	"github.com/casbin/lego/v4/providers/dns/bluecat"
	"github.com/casbin/lego/v4/providers/dns/checkdomain"
	"github.com/casbin/lego/v4/providers/dns/clouddns"
	"github.com/casbin/lego/v4/providers/dns/cloudflare"
	"github.com/casbin/lego/v4/providers/dns/cloudns"
	"github.com/casbin/lego/v4/providers/dns/cloudxns"
	"github.com/casbin/lego/v4/providers/dns/conoha"
	"github.com/casbin/lego/v4/providers/dns/constellix"
	"github.com/casbin/lego/v4/providers/dns/desec"
	"github.com/casbin/lego/v4/providers/dns/designate"
	"github.com/casbin/lego/v4/providers/dns/digitalocean"
	"github.com/casbin/lego/v4/providers/dns/dnsimple"
	"github.com/casbin/lego/v4/providers/dns/dnsmadeeasy"
	"github.com/casbin/lego/v4/providers/dns/dnspod"
	"github.com/casbin/lego/v4/providers/dns/dode"
	"github.com/casbin/lego/v4/providers/dns/domeneshop"
	"github.com/casbin/lego/v4/providers/dns/dreamhost"
	"github.com/casbin/lego/v4/providers/dns/duckdns"
	"github.com/casbin/lego/v4/providers/dns/dyn"
	"github.com/casbin/lego/v4/providers/dns/dynu"
	"github.com/casbin/lego/v4/providers/dns/easydns"
	"github.com/casbin/lego/v4/providers/dns/edgedns"
	"github.com/casbin/lego/v4/providers/dns/epik"
	"github.com/casbin/lego/v4/providers/dns/exec"
	"github.com/casbin/lego/v4/providers/dns/exoscale"
	"github.com/casbin/lego/v4/providers/dns/freemyip"
	"github.com/casbin/lego/v4/providers/dns/gandi"
	"github.com/casbin/lego/v4/providers/dns/gandiv5"
	"github.com/casbin/lego/v4/providers/dns/gcloud"
	"github.com/casbin/lego/v4/providers/dns/gcore"
	"github.com/casbin/lego/v4/providers/dns/glesys"
	"github.com/casbin/lego/v4/providers/dns/godaddy"
	"github.com/casbin/lego/v4/providers/dns/hetzner"
	"github.com/casbin/lego/v4/providers/dns/hostingde"
	"github.com/casbin/lego/v4/providers/dns/hosttech"
	"github.com/casbin/lego/v4/providers/dns/httpreq"
	"github.com/casbin/lego/v4/providers/dns/hurricane"
	"github.com/casbin/lego/v4/providers/dns/hyperone"
	"github.com/casbin/lego/v4/providers/dns/ibmcloud"
	"github.com/casbin/lego/v4/providers/dns/iij"
	"github.com/casbin/lego/v4/providers/dns/infoblox"
	"github.com/casbin/lego/v4/providers/dns/infomaniak"
	"github.com/casbin/lego/v4/providers/dns/internetbs"
	"github.com/casbin/lego/v4/providers/dns/inwx"
	"github.com/casbin/lego/v4/providers/dns/ionos"
	"github.com/casbin/lego/v4/providers/dns/joker"
	"github.com/casbin/lego/v4/providers/dns/lightsail"
	"github.com/casbin/lego/v4/providers/dns/linode"
	"github.com/casbin/lego/v4/providers/dns/liquidweb"
	"github.com/casbin/lego/v4/providers/dns/loopia"
	"github.com/casbin/lego/v4/providers/dns/luadns"
	"github.com/casbin/lego/v4/providers/dns/mydnsjp"
	"github.com/casbin/lego/v4/providers/dns/mythicbeasts"
	"github.com/casbin/lego/v4/providers/dns/namecheap"
	"github.com/casbin/lego/v4/providers/dns/namedotcom"
	"github.com/casbin/lego/v4/providers/dns/namesilo"
	"github.com/casbin/lego/v4/providers/dns/netcup"
	"github.com/casbin/lego/v4/providers/dns/netlify"
	"github.com/casbin/lego/v4/providers/dns/nicmanager"
	"github.com/casbin/lego/v4/providers/dns/nifcloud"
	"github.com/casbin/lego/v4/providers/dns/njalla"
	"github.com/casbin/lego/v4/providers/dns/ns1"
	"github.com/casbin/lego/v4/providers/dns/oraclecloud"
	"github.com/casbin/lego/v4/providers/dns/otc"
	"github.com/casbin/lego/v4/providers/dns/ovh"
	"github.com/casbin/lego/v4/providers/dns/pdns"
	"github.com/casbin/lego/v4/providers/dns/porkbun"
	"github.com/casbin/lego/v4/providers/dns/rackspace"
	"github.com/casbin/lego/v4/providers/dns/regru"
	"github.com/casbin/lego/v4/providers/dns/rfc2136"
	"github.com/casbin/lego/v4/providers/dns/rimuhosting"
	"github.com/casbin/lego/v4/providers/dns/route53"
	"github.com/casbin/lego/v4/providers/dns/sakuracloud"
	"github.com/casbin/lego/v4/providers/dns/scaleway"
	"github.com/casbin/lego/v4/providers/dns/selectel"
	"github.com/casbin/lego/v4/providers/dns/servercow"
	"github.com/casbin/lego/v4/providers/dns/simply"
	"github.com/casbin/lego/v4/providers/dns/sonic"
	"github.com/casbin/lego/v4/providers/dns/stackpath"
	"github.com/casbin/lego/v4/providers/dns/transip"
	"github.com/casbin/lego/v4/providers/dns/vegadns"
	"github.com/casbin/lego/v4/providers/dns/versio"
	"github.com/casbin/lego/v4/providers/dns/vinyldns"
	"github.com/casbin/lego/v4/providers/dns/vscale"
	"github.com/casbin/lego/v4/providers/dns/vultr"
	"github.com/casbin/lego/v4/providers/dns/wedos"
	"github.com/casbin/lego/v4/providers/dns/yandex"
	"github.com/casbin/lego/v4/providers/dns/zoneee"
	"github.com/casbin/lego/v4/providers/dns/zonomi"
)

// NewDNSChallengeProviderByName Factory for DNS providers.
func NewDNSChallengeProviderByName(name string) (challenge.Provider, error) {
	switch name {
	case "acme-dns":
		return acmedns.NewDNSProvider()
	case "alidns":
		return alidns.NewDNSProvider(&alidns.Config{})
	case "allinkl":
		return allinkl.NewDNSProvider()
	case "arvancloud":
		return arvancloud.NewDNSProvider()
	case "azure":
		return azure.NewDNSProvider()
	case "auroradns":
		return auroradns.NewDNSProvider()
	case "autodns":
		return autodns.NewDNSProvider()
	case "bindman":
		return bindman.NewDNSProvider()
	case "bluecat":
		return bluecat.NewDNSProvider()
	case "checkdomain":
		return checkdomain.NewDNSProvider()
	case "clouddns":
		return clouddns.NewDNSProvider()
	case "cloudflare":
		return cloudflare.NewDNSProvider()
	case "cloudns":
		return cloudns.NewDNSProvider()
	case "cloudxns":
		return cloudxns.NewDNSProvider()
	case "conoha":
		return conoha.NewDNSProvider()
	case "constellix":
		return constellix.NewDNSProvider()
	case "desec":
		return desec.NewDNSProvider()
	case "designate":
		return designate.NewDNSProvider()
	case "digitalocean":
		return digitalocean.NewDNSProvider()
	case "dnsimple":
		return dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		return dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		return dnspod.NewDNSProvider()
	case "dode":
		return dode.NewDNSProvider()
	case "domeneshop", "domainnameshop":
		return domeneshop.NewDNSProvider()
	case "dreamhost":
		return dreamhost.NewDNSProvider()
	case "duckdns":
		return duckdns.NewDNSProvider()
	case "dyn":
		return dyn.NewDNSProvider()
	case "dynu":
		return dynu.NewDNSProvider()
	case "easydns":
		return easydns.NewDNSProvider()
	case "edgedns", "fastdns": // "fastdns" is for compatibility with v3, must be dropped in v5
		return edgedns.NewDNSProvider()
	case "epik":
		return epik.NewDNSProvider()
	case "exec":
		return exec.NewDNSProvider()
	case "exoscale":
		return exoscale.NewDNSProvider()
	case "freemyip":
		return freemyip.NewDNSProvider()
	case "gandi":
		return gandi.NewDNSProvider()
	case "gandiv5":
		return gandiv5.NewDNSProvider()
	case "gcloud":
		return gcloud.NewDNSProvider()
	case "gcore":
		return gcore.NewDNSProvider()
	case "glesys":
		return glesys.NewDNSProvider()
	case "godaddy":
		return godaddy.NewDNSProvider()
	case "hetzner":
		return hetzner.NewDNSProvider()
	case "hostingde":
		return hostingde.NewDNSProvider()
	case "hosttech":
		return hosttech.NewDNSProvider()
	case "httpreq":
		return httpreq.NewDNSProvider()
	case "hurricane":
		return hurricane.NewDNSProvider()
	case "hyperone":
		return hyperone.NewDNSProvider()
	case "ibmcloud":
		return ibmcloud.NewDNSProvider()
	case "iij":
		return iij.NewDNSProvider()
	case "infoblox":
		return infoblox.NewDNSProvider()
	case "infomaniak":
		return infomaniak.NewDNSProvider()
	case "internetbs":
		return internetbs.NewDNSProvider()
	case "inwx":
		return inwx.NewDNSProvider()
	case "ionos":
		return ionos.NewDNSProvider()
	case "joker":
		return joker.NewDNSProvider()
	case "lightsail":
		return lightsail.NewDNSProvider()
	case "linode", "linodev4": // "linodev4" is for compatibility with v3, must be dropped in v5
		return linode.NewDNSProvider()
	case "liquidweb":
		return liquidweb.NewDNSProvider()
	case "luadns":
		return luadns.NewDNSProvider()
	case "loopia":
		return loopia.NewDNSProvider()
	case "manual":
		return dns01.NewDNSProviderManual()
	case "mydnsjp":
		return mydnsjp.NewDNSProvider()
	case "mythicbeasts":
		return mythicbeasts.NewDNSProvider()
	case "namecheap":
		return namecheap.NewDNSProvider()
	case "namedotcom":
		return namedotcom.NewDNSProvider()
	case "namesilo":
		return namesilo.NewDNSProvider()
	case "netcup":
		return netcup.NewDNSProvider()
	case "netlify":
		return netlify.NewDNSProvider()
	case "nicmanager":
		return nicmanager.NewDNSProvider()
	case "nifcloud":
		return nifcloud.NewDNSProvider()
	case "njalla":
		return njalla.NewDNSProvider()
	case "ns1":
		return ns1.NewDNSProvider()
	case "oraclecloud":
		return oraclecloud.NewDNSProvider()
	case "otc":
		return otc.NewDNSProvider()
	case "ovh":
		return ovh.NewDNSProvider()
	case "pdns":
		return pdns.NewDNSProvider()
	case "porkbun":
		return porkbun.NewDNSProvider()
	case "rackspace":
		return rackspace.NewDNSProvider()
	case "regru":
		return regru.NewDNSProvider()
	case "rfc2136":
		return rfc2136.NewDNSProvider()
	case "rimuhosting":
		return rimuhosting.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	case "sakuracloud":
		return sakuracloud.NewDNSProvider()
	case "scaleway":
		return scaleway.NewDNSProvider()
	case "selectel":
		return selectel.NewDNSProvider()
	case "servercow":
		return servercow.NewDNSProvider()
	case "simply":
		return simply.NewDNSProvider()
	case "sonic":
		return sonic.NewDNSProvider()
	case "stackpath":
		return stackpath.NewDNSProvider()
	case "transip":
		return transip.NewDNSProvider()
	case "vegadns":
		return vegadns.NewDNSProvider()
	case "versio":
		return versio.NewDNSProvider()
	case "vinyldns":
		return vinyldns.NewDNSProvider()
	case "vultr":
		return vultr.NewDNSProvider()
	case "vscale":
		return vscale.NewDNSProvider()
	case "wedos":
		return wedos.NewDNSProvider()
	case "yandex":
		return yandex.NewDNSProvider()
	case "zoneee":
		return zoneee.NewDNSProvider()
	case "zonomi":
		return zonomi.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
