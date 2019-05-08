package ast

type NodeFlag uint8

//go:generate stringer -type=NodeFlag -trimprefix=NodeFlag -output ./nodeflag_string.go
const (
	NodeFlagRef NodeFlag = 1 << iota
	NodeFlagVariadic
	NodeFlagStatic
	NodeFlagAltSyntax
)

func (nf NodeFlag) GetFlagNames() []string {
	flags := make([]string, 0, 4)

	for _, flag := range [...]NodeFlag{NodeFlagRef, NodeFlagVariadic, NodeFlagStatic, NodeFlagAltSyntax} {
		if nf&flag != 0 {
			flags = append(flags, flag.String())
		}
	}

	return flags
}

type NodeClassType uint16

const (
	NodeClassTypeValue  NodeClassType = 1 << 8
	NodeClassTypeScalar NodeClassType = 1 << 9
	NodeClassTypeName   NodeClassType = 1 << 10
	NodeClassTypeStmt   NodeClassType = 1 << 11
	NodeClassTypeExpr   NodeClassType = 1 << 12
	NodeClassTypeAssign NodeClassType = 1<<13 | NodeClassTypeExpr
	NodeClassTypeBinary NodeClassType = 1<<14 | NodeClassTypeExpr
	NodeClassTypeCast   NodeClassType = 1<<15 | NodeClassTypeExpr
)

type NodeType uint16

