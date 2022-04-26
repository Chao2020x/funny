package port

import (
	"errors"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// ShaderData 着色器数据
type ShaderData struct {
	Program uint32 // 着色器程序
	// 默认着色器
	Vertex   string // 顶点着色器
	Fragment string // 片面着色器
	// 附加着色器
	Geometry   string // 几何着色器
	IsGeometry bool   // 使用几何着色器
}

// Init 初始化
func (S *ShaderData) Init() (err error) {
	var Svertex, Sfragment, Sgeometry uint32
	// 注册流程
	Svertex, err = ShaderCompile(S.Vertex, gl.VERTEX_SHADER)
	if Svertex == 0 {
		return errors.New("顶点着色器编译失败: " + err.Error())
	}
	defer gl.DeleteShader(Svertex)
	Sfragment, err = ShaderCompile(S.Fragment, gl.FRAGMENT_SHADER)
	if Sfragment == 0 {
		return errors.New("片段着色器编译失败: " + err.Error())
	}
	defer gl.DeleteShader(Sfragment)
	if S.IsGeometry {
		Sgeometry, err = ShaderCompile(S.Geometry, gl.GEOMETRY_SHADER)
		if Sgeometry == 0 {
			return errors.New("几何着色器编译失败: " + err.Error())
		}
		defer gl.DeleteShader(Sgeometry)
	}
	S.Program, err = ShaderLinkProgram(Svertex, Sfragment, Sgeometry)
	return err
}

// GetProgram 得到着色器程序
func (S *ShaderData) GetProgram() uint32 {
	return S.Program
}

// Delete 销毁
func (S *ShaderData) Delete() {
	gl.DeleteProgram(S.Program)
}

// ShaderCompile  创建着色器
func ShaderCompile(source string, shaderType uint32) (uint32, error) {
	// 创建着色器
	shader := gl.CreateShader(shaderType)
	// 获得指针
	csource, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csource, nil)
	// 销毁缓存
	free()
	// 编译
	gl.CompileShader(shader)
	// 获得错误
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return shader, GetLog(shader)
	}
	return shader, nil
}

// ShaderLinkProgram 链接着色器程序
func ShaderLinkProgram(Shader ...uint32) (uint32, error) {
	// 着色器程序
	shaderProgram := gl.CreateProgram()
	// 设置
	for _, P := range Shader {
		if P != 0 {
			gl.AttachShader(shaderProgram, P)
		}
	}
	// 链接
	gl.LinkProgram(shaderProgram)
	// 获得错误
	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return shaderProgram, GetLog(shaderProgram)
	}
	return shaderProgram, nil
}
