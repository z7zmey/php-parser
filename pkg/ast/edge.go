package ast

type EdgeClassType uint8

const EdgeClassTypeMultiple EdgeClassType = 1 << 7

type EdgeType uint8

//go:generate stringer -type=EdgeType -trimprefix=EdgeType -output ./edgetype_string.go
const (
	EdgeTypeNil EdgeType = iota
	EdgeTypeExpr
	EdgeTypeVarType
	EdgeTypeVar
	EdgeTypeDefaultValue
	EdgeTypeStmt
	EdgeTypeElseIf
	EdgeTypeMethodName
	EdgeTypeReturnType
	EdgeTypeClassName
	EdgeTypeExtends
	EdgeTypeImplements
	EdgeTypeConstantName
	EdgeTypeKey
	EdgeTypeFunctionName
	EdgeTypeLabel
	EdgeTypeUseType
	EdgeTypePrefix
	EdgeTypeInterfaceName
	EdgeTypeLabelName
	EdgeTypeNamespaceName
	EdgeTypeCaseList
	EdgeTypeTrait
	EdgeTypeMethod
	EdgeTypeRef
	EdgeTypeModifier
	EdgeTypeAlias
	EdgeTypeTraitAdaptationList
	EdgeTypeTraitName
	EdgeTypeFinally
	EdgeTypeUse
	EdgeTypeLeft
	EdgeTypeRight
	EdgeTypeDim
	EdgeTypeVal
	EdgeTypeClass
	EdgeTypeClosureUse
	EdgeTypeConstant
	EdgeTypeFunction
	EdgeTypeProperty
	EdgeTypeCall
	EdgeTypeIfTrue
	EdgeTypeIfFalse
	EdgeTypeVarName

	EdgeTypeStmts EdgeType = iota | EdgeType(EdgeClassTypeMultiple)
	EdgeTypeParts
	EdgeTypeUses
	EdgeTypeConsts
	EdgeTypeUseList
	EdgeTypeLoop
	EdgeTypeCond
	EdgeTypeInit
	EdgeTypeVars
	EdgeTypeExprs
	EdgeTypeCatches
	EdgeTypeTypes
	EdgeTypeParams
	EdgeTypeModifiers
	EdgeTypeInterfaceNames
	EdgeTypeItems
	EdgeTypeCases
	EdgeTypeArguments
	EdgeTypeProperties
	EdgeTypeTraits
	EdgeTypeAdaptations
	EdgeTypeInsteadof
	EdgeTypeArgumentList
	EdgeTypeElse
)

func (et EdgeType) Is(ect EdgeClassType) bool {
	return uint8(et)&uint8(ect) == uint8(ect)
}
