package port

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// GameVertexStructure 顶点结构定义
//  设定着色器顶点默认启用
type GameVertexStructure struct {
	CountSize int32 // 单个顶点占用元素数量大小
	// 设置信息
	IsPositions bool // 位置
	IsNormals   bool // 法线
	IsColor     bool // 颜色
	IsUv        bool // Uv
	IsTangent   bool // 切线
	IsBitangent bool // 副切线
}

// New 创建VAO对象
//  VBO 顶点缓冲区
func (G *GameVertexStructure) New(VBO uint32) (VAO uint32) {
	gl.GenVertexArrays(1, &(VAO))
	gl.BindVertexArray(VAO)
	// 绑定数据
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	// 设置结构
	{
		// 得到占用大小
		var posltion int32
		var getSize = func(b bool, s int32) {
			if b {
				posltion += s
			}
		}
		getSize(G.IsPositions, 12)
		getSize(G.IsNormals, 12)
		getSize(G.IsColor, 12)
		getSize(G.IsUv, 8)
		getSize(G.IsTangent, 12)
		getSize(G.IsBitangent, 12)
		// 设置结构
		var offset int
		var setVertex = func(b bool, l uint32, s int32) {
			if b {
				// 绑定索引 元素数量 元素类型 不标准化 结构大小 开始偏移
				gl.VertexAttribPointer(l, s, gl.FLOAT, false, posltion, gl.PtrOffset(offset)) //- 设置顶点结构
				gl.EnableVertexAttribArray(l)                                                 //- 设置顶点layout索引
				offset += int(s) * 4
				G.CountSize += s
			}
		}
		G.CountSize = 0
		setVertex(G.IsPositions, 0, 3)
		setVertex(G.IsNormals, 1, 3)
		setVertex(G.IsColor, 2, 3)
		setVertex(G.IsUv, 3, 2)
		setVertex(G.IsTangent, 4, 3)
		setVertex(G.IsBitangent, 5, 3)
	}
	// 解除绑定
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	return VAO
}
