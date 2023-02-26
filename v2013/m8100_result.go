package v2013

type M8100Result uint8

const (
	// M8100Success 成功
	M8100Success M8100Result = iota

	// M8100CarRegistered 车辆已被注册
	M8100CarRegistered

	// M8100CarNotInDB 数据库中无该车辆
	M8100CarNotInDB

	// M8100TermRegistered 终端已被注册
	M8100TermRegistered

	// M8100TermNotInDB 数据库中无该终端
	M8100TermNotInDB
)