//go:generate stringer -type=NodeType -trimprefix=NodeType -output ./nodetype_string.go
const (
	NodeTypeRoot NodeType = iota

	NodeTypeParameter
	NodeTypeArgument
	NodeTypeArgumentList
	NodeTypeNullable
	NodeTypeIdentifier NodeType = iota | NodeType(NodeClassTypeValue)

	NodeTypeNameNamePart NodeType = iota | NodeType(NodeClassTypeName) | NodeType(NodeClassTypeValue)
	NodeTypeNameName     NodeType = iota | NodeType(NodeClassTypeName)
	NodeTypeNameFullyQualified
	NodeTypeNameRelative

	NodeTypeScalarEncapsedStringPart NodeType = iota | NodeType(NodeClassTypeScalar) | NodeType(NodeClassTypeValue)
	NodeTypeScalarDnumber
	NodeTypeScalarMagicConstant
	NodeTypeScalarLnumber
	NodeTypeScalarString
	NodeTypeScalarHeredoc NodeType = iota | NodeType(NodeClassTypeScalar)
	NodeTypeScalarEncapsed

	NodeTypeStmtInlineHtml NodeType = iota | NodeType(NodeClassTypeStmt) | NodeType(NodeClassTypeValue)
	NodeTypeStmtFinally    NodeType = iota | NodeType(NodeClassTypeStmt)
	NodeTypeStmtNop
	NodeTypeStmtUnset
	NodeTypeStmtStatic
	NodeTypeStmtProperty
	NodeTypeStmtGroupUse
	NodeTypeStmtAltElseIf
	NodeTypeStmtElse
	NodeTypeStmtCase
	NodeTypeStmtConstList
	NodeTypeStmtExpression
	NodeTypeStmtReturn
	NodeTypeStmtAltForeach
	NodeTypeStmtFor
	NodeTypeStmtClassConstList
	NodeTypeStmtInterface
	NodeTypeStmtEcho
	NodeTypeStmtInterfaceExtends
	NodeTypeStmtTrait
	NodeTypeStmtContinue
	NodeTypeStmtClassExtends
	NodeTypeStmtUseList
	NodeTypeStmtTraitUse
	NodeTypeStmtTraitUsePrecedence
	NodeTypeStmtTry
	NodeTypeStmtDeclare
	NodeTypeStmtDefault
	NodeTypeStmtForeach
	NodeTypeStmtTraitAdaptationList
	NodeTypeStmtGlobal
	NodeTypeStmtWhile
	NodeTypeStmtIf
	NodeTypeStmtAltElse
	NodeTypeStmtGoto
	NodeTypeStmtClass
	NodeTypeStmtLabel
	NodeTypeStmtThrow
	NodeTypeStmtClassImplements
	NodeTypeStmtUse
	NodeTypeStmtPropertyList
	NodeTypeStmtAltWhile
	NodeTypeStmtAltIf
	NodeTypeStmtCatch
	NodeTypeStmtConstant
	NodeTypeStmtTraitUseAlias
	NodeTypeStmtNamespace
	NodeTypeStmtSwitch
	NodeTypeStmtBreak
	NodeTypeStmtClassMethod
	NodeTypeStmtFunction
	NodeTypeStmtAltSwitch
	NodeTypeStmtHaltCompiler
	NodeTypeStmtTraitMethodRef
	NodeTypeStmtDo
	NodeTypeStmtStaticVar
	NodeTypeStmtCaseList
	NodeTypeStmtElseIf
	NodeTypeStmtAltFor
	NodeTypeStmtStmtList
	NodeTypeStmtReturnType

	NodeTypeExprInstanceOf NodeType = iota | NodeType(NodeClassTypeExpr)
	NodeTypeExprBooleanNot
	NodeTypeExprNew
	NodeTypeExprClosureUse
	NodeTypeExprPreDec
	NodeTypeExprEmpty
	NodeTypeExprErrorSuppress
	NodeTypeExprEval
	NodeTypeExprIsset
	NodeTypeExprClassConstFetch
	NodeTypeExprExit
	NodeTypeExprShortList
	NodeTypeExprIncludeOnce
	NodeTypeExprPostInc
	NodeTypeExprConstFetch
	NodeTypeExprTernary
	NodeTypeExprVariable
	NodeTypeExprRequireOnce
	NodeTypeExprPostDec
	NodeTypeExprBitwiseNot
	NodeTypeExprInclude
	NodeTypeExprList
	NodeTypeExprShellExec
	NodeTypeExprMethodCall
	NodeTypeExprReference
	NodeTypeExprPrint
	NodeTypeExprUnaryPlus
	NodeTypeExprPropertyFetch
	NodeTypeExprArray
	NodeTypeExprStaticPropertyFetch
	NodeTypeExprUnaryMinus
	NodeTypeExprRequire
	NodeTypeExprShortArray
	NodeTypeExprPreInc
	NodeTypeExprYieldFrom
	NodeTypeExprYield
	NodeTypeExprClone
	NodeTypeExprStaticCall
	NodeTypeExprClosure
	NodeTypeExprArrayDimFetch
	NodeTypeExprArrayItem
	NodeTypeExprFunctionCall

	NodeTypeCastUnset NodeType = iota | NodeType(NodeClassTypeCast)
	NodeTypeCastDouble
	NodeTypeCastString
	NodeTypeCastBool
	NodeTypeCastArray
	NodeTypeCastInt
	NodeTypeCastObject

	NodeTypeAssignAssign NodeType = iota | NodeType(NodeClassTypeAssign)
	NodeTypeAssignMul
	NodeTypeAssignMinus
	NodeTypeAssignShiftLeft
	NodeTypeAssignMod
	NodeTypeAssignPow
	NodeTypeAssignBitwiseXor
	NodeTypeAssignConcat
	NodeTypeAssignDiv
	NodeTypeAssignPlus
	NodeTypeAssignReference
	NodeTypeAssignBitwiseAnd
	NodeTypeAssignBitwiseOr
	NodeTypeAssignShiftRight

	NodeTypeBinaryBooleanOr NodeType = iota | NodeType(NodeClassTypeBinary)
	NodeTypeBinaryMul
	NodeTypeBinaryMinus
	NodeTypeBinaryEqual
	NodeTypeBinaryIdentical
	NodeTypeBinaryLogicalXor
	NodeTypeBinaryCoalesce
	NodeTypeBinaryBooleanAnd
	NodeTypeBinarySmaller
	NodeTypeBinarySmallerOrEqual
	NodeTypeBinaryShiftLeft
	NodeTypeBinaryMod
	NodeTypeBinaryGreaterOrEqual
	NodeTypeBinaryLogicalAnd
	NodeTypeBinaryLogicalOr
	NodeTypeBinaryPow
	NodeTypeBinaryBitwiseXor
	NodeTypeBinaryConcat
	NodeTypeBinaryDiv
	NodeTypeBinaryPlus
	NodeTypeBinaryGreater
	NodeTypeBinaryBitwiseAnd
	NodeTypeBinaryBitwiseOr
	NodeTypeBinarySpaceship
	NodeTypeBinaryShiftRight
	NodeTypeBinaryNotIdentical
	NodeTypeBinaryNotEqual
)

func (nt NodeType) Is(nct NodeClassType) bool {
	return uint16(nt)&uint16(nct) == uint16(nct)
}
