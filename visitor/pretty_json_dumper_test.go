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
	php7parser.WithMeta()
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
	//   "meta": [
	//     {
	//       "type": "*meta.WhiteSpaceType",
	//       "value": "\n\t\t",
	//       "tokenName": "NodeEnd"
	//     }
	//   ],
	//   "Stmts": [
	//     {
	//       "type": "*stmt.Namespace",
	//       "position": {
	//         "startPos": 10,
	//         "endPos": 198,
	//         "startLine": 3,
	//         "endLine": 16
	//       },
	//       "meta": [
	//         {
	//           "type": "*meta.TokenType",
	//           "value": "<?php",
	//           "tokenName": "NamespaceToken"
	//         },
	//         {
	//           "type": "*meta.WhiteSpaceType",
	//           "value": "\n\n\t\t",
	//           "tokenName": "NamespaceToken"
	//         },
	//         {
	//           "type": "*meta.WhiteSpaceType",
	//           "value": " ",
	//           "tokenName": "OpenCurlyBracesToken"
	//         },
	//         {
	//           "type": "*meta.WhiteSpaceType",
	//           "value": "\n\t\t",
	//           "tokenName": "CloseCurlyBracesToken"
	//         }
	//       ],
	//       "NamespaceName": {
	//         "type": "*name.Name",
	//         "position": {
	//           "startPos": 20,
	//           "endPos": 22,
	//           "startLine": 3,
	//           "endLine": 3
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
	//             "meta": [
	//               {
	//                 "type": "*meta.WhiteSpaceType",
	//                 "value": " ",
	//                 "tokenName": "StringToken"
	//               }
	//             ],
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
	//           "meta": [
	//             {
	//               "type": "*meta.WhiteSpaceType",
	//               "value": "\n\t\t\t",
	//               "tokenName": "ClassToken"
	//             },
	//             {
	//               "type": "*meta.WhiteSpaceType",
	//               "value": " ",
	//               "tokenName": "OpenCurlyBracesToken"
	//             },
	//             {
	//               "type": "*meta.WhiteSpaceType",
	//               "value": "\n\t\t\t",
	//               "tokenName": "CloseCurlyBracesToken"
	//             }
	//           ],
	//           "PhpDocComment": "",
	//           "ClassName": {
	//             "type": "*node.Identifier",
	//             "position": {
	//               "startPos": 35,
	//               "endPos": 37,
	//               "startLine": 4,
	//               "endLine": 4
	//             },
	//             "meta": [
	//               {
	//                 "type": "*meta.WhiteSpaceType",
	//                 "value": " ",
	//                 "tokenName": "NodeStart"
	//               }
	//             ],
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
	//               "meta": [
	//                 {
	//                   "type": "*meta.WhiteSpaceType",
	//                   "value": " ",
	//                   "tokenName": "FunctionToken"
	//                 }
	//               ],
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
	//                 "meta": [
	//                   {
	//                     "type": "*meta.WhiteSpaceType",
	//                     "value": " ",
	//                     "tokenName": "NodeStart"
	//                   }
	//                 ],
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
	//                   "meta": [
	//                     {
	//                       "type": "*meta.WhiteSpaceType",
	//                       "value": "\n\t\t\t\t",
	//                       "tokenName": "NodeStart"
	//                     }
	//                   ],
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
	//                   "meta": [
	//                     {
	//                       "type": "*meta.WhiteSpaceType",
	//                       "value": " ",
	//                       "tokenName": "EqualToken"
	//                     }
	//                   ],
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
	//                     "meta": [
	//                       {
	//                         "type": "*meta.WhiteSpaceType",
	//                         "value": " ",
	//                         "tokenName": "NodeStart"
	//                       },
	//                       {
	//                         "type": "*meta.TokenType",
	//                         "value": "$",
	//                         "tokenName": "NodeStart"
	//                       }
	//                     ],
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
	//                           "meta": [
	//                             {
	//                               "type": "*meta.WhiteSpaceType",
	//                               "value": " ",
	//                               "tokenName": "StringToken"
	//                             }
	//                           ],
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
	//                 "meta": [
	//                   {
	//                     "type": "*meta.WhiteSpaceType",
	//                     "value": "\n\t\t\t\t",
	//                     "tokenName": "OpenCurlyBracesToken"
	//                   },
	//                   {
	//                     "type": "*meta.WhiteSpaceType",
	//                     "value": "\n\t\t\t\t",
	//                     "tokenName": "CloseCurlyBracesToken"
	//                   }
	//                 ],
	//                 "Stmts": [
	//                   {
	//                     "type": "*stmt.Expression",
	//                     "position": {
	//                       "startPos": 147,
	//                       "endPos": 151,
	//                       "startLine": 9,
	//                       "endLine": 9
	//                     },
	//                     "meta": [
	//                       {
	//                         "type": "*meta.TokenType",
	//                         "value": ";",
	//                         "tokenName": "SemiColonToken"
	//                       }
	//                     ],
	//                     "Expr": {
	//                       "type": "*expr.Variable",
	//                       "position": {
	//                         "startPos": 147,
	//                         "endPos": 150,
	//                         "startLine": 9,
	//                         "endLine": 9
	//                       },
	//                       "meta": [
	//                         {
	//                           "type": "*meta.WhiteSpaceType",
	//                           "value": "\n\t\t\t\t\t",
	//                           "tokenName": "NodeStart"
	//                         },
	//                         {
	//                           "type": "*meta.CommentType",
	//                           "value": "// some comment\n",
	//                           "tokenName": "NodeStart"
	//                         },
	//                         {
	//                           "type": "*meta.WhiteSpaceType",
	//                           "value": "\t\t\t\t\t",
	//                           "tokenName": "NodeStart"
	//                         },
	//                         {
	//                           "type": "*meta.CommentType",
	//                           "value": "// second comment\n",
	//                           "tokenName": "NodeStart"
	//                         },
	//                         {
	//                           "type": "*meta.WhiteSpaceType",
	//                           "value": "\t\t\t\t\t",
	//                           "tokenName": "NodeStart"
	//                         },
	//                         {
	//                           "type": "*meta.TokenType",
	//                           "value": "$",
	//                           "tokenName": "NodeStart"
	//                         }
	//                       ],
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
	//           "meta": [
	//             {
	//               "type": "*meta.WhiteSpaceType",
	//               "value": "\n\n\t\t\t",
	//               "tokenName": "FunctionToken"
	//             },
	//             {
	//               "type": "*meta.WhiteSpaceType",
	//               "value": " ",
	//               "tokenName": "OpenCurlyBracesToken"
	//             },
	//             {
	//               "type": "*meta.WhiteSpaceType",
	//               "value": "\n\t\t\t",
	//               "tokenName": "CloseCurlyBracesToken"
	//             }
	//           ],
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
	//             "meta": [
	//               {
	//                 "type": "*meta.WhiteSpaceType",
	//                 "value": " ",
	//                 "tokenName": "NodeStart"
	//               }
	//             ],
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
	//               "meta": [
	//                 {
	//                   "type": "*meta.WhiteSpaceType",
	//                   "value": "\n\t\t\t\t",
	//                   "tokenName": "SemiColonToken"
	//                 },
	//                 {
	//                   "type": "*meta.TokenType",
	//                   "value": ";",
	//                   "tokenName": "SemiColonToken"
	//                 }
	//               ]
	//             }
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }
}
