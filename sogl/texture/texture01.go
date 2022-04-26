package texture

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type TextureObject struct {
	*TextureData
}

// * File   材质文件
// * UnitID 文理单元, 0~31 号可用
// * Target 纹理类型
func NewTextureFromFile(file string, UnitID uint32, Target uint32) (*TextureObject, error) {

	textureXXX := &TextureObject{&TextureData{}}
	var err = textureXXX.New(file, UnitID, Target)
	if err != nil {
		return nil, err
	}
	err = textureXXX.Init()
	return textureXXX, err
}

func (tex *TextureObject) Bind(texUnit uint32) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(tex.Target, tex.ID)
	tex.UnitID = texUnit
}

func (tex *TextureObject) UnBind() {
	tex.UnitID = 0
	gl.BindTexture(tex.Target, 0)
}

func (tex *TextureObject) SetUniform(uniformLoc int32) error {
	if tex.UnitID == 0 {
		return fmt.Errorf("texture not bound")
	}
	gl.Uniform1i(uniformLoc, int32(tex.UnitID-gl.TEXTURE0))
	return nil
}
