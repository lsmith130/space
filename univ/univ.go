package univ

import (
	"fmt"
	"log"
	"time"
	"unsafe"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lsmith130/space/draw"
	assimp "github.com/tbogdala/assimp-go"
)

// DefaultRefreshRate is the default refresh rate
const DefaultRefreshRate = time.Millisecond * 16

// Universe is a group of Bodies drawn and updated together. It is the base object of
// the univ package, and all Bodies are created in a Universe.
type Universe struct {
	// bodies is a set of bodies
	bodies map[*Body]struct{}
	window *draw.Window
}

// NewUniverse constructs a new empty Universe
func NewUniverse(window *draw.Window, updateRate time.Duration) *Universe {

	u := &Universe{
		bodies: make(map[*Body]struct{}),
		window: window,
	}

	return u
}

// NewBody constructs a new body in u with a given model and shader
func (u *Universe) NewBody(modelPath string, programType draw.ProgramType, textures []*draw.Texture) (*Body, error) {

	meshes, err := assimp.ParseFile(modelPath)
	if err != nil {
		return nil, fmt.Errorf("load model %s: %v", modelPath, err)
	}
	if len(textures) > 0 && len(textures) != len(meshes) {
		return nil, fmt.Errorf("%d textures dosen't match %d meshes", len(textures), len(meshes))
	}

	body := &Body{
		meshes:   make([]*draw.Mesh, len(meshes)),
		rotation: mgl32.QuatIdent(),
		program:  u.window.GetProgram(programType),
	}
	// log.Println(modelPath)
	for i, mesh := range meshes {
		log.Printf("%+v", mesh)
		faces := *(*[]draw.MeshFace)(unsafe.Pointer(&mesh.Faces))
		body.meshes[i] = body.program.NewMesh(mesh.Vertices, faces, mesh.UVChannels[0], mesh.Normals)
		body.meshes[i].SetTexture(textures[i])
	}

	body.program.AddDrawable(body)
	u.bodies[body] = struct{}{}

	return body, nil
}

// RemoveBody removes a body from u, such that it will no longer be drawn or recieve updates
func (u *Universe) RemoveBody(body *Body) {
	body.program.RemoveDrawable(body)
	delete(u.bodies, body)
}
