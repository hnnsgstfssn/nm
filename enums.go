package nm

import "github.com/hnnsgstfssn/nm/internal/nmtype"

// The enums live in an internal package so that seventeen files of stringer
// output stay out of this one's list. They are aliased rather than
// re-declared, so nm.DeviceState and nmtype.DeviceState are the same type
// and no caller has to name the internal package.
type (
	Connectivity                = nmtype.Connectivity
	State                       = nmtype.State
	CheckpointCreateFlags       = nmtype.CheckpointCreateFlags
	Capability                  = nmtype.Capability
	Metered                     = nmtype.Metered
	DeviceState                 = nmtype.DeviceState
	DeviceStateReason           = nmtype.DeviceStateReason
	DeviceInterfaceFlags        = nmtype.DeviceInterfaceFlags
	ActiveConnectionState       = nmtype.ActiveConnectionState
	ActiveConnectionStateReason = nmtype.ActiveConnectionStateReason
	ActivationStateFlag         = nmtype.ActivationStateFlag
	DeviceType                  = nmtype.DeviceType
	APFlags                     = nmtype.APFlags
	APSecurity                  = nmtype.APSecurity
	WifiMode                    = nmtype.WifiMode
	VpnConnectionState          = nmtype.VpnConnectionState
	RollbackResult              = nmtype.RollbackResult
)

const (
	ConnectivityUnknown = nmtype.ConnectivityUnknown
	ConnectivityNone    = nmtype.ConnectivityNone
	ConnectivityPortal  = nmtype.ConnectivityPortal
	ConnectivityLimited = nmtype.ConnectivityLimited
	ConnectivityFull    = nmtype.ConnectivityFull
)

const (
	StateUnknown         = nmtype.StateUnknown
	StateAsleep          = nmtype.StateAsleep
	StateDisconnected    = nmtype.StateDisconnected
	StateDisconnecting   = nmtype.StateDisconnecting
	StateConnecting      = nmtype.StateConnecting
	StateConnectedLocal  = nmtype.StateConnectedLocal
	StateConnectedSite   = nmtype.StateConnectedSite
	StateConnectedGlobal = nmtype.StateConnectedGlobal
)

const (
	CheckpointCreateFlagsNone                 = nmtype.CheckpointCreateFlagsNone
	CheckpointCreateFlagsDestroyAll           = nmtype.CheckpointCreateFlagsDestroyAll
	CheckpointCreateFlagsDeleteNewConnections = nmtype.CheckpointCreateFlagsDeleteNewConnections
	CheckpointCreateFlagsDisconnectNewDevices = nmtype.CheckpointCreateFlagsDisconnectNewDevices
	CheckpointCreateFlagsAllowOverlapping     = nmtype.CheckpointCreateFlagsAllowOverlapping
)

const (
	CapabilityTeam = nmtype.CapabilityTeam
)

const (
	MeteredUnknown  = nmtype.MeteredUnknown
	MeteredYes      = nmtype.MeteredYes
	MeteredNo       = nmtype.MeteredNo
	MeteredGuessYes = nmtype.MeteredGuessYes
	MeteredGuessNo  = nmtype.MeteredGuessNo
)

const (
	DeviceStateUnknown      = nmtype.DeviceStateUnknown
	DeviceStateUnmanaged    = nmtype.DeviceStateUnmanaged
	DeviceStateUnavailable  = nmtype.DeviceStateUnavailable
	DeviceStateDisconnected = nmtype.DeviceStateDisconnected
	DeviceStatePrepare      = nmtype.DeviceStatePrepare
	DeviceStateConfig       = nmtype.DeviceStateConfig
	DeviceStateNeedAuth     = nmtype.DeviceStateNeedAuth
	DeviceStateIPConfig     = nmtype.DeviceStateIPConfig
	DeviceStateIPCheck      = nmtype.DeviceStateIPCheck
	DeviceStateSecondaries  = nmtype.DeviceStateSecondaries
	DeviceStateActivated    = nmtype.DeviceStateActivated
	DeviceStateDeactivating = nmtype.DeviceStateDeactivating
	DeviceStateFailed       = nmtype.DeviceStateFailed
)

