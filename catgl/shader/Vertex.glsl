//? 默认数据顶点
layout (location = 0) in vec3 Positions; //* 位置
layout (location = 1) in vec3 Normals;	 //* 法线
layout (location = 2) in vec3 Color;	 //* 颜色
layout (location = 3) in vec2 Uv;		 //* Uv
layout (location = 4) in vec3 Tangent;	 //* 切线
layout (location = 5) in vec3 Bitangent; //* 副切线

//? 引擎数据
layout (std140, binding = 0) uniform GlobalMatrix //* 全局变量
{
    mat4 ProjectionMatrix;     //* 投影矩阵 4*16
    mat4 ViewMatrix;           //* 观察矩阵 4*16
};
uniform mat4 Model;            //* 模型位置(Vertex类)

//? 内部函数
vec4 Position() {
	return ProjectionMatrix * ViewMatrix *  Model * vec4(Positions,1.0);
}