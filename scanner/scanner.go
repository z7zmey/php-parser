//line scanner/scanner.rl:1
package scanner

import (
	"fmt"

	"github.com/z7zmey/php-parser/freefloating"
)

//line scanner/scanner.go:13
const lexer_start int = 111
const lexer_first_final int = 111
const lexer_error int = 0

const lexer_en_main int = 111
const lexer_en_php int = 118
const lexer_en_property int = 466
const lexer_en_nowdoc int = 472
const lexer_en_heredoc int = 475
const lexer_en_backqote int = 481
const lexer_en_template_string int = 487
const lexer_en_heredoc_end int = 493
const lexer_en_string_var int = 495
const lexer_en_string_var_index int = 500
const lexer_en_string_var_name int = 511
const lexer_en_halt_compiller_open_parenthesis int = 513
const lexer_en_halt_compiller_close_parenthesis int = 517
const lexer_en_halt_compiller_close_semicolon int = 521
const lexer_en_halt_compiller_end int = 525

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

	lblStart := 0
	lblEnd := 0

	_, _ = lblStart, lblEnd

//line scanner/scanner.go:72
	{
		var _widec int16
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 111:
			goto st111
		case 112:
			goto st112
		case 113:
			goto st113
		case 114:
			goto st114
		case 115:
			goto st115
		case 116:
			goto st116
		case 1:
			goto st1
		case 2:
			goto st2
		case 3:
			goto st3
		case 117:
			goto st117
		case 4:
			goto st4
		case 118:
			goto st118
		case 119:
			goto st119
		case 120:
			goto st120
		case 5:
			goto st5
		case 121:
			goto st121
		case 122:
			goto st122
		case 123:
			goto st123
		case 124:
			goto st124
		case 6:
			goto st6
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 11:
			goto st11
		case 125:
			goto st125
		case 126:
			goto st126
		case 127:
			goto st127
		case 128:
			goto st128
		case 129:
			goto st129
		case 130:
			goto st130
		case 131:
			goto st131
		case 12:
			goto st12
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 132:
			goto st132
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
		case 66:
			goto st66
		case 67:
			goto st67
		case 68:
			goto st68
		case 69:
			goto st69
		case 133:
			goto st133
		case 134:
			goto st134
		case 135:
			goto st135
		case 136:
			goto st136
		case 137:
			goto st137
		case 70:
			goto st70
		case 138:
			goto st138
		case 71:
			goto st71
		case 72:
			goto st72
		case 139:
			goto st139
		case 140:
			goto st140
		case 73:
			goto st73
		case 74:
			goto st74
		case 75:
			goto st75
		case 141:
			goto st141
		case 142:
			goto st142
		case 76:
			goto st76
		case 143:
			goto st143
		case 77:
			goto st77
		case 144:
			goto st144
		case 145:
			goto st145
		case 146:
			goto st146
		case 147:
			goto st147
		case 148:
			goto st148
		case 149:
			goto st149
		case 150:
			goto st150
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 151:
			goto st151
		case 152:
			goto st152
		case 82:
			goto st82
		case 153:
			goto st153
		case 154:
			goto st154
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 86:
			goto st86
		case 155:
			goto st155
		case 87:
			goto st87
		case 88:
			goto st88
		case 89:
			goto st89
		case 90:
			goto st90
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
		case 91:
			goto st91
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
		case 92:
			goto st92
		case 93:
			goto st93
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
		case 94:
			goto st94
		case 95:
			goto st95
		case 96:
			goto st96
		case 97:
			goto st97
		case 98:
			goto st98
		case 99:
			goto st99
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
		case 462:
			goto st462
		case 463:
			goto st463
		case 464:
			goto st464
		case 465:
			goto st465
		case 466:
			goto st466
		case 467:
			goto st467
		case 468:
			goto st468
		case 100:
			goto st100
		case 469:
			goto st469
		case 470:
			goto st470
		case 471:
			goto st471
		case 472:
			goto st472
		case 0:
			goto st0
		case 473:
			goto st473
		case 474:
			goto st474
		case 475:
			goto st475
		case 476:
			goto st476
		case 101:
			goto st101
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
		case 482:
			goto st482
		case 102:
			goto st102
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
		case 103:
			goto st103
		case 489:
			goto st489
		case 490:
			goto st490
		case 491:
			goto st491
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
		case 104:
			goto st104
		case 499:
			goto st499
		case 500:
			goto st500
		case 501:
			goto st501
		case 502:
			goto st502
		case 503:
			goto st503
		case 504:
			goto st504
		case 505:
			goto st505
		case 506:
			goto st506
		case 105:
			goto st105
		case 507:
			goto st507
		case 106:
			goto st106
		case 508:
			goto st508
		case 509:
			goto st509
		case 510:
			goto st510
		case 511:
			goto st511
		case 512:
			goto st512
		case 107:
			goto st107
		case 513:
			goto st513
		case 514:
			goto st514
		case 515:
			goto st515
		case 108:
			goto st108
		case 516:
			goto st516
		case 517:
			goto st517
		case 518:
			goto st518
		case 519:
			goto st519
		case 109:
			goto st109
		case 520:
			goto st520
		case 521:
			goto st521
		case 522:
			goto st522
		case 523:
			goto st523
		case 110:
			goto st110
		case 524:
			goto st524
		case 525:
			goto st525
		case 526:
			goto st526
		case 527:
			goto st527
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 111:
			goto st_case_111
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 117:
			goto st_case_117
		case 4:
			goto st_case_4
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 5:
			goto st_case_5
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 132:
			goto st_case_132
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
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 70:
			goto st_case_70
		case 138:
			goto st_case_138
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 76:
			goto st_case_76
		case 143:
			goto st_case_143
		case 77:
			goto st_case_77
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 82:
			goto st_case_82
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 155:
			goto st_case_155
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
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
		case 91:
			goto st_case_91
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
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
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
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
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
		case 462:
			goto st_case_462
		case 463:
			goto st_case_463
		case 464:
			goto st_case_464
		case 465:
			goto st_case_465
		case 466:
			goto st_case_466
		case 467:
			goto st_case_467
		case 468:
			goto st_case_468
		case 100:
			goto st_case_100
		case 469:
			goto st_case_469
		case 470:
			goto st_case_470
		case 471:
			goto st_case_471
		case 472:
			goto st_case_472
		case 0:
			goto st_case_0
		case 473:
			goto st_case_473
		case 474:
			goto st_case_474
		case 475:
			goto st_case_475
		case 476:
			goto st_case_476
		case 101:
			goto st_case_101
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
		case 482:
			goto st_case_482
		case 102:
			goto st_case_102
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
		case 103:
			goto st_case_103
		case 489:
			goto st_case_489
		case 490:
			goto st_case_490
		case 491:
			goto st_case_491
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
		case 104:
			goto st_case_104
		case 499:
			goto st_case_499
		case 500:
			goto st_case_500
		case 501:
			goto st_case_501
		case 502:
			goto st_case_502
		case 503:
			goto st_case_503
		case 504:
			goto st_case_504
		case 505:
			goto st_case_505
		case 506:
			goto st_case_506
		case 105:
			goto st_case_105
		case 507:
			goto st_case_507
		case 106:
			goto st_case_106
		case 508:
			goto st_case_508
		case 509:
			goto st_case_509
		case 510:
			goto st_case_510
		case 511:
			goto st_case_511
		case 512:
			goto st_case_512
		case 107:
			goto st_case_107
		case 513:
			goto st_case_513
		case 514:
			goto st_case_514
		case 515:
			goto st_case_515
		case 108:
			goto st_case_108
		case 516:
			goto st_case_516
		case 517:
			goto st_case_517
		case 518:
			goto st_case_518
		case 519:
			goto st_case_519
		case 109:
			goto st_case_109
		case 520:
			goto st_case_520
		case 521:
			goto st_case_521
		case 522:
			goto st_case_522
		case 523:
			goto st_case_523
		case 110:
			goto st_case_110
		case 524:
			goto st_case_524
		case 525:
			goto st_case_525
		case 526:
			goto st_case_526
		case 527:
			goto st_case_527
		}
		goto st_out
	tr0:
		lex.cs = 111
//line scanner/scanner.rl:113
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			lex.cs = 118
		}
		goto _again
	tr3:
		lex.cs = 111
//line scanner/scanner.rl:117
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.ts+5)
			lex.cs = 118
		}
		goto _again
	tr177:
//line scanner/scanner.rl:107
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.createToken(lval)
			tok = T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 111
				goto _out
			}
		}
		goto st111
	tr179:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:107
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.createToken(lval)
			tok = T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 111
				goto _out
			}
		}
		goto st111
	tr184:
		lex.cs = 111
//line scanner/scanner.rl:113
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			lex.cs = 118
		}
		goto _again
	tr185:
		lex.cs = 111
//line scanner/scanner.rl:122
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_ECHO
			lex.cs = 118
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr187:
		lex.cs = 111
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:117
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.ts+5)
			lex.cs = 118
		}
		goto _again
	st111:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof111
		}
	st_case_111:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:2291
		switch lex.data[(lex.p)] {
		case 10:
			goto st113
		case 60:
			goto st115
		}
		goto st112
	tr180:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st112
	st112:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof112
		}
	st_case_112:
//line scanner/scanner.go:2308
		switch lex.data[(lex.p)] {
		case 10:
			goto st113
		case 60:
			goto st114
		}
		goto st112
	tr181:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st113
	st113:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof113
		}
	st_case_113:
//line scanner/scanner.go:2325
		switch lex.data[(lex.p)] {
		case 10:
			goto tr181
		case 60:
			goto tr182
		}
		goto tr180
	tr182:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st114
	st114:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof114
		}
	st_case_114:
//line scanner/scanner.go:2342
		switch lex.data[(lex.p)] {
		case 10:
			goto st113
		case 60:
			goto st114
		case 63:
			goto tr177
		}
		goto st112
	st115:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch lex.data[(lex.p)] {
		case 10:
			goto st113
		case 60:
			goto st114
		case 63:
			goto tr183
		}
		goto st112
	tr183:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st116
	st116:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof116
		}
	st_case_116:
//line scanner/scanner.go:2376
		switch lex.data[(lex.p)] {
		case 61:
			goto tr185
		case 80:
			goto st1
		case 112:
			goto st1
		}
		goto tr184
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
			goto st117
		case 13:
			goto st4
		case 32:
			goto tr3
		}
		goto tr0
	st117:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		goto tr187
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		if lex.data[(lex.p)] == 10 {
			goto st117
		}
		goto tr0
	tr6:
//line scanner/scanner.rl:131
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st118
	tr8:
		lex.cs = 118
//line NONE:1
		switch lex.act {
		case 8:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
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
					lex.createToken(lval)
					tok = T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}
				lex.createToken(lval)
				tok = T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 12:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ABSTRACT
				{
					(lex.p)++
					goto _out
				}
			}
		case 13:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ARRAY
				{
					(lex.p)++
					goto _out
				}
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_AS
				{
					(lex.p)++
					goto _out
				}
			}
		case 15:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_BREAK
				{
					(lex.p)++
					goto _out
				}
			}
		case 16:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CALLABLE
				{
					(lex.p)++
					goto _out
				}
			}
		case 17:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CASE
				{
					(lex.p)++
					goto _out
				}
			}
		case 18:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CATCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 19:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CLASS
				{
					(lex.p)++
					goto _out
				}
			}
		case 20:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CLONE
				{
					(lex.p)++
					goto _out
				}
			}
		case 21:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CONST
				{
					(lex.p)++
					goto _out
				}
			}
		case 22:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CONTINUE
				{
					(lex.p)++
					goto _out
				}
			}
		case 23:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_DECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 24:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_DEFAULT
				{
					(lex.p)++
					goto _out
				}
			}
		case 25:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_DO
				{
					(lex.p)++
					goto _out
				}
			}
		case 26:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ECHO
				{
					(lex.p)++
					goto _out
				}
			}
		case 28:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ELSEIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 29:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_EMPTY
				{
					(lex.p)++
					goto _out
				}
			}
		case 30:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ENDDECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 32:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ENDFOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 33:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ENDIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 34:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ENDSWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 35:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ENDWHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 36:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_EVAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 37:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_EXIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 38:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_EXTENDS
				{
					(lex.p)++
					goto _out
				}
			}
		case 40:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_FINALLY
				{
					(lex.p)++
					goto _out
				}
			}
		case 42:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_FOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 43:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_FUNCTION
				{
					(lex.p)++
					goto _out
				}
			}
		case 44:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_GLOBAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 45:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_GOTO
				{
					(lex.p)++
					goto _out
				}
			}
		case 46:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_IF
				{
					(lex.p)++
					goto _out
				}
			}
		case 47:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_ISSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 48:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_IMPLEMENTS
				{
					(lex.p)++
					goto _out
				}
			}
		case 49:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_INSTANCEOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 50:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_INSTEADOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 51:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_INTERFACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 52:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_LIST
				{
					(lex.p)++
					goto _out
				}
			}
		case 53:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_NAMESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 54:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_PRIVATE
				{
					(lex.p)++
					goto _out
				}
			}
		case 55:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_PUBLIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 56:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_PRINT
				{
					(lex.p)++
					goto _out
				}
			}
		case 57:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_PROTECTED
				{
					(lex.p)++
					goto _out
				}
			}
		case 58:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_RETURN
				{
					(lex.p)++
					goto _out
				}
			}
		case 59:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_STATIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 60:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_SWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 61:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_THROW
				{
					(lex.p)++
					goto _out
				}
			}
		case 62:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_TRAIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 63:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_TRY
				{
					(lex.p)++
					goto _out
				}
			}
		case 64:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_UNSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 65:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_USE
				{
					(lex.p)++
					goto _out
				}
			}
		case 66:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_VAR
				{
					(lex.p)++
					goto _out
				}
			}
		case 67:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_WHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 68:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_YIELD_FROM
				{
					(lex.p)++
					goto _out
				}
			}
		case 71:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_INCLUDE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 73:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_REQUIRE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 74:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_CLASS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 75:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_DIR
				{
					(lex.p)++
					goto _out
				}
			}
		case 76:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_FILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 77:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_FUNC_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 78:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_LINE
				{
					(lex.p)++
					goto _out
				}
			}
		case 79:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_NS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 80:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_METHOD_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 81:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_TRAIT_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 82:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_HALT_COMPILER
				lex.cs = 513
				{
					(lex.p)++
					goto _out
				}
			}
		case 83:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_NEW
				{
					(lex.p)++
					goto _out
				}
			}
		case 84:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_LOGICAL_AND
				{
					(lex.p)++
					goto _out
				}
			}
		case 85:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_LOGICAL_OR
				{
					(lex.p)++
					goto _out
				}
			}
		case 86:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_LOGICAL_XOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 115:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_SL
				{
					(lex.p)++
					goto _out
				}
			}
		case 131:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_STRING
				{
					(lex.p)++
					goto _out
				}
			}
		case 136:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = TokenID(int('"'))
				lex.cs = 487
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr12:
//line scanner/scanner.rl:316
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_CONSTANT_ENCAPSED_STRING
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr21:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:316
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_CONSTANT_ENCAPSED_STRING
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr26:
//line scanner/scanner.rl:339
		(lex.p) = (lex.te) - 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st118
	tr38:
//line scanner/scanner.rl:301
		(lex.p) = (lex.te) - 1
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.createToken(lval)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr53:
//line scanner/scanner.rl:277
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_ARRAY_CAST
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr60:
//line scanner/scanner.rl:282
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_STRING_CAST
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr64:
//line scanner/scanner.rl:278
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_BOOL_CAST
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr72:
//line scanner/scanner.rl:279
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_DOUBLE_CAST
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr79:
//line scanner/scanner.rl:280
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_INT_CAST
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr88:
//line scanner/scanner.rl:281
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_OBJECT_CAST
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr99:
//line scanner/scanner.rl:283
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_UNSET_CAST
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr100:
//line scanner/scanner.rl:246
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_ELLIPSIS
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr109:
//line scanner/scanner.rl:289
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
		goto st118
	tr110:
//line scanner/scanner.rl:149
		(lex.p) = (lex.te) - 1
		{
			if lex.te-lex.ts < 20 {
				lex.createToken(lval)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 118
					goto _out
				}
			}
			lex.createToken(lval)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr122:
		lex.cs = 118
//line scanner/scanner.rl:133
		(lex.p) = (lex.te) - 1
		{
			lex.createToken(lval)
			tok = TokenID(int(';'))
			lex.cs = 111
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr138:
		lex.cs = 118
//line scanner/scanner.rl:132
		(lex.p) = (lex.te) - 1
		{
			lex.createToken(lval)
			tok = TokenID(int(';'))
			lex.cs = 111
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr140:
//line scanner/scanner.rl:312
		(lex.p) = (lex.te) - 1
		{
			lex.createToken(lval)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr142:
//line scanner/scanner.rl:227
		(lex.p) = (lex.te) - 1
		{
			lex.createToken(lval)
			tok = T_YIELD
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr153:
//line scanner/scanner.rl:226
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_YIELD_FROM
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr188:
//line scanner/scanner.rl:339
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st118
	tr199:
//line scanner/scanner.rl:301
		lex.te = (lex.p) + 1
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.createToken(lval)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr235:
//line scanner/scanner.rl:245
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_NS_SEPARATOR
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr238:
		lex.cs = 118
//line scanner/scanner.rl:336
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int('`'))
			lex.cs = 481
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr239:
//line scanner/scanner.rl:309
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int('{'))
			lex.call(118, 118)
			goto _out
		}
		goto st118
	tr241:
//line scanner/scanner.rl:310
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int('}'))
			lex.ret(1)
			lex.PhpDocComment = ""
			goto _out
		}
		goto st118
	tr242:
//line scanner/scanner.rl:131
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st118
	tr244:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:131
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st118
	tr248:
//line scanner/scanner.rl:339
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st118
	tr249:
//line scanner/scanner.rl:301
		lex.te = (lex.p)
		(lex.p)--
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.createToken(lval)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr251:
//line scanner/scanner.rl:264
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr252:
//line scanner/scanner.rl:265
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_IS_NOT_IDENTICAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr253:
		lex.cs = 118
//line scanner/scanner.rl:337
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = TokenID(int('"'))
			lex.cs = 487
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr254:
//line scanner/scanner.rl:285
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
		}
		goto st118
	tr256:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:285
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
		}
		goto st118
	tr260:
//line scanner/scanner.rl:311
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr261:
//line scanner/scanner.rl:259
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_MOD_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr262:
//line scanner/scanner.rl:248
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_BOOLEAN_AND
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr263:
//line scanner/scanner.rl:250
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_AND_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr265:
//line scanner/scanner.rl:253
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_MUL_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr266:
//line scanner/scanner.rl:272
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_POW
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr267:
//line scanner/scanner.rl:254
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_POW_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr268:
//line scanner/scanner.rl:261
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_INC
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr269:
//line scanner/scanner.rl:256
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_PLUS_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr270:
//line scanner/scanner.rl:260
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_DEC
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr271:
//line scanner/scanner.rl:257
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_MINUS_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr272:
		lex.cs = 118
