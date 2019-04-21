package linkedtree

type NodeID uint32
type NodeType uint16
type NodeClassType uint16
type NodeFlag uint8

const (
	NodeFlagRef NodeFlag = 1 << iota
	NodeFlagVariadic
	NodeFlagStatic
	NodeFlagAltSyntax
)

const (
	NodeClassTypeGeneral NodeClassType = 1 << 8
	NodeClassTypeScalar  NodeClassType = 1 << 9
	NodeClassTypeName    NodeClassType = 1 << 10
	NodeClassTypeStmt    NodeClassType = 1 << 11
	NodeClassTypeExpr    NodeClassType = 1 << 12
	NodeClassTypeAssign  NodeClassType = 1<<13 | NodeClassTypeExpr
	NodeClassTypeBinary  NodeClassType = 1<<14 | NodeClassTypeExpr
	NodeClassTypeCast    NodeClassType = 1<<15 | NodeClassTypeExpr
)

//go:generate stringer -type=NodeType -trimprefix=NodeType -output ./nodetype_string.go
const (
	NodeTypeRoot NodeType = 1 | NodeType(NodeClassTypeGeneral)

	NodeTypeIdentifier   NodeType = 2 | NodeType(NodeClassTypeGeneral)
	NodeTypeParameter    NodeType = 3 | NodeType(NodeClassTypeGeneral)
	NodeTypeArgument     NodeType = 4 | NodeType(NodeClassTypeGeneral)
	NodeTypeArgumentList NodeType = 5 | NodeType(NodeClassTypeGeneral)
	NodeTypeNullable     NodeType = 6 | NodeType(NodeClassTypeGeneral)

	NodeTypeNameNamePart       NodeType = 7 | NodeType(NodeClassTypeName)
	NodeTypeNameName           NodeType = 8 | NodeType(NodeClassTypeName)
	NodeTypeNameFullyQualified NodeType = 9 | NodeType(NodeClassTypeName)
	NodeTypeNameRelative       NodeType = 10 | NodeType(NodeClassTypeName)

	NodeTypeScalarEncapsedStringPart NodeType = 11 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarDnumber            NodeType = 12 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarHeredoc            NodeType = 13 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarMagicConstant      NodeType = 14 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarLnumber            NodeType = 15 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarEncapsed           NodeType = 16 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarString             NodeType = 17 | NodeType(NodeClassTypeScalar)

	NodeTypeStmtFinally             NodeType = 18 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtNop                 NodeType = 19 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtUnset               NodeType = 20 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtStatic              NodeType = 21 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtProperty            NodeType = 22 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtGroupUse            NodeType = 23 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltElseIf           NodeType = 24 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtElse                NodeType = 25 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtCase                NodeType = 26 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtConstList           NodeType = 27 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtExpression          NodeType = 28 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtReturn              NodeType = 29 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltForeach          NodeType = 30 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtFor                 NodeType = 31 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassConstList      NodeType = 32 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtInterface           NodeType = 33 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtEcho                NodeType = 34 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtInterfaceExtends    NodeType = 35 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTrait               NodeType = 36 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtContinue            NodeType = 37 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtInlineHtml          NodeType = 38 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassExtends        NodeType = 39 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtUseList             NodeType = 40 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitUse            NodeType = 41 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitUsePrecedence  NodeType = 42 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTry                 NodeType = 43 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtDeclare             NodeType = 44 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtDefault             NodeType = 45 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtForeach             NodeType = 46 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitAdaptationList NodeType = 47 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtGlobal              NodeType = 48 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtWhile               NodeType = 49 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtIf                  NodeType = 50 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltElse             NodeType = 51 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtGoto                NodeType = 52 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClass               NodeType = 53 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtLabel               NodeType = 54 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtThrow               NodeType = 55 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassImplements     NodeType = 56 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtUse                 NodeType = 57 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtPropertyList        NodeType = 58 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltWhile            NodeType = 59 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltIf               NodeType = 60 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtCatch               NodeType = 61 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtConstant            NodeType = 62 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitUseAlias       NodeType = 63 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtNamespace           NodeType = 64 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtSwitch              NodeType = 65 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtBreak               NodeType = 66 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassMethod         NodeType = 67 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtFunction            NodeType = 68 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltSwitch           NodeType = 69 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtHaltCompiler        NodeType = 70 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitMethodRef      NodeType = 71 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtDo                  NodeType = 72 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtStaticVar           NodeType = 73 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtCaseList            NodeType = 74 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtElseIf              NodeType = 75 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltFor              NodeType = 76 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtStmtList            NodeType = 77 | NodeType(NodeClassTypeStmt)

	NodeTypeExprInstanceOf          NodeType = 78 | NodeType(NodeClassTypeExpr)
	NodeTypeExprBooleanNot          NodeType = 79 | NodeType(NodeClassTypeExpr)
	NodeTypeExprNew                 NodeType = 80 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClosureUse          NodeType = 81 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPreDec              NodeType = 82 | NodeType(NodeClassTypeExpr)
	NodeTypeExprEmpty               NodeType = 83 | NodeType(NodeClassTypeExpr)
	NodeTypeExprErrorSuppress       NodeType = 84 | NodeType(NodeClassTypeExpr)
	NodeTypeExprEval                NodeType = 85 | NodeType(NodeClassTypeExpr)
	NodeTypeExprIsset               NodeType = 86 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClassConstFetch     NodeType = 87 | NodeType(NodeClassTypeExpr)
	NodeTypeExprExit                NodeType = 88 | NodeType(NodeClassTypeExpr)
	NodeTypeExprShortList           NodeType = 89 | NodeType(NodeClassTypeExpr)
	NodeTypeExprIncludeOnce         NodeType = 90 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPostInc             NodeType = 91 | NodeType(NodeClassTypeExpr)
	NodeTypeExprConstFetch          NodeType = 92 | NodeType(NodeClassTypeExpr)
	NodeTypeExprTernary             NodeType = 93 | NodeType(NodeClassTypeExpr)
	NodeTypeExprVariable            NodeType = 94 | NodeType(NodeClassTypeExpr)
	NodeTypeExprRequireOnce         NodeType = 95 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPostDec             NodeType = 96 | NodeType(NodeClassTypeExpr)
	NodeTypeExprBitwiseNot          NodeType = 97 | NodeType(NodeClassTypeExpr)
	NodeTypeExprInclude             NodeType = 98 | NodeType(NodeClassTypeExpr)
	NodeTypeExprList                NodeType = 99 | NodeType(NodeClassTypeExpr)
	NodeTypeExprShellExec           NodeType = 100 | NodeType(NodeClassTypeExpr)
	NodeTypeExprMethodCall          NodeType = 101 | NodeType(NodeClassTypeExpr)
	NodeTypeExprReference           NodeType = 102 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPrint               NodeType = 103 | NodeType(NodeClassTypeExpr)
	NodeTypeExprUnaryPlus           NodeType = 104 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPropertyFetch       NodeType = 105 | NodeType(NodeClassTypeExpr)
	NodeTypeExprArray               NodeType = 106 | NodeType(NodeClassTypeExpr)
	NodeTypeExprStaticPropertyFetch NodeType = 107 | NodeType(NodeClassTypeExpr)
	NodeTypeExprUnaryMinus          NodeType = 108 | NodeType(NodeClassTypeExpr)
	NodeTypeExprRequire             NodeType = 109 | NodeType(NodeClassTypeExpr)
	NodeTypeExprShortArray          NodeType = 110 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPreInc              NodeType = 111 | NodeType(NodeClassTypeExpr)
	NodeTypeExprYieldFrom           NodeType = 112 | NodeType(NodeClassTypeExpr)
	NodeTypeExprYield               NodeType = 113 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClone               NodeType = 114 | NodeType(NodeClassTypeExpr)
	NodeTypeExprStaticCall          NodeType = 115 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClosure             NodeType = 116 | NodeType(NodeClassTypeExpr)
	NodeTypeExprArrayDimFetch       NodeType = 117 | NodeType(NodeClassTypeExpr)
	NodeTypeExprArrayItem           NodeType = 118 | NodeType(NodeClassTypeExpr)
	NodeTypeExprFunctionCall        NodeType = 119 | NodeType(NodeClassTypeExpr)

	NodeTypeCastUnset  NodeType = 120 | NodeType(NodeClassTypeCast)
	NodeTypeCastDouble NodeType = 121 | NodeType(NodeClassTypeCast)
	NodeTypeCastString NodeType = 122 | NodeType(NodeClassTypeCast)
	NodeTypeCastBool   NodeType = 123 | NodeType(NodeClassTypeCast)
	NodeTypeCastArray  NodeType = 124 | NodeType(NodeClassTypeCast)
	NodeTypeCastInt    NodeType = 125 | NodeType(NodeClassTypeCast)
	NodeTypeCastObject NodeType = 126 | NodeType(NodeClassTypeCast)

	NodeTypeAssignAssign     NodeType = 127 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignMul        NodeType = 128 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignMinus      NodeType = 129 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignShiftLeft  NodeType = 130 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignMod        NodeType = 131 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignPow        NodeType = 132 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignBitwiseXor NodeType = 133 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignConcat     NodeType = 134 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignDiv        NodeType = 135 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignPlus       NodeType = 136 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignReference  NodeType = 137 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignBitwiseAnd NodeType = 138 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignBitwiseOr  NodeType = 139 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignShiftRight NodeType = 140 | NodeType(NodeClassTypeAssign)

	NodeTypeBinaryBooleanOr      NodeType = 141 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryMul            NodeType = 142 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryMinus          NodeType = 143 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryEqual          NodeType = 144 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryIdentical      NodeType = 145 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryLogicalXor     NodeType = 146 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryCoalesce       NodeType = 147 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBooleanAnd     NodeType = 148 | NodeType(NodeClassTypeBinary)
	NodeTypeBinarySmaller        NodeType = 149 | NodeType(NodeClassTypeBinary)
	NodeTypeBinarySmallerOrEqual NodeType = 150 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryShiftLeft      NodeType = 151 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryMod            NodeType = 152 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryGreaterOrEqual NodeType = 153 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryLogicalAnd     NodeType = 154 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryLogicalOr      NodeType = 155 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryPow            NodeType = 156 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBitwiseXor     NodeType = 157 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryConcat         NodeType = 158 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryDiv            NodeType = 159 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryPlus           NodeType = 160 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryGreater        NodeType = 161 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBitwiseAnd     NodeType = 162 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBitwiseOr      NodeType = 163 | NodeType(NodeClassTypeBinary)
	NodeTypeBinarySpaceship      NodeType = 164 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryShiftRight     NodeType = 165 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryNotIdentical   NodeType = 166 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryNotEqual       NodeType = 167 | NodeType(NodeClassTypeBinary)
)

