//? 全局灯光对象
struct lightObject {
	float PositionX,PositionY,PositionZ; 		// 位置
	//* 颜色
	float AmbientX,AmbientY,AmbientZ;			// 环境色
	float DiffuseX,DiffuseY,DiffuseZ;  			// 漫反射
	float SpecularX,SpecularY,SpecularZ;		// 高光反射
	//? 物理参数
	float Linear;       						// 线性
	float Constant;     						// 常量
	float Quadratic;    						// 二次方程式
	// 灯光是否生效
	bool  IsLight; 
};

//? 全局灯光
layout (std140,binding = 1) uniform GlobalLight //* 全局变量
{
	lightObject []Light;
};