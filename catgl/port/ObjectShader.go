package port

import (
	"github.com/Chao2020x/funny/lib/mgl32"

	"github.com/go-gl/gl/v4.1-core/gl"
)

///
// 着色器数据包
// ? 日志
//! 2019年11月1日 重写
///

// Color 颜色
type Color struct {
	Color    mgl32.Vec3 //* 颜色
	Texture  int32      //* 绑定贴图(索引)
	Strength float32    //* 强度
}

// UpdateShader 更新颜色
func (color *Color) UpdateShader(Program uint32, Name string) {
	gl.Uniform3fv(gl.GetUniformLocation(Program, gl.Str(Name+".Color")), 1, &color.Color[0])
	gl.Uniform1i(gl.GetUniformLocation(Program, gl.Str(Name+".Texture")), color.Texture)
	gl.Uniform1f(gl.GetUniformLocation(Program, gl.Str(Name+".Strength")), color.Strength)
}