const (
	DeviceInterfaceFlagsNone              = nmtype.DeviceInterfaceFlagsNone
	DeviceInterfaceFlagsUp                = nmtype.DeviceInterfaceFlagsUp
	DeviceInterfaceFlagsLowerUp           = nmtype.DeviceInterfaceFlagsLowerUp
	DeviceInterfaceFlagsPromisc           = nmtype.DeviceInterfaceFlagsPromisc
	DeviceInterfaceFlagsCarrier           = nmtype.DeviceInterfaceFlagsCarrier
	DeviceInterfaceFlagsLldpClientEnabled = nmtype.DeviceInterfaceFlagsLldpClientEnabled
)

const (
	ActiveConnectionStateUnknown      = nmtype.ActiveConnectionStateUnknown
	ActiveConnectionStateActivating   = nmtype.ActiveConnectionStateActivating
	ActiveConnectionStateActivated    = nmtype.ActiveConnectionStateActivated
	ActiveConnectionStateDeactivating = nmtype.ActiveConnectionStateDeactivating
	ActiveConnectionStateDeactivated  = nmtype.ActiveConnectionStateDeactivated
)

const (
	ActivationStateFlagNone                             = nmtype.ActivationStateFlagNone
	ActivationStateFlagIsMaster                         = nmtype.ActivationStateFlagIsMaster
	ActivationStateFlagIsSlave                          = nmtype.ActivationStateFlagIsSlave
	ActivationStateFlagLayer2Ready                      = nmtype.ActivationStateFlagLayer2Ready
	ActivationStateFlagIP4Ready                         = nmtype.ActivationStateFlagIP4Ready
	ActivationStateFlagIP6Ready                         = nmtype.ActivationStateFlagIP6Ready
	ActivationStateFlagMasterHasSlaves                  = nmtype.ActivationStateFlagMasterHasSlaves
	ActivationStateFlagLifetimeBoundToProfileVisibility = nmtype.ActivationStateFlagLifetimeBoundToProfileVisibility
)

const (
	DeviceTypeUnknown      = nmtype.DeviceTypeUnknown
	DeviceTypeEthernet     = nmtype.DeviceTypeEthernet
	DeviceTypeWifi         = nmtype.DeviceTypeWifi
	DeviceTypeUnused1      = nmtype.DeviceTypeUnused1
	DeviceTypeUnused2      = nmtype.DeviceTypeUnused2
	DeviceTypeBt           = nmtype.DeviceTypeBt
	DeviceTypeOlpcMesh     = nmtype.DeviceTypeOlpcMesh
	DeviceTypeWimax        = nmtype.DeviceTypeWimax
	DeviceTypeModem        = nmtype.DeviceTypeModem
	DeviceTypeInfiniband   = nmtype.DeviceTypeInfiniband
	DeviceTypeBond         = nmtype.DeviceTypeBond
	DeviceTypeVlan         = nmtype.DeviceTypeVlan
	DeviceTypeAdsl         = nmtype.DeviceTypeAdsl
	DeviceTypeBridge       = nmtype.DeviceTypeBridge
	DeviceTypeGeneric      = nmtype.DeviceTypeGeneric
	DeviceTypeTeam         = nmtype.DeviceTypeTeam
	DeviceTypeTun          = nmtype.DeviceTypeTun
	DeviceTypeIPTunnel     = nmtype.DeviceTypeIPTunnel
	DeviceTypeMacvlan      = nmtype.DeviceTypeMacvlan
	DeviceTypeVxlan        = nmtype.DeviceTypeVxlan
	DeviceTypeVeth         = nmtype.DeviceTypeVeth
	DeviceTypeMacsec       = nmtype.DeviceTypeMacsec
	DeviceTypeDummy        = nmtype.DeviceTypeDummy
	DeviceTypePpp          = nmtype.DeviceTypePpp
	DeviceTypeOvsInterface = nmtype.DeviceTypeOvsInterface
	DeviceTypeOvsPort      = nmtype.DeviceTypeOvsPort
	DeviceTypeOvsBridge    = nmtype.DeviceTypeOvsBridge
	DeviceTypeWpan         = nmtype.DeviceTypeWpan
	DeviceType6lowpan      = nmtype.DeviceType6lowpan
	DeviceTypeWireguard    = nmtype.DeviceTypeWireguard
	DeviceTypeWifiP2p      = nmtype.DeviceTypeWifiP2p
)

