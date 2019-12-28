// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is auto-generated from our dictionaries.

package avp

// Diameter AVP types.
const (
	AccessNetworkChargingAddress      = 501
	AccessNetworkChargingIdentifierGx = 1022
	AccessNetworkType                 = 306
	AccountingRealtimeRequired        = 483
	AccountingRecordNumber            = 485
	AccountingRecordType              = 480
	AccountingSessionID               = 44
	AccountingSubSessionID            = 287
	AcctApplicationID                 = 259
	AcctInterimInterval               = 85
	AcctMultiSessionID                = 50
	AddressRealm                      = 301
	AFApplicationIdentifier           = 504
	AggregationNetworkType            = 307
	ANGWAddress                       = 1050
	AuthApplicationID                 = 258
	AuthGracePeriod                   = 276
	AuthorizationLifetime             = 291
	AuthRequestType                   = 274
	AuthSessionState                  = 277
	BearerUsage                       = 1000
	CCCorrelationID                   = 411
	CCInputOctets                     = 412
	CCMoney                           = 413
	CCOutputOctets                    = 414
	CCRequestNumber                   = 415
	CCRequestType                     = 416
	CCServiceSpecificUnits            = 417
	CCSessionFailover                 = 418
	CCSubSessionID                    = 419
	CCTime                            = 420
	CCTotalOctets                     = 421
	CCUnitType                        = 454
	ChargingRuleBaseName              = 1004
	ChargingRuleDefinition            = 1003
	ChargingRuleInstall               = 1001
	ChargingRuleName                  = 1005
	ChargingRuleRemove                = 1002
	CheckBalanceResult                = 422
	CivicLocation                     = 355
	Class                             = 25
	CostInformation                   = 423
	CostUnit                          = 424
	CreditControl                     = 426
	CreditControlFailureHandling      = 427
	CurrencyCode                      = 425
	DefaultEPSBearerQoS               = 1049
	DestinationHost                   = 293
	DestinationRealm                  = 283
	DirectDebitingFailureHandling     = 428
	DisconnectCause                   = 273
	DRMP                              = 301
	EmergencyCallRoutingInfo          = 361
	ErrorMessage                      = 281
	ErrorReportingHost                = 294
	EventTimestamp                    = 55
	EventTrigger                      = 1006
	EventType                         = 354
	ExperimentalResult                = 297
	ExperimentalResultCode            = 298
	ExpiryTime                        = 709
	Exponent                          = 429
	FailedAVP                         = 279
	FinalUnitAction                   = 449
	FinalUnitIndication               = 430
	FirmwareRevision                  = 267
	FixedAccessID                     = 358
	FlowDescription                   = 507
	FlowDirection                     = 1080
	FlowInformation                   = 1058
	FlowLabel                         = 1057
	FramedIPAddress                   = 8
	FramedIPv6Prefix                  = 97
	GeospatialLocation                = 356
	GlobalAccessID                    = 357
	GloballyUniqueAddress             = 300
	GrantedServiceUnit                = 431
	GSUPoolIdentifier                 = 453
	GSUPoolReference                  = 457
	HostIPAddress                     = 257
	InbandSecurityID                  = 299
	InitialGateSetting                = 303
	IPCANType                         = 1027
	IPConnectivityStatus              = 305
	LineIdentifier                    = 500
	LocationInformation               = 350
	LogicalAccessID                   = 302
	MaxRequestedBandwidthDL           = 515
	MaxRequestedBandwidthUL           = 516
	MonitoringKey                     = 1066
	MultipleServicesCreditControl     = 456
	MultipleServicesIndicator         = 455
	MultiRoundTimeOut                 = 272
	NetworkRequestSupport             = 1024
	OCFeatureVector                   = 622
	OCOLR                             = 623
	OCReductionPercentage             = 627
	OCReportType                      = 626
	OCSequenceNumber                  = 624
	OCSupportedFeatures               = 621
	OCValidityDuration                = 625
	Offline                           = 1008
	Online                            = 1009
	OperatorSpecificGI                = 360
	OriginHost                        = 264
	OriginRealm                       = 296
	OriginStateID                     = 278
	PacketFilterIdentifier            = 1060
	PacketFilterUsage                 = 1072
	PIDFLocationObject                = 363
	PolicyControlContactPoint         = 351
	PortNumber                        = 362
	Precedence                        = 1010
	ProductName                       = 269
	ProxyHost                         = 280
	ProxyInfo                         = 284
	ProxyState                        = 33
	QoSProfile                        = 304
	RatingGroup                       = 432
	ReAuthRequestType                 = 285
	RedirectAddressType               = 433
	RedirectHost                      = 292
	RedirectHostUsage                 = 261
	RedirectInformation               = 1085
	RedirectMaxCacheTime              = 262
	RedirectServer                    = 434
	RedirectServerAddress             = 435
	RedirectSupport                   = 1086
	RelayAgent                        = 359
	RequestedAction                   = 436
	RequestedInformation              = 353
	RequestedServiceUnit              = 437
	RestrictionFilterRule             = 438
	ResultCode                        = 268
	RevalidationTime                  = 1042
	RouteRecord                       = 282
	RuleActivationTime                = 1043
	RuleDeactivationTime              = 1044
	SecurityParameterIndex            = 1056
	ServiceContextID                  = 461
	ServiceIdentifier                 = 439
	ServiceParameterInfo              = 440
	ServiceParameterType              = 441
	ServiceParameterValue             = 442
	SessionBinding                    = 270
	SessionID                         = 263
	SessionServerFailover             = 271
	SessionTimeout                    = 27
	SLRequestType                     = 2904
	SubscriptionID                    = 443
	SubscriptionIDData                = 444
	SubscriptionIDType                = 450
	SubsReqType                       = 705
	SupportedVendorID                 = 265
	TariffChangeUsage                 = 452
	TariffTimeChange                  = 451
	TerminalType                      = 352
	TerminationCause                  = 295
	TFTFilter                         = 1012
	TFTPacketFilterInformation        = 1013
	TGPPGGSNAddress                   = 7
	TGPPMSTimeZone                    = 23
	TGPPSGSNAddress                   = 6
	TGPPUserLocationInfo              = 22
	ToSTrafficClass                   = 1014
	UnitValue                         = 445
	UsageMonitoringInformation        = 1067
	UsageMonitoringLevel              = 1068
	UsedServiceUnit                   = 446
	UserEquipmentInfo                 = 458
	UserEquipmentInfoType             = 459
	UserEquipmentInfoValue            = 460
	UserName                          = 1
	ValidityTime                      = 448
	ValueDigits                       = 447
	VendorID                          = 266
	VendorSpecificApplicationID       = 260
)
