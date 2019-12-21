//line scanner/scanner.rl:1
package scanner

import (
	"fmt"

	"github.com/z7zmey/php-parser/freefloating"
)

//line scanner/scanner.go:13
const lexer_start int = 107
const lexer_first_final int = 107
const lexer_error int = 0

const lexer_en_main int = 107
const lexer_en_php int = 114
const lexer_en_property int = 459
const lexer_en_nowdoc int = 465
const lexer_en_heredoc int = 468
const lexer_en_backqote int = 474
const lexer_en_template_string int = 480
const lexer_en_heredoc_end int = 486
const lexer_en_string_var int = 488
const lexer_en_string_var_index int = 493
const lexer_en_string_var_name int = 503
const lexer_en_halt_compiller_open_parenthesis int = 505
const lexer_en_halt_compiller_close_parenthesis int = 509
const lexer_en_halt_compiller_close_semicolon int = 513
const lexer_en_halt_compiller_end int = 517

//line scanner/scanner.rl:15

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{
		data:  data,
		pe:    len(data),
		stack: make([]int, 0),

		TokenPool: &TokenPool{},
		NewLines:  NewLines{make([]int, 0, 128)},
	}

//line scanner/scanner.go:48
	{
		lex.cs = lexer_start
		lex.top = 0
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line scanner/scanner.rl:27
	return lex
}

func (lex *Lexer) Lex(lval Lval) int {
	lex.FreeFloating = nil
	eof := lex.pe
	var tok TokenID

	token := lex.TokenPool.Get()
	token.FreeFloating = lex.FreeFloating
	token.Value = string(lex.data[0:0])

	lblStart := 0
	lblEnd := 0

	_, _ = lblStart, lblEnd

//line scanner/scanner.go:76
	{
		var _widec int16
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 107:
			goto st107
		case 108:
			goto st108
		case 109:
			goto st109
		case 110:
			goto st110
		case 111:
			goto st111
		case 112:
			goto st112
		case 1:
			goto st1
		case 2:
			goto st2
		case 3:
			goto st3
		case 113:
			goto st113
		case 4:
			goto st4
		case 114:
			goto st114
		case 115:
			goto st115
		case 116:
			goto st116
		case 5:
			goto st5
		case 117:
			goto st117
		case 118:
			goto st118
		case 119:
			goto st119
		case 120:
			goto st120
		case 6:
			goto st6
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 121:
			goto st121
		case 122:
			goto st122
		case 123:
			goto st123
		case 124:
			goto st124
		case 125:
			goto st125
		case 126:
			goto st126
		case 127:
			goto st127
		case 10:
			goto st10
		case 11:
			goto st11
		case 128:
			goto st128
		case 12:
			goto st12
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto st40
		case 41:
			goto st41
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 45:
			goto st45
		case 46:
			goto st46
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 129:
			goto st129
		case 130:
			goto st130
		case 131:
			goto st131
		case 132:
			goto st132
		case 133:
			goto st133
		case 66:
			goto st66
		case 134:
			goto st134
		case 67:
			goto st67
		case 68:
			goto st68
		case 135:
			goto st135
		case 136:
			goto st136
		case 69:
			goto st69
		case 70:
			goto st70
		case 71:
			goto st71
		case 137:
			goto st137
		case 138:
			goto st138
		case 72:
			goto st72
		case 139:
			goto st139
		case 73:
			goto st73
		case 140:
			goto st140
		case 141:
			goto st141
		case 142:
			goto st142
		case 74:
			goto st74
		case 75:
			goto st75
		case 76:
			goto st76
		case 77:
			goto st77
		case 143:
			goto st143
		case 144:
			goto st144
		case 78:
			goto st78
		case 145:
			goto st145
		case 146:
			goto st146
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 82:
			goto st82
		case 147:
			goto st147
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 86:
			goto st86
		case 148:
			goto st148
		case 149:
			goto st149
		case 150:
			goto st150
		case 151:
			goto st151
		case 152:
			goto st152
		case 153:
			goto st153
		case 154:
			goto st154
		case 155:
			goto st155
		case 87:
			goto st87
		case 156:
			goto st156
		case 157:
			goto st157
		case 158:
			goto st158
		case 159:
			goto st159
		case 160:
			goto st160
		case 161:
			goto st161
		case 162:
			goto st162
		case 163:
			goto st163
		case 164:
			goto st164
		case 165:
			goto st165
		case 166:
			goto st166
		case 167:
			goto st167
		case 168:
			goto st168
		case 88:
			goto st88
		case 89:
			goto st89
		case 169:
			goto st169
		case 170:
			goto st170
		case 171:
			goto st171
		case 172:
			goto st172
		case 173:
			goto st173
		case 174:
			goto st174
		case 175:
			goto st175
		case 176:
			goto st176
		case 177:
			goto st177
		case 178:
			goto st178
		case 179:
			goto st179
		case 180:
			goto st180
		case 181:
			goto st181
		case 182:
			goto st182
		case 183:
			goto st183
		case 184:
			goto st184
		case 185:
			goto st185
		case 186:
			goto st186
		case 187:
			goto st187
		case 188:
			goto st188
		case 189:
			goto st189
		case 190:
			goto st190
		case 191:
			goto st191
		case 192:
			goto st192
		case 193:
			goto st193
		case 194:
			goto st194
		case 195:
			goto st195
		case 196:
			goto st196
		case 197:
			goto st197
		case 198:
			goto st198
		case 199:
			goto st199
		case 200:
			goto st200
		case 201:
			goto st201
		case 202:
			goto st202
		case 203:
			goto st203
		case 204:
			goto st204
		case 205:
			goto st205
		case 206:
			goto st206
		case 207:
			goto st207
		case 208:
			goto st208
		case 209:
			goto st209
		case 210:
			goto st210
		case 211:
			goto st211
		case 212:
			goto st212
		case 213:
			goto st213
		case 214:
			goto st214
		case 215:
			goto st215
		case 216:
			goto st216
		case 217:
			goto st217
		case 218:
			goto st218
		case 219:
			goto st219
		case 220:
			goto st220
		case 221:
			goto st221
		case 222:
			goto st222
		case 223:
			goto st223
		case 224:
			goto st224
		case 225:
			goto st225
		case 226:
			goto st226
		case 227:
			goto st227
		case 228:
			goto st228
		case 229:
			goto st229
		case 230:
			goto st230
		case 231:
			goto st231
		case 232:
			goto st232
		case 233:
			goto st233
		case 234:
			goto st234
		case 235:
			goto st235
		case 236:
			goto st236
		case 237:
			goto st237
		case 238:
			goto st238
		case 239:
			goto st239
		case 240:
			goto st240
		case 241:
			goto st241
		case 242:
			goto st242
		case 243:
			goto st243
		case 244:
			goto st244
		case 245:
			goto st245
		case 246:
			goto st246
		case 247:
			goto st247
		case 248:
			goto st248
		case 249:
			goto st249
		case 250:
			goto st250
		case 251:
			goto st251
		case 252:
			goto st252
		case 253:
			goto st253
		case 254:
			goto st254
		case 255:
			goto st255
		case 256:
			goto st256
		case 257:
			goto st257
		case 258:
			goto st258
		case 259:
			goto st259
		case 260:
			goto st260
		case 261:
			goto st261
		case 262:
			goto st262
		case 263:
			goto st263
		case 264:
			goto st264
		case 265:
			goto st265
		case 266:
			goto st266
		case 267:
			goto st267
		case 268:
			goto st268
		case 269:
			goto st269
		case 270:
			goto st270
		case 271:
			goto st271
		case 272:
			goto st272
		case 273:
			goto st273
		case 274:
			goto st274
		case 275:
			goto st275
		case 276:
			goto st276
		case 277:
			goto st277
		case 278:
			goto st278
		case 279:
			goto st279
		case 280:
			goto st280
		case 281:
			goto st281
		case 282:
			goto st282
		case 283:
			goto st283
		case 284:
			goto st284
		case 285:
			goto st285
		case 286:
			goto st286
		case 287:
			goto st287
		case 288:
			goto st288
		case 289:
			goto st289
		case 290:
			goto st290
		case 291:
			goto st291
		case 292:
			goto st292
		case 293:
			goto st293
		case 294:
			goto st294
		case 295:
			goto st295
		case 296:
			goto st296
		case 297:
			goto st297
		case 298:
			goto st298
		case 299:
			goto st299
		case 300:
			goto st300
		case 301:
			goto st301
		case 302:
			goto st302
		case 303:
			goto st303
		case 304:
			goto st304
		case 305:
			goto st305
		case 306:
			goto st306
		case 307:
			goto st307
		case 308:
			goto st308
		case 309:
			goto st309
		case 310:
			goto st310
		case 311:
			goto st311
		case 312:
			goto st312
		case 313:
			goto st313
		case 314:
			goto st314
		case 315:
			goto st315
		case 316:
			goto st316
		case 317:
			goto st317
		case 318:
			goto st318
		case 319:
			goto st319
		case 320:
			goto st320
		case 321:
			goto st321
		case 322:
			goto st322
		case 323:
			goto st323
		case 324:
			goto st324
		case 325:
			goto st325
		case 326:
			goto st326
		case 327:
			goto st327
		case 328:
			goto st328
		case 329:
			goto st329
		case 330:
			goto st330
		case 331:
			goto st331
		case 332:
			goto st332
		case 333:
			goto st333
		case 334:
			goto st334
		case 335:
			goto st335
		case 336:
			goto st336
		case 337:
			goto st337
		case 338:
			goto st338
		case 339:
			goto st339
		case 340:
			goto st340
		case 341:
			goto st341
		case 342:
			goto st342
		case 343:
			goto st343
		case 344:
			goto st344
		case 345:
			goto st345
		case 346:
			goto st346
		case 347:
			goto st347
		case 348:
			goto st348
		case 349:
			goto st349
		case 350:
			goto st350
		case 351:
			goto st351
		case 352:
			goto st352
		case 353:
			goto st353
		case 354:
			goto st354
		case 355:
			goto st355
		case 356:
			goto st356
		case 357:
			goto st357
		case 358:
			goto st358
		case 359:
			goto st359
		case 360:
			goto st360
		case 361:
			goto st361
		case 362:
			goto st362
		case 363:
			goto st363
		case 364:
			goto st364
		case 365:
			goto st365
		case 366:
			goto st366
		case 367:
			goto st367
		case 368:
			goto st368
		case 369:
			goto st369
		case 370:
			goto st370
		case 371:
			goto st371
		case 372:
			goto st372
		case 373:
			goto st373
		case 374:
			goto st374
		case 375:
			goto st375
		case 376:
			goto st376
		case 377:
			goto st377
		case 378:
			goto st378
		case 379:
			goto st379
		case 380:
			goto st380
		case 381:
			goto st381
		case 382:
			goto st382
		case 383:
			goto st383
		case 384:
			goto st384
		case 385:
			goto st385
		case 386:
			goto st386
		case 387:
			goto st387
		case 388:
			goto st388
		case 90:
			goto st90
		case 91:
			goto st91
		case 92:
			goto st92
		case 93:
			goto st93
		case 94:
			goto st94
		case 95:
			goto st95
		case 389:
			goto st389
		case 390:
			goto st390
		case 391:
			goto st391
		case 392:
			goto st392
		case 393:
			goto st393
		case 394:
			goto st394
		case 395:
			goto st395
		case 396:
			goto st396
		case 397:
			goto st397
		case 398:
			goto st398
		case 399:
			goto st399
		case 400:
			goto st400
		case 401:
			goto st401
		case 402:
			goto st402
		case 403:
			goto st403
		case 404:
			goto st404
		case 405:
			goto st405
		case 406:
			goto st406
		case 407:
			goto st407
		case 408:
			goto st408
		case 409:
			goto st409
		case 410:
			goto st410
		case 411:
			goto st411
		case 412:
			goto st412
		case 413:
			goto st413
		case 414:
			goto st414
		case 415:
			goto st415
		case 416:
			goto st416
		case 417:
			goto st417
		case 418:
			goto st418
		case 419:
			goto st419
		case 420:
			goto st420
		case 421:
			goto st421
		case 422:
			goto st422
		case 423:
			goto st423
		case 424:
			goto st424
		case 425:
			goto st425
		case 426:
			goto st426
		case 427:
			goto st427
		case 428:
			goto st428
		case 429:
			goto st429
		case 430:
			goto st430
		case 431:
			goto st431
		case 432:
			goto st432
		case 433:
			goto st433
		case 434:
			goto st434
		case 435:
			goto st435
		case 436:
			goto st436
		case 437:
			goto st437
		case 438:
			goto st438
		case 439:
			goto st439
		case 440:
			goto st440
		case 441:
			goto st441
		case 442:
			goto st442
		case 443:
			goto st443
		case 444:
			goto st444
		case 445:
			goto st445
		case 446:
			goto st446
		case 447:
			goto st447
		case 448:
			goto st448
		case 449:
			goto st449
		case 450:
			goto st450
		case 451:
			goto st451
		case 452:
			goto st452
		case 453:
			goto st453
		case 454:
			goto st454
		case 455:
			goto st455
		case 456:
			goto st456
		case 457:
			goto st457
		case 458:
			goto st458
		case 459:
			goto st459
		case 460:
			goto st460
		case 461:
			goto st461
		case 96:
			goto st96
		case 462:
			goto st462
		case 463:
			goto st463
		case 464:
			goto st464
		case 465:
			goto st465
		case 0:
			goto st0
		case 466:
			goto st466
		case 467:
			goto st467
		case 468:
			goto st468
		case 469:
			goto st469
		case 97:
			goto st97
		case 470:
			goto st470
		case 471:
			goto st471
		case 472:
			goto st472
		case 473:
			goto st473
		case 474:
			goto st474
		case 475:
			goto st475
		case 98:
			goto st98
		case 476:
			goto st476
		case 477:
			goto st477
		case 478:
			goto st478
		case 479:
			goto st479
		case 480:
			goto st480
		case 481:
			goto st481
		case 99:
			goto st99
		case 482:
			goto st482
		case 483:
			goto st483
		case 484:
			goto st484
		case 485:
			goto st485
		case 486:
			goto st486
		case 487:
			goto st487
		case 488:
			goto st488
		case 489:
			goto st489
		case 490:
			goto st490
		case 491:
			goto st491
		case 100:
			goto st100
		case 492:
			goto st492
		case 493:
			goto st493
		case 494:
			goto st494
		case 495:
			goto st495
		case 496:
			goto st496
		case 497:
			goto st497
		case 498:
			goto st498
		case 499:
			goto st499
		case 101:
			goto st101
		case 500:
			goto st500
		case 102:
			goto st102
		case 501:
			goto st501
		case 502:
			goto st502
		case 503:
			goto st503
		case 504:
			goto st504
		case 103:
			goto st103
		case 505:
			goto st505
		case 506:
			goto st506
		case 507:
			goto st507
		case 104:
			goto st104
		case 508:
			goto st508
		case 509:
			goto st509
		case 510:
			goto st510
		case 511:
			goto st511
		case 105:
			goto st105
		case 512:
			goto st512
		case 513:
			goto st513
		case 514:
			goto st514
		case 515:
			goto st515
		case 106:
			goto st106
		case 516:
			goto st516
		case 517:
			goto st517
		case 518:
			goto st518
		case 519:
			goto st519
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 112:
			goto st_case_112
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 113:
			goto st_case_113
		case 4:
			goto st_case_4
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 5:
			goto st_case_5
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 128:
			goto st_case_128
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 66:
			goto st_case_66
		case 134:
			goto st_case_134
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 72:
			goto st_case_72
		case 139:
			goto st_case_139
		case 73:
			goto st_case_73
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 78:
			goto st_case_78
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 147:
			goto st_case_147
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 87:
			goto st_case_87
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 158:
			goto st_case_158
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 166:
			goto st_case_166
		case 167:
			goto st_case_167
		case 168:
			goto st_case_168
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 173:
			goto st_case_173
		case 174:
			goto st_case_174
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 181:
			goto st_case_181
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
		case 192:
			goto st_case_192
		case 193:
			goto st_case_193
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 201:
			goto st_case_201
		case 202:
			goto st_case_202
		case 203:
			goto st_case_203
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 206:
			goto st_case_206
		case 207:
			goto st_case_207
		case 208:
			goto st_case_208
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 213:
			goto st_case_213
		case 214:
			goto st_case_214
		case 215:
			goto st_case_215
		case 216:
			goto st_case_216
		case 217:
			goto st_case_217
		case 218:
			goto st_case_218
		case 219:
			goto st_case_219
		case 220:
			goto st_case_220
		case 221:
			goto st_case_221
		case 222:
			goto st_case_222
		case 223:
			goto st_case_223
		case 224:
			goto st_case_224
		case 225:
			goto st_case_225
		case 226:
			goto st_case_226
		case 227:
			goto st_case_227
		case 228:
			goto st_case_228
		case 229:
			goto st_case_229
		case 230:
			goto st_case_230
		case 231:
			goto st_case_231
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 236:
			goto st_case_236
		case 237:
			goto st_case_237
		case 238:
			goto st_case_238
		case 239:
			goto st_case_239
		case 240:
			goto st_case_240
		case 241:
			goto st_case_241
		case 242:
			goto st_case_242
		case 243:
			goto st_case_243
		case 244:
			goto st_case_244
		case 245:
			goto st_case_245
		case 246:
			goto st_case_246
		case 247:
			goto st_case_247
		case 248:
			goto st_case_248
		case 249:
			goto st_case_249
		case 250:
			goto st_case_250
		case 251:
			goto st_case_251
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 256:
			goto st_case_256
		case 257:
			goto st_case_257
		case 258:
			goto st_case_258
		case 259:
			goto st_case_259
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 262:
			goto st_case_262
		case 263:
			goto st_case_263
		case 264:
			goto st_case_264
		case 265:
			goto st_case_265
		case 266:
			goto st_case_266
		case 267:
			goto st_case_267
		case 268:
			goto st_case_268
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
		case 271:
			goto st_case_271
		case 272:
			goto st_case_272
		case 273:
			goto st_case_273
		case 274:
			goto st_case_274
		case 275:
			goto st_case_275
		case 276:
			goto st_case_276
		case 277:
			goto st_case_277
		case 278:
			goto st_case_278
		case 279:
			goto st_case_279
		case 280:
			goto st_case_280
		case 281:
			goto st_case_281
		case 282:
			goto st_case_282
		case 283:
			goto st_case_283
		case 284:
			goto st_case_284
		case 285:
			goto st_case_285
		case 286:
			goto st_case_286
		case 287:
			goto st_case_287
		case 288:
			goto st_case_288
		case 289:
			goto st_case_289
		case 290:
			goto st_case_290
		case 291:
			goto st_case_291
		case 292:
			goto st_case_292
		case 293:
			goto st_case_293
		case 294:
			goto st_case_294
		case 295:
			goto st_case_295
		case 296:
			goto st_case_296
		case 297:
			goto st_case_297
		case 298:
			goto st_case_298
		case 299:
			goto st_case_299
		case 300:
			goto st_case_300
		case 301:
			goto st_case_301
		case 302:
			goto st_case_302
		case 303:
			goto st_case_303
		case 304:
			goto st_case_304
		case 305:
			goto st_case_305
		case 306:
			goto st_case_306
		case 307:
			goto st_case_307
		case 308:
			goto st_case_308
		case 309:
			goto st_case_309
		case 310:
			goto st_case_310
		case 311:
			goto st_case_311
		case 312:
			goto st_case_312
		case 313:
			goto st_case_313
		case 314:
			goto st_case_314
		case 315:
			goto st_case_315
		case 316:
			goto st_case_316
		case 317:
			goto st_case_317
		case 318:
			goto st_case_318
		case 319:
			goto st_case_319
		case 320:
			goto st_case_320
		case 321:
			goto st_case_321
		case 322:
			goto st_case_322
		case 323:
			goto st_case_323
		case 324:
			goto st_case_324
		case 325:
			goto st_case_325
		case 326:
			goto st_case_326
		case 327:
			goto st_case_327
		case 328:
			goto st_case_328
		case 329:
			goto st_case_329
		case 330:
			goto st_case_330
		case 331:
			goto st_case_331
		case 332:
			goto st_case_332
		case 333:
			goto st_case_333
		case 334:
			goto st_case_334
		case 335:
			goto st_case_335
		case 336:
			goto st_case_336
		case 337:
			goto st_case_337
		case 338:
			goto st_case_338
		case 339:
			goto st_case_339
		case 340:
			goto st_case_340
		case 341:
			goto st_case_341
		case 342:
			goto st_case_342
		case 343:
			goto st_case_343
		case 344:
			goto st_case_344
		case 345:
			goto st_case_345
		case 346:
			goto st_case_346
		case 347:
			goto st_case_347
		case 348:
			goto st_case_348
		case 349:
			goto st_case_349
		case 350:
			goto st_case_350
		case 351:
			goto st_case_351
		case 352:
			goto st_case_352
		case 353:
			goto st_case_353
		case 354:
			goto st_case_354
		case 355:
			goto st_case_355
		case 356:
			goto st_case_356
		case 357:
			goto st_case_357
		case 358:
			goto st_case_358
		case 359:
			goto st_case_359
		case 360:
			goto st_case_360
		case 361:
			goto st_case_361
		case 362:
			goto st_case_362
		case 363:
			goto st_case_363
		case 364:
			goto st_case_364
		case 365:
			goto st_case_365
		case 366:
			goto st_case_366
		case 367:
			goto st_case_367
		case 368:
			goto st_case_368
		case 369:
			goto st_case_369
		case 370:
			goto st_case_370
		case 371:
			goto st_case_371
		case 372:
			goto st_case_372
		case 373:
			goto st_case_373
		case 374:
			goto st_case_374
		case 375:
			goto st_case_375
		case 376:
			goto st_case_376
		case 377:
			goto st_case_377
		case 378:
			goto st_case_378
		case 379:
			goto st_case_379
		case 380:
			goto st_case_380
		case 381:
			goto st_case_381
		case 382:
			goto st_case_382
		case 383:
			goto st_case_383
		case 384:
			goto st_case_384
		case 385:
			goto st_case_385
		case 386:
			goto st_case_386
		case 387:
			goto st_case_387
		case 388:
			goto st_case_388
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 389:
			goto st_case_389
		case 390:
			goto st_case_390
		case 391:
			goto st_case_391
		case 392:
			goto st_case_392
		case 393:
			goto st_case_393
		case 394:
			goto st_case_394
		case 395:
			goto st_case_395
		case 396:
			goto st_case_396
		case 397:
			goto st_case_397
		case 398:
			goto st_case_398
		case 399:
			goto st_case_399
		case 400:
			goto st_case_400
		case 401:
			goto st_case_401
		case 402:
			goto st_case_402
		case 403:
			goto st_case_403
		case 404:
			goto st_case_404
		case 405:
			goto st_case_405
		case 406:
			goto st_case_406
		case 407:
			goto st_case_407
		case 408:
			goto st_case_408
		case 409:
			goto st_case_409
		case 410:
			goto st_case_410
		case 411:
			goto st_case_411
		case 412:
			goto st_case_412
		case 413:
			goto st_case_413
		case 414:
			goto st_case_414
		case 415:
			goto st_case_415
		case 416:
			goto st_case_416
		case 417:
			goto st_case_417
		case 418:
			goto st_case_418
		case 419:
			goto st_case_419
		case 420:
			goto st_case_420
		case 421:
			goto st_case_421
		case 422:
			goto st_case_422
		case 423:
			goto st_case_423
		case 424:
			goto st_case_424
		case 425:
			goto st_case_425
		case 426:
			goto st_case_426
		case 427:
			goto st_case_427
		case 428:
			goto st_case_428
		case 429:
			goto st_case_429
		case 430:
			goto st_case_430
		case 431:
			goto st_case_431
		case 432:
			goto st_case_432
		case 433:
			goto st_case_433
		case 434:
			goto st_case_434
		case 435:
			goto st_case_435
		case 436:
			goto st_case_436
		case 437:
			goto st_case_437
		case 438:
			goto st_case_438
		case 439:
			goto st_case_439
		case 440:
			goto st_case_440
		case 441:
			goto st_case_441
		case 442:
			goto st_case_442
		case 443:
			goto st_case_443
		case 444:
			goto st_case_444
		case 445:
			goto st_case_445
		case 446:
			goto st_case_446
		case 447:
			goto st_case_447
		case 448:
			goto st_case_448
		case 449:
			goto st_case_449
		case 450:
			goto st_case_450
		case 451:
			goto st_case_451
		case 452:
			goto st_case_452
		case 453:
			goto st_case_453
		case 454:
			goto st_case_454
		case 455:
			goto st_case_455
		case 456:
			goto st_case_456
		case 457:
			goto st_case_457
		case 458:
			goto st_case_458
		case 459:
			goto st_case_459
		case 460:
			goto st_case_460
		case 461:
			goto st_case_461
		case 96:
			goto st_case_96
		case 462:
			goto st_case_462
		case 463:
			goto st_case_463
		case 464:
			goto st_case_464
		case 465:
			goto st_case_465
		case 0:
			goto st_case_0
		case 466:
			goto st_case_466
		case 467:
			goto st_case_467
		case 468:
			goto st_case_468
		case 469:
			goto st_case_469
		case 97:
			goto st_case_97
		case 470:
			goto st_case_470
		case 471:
			goto st_case_471
		case 472:
			goto st_case_472
		case 473:
			goto st_case_473
		case 474:
			goto st_case_474
		case 475:
			goto st_case_475
		case 98:
			goto st_case_98
		case 476:
			goto st_case_476
		case 477:
			goto st_case_477
		case 478:
			goto st_case_478
		case 479:
			goto st_case_479
		case 480:
			goto st_case_480
		case 481:
			goto st_case_481
		case 99:
			goto st_case_99
		case 482:
			goto st_case_482
		case 483:
			goto st_case_483
		case 484:
			goto st_case_484
		case 485:
			goto st_case_485
		case 486:
			goto st_case_486
		case 487:
			goto st_case_487
		case 488:
			goto st_case_488
		case 489:
			goto st_case_489
		case 490:
			goto st_case_490
		case 491:
			goto st_case_491
		case 100:
			goto st_case_100
		case 492:
			goto st_case_492
		case 493:
			goto st_case_493
		case 494:
			goto st_case_494
		case 495:
			goto st_case_495
		case 496:
			goto st_case_496
		case 497:
			goto st_case_497
		case 498:
			goto st_case_498
		case 499:
			goto st_case_499
		case 101:
			goto st_case_101
		case 500:
			goto st_case_500
		case 102:
			goto st_case_102
		case 501:
			goto st_case_501
		case 502:
			goto st_case_502
		case 503:
			goto st_case_503
		case 504:
			goto st_case_504
		case 103:
			goto st_case_103
		case 505:
			goto st_case_505
		case 506:
			goto st_case_506
		case 507:
			goto st_case_507
		case 104:
			goto st_case_104
		case 508:
			goto st_case_508
		case 509:
			goto st_case_509
		case 510:
			goto st_case_510
		case 511:
			goto st_case_511
		case 105:
			goto st_case_105
		case 512:
			goto st_case_512
		case 513:
			goto st_case_513
		case 514:
			goto st_case_514
		case 515:
			goto st_case_515
		case 106:
			goto st_case_106
		case 516:
			goto st_case_516
		case 517:
			goto st_case_517
		case 518:
			goto st_case_518
		case 519:
			goto st_case_519
		}
		goto st_out
	tr0:
		lex.cs = 107
//line scanner/scanner.rl:141
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			lex.cs = 114
		}
		goto _again
	tr3:
		lex.cs = 107