const (
	APFlagsNone    = nmtype.APFlagsNone
	APFlagsPrivacy = nmtype.APFlagsPrivacy
)

const (
	APSecurityNone         = nmtype.APSecurityNone
	APSecurityPairWEP40    = nmtype.APSecurityPairWEP40
	APSecurityPairWEP104   = nmtype.APSecurityPairWEP104
	APSecurityPairTKIP     = nmtype.APSecurityPairTKIP
	APSecurityPairCCMP     = nmtype.APSecurityPairCCMP
	APSecurityGroupWEP40   = nmtype.APSecurityGroupWEP40
	APSecurityGroupWEP104  = nmtype.APSecurityGroupWEP104
	APSecurityGroupTKIP    = nmtype.APSecurityGroupTKIP
	APSecurityGroupCCMP    = nmtype.APSecurityGroupCCMP
	APSecurityKeyMgmtPSK   = nmtype.APSecurityKeyMgmtPSK
	APSecurityKeyMgmt8021X = nmtype.APSecurityKeyMgmt8021X
	APSecurityKeyMgmtSAE   = nmtype.APSecurityKeyMgmtSAE
	APSecurityKeyMgmtOWE   = nmtype.APSecurityKeyMgmtOWE
	APSecurityKeyMgmtOWETM = nmtype.APSecurityKeyMgmtOWETM
)

const (
	WifiModeUnknown = nmtype.WifiModeUnknown
	WifiModeAdhoc   = nmtype.WifiModeAdhoc
	WifiModeInfra   = nmtype.WifiModeInfra
	WifiModeAP      = nmtype.WifiModeAP
)

const (
	VpnConnectionUnknown      = nmtype.VpnConnectionUnknown
	VpnConnectionPrepare      = nmtype.VpnConnectionPrepare
	VpnConnectionNeedAuth     = nmtype.VpnConnectionNeedAuth
	VpnConnectionConnect      = nmtype.VpnConnectionConnect
	VpnConnectionIPConfigGet  = nmtype.VpnConnectionIPConfigGet
	VpnConnectionActivated    = nmtype.VpnConnectionActivated
	VpnConnectionFailed       = nmtype.VpnConnectionFailed
	VpnConnectionDisconnected = nmtype.VpnConnectionDisconnected
)

