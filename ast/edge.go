package ast

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
