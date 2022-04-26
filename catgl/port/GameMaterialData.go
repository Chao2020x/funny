package port

// GameMaterialData 游戏对象材质信息
type GameMaterialData struct {
	ShaderProgram ID //* 目标着色器
	//? 颜色贴图
	Ambient  Color //* 环境色
	Diffuse  Color //* 漫反射
	Specular Color //* 高光反射
	//? 置换贴图
	BumpMapping         Color //* 凹凸贴图
	NormalMap           Color //* 法线贴图
	ParallaxMapping     Color //* 视差贴图
	DisplacementMapping Color //* 置换贴图
}
