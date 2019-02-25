package visitor_test

import (
	"bytes"
	"os"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

func ExamplePrettyJsonDumper() {
	src := `<?php

		namespace Foo {
			class Bar {
				public function FunctionName(Type $var = null)
				{
					// some comment
					// second comment
					$var;
				}
			}

			function foo() {
				;
			}
		}
		`

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.WithFreeFloating()
	php7parser.Parse()
	nodes := php7parser.GetRootNode()

	nsResolver := visitor.NewNamespaceResolver()
	nodes.Walk(nsResolver)

	dumper := visitor.NewPrettyJsonDumper(os.Stdout, nsResolver)
	nodes.Walk(dumper)

	// Unordered output:
	// {
	//   "type": "*node.Root",
	//   "position": {
	//     "startPos": 10,
	//     "endPos": 198,
	//     "startLine": 3,
	//     "endLine": 16
	//   },
	//   "freefloating": {
	//     "End": [
	//       {
	//         "type": "freefloating.WhiteSpaceType",
	//         "value": "\n\t\t"
	//       }
	//     ]
	//   },
	//   "Stmts": [
	//     {
	//       "type": "*stmt.Namespace",
	//       "position": {
	//         "startPos": 10,
	//         "endPos": 198,
	//         "startLine": 3,
	//         "endLine": 16
	//       },
	//       "freefloating": {
	//         "Start": [
	//           {
	//             "type": "freefloating.TokenType",
	//             "value": "<?php"
	//           },
	//           {
	//             "type": "freefloating.WhiteSpaceType",
	//             "value": "\n\n\t\t"
	//           }
	//         ],
	//         "Stmts": [
	//           {
	//             "type": "freefloating.WhiteSpaceType",
	//             "value": "\n\t\t"
	//           }
	//         ]
	//       },
	//       "NamespaceName": {
	//         "type": "*name.Name",
	//         "position": {
	//           "startPos": 20,
	//           "endPos": 22,
	//           "startLine": 3,
	//           "endLine": 3
	//         },
	//         "freefloating": {
	//           "End": [
	//             {
	//               "type": "freefloating.WhiteSpaceType",
	//               "value": " "
	//             }
	//           ],
	//           "Start": [
	//             {
	//               "type": "freefloating.WhiteSpaceType",
	//               "value": " "
	//             }
	//           ]
	//         },
	//         "Parts": [
	//           {
	//             "type": "*name.NamePart",
	//             "position": {
	//               "startPos": 20,
	//               "endPos": 22,
	//               "startLine": 3,
	//               "endLine": 3
	//             },
	//             "Value": "Foo"
	//           }
	//         ]
	//       },
	//       "Stmts": [
	//         {
	//           "type": "*stmt.Class",
	//           "position": {
	//             "startPos": 29,
	//             "endPos": 162,
	//             "startLine": 4,
	//             "endLine": 11
	//           },
	//           "namespacedName": "Foo\\Bar",
	//           "freefloating": {
	//             "Start": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": "\n\t\t\t"
	//               }
	//             ],
	//             "Name": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": " "
	//               }
	//             ],
	//             "Stmts": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": "\n\t\t\t"
	//               }
	//             ]
	//           },
	//           "PhpDocComment": "",
	//           "ClassName": {
	//             "type": "*node.Identifier",
	//             "position": {
	//               "startPos": 35,
	//               "endPos": 37,
	//               "startLine": 4,
	//               "endLine": 4
	//             },
	//             "freefloating": {
	//               "Start": [
	//                 {
	//                   "type": "freefloating.WhiteSpaceType",
	//                   "value": " "
	//                 }
	//               ]
	//             },
	//             "Value": "Bar"
	//           },
	//           "Stmts": [
	//             {
	//               "type": "*stmt.ClassMethod",
	//               "position": {
	//                 "startPos": 45,
	//                 "endPos": 157,
	//                 "startLine": 5,
	//                 "endLine": 10
	//               },
	//               "freefloating": {
	//                 "Start": [
	//                   {
	//                     "type": "freefloating.WhiteSpaceType",
	//                     "value": "\n\t\t\t\t"
	//                   }
	//                 ],
	//                 "ModifierList": [
	//                   {
	//                     "type": "freefloating.WhiteSpaceType",
	//                     "value": " "
	//                   }
	//                 ],
	//                 "Function": [
	//                   {
	//                     "type": "freefloating.WhiteSpaceType",
	//                     "value": " "
	//                   }
	//                 ],
	//                 "Name": [

	//                 ],
	//                 "ParameterList": [

	//                 ]
	//               },
	//               "ReturnsRef": false,
	//               "PhpDocComment": "",
	//               "MethodName": {
	//                 "type": "*node.Identifier",
	//                 "position": {
	//                   "startPos": 61,
	//                   "endPos": 72,
	//                   "startLine": 5,
	//                   "endLine": 5
	//                 },
	//                 "Value": "FunctionName"
	//               },
	//               "Modifiers": [
	//                 {
	//                   "type": "*node.Identifier",
	//                   "position": {
	//                     "startPos": 45,
	//                     "endPos": 50,
	//                     "startLine": 5,
	//                     "endLine": 5
	//                   },
	//                   "Value": "public"
	//                 }
	//               ],
	//               "Params": [
	//                 {
	//                   "type": "*node.Parameter",
	//                   "position": {
	//                     "startPos": 74,
	//                     "endPos": 89,
	//                     "startLine": 5,
	//                     "endLine": 5
	//                   },
	//                   "freefloating": {
	//                     "Ref": [

	//                     ],
	//                     "OptionalType": [
	//                       {
	//                         "type": "freefloating.WhiteSpaceType",
	//                         "value": " "
	//                       }
	//                     ],
	//                     "Start": [

	//                     ],
	//                     "Variadic": [

	//                     ],
	//                     "Var": [
	//                       {
	//                         "type": "freefloating.WhiteSpaceType",
	//                         "value": " "
	//                       }
	//                     ]
	//                   },
	//                   "ByRef": false,
	//                   "Variadic": false,
	//                   "VariableType": {
	//                     "type": "*name.Name",
	//                     "position": {
	//                       "startPos": 74,
	//                       "endPos": 77,
	//                       "startLine": 5,
	//                       "endLine": 5
	//                     },
	//                     "namespacedName": "Foo\\Type",
	//                     "Parts": [
	//                       {
	//                         "type": "*name.NamePart",
	//                         "position": {
	//                           "startPos": 74,
	//                           "endPos": 77,
	//                           "startLine": 5,
	//                           "endLine": 5
	//                         },
	//                         "Value": "Type"
	//                       }
	//                     ]
	//                   },
	//                   "Variable": {
	//                     "type": "*expr.Variable",
	//                     "position": {
	//                       "startPos": 79,
	//                       "endPos": 82,
	//                       "startLine": 5,
	//                       "endLine": 5
	//                     },
	//                     "freefloating": {
	//                       "Dollar": [
	//                         {
	//                           "type": "freefloating.TokenType",
	//                           "value": "$"
	//                         }
	//                       ]
	//                     },
	//                     "VarName": {
	//                       "type": "*node.Identifier",
	//                       "position": {
	//                         "startPos": 79,
	//                         "endPos": 82,
	//                         "startLine": 5,
	//                         "endLine": 5
	//                       },
	//                       "Value": "var"
	//                     }
	//                   },
	//                   "DefaultValue": {
	//                     "type": "*expr.ConstFetch",
	//                     "position": {
	//                       "startPos": 86,
	//                       "endPos": 89,
	//                       "startLine": 5,
	//                       "endLine": 5
	//                     },
	//                     "freefloating": {
	//                       "Start": [
	//                         {
	//                           "type": "freefloating.WhiteSpaceType",
	//                           "value": " "
	//                         }
	//                       ]
	//                     },
	//                     "Constant": {
	//                       "type": "*name.Name",
	//                       "position": {
	//                         "startPos": 86,
	//                         "endPos": 89,
	//                         "startLine": 5,
	//                         "endLine": 5
	//                       },
	//                       "namespacedName": "null",
	//                       "Parts": [
	//                         {
	//                           "type": "*name.NamePart",
	//                           "position": {
	//                             "startPos": 86,
	//                             "endPos": 89,
	//                             "startLine": 5,
	//                             "endLine": 5
	//                           },
	//                           "Value": "null"
	//                         }
	//                       ]
	//                     }
	//                   }
	//                 }
	//               ],
	//               "Stmt": {
	//                 "type": "*stmt.StmtList",
	//                 "position": {
	//                   "startPos": 96,
	//                   "endPos": 157,
	//                   "startLine": 6,
	//                   "endLine": 10
	//                 },
	//                 "freefloating": {
	//                   "Start": [
	//                     {
	//                       "type": "freefloating.WhiteSpaceType",
	//                       "value": "\n\t\t\t\t"
	//                     }
	//                   ],
	//                   "Stmts": [
	//                     {
	//                       "type": "freefloating.WhiteSpaceType",
	//                       "value": "\n\t\t\t\t"
	//                     }
	//                   ]
	//                 },
	//                 "Stmts": [
	//                   {
	//                     "type": "*stmt.Expression",
	//                     "position": {
	//                       "startPos": 147,
	//                       "endPos": 151,
	//                       "startLine": 9,
	//                       "endLine": 9
	//                     },
	//                     "freefloating": {
	//                       "Start": [
	//                         {
	//                           "type": "freefloating.WhiteSpaceType",
	//                           "value": "\n\t\t\t\t\t"
	//                         },
	//                         {
	//                           "type": "freefloating.CommentType",
	//                           "value": "// some comment\n"
	//                         },
	//                         {
	//                           "type": "freefloating.WhiteSpaceType",
	//                           "value": "\t\t\t\t\t"
	//                         },
	//                         {
	//                           "type": "freefloating.CommentType",
	//                           "value": "// second comment\n"
	//                         },
	//                         {
	//                           "type": "freefloating.WhiteSpaceType",
	//                           "value": "\t\t\t\t\t"
	//                         }
	//                       ],
	//                       "Expr": [

	//                       ],
	//                       "SemiColon": [
	//                         {
	//                           "type": "freefloating.TokenType",
	//                           "value": ";"
	//                         }
	//                       ]
	//                     },
	//                     "Expr": {
	//                       "type": "*expr.Variable",
	//                       "position": {
	//                         "startPos": 147,
	//                         "endPos": 150,
	//                         "startLine": 9,
	//                         "endLine": 9
	//                       },
	//                       "freefloating": {
	//                         "Start": [

	//                         ],
	//                         "Dollar": [
	//                           {
	//                             "type": "freefloating.TokenType",
	//                             "value": "$"
	//                           }
	//                         ]
	//                       },
	//                       "VarName": {
	//                         "type": "*node.Identifier",
	//                         "position": {
	//                           "startPos": 147,
	//                           "endPos": 150,
	//                           "startLine": 9,
	//                           "endLine": 9
	//                         },
	//                         "Value": "var"
	//                       }
	//                     }
	//                   }
	//                 ]
	//               }
	//             }
	//           ]
	//         },
	//         {
	//           "type": "*stmt.Function",
	//           "position": {
	//             "startPos": 168,
	//             "endPos": 194,
	//             "startLine": 13,
	//             "endLine": 15
	//           },
	//           "namespacedName": "Foo\\foo",
	//           "freefloating": {
	//             "Start": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": "\n\n\t\t\t"
	//               }
	//             ],
	//             "Function": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": " "
	//               }
	//             ],
	//             "Name": [

	//             ],
	//             "ParamList": [

	//             ],
	//             "ReturnType": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": " "
	//               }
	//             ],
	//             "Stmts": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": "\n\t\t\t"
	//               }
	//             ]
	//           },
	//           "ReturnsRef": false,
	//           "PhpDocComment": "",
	//           "FunctionName": {
	//             "type": "*node.Identifier",
	//             "position": {
	//               "startPos": 177,
	//               "endPos": 179,
	//               "startLine": 13,
	//               "endLine": 13
	//             },
	//             "Value": "foo"
	//           },
	//           "Stmts": [
	//             {
	//               "type": "*stmt.Nop",
	//               "position": {
	//                 "startPos": 189,
	//                 "endPos": 189,
	//                 "startLine": 14,
	//                 "endLine": 14
	//               },
	//               "freefloating": {
	//                 "Start": [
	//                   {
	//                     "type": "freefloating.WhiteSpaceType",
	//                     "value": "\n\t\t\t\t"
	//                   }
	//                 ],
	//                 "SemiColon": [
	//                   {
	//                     "type": "freefloating.TokenType",
	//                     "value": ";"
	//                   }
	//                 ]
	//               }
	//             }
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }
}
