package model

import (
	"fmt"
	"strings"

	"github.com/Chao2020x/funny/lib/assimp"
	"github.com/Chao2020x/funny/sogl/shader"
	"github.com/Chao2020x/funny/sogl/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type ModelObject struct {
	Textures_loaded []*MeshTexture
	Meshes          []*Mesh
	Directory       string
	GammaCorrection bool
}

func NewModelObject(path string, gamma bool) *ModelObject {

	m := new(ModelObject)
	m.Textures_loaded = make([]*MeshTexture, 0)
	m.Meshes = make([]*Mesh, 0)
	m.GammaCorrection = gamma
	m.LoadModel(path)
	return m
}

func (m *ModelObject) LoadModel(path string) error {
	var scene11 = assimp.ImportFile(path, assimp.Process_Triangulate|assimp.Process_FlipUVs|assimp.Process_CalcTangentSpace)

	if scene11 == nil || (scene11.Flags() & assimp.SceneFlags_Incomplete) || (scene11.RootNode() == nil) {
		return fmt.Errorf(" ERROR::ASSIMP:: " + assimp.GetErrorString())
	}
	index := strings.LastIndex(m.Directory, "/")
	m.Directory = m.Directory[:index]
	return nil
}

func (m *ModelObject) Draw(shader111 *shader.ShaderObject) {
	for i := range m.Meshes {
		m.Meshes[i].Draw(shader111)
	}
}

func (m *ModelObject) processNode(aiNode *assimp.Node, aiScene *assimp.Scene) {

	var mNumMeshes11 = aiNode.NumMeshes()
	var SceneMeshSlice = aiScene.Meshes()
	var NodeMeshSlice = aiNode.Meshes()
	for i := 0; i < mNumMeshes11; i++ {
		var mesh = SceneMeshSlice[NodeMeshSlice[i]]
		m.Meshes = append(m.Meshes, m.processMesh(mesh, aiScene))
	}
	var mNumChildren = aiNode.NumChildren()
	var childrenSlice = aiNode.Children()
	for i := 0; i < mNumChildren; i++ {
		m.processNode(childrenSlice[i], aiScene)
	}
}

func (m *ModelObject) processMesh(aiMesh *assimp.Mesh, aiScene *assimp.Scene) *Mesh {

	var (
		vertices []*Vertex
		indices  []uint32
		textures []*MeshTexture

		mNumVertices      = aiMesh.NumVertices()
		MeshVerticesSlice = aiMesh.Vertices()
		MeshNormalsSlice  = aiMesh.Normals()
	)

	// 顶点&&索引

	for i := 0; i < mNumVertices; i++ {
		vertex := &Vertex{}

		vertex.Position[0] = MeshVerticesSlice[i].X()
		vertex.Position[1] = MeshVerticesSlice[i].Y()
		vertex.Position[2] = MeshVerticesSlice[i].Z()

		vertex.Normal[0] = MeshNormalsSlice[i].X()
		vertex.Normal[1] = MeshNormalsSlice[i].Y()
		vertex.Normal[2] = MeshNormalsSlice[i].Z()

		var MeshTextureCoordsSlice = aiMesh.TextureCoords(0)
		if len(MeshTextureCoordsSlice) != 0 {
			vertex.TexCoords[0] = MeshTextureCoordsSlice[i].X()
			vertex.TexCoords[1] = MeshTextureCoordsSlice[i].Y()
		} else {
			vertex.TexCoords[0] = 0.0
			vertex.TexCoords[1] = 0.0
		}

		var MeshTangentsSlice = aiMesh.Tangents()
		vertex.Tangent[0] = MeshTangentsSlice[i].X()
		vertex.Tangent[1] = MeshTangentsSlice[i].Y()
		vertex.Tangent[2] = MeshTangentsSlice[i].Z()

		var MeshBitangentsSlice = aiMesh.Bitangents()
		vertex.Bitangent[0] = MeshBitangentsSlice[i].X()
		vertex.Bitangent[1] = MeshBitangentsSlice[i].Y()
		vertex.Bitangent[2] = MeshBitangentsSlice[i].Z()

		vertices = append(vertices, vertex)
	}

	var meshNumFaces = aiMesh.NumFaces()
	var meshFacesSlice = aiMesh.Faces()
	for i := 0; i < meshNumFaces; i++ {
		var face = meshFacesSlice[i]
		var mNumIndices = face.NumIndices()
		var faceIndicesSlice = face.CopyIndices()
		for j := uint32(0); j < mNumIndices; j++ {
			indices = append(indices, faceIndicesSlice[j])
		}
	}

	// 材质
	material := aiScene.Materials()[aiMesh.MaterialIndex()]

	diffuseMaps := m.loadMaterialTextures(material, assimp.TextureMapping_Diffuse, "texture_diffuse")
	textures = append(textures, diffuseMaps...)

	specularMaps := m.loadMaterialTextures(material, assimp.TextureMapping_Specular, "texture_specular")
	textures = append(textures, specularMaps...)

	normalMaps := m.loadMaterialTextures(material, assimp.TextureMapping_Height, "texture_normal")
	textures = append(textures, normalMaps...)

	heightMaps := m.loadMaterialTextures(material, assimp.TextureMapping_Ambient, "texture_height")
	textures = append(textures, heightMaps...)

	return &Mesh{
		Vertices: vertices,
		Indices:  indices,
		Textures: textures,
	}
}

func (m *ModelObject) loadMaterialTextures(aiMat *assimp.Material, aiType *assimp.TextureType, typeName string) []*MeshTexture {

	var textures []*MeshTexture = make([]*MeshTexture, 0)
	var matTectureCount = aiMat.GetMaterialTextureCount(*aiType)
	for i := 0; i < matTectureCount; i++ {
		str11, _, _, _, _, _, _, _ := aiMat.GetMaterialTexture(*aiType, i)
		skip := false

		var textureLoadSize = len(m.Textures_loaded)
		for j := 0; j < textureLoadSize; j++ {
			if m.Textures_loaded[j].Path == str11 {
				textures = append(textures, m.Textures_loaded[j])
				skip = true
				break
			}
		}

		//优化方案：如果纹理还没有被加载，则加载它
		if !skip {
			var texture11, err = texture.NewTextureFromFile(m.Directory+str11, 0, gl.TEXTURE_2D)
			if err == nil {
				meshtexture := &MeshTexture{Type: typeName,
					BaseTexture: texture11,
				}
				meshtexture.ID = meshtexture.BaseTexture.ID
				m.Textures_loaded = append(m.Textures_loaded, meshtexture)
			}
		}

	}

	return textures
}