//line scanner/scanner.rl:314
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_OBJECT_OPERATOR
			lex.cs = 466
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr275:
//line scanner/scanner.rl:252
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_CONCAT_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr276:
//line scanner/scanner.rl:135
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr278:
//line scanner/scanner.rl:255
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_DIV_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr279:
//line scanner/scanner.rl:149
		lex.te = (lex.p)
		(lex.p)--
		{
			if lex.te-lex.ts < 20 {
				lex.createToken(lval)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 118
					goto _out
				}
			}
			lex.createToken(lval)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr282:
//line scanner/scanner.rl:136
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
				lex.createToken(lval)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 118
					goto _out
				}
			}
			lex.createToken(lval)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr283:
//line scanner/scanner.rl:155
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
				lex.createToken(lval)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 118
					goto _out
				}
			}
			lex.createToken(lval)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr285:
//line scanner/scanner.rl:312
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr287:
//line scanner/scanner.rl:247
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_PAAMAYIM_NEKUDOTAYIM
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr288:
		lex.cs = 118
//line scanner/scanner.rl:133
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = TokenID(int(';'))
			lex.cs = 111
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr290:
		lex.cs = 118
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:133
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = TokenID(int(';'))
			lex.cs = 111
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr293:
//line scanner/scanner.rl:264
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr294:
//line scanner/scanner.rl:273
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_SL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr295:
//line scanner/scanner.rl:268
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_SL_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr296:
		lex.cs = 118
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:322
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.heredocLabel = lex.data[lblStart:lblEnd]
			lex.createToken(lval)
			tok = T_START_HEREDOC

			if lex.isHeredocEnd(lex.p + 1) {
				lex.cs = 493
			} else if lex.data[lblStart-1] == '\'' {
				lex.cs = 472
			} else {
				lex.cs = 475
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr297:
//line scanner/scanner.rl:271
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_IS_SMALLER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr298:
//line scanner/scanner.rl:263
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_SPACESHIP
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr300:
//line scanner/scanner.rl:262
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_DOUBLE_ARROW
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr301:
//line scanner/scanner.rl:266
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_IS_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr302:
//line scanner/scanner.rl:267
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_IS_IDENTICAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr303:
//line scanner/scanner.rl:270
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_IS_GREATER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr305:
//line scanner/scanner.rl:274
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_SR
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr306:
//line scanner/scanner.rl:269
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_SR_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr308:
//line scanner/scanner.rl:275
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_COALESCE
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr309:
		lex.cs = 118
//line scanner/scanner.rl:132
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = TokenID(int(';'))
			lex.cs = 111
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr311:
		lex.cs = 118
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:132
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = TokenID(int(';'))
			lex.cs = 111
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr391:
//line scanner/scanner.rl:185
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ELSE
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr411:
//line scanner/scanner.rl:189
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENDFOR
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr439:
//line scanner/scanner.rl:197
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_FINAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr443:
//line scanner/scanner.rl:199
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_FOR
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr475:
//line scanner/scanner.rl:228
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_INCLUDE
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr543:
//line scanner/scanner.rl:230
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_REQUIRE
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr589:
//line scanner/scanner.rl:227
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_YIELD
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr594:
//line scanner/scanner.rl:258
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_XOR_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr668:
//line scanner/scanner.rl:251
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	tr669:
//line scanner/scanner.rl:249
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_BOOLEAN_OR
			{
				(lex.p)++
				lex.cs = 118
				goto _out
			}
		}
		goto st118
	st118:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof118
		}
	st_case_118:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:3245
		switch lex.data[(lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto st121
		case 32:
			goto tr189
		case 33:
			goto st122
		case 34:
			goto tr192
		case 35:
			goto st125
		case 36:
			goto st127
		case 37:
			goto st129
		case 38:
			goto st130
		case 39:
			goto tr197
		case 40:
			goto tr198
		case 42:
			goto st133
		case 43:
			goto st135
		case 45:
			goto st136
		case 46:
			goto tr203
		case 47:
			goto tr204
		case 48:
			goto tr205
		case 55:
			goto st145
		case 58:
			goto st149
		case 59:
			goto tr209
		case 60:
			goto st153
		case 61:
			goto st157
		case 62:
			goto st159
		case 63:
			goto st161
		case 64:
			goto tr199
		case 65:
			goto st164
		case 66:
			goto tr215
		case 67:
			goto st179
		case 68:
			goto st208
		case 69:
			goto st219
		case 70:
			goto st261
		case 71:
			goto st272
		case 73:
			goto st279
		case 76:
			goto st318
		case 78:
			goto st321
		case 79:
			goto st330
		case 80:
			goto st331
		case 82:
			goto st348
		case 83:
			goto st362
		case 84:
			goto st371
		case 85:
			goto st378
		case 86:
			goto st383
		case 87:
			goto st385
		case 88:
			goto st389
		case 89:
			goto st391
		case 91:
			goto tr199
		case 92:
			goto tr235
		case 93:
			goto tr199
		case 94:
			goto st399
		case 95:
			goto st400
		case 96:
			goto tr238
		case 97:
			goto st164
		case 98:
			goto tr215
		case 99:
			goto st179
		case 100:
			goto st208
		case 101:
			goto st219
		case 102:
			goto st261
		case 103:
			goto st272
		case 105:
			goto st279
		case 108:
			goto st318
		case 110:
			goto st321
		case 111:
			goto st330
		case 112:
			goto st331
		case 114:
			goto st348
		case 115:
			goto st362
		case 116:
			goto st371
		case 117:
			goto st378
		case 118:
			goto st383
		case 119:
			goto st385
		case 120:
			goto st389
		case 121:
			goto st391
		case 123:
			goto tr239
		case 124:
			goto st465
		case 125:
			goto tr241
		case 126:
			goto tr199
		}
		switch {
		case lex.data[(lex.p)] < 41:
			if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
				goto tr189
			}
		case lex.data[(lex.p)] > 44:
			switch {
			case lex.data[(lex.p)] > 57:
				if 72 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
					goto tr221
				}
			case lex.data[(lex.p)] >= 49:
				goto tr206
			}
		default:
			goto tr199
		}
		goto tr188
	tr189:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st119
	tr245:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st119
	st119:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof119
		}
	st_case_119:
//line scanner/scanner.go:3433
		switch lex.data[(lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto st5
		case 32:
			goto tr189
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr189
		}
		goto tr242
	tr7:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st120
	tr246:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st120
	st120:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof120
		}
	st_case_120:
//line scanner/scanner.go:3463
		switch lex.data[(lex.p)] {
		case 10:
			goto tr246
		case 13:
			goto tr247
		case 32:
			goto tr245
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr245
		}
		goto tr244
	tr247:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line scanner/scanner.go:3485
		if lex.data[(lex.p)] == 10 {
			goto tr7
		}
		goto tr6
	st121:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof121
		}
	st_case_121:
		if lex.data[(lex.p)] == 10 {
			goto tr7
		}
		goto tr248
	st122:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof122
		}
	st_case_122:
		if lex.data[(lex.p)] == 61 {
			goto st123
		}
		goto tr249
	st123:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof123
		}
	st_case_123:
		if lex.data[(lex.p)] == 61 {
			goto tr252
		}
		goto tr251
	tr192:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:337
		lex.act = 136
		goto st124
	st124:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof124
		}
	st_case_124:
//line scanner/scanner.go:3529
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto tr11
		case 34:
			goto tr12
		case 36:
			goto st7
		case 92:
			goto st8
		case 123:
			goto st10
		}
		goto st6
	tr10:
//line scanner/scanner.rl:86
		lex.NewLines.Append(lex.p)
		goto st6
	tr11:
//line scanner/scanner.rl:85
		if lex.p+1 != eof && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}
		goto st6
	tr18:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st6
	tr19:
//line scanner/scanner.rl:86
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st6
	tr20:
//line scanner/scanner.rl:85
		if lex.p+1 != eof && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line scanner/scanner.go:3574
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto tr11
		case 34:
			goto tr12
		case 36:
			goto st7
		case 92:
			goto st8
		case 123:
			goto st10
		}
		goto st6
	tr22:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st7
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line scanner/scanner.go:3599
		switch lex.data[(lex.p)] {
		case 34:
			goto tr12
		case 55:
			goto tr8
		case 92:
			goto st8
		case 95:
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 123 {
				goto tr8
			}
		case lex.data[(lex.p)] >= 65:
			goto tr8
		}
		goto st6
	tr23:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st8
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
//line scanner/scanner.go:3628
		switch lex.data[(lex.p)] {
		case 10:
			goto st9
		case 13:
			goto st11
		}
		goto st6
	tr25:
//line scanner/scanner.rl:86
		lex.NewLines.Append(lex.p)
		goto st9
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line scanner/scanner.go:3645
		switch lex.data[(lex.p)] {
		case 10:
			goto tr19
		case 13:
			goto tr20
		case 34:
			goto tr21
		case 36:
			goto tr22
		case 92:
			goto tr23
		case 123:
			goto tr24
		}
		goto tr18
	tr24:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st10
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
//line scanner/scanner.go:3670
		switch lex.data[(lex.p)] {
		case 34:
			goto tr12
		case 36:
			goto tr8
		case 92:
			goto st8
		}
		goto st6
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr25
		case 13:
			goto tr11
		case 34:
			goto tr12
		case 36:
			goto st7
		case 92:
			goto st8
		case 123:
			goto st10
		}
		goto st6
	tr257:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st125
	st125:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof125
		}
	st_case_125:
//line scanner/scanner.go:3709
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
			goto st126
		}
		if 512 <= _widec && _widec <= 767 {
			goto st125
		}
		goto tr254
	tr258:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st126
	st126:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof126
		}
	st_case_126:
//line scanner/scanner.go:3764
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
			goto tr258
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr257
		}
		goto tr256
	st127:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof127
		}
	st_case_127:
		switch lex.data[(lex.p)] {
		case 55:
			goto st128
		case 95:
			goto st128
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st128
			}
		case lex.data[(lex.p)] >= 65:
			goto st128
		}
		goto tr249
	st128:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof128
		}
	st_case_128:
		if lex.data[(lex.p)] == 95 {
			goto st128
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st128
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st128
			}
		default:
			goto st128
		}
		goto tr260
	st129:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof129
		}
	st_case_129:
		if lex.data[(lex.p)] == 61 {
			goto tr261
		}
		goto tr249
	st130:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof130
		}
	st_case_130:
		switch lex.data[(lex.p)] {
		case 38:
			goto tr262
		case 61:
			goto tr263
		}
		goto tr249
	tr197:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st131
	st131:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof131
		}
	st_case_131:
//line scanner/scanner.go:3882
		switch lex.data[(lex.p)] {
		case 10:
			goto tr28
		case 13:
			goto tr29
		case 39:
			goto tr12
		case 92:
			goto st13
		}
		goto st12
	tr33:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st12
	tr28:
//line scanner/scanner.rl:76
		lex.NewLines.Append(lex.p)
		goto st12
	tr29:
//line scanner/scanner.rl:75
		if lex.p+1 != eof && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}
		goto st12
	tr34:
//line scanner/scanner.rl:76
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st12
	tr35:
//line scanner/scanner.rl:75
		if lex.p+1 != eof && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st12
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
//line scanner/scanner.go:3923
		switch lex.data[(lex.p)] {
		case 10:
			goto tr28
		case 13:
			goto tr29
		case 39:
			goto tr12
		case 92:
			goto st13
		}
		goto st12
	tr36:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st13
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
//line scanner/scanner.go:3944
		switch lex.data[(lex.p)] {
		case 10:
			goto st14
		case 13:
			goto st15
		}
		goto st12
	tr37:
//line scanner/scanner.rl:76
		lex.NewLines.Append(lex.p)
		goto st14
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
//line scanner/scanner.go:3961
		switch lex.data[(lex.p)] {
		case 10:
			goto tr34
		case 13:
			goto tr35
		case 39:
			goto tr21
		case 92:
			goto tr36
		}
		goto tr33
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr37
		case 13:
			goto tr29
		case 39:
			goto tr12
		case 92:
			goto st13
		}
		goto st12
	tr198:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st132
	st132:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof132
		}
	st_case_132:
//line scanner/scanner.go:3999
		switch lex.data[(lex.p)] {
		case 9:
			goto st16
		case 32:
			goto st16
		case 65:
			goto st17
		case 66:
			goto st22
		case 68:
			goto st34
		case 70:
			goto st40
		case 73:
			goto st44
		case 79:
			goto st51
		case 82:
			goto st57
		case 83:
			goto st60
		case 85:
			goto st65
		case 97:
			goto st17
		case 98:
			goto st22
		case 100:
			goto st34
		case 102:
			goto st40
		case 105:
			goto st44
		case 111:
			goto st51
		case 114:
			goto st57
		case 115:
			goto st60
		case 117:
			goto st65
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st16
		}
		goto tr249
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch lex.data[(lex.p)] {
		case 9:
			goto st16
		case 32:
			goto st16
		case 65:
			goto st17
		case 66:
			goto st22
		case 68:
			goto st34
		case 70:
			goto st40
		case 73:
			goto st44
		case 79:
			goto st51
		case 82:
			goto st57
		case 83:
			goto st60
		case 85:
			goto st65
		case 97:
			goto st17
		case 98:
			goto st22
		case 100:
			goto st34
		case 102:
			goto st40
		case 105:
			goto st44
		case 111:
			goto st51
		case 114:
			goto st57
		case 115:
			goto st60
		case 117:
			goto st65
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st16
		}
		goto tr38
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch lex.data[(lex.p)] {
		case 82:
			goto st18
		case 114:
			goto st18
		}
		goto tr38
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 82:
			goto st19
		case 114:
			goto st19
		}
		goto tr38
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 65:
			goto st20
		case 97:
			goto st20
		}
		goto tr38
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 89:
			goto st21
		case 121:
			goto st21
		}
		goto tr38
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch lex.data[(lex.p)] {
		case 9:
			goto st21
		case 32:
			goto st21
		case 41:
			goto tr53
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st21
		}
		goto tr38
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch lex.data[(lex.p)] {
		case 73:
			goto st23
		case 79:
			goto st28
		case 105:
			goto st23
		case 111:
			goto st28
		}
		goto tr38
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 78:
			goto st24
		case 110:
			goto st24
		}
		goto tr38
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 65:
			goto st25
		case 97:
			goto st25
		}
		goto tr38
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 82:
			goto st26
		case 114:
			goto st26
		}
		goto tr38
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 89:
			goto st27
		case 121:
			goto st27
		}
		goto tr38
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
			goto tr60
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st27
		}
		goto tr38
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 79:
			goto st29
		case 111:
			goto st29
		}
		goto tr38
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 76:
			goto st30
		case 108:
			goto st30
		}
		goto tr38
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 9:
			goto st31
		case 32:
			goto st31
		case 41:
			goto tr64
		case 69:
			goto st32
		case 101:
			goto st32
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st31
		}
		goto tr38
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 9:
			goto st31
		case 32:
			goto st31
		case 41:
			goto tr64
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st31
		}
		goto tr38
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 65:
			goto st33
		case 97:
			goto st33
		}
		goto tr38
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 78:
			goto st31
		case 110:
			goto st31
		}
		goto tr38
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 79:
			goto st35
		case 111:
			goto st35
		}
		goto tr38
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 85:
			goto st36
		case 117:
			goto st36
		}
		goto tr38
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 66:
			goto st37
		case 98:
			goto st37
		}
		goto tr38
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch lex.data[(lex.p)] {
		case 76:
			goto st38
		case 108:
			goto st38
		}
		goto tr38
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 69:
			goto st39
		case 101:
			goto st39
		}
		goto tr38
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 9:
			goto st39
		case 32:
			goto st39
		case 41:
			goto tr72
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st39
		}
		goto tr38
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 76:
			goto st41
		case 108:
			goto st41
		}
		goto tr38
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 79:
			goto st42
		case 111:
			goto st42
		}
		goto tr38
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch lex.data[(lex.p)] {
		case 65:
			goto st43
		case 97:
			goto st43
		}
		goto tr38
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch lex.data[(lex.p)] {
		case 84:
			goto st39
		case 116:
			goto st39
		}
		goto tr38
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch lex.data[(lex.p)] {
		case 78:
			goto st45
		case 110:
			goto st45
		}
		goto tr38
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 84:
			goto st46
		case 116:
			goto st46
		}
		goto tr38
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch lex.data[(lex.p)] {
		case 9:
			goto st47
		case 32:
			goto st47
		case 41:
			goto tr79
		case 69:
			goto st48
		case 101:
			goto st48
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st47
		}
		goto tr38
	st47:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch lex.data[(lex.p)] {
		case 9:
			goto st47
		case 32:
			goto st47
		case 41:
			goto tr79
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st47
		}
		goto tr38
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch lex.data[(lex.p)] {
		case 71:
			goto st49
		case 103:
			goto st49
		}
		goto tr38
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
		goto tr38
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch lex.data[(lex.p)] {
		case 82:
			goto st47
		case 114:
			goto st47
		}
		goto tr38
	st51:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch lex.data[(lex.p)] {
		case 66:
			goto st52
		case 98:
			goto st52
		}
		goto tr38
	st52:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch lex.data[(lex.p)] {
		case 74:
			goto st53
		case 106:
			goto st53
		}
		goto tr38
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
		goto tr38
	st54:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch lex.data[(lex.p)] {
		case 67:
			goto st55
		case 99:
			goto st55
		}
		goto tr38
	st55:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch lex.data[(lex.p)] {
		case 84:
			goto st56
		case 116:
			goto st56
		}
		goto tr38
	st56:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch lex.data[(lex.p)] {
		case 9:
			goto st56
		case 32:
			goto st56
		case 41:
			goto tr88
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st56
		}
		goto tr38
	st57:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch lex.data[(lex.p)] {
		case 69:
			goto st58
		case 101:
			goto st58
		}
		goto tr38
	st58:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch lex.data[(lex.p)] {
		case 65:
			goto st59
		case 97:
			goto st59
		}
		goto tr38
	st59:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch lex.data[(lex.p)] {
		case 76:
			goto st39
		case 108:
			goto st39
		}
		goto tr38
	st60:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch lex.data[(lex.p)] {
		case 84:
			goto st61
		case 116:
			goto st61
		}
		goto tr38
	st61:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch lex.data[(lex.p)] {
		case 82:
			goto st62
		case 114:
			goto st62
		}
		goto tr38
	st62:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch lex.data[(lex.p)] {
		case 73:
			goto st63
		case 105:
			goto st63
		}
		goto tr38
	st63:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch lex.data[(lex.p)] {
		case 78:
			goto st64
		case 110:
			goto st64
		}
		goto tr38
	st64:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch lex.data[(lex.p)] {
		case 71:
			goto st27
		case 103:
			goto st27
		}
		goto tr38
	st65:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch lex.data[(lex.p)] {
		case 78:
			goto st66
		case 110:
			goto st66
		}
		goto tr38
	st66:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch lex.data[(lex.p)] {
		case 83:
			goto st67
		case 115:
			goto st67
		}
		goto tr38
	st67:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch lex.data[(lex.p)] {
		case 69:
			goto st68
		case 101:
			goto st68
		}
		goto tr38
	st68:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch lex.data[(lex.p)] {
		case 84:
			goto st69
		case 116:
			goto st69
		}
		goto tr38
	st69:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch lex.data[(lex.p)] {
		case 9:
			goto st69
		case 32:
			goto st69
		case 41:
			goto tr99
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st69
		}
		goto tr38
	st133:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof133
		}
	st_case_133:
		switch lex.data[(lex.p)] {
		case 42:
			goto st134
		case 61:
			goto tr265
		}
		goto tr249
	st134:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof134
		}
	st_case_134:
		if lex.data[(lex.p)] == 61 {
			goto tr267
		}
		goto tr266
	st135:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof135
		}
	st_case_135:
		switch lex.data[(lex.p)] {
		case 43:
			goto tr268
		case 61:
			goto tr269
		}
		goto tr249
	st136:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch lex.data[(lex.p)] {
		case 45:
			goto tr270
		case 61:
			goto tr271
		case 62:
			goto tr272
		}
		goto tr249
	tr203:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st137
	st137:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof137
		}
	st_case_137:
//line scanner/scanner.go:4847
		switch lex.data[(lex.p)] {
		case 46:
			goto st70
		case 61:
			goto tr275
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr274
		}
		goto tr249
	st70:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		if lex.data[(lex.p)] == 46 {
			goto tr100
		}
		goto tr38
	tr274:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:135
		lex.act = 8
		goto st138
	st138:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof138
		}
	st_case_138:
//line scanner/scanner.go:4879
		switch lex.data[(lex.p)] {
		case 69:
			goto st71
		case 101:
			goto st71
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr274
		}
		goto tr276
	st71:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch lex.data[(lex.p)] {
		case 43:
			goto st72
		case 45:
			goto st72
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st139
		}
		goto tr8
	st72:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st139
		}
		goto tr8
	st139:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof139
		}
	st_case_139:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st139
		}
		goto tr276
	tr204:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st140
	st140:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof140
		}
	st_case_140:
//line scanner/scanner.go:4933
		switch lex.data[(lex.p)] {
		case 42:
			goto st73
		case 47:
			goto st125
		case 61:
			goto tr278
		}
		goto tr249
	tr106:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st73
	st73:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof73
		}
	st_case_73:
//line scanner/scanner.go:4952
		switch lex.data[(lex.p)] {
		case 10:
			goto st74
		case 42:
			goto st75
		}
		goto st73
	tr107:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st74
	st74:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof74
		}
	st_case_74:
//line scanner/scanner.go:4969
		switch lex.data[(lex.p)] {
		case 10:
			goto tr107
		case 42:
			goto tr108
		}
		goto tr106
	tr108:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st75
	st75:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof75
		}
	st_case_75:
//line scanner/scanner.go:4986
		switch lex.data[(lex.p)] {
		case 10:
			goto st74
		case 42:
			goto st75
		case 47:
			goto tr109
		}
		goto st73
	tr205:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:149
		lex.act = 10
		goto st141
	st141:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof141
		}
	st_case_141:
//line scanner/scanner.go:5008
		switch lex.data[(lex.p)] {
		case 46:
			goto tr274
		case 69:
			goto st71
		case 98:
			goto st76
		case 101:
			goto st71
		case 120:
			goto st77
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr206
		}
		goto tr279
	tr206:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:149
		lex.act = 10
		goto st142
	st142:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof142
		}
	st_case_142:
//line scanner/scanner.go:5037
		switch lex.data[(lex.p)] {
		case 46:
			goto tr274
		case 69:
			goto st71
		case 101:
			goto st71
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr206
		}
		goto tr279
	st76:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st143
		}
		goto tr110
	st143:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof143
		}
	st_case_143:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st143
		}
		goto tr282
	st77:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st144
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st144
			}
		default:
			goto st144
		}
		goto tr110
	st144:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof144
		}
	st_case_144:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st144
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st144
			}
		default:
			goto st144
		}
		goto tr283
	st145:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof145
		}
	st_case_145:
		switch lex.data[(lex.p)] {
		case 46:
			goto tr274
		case 69:
			goto tr284
		case 95:
			goto tr221
		case 101:
			goto tr284
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st145
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr279
	tr221:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:312
		lex.act = 131
		goto st146
	tr315:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:172
		lex.act = 14
		goto st146
	tr321:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:170
		lex.act = 12
		goto st146
	tr322:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:242
		lex.act = 84
		goto st146
	tr325:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:171
		lex.act = 13
		goto st146
	tr330:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:173
		lex.act = 15
		goto st146
	tr342:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:174
		lex.act = 16
		goto st146
	tr343:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:175
		lex.act = 17
		goto st146
	tr345:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:176
		lex.act = 18
		goto st146
	tr352:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:201
		lex.act = 43
		goto st146
	tr356:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:177
		lex.act = 19
		goto st146
	tr358:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:178
		lex.act = 20
		goto st146
	tr362:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:179
		lex.act = 21
		goto st146
	tr366:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:180
		lex.act = 22
		goto st146
	tr369:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:183
		lex.act = 25
		goto st146
	tr375:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:181
		lex.act = 23
		goto st146
	tr379:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:182
		lex.act = 24
		goto st146
	tr380:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:195
		lex.act = 37
		goto st146
	tr388:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:184
		lex.act = 26
		goto st146
	tr393:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:186
		lex.act = 28
		goto st146
	tr396:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:187
		lex.act = 29
		goto st146
	tr408:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:188
		lex.act = 30
		goto st146
	tr415:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:190
		lex.act = 32
		goto st146
	tr416:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:191
		lex.act = 33
		goto st146
	tr421:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:192
		lex.act = 34
		goto st146
	tr425:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:193
		lex.act = 35
		goto st146
	tr427:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:194
		lex.act = 36
		goto st146
	tr433:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:196
		lex.act = 38
		goto st146
	tr441:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:198
		lex.act = 40
		goto st146
	tr447:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:200
		lex.act = 42
		goto st146
	tr453:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:202
		lex.act = 44
		goto st146
	tr455:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:203
		lex.act = 45
		goto st146
	tr456:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:204
		lex.act = 46
		goto st146
	tr467:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:206
		lex.act = 48
		goto st146
	tr480:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:229
		lex.act = 71
		goto st146
	tr488:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:207
		lex.act = 49
		goto st146
	tr492:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:208
		lex.act = 50
		goto st146
	tr498:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:209
		lex.act = 51
		goto st146
	tr501:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:205
		lex.act = 47
		goto st146
	tr504:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:210
		lex.act = 52
		goto st146
	tr513:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:211
		lex.act = 53
		goto st146
	tr514:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:241
		lex.act = 83
		goto st146
	tr515:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:243
		lex.act = 85
		goto st146
	tr522:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:214
		lex.act = 56
		goto st146
	tr525:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:212
		lex.act = 54
		goto st146
	tr531:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:215
		lex.act = 57
		goto st146
	tr535:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:213
		lex.act = 55
		goto st146
	tr548:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:231
		lex.act = 73
		goto st146
	tr551:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:216
		lex.act = 58
		goto st146
	tr557:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:217
		lex.act = 59
		goto st146
	tr561:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:218
		lex.act = 60
		goto st146
	tr566:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:219
		lex.act = 61
		goto st146
	tr568:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:221
		lex.act = 63
		goto st146
	tr570:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:220
		lex.act = 62
		goto st146
	tr575:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:222
		lex.act = 64
		goto st146
	tr576:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:223
		lex.act = 65
		goto st146
	tr578:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:224
		lex.act = 66
		goto st146
	tr582:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:225
		lex.act = 67
		goto st146
	tr584:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:244
		lex.act = 86
		goto st146
	tr593:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:226
		lex.act = 68
		goto st146
	tr609:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:232
		lex.act = 74
		goto st146
	tr613:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:233
		lex.act = 75
		goto st146
	tr619:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:234
		lex.act = 76
		goto st146
	tr627:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:235
		lex.act = 77
		goto st146
	tr639:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:240
		lex.act = 82
		goto st146
	tr644:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:236
		lex.act = 78
		goto st146
	tr651:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:238
		lex.act = 80
		goto st146
	tr661:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:237
		lex.act = 79
		goto st146
	tr667:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:239
		lex.act = 81
		goto st146
	st146:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof146
		}
	st_case_146:
//line scanner/scanner.go:5620
		if lex.data[(lex.p)] == 95 {
			goto tr221
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr8
	tr284:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:312
		lex.act = 131
		goto st147
	st147:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof147
		}
	st_case_147:
//line scanner/scanner.go:5649
		switch lex.data[(lex.p)] {
		case 43:
			goto st72
		case 45:
			goto st72
		case 95:
			goto tr221
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st148
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st148:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof148
		}
	st_case_148:
		if lex.data[(lex.p)] == 95 {
			goto tr221
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st148
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr276
	st149:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof149
		}
	st_case_149:
		if lex.data[(lex.p)] == 58 {
			goto tr287
		}
		goto tr249
	tr209:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st150
	st150:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof150
		}
	st_case_150:
//line scanner/scanner.go:5711
		switch lex.data[(lex.p)] {
		case 10:
			goto st79
		case 13:
			goto st80
		case 32:
			goto st78
		case 63:
			goto st81
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st78
		}
		goto tr249
	tr117:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st78
	st78:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof78
		}
	st_case_78:
//line scanner/scanner.go:5735
		switch lex.data[(lex.p)] {
		case 10:
			goto st79
		case 13:
			goto st80
		case 32:
			goto st78
		case 63:
			goto st81
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st78
		}
		goto tr38
	tr118:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st79
	st79:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof79
		}
	st_case_79:
//line scanner/scanner.go:5759
		switch lex.data[(lex.p)] {
		case 10:
			goto tr118
		case 13:
			goto tr119
		case 32:
			goto tr117
		case 63:
			goto tr120
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr117
		}
		goto tr38
	tr119:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st80
	st80:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof80
		}
	st_case_80:
//line scanner/scanner.go:5783
		if lex.data[(lex.p)] == 10 {
			goto st79
		}
		goto tr38
	tr120:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st81
	st81:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof81
		}
	st_case_81:
//line scanner/scanner.go:5797
		if lex.data[(lex.p)] == 62 {
			goto tr121
		}
		goto tr38
	tr121:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st151
	st151:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof151
		}
	st_case_151:
//line scanner/scanner.go:5812
		switch lex.data[(lex.p)] {
		case 10:
			goto st152
		case 13:
			goto st82
		}
		goto tr288
	st152:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof152
		}
	st_case_152:
		goto tr290
	st82:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		if lex.data[(lex.p)] == 10 {
			goto st152
		}
		goto tr122
	st153:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof153
		}
	st_case_153:
		switch lex.data[(lex.p)] {
		case 60:
			goto tr291
		case 61:
			goto st156
		case 62:
			goto tr293
		}
		goto tr249
	tr291:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:273
		lex.act = 115
		goto st154
	st154:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof154
		}
	st_case_154:
//line scanner/scanner.go:5861
		switch lex.data[(lex.p)] {
		case 60:
			goto st83
		case 61:
			goto tr295
		}
		goto tr294
	st83:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch lex.data[(lex.p)] {
		case 9:
			goto st83
		case 32:
			goto st83
		case 34:
			goto st84
		case 39:
			goto st88
		case 55:
			goto tr127
		case 95:
			goto tr127
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr127
			}
		case lex.data[(lex.p)] >= 65:
			goto tr127
		}
		goto tr8
	st84:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch lex.data[(lex.p)] {
		case 55:
			goto tr128
		case 95:
			goto tr128
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr128
			}
		case lex.data[(lex.p)] >= 65:
			goto tr128
		}
		goto tr8
	tr128:
//line scanner/scanner.rl:41
		lblStart = lex.p
		goto st85
	st85:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof85
		}
	st_case_85:
//line scanner/scanner.go:5926
		switch lex.data[(lex.p)] {
		case 34:
			goto tr129
		case 95:
			goto st85
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st85
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st85
			}
		default:
			goto st85
		}
		goto tr8
	tr129:
//line scanner/scanner.rl:42
		lblEnd = lex.p
		goto st86
	st86:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof86
		}
	st_case_86:
//line scanner/scanner.go:5955
		switch lex.data[(lex.p)] {
		case 10:
			goto st155
		case 13:
			goto st87
		}
		goto tr8
	tr135:
//line scanner/scanner.rl:42
		lblEnd = lex.p
		goto st155
	st155:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof155
		}
	st_case_155:
//line scanner/scanner.go:5972
		goto tr296
	tr136:
//line scanner/scanner.rl:42
		lblEnd = lex.p
		goto st87
	st87:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof87
		}
	st_case_87:
//line scanner/scanner.go:5983
		if lex.data[(lex.p)] == 10 {
			goto st155
		}
		goto tr8
	st88:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch lex.data[(lex.p)] {
		case 55:
			goto tr133
		case 95:
			goto tr133
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr133
			}
		case lex.data[(lex.p)] >= 65:
			goto tr133
		}
		goto tr8
	tr133:
//line scanner/scanner.rl:41
		lblStart = lex.p
		goto st89
	st89:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof89
		}
	st_case_89:
//line scanner/scanner.go:6017
		switch lex.data[(lex.p)] {
		case 39:
			goto tr129
		case 95:
			goto st89
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st89
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st89
			}
		default:
			goto st89
		}
		goto tr8
	tr127:
//line scanner/scanner.rl:41
		lblStart = lex.p
		goto st90
	st90:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof90
		}
	st_case_90:
//line scanner/scanner.go:6046
		switch lex.data[(lex.p)] {
		case 10:
			goto tr135
		case 13:
			goto tr136
		case 95:
			goto st90
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st90
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st90
			}
		default:
			goto st90
		}
		goto tr8
	st156:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof156
		}
	st_case_156:
		if lex.data[(lex.p)] == 62 {
			goto tr298
		}
		goto tr297
	st157:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof157
		}
	st_case_157:
		switch lex.data[(lex.p)] {
		case 61:
			goto st158
		case 62:
			goto tr300
		}
		goto tr249
	st158:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof158
		}
	st_case_158:
		if lex.data[(lex.p)] == 61 {
			goto tr302
		}
		goto tr301
	st159:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof159
		}
	st_case_159:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr303
		case 62:
			goto st160
		}
		goto tr249
	st160:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof160
		}
	st_case_160:
		if lex.data[(lex.p)] == 61 {
			goto tr306
		}
		goto tr305
	st161:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch lex.data[(lex.p)] {
		case 62:
			goto tr307
		case 63:
			goto tr308
		}
		goto tr249
	tr307:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st162
	st162:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof162
		}
	st_case_162:
//line scanner/scanner.go:6141
		switch lex.data[(lex.p)] {
		case 10:
			goto st163
		case 13:
			goto st91
		}
		goto tr309
	st163:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof163
		}
	st_case_163:
		goto tr311
	st91:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		if lex.data[(lex.p)] == 10 {
			goto st163
		}
		goto tr138
	st164:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		switch lex.data[(lex.p)] {
		case 66:
			goto st165
		case 78:
			goto st171
		case 82:
			goto st172
		case 83:
			goto tr315
		case 95:
			goto tr221
		case 98:
			goto st165
		case 110:
			goto st171
		case 114:
			goto st172
		case 115:
			goto tr315
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st165:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof165
		}
	st_case_165:
		switch lex.data[(lex.p)] {
		case 83:
			goto st166
		case 95:
			goto tr221
		case 115:
			goto st166
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st166:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof166
		}
	st_case_166:
		switch lex.data[(lex.p)] {
		case 84:
			goto st167
		case 95:
			goto tr221
		case 116:
			goto st167
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st167:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof167
		}
	st_case_167:
		switch lex.data[(lex.p)] {
		case 82:
			goto st168
		case 95:
			goto tr221
		case 114:
			goto st168
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st168:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof168
		}
	st_case_168:
		switch lex.data[(lex.p)] {
		case 65:
			goto st169
		case 95:
			goto tr221
		case 97:
			goto st169
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st169:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch lex.data[(lex.p)] {
		case 67:
			goto st170
		case 95:
			goto tr221
		case 99:
			goto st170
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st170:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr321
		case 95:
			goto tr221
		case 116:
			goto tr321
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st171:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr322
		case 95:
			goto tr221
		case 100:
			goto tr322
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st172:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof172
		}
	st_case_172:
		switch lex.data[(lex.p)] {
		case 82:
			goto st173
		case 95:
			goto tr221
		case 114:
			goto st173
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st173:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof173
		}
	st_case_173:
		switch lex.data[(lex.p)] {
		case 65:
			goto st174
		case 95:
			goto tr221
		case 97:
			goto st174
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st174:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof174
		}
	st_case_174:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr325
		case 95:
			goto tr221
		case 121:
			goto tr325
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	tr215:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:312
		lex.act = 131
		goto st175
	st175:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof175
		}
	st_case_175:
//line scanner/scanner.go:6474
		switch lex.data[(lex.p)] {
		case 34:
			goto st6
		case 60:
			goto st92
		case 82:
			goto st176
		case 95:
			goto tr221
		case 114:
			goto st176
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st92:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		if lex.data[(lex.p)] == 60 {
			goto st93
		}
		goto tr140
	st93:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		if lex.data[(lex.p)] == 60 {
			goto st83
		}
		goto tr140
	st176:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof176
		}
	st_case_176:
		switch lex.data[(lex.p)] {
		case 69:
			goto st177
		case 95:
			goto tr221
		case 101:
			goto st177
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st177:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof177
		}
	st_case_177:
		switch lex.data[(lex.p)] {
		case 65:
			goto st178
		case 95:
			goto tr221
		case 97:
			goto st178
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st178:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof178
		}
	st_case_178:
		switch lex.data[(lex.p)] {
		case 75:
			goto tr330
		case 95:
			goto tr221
		case 107:
			goto tr330
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st179:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof179
		}
	st_case_179:
		switch lex.data[(lex.p)] {
		case 65:
			goto st180
		case 70:
			goto st189
		case 76:
			goto st196
		case 79:
			goto st201
		case 95:
			goto tr221
		case 97:
			goto st180
		case 102:
			goto st189
		case 108:
			goto st196
		case 111:
			goto st201
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st180:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof180
		}
	st_case_180:
		switch lex.data[(lex.p)] {
		case 76:
			goto st181
		case 83:
			goto st186
		case 84:
			goto st187
		case 95:
			goto tr221
		case 108:
			goto st181
		case 115:
			goto st186
		case 116:
			goto st187
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st181:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof181
		}
	st_case_181:
		switch lex.data[(lex.p)] {
		case 76:
			goto st182
		case 95:
			goto tr221
		case 108:
			goto st182
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st182:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof182
		}
	st_case_182:
		switch lex.data[(lex.p)] {
		case 65:
			goto st183
		case 95:
			goto tr221
		case 97:
			goto st183
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st183:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof183
		}
	st_case_183:
		switch lex.data[(lex.p)] {
		case 66:
			goto st184
		case 95:
			goto tr221
		case 98:
			goto st184
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st184:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof184
		}
	st_case_184:
		switch lex.data[(lex.p)] {
		case 76:
			goto st185
		case 95:
			goto tr221
		case 108:
			goto st185
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st185:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr342
		case 95:
			goto tr221
		case 101:
			goto tr342
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st186:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof186
		}
	st_case_186:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr343
		case 95:
			goto tr221
		case 101:
			goto tr343
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st187:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof187
		}
	st_case_187:
		switch lex.data[(lex.p)] {
		case 67:
			goto st188
		case 95:
			goto tr221
		case 99:
			goto st188
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st188:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof188
		}
	st_case_188:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr345
		case 95:
			goto tr221
		case 104:
			goto tr345
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st189:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof189
		}
	st_case_189:
		switch lex.data[(lex.p)] {
		case 85:
			goto st190
		case 95:
			goto tr221
		case 117:
			goto st190
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st190:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof190
		}
	st_case_190:
		switch lex.data[(lex.p)] {
		case 78:
			goto st191
		case 95:
			goto tr221
		case 110:
			goto st191
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st191:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof191
		}
	st_case_191:
		switch lex.data[(lex.p)] {
		case 67:
			goto st192
		case 95:
			goto tr221
		case 99:
			goto st192
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st192:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof192
		}
	st_case_192:
		switch lex.data[(lex.p)] {
		case 84:
			goto st193
		case 95:
			goto tr221
		case 116:
			goto st193
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st193:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof193
		}
	st_case_193:
		switch lex.data[(lex.p)] {
		case 73:
			goto st194
		case 95:
			goto tr221
		case 105:
			goto st194
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st194:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof194
		}
	st_case_194:
		switch lex.data[(lex.p)] {
		case 79:
			goto st195
		case 95:
			goto tr221
		case 111:
			goto st195
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st195:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof195
		}
	st_case_195:
		switch lex.data[(lex.p)] {
		case 78:
			goto tr352
		case 95:
			goto tr221
		case 110:
			goto tr352
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st196:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof196
		}
	st_case_196:
		switch lex.data[(lex.p)] {
		case 65:
			goto st197
		case 79:
			goto st199
		case 95:
			goto tr221
		case 97:
			goto st197
		case 111:
			goto st199
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st197:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof197
		}
	st_case_197:
		switch lex.data[(lex.p)] {
		case 83:
			goto st198
		case 95:
			goto tr221
		case 115:
			goto st198
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st198:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof198
		}
	st_case_198:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr356
		case 95:
			goto tr221
		case 115:
			goto tr356
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st199:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof199
		}
	st_case_199:
		switch lex.data[(lex.p)] {
		case 78:
			goto st200
		case 95:
			goto tr221
		case 110:
			goto st200
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st200:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof200
		}
	st_case_200:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr358
		case 95:
			goto tr221
		case 101:
			goto tr358
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st201:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof201
		}
	st_case_201:
		switch lex.data[(lex.p)] {
		case 78:
			goto st202
		case 95:
			goto tr221
		case 110:
			goto st202
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st202:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof202
		}
	st_case_202:
		switch lex.data[(lex.p)] {
		case 83:
			goto st203
		case 84:
			goto st204
		case 95:
			goto tr221
		case 115:
			goto st203
		case 116:
			goto st204
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st203:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof203
		}
	st_case_203:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr362
		case 95:
			goto tr221
		case 116:
			goto tr362
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st204:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof204
		}
	st_case_204:
		switch lex.data[(lex.p)] {
		case 73:
			goto st205
		case 95:
			goto tr221
		case 105:
			goto st205
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st205:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof205
		}
	st_case_205:
		switch lex.data[(lex.p)] {
		case 78:
			goto st206
		case 95:
			goto tr221
		case 110:
			goto st206
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st206:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof206
		}
	st_case_206:
		switch lex.data[(lex.p)] {
		case 85:
			goto st207
		case 95:
			goto tr221
		case 117:
			goto st207
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st207:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof207
		}
	st_case_207:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr366
		case 95:
			goto tr221
		case 101:
			goto tr366
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st208:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof208
		}
	st_case_208:
		switch lex.data[(lex.p)] {
		case 69:
			goto st209
		case 73:
			goto st218
		case 79:
			goto tr369
		case 95:
			goto tr221
		case 101:
			goto st209
		case 105:
			goto st218
		case 111:
			goto tr369
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st209:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof209
		}
	st_case_209:
		switch lex.data[(lex.p)] {
		case 67:
			goto st210
		case 70:
			goto st214
		case 95:
			goto tr221
		case 99:
			goto st210
		case 102:
			goto st214
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st210:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof210
		}
	st_case_210:
		switch lex.data[(lex.p)] {
		case 76:
			goto st211
		case 95:
			goto tr221
		case 108:
			goto st211
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st211:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof211
		}
	st_case_211:
		switch lex.data[(lex.p)] {
		case 65:
			goto st212
		case 95:
			goto tr221
		case 97:
			goto st212
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st212:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof212
		}
	st_case_212:
		switch lex.data[(lex.p)] {
		case 82:
			goto st213
		case 95:
			goto tr221
		case 114:
			goto st213
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st213:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof213
		}
	st_case_213:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr375
		case 95:
			goto tr221
		case 101:
			goto tr375
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st214:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof214
		}
	st_case_214:
		switch lex.data[(lex.p)] {
		case 65:
			goto st215
		case 95:
			goto tr221
		case 97:
			goto st215
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st215:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof215
		}
	st_case_215:
		switch lex.data[(lex.p)] {
		case 85:
			goto st216
		case 95:
			goto tr221
		case 117:
			goto st216
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st216:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof216
		}
	st_case_216:
		switch lex.data[(lex.p)] {
		case 76:
			goto st217
		case 95:
			goto tr221
		case 108:
			goto st217
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st217:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof217
		}
	st_case_217:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr379
		case 95:
			goto tr221
		case 116:
			goto tr379
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st218:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof218
		}
	st_case_218:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr380
		case 95:
			goto tr221
		case 101:
			goto tr380
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st219:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof219
		}
	st_case_219:
		switch lex.data[(lex.p)] {
		case 67:
			goto st220
		case 76:
			goto st222
		case 77:
			goto st226
		case 78:
			goto st229
		case 86:
			goto st253
		case 88:
			goto st255
		case 95:
			goto tr221
		case 99:
			goto st220
		case 108:
			goto st222
		case 109:
			goto st226
		case 110:
			goto st229
		case 118:
			goto st253
		case 120:
			goto st255
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st220:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof220
		}
	st_case_220:
		switch lex.data[(lex.p)] {
		case 72:
			goto st221
		case 95:
			goto tr221
		case 104:
			goto st221
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st221:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof221
		}
	st_case_221:
		switch lex.data[(lex.p)] {
		case 79:
			goto tr388
		case 95:
			goto tr221
		case 111:
			goto tr388
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st222:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof222
		}
	st_case_222:
		switch lex.data[(lex.p)] {
		case 83:
			goto st223
		case 95:
			goto tr221
		case 115:
			goto st223
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st223:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof223
		}
	st_case_223:
		switch lex.data[(lex.p)] {
		case 69:
			goto st224
		case 95:
			goto tr221
		case 101:
			goto st224
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st224:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof224
		}
	st_case_224:
		switch lex.data[(lex.p)] {
		case 73:
			goto st225
		case 95:
			goto tr221
		case 105:
			goto st225
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr391
	st225:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof225
		}
	st_case_225:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr393
		case 95:
			goto tr221
		case 102:
			goto tr393
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st226:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof226
		}
	st_case_226:
		switch lex.data[(lex.p)] {
		case 80:
			goto st227
		case 95:
			goto tr221
		case 112:
			goto st227
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st227:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof227
		}
	st_case_227:
		switch lex.data[(lex.p)] {
		case 84:
			goto st228
		case 95:
			goto tr221
		case 116:
			goto st228
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st228:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof228
		}
	st_case_228:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr396
		case 95:
			goto tr221
		case 121:
			goto tr396
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st229:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof229
		}
	st_case_229:
		switch lex.data[(lex.p)] {
		case 68:
			goto st230
		case 95:
			goto tr221
		case 100:
			goto st230
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st230:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof230
		}
	st_case_230:
		switch lex.data[(lex.p)] {
		case 68:
			goto st231
		case 70:
			goto st237
		case 73:
			goto st243
		case 83:
			goto st244
		case 87:
			goto st249
		case 95:
			goto tr221
		case 100:
			goto st231
		case 102:
			goto st237
		case 105:
			goto st243
		case 115:
			goto st244
		case 119:
			goto st249
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st231:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof231
		}
	st_case_231:
		switch lex.data[(lex.p)] {
		case 69:
			goto st232
		case 95:
			goto tr221
		case 101:
			goto st232
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st232:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof232
		}
	st_case_232:
		switch lex.data[(lex.p)] {
		case 67:
			goto st233
		case 95:
			goto tr221
		case 99:
			goto st233
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st233:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof233
		}
	st_case_233:
		switch lex.data[(lex.p)] {
		case 76:
			goto st234
		case 95:
			goto tr221
		case 108:
			goto st234
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st234:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof234
		}
	st_case_234:
		switch lex.data[(lex.p)] {
		case 65:
			goto st235
		case 95:
			goto tr221
		case 97:
			goto st235
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st235:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof235
		}
	st_case_235:
		switch lex.data[(lex.p)] {
		case 82:
			goto st236
		case 95:
			goto tr221
		case 114:
			goto st236
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st236:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof236
		}
	st_case_236:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr408
		case 95:
			goto tr221
		case 101:
			goto tr408
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st237:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof237
		}
	st_case_237:
		switch lex.data[(lex.p)] {
		case 79:
			goto st238
		case 95:
			goto tr221
		case 111:
			goto st238
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st238:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof238
		}
	st_case_238:
		switch lex.data[(lex.p)] {
		case 82:
			goto st239
		case 95:
			goto tr221
		case 114:
			goto st239
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st239:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof239
		}
	st_case_239:
		switch lex.data[(lex.p)] {
		case 69:
			goto st240
		case 95:
			goto tr221
		case 101:
			goto st240
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr411
	st240:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof240
		}
	st_case_240:
		switch lex.data[(lex.p)] {
		case 65:
			goto st241
		case 95:
			goto tr221
		case 97:
			goto st241
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st241:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof241
		}
	st_case_241:
		switch lex.data[(lex.p)] {
		case 67:
			goto st242
		case 95:
			goto tr221
		case 99:
			goto st242
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st242:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof242
		}
	st_case_242:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr415
		case 95:
			goto tr221
		case 104:
			goto tr415
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st243:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof243
		}
	st_case_243:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr416
		case 95:
			goto tr221
		case 102:
			goto tr416
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st244:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof244
		}
	st_case_244:
		switch lex.data[(lex.p)] {
		case 87:
			goto st245
		case 95:
			goto tr221
		case 119:
			goto st245
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st245:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof245
		}
	st_case_245:
		switch lex.data[(lex.p)] {
		case 73:
			goto st246
		case 95:
			goto tr221
		case 105:
			goto st246
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st246:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof246
		}
	st_case_246:
		switch lex.data[(lex.p)] {
		case 84:
			goto st247
		case 95:
			goto tr221
		case 116:
			goto st247
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st247:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof247
		}
	st_case_247:
		switch lex.data[(lex.p)] {
		case 67:
			goto st248
		case 95:
			goto tr221
		case 99:
			goto st248
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st248:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof248
		}
	st_case_248:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr421
		case 95:
			goto tr221
		case 104:
			goto tr421
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st249:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof249
		}
	st_case_249:
		switch lex.data[(lex.p)] {
		case 72:
			goto st250
		case 95:
			goto tr221
		case 104:
			goto st250
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st250:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof250
		}
	st_case_250:
		switch lex.data[(lex.p)] {
		case 73:
			goto st251
		case 95:
			goto tr221
		case 105:
			goto st251
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st251:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof251
		}
	st_case_251:
		switch lex.data[(lex.p)] {
		case 76:
			goto st252
		case 95:
			goto tr221
		case 108:
			goto st252
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st252:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof252
		}
	st_case_252:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr425
		case 95:
			goto tr221
		case 101:
			goto tr425
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st253:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof253
		}
	st_case_253:
		switch lex.data[(lex.p)] {
		case 65:
			goto st254
		case 95:
			goto tr221
		case 97:
			goto st254
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st254:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof254
		}
	st_case_254:
		switch lex.data[(lex.p)] {
		case 76:
			goto tr427
		case 95:
			goto tr221
		case 108:
			goto tr427
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st255:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof255
		}
	st_case_255:
		switch lex.data[(lex.p)] {
		case 73:
			goto st256
		case 84:
			goto st257
		case 95:
			goto tr221
		case 105:
			goto st256
		case 116:
			goto st257
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st256:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof256
		}
	st_case_256:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr380
		case 95:
			goto tr221
		case 116:
			goto tr380
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st257:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof257
		}
	st_case_257:
		switch lex.data[(lex.p)] {
		case 69:
			goto st258
		case 95:
			goto tr221
		case 101:
			goto st258
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st258:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof258
		}
	st_case_258:
		switch lex.data[(lex.p)] {
		case 78:
			goto st259
		case 95:
			goto tr221
		case 110:
			goto st259
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st259:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof259
		}
	st_case_259:
		switch lex.data[(lex.p)] {
		case 68:
			goto st260
		case 95:
			goto tr221
		case 100:
			goto st260
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st260:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof260
		}
	st_case_260:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr433
		case 95:
			goto tr221
		case 115:
			goto tr433
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st261:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof261
		}
	st_case_261:
		switch lex.data[(lex.p)] {
		case 73:
			goto st262
		case 79:
			goto st267
		case 85:
			goto st190
		case 95:
			goto tr221
		case 105:
			goto st262
		case 111:
			goto st267
		case 117:
			goto st190
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st262:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof262
		}
	st_case_262:
		switch lex.data[(lex.p)] {
		case 78:
			goto st263
		case 95:
			goto tr221
		case 110:
			goto st263
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st263:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof263
		}
	st_case_263:
		switch lex.data[(lex.p)] {
		case 65:
			goto st264
		case 95:
			goto tr221
		case 97:
			goto st264
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st264:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof264
		}
	st_case_264:
		switch lex.data[(lex.p)] {
		case 76:
			goto st265
		case 95:
			goto tr221
		case 108:
			goto st265
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st265:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof265
		}
	st_case_265:
		switch lex.data[(lex.p)] {
		case 76:
			goto st266
		case 95:
			goto tr221
		case 108:
			goto st266
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr439
	st266:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof266
		}
	st_case_266:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr441
		case 95:
			goto tr221
		case 121:
			goto tr441
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st267:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof267
		}
	st_case_267:
		switch lex.data[(lex.p)] {
		case 82:
			goto st268
		case 95:
			goto tr221
		case 114:
			goto st268
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st268:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof268
		}
	st_case_268:
		switch lex.data[(lex.p)] {
		case 69:
			goto st269
		case 95:
			goto tr221
		case 101:
			goto st269
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr443
	st269:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof269
		}
	st_case_269:
		switch lex.data[(lex.p)] {
		case 65:
			goto st270
		case 95:
			goto tr221
		case 97:
			goto st270
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st270:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof270
		}
	st_case_270:
		switch lex.data[(lex.p)] {
		case 67:
			goto st271
		case 95:
			goto tr221
		case 99:
			goto st271
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st271:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof271
		}
	st_case_271:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr447
		case 95:
			goto tr221
		case 104:
			goto tr447
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st272:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof272
		}
	st_case_272:
		switch lex.data[(lex.p)] {
		case 76:
			goto st273
		case 79:
			goto st277
		case 95:
			goto tr221
		case 108:
			goto st273
		case 111:
			goto st277
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st273:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof273
		}
	st_case_273:
		switch lex.data[(lex.p)] {
		case 79:
			goto st274
		case 95:
			goto tr221
		case 111:
			goto st274
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st274:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof274
		}
	st_case_274:
		switch lex.data[(lex.p)] {
		case 66:
			goto st275
		case 95:
			goto tr221
		case 98:
			goto st275
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st275:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof275
		}
	st_case_275:
		switch lex.data[(lex.p)] {
		case 65:
			goto st276
		case 95:
			goto tr221
		case 97:
			goto st276
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st276:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof276
		}
	st_case_276:
		switch lex.data[(lex.p)] {
		case 76:
			goto tr453
		case 95:
			goto tr221
		case 108:
			goto tr453
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st277:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof277
		}
	st_case_277:
		switch lex.data[(lex.p)] {
		case 84:
			goto st278
		case 95:
			goto tr221
		case 116:
			goto st278
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st278:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof278
		}
	st_case_278:
		switch lex.data[(lex.p)] {
		case 79:
			goto tr455
		case 95:
			goto tr221
		case 111:
			goto tr455
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st279:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof279
		}
	st_case_279:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr456
		case 77:
			goto st280
		case 78:
			goto st288
		case 83:
			goto st315
		case 95:
			goto tr221
		case 102:
			goto tr456
		case 109:
			goto st280
		case 110:
			goto st288
		case 115:
			goto st315
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st280:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof280
		}
	st_case_280:
		switch lex.data[(lex.p)] {
		case 80:
			goto st281
		case 95:
			goto tr221
		case 112:
			goto st281
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st281:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof281
		}
	st_case_281:
		switch lex.data[(lex.p)] {
		case 76:
			goto st282
		case 95:
			goto tr221
		case 108:
			goto st282
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st282:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof282
		}
	st_case_282:
		switch lex.data[(lex.p)] {
		case 69:
			goto st283
		case 95:
			goto tr221
		case 101:
			goto st283
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st283:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof283
		}
	st_case_283:
		switch lex.data[(lex.p)] {
		case 77:
			goto st284
		case 95:
			goto tr221
		case 109:
			goto st284
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st284:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof284
		}
	st_case_284:
		switch lex.data[(lex.p)] {
		case 69:
			goto st285
		case 95:
			goto tr221
		case 101:
			goto st285
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st285:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof285
		}
	st_case_285:
		switch lex.data[(lex.p)] {
		case 78:
			goto st286
		case 95:
			goto tr221
		case 110:
			goto st286
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st286:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof286
		}
	st_case_286:
		switch lex.data[(lex.p)] {
		case 84:
			goto st287
		case 95:
			goto tr221
		case 116:
			goto st287
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st287:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof287
		}
	st_case_287:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr467
		case 95:
			goto tr221
		case 115:
			goto tr467
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st288:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof288
		}
	st_case_288:
		switch lex.data[(lex.p)] {
		case 67:
			goto st289
		case 83:
			goto st298
		case 84:
			goto st309
		case 95:
			goto tr221
		case 99:
			goto st289
		case 115:
			goto st298
		case 116:
			goto st309
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st289:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof289
		}
	st_case_289:
		switch lex.data[(lex.p)] {
		case 76:
			goto st290
		case 95:
			goto tr221
		case 108:
			goto st290
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st290:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof290
		}
	st_case_290:
		switch lex.data[(lex.p)] {
		case 85:
			goto st291
		case 95:
			goto tr221
		case 117:
			goto st291
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st291:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof291
		}
	st_case_291:
		switch lex.data[(lex.p)] {
		case 68:
			goto st292
		case 95:
			goto tr221
		case 100:
			goto st292
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st292:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof292
		}
	st_case_292:
		switch lex.data[(lex.p)] {
		case 69:
			goto st293
		case 95:
			goto tr221
		case 101:
			goto st293
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st293:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof293
		}
	st_case_293:
		if lex.data[(lex.p)] == 95 {
			goto st294
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr475
	st294:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof294
		}
	st_case_294:
		switch lex.data[(lex.p)] {
		case 79:
			goto st295
		case 95:
			goto tr221
		case 111:
			goto st295
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st295:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof295
		}
	st_case_295:
		switch lex.data[(lex.p)] {
		case 78:
			goto st296
		case 95:
			goto tr221
		case 110:
			goto st296
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st296:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof296
		}
	st_case_296:
		switch lex.data[(lex.p)] {
		case 67:
			goto st297
		case 95:
			goto tr221
		case 99:
			goto st297
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st297:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof297
		}
	st_case_297:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr480
		case 95:
			goto tr221
		case 101:
			goto tr480
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st298:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof298
		}
	st_case_298:
		switch lex.data[(lex.p)] {
		case 84:
			goto st299
		case 95:
			goto tr221
		case 116:
			goto st299
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st299:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof299
		}
	st_case_299:
		switch lex.data[(lex.p)] {
		case 65:
			goto st300
		case 69:
			goto st305
		case 95:
			goto tr221
		case 97:
			goto st300
		case 101:
			goto st305
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st300:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof300
		}
	st_case_300:
		switch lex.data[(lex.p)] {
		case 78:
			goto st301
		case 95:
			goto tr221
		case 110:
			goto st301
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st301:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof301
		}
	st_case_301:
		switch lex.data[(lex.p)] {
		case 67:
			goto st302
		case 95:
			goto tr221
		case 99:
			goto st302
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st302:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof302
		}
	st_case_302:
		switch lex.data[(lex.p)] {
		case 69:
			goto st303
		case 95:
			goto tr221
		case 101:
			goto st303
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st303:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof303
		}
	st_case_303:
		switch lex.data[(lex.p)] {
		case 79:
			goto st304
		case 95:
			goto tr221
		case 111:
			goto st304
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st304:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof304
		}
	st_case_304:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr488
		case 95:
			goto tr221
		case 102:
			goto tr488
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st305:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof305
		}
	st_case_305:
		switch lex.data[(lex.p)] {
		case 65:
			goto st306
		case 95:
			goto tr221
		case 97:
			goto st306
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st306:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof306
		}
	st_case_306:
		switch lex.data[(lex.p)] {
		case 68:
			goto st307
		case 95:
			goto tr221
		case 100:
			goto st307
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st307:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof307
		}
	st_case_307:
		switch lex.data[(lex.p)] {
		case 79:
			goto st308
		case 95:
			goto tr221
		case 111:
			goto st308
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st308:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof308
		}
	st_case_308:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr492
		case 95:
			goto tr221
		case 102:
			goto tr492
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st309:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof309
		}
	st_case_309:
		switch lex.data[(lex.p)] {
		case 69:
			goto st310
		case 95:
			goto tr221
		case 101:
			goto st310
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st310:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof310
		}
	st_case_310:
		switch lex.data[(lex.p)] {
		case 82:
			goto st311
		case 95:
			goto tr221
		case 114:
			goto st311
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st311:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof311
		}
	st_case_311:
		switch lex.data[(lex.p)] {
		case 70:
			goto st312
		case 95:
			goto tr221
		case 102:
			goto st312
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st312:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof312
		}
	st_case_312:
		switch lex.data[(lex.p)] {
		case 65:
			goto st313
		case 95:
			goto tr221
		case 97:
			goto st313
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st313:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof313
		}
	st_case_313:
		switch lex.data[(lex.p)] {
		case 67:
			goto st314
		case 95:
			goto tr221
		case 99:
			goto st314
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st314:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof314
		}
	st_case_314:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr498
		case 95:
			goto tr221
		case 101:
			goto tr498
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st315:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof315
		}
	st_case_315:
		switch lex.data[(lex.p)] {
		case 83:
			goto st316
		case 95:
			goto tr221
		case 115:
			goto st316
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st316:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof316
		}
	st_case_316:
		switch lex.data[(lex.p)] {
		case 69:
			goto st317
		case 95:
			goto tr221
		case 101:
			goto st317
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st317:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof317
		}
	st_case_317:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr501
		case 95:
			goto tr221
		case 116:
			goto tr501
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st318:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof318
		}
	st_case_318:
		switch lex.data[(lex.p)] {
		case 73:
			goto st319
		case 95:
			goto tr221
		case 105:
			goto st319
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st319:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof319
		}
	st_case_319:
		switch lex.data[(lex.p)] {
		case 83:
			goto st320
		case 95:
			goto tr221
		case 115:
			goto st320
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st320:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof320
		}
	st_case_320:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr504
		case 95:
			goto tr221
		case 116:
			goto tr504
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st321:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof321
		}
	st_case_321:
		switch lex.data[(lex.p)] {
		case 65:
			goto st322
		case 69:
			goto st329
		case 95:
			goto tr221
		case 97:
			goto st322
		case 101:
			goto st329
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st322:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof322
		}
	st_case_322:
		switch lex.data[(lex.p)] {
		case 77:
			goto st323
		case 95:
			goto tr221
		case 109:
			goto st323
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st323:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof323
		}
	st_case_323:
		switch lex.data[(lex.p)] {
		case 69:
			goto st324
		case 95:
			goto tr221
		case 101:
			goto st324
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st324:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof324
		}
	st_case_324:
		switch lex.data[(lex.p)] {
		case 83:
			goto st325
		case 95:
			goto tr221
		case 115:
			goto st325
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st325:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof325
		}
	st_case_325:
		switch lex.data[(lex.p)] {
		case 80:
			goto st326
		case 95:
			goto tr221
		case 112:
			goto st326
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st326:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof326
		}
	st_case_326:
		switch lex.data[(lex.p)] {
		case 65:
			goto st327
		case 95:
			goto tr221
		case 97:
			goto st327
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st327:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof327
		}
	st_case_327:
		switch lex.data[(lex.p)] {
		case 67:
			goto st328
		case 95:
			goto tr221
		case 99:
			goto st328
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st328:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof328
		}
	st_case_328:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr513
		case 95:
			goto tr221
		case 101:
			goto tr513
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st329:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof329
		}
	st_case_329:
		switch lex.data[(lex.p)] {
		case 87:
			goto tr514
		case 95:
			goto tr221
		case 119:
			goto tr514
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st330:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof330
		}
	st_case_330:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr515
		case 95:
			goto tr221
		case 114:
			goto tr515
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st331:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof331
		}
	st_case_331:
		switch lex.data[(lex.p)] {
		case 82:
			goto st332
		case 85:
			goto st344
		case 95:
			goto tr221
		case 114:
			goto st332
		case 117:
			goto st344
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st332:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof332
		}
	st_case_332:
		switch lex.data[(lex.p)] {
		case 73:
			goto st333
		case 79:
			goto st338
		case 95:
			goto tr221
		case 105:
			goto st333
		case 111:
			goto st338
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st333:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof333
		}
	st_case_333:
		switch lex.data[(lex.p)] {
		case 78:
			goto st334
		case 86:
			goto st335
		case 95:
			goto tr221
		case 110:
			goto st334
		case 118:
			goto st335
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st334:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof334
		}
	st_case_334:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr522
		case 95:
			goto tr221
		case 116:
			goto tr522
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st335:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof335
		}
	st_case_335:
		switch lex.data[(lex.p)] {
		case 65:
			goto st336
		case 95:
			goto tr221
		case 97:
			goto st336
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st336:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof336
		}
	st_case_336:
		switch lex.data[(lex.p)] {
		case 84:
			goto st337
		case 95:
			goto tr221
		case 116:
			goto st337
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st337:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof337
		}
	st_case_337:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr525
		case 95:
			goto tr221
		case 101:
			goto tr525
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st338:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof338
		}
	st_case_338:
		switch lex.data[(lex.p)] {
		case 84:
			goto st339
		case 95:
			goto tr221
		case 116:
			goto st339
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st339:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof339
		}
	st_case_339:
		switch lex.data[(lex.p)] {
		case 69:
			goto st340
		case 95:
			goto tr221
		case 101:
			goto st340
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st340:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof340
		}
	st_case_340:
		switch lex.data[(lex.p)] {
		case 67:
			goto st341
		case 95:
			goto tr221
		case 99:
			goto st341
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st341:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof341
		}
	st_case_341:
		switch lex.data[(lex.p)] {
		case 84:
			goto st342
		case 95:
			goto tr221
		case 116:
			goto st342
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st342:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof342
		}
	st_case_342:
		switch lex.data[(lex.p)] {
		case 69:
			goto st343
		case 95:
			goto tr221
		case 101:
			goto st343
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st343:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof343
		}
	st_case_343:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr531
		case 95:
			goto tr221
		case 100:
			goto tr531
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st344:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof344
		}
	st_case_344:
		switch lex.data[(lex.p)] {
		case 66:
			goto st345
		case 95:
			goto tr221
		case 98:
			goto st345
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st345:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof345
		}
	st_case_345:
		switch lex.data[(lex.p)] {
		case 76:
			goto st346
		case 95:
			goto tr221
		case 108:
			goto st346
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st346:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof346
		}
	st_case_346:
		switch lex.data[(lex.p)] {
		case 73:
			goto st347
		case 95:
			goto tr221
		case 105:
			goto st347
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st347:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof347
		}
	st_case_347:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr535
		case 95:
			goto tr221
		case 99:
			goto tr535
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st348:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof348
		}
	st_case_348:
		switch lex.data[(lex.p)] {
		case 69:
			goto st349
		case 95:
			goto tr221
		case 101:
			goto st349
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st349:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof349
		}
	st_case_349:
		switch lex.data[(lex.p)] {
		case 81:
			goto st350
		case 84:
			goto st359
		case 95:
			goto tr221
		case 113:
			goto st350
		case 116:
			goto st359
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st350:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof350
		}
	st_case_350:
		switch lex.data[(lex.p)] {
		case 85:
			goto st351
		case 95:
			goto tr221
		case 117:
			goto st351
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st351:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof351
		}
	st_case_351:
		switch lex.data[(lex.p)] {
		case 73:
			goto st352
		case 95:
			goto tr221
		case 105:
			goto st352
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st352:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof352
		}
	st_case_352:
		switch lex.data[(lex.p)] {
		case 82:
			goto st353
		case 95:
			goto tr221
		case 114:
			goto st353
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st353:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof353
		}
	st_case_353:
		switch lex.data[(lex.p)] {
		case 69:
			goto st354
		case 95:
			goto tr221
		case 101:
			goto st354
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st354:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof354
		}
	st_case_354:
		if lex.data[(lex.p)] == 95 {
			goto st355
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr543
	st355:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof355
		}
	st_case_355:
		switch lex.data[(lex.p)] {
		case 79:
			goto st356
		case 95:
			goto tr221
		case 111:
			goto st356
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st356:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof356
		}
	st_case_356:
		switch lex.data[(lex.p)] {
		case 78:
			goto st357
		case 95:
			goto tr221
		case 110:
			goto st357
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st357:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof357
		}
	st_case_357:
		switch lex.data[(lex.p)] {
		case 67:
			goto st358
		case 95:
			goto tr221
		case 99:
			goto st358
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st358:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof358
		}
	st_case_358:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr548
		case 95:
			goto tr221
		case 101:
			goto tr548
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st359:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof359
		}
	st_case_359:
		switch lex.data[(lex.p)] {
		case 85:
			goto st360
		case 95:
			goto tr221
		case 117:
			goto st360
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st360:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof360
		}
	st_case_360:
		switch lex.data[(lex.p)] {
		case 82:
			goto st361
		case 95:
			goto tr221
		case 114:
			goto st361
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st361:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof361
		}
	st_case_361:
		switch lex.data[(lex.p)] {
		case 78:
			goto tr551
		case 95:
			goto tr221
		case 110:
			goto tr551
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st362:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof362
		}
	st_case_362:
		switch lex.data[(lex.p)] {
		case 84:
			goto st363
		case 87:
			goto st367
		case 95:
			goto tr221
		case 116:
			goto st363
		case 119:
			goto st367
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st363:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof363
		}
	st_case_363:
		switch lex.data[(lex.p)] {
		case 65:
			goto st364
		case 95:
			goto tr221
		case 97:
			goto st364
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st364:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof364
		}
	st_case_364:
		switch lex.data[(lex.p)] {
		case 84:
			goto st365
		case 95:
			goto tr221
		case 116:
			goto st365
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st365:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof365
		}
	st_case_365:
		switch lex.data[(lex.p)] {
		case 73:
			goto st366
		case 95:
			goto tr221
		case 105:
			goto st366
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st366:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof366
		}
	st_case_366:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr557
		case 95:
			goto tr221
		case 99:
			goto tr557
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st367:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof367
		}
	st_case_367:
		switch lex.data[(lex.p)] {
		case 73:
			goto st368
		case 95:
			goto tr221
		case 105:
			goto st368
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st368:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof368
		}
	st_case_368:
		switch lex.data[(lex.p)] {
		case 84:
			goto st369
		case 95:
			goto tr221
		case 116:
			goto st369
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st369:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof369
		}
	st_case_369:
		switch lex.data[(lex.p)] {
		case 67:
			goto st370
		case 95:
			goto tr221
		case 99:
			goto st370
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st370:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof370
		}
	st_case_370:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr561
		case 95:
			goto tr221
		case 104:
			goto tr561
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st371:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof371
		}
	st_case_371:
		switch lex.data[(lex.p)] {
		case 72:
			goto st372
		case 82:
			goto st375
		case 95:
			goto tr221
		case 104:
			goto st372
		case 114:
			goto st375
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st372:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof372
		}
	st_case_372:
		switch lex.data[(lex.p)] {
		case 82:
			goto st373
		case 95:
			goto tr221
		case 114:
			goto st373
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st373:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof373
		}
	st_case_373:
		switch lex.data[(lex.p)] {
		case 79:
			goto st374
		case 95:
			goto tr221
		case 111:
			goto st374
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st374:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof374
		}
	st_case_374:
		switch lex.data[(lex.p)] {
		case 87:
			goto tr566
		case 95:
			goto tr221
		case 119:
			goto tr566
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st375:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof375
		}
	st_case_375:
		switch lex.data[(lex.p)] {
		case 65:
			goto st376
		case 89:
			goto tr568
		case 95:
			goto tr221
		case 97:
			goto st376
		case 121:
			goto tr568
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st376:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof376
		}
	st_case_376:
		switch lex.data[(lex.p)] {
		case 73:
			goto st377
		case 95:
			goto tr221
		case 105:
			goto st377
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st377:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof377
		}
	st_case_377:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr570
		case 95:
			goto tr221
		case 116:
			goto tr570
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st378:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof378
		}
	st_case_378:
		switch lex.data[(lex.p)] {
		case 78:
			goto st379
		case 83:
			goto st382
		case 95:
			goto tr221
		case 110:
			goto st379
		case 115:
			goto st382
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st379:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof379
		}
	st_case_379:
		switch lex.data[(lex.p)] {
		case 83:
			goto st380
		case 95:
			goto tr221
		case 115:
			goto st380
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st380:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof380
		}
	st_case_380:
		switch lex.data[(lex.p)] {
		case 69:
			goto st381
		case 95:
			goto tr221
		case 101:
			goto st381
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st381:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof381
		}
	st_case_381:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr575
		case 95:
			goto tr221
		case 116:
			goto tr575
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st382:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof382
		}
	st_case_382:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr576
		case 95:
			goto tr221
		case 101:
			goto tr576
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st383:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof383
		}
	st_case_383:
		switch lex.data[(lex.p)] {
		case 65:
			goto st384
		case 95:
			goto tr221
		case 97:
			goto st384
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st384:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof384
		}
	st_case_384:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr578
		case 95:
			goto tr221
		case 114:
			goto tr578
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st385:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof385
		}
	st_case_385:
		switch lex.data[(lex.p)] {
		case 72:
			goto st386
		case 95:
			goto tr221
		case 104:
			goto st386
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st386:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof386
		}
	st_case_386:
		switch lex.data[(lex.p)] {
		case 73:
			goto st387
		case 95:
			goto tr221
		case 105:
			goto st387
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st387:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof387
		}
	st_case_387:
		switch lex.data[(lex.p)] {
		case 76:
			goto st388
		case 95:
			goto tr221
		case 108:
			goto st388
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st388:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof388
		}
	st_case_388:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr582
		case 95:
			goto tr221
		case 101:
			goto tr582
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st389:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof389
		}
	st_case_389:
		switch lex.data[(lex.p)] {
		case 79:
			goto st390
		case 95:
			goto tr221
		case 111:
			goto st390
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st390:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof390
		}
	st_case_390:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr584
		case 95:
			goto tr221
		case 114:
			goto tr584
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st391:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof391
		}
	st_case_391:
		switch lex.data[(lex.p)] {
		case 73:
			goto st392
		case 95:
			goto tr221
		case 105:
			goto st392
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st392:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof392
		}
	st_case_392:
		switch lex.data[(lex.p)] {
		case 69:
			goto st393
		case 95:
			goto tr221
		case 101:
			goto st393
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st393:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof393
		}
	st_case_393:
		switch lex.data[(lex.p)] {
		case 76:
			goto st394
		case 95:
			goto tr221
		case 108:
			goto st394
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st394:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof394
		}
	st_case_394:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr588
		case 95:
			goto tr221
		case 100:
			goto tr588
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	tr588:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st395
	st395:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof395
		}
	st_case_395:
//line scanner/scanner.go:12364
		switch lex.data[(lex.p)] {
		case 10:
			goto st95
		case 13:
			goto st96
		case 32:
			goto st94
		case 70:
			goto st396
		case 95:
			goto tr221
		case 102:
			goto st396
		}
		switch {
		case lex.data[(lex.p)] < 48:
			if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
				goto st94
			}
		case lex.data[(lex.p)] > 57:
			switch {
			case lex.data[(lex.p)] > 90:
				if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
					goto tr221
				}
			case lex.data[(lex.p)] >= 65:
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr589
	tr147:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st94
	st94:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof94
		}
	st_case_94:
//line scanner/scanner.go:12406
		switch lex.data[(lex.p)] {
		case 10:
			goto st95
		case 13:
			goto st96
		case 32:
			goto st94
		case 70:
			goto st97
		case 102:
			goto st97
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st94
		}
		goto tr142
	tr148:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st95
	st95:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof95
		}
	st_case_95:
//line scanner/scanner.go:12432
		switch lex.data[(lex.p)] {
		case 10:
			goto tr148
		case 13:
			goto tr149
		case 32:
			goto tr147
		case 70:
			goto tr150
		case 102:
			goto tr150
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr147
		}
		goto tr142
	tr149:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st96
	st96:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof96
		}
	st_case_96:
//line scanner/scanner.go:12458
		if lex.data[(lex.p)] == 10 {
			goto st95
		}
		goto tr142
	tr150:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st97
	st97:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof97
		}
	st_case_97:
//line scanner/scanner.go:12472
		switch lex.data[(lex.p)] {
		case 82:
			goto st98
		case 114:
			goto st98
		}
		goto tr142
	st98:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch lex.data[(lex.p)] {
		case 79:
			goto st99
		case 111:
			goto st99
		}
		goto tr142
	st99:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr153
		case 109:
			goto tr153
		}
		goto tr142
	st396:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof396
		}
	st_case_396:
		switch lex.data[(lex.p)] {
		case 82:
			goto st397
		case 95:
			goto tr221
		case 114:
			goto st397
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st397:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof397
		}
	st_case_397:
		switch lex.data[(lex.p)] {
		case 79:
			goto st398
		case 95:
			goto tr221
		case 111:
			goto st398
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st398:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof398
		}
	st_case_398:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr593
		case 95:
			goto tr221
		case 109:
			goto tr593
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st399:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof399
		}
	st_case_399:
		if lex.data[(lex.p)] == 61 {
			goto tr594
		}
		goto tr249
	st400:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof400
		}
	st_case_400:
		if lex.data[(lex.p)] == 95 {
			goto st401
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st401:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof401
		}
	st_case_401:
		switch lex.data[(lex.p)] {
		case 67:
			goto st402
		case 68:
			goto st408
		case 70:
			goto st412
		case 72:
			goto st425
		case 76:
			goto st437
		case 77:
			goto st442
		case 78:
			goto st449
		case 84:
			goto st459
		case 95:
			goto tr221
		case 99:
			goto st402
		case 100:
			goto st408
		case 102:
			goto st412
		case 104:
			goto st425
		case 108:
			goto st437
		case 109:
			goto st442
		case 110:
			goto st449
		case 116:
			goto st459
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st402:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof402
		}
	st_case_402:
		switch lex.data[(lex.p)] {
		case 76:
			goto st403
		case 95:
			goto tr221
		case 108:
			goto st403
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st403:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof403
		}
	st_case_403:
		switch lex.data[(lex.p)] {
		case 65:
			goto st404
		case 95:
			goto tr221
		case 97:
			goto st404
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st404:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof404
		}
	st_case_404:
		switch lex.data[(lex.p)] {
		case 83:
			goto st405
		case 95:
			goto tr221
		case 115:
			goto st405
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st405:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof405
		}
	st_case_405:
		switch lex.data[(lex.p)] {
		case 83:
			goto st406
		case 95:
			goto tr221
		case 115:
			goto st406
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st406:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof406
		}
	st_case_406:
		if lex.data[(lex.p)] == 95 {
			goto st407
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st407:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof407
		}
	st_case_407:
		if lex.data[(lex.p)] == 95 {
			goto tr609
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st408:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof408
		}
	st_case_408:
		switch lex.data[(lex.p)] {
		case 73:
			goto st409
		case 95:
			goto tr221
		case 105:
			goto st409
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st409:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof409
		}
	st_case_409:
		switch lex.data[(lex.p)] {
		case 82:
			goto st410
		case 95:
			goto tr221
		case 114:
			goto st410
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st410:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof410
		}
	st_case_410:
		if lex.data[(lex.p)] == 95 {
			goto st411
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st411:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof411
		}
	st_case_411:
		if lex.data[(lex.p)] == 95 {
			goto tr613
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st412:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof412
		}
	st_case_412:
		switch lex.data[(lex.p)] {
		case 73:
			goto st413
		case 85:
			goto st417
		case 95:
			goto tr221
		case 105:
			goto st413
		case 117:
			goto st417
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st413:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof413
		}
	st_case_413:
		switch lex.data[(lex.p)] {
		case 76:
			goto st414
		case 95:
			goto tr221
		case 108:
			goto st414
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st414:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof414
		}
	st_case_414:
		switch lex.data[(lex.p)] {
		case 69:
			goto st415
		case 95:
			goto tr221
		case 101:
			goto st415
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st415:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof415
		}
	st_case_415:
		if lex.data[(lex.p)] == 95 {
			goto st416
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st416:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof416
		}
	st_case_416:
		if lex.data[(lex.p)] == 95 {
			goto tr619
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st417:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof417
		}
	st_case_417:
		switch lex.data[(lex.p)] {
		case 78:
			goto st418
		case 95:
			goto tr221
		case 110:
			goto st418
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st418:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof418
		}
	st_case_418:
		switch lex.data[(lex.p)] {
		case 67:
			goto st419
		case 95:
			goto tr221
		case 99:
			goto st419
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st419:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof419
		}
	st_case_419:
		switch lex.data[(lex.p)] {
		case 84:
			goto st420
		case 95:
			goto tr221
		case 116:
			goto st420
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st420:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof420
		}
	st_case_420:
		switch lex.data[(lex.p)] {
		case 73:
			goto st421
		case 95:
			goto tr221
		case 105:
			goto st421
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st421:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof421
		}
	st_case_421:
		switch lex.data[(lex.p)] {
		case 79:
			goto st422
		case 95:
			goto tr221
		case 111:
			goto st422
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st422:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof422
		}
	st_case_422:
		switch lex.data[(lex.p)] {
		case 78:
			goto st423
		case 95:
			goto tr221
		case 110:
			goto st423
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st423:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof423
		}
	st_case_423:
		if lex.data[(lex.p)] == 95 {
			goto st424
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st424:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof424
		}
	st_case_424:
		if lex.data[(lex.p)] == 95 {
			goto tr627
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st425:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof425
		}
	st_case_425:
		switch lex.data[(lex.p)] {
		case 65:
			goto st426
		case 95:
			goto tr221
		case 97:
			goto st426
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st426:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof426
		}
	st_case_426:
		switch lex.data[(lex.p)] {
		case 76:
			goto st427
		case 95:
			goto tr221
		case 108:
			goto st427
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st427:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof427
		}
	st_case_427:
		switch lex.data[(lex.p)] {
		case 84:
			goto st428
		case 95:
			goto tr221
		case 116:
			goto st428
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st428:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof428
		}
	st_case_428:
		if lex.data[(lex.p)] == 95 {
			goto st429
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st429:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof429
		}
	st_case_429:
		switch lex.data[(lex.p)] {
		case 67:
			goto st430
		case 95:
			goto tr221
		case 99:
			goto st430
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st430:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof430
		}
	st_case_430:
		switch lex.data[(lex.p)] {
		case 79:
			goto st431
		case 95:
			goto tr221
		case 111:
			goto st431
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st431:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof431
		}
	st_case_431:
		switch lex.data[(lex.p)] {
		case 77:
			goto st432
		case 95:
			goto tr221
		case 109:
			goto st432
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st432:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof432
		}
	st_case_432:
		switch lex.data[(lex.p)] {
		case 80:
			goto st433
		case 95:
			goto tr221
		case 112:
			goto st433
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st433:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof433
		}
	st_case_433:
		switch lex.data[(lex.p)] {
		case 73:
			goto st434
		case 95:
			goto tr221
		case 105:
			goto st434
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st434:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof434
		}
	st_case_434:
		switch lex.data[(lex.p)] {
		case 76:
			goto st435
		case 95:
			goto tr221
		case 108:
			goto st435
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st435:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof435
		}
	st_case_435:
		switch lex.data[(lex.p)] {
		case 69:
			goto st436
		case 95:
			goto tr221
		case 101:
			goto st436
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st436:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof436
		}
	st_case_436:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr639
		case 95:
			goto tr221
		case 114:
			goto tr639
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st437:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof437
		}
	st_case_437:
		switch lex.data[(lex.p)] {
		case 73:
			goto st438
		case 95:
			goto tr221
		case 105:
			goto st438
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st438:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof438
		}
	st_case_438:
		switch lex.data[(lex.p)] {
		case 78:
			goto st439
		case 95:
			goto tr221
		case 110:
			goto st439
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st439:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof439
		}
	st_case_439:
		switch lex.data[(lex.p)] {
		case 69:
			goto st440
		case 95:
			goto tr221
		case 101:
			goto st440
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st440:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof440
		}
	st_case_440:
		if lex.data[(lex.p)] == 95 {
			goto st441
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st441:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof441
		}
	st_case_441:
		if lex.data[(lex.p)] == 95 {
			goto tr644
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st442:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof442
		}
	st_case_442:
		switch lex.data[(lex.p)] {
		case 69:
			goto st443
		case 95:
			goto tr221
		case 101:
			goto st443
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st443:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof443
		}
	st_case_443:
		switch lex.data[(lex.p)] {
		case 84:
			goto st444
		case 95:
			goto tr221
		case 116:
			goto st444
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st444:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof444
		}
	st_case_444:
		switch lex.data[(lex.p)] {
		case 72:
			goto st445
		case 95:
			goto tr221
		case 104:
			goto st445
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st445:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof445
		}
	st_case_445:
		switch lex.data[(lex.p)] {
		case 79:
			goto st446
		case 95:
			goto tr221
		case 111:
			goto st446
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st446:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof446
		}
	st_case_446:
		switch lex.data[(lex.p)] {
		case 68:
			goto st447
		case 95:
			goto tr221
		case 100:
			goto st447
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st447:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof447
		}
	st_case_447:
		if lex.data[(lex.p)] == 95 {
			goto st448
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st448:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof448
		}
	st_case_448:
		if lex.data[(lex.p)] == 95 {
			goto tr651
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st449:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof449
		}
	st_case_449:
		switch lex.data[(lex.p)] {
		case 65:
			goto st450
		case 95:
			goto tr221
		case 97:
			goto st450
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st450:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof450
		}
	st_case_450:
		switch lex.data[(lex.p)] {
		case 77:
			goto st451
		case 95:
			goto tr221
		case 109:
			goto st451
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st451:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof451
		}
	st_case_451:
		switch lex.data[(lex.p)] {
		case 69:
			goto st452
		case 95:
			goto tr221
		case 101:
			goto st452
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st452:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof452
		}
	st_case_452:
		switch lex.data[(lex.p)] {
		case 83:
			goto st453
		case 95:
			goto tr221
		case 115:
			goto st453
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st453:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof453
		}
	st_case_453:
		switch lex.data[(lex.p)] {
		case 80:
			goto st454
		case 95:
			goto tr221
		case 112:
			goto st454
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st454:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof454
		}
	st_case_454:
		switch lex.data[(lex.p)] {
		case 65:
			goto st455
		case 95:
			goto tr221
		case 97:
			goto st455
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st455:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof455
		}
	st_case_455:
		switch lex.data[(lex.p)] {
		case 67:
			goto st456
		case 95:
			goto tr221
		case 99:
			goto st456
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st456:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof456
		}
	st_case_456:
		switch lex.data[(lex.p)] {
		case 69:
			goto st457
		case 95:
			goto tr221
		case 101:
			goto st457
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st457:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof457
		}
	st_case_457:
		if lex.data[(lex.p)] == 95 {
			goto st458
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st458:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof458
		}
	st_case_458:
		if lex.data[(lex.p)] == 95 {
			goto tr661
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st459:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof459
		}
	st_case_459:
		switch lex.data[(lex.p)] {
		case 82:
			goto st460
		case 95:
			goto tr221
		case 114:
			goto st460
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st460:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof460
		}
	st_case_460:
		switch lex.data[(lex.p)] {
		case 65:
			goto st461
		case 95:
			goto tr221
		case 97:
			goto st461
		}
		switch {
		case lex.data[(lex.p)] < 66:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st461:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof461
		}
	st_case_461:
		switch lex.data[(lex.p)] {
		case 73:
			goto st462
		case 95:
			goto tr221
		case 105:
			goto st462
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st462:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof462
		}
	st_case_462:
		switch lex.data[(lex.p)] {
		case 84:
			goto st463
		case 95:
			goto tr221
		case 116:
			goto st463
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st463:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof463
		}
	st_case_463:
		if lex.data[(lex.p)] == 95 {
			goto st464
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st464:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof464
		}
	st_case_464:
		if lex.data[(lex.p)] == 95 {
			goto tr667
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr221
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr285
	st465:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof465
		}
	st_case_465:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr668
		case 124:
			goto tr669
		}
		goto tr249
	tr154:
//line scanner/scanner.rl:346
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st466
	tr670:
//line scanner/scanner.rl:349
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				goto st118
			}
		}
		goto st466
	tr675:
//line scanner/scanner.rl:346
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st466
	tr677:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:346
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st466
	tr681:
//line scanner/scanner.rl:349
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				goto st118
			}
		}
		goto st466
	tr682:
//line scanner/scanner.rl:347
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 466
				goto _out
			}
		}
		goto st466
	tr683:
		lex.cs = 466
//line scanner/scanner.rl:348
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_STRING
			lex.cs = 118
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st466:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof466
		}
	st_case_466:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:14288
		switch lex.data[(lex.p)] {
		case 10:
			goto tr155
		case 13:
			goto st469
		case 32:
			goto tr671
		case 45:
			goto st470
		case 55:
			goto st471
		case 95:
			goto st471
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
				goto tr671
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st471
			}
		default:
			goto st471
		}
		goto tr670
	tr671:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st467
	tr678:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st467
	st467:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof467
		}
	st_case_467:
//line scanner/scanner.go:14333
		switch lex.data[(lex.p)] {
		case 10:
			goto tr155
		case 13:
			goto st100
		case 32:
			goto tr671
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr671
		}
		goto tr675
	tr155:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st468
	tr679:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st468
	st468:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof468
		}
	st_case_468:
//line scanner/scanner.go:14363
		switch lex.data[(lex.p)] {
		case 10:
			goto tr679
		case 13:
			goto tr680
		case 32:
			goto tr678
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr678
		}
		goto tr677
	tr680:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st100
	st100:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof100
		}
	st_case_100:
//line scanner/scanner.go:14385
		if lex.data[(lex.p)] == 10 {
			goto tr155
		}
		goto tr154
	st469:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof469
		}
	st_case_469:
		if lex.data[(lex.p)] == 10 {
			goto tr155
		}
		goto tr681
	st470:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof470
		}
	st_case_470:
		if lex.data[(lex.p)] == 62 {
			goto tr682
		}
		goto tr681
	st471:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof471
		}
	st_case_471:
		if lex.data[(lex.p)] == 95 {
			goto st471
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st471
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st471
			}
		default:
			goto st471
		}
		goto tr683
	tr686:
		lex.cs = 472
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

				lex.createToken(lval)
				tok = T_ENCAPSED_AND_WHITESPACE
				lex.cs = 493
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr687:
		lex.cs = 472
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:353
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.cs = 493
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st472:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof472
		}
	st_case_472:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:14474
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
			goto st474
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr684
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	tr684:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:353
		lex.act = 142
		goto st473
	tr688:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:353
		lex.act = 142
		goto st473
	st473:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof473
		}
	st_case_473:
//line scanner/scanner.go:14545
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
			goto st474
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr684
		}
		goto tr686
	tr689:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st474
	st474:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof474
		}
	st_case_474:
//line scanner/scanner.go:14600
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
			goto tr689
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr688
		}
		goto tr687
	tr156:
//line scanner/scanner.rl:362
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.createToken(lval)
			tok = T_CURLY_OPEN
			lex.call(475, 118)
			goto _out
		}
		goto st475
	tr696:
//line scanner/scanner.rl:364
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 475
					lex.top++
					goto st495
				}
			}
		}
		goto st475
	tr697:
//line scanner/scanner.rl:363
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(475, 511)
			goto _out
		}
		goto st475
	tr698:
		lex.cs = 475
//line NONE:1
		switch lex.act {
		case 143:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.createToken(lval)
				tok = T_CURLY_OPEN
				lex.call(475, 118)
				goto _out
			}
		case 144:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(475, 511)
				goto _out
			}
		case 146:
			{
				(lex.p) = (lex.te) - 1

				lex.createToken(lval)
				tok = T_ENCAPSED_AND_WHITESPACE

				if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
					lex.cs = 493
				}
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr699:
		lex.cs = 475
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:365
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE

			if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 493
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr703:
		lex.cs = 475
//line scanner/scanner.rl:365
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE

			if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 493
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st475:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof475
		}
	st_case_475:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:14729
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
			goto st476
		case 1403:
			goto st101
		case 1546:
			goto st478
		case 1572:
			goto st479
		case 1659:
			goto st480
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr692
		}
		goto st0
	st476:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof476
		}
	st_case_476:
		if lex.data[(lex.p)] == 123 {
			goto tr697
		}
		goto tr696
	st101:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if lex.data[(lex.p)] == 36 {
			goto tr156
		}
		goto st0
	tr692:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:365
		lex.act = 146
		goto st477
	tr700:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:365
		lex.act = 146
		goto st477
	tr702:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:363
		lex.act = 144
		goto st477
	tr704:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:362
		lex.act = 143
		goto st477
	st477:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof477
		}
	st_case_477:
//line scanner/scanner.go:14837
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
			goto st478
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr692
		}
		goto tr698
	tr701:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st478
	st478:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof478
		}
	st_case_478:
//line scanner/scanner.go:14892
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
			goto tr701
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr700
		}
		goto tr699
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
			goto tr697
		case 1546:
			goto st478
		case 1659:
			goto tr702
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr692
		}
		goto tr696
	st480:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof480
		}
	st_case_480:
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
			goto tr156
		case 1546:
			goto st478
		case 1572:
			goto tr704
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr692
		}
		goto tr703
	tr158:
//line scanner/scanner.rl:377
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.createToken(lval)
			tok = T_CURLY_OPEN
			lex.call(481, 118)
			goto _out
		}
		goto st481
	tr706:
		lex.cs = 481
//line scanner/scanner.rl:380
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int('`'))
			lex.cs = 118
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr713:
//line scanner/scanner.rl:379
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 481
					lex.top++
					goto st495
				}
			}
		}
		goto st481
	tr714:
//line scanner/scanner.rl:378
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(481, 511)
			goto _out
		}
		goto st481
	tr715:
		lex.cs = 481
//line NONE:1
		switch lex.act {
		case 147:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.createToken(lval)
				tok = T_CURLY_OPEN
				lex.call(481, 118)
				goto _out
			}
		case 148:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(481, 511)
				goto _out
			}
		case 150:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = TokenID(int('`'))
				lex.cs = 118
				{
					(lex.p)++
					goto _out
				}
			}
		case 151:
			{
				(lex.p) = (lex.te) - 1

				lex.createToken(lval)
				tok = T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr716:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:381
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 481
				goto _out
			}
		}
		goto st481
	tr720:
//line scanner/scanner.rl:381
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 481
				goto _out
			}
		}
		goto st481
	st481:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof481
		}
	st_case_481:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:15126
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
			goto st482
		case 1888:
			goto tr706
		case 1915:
			goto st102
		case 2058:
			goto st484
		case 2084:
			goto st485
		case 2144:
			goto tr711
		case 2171:
			goto st486
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr708
		}
		goto st0
	st482:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof482
		}
	st_case_482:
		if lex.data[(lex.p)] == 123 {
			goto tr714
		}
		goto tr713
	st102:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		if lex.data[(lex.p)] == 36 {
			goto tr158
		}
		goto st0
	tr708:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:381
		lex.act = 151
		goto st483
	tr711:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:380
		lex.act = 150
		goto st483
	tr717:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:381
		lex.act = 151
		goto st483
	tr719:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:378
		lex.act = 148
		goto st483
	tr721:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:377
		lex.act = 147
		goto st483
	st483:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof483
		}
	st_case_483:
//line scanner/scanner.go:15245
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
			goto st484
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr708
		}
		goto tr715
	tr718:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st484
	st484:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof484
		}
	st_case_484:
