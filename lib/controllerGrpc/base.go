package contents

import (
	contentsGrpc "github.com/hanherb/omdb-api/omdb-generate/contents"
)

var _ contentsGrpc.ContentServiceServer = &Controller{}

type Controller struct {
	contentsGrpc.UnimplementedContentServiceServer
}
