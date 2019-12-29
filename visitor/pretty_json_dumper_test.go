package visitor_test

import (
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

	php7parser := php7.NewParser([]byte(src), "7.4")
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
	//     "startPos": 9,
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
	//         "startPos": 9,
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
	//           "startPos": 19,
	//           "endPos": 22,
	//           "startLine": 3,
	//           "endLine": 3
	//         },
	//         "freefloating": {
	//           "Start": [
	//             {
	//               "type": "freefloating.WhiteSpaceType",
	//               "value": " "
	//             }
	//           ],
	//           "End": [
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
	//               "startPos": 19,
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
	//             "startPos": 28,
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
	//               "startPos": 34,
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
	//                 "startPos": 44,
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
	//                 ]
	//               },
	//               "ReturnsRef": false,
	//               "PhpDocComment": "",
	//               "MethodName": {
	//                 "type": "*node.Identifier",
	//                 "position": {
	//                   "startPos": 60,
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
	//                     "startPos": 44,
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
	//                     "startPos": 73,
	//                     "endPos": 89,
	//                     "startLine": 5,
	//                     "endLine": 5
	//                   },
	//                   "freefloating": {
	//                     "OptionalType": [
	//                       {
	//                         "type": "freefloating.WhiteSpaceType",
	//                         "value": " "
	//                       }
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
	//                       "startPos": 73,
	//                       "endPos": 77,
	//                       "startLine": 5,
	//                       "endLine": 5
	//                     },
	//                     "namespacedName": "Foo\\Type",
	//                     "Parts": [
	//                       {
	//                         "type": "*name.NamePart",
	//                         "position": {
	//                           "startPos": 73,
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
	//                       "startPos": 78,
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
	//                         "startPos": 78,
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
	//                       "startPos": 85,
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
	//                         "startPos": 85,
	//                         "endPos": 89,
	//                         "startLine": 5,
	//                         "endLine": 5
	//                       },
	//                       "namespacedName": "null",
	//                       "Parts": [
	//                         {
	//                           "type": "*name.NamePart",
	//                           "position": {
	//                             "startPos": 85,
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
	//                   "startPos": 95,
	//                   "endPos": 157,
	//                   "startLine": 6,
	//                   "endLine": 10
	//                 },
	//                 "freefloating": {
	//                   "Stmts": [
	//                     {
	//                       "type": "freefloating.WhiteSpaceType",
	//                       "value": "\n\t\t\t\t"
	//                     }
	//                   ],
	//                   "Start": [
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
	//                       "startPos": 146,
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
	//                         "startPos": 146,
	//                         "endPos": 150,
	//                         "startLine": 9,
	//                         "endLine": 9
	//                       },
	//                       "freefloating": {
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
	//                           "startPos": 146,
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
	//             "startPos": 167,
	//             "endPos": 194,
	//             "startLine": 13,
	//             "endLine": 15
	//           },
	//           "namespacedName": "Foo\\foo",
	//           "freefloating": {
	//             "Params": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": " "
	//               }
	//             ],
	//             "Start": [
	//               {
	//                 "type": "freefloating.WhiteSpaceType",
	//                 "value": "\n\n\t\t\t"
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
	//               "startPos": 176,
	//               "endPos": 179,
	//               "startLine": 13,
	//               "endLine": 13
	//             },
	//             "freefloating": {
	//               "Start": [
	//                 {
	//                   "type": "freefloating.WhiteSpaceType",
	//                   "value": " "
	//                 }
	//               ]
	//             },
	//             "Value": "foo"
	//           },
	//           "Stmts": [
	//             {
	//               "type": "*stmt.Nop",
	//               "position": {
	//                 "startPos": 188,
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