//line scanner/scanner.go:15300
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
			goto tr718
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr717
		}
		goto tr716
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
			goto tr714
		case 2058:
			goto st484
		case 2171:
			goto tr719
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr708
		}
		goto tr713
	st486:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof486
		}
	st_case_486:
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
			goto tr158
		case 2058:
			goto st484
		case 2084:
			goto tr721
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr708
		}
		goto tr720
	tr159:
//line scanner/scanner.rl:389
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.createToken(lval)
			tok = T_CURLY_OPEN
			lex.call(487, 118)
			goto _out
		}
		goto st487
	tr722:
		lex.cs = 487
//line scanner/scanner.rl:392
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int('"'))
			lex.cs = 118
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr730:
//line scanner/scanner.rl:391
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 487
					lex.top++
					goto st495
				}
			}
		}
		goto st487
	tr731:
//line scanner/scanner.rl:390
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(487, 511)
			goto _out
		}
		goto st487
	tr732:
		lex.cs = 487
//line NONE:1
		switch lex.act {
		case 152:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.createToken(lval)
				tok = T_CURLY_OPEN
				lex.call(487, 118)
				goto _out
			}
		case 153:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(487, 511)
				goto _out
			}
		case 155:
			{
				(lex.p) = (lex.te) - 1
				lex.createToken(lval)
				tok = TokenID(int('"'))
				lex.cs = 118
				{
					(lex.p)++
					goto _out
				}
			}
		case 156:
			{
				(lex.p) = (lex.te) - 1

				lex.createToken(lval)
				tok = T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr733:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:393
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 487
				goto _out
			}
		}
		goto st487
	tr737:
//line scanner/scanner.rl:393
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 487
				goto _out
			}
		}
		goto st487
	st487:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof487
		}
	st_case_487:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:15534
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
			goto tr722
		case 2340:
			goto st488
		case 2427:
			goto st103
		case 2570:
			goto st490
		case 2594:
			goto tr727
		case 2596:
			goto st491
		case 2683:
			goto st492
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr725
		}
		goto st0
	st488:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof488
		}
	st_case_488:
		if lex.data[(lex.p)] == 123 {
			goto tr731
		}
		goto tr730
	st103:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		if lex.data[(lex.p)] == 36 {
			goto tr159
		}
		goto st0
	tr725:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:393
		lex.act = 156
		goto st489
	tr727:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:392
		lex.act = 155
		goto st489
	tr734:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:393
		lex.act = 156
		goto st489
	tr736:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:390
		lex.act = 153
		goto st489
	tr738:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:389
		lex.act = 152
		goto st489
	st489:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof489
		}
	st_case_489:
//line scanner/scanner.go:15653
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
			goto st490
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr725
		}
		goto tr732
	tr735:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st490
	st490:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof490
		}
	st_case_490:
//line scanner/scanner.go:15708
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
			goto tr735
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr734
		}
		goto tr733
	st491:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof491
		}
	st_case_491:
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
			goto tr731
		case 2570:
			goto st490
		case 2683:
			goto tr736
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr725
		}
		goto tr730
	st492:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof492
		}
	st_case_492:
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
			goto tr159
		case 2570:
			goto st490
		case 2596:
			goto tr738
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr725
		}
		goto tr737
	tr740:
		lex.cs = 493
//line scanner/scanner.rl:401
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_END_HEREDOC
			lex.cs = 118
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st493:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof493
		}
	st_case_493:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:15887
		switch lex.data[(lex.p)] {
		case 55:
			goto st494
		case 95:
			goto st494
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st494
			}
		case lex.data[(lex.p)] >= 65:
			goto st494
		}
		goto st0
	st494:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof494
		}
	st_case_494:
		if lex.data[(lex.p)] == 95 {
			goto st494
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st494
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st494
			}
		default:
			goto st494
		}
		goto tr740
	tr160:
//line scanner/scanner.rl:420
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st495
	tr161:
//line scanner/scanner.rl:417
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.createToken(lval)
			tok = T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 495
				goto _out
			}
		}
		goto st495
	tr741:
//line scanner/scanner.rl:420
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st495
	tr745:
//line scanner/scanner.rl:419
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int('['))
			lex.call(495, 500)
			goto _out
		}
		goto st495
	tr746:
//line scanner/scanner.rl:420
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
		goto st495
	tr748:
//line scanner/scanner.rl:416
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 495
				goto _out
			}
		}
		goto st495
	tr750:
//line scanner/scanner.rl:418
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 495
				goto _out
			}
		}
		goto st495
	st495:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof495
		}
	st_case_495:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:15973
		switch lex.data[(lex.p)] {
		case 36:
			goto st496
		case 45:
			goto tr743
		case 55:
			goto st499
		case 91:
			goto tr745
		case 95:
			goto st499
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st499
			}
		case lex.data[(lex.p)] >= 65:
			goto st499
		}
		goto tr741
	st496:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof496
		}
	st_case_496:
		switch lex.data[(lex.p)] {
		case 55:
			goto st497
		case 95:
			goto st497
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st497
			}
		case lex.data[(lex.p)] >= 65:
			goto st497
		}
		goto tr746
	st497:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof497
		}
	st_case_497:
		if lex.data[(lex.p)] == 95 {
			goto st497
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st497
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st497
			}
		default:
			goto st497
		}
		goto tr748
	tr743:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st498
	st498:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof498
		}
	st_case_498:
//line scanner/scanner.go:16046
		if lex.data[(lex.p)] == 62 {
			goto st104
		}
		goto tr746
	st104:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch lex.data[(lex.p)] {
		case 55:
			goto tr161
		case 95:
			goto tr161
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr161
			}
		case lex.data[(lex.p)] >= 65:
			goto tr161
		}
		goto tr160
	st499:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof499
		}
	st_case_499:
		if lex.data[(lex.p)] == 95 {
			goto st499
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st499
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st499
			}
		default:
			goto st499
		}
		goto tr750
	tr162:
//line scanner/scanner.rl:424
		(lex.p) = (lex.te) - 1
		{
			lex.createToken(lval)
			tok = T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 500
				goto _out
			}
		}
		goto st500
	tr751:
//line scanner/scanner.rl:430
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st500
	tr752:
//line scanner/scanner.rl:427
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st500
	tr755:
//line scanner/scanner.rl:428
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 500
				goto _out
			}
		}
		goto st500
	tr761:
//line scanner/scanner.rl:429
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int(']'))
			lex.ret(2)
			goto _out
		}
		goto st500
	tr762:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:427
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st500
	tr763:
//line scanner/scanner.rl:430
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st500
	tr764:
//line scanner/scanner.rl:428
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 500
				goto _out
			}
		}
		goto st500
	tr766:
//line scanner/scanner.rl:425
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 500
				goto _out
			}
		}
		goto st500
	tr767:
//line scanner/scanner.rl:424
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 500
				goto _out
			}
		}
		goto st500
	tr770:
//line scanner/scanner.rl:426
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.createToken(lval)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 500
				goto _out
			}
		}
		goto st500
	st500:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof500
		}
	st_case_500:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16172
		switch lex.data[(lex.p)] {
		case 10:
			goto st501
		case 13:
			goto st502
		case 32:
			goto tr752
		case 33:
			goto tr755
		case 35:
			goto tr752
		case 36:
			goto st503
		case 39:
			goto tr752
		case 48:
			goto tr757
		case 55:
			goto st509
		case 91:
			goto tr755
		case 92:
			goto tr752
		case 93:
			goto tr761
		case 94:
			goto tr755
		case 124:
			goto tr755
		case 126:
			goto tr755
		}
		switch {
		case lex.data[(lex.p)] < 49:
			switch {
			case lex.data[(lex.p)] > 12:
				if 37 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 47 {
					goto tr755
				}
			case lex.data[(lex.p)] >= 9:
				goto tr752
			}
		case lex.data[(lex.p)] > 57:
			switch {
			case lex.data[(lex.p)] < 65:
				if 58 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 64 {
					goto tr755
				}
			case lex.data[(lex.p)] > 95:
				if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
					goto st510
				}
			default:
				goto st510
			}
		default:
			goto st506
		}
		goto tr751
	st501:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof501
		}
	st_case_501:
		goto tr762
	st502:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof502
		}
	st_case_502:
		if lex.data[(lex.p)] == 10 {
			goto st501
		}
		goto tr763
	st503:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof503
		}
	st_case_503:
		switch lex.data[(lex.p)] {
		case 55:
			goto st504
		case 95:
			goto st504
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st504
			}
		case lex.data[(lex.p)] >= 65:
			goto st504
		}
		goto tr764
	st504:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof504
		}
	st_case_504:
		if lex.data[(lex.p)] == 95 {
			goto st504
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st504
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st504
			}
		default:
			goto st504
		}
		goto tr766
	tr757:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st505
	st505:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof505
		}
	st_case_505:
//line scanner/scanner.go:16298
		switch lex.data[(lex.p)] {
		case 98:
			goto st105
		case 120:
			goto st106
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st506
		}
		goto tr767
	st506:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof506
		}
	st_case_506:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st506
		}
		goto tr767
	st105:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st507
		}
		goto tr162
	st507:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof507
		}
	st_case_507:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto st507
		}
		goto tr767
	st106:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st508
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st508
			}
		default:
			goto st508
		}
		goto tr162
	st508:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof508
		}
	st_case_508:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st508
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto st508
			}
		default:
			goto st508
		}
		goto tr767
	st509:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof509
		}
	st_case_509:
		if lex.data[(lex.p)] == 95 {
			goto st510
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st509
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st510
			}
		default:
			goto st510
		}
		goto tr767
	st510:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof510
		}
	st_case_510:
		if lex.data[(lex.p)] == 95 {
			goto st510
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st510
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st510
			}
		default:
			goto st510
		}
		goto tr770
	tr165:
		lex.cs = 511
//line scanner/scanner.rl:438
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	tr167:
		lex.cs = 511
//line scanner/scanner.rl:437
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.createToken(lval)
			tok = T_STRING_VARNAME
			lex.cs = 118
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr771:
		lex.cs = 511
//line scanner/scanner.rl:438
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	tr773:
		lex.cs = 511
//line scanner/scanner.rl:438
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	st511:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof511
		}
	st_case_511:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16450
		switch lex.data[(lex.p)] {
		case 55:
			goto tr772
		case 95:
			goto tr772
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr772
			}
		case lex.data[(lex.p)] >= 65:
			goto tr772
		}
		goto tr771
	tr772:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st512
	st512:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof512
		}
	st_case_512:
//line scanner/scanner.go:16476
		switch lex.data[(lex.p)] {
		case 91:
			goto tr167
		case 95:
			goto st107
		case 125:
			goto tr167
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st107
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st107
			}
		default:
			goto st107
		}
		goto tr773
	st107:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch lex.data[(lex.p)] {
		case 91:
			goto tr167
		case 95:
			goto st107
		case 125:
			goto tr167
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st107
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st107
			}
		default:
			goto st107
		}
		goto tr165
	tr168:
//line scanner/scanner.rl:442
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st513
	tr774:
		lex.cs = 513
//line scanner/scanner.rl:444
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	tr777:
		lex.cs = 513
//line scanner/scanner.rl:443
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int('('))
			lex.cs = 517
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr778:
//line scanner/scanner.rl:442
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st513
	tr780:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:442
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st513
	tr784:
		lex.cs = 513
//line scanner/scanner.rl:444
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 118
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

//line scanner/scanner.go:16573
		switch lex.data[(lex.p)] {
		case 10:
			goto tr169
		case 13:
			goto st516
		case 32:
			goto tr775
		case 40:
			goto tr777
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr775
		}
		goto tr774
	tr775:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st514
	tr781:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st514
	st514:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof514
		}
	st_case_514:
//line scanner/scanner.go:16605
		switch lex.data[(lex.p)] {
		case 10:
			goto tr169
		case 13:
			goto st108
		case 32:
			goto tr775
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr775
		}
		goto tr778
	tr169:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st515
	tr782:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st515
	st515:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof515
		}
	st_case_515:
//line scanner/scanner.go:16635
		switch lex.data[(lex.p)] {
		case 10:
			goto tr782
		case 13:
			goto tr783
		case 32:
			goto tr781
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr781
		}
		goto tr780
	tr783:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st108
	st108:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof108
		}
	st_case_108:
//line scanner/scanner.go:16657
		if lex.data[(lex.p)] == 10 {
			goto tr169
		}
		goto tr168
	st516:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof516
		}
	st_case_516:
		if lex.data[(lex.p)] == 10 {
			goto tr169
		}
		goto tr784
	tr170:
//line scanner/scanner.rl:448
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st517
	tr785:
		lex.cs = 517
//line scanner/scanner.rl:450
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	tr788:
		lex.cs = 517
//line scanner/scanner.rl:449
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int(')'))
			lex.cs = 521
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr789:
//line scanner/scanner.rl:448
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st517
	tr791:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:448
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st517
	tr795:
		lex.cs = 517
//line scanner/scanner.rl:450
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	st517:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof517
		}
	st_case_517:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16720
		switch lex.data[(lex.p)] {
		case 10:
			goto tr171
		case 13:
			goto st520
		case 32:
			goto tr786
		case 41:
			goto tr788
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr786
		}
		goto tr785
	tr786:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st518
	tr792:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st518
	st518:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof518
		}
	st_case_518:
//line scanner/scanner.go:16752
		switch lex.data[(lex.p)] {
		case 10:
			goto tr171
		case 13:
			goto st109
		case 32:
			goto tr786
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr786
		}
		goto tr789
	tr171:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st519
	tr793:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st519
	st519:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof519
		}
	st_case_519:
//line scanner/scanner.go:16782
		switch lex.data[(lex.p)] {
		case 10:
			goto tr793
		case 13:
			goto tr794
		case 32:
			goto tr792
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr792
		}
		goto tr791
	tr794:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st109
	st109:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof109
		}
	st_case_109:
//line scanner/scanner.go:16804
		if lex.data[(lex.p)] == 10 {
			goto tr171
		}
		goto tr170
	st520:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof520
		}
	st_case_520:
		if lex.data[(lex.p)] == 10 {
			goto tr171
		}
		goto tr795
	tr172:
//line scanner/scanner.rl:454
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st521
	tr796:
		lex.cs = 521
//line scanner/scanner.rl:456
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	tr799:
		lex.cs = 521
//line scanner/scanner.rl:455
		lex.te = (lex.p) + 1
		{
			lex.createToken(lval)
			tok = TokenID(int(';'))
			lex.cs = 525
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr800:
//line scanner/scanner.rl:454
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st521
	tr802:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:454
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st521
	tr806:
		lex.cs = 521
//line scanner/scanner.rl:456
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 118
		}
		goto _again
	st521:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof521
		}
	st_case_521:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16867
		switch lex.data[(lex.p)] {
		case 10:
			goto tr173
		case 13:
			goto st524
		case 32:
			goto tr797
		case 59:
			goto tr799
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr797
		}
		goto tr796
	tr797:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st522
	tr803:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st522
	st522:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof522
		}
	st_case_522:
//line scanner/scanner.go:16899
		switch lex.data[(lex.p)] {
		case 10:
			goto tr173
		case 13:
			goto st110
		case 32:
			goto tr797
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr797
		}
		goto tr800
	tr173:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st523
	tr804:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st523
	st523:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof523
		}
	st_case_523:
//line scanner/scanner.go:16929
		switch lex.data[(lex.p)] {
		case 10:
			goto tr804
		case 13:
			goto tr805
		case 32:
			goto tr803
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr803
		}
		goto tr802
	tr805:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st110
	st110:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof110
		}
	st_case_110:
//line scanner/scanner.go:16951
		if lex.data[(lex.p)] == 10 {
			goto tr173
		}
		goto tr172
	st524:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof524
		}
	st_case_524:
		if lex.data[(lex.p)] == 10 {
			goto tr173
		}
		goto tr806
	tr809:
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

		goto st525
	tr810:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:460
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
		}
		goto st525
	st525:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof525
		}
	st_case_525:
//line NONE:1
		lex.ts = (lex.p)

//line scanner/scanner.go:16998
		if lex.data[(lex.p)] == 10 {
			goto st527
		}
		goto tr807
	tr807:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:460
		lex.act = 182
		goto st526
	tr811:
//line NONE:1
		lex.te = (lex.p) + 1

//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
//line scanner/scanner.rl:460
		lex.act = 182
		goto st526
	st526:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof526
		}
	st_case_526:
//line scanner/scanner.go:17024
		if lex.data[(lex.p)] == 10 {
			goto st527
		}
		goto tr807
	tr812:
//line scanner/scanner.rl:50
		lex.NewLines.Append(lex.p)
		goto st527
	st527:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof527
		}
	st_case_527:
//line scanner/scanner.go:17038
		if lex.data[(lex.p)] == 10 {
			goto tr812
		}
		goto tr811
	st_out:
	_test_eof111:
		lex.cs = 111
		goto _test_eof
	_test_eof112:
		lex.cs = 112
		goto _test_eof
	_test_eof113:
		lex.cs = 113
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
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof117:
		lex.cs = 117
		goto _test_eof
	_test_eof4:
		lex.cs = 4
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
	_test_eof5:
		lex.cs = 5
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
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof11:
		lex.cs = 11
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
	_test_eof128:
		lex.cs = 128
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
	_test_eof132:
		lex.cs = 132
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
	_test_eof66:
		lex.cs = 66
		goto _test_eof
	_test_eof67:
		lex.cs = 67
		goto _test_eof
	_test_eof68:
		lex.cs = 68
		goto _test_eof
	_test_eof69:
		lex.cs = 69
		goto _test_eof
	_test_eof133:
		lex.cs = 133
		goto _test_eof
	_test_eof134:
		lex.cs = 134
		goto _test_eof
	_test_eof135:
		lex.cs = 135
		goto _test_eof
	_test_eof136:
		lex.cs = 136
		goto _test_eof
	_test_eof137:
		lex.cs = 137
		goto _test_eof
	_test_eof70:
		lex.cs = 70
		goto _test_eof
	_test_eof138:
		lex.cs = 138
		goto _test_eof
	_test_eof71:
		lex.cs = 71
		goto _test_eof
	_test_eof72:
		lex.cs = 72
		goto _test_eof
	_test_eof139:
		lex.cs = 139
		goto _test_eof
	_test_eof140:
		lex.cs = 140
		goto _test_eof
	_test_eof73:
		lex.cs = 73
		goto _test_eof
	_test_eof74:
		lex.cs = 74
		goto _test_eof
	_test_eof75:
		lex.cs = 75
		goto _test_eof
	_test_eof141:
		lex.cs = 141
		goto _test_eof
	_test_eof142:
		lex.cs = 142
		goto _test_eof
	_test_eof76:
		lex.cs = 76
		goto _test_eof
	_test_eof143:
		lex.cs = 143
		goto _test_eof
	_test_eof77:
		lex.cs = 77
		goto _test_eof
	_test_eof144:
		lex.cs = 144
		goto _test_eof
	_test_eof145:
		lex.cs = 145
		goto _test_eof
	_test_eof146:
		lex.cs = 146
		goto _test_eof
	_test_eof147:
		lex.cs = 147
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
	_test_eof78:
		lex.cs = 78
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
	_test_eof151:
		lex.cs = 151
		goto _test_eof
	_test_eof152:
		lex.cs = 152
		goto _test_eof
	_test_eof82:
		lex.cs = 82
		goto _test_eof
	_test_eof153:
		lex.cs = 153
		goto _test_eof
	_test_eof154:
		lex.cs = 154
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
	_test_eof155:
		lex.cs = 155
		goto _test_eof
	_test_eof87:
		lex.cs = 87
		goto _test_eof
	_test_eof88:
		lex.cs = 88
		goto _test_eof
	_test_eof89:
		lex.cs = 89
		goto _test_eof
	_test_eof90:
		lex.cs = 90
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
	_test_eof91:
		lex.cs = 91
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
	_test_eof92:
		lex.cs = 92
		goto _test_eof
	_test_eof93:
		lex.cs = 93
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
	_test_eof94:
		lex.cs = 94
		goto _test_eof
	_test_eof95:
		lex.cs = 95
		goto _test_eof
	_test_eof96:
		lex.cs = 96
		goto _test_eof
	_test_eof97:
		lex.cs = 97
		goto _test_eof
	_test_eof98:
		lex.cs = 98
		goto _test_eof
	_test_eof99:
		lex.cs = 99
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
	_test_eof100:
		lex.cs = 100
		goto _test_eof
	_test_eof469:
		lex.cs = 469
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
	_test_eof476:
		lex.cs = 476
		goto _test_eof
	_test_eof101:
		lex.cs = 101
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
	_test_eof482:
		lex.cs = 482
		goto _test_eof
	_test_eof102:
		lex.cs = 102
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
	_test_eof103:
		lex.cs = 103
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
	_test_eof104:
		lex.cs = 104
		goto _test_eof
	_test_eof499:
		lex.cs = 499
		goto _test_eof
	_test_eof500:
		lex.cs = 500
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
	_test_eof505:
		lex.cs = 505
		goto _test_eof
	_test_eof506:
		lex.cs = 506
		goto _test_eof
	_test_eof105:
		lex.cs = 105
		goto _test_eof
	_test_eof507:
		lex.cs = 507
		goto _test_eof
	_test_eof106:
		lex.cs = 106
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
	_test_eof512:
		lex.cs = 512
		goto _test_eof
	_test_eof107:
		lex.cs = 107
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
	_test_eof108:
		lex.cs = 108
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
	_test_eof109:
		lex.cs = 109
		goto _test_eof
	_test_eof520:
		lex.cs = 520
		goto _test_eof
	_test_eof521:
		lex.cs = 521
		goto _test_eof
	_test_eof522:
		lex.cs = 522
		goto _test_eof
	_test_eof523:
		lex.cs = 523
		goto _test_eof
	_test_eof110:
		lex.cs = 110
		goto _test_eof
	_test_eof524:
		lex.cs = 524
		goto _test_eof
	_test_eof525:
		lex.cs = 525
		goto _test_eof
	_test_eof526:
		lex.cs = 526
		goto _test_eof
	_test_eof527:
		lex.cs = 527
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 112:
				goto tr177
			case 113:
				goto tr179
			case 114:
				goto tr177
			case 115:
				goto tr177
			case 116:
				goto tr184
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr0
			case 117:
				goto tr187
			case 4:
				goto tr0
			case 119:
				goto tr242
			case 120:
				goto tr244
			case 5:
				goto tr6
			case 121:
				goto tr248
			case 122:
				goto tr249
			case 123:
				goto tr251
			case 124:
				goto tr253
			case 6:
				goto tr8
			case 7:
				goto tr8
			case 8:
				goto tr8
			case 9:
				goto tr8
			case 10:
				goto tr8
			case 11:
				goto tr8
			case 125:
				goto tr254
			case 126:
				goto tr256
			case 127:
				goto tr249
			case 128:
				goto tr260
			case 129:
				goto tr249
			case 130:
				goto tr249
			case 131:
				goto tr248
			case 12:
				goto tr26
			case 13:
				goto tr26
			case 14:
				goto tr26
			case 15:
				goto tr26
			case 132:
				goto tr249
			case 16:
				goto tr38
			case 17:
				goto tr38
			case 18:
				goto tr38
			case 19:
				goto tr38
			case 20:
				goto tr38
			case 21:
				goto tr38
			case 22:
				goto tr38
			case 23:
				goto tr38
			case 24:
				goto tr38
			case 25:
				goto tr38
			case 26:
				goto tr38
			case 27:
				goto tr38
			case 28:
				goto tr38
			case 29:
				goto tr38
			case 30:
				goto tr38
			case 31:
				goto tr38
			case 32:
				goto tr38
			case 33:
				goto tr38
			case 34:
				goto tr38
			case 35:
				goto tr38
			case 36:
				goto tr38
			case 37:
				goto tr38
			case 38:
				goto tr38
			case 39:
				goto tr38
			case 40:
				goto tr38
			case 41:
				goto tr38
			case 42:
				goto tr38
			case 43:
				goto tr38
			case 44:
				goto tr38
			case 45:
				goto tr38
			case 46:
				goto tr38
			case 47:
				goto tr38
			case 48:
				goto tr38
			case 49:
				goto tr38
			case 50:
				goto tr38
			case 51:
				goto tr38
			case 52:
				goto tr38
			case 53:
				goto tr38
			case 54:
				goto tr38
			case 55:
				goto tr38
			case 56:
				goto tr38
			case 57:
				goto tr38
			case 58:
				goto tr38
			case 59:
				goto tr38
			case 60:
				goto tr38
			case 61:
				goto tr38
			case 62:
				goto tr38
			case 63:
				goto tr38
			case 64:
				goto tr38
			case 65:
				goto tr38
			case 66:
				goto tr38
			case 67:
				goto tr38
			case 68:
				goto tr38
			case 69:
				goto tr38
			case 133:
				goto tr249
			case 134:
				goto tr266
			case 135:
				goto tr249
			case 136:
				goto tr249
			case 137:
				goto tr249
			case 70:
				goto tr38
			case 138:
				goto tr276
			case 71:
				goto tr8
			case 72:
				goto tr8
			case 139:
				goto tr276
			case 140:
				goto tr249
			case 73:
				goto tr38
			case 74:
				goto tr38
			case 75:
				goto tr38
			case 141:
				goto tr279
			case 142:
				goto tr279
			case 76:
				goto tr110
			case 143:
				goto tr282
			case 77:
				goto tr110
			case 144:
				goto tr283
			case 145:
				goto tr279
			case 146:
				goto tr8
			case 147:
				goto tr285
			case 148:
				goto tr276
			case 149:
				goto tr249
			case 150:
				goto tr249
			case 78:
				goto tr38
			case 79:
				goto tr38
			case 80:
				goto tr38
			case 81:
				goto tr38
			case 151:
				goto tr288
			case 152:
				goto tr290
			case 82:
				goto tr122
			case 153:
				goto tr249
			case 154:
				goto tr294
			case 83:
				goto tr8
			case 84:
				goto tr8
			case 85:
				goto tr8
			case 86:
				goto tr8
			case 155:
				goto tr296
			case 87:
				goto tr8
			case 88:
				goto tr8
			case 89:
				goto tr8
			case 90:
				goto tr8
			case 156:
				goto tr297
			case 157:
				goto tr249
			case 158:
				goto tr301
			case 159:
				goto tr249
			case 160:
				goto tr305
			case 161:
				goto tr249
			case 162:
				goto tr309
			case 163:
				goto tr311
			case 91:
				goto tr138
			case 164:
				goto tr285
			case 165:
				goto tr285
			case 166:
				goto tr285
			case 167:
				goto tr285
			case 168:
				goto tr285
			case 169:
				goto tr285
			case 170:
				goto tr285
			case 171:
				goto tr285
			case 172:
				goto tr285
			case 173:
				goto tr285
			case 174:
				goto tr285
			case 175:
				goto tr285
			case 92:
				goto tr140
			case 93:
				goto tr140
			case 176:
				goto tr285
			case 177:
				goto tr285
			case 178:
				goto tr285
			case 179:
				goto tr285
			case 180:
				goto tr285
			case 181:
				goto tr285
			case 182:
				goto tr285
			case 183:
				goto tr285
			case 184:
				goto tr285
			case 185:
				goto tr285
			case 186:
				goto tr285
			case 187:
				goto tr285
			case 188:
				goto tr285
			case 189:
				goto tr285
			case 190:
				goto tr285
			case 191:
				goto tr285
			case 192:
				goto tr285
			case 193:
				goto tr285
			case 194:
				goto tr285
			case 195:
				goto tr285
			case 196:
				goto tr285
			case 197:
				goto tr285
			case 198:
				goto tr285
			case 199:
				goto tr285
			case 200:
				goto tr285
			case 201:
				goto tr285
			case 202:
				goto tr285
			case 203:
				goto tr285
			case 204:
				goto tr285
			case 205:
				goto tr285
			case 206:
				goto tr285
			case 207:
				goto tr285
			case 208:
				goto tr285
			case 209:
				goto tr285
			case 210:
				goto tr285
			case 211:
				goto tr285
			case 212:
				goto tr285
			case 213:
				goto tr285
			case 214:
				goto tr285
			case 215:
				goto tr285
			case 216:
				goto tr285
			case 217:
				goto tr285
			case 218:
				goto tr285
			case 219:
				goto tr285
			case 220:
				goto tr285
			case 221:
				goto tr285
			case 222:
				goto tr285
			case 223:
				goto tr285
			case 224:
				goto tr391
			case 225:
				goto tr285
			case 226:
				goto tr285
			case 227:
				goto tr285
			case 228:
				goto tr285
			case 229:
				goto tr285
			case 230:
				goto tr285
			case 231:
				goto tr285
			case 232:
				goto tr285
			case 233:
				goto tr285
			case 234:
				goto tr285
			case 235:
				goto tr285
			case 236:
				goto tr285
			case 237:
				goto tr285
			case 238:
				goto tr285
			case 239:
				goto tr411
			case 240:
				goto tr285
			case 241:
				goto tr285
			case 242:
				goto tr285
			case 243:
				goto tr285
			case 244:
				goto tr285
			case 245:
				goto tr285
			case 246:
				goto tr285
			case 247:
				goto tr285
			case 248:
				goto tr285
			case 249:
				goto tr285
			case 250:
				goto tr285
			case 251:
				goto tr285
			case 252:
				goto tr285
			case 253:
				goto tr285
			case 254:
				goto tr285
			case 255:
				goto tr285
			case 256:
				goto tr285
			case 257:
				goto tr285
			case 258:
				goto tr285
			case 259:
				goto tr285
			case 260:
				goto tr285
			case 261:
				goto tr285
			case 262:
				goto tr285
			case 263:
				goto tr285
			case 264:
				goto tr285
			case 265:
				goto tr439
			case 266:
				goto tr285
			case 267:
				goto tr285
			case 268:
				goto tr443
			case 269:
				goto tr285
			case 270:
				goto tr285
			case 271:
				goto tr285
			case 272:
				goto tr285
			case 273:
				goto tr285
			case 274:
				goto tr285
			case 275:
				goto tr285
			case 276:
				goto tr285
			case 277:
				goto tr285
			case 278:
				goto tr285
			case 279:
				goto tr285
			case 280:
				goto tr285
			case 281:
				goto tr285
			case 282:
				goto tr285
			case 283:
				goto tr285
			case 284:
				goto tr285
			case 285:
				goto tr285
			case 286:
				goto tr285
			case 287:
				goto tr285
			case 288:
				goto tr285
			case 289:
				goto tr285
			case 290:
				goto tr285
			case 291:
				goto tr285
			case 292:
				goto tr285
			case 293:
				goto tr475
			case 294:
				goto tr285
			case 295:
				goto tr285
			case 296:
				goto tr285
			case 297:
				goto tr285
			case 298:
				goto tr285
			case 299:
				goto tr285
			case 300:
				goto tr285
			case 301:
				goto tr285
			case 302:
				goto tr285
			case 303:
				goto tr285
			case 304:
				goto tr285
			case 305:
				goto tr285
			case 306:
				goto tr285
			case 307:
				goto tr285
			case 308:
				goto tr285
			case 309:
				goto tr285
			case 310:
				goto tr285
			case 311:
				goto tr285
			case 312:
				goto tr285
			case 313:
				goto tr285
			case 314:
				goto tr285
			case 315:
				goto tr285
			case 316:
				goto tr285
			case 317:
				goto tr285
			case 318:
				goto tr285
			case 319:
				goto tr285
			case 320:
				goto tr285
			case 321:
				goto tr285
			case 322:
				goto tr285
			case 323:
				goto tr285
			case 324:
				goto tr285
			case 325:
				goto tr285
			case 326:
				goto tr285
			case 327:
				goto tr285
			case 328:
				goto tr285
			case 329:
				goto tr285
			case 330:
				goto tr285
			case 331:
				goto tr285
			case 332:
				goto tr285
			case 333:
				goto tr285
			case 334:
				goto tr285
			case 335:
				goto tr285
			case 336:
				goto tr285
			case 337:
				goto tr285
			case 338:
				goto tr285
			case 339:
				goto tr285
			case 340:
				goto tr285
			case 341:
				goto tr285
			case 342:
				goto tr285
			case 343:
				goto tr285
			case 344:
				goto tr285
			case 345:
				goto tr285
			case 346:
				goto tr285
			case 347:
				goto tr285
			case 348:
				goto tr285
			case 349:
				goto tr285
			case 350:
				goto tr285
			case 351:
				goto tr285
			case 352:
				goto tr285
			case 353:
				goto tr285
			case 354:
				goto tr543
			case 355:
				goto tr285
			case 356:
				goto tr285
			case 357:
				goto tr285
			case 358:
				goto tr285
			case 359:
				goto tr285
			case 360:
				goto tr285
			case 361:
				goto tr285
			case 362:
				goto tr285
			case 363:
				goto tr285
			case 364:
				goto tr285
			case 365:
				goto tr285
			case 366:
				goto tr285
			case 367:
				goto tr285
			case 368:
				goto tr285
			case 369:
				goto tr285
			case 370:
				goto tr285
			case 371:
				goto tr285
			case 372:
				goto tr285
			case 373:
				goto tr285
			case 374:
				goto tr285
			case 375:
				goto tr285
			case 376:
				goto tr285
			case 377:
				goto tr285
			case 378:
				goto tr285
			case 379:
				goto tr285
			case 380:
				goto tr285
			case 381:
				goto tr285
			case 382:
				goto tr285
			case 383:
				goto tr285
			case 384:
				goto tr285
			case 385:
				goto tr285
			case 386:
				goto tr285
			case 387:
				goto tr285
			case 388:
				goto tr285
			case 389:
				goto tr285
			case 390:
				goto tr285
			case 391:
				goto tr285
			case 392:
				goto tr285
			case 393:
				goto tr285
			case 394:
				goto tr285
			case 395:
				goto tr589
			case 94:
				goto tr142
			case 95:
				goto tr142
			case 96:
				goto tr142
			case 97:
				goto tr142
			case 98:
				goto tr142
			case 99:
				goto tr142
			case 396:
				goto tr285
			case 397:
				goto tr285
			case 398:
				goto tr285
			case 399:
				goto tr249
			case 400:
				goto tr285
			case 401:
				goto tr285
			case 402:
				goto tr285
			case 403:
				goto tr285
			case 404:
				goto tr285
			case 405:
				goto tr285
			case 406:
				goto tr285
			case 407:
				goto tr285
			case 408:
				goto tr285
			case 409:
				goto tr285
			case 410:
				goto tr285
			case 411:
				goto tr285
			case 412:
				goto tr285
			case 413:
				goto tr285
			case 414:
				goto tr285
			case 415:
				goto tr285
			case 416:
				goto tr285
			case 417:
				goto tr285
			case 418:
				goto tr285
			case 419:
				goto tr285
			case 420:
				goto tr285
			case 421:
				goto tr285
			case 422:
				goto tr285
			case 423:
				goto tr285
			case 424:
				goto tr285
			case 425:
				goto tr285
			case 426:
				goto tr285
			case 427:
				goto tr285
			case 428:
				goto tr285
			case 429:
				goto tr285
			case 430:
				goto tr285
			case 431:
				goto tr285
			case 432:
				goto tr285
			case 433:
				goto tr285
			case 434:
				goto tr285
			case 435:
				goto tr285
			case 436:
				goto tr285
			case 437:
				goto tr285
			case 438:
				goto tr285
			case 439:
				goto tr285
			case 440:
				goto tr285
			case 441:
				goto tr285
			case 442:
				goto tr285
			case 443:
				goto tr285
			case 444:
				goto tr285
			case 445:
				goto tr285
			case 446:
				goto tr285
			case 447:
				goto tr285
			case 448:
				goto tr285
			case 449:
				goto tr285
			case 450:
				goto tr285
			case 451:
				goto tr285
			case 452:
				goto tr285
			case 453:
				goto tr285
			case 454:
				goto tr285
			case 455:
				goto tr285
			case 456:
				goto tr285
			case 457:
				goto tr285
			case 458:
				goto tr285
			case 459:
				goto tr285
			case 460:
				goto tr285
			case 461:
				goto tr285
			case 462:
				goto tr285
			case 463:
				goto tr285
			case 464:
				goto tr285
			case 465:
				goto tr249
			case 467:
				goto tr675
			case 468:
				goto tr677
			case 100:
				goto tr154
			case 469:
				goto tr681
			case 470:
				goto tr681
			case 471:
				goto tr683
			case 473:
				goto tr686
			case 474:
				goto tr687
			case 476:
				goto tr696
			case 477:
				goto tr698
			case 478:
				goto tr699
			case 479:
				goto tr696
			case 480:
				goto tr703
			case 482:
				goto tr713
			case 483:
				goto tr715
			case 484:
				goto tr716
			case 485:
				goto tr713
			case 486:
				goto tr720
			case 488:
				goto tr730
			case 489:
				goto tr732
			case 490:
				goto tr733
			case 491:
				goto tr730
			case 492:
				goto tr737
			case 494:
				goto tr740
			case 496:
				goto tr746
			case 497:
				goto tr748
			case 498:
				goto tr746
			case 104:
				goto tr160
			case 499:
				goto tr750
			case 501:
				goto tr762
			case 502:
				goto tr763
			case 503:
				goto tr764
			case 504:
				goto tr766
			case 505:
				goto tr767
			case 506:
				goto tr767
			case 105:
				goto tr162
			case 507:
				goto tr767
			case 106:
				goto tr162
			case 508:
				goto tr767
			case 509:
				goto tr767
			case 510:
				goto tr770
			case 512:
				goto tr773
			case 107:
				goto tr165
			case 514:
				goto tr778
			case 515:
				goto tr780
			case 108:
				goto tr168
			case 516:
				goto tr784
			case 518:
				goto tr789
			case 519:
				goto tr791
			case 109:
				goto tr170
			case 520:
				goto tr795
			case 522:
				goto tr800
			case 523:
				goto tr802
			case 110:
				goto tr172
			case 524:
				goto tr806
			case 526:
				goto tr809
			case 527:
				goto tr810
			}
		}

	_out:
		{
		}
	}

//line scanner/scanner.rl:464

	// always return same $end token
	if tok == 0 {
		if lex.lastToken == nil {
			lex.ts, lex.te = 0, 0
			lex.lastToken = lex.createToken(lval)
		}
		lval.Token(lex.lastToken)
	}

	return int(tok)
}