type EdgeType uint8

//go:generate stringer -type=EdgeType -trimprefix=EdgeType -output ./edgetype_string.go
const (
	EdgeTypeNil EdgeType = iota
	EdgeTypeExpr
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

func (nt NodeType) Is(nct NodeClassType) bool {
	return uint16(nt)&uint16(nct) == uint16(nct)
}

type Node struct {
	Type NodeType
	Flag NodeFlag

	Parent NodeID
	Child  NodeID
	Next   NodeID

	Key EdgeType

	Pos PositionID
}

func (n Node) IsNode() bool {
	return true
}

// NodeStorage store nodes
type NodeStorage struct {
	buf []Node
}

// NewNodeStorage creates new NodeStorage
func NewNodeStorage(buf []Node) *NodeStorage {
	return &NodeStorage{buf}
}

// Reset storage
func (b *NodeStorage) Reset() {
	b.buf = b.buf[:0]
}

// Create saves new Node in store
func (b *NodeStorage) Create(n Node) NodeID {
	b.buf = append(b.buf, n)
	return NodeID(len(b.buf))
}

// Save modified Node
func (b *NodeStorage) Save(id NodeID, n Node) {
	b.buf[id-1] = n
}

// Get returns Node by NodeID
func (b NodeStorage) Get(id NodeID) Node {
	return b.buf[id-1]
}

// GetAll returns all Nodes
func (b NodeStorage) GetAll() []Node {
	return b.buf
}
