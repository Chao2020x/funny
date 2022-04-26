package camera

import (
	"github.com/Chao2020x/funny/lib/mgl32"
)

const (
	YAW         = -90.0
	PITCH       = 0.0
	SPEED       = 2.5
	SENSITIVITY = 0.1
	ZOOM        = 45.0
)

type Camera_Movement int

const (
	FORWARD Camera_Movement = iota
	BACKWARD
	LEFT
	RIGHT
)

//
type CameraObject struct {
	Position mgl32.Vec3 //摄像头位置
	Forward  mgl32.Vec3 //摄像头方向
	Up       mgl32.Vec3 //摄像头垂直上方向
	Right    mgl32.Vec3 //摄像头右方向
	WorldUp  mgl32.Vec3 //正上方向

	Yaw   float32
	Pitch float32

	MovementSpeed    float32 //移动速度
	MouseSensitivity float32 //鼠标灵敏度
	Zoom             float32 //放缩

	Flip_y bool
}

//GetCamera Camera的构造函数
//pos=mgl32.Vec3{0.0,0.0,0.0}
//默认
//up=mgl32.Vec3{0.0,1.0,0.0}
//yaw=YAW,pitch=PITCH float32
func NewCameraObject(pos mgl32.Vec3) *CameraObject {
	c := &CameraObject{
		Position:         pos,
		Forward:          mgl32.Vec3{0.0, 0.0, -1.0},
		Up:               mgl32.Vec3{},
		Right:            mgl32.Vec3{},
		WorldUp:          mgl32.Vec3{0.0, 1.0, 0.0},
		Yaw:              YAW,
		Pitch:            PITCH,
		MovementSpeed:    SPEED,
		MouseSensitivity: SENSITIVITY,
		Zoom:             ZOOM,

		Flip_y: false,
	}
	c.updateCameraVectors()
	return c
}

func (c *CameraObject) GetViewMatrix() mgl32.Mat4 {

	return mgl32.LookAtV(c.Position, c.Position.Add(c.Forward), c.Up)
}

//ProcessKeyboard 对应键盘移动事件
func (c *CameraObject) ProcessKeyboard(direction Camera_Movement, deltaTime float64) {
	velocity := c.MovementSpeed * float32(deltaTime)
	switch direction {
	case FORWARD:
		c.Position = c.Position.Add(c.Forward.Mul(velocity))
	case BACKWARD:
		c.Position = c.Position.Sub(c.Forward.Mul(velocity))
	case LEFT:
		c.Position = c.Position.Sub(c.Right.Mul(velocity))
	case RIGHT:
		c.Position = c.Position.Add(c.Right.Mul(velocity))
	}
}

//ProcessMouseMovement 对应鼠标移动事件
func (c *CameraObject) ProcessMouseMovement(xoffset, yoffset float64, constrainPitch bool) {
	xoffset *= float64(c.MouseSensitivity)
	yoffset *= float64(c.MouseSensitivity)

	c.Yaw += float32(xoffset)
	c.Pitch += float32(yoffset)

	if constrainPitch {
		if c.Pitch > 89.0 {
			c.Pitch = 89.0
		} else if c.Pitch < -89.0 {
			c.Pitch = -89.0
		}
	}
	c.updateCameraVectors()
}

//ProcessMouseScroll 对应鼠标滚轮事件
func (c *CameraObject) ProcessMouseScroll(yoffset float64) {
	if c.Zoom >= 1.0 && c.Zoom <= 45.0 {
		c.Zoom -= float32(yoffset)
	}
	if c.Zoom <= 1.0 {
		c.Zoom = 1.0
	}
	if c.Zoom >= 45.0 {
		c.Zoom = 45.0
	}
}

// updateCameraVectors 更新摄像机对应的向量
func (c *CameraObject) updateCameraVectors() {

	x := mgl32.Cos(mgl32.DegToRad(c.Yaw)) * mgl32.Cos(mgl32.DegToRad(c.Pitch))
	y := mgl32.Sin(mgl32.DegToRad(c.Pitch))
	z := mgl32.Sin(mgl32.DegToRad(c.Yaw)) * mgl32.Cos(mgl32.DegToRad(c.Pitch))

	var front = mgl32.Vec3{x, y, z}
	c.Forward = front.Normalize()

	c.Right = c.Forward.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Forward).Normalize()
}
