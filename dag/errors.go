package dag

import "github.com/pkg/errors"

const (
	ErrorFormat         = "%w nodeId %s"
	ErrorFormatTwoNodes = "%w nodeId %s %s"
)

var (
	NodeAlreadyExistsError    = errors.Errorf("error : node already exists")
	NodeNotFound              = errors.Errorf("error : node not found")
	NodeRelationNotFound      = errors.Errorf("error : node relation not present")
	CyclicDependencyError     = errors.Errorf("error : can not create cyclic dependency")
	AncestorsComputationError = errors.Errorf("error : ancestors could not be computed ")
)
