package network

import (
	"net"
	"net/netip"

	"github.com/sagernet/sing/common"
	M "github.com/sagernet/sing/common/metadata"
)

func LocalAddrs() ([]netip.Addr, error) {
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	return common.Map(interfaceAddrs, func(addr net.Addr) netip.Addr {
		return M.AddrFromNetAddr(addr)
	}), nil
}

func IsPublicAddr(addr netip.Addr) bool {
	return !(addr.IsPrivate() || addr.IsLoopback() || addr.IsMulticast() || addr.IsGlobalUnicast() || addr.IsLinkLocalUnicast() || addr.IsInterfaceLocalMulticast())
}

func LocalPublicAddrs() ([]netip.Addr, error) {
	publicAddrs, err := LocalAddrs()
	if err != nil {
		return nil, err
	}
	return common.Filter(publicAddrs, func(addr netip.Addr) bool {
		return IsPublicAddr(addr)
	}), nil
}
