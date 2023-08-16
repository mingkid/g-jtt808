package common

// RealtimeStreamEncodingMode 实时流编码模式
type RealtimeStreamEncodingMode uint8

const (
	// RealtimeStreamCBR 恒定比特率编码模式
	RealtimeStreamCBR RealtimeStreamEncodingMode = iota

	// RealtimeStreamVBR 可变比特率编码模式
	RealtimeStreamVBR

	// RealtimeStreamABR 平均比特率编码模式
	RealtimeStreamABR
)

// RealtimeStreamResolution 实时流分辨率
type RealtimeStreamResolution uint8

const (
	// RealtimeStreamQCIF QCIF分辨率
	RealtimeStreamQCIF RealtimeStreamResolution = iota

	// RealtimeStreamCIF CIF分辨率
	RealtimeStreamCIF

	// RealtimeStreamWCIF WCIF分辨率
	RealtimeStreamWCIF

	// RealtimeStreamDI DI分辨率
	RealtimeStreamDI

	// RealtimeStreamWDI WDI分辨率
	RealtimeStreamWDI

	// RealtimeStream720P 720P分辨率
	RealtimeStream720P

	// RealtimeStream1080P 1080P分辨率
	RealtimeStream1080P
)

// StorageStreamEncodingMode 存储流编码模式枚举
type StorageStreamEncodingMode uint8

const (
	// StorageStreamEncodingModeCBR 恒定比特率编码模式
	StorageStreamEncodingModeCBR StorageStreamEncodingMode = iota

	// StorageStreamEncodingModeVBR 可变比特率编码模式
	StorageStreamEncodingModeVBR

	// StorageStreamEncodingModeABR 平均比特率编码模式
	StorageStreamEncodingModeABR
)

// StorageStreamResolution 存储流分辨率枚举
type StorageStreamResolution uint8

const (
	// StorageStreamResolutionQCIF QCIF分辨率
	StorageStreamResolutionQCIF StorageStreamResolution = iota

	// StorageStreamResolutionCIF CIF分辨率
	StorageStreamResolutionCIF

	// StorageStreamResolutionWCIF WCIF分辨率
	StorageStreamResolutionWCIF

	// StorageStreamResolutionD1 D1分辨率
	StorageStreamResolutionD1

	// StorageStreamResolutionWD1 WD1分辨率
	StorageStreamResolutionWD1

	// StorageStreamResolution720P 720P分辨率
	StorageStreamResolution720P

	// StorageStreamResolution1080P 1080P分辨率
	StorageStreamResolution1080P
)

// OSDOverlaySettingsBitmask OSD字幕叠加设置位掩码
type OSDOverlaySettingsBitmask uint16

const (
	// OverlayDateAndTime 日期和时间叠加
	OverlayDateAndTime OSDOverlaySettingsBitmask = 1 << iota

	// OverlayLicensePlateNumber 车牌号叠加
	OverlayLicensePlateNumber

	// OverlayLogicalChannel 逻辑通道号叠加
	OverlayLogicalChannel

	// OverlayLatitudeLongitude 经纬度叠加
	OverlayLatitudeLongitude

	// OverlayDrivingSpeed 行驶速度叠加
	OverlayDrivingSpeed

	// OverlaySatelliteSpeed 卫星速度叠加
	OverlaySatelliteSpeed

	// OverlayContinuousDrivingTime 连续驾驶时间叠加
	OverlayContinuousDrivingTime
)
