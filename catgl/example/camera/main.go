package main

import (
	"fmt"

	"github.com/Chao2020x/funny/catgl"
	"github.com/Chao2020x/funny/catgl/port"
	"github.com/Chao2020x/funny/lib/mgl32"
)

func main() {
	// 初始化
	catgl.NOScreenRender()
	// 创建场景
	Scene := &catgl.Scene{}
	// 场景默认对象
	Scene.Light = &port.LightData{Quantity: 1}
	// 设置相机
	SceneCamera(Scene)
	// 渲染窗口
	R := catgl.Renderer{Width: 500, Height: 500, Title: "测试窗口", Scene: Scene}
	// 创建窗口
	err := R.New()
	if errs := <-err; errs != nil {
		fmt.Println("窗口初始化错误:", errs)
		return
	}
	// 游戏元素初始化
	Game(Scene)
	// 等待结束
	fmt.Println("程序关闭", <-err)
}

// Game 游戏对象
func Game(Scene *catgl.Scene) {
	// 默认着色器
	shader := &port.ShaderData{
		Vertex: `
		#version 420 core
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
			return ProjectionMatrix * ViewMatrix * vec4(Positions,1.0);
		}

		out vec3 ourColor; 
		out vec2 TexCoord;

		void main()
		{
			gl_Position = Position();
			
			ourColor = Color; 
			TexCoord = Uv;
		}
		`,
		Fragment: `
		#version 420 core
		out vec4 color;

		in vec3 ourColor;
		in vec2 TexCoord; 

		uniform sampler2D Texture[31];
		uniform bool IsTexture[31];

		void main()
		{
			if (IsTexture[0]) {
				color = texture(Texture[0], TexCoord) * vec4(ourColor, 1.0);
			} else {
				color = vec4(ourColor, 1.0);
			}
		}
		`,
	}
	fmt.Println("编译着色器", Scene.AddObject(shader))
	Scene.ShaderObjectList["默认着色器"] = shader
	// 默认游戏对象
	game := &GameObject{
		ShaderProgram: shader.Program,
		Vertices: []float32{
			-0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0,
			0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0,
			0.0, 0.5, 0.0, 0.0, 0.0, 1.0, 0.5, 1.0,
		},
		Texture: &port.TextureData{},
	}
	game.Texture.New("./UV.jpg", 0, port.TEXTURE2D)
	fmt.Println("初始化模型", Scene.AddObject(game))
	fmt.Println("初始化材质", Scene.AddObject(game.Texture))
	Scene.GameObjectList["三角形"] = game
}

// SceneCamera 场景相机
func SceneCamera(Scene *catgl.Scene) {
	Camera := &port.CameraData{}
	// 设置相机参数
	Camera.Up = mgl32.Vec3{0, 1, 0}     // 朝向
	Camera.Eye = mgl32.Vec3{0, 0, 3.0}  // 位置
	Camera.Center = mgl32.Vec3{0, 0, 0} // 视点
	Camera.Fovy = 45                    // 视角
	Camera.Aspect = 1                   // 屏幕宽高比
	Camera.ZNear = 1                    // 近平面距离
	Camera.ZFar = 100                   // 远平面距离
	Scene.Camera = Camera
}