const (
	DeviceStateReasonNone                        = nmtype.DeviceStateReasonNone
	DeviceStateReasonUnknown                     = nmtype.DeviceStateReasonUnknown
	DeviceStateReasonNowManaged                  = nmtype.DeviceStateReasonNowManaged
	DeviceStateReasonNowUnmanaged                = nmtype.DeviceStateReasonNowUnmanaged
	DeviceStateReasonConfigFailed                = nmtype.DeviceStateReasonConfigFailed
	DeviceStateReasonIPConfigUnavailable         = nmtype.DeviceStateReasonIPConfigUnavailable
	DeviceStateReasonIPConfigExpired             = nmtype.DeviceStateReasonIPConfigExpired
	DeviceStateReasonNoSecrets                   = nmtype.DeviceStateReasonNoSecrets
	DeviceStateReasonSupplicantDisconnect        = nmtype.DeviceStateReasonSupplicantDisconnect
	DeviceStateReasonSupplicantConfigFailed      = nmtype.DeviceStateReasonSupplicantConfigFailed
	DeviceStateReasonSupplicantFailed            = nmtype.DeviceStateReasonSupplicantFailed
	DeviceStateReasonSupplicantTimeout           = nmtype.DeviceStateReasonSupplicantTimeout
	DeviceStateReasonPppStartFailed              = nmtype.DeviceStateReasonPppStartFailed
	DeviceStateReasonPppDisconnect               = nmtype.DeviceStateReasonPppDisconnect
	DeviceStateReasonPppFailed                   = nmtype.DeviceStateReasonPppFailed
	DeviceStateReasonDhcpStartFailed             = nmtype.DeviceStateReasonDhcpStartFailed
	DeviceStateReasonDhcpError                   = nmtype.DeviceStateReasonDhcpError
	DeviceStateReasonDhcpFailed                  = nmtype.DeviceStateReasonDhcpFailed
	DeviceStateReasonSharedStartFailed           = nmtype.DeviceStateReasonSharedStartFailed
	DeviceStateReasonSharedFailed                = nmtype.DeviceStateReasonSharedFailed
	DeviceStateReasonAutoipStartFailed           = nmtype.DeviceStateReasonAutoipStartFailed
	DeviceStateReasonAutoipError                 = nmtype.DeviceStateReasonAutoipError
	DeviceStateReasonAutoipFailed                = nmtype.DeviceStateReasonAutoipFailed
	DeviceStateReasonModemBusy                   = nmtype.DeviceStateReasonModemBusy
	DeviceStateReasonModemNoDialTone             = nmtype.DeviceStateReasonModemNoDialTone
	DeviceStateReasonModemNoCarrier              = nmtype.DeviceStateReasonModemNoCarrier
	DeviceStateReasonModemDialTimeout            = nmtype.DeviceStateReasonModemDialTimeout
	DeviceStateReasonModemDialFailed             = nmtype.DeviceStateReasonModemDialFailed
	DeviceStateReasonModemInitFailed             = nmtype.DeviceStateReasonModemInitFailed
	DeviceStateReasonGsmApnFailed                = nmtype.DeviceStateReasonGsmApnFailed
	DeviceStateReasonGsmRegistrationNotSearching = nmtype.DeviceStateReasonGsmRegistrationNotSearching
	DeviceStateReasonGsmRegistrationDenied       = nmtype.DeviceStateReasonGsmRegistrationDenied
	DeviceStateReasonGsmRegistrationTimeout      = nmtype.DeviceStateReasonGsmRegistrationTimeout
	DeviceStateReasonGsmRegistrationFailed       = nmtype.DeviceStateReasonGsmRegistrationFailed
	DeviceStateReasonGsmPinCheckFailed           = nmtype.DeviceStateReasonGsmPinCheckFailed
	DeviceStateReasonFirmwareMissing             = nmtype.DeviceStateReasonFirmwareMissing
	DeviceStateReasonRemoved                     = nmtype.DeviceStateReasonRemoved
	DeviceStateReasonSleeping                    = nmtype.DeviceStateReasonSleeping
	DeviceStateReasonConnectionRemoved           = nmtype.DeviceStateReasonConnectionRemoved
	DeviceStateReasonUserRequested               = nmtype.DeviceStateReasonUserRequested
	DeviceStateReasonCarrier                     = nmtype.DeviceStateReasonCarrier
	DeviceStateReasonConnectionAssumed           = nmtype.DeviceStateReasonConnectionAssumed
	DeviceStateReasonSupplicantAvailable         = nmtype.DeviceStateReasonSupplicantAvailable
	DeviceStateReasonModemNotFound               = nmtype.DeviceStateReasonModemNotFound
	DeviceStateReasonBtFailed                    = nmtype.DeviceStateReasonBtFailed
	DeviceStateReasonGsmSimNotInserted           = nmtype.DeviceStateReasonGsmSimNotInserted
	DeviceStateReasonGsmSimPinRequired           = nmtype.DeviceStateReasonGsmSimPinRequired
	DeviceStateReasonGsmSimPukRequired           = nmtype.DeviceStateReasonGsmSimPukRequired
	DeviceStateReasonGsmSimWrong                 = nmtype.DeviceStateReasonGsmSimWrong
	DeviceStateReasonInfinibandMode              = nmtype.DeviceStateReasonInfinibandMode
	DeviceStateReasonDependencyFailed            = nmtype.DeviceStateReasonDependencyFailed
	DeviceStateReasonBr2684Failed                = nmtype.DeviceStateReasonBr2684Failed
	DeviceStateReasonModemManagerUnavailable     = nmtype.DeviceStateReasonModemManagerUnavailable
	DeviceStateReasonSsidNotFound                = nmtype.DeviceStateReasonSsidNotFound
	DeviceStateReasonSecondaryConnectionFailed   = nmtype.DeviceStateReasonSecondaryConnectionFailed
	DeviceStateReasonDcbFcoeFailed               = nmtype.DeviceStateReasonDcbFcoeFailed
	DeviceStateReasonTeamdControlFailed          = nmtype.DeviceStateReasonTeamdControlFailed
	DeviceStateReasonModemFailed                 = nmtype.DeviceStateReasonModemFailed
	DeviceStateReasonModemAvailable              = nmtype.DeviceStateReasonModemAvailable
	DeviceStateReasonSimPinIncorrect             = nmtype.DeviceStateReasonSimPinIncorrect
	DeviceStateReasonNewActivation               = nmtype.DeviceStateReasonNewActivation
	DeviceStateReasonParentChanged               = nmtype.DeviceStateReasonParentChanged
	DeviceStateReasonParentManagedChanged        = nmtype.DeviceStateReasonParentManagedChanged
	DeviceStateReasonOvsdbFailed                 = nmtype.DeviceStateReasonOvsdbFailed
	DeviceStateReasonIPAddressDuplicate          = nmtype.DeviceStateReasonIPAddressDuplicate
	DeviceStateReasonIPMethodUnsupported         = nmtype.DeviceStateReasonIPMethodUnsupported
	DeviceStateReasonSriovConfigurationFailed    = nmtype.DeviceStateReasonSriovConfigurationFailed
	DeviceStateReasonPeerNotFound                = nmtype.DeviceStateReasonPeerNotFound
)

