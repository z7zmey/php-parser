/*

Names:

	name:
		namespace_name
		T_NAMESPACE T_NS_SEPARATOR namespace_name
		T_NS_SEPARATOR namespace_name

	namespace_name:
		T_STRING
		namespace_name T_NS_SEPARATOR T_STRING

Statement

	start:
		top_statement_list

	top_statement_list:
		top_statement_list top_statement
		// empty

	top_statement:
		statement
		function_declaration_statement
		class_declaration_statement
		trait_declaration_statement
		interface_declaration_statement
		T_HALT_COMPILER '(' ')' ';'
		T_NAMESPACE namespace_name ';'
		T_NAMESPACE namespace_name '{' top_statement_list '}'
		T_NAMESPACE '{' top_statement_list '}'
		T_USE mixed_group_use_declaration ';'
		T_USE use_type group_use_declaration ';'
		T_USE use_declarations ';'
		T_USE use_type use_declarations ';'
		T_CONST const_list ';'

	statement:
		'{' inner_statement_list '}'
		if_stmt
		alt_if_stmt
		T_WHILE '(' expr ')' while_statement
		T_DO statement T_WHILE '(' expr ')' ';'
		T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
		T_SWITCH '(' expr ')' switch_case_list
		T_BREAK optional_expr ';'
		T_CONTINUE optional_expr ';'
		T_RETURN optional_expr ';'
		T_GLOBAL global_var_list ';'
		T_STATIC static_var_list ';'
		T_ECHO echo_expr_list ';'
		T_INLINE_HTML
		expr ';'
		T_UNSET '(' unset_variables possible_comma ')' ';'
		T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
		T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
		T_DECLARE '(' const_list ')' declare_statement
		';'
		T_TRY '{' inner_statement_list '}' catch_list finally_statement
		T_THROW expr ';'
		T_GOTO T_STRING ';'
		T_STRING ':'

	inner_statement_list:
		inner_statement_list inner_statement
		// empty

	inner_statement:
		statement
		function_declaration_statement
		class_declaration_statement
		trait_declaration_statement
		interface_declaration_statement
		T_HALT_COMPILER '(' ')' ';'

	if_stmt:
		if_stmt_without_else %prec T_NOELSE
		if_stmt_without_else T_ELSE statement

	if_stmt_without_else:
		T_IF '(' expr ')' statement
		if_stmt_without_else T_ELSEIF '(' expr ')' statement

	alt_if_stmt:
		alt_if_stmt_without_else T_ENDIF ';'
		alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'

	alt_if_stmt_without_else:
		T_IF '(' expr ')' ':' inner_statement_list
		alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list

	while_statement:
		statement
		':' inner_statement_list T_ENDWHILE ';'

	for_statement:
		statement
		':' inner_statement_list T_ENDFOR ';'

 	for_exprs:
		// empty
		non_empty_for_exprs

	non_empty_for_exprs:
		non_empty_for_exprs ',' expr
		expr

	switch_case_list:
		'{' case_list '}'
		'{' ';' case_list '}'
		':' case_list T_ENDSWITCH ';'
		':' ';' case_list T_ENDSWITCH ';'

	case_list:
		// empty
		case_list T_CASE expr case_separator inner_statement_list
		case_list T_DEFAULT case_separator inner_statement_list

	case_separator:
		':'
		';'

	global_var_list:
		global_var_list ',' global_var
		global_var

	global_var:
		simple_variable

	static_var_list:
		static_var_list ',' static_var
		static_var

	static_var:
		T_VARIABLE
		T_VARIABLE '=' expr

	echo_expr_list:
		echo_expr_list ',' echo_expr
		echo_expr

	echo_expr:
		expr

	unset_variables:
		unset_variable
		unset_variables ',' unset_variable

	unset_variable:
		variable

	foreach_statement:
		statement
		':' inner_statement_list T_ENDFOREACH ';'

	foreach_variable:
		variable
		'&' variable
		T_LIST '(' array_pair_list ')'
		'[' array_pair_list ']'

	declare_statement:
		statement
		':' inner_statement_list T_ENDDECLARE ';'

	catch_list:
		// empty
		catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'

	catch_name_list:
		name
		catch_name_list '|' name

	finally_statement:
		// empty
		T_FINALLY '{' inner_statement_list '}'

Function declaration

	function_declaration_statement:
		T_FUNCTION returns_ref T_STRING backup_doc_comment '(' parameter_list ')' return_type '{' inner_statement_list '}'

	returns_ref:
		// empty
		'&'

	return_type:
		// empty
		':' type_expr

	type_expr:
		type
		'?' type

	type:
		T_ARRAY
		T_CALLABLE
		name

	parameter_list:
		non_empty_parameter_list
		// empty

	non_empty_parameter_list:
		parameter
		non_empty_parameter_list ',' parameter

	parameter:
		optional_type is_reference is_variadic T_VARIABLE
		optional_type is_reference is_variadic T_VARIABLE '=' expr

	optional_type:
		// empty
		type_expr

	is_reference:
		// empty
		'&'

	is_variadic:
		// empty
		T_ELLIPSIS

Class/Trait/Interface declaration

	class_declaration_statement:
		class_modifiers T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
		T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'

	trait_declaration_statement:
		T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'

	interface_declaration_statement:
		T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'

	class_modifiers:
		class_modifier
		class_modifiers class_modifier

	class_modifier:
		T_ABSTRACT
		T_FINAL

	extends_from:
		// empty
		T_EXTENDS name

	interface_extends_list:
		// empty
		T_EXTENDS name_list

	implements_list:
		// empty
		T_IMPLEMENTS name_list

	name_list:
		name
		name_list ',' name

Class statement

	class_statement_list:
		class_statement_list class_statement
		// empty

	class_statement:
		variable_modifiers property_list ';'
		method_modifiers T_CONST class_const_list ';'
		T_USE name_list trait_adaptations
		method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body

	method_modifiers:
		// empty
		non_empty_member_modifiers

	variable_modifiers:
		non_empty_member_modifiers
		T_VAR

	non_empty_member_modifiers:
		member_modifier
		non_empty_member_modifiers member_modifier

	member_modifier:
		T_PUBLIC
		T_PROTECTED
		T_PRIVATE
		T_STATIC
		T_ABSTRACT
		T_FINAL

	property_list:
		property_list ',' property

	property:
		T_VARIABLE backup_doc_comment
		T_VARIABLE '=' expr backup_doc_comment

	class_const_list:
		class_const_list ',' class_const_decl
		class_const_decl

	class_const_decl:
		identifier '=' expr backup_doc_comment

	method_body:
	';' // abstract method
	'{' inner_statement_list '}'

trait adaptation

	trait_adaptations:
		';'
		'{' '}'
		'{' trait_adaptation_list '}'

	trait_adaptation_list:
		trait_adaptation
		trait_adaptation_list trait_adaptation

	trait_adaptation:
		trait_precedence ';'
		trait_alias ';'

	trait_precedence:
		absolute_trait_method_reference T_INSTEADOF name_list

	trait_alias:
		trait_method_reference T_AS T_STRING
		trait_method_reference T_AS reserved_non_modifiers
		trait_method_reference T_AS member_modifier identifier
		trait_method_reference T_AS member_modifier

	trait_method_reference:
		identifier
		absolute_trait_method_reference

	absolute_trait_method_reference:
		name T_PAAMAYIM_NEKUDOTAYIM identifier

Use statements:

	mixed_group_use_declaration:
		namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
		T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'

	inline_use_declarations:
		inline_use_declarations ',' inline_use_declaration
		inline_use_declaration

	inline_use_declaration:
		unprefixed_use_declaration
		use_type unprefixed_use_declaration


	group_use_declaration:
		namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
		T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'

	unprefixed_use_declarations:
		unprefixed_use_declarations ',' unprefixed_use_declaration
		unprefixed_use_declaration


	use_declarations:
		use_declarations ',' use_declaration

	use_declaration:
		unprefixed_use_declaration
		T_NS_SEPARATOR unprefixed_use_declaration

	unprefixed_use_declaration:
		namespace_name
		namespace_name T_AS T_STRING


Expression

	expr:
		variable
		expr_without_variable

	expr_without_variable:
		T_LIST '(' array_pair_list ')' '=' expr
		[' array_pair_list ']' '=' expr
		variable '=' expr
		variable '=' '&' expr
		T_CLONE expr
		variable T_PLUS_EQUAL expr
		variable T_MINUS_EQUAL expr
		variable T_MUL_EQUAL expr
		variable T_POW_EQUAL expr
		variable T_DIV_EQUAL expr
		variable T_CONCAT_EQUAL expr
		variable T_MOD_EQUAL expr
		variable T_AND_EQUAL expr
		variable T_OR_EQUAL expr
		variable T_XOR_EQUAL expr
		variable T_SL_EQUAL expr
		variable T_SR_EQUAL expr
		variable T_INC
		T_INC variable
		variable T_DEC
		T_DEC variable
		expr T_BOOLEAN_OR expr
		expr T_BOOLEAN_AND expr
		expr T_LOGICAL_OR expr
		expr T_LOGICAL_AND expr
		expr T_LOGICAL_XOR expr
		expr '|' expr
		expr '&' expr
		expr '^' expr
		expr '.' expr
		expr '+' expr
		expr '-' expr
		expr '*' expr
		expr T_POW expr
		expr '/' expr
		expr '%' expr
		expr T_SL expr
		expr T_SR expr
		'+' expr %prec T_INC
		'-' expr %prec T_INC
		'!' expr
		'~' expr
		expr T_IS_IDENTICAL expr
		expr T_IS_NOT_IDENTICAL expr
		expr T_IS_EQUAL expr
		expr T_IS_NOT_EQUAL expr
		expr '<' expr
		expr T_IS_SMALLER_OR_EQUAL expr
		expr '>' expr
		expr T_IS_GREATER_OR_EQUAL expr
		expr T_SPACESHIP expr
		expr T_INSTANCEOF class_name_reference
		'(' expr ')'
		new_expr
		expr '?' expr ':' expr
		expr '?' ':' expr
		expr T_COALESCE expr
		internal_functions_in_yacc
		T_INT_CAST expr
		T_DOUBLE_CAST expr
		T_STRING_CAST expr
		T_ARRAY_CAST expr
		T_OBJECT_CAST expr
		T_BOOL_CAST expr
		T_UNSET_CAST expr
		T_EXIT exit_expr
		'@' expr
		scalar
		'`' backticks_expr '`'
		T_PRINT expr
		T_YIELD
		T_YIELD expr
		T_YIELD expr T_DOUBLE_ARROW expr
		T_YIELD_FROM expr
		T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
		T_STATIC T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'

	optional_expr:
		// empty
		expr

	exit_expr:
		// empty
		'(' optional_expr ')'

	internal_functions_in_yacc:
		T_ISSET '(' isset_variables possible_comma ')'
		T_EMPTY '(' expr ')'
		T_INCLUDE expr
		T_INCLUDE_ONCE expr
		T_EVAL '(' expr ')'
		T_REQUIRE expr
		T_REQUIRE_ONCE expr

	isset_variables
		isset_variable
		isset_variables ',' isset_variable

	isset_variable:
		expr

	new_expr:
		T_NEW class_name_reference ctor_arguments
		T_NEW anonymous_class

	anonymous_class:
		T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'

	class_name_reference:
		class_name
		new_variable

	new_variable:
		simple_variable
		new_variable '[' optional_expr ']'
		new_variable '{' expr '}'
		new_variable T_OBJECT_OPERATOR property_name
		class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
		new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable

	ctor_arguments:
		// empty
		argument_list

	lexical_vars:
		// empty
		T_USE '(' lexical_var_list ')'

	lexical_var_list:
		lexical_var_list ',' lexical_var
		lexical_var

	lexical_var:
		T_VARIABLE
		'&' T_VARIABLE

Scalars:

	scalar:
		T_LNUMBER
		T_DNUMBER
		T_LINE
		T_FILE
		T_DIR
		T_TRAIT_C
		T_METHOD_C
		T_FUNC_C
		T_NS_C
		T_CLASS_C
		T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
		T_START_HEREDOC T_END_HEREDOC
		'"' encaps_list '"'
		T_START_HEREDOC encaps_list T_END_HEREDOC
		dereferencable_scalar
		constant

	encaps_list:
		encaps_list encaps_var
		encaps_list T_ENCAPSED_AND_WHITESPACE
		encaps_var
		T_ENCAPSED_AND_WHITESPACE encaps_var

	encaps_var:
		T_VARIABLE
		T_VARIABLE '[' encaps_var_offset ']'
		T_VARIABLE T_OBJECT_OPERATOR T_STRING
		T_DOLLAR_OPEN_CURLY_BRACES expr '}'
		T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
		T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
		T_CURLY_OPEN variable '}'

	encaps_var_offset:
		T_STRING
		T_NUM_STRING
		'-' T_NUM_STRING
		T_VARIABLE

Variables:

	variable:
		callable_variable
		static_member
		dereferencable T_OBJECT_OPERATOR property_name

	callable_variable:
		simple_variable
		dereferencable '[' optional_expr ']'
		constant '[' optional_expr ']'
		dereferencable '{' expr '}'
		dereferencable T_OBJECT_OPERATOR property_name argument_list
		function_call

	constant:
		name
		class_name T_PAAMAYIM_NEKUDOTAYIM identifier
		variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier

	property_name:
		T_STRING
		'{' expr '}'
		simple_variable

	simple_variable:
		T_VARIABLE
		'$' '{' expr '}'
		'$' simple_variable

	static_member:
		class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
		variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable

	class_name:
		T_STATIC
		name

	variable_class_name:
		dereferencable

	dereferencable:
		variable
		'(' expr ')'
		dereferencable_scalar

	dereferencable_scalar:
		T_ARRAY '(' array_pair_list ')'
		'[' array_pair_list ']'
		T_CONSTANT_ENCAPSED_STRING

Function call:

	function_call:
		name argument_list
		class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
		variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
		callable_expr argument_list

	member_name:
		identifier
		'{' expr '}'
		simple_variable

	callable_expr:
		callable_variable
		'(' expr ')'
		dereferencable_scalar

Arguments:

	argument_list:
		'(' ')'
		'(' non_empty_argument_list possible_comma ')'

	non_empty_argument_list:
		argument
		non_empty_argument_list ',' argument

	argument:
		expr
		T_ELLIPSIS expr

Array pair

	array_pair_list:
		non_empty_array_pair_list

	non_empty_array_pair_list:
		non_empty_array_pair_list ',' possible_array_pair
		possible_array_pair

	possible_array_pair:
		// empty
		array_pair

	array_pair:
		expr T_DOUBLE_ARROW expr
		expr
		expr T_DOUBLE_ARROW '&' variable
		'&' variable
		expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
		T_LIST '(' array_pair_list ')'
*/
package php7 // import "github.com/z7zmey/php-parser/internal/parser/php7"