//line scanner/scanner.rl:145
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.ts+5)
			lex.cs = 114
		}
		goto _again
	tr158:
//line scanner/scanner.rl:135
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(token)
			tok = T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 107
				goto _out
			}
		}
		goto st107
	tr160:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:135
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(token)
			tok = T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 107
				goto _out
			}
		}
		goto st107
	tr165:
		lex.cs = 107
//line scanner/scanner.rl:141
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			lex.cs = 114
		}
		goto _again
	tr166:
		lex.cs = 107
//line scanner/scanner.rl:150
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ECHO
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr168:
		lex.cs = 107
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:145
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.ts+5)
			lex.cs = 114
		}
		goto _again
	st107:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof107
		}
	st_case_107:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:2263
		switch lex.data[(lex.p)] {
		case 10:
			goto st109
		case 60:
			goto st111
		}
		goto st108
	tr161:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st108
	st108:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof108
		}
	st_case_108:
//line scanner/scanner.go:2280
		switch lex.data[(lex.p)] {
		case 10:
			goto st109
		case 60:
			goto st110
		}
		goto st108
	tr162:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st109
	st109:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof109
		}
	st_case_109:
//line scanner/scanner.go:2297
		switch lex.data[(lex.p)] {
		case 10:
			goto tr162
		case 60:
			goto tr163
		}
		goto tr161
	tr163:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st110
	st110:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof110
		}
	st_case_110:
//line scanner/scanner.go:2314
		switch lex.data[(lex.p)] {
		case 10:
			goto st109
		case 60:
			goto st110
		case 63:
			goto tr158
		}
		goto st108
	st111:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof111
		}
	st_case_111:
		switch lex.data[(lex.p)] {
		case 10:
			goto st109
		case 60:
			goto st110
		case 63:
			goto tr164
		}
		goto st108
	tr164:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st112
	st112:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof112
		}
	st_case_112:
//line scanner/scanner.go:2348
		switch lex.data[(lex.p)] {
		case 61:
			goto tr166
		case 80:
			goto st1
		case 112:
			goto st1
		}
		goto tr165
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch lex.data[(lex.p)] {
		case 72:
			goto st2
		case 104:
			goto st2
		}
		goto tr0
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 80:
			goto st3
		case 112:
			goto st3
		}
		goto tr0
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 9:
			goto tr3
		case 10:
			goto st113
		case 13:
			goto st4
		case 32:
			goto tr3
		}
		goto tr0
	st113:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		goto tr168
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		if lex.data[(lex.p)] == 10 {
			goto st113
		}
		goto tr0
	tr6:
//line scanner/scanner.rl:159
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st114
	tr8:
		lex.cs = 114
//line NONE:1
		switch lex.act {
		case 8:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 10:
			{
				(lex.p) = (lex.te) - 1

				if lex.te-lex.ts < 20 {
					lex.setTokenPosition(token)
					tok = T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}
				lex.setTokenPosition(token)
				tok = T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 12:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ABSTRACT
				{
					(lex.p)++
					goto _out
				}
			}
		case 13:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ARRAY
				{
					(lex.p)++
					goto _out
				}
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_AS
				{
					(lex.p)++
					goto _out
				}
			}
		case 15:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_BREAK
				{
					(lex.p)++
					goto _out
				}
			}
		case 16:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CALLABLE
				{
					(lex.p)++
					goto _out
				}
			}
		case 17:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CASE
				{
					(lex.p)++
					goto _out
				}
			}
		case 18:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CATCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 19:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CLASS
				{
					(lex.p)++
					goto _out
				}
			}
		case 20:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CLONE
				{
					(lex.p)++
					goto _out
				}
			}
		case 21:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CONST
				{
					(lex.p)++
					goto _out
				}
			}
		case 22:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CONTINUE
				{
					(lex.p)++
					goto _out
				}
			}
		case 23:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 24:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DEFAULT
				{
					(lex.p)++
					goto _out
				}
			}
		case 25:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DO
				{
					(lex.p)++
					goto _out
				}
			}
		case 26:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ECHO
				{
					(lex.p)++
					goto _out
				}
			}
		case 28:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ELSEIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 29:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EMPTY
				{
					(lex.p)++
					goto _out
				}
			}
		case 30:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDDECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 32:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDFOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 33:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 34:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDSWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 35:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDWHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 36:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EVAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 37:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EXIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 38:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EXTENDS
				{
					(lex.p)++
					goto _out
				}
			}
		case 40:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FINALLY
				{
					(lex.p)++
					goto _out
				}
			}
		case 42:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 43:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FUNCTION
				{
					(lex.p)++
					goto _out
				}
			}
		case 44:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_GLOBAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 45:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_GOTO
				{
					(lex.p)++
					goto _out
				}
			}
		case 46:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_IF
				{
					(lex.p)++
					goto _out
				}
			}
		case 47:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ISSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 48:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_IMPLEMENTS
				{
					(lex.p)++
					goto _out
				}
			}
		case 49:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INSTANCEOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 50:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INSTEADOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 51:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INTERFACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 52:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LIST
				{
					(lex.p)++
					goto _out
				}
			}
		case 53:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_NAMESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 54:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PRIVATE
				{
					(lex.p)++
					goto _out
				}
			}
		case 55:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PUBLIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 56:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PRINT
				{
					(lex.p)++
					goto _out
				}
			}
		case 57:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PROTECTED
				{
					(lex.p)++
					goto _out
				}
			}
		case 58:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_RETURN
				{
					(lex.p)++
					goto _out
				}
			}
		case 59:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_STATIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 60:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_SWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 61:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_THROW
				{
					(lex.p)++
					goto _out
				}
			}
		case 62:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_TRAIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 63:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_TRY
				{
					(lex.p)++
					goto _out
				}
			}
		case 64:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_UNSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 65:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_USE
				{
					(lex.p)++
					goto _out
				}
			}
		case 66:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_VAR
				{
					(lex.p)++
					goto _out
				}
			}
		case 67:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_WHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 68:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_YIELD_FROM
				{
					(lex.p)++
					goto _out
				}
			}
		case 71:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INCLUDE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 73:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_REQUIRE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 74:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CLASS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 75:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DIR
				{
					(lex.p)++
					goto _out
				}
			}
		case 76:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 77:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FUNC_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 78:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LINE
				{
					(lex.p)++
					goto _out
				}
			}
		case 79:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_NS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 80:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_METHOD_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 81:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_TRAIT_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 82:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_HALT_COMPILER
				lex.cs = 505
				{
					(lex.p)++
					goto _out
				}
			}
		case 83:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_NEW
				{
					(lex.p)++
					goto _out
				}
			}
		case 84:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LOGICAL_AND
				{
					(lex.p)++
					goto _out
				}
			}
		case 85:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LOGICAL_OR
				{
					(lex.p)++
					goto _out
				}
			}
		case 86:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LOGICAL_XOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 115:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_SL
				{
					(lex.p)++
					goto _out
				}
			}
		case 131:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_STRING
				{
					(lex.p)++
					goto _out
				}
			}
		case 136:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = TokenID(int('"'))
				lex.cs = 480
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr11:
//line scanner/scanner.rl:344
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_CONSTANT_ENCAPSED_STRING
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr15:
//line scanner/scanner.rl:367
		(lex.p) = (lex.te) - 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st114
	tr19:
//line scanner/scanner.rl:329
		(lex.p) = (lex.te) - 1
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr34:
//line scanner/scanner.rl:305
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ARRAY_CAST
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr41:
//line scanner/scanner.rl:310
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_STRING_CAST
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr45:
//line scanner/scanner.rl:306
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_BOOL_CAST
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr53:
//line scanner/scanner.rl:307
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOUBLE_CAST
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr60:
//line scanner/scanner.rl:308
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_INT_CAST
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr69:
//line scanner/scanner.rl:309
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OBJECT_CAST
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr80:
//line scanner/scanner.rl:311
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_UNSET_CAST
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr81:
//line scanner/scanner.rl:274
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ELLIPSIS
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr90:
//line scanner/scanner.rl:317
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
				isDocComment = true
			}
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)

			if isDocComment {
				lex.PhpDocComment = string(lex.data[lex.ts:lex.te])
			}
		}
		goto st114
	tr91:
//line scanner/scanner.rl:177
		(lex.p) = (lex.te) - 1
		{
			if lex.te-lex.ts < 20 {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 114
					goto _out
				}
			}
			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr103:
		lex.cs = 114
//line scanner/scanner.rl:161
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 107
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr119:
		lex.cs = 114
//line scanner/scanner.rl:160
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 107
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr121:
//line scanner/scanner.rl:340
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr123:
//line scanner/scanner.rl:255
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = T_YIELD
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr134:
//line scanner/scanner.rl:254
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_YIELD_FROM
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr169:
//line scanner/scanner.rl:367
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st114
	tr180:
//line scanner/scanner.rl:329
		lex.te = (lex.p) + 1
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr215:
//line scanner/scanner.rl:273
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_NS_SEPARATOR
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr218:
		lex.cs = 114
//line scanner/scanner.rl:364
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('`'))
			lex.cs = 474
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr219:
//line scanner/scanner.rl:337
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('{'))
			lex.call(114, 114)
			goto _out
		}
		goto st114
	tr221:
//line scanner/scanner.rl:338
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('}'))
			lex.ret(1)
			lex.PhpDocComment = ""
			goto _out
		}
		goto st114
	tr222:
//line scanner/scanner.rl:159
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st114
	tr224:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:159
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st114
	tr228:
//line scanner/scanner.rl:367
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st114
	tr229:
//line scanner/scanner.rl:329
		lex.te = (lex.p)
		(lex.p)--
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr231:
//line scanner/scanner.rl:292
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr232:
//line scanner/scanner.rl:293
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_NOT_IDENTICAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr233:
		lex.cs = 114
//line scanner/scanner.rl:365
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('"'))
			lex.cs = 480
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr234:
//line scanner/scanner.rl:313
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
		}
		goto st114
	tr236:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:313
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
		}
		goto st114
	tr240:
//line scanner/scanner.rl:339
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr241:
//line scanner/scanner.rl:287
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_MOD_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr242:
//line scanner/scanner.rl:276
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_BOOLEAN_AND
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr243:
//line scanner/scanner.rl:278
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_AND_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr245:
//line scanner/scanner.rl:281
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_MUL_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr246:
//line scanner/scanner.rl:300
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_POW
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr247:
//line scanner/scanner.rl:282
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_POW_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr248:
//line scanner/scanner.rl:289
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_INC
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr249:
//line scanner/scanner.rl:284
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_PLUS_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr250:
//line scanner/scanner.rl:288
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DEC
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr251:
//line scanner/scanner.rl:285
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_MINUS_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr252:
		lex.cs = 114
//line scanner/scanner.rl:342
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OBJECT_OPERATOR
			lex.cs = 459
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr255:
//line scanner/scanner.rl:280
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_CONCAT_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr256:
//line scanner/scanner.rl:163
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr258:
//line scanner/scanner.rl:283
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DIV_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr259:
//line scanner/scanner.rl:177
		lex.te = (lex.p)
		(lex.p)--
		{
			if lex.te-lex.ts < 20 {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 114
					goto _out
				}
			}
			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr262:
//line scanner/scanner.rl:164
		lex.te = (lex.p)
		(lex.p)--
		{
			firstNum := 2
			for i := lex.ts + 2; i < lex.te; i++ {
				if lex.data[i] == '0' {
					firstNum++
				}
			}

			if lex.te-lex.ts-firstNum < 64 {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 114
					goto _out
				}
			}
			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr263:
//line scanner/scanner.rl:183
		lex.te = (lex.p)
		(lex.p)--
		{
			firstNum := lex.ts + 2
			for i := lex.ts + 2; i < lex.te; i++ {
				if lex.data[i] == '0' {
					firstNum++
				}
			}

			length := lex.te - firstNum
			if length < 16 || (length == 16 && lex.data[firstNum] <= '7') {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 114
					goto _out
				}
			}
			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr264:
//line scanner/scanner.rl:275
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_PAAMAYIM_NEKUDOTAYIM
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr265:
		lex.cs = 114
//line scanner/scanner.rl:161
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 107
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr267:
		lex.cs = 114
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:161
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 107
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr270:
//line scanner/scanner.rl:292
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr271:
//line scanner/scanner.rl:301
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_SL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr272:
//line scanner/scanner.rl:296
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_SL_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr273:
		lex.cs = 114
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:350
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.heredocLabel = lex.data[lblStart:lblEnd]
			lex.setTokenPosition(token)
			tok = T_START_HEREDOC

			if lex.isHeredocEnd(lex.p + 1) {
				lex.cs = 486
			} else if lex.data[lblStart-1] == '\'' {
				lex.cs = 465
			} else {
				lex.cs = 468
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr274:
//line scanner/scanner.rl:299
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_IS_SMALLER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr275:
//line scanner/scanner.rl:291
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_SPACESHIP
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr277:
//line scanner/scanner.rl:290
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOUBLE_ARROW
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr278:
//line scanner/scanner.rl:294
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_IS_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr279:
//line scanner/scanner.rl:295
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_IDENTICAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr280:
//line scanner/scanner.rl:298
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_GREATER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr282:
//line scanner/scanner.rl:302
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_SR
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr283:
//line scanner/scanner.rl:297
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_SR_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr285:
//line scanner/scanner.rl:303
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_COALESCE
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr286:
		lex.cs = 114
//line scanner/scanner.rl:160
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 107
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr288:
		lex.cs = 114
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:160
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 107
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr289:
//line scanner/scanner.rl:340
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr369:
//line scanner/scanner.rl:213
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ELSE
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr389:
//line scanner/scanner.rl:217
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENDFOR
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr417:
//line scanner/scanner.rl:225
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_FINAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr421:
//line scanner/scanner.rl:227
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_FOR
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr453:
//line scanner/scanner.rl:256
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_INCLUDE
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr521:
//line scanner/scanner.rl:258
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_REQUIRE
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr567:
//line scanner/scanner.rl:255
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_YIELD
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr572:
//line scanner/scanner.rl:286
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_XOR_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr646:
//line scanner/scanner.rl:279
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr647:
//line scanner/scanner.rl:277
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_BOOLEAN_OR
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	st114:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof114
		}
	st_case_114:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:3206
		switch lex.data[(lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto st117
		case 32:
			goto tr170
		case 33:
			goto st118
		case 34:
			goto tr173
		case 35:
			goto st121
		case 36:
			goto st123
		case 37:
			goto st125
		case 38:
			goto st126
		case 39:
			goto tr178
		case 40:
			goto tr179
		case 42:
			goto st129
		case 43:
			goto st131
		case 45:
			goto st132
		case 46:
			goto tr184
		case 47:
			goto tr185
		case 48:
			goto tr186
		case 58:
			goto st141
		case 59:
			goto tr189
		case 60:
			goto st145
		case 61:
			goto st149
		case 62:
			goto st151
		case 63:
			goto st153
		case 64:
			goto tr180
		case 65:
			goto st156
		case 66:
			goto tr195
		case 67:
			goto st172
		case 68:
			goto st201
		case 69:
			goto st212
		case 70:
			goto st254
		case 71:
			goto st265
		case 73:
			goto st272
		case 76:
			goto st311
		case 78:
			goto st314
		case 79:
			goto st323
		case 80:
			goto st324
		case 82:
			goto st341
		case 83:
			goto st355
		case 84:
			goto st364
		case 85:
			goto st371
		case 86:
			goto st376
		case 87:
			goto st378
		case 88:
			goto st382
		case 89:
			goto st384
		case 92:
			goto tr215
		case 94:
			goto st392
		case 95:
			goto st393
		case 96:
			goto tr218
		case 97:
			goto st156
		case 98:
			goto tr195
		case 99:
			goto st172
		case 100:
			goto st201
		case 101:
			goto st212
		case 102:
			goto st254
		case 103:
			goto st265
		case 105:
			goto st272
		case 108:
			goto st311
		case 110:
			goto st314
		case 111:
			goto st323
		case 112:
			goto st324
		case 114:
			goto st341
		case 115:
			goto st355
		case 116:
			goto st364
		case 117:
			goto st371
		case 118:
			goto st376
		case 119:
			goto st378
		case 120:
			goto st382
		case 121:
			goto st384
		case 123:
			goto tr219
		case 124:
			goto st458
		case 125:
			goto tr221
		case 126:
			goto tr180
		case 127:
			goto tr169
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr170
				}
			default:
				goto tr169
			}
		case lex.data[(lex.p)] > 31:
			switch {
			case lex.data[(lex.p)] < 49:
				if 41 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 44 {
					goto tr180
				}
			case lex.data[(lex.p)] > 57:
				if 91 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 93 {
					goto tr180
				}
			default:
				goto tr187
			}
		default:
			goto tr169
		}
		goto tr201
	tr170:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st115
	tr225:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st115
	st115:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof115
		}
	st_case_115:
//line scanner/scanner.go:3399
		switch lex.data[(lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto st5
		case 32:
			goto tr170
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr170
		}
		goto tr222
	tr7:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st116
	tr226:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st116
	st116:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof116
		}
	st_case_116:
//line scanner/scanner.go:3429
		switch lex.data[(lex.p)] {
		case 10:
			goto tr226
		case 13:
			goto tr227
		case 32:
			goto tr225
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr225
		}
		goto tr224
	tr227:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line scanner/scanner.go:3451
		if lex.data[(lex.p)] == 10 {
			goto tr7
		}
		goto tr6
	st117:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		if lex.data[(lex.p)] == 10 {
			goto tr7
		}
		goto tr228
	st118:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof118
		}
	st_case_118:
		if lex.data[(lex.p)] == 61 {
			goto st119
		}
		goto tr229
	st119:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof119
		}
	st_case_119:
		if lex.data[(lex.p)] == 61 {
			goto tr232
		}
		goto tr231
	tr173:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:365
		lex.act = 136
		goto st120
	st120:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof120
		}
	st_case_120:
//line scanner/scanner.go:3495
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto tr10
		case 34:
			goto tr11
		case 36:
			goto st7
		case 92:
			goto st8
		case 123:
			goto st9
		}
		goto st6
	tr10:
//line scanner/scanner.rl:48

		if lex.data[lex.p] == '\n' {
			lex.NewLines.Append(lex.p)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}

		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line scanner/scanner.go:3528
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto tr10
		case 34:
			goto tr11
		case 36:
			goto st7
		case 92:
			goto st8
		case 123:
			goto st9
		}
		goto st6
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto tr10
		case 34:
			goto tr11
		case 92:
			goto st8
		case 96:
			goto st6
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st6
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st6
			}
		default:
			goto st6
		}
		goto tr8
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto tr10
		}
		goto st6
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto tr10
		case 34:
			goto tr11
		case 36:
			goto tr8
		}
		goto st6
	tr237:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st121
	st121:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof121
		}
	st_case_121:
//line scanner/scanner.go:3611
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		if _widec == 522 {
			goto st122
		}
		if 512 <= _widec && _widec <= 767 {
			goto st121
		}
		goto tr234
	tr238:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st122
	st122:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof122
		}
	st_case_122:
//line scanner/scanner.go:3666
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		if _widec == 522 {
			goto tr238
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr237
		}
		goto tr236
	st123:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof123
		}
	st_case_123:
		if lex.data[(lex.p)] == 96 {
			goto tr229
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr229
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr229
			}
		default:
			goto tr229
		}
		goto st124
	st124:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		if lex.data[(lex.p)] == 96 {
			goto tr240
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr240
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr240
				}
			case lex.data[(lex.p)] >= 91:
				goto tr240
			}
		default:
			goto tr240
		}
		goto st124
	st125:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof125
		}
	st_case_125:
		if lex.data[(lex.p)] == 61 {
			goto tr241
		}
		goto tr229
	st126:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof126
		}
	st_case_126:
		switch lex.data[(lex.p)] {
		case 38:
			goto tr242
		case 61:
			goto tr243
		}
		goto tr229
	tr178:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st127
	st127:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof127
		}
	st_case_127:
//line scanner/scanner.go:3790
		switch lex.data[(lex.p)] {
		case 10:
			goto tr17
		case 13:
			goto tr17
		case 39:
			goto tr11
		case 92:
			goto st11
		}
		goto st10
	tr17:
//line scanner/scanner.rl:48

		if lex.data[lex.p] == '\n' {
			lex.NewLines.Append(lex.p)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}

		goto st10
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
//line scanner/scanner.go:3819
		switch lex.data[(lex.p)] {
		case 10:
			goto tr17
		case 13:
			goto tr17
		case 39:
			goto tr11
		case 92:
			goto st11
		}
		goto st10
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr17
		case 13:
			goto tr17
		}
		goto st10
	tr179:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st128
	st128:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof128
		}
	st_case_128:
//line scanner/scanner.go:3853
		switch lex.data[(lex.p)] {
		case 9:
			goto st12
		case 32:
			goto st12
		case 65:
			goto st13
		case 66:
			goto st18
		case 68:
			goto st30
		case 70:
			goto st36
		case 73:
			goto st40
		case 79:
			goto st47
		case 82:
			goto st53
		case 83:
			goto st56
		case 85:
			goto st61
		case 97:
			goto st13
		case 98:
			goto st18
		case 100:
			goto st30
		case 102:
			goto st36
		case 105:
			goto st40
		case 111:
			goto st47
		case 114:
			goto st53
		case 115:
			goto st56
		case 117:
			goto st61
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st12
		}
		goto tr229
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch lex.data[(lex.p)] {
		case 9:
			goto st12
		case 32:
			goto st12
		case 65:
			goto st13
		case 66:
			goto st18
		case 68:
			goto st30
		case 70:
			goto st36
		case 73:
			goto st40
		case 79:
			goto st47
		case 82:
			goto st53
		case 83:
			goto st56
		case 85:
			goto st61
		case 97:
			goto st13
		case 98:
			goto st18
		case 100:
			goto st30
		case 102:
			goto st36
		case 105:
			goto st40
		case 111:
			goto st47
		case 114:
			goto st53
		case 115:
			goto st56
		case 117:
			goto st61
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st12
		}
		goto tr19
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch lex.data[(lex.p)] {
		case 82:
			goto st14
		case 114:
			goto st14
		}
		goto tr19
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch lex.data[(lex.p)] {
		case 82:
			goto st15
		case 114:
			goto st15
		}
		goto tr19
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 65:
			goto st16
		case 97:
			goto st16
		}
		goto tr19
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch lex.data[(lex.p)] {
		case 89:
			goto st17
		case 121:
			goto st17
		}
		goto tr19
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch lex.data[(lex.p)] {
		case 9:
			goto st17
		case 32:
			goto st17
		case 41:
			goto tr34
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st17
		}
		goto tr19
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 73:
			goto st19
		case 79:
			goto st24
		case 105:
			goto st19
		case 111:
			goto st24
		}
		goto tr19
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 78:
			goto st20
		case 110:
			goto st20
		}
		goto tr19
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 65:
			goto st21
		case 97:
			goto st21
		}
		goto tr19
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch lex.data[(lex.p)] {
		case 82:
			goto st22
		case 114:
			goto st22
		}
		goto tr19
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch lex.data[(lex.p)] {
		case 89:
			goto st23
		case 121:
			goto st23
		}
		goto tr19
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 9:
			goto st23
		case 32:
			goto st23
		case 41:
			goto tr41
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st23
		}
		goto tr19
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 79:
			goto st25
		case 111:
			goto st25
		}
		goto tr19
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 76:
			goto st26
		case 108:
			goto st26
		}
		goto tr19
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 9:
			goto st27
		case 32:
			goto st27
		case 41:
			goto tr45
		case 69:
			goto st28
		case 101:
			goto st28
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st27
		}
		goto tr19
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 9:
			goto st27
		case 32:
			goto st27
		case 41:
			goto tr45
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st27
		}
		goto tr19
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 65:
			goto st29
		case 97:
			goto st29
		}
		goto tr19
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 78:
			goto st27
		case 110:
			goto st27
		}
		goto tr19
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 79:
			goto st31
		case 111:
			goto st31
		}
		goto tr19
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 85:
			goto st32
		case 117:
			goto st32
		}
		goto tr19
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 66:
			goto st33
		case 98:
			goto st33
		}
		goto tr19
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 76:
			goto st34
		case 108:
			goto st34
		}
		goto tr19
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 69:
			goto st35
		case 101:
			goto st35
		}
		goto tr19
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 9:
			goto st35
		case 32:
			goto st35
		case 41:
			goto tr53
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st35
		}
		goto tr19
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 76:
			goto st37
		case 108:
			goto st37
		}
		goto tr19
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch lex.data[(lex.p)] {
		case 79:
			goto st38
		case 111:
			goto st38
		}
		goto tr19
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 65:
			goto st39
		case 97:
			goto st39
		}
		goto tr19
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 84:
			goto st35
		case 116:
			goto st35
		}
		goto tr19
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 78:
			goto st41
		case 110:
			goto st41
		}
		goto tr19
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 84:
			goto st42
		case 116:
			goto st42
		}
		goto tr19
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch lex.data[(lex.p)] {
		case 9:
			goto st43
		case 32:
			goto st43
		case 41:
			goto tr60
		case 69:
			goto st44
		case 101:
			goto st44
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st43
		}
		goto tr19
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch lex.data[(lex.p)] {
		case 9:
			goto st43
		case 32:
			goto st43
		case 41:
			goto tr60
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st43
		}
		goto tr19
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch lex.data[(lex.p)] {
		case 71:
			goto st45
		case 103:
			goto st45
		}
		goto tr19
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 69:
			goto st46
		case 101:
			goto st46
		}
		goto tr19
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch lex.data[(lex.p)] {
		case 82:
			goto st43
		case 114:
			goto st43
		}
		goto tr19
	st47:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch lex.data[(lex.p)] {
		case 66:
			goto st48
		case 98:
			goto st48
		}
		goto tr19
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch lex.data[(lex.p)] {
		case 74:
			goto st49
		case 106:
			goto st49
		}
		goto tr19
	st49:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch lex.data[(lex.p)] {
		case 69:
			goto st50
		case 101:
			goto st50
		}
		goto tr19
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch lex.data[(lex.p)] {
		case 67:
			goto st51
		case 99:
			goto st51
		}
		goto tr19
	st51:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch lex.data[(lex.p)] {
		case 84:
			goto st52
		case 116:
			goto st52
		}
		goto tr19
	st52:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch lex.data[(lex.p)] {
		case 9:
			goto st52
		case 32:
			goto st52
		case 41:
			goto tr69
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st52
		}
		goto tr19
	st53:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch lex.data[(lex.p)] {
		case 69:
			goto st54
		case 101:
			goto st54
		}
		goto tr19
	st54:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch lex.data[(lex.p)] {
		case 65:
			goto st55
		case 97:
			goto st55
		}
		goto tr19
	st55:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch lex.data[(lex.p)] {
		case 76:
			goto st35
		case 108:
			goto st35
		}
		goto tr19
	st56:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch lex.data[(lex.p)] {
		case 84:
			goto st57
		case 116:
			goto st57
		}
		goto tr19
	st57:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch lex.data[(lex.p)] {
		case 82:
			goto st58
		case 114:
			goto st58
		}
		goto tr19
	st58:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch lex.data[(lex.p)] {
		case 73:
			goto st59
		case 105:
			goto st59
		}
		goto tr19
	st59:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch lex.data[(lex.p)] {
		case 78:
			goto st60
		case 110:
			goto st60
		}
		goto tr19
	st60:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch lex.data[(lex.p)] {
		case 71:
			goto st23
		case 103:
			goto st23
		}
		goto tr19
	st61:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch lex.data[(lex.p)] {
		case 78:
			goto st62
		case 110:
			goto st62
		}
		goto tr19
	st62:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch lex.data[(lex.p)] {
		case 83:
			goto st63
		case 115:
			goto st63
		}
		goto tr19
	st63:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch lex.data[(lex.p)] {
		case 69:
			goto st64
		case 101:
			goto st64
		}
		goto tr19
	st64:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch lex.data[(lex.p)] {
		case 84:
			goto st65
		case 116:
			goto st65
		}
		goto tr19
	st65:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch lex.data[(lex.p)] {
		case 9:
			goto st65
		case 32:
			goto st65
		case 41:
			goto tr80
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st65
		}
		goto tr19
	st129:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof129
		}
	st_case_129:
		switch lex.data[(lex.p)] {
		case 42:
			goto st130
		case 61:
			goto tr245
		}
		goto tr229
	st130:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof130
		}
	st_case_130:
		if lex.data[(lex.p)] == 61 {
			goto tr247
		}
		goto tr246
	st131:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof131
		}
	st_case_131:
		switch lex.data[(lex.p)] {
		case 43:
			goto tr248
		case 61:
			goto tr249
		}
		goto tr229
	st132:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		switch lex.data[(lex.p)] {
		case 45:
			goto tr250
		case 61:
			goto tr251
		case 62:
			goto tr252
		}
		goto tr229
	tr184:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st133
	st133:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof133
		}
	st_case_133:
//line scanner/scanner.go:4701
		switch lex.data[(lex.p)] {
		case 46:
			goto st66
		case 61:
			goto tr255
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr254
		}
		goto tr229
	st66:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		if lex.data[(lex.p)] == 46 {
			goto tr81
		}
		goto tr19
	tr254:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:163
		lex.act = 8
		goto st134
	st134:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof134
		}
	st_case_134:
//line scanner/scanner.go:4733
		switch lex.data[(lex.p)] {
		case 69:
			goto st67
		case 101:
			goto st67
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr254
		}
		goto tr256
	st67:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch lex.data[(lex.p)] {
		case 43:
			goto st68
		case 45:
			goto st68
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st135
		}
		goto tr8
	st68:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st135
		}
		goto tr8
	st135:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof135
		}
	st_case_135:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st135
		}
		goto tr256
	tr185:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st136
	st136:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof136
		}
	st_case_136:
//line scanner/scanner.go:4787
		switch lex.data[(lex.p)] {
		case 42:
			goto st69
		case 47:
			goto st121
		case 61:
			goto tr258
		}
		goto tr229
	tr87:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st69
	st69:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof69
		}
	st_case_69:
//line scanner/scanner.go:4806
		switch lex.data[(lex.p)] {
		case 10:
			goto st70
		case 42:
			goto st71
		}
		goto st69
	tr88:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st70
	st70:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof70
		}
	st_case_70:
//line scanner/scanner.go:4823
		switch lex.data[(lex.p)] {
		case 10:
			goto tr88
		case 42:
			goto tr89
		}
		goto tr87
	tr89:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st71
	st71:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof71
		}
	st_case_71:
//line scanner/scanner.go:4840
		switch lex.data[(lex.p)] {
		case 10:
			goto st70
		case 42:
			goto st71
		case 47:
			goto tr90
		}
		goto st69
	tr186:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:177
		lex.act = 10
		goto st137
	st137:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof137
		}
	st_case_137:
//line scanner/scanner.go:4862
		switch lex.data[(lex.p)] {
		case 46:
			goto tr254
		case 69:
			goto st67
		case 98:
			goto st72
		case 101:
			goto st67
		case 120:
			goto st73
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr187
		}
		goto tr259
	tr187:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:177
		lex.act = 10
		goto st138
	st138:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof138
		}
	st_case_138:
//line scanner/scanner.go:4891
		switch lex.data[(lex.p)] {
		case 46:
			goto tr254
		case 69:
			goto st67
		case 101:
			goto st67
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr187
		}
		goto tr259
	st72:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st139
		}
		goto tr91
	st139:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof139
		}
	st_case_139:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st139
		}
		goto tr262
	st73:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st140
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st140
			}
		default:
			goto st140
		}
		goto tr91
	st140:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof140
		}
	st_case_140:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st140
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st140
			}
		default:
			goto st140
		}
		goto tr263
	st141:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof141
		}
	st_case_141:
		if lex.data[(lex.p)] == 58 {
			goto tr264
		}
		goto tr229
	tr189:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st142
	st142:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof142
		}
	st_case_142:
//line scanner/scanner.go:4977
		switch lex.data[(lex.p)] {
		case 10:
			goto st75
		case 13:
			goto st76
		case 32:
			goto st74
		case 63:
			goto st77
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st74
		}
		goto tr229
	tr98:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st74
	st74:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof74
		}
	st_case_74:
//line scanner/scanner.go:5001
		switch lex.data[(lex.p)] {
		case 10:
			goto st75
		case 13:
			goto st76
		case 32:
			goto st74
		case 63:
			goto st77
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st74
		}
		goto tr19
	tr99:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st75
	st75:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof75
		}
	st_case_75:
//line scanner/scanner.go:5025
		switch lex.data[(lex.p)] {
		case 10:
			goto tr99
		case 13:
			goto tr100
		case 32:
			goto tr98
		case 63:
			goto tr101
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr98
		}
		goto tr19
	tr100:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st76
	st76:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof76
		}
	st_case_76:
//line scanner/scanner.go:5049
		if lex.data[(lex.p)] == 10 {
			goto st75
		}
		goto tr19
	tr101:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st77
	st77:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof77
		}
	st_case_77:
//line scanner/scanner.go:5063
		if lex.data[(lex.p)] == 62 {
			goto tr102
		}
		goto tr19
	tr102:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st143
	st143:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof143
		}
	st_case_143:
//line scanner/scanner.go:5078
		switch lex.data[(lex.p)] {
		case 10:
			goto st144
		case 13:
			goto st78
		}
		goto tr265
	st144:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof144
		}
	st_case_144:
		goto tr267
	st78:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		if lex.data[(lex.p)] == 10 {
			goto st144
		}
		goto tr103
	st145:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof145
		}
	st_case_145:
		switch lex.data[(lex.p)] {
		case 60:
			goto tr268
		case 61:
			goto st148
		case 62:
			goto tr270
		}
		goto tr229
	tr268:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:301
		lex.act = 115
		goto st146
	st146:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof146
		}
	st_case_146:
//line scanner/scanner.go:5127
		switch lex.data[(lex.p)] {
		case 60:
			goto st79
		case 61:
			goto tr272
		}
		goto tr271
	st79:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch lex.data[(lex.p)] {
		case 9:
			goto st79
		case 32:
			goto st79
		case 34:
			goto st80
		case 39:
			goto st84
		case 96:
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr8
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr108
	st80:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		if lex.data[(lex.p)] == 96 {
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr8
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr109
	tr109:
//line scanner/scanner.rl:45
		lblStart = lex.p
		goto st81
	st81:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof81
		}
	st_case_81:
//line scanner/scanner.go:5195
		switch lex.data[(lex.p)] {
		case 34:
			goto tr110
		case 96:
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr8
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr8
				}
			case lex.data[(lex.p)] >= 91:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st81
	tr110:
//line scanner/scanner.rl:46
		lblEnd = lex.p
		goto st82
	st82:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof82
		}
	st_case_82:
//line scanner/scanner.go:5229
		switch lex.data[(lex.p)] {
		case 10:
			goto st147
		case 13:
			goto st83
		}
		goto tr8
	tr116:
//line scanner/scanner.rl:46
		lblEnd = lex.p
		goto st147
	st147:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof147
		}
	st_case_147:
//line scanner/scanner.go:5246
		goto tr273
	tr117:
//line scanner/scanner.rl:46
		lblEnd = lex.p
		goto st83
	st83:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof83
		}
	st_case_83:
//line scanner/scanner.go:5257
		if lex.data[(lex.p)] == 10 {
			goto st147
		}
		goto tr8
	st84:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		if lex.data[(lex.p)] == 96 {
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr8
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr114
	tr114:
//line scanner/scanner.rl:45
		lblStart = lex.p
		goto st85
	st85:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof85
		}
	st_case_85:
//line scanner/scanner.go:5292
		switch lex.data[(lex.p)] {
		case 39:
			goto tr110
		case 96:
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr8
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr8
				}
			case lex.data[(lex.p)] >= 91:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st85
	tr108:
//line scanner/scanner.rl:45
		lblStart = lex.p
		goto st86
	st86:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof86
		}
	st_case_86:
//line scanner/scanner.go:5326
		switch lex.data[(lex.p)] {
		case 10:
			goto tr116
		case 13:
			goto tr117
		case 96:
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr8
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr8
				}
			case lex.data[(lex.p)] >= 91:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st86
	st148:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof148
		}
	st_case_148:
		if lex.data[(lex.p)] == 62 {
			goto tr275
		}
		goto tr274
	st149:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof149
		}
	st_case_149:
		switch lex.data[(lex.p)] {
		case 61:
			goto st150
		case 62:
			goto tr277
		}
		goto tr229
	st150:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof150
		}
	st_case_150:
		if lex.data[(lex.p)] == 61 {
			goto tr279
		}
		goto tr278
	st151:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof151
		}
	st_case_151:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr280
		case 62:
			goto st152
		}
		goto tr229
	st152:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof152
		}
	st_case_152:
		if lex.data[(lex.p)] == 61 {
			goto tr283
		}
		goto tr282
	st153:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof153
		}
	st_case_153:
		switch lex.data[(lex.p)] {
		case 62:
			goto tr284
		case 63:
			goto tr285
		}
		goto tr229
	tr284:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st154
	st154:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof154
		}
	st_case_154:
//line scanner/scanner.go:5426
		switch lex.data[(lex.p)] {
		case 10:
			goto st155
		case 13:
			goto st87
		}
		goto tr286
	st155:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof155
		}
	st_case_155:
		goto tr288
	st87:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		if lex.data[(lex.p)] == 10 {
			goto st155
		}
		goto tr119
	st156:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof156
		}
	st_case_156:
		switch lex.data[(lex.p)] {
		case 66:
			goto st158
		case 78:
			goto st164
		case 82:
			goto st165
		case 83:
			goto tr293
		case 96:
			goto tr289
		case 98:
			goto st158
		case 110:
			goto st164
		case 114:
			goto st165
		case 115:
			goto tr293
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	tr201:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:340
		lex.act = 131
		goto st157
	tr293:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:200
		lex.act = 14
		goto st157
	tr299:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:198
		lex.act = 12
		goto st157
	tr300:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:270
		lex.act = 84
		goto st157
	tr303:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:199
		lex.act = 13
		goto st157
	tr308:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:201
		lex.act = 15
		goto st157
	tr320:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:202
		lex.act = 16
		goto st157
	tr321:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:203
		lex.act = 17
		goto st157
	tr323:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:204
		lex.act = 18
		goto st157
	tr330:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:229
		lex.act = 43
		goto st157
	tr334:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:205
		lex.act = 19
		goto st157
	tr336:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:206
		lex.act = 20
		goto st157
	tr340:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:207
		lex.act = 21
		goto st157
	tr344:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:208
		lex.act = 22
		goto st157
	tr347:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:211
		lex.act = 25
		goto st157
	tr353:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:209
		lex.act = 23
		goto st157
	tr357:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:210
		lex.act = 24
		goto st157
	tr358:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:223
		lex.act = 37
		goto st157
	tr366:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:212
		lex.act = 26
		goto st157
	tr371:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:214
		lex.act = 28
		goto st157
	tr374:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:215
		lex.act = 29
		goto st157
	tr386:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:216
		lex.act = 30
		goto st157
	tr393:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:218
		lex.act = 32
		goto st157
	tr394:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:219
		lex.act = 33
		goto st157
	tr399:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:220
		lex.act = 34
		goto st157
	tr403:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:221
		lex.act = 35
		goto st157
	tr405:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:222
		lex.act = 36
		goto st157
	tr411:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:224
		lex.act = 38
		goto st157
	tr419:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:226
		lex.act = 40
		goto st157
	tr425:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:228
		lex.act = 42
		goto st157
	tr431:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:230
		lex.act = 44
		goto st157
	tr433:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:231
		lex.act = 45
		goto st157
	tr434:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:232
		lex.act = 46
		goto st157
	tr445:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:234
		lex.act = 48
		goto st157
	tr458:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:257
		lex.act = 71
		goto st157
	tr466:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:235
		lex.act = 49
		goto st157
	tr470:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:236
		lex.act = 50
		goto st157
	tr476:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:237
		lex.act = 51
		goto st157
	tr479:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:233
		lex.act = 47
		goto st157
	tr482:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:238
		lex.act = 52
		goto st157
	tr491:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:239
		lex.act = 53
		goto st157
	tr492:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:269
		lex.act = 83
		goto st157
	tr493:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:271
		lex.act = 85
		goto st157
	tr500:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:242
		lex.act = 56
		goto st157
	tr503:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:240
		lex.act = 54
		goto st157
	tr509:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:243
		lex.act = 57
		goto st157
	tr513:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:241
		lex.act = 55
		goto st157
	tr526:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:259
		lex.act = 73
		goto st157
	tr529:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:244
		lex.act = 58
		goto st157
	tr535:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:245
		lex.act = 59
		goto st157
	tr539:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:246
		lex.act = 60
		goto st157
	tr544:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:247
		lex.act = 61
		goto st157
	tr546:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:249
		lex.act = 63
		goto st157
	tr548:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:248
		lex.act = 62
		goto st157
	tr553:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:250
		lex.act = 64
		goto st157
	tr554:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:251
		lex.act = 65
		goto st157
	tr556:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:252
		lex.act = 66
		goto st157
	tr560:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:253
		lex.act = 67
		goto st157
	tr562:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:272
		lex.act = 86
		goto st157
	tr571:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:254
		lex.act = 68
		goto st157
	tr587:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:260
		lex.act = 74
		goto st157
	tr591:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:261
		lex.act = 75
		goto st157
	tr597:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:262
		lex.act = 76
		goto st157
	tr605:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:263
		lex.act = 77
		goto st157
	tr617:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:268
		lex.act = 82
		goto st157
	tr622:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:264
		lex.act = 78
		goto st157
	tr629:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:266
		lex.act = 80
		goto st157
	tr639:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:265
		lex.act = 79
		goto st157
	tr645:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:267
		lex.act = 81
		goto st157
	st157:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof157
		}
	st_case_157:
//line scanner/scanner.go:5980
		if lex.data[(lex.p)] == 96 {
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr8
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr8
				}
			case lex.data[(lex.p)] >= 91:
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr201
	st158:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof158
		}
	st_case_158:
		switch lex.data[(lex.p)] {
		case 83:
			goto st159
		case 96:
			goto tr289
		case 115:
			goto st159
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st159:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof159
		}
	st_case_159:
		switch lex.data[(lex.p)] {
		case 84:
			goto st160
		case 96:
			goto tr289
		case 116:
			goto st160
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st160:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof160
		}
	st_case_160:
		switch lex.data[(lex.p)] {
		case 82:
			goto st161
		case 96:
			goto tr289
		case 114:
			goto st161
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st161:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch lex.data[(lex.p)] {
		case 65:
			goto st162
		case 96:
			goto tr289
		case 97:
			goto st162
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st162:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof162
		}
	st_case_162:
		switch lex.data[(lex.p)] {
		case 67:
			goto st163
		case 96:
			goto tr289
		case 99:
			goto st163
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st163:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof163
		}
	st_case_163:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr299
		case 96:
			goto tr289
		case 116:
			goto tr299
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st164:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr300
		case 96:
			goto tr289
		case 100:
			goto tr300
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st165:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof165
		}
	st_case_165:
		switch lex.data[(lex.p)] {
		case 82:
			goto st166
		case 96:
			goto tr289
		case 114:
			goto st166
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st166:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof166
		}
	st_case_166:
		switch lex.data[(lex.p)] {
		case 65:
			goto st167
		case 96:
			goto tr289
		case 97:
			goto st167
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st167:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof167
		}
	st_case_167:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr303
		case 96:
			goto tr289
		case 121:
			goto tr303
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	tr195:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:340
		lex.act = 131
		goto st168
	st168:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof168
		}
	st_case_168:
//line scanner/scanner.go:6324
		switch lex.data[(lex.p)] {
		case 34:
			goto st6
		case 60:
			goto st88
		case 82:
			goto st169
		case 96:
			goto tr289
		case 114:
			goto st169
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st88:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		if lex.data[(lex.p)] == 60 {
			goto st89
		}
		goto tr121
	st89:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		if lex.data[(lex.p)] == 60 {
			goto st79
		}
		goto tr121
	st169:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch lex.data[(lex.p)] {
		case 69:
			goto st170
		case 96:
			goto tr289
		case 101:
			goto st170
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st170:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch lex.data[(lex.p)] {
		case 65:
			goto st171
		case 96:
			goto tr289
		case 97:
			goto st171
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st171:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch lex.data[(lex.p)] {
		case 75:
			goto tr308
		case 96:
			goto tr289
		case 107:
			goto tr308
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st172:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof172
		}
	st_case_172:
		switch lex.data[(lex.p)] {
		case 65:
			goto st173
		case 70:
			goto st182
		case 76:
			goto st189
		case 79:
			goto st194
		case 96:
			goto tr289
		case 97:
			goto st173
		case 102:
			goto st182
		case 108:
			goto st189
		case 111:
			goto st194
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st173:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof173
		}
	st_case_173:
		switch lex.data[(lex.p)] {
		case 76:
			goto st174
		case 83:
			goto st179
		case 84:
			goto st180
		case 96:
			goto tr289
		case 108:
			goto st174
		case 115:
			goto st179
		case 116:
			goto st180
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st174:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof174
		}
	st_case_174:
		switch lex.data[(lex.p)] {
		case 76:
			goto st175
		case 96:
			goto tr289
		case 108:
			goto st175
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st175:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof175
		}
	st_case_175:
		switch lex.data[(lex.p)] {
		case 65:
			goto st176
		case 96:
			goto tr289
		case 97:
			goto st176
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st176:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof176
		}
	st_case_176:
		switch lex.data[(lex.p)] {
		case 66:
			goto st177
		case 96:
			goto tr289
		case 98:
			goto st177
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st177:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof177
		}
	st_case_177:
		switch lex.data[(lex.p)] {
		case 76:
			goto st178
		case 96:
			goto tr289
		case 108:
			goto st178
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st178:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof178
		}
	st_case_178:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr320
		case 96:
			goto tr289
		case 101:
			goto tr320
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st179:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof179
		}
	st_case_179:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr321
		case 96:
			goto tr289
		case 101:
			goto tr321
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st180:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof180
		}
	st_case_180:
		switch lex.data[(lex.p)] {
		case 67:
			goto st181
		case 96:
			goto tr289
		case 99:
			goto st181
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st181:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof181
		}
	st_case_181:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr323
		case 96:
			goto tr289
		case 104:
			goto tr323
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st182:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof182
		}
	st_case_182:
		switch lex.data[(lex.p)] {
		case 85:
			goto st183
		case 96:
			goto tr289
		case 117:
			goto st183
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st183:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof183
		}
	st_case_183:
		switch lex.data[(lex.p)] {
		case 78:
			goto st184
		case 96:
			goto tr289
		case 110:
			goto st184
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st184:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof184
		}
	st_case_184:
		switch lex.data[(lex.p)] {
		case 67:
			goto st185
		case 96:
			goto tr289
		case 99:
			goto st185
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st185:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch lex.data[(lex.p)] {
		case 84:
			goto st186
		case 96:
			goto tr289
		case 116:
			goto st186
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st186:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof186
		}
	st_case_186:
		switch lex.data[(lex.p)] {
		case 73:
			goto st187
		case 96:
			goto tr289
		case 105:
			goto st187
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st187:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof187
		}
	st_case_187:
		switch lex.data[(lex.p)] {
		case 79:
			goto st188
		case 96:
			goto tr289
		case 111:
			goto st188
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st188:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof188
		}
	st_case_188:
		switch lex.data[(lex.p)] {
		case 78:
			goto tr330
		case 96:
			goto tr289
		case 110:
			goto tr330
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st189:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof189
		}
	st_case_189:
		switch lex.data[(lex.p)] {
		case 65:
			goto st190
		case 79:
			goto st192
		case 96:
			goto tr289
		case 97:
			goto st190
		case 111:
			goto st192
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st190:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof190
		}
	st_case_190:
		switch lex.data[(lex.p)] {
		case 83:
			goto st191
		case 96:
			goto tr289
		case 115:
			goto st191
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st191:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof191
		}
	st_case_191:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr334
		case 96:
			goto tr289
		case 115:
			goto tr334
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st192:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof192
		}
	st_case_192:
		switch lex.data[(lex.p)] {
		case 78:
			goto st193
		case 96:
			goto tr289
		case 110:
			goto st193
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st193:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof193
		}
	st_case_193:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr336
		case 96:
			goto tr289
		case 101:
			goto tr336
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st194:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof194
		}
	st_case_194:
		switch lex.data[(lex.p)] {
		case 78:
			goto st195
		case 96:
			goto tr289
		case 110:
			goto st195
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st195:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof195
		}
	st_case_195:
		switch lex.data[(lex.p)] {
		case 83:
			goto st196
		case 84:
			goto st197
		case 96:
			goto tr289
		case 115:
			goto st196
		case 116:
			goto st197
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st196:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof196
		}
	st_case_196:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr340
		case 96:
			goto tr289
		case 116:
			goto tr340
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st197:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof197
		}
	st_case_197:
		switch lex.data[(lex.p)] {
		case 73:
			goto st198
		case 96:
			goto tr289
		case 105:
			goto st198
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st198:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof198
		}
	st_case_198:
		switch lex.data[(lex.p)] {
		case 78:
			goto st199
		case 96:
			goto tr289
		case 110:
			goto st199
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st199:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof199
		}
	st_case_199:
		switch lex.data[(lex.p)] {
		case 85:
			goto st200
		case 96:
			goto tr289
		case 117:
			goto st200
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st200:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof200
		}
	st_case_200:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr344
		case 96:
			goto tr289
		case 101:
			goto tr344
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st201:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof201
		}
	st_case_201:
		switch lex.data[(lex.p)] {
		case 69:
			goto st202
		case 73:
			goto st211
		case 79:
			goto tr347
		case 96:
			goto tr289
		case 101:
			goto st202
		case 105:
			goto st211
		case 111:
			goto tr347
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st202:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof202
		}
	st_case_202:
		switch lex.data[(lex.p)] {
		case 67:
			goto st203
		case 70:
			goto st207
		case 96:
			goto tr289
		case 99:
			goto st203
		case 102:
			goto st207
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st203:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof203
		}
	st_case_203:
		switch lex.data[(lex.p)] {
		case 76:
			goto st204
		case 96:
			goto tr289
		case 108:
			goto st204
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st204:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof204
		}
	st_case_204:
		switch lex.data[(lex.p)] {
		case 65:
			goto st205
		case 96:
			goto tr289
		case 97:
			goto st205
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st205:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof205
		}
	st_case_205:
		switch lex.data[(lex.p)] {
		case 82:
			goto st206
		case 96:
			goto tr289
		case 114:
			goto st206
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st206:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof206
		}
	st_case_206:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr353
		case 96:
			goto tr289
		case 101:
			goto tr353
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st207:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof207
		}
	st_case_207:
		switch lex.data[(lex.p)] {
		case 65:
			goto st208
		case 96:
			goto tr289
		case 97:
			goto st208
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st208:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof208
		}
	st_case_208:
		switch lex.data[(lex.p)] {
		case 85:
			goto st209
		case 96:
			goto tr289
		case 117:
			goto st209
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st209:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof209
		}
	st_case_209:
		switch lex.data[(lex.p)] {
		case 76:
			goto st210
		case 96:
			goto tr289
		case 108:
			goto st210
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st210:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof210
		}
	st_case_210:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr357
		case 96:
			goto tr289
		case 116:
			goto tr357
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st211:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof211
		}
	st_case_211:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr358
		case 96:
			goto tr289
		case 101:
			goto tr358
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st212:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof212
		}
	st_case_212:
		switch lex.data[(lex.p)] {
		case 67:
			goto st213
		case 76:
			goto st215
		case 77:
			goto st219
		case 78:
			goto st222
		case 86:
			goto st246
		case 88:
			goto st248
		case 96:
			goto tr289
		case 99:
			goto st213
		case 108:
			goto st215
		case 109:
			goto st219
		case 110:
			goto st222
		case 118:
			goto st246
		case 120:
			goto st248
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st213:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof213
		}
	st_case_213:
		switch lex.data[(lex.p)] {
		case 72:
			goto st214
		case 96:
			goto tr289
		case 104:
			goto st214
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st214:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof214
		}
	st_case_214:
		switch lex.data[(lex.p)] {
		case 79:
			goto tr366
		case 96:
			goto tr289
		case 111:
			goto tr366
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st215:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof215
		}
	st_case_215:
		switch lex.data[(lex.p)] {
		case 83:
			goto st216
		case 96:
			goto tr289
		case 115:
			goto st216
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st216:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof216
		}
	st_case_216:
		switch lex.data[(lex.p)] {
		case 69:
			goto st217
		case 96:
			goto tr289
		case 101:
			goto st217
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st217:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof217
		}
	st_case_217:
		switch lex.data[(lex.p)] {
		case 73:
			goto st218
		case 96:
			goto tr369
		case 105:
			goto st218
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr369
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr369
				}
			case lex.data[(lex.p)] >= 91:
				goto tr369
			}
		default:
			goto tr369
		}
		goto tr201
	st218:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof218
		}
	st_case_218:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr371
		case 96:
			goto tr289
		case 102:
			goto tr371
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st219:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof219
		}
	st_case_219:
		switch lex.data[(lex.p)] {
		case 80:
			goto st220
		case 96:
			goto tr289
		case 112:
			goto st220
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st220:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof220
		}
	st_case_220:
		switch lex.data[(lex.p)] {
		case 84:
			goto st221
		case 96:
			goto tr289
		case 116:
			goto st221
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st221:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof221
		}
	st_case_221:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr374
		case 96:
			goto tr289
		case 121:
			goto tr374
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st222:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof222
		}
	st_case_222:
		switch lex.data[(lex.p)] {
		case 68:
			goto st223
		case 96:
			goto tr289
		case 100:
			goto st223
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st223:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof223
		}
	st_case_223:
		switch lex.data[(lex.p)] {
		case 68:
			goto st224
		case 70:
			goto st230
		case 73:
			goto st236
		case 83:
			goto st237
		case 87:
			goto st242
		case 96:
			goto tr289
		case 100:
			goto st224
		case 102:
			goto st230
		case 105:
			goto st236
		case 115:
			goto st237
		case 119:
			goto st242
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st224:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof224
		}
	st_case_224:
		switch lex.data[(lex.p)] {
		case 69:
			goto st225
		case 96:
			goto tr289
		case 101:
			goto st225
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st225:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof225
		}
	st_case_225:
		switch lex.data[(lex.p)] {
		case 67:
			goto st226
		case 96:
			goto tr289
		case 99:
			goto st226
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st226:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof226
		}
	st_case_226:
		switch lex.data[(lex.p)] {
		case 76:
			goto st227
		case 96:
			goto tr289
		case 108:
			goto st227
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st227:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof227
		}
	st_case_227:
		switch lex.data[(lex.p)] {
		case 65:
			goto st228
		case 96:
			goto tr289
		case 97:
			goto st228
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st228:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof228
		}
	st_case_228:
		switch lex.data[(lex.p)] {
		case 82:
			goto st229
		case 96:
			goto tr289
		case 114:
			goto st229
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st229:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof229
		}
	st_case_229:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr386
		case 96:
			goto tr289
		case 101:
			goto tr386
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st230:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof230
		}
	st_case_230:
		switch lex.data[(lex.p)] {
		case 79:
			goto st231
		case 96:
			goto tr289
		case 111:
			goto st231
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st231:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof231
		}
	st_case_231:
		switch lex.data[(lex.p)] {
		case 82:
			goto st232
		case 96:
			goto tr289
		case 114:
			goto st232
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st232:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof232
		}
	st_case_232:
		switch lex.data[(lex.p)] {
		case 69:
			goto st233
		case 96:
			goto tr389
		case 101:
			goto st233
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr389
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr389
				}
			case lex.data[(lex.p)] >= 91:
				goto tr389
			}
		default:
			goto tr389
		}
		goto tr201
	st233:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof233
		}
	st_case_233:
		switch lex.data[(lex.p)] {
		case 65:
			goto st234
		case 96:
			goto tr289
		case 97:
			goto st234
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st234:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof234
		}
	st_case_234:
		switch lex.data[(lex.p)] {
		case 67:
			goto st235
		case 96:
			goto tr289
		case 99:
			goto st235
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st235:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof235
		}
	st_case_235:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr393
		case 96:
			goto tr289
		case 104:
			goto tr393
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st236:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof236
		}
	st_case_236:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr394
		case 96:
			goto tr289
		case 102:
			goto tr394
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st237:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof237
		}
	st_case_237:
		switch lex.data[(lex.p)] {
		case 87:
			goto st238
		case 96:
			goto tr289
		case 119:
			goto st238
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st238:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof238
		}
	st_case_238:
		switch lex.data[(lex.p)] {
		case 73:
			goto st239
		case 96:
			goto tr289
		case 105:
			goto st239
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st239:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof239
		}
	st_case_239:
		switch lex.data[(lex.p)] {
		case 84:
			goto st240
		case 96:
			goto tr289
		case 116:
			goto st240
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st240:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof240
		}
	st_case_240:
		switch lex.data[(lex.p)] {
		case 67:
			goto st241
		case 96:
			goto tr289
		case 99:
			goto st241
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st241:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof241
		}
	st_case_241:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr399
		case 96:
			goto tr289
		case 104:
			goto tr399
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st242:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof242
		}
	st_case_242:
		switch lex.data[(lex.p)] {
		case 72:
			goto st243
		case 96:
			goto tr289
		case 104:
			goto st243
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st243:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof243
		}
	st_case_243:
		switch lex.data[(lex.p)] {
		case 73:
			goto st244
		case 96:
			goto tr289
		case 105:
			goto st244
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st244:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof244
		}
	st_case_244:
		switch lex.data[(lex.p)] {
		case 76:
			goto st245
		case 96:
			goto tr289
		case 108:
			goto st245
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st245:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof245
		}
	st_case_245:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr403
		case 96:
			goto tr289
		case 101:
			goto tr403
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st246:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof246
		}
	st_case_246:
		switch lex.data[(lex.p)] {
		case 65:
			goto st247
		case 96:
			goto tr289
		case 97:
			goto st247
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st247:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof247
		}
	st_case_247:
		switch lex.data[(lex.p)] {
		case 76:
			goto tr405
		case 96:
			goto tr289
		case 108:
			goto tr405
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st248:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof248
		}
	st_case_248:
		switch lex.data[(lex.p)] {
		case 73:
			goto st249
		case 84:
			goto st250
		case 96:
			goto tr289
		case 105:
			goto st249
		case 116:
			goto st250
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st249:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof249
		}
	st_case_249:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr358
		case 96:
			goto tr289
		case 116:
			goto tr358
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st250:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof250
		}
	st_case_250:
		switch lex.data[(lex.p)] {
		case 69:
			goto st251
		case 96:
			goto tr289
		case 101:
			goto st251
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st251:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof251
		}
	st_case_251:
		switch lex.data[(lex.p)] {
		case 78:
			goto st252
		case 96:
			goto tr289
		case 110:
			goto st252
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st252:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof252
		}
	st_case_252:
		switch lex.data[(lex.p)] {
		case 68:
			goto st253
		case 96:
			goto tr289
		case 100:
			goto st253
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st253:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof253
		}
	st_case_253:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr411
		case 96:
			goto tr289
		case 115:
			goto tr411
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st254:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof254
		}
	st_case_254:
		switch lex.data[(lex.p)] {
		case 73:
			goto st255
		case 79:
			goto st260
		case 85:
			goto st183
		case 96:
			goto tr289
		case 105:
			goto st255
		case 111:
			goto st260
		case 117:
			goto st183
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st255:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof255
		}
	st_case_255:
		switch lex.data[(lex.p)] {
		case 78:
			goto st256
		case 96:
			goto tr289
		case 110:
			goto st256
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st256:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof256
		}
	st_case_256:
		switch lex.data[(lex.p)] {
		case 65:
			goto st257
		case 96:
			goto tr289
		case 97:
			goto st257
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st257:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof257
		}
	st_case_257:
		switch lex.data[(lex.p)] {
		case 76:
			goto st258
		case 96:
			goto tr289
		case 108:
			goto st258
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st258:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof258
		}
	st_case_258:
		switch lex.data[(lex.p)] {
		case 76:
			goto st259
		case 96:
			goto tr417
		case 108:
			goto st259
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr417
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr417
				}
			case lex.data[(lex.p)] >= 91:
				goto tr417
			}
		default:
			goto tr417
		}
		goto tr201
	st259:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof259
		}
	st_case_259:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr419
		case 96:
			goto tr289
		case 121:
			goto tr419
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st260:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof260
		}
	st_case_260:
		switch lex.data[(lex.p)] {
		case 82:
			goto st261
		case 96:
			goto tr289
		case 114:
			goto st261
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st261:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof261
		}
	st_case_261:
		switch lex.data[(lex.p)] {
		case 69:
			goto st262
		case 96:
			goto tr421
		case 101:
			goto st262
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr421
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr421
				}
			case lex.data[(lex.p)] >= 91:
				goto tr421
			}
		default:
			goto tr421
		}
		goto tr201
	st262:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof262
		}
	st_case_262:
		switch lex.data[(lex.p)] {
		case 65:
			goto st263
		case 96:
			goto tr289
		case 97:
			goto st263
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st263:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof263
		}
	st_case_263:
		switch lex.data[(lex.p)] {
		case 67:
			goto st264
		case 96:
			goto tr289
		case 99:
			goto st264
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st264:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof264
		}
	st_case_264:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr425
		case 96:
			goto tr289
		case 104:
			goto tr425
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st265:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof265
		}
	st_case_265:
		switch lex.data[(lex.p)] {
		case 76:
			goto st266
		case 79:
			goto st270
		case 96:
			goto tr289
		case 108:
			goto st266
		case 111:
			goto st270
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st266:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof266
		}
	st_case_266:
		switch lex.data[(lex.p)] {
		case 79:
			goto st267
		case 96:
			goto tr289
		case 111:
			goto st267
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st267:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof267
		}
	st_case_267:
		switch lex.data[(lex.p)] {
		case 66:
			goto st268
		case 96:
			goto tr289
		case 98:
			goto st268
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st268:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof268
		}
	st_case_268:
		switch lex.data[(lex.p)] {
		case 65:
			goto st269
		case 96:
			goto tr289
		case 97:
			goto st269
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st269:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof269
		}
	st_case_269:
		switch lex.data[(lex.p)] {
		case 76:
			goto tr431
		case 96:
			goto tr289
		case 108:
			goto tr431
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st270:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof270
		}
	st_case_270:
		switch lex.data[(lex.p)] {
		case 84:
			goto st271
		case 96:
			goto tr289
		case 116:
			goto st271
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st271:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof271
		}
	st_case_271:
		switch lex.data[(lex.p)] {
		case 79:
			goto tr433
		case 96:
			goto tr289
		case 111:
			goto tr433
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st272:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof272
		}
	st_case_272:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr434
		case 77:
			goto st273
		case 78:
			goto st281
		case 83:
			goto st308
		case 96:
			goto tr289
		case 102:
			goto tr434
		case 109:
			goto st273
		case 110:
			goto st281
		case 115:
			goto st308
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st273:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof273
		}
	st_case_273:
		switch lex.data[(lex.p)] {
		case 80:
			goto st274
		case 96:
			goto tr289
		case 112:
			goto st274
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st274:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof274
		}
	st_case_274:
		switch lex.data[(lex.p)] {
		case 76:
			goto st275
		case 96:
			goto tr289
		case 108:
			goto st275
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st275:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof275
		}
	st_case_275:
		switch lex.data[(lex.p)] {
		case 69:
			goto st276
		case 96:
			goto tr289
		case 101:
			goto st276
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st276:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof276
		}
	st_case_276:
		switch lex.data[(lex.p)] {
		case 77:
			goto st277
		case 96:
			goto tr289
		case 109:
			goto st277
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st277:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof277
		}
	st_case_277:
		switch lex.data[(lex.p)] {
		case 69:
			goto st278
		case 96:
			goto tr289
		case 101:
			goto st278
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st278:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof278
		}
	st_case_278:
		switch lex.data[(lex.p)] {
		case 78:
			goto st279
		case 96:
			goto tr289
		case 110:
			goto st279
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st279:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof279
		}
	st_case_279:
		switch lex.data[(lex.p)] {
		case 84:
			goto st280
		case 96:
			goto tr289
		case 116:
			goto st280
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st280:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof280
		}
	st_case_280:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr445
		case 96:
			goto tr289
		case 115:
			goto tr445
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st281:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof281
		}
	st_case_281:
		switch lex.data[(lex.p)] {
		case 67:
			goto st282
		case 83:
			goto st291
		case 84:
			goto st302
		case 96:
			goto tr289
		case 99:
			goto st282
		case 115:
			goto st291
		case 116:
			goto st302
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st282:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof282
		}
	st_case_282:
		switch lex.data[(lex.p)] {
		case 76:
			goto st283
		case 96:
			goto tr289
		case 108:
			goto st283
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st283:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof283
		}
	st_case_283:
		switch lex.data[(lex.p)] {
		case 85:
			goto st284
		case 96:
			goto tr289
		case 117:
			goto st284
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st284:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof284
		}
	st_case_284:
		switch lex.data[(lex.p)] {
		case 68:
			goto st285
		case 96:
			goto tr289
		case 100:
			goto st285
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st285:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof285
		}
	st_case_285:
		switch lex.data[(lex.p)] {
		case 69:
			goto st286
		case 96:
			goto tr289
		case 101:
			goto st286
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st286:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof286
		}
	st_case_286:
		if lex.data[(lex.p)] == 95 {
			goto st287
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr453
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr453
				}
			case lex.data[(lex.p)] >= 91:
				goto tr453
			}
		default:
			goto tr453
		}
		goto tr201
	st287:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof287
		}
	st_case_287:
		switch lex.data[(lex.p)] {
		case 79:
			goto st288
		case 96:
			goto tr289
		case 111:
			goto st288
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st288:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof288
		}
	st_case_288:
		switch lex.data[(lex.p)] {
		case 78:
			goto st289
		case 96:
			goto tr289
		case 110:
			goto st289
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st289:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof289
		}
	st_case_289:
		switch lex.data[(lex.p)] {
		case 67:
			goto st290
		case 96:
			goto tr289
		case 99:
			goto st290
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st290:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof290
		}
	st_case_290:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr458
		case 96:
			goto tr289
		case 101:
			goto tr458
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st291:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof291
		}
	st_case_291:
		switch lex.data[(lex.p)] {
		case 84:
			goto st292
		case 96:
			goto tr289
		case 116:
			goto st292
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st292:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof292
		}
	st_case_292:
		switch lex.data[(lex.p)] {
		case 65:
			goto st293
		case 69:
			goto st298
		case 96:
			goto tr289
		case 97:
			goto st293
		case 101:
			goto st298
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st293:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof293
		}
	st_case_293:
		switch lex.data[(lex.p)] {
		case 78:
			goto st294
		case 96:
			goto tr289
		case 110:
			goto st294
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st294:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof294
		}
	st_case_294:
		switch lex.data[(lex.p)] {
		case 67:
			goto st295
		case 96:
			goto tr289
		case 99:
			goto st295
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st295:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof295
		}
	st_case_295:
		switch lex.data[(lex.p)] {
		case 69:
			goto st296
		case 96:
			goto tr289
		case 101:
			goto st296
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st296:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof296
		}
	st_case_296:
		switch lex.data[(lex.p)] {
		case 79:
			goto st297
		case 96:
			goto tr289
		case 111:
			goto st297
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st297:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof297
		}
	st_case_297:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr466
		case 96:
			goto tr289
		case 102:
			goto tr466
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st298:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof298
		}
	st_case_298:
		switch lex.data[(lex.p)] {
		case 65:
			goto st299
		case 96:
			goto tr289
		case 97:
			goto st299
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st299:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof299
		}
	st_case_299:
		switch lex.data[(lex.p)] {
		case 68:
			goto st300
		case 96:
			goto tr289
		case 100:
			goto st300
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st300:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof300
		}
	st_case_300:
		switch lex.data[(lex.p)] {
		case 79:
			goto st301
		case 96:
			goto tr289
		case 111:
			goto st301
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st301:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof301
		}
	st_case_301:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr470
		case 96:
			goto tr289
		case 102:
			goto tr470
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st302:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof302
		}
	st_case_302:
		switch lex.data[(lex.p)] {
		case 69:
			goto st303
		case 96:
			goto tr289
		case 101:
			goto st303
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st303:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof303
		}
	st_case_303:
		switch lex.data[(lex.p)] {
		case 82:
			goto st304
		case 96:
			goto tr289
		case 114:
			goto st304
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st304:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof304
		}
	st_case_304:
		switch lex.data[(lex.p)] {
		case 70:
			goto st305
		case 96:
			goto tr289
		case 102:
			goto st305
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st305:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof305
		}
	st_case_305:
		switch lex.data[(lex.p)] {
		case 65:
			goto st306
		case 96:
			goto tr289
		case 97:
			goto st306
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st306:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof306
		}
	st_case_306:
		switch lex.data[(lex.p)] {
		case 67:
			goto st307
		case 96:
			goto tr289
		case 99:
			goto st307
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st307:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof307
		}
	st_case_307:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr476
		case 96:
			goto tr289
		case 101:
			goto tr476
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st308:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof308
		}
	st_case_308:
		switch lex.data[(lex.p)] {
		case 83:
			goto st309
		case 96:
			goto tr289
		case 115:
			goto st309
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st309:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof309
		}
	st_case_309:
		switch lex.data[(lex.p)] {
		case 69:
			goto st310
		case 96:
			goto tr289
		case 101:
			goto st310
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st310:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof310
		}
	st_case_310:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr479
		case 96:
			goto tr289
		case 116:
			goto tr479
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st311:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof311
		}
	st_case_311:
		switch lex.data[(lex.p)] {
		case 73:
			goto st312
		case 96:
			goto tr289
		case 105:
			goto st312
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st312:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof312
		}
	st_case_312:
		switch lex.data[(lex.p)] {
		case 83:
			goto st313
		case 96:
			goto tr289
		case 115:
			goto st313
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st313:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof313
		}
	st_case_313:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr482
		case 96:
			goto tr289
		case 116:
			goto tr482
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st314:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof314
		}
	st_case_314:
		switch lex.data[(lex.p)] {
		case 65:
			goto st315
		case 69:
			goto st322
		case 96:
			goto tr289
		case 97:
			goto st315
		case 101:
			goto st322
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st315:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof315
		}
	st_case_315:
		switch lex.data[(lex.p)] {
		case 77:
			goto st316
		case 96:
			goto tr289
		case 109:
			goto st316
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st316:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof316
		}
	st_case_316:
		switch lex.data[(lex.p)] {
		case 69:
			goto st317
		case 96:
			goto tr289
		case 101:
			goto st317
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st317:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof317
		}
	st_case_317:
		switch lex.data[(lex.p)] {
		case 83:
			goto st318
		case 96:
			goto tr289
		case 115:
			goto st318
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st318:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof318
		}
	st_case_318:
		switch lex.data[(lex.p)] {
		case 80:
			goto st319
		case 96:
			goto tr289
		case 112:
			goto st319
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st319:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof319
		}
	st_case_319:
		switch lex.data[(lex.p)] {
		case 65:
			goto st320
		case 96:
			goto tr289
		case 97:
			goto st320
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st320:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof320
		}
	st_case_320:
		switch lex.data[(lex.p)] {
		case 67:
			goto st321
		case 96:
			goto tr289
		case 99:
			goto st321
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st321:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof321
		}
	st_case_321:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr491
		case 96:
			goto tr289
		case 101:
			goto tr491
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st322:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof322
		}
	st_case_322:
		switch lex.data[(lex.p)] {
		case 87:
			goto tr492
		case 96:
			goto tr289
		case 119:
			goto tr492
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st323:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof323
		}
	st_case_323:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr493
		case 96:
			goto tr289
		case 114:
			goto tr493
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st324:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof324
		}
	st_case_324:
		switch lex.data[(lex.p)] {
		case 82:
			goto st325
		case 85:
			goto st337
		case 96:
			goto tr289
		case 114:
			goto st325
		case 117:
			goto st337
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st325:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof325
		}
	st_case_325:
		switch lex.data[(lex.p)] {
		case 73:
			goto st326
		case 79:
			goto st331
		case 96:
			goto tr289
		case 105:
			goto st326
		case 111:
			goto st331
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st326:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof326
		}
	st_case_326:
		switch lex.data[(lex.p)] {
		case 78:
			goto st327
		case 86:
			goto st328
		case 96:
			goto tr289
		case 110:
			goto st327
		case 118:
			goto st328
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st327:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof327
		}
	st_case_327:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr500
		case 96:
			goto tr289
		case 116:
			goto tr500
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st328:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof328
		}
	st_case_328:
		switch lex.data[(lex.p)] {
		case 65:
			goto st329
		case 96:
			goto tr289
		case 97:
			goto st329
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st329:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof329
		}
	st_case_329:
		switch lex.data[(lex.p)] {
		case 84:
			goto st330
		case 96:
			goto tr289
		case 116:
			goto st330
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st330:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof330
		}
	st_case_330:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr503
		case 96:
			goto tr289
		case 101:
			goto tr503
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st331:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof331
		}
	st_case_331:
		switch lex.data[(lex.p)] {
		case 84:
			goto st332
		case 96:
			goto tr289
		case 116:
			goto st332
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st332:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof332
		}
	st_case_332:
		switch lex.data[(lex.p)] {
		case 69:
			goto st333
		case 96:
			goto tr289
		case 101:
			goto st333
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st333:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof333
		}
	st_case_333:
		switch lex.data[(lex.p)] {
		case 67:
			goto st334
		case 96:
			goto tr289
		case 99:
			goto st334
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st334:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof334
		}
	st_case_334:
		switch lex.data[(lex.p)] {
		case 84:
			goto st335
		case 96:
			goto tr289
		case 116:
			goto st335
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st335:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof335
		}
	st_case_335:
		switch lex.data[(lex.p)] {
		case 69:
			goto st336
		case 96:
			goto tr289
		case 101:
			goto st336
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st336:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof336
		}
	st_case_336:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr509
		case 96:
			goto tr289
		case 100:
			goto tr509
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st337:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof337
		}
	st_case_337:
		switch lex.data[(lex.p)] {
		case 66:
			goto st338
		case 96:
			goto tr289
		case 98:
			goto st338
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st338:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof338
		}
	st_case_338:
		switch lex.data[(lex.p)] {
		case 76:
			goto st339
		case 96:
			goto tr289
		case 108:
			goto st339
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st339:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof339
		}
	st_case_339:
		switch lex.data[(lex.p)] {
		case 73:
			goto st340
		case 96:
			goto tr289
		case 105:
			goto st340
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st340:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof340
		}
	st_case_340:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr513
		case 96:
			goto tr289
		case 99:
			goto tr513
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st341:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof341
		}
	st_case_341:
		switch lex.data[(lex.p)] {
		case 69:
			goto st342
		case 96:
			goto tr289
		case 101:
			goto st342
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st342:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof342
		}
	st_case_342:
		switch lex.data[(lex.p)] {
		case 81:
			goto st343
		case 84:
			goto st352
		case 96:
			goto tr289
		case 113:
			goto st343
		case 116:
			goto st352
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st343:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof343
		}
	st_case_343:
		switch lex.data[(lex.p)] {
		case 85:
			goto st344
		case 96:
			goto tr289
		case 117:
			goto st344
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st344:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof344
		}
	st_case_344:
		switch lex.data[(lex.p)] {
		case 73:
			goto st345
		case 96:
			goto tr289
		case 105:
			goto st345
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st345:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof345
		}
	st_case_345:
		switch lex.data[(lex.p)] {
		case 82:
			goto st346
		case 96:
			goto tr289
		case 114:
			goto st346
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st346:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof346
		}
	st_case_346:
		switch lex.data[(lex.p)] {
		case 69:
			goto st347
		case 96:
			goto tr289
		case 101:
			goto st347
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st347:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof347
		}
	st_case_347:
		if lex.data[(lex.p)] == 95 {
			goto st348
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr521
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr521
				}
			case lex.data[(lex.p)] >= 91:
				goto tr521
			}
		default:
			goto tr521
		}
		goto tr201
	st348:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof348
		}
	st_case_348:
		switch lex.data[(lex.p)] {
		case 79:
			goto st349
		case 96:
			goto tr289
		case 111:
			goto st349
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st349:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof349
		}
	st_case_349:
		switch lex.data[(lex.p)] {
		case 78:
			goto st350
		case 96:
			goto tr289
		case 110:
			goto st350
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st350:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof350
		}
	st_case_350:
		switch lex.data[(lex.p)] {
		case 67:
			goto st351
		case 96:
			goto tr289
		case 99:
			goto st351
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st351:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof351
		}
	st_case_351:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr526
		case 96:
			goto tr289
		case 101:
			goto tr526
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st352:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof352
		}
	st_case_352:
		switch lex.data[(lex.p)] {
		case 85:
			goto st353
		case 96:
			goto tr289
		case 117:
			goto st353
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st353:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof353
		}
	st_case_353:
		switch lex.data[(lex.p)] {
		case 82:
			goto st354
		case 96:
			goto tr289
		case 114:
			goto st354
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st354:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof354
		}
	st_case_354:
		switch lex.data[(lex.p)] {
		case 78:
			goto tr529
		case 96:
			goto tr289
		case 110:
			goto tr529
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st355:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof355
		}
	st_case_355:
		switch lex.data[(lex.p)] {
		case 84:
			goto st356
		case 87:
			goto st360
		case 96:
			goto tr289
		case 116:
			goto st356
		case 119:
			goto st360
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st356:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof356
		}
	st_case_356:
		switch lex.data[(lex.p)] {
		case 65:
			goto st357
		case 96:
			goto tr289
		case 97:
			goto st357
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st357:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof357
		}
	st_case_357:
		switch lex.data[(lex.p)] {
		case 84:
			goto st358
		case 96:
			goto tr289
		case 116:
			goto st358
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st358:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof358
		}
	st_case_358:
		switch lex.data[(lex.p)] {
		case 73:
			goto st359
		case 96:
			goto tr289
		case 105:
			goto st359
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st359:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof359
		}
	st_case_359:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr535
		case 96:
			goto tr289
		case 99:
			goto tr535
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st360:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof360
		}
	st_case_360:
		switch lex.data[(lex.p)] {
		case 73:
			goto st361
		case 96:
			goto tr289
		case 105:
			goto st361
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st361:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof361
		}
	st_case_361:
		switch lex.data[(lex.p)] {
		case 84:
			goto st362
		case 96:
			goto tr289
		case 116:
			goto st362
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st362:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof362
		}
	st_case_362:
		switch lex.data[(lex.p)] {
		case 67:
			goto st363
		case 96:
			goto tr289
		case 99:
			goto st363
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st363:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof363
		}
	st_case_363:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr539
		case 96:
			goto tr289
		case 104:
			goto tr539
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st364:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof364
		}
	st_case_364:
		switch lex.data[(lex.p)] {
		case 72:
			goto st365
		case 82:
			goto st368
		case 96:
			goto tr289
		case 104:
			goto st365
		case 114:
			goto st368
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st365:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof365
		}
	st_case_365:
		switch lex.data[(lex.p)] {
		case 82:
			goto st366
		case 96:
			goto tr289
		case 114:
			goto st366
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st366:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof366
		}
	st_case_366:
		switch lex.data[(lex.p)] {
		case 79:
			goto st367
		case 96:
			goto tr289
		case 111:
			goto st367
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st367:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof367
		}
	st_case_367:
		switch lex.data[(lex.p)] {
		case 87:
			goto tr544
		case 96:
			goto tr289
		case 119:
			goto tr544
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st368:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof368
		}
	st_case_368:
		switch lex.data[(lex.p)] {
		case 65:
			goto st369
		case 89:
			goto tr546
		case 96:
			goto tr289
		case 97:
			goto st369
		case 121:
			goto tr546
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st369:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof369
		}
	st_case_369:
		switch lex.data[(lex.p)] {
		case 73:
			goto st370
		case 96:
			goto tr289
		case 105:
			goto st370
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st370:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof370
		}
	st_case_370:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr548
		case 96:
			goto tr289
		case 116:
			goto tr548
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st371:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof371
		}
	st_case_371:
		switch lex.data[(lex.p)] {
		case 78:
			goto st372
		case 83:
			goto st375
		case 96:
			goto tr289
		case 110:
			goto st372
		case 115:
			goto st375
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st372:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof372
		}
	st_case_372:
		switch lex.data[(lex.p)] {
		case 83:
			goto st373
		case 96:
			goto tr289
		case 115:
			goto st373
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st373:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof373
		}
	st_case_373:
		switch lex.data[(lex.p)] {
		case 69:
			goto st374
		case 96:
			goto tr289
		case 101:
			goto st374
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st374:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof374
		}
	st_case_374:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr553
		case 96:
			goto tr289
		case 116:
			goto tr553
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st375:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof375
		}
	st_case_375:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr554
		case 96:
			goto tr289
		case 101:
			goto tr554
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st376:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof376
		}
	st_case_376:
		switch lex.data[(lex.p)] {
		case 65:
			goto st377
		case 96:
			goto tr289
		case 97:
			goto st377
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st377:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof377
		}
	st_case_377:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr556
		case 96:
			goto tr289
		case 114:
			goto tr556
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st378:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof378
		}
	st_case_378:
		switch lex.data[(lex.p)] {
		case 72:
			goto st379
		case 96:
			goto tr289
		case 104:
			goto st379
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st379:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof379
		}
	st_case_379:
		switch lex.data[(lex.p)] {
		case 73:
			goto st380
		case 96:
			goto tr289
		case 105:
			goto st380
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st380:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof380
		}
	st_case_380:
		switch lex.data[(lex.p)] {
		case 76:
			goto st381
		case 96:
			goto tr289
		case 108:
			goto st381
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st381:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof381
		}
	st_case_381:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr560
		case 96:
			goto tr289
		case 101:
			goto tr560
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st382:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof382
		}
	st_case_382:
		switch lex.data[(lex.p)] {
		case 79:
			goto st383
		case 96:
			goto tr289
		case 111:
			goto st383
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st383:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof383
		}
	st_case_383:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr562
		case 96:
			goto tr289
		case 114:
			goto tr562
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st384:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof384
		}
	st_case_384:
		switch lex.data[(lex.p)] {
		case 73:
			goto st385
		case 96:
			goto tr289
		case 105:
			goto st385
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st385:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof385
		}
	st_case_385:
		switch lex.data[(lex.p)] {
		case 69:
			goto st386
		case 96:
			goto tr289
		case 101:
			goto st386
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st386:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof386
		}
	st_case_386:
		switch lex.data[(lex.p)] {
		case 76:
			goto st387
		case 96:
			goto tr289
		case 108:
			goto st387
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st387:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof387
		}
	st_case_387:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr566
		case 96:
			goto tr289
		case 100:
			goto tr566
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	tr566:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st388
	st388:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof388
		}
	st_case_388:
//line scanner/scanner.go:13314
		switch lex.data[(lex.p)] {
		case 10:
			goto st91
		case 13:
			goto st92
		case 32:
			goto st90
		case 70:
			goto st389
		case 96:
			goto tr567
		case 102:
			goto st389
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st90
				}
			default:
				goto tr567
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 91:
				if 58 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 64 {
					goto tr567
				}
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr567
				}
			default:
				goto tr567
			}
		default:
			goto tr567
		}
		goto tr201
	tr128:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st90
	st90:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof90
		}
	st_case_90:
//line scanner/scanner.go:13365
		switch lex.data[(lex.p)] {
		case 10:
			goto st91
		case 13:
			goto st92
		case 32:
			goto st90
		case 70:
			goto st93
		case 102:
			goto st93
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st90
		}
		goto tr123
	tr129:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st91
	st91:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof91
		}
	st_case_91:
//line scanner/scanner.go:13391
		switch lex.data[(lex.p)] {
		case 10:
			goto tr129
		case 13:
			goto tr130
		case 32:
			goto tr128
		case 70:
			goto tr131
		case 102:
			goto tr131
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr128
		}
		goto tr123
	tr130:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st92
	st92:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof92
		}
	st_case_92:
//line scanner/scanner.go:13417
		if lex.data[(lex.p)] == 10 {
			goto st91
		}
		goto tr123
	tr131:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st93
	st93:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof93
		}
	st_case_93:
//line scanner/scanner.go:13431
		switch lex.data[(lex.p)] {
		case 82:
			goto st94
		case 114:
			goto st94
		}
		goto tr123
	st94:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch lex.data[(lex.p)] {
		case 79:
			goto st95
		case 111:
			goto st95
		}
		goto tr123
	st95:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr134
		case 109:
			goto tr134
		}
		goto tr123
	st389:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof389
		}
	st_case_389:
		switch lex.data[(lex.p)] {
		case 82:
			goto st390
		case 96:
			goto tr289
		case 114:
			goto st390
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st390:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof390
		}
	st_case_390:
		switch lex.data[(lex.p)] {
		case 79:
			goto st391
		case 96:
			goto tr289
		case 111:
			goto st391
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st391:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof391
		}
	st_case_391:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr571
		case 96:
			goto tr289
		case 109:
			goto tr571
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st392:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof392
		}
	st_case_392:
		if lex.data[(lex.p)] == 61 {
			goto tr572
		}
		goto tr229
	st393:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof393
		}
	st_case_393:
		if lex.data[(lex.p)] == 95 {
			goto st394
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st394:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof394
		}
	st_case_394:
		switch lex.data[(lex.p)] {
		case 67:
			goto st395
		case 68:
			goto st401
		case 70:
			goto st405
		case 72:
			goto st418
		case 76:
			goto st430
		case 77:
			goto st435
		case 78:
			goto st442
		case 84:
			goto st452
		case 96:
			goto tr289
		case 99:
			goto st395
		case 100:
			goto st401
		case 102:
			goto st405
		case 104:
			goto st418
		case 108:
			goto st430
		case 109:
			goto st435
		case 110:
			goto st442
		case 116:
			goto st452
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st395:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof395
		}
	st_case_395:
		switch lex.data[(lex.p)] {
		case 76:
			goto st396
		case 96:
			goto tr289
		case 108:
			goto st396
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st396:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof396
		}
	st_case_396:
		switch lex.data[(lex.p)] {
		case 65:
			goto st397
		case 96:
			goto tr289
		case 97:
			goto st397
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st397:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof397
		}
	st_case_397:
		switch lex.data[(lex.p)] {
		case 83:
			goto st398
		case 96:
			goto tr289
		case 115:
			goto st398
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st398:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof398
		}
	st_case_398:
		switch lex.data[(lex.p)] {
		case 83:
			goto st399
		case 96:
			goto tr289
		case 115:
			goto st399
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st399:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof399
		}
	st_case_399:
		if lex.data[(lex.p)] == 95 {
			goto st400
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st400:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof400
		}
	st_case_400:
		if lex.data[(lex.p)] == 95 {
			goto tr587
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st401:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof401
		}
	st_case_401:
		switch lex.data[(lex.p)] {
		case 73:
			goto st402
		case 96:
			goto tr289
		case 105:
			goto st402
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st402:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof402
		}
	st_case_402:
		switch lex.data[(lex.p)] {
		case 82:
			goto st403
		case 96:
			goto tr289
		case 114:
			goto st403
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st403:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof403
		}
	st_case_403:
		if lex.data[(lex.p)] == 95 {
			goto st404
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st404:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof404
		}
	st_case_404:
		if lex.data[(lex.p)] == 95 {
			goto tr591
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st405:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof405
		}
	st_case_405:
		switch lex.data[(lex.p)] {
		case 73:
			goto st406
		case 85:
			goto st410
		case 96:
			goto tr289
		case 105:
			goto st406
		case 117:
			goto st410
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st406:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof406
		}
	st_case_406:
		switch lex.data[(lex.p)] {
		case 76:
			goto st407
		case 96:
			goto tr289
		case 108:
			goto st407
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st407:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof407
		}
	st_case_407:
		switch lex.data[(lex.p)] {
		case 69:
			goto st408
		case 96:
			goto tr289
		case 101:
			goto st408
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st408:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof408
		}
	st_case_408:
		if lex.data[(lex.p)] == 95 {
			goto st409
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st409:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof409
		}
	st_case_409:
		if lex.data[(lex.p)] == 95 {
			goto tr597
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st410:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof410
		}
	st_case_410:
		switch lex.data[(lex.p)] {
		case 78:
			goto st411
		case 96:
			goto tr289
		case 110:
			goto st411
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st411:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof411
		}
	st_case_411:
		switch lex.data[(lex.p)] {
		case 67:
			goto st412
		case 96:
			goto tr289
		case 99:
			goto st412
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st412:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof412
		}
	st_case_412:
		switch lex.data[(lex.p)] {
		case 84:
			goto st413
		case 96:
			goto tr289
		case 116:
			goto st413
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st413:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof413
		}
	st_case_413:
		switch lex.data[(lex.p)] {
		case 73:
			goto st414
		case 96:
			goto tr289
		case 105:
			goto st414
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st414:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof414
		}
	st_case_414:
		switch lex.data[(lex.p)] {
		case 79:
			goto st415
		case 96:
			goto tr289
		case 111:
			goto st415
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st415:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof415
		}
	st_case_415:
		switch lex.data[(lex.p)] {
		case 78:
			goto st416
		case 96:
			goto tr289
		case 110:
			goto st416
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st416:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof416
		}
	st_case_416:
		if lex.data[(lex.p)] == 95 {
			goto st417
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st417:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof417
		}
	st_case_417:
		if lex.data[(lex.p)] == 95 {
			goto tr605
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st418:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof418
		}
	st_case_418:
		switch lex.data[(lex.p)] {
		case 65:
			goto st419
		case 96:
			goto tr289
		case 97:
			goto st419
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st419:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof419
		}
	st_case_419:
		switch lex.data[(lex.p)] {
		case 76:
			goto st420
		case 96:
			goto tr289
		case 108:
			goto st420
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st420:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof420
		}
	st_case_420:
		switch lex.data[(lex.p)] {
		case 84:
			goto st421
		case 96:
			goto tr289
		case 116:
			goto st421
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st421:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof421
		}
	st_case_421:
		if lex.data[(lex.p)] == 95 {
			goto st422
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st422:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof422
		}
	st_case_422:
		switch lex.data[(lex.p)] {
		case 67:
			goto st423
		case 96:
			goto tr289
		case 99:
			goto st423
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st423:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof423
		}
	st_case_423:
		switch lex.data[(lex.p)] {
		case 79:
			goto st424
		case 96:
			goto tr289
		case 111:
			goto st424
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st424:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof424
		}
	st_case_424:
		switch lex.data[(lex.p)] {
		case 77:
			goto st425
		case 96:
			goto tr289
		case 109:
			goto st425
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st425:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof425
		}
	st_case_425:
		switch lex.data[(lex.p)] {
		case 80:
			goto st426
		case 96:
			goto tr289
		case 112:
			goto st426
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st426:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof426
		}
	st_case_426:
		switch lex.data[(lex.p)] {
		case 73:
			goto st427
		case 96:
			goto tr289
		case 105:
			goto st427
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st427:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof427
		}
	st_case_427:
		switch lex.data[(lex.p)] {
		case 76:
			goto st428
		case 96:
			goto tr289
		case 108:
			goto st428
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st428:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof428
		}
	st_case_428:
		switch lex.data[(lex.p)] {
		case 69:
			goto st429
		case 96:
			goto tr289
		case 101:
			goto st429
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st429:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof429
		}
	st_case_429:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr617
		case 96:
			goto tr289
		case 114:
			goto tr617
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st430:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof430
		}
	st_case_430:
		switch lex.data[(lex.p)] {
		case 73:
			goto st431
		case 96:
			goto tr289
		case 105:
			goto st431
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st431:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof431
		}
	st_case_431:
		switch lex.data[(lex.p)] {
		case 78:
			goto st432
		case 96:
			goto tr289
		case 110:
			goto st432
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st432:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof432
		}
	st_case_432:
		switch lex.data[(lex.p)] {
		case 69:
			goto st433
		case 96:
			goto tr289
		case 101:
			goto st433
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st433:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof433
		}
	st_case_433:
		if lex.data[(lex.p)] == 95 {
			goto st434
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st434:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof434
		}
	st_case_434:
		if lex.data[(lex.p)] == 95 {
			goto tr622
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st435:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof435
		}
	st_case_435:
		switch lex.data[(lex.p)] {
		case 69:
			goto st436
		case 96:
			goto tr289
		case 101:
			goto st436
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st436:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof436
		}
	st_case_436:
		switch lex.data[(lex.p)] {
		case 84:
			goto st437
		case 96:
			goto tr289
		case 116:
			goto st437
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st437:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof437
		}
	st_case_437:
		switch lex.data[(lex.p)] {
		case 72:
			goto st438
		case 96:
			goto tr289
		case 104:
			goto st438
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st438:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof438
		}
	st_case_438:
		switch lex.data[(lex.p)] {
		case 79:
			goto st439
		case 96:
			goto tr289
		case 111:
			goto st439
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st439:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof439
		}
	st_case_439:
		switch lex.data[(lex.p)] {
		case 68:
			goto st440
		case 96:
			goto tr289
		case 100:
			goto st440
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st440:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof440
		}
	st_case_440:
		if lex.data[(lex.p)] == 95 {
			goto st441
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st441:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof441
		}
	st_case_441:
		if lex.data[(lex.p)] == 95 {
			goto tr629
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st442:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof442
		}
	st_case_442:
		switch lex.data[(lex.p)] {
		case 65:
			goto st443
		case 96:
			goto tr289
		case 97:
			goto st443
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st443:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof443
		}
	st_case_443:
		switch lex.data[(lex.p)] {
		case 77:
			goto st444
		case 96:
			goto tr289
		case 109:
			goto st444
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st444:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof444
		}
	st_case_444:
		switch lex.data[(lex.p)] {
		case 69:
			goto st445
		case 96:
			goto tr289
		case 101:
			goto st445
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st445:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof445
		}
	st_case_445:
		switch lex.data[(lex.p)] {
		case 83:
			goto st446
		case 96:
			goto tr289
		case 115:
			goto st446
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st446:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof446
		}
	st_case_446:
		switch lex.data[(lex.p)] {
		case 80:
			goto st447
		case 96:
			goto tr289
		case 112:
			goto st447
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st447:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof447
		}
	st_case_447:
		switch lex.data[(lex.p)] {
		case 65:
			goto st448
		case 96:
			goto tr289
		case 97:
			goto st448
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st448:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof448
		}
	st_case_448:
		switch lex.data[(lex.p)] {
		case 67:
			goto st449
		case 96:
			goto tr289
		case 99:
			goto st449
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st449:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof449
		}
	st_case_449:
		switch lex.data[(lex.p)] {
		case 69:
			goto st450
		case 96:
			goto tr289
		case 101:
			goto st450
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st450:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof450
		}
	st_case_450:
		if lex.data[(lex.p)] == 95 {
			goto st451
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st451:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof451
		}
	st_case_451:
		if lex.data[(lex.p)] == 95 {
			goto tr639
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st452:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof452
		}
	st_case_452:
		switch lex.data[(lex.p)] {
		case 82:
			goto st453
		case 96:
			goto tr289
		case 114:
			goto st453
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st453:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof453
		}
	st_case_453:
		switch lex.data[(lex.p)] {
		case 65:
			goto st454
		case 96:
			goto tr289
		case 97:
			goto st454
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st454:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof454
		}
	st_case_454:
		switch lex.data[(lex.p)] {
		case 73:
			goto st455
		case 96:
			goto tr289
		case 105:
			goto st455
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st455:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof455
		}
	st_case_455:
		switch lex.data[(lex.p)] {
		case 84:
			goto st456
		case 96:
			goto tr289
		case 116:
			goto st456
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st456:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof456
		}
	st_case_456:
		if lex.data[(lex.p)] == 95 {
			goto st457
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st457:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof457
		}
	st_case_457:
		if lex.data[(lex.p)] == 95 {
			goto tr645
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr289
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr289
				}
			case lex.data[(lex.p)] >= 91:
				goto tr289
			}
		default:
			goto tr289
		}
		goto tr201
	st458:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof458
		}
	st_case_458:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr646
		case 124:
			goto tr647
		}
		goto tr229
	tr135:
//line scanner/scanner.rl:374
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st459
	tr648:
//line scanner/scanner.rl:377
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				goto st114
			}
		}
		goto st459
	tr653:
//line scanner/scanner.rl:374
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st459
	tr655:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:374
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st459
	tr659:
//line scanner/scanner.rl:377
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				goto st114
			}
		}
		goto st459
	tr660:
//line scanner/scanner.rl:375
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 459
				goto _out
			}
		}
		goto st459
	tr661:
		lex.cs = 459
//line scanner/scanner.rl:376
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st459:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof459
		}
	st_case_459:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:15587
		switch lex.data[(lex.p)] {
		case 10:
			goto tr136
		case 13:
			goto st462
		case 32:
			goto tr649
		case 45:
			goto st463
		case 96:
			goto tr648
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr649
				}
			default:
				goto tr648
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr648
				}
			case lex.data[(lex.p)] >= 91:
				goto tr648
			}
		default:
			goto tr648
		}
		goto st464
	tr649:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st460
	tr656:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st460
	st460:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof460
		}
	st_case_460:
//line scanner/scanner.go:15640
		switch lex.data[(lex.p)] {
		case 10:
			goto tr136
		case 13:
			goto st96
		case 32:
			goto tr649
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr649
		}
		goto tr653
	tr136:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st461
	tr657:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st461
	st461:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof461
		}
	st_case_461:
//line scanner/scanner.go:15670
		switch lex.data[(lex.p)] {
		case 10:
			goto tr657
		case 13:
			goto tr658
		case 32:
			goto tr656
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr656
		}
		goto tr655
	tr658:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st96
	st96:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof96
		}
	st_case_96:
//line scanner/scanner.go:15692
		if lex.data[(lex.p)] == 10 {
			goto tr136
		}
		goto tr135
	st462:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof462
		}
	st_case_462:
		if lex.data[(lex.p)] == 10 {
			goto tr136
		}
		goto tr659
	st463:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof463
		}
	st_case_463:
		if lex.data[(lex.p)] == 62 {
			goto tr660
		}
		goto tr659
	st464:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof464
		}
	st_case_464:
		if lex.data[(lex.p)] == 96 {
			goto tr661
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr661
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr661
				}
			case lex.data[(lex.p)] >= 91:
				goto tr661
			}
		default:
			goto tr661
		}
		goto st464
	tr664:
		lex.cs = 465
//line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 142:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(token)
				tok = T_ENCAPSED_AND_WHITESPACE
				lex.cs = 486
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr665:
		lex.cs = 465
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:381
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.cs = 486
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st465:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof465
		}
	st_case_465:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:15786
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		if _widec == 1034 {
			goto st467
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr662
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	tr662:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:381
		lex.act = 142
		goto st466
	tr666:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:381
		lex.act = 142
		goto st466
	st466:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof466
		}
	st_case_466:
//line scanner/scanner.go:15857
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		if _widec == 1034 {
			goto st467
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr662
		}
		goto tr664
	tr667:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st467
	st467:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof467
		}
	st_case_467:
//line scanner/scanner.go:15912
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		if _widec == 1034 {
			goto tr667
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr666
		}
		goto tr665
	tr137:
//line scanner/scanner.rl:390
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_CURLY_OPEN
			lex.call(468, 114)
			goto _out
		}
		goto st468
	tr674:
//line scanner/scanner.rl:392
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 468
					lex.top++
					goto st488
				}
			}
		}
		goto st468
	tr675:
//line scanner/scanner.rl:391
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(468, 503)
			goto _out
		}
		goto st468
	tr676:
		lex.cs = 468
//line NONE:1
		switch lex.act {
		case 143:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(token)
				tok = T_CURLY_OPEN
				lex.call(468, 114)
				goto _out
			}
		case 144:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(468, 503)
				goto _out
			}
		case 146:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(token)
				tok = T_ENCAPSED_AND_WHITESPACE

				if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
					lex.cs = 486
				}
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr677:
		lex.cs = 468
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:393
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE

			if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 486
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr681:
		lex.cs = 468
//line scanner/scanner.rl:393
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE

			if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 486
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st468:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof468
		}
	st_case_468:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16041
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto st469
		case 1403:
			goto st97
		case 1546:
			goto st471
		case 1572:
			goto st472
		case 1659:
			goto st473
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr670
		}
		goto st0
	st469:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof469
		}
	st_case_469:
		if lex.data[(lex.p)] == 123 {
			goto tr675
		}
		goto tr674
	st97:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		if lex.data[(lex.p)] == 36 {
			goto tr137
		}
		goto st0
	tr670:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:393
		lex.act = 146
		goto st470
	tr678:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:393
		lex.act = 146
		goto st470
	tr680:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:391
		lex.act = 144
		goto st470
	tr682:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:390
		lex.act = 143
		goto st470
	st470:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof470
		}
	st_case_470:
//line scanner/scanner.go:16149
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		if _widec == 1546 {
			goto st471
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr670
		}
		goto tr676
	tr679:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st471
	st471:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof471
		}
	st_case_471:
//line scanner/scanner.go:16204
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		if _widec == 1546 {
			goto tr679
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr678
		}
		goto tr677
	st472:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof472
		}
	st_case_472:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1403:
			goto tr675
		case 1546:
			goto st471
		case 1659:
			goto tr680
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr670
		}
		goto tr674
	st473:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof473
		}
	st_case_473:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto tr137
		case 1546:
			goto st471
		case 1572:
			goto tr682
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr670
		}
		goto tr681
	tr139:
//line scanner/scanner.rl:405
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_CURLY_OPEN
			lex.call(474, 114)
			goto _out
		}
		goto st474
	tr684:
		lex.cs = 474
//line scanner/scanner.rl:408
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('`'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr691:
//line scanner/scanner.rl:407
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 474
					lex.top++
					goto st488
				}
			}
		}
		goto st474
	tr692:
//line scanner/scanner.rl:406
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(474, 503)
			goto _out
		}
		goto st474
	tr693:
		lex.cs = 474
//line NONE:1
		switch lex.act {
		case 147:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(token)
				tok = T_CURLY_OPEN
				lex.call(474, 114)
				goto _out
			}
		case 148:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(474, 503)
				goto _out
			}
		case 150:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = TokenID(int('`'))
				lex.cs = 114
				{
					(lex.p)++
					goto _out
				}
			}
		case 151:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(token)
				tok = T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr694:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:409
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 474
				goto _out
			}
		}
		goto st474
	tr698:
//line scanner/scanner.rl:409
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 474
				goto _out
			}
		}
		goto st474
	st474:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof474
		}
	st_case_474:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16438
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto st475
		case 1888:
			goto tr684
		case 1915:
			goto st98
		case 2058:
			goto st477
		case 2084:
			goto st478
		case 2144:
			goto tr689
		case 2171:
			goto st479
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr686
		}
		goto st0
	st475:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof475
		}
	st_case_475:
		if lex.data[(lex.p)] == 123 {
			goto tr692
		}
		goto tr691
	st98:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		if lex.data[(lex.p)] == 36 {
			goto tr139
		}
		goto st0
	tr686:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:409
		lex.act = 151
		goto st476
	tr689:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:408
		lex.act = 150
		goto st476
	tr695:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:409
		lex.act = 151
		goto st476
	tr697:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:406
		lex.act = 148
		goto st476
	tr699:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:405
		lex.act = 147
		goto st476
	st476:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof476
		}
	st_case_476:
//line scanner/scanner.go:16557
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		if _widec == 2058 {
			goto st477
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr686
		}
		goto tr693
	tr696:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st477
	st477:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof477
		}
	st_case_477:
//line scanner/scanner.go:16612
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		if _widec == 2058 {
			goto tr696
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr695
		}
		goto tr694
	st478:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof478
		}
	st_case_478:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1915:
			goto tr692
		case 2058:
			goto st477
		case 2171:
			goto tr697
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr686
		}
		goto tr691
	st479:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof479
		}
	st_case_479:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto tr139
		case 2058:
			goto st477
		case 2084:
			goto tr699
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr686
		}
		goto tr698
	tr140:
//line scanner/scanner.rl:417
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_CURLY_OPEN
			lex.call(480, 114)
			goto _out
		}
		goto st480
	tr700:
		lex.cs = 480
//line scanner/scanner.rl:420
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('"'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr708:
//line scanner/scanner.rl:419
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 480
					lex.top++
					goto st488
				}
			}
		}
		goto st480
	tr709:
//line scanner/scanner.rl:418
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(480, 503)
			goto _out
		}
		goto st480
	tr710:
		lex.cs = 480
//line NONE:1
		switch lex.act {
		case 152:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(token)
				tok = T_CURLY_OPEN
				lex.call(480, 114)
				goto _out
			}
		case 153:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(480, 503)
				goto _out
			}
		case 155:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = TokenID(int('"'))
				lex.cs = 114
				{
					(lex.p)++
					goto _out
				}
			}
		case 156:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(token)
				tok = T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr711:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:421
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 480
				goto _out
			}
		}
		goto st480
	tr715:
//line scanner/scanner.rl:421
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 480
				goto _out
			}
		}
		goto st480
	st480:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof480
		}
	st_case_480:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16846
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2338:
			goto tr700
		case 2340:
			goto st481
		case 2427:
			goto st99
		case 2570:
			goto st483
		case 2594:
			goto tr705
		case 2596:
			goto st484
		case 2683:
			goto st485
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr703
		}
		goto st0
	st481:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof481
		}
	st_case_481:
		if lex.data[(lex.p)] == 123 {
			goto tr709
		}
		goto tr708
	st99:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		if lex.data[(lex.p)] == 36 {
			goto tr140
		}
		goto st0
	tr703:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:421
		lex.act = 156
		goto st482
	tr705:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:420
		lex.act = 155
		goto st482
	tr712:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:421
		lex.act = 156
		goto st482
	tr714:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:418
		lex.act = 153
		goto st482
	tr716:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:417
		lex.act = 152
		goto st482
	st482:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof482
		}
	st_case_482:
//line scanner/scanner.go:16965
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		if _widec == 2570 {
			goto st483
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr703
		}
		goto tr710
	tr713:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st483
	st483:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof483
		}
	st_case_483:
//line scanner/scanner.go:17020
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		if _widec == 2570 {
			goto tr713
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr712
		}
		goto tr711
	st484:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof484
		}
	st_case_484:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2427:
			goto tr709
		case 2570:
			goto st483
		case 2683:
			goto tr714
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr703
		}
		goto tr708
	st485:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof485
		}
	st_case_485:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2340:
			goto tr140
		case 2570:
			goto st483
		case 2596:
			goto tr716
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr703
		}
		goto tr715
	tr718:
		lex.cs = 486
//line scanner/scanner.rl:429
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_END_HEREDOC
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st486:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof486
		}
	st_case_486:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:17199
		if lex.data[(lex.p)] == 96 {
			goto st0
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto st487
	st487:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof487
		}
	st_case_487:
		if lex.data[(lex.p)] == 96 {
			goto tr718
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr718
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr718
				}
			case lex.data[(lex.p)] >= 91:
				goto tr718
			}
		default:
			goto tr718
		}
		goto st487
	tr141:
//line scanner/scanner.rl:448
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st488
	tr142:
//line scanner/scanner.rl:445
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 488
				goto _out
			}
		}
		goto st488
	tr719:
//line scanner/scanner.rl:448
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st488
	tr723:
//line scanner/scanner.rl:447
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('['))
			lex.call(488, 493)
			goto _out
		}
		goto st488
	tr724:
//line scanner/scanner.rl:448
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st488
	tr726:
//line scanner/scanner.rl:444
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 488
				goto _out
			}
		}
		goto st488
	tr728:
//line scanner/scanner.rl:446
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 488
				goto _out
			}
		}
		goto st488
	st488:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof488
		}
	st_case_488:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:17291
		switch lex.data[(lex.p)] {
		case 36:
			goto st489
		case 45:
			goto tr721
		case 91:
			goto tr723
		case 96:
			goto tr719
		}
		switch {
		case lex.data[(lex.p)] < 92:
			if lex.data[(lex.p)] <= 64 {
				goto tr719
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr719
			}
		default:
			goto tr719
		}
		goto st492
	st489:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof489
		}
	st_case_489:
		if lex.data[(lex.p)] == 96 {
			goto tr724
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr724
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr724
			}
		default:
			goto tr724
		}
		goto st490
	st490:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof490
		}
	st_case_490:
		if lex.data[(lex.p)] == 96 {
			goto tr726
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr726
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr726
				}
			case lex.data[(lex.p)] >= 91:
				goto tr726
			}
		default:
			goto tr726
		}
		goto st490
	tr721:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st491
	st491:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof491
		}
	st_case_491:
//line scanner/scanner.go:17372
		if lex.data[(lex.p)] == 62 {
			goto st100
		}
		goto tr724
	st100:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		if lex.data[(lex.p)] == 96 {
			goto tr141
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr141
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr141
			}
		default:
			goto tr141
		}
		goto tr142
	st492:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof492
		}
	st_case_492:
		if lex.data[(lex.p)] == 96 {
			goto tr728
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr728
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr728
				}
			case lex.data[(lex.p)] >= 91:
				goto tr728
			}
		default:
			goto tr728
		}
		goto st492
	tr143:
//line scanner/scanner.rl:452
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 493
				goto _out
			}
		}
		goto st493
	tr729:
//line scanner/scanner.rl:458
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st493
	tr730:
//line scanner/scanner.rl:455
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st493
	tr733:
//line scanner/scanner.rl:456
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 493
				goto _out
			}
		}
		goto st493
	tr738:
//line scanner/scanner.rl:457
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(']'))
			lex.ret(2)
			goto _out
		}
		goto st493
	tr739:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:455
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st493
	tr740:
//line scanner/scanner.rl:458
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st493
	tr741:
//line scanner/scanner.rl:456
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 493
				goto _out
			}
		}
		goto st493
	tr743:
//line scanner/scanner.rl:453
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 493
				goto _out
			}
		}
		goto st493
	tr744:
//line scanner/scanner.rl:452
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 493
				goto _out
			}
		}
		goto st493
	tr747:
//line scanner/scanner.rl:454
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 493
				goto _out
			}
		}
		goto st493
	st493:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof493
		}
	st_case_493:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:17504
		switch lex.data[(lex.p)] {
		case 10:
			goto st494
		case 13:
			goto st495
		case 32:
			goto tr730
		case 33:
			goto tr733
		case 35:
			goto tr730
		case 36:
			goto st496
		case 39:
			goto tr730
		case 48:
			goto tr735
		case 92:
			goto tr730
		case 93:
			goto tr738
		case 96:
			goto tr729
		case 124:
			goto tr733
		case 126:
			goto tr733
		}
		switch {
		case lex.data[(lex.p)] < 37:
			switch {
			case lex.data[(lex.p)] < 9:
				if lex.data[(lex.p)] <= 8 {
					goto tr729
				}
			case lex.data[(lex.p)] > 12:
				if 14 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 34 {
					goto tr729
				}
			default:
				goto tr730
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 58:
				if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
					goto st499
				}
			case lex.data[(lex.p)] > 64:
				switch {
				case lex.data[(lex.p)] > 94:
					if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
						goto tr729
					}
				case lex.data[(lex.p)] >= 91:
					goto tr733
				}
			default:
				goto tr733
			}
		default:
			goto tr733
		}
		goto st502
	st494:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof494
		}
	st_case_494:
		goto tr739
	st495:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof495
		}
	st_case_495:
		if lex.data[(lex.p)] == 10 {
			goto st494
		}
		goto tr740
	st496:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof496
		}
	st_case_496:
		if lex.data[(lex.p)] == 96 {
			goto tr741
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr741
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr741
			}
		default:
			goto tr741
		}
		goto st497
	st497:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof497
		}
	st_case_497:
		if lex.data[(lex.p)] == 96 {
			goto tr743
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr743
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr743
				}
			case lex.data[(lex.p)] >= 91:
				goto tr743
			}
		default:
			goto tr743
		}
		goto st497
	tr735:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st498
	st498:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof498
		}
	st_case_498:
//line scanner/scanner.go:17641
		switch lex.data[(lex.p)] {
		case 98:
			goto st101
		case 120:
			goto st102
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st499
		}
		goto tr744
	st499:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof499
		}
	st_case_499:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st499
		}
		goto tr744
	st101:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st500
		}
		goto tr143
	st500:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof500
		}
	st_case_500:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st500
		}
		goto tr744
	st102:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st501
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st501
			}
		default:
			goto st501
		}
		goto tr143
	st501:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof501
		}
	st_case_501:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st501
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st501
			}
		default:
			goto st501
		}
		goto tr744
	st502:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof502
		}
	st_case_502:
		if lex.data[(lex.p)] == 96 {
			goto tr747
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr747
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr747
				}
			case lex.data[(lex.p)] >= 91:
				goto tr747
			}
		default:
			goto tr747
		}
		goto st502
	tr146:
		lex.cs = 503
//line scanner/scanner.rl:466
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	tr148:
		lex.cs = 503
//line scanner/scanner.rl:465
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_STRING_VARNAME
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr748:
		lex.cs = 503
//line scanner/scanner.rl:466
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	tr750:
		lex.cs = 503
//line scanner/scanner.rl:466
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	st503:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof503
		}
	st_case_503:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:17777
		if lex.data[(lex.p)] == 96 {
			goto tr748
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr748
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr748
			}
		default:
			goto tr748
		}
		goto tr749
	tr749:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st504
	st504:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof504
		}
	st_case_504:
//line scanner/scanner.go:17804
		switch lex.data[(lex.p)] {
		case 91:
			goto tr148
		case 96:
			goto tr750
		case 125:
			goto tr148
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr750
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr750
				}
			case lex.data[(lex.p)] >= 92:
				goto tr750
			}
		default:
			goto tr750
		}
		goto st103
	st103:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch lex.data[(lex.p)] {
		case 91:
			goto tr148
		case 96:
			goto tr146
		case 125:
			goto tr148
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr146
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr146
				}
			case lex.data[(lex.p)] >= 92:
				goto tr146
			}
		default:
			goto tr146
		}
		goto st103
	tr149:
//line scanner/scanner.rl:470
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st505
	tr751:
		lex.cs = 505
//line scanner/scanner.rl:472
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	tr754:
		lex.cs = 505
//line scanner/scanner.rl:471
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('('))
			lex.cs = 509
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr755:
//line scanner/scanner.rl:470
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st505
	tr757:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:470
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st505
	tr761:
		lex.cs = 505
//line scanner/scanner.rl:472
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	st505:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof505
		}
	st_case_505:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:17911
		switch lex.data[(lex.p)] {
		case 10:
			goto tr150
		case 13:
			goto st508
		case 32:
			goto tr752
		case 40:
			goto tr754
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr752
		}
		goto tr751
	tr752:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st506
	tr758:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st506
	st506:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof506
		}
	st_case_506:
//line scanner/scanner.go:17943
		switch lex.data[(lex.p)] {
		case 10:
			goto tr150
		case 13:
			goto st104
		case 32:
			goto tr752
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr752
		}
		goto tr755
	tr150:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st507
	tr759:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st507
	st507:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof507
		}
	st_case_507:
//line scanner/scanner.go:17973
		switch lex.data[(lex.p)] {
		case 10:
			goto tr759
		case 13:
			goto tr760
		case 32:
			goto tr758
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr758
		}
		goto tr757
	tr760:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st104
	st104:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof104
		}
	st_case_104:
//line scanner/scanner.go:17995
		if lex.data[(lex.p)] == 10 {
			goto tr150
		}
		goto tr149
	st508:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof508
		}
	st_case_508:
		if lex.data[(lex.p)] == 10 {
			goto tr150
		}
		goto tr761
	tr151:
//line scanner/scanner.rl:476
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st509
	tr762:
		lex.cs = 509
//line scanner/scanner.rl:478
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	tr765:
		lex.cs = 509
//line scanner/scanner.rl:477
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(')'))
			lex.cs = 513
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr766:
//line scanner/scanner.rl:476
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st509
	tr768:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:476
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st509
	tr772:
		lex.cs = 509
//line scanner/scanner.rl:478
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	st509:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof509
		}
	st_case_509:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:18058
		switch lex.data[(lex.p)] {
		case 10:
			goto tr152
		case 13:
			goto st512
		case 32:
			goto tr763
		case 41:
			goto tr765
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr763
		}
		goto tr762
	tr763:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st510
	tr769:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st510
	st510:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof510
		}
	st_case_510:
//line scanner/scanner.go:18090
		switch lex.data[(lex.p)] {
		case 10:
			goto tr152
		case 13:
			goto st105
		case 32:
			goto tr763
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr763
		}
		goto tr766
	tr152:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st511
	tr770:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st511
	st511:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof511
		}
	st_case_511:
//line scanner/scanner.go:18120
		switch lex.data[(lex.p)] {
		case 10:
			goto tr770
		case 13:
			goto tr771
		case 32:
			goto tr769
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr769
		}
		goto tr768
	tr771:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st105
	st105:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof105
		}
	st_case_105:
//line scanner/scanner.go:18142
		if lex.data[(lex.p)] == 10 {
			goto tr152
		}
		goto tr151
	st512:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof512
		}
	st_case_512:
		if lex.data[(lex.p)] == 10 {
			goto tr152
		}
		goto tr772
	tr153:
//line scanner/scanner.rl:482
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st513
	tr773:
		lex.cs = 513
//line scanner/scanner.rl:484
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	tr776:
		lex.cs = 513
//line scanner/scanner.rl:483
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 517
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr777:
//line scanner/scanner.rl:482
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st513
	tr779:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:482
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st513
	tr783:
		lex.cs = 513
//line scanner/scanner.rl:484
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 114
		}
		goto _again
	st513:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof513
		}
	st_case_513:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:18205
		switch lex.data[(lex.p)] {
		case 10:
			goto tr154
		case 13:
			goto st516
		case 32:
			goto tr774
		case 59:
			goto tr776
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr774
		}
		goto tr773
	tr774:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st514
	tr780:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st514
	st514:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof514
		}
	st_case_514:
//line scanner/scanner.go:18237
		switch lex.data[(lex.p)] {
		case 10:
			goto tr154
		case 13:
			goto st106
		case 32:
			goto tr774
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr774
		}
		goto tr777
	tr154:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st515
	tr781:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st515
	st515:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof515
		}
	st_case_515:
//line scanner/scanner.go:18267
		switch lex.data[(lex.p)] {
		case 10:
			goto tr781
		case 13:
			goto tr782
		case 32:
			goto tr780
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr780
		}
		goto tr779
	tr782:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st106
	st106:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof106
		}
	st_case_106:
//line scanner/scanner.go:18289
		if lex.data[(lex.p)] == 10 {
			goto tr154
		}
		goto tr153
	st516:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof516
		}
	st_case_516:
		if lex.data[(lex.p)] == 10 {
			goto tr154
		}
		goto tr783
	tr786:
//line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 182:
			{
				(lex.p) = (lex.te) - 1
				lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			}
		}

		goto st517
	tr787:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:488
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
		}
		goto st517
	st517:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof517
		}
	st_case_517:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:18336
		if lex.data[(lex.p)] == 10 {
			goto st519
		}
		goto tr784
	tr784:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:488
		lex.act = 182
		goto st518
	tr788:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:488
		lex.act = 182
		goto st518
	st518:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof518
		}
	st_case_518:
//line scanner/scanner.go:18362
		if lex.data[(lex.p)] == 10 {
			goto st519
		}
		goto tr784
	tr789:
//line scanner/scanner.rl:64
		lex.NewLines.Append(lex.p)
		goto st519
	st519:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof519
		}
	st_case_519:
//line scanner/scanner.go:18376
		if lex.data[(lex.p)] == 10 {
			goto tr789
		}
		goto tr788
	st_out:
	_test_eof107:
		lex.cs = 107
		goto _test_eof
	_test_eof108:
		lex.cs = 108
		goto _test_eof
	_test_eof109:
		lex.cs = 109
		goto _test_eof
	_test_eof110:
		lex.cs = 110
		goto _test_eof
	_test_eof111:
		lex.cs = 111
		goto _test_eof
	_test_eof112:
		lex.cs = 112
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof113:
		lex.cs = 113
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof114:
		lex.cs = 114
		goto _test_eof
	_test_eof115:
		lex.cs = 115
		goto _test_eof
	_test_eof116:
		lex.cs = 116
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof117:
		lex.cs = 117
		goto _test_eof
	_test_eof118:
		lex.cs = 118
		goto _test_eof
	_test_eof119:
		lex.cs = 119
		goto _test_eof
	_test_eof120:
		lex.cs = 120
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof121:
		lex.cs = 121
		goto _test_eof
	_test_eof122:
		lex.cs = 122
		goto _test_eof
	_test_eof123:
		lex.cs = 123
		goto _test_eof
	_test_eof124:
		lex.cs = 124
		goto _test_eof
	_test_eof125:
		lex.cs = 125
		goto _test_eof
	_test_eof126:
		lex.cs = 126
		goto _test_eof
	_test_eof127:
		lex.cs = 127
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof
	_test_eof128:
		lex.cs = 128
		goto _test_eof
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof18:
		lex.cs = 18
		goto _test_eof
	_test_eof19:
		lex.cs = 19
		goto _test_eof
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof22:
		lex.cs = 22
		goto _test_eof
	_test_eof23:
		lex.cs = 23
		goto _test_eof
	_test_eof24:
		lex.cs = 24
		goto _test_eof
	_test_eof25:
		lex.cs = 25
		goto _test_eof
	_test_eof26:
		lex.cs = 26
		goto _test_eof
	_test_eof27:
		lex.cs = 27
		goto _test_eof
	_test_eof28:
		lex.cs = 28
		goto _test_eof
	_test_eof29:
		lex.cs = 29
		goto _test_eof
	_test_eof30:
		lex.cs = 30
		goto _test_eof
	_test_eof31:
		lex.cs = 31
		goto _test_eof
	_test_eof32:
		lex.cs = 32
		goto _test_eof
	_test_eof33:
		lex.cs = 33
		goto _test_eof
	_test_eof34:
		lex.cs = 34
		goto _test_eof
	_test_eof35:
		lex.cs = 35
		goto _test_eof
	_test_eof36:
		lex.cs = 36
		goto _test_eof
	_test_eof37:
		lex.cs = 37
		goto _test_eof
	_test_eof38:
		lex.cs = 38
		goto _test_eof
	_test_eof39:
		lex.cs = 39
		goto _test_eof
	_test_eof40:
		lex.cs = 40
		goto _test_eof
	_test_eof41:
		lex.cs = 41
		goto _test_eof
	_test_eof42:
		lex.cs = 42
		goto _test_eof
	_test_eof43:
		lex.cs = 43
		goto _test_eof
	_test_eof44:
		lex.cs = 44
		goto _test_eof
	_test_eof45:
		lex.cs = 45
		goto _test_eof
	_test_eof46:
		lex.cs = 46
		goto _test_eof
	_test_eof47:
		lex.cs = 47
		goto _test_eof
	_test_eof48:
		lex.cs = 48
		goto _test_eof
	_test_eof49:
		lex.cs = 49
		goto _test_eof
	_test_eof50:
		lex.cs = 50
		goto _test_eof
	_test_eof51:
		lex.cs = 51
		goto _test_eof
	_test_eof52:
		lex.cs = 52
		goto _test_eof
	_test_eof53:
		lex.cs = 53
		goto _test_eof
	_test_eof54:
		lex.cs = 54
		goto _test_eof
	_test_eof55:
		lex.cs = 55
		goto _test_eof
	_test_eof56:
		lex.cs = 56
		goto _test_eof
	_test_eof57:
		lex.cs = 57
		goto _test_eof
	_test_eof58:
		lex.cs = 58
		goto _test_eof
	_test_eof59:
		lex.cs = 59
		goto _test_eof
	_test_eof60:
		lex.cs = 60
		goto _test_eof
	_test_eof61:
		lex.cs = 61
		goto _test_eof
	_test_eof62:
		lex.cs = 62
		goto _test_eof
	_test_eof63:
		lex.cs = 63
		goto _test_eof
	_test_eof64:
		lex.cs = 64
		goto _test_eof
	_test_eof65:
		lex.cs = 65
		goto _test_eof
	_test_eof129:
		lex.cs = 129
		goto _test_eof
	_test_eof130:
		lex.cs = 130
		goto _test_eof
	_test_eof131:
		lex.cs = 131
		goto _test_eof
	_test_eof132:
		lex.cs = 132
		goto _test_eof
	_test_eof133:
		lex.cs = 133
		goto _test_eof
	_test_eof66:
		lex.cs = 66
		goto _test_eof
	_test_eof134:
		lex.cs = 134
		goto _test_eof
	_test_eof67:
		lex.cs = 67
		goto _test_eof
	_test_eof68:
		lex.cs = 68
		goto _test_eof
	_test_eof135:
		lex.cs = 135
		goto _test_eof
	_test_eof136:
		lex.cs = 136
		goto _test_eof
	_test_eof69:
		lex.cs = 69
		goto _test_eof
	_test_eof70:
		lex.cs = 70
		goto _test_eof
	_test_eof71:
		lex.cs = 71
		goto _test_eof
	_test_eof137:
		lex.cs = 137
		goto _test_eof
	_test_eof138:
		lex.cs = 138
		goto _test_eof
	_test_eof72:
		lex.cs = 72
		goto _test_eof
	_test_eof139:
		lex.cs = 139
		goto _test_eof
	_test_eof73:
		lex.cs = 73
		goto _test_eof
	_test_eof140:
		lex.cs = 140
		goto _test_eof
	_test_eof141:
		lex.cs = 141
		goto _test_eof
	_test_eof142:
		lex.cs = 142
		goto _test_eof
	_test_eof74:
		lex.cs = 74
		goto _test_eof
	_test_eof75:
		lex.cs = 75
		goto _test_eof
	_test_eof76:
		lex.cs = 76
		goto _test_eof
	_test_eof77:
		lex.cs = 77
		goto _test_eof
	_test_eof143:
		lex.cs = 143
		goto _test_eof
	_test_eof144:
		lex.cs = 144
		goto _test_eof
	_test_eof78:
		lex.cs = 78
		goto _test_eof
	_test_eof145:
		lex.cs = 145
		goto _test_eof
	_test_eof146:
		lex.cs = 146
		goto _test_eof
	_test_eof79:
		lex.cs = 79
		goto _test_eof
	_test_eof80:
		lex.cs = 80
		goto _test_eof
	_test_eof81:
		lex.cs = 81
		goto _test_eof
	_test_eof82:
		lex.cs = 82
		goto _test_eof
	_test_eof147:
		lex.cs = 147
		goto _test_eof
	_test_eof83:
		lex.cs = 83
		goto _test_eof
	_test_eof84:
		lex.cs = 84
		goto _test_eof
	_test_eof85:
		lex.cs = 85
		goto _test_eof
	_test_eof86:
		lex.cs = 86
		goto _test_eof
	_test_eof148:
		lex.cs = 148
		goto _test_eof
	_test_eof149:
		lex.cs = 149
		goto _test_eof
	_test_eof150:
		lex.cs = 150
		goto _test_eof
	_test_eof151:
		lex.cs = 151
		goto _test_eof
	_test_eof152:
		lex.cs = 152
		goto _test_eof
	_test_eof153:
		lex.cs = 153
		goto _test_eof
	_test_eof154:
		lex.cs = 154
		goto _test_eof
	_test_eof155:
		lex.cs = 155
		goto _test_eof
	_test_eof87:
		lex.cs = 87
		goto _test_eof
	_test_eof156:
		lex.cs = 156
		goto _test_eof
	_test_eof157:
		lex.cs = 157
		goto _test_eof
	_test_eof158:
		lex.cs = 158
		goto _test_eof
	_test_eof159:
		lex.cs = 159
		goto _test_eof
	_test_eof160:
		lex.cs = 160
		goto _test_eof
	_test_eof161:
		lex.cs = 161
		goto _test_eof
	_test_eof162:
		lex.cs = 162
		goto _test_eof
	_test_eof163:
		lex.cs = 163
		goto _test_eof
	_test_eof164:
		lex.cs = 164
		goto _test_eof
	_test_eof165:
		lex.cs = 165
		goto _test_eof
	_test_eof166:
		lex.cs = 166
		goto _test_eof
	_test_eof167:
		lex.cs = 167
		goto _test_eof
	_test_eof168:
		lex.cs = 168
		goto _test_eof
	_test_eof88:
		lex.cs = 88
		goto _test_eof
	_test_eof89:
		lex.cs = 89
		goto _test_eof
	_test_eof169:
		lex.cs = 169
		goto _test_eof
	_test_eof170:
		lex.cs = 170
		goto _test_eof
	_test_eof171:
		lex.cs = 171
		goto _test_eof
	_test_eof172:
		lex.cs = 172
		goto _test_eof
	_test_eof173:
		lex.cs = 173
		goto _test_eof
	_test_eof174:
		lex.cs = 174
		goto _test_eof
	_test_eof175:
		lex.cs = 175
		goto _test_eof
	_test_eof176:
		lex.cs = 176
		goto _test_eof
	_test_eof177:
		lex.cs = 177
		goto _test_eof
	_test_eof178:
		lex.cs = 178
		goto _test_eof
	_test_eof179:
		lex.cs = 179
		goto _test_eof
	_test_eof180:
		lex.cs = 180
		goto _test_eof
	_test_eof181:
		lex.cs = 181
		goto _test_eof
	_test_eof182:
		lex.cs = 182
		goto _test_eof
	_test_eof183:
		lex.cs = 183
		goto _test_eof
	_test_eof184:
		lex.cs = 184
		goto _test_eof
	_test_eof185:
		lex.cs = 185
		goto _test_eof
	_test_eof186:
		lex.cs = 186
		goto _test_eof
	_test_eof187:
		lex.cs = 187
		goto _test_eof
	_test_eof188:
		lex.cs = 188
		goto _test_eof
	_test_eof189:
		lex.cs = 189
		goto _test_eof
	_test_eof190:
		lex.cs = 190
		goto _test_eof
	_test_eof191:
		lex.cs = 191
		goto _test_eof
	_test_eof192:
		lex.cs = 192
		goto _test_eof
	_test_eof193:
		lex.cs = 193
		goto _test_eof
	_test_eof194:
		lex.cs = 194
		goto _test_eof
	_test_eof195:
		lex.cs = 195
		goto _test_eof
	_test_eof196:
		lex.cs = 196
		goto _test_eof
	_test_eof197:
		lex.cs = 197
		goto _test_eof
	_test_eof198:
		lex.cs = 198
		goto _test_eof
	_test_eof199:
		lex.cs = 199
		goto _test_eof
	_test_eof200:
		lex.cs = 200
		goto _test_eof
	_test_eof201:
		lex.cs = 201
		goto _test_eof
	_test_eof202:
		lex.cs = 202
		goto _test_eof
	_test_eof203:
		lex.cs = 203
		goto _test_eof
	_test_eof204:
		lex.cs = 204
		goto _test_eof
	_test_eof205:
		lex.cs = 205
		goto _test_eof
	_test_eof206:
		lex.cs = 206
		goto _test_eof
	_test_eof207:
		lex.cs = 207
		goto _test_eof
	_test_eof208:
		lex.cs = 208
		goto _test_eof
	_test_eof209:
		lex.cs = 209
		goto _test_eof
	_test_eof210:
		lex.cs = 210
		goto _test_eof
	_test_eof211:
		lex.cs = 211
		goto _test_eof
	_test_eof212:
		lex.cs = 212
		goto _test_eof
	_test_eof213:
		lex.cs = 213
		goto _test_eof
	_test_eof214:
		lex.cs = 214
		goto _test_eof
	_test_eof215:
		lex.cs = 215
		goto _test_eof
	_test_eof216:
		lex.cs = 216
		goto _test_eof
	_test_eof217:
		lex.cs = 217
		goto _test_eof
	_test_eof218:
		lex.cs = 218
		goto _test_eof
	_test_eof219:
		lex.cs = 219
		goto _test_eof
	_test_eof220:
		lex.cs = 220
		goto _test_eof
	_test_eof221:
		lex.cs = 221
		goto _test_eof
	_test_eof222:
		lex.cs = 222
		goto _test_eof
	_test_eof223:
		lex.cs = 223
		goto _test_eof
	_test_eof224:
		lex.cs = 224
		goto _test_eof
	_test_eof225:
		lex.cs = 225
		goto _test_eof
	_test_eof226:
		lex.cs = 226
		goto _test_eof
	_test_eof227:
		lex.cs = 227
		goto _test_eof
	_test_eof228:
		lex.cs = 228
		goto _test_eof
	_test_eof229:
		lex.cs = 229
		goto _test_eof
	_test_eof230:
		lex.cs = 230
		goto _test_eof
	_test_eof231:
		lex.cs = 231
		goto _test_eof
	_test_eof232:
		lex.cs = 232
		goto _test_eof
	_test_eof233:
		lex.cs = 233
		goto _test_eof
	_test_eof234:
		lex.cs = 234
		goto _test_eof
	_test_eof235:
		lex.cs = 235
		goto _test_eof
	_test_eof236:
		lex.cs = 236
		goto _test_eof
	_test_eof237:
		lex.cs = 237
		goto _test_eof
	_test_eof238:
		lex.cs = 238
		goto _test_eof
	_test_eof239:
		lex.cs = 239
		goto _test_eof
	_test_eof240:
		lex.cs = 240
		goto _test_eof
	_test_eof241:
		lex.cs = 241
		goto _test_eof
	_test_eof242:
		lex.cs = 242
		goto _test_eof
	_test_eof243:
		lex.cs = 243
		goto _test_eof
	_test_eof244:
		lex.cs = 244
		goto _test_eof
	_test_eof245:
		lex.cs = 245
		goto _test_eof
	_test_eof246:
		lex.cs = 246
		goto _test_eof
	_test_eof247:
		lex.cs = 247
		goto _test_eof
	_test_eof248:
		lex.cs = 248
		goto _test_eof
	_test_eof249:
		lex.cs = 249
		goto _test_eof
	_test_eof250:
		lex.cs = 250
		goto _test_eof
	_test_eof251:
		lex.cs = 251
		goto _test_eof
	_test_eof252:
		lex.cs = 252
		goto _test_eof
	_test_eof253:
		lex.cs = 253
		goto _test_eof
	_test_eof254:
		lex.cs = 254
		goto _test_eof
	_test_eof255:
		lex.cs = 255
		goto _test_eof
	_test_eof256:
		lex.cs = 256
		goto _test_eof
	_test_eof257:
		lex.cs = 257
		goto _test_eof
	_test_eof258:
		lex.cs = 258
		goto _test_eof
	_test_eof259:
		lex.cs = 259
		goto _test_eof
	_test_eof260:
		lex.cs = 260
		goto _test_eof
	_test_eof261:
		lex.cs = 261
		goto _test_eof
	_test_eof262:
		lex.cs = 262
		goto _test_eof
	_test_eof263:
		lex.cs = 263
		goto _test_eof
	_test_eof264:
		lex.cs = 264
		goto _test_eof
	_test_eof265:
		lex.cs = 265
		goto _test_eof
	_test_eof266:
		lex.cs = 266
		goto _test_eof
	_test_eof267:
		lex.cs = 267
		goto _test_eof
	_test_eof268:
		lex.cs = 268
		goto _test_eof
	_test_eof269:
		lex.cs = 269
		goto _test_eof
	_test_eof270:
		lex.cs = 270
		goto _test_eof
	_test_eof271:
		lex.cs = 271
		goto _test_eof
	_test_eof272:
		lex.cs = 272
		goto _test_eof
	_test_eof273:
		lex.cs = 273
		goto _test_eof
	_test_eof274:
		lex.cs = 274
		goto _test_eof
	_test_eof275:
		lex.cs = 275
		goto _test_eof
	_test_eof276:
		lex.cs = 276
		goto _test_eof
	_test_eof277:
		lex.cs = 277
		goto _test_eof
	_test_eof278:
		lex.cs = 278
		goto _test_eof
	_test_eof279:
		lex.cs = 279
		goto _test_eof
	_test_eof280:
		lex.cs = 280
		goto _test_eof
	_test_eof281:
		lex.cs = 281
		goto _test_eof
	_test_eof282:
		lex.cs = 282
		goto _test_eof
	_test_eof283:
		lex.cs = 283
		goto _test_eof
	_test_eof284:
		lex.cs = 284
		goto _test_eof
	_test_eof285:
		lex.cs = 285
		goto _test_eof
	_test_eof286:
		lex.cs = 286
		goto _test_eof
	_test_eof287:
		lex.cs = 287
		goto _test_eof
	_test_eof288:
		lex.cs = 288
		goto _test_eof
	_test_eof289:
		lex.cs = 289
		goto _test_eof
	_test_eof290:
		lex.cs = 290
		goto _test_eof
	_test_eof291:
		lex.cs = 291
		goto _test_eof
	_test_eof292:
		lex.cs = 292
		goto _test_eof
	_test_eof293:
		lex.cs = 293
		goto _test_eof
	_test_eof294:
		lex.cs = 294
		goto _test_eof
	_test_eof295:
		lex.cs = 295
		goto _test_eof
	_test_eof296:
		lex.cs = 296
		goto _test_eof
	_test_eof297:
		lex.cs = 297
		goto _test_eof
	_test_eof298:
		lex.cs = 298
		goto _test_eof
	_test_eof299:
		lex.cs = 299
		goto _test_eof
	_test_eof300:
		lex.cs = 300
		goto _test_eof
	_test_eof301:
		lex.cs = 301
		goto _test_eof
	_test_eof302:
		lex.cs = 302
		goto _test_eof
	_test_eof303:
		lex.cs = 303
		goto _test_eof
	_test_eof304:
		lex.cs = 304
		goto _test_eof
	_test_eof305:
		lex.cs = 305
		goto _test_eof
	_test_eof306:
		lex.cs = 306
		goto _test_eof
	_test_eof307:
		lex.cs = 307
		goto _test_eof
	_test_eof308:
		lex.cs = 308
		goto _test_eof
	_test_eof309:
		lex.cs = 309
		goto _test_eof
	_test_eof310:
		lex.cs = 310
		goto _test_eof
	_test_eof311:
		lex.cs = 311
		goto _test_eof
	_test_eof312:
		lex.cs = 312
		goto _test_eof
	_test_eof313:
		lex.cs = 313
		goto _test_eof
	_test_eof314:
		lex.cs = 314
		goto _test_eof
	_test_eof315:
		lex.cs = 315
		goto _test_eof
	_test_eof316:
		lex.cs = 316
		goto _test_eof
	_test_eof317:
		lex.cs = 317
		goto _test_eof
	_test_eof318:
		lex.cs = 318
		goto _test_eof
	_test_eof319:
		lex.cs = 319
		goto _test_eof
	_test_eof320:
		lex.cs = 320
		goto _test_eof
	_test_eof321:
		lex.cs = 321
		goto _test_eof
	_test_eof322:
		lex.cs = 322
		goto _test_eof
	_test_eof323:
		lex.cs = 323
		goto _test_eof
	_test_eof324:
		lex.cs = 324
		goto _test_eof
	_test_eof325:
		lex.cs = 325
		goto _test_eof
	_test_eof326:
		lex.cs = 326
		goto _test_eof
	_test_eof327:
		lex.cs = 327
		goto _test_eof
	_test_eof328:
		lex.cs = 328
		goto _test_eof
	_test_eof329:
		lex.cs = 329
		goto _test_eof
	_test_eof330:
		lex.cs = 330
		goto _test_eof
	_test_eof331:
		lex.cs = 331
		goto _test_eof
	_test_eof332:
		lex.cs = 332
		goto _test_eof
	_test_eof333:
		lex.cs = 333
		goto _test_eof
	_test_eof334:
		lex.cs = 334
		goto _test_eof
	_test_eof335:
		lex.cs = 335
		goto _test_eof
	_test_eof336:
		lex.cs = 336
		goto _test_eof
	_test_eof337:
		lex.cs = 337
		goto _test_eof
	_test_eof338:
		lex.cs = 338
		goto _test_eof
	_test_eof339:
		lex.cs = 339
		goto _test_eof
	_test_eof340:
		lex.cs = 340
		goto _test_eof
	_test_eof341:
		lex.cs = 341
		goto _test_eof
	_test_eof342:
		lex.cs = 342
		goto _test_eof
	_test_eof343:
		lex.cs = 343
		goto _test_eof
	_test_eof344:
		lex.cs = 344
		goto _test_eof
	_test_eof345:
		lex.cs = 345
		goto _test_eof
	_test_eof346:
		lex.cs = 346
		goto _test_eof
	_test_eof347:
		lex.cs = 347
		goto _test_eof
	_test_eof348:
		lex.cs = 348
		goto _test_eof
	_test_eof349:
		lex.cs = 349
		goto _test_eof
	_test_eof350:
		lex.cs = 350
		goto _test_eof
	_test_eof351:
		lex.cs = 351
		goto _test_eof
	_test_eof352:
		lex.cs = 352
		goto _test_eof
	_test_eof353:
		lex.cs = 353
		goto _test_eof
	_test_eof354:
		lex.cs = 354
		goto _test_eof
	_test_eof355:
		lex.cs = 355
		goto _test_eof
	_test_eof356:
		lex.cs = 356
		goto _test_eof
	_test_eof357:
		lex.cs = 357
		goto _test_eof
	_test_eof358:
		lex.cs = 358
		goto _test_eof
	_test_eof359:
		lex.cs = 359
		goto _test_eof
	_test_eof360:
		lex.cs = 360
		goto _test_eof
	_test_eof361:
		lex.cs = 361
		goto _test_eof
	_test_eof362:
		lex.cs = 362
		goto _test_eof
	_test_eof363:
		lex.cs = 363
		goto _test_eof
	_test_eof364:
		lex.cs = 364
		goto _test_eof
	_test_eof365:
		lex.cs = 365
		goto _test_eof
	_test_eof366:
		lex.cs = 366
		goto _test_eof
	_test_eof367:
		lex.cs = 367
		goto _test_eof
	_test_eof368:
		lex.cs = 368
		goto _test_eof
	_test_eof369:
		lex.cs = 369
		goto _test_eof
	_test_eof370:
		lex.cs = 370
		goto _test_eof
	_test_eof371:
		lex.cs = 371
		goto _test_eof
	_test_eof372:
		lex.cs = 372
		goto _test_eof
	_test_eof373:
		lex.cs = 373
		goto _test_eof
	_test_eof374:
		lex.cs = 374
		goto _test_eof
	_test_eof375:
		lex.cs = 375
		goto _test_eof
	_test_eof376:
		lex.cs = 376
		goto _test_eof
	_test_eof377:
		lex.cs = 377
		goto _test_eof
	_test_eof378:
		lex.cs = 378
		goto _test_eof
	_test_eof379:
		lex.cs = 379
		goto _test_eof
	_test_eof380:
		lex.cs = 380
		goto _test_eof
	_test_eof381:
		lex.cs = 381
		goto _test_eof
	_test_eof382:
		lex.cs = 382
		goto _test_eof
	_test_eof383:
		lex.cs = 383
		goto _test_eof
	_test_eof384:
		lex.cs = 384
		goto _test_eof
	_test_eof385:
		lex.cs = 385
		goto _test_eof
	_test_eof386:
		lex.cs = 386
		goto _test_eof
	_test_eof387:
		lex.cs = 387
		goto _test_eof
	_test_eof388:
		lex.cs = 388
		goto _test_eof
	_test_eof90:
		lex.cs = 90
		goto _test_eof
	_test_eof91:
		lex.cs = 91
		goto _test_eof
	_test_eof92:
		lex.cs = 92
		goto _test_eof
	_test_eof93:
		lex.cs = 93
		goto _test_eof
	_test_eof94:
		lex.cs = 94
		goto _test_eof
	_test_eof95:
		lex.cs = 95
		goto _test_eof
	_test_eof389:
		lex.cs = 389
		goto _test_eof
	_test_eof390:
		lex.cs = 390
		goto _test_eof
	_test_eof391:
		lex.cs = 391
		goto _test_eof
	_test_eof392:
		lex.cs = 392
		goto _test_eof
	_test_eof393:
		lex.cs = 393
		goto _test_eof
	_test_eof394:
		lex.cs = 394
		goto _test_eof
	_test_eof395:
		lex.cs = 395
		goto _test_eof
	_test_eof396:
		lex.cs = 396
		goto _test_eof
	_test_eof397:
		lex.cs = 397
		goto _test_eof
	_test_eof398:
		lex.cs = 398
		goto _test_eof
	_test_eof399:
		lex.cs = 399
		goto _test_eof
	_test_eof400:
		lex.cs = 400
		goto _test_eof
	_test_eof401:
		lex.cs = 401
		goto _test_eof
	_test_eof402:
		lex.cs = 402
		goto _test_eof
	_test_eof403:
		lex.cs = 403
		goto _test_eof
	_test_eof404:
		lex.cs = 404
		goto _test_eof
	_test_eof405:
		lex.cs = 405
		goto _test_eof
	_test_eof406:
		lex.cs = 406
		goto _test_eof
	_test_eof407:
		lex.cs = 407
		goto _test_eof
	_test_eof408:
		lex.cs = 408
		goto _test_eof
	_test_eof409:
		lex.cs = 409
		goto _test_eof
	_test_eof410:
		lex.cs = 410
		goto _test_eof
	_test_eof411:
		lex.cs = 411
		goto _test_eof
	_test_eof412:
		lex.cs = 412
		goto _test_eof
	_test_eof413:
		lex.cs = 413
		goto _test_eof
	_test_eof414:
		lex.cs = 414
		goto _test_eof
	_test_eof415:
		lex.cs = 415
		goto _test_eof
	_test_eof416:
		lex.cs = 416
		goto _test_eof
	_test_eof417:
		lex.cs = 417
		goto _test_eof
	_test_eof418:
		lex.cs = 418
		goto _test_eof
	_test_eof419:
		lex.cs = 419
		goto _test_eof
	_test_eof420:
		lex.cs = 420
		goto _test_eof
	_test_eof421:
		lex.cs = 421
		goto _test_eof
	_test_eof422:
		lex.cs = 422
		goto _test_eof
	_test_eof423:
		lex.cs = 423
		goto _test_eof
	_test_eof424:
		lex.cs = 424
		goto _test_eof
	_test_eof425:
		lex.cs = 425
		goto _test_eof
	_test_eof426:
		lex.cs = 426
		goto _test_eof
	_test_eof427:
		lex.cs = 427
		goto _test_eof
	_test_eof428:
		lex.cs = 428
		goto _test_eof
	_test_eof429:
		lex.cs = 429
		goto _test_eof
	_test_eof430:
		lex.cs = 430
		goto _test_eof
	_test_eof431:
		lex.cs = 431
		goto _test_eof
	_test_eof432:
		lex.cs = 432
		goto _test_eof
	_test_eof433:
		lex.cs = 433
		goto _test_eof
	_test_eof434:
		lex.cs = 434
		goto _test_eof
	_test_eof435:
		lex.cs = 435
		goto _test_eof
	_test_eof436:
		lex.cs = 436
		goto _test_eof
	_test_eof437:
		lex.cs = 437
		goto _test_eof
	_test_eof438:
		lex.cs = 438
		goto _test_eof
	_test_eof439:
		lex.cs = 439
		goto _test_eof
	_test_eof440:
		lex.cs = 440
		goto _test_eof
	_test_eof441:
		lex.cs = 441
		goto _test_eof
	_test_eof442:
		lex.cs = 442
		goto _test_eof
	_test_eof443:
		lex.cs = 443
		goto _test_eof
	_test_eof444:
		lex.cs = 444
		goto _test_eof
	_test_eof445:
		lex.cs = 445
		goto _test_eof
	_test_eof446:
		lex.cs = 446
		goto _test_eof
	_test_eof447:
		lex.cs = 447
		goto _test_eof
	_test_eof448:
		lex.cs = 448
		goto _test_eof
	_test_eof449:
		lex.cs = 449
		goto _test_eof
	_test_eof450:
		lex.cs = 450
		goto _test_eof
	_test_eof451:
		lex.cs = 451
		goto _test_eof
	_test_eof452:
		lex.cs = 452
		goto _test_eof
	_test_eof453:
		lex.cs = 453
		goto _test_eof
	_test_eof454:
		lex.cs = 454
		goto _test_eof
	_test_eof455:
		lex.cs = 455
		goto _test_eof
	_test_eof456:
		lex.cs = 456
		goto _test_eof
	_test_eof457:
		lex.cs = 457
		goto _test_eof
	_test_eof458:
		lex.cs = 458
		goto _test_eof
	_test_eof459:
		lex.cs = 459
		goto _test_eof
	_test_eof460:
		lex.cs = 460
		goto _test_eof
	_test_eof461:
		lex.cs = 461
		goto _test_eof
	_test_eof96:
		lex.cs = 96
		goto _test_eof
	_test_eof462:
		lex.cs = 462
		goto _test_eof
	_test_eof463:
		lex.cs = 463
		goto _test_eof
	_test_eof464:
		lex.cs = 464
		goto _test_eof
	_test_eof465:
		lex.cs = 465
		goto _test_eof
	_test_eof466:
		lex.cs = 466
		goto _test_eof
	_test_eof467:
		lex.cs = 467
		goto _test_eof
	_test_eof468:
		lex.cs = 468
		goto _test_eof
	_test_eof469:
		lex.cs = 469
		goto _test_eof
	_test_eof97:
		lex.cs = 97
		goto _test_eof
	_test_eof470:
		lex.cs = 470
		goto _test_eof
	_test_eof471:
		lex.cs = 471
		goto _test_eof
	_test_eof472:
		lex.cs = 472
		goto _test_eof
	_test_eof473:
		lex.cs = 473
		goto _test_eof
	_test_eof474:
		lex.cs = 474
		goto _test_eof
	_test_eof475:
		lex.cs = 475
		goto _test_eof
	_test_eof98:
		lex.cs = 98
		goto _test_eof
	_test_eof476:
		lex.cs = 476
		goto _test_eof
	_test_eof477:
		lex.cs = 477
		goto _test_eof
	_test_eof478:
		lex.cs = 478
		goto _test_eof
	_test_eof479:
		lex.cs = 479
		goto _test_eof
	_test_eof480:
		lex.cs = 480
		goto _test_eof
	_test_eof481:
		lex.cs = 481
		goto _test_eof
	_test_eof99:
		lex.cs = 99
		goto _test_eof
	_test_eof482:
		lex.cs = 482
		goto _test_eof
	_test_eof483:
		lex.cs = 483
		goto _test_eof
	_test_eof484:
		lex.cs = 484
		goto _test_eof
	_test_eof485:
		lex.cs = 485
		goto _test_eof
	_test_eof486:
		lex.cs = 486
		goto _test_eof
	_test_eof487:
		lex.cs = 487
		goto _test_eof
	_test_eof488:
		lex.cs = 488
		goto _test_eof
	_test_eof489:
		lex.cs = 489
		goto _test_eof
	_test_eof490:
		lex.cs = 490
		goto _test_eof
	_test_eof491:
		lex.cs = 491
		goto _test_eof
	_test_eof100:
		lex.cs = 100
		goto _test_eof
	_test_eof492:
		lex.cs = 492
		goto _test_eof
	_test_eof493:
		lex.cs = 493
		goto _test_eof
	_test_eof494:
		lex.cs = 494
		goto _test_eof
	_test_eof495:
		lex.cs = 495
		goto _test_eof
	_test_eof496:
		lex.cs = 496
		goto _test_eof
	_test_eof497:
		lex.cs = 497
		goto _test_eof
	_test_eof498:
		lex.cs = 498
		goto _test_eof
	_test_eof499:
		lex.cs = 499
		goto _test_eof
	_test_eof101:
		lex.cs = 101
		goto _test_eof
	_test_eof500:
		lex.cs = 500
		goto _test_eof
	_test_eof102:
		lex.cs = 102
		goto _test_eof
	_test_eof501:
		lex.cs = 501
		goto _test_eof
	_test_eof502:
		lex.cs = 502
		goto _test_eof
	_test_eof503:
		lex.cs = 503
		goto _test_eof
	_test_eof504:
		lex.cs = 504
		goto _test_eof
	_test_eof103:
		lex.cs = 103
		goto _test_eof
	_test_eof505:
		lex.cs = 505
		goto _test_eof
	_test_eof506:
		lex.cs = 506
		goto _test_eof
	_test_eof507:
		lex.cs = 507
		goto _test_eof
	_test_eof104:
		lex.cs = 104
		goto _test_eof
	_test_eof508:
		lex.cs = 508
		goto _test_eof
	_test_eof509:
		lex.cs = 509
		goto _test_eof
	_test_eof510:
		lex.cs = 510
		goto _test_eof
	_test_eof511:
		lex.cs = 511
		goto _test_eof
	_test_eof105:
		lex.cs = 105
		goto _test_eof
	_test_eof512:
		lex.cs = 512
		goto _test_eof
	_test_eof513:
		lex.cs = 513
		goto _test_eof
	_test_eof514:
		lex.cs = 514
		goto _test_eof
	_test_eof515:
		lex.cs = 515
		goto _test_eof
	_test_eof106:
		lex.cs = 106
		goto _test_eof
	_test_eof516:
		lex.cs = 516
		goto _test_eof
	_test_eof517:
		lex.cs = 517
		goto _test_eof
	_test_eof518:
		lex.cs = 518
		goto _test_eof
	_test_eof519:
		lex.cs = 519
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 108:
				goto tr158
			case 109:
				goto tr160
			case 110:
				goto tr158
			case 111:
				goto tr158
			case 112:
				goto tr165
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr0
			case 113:
				goto tr168
			case 4:
				goto tr0
			case 115:
				goto tr222
			case 116:
				goto tr224
			case 5:
				goto tr6
			case 117:
				goto tr228
			case 118:
				goto tr229
			case 119:
				goto tr231
			case 120:
				goto tr233
			case 6:
				goto tr8
			case 7:
				goto tr8
			case 8:
				goto tr8
			case 9:
				goto tr8
			case 121:
				goto tr234
			case 122:
				goto tr236
			case 123:
				goto tr229
			case 124:
				goto tr240
			case 125:
				goto tr229
			case 126:
				goto tr229
			case 127:
				goto tr228
			case 10:
				goto tr15
			case 11:
				goto tr15
			case 128:
				goto tr229
			case 12:
				goto tr19
			case 13:
				goto tr19
			case 14:
				goto tr19
			case 15:
				goto tr19
			case 16:
				goto tr19
			case 17:
				goto tr19
			case 18:
				goto tr19
			case 19:
				goto tr19
			case 20:
				goto tr19
			case 21:
				goto tr19
			case 22:
				goto tr19
			case 23:
				goto tr19
			case 24:
				goto tr19
			case 25:
				goto tr19
			case 26:
				goto tr19
			case 27:
				goto tr19
			case 28:
				goto tr19
			case 29:
				goto tr19
			case 30:
				goto tr19
			case 31:
				goto tr19
			case 32:
				goto tr19
			case 33:
				goto tr19
			case 34:
				goto tr19
			case 35:
				goto tr19
			case 36:
				goto tr19
			case 37:
				goto tr19
			case 38:
				goto tr19
			case 39:
				goto tr19
			case 40:
				goto tr19
			case 41:
				goto tr19
			case 42:
				goto tr19
			case 43:
				goto tr19
			case 44:
				goto tr19
			case 45:
				goto tr19
			case 46:
				goto tr19
			case 47:
				goto tr19
			case 48:
				goto tr19
			case 49:
				goto tr19
			case 50:
				goto tr19
			case 51:
				goto tr19
			case 52:
				goto tr19
			case 53:
				goto tr19
			case 54:
				goto tr19
			case 55:
				goto tr19
			case 56:
				goto tr19
			case 57:
				goto tr19
			case 58:
				goto tr19
			case 59:
				goto tr19
			case 60:
				goto tr19
			case 61:
				goto tr19
			case 62:
				goto tr19
			case 63:
				goto tr19
			case 64:
				goto tr19
			case 65:
				goto tr19
			case 129:
				goto tr229
			case 130:
				goto tr246
			case 131:
				goto tr229
			case 132:
				goto tr229
			case 133:
				goto tr229
			case 66:
				goto tr19
			case 134:
				goto tr256
			case 67:
				goto tr8
			case 68:
				goto tr8
			case 135:
				goto tr256
			case 136:
				goto tr229
			case 69:
				goto tr19
			case 70:
				goto tr19
			case 71:
				goto tr19
			case 137:
				goto tr259
			case 138:
				goto tr259
			case 72:
				goto tr91
			case 139:
				goto tr262
			case 73:
				goto tr91
			case 140:
				goto tr263
			case 141:
				goto tr229
			case 142:
				goto tr229
			case 74:
				goto tr19
			case 75:
				goto tr19
			case 76:
				goto tr19
			case 77:
				goto tr19
			case 143:
				goto tr265
			case 144:
				goto tr267
			case 78:
				goto tr103
			case 145:
				goto tr229
			case 146:
				goto tr271
			case 79:
				goto tr8
			case 80:
				goto tr8
			case 81:
				goto tr8
			case 82:
				goto tr8
			case 147:
				goto tr273
			case 83:
				goto tr8
			case 84:
				goto tr8
			case 85:
				goto tr8
			case 86:
				goto tr8
			case 148:
				goto tr274
			case 149:
				goto tr229
			case 150:
				goto tr278
			case 151:
				goto tr229
			case 152:
				goto tr282
			case 153:
				goto tr229
			case 154:
				goto tr286
			case 155:
				goto tr288
			case 87:
				goto tr119
			case 156:
				goto tr289
			case 157:
				goto tr8
			case 158:
				goto tr289
			case 159:
				goto tr289
			case 160:
				goto tr289
			case 161:
				goto tr289
			case 162:
				goto tr289
			case 163:
				goto tr289
			case 164:
				goto tr289
			case 165:
				goto tr289
			case 166:
				goto tr289
			case 167:
				goto tr289
			case 168:
				goto tr289
			case 88:
				goto tr121
			case 89:
				goto tr121
			case 169:
				goto tr289
			case 170:
				goto tr289
			case 171:
				goto tr289
			case 172:
				goto tr289
			case 173:
				goto tr289
			case 174:
				goto tr289
			case 175:
				goto tr289
			case 176:
				goto tr289
			case 177:
				goto tr289
			case 178:
				goto tr289
			case 179:
				goto tr289
			case 180:
				goto tr289
			case 181:
				goto tr289
			case 182:
				goto tr289
			case 183:
				goto tr289
			case 184:
				goto tr289
			case 185:
				goto tr289
			case 186:
				goto tr289
			case 187:
				goto tr289
			case 188:
				goto tr289
			case 189:
				goto tr289
			case 190:
				goto tr289
			case 191:
				goto tr289
			case 192:
				goto tr289
			case 193:
				goto tr289
			case 194:
				goto tr289
			case 195:
				goto tr289
			case 196:
				goto tr289
			case 197:
				goto tr289
			case 198:
				goto tr289
			case 199:
				goto tr289
			case 200:
				goto tr289
			case 201:
				goto tr289
			case 202:
				goto tr289
			case 203:
				goto tr289
			case 204:
				goto tr289
			case 205:
				goto tr289
			case 206:
				goto tr289
			case 207:
				goto tr289
			case 208:
				goto tr289
			case 209:
				goto tr289
			case 210:
				goto tr289
			case 211:
				goto tr289
			case 212:
				goto tr289
			case 213:
				goto tr289
			case 214:
				goto tr289
			case 215:
				goto tr289
			case 216:
				goto tr289
			case 217:
				goto tr369
			case 218:
				goto tr289
			case 219:
				goto tr289
			case 220:
				goto tr289
			case 221:
				goto tr289
			case 222:
				goto tr289
			case 223:
				goto tr289
			case 224:
				goto tr289
			case 225:
				goto tr289
			case 226:
				goto tr289
			case 227:
				goto tr289
			case 228:
				goto tr289
			case 229:
				goto tr289
			case 230:
				goto tr289
			case 231:
				goto tr289
			case 232:
				goto tr389
			case 233:
				goto tr289
			case 234:
				goto tr289
			case 235:
				goto tr289
			case 236:
				goto tr289
			case 237:
				goto tr289
			case 238:
				goto tr289
			case 239:
				goto tr289
			case 240:
				goto tr289
			case 241:
				goto tr289
			case 242:
				goto tr289
			case 243:
				goto tr289
			case 244:
				goto tr289
			case 245:
				goto tr289
			case 246:
				goto tr289
			case 247:
				goto tr289
			case 248:
				goto tr289
			case 249:
				goto tr289
			case 250:
				goto tr289
			case 251:
				goto tr289
			case 252:
				goto tr289
			case 253:
				goto tr289
			case 254:
				goto tr289
			case 255:
				goto tr289
			case 256:
				goto tr289
			case 257:
				goto tr289
			case 258:
				goto tr417
			case 259:
				goto tr289
			case 260:
				goto tr289
			case 261:
				goto tr421
			case 262:
				goto tr289
			case 263:
				goto tr289
			case 264:
				goto tr289
			case 265:
				goto tr289
			case 266:
				goto tr289
			case 267:
				goto tr289
			case 268:
				goto tr289
			case 269:
				goto tr289
			case 270:
				goto tr289
			case 271:
				goto tr289
			case 272:
				goto tr289
			case 273:
				goto tr289
			case 274:
				goto tr289
			case 275:
				goto tr289
			case 276:
				goto tr289
			case 277:
				goto tr289
			case 278:
				goto tr289
			case 279:
				goto tr289
			case 280:
				goto tr289
			case 281:
				goto tr289
			case 282:
				goto tr289
			case 283:
				goto tr289
			case 284:
				goto tr289
			case 285:
				goto tr289
			case 286:
				goto tr453
			case 287:
				goto tr289
			case 288:
				goto tr289
			case 289:
				goto tr289
			case 290:
				goto tr289
			case 291:
				goto tr289
			case 292:
				goto tr289
			case 293:
				goto tr289
			case 294:
				goto tr289
			case 295:
				goto tr289
			case 296:
				goto tr289
			case 297:
				goto tr289
			case 298:
				goto tr289
			case 299:
				goto tr289
			case 300:
				goto tr289
			case 301:
				goto tr289
			case 302:
				goto tr289
			case 303:
				goto tr289
			case 304:
				goto tr289
			case 305:
				goto tr289
			case 306:
				goto tr289
			case 307:
				goto tr289
			case 308:
				goto tr289
			case 309:
				goto tr289
			case 310:
				goto tr289
			case 311:
				goto tr289
			case 312:
				goto tr289
			case 313:
				goto tr289
			case 314:
				goto tr289
			case 315:
				goto tr289
			case 316:
				goto tr289
			case 317:
				goto tr289
			case 318:
				goto tr289
			case 319:
				goto tr289
			case 320:
				goto tr289
			case 321:
				goto tr289
			case 322:
				goto tr289
			case 323:
				goto tr289
			case 324:
				goto tr289
			case 325:
				goto tr289
			case 326:
				goto tr289
			case 327:
				goto tr289
			case 328:
				goto tr289
			case 329:
				goto tr289
			case 330:
				goto tr289
			case 331:
				goto tr289
			case 332:
				goto tr289
			case 333:
				goto tr289
			case 334:
				goto tr289
			case 335:
				goto tr289
			case 336:
				goto tr289
			case 337:
				goto tr289
			case 338:
				goto tr289
			case 339:
				goto tr289
			case 340:
				goto tr289
			case 341:
				goto tr289
			case 342:
				goto tr289
			case 343:
				goto tr289
			case 344:
				goto tr289
			case 345:
				goto tr289
			case 346:
				goto tr289
			case 347:
				goto tr521
			case 348:
				goto tr289
			case 349:
				goto tr289
			case 350:
				goto tr289
			case 351:
				goto tr289
			case 352:
				goto tr289
			case 353:
				goto tr289
			case 354:
				goto tr289
			case 355:
				goto tr289
			case 356:
				goto tr289
			case 357:
				goto tr289
			case 358:
				goto tr289
			case 359:
				goto tr289
			case 360:
				goto tr289
			case 361:
				goto tr289
			case 362:
				goto tr289
			case 363:
				goto tr289
			case 364:
				goto tr289
			case 365:
				goto tr289
			case 366:
				goto tr289
			case 367:
				goto tr289
			case 368:
				goto tr289
			case 369:
				goto tr289
			case 370:
				goto tr289
			case 371:
				goto tr289
			case 372:
				goto tr289
			case 373:
				goto tr289
			case 374:
				goto tr289
			case 375:
				goto tr289
			case 376:
				goto tr289
			case 377:
				goto tr289
			case 378:
				goto tr289
			case 379:
				goto tr289
			case 380:
				goto tr289
			case 381:
				goto tr289
			case 382:
				goto tr289
			case 383:
				goto tr289
			case 384:
				goto tr289
			case 385:
				goto tr289
			case 386:
				goto tr289
			case 387:
				goto tr289
			case 388:
				goto tr567
			case 90:
				goto tr123
			case 91:
				goto tr123
			case 92:
				goto tr123
			case 93:
				goto tr123
			case 94:
				goto tr123
			case 95:
				goto tr123
			case 389:
				goto tr289
			case 390:
				goto tr289
			case 391:
				goto tr289
			case 392:
				goto tr229
			case 393:
				goto tr289
			case 394:
				goto tr289
			case 395:
				goto tr289
			case 396:
				goto tr289
			case 397:
				goto tr289
			case 398:
				goto tr289
			case 399:
				goto tr289
			case 400:
				goto tr289
			case 401:
				goto tr289
			case 402:
				goto tr289
			case 403:
				goto tr289
			case 404:
				goto tr289
			case 405:
				goto tr289
			case 406:
				goto tr289
			case 407:
				goto tr289
			case 408:
				goto tr289
			case 409:
				goto tr289
			case 410:
				goto tr289
			case 411:
				goto tr289
			case 412:
				goto tr289
			case 413:
				goto tr289
			case 414:
				goto tr289
			case 415:
				goto tr289
			case 416:
				goto tr289
			case 417:
				goto tr289
			case 418:
				goto tr289
			case 419:
				goto tr289
			case 420:
				goto tr289
			case 421:
				goto tr289
			case 422:
				goto tr289
			case 423:
				goto tr289
			case 424:
				goto tr289
			case 425:
				goto tr289
			case 426:
				goto tr289
			case 427:
				goto tr289
			case 428:
				goto tr289
			case 429:
				goto tr289
			case 430:
				goto tr289
			case 431:
				goto tr289
			case 432:
				goto tr289
			case 433:
				goto tr289
			case 434:
				goto tr289
			case 435:
				goto tr289
			case 436:
				goto tr289
			case 437:
				goto tr289
			case 438:
				goto tr289
			case 439:
				goto tr289
			case 440:
				goto tr289
			case 441:
				goto tr289
			case 442:
				goto tr289
			case 443:
				goto tr289
			case 444:
				goto tr289
			case 445:
				goto tr289
			case 446:
				goto tr289
			case 447:
				goto tr289
			case 448:
				goto tr289
			case 449:
				goto tr289
			case 450:
				goto tr289
			case 451:
				goto tr289
			case 452:
				goto tr289
			case 453:
				goto tr289
			case 454:
				goto tr289
			case 455:
				goto tr289
			case 456:
				goto tr289
			case 457:
				goto tr289
			case 458:
				goto tr229
			case 460:
				goto tr653
			case 461:
				goto tr655
			case 96:
				goto tr135
			case 462:
				goto tr659
			case 463:
				goto tr659
			case 464:
				goto tr661
			case 466:
				goto tr664
			case 467:
				goto tr665
			case 469:
				goto tr674
			case 470:
				goto tr676
			case 471:
				goto tr677
			case 472:
				goto tr674
			case 473:
				goto tr681
			case 475:
				goto tr691
			case 476:
				goto tr693
			case 477:
				goto tr694
			case 478:
				goto tr691
			case 479:
				goto tr698
			case 481:
				goto tr708
			case 482:
				goto tr710
			case 483:
				goto tr711
			case 484:
				goto tr708
			case 485:
				goto tr715
			case 487:
				goto tr718
			case 489:
				goto tr724
			case 490:
				goto tr726
			case 491:
				goto tr724
			case 100:
				goto tr141
			case 492:
				goto tr728
			case 494:
				goto tr739
			case 495:
				goto tr740
			case 496:
				goto tr741
			case 497:
				goto tr743
			case 498:
				goto tr744
			case 499:
				goto tr744
			case 101:
				goto tr143
			case 500:
				goto tr744
			case 102:
				goto tr143
			case 501:
				goto tr744
			case 502:
				goto tr747
			case 504:
				goto tr750
			case 103:
				goto tr146
			case 506:
				goto tr755
			case 507:
				goto tr757
			case 104:
				goto tr149
			case 508:
				goto tr761
			case 510:
				goto tr766
			case 511:
				goto tr768
			case 105:
				goto tr151
			case 512:
				goto tr772
			case 514:
				goto tr777
			case 515:
				goto tr779
			case 106:
				goto tr153
			case 516:
				goto tr783
			case 518:
				goto tr786
			case 519:
				goto tr787
			}
		}

	_out:
		{
		}
	}

//line scanner/scanner.rl:492

	token.FreeFloating = lex.FreeFloating
	token.Value = string(lex.data[lex.ts:lex.te])

	lval.Token(token)

	return int(tok)
}
