package texture

///
// 文理数据加载包
//  材质加载
// ? 日志
// !  2019-8-17 重写
///
import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"path"
	"strconv"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/Chao2020x/funny/lib/image/bmp"
	"github.com/Chao2020x/funny/lib/image/tga"
)

// 纹理变量
const (
	// 纹理单元
	TEXTURE   = 0x1702
	TEXTURE0  = 0x84C0
	TEXTURE1  = 0x84C1
	TEXTURE10 = 0x84CA
	TEXTURE11 = 0x84CB
	TEXTURE12 = 0x84CC
	TEXTURE13 = 0x84CD
	TEXTURE14 = 0x84CE
	TEXTURE15 = 0x84CF
	TEXTURE16 = 0x84D0
	TEXTURE17 = 0x84D1
	TEXTURE18 = 0x84D2
	TEXTURE19 = 0x84D3
	TEXTURE2  = 0x84C2
	TEXTURE20 = 0x84D4
	TEXTURE21 = 0x84D5
	TEXTURE22 = 0x84D6
	TEXTURE23 = 0x84D7
	TEXTURE24 = 0x84D8
	TEXTURE25 = 0x84D9
	TEXTURE26 = 0x84DA
	TEXTURE27 = 0x84DB
	TEXTURE28 = 0x84DC
	TEXTURE29 = 0x84DD
	TEXTURE3  = 0x84C3
	TEXTURE30 = 0x84DE
	TEXTURE31 = 0x84DF
	TEXTURE4  = 0x84C4
	TEXTURE5  = 0x84C5
	TEXTURE6  = 0x84C6
	TEXTURE7  = 0x84C7
	TEXTURE8  = 0x84C8
	TEXTURE9  = 0x84C9
	// 纹理类型
	TEXTURE1D                 = 0x0DE0
	TEXTURE1DARRAY            = 0x8C18
	TEXTURE2D                 = 0x0DE1
	TEXTURE2DARRAY            = 0x8C1A
	TEXTURE2DMULTISAMPLE      = 0x9100
	TEXTURE2DMULTISAMPLEARRAY = 0x9102
	TEXTURE3D                 = 0x806F
)

// ImageDecode 图片解码
var ImageDecode map[string]func(r io.Reader) (image.Image, error)

func init() {
	ImageDecode = make(map[string]func(r io.Reader) (image.Image, error))
	ImageDecode[".png"] = png.Decode
	ImageDecode[".jpg"] = jpeg.Decode
	ImageDecode[".tga"] = tga.Decode
	ImageDecode[".bmp"] = bmp.Decode
}

// TextureData 纹理数据结构体
type TextureData struct {
	// 内部用
	ID     uint32 //? 文理ID
	Name   *uint8 //? 参数名称(材质)
	IsName *uint8 //? 参数名称(标记)
	UnitID uint32 //? 纹理单元
	Target uint32 //? 纹理类型
	// 外部信息
	File      string //? 纹理路径
	IsTexture bool   //? 是否生效
}

// Init 初始化
func (T *TextureData) Init() error {
	// 初始化变量
	File, UnitID, Target := T.File, T.UnitID, T.Target
	T.Name = gl.Str("Texture[" + strconv.FormatUint(uint64(UnitID), 10) + "]\x00")
	T.IsName = gl.Str("IsTexture[" + strconv.FormatUint(uint64(UnitID), 10) + "]\x00")
	// 读文件
	imgdata, err := ioutil.ReadFile(File)
	if err != nil {
		return fmt.Errorf("材质文件 %q 打开失败: %v", File, err)
	}
	// 解码图片
	var img image.Image
	stamp := path.Ext(File) //? 得到文件名后戳
	if imgD, ok := ImageDecode[stamp]; ok {
		img, err = imgD(bytes.NewReader(imgdata))
		if err != nil {
			return fmt.Errorf("图片解码失败: " + err.Error())
		}
	} else {
		return fmt.Errorf("未知图片格式")
	}
	// 得到图片通道信息
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return fmt.Errorf("材质大小不支持")
	}
	// 转换格式
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	// 创建纹理
	gl.GenTextures(1, &T.ID)
	gl.BindTexture(Target, T.ID)
	// 纹理参数
	gl.TexParameteri(Target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(Target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(Target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(Target, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	// 添加纹理
	gl.TexImage2D(Target, 0, gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	// 解除纹理
	gl.BindTexture(Target, 0)
	// 设置生效
	T.IsTexture = true
	return nil
}

// Delete 销毁
func (T *TextureData) Delete() {
	gl.DeleteTextures(1, &T.ID)
}

// New 创建
// * File   材质文件
// * UnitID 文理单元, 0~31 号可用
// * Target 纹理类型
func (T *TextureData) New(File string, UnitID uint32, Target uint32) error {
	if UnitID > 32 {
		return fmt.Errorf("材质索引错误")
	}
	// 记录变量
	T.Target = Target
	T.UnitID = UnitID
	T.File = File
	return nil
}

// Update 更新
//? Program 着色器
func (T *TextureData) Update(Program uint32) {
	if T.IsTexture {
		// 设置材质
		gl.ActiveTexture(TEXTURE0 + T.UnitID)                                 //? 激活纹理单元
		gl.BindTexture(T.Target, T.ID)                                        //? 绑定纹理
		gl.Uniform1i(gl.GetUniformLocation(Program, T.Name), int32(T.UnitID)) //? 设置材质
		// 设置材质参数
		gl.Uniform1i(gl.GetUniformLocation(Program, T.IsName), 1) //? 设置材质是否生效
	} else {
		gl.Uniform1i(gl.GetUniformLocation(Program, T.IsName), 0) //? 设置材质是否生效
	}
}
