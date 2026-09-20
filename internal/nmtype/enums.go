package nmtype

import "encoding/json"

//go:generate go run golang.org/x/tools/cmd/stringer -type=Connectivity
type Connectivity uint32

const (
	ConnectivityUnknown Connectivity = 0
	ConnectivityNone    Connectivity = 1
	ConnectivityPortal  Connectivity = 2
	ConnectivityLimited Connectivity = 3
	ConnectivityFull    Connectivity = 4
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=State
type State uint32

const (
	StateUnknown         State = 0
	StateAsleep          State = 10
	StateDisconnected    State = 20
	StateDisconnecting   State = 30
	StateConnecting      State = 40
	StateConnectedLocal  State = 50
	StateConnectedSite   State = 60
	StateConnectedGlobal State = 70
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=CheckpointCreateFlags
type CheckpointCreateFlags uint32

const (
	CheckpointCreateFlagsNone                 CheckpointCreateFlags = 0
	CheckpointCreateFlagsDestroyAll           CheckpointCreateFlags = 0x01
	CheckpointCreateFlagsDeleteNewConnections CheckpointCreateFlags = 0x02
	CheckpointCreateFlagsDisconnectNewDevices CheckpointCreateFlags = 0x04
	CheckpointCreateFlagsAllowOverlapping     CheckpointCreateFlags = 0x08
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Capability
type Capability uint32

const (
	CapabilityTeam Capability = 1
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Metered
type Metered uint32

const (
	MeteredUnknown  Metered = 0
	MeteredYes      Metered = 1
	MeteredNo       Metered = 2
	MeteredGuessYes Metered = 3
	MeteredGuessNo  Metered = 4
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=DeviceState
type DeviceState uint32

const (
	DeviceStateUnknown      DeviceState = 0
	DeviceStateUnmanaged    DeviceState = 10
	DeviceStateUnavailable  DeviceState = 20
	DeviceStateDisconnected DeviceState = 30
	DeviceStatePrepare      DeviceState = 40
	DeviceStateConfig       DeviceState = 50
	DeviceStateNeedAuth     DeviceState = 60
	DeviceStateIPConfig     DeviceState = 70
	DeviceStateIPCheck      DeviceState = 80
	DeviceStateSecondaries  DeviceState = 90
	DeviceStateActivated    DeviceState = 100
	DeviceStateDeactivating DeviceState = 110
	DeviceStateFailed       DeviceState = 120
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=DeviceInterfaceFlags
type DeviceInterfaceFlags uint32

const (
	DeviceInterfaceFlagsNone              DeviceInterfaceFlags = 0
	DeviceInterfaceFlagsUp                DeviceInterfaceFlags = 0x1
	DeviceInterfaceFlagsLowerUp           DeviceInterfaceFlags = 0x2
	DeviceInterfaceFlagsPromisc           DeviceInterfaceFlags = 0x4
	DeviceInterfaceFlagsCarrier           DeviceInterfaceFlags = 0x10000
	DeviceInterfaceFlagsLldpClientEnabled DeviceInterfaceFlags = 0x20000
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ActiveConnectionState
type ActiveConnectionState uint32

const (
	ActiveConnectionStateUnknown      ActiveConnectionState = 0
	ActiveConnectionStateActivating   ActiveConnectionState = 1
	ActiveConnectionStateActivated    ActiveConnectionState = 2
	ActiveConnectionStateDeactivating ActiveConnectionState = 3
	ActiveConnectionStateDeactivated  ActiveConnectionState = 4
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ActivationStateFlag
type ActivationStateFlag uint32

const (
	ActivationStateFlagNone                             ActivationStateFlag = 0x00
	ActivationStateFlagIsMaster                         ActivationStateFlag = 0x01
	ActivationStateFlagIsSlave                          ActivationStateFlag = 0x02
	ActivationStateFlagLayer2Ready                      ActivationStateFlag = 0x04
	ActivationStateFlagIP4Ready                         ActivationStateFlag = 0x08
	ActivationStateFlagIP6Ready                         ActivationStateFlag = 0x10
	ActivationStateFlagMasterHasSlaves                  ActivationStateFlag = 0x20
	ActivationStateFlagLifetimeBoundToProfileVisibility ActivationStateFlag = 0x40
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=DeviceType
type DeviceType uint32

const (
	DeviceTypeUnknown      DeviceType = 0
	DeviceTypeEthernet     DeviceType = 1
	DeviceTypeWifi         DeviceType = 2
	DeviceTypeUnused1      DeviceType = 3
	DeviceTypeUnused2      DeviceType = 4
	DeviceTypeBt           DeviceType = 5
	DeviceTypeOlpcMesh     DeviceType = 6
	DeviceTypeWimax        DeviceType = 7
	DeviceTypeModem        DeviceType = 8
	DeviceTypeInfiniband   DeviceType = 9
	DeviceTypeBond         DeviceType = 10
	DeviceTypeVlan         DeviceType = 11
	DeviceTypeAdsl         DeviceType = 12
	DeviceTypeBridge       DeviceType = 13
	DeviceTypeGeneric      DeviceType = 14
	DeviceTypeTeam         DeviceType = 15
	DeviceTypeTun          DeviceType = 16
	DeviceTypeIPTunnel     DeviceType = 17
	DeviceTypeMacvlan      DeviceType = 18
	DeviceTypeVxlan        DeviceType = 19
	DeviceTypeVeth         DeviceType = 20
	DeviceTypeMacsec       DeviceType = 21
	DeviceTypeDummy        DeviceType = 22
	DeviceTypePpp          DeviceType = 23
	DeviceTypeOvsInterface DeviceType = 24
	DeviceTypeOvsPort      DeviceType = 25
	DeviceTypeOvsBridge    DeviceType = 26
	DeviceTypeWpan         DeviceType = 27
	DeviceType6lowpan      DeviceType = 28
	DeviceTypeWireguard    DeviceType = 29
	DeviceTypeWifiP2p      DeviceType = 30
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=APFlags
type APFlags uint32

const (
	APFlagsNone    APFlags = 0x0
	APFlagsPrivacy APFlags = 0x1
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=APSecurity
type APSecurity uint32

const (
	APSecurityNone         APSecurity = 0x0
	APSecurityPairWEP40    APSecurity = 0x1
	APSecurityPairWEP104   APSecurity = 0x2
	APSecurityPairTKIP     APSecurity = 0x4
	APSecurityPairCCMP     APSecurity = 0x8
	APSecurityGroupWEP40   APSecurity = 0x10
	APSecurityGroupWEP104  APSecurity = 0x20
	APSecurityGroupTKIP    APSecurity = 0x40
	APSecurityGroupCCMP    APSecurity = 0x80
	APSecurityKeyMgmtPSK   APSecurity = 0x100
	APSecurityKeyMgmt8021X APSecurity = 0x200
	APSecurityKeyMgmtSAE   APSecurity = 0x400
	APSecurityKeyMgmtOWE   APSecurity = 0x800
	APSecurityKeyMgmtOWETM APSecurity = 0x1000
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=WifiMode
type WifiMode uint32

const (
	WifiModeUnknown WifiMode = 0
	WifiModeAdhoc   WifiMode = 1
	WifiModeInfra   WifiMode = 2
	WifiModeAP      WifiMode = 3
)

// MarshalJSON encodes a WifiMode as its string representation.
func (m WifiMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=VpnConnectionState
type VpnConnectionState uint32

const (
	VpnConnectionUnknown      VpnConnectionState = 0
	VpnConnectionPrepare      VpnConnectionState = 1
	VpnConnectionNeedAuth     VpnConnectionState = 2
	VpnConnectionConnect      VpnConnectionState = 3
	VpnConnectionIPConfigGet  VpnConnectionState = 4
	VpnConnectionActivated    VpnConnectionState = 5
	VpnConnectionFailed       VpnConnectionState = 6
	VpnConnectionDisconnected VpnConnectionState = 7
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=DeviceStateReason
type DeviceStateReason uint32

const (
	DeviceStateReasonNone                        DeviceStateReason = 0
	DeviceStateReasonUnknown                     DeviceStateReason = 1
	DeviceStateReasonNowManaged                  DeviceStateReason = 2
	DeviceStateReasonNowUnmanaged                DeviceStateReason = 3
	DeviceStateReasonConfigFailed                DeviceStateReason = 4
	DeviceStateReasonIPConfigUnavailable         DeviceStateReason = 5
	DeviceStateReasonIPConfigExpired             DeviceStateReason = 6
	DeviceStateReasonNoSecrets                   DeviceStateReason = 7
	DeviceStateReasonSupplicantDisconnect        DeviceStateReason = 8
	DeviceStateReasonSupplicantConfigFailed      DeviceStateReason = 9
	DeviceStateReasonSupplicantFailed            DeviceStateReason = 10
	DeviceStateReasonSupplicantTimeout           DeviceStateReason = 11
	DeviceStateReasonPppStartFailed              DeviceStateReason = 12
	DeviceStateReasonPppDisconnect               DeviceStateReason = 13
	DeviceStateReasonPppFailed                   DeviceStateReason = 14
	DeviceStateReasonDhcpStartFailed             DeviceStateReason = 15
	DeviceStateReasonDhcpError                   DeviceStateReason = 16
	DeviceStateReasonDhcpFailed                  DeviceStateReason = 17
	DeviceStateReasonSharedStartFailed           DeviceStateReason = 18
	DeviceStateReasonSharedFailed                DeviceStateReason = 19
	DeviceStateReasonAutoipStartFailed           DeviceStateReason = 20
	DeviceStateReasonAutoipError                 DeviceStateReason = 21
	DeviceStateReasonAutoipFailed                DeviceStateReason = 22
	DeviceStateReasonModemBusy                   DeviceStateReason = 23
	DeviceStateReasonModemNoDialTone             DeviceStateReason = 24
	DeviceStateReasonModemNoCarrier              DeviceStateReason = 25
	DeviceStateReasonModemDialTimeout            DeviceStateReason = 26
	DeviceStateReasonModemDialFailed             DeviceStateReason = 27
	DeviceStateReasonModemInitFailed             DeviceStateReason = 28
	DeviceStateReasonGsmApnFailed                DeviceStateReason = 29
	DeviceStateReasonGsmRegistrationNotSearching DeviceStateReason = 30
	DeviceStateReasonGsmRegistrationDenied       DeviceStateReason = 31
	DeviceStateReasonGsmRegistrationTimeout      DeviceStateReason = 32
	DeviceStateReasonGsmRegistrationFailed       DeviceStateReason = 33
	DeviceStateReasonGsmPinCheckFailed           DeviceStateReason = 34
	DeviceStateReasonFirmwareMissing             DeviceStateReason = 35
	DeviceStateReasonRemoved                     DeviceStateReason = 36
	DeviceStateReasonSleeping                    DeviceStateReason = 37
	DeviceStateReasonConnectionRemoved           DeviceStateReason = 38
	DeviceStateReasonUserRequested               DeviceStateReason = 39
	DeviceStateReasonCarrier                     DeviceStateReason = 40
	DeviceStateReasonConnectionAssumed           DeviceStateReason = 41
	DeviceStateReasonSupplicantAvailable         DeviceStateReason = 42
	DeviceStateReasonModemNotFound               DeviceStateReason = 43
	DeviceStateReasonBtFailed                    DeviceStateReason = 44
	DeviceStateReasonGsmSimNotInserted           DeviceStateReason = 45
	DeviceStateReasonGsmSimPinRequired           DeviceStateReason = 46
	DeviceStateReasonGsmSimPukRequired           DeviceStateReason = 47
	DeviceStateReasonGsmSimWrong                 DeviceStateReason = 48
	DeviceStateReasonInfinibandMode              DeviceStateReason = 49
	DeviceStateReasonDependencyFailed            DeviceStateReason = 50
	DeviceStateReasonBr2684Failed                DeviceStateReason = 51
	DeviceStateReasonModemManagerUnavailable     DeviceStateReason = 52
	DeviceStateReasonSsidNotFound                DeviceStateReason = 53
	DeviceStateReasonSecondaryConnectionFailed   DeviceStateReason = 54
	DeviceStateReasonDcbFcoeFailed               DeviceStateReason = 55
	DeviceStateReasonTeamdControlFailed          DeviceStateReason = 56
	DeviceStateReasonModemFailed                 DeviceStateReason = 57
	DeviceStateReasonModemAvailable              DeviceStateReason = 58
	DeviceStateReasonSimPinIncorrect             DeviceStateReason = 59
	DeviceStateReasonNewActivation               DeviceStateReason = 60
	DeviceStateReasonParentChanged               DeviceStateReason = 61
	DeviceStateReasonParentManagedChanged        DeviceStateReason = 62
	DeviceStateReasonOvsdbFailed                 DeviceStateReason = 63
	DeviceStateReasonIPAddressDuplicate          DeviceStateReason = 64
	DeviceStateReasonIPMethodUnsupported         DeviceStateReason = 65
	DeviceStateReasonSriovConfigurationFailed    DeviceStateReason = 66
	DeviceStateReasonPeerNotFound                DeviceStateReason = 67
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ActiveConnectionStateReason
type ActiveConnectionStateReason uint32

const (
	ActiveConnectionStateReasonUnknown             ActiveConnectionStateReason = 0
	ActiveConnectionStateReasonNone                ActiveConnectionStateReason = 1
	ActiveConnectionStateReasonUserDisconnected    ActiveConnectionStateReason = 2
	ActiveConnectionStateReasonDeviceDisconnected  ActiveConnectionStateReason = 3
	ActiveConnectionStateReasonServiceStopped      ActiveConnectionStateReason = 4
	ActiveConnectionStateReasonIPConfigInvalid     ActiveConnectionStateReason = 5
	ActiveConnectionStateReasonConnectTimeout      ActiveConnectionStateReason = 6
	ActiveConnectionStateReasonServiceStartTimeout ActiveConnectionStateReason = 7
	ActiveConnectionStateReasonServiceStartFailed  ActiveConnectionStateReason = 8
	ActiveConnectionStateReasonNoSecrets           ActiveConnectionStateReason = 9
	ActiveConnectionStateReasonLoginFailed         ActiveConnectionStateReason = 10
	ActiveConnectionStateReasonConnectionRemoved   ActiveConnectionStateReason = 11
	ActiveConnectionStateReasonDependencyFailed    ActiveConnectionStateReason = 12
	ActiveConnectionStateReasonDeviceRealizeFailed ActiveConnectionStateReason = 13
	ActiveConnectionStateReasonDeviceRemoved       ActiveConnectionStateReason = 14
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=RollbackResult
type RollbackResult uint32

const (
	RollbackResultOk                 RollbackResult = 0
	RollbackResultErrNoDevice        RollbackResult = 1
	RollbackResultErrDeviceUnmanaged RollbackResult = 2
	RollbackResultErrFailed          RollbackResult = 3
)
