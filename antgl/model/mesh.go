package model

import (
	"strconv"
	"unsafe"

	"github.com/Chao2020x/funny/antgl/shader"
	"github.com/Chao2020x/funny/antgl/texture"
	"github.com/Chao2020x/funny/lib/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Vertex struct {
	Position  mgl32.Vec3
	Normal    mgl32.Vec3
	TexCoords mgl32.Vec2
	Tangent   mgl32.Vec3
	Bitangent mgl32.Vec3
}

type MeshTexture struct {
	ID          uint32
	Type        string
	Path        string
	BaseTexture *texture.TextureObject
}

type Mesh struct {
	Vertices []*Vertex
	Indices  []uint32
	Textures []*MeshTexture
	VAO      uint32
	VBO      uint32
	EBO      uint32
}

func NewMesh() *Mesh {
	m := &Mesh{}
	m.setupMesh()
	return m
}

func (m *Mesh) Draw(shader111 *shader.ShaderObject) {
	diffuseNr, specularNr, normalNr, heightNr := 1, 1, 1, 1
	for i := range m.Textures {
		gl.ActiveTexture(gl.TEXTURE0 + uint32(i))
		var (
			number string
			name   = m.Textures[i].Type
		)
		if name == "texture_diffuse" {
			diffuseNr++
			number = strconv.Itoa(diffuseNr)
		} else if name == "texture_specular" {
			specularNr++
			number = strconv.Itoa(specularNr)
		} else if name == "texture_normal" {
			normalNr++
			number = strconv.Itoa(normalNr)
		} else if name == "texture_height" {
			heightNr++
			number = strconv.Itoa(heightNr)
		}
		shader111.SetInt(name+number, int32(i))
		gl.BindTexture(gl.TEXTURE_2D, m.Textures[i].ID)
	}

	gl.BindVertexArray(m.VAO)
	gl.DrawElements(gl.TRIANGLES, int32(len(m.Indices)), gl.UNSIGNED_INT, unsafe.Pointer(&m.Indices))
	gl.BindVertexArray(0)
	gl.ActiveTexture(gl.TEXTURE0)
}

func (m *Mesh) setupMesh() {
	gl.GenVertexArrays(1, &m.VAO)
	gl.GenBuffers(1, &m.VBO)
	gl.GenBuffers(1, &m.EBO)
	gl.BindVertexArray(m.VAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.VBO)

	gl.BufferData(gl.ARRAY_BUFFER, len(m.Vertices)*int(unsafe.Sizeof(Vertex{})), gl.Ptr(m.Vertices[0]), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.Indices)*int(unsafe.Sizeof(uint32(0))), gl.Ptr(m.Indices[0]), gl.STATIC_DRAW)

	var stride int32 = 3*4 + 3*4 + 2*4 + 3*4 + 3*4

	var offset int = 0
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))

	offset += 3 * 4
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))

	offset += 2 * 4
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))

	offset += 3 * 4
	gl.EnableVertexAttribArray(3)
	gl.VertexAttribPointer(3, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))

	offset += 3 * 4
	gl.EnableVertexAttribArray(4)
	gl.VertexAttribPointer(4, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))

	gl.BindVertexArray(0)

}
