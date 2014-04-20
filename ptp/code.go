package ptp

type FunctionalMode uint16
type OperationCode uint16
type OperationResponseCode uint16
type ErrorCode uint16
type DevicePropCode uint16
type ObjectFormatCode uint16
type StorageTypes uint16
type FilesysytemTypes uint16
type AccessCapability uint16
type ProtectionStatus uint16
type AssociationTypes uint16
type DevicePropDesc uint16
type WhiteBalance uint16
type FocusMode uint16
type ExposureMeteringMode uint16
type FlashMode uint16
type ExposureProgramMode uint16
type StillCaptureMode uint16
type EffectMode uint16
type FocusMeteringMode uint16

const (
	FM_StandardMode FunctionalMode = 0x0000
	FM_SleepState   FunctionalMode = 0x0001

	OC_Undefinded           OperationCode = 0x1000
	OC_GetDeviceInfo        OperationCode = 0x1001
	OC_OpenSession          OperationCode = 0x1002
	OC_CloseSession         OperationCode = 0x1003
	OC_GetStorageIDs        OperationCode = 0x1004
	OC_GetStorageInfo       OperationCode = 0x1005
	OC_GetNumObjects        OperationCode = 0x1006
	OC_GetObjectHandles     OperationCode = 0x1007
	OC_GetObjectInfo        OperationCode = 0x1008
	OC_GetObject            OperationCode = 0x1009
	OC_GetThumb             OperationCode = 0x100A
	OC_DeleteObject         OperationCode = 0x100B
	OC_SendObjectInfo       OperationCode = 0x100C
	OC_SendObject           OperationCode = 0x100D
	OC_InitiateCapture      OperationCode = 0x100E
	OC_FormatStore          OperationCode = 0x100F
	OC_ResetDevice          OperationCode = 0x1010
	OC_SelfTest             OperationCode = 0x1011
	OC_SetObjectProtection  OperationCode = 0x1012
	OC_PowerDown            OperationCode = 0x1013
	OC_GetDevicePropDesc    OperationCode = 0x1014
	OC_GetDevicePropValue   OperationCode = 0x1015
	OC_SetDevicePropValue   OperationCode = 0x1016
	OC_ResetDevicePropValue OperationCode = 0x1017
	OC_TerminateOpenCapture OperationCode = 0x1018
	OC_MoveObject           OperationCode = 0x1019
	OC_CopyObject           OperationCode = 0x101A
	OC_GetPartialObject     OperationCode = 0x101B
	OC_InitiateOpenCapture  OperationCode = 0x101C

	RC_Undefined                             OperationResponseCode = 0x2000
	RC_OK                                    OperationResponseCode = 0x2001
	RC_GeneralError                          OperationResponseCode = 0x2002
	RC_SessionNotOpen                        OperationResponseCode = 0x2003
	RC_InvalidTransactionID                  OperationResponseCode = 0x2004
	RC_OperationNotSupported                 OperationResponseCode = 0x2005
	RC_ParameterNotSupported                 OperationResponseCode = 0x2006
	RC_IncompleteTransfer                    OperationResponseCode = 0x2007
	RC_InvalidStorageID                      OperationResponseCode = 0x2008
	RC_InvalidObjectHandle                   OperationResponseCode = 0x2009
	RC_DevicePropNotSupported                OperationResponseCode = 0x200a
	RC_InvalidObjectFormatCode               OperationResponseCode = 0x200b
	RC_StoreFull                             OperationResponseCode = 0x200c
	RC_ObjectWriteProtected                  OperationResponseCode = 0x200d
	RC_StoreReadOnly                         OperationResponseCode = 0x200e
	RC_AcceddDenied                          OperationResponseCode = 0x200f
	RC_NoThumbnailPresent                    OperationResponseCode = 0x2010
	RC_SelfTestFailed                        OperationResponseCode = 0x2011
	RC_PartialDeletion                       OperationResponseCode = 0x2012
	RC_StoreNotAvailable                     OperationResponseCode = 0x2013
	RC_SpecificationByFormatUnsupported      OperationResponseCode = 0x2014
	RC_NoValidObjectInfo                     OperationResponseCode = 0x2015
	RC_InvalidCodeFormat                     OperationResponseCode = 0x2016
	RC_UnknownVendorCode                     OperationResponseCode = 0x2017
	RC_CaptureAlreadyTerminated              OperationResponseCode = 0x2018
	RC_DeviceBusy                            OperationResponseCode = 0x2019
	RC_InvalidParentObject                   OperationResponseCode = 0x201a
	RC_InvalidDevicePropFormat               OperationResponseCode = 0x201b
	RC_InvalidDevicePropValue                OperationResponseCode = 0x201c
	RC_InvalidParameter                      OperationResponseCode = 0x201d
	RC_SessionAlreadyOpen                    OperationResponseCode = 0x201e
	RC_TransactionCancelled                  OperationResponseCode = 0x201f
	RC_SpecificationofDestinationUnsupported OperationResponseCode = 0x2020

	EC_Undefined             ErrorCode = 0x4000
	EC_CancelTransaction     ErrorCode = 0x4001
	EC_ObjectAdded           ErrorCode = 0x4002
	EC_ObjectRemoved         ErrorCode = 0x4003
	EC_StoreAdded            ErrorCode = 0x4004
	EC_StoreRemoved          ErrorCode = 0x4005
	EC_DevicePropChanged     ErrorCode = 0x4006
	EC_ObjectInfoChanged     ErrorCode = 0x4007
	EC_DeviceInfoChanged     ErrorCode = 0x4008
	EC_RequestObjectTransfer ErrorCode = 0x4009
	EC_StoreFull             ErrorCode = 0x400a
	EC_DeviceReset           ErrorCode = 0x400b
	EC_StorageInfoChanged    ErrorCode = 0x400c
	EC_CaptureComplete       ErrorCode = 0x400d
	EC_UnreportedStatus      ErrorCode = 0x400e

	DPC_Undefined                DevicePropCode = 0x5000
	DPC_BatteryLevel             DevicePropCode = 0x5001
	DPC_FunctionalMode           DevicePropCode = 0x5002
	DPC_ImageSize                DevicePropCode = 0x5003
	DPC_CompressionSetting       DevicePropCode = 0x5004
	DPC_WiteBalance              DevicePropCode = 0x5005
	DPC_RGBGain                  DevicePropCode = 0x5006
	DPC_FNumber                  DevicePropCode = 0x5007
	DPC_FocalLength              DevicePropCode = 0x5008
	DPC_FocusDistance            DevicePropCode = 0x5009
	DPC_FocusMode                DevicePropCode = 0x500a
	DPC_ExposureMeteringMode     DevicePropCode = 0x500b
	DPC_FlashMode                DevicePropCode = 0x500c
	DPC_ExposureTime             DevicePropCode = 0x500d
	DPC_ExposureProgramMode      DevicePropCode = 0x500e
	DPC_ExposureIndex            DevicePropCode = 0x500f
	DPC_ExposureBiasCompensation DevicePropCode = 0x5010
	DPC_DateTime                 DevicePropCode = 0x5011
	DPC_CaptureDelay             DevicePropCode = 0x5012
	DPC_StillCaptureMode         DevicePropCode = 0x5013
	DPC_Contrast                 DevicePropCode = 0x5014
	DPC_Sharpness                DevicePropCode = 0x5015
	DPC_DigitalZoom              DevicePropCode = 0x5016
	DPC_EffectMode               DevicePropCode = 0x5017
	DPC_BurstNumber              DevicePropCode = 0x5018
	DPC_BurstInterval            DevicePropCode = 0x5019
	DPC_TimelapseNumber          DevicePropCode = 0x501a
	DPC_TimelapseInterval        DevicePropCode = 0x501b
	DPC_FocusMeteringMode        DevicePropCode = 0x501c
	DPC_UploadURL                DevicePropCode = 0x501d
	DPC_Artist                   DevicePropCode = 0x501e
	DPC_CopyrightInfo            DevicePropCode = 0X501F

	OFC_Undefined         ObjectFormatCode = 0x3000
	OFC_Association       ObjectFormatCode = 0x3001
	OFC_Script            ObjectFormatCode = 0x3002
	OFC_Executable        ObjectFormatCode = 0x3003
	OFC_Text              ObjectFormatCode = 0x3004
	OFC_HTML              ObjectFormatCode = 0x3005
	OFC_DPOF              ObjectFormatCode = 0x3006
	OFC_AIFF              ObjectFormatCode = 0x3007
	OFC_WAV               ObjectFormatCode = 0x3008
	OFC_MP3               ObjectFormatCode = 0x3009
	OFC_AVI               ObjectFormatCode = 0x300a
	OFC_MPEG              ObjectFormatCode = 0x300b
	OFC_ASF               ObjectFormatCode = 0x300c
	OFC_Unknown           ObjectFormatCode = 0x3800
	OFC_EXIF_JPEG         ObjectFormatCode = 0x3801
	OFC_TIFF_EP           ObjectFormatCode = 0x3802
	OFC_FlashPix          ObjectFormatCode = 0x3803
	OFC_BMP               ObjectFormatCode = 0x3804
	OFC_CIFF              ObjectFormatCode = 0x3805
	OFC_GIF               ObjectFormatCode = 0x3807
	OFC_JFIF              ObjectFormatCode = 0x3808
	OFC_PCD               ObjectFormatCode = 0x3809
	OFC_PICT              ObjectFormatCode = 0x380a
	OFC_PNG               ObjectFormatCode = 0x380b
	OFC_TIFF              ObjectFormatCode = 0x380d
	OFC_TIFF_IT           ObjectFormatCode = 0x380e
	OFC_JP2               ObjectFormatCode = 0x380f
	OFC_JPX               ObjectFormatCode = 0x3810
	OFC_AncillaryDataFile ObjectFormatCode = 0x3000
	OFC_ImageFile         ObjectFormatCode = 0x3800

	ST_Undefined    StorageTypes = 0x0000
	ST_FixedROM     StorageTypes = 0x0001
	ST_RemovableROM StorageTypes = 0x0002
	ST_FixedRAM     StorageTypes = 0x0003
	ST_RemovableRAM StorageTypes = 0x0004

	FT_Undefined           FilesysytemTypes = 0x0000
	FT_GenericFlat         FilesysytemTypes = 0x0001
	FT_GenericHierarchical FilesysytemTypes = 0x0002
	FT_DCF                 FilesysytemTypes = 0x0003

	AC_ReadWrite         AccessCapability = 0x0000
	AC_ReadOnly          AccessCapability = 0x0001
	AC_ReadOnly_Deletion AccessCapability = 0x0002

	PS_NoProtection ProtectionStatus = 0x0000
	PS_ReadOnly     ProtectionStatus = 0x0001

	AT_Undefined           AssociationTypes = 0x0000
	AT_GenericFolder       AssociationTypes = 0x0001
	AT_Album               AssociationTypes = 0x0002
	AT_TimeSequence        AssociationTypes = 0x0003
	AT_HorizontalPanoramic AssociationTypes = 0x0004
	AT_VerticalPanoramic   AssociationTypes = 0x0005
	AT_2DPanoramic         AssociationTypes = 0x0006
	AT_AncillaryData       AssociationTypes = 0x0007

	DPD_Get            DevicePropDesc = 0x00
	DPD_Set            DevicePropDesc = 0x01
	DPD_FormFlag_None  DevicePropDesc = 0x00
	DPD_FormFlag_Range DevicePropDesc = 0x01
	DPD_FormFlag_Enum  DevicePropDesc = 0x02

	WB_Undefined        WhiteBalance = 0x0000
	WB_Manual           WhiteBalance = 0x0001
	WB_Automatic        WhiteBalance = 0x0002
	WB_OnePushAutomatic WhiteBalance = 0x0003
	WB_Daylight         WhiteBalance = 0x0004
	WB_Florescent       WhiteBalance = 0x0005
	WB_Tungsten         WhiteBalance = 0x0006
	WB_Flush            WhiteBalance = 0x0007

	FCM_Undefined      FocusMode = 0x0000
	FCM_Manual         FocusMode = 0x0001
	FCM_Automatic      FocusMode = 0x0002
	FCM_AutomaticMacro FocusMode = 0x0003

	EMM_Undefined             ExposureMeteringMode = 0x0000
	EMM_Avarage               ExposureMeteringMode = 0x0001
	EMM_CenterWeightedAvarage ExposureMeteringMode = 0x0002
	EMM_MultiSpot             ExposureMeteringMode = 0x0003
	EMM_CenterSpot            ExposureMeteringMode = 0x0004

	FLM_Undefined    FlashMode = 0x0000
	FLM_AutoFlash    FlashMode = 0x0001
	FLM_FlashOff     FlashMode = 0x0002
	FLM_FillFlash    FlashMode = 0x0003
	FLM_RedEyeAuto   FlashMode = 0x0004
	FLM_RedEyeFill   FlashMode = 0x0005
	FLM_ExternalSync FlashMode = 0x0006

	EPM_Undefined        ExposureProgramMode = 0x0000
	EPM_Manual           ExposureProgramMode = 0x0001
	EPM_Automatic        ExposureProgramMode = 0x0002
	EPM_AperturePriority ExposureProgramMode = 0x0003
	EPM_SutterPriority   ExposureProgramMode = 0x0004
	EPM_ProgramCreative  ExposureProgramMode = 0x0005
	EPM_ProgramAction    ExposureProgramMode = 0x0006
	EPM_Portrait         ExposureProgramMode = 0x0007

	SCM_Undefined StillCaptureMode = 0x0000
	SCM_Normal    StillCaptureMode = 0x0001
	SCM_Burst     StillCaptureMode = 0x0002
	SCM_Timelapse StillCaptureMode = 0x0003

	EM_Undefined  EffectMode = 0x0000
	EM_Standard   EffectMode = 0x0001
	EM_BlackWhite EffectMode = 0x0002
	EM_Sepia      EffectMode = 0x0003

	FMM_Undefined  FocusMeteringMode = 0x0000
	FMM_CenterSpot FocusMeteringMode = 0x0001
)
