package port

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// GameVertexVAO 顶点数组对象
type GameVertexVAO struct {
	Pointer uint32 // 对象指针
	IsEBO   bool   // 绑定EBO
	First   int32  // 起始位置
	Count   int32  // 顶点数量
}

// Update 更新顶点
func (V *GameVertexVAO) Update() {
	gl.BindVertexArray(V.Pointer)
	if V.IsEBO {
		gl.DrawElements(gl.TRIANGLES, V.Count, gl.UNSIGNED_INT, gl.PtrOffset(int(V.First)))
	} else {
		gl.DrawArrays(gl.TRIANGLES, V.First, V.Count)
	}
	gl.BindVertexArray(0)
}

// Delete 销毁
func (V *GameVertexVAO) Delete() {
	gl.DeleteVertexArrays(1, &V.Pointer)
}

// GameVertexVBO 顶点缓冲区对象
type GameVertexVBO struct {
	Data    []float32
	Pointer uint32
}

// New 创建
func (V *GameVertexVBO) New(T GameVertexlBuffersType) {
	if V.Pointer != 0 {
		return
	}
	var VBO uint32
	gl.GenBuffers(1, &(VBO))                                                 //? 创建 VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)                                      //? 激活缓冲区
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(V.Data), gl.Ptr(V.Data), uint32(T)) //? 添加缓冲区数据
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	V.Pointer = VBO
}

// Delete 销毁
func (V *GameVertexVBO) Delete() {
	gl.DeleteBuffers(1, &V.Pointer)
}

// GameVertexEBO 顶点数组对象
type GameVertexEBO struct {
	Data    []uint32
	Pointer uint32
}

// New 创建
func (E *GameVertexEBO) New(T GameVertexlBuffersType) {
	if E.Pointer != 0 {
		return
	}
	var EBO uint32
	gl.GenBuffers(1, &(EBO))                                                         //? 创建 EBO
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)                                      //? 激活缓冲区
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(E.Data), gl.Ptr(E.Data), uint32(T)) //? 添加缓冲区数据
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	E.Pointer = EBO
}

// Delete 销毁
func (E *GameVertexEBO) Delete() {
	gl.DeleteBuffers(1, &E.Pointer)
}

// GameVertexObject 顶点对象
type GameVertexObject struct {
	GameVertexVAO
	VBO GameVertexVBO
	EBO GameVertexEBO
}

// New 创建
func (V *GameVertexObject) New(T GameVertexlBuffersType, Gvs GameVertexStructure) {
	if V.Pointer != 0 {
		return
	}
	V.VBO.New(T)
	V.Pointer = Gvs.New(V.VBO.Pointer)
	if V.IsEBO {
		// 初始化
		V.EBO.New(T)
		V.Count = int32(len(V.EBO.Data))
		// 绑定对象
		gl.BindVertexArray(V.Pointer)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, V.EBO.Pointer)
		gl.BindVertexArray(0)
	} else {
		V.Count = int32(len(V.VBO.Data)) / Gvs.CountSize
	}
}

// Delete 销毁
func (V *GameVertexObject) Delete() {
	V.VBO.Delete()
	if V.IsEBO {
		V.EBO.Delete()
	}
	V.GameVertexVAO.Delete()
}
