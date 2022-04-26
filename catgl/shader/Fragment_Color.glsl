//! 颜色结构体
struct Color {
    vec3  Color;     //* 颜色
    int   Texture;   //* 绑定贴图(索引)
	float Strength;  //* 强度
};

//? 颜色贴图
uniform Color Ambient;   //* 环境色
uniform Color Diffuse;   //* 漫反射
uniform	Color Specular;  //* 高光反射
//? 置换贴图
uniform	Color BumpMapping;          //* 凹凸贴图
uniform	Color NormalMap;            //* 法线贴图
uniform	Color ParallaxMapping;      //* 视差贴图
uniform	Color DisplacementMapping;  //* 置换贴图
//? 深度贴图
uniform	Color DepthMap; 