const (
	ActiveConnectionStateReasonUnknown             = nmtype.ActiveConnectionStateReasonUnknown
	ActiveConnectionStateReasonNone                = nmtype.ActiveConnectionStateReasonNone
	ActiveConnectionStateReasonUserDisconnected    = nmtype.ActiveConnectionStateReasonUserDisconnected
	ActiveConnectionStateReasonDeviceDisconnected  = nmtype.ActiveConnectionStateReasonDeviceDisconnected
	ActiveConnectionStateReasonServiceStopped      = nmtype.ActiveConnectionStateReasonServiceStopped
	ActiveConnectionStateReasonIPConfigInvalid     = nmtype.ActiveConnectionStateReasonIPConfigInvalid
	ActiveConnectionStateReasonConnectTimeout      = nmtype.ActiveConnectionStateReasonConnectTimeout
	ActiveConnectionStateReasonServiceStartTimeout = nmtype.ActiveConnectionStateReasonServiceStartTimeout
	ActiveConnectionStateReasonServiceStartFailed  = nmtype.ActiveConnectionStateReasonServiceStartFailed
	ActiveConnectionStateReasonNoSecrets           = nmtype.ActiveConnectionStateReasonNoSecrets
	ActiveConnectionStateReasonLoginFailed         = nmtype.ActiveConnectionStateReasonLoginFailed
	ActiveConnectionStateReasonConnectionRemoved   = nmtype.ActiveConnectionStateReasonConnectionRemoved
	ActiveConnectionStateReasonDependencyFailed    = nmtype.ActiveConnectionStateReasonDependencyFailed
	ActiveConnectionStateReasonDeviceRealizeFailed = nmtype.ActiveConnectionStateReasonDeviceRealizeFailed
	ActiveConnectionStateReasonDeviceRemoved       = nmtype.ActiveConnectionStateReasonDeviceRemoved
)

const (
	RollbackResultOk                 = nmtype.RollbackResultOk
	RollbackResultErrNoDevice        = nmtype.RollbackResultErrNoDevice
	RollbackResultErrDeviceUnmanaged = nmtype.RollbackResultErrDeviceUnmanaged
	RollbackResultErrFailed          = nmtype.RollbackResultErrFailed
)
