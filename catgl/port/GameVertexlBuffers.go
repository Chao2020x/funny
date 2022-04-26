package port

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// GameVertexlBuffersType 游戏顶点缓冲区类型
type GameVertexlBuffersType int

const (
	// GameVertexSTATIC 表示该缓存区不会被修改
	GameVertexSTATIC GameVertexlBuffersType = gl.STATIC_DRAW
	// GameVertexDyNAMIC 表示该缓存区会被周期性更改
	GameVertexDyNAMIC GameVertexlBuffersType = gl.DYNAMIC_DRAW
	// GameVertexSTREAM 表示该缓存区会被频繁更改
	GameVertexSTREAM GameVertexlBuffersType = gl.STREAM_DRAW
)

// GameVertexBuffersSync 顶点缓冲区同步标记
type GameVertexBuffersSync int

const (
	// GameVertexSyncDefault 默认
	GameVertexSyncDefault GameVertexBuffersSync = 0
	// GameVertexSyncRetrieve 取回
	GameVertexSyncRetrieve GameVertexBuffersSync = 1
	// GameVertexSyncDelivery 发送
	GameVertexSyncDelivery GameVertexBuffersSync = 2
)
