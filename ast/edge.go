package ast

type EdgeID uint32
type EdgeType uint8

const (
	EdgeTypeExpr EdgeType = iota
	EdgeTypeVarType
	EdgeTypeVar
	EdgeTypeDefaultValue
	EdgeTypeArguments
	EdgeTypeStmts
	EdgeTypeParts
	EdgeTypeCond
	EdgeTypeStmt
	EdgeTypeElseIf
	EdgeTypeElse
	EdgeTypeTypes
	EdgeTypeModifiers
	EdgeTypeConsts
	EdgeTypeMethodName
	EdgeTypeParams
	EdgeTypeReturnType
	EdgeTypeClassName
	EdgeTypeArgumentList
	EdgeTypeExtends
	EdgeTypeImplements
	EdgeTypeConstantName
	EdgeTypeExprs
	EdgeTypeInit
	EdgeTypeLoop
	EdgeTypeKey
	EdgeTypeFunctionName
	EdgeTypeVars
	EdgeTypeLabel
	EdgeTypeUseType
	EdgeTypePrefix
	EdgeTypeUseList
	EdgeTypeInterfaceName
	EdgeTypeLabelName
	EdgeTypeNamespaceName
	EdgeTypeProperties
	EdgeTypeCaseList
	EdgeTypeTrait
	EdgeTypeMethod
	EdgeTypeRef
	EdgeTypeModifier
	EdgeTypeAlias
	EdgeTypeInsteadof
	EdgeTypeTraits
	EdgeTypeTraitAdaptationList
	EdgeTypeTraitName
	EdgeTypeCatches
	EdgeTypeFinally
	EdgeTypeUses
	EdgeTypeUse
	EdgeTypeCases
	EdgeTypeAdaptations
	EdgeTypeInterfaceNames
	EdgeTypeLeft
	EdgeTypeRight
	EdgeTypeDim
	EdgeTypeVal
	EdgeTypeItems
	EdgeTypeClass
	EdgeTypeClosureUse
	EdgeTypeConstant
	EdgeTypeFunction
	EdgeTypeProperty
	EdgeTypeCall
	EdgeTypeIfTrue
	EdgeTypeIfFalse
	EdgeTypeVarName
)

// Edge represent relation between nodes
type Edge struct {
	Type     EdgeType
	From, To NodeID
}

//EdgeStorage stores edges
type EdgeStorage struct {
	buf []Edge
}

// NewEdgeStorage creates new EdgeStorage
func NewEdgeStorage(buf []Edge) *EdgeStorage {
	return &EdgeStorage{buf}
}

// Reset storage
func (b *EdgeStorage) Reset() {
	b.buf = b.buf[:0]
}

// Create new Edge in store
func (b *EdgeStorage) Create(e Edge) EdgeID {
	b.buf = append(b.buf, e)
	return EdgeID(len(b.buf))
}

// Save modified Edge
func (b *EdgeStorage) Save(id EdgeID, e Edge) {
	b.buf[id-1] = e
}

// Get returns Edge by EdgeID
func (b *EdgeStorage) Get(id EdgeID) Edge {
	return b.buf[id-1]
}

// GetAll returns all Edges
func (b *EdgeStorage) GetAll() []Edge {
	return b.buf
}

// Chidren creates edges from parent to children nodes with specified EdgeType
func (b *EdgeStorage) Children(parent NodeID, edgeType EdgeType, children ...NodeID) {
	for _, child := range children {
		if child != 0 {
			b.Create(Edge{
				Type: edgeType,
				From: parent,
				To:   child,
			})
		}
	}
}
