package ast

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

const (
	NodeTypeRoot = 1 | NodeType(NodeClassTypeGeneral)

	NodeTypeIdentifier   = 2 | NodeType(NodeClassTypeGeneral)
	NodeTypeParameter    = 3 | NodeType(NodeClassTypeGeneral)
	NodeTypeArgument     = 4 | NodeType(NodeClassTypeGeneral)
	NodeTypeArgumentList = 5 | NodeType(NodeClassTypeGeneral)
	NodeTypeNullable     = 6 | NodeType(NodeClassTypeGeneral)

	NodeTypeNameNamePart       = 7 | NodeType(NodeClassTypeName)
	NodeTypeNameName           = 8 | NodeType(NodeClassTypeName)
	NodeTypeNameFullyQualified = 9 | NodeType(NodeClassTypeName)
	NodeTypeNameRelative       = 10 | NodeType(NodeClassTypeName)

	NodeTypeScalarEncapsedStringPart = 11 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarDnumber            = 12 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarHeredoc            = 13 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarMagicConstant      = 14 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarLnumber            = 15 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarEncapsed           = 16 | NodeType(NodeClassTypeScalar)
	NodeTypeScalarString             = 17 | NodeType(NodeClassTypeScalar)

	NodeTypeStmtFinally             = 18 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtNop                 = 19 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtUnset               = 20 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtStatic              = 21 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtProperty            = 22 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtGroupUse            = 23 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltElseIf           = 24 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtElse                = 25 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtCase                = 26 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtConstList           = 27 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtExpression          = 28 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtReturn              = 29 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltForeach          = 30 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtFor                 = 31 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassConstList      = 32 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtInterface           = 33 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtEcho                = 34 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtInterfaceExtends    = 35 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTrait               = 36 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtContinue            = 37 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtInlineHtml          = 38 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassExtends        = 39 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtUseList             = 40 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitUse            = 41 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitUsePrecedence  = 42 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTry                 = 43 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtDeclare             = 44 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtDefault             = 45 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtForeach             = 46 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitAdaptationList = 47 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtGlobal              = 48 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtWhile               = 49 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtIf                  = 50 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltElse             = 51 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtGoto                = 52 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClass               = 53 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtLabel               = 54 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtThrow               = 55 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassImplements     = 56 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtUse                 = 57 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtPropertyList        = 58 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltWhile            = 59 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltIf               = 60 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtCatch               = 61 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtConstant            = 62 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitUseAlias       = 63 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtNamespace           = 64 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtSwitch              = 65 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtBreak               = 66 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtClassMethod         = 67 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtFunction            = 68 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltSwitch           = 69 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtHaltCompiler        = 70 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtTraitMethodRef      = 71 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtDo                  = 72 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtStaticVar           = 73 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtCaseList            = 74 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtElseIf              = 75 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtAltFor              = 76 | NodeType(NodeClassTypeStmt)
	NodeTypeStmtStmtList            = 77 | NodeType(NodeClassTypeStmt)

	NodeTypeExprInstanceOf          = 78 | NodeType(NodeClassTypeExpr)
	NodeTypeExprBooleanNot          = 79 | NodeType(NodeClassTypeExpr)
	NodeTypeExprNew                 = 80 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClosureUse          = 81 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPreDec              = 82 | NodeType(NodeClassTypeExpr)
	NodeTypeExprEmpty               = 83 | NodeType(NodeClassTypeExpr)
	NodeTypeExprErrorSuppress       = 84 | NodeType(NodeClassTypeExpr)
	NodeTypeExprEval                = 85 | NodeType(NodeClassTypeExpr)
	NodeTypeExprIsset               = 86 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClassConstFetch     = 87 | NodeType(NodeClassTypeExpr)
	NodeTypeExprExit                = 88 | NodeType(NodeClassTypeExpr)
	NodeTypeExprShortList           = 89 | NodeType(NodeClassTypeExpr)
	NodeTypeExprIncludeOnce         = 90 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPostInc             = 91 | NodeType(NodeClassTypeExpr)
	NodeTypeExprConstFetch          = 92 | NodeType(NodeClassTypeExpr)
	NodeTypeExprTernary             = 93 | NodeType(NodeClassTypeExpr)
	NodeTypeExprVariable            = 94 | NodeType(NodeClassTypeExpr)
	NodeTypeExprRequireOnce         = 95 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPostDec             = 96 | NodeType(NodeClassTypeExpr)
	NodeTypeExprBitwiseNot          = 97 | NodeType(NodeClassTypeExpr)
	NodeTypeExprInclude             = 98 | NodeType(NodeClassTypeExpr)
	NodeTypeExprList                = 99 | NodeType(NodeClassTypeExpr)
	NodeTypeExprShellExec           = 100 | NodeType(NodeClassTypeExpr)
	NodeTypeExprMethodCall          = 101 | NodeType(NodeClassTypeExpr)
	NodeTypeExprReference           = 102 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPrint               = 103 | NodeType(NodeClassTypeExpr)
	NodeTypeExprUnaryPlus           = 104 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPropertyFetch       = 105 | NodeType(NodeClassTypeExpr)
	NodeTypeExprArray               = 106 | NodeType(NodeClassTypeExpr)
	NodeTypeExprStaticPropertyFetch = 107 | NodeType(NodeClassTypeExpr)
	NodeTypeExprUnaryMinus          = 108 | NodeType(NodeClassTypeExpr)
	NodeTypeExprRequire             = 109 | NodeType(NodeClassTypeExpr)
	NodeTypeExprShortArray          = 110 | NodeType(NodeClassTypeExpr)
	NodeTypeExprPreInc              = 111 | NodeType(NodeClassTypeExpr)
	NodeTypeExprYieldFrom           = 112 | NodeType(NodeClassTypeExpr)
	NodeTypeExprYield               = 113 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClone               = 114 | NodeType(NodeClassTypeExpr)
	NodeTypeExprStaticCall          = 115 | NodeType(NodeClassTypeExpr)
	NodeTypeExprClosure             = 116 | NodeType(NodeClassTypeExpr)
	NodeTypeExprArrayDimFetch       = 117 | NodeType(NodeClassTypeExpr)
	NodeTypeExprArrayItem           = 118 | NodeType(NodeClassTypeExpr)
	NodeTypeExprFunctionCall        = 119 | NodeType(NodeClassTypeExpr)

	NodeTypeCastUnset  = 120 | NodeType(NodeClassTypeCast)
	NodeTypeCastDouble = 121 | NodeType(NodeClassTypeCast)
	NodeTypeCastString = 122 | NodeType(NodeClassTypeCast)
	NodeTypeCastBool   = 123 | NodeType(NodeClassTypeCast)
	NodeTypeCastArray  = 124 | NodeType(NodeClassTypeCast)
	NodeTypeCastInt    = 125 | NodeType(NodeClassTypeCast)
	NodeTypeCastObject = 126 | NodeType(NodeClassTypeCast)

	NodeTypeAssignAssign     = 127 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignMul        = 128 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignMinus      = 129 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignShiftLeft  = 130 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignMod        = 131 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignPow        = 132 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignBitwiseXor = 133 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignConcat     = 134 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignDiv        = 135 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignPlus       = 136 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignReference  = 137 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignBitwiseAnd = 138 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignBitwiseOr  = 139 | NodeType(NodeClassTypeAssign)
	NodeTypeAssignShiftRight = 140 | NodeType(NodeClassTypeAssign)

	NodeTypeBinaryBooleanOr      = 141 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryMul            = 142 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryMinus          = 143 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryEqual          = 144 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryIdentical      = 145 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryLogicalXor     = 146 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryCoalesce       = 147 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBooleanAnd     = 148 | NodeType(NodeClassTypeBinary)
	NodeTypeBinarySmaller        = 149 | NodeType(NodeClassTypeBinary)
	NodeTypeBinarySmallerOrEqual = 150 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryShiftLeft      = 151 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryMod            = 152 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryGreaterOrEqual = 153 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryLogicalAnd     = 154 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryLogicalOr      = 155 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryPow            = 156 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBitwiseXor     = 157 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryConcat         = 158 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryDiv            = 159 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryPlus           = 160 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryGreater        = 161 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBitwiseAnd     = 162 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryBitwiseOr      = 163 | NodeType(NodeClassTypeBinary)
	NodeTypeBinarySpaceship      = 164 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryShiftRight     = 165 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryNotIdentical   = 166 | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryNotEqual       = 167 | NodeType(NodeClassTypeBinary)
)

func (nt NodeType) Is(nct NodeClassType) bool {
	return uint16(nt)&uint16(nct) == uint16(nct)
}

type Node struct {
	Parent NodeID
	Type   NodeType
	Flag   NodeFlag

	Child EdgeID

	Pos PositionID
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
