package dns

import (
	"fmt"

	"github.com/casbin/lego/v4/challenge"
	"github.com/casbin/lego/v4/providers/dns/alidns"
	"github.com/casbin/lego/v4/providers/dns/godaddy"
)

// NewDNSChallengeProviderByName Factory for DNS providers.
func NewDNSChallengeProviderByName(name string) (challenge.Provider, error) {
	switch name {
	case "alidns":
		return alidns.NewDNSProvider(&alidns.Config{})
	case "godaddy":
		return godaddy.NewDNSProvider(&godaddy.Config{})
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
