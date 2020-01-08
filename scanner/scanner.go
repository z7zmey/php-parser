// line scanner/scanner.rl:1
package scanner

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/z7zmey/php-parser/freefloating"
)

// line scanner/scanner.go:15
const lexer_start int = 111
const lexer_first_final int = 111
const lexer_error int = 0

const lexer_en_main int = 111
const lexer_en_html int = 114
const lexer_en_php int = 121
const lexer_en_property int = 468
const lexer_en_nowdoc int = 474
const lexer_en_heredoc int = 477
const lexer_en_backqote int = 483
const lexer_en_template_string int = 489
const lexer_en_heredoc_end int = 495
const lexer_en_string_var int = 497
const lexer_en_string_var_index int = 502
const lexer_en_string_var_name int = 512
const lexer_en_halt_compiller_open_parenthesis int = 514
const lexer_en_halt_compiller_close_parenthesis int = 518
const lexer_en_halt_compiller_close_semicolon int = 522
const lexer_en_halt_compiller_end int = 526

// line scanner/scanner.rl:17

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{
		data:  data,
		pe:    len(data),
		stack: make([]int, 0),

		TokenPool: &TokenPool{},
		NewLines:  NewLines{make([]int, 0, 128)},
	}

// line scanner/scanner.go:51
	{
		lex.cs = lexer_start
		lex.top = 0
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

// line scanner/scanner.rl:29
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

// line scanner/scanner.go:79
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
		case 1:
			goto st1
		case 113:
			goto st113
		case 114:
			goto st114
		case 115:
			goto st115
		case 116:
			goto st116
		case 117:
			goto st117
		case 118:
			goto st118
		case 119:
			goto st119
		case 2:
			goto st2
		case 3:
			goto st3
		case 4:
			goto st4
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
		case 6:
			goto st6
		case 124:
			goto st124
		case 125:
			goto st125
		case 126:
			goto st126
		case 127:
			goto st127
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 128:
			goto st128
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
		case 134:
			goto st134
		case 11:
			goto st11
		case 12:
			goto st12
		case 135:
			goto st135
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
		case 66:
			goto st66
		case 136:
			goto st136
		case 137:
			goto st137
		case 138:
			goto st138
		case 139:
			goto st139
		case 140:
			goto st140
		case 67:
			goto st67
		case 141:
			goto st141
		case 68:
			goto st68
		case 69:
			goto st69
		case 142:
			goto st142
		case 70:
			goto st70
		case 143:
			goto st143
		case 71:
			goto st71
		case 72:
			goto st72
		case 73:
			goto st73
		case 144:
			goto st144
		case 145:
			goto st145
		case 146:
			goto st146
		case 74:
			goto st74
		case 75:
			goto st75
		case 147:
			goto st147
		case 76:
			goto st76
		case 148:
			goto st148
		case 149:
			goto st149
		case 150:
			goto st150
		case 77:
			goto st77
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 151:
			goto st151
		case 152:
			goto st152
		case 81:
			goto st81
		case 153:
			goto st153
		case 154:
			goto st154
		case 82:
			goto st82
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 155:
			goto st155
		case 86:
			goto st86
		case 87:
			goto st87
		case 88:
			goto st88
		case 89:
			goto st89
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
		case 90:
			goto st90
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
		case 176:
			goto st176
		case 177:
			goto st177
		case 91:
			goto st91
		case 92:
			goto st92
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
		case 396:
			goto st396
		case 397:
			goto st397
		case 93:
			goto st93
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
		case 469:
			goto st469
		case 470:
			goto st470
		case 99:
			goto st99
		case 471:
			goto st471
		case 472:
			goto st472
		case 473:
			goto st473
		case 474:
			goto st474
		case 0:
			goto st0
		case 475:
			goto st475
		case 476:
			goto st476
		case 477:
			goto st477
		case 478:
			goto st478
		case 100:
			goto st100
		case 479:
			goto st479
		case 480:
			goto st480
		case 481:
			goto st481
		case 482:
			goto st482
		case 483:
			goto st483
		case 484:
			goto st484
		case 101:
			goto st101
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
		case 102:
			goto st102
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
		case 499:
			goto st499
		case 500:
			goto st500
		case 103:
			goto st103
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
		case 507:
			goto st507
		case 508:
			goto st508
		case 104:
			goto st104
		case 105:
			goto st105
		case 509:
			goto st509
		case 106:
			goto st106
		case 510:
			goto st510
		case 511:
			goto st511
		case 512:
			goto st512
		case 513:
			goto st513
		case 107:
			goto st107
		case 514:
			goto st514
		case 515:
			goto st515
		case 516:
			goto st516
		case 108:
			goto st108
		case 517:
			goto st517
		case 518:
			goto st518
		case 519:
			goto st519
		case 520:
			goto st520
		case 109:
			goto st109
		case 521:
			goto st521
		case 522:
			goto st522
		case 523:
			goto st523
		case 524:
			goto st524
		case 110:
			goto st110
		case 525:
			goto st525
		case 526:
			goto st526
		case 527:
			goto st527
		case 528:
			goto st528
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
		case 1:
			goto st_case_1
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
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
		case 6:
			goto st_case_6
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 128:
			goto st_case_128
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
		case 134:
			goto st_case_134
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 135:
			goto st_case_135
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
		case 66:
			goto st_case_66
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 67:
			goto st_case_67
		case 141:
			goto st_case_141
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 142:
			goto st_case_142
		case 70:
			goto st_case_70
		case 143:
			goto st_case_143
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 147:
			goto st_case_147
		case 76:
			goto st_case_76
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 81:
			goto st_case_81
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 155:
			goto st_case_155
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
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
		case 90:
			goto st_case_90
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
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
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
		case 396:
			goto st_case_396
		case 397:
			goto st_case_397
		case 93:
			goto st_case_93
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
		case 469:
			goto st_case_469
		case 470:
			goto st_case_470
		case 99:
			goto st_case_99
		case 471:
			goto st_case_471
		case 472:
			goto st_case_472
		case 473:
			goto st_case_473
		case 474:
			goto st_case_474
		case 0:
			goto st_case_0
		case 475:
			goto st_case_475
		case 476:
			goto st_case_476
		case 477:
			goto st_case_477
		case 478:
			goto st_case_478
		case 100:
			goto st_case_100
		case 479:
			goto st_case_479
		case 480:
			goto st_case_480
		case 481:
			goto st_case_481
		case 482:
			goto st_case_482
		case 483:
			goto st_case_483
		case 484:
			goto st_case_484
		case 101:
			goto st_case_101
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
		case 102:
			goto st_case_102
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
		case 499:
			goto st_case_499
		case 500:
			goto st_case_500
		case 103:
			goto st_case_103
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
		case 507:
			goto st_case_507
		case 508:
			goto st_case_508
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 509:
			goto st_case_509
		case 106:
			goto st_case_106
		case 510:
			goto st_case_510
		case 511:
			goto st_case_511
		case 512:
			goto st_case_512
		case 513:
			goto st_case_513
		case 107:
			goto st_case_107
		case 514:
			goto st_case_514
		case 515:
			goto st_case_515
		case 516:
			goto st_case_516
		case 108:
			goto st_case_108
		case 517:
			goto st_case_517
		case 518:
			goto st_case_518
		case 519:
			goto st_case_519
		case 520:
			goto st_case_520
		case 109:
			goto st_case_109
		case 521:
			goto st_case_521
		case 522:
			goto st_case_522
		case 523:
			goto st_case_523
		case 524:
			goto st_case_524
		case 110:
			goto st_case_110
		case 525:
			goto st_case_525
		case 526:
			goto st_case_526
		case 527:
			goto st_case_527
		case 528:
			goto st_case_528
		}
		goto st_out
	tr0:
		lex.cs = 111
// line scanner/scanner.rl:141
		(lex.p) = (lex.te) - 1
		{
			lex.cs = 114
			lex.ungetCnt(1)
		}
		goto _again
	tr162:
		lex.cs = 111
// line scanner/scanner.rl:141
		lex.te = (lex.p) + 1
		{
			lex.cs = 114
			lex.ungetCnt(1)
		}
		goto _again
	tr164:
		lex.cs = 111
// line scanner/scanner.rl:141
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.cs = 114
			lex.ungetCnt(1)
		}
		goto _again
	tr165:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:138
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
		}
		goto st111
	st111:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof111
		}
	st_case_111:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:2263
		if lex.data[(lex.p)] == 35 {
			goto tr163
		}
		goto tr162
	tr163:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st112
	st112:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof112
		}
	st_case_112:
// line scanner/scanner.go:2278
		if lex.data[(lex.p)] == 33 {
			goto st1
		}
		goto tr164
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if lex.data[(lex.p)] == 10 {
			goto st113
		}
		goto st1
	st113:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		goto tr165
	tr3:
		lex.cs = 114
// line scanner/scanner.rl:154
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			lex.cs = 121
		}
		goto _again
	tr6:
		lex.cs = 114
// line scanner/scanner.rl:158
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.ts+5)
			lex.cs = 121
		}
		goto _again
	tr169:
// line scanner/scanner.rl:148
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(token)
			tok = T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr171:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:148
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(token)
			tok = T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 114
				goto _out
			}
		}
		goto st114
	tr176:
		lex.cs = 114
// line scanner/scanner.rl:154
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			lex.cs = 121
		}
		goto _again
	tr177:
		lex.cs = 114
// line scanner/scanner.rl:163
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ECHO
			lex.cs = 121
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr179:
		lex.cs = 114
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:158
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.ts+5)
			lex.cs = 121
		}
		goto _again
	st114:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof114
		}
	st_case_114:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:2386
		switch lex.data[(lex.p)] {
		case 10:
			goto st116
		case 60:
			goto st118
		}
		goto st115
	tr172:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st115
	st115:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof115
		}
	st_case_115:
// line scanner/scanner.go:2403
		switch lex.data[(lex.p)] {
		case 10:
			goto st116
		case 60:
			goto st117
		}
		goto st115
	tr173:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st116
	st116:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof116
		}
	st_case_116:
// line scanner/scanner.go:2420
		switch lex.data[(lex.p)] {
		case 10:
			goto tr173
		case 60:
			goto tr174
		}
		goto tr172
	tr174:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st117
	st117:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof117
		}
	st_case_117:
// line scanner/scanner.go:2437
		switch lex.data[(lex.p)] {
		case 10:
			goto st116
		case 60:
			goto st117
		case 63:
			goto tr169
		}
		goto st115
	st118:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof118
		}
	st_case_118:
		switch lex.data[(lex.p)] {
		case 10:
			goto st116
		case 60:
			goto st117
		case 63:
			goto tr175
		}
		goto st115
	tr175:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st119
	st119:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof119
		}
	st_case_119:
// line scanner/scanner.go:2471
		switch lex.data[(lex.p)] {
		case 61:
			goto tr177
		case 80:
			goto st2
		case 112:
			goto st2
		}
		goto tr176
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 72:
			goto st3
		case 104:
			goto st3
		}
		goto tr3
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 80:
			goto st4
		case 112:
			goto st4
		}
		goto tr3
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 9:
			goto tr6
		case 10:
			goto st120
		case 13:
			goto st5
		case 32:
			goto tr6
		}
		goto tr3
	st120:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof120
		}
	st_case_120:
		goto tr179
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if lex.data[(lex.p)] == 10 {
			goto st120
		}
		goto tr3
	tr9:
// line scanner/scanner.rl:172
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st121
	tr11:
		lex.cs = 121
// line NONE:1
		switch lex.act {
		case 10:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 11:
			{
				(lex.p) = (lex.te) - 1

				s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, 2, 0)

				if err == nil {
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

				base := 10
				if lex.data[lex.ts] == '0' {
					base = 8
				}

				s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, base, 0)

				if err == nil {
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
		case 13:
			{
				(lex.p) = (lex.te) - 1

				s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, 16, 0)

				if err == nil {
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
		case 14:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ABSTRACT
				{
					(lex.p)++
					goto _out
				}
			}
		case 15:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ARRAY
				{
					(lex.p)++
					goto _out
				}
			}
		case 16:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_AS
				{
					(lex.p)++
					goto _out
				}
			}
		case 17:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_BREAK
				{
					(lex.p)++
					goto _out
				}
			}
		case 18:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CALLABLE
				{
					(lex.p)++
					goto _out
				}
			}
		case 19:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CASE
				{
					(lex.p)++
					goto _out
				}
			}
		case 20:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CATCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 21:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CLASS
				{
					(lex.p)++
					goto _out
				}
			}
		case 22:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CLONE
				{
					(lex.p)++
					goto _out
				}
			}
		case 23:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CONST
				{
					(lex.p)++
					goto _out
				}
			}
		case 24:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CONTINUE
				{
					(lex.p)++
					goto _out
				}
			}
		case 25:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 26:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DEFAULT
				{
					(lex.p)++
					goto _out
				}
			}
		case 27:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DO
				{
					(lex.p)++
					goto _out
				}
			}
		case 28:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ECHO
				{
					(lex.p)++
					goto _out
				}
			}
		case 30:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ELSEIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 31:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EMPTY
				{
					(lex.p)++
					goto _out
				}
			}
		case 32:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDDECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 34:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDFOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 35:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 36:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDSWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 37:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ENDWHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 38:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EVAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 39:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EXIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 40:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_EXTENDS
				{
					(lex.p)++
					goto _out
				}
			}
		case 42:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FINALLY
				{
					(lex.p)++
					goto _out
				}
			}
		case 44:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 45:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FUNCTION
				{
					(lex.p)++
					goto _out
				}
			}
		case 46:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FN
				{
					(lex.p)++
					goto _out
				}
			}
		case 47:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_GLOBAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 48:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_GOTO
				{
					(lex.p)++
					goto _out
				}
			}
		case 49:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_IF
				{
					(lex.p)++
					goto _out
				}
			}
		case 50:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_ISSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 51:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_IMPLEMENTS
				{
					(lex.p)++
					goto _out
				}
			}
		case 52:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INSTANCEOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 53:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INSTEADOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 54:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INTERFACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 55:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LIST
				{
					(lex.p)++
					goto _out
				}
			}
		case 56:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_NAMESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 57:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PRIVATE
				{
					(lex.p)++
					goto _out
				}
			}
		case 58:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PUBLIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 59:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PRINT
				{
					(lex.p)++
					goto _out
				}
			}
		case 60:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_PROTECTED
				{
					(lex.p)++
					goto _out
				}
			}
		case 61:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_RETURN
				{
					(lex.p)++
					goto _out
				}
			}
		case 62:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_STATIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 63:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_SWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 64:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_THROW
				{
					(lex.p)++
					goto _out
				}
			}
		case 65:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_TRAIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 66:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_TRY
				{
					(lex.p)++
					goto _out
				}
			}
		case 67:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_UNSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 68:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_USE
				{
					(lex.p)++
					goto _out
				}
			}
		case 69:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_VAR
				{
					(lex.p)++
					goto _out
				}
			}
		case 70:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_WHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 71:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_YIELD_FROM
				{
					(lex.p)++
					goto _out
				}
			}
		case 74:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_INCLUDE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 76:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_REQUIRE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 77:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_CLASS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 78:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DIR
				{
					(lex.p)++
					goto _out
				}
			}
		case 79:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 80:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_FUNC_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 81:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LINE
				{
					(lex.p)++
					goto _out
				}
			}
		case 82:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_NS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 83:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_METHOD_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 84:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_TRAIT_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 85:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_HALT_COMPILER
				lex.cs = 514
				{
					(lex.p)++
					goto _out
				}
			}
		case 86:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_NEW
				{
					(lex.p)++
					goto _out
				}
			}
		case 87:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LOGICAL_AND
				{
					(lex.p)++
					goto _out
				}
			}
		case 88:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LOGICAL_OR
				{
					(lex.p)++
					goto _out
				}
			}
		case 89:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_LOGICAL_XOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 118:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_SL
				{
					(lex.p)++
					goto _out
				}
			}
		case 135:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_STRING
				{
					(lex.p)++
					goto _out
				}
			}
		case 140:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = TokenID(int('"'))
				lex.cs = 489
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr14:
// line scanner/scanner.rl:361
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_CONSTANT_ENCAPSED_STRING
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr18:
// line scanner/scanner.rl:384
		(lex.p) = (lex.te) - 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st121
	tr22:
// line scanner/scanner.rl:346
		(lex.p) = (lex.te) - 1
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr37:
// line scanner/scanner.rl:322
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ARRAY_CAST
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr44:
// line scanner/scanner.rl:327
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_STRING_CAST
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr48:
// line scanner/scanner.rl:323
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_BOOL_CAST
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr56:
// line scanner/scanner.rl:324
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOUBLE_CAST
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr63:
// line scanner/scanner.rl:325
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_INT_CAST
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr72:
// line scanner/scanner.rl:326
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OBJECT_CAST
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr83:
// line scanner/scanner.rl:328
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_UNSET_CAST
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr84:
// line scanner/scanner.rl:290
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ELLIPSIS
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr87:
// line scanner/scanner.rl:176
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr95:
// line scanner/scanner.rl:334
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
		goto st121
	tr96:
// line scanner/scanner.rl:187
		(lex.p) = (lex.te) - 1
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 121
					goto _out
				}
			}

			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr109:
		lex.cs = 121
// line scanner/scanner.rl:174
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr125:
		lex.cs = 121
// line scanner/scanner.rl:173
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr127:
// line scanner/scanner.rl:357
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr129:
// line scanner/scanner.rl:271
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = T_YIELD
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr140:
// line scanner/scanner.rl:270
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_YIELD_FROM
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr180:
// line scanner/scanner.rl:384
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st121
	tr191:
// line scanner/scanner.rl:346
		lex.te = (lex.p) + 1
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr225:
// line scanner/scanner.rl:289
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_NS_SEPARATOR
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr228:
		lex.cs = 121
// line scanner/scanner.rl:381
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('`'))
			lex.cs = 483
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr229:
// line scanner/scanner.rl:354
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('{'))
			lex.call(121, 121)
			goto _out
		}
		goto st121
	tr231:
// line scanner/scanner.rl:355
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('}'))
			lex.ret(1)
			lex.PhpDocComment = ""
			goto _out
		}
		goto st121
	tr232:
// line scanner/scanner.rl:172
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st121
	tr234:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:172
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st121
	tr238:
// line scanner/scanner.rl:384
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st121
	tr239:
// line scanner/scanner.rl:346
		lex.te = (lex.p)
		(lex.p)--
		{
			// rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
			// tok = TokenID(Rune2Class(rune));
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr241:
// line scanner/scanner.rl:308
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr242:
// line scanner/scanner.rl:309
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_NOT_IDENTICAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr243:
		lex.cs = 121
// line scanner/scanner.rl:382
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('"'))
			lex.cs = 489
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr244:
// line scanner/scanner.rl:330
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
		}
		goto st121
	tr246:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:330
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
		}
		goto st121
	tr250:
// line scanner/scanner.rl:356
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr251:
// line scanner/scanner.rl:303
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_MOD_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr252:
// line scanner/scanner.rl:292
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_BOOLEAN_AND
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr253:
// line scanner/scanner.rl:294
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_AND_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr255:
// line scanner/scanner.rl:297
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_MUL_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr256:
// line scanner/scanner.rl:316
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_POW
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr257:
// line scanner/scanner.rl:298
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_POW_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr258:
// line scanner/scanner.rl:305
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_INC
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr259:
// line scanner/scanner.rl:300
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_PLUS_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr260:
// line scanner/scanner.rl:304
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DEC
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr261:
// line scanner/scanner.rl:301
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_MINUS_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr262:
		lex.cs = 121
// line scanner/scanner.rl:359
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OBJECT_OPERATOR
			lex.cs = 468
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr264:
// line scanner/scanner.rl:296
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_CONCAT_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr265:
// line scanner/scanner.rl:176
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr268:
// line scanner/scanner.rl:299
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DIV_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr269:
// line scanner/scanner.rl:187
		lex.te = (lex.p)
		(lex.p)--
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 121
					goto _out
				}
			}

			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr274:
// line scanner/scanner.rl:177
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, 2, 0)

			if err == nil {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 121
					goto _out
				}
			}

			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr275:
// line scanner/scanner.rl:202
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, 16, 0)

			if err == nil {
				lex.setTokenPosition(token)
				tok = T_LNUMBER
				{
					(lex.p)++
					lex.cs = 121
					goto _out
				}
			}

			lex.setTokenPosition(token)
			tok = T_DNUMBER
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr276:
// line scanner/scanner.rl:291
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_PAAMAYIM_NEKUDOTAYIM
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr277:
		lex.cs = 121
// line scanner/scanner.rl:174
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr279:
		lex.cs = 121
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:174
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr282:
// line scanner/scanner.rl:308
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr283:
// line scanner/scanner.rl:317
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_SL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr284:
// line scanner/scanner.rl:312
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_SL_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr285:
		lex.cs = 121
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:367
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.heredocLabel = lex.data[lblStart:lblEnd]
			lex.setTokenPosition(token)
			tok = T_START_HEREDOC

			if lex.isHeredocEnd(lex.p + 1) {
				lex.cs = 495
			} else if lex.data[lblStart-1] == '\'' {
				lex.cs = 474
			} else {
				lex.cs = 477
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr286:
// line scanner/scanner.rl:315
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_IS_SMALLER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr287:
// line scanner/scanner.rl:307
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_SPACESHIP
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr289:
// line scanner/scanner.rl:306
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOUBLE_ARROW
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr290:
// line scanner/scanner.rl:310
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_IS_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr291:
// line scanner/scanner.rl:311
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_IDENTICAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr292:
// line scanner/scanner.rl:314
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_IS_GREATER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr294:
// line scanner/scanner.rl:318
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_SR
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr295:
// line scanner/scanner.rl:313
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_SR_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr298:
		lex.cs = 121
// line scanner/scanner.rl:173
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr300:
		lex.cs = 121
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:173
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 114
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr301:
// line scanner/scanner.rl:319
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_COALESCE
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr302:
// line scanner/scanner.rl:320
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_COALESCE_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr303:
// line scanner/scanner.rl:357
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr383:
// line scanner/scanner.rl:228
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ELSE
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr403:
// line scanner/scanner.rl:232
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENDFOR
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr432:
// line scanner/scanner.rl:240
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_FINAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr436:
// line scanner/scanner.rl:242
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_FOR
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr468:
// line scanner/scanner.rl:272
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_INCLUDE
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr536:
// line scanner/scanner.rl:274
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_REQUIRE
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr582:
// line scanner/scanner.rl:271
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_YIELD
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr587:
// line scanner/scanner.rl:302
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_XOR_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr661:
// line scanner/scanner.rl:295
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	tr662:
// line scanner/scanner.rl:293
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_BOOLEAN_OR
			{
				(lex.p)++
				lex.cs = 121
				goto _out
			}
		}
		goto st121
	st121:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof121
		}
	st_case_121:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:3387
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto st124
		case 32:
			goto tr181
		case 33:
			goto st125
		case 34:
			goto tr184
		case 35:
			goto st128
		case 36:
			goto st130
		case 37:
			goto st132
		case 38:
			goto st133
		case 39:
			goto tr189
		case 40:
			goto tr190
		case 42:
			goto st136
		case 43:
			goto st138
		case 45:
			goto st139
		case 46:
			goto tr195
		case 47:
			goto tr196
		case 48:
			goto tr197
		case 58:
			goto st149
		case 59:
			goto tr199
		case 60:
			goto st153
		case 61:
			goto st157
		case 62:
			goto st159
		case 63:
			goto st161
		case 64:
			goto tr191
		case 65:
			goto st165
		case 66:
			goto tr205
		case 67:
			goto st181
		case 68:
			goto st210
		case 69:
			goto st221
		case 70:
			goto st263
		case 71:
			goto st274
		case 73:
			goto st281
		case 76:
			goto st320
		case 78:
			goto st323
		case 79:
			goto st332
		case 80:
			goto st333
		case 82:
			goto st350
		case 83:
			goto st364
		case 84:
			goto st373
		case 85:
			goto st380
		case 86:
			goto st385
		case 87:
			goto st387
		case 88:
			goto st391
		case 89:
			goto st393
		case 92:
			goto tr225
		case 94:
			goto st401
		case 95:
			goto st402
		case 96:
			goto tr228
		case 97:
			goto st165
		case 98:
			goto tr205
		case 99:
			goto st181
		case 100:
			goto st210
		case 101:
			goto st221
		case 102:
			goto st263
		case 103:
			goto st274
		case 105:
			goto st281
		case 108:
			goto st320
		case 110:
			goto st323
		case 111:
			goto st332
		case 112:
			goto st333
		case 114:
			goto st350
		case 115:
			goto st364
		case 116:
			goto st373
		case 117:
			goto st380
		case 118:
			goto st385
		case 119:
			goto st387
		case 120:
			goto st391
		case 121:
			goto st393
		case 123:
			goto tr229
		case 124:
			goto st467
		case 125:
			goto tr231
		case 126:
			goto tr191
		case 127:
			goto tr180
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr181
				}
			default:
				goto tr180
			}
		case lex.data[(lex.p)] > 31:
			switch {
			case lex.data[(lex.p)] < 49:
				if 41 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 44 {
					goto tr191
				}
			case lex.data[(lex.p)] > 57:
				if 91 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 93 {
					goto tr191
				}
			default:
				goto tr97
			}
		default:
			goto tr180
		}
		goto tr211
	tr181:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st122
	tr235:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st122
	st122:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof122
		}
	st_case_122:
// line scanner/scanner.go:3580
		switch lex.data[(lex.p)] {
		case 10:
			goto tr10
		case 13:
			goto st6
		case 32:
			goto tr181
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr181
		}
		goto tr232
	tr10:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st123
	tr236:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st123
	st123:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof123
		}
	st_case_123:
// line scanner/scanner.go:3610
		switch lex.data[(lex.p)] {
		case 10:
			goto tr236
		case 13:
			goto tr237
		case 32:
			goto tr235
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr235
		}
		goto tr234
	tr237:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
// line scanner/scanner.go:3632
		if lex.data[(lex.p)] == 10 {
			goto tr10
		}
		goto tr9
	st124:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		if lex.data[(lex.p)] == 10 {
			goto tr10
		}
		goto tr238
	st125:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof125
		}
	st_case_125:
		if lex.data[(lex.p)] == 61 {
			goto st126
		}
		goto tr239
	st126:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof126
		}
	st_case_126:
		if lex.data[(lex.p)] == 61 {
			goto tr242
		}
		goto tr241
	tr184:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:382
		lex.act = 140
		goto st127
	st127:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof127
		}
	st_case_127:
// line scanner/scanner.go:3676
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 34:
			goto tr14
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
	tr13:
// line scanner/scanner.rl:50

		if lex.data[lex.p] == '\n' {
			lex.NewLines.Append(lex.p)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}

		goto st7
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
// line scanner/scanner.go:3709
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 34:
			goto tr14
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 34:
			goto tr14
		case 92:
			goto st9
		case 96:
			goto st7
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st7
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st7
			}
		default:
			goto st7
		}
		goto tr11
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		}
		goto st7
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 34:
			goto tr14
		case 36:
			goto tr11
		case 92:
			goto st9
		}
		goto st7
	tr247:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st128
	st128:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof128
		}
	st_case_128:
// line scanner/scanner.go:3794
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
			goto st129
		}
		if 512 <= _widec && _widec <= 767 {
			goto st128
		}
		goto tr244
	tr248:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st129
	st129:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof129
		}
	st_case_129:
// line scanner/scanner.go:3849
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
			goto tr248
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr247
		}
		goto tr246
	st130:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof130
		}
	st_case_130:
		if lex.data[(lex.p)] == 96 {
			goto tr239
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr239
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr239
			}
		default:
			goto tr239
		}
		goto st131
	st131:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof131
		}
	st_case_131:
		if lex.data[(lex.p)] == 96 {
			goto tr250
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr250
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr250
				}
			case lex.data[(lex.p)] >= 91:
				goto tr250
			}
		default:
			goto tr250
		}
		goto st131
	st132:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		if lex.data[(lex.p)] == 61 {
			goto tr251
		}
		goto tr239
	st133:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof133
		}
	st_case_133:
		switch lex.data[(lex.p)] {
		case 38:
			goto tr252
		case 61:
			goto tr253
		}
		goto tr239
	tr189:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st134
	st134:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof134
		}
	st_case_134:
// line scanner/scanner.go:3973
		switch lex.data[(lex.p)] {
		case 10:
			goto tr20
		case 13:
			goto tr20
		case 39:
			goto tr14
		case 92:
			goto st12
		}
		goto st11
	tr20:
// line scanner/scanner.rl:50

		if lex.data[lex.p] == '\n' {
			lex.NewLines.Append(lex.p)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.NewLines.Append(lex.p)
		}

		goto st11
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
// line scanner/scanner.go:4002
		switch lex.data[(lex.p)] {
		case 10:
			goto tr20
		case 13:
			goto tr20
		case 39:
			goto tr14
		case 92:
			goto st12
		}
		goto st11
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr20
		case 13:
			goto tr20
		}
		goto st11
	tr190:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st135
	st135:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof135
		}
	st_case_135:
// line scanner/scanner.go:4036
		switch lex.data[(lex.p)] {
		case 9:
			goto st13
		case 32:
			goto st13
		case 65:
			goto st14
		case 66:
			goto st19
		case 68:
			goto st31
		case 70:
			goto st37
		case 73:
			goto st41
		case 79:
			goto st48
		case 82:
			goto st54
		case 83:
			goto st57
		case 85:
			goto st62
		case 97:
			goto st14
		case 98:
			goto st19
		case 100:
			goto st31
		case 102:
			goto st37
		case 105:
			goto st41
		case 111:
			goto st48
		case 114:
			goto st54
		case 115:
			goto st57
		case 117:
			goto st62
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st13
		}
		goto tr239
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch lex.data[(lex.p)] {
		case 9:
			goto st13
		case 32:
			goto st13
		case 65:
			goto st14
		case 66:
			goto st19
		case 68:
			goto st31
		case 70:
			goto st37
		case 73:
			goto st41
		case 79:
			goto st48
		case 82:
			goto st54
		case 83:
			goto st57
		case 85:
			goto st62
		case 97:
			goto st14
		case 98:
			goto st19
		case 100:
			goto st31
		case 102:
			goto st37
		case 105:
			goto st41
		case 111:
			goto st48
		case 114:
			goto st54
		case 115:
			goto st57
		case 117:
			goto st62
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st13
		}
		goto tr22
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
		goto tr22
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 82:
			goto st16
		case 114:
			goto st16
		}
		goto tr22
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch lex.data[(lex.p)] {
		case 65:
			goto st17
		case 97:
			goto st17
		}
		goto tr22
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch lex.data[(lex.p)] {
		case 89:
			goto st18
		case 121:
			goto st18
		}
		goto tr22
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 9:
			goto st18
		case 32:
			goto st18
		case 41:
			goto tr37
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st18
		}
		goto tr22
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 73:
			goto st20
		case 79:
			goto st25
		case 105:
			goto st20
		case 111:
			goto st25
		}
		goto tr22
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 78:
			goto st21
		case 110:
			goto st21
		}
		goto tr22
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch lex.data[(lex.p)] {
		case 65:
			goto st22
		case 97:
			goto st22
		}
		goto tr22
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch lex.data[(lex.p)] {
		case 82:
			goto st23
		case 114:
			goto st23
		}
		goto tr22
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 89:
			goto st24
		case 121:
			goto st24
		}
		goto tr22
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 9:
			goto st24
		case 32:
			goto st24
		case 41:
			goto tr44
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st24
		}
		goto tr22
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 79:
			goto st26
		case 111:
			goto st26
		}
		goto tr22
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 76:
			goto st27
		case 108:
			goto st27
		}
		goto tr22
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 41:
			goto tr48
		case 69:
			goto st29
		case 101:
			goto st29
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st28
		}
		goto tr22
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 41:
			goto tr48
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st28
		}
		goto tr22
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 65:
			goto st30
		case 97:
			goto st30
		}
		goto tr22
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 78:
			goto st28
		case 110:
			goto st28
		}
		goto tr22
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 79:
			goto st32
		case 111:
			goto st32
		}
		goto tr22
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 85:
			goto st33
		case 117:
			goto st33
		}
		goto tr22
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 66:
			goto st34
		case 98:
			goto st34
		}
		goto tr22
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 76:
			goto st35
		case 108:
			goto st35
		}
		goto tr22
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 69:
			goto st36
		case 101:
			goto st36
		}
		goto tr22
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 9:
			goto st36
		case 32:
			goto st36
		case 41:
			goto tr56
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st36
		}
		goto tr22
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
		goto tr22
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 79:
			goto st39
		case 111:
			goto st39
		}
		goto tr22
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 65:
			goto st40
		case 97:
			goto st40
		}
		goto tr22
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 84:
			goto st36
		case 116:
			goto st36
		}
		goto tr22
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 78:
			goto st42
		case 110:
			goto st42
		}
		goto tr22
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch lex.data[(lex.p)] {
		case 84:
			goto st43
		case 116:
			goto st43
		}
		goto tr22
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch lex.data[(lex.p)] {
		case 9:
			goto st44
		case 32:
			goto st44
		case 41:
			goto tr63
		case 69:
			goto st45
		case 101:
			goto st45
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st44
		}
		goto tr22
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch lex.data[(lex.p)] {
		case 9:
			goto st44
		case 32:
			goto st44
		case 41:
			goto tr63
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st44
		}
		goto tr22
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 71:
			goto st46
		case 103:
			goto st46
		}
		goto tr22
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch lex.data[(lex.p)] {
		case 69:
			goto st47
		case 101:
			goto st47
		}
		goto tr22
	st47:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch lex.data[(lex.p)] {
		case 82:
			goto st44
		case 114:
			goto st44
		}
		goto tr22
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch lex.data[(lex.p)] {
		case 66:
			goto st49
		case 98:
			goto st49
		}
		goto tr22
	st49:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch lex.data[(lex.p)] {
		case 74:
			goto st50
		case 106:
			goto st50
		}
		goto tr22
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch lex.data[(lex.p)] {
		case 69:
			goto st51
		case 101:
			goto st51
		}
		goto tr22
	st51:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch lex.data[(lex.p)] {
		case 67:
			goto st52
		case 99:
			goto st52
		}
		goto tr22
	st52:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch lex.data[(lex.p)] {
		case 84:
			goto st53
		case 116:
			goto st53
		}
		goto tr22
	st53:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch lex.data[(lex.p)] {
		case 9:
			goto st53
		case 32:
			goto st53
		case 41:
			goto tr72
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st53
		}
		goto tr22
	st54:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch lex.data[(lex.p)] {
		case 69:
			goto st55
		case 101:
			goto st55
		}
		goto tr22
	st55:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch lex.data[(lex.p)] {
		case 65:
			goto st56
		case 97:
			goto st56
		}
		goto tr22
	st56:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch lex.data[(lex.p)] {
		case 76:
			goto st36
		case 108:
			goto st36
		}
		goto tr22
	st57:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch lex.data[(lex.p)] {
		case 84:
			goto st58
		case 116:
			goto st58
		}
		goto tr22
	st58:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch lex.data[(lex.p)] {
		case 82:
			goto st59
		case 114:
			goto st59
		}
		goto tr22
	st59:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch lex.data[(lex.p)] {
		case 73:
			goto st60
		case 105:
			goto st60
		}
		goto tr22
	st60:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch lex.data[(lex.p)] {
		case 78:
			goto st61
		case 110:
			goto st61
		}
		goto tr22
	st61:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch lex.data[(lex.p)] {
		case 71:
			goto st24
		case 103:
			goto st24
		}
		goto tr22
	st62:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch lex.data[(lex.p)] {
		case 78:
			goto st63
		case 110:
			goto st63
		}
		goto tr22
	st63:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch lex.data[(lex.p)] {
		case 83:
			goto st64
		case 115:
			goto st64
		}
		goto tr22
	st64:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch lex.data[(lex.p)] {
		case 69:
			goto st65
		case 101:
			goto st65
		}
		goto tr22
	st65:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch lex.data[(lex.p)] {
		case 84:
			goto st66
		case 116:
			goto st66
		}
		goto tr22
	st66:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch lex.data[(lex.p)] {
		case 9:
			goto st66
		case 32:
			goto st66
		case 41:
			goto tr83
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st66
		}
		goto tr22
	st136:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch lex.data[(lex.p)] {
		case 42:
			goto st137
		case 61:
			goto tr255
		}
		goto tr239
	st137:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof137
		}
	st_case_137:
		if lex.data[(lex.p)] == 61 {
			goto tr257
		}
		goto tr256
	st138:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof138
		}
	st_case_138:
		switch lex.data[(lex.p)] {
		case 43:
			goto tr258
		case 61:
			goto tr259
		}
		goto tr239
	st139:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof139
		}
	st_case_139:
		switch lex.data[(lex.p)] {
		case 45:
			goto tr260
		case 61:
			goto tr261
		case 62:
			goto tr262
		}
		goto tr239
	tr195:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st140
	st140:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof140
		}
	st_case_140:
// line scanner/scanner.go:4884
		switch lex.data[(lex.p)] {
		case 46:
			goto st67
		case 61:
			goto tr264
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr88
		}
		goto tr239
	st67:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		if lex.data[(lex.p)] == 46 {
			goto tr84
		}
		goto tr22
	tr88:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:176
		lex.act = 10
		goto st141
	st141:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof141
		}
	st_case_141:
// line scanner/scanner.go:4916
		switch lex.data[(lex.p)] {
		case 69:
			goto st68
		case 95:
			goto st70
		case 101:
			goto st68
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr88
		}
		goto tr265
	st68:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch lex.data[(lex.p)] {
		case 43:
			goto st69
		case 45:
			goto st69
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr86
		}
		goto tr11
	st69:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr86
		}
		goto tr11
	tr86:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:176
		lex.act = 10
		goto st142
	st142:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof142
		}
	st_case_142:
// line scanner/scanner.go:4965
		if lex.data[(lex.p)] == 95 {
			goto st69
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr86
		}
		goto tr265
	st70:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr88
		}
		goto tr87
	tr196:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st143
	st143:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof143
		}
	st_case_143:
// line scanner/scanner.go:4992
		switch lex.data[(lex.p)] {
		case 42:
			goto st71
		case 47:
			goto st128
		case 61:
			goto tr268
		}
		goto tr239
	tr92:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st71
	st71:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof71
		}
	st_case_71:
// line scanner/scanner.go:5011
		switch lex.data[(lex.p)] {
		case 10:
			goto st72
		case 42:
			goto st73
		}
		goto st71
	tr93:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st72
	st72:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof72
		}
	st_case_72:
// line scanner/scanner.go:5028
		switch lex.data[(lex.p)] {
		case 10:
			goto tr93
		case 42:
			goto tr94
		}
		goto tr92
	tr94:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st73
	st73:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof73
		}
	st_case_73:
// line scanner/scanner.go:5045
		switch lex.data[(lex.p)] {
		case 10:
			goto st72
		case 42:
			goto st73
		case 47:
			goto tr95
		}
		goto st71
	tr197:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:187
		lex.act = 12
		goto st144
	st144:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof144
		}
	st_case_144:
// line scanner/scanner.go:5067
		switch lex.data[(lex.p)] {
		case 46:
			goto tr270
		case 69:
			goto st68
		case 95:
			goto st74
		case 98:
			goto st75
		case 101:
			goto st68
		case 120:
			goto st76
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr97
		}
		goto tr269
	tr270:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:176
		lex.act = 10
		goto st145
	st145:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof145
		}
	st_case_145:
// line scanner/scanner.go:5098
		switch lex.data[(lex.p)] {
		case 69:
			goto st68
		case 101:
			goto st68
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr88
		}
		goto tr265
	tr97:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:187
		lex.act = 12
		goto st146
	st146:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof146
		}
	st_case_146:
// line scanner/scanner.go:5121
		switch lex.data[(lex.p)] {
		case 46:
			goto tr270
		case 69:
			goto st68
		case 95:
			goto st74
		case 101:
			goto st68
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr97
		}
		goto tr269
	st74:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr97
		}
		goto tr96
	st75:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr98
		}
		goto tr11
	tr98:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:177
		lex.act = 11
		goto st147
	st147:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof147
		}
	st_case_147:
// line scanner/scanner.go:5166
		if lex.data[(lex.p)] == 95 {
			goto st75
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr98
		}
		goto tr274
	st76:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr99
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr99
			}
		default:
			goto tr99
		}
		goto tr11
	tr99:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:202
		lex.act = 13
		goto st148
	st148:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof148
		}
	st_case_148:
// line scanner/scanner.go:5204
		if lex.data[(lex.p)] == 95 {
			goto st76
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr99
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr99
			}
		default:
			goto tr99
		}
		goto tr275
	st149:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof149
		}
	st_case_149:
		if lex.data[(lex.p)] == 58 {
			goto tr276
		}
		goto tr239
	tr199:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st150
	st150:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof150
		}
	st_case_150:
// line scanner/scanner.go:5240
		switch lex.data[(lex.p)] {
		case 10:
			goto st78
		case 13:
			goto st79
		case 32:
			goto st77
		case 63:
			goto st80
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st77
		}
		goto tr239
	tr104:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st77
	st77:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof77
		}
	st_case_77:
// line scanner/scanner.go:5264
		switch lex.data[(lex.p)] {
		case 10:
			goto st78
		case 13:
			goto st79
		case 32:
			goto st77
		case 63:
			goto st80
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st77
		}
		goto tr22
	tr105:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st78
	st78:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof78
		}
	st_case_78:
// line scanner/scanner.go:5288
		switch lex.data[(lex.p)] {
		case 10:
			goto tr105
		case 13:
			goto tr106
		case 32:
			goto tr104
		case 63:
			goto tr107
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr104
		}
		goto tr22
	tr106:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st79
	st79:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof79
		}
	st_case_79:
// line scanner/scanner.go:5312
		if lex.data[(lex.p)] == 10 {
			goto st78
		}
		goto tr22
	tr107:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st80
	st80:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof80
		}
	st_case_80:
// line scanner/scanner.go:5326
		if lex.data[(lex.p)] == 62 {
			goto tr108
		}
		goto tr22
	tr108:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st151
	st151:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof151
		}
	st_case_151:
// line scanner/scanner.go:5341
		switch lex.data[(lex.p)] {
		case 10:
			goto st152
		case 13:
			goto st81
		}
		goto tr277
	st152:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof152
		}
	st_case_152:
		goto tr279
	st81:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		if lex.data[(lex.p)] == 10 {
			goto st152
		}
		goto tr109
	st153:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof153
		}
	st_case_153:
		switch lex.data[(lex.p)] {
		case 60:
			goto tr280
		case 61:
			goto st156
		case 62:
			goto tr282
		}
		goto tr239
	tr280:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:317
		lex.act = 118
		goto st154
	st154:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof154
		}
	st_case_154:
// line scanner/scanner.go:5390
		switch lex.data[(lex.p)] {
		case 60:
			goto st82
		case 61:
			goto tr284
		}
		goto tr283
	st82:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		switch lex.data[(lex.p)] {
		case 9:
			goto st82
		case 32:
			goto st82
		case 34:
			goto st83
		case 39:
			goto st87
		case 96:
			goto tr11
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr11
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr114
	st83:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		if lex.data[(lex.p)] == 96 {
			goto tr11
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr11
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr115
	tr115:
// line scanner/scanner.rl:47
		lblStart = lex.p
		goto st84
	st84:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof84
		}
	st_case_84:
// line scanner/scanner.go:5458
		switch lex.data[(lex.p)] {
		case 34:
			goto tr116
		case 96:
			goto tr11
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr11
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr11
				}
			case lex.data[(lex.p)] >= 91:
				goto tr11
			}
		default:
			goto tr11
		}
		goto st84
	tr116:
// line scanner/scanner.rl:48
		lblEnd = lex.p
		goto st85
	st85:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof85
		}
	st_case_85:
// line scanner/scanner.go:5492
		switch lex.data[(lex.p)] {
		case 10:
			goto st155
		case 13:
			goto st86
		}
		goto tr11
	tr122:
// line scanner/scanner.rl:48
		lblEnd = lex.p
		goto st155
	st155:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof155
		}
	st_case_155:
// line scanner/scanner.go:5509
		goto tr285
	tr123:
// line scanner/scanner.rl:48
		lblEnd = lex.p
		goto st86
	st86:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof86
		}
	st_case_86:
// line scanner/scanner.go:5520
		if lex.data[(lex.p)] == 10 {
			goto st155
		}
		goto tr11
	st87:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		if lex.data[(lex.p)] == 96 {
			goto tr11
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr11
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr120
	tr120:
// line scanner/scanner.rl:47
		lblStart = lex.p
		goto st88
	st88:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof88
		}
	st_case_88:
// line scanner/scanner.go:5555
		switch lex.data[(lex.p)] {
		case 39:
			goto tr116
		case 96:
			goto tr11
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr11
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr11
				}
			case lex.data[(lex.p)] >= 91:
				goto tr11
			}
		default:
			goto tr11
		}
		goto st88
	tr114:
// line scanner/scanner.rl:47
		lblStart = lex.p
		goto st89
	st89:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof89
		}
	st_case_89:
// line scanner/scanner.go:5589
		switch lex.data[(lex.p)] {
		case 10:
			goto tr122
		case 13:
			goto tr123
		case 96:
			goto tr11
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr11
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr11
				}
			case lex.data[(lex.p)] >= 91:
				goto tr11
			}
		default:
			goto tr11
		}
		goto st89
	st156:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof156
		}
	st_case_156:
		if lex.data[(lex.p)] == 62 {
			goto tr287
		}
		goto tr286
	st157:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof157
		}
	st_case_157:
		switch lex.data[(lex.p)] {
		case 61:
			goto st158
		case 62:
			goto tr289
		}
		goto tr239
	st158:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof158
		}
	st_case_158:
		if lex.data[(lex.p)] == 61 {
			goto tr291
		}
		goto tr290
	st159:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof159
		}
	st_case_159:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr292
		case 62:
			goto st160
		}
		goto tr239
	st160:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof160
		}
	st_case_160:
		if lex.data[(lex.p)] == 61 {
			goto tr295
		}
		goto tr294
	st161:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch lex.data[(lex.p)] {
		case 62:
			goto tr296
		case 63:
			goto st164
		}
		goto tr239
	tr296:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st162
	st162:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof162
		}
	st_case_162:
// line scanner/scanner.go:5689
		switch lex.data[(lex.p)] {
		case 10:
			goto st163
		case 13:
			goto st90
		}
		goto tr298
	st163:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof163
		}
	st_case_163:
		goto tr300
	st90:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		if lex.data[(lex.p)] == 10 {
			goto st163
		}
		goto tr125
	st164:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		if lex.data[(lex.p)] == 61 {
			goto tr302
		}
		goto tr301
	st165:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof165
		}
	st_case_165:
		switch lex.data[(lex.p)] {
		case 66:
			goto st167
		case 78:
			goto st173
		case 82:
			goto st174
		case 83:
			goto tr307
		case 96:
			goto tr303
		case 98:
			goto st167
		case 110:
			goto st173
		case 114:
			goto st174
		case 115:
			goto tr307
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	tr211:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:357
		lex.act = 135
		goto st166
	tr307:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:215
		lex.act = 16
		goto st166
	tr313:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:213
		lex.act = 14
		goto st166
	tr314:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:286
		lex.act = 87
		goto st166
	tr317:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:214
		lex.act = 15
		goto st166
	tr322:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:216
		lex.act = 17
		goto st166
	tr334:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:217
		lex.act = 18
		goto st166
	tr335:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:218
		lex.act = 19
		goto st166
	tr337:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:219
		lex.act = 20
		goto st166
	tr344:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:244
		lex.act = 45
		goto st166
	tr348:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:220
		lex.act = 21
		goto st166
	tr350:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:221
		lex.act = 22
		goto st166
	tr354:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:222
		lex.act = 23
		goto st166
	tr358:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:223
		lex.act = 24
		goto st166
	tr361:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:226
		lex.act = 27
		goto st166
	tr367:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:224
		lex.act = 25
		goto st166
	tr371:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:225
		lex.act = 26
		goto st166
	tr372:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:238
		lex.act = 39
		goto st166
	tr380:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:227
		lex.act = 28
		goto st166
	tr385:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:229
		lex.act = 30
		goto st166
	tr388:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:230
		lex.act = 31
		goto st166
	tr400:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:231
		lex.act = 32
		goto st166
	tr407:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:233
		lex.act = 34
		goto st166
	tr408:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:234
		lex.act = 35
		goto st166
	tr413:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:235
		lex.act = 36
		goto st166
	tr417:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:236
		lex.act = 37
		goto st166
	tr419:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:237
		lex.act = 38
		goto st166
	tr425:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:239
		lex.act = 40
		goto st166
	tr427:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:245
		lex.act = 46
		goto st166
	tr434:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:241
		lex.act = 42
		goto st166
	tr440:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:243
		lex.act = 44
		goto st166
	tr446:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:246
		lex.act = 47
		goto st166
	tr448:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:247
		lex.act = 48
		goto st166
	tr449:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:248
		lex.act = 49
		goto st166
	tr460:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:250
		lex.act = 51
		goto st166
	tr473:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:273
		lex.act = 74
		goto st166
	tr481:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:251
		lex.act = 52
		goto st166
	tr485:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:252
		lex.act = 53
		goto st166
	tr491:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:253
		lex.act = 54
		goto st166
	tr494:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:249
		lex.act = 50
		goto st166
	tr497:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:254
		lex.act = 55
		goto st166
	tr506:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:255
		lex.act = 56
		goto st166
	tr507:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:285
		lex.act = 86
		goto st166
	tr508:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:287
		lex.act = 88
		goto st166
	tr515:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:258
		lex.act = 59
		goto st166
	tr518:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:256
		lex.act = 57
		goto st166
	tr524:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:259
		lex.act = 60
		goto st166
	tr528:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:257
		lex.act = 58
		goto st166
	tr541:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:275
		lex.act = 76
		goto st166
	tr544:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:260
		lex.act = 61
		goto st166
	tr550:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:261
		lex.act = 62
		goto st166
	tr554:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:262
		lex.act = 63
		goto st166
	tr559:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:263
		lex.act = 64
		goto st166
	tr561:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:265
		lex.act = 66
		goto st166
	tr563:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:264
		lex.act = 65
		goto st166
	tr568:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:266
		lex.act = 67
		goto st166
	tr569:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:267
		lex.act = 68
		goto st166
	tr571:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:268
		lex.act = 69
		goto st166
	tr575:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:269
		lex.act = 70
		goto st166
	tr577:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:288
		lex.act = 89
		goto st166
	tr586:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:270
		lex.act = 71
		goto st166
	tr602:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:276
		lex.act = 77
		goto st166
	tr606:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:277
		lex.act = 78
		goto st166
	tr612:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:278
		lex.act = 79
		goto st166
	tr620:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:279
		lex.act = 80
		goto st166
	tr632:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:284
		lex.act = 85
		goto st166
	tr637:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:280
		lex.act = 81
		goto st166
	tr644:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:282
		lex.act = 83
		goto st166
	tr654:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:281
		lex.act = 82
		goto st166
	tr660:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:283
		lex.act = 84
		goto st166
	st166:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof166
		}
	st_case_166:
// line scanner/scanner.go:6259
		if lex.data[(lex.p)] == 96 {
			goto tr11
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr11
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr11
				}
			case lex.data[(lex.p)] >= 91:
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr211
	st167:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof167
		}
	st_case_167:
		switch lex.data[(lex.p)] {
		case 83:
			goto st168
		case 96:
			goto tr303
		case 115:
			goto st168
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st168:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof168
		}
	st_case_168:
		switch lex.data[(lex.p)] {
		case 84:
			goto st169
		case 96:
			goto tr303
		case 116:
			goto st169
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st169:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch lex.data[(lex.p)] {
		case 82:
			goto st170
		case 96:
			goto tr303
		case 114:
			goto st170
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st170:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch lex.data[(lex.p)] {
		case 65:
			goto st171
		case 96:
			goto tr303
		case 97:
			goto st171
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st171:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch lex.data[(lex.p)] {
		case 67:
			goto st172
		case 96:
			goto tr303
		case 99:
			goto st172
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st172:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof172
		}
	st_case_172:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr313
		case 96:
			goto tr303
		case 116:
			goto tr313
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st173:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof173
		}
	st_case_173:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr314
		case 96:
			goto tr303
		case 100:
			goto tr314
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st174:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof174
		}
	st_case_174:
		switch lex.data[(lex.p)] {
		case 82:
			goto st175
		case 96:
			goto tr303
		case 114:
			goto st175
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st175:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof175
		}
	st_case_175:
		switch lex.data[(lex.p)] {
		case 65:
			goto st176
		case 96:
			goto tr303
		case 97:
			goto st176
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st176:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof176
		}
	st_case_176:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr317
		case 96:
			goto tr303
		case 121:
			goto tr317
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	tr205:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:357
		lex.act = 135
		goto st177
	st177:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof177
		}
	st_case_177:
// line scanner/scanner.go:6603
		switch lex.data[(lex.p)] {
		case 34:
			goto st7
		case 60:
			goto st91
		case 82:
			goto st178
		case 96:
			goto tr303
		case 114:
			goto st178
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st91:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		if lex.data[(lex.p)] == 60 {
			goto st92
		}
		goto tr127
	st92:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		if lex.data[(lex.p)] == 60 {
			goto st82
		}
		goto tr127
	st178:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof178
		}
	st_case_178:
		switch lex.data[(lex.p)] {
		case 69:
			goto st179
		case 96:
			goto tr303
		case 101:
			goto st179
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st179:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof179
		}
	st_case_179:
		switch lex.data[(lex.p)] {
		case 65:
			goto st180
		case 96:
			goto tr303
		case 97:
			goto st180
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st180:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof180
		}
	st_case_180:
		switch lex.data[(lex.p)] {
		case 75:
			goto tr322
		case 96:
			goto tr303
		case 107:
			goto tr322
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st181:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof181
		}
	st_case_181:
		switch lex.data[(lex.p)] {
		case 65:
			goto st182
		case 70:
			goto st191
		case 76:
			goto st198
		case 79:
			goto st203
		case 96:
			goto tr303
		case 97:
			goto st182
		case 102:
			goto st191
		case 108:
			goto st198
		case 111:
			goto st203
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st182:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof182
		}
	st_case_182:
		switch lex.data[(lex.p)] {
		case 76:
			goto st183
		case 83:
			goto st188
		case 84:
			goto st189
		case 96:
			goto tr303
		case 108:
			goto st183
		case 115:
			goto st188
		case 116:
			goto st189
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st183:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof183
		}
	st_case_183:
		switch lex.data[(lex.p)] {
		case 76:
			goto st184
		case 96:
			goto tr303
		case 108:
			goto st184
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st184:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof184
		}
	st_case_184:
		switch lex.data[(lex.p)] {
		case 65:
			goto st185
		case 96:
			goto tr303
		case 97:
			goto st185
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st185:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch lex.data[(lex.p)] {
		case 66:
			goto st186
		case 96:
			goto tr303
		case 98:
			goto st186
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st186:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof186
		}
	st_case_186:
		switch lex.data[(lex.p)] {
		case 76:
			goto st187
		case 96:
			goto tr303
		case 108:
			goto st187
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st187:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof187
		}
	st_case_187:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr334
		case 96:
			goto tr303
		case 101:
			goto tr334
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st188:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof188
		}
	st_case_188:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr335
		case 96:
			goto tr303
		case 101:
			goto tr335
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st189:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof189
		}
	st_case_189:
		switch lex.data[(lex.p)] {
		case 67:
			goto st190
		case 96:
			goto tr303
		case 99:
			goto st190
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st190:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof190
		}
	st_case_190:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr337
		case 96:
			goto tr303
		case 104:
			goto tr337
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st191:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof191
		}
	st_case_191:
		switch lex.data[(lex.p)] {
		case 85:
			goto st192
		case 96:
			goto tr303
		case 117:
			goto st192
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st192:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof192
		}
	st_case_192:
		switch lex.data[(lex.p)] {
		case 78:
			goto st193
		case 96:
			goto tr303
		case 110:
			goto st193
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st193:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof193
		}
	st_case_193:
		switch lex.data[(lex.p)] {
		case 67:
			goto st194
		case 96:
			goto tr303
		case 99:
			goto st194
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st194:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof194
		}
	st_case_194:
		switch lex.data[(lex.p)] {
		case 84:
			goto st195
		case 96:
			goto tr303
		case 116:
			goto st195
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st195:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof195
		}
	st_case_195:
		switch lex.data[(lex.p)] {
		case 73:
			goto st196
		case 96:
			goto tr303
		case 105:
			goto st196
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st196:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof196
		}
	st_case_196:
		switch lex.data[(lex.p)] {
		case 79:
			goto st197
		case 96:
			goto tr303
		case 111:
			goto st197
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st197:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof197
		}
	st_case_197:
		switch lex.data[(lex.p)] {
		case 78:
			goto tr344
		case 96:
			goto tr303
		case 110:
			goto tr344
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st198:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof198
		}
	st_case_198:
		switch lex.data[(lex.p)] {
		case 65:
			goto st199
		case 79:
			goto st201
		case 96:
			goto tr303
		case 97:
			goto st199
		case 111:
			goto st201
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st199:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof199
		}
	st_case_199:
		switch lex.data[(lex.p)] {
		case 83:
			goto st200
		case 96:
			goto tr303
		case 115:
			goto st200
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st200:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof200
		}
	st_case_200:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr348
		case 96:
			goto tr303
		case 115:
			goto tr348
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st201:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof201
		}
	st_case_201:
		switch lex.data[(lex.p)] {
		case 78:
			goto st202
		case 96:
			goto tr303
		case 110:
			goto st202
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st202:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof202
		}
	st_case_202:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr350
		case 96:
			goto tr303
		case 101:
			goto tr350
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st203:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof203
		}
	st_case_203:
		switch lex.data[(lex.p)] {
		case 78:
			goto st204
		case 96:
			goto tr303
		case 110:
			goto st204
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st204:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof204
		}
	st_case_204:
		switch lex.data[(lex.p)] {
		case 83:
			goto st205
		case 84:
			goto st206
		case 96:
			goto tr303
		case 115:
			goto st205
		case 116:
			goto st206
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st205:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof205
		}
	st_case_205:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr354
		case 96:
			goto tr303
		case 116:
			goto tr354
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st206:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof206
		}
	st_case_206:
		switch lex.data[(lex.p)] {
		case 73:
			goto st207
		case 96:
			goto tr303
		case 105:
			goto st207
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st207:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof207
		}
	st_case_207:
		switch lex.data[(lex.p)] {
		case 78:
			goto st208
		case 96:
			goto tr303
		case 110:
			goto st208
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st208:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof208
		}
	st_case_208:
		switch lex.data[(lex.p)] {
		case 85:
			goto st209
		case 96:
			goto tr303
		case 117:
			goto st209
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st209:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof209
		}
	st_case_209:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr358
		case 96:
			goto tr303
		case 101:
			goto tr358
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st210:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof210
		}
	st_case_210:
		switch lex.data[(lex.p)] {
		case 69:
			goto st211
		case 73:
			goto st220
		case 79:
			goto tr361
		case 96:
			goto tr303
		case 101:
			goto st211
		case 105:
			goto st220
		case 111:
			goto tr361
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st211:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof211
		}
	st_case_211:
		switch lex.data[(lex.p)] {
		case 67:
			goto st212
		case 70:
			goto st216
		case 96:
			goto tr303
		case 99:
			goto st212
		case 102:
			goto st216
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st212:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof212
		}
	st_case_212:
		switch lex.data[(lex.p)] {
		case 76:
			goto st213
		case 96:
			goto tr303
		case 108:
			goto st213
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st213:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof213
		}
	st_case_213:
		switch lex.data[(lex.p)] {
		case 65:
			goto st214
		case 96:
			goto tr303
		case 97:
			goto st214
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st214:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof214
		}
	st_case_214:
		switch lex.data[(lex.p)] {
		case 82:
			goto st215
		case 96:
			goto tr303
		case 114:
			goto st215
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st215:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof215
		}
	st_case_215:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr367
		case 96:
			goto tr303
		case 101:
			goto tr367
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st216:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof216
		}
	st_case_216:
		switch lex.data[(lex.p)] {
		case 65:
			goto st217
		case 96:
			goto tr303
		case 97:
			goto st217
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st217:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof217
		}
	st_case_217:
		switch lex.data[(lex.p)] {
		case 85:
			goto st218
		case 96:
			goto tr303
		case 117:
			goto st218
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st218:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof218
		}
	st_case_218:
		switch lex.data[(lex.p)] {
		case 76:
			goto st219
		case 96:
			goto tr303
		case 108:
			goto st219
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st219:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof219
		}
	st_case_219:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr371
		case 96:
			goto tr303
		case 116:
			goto tr371
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st220:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof220
		}
	st_case_220:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr372
		case 96:
			goto tr303
		case 101:
			goto tr372
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st221:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof221
		}
	st_case_221:
		switch lex.data[(lex.p)] {
		case 67:
			goto st222
		case 76:
			goto st224
		case 77:
			goto st228
		case 78:
			goto st231
		case 86:
			goto st255
		case 88:
			goto st257
		case 96:
			goto tr303
		case 99:
			goto st222
		case 108:
			goto st224
		case 109:
			goto st228
		case 110:
			goto st231
		case 118:
			goto st255
		case 120:
			goto st257
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st222:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof222
		}
	st_case_222:
		switch lex.data[(lex.p)] {
		case 72:
			goto st223
		case 96:
			goto tr303
		case 104:
			goto st223
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st223:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof223
		}
	st_case_223:
		switch lex.data[(lex.p)] {
		case 79:
			goto tr380
		case 96:
			goto tr303
		case 111:
			goto tr380
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st224:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof224
		}
	st_case_224:
		switch lex.data[(lex.p)] {
		case 83:
			goto st225
		case 96:
			goto tr303
		case 115:
			goto st225
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st225:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof225
		}
	st_case_225:
		switch lex.data[(lex.p)] {
		case 69:
			goto st226
		case 96:
			goto tr303
		case 101:
			goto st226
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st226:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof226
		}
	st_case_226:
		switch lex.data[(lex.p)] {
		case 73:
			goto st227
		case 96:
			goto tr383
		case 105:
			goto st227
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr383
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr383
				}
			case lex.data[(lex.p)] >= 91:
				goto tr383
			}
		default:
			goto tr383
		}
		goto tr211
	st227:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof227
		}
	st_case_227:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr385
		case 96:
			goto tr303
		case 102:
			goto tr385
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st228:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof228
		}
	st_case_228:
		switch lex.data[(lex.p)] {
		case 80:
			goto st229
		case 96:
			goto tr303
		case 112:
			goto st229
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st229:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof229
		}
	st_case_229:
		switch lex.data[(lex.p)] {
		case 84:
			goto st230
		case 96:
			goto tr303
		case 116:
			goto st230
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st230:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof230
		}
	st_case_230:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr388
		case 96:
			goto tr303
		case 121:
			goto tr388
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st231:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof231
		}
	st_case_231:
		switch lex.data[(lex.p)] {
		case 68:
			goto st232
		case 96:
			goto tr303
		case 100:
			goto st232
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st232:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof232
		}
	st_case_232:
		switch lex.data[(lex.p)] {
		case 68:
			goto st233
		case 70:
			goto st239
		case 73:
			goto st245
		case 83:
			goto st246
		case 87:
			goto st251
		case 96:
			goto tr303
		case 100:
			goto st233
		case 102:
			goto st239
		case 105:
			goto st245
		case 115:
			goto st246
		case 119:
			goto st251
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st233:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof233
		}
	st_case_233:
		switch lex.data[(lex.p)] {
		case 69:
			goto st234
		case 96:
			goto tr303
		case 101:
			goto st234
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st234:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof234
		}
	st_case_234:
		switch lex.data[(lex.p)] {
		case 67:
			goto st235
		case 96:
			goto tr303
		case 99:
			goto st235
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st235:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof235
		}
	st_case_235:
		switch lex.data[(lex.p)] {
		case 76:
			goto st236
		case 96:
			goto tr303
		case 108:
			goto st236
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st236:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof236
		}
	st_case_236:
		switch lex.data[(lex.p)] {
		case 65:
			goto st237
		case 96:
			goto tr303
		case 97:
			goto st237
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st237:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof237
		}
	st_case_237:
		switch lex.data[(lex.p)] {
		case 82:
			goto st238
		case 96:
			goto tr303
		case 114:
			goto st238
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st238:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof238
		}
	st_case_238:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr400
		case 96:
			goto tr303
		case 101:
			goto tr400
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st239:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof239
		}
	st_case_239:
		switch lex.data[(lex.p)] {
		case 79:
			goto st240
		case 96:
			goto tr303
		case 111:
			goto st240
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st240:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof240
		}
	st_case_240:
		switch lex.data[(lex.p)] {
		case 82:
			goto st241
		case 96:
			goto tr303
		case 114:
			goto st241
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st241:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof241
		}
	st_case_241:
		switch lex.data[(lex.p)] {
		case 69:
			goto st242
		case 96:
			goto tr403
		case 101:
			goto st242
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr403
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr403
				}
			case lex.data[(lex.p)] >= 91:
				goto tr403
			}
		default:
			goto tr403
		}
		goto tr211
	st242:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof242
		}
	st_case_242:
		switch lex.data[(lex.p)] {
		case 65:
			goto st243
		case 96:
			goto tr303
		case 97:
			goto st243
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st243:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof243
		}
	st_case_243:
		switch lex.data[(lex.p)] {
		case 67:
			goto st244
		case 96:
			goto tr303
		case 99:
			goto st244
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st244:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof244
		}
	st_case_244:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr407
		case 96:
			goto tr303
		case 104:
			goto tr407
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st245:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof245
		}
	st_case_245:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr408
		case 96:
			goto tr303
		case 102:
			goto tr408
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st246:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof246
		}
	st_case_246:
		switch lex.data[(lex.p)] {
		case 87:
			goto st247
		case 96:
			goto tr303
		case 119:
			goto st247
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st247:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof247
		}
	st_case_247:
		switch lex.data[(lex.p)] {
		case 73:
			goto st248
		case 96:
			goto tr303
		case 105:
			goto st248
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st248:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof248
		}
	st_case_248:
		switch lex.data[(lex.p)] {
		case 84:
			goto st249
		case 96:
			goto tr303
		case 116:
			goto st249
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st249:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof249
		}
	st_case_249:
		switch lex.data[(lex.p)] {
		case 67:
			goto st250
		case 96:
			goto tr303
		case 99:
			goto st250
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st250:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof250
		}
	st_case_250:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr413
		case 96:
			goto tr303
		case 104:
			goto tr413
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st251:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof251
		}
	st_case_251:
		switch lex.data[(lex.p)] {
		case 72:
			goto st252
		case 96:
			goto tr303
		case 104:
			goto st252
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st252:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof252
		}
	st_case_252:
		switch lex.data[(lex.p)] {
		case 73:
			goto st253
		case 96:
			goto tr303
		case 105:
			goto st253
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st253:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof253
		}
	st_case_253:
		switch lex.data[(lex.p)] {
		case 76:
			goto st254
		case 96:
			goto tr303
		case 108:
			goto st254
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st254:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof254
		}
	st_case_254:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr417
		case 96:
			goto tr303
		case 101:
			goto tr417
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st255:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof255
		}
	st_case_255:
		switch lex.data[(lex.p)] {
		case 65:
			goto st256
		case 96:
			goto tr303
		case 97:
			goto st256
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st256:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof256
		}
	st_case_256:
		switch lex.data[(lex.p)] {
		case 76:
			goto tr419
		case 96:
			goto tr303
		case 108:
			goto tr419
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st257:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof257
		}
	st_case_257:
		switch lex.data[(lex.p)] {
		case 73:
			goto st258
		case 84:
			goto st259
		case 96:
			goto tr303
		case 105:
			goto st258
		case 116:
			goto st259
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st258:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof258
		}
	st_case_258:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr372
		case 96:
			goto tr303
		case 116:
			goto tr372
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st259:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof259
		}
	st_case_259:
		switch lex.data[(lex.p)] {
		case 69:
			goto st260
		case 96:
			goto tr303
		case 101:
			goto st260
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st260:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof260
		}
	st_case_260:
		switch lex.data[(lex.p)] {
		case 78:
			goto st261
		case 96:
			goto tr303
		case 110:
			goto st261
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st261:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof261
		}
	st_case_261:
		switch lex.data[(lex.p)] {
		case 68:
			goto st262
		case 96:
			goto tr303
		case 100:
			goto st262
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st262:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof262
		}
	st_case_262:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr425
		case 96:
			goto tr303
		case 115:
			goto tr425
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st263:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof263
		}
	st_case_263:
		switch lex.data[(lex.p)] {
		case 73:
			goto st264
		case 78:
			goto tr427
		case 79:
			goto st269
		case 85:
			goto st192
		case 96:
			goto tr303
		case 105:
			goto st264
		case 110:
			goto tr427
		case 111:
			goto st269
		case 117:
			goto st192
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st264:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof264
		}
	st_case_264:
		switch lex.data[(lex.p)] {
		case 78:
			goto st265
		case 96:
			goto tr303
		case 110:
			goto st265
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st265:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof265
		}
	st_case_265:
		switch lex.data[(lex.p)] {
		case 65:
			goto st266
		case 96:
			goto tr303
		case 97:
			goto st266
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st266:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof266
		}
	st_case_266:
		switch lex.data[(lex.p)] {
		case 76:
			goto st267
		case 96:
			goto tr303
		case 108:
			goto st267
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st267:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof267
		}
	st_case_267:
		switch lex.data[(lex.p)] {
		case 76:
			goto st268
		case 96:
			goto tr432
		case 108:
			goto st268
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr432
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr432
				}
			case lex.data[(lex.p)] >= 91:
				goto tr432
			}
		default:
			goto tr432
		}
		goto tr211
	st268:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof268
		}
	st_case_268:
		switch lex.data[(lex.p)] {
		case 89:
			goto tr434
		case 96:
			goto tr303
		case 121:
			goto tr434
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st269:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof269
		}
	st_case_269:
		switch lex.data[(lex.p)] {
		case 82:
			goto st270
		case 96:
			goto tr303
		case 114:
			goto st270
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st270:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof270
		}
	st_case_270:
		switch lex.data[(lex.p)] {
		case 69:
			goto st271
		case 96:
			goto tr436
		case 101:
			goto st271
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr436
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr436
				}
			case lex.data[(lex.p)] >= 91:
				goto tr436
			}
		default:
			goto tr436
		}
		goto tr211
	st271:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof271
		}
	st_case_271:
		switch lex.data[(lex.p)] {
		case 65:
			goto st272
		case 96:
			goto tr303
		case 97:
			goto st272
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st272:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof272
		}
	st_case_272:
		switch lex.data[(lex.p)] {
		case 67:
			goto st273
		case 96:
			goto tr303
		case 99:
			goto st273
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st273:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof273
		}
	st_case_273:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr440
		case 96:
			goto tr303
		case 104:
			goto tr440
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st274:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof274
		}
	st_case_274:
		switch lex.data[(lex.p)] {
		case 76:
			goto st275
		case 79:
			goto st279
		case 96:
			goto tr303
		case 108:
			goto st275
		case 111:
			goto st279
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st275:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof275
		}
	st_case_275:
		switch lex.data[(lex.p)] {
		case 79:
			goto st276
		case 96:
			goto tr303
		case 111:
			goto st276
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st276:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof276
		}
	st_case_276:
		switch lex.data[(lex.p)] {
		case 66:
			goto st277
		case 96:
			goto tr303
		case 98:
			goto st277
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st277:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof277
		}
	st_case_277:
		switch lex.data[(lex.p)] {
		case 65:
			goto st278
		case 96:
			goto tr303
		case 97:
			goto st278
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st278:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof278
		}
	st_case_278:
		switch lex.data[(lex.p)] {
		case 76:
			goto tr446
		case 96:
			goto tr303
		case 108:
			goto tr446
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st279:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof279
		}
	st_case_279:
		switch lex.data[(lex.p)] {
		case 84:
			goto st280
		case 96:
			goto tr303
		case 116:
			goto st280
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st280:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof280
		}
	st_case_280:
		switch lex.data[(lex.p)] {
		case 79:
			goto tr448
		case 96:
			goto tr303
		case 111:
			goto tr448
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st281:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof281
		}
	st_case_281:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr449
		case 77:
			goto st282
		case 78:
			goto st290
		case 83:
			goto st317
		case 96:
			goto tr303
		case 102:
			goto tr449
		case 109:
			goto st282
		case 110:
			goto st290
		case 115:
			goto st317
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st282:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof282
		}
	st_case_282:
		switch lex.data[(lex.p)] {
		case 80:
			goto st283
		case 96:
			goto tr303
		case 112:
			goto st283
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st283:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof283
		}
	st_case_283:
		switch lex.data[(lex.p)] {
		case 76:
			goto st284
		case 96:
			goto tr303
		case 108:
			goto st284
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st284:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof284
		}
	st_case_284:
		switch lex.data[(lex.p)] {
		case 69:
			goto st285
		case 96:
			goto tr303
		case 101:
			goto st285
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st285:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof285
		}
	st_case_285:
		switch lex.data[(lex.p)] {
		case 77:
			goto st286
		case 96:
			goto tr303
		case 109:
			goto st286
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st286:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof286
		}
	st_case_286:
		switch lex.data[(lex.p)] {
		case 69:
			goto st287
		case 96:
			goto tr303
		case 101:
			goto st287
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st287:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof287
		}
	st_case_287:
		switch lex.data[(lex.p)] {
		case 78:
			goto st288
		case 96:
			goto tr303
		case 110:
			goto st288
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st288:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof288
		}
	st_case_288:
		switch lex.data[(lex.p)] {
		case 84:
			goto st289
		case 96:
			goto tr303
		case 116:
			goto st289
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st289:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof289
		}
	st_case_289:
		switch lex.data[(lex.p)] {
		case 83:
			goto tr460
		case 96:
			goto tr303
		case 115:
			goto tr460
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st290:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof290
		}
	st_case_290:
		switch lex.data[(lex.p)] {
		case 67:
			goto st291
		case 83:
			goto st300
		case 84:
			goto st311
		case 96:
			goto tr303
		case 99:
			goto st291
		case 115:
			goto st300
		case 116:
			goto st311
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st291:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof291
		}
	st_case_291:
		switch lex.data[(lex.p)] {
		case 76:
			goto st292
		case 96:
			goto tr303
		case 108:
			goto st292
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st292:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof292
		}
	st_case_292:
		switch lex.data[(lex.p)] {
		case 85:
			goto st293
		case 96:
			goto tr303
		case 117:
			goto st293
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st293:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof293
		}
	st_case_293:
		switch lex.data[(lex.p)] {
		case 68:
			goto st294
		case 96:
			goto tr303
		case 100:
			goto st294
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st294:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof294
		}
	st_case_294:
		switch lex.data[(lex.p)] {
		case 69:
			goto st295
		case 96:
			goto tr303
		case 101:
			goto st295
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st295:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof295
		}
	st_case_295:
		if lex.data[(lex.p)] == 95 {
			goto st296
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr468
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr468
				}
			case lex.data[(lex.p)] >= 91:
				goto tr468
			}
		default:
			goto tr468
		}
		goto tr211
	st296:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof296
		}
	st_case_296:
		switch lex.data[(lex.p)] {
		case 79:
			goto st297
		case 96:
			goto tr303
		case 111:
			goto st297
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st297:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof297
		}
	st_case_297:
		switch lex.data[(lex.p)] {
		case 78:
			goto st298
		case 96:
			goto tr303
		case 110:
			goto st298
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st298:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof298
		}
	st_case_298:
		switch lex.data[(lex.p)] {
		case 67:
			goto st299
		case 96:
			goto tr303
		case 99:
			goto st299
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st299:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof299
		}
	st_case_299:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr473
		case 96:
			goto tr303
		case 101:
			goto tr473
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st300:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof300
		}
	st_case_300:
		switch lex.data[(lex.p)] {
		case 84:
			goto st301
		case 96:
			goto tr303
		case 116:
			goto st301
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st301:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof301
		}
	st_case_301:
		switch lex.data[(lex.p)] {
		case 65:
			goto st302
		case 69:
			goto st307
		case 96:
			goto tr303
		case 97:
			goto st302
		case 101:
			goto st307
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st302:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof302
		}
	st_case_302:
		switch lex.data[(lex.p)] {
		case 78:
			goto st303
		case 96:
			goto tr303
		case 110:
			goto st303
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st303:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof303
		}
	st_case_303:
		switch lex.data[(lex.p)] {
		case 67:
			goto st304
		case 96:
			goto tr303
		case 99:
			goto st304
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st304:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof304
		}
	st_case_304:
		switch lex.data[(lex.p)] {
		case 69:
			goto st305
		case 96:
			goto tr303
		case 101:
			goto st305
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st305:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof305
		}
	st_case_305:
		switch lex.data[(lex.p)] {
		case 79:
			goto st306
		case 96:
			goto tr303
		case 111:
			goto st306
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st306:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof306
		}
	st_case_306:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr481
		case 96:
			goto tr303
		case 102:
			goto tr481
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st307:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof307
		}
	st_case_307:
		switch lex.data[(lex.p)] {
		case 65:
			goto st308
		case 96:
			goto tr303
		case 97:
			goto st308
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st308:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof308
		}
	st_case_308:
		switch lex.data[(lex.p)] {
		case 68:
			goto st309
		case 96:
			goto tr303
		case 100:
			goto st309
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st309:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof309
		}
	st_case_309:
		switch lex.data[(lex.p)] {
		case 79:
			goto st310
		case 96:
			goto tr303
		case 111:
			goto st310
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st310:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof310
		}
	st_case_310:
		switch lex.data[(lex.p)] {
		case 70:
			goto tr485
		case 96:
			goto tr303
		case 102:
			goto tr485
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st311:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof311
		}
	st_case_311:
		switch lex.data[(lex.p)] {
		case 69:
			goto st312
		case 96:
			goto tr303
		case 101:
			goto st312
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st312:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof312
		}
	st_case_312:
		switch lex.data[(lex.p)] {
		case 82:
			goto st313
		case 96:
			goto tr303
		case 114:
			goto st313
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st313:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof313
		}
	st_case_313:
		switch lex.data[(lex.p)] {
		case 70:
			goto st314
		case 96:
			goto tr303
		case 102:
			goto st314
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st314:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof314
		}
	st_case_314:
		switch lex.data[(lex.p)] {
		case 65:
			goto st315
		case 96:
			goto tr303
		case 97:
			goto st315
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st315:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof315
		}
	st_case_315:
		switch lex.data[(lex.p)] {
		case 67:
			goto st316
		case 96:
			goto tr303
		case 99:
			goto st316
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st316:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof316
		}
	st_case_316:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr491
		case 96:
			goto tr303
		case 101:
			goto tr491
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st317:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof317
		}
	st_case_317:
		switch lex.data[(lex.p)] {
		case 83:
			goto st318
		case 96:
			goto tr303
		case 115:
			goto st318
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st318:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof318
		}
	st_case_318:
		switch lex.data[(lex.p)] {
		case 69:
			goto st319
		case 96:
			goto tr303
		case 101:
			goto st319
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st319:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof319
		}
	st_case_319:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr494
		case 96:
			goto tr303
		case 116:
			goto tr494
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st320:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof320
		}
	st_case_320:
		switch lex.data[(lex.p)] {
		case 73:
			goto st321
		case 96:
			goto tr303
		case 105:
			goto st321
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st321:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof321
		}
	st_case_321:
		switch lex.data[(lex.p)] {
		case 83:
			goto st322
		case 96:
			goto tr303
		case 115:
			goto st322
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st322:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof322
		}
	st_case_322:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr497
		case 96:
			goto tr303
		case 116:
			goto tr497
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st323:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof323
		}
	st_case_323:
		switch lex.data[(lex.p)] {
		case 65:
			goto st324
		case 69:
			goto st331
		case 96:
			goto tr303
		case 97:
			goto st324
		case 101:
			goto st331
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st324:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof324
		}
	st_case_324:
		switch lex.data[(lex.p)] {
		case 77:
			goto st325
		case 96:
			goto tr303
		case 109:
			goto st325
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st325:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof325
		}
	st_case_325:
		switch lex.data[(lex.p)] {
		case 69:
			goto st326
		case 96:
			goto tr303
		case 101:
			goto st326
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st326:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof326
		}
	st_case_326:
		switch lex.data[(lex.p)] {
		case 83:
			goto st327
		case 96:
			goto tr303
		case 115:
			goto st327
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st327:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof327
		}
	st_case_327:
		switch lex.data[(lex.p)] {
		case 80:
			goto st328
		case 96:
			goto tr303
		case 112:
			goto st328
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st328:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof328
		}
	st_case_328:
		switch lex.data[(lex.p)] {
		case 65:
			goto st329
		case 96:
			goto tr303
		case 97:
			goto st329
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st329:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof329
		}
	st_case_329:
		switch lex.data[(lex.p)] {
		case 67:
			goto st330
		case 96:
			goto tr303
		case 99:
			goto st330
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st330:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof330
		}
	st_case_330:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr506
		case 96:
			goto tr303
		case 101:
			goto tr506
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st331:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof331
		}
	st_case_331:
		switch lex.data[(lex.p)] {
		case 87:
			goto tr507
		case 96:
			goto tr303
		case 119:
			goto tr507
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st332:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof332
		}
	st_case_332:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr508
		case 96:
			goto tr303
		case 114:
			goto tr508
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st333:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof333
		}
	st_case_333:
		switch lex.data[(lex.p)] {
		case 82:
			goto st334
		case 85:
			goto st346
		case 96:
			goto tr303
		case 114:
			goto st334
		case 117:
			goto st346
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st334:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof334
		}
	st_case_334:
		switch lex.data[(lex.p)] {
		case 73:
			goto st335
		case 79:
			goto st340
		case 96:
			goto tr303
		case 105:
			goto st335
		case 111:
			goto st340
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st335:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof335
		}
	st_case_335:
		switch lex.data[(lex.p)] {
		case 78:
			goto st336
		case 86:
			goto st337
		case 96:
			goto tr303
		case 110:
			goto st336
		case 118:
			goto st337
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st336:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof336
		}
	st_case_336:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr515
		case 96:
			goto tr303
		case 116:
			goto tr515
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st337:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof337
		}
	st_case_337:
		switch lex.data[(lex.p)] {
		case 65:
			goto st338
		case 96:
			goto tr303
		case 97:
			goto st338
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st338:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof338
		}
	st_case_338:
		switch lex.data[(lex.p)] {
		case 84:
			goto st339
		case 96:
			goto tr303
		case 116:
			goto st339
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st339:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof339
		}
	st_case_339:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr518
		case 96:
			goto tr303
		case 101:
			goto tr518
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st340:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof340
		}
	st_case_340:
		switch lex.data[(lex.p)] {
		case 84:
			goto st341
		case 96:
			goto tr303
		case 116:
			goto st341
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st341:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof341
		}
	st_case_341:
		switch lex.data[(lex.p)] {
		case 69:
			goto st342
		case 96:
			goto tr303
		case 101:
			goto st342
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st342:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof342
		}
	st_case_342:
		switch lex.data[(lex.p)] {
		case 67:
			goto st343
		case 96:
			goto tr303
		case 99:
			goto st343
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st343:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof343
		}
	st_case_343:
		switch lex.data[(lex.p)] {
		case 84:
			goto st344
		case 96:
			goto tr303
		case 116:
			goto st344
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st344:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof344
		}
	st_case_344:
		switch lex.data[(lex.p)] {
		case 69:
			goto st345
		case 96:
			goto tr303
		case 101:
			goto st345
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st345:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof345
		}
	st_case_345:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr524
		case 96:
			goto tr303
		case 100:
			goto tr524
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st346:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof346
		}
	st_case_346:
		switch lex.data[(lex.p)] {
		case 66:
			goto st347
		case 96:
			goto tr303
		case 98:
			goto st347
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st347:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof347
		}
	st_case_347:
		switch lex.data[(lex.p)] {
		case 76:
			goto st348
		case 96:
			goto tr303
		case 108:
			goto st348
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st348:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof348
		}
	st_case_348:
		switch lex.data[(lex.p)] {
		case 73:
			goto st349
		case 96:
			goto tr303
		case 105:
			goto st349
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st349:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof349
		}
	st_case_349:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr528
		case 96:
			goto tr303
		case 99:
			goto tr528
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st350:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof350
		}
	st_case_350:
		switch lex.data[(lex.p)] {
		case 69:
			goto st351
		case 96:
			goto tr303
		case 101:
			goto st351
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st351:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof351
		}
	st_case_351:
		switch lex.data[(lex.p)] {
		case 81:
			goto st352
		case 84:
			goto st361
		case 96:
			goto tr303
		case 113:
			goto st352
		case 116:
			goto st361
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st352:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof352
		}
	st_case_352:
		switch lex.data[(lex.p)] {
		case 85:
			goto st353
		case 96:
			goto tr303
		case 117:
			goto st353
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st353:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof353
		}
	st_case_353:
		switch lex.data[(lex.p)] {
		case 73:
			goto st354
		case 96:
			goto tr303
		case 105:
			goto st354
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st354:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof354
		}
	st_case_354:
		switch lex.data[(lex.p)] {
		case 82:
			goto st355
		case 96:
			goto tr303
		case 114:
			goto st355
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st355:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof355
		}
	st_case_355:
		switch lex.data[(lex.p)] {
		case 69:
			goto st356
		case 96:
			goto tr303
		case 101:
			goto st356
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st356:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof356
		}
	st_case_356:
		if lex.data[(lex.p)] == 95 {
			goto st357
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr536
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr536
				}
			case lex.data[(lex.p)] >= 91:
				goto tr536
			}
		default:
			goto tr536
		}
		goto tr211
	st357:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof357
		}
	st_case_357:
		switch lex.data[(lex.p)] {
		case 79:
			goto st358
		case 96:
			goto tr303
		case 111:
			goto st358
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st358:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof358
		}
	st_case_358:
		switch lex.data[(lex.p)] {
		case 78:
			goto st359
		case 96:
			goto tr303
		case 110:
			goto st359
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st359:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof359
		}
	st_case_359:
		switch lex.data[(lex.p)] {
		case 67:
			goto st360
		case 96:
			goto tr303
		case 99:
			goto st360
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st360:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof360
		}
	st_case_360:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr541
		case 96:
			goto tr303
		case 101:
			goto tr541
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st361:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof361
		}
	st_case_361:
		switch lex.data[(lex.p)] {
		case 85:
			goto st362
		case 96:
			goto tr303
		case 117:
			goto st362
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st362:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof362
		}
	st_case_362:
		switch lex.data[(lex.p)] {
		case 82:
			goto st363
		case 96:
			goto tr303
		case 114:
			goto st363
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st363:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof363
		}
	st_case_363:
		switch lex.data[(lex.p)] {
		case 78:
			goto tr544
		case 96:
			goto tr303
		case 110:
			goto tr544
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st364:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof364
		}
	st_case_364:
		switch lex.data[(lex.p)] {
		case 84:
			goto st365
		case 87:
			goto st369
		case 96:
			goto tr303
		case 116:
			goto st365
		case 119:
			goto st369
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st365:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof365
		}
	st_case_365:
		switch lex.data[(lex.p)] {
		case 65:
			goto st366
		case 96:
			goto tr303
		case 97:
			goto st366
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st366:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof366
		}
	st_case_366:
		switch lex.data[(lex.p)] {
		case 84:
			goto st367
		case 96:
			goto tr303
		case 116:
			goto st367
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st367:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof367
		}
	st_case_367:
		switch lex.data[(lex.p)] {
		case 73:
			goto st368
		case 96:
			goto tr303
		case 105:
			goto st368
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st368:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof368
		}
	st_case_368:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr550
		case 96:
			goto tr303
		case 99:
			goto tr550
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st369:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof369
		}
	st_case_369:
		switch lex.data[(lex.p)] {
		case 73:
			goto st370
		case 96:
			goto tr303
		case 105:
			goto st370
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st370:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof370
		}
	st_case_370:
		switch lex.data[(lex.p)] {
		case 84:
			goto st371
		case 96:
			goto tr303
		case 116:
			goto st371
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st371:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof371
		}
	st_case_371:
		switch lex.data[(lex.p)] {
		case 67:
			goto st372
		case 96:
			goto tr303
		case 99:
			goto st372
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st372:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof372
		}
	st_case_372:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr554
		case 96:
			goto tr303
		case 104:
			goto tr554
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st373:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof373
		}
	st_case_373:
		switch lex.data[(lex.p)] {
		case 72:
			goto st374
		case 82:
			goto st377
		case 96:
			goto tr303
		case 104:
			goto st374
		case 114:
			goto st377
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st374:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof374
		}
	st_case_374:
		switch lex.data[(lex.p)] {
		case 82:
			goto st375
		case 96:
			goto tr303
		case 114:
			goto st375
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st375:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof375
		}
	st_case_375:
		switch lex.data[(lex.p)] {
		case 79:
			goto st376
		case 96:
			goto tr303
		case 111:
			goto st376
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st376:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof376
		}
	st_case_376:
		switch lex.data[(lex.p)] {
		case 87:
			goto tr559
		case 96:
			goto tr303
		case 119:
			goto tr559
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st377:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof377
		}
	st_case_377:
		switch lex.data[(lex.p)] {
		case 65:
			goto st378
		case 89:
			goto tr561
		case 96:
			goto tr303
		case 97:
			goto st378
		case 121:
			goto tr561
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st378:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof378
		}
	st_case_378:
		switch lex.data[(lex.p)] {
		case 73:
			goto st379
		case 96:
			goto tr303
		case 105:
			goto st379
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st379:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof379
		}
	st_case_379:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr563
		case 96:
			goto tr303
		case 116:
			goto tr563
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st380:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof380
		}
	st_case_380:
		switch lex.data[(lex.p)] {
		case 78:
			goto st381
		case 83:
			goto st384
		case 96:
			goto tr303
		case 110:
			goto st381
		case 115:
			goto st384
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st381:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof381
		}
	st_case_381:
		switch lex.data[(lex.p)] {
		case 83:
			goto st382
		case 96:
			goto tr303
		case 115:
			goto st382
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st382:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof382
		}
	st_case_382:
		switch lex.data[(lex.p)] {
		case 69:
			goto st383
		case 96:
			goto tr303
		case 101:
			goto st383
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st383:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof383
		}
	st_case_383:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr568
		case 96:
			goto tr303
		case 116:
			goto tr568
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st384:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof384
		}
	st_case_384:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr569
		case 96:
			goto tr303
		case 101:
			goto tr569
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st385:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof385
		}
	st_case_385:
		switch lex.data[(lex.p)] {
		case 65:
			goto st386
		case 96:
			goto tr303
		case 97:
			goto st386
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st386:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof386
		}
	st_case_386:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr571
		case 96:
			goto tr303
		case 114:
			goto tr571
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st387:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof387
		}
	st_case_387:
		switch lex.data[(lex.p)] {
		case 72:
			goto st388
		case 96:
			goto tr303
		case 104:
			goto st388
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st388:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof388
		}
	st_case_388:
		switch lex.data[(lex.p)] {
		case 73:
			goto st389
		case 96:
			goto tr303
		case 105:
			goto st389
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st389:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof389
		}
	st_case_389:
		switch lex.data[(lex.p)] {
		case 76:
			goto st390
		case 96:
			goto tr303
		case 108:
			goto st390
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st390:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof390
		}
	st_case_390:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr575
		case 96:
			goto tr303
		case 101:
			goto tr575
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st391:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof391
		}
	st_case_391:
		switch lex.data[(lex.p)] {
		case 79:
			goto st392
		case 96:
			goto tr303
		case 111:
			goto st392
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st392:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof392
		}
	st_case_392:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr577
		case 96:
			goto tr303
		case 114:
			goto tr577
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st393:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof393
		}
	st_case_393:
		switch lex.data[(lex.p)] {
		case 73:
			goto st394
		case 96:
			goto tr303
		case 105:
			goto st394
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st394:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof394
		}
	st_case_394:
		switch lex.data[(lex.p)] {
		case 69:
			goto st395
		case 96:
			goto tr303
		case 101:
			goto st395
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st395:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof395
		}
	st_case_395:
		switch lex.data[(lex.p)] {
		case 76:
			goto st396
		case 96:
			goto tr303
		case 108:
			goto st396
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st396:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof396
		}
	st_case_396:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr581
		case 96:
			goto tr303
		case 100:
			goto tr581
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	tr581:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st397
	st397:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof397
		}
	st_case_397:
// line scanner/scanner.go:13597
		switch lex.data[(lex.p)] {
		case 10:
			goto st94
		case 13:
			goto st95
		case 32:
			goto st93
		case 70:
			goto st398
		case 96:
			goto tr582
		case 102:
			goto st398
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st93
				}
			default:
				goto tr582
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 91:
				if 58 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 64 {
					goto tr582
				}
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr582
				}
			default:
				goto tr582
			}
		default:
			goto tr582
		}
		goto tr211
	tr134:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st93
	st93:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof93
		}
	st_case_93:
// line scanner/scanner.go:13648
		switch lex.data[(lex.p)] {
		case 10:
			goto st94
		case 13:
			goto st95
		case 32:
			goto st93
		case 70:
			goto st96
		case 102:
			goto st96
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st93
		}
		goto tr129
	tr135:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st94
	st94:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof94
		}
	st_case_94:
// line scanner/scanner.go:13674
		switch lex.data[(lex.p)] {
		case 10:
			goto tr135
		case 13:
			goto tr136
		case 32:
			goto tr134
		case 70:
			goto tr137
		case 102:
			goto tr137
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr134
		}
		goto tr129
	tr136:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st95
	st95:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof95
		}
	st_case_95:
// line scanner/scanner.go:13700
		if lex.data[(lex.p)] == 10 {
			goto st94
		}
		goto tr129
	tr137:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st96
	st96:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof96
		}
	st_case_96:
// line scanner/scanner.go:13714
		switch lex.data[(lex.p)] {
		case 82:
			goto st97
		case 114:
			goto st97
		}
		goto tr129
	st97:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch lex.data[(lex.p)] {
		case 79:
			goto st98
		case 111:
			goto st98
		}
		goto tr129
	st98:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr140
		case 109:
			goto tr140
		}
		goto tr129
	st398:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof398
		}
	st_case_398:
		switch lex.data[(lex.p)] {
		case 82:
			goto st399
		case 96:
			goto tr303
		case 114:
			goto st399
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st399:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof399
		}
	st_case_399:
		switch lex.data[(lex.p)] {
		case 79:
			goto st400
		case 96:
			goto tr303
		case 111:
			goto st400
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st400:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof400
		}
	st_case_400:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr586
		case 96:
			goto tr303
		case 109:
			goto tr586
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st401:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof401
		}
	st_case_401:
		if lex.data[(lex.p)] == 61 {
			goto tr587
		}
		goto tr239
	st402:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof402
		}
	st_case_402:
		if lex.data[(lex.p)] == 95 {
			goto st403
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st403:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof403
		}
	st_case_403:
		switch lex.data[(lex.p)] {
		case 67:
			goto st404
		case 68:
			goto st410
		case 70:
			goto st414
		case 72:
			goto st427
		case 76:
			goto st439
		case 77:
			goto st444
		case 78:
			goto st451
		case 84:
			goto st461
		case 96:
			goto tr303
		case 99:
			goto st404
		case 100:
			goto st410
		case 102:
			goto st414
		case 104:
			goto st427
		case 108:
			goto st439
		case 109:
			goto st444
		case 110:
			goto st451
		case 116:
			goto st461
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st404:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof404
		}
	st_case_404:
		switch lex.data[(lex.p)] {
		case 76:
			goto st405
		case 96:
			goto tr303
		case 108:
			goto st405
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st405:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof405
		}
	st_case_405:
		switch lex.data[(lex.p)] {
		case 65:
			goto st406
		case 96:
			goto tr303
		case 97:
			goto st406
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st406:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof406
		}
	st_case_406:
		switch lex.data[(lex.p)] {
		case 83:
			goto st407
		case 96:
			goto tr303
		case 115:
			goto st407
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st407:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof407
		}
	st_case_407:
		switch lex.data[(lex.p)] {
		case 83:
			goto st408
		case 96:
			goto tr303
		case 115:
			goto st408
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
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
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st409:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof409
		}
	st_case_409:
		if lex.data[(lex.p)] == 95 {
			goto tr602
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st410:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof410
		}
	st_case_410:
		switch lex.data[(lex.p)] {
		case 73:
			goto st411
		case 96:
			goto tr303
		case 105:
			goto st411
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st411:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof411
		}
	st_case_411:
		switch lex.data[(lex.p)] {
		case 82:
			goto st412
		case 96:
			goto tr303
		case 114:
			goto st412
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st412:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof412
		}
	st_case_412:
		if lex.data[(lex.p)] == 95 {
			goto st413
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st413:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof413
		}
	st_case_413:
		if lex.data[(lex.p)] == 95 {
			goto tr606
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st414:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof414
		}
	st_case_414:
		switch lex.data[(lex.p)] {
		case 73:
			goto st415
		case 85:
			goto st419
		case 96:
			goto tr303
		case 105:
			goto st415
		case 117:
			goto st419
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st415:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof415
		}
	st_case_415:
		switch lex.data[(lex.p)] {
		case 76:
			goto st416
		case 96:
			goto tr303
		case 108:
			goto st416
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st416:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof416
		}
	st_case_416:
		switch lex.data[(lex.p)] {
		case 69:
			goto st417
		case 96:
			goto tr303
		case 101:
			goto st417
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st417:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof417
		}
	st_case_417:
		if lex.data[(lex.p)] == 95 {
			goto st418
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st418:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof418
		}
	st_case_418:
		if lex.data[(lex.p)] == 95 {
			goto tr612
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st419:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof419
		}
	st_case_419:
		switch lex.data[(lex.p)] {
		case 78:
			goto st420
		case 96:
			goto tr303
		case 110:
			goto st420
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st420:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof420
		}
	st_case_420:
		switch lex.data[(lex.p)] {
		case 67:
			goto st421
		case 96:
			goto tr303
		case 99:
			goto st421
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st421:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof421
		}
	st_case_421:
		switch lex.data[(lex.p)] {
		case 84:
			goto st422
		case 96:
			goto tr303
		case 116:
			goto st422
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st422:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof422
		}
	st_case_422:
		switch lex.data[(lex.p)] {
		case 73:
			goto st423
		case 96:
			goto tr303
		case 105:
			goto st423
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st423:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof423
		}
	st_case_423:
		switch lex.data[(lex.p)] {
		case 79:
			goto st424
		case 96:
			goto tr303
		case 111:
			goto st424
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st424:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof424
		}
	st_case_424:
		switch lex.data[(lex.p)] {
		case 78:
			goto st425
		case 96:
			goto tr303
		case 110:
			goto st425
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st425:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof425
		}
	st_case_425:
		if lex.data[(lex.p)] == 95 {
			goto st426
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st426:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof426
		}
	st_case_426:
		if lex.data[(lex.p)] == 95 {
			goto tr620
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st427:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof427
		}
	st_case_427:
		switch lex.data[(lex.p)] {
		case 65:
			goto st428
		case 96:
			goto tr303
		case 97:
			goto st428
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st428:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof428
		}
	st_case_428:
		switch lex.data[(lex.p)] {
		case 76:
			goto st429
		case 96:
			goto tr303
		case 108:
			goto st429
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st429:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof429
		}
	st_case_429:
		switch lex.data[(lex.p)] {
		case 84:
			goto st430
		case 96:
			goto tr303
		case 116:
			goto st430
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st430:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof430
		}
	st_case_430:
		if lex.data[(lex.p)] == 95 {
			goto st431
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st431:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof431
		}
	st_case_431:
		switch lex.data[(lex.p)] {
		case 67:
			goto st432
		case 96:
			goto tr303
		case 99:
			goto st432
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st432:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof432
		}
	st_case_432:
		switch lex.data[(lex.p)] {
		case 79:
			goto st433
		case 96:
			goto tr303
		case 111:
			goto st433
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st433:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof433
		}
	st_case_433:
		switch lex.data[(lex.p)] {
		case 77:
			goto st434
		case 96:
			goto tr303
		case 109:
			goto st434
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st434:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof434
		}
	st_case_434:
		switch lex.data[(lex.p)] {
		case 80:
			goto st435
		case 96:
			goto tr303
		case 112:
			goto st435
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st435:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof435
		}
	st_case_435:
		switch lex.data[(lex.p)] {
		case 73:
			goto st436
		case 96:
			goto tr303
		case 105:
			goto st436
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st436:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof436
		}
	st_case_436:
		switch lex.data[(lex.p)] {
		case 76:
			goto st437
		case 96:
			goto tr303
		case 108:
			goto st437
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st437:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof437
		}
	st_case_437:
		switch lex.data[(lex.p)] {
		case 69:
			goto st438
		case 96:
			goto tr303
		case 101:
			goto st438
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st438:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof438
		}
	st_case_438:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr632
		case 96:
			goto tr303
		case 114:
			goto tr632
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st439:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof439
		}
	st_case_439:
		switch lex.data[(lex.p)] {
		case 73:
			goto st440
		case 96:
			goto tr303
		case 105:
			goto st440
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st440:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof440
		}
	st_case_440:
		switch lex.data[(lex.p)] {
		case 78:
			goto st441
		case 96:
			goto tr303
		case 110:
			goto st441
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st441:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof441
		}
	st_case_441:
		switch lex.data[(lex.p)] {
		case 69:
			goto st442
		case 96:
			goto tr303
		case 101:
			goto st442
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st442:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof442
		}
	st_case_442:
		if lex.data[(lex.p)] == 95 {
			goto st443
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st443:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof443
		}
	st_case_443:
		if lex.data[(lex.p)] == 95 {
			goto tr637
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st444:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof444
		}
	st_case_444:
		switch lex.data[(lex.p)] {
		case 69:
			goto st445
		case 96:
			goto tr303
		case 101:
			goto st445
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st445:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof445
		}
	st_case_445:
		switch lex.data[(lex.p)] {
		case 84:
			goto st446
		case 96:
			goto tr303
		case 116:
			goto st446
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st446:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof446
		}
	st_case_446:
		switch lex.data[(lex.p)] {
		case 72:
			goto st447
		case 96:
			goto tr303
		case 104:
			goto st447
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st447:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof447
		}
	st_case_447:
		switch lex.data[(lex.p)] {
		case 79:
			goto st448
		case 96:
			goto tr303
		case 111:
			goto st448
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st448:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof448
		}
	st_case_448:
		switch lex.data[(lex.p)] {
		case 68:
			goto st449
		case 96:
			goto tr303
		case 100:
			goto st449
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st449:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof449
		}
	st_case_449:
		if lex.data[(lex.p)] == 95 {
			goto st450
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st450:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof450
		}
	st_case_450:
		if lex.data[(lex.p)] == 95 {
			goto tr644
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st451:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof451
		}
	st_case_451:
		switch lex.data[(lex.p)] {
		case 65:
			goto st452
		case 96:
			goto tr303
		case 97:
			goto st452
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st452:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof452
		}
	st_case_452:
		switch lex.data[(lex.p)] {
		case 77:
			goto st453
		case 96:
			goto tr303
		case 109:
			goto st453
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st453:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof453
		}
	st_case_453:
		switch lex.data[(lex.p)] {
		case 69:
			goto st454
		case 96:
			goto tr303
		case 101:
			goto st454
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st454:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof454
		}
	st_case_454:
		switch lex.data[(lex.p)] {
		case 83:
			goto st455
		case 96:
			goto tr303
		case 115:
			goto st455
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st455:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof455
		}
	st_case_455:
		switch lex.data[(lex.p)] {
		case 80:
			goto st456
		case 96:
			goto tr303
		case 112:
			goto st456
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st456:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof456
		}
	st_case_456:
		switch lex.data[(lex.p)] {
		case 65:
			goto st457
		case 96:
			goto tr303
		case 97:
			goto st457
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st457:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof457
		}
	st_case_457:
		switch lex.data[(lex.p)] {
		case 67:
			goto st458
		case 96:
			goto tr303
		case 99:
			goto st458
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st458:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof458
		}
	st_case_458:
		switch lex.data[(lex.p)] {
		case 69:
			goto st459
		case 96:
			goto tr303
		case 101:
			goto st459
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st459:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof459
		}
	st_case_459:
		if lex.data[(lex.p)] == 95 {
			goto st460
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st460:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof460
		}
	st_case_460:
		if lex.data[(lex.p)] == 95 {
			goto tr654
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st461:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof461
		}
	st_case_461:
		switch lex.data[(lex.p)] {
		case 82:
			goto st462
		case 96:
			goto tr303
		case 114:
			goto st462
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st462:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof462
		}
	st_case_462:
		switch lex.data[(lex.p)] {
		case 65:
			goto st463
		case 96:
			goto tr303
		case 97:
			goto st463
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st463:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof463
		}
	st_case_463:
		switch lex.data[(lex.p)] {
		case 73:
			goto st464
		case 96:
			goto tr303
		case 105:
			goto st464
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st464:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof464
		}
	st_case_464:
		switch lex.data[(lex.p)] {
		case 84:
			goto st465
		case 96:
			goto tr303
		case 116:
			goto st465
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st465:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof465
		}
	st_case_465:
		if lex.data[(lex.p)] == 95 {
			goto st466
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st466:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof466
		}
	st_case_466:
		if lex.data[(lex.p)] == 95 {
			goto tr660
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr303
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr303
				}
			case lex.data[(lex.p)] >= 91:
				goto tr303
			}
		default:
			goto tr303
		}
		goto tr211
	st467:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof467
		}
	st_case_467:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr661
		case 124:
			goto tr662
		}
		goto tr239
	tr141:
// line scanner/scanner.rl:391
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st468
	tr663:
// line scanner/scanner.rl:394
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				goto st121
			}
		}
		goto st468
	tr668:
// line scanner/scanner.rl:391
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st468
	tr670:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:391
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st468
	tr674:
// line scanner/scanner.rl:394
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				goto st121
			}
		}
		goto st468
	tr675:
// line scanner/scanner.rl:392
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 468
				goto _out
			}
		}
		goto st468
	tr676:
		lex.cs = 468
// line scanner/scanner.rl:393
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			lex.cs = 121
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st468:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof468
		}
	st_case_468:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:15870
		switch lex.data[(lex.p)] {
		case 10:
			goto tr142
		case 13:
			goto st471
		case 32:
			goto tr664
		case 45:
			goto st472
		case 96:
			goto tr663
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr664
				}
			default:
				goto tr663
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr663
				}
			case lex.data[(lex.p)] >= 91:
				goto tr663
			}
		default:
			goto tr663
		}
		goto st473
	tr664:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st469
	tr671:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st469
	st469:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof469
		}
	st_case_469:
// line scanner/scanner.go:15923
		switch lex.data[(lex.p)] {
		case 10:
			goto tr142
		case 13:
			goto st99
		case 32:
			goto tr664
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr664
		}
		goto tr668
	tr142:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st470
	tr672:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st470
	st470:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof470
		}
	st_case_470:
// line scanner/scanner.go:15953
		switch lex.data[(lex.p)] {
		case 10:
			goto tr672
		case 13:
			goto tr673
		case 32:
			goto tr671
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr671
		}
		goto tr670
	tr673:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st99
	st99:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof99
		}
	st_case_99:
// line scanner/scanner.go:15975
		if lex.data[(lex.p)] == 10 {
			goto tr142
		}
		goto tr141
	st471:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof471
		}
	st_case_471:
		if lex.data[(lex.p)] == 10 {
			goto tr142
		}
		goto tr674
	st472:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof472
		}
	st_case_472:
		if lex.data[(lex.p)] == 62 {
			goto tr675
		}
		goto tr674
	st473:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof473
		}
	st_case_473:
		if lex.data[(lex.p)] == 96 {
			goto tr676
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr676
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr676
				}
			case lex.data[(lex.p)] >= 91:
				goto tr676
			}
		default:
			goto tr676
		}
		goto st473
	tr679:
		lex.cs = 474
// line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 146:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(token)
				tok = T_ENCAPSED_AND_WHITESPACE
				lex.cs = 495
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr680:
		lex.cs = 474
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:398
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.cs = 495
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st474:
// line NONE:1
		lex.ts = 0

// line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof474
		}
	st_case_474:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:16069
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
			goto st476
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr677
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	tr677:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:398
		lex.act = 146
		goto st475
	tr681:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:398
		lex.act = 146
		goto st475
	st475:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof475
		}
	st_case_475:
// line scanner/scanner.go:16140
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
			goto st476
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr677
		}
		goto tr679
	tr682:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st476
	st476:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof476
		}
	st_case_476:
// line scanner/scanner.go:16195
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
			goto tr682
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr681
		}
		goto tr680
	tr143:
// line scanner/scanner.rl:407
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_CURLY_OPEN
			lex.call(477, 121)
			goto _out
		}
		goto st477
	tr689:
// line scanner/scanner.rl:409
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 477
					lex.top++
					goto st497
				}
			}
		}
		goto st477
	tr690:
// line scanner/scanner.rl:408
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(477, 512)
			goto _out
		}
		goto st477
	tr691:
		lex.cs = 477
// line NONE:1
		switch lex.act {
		case 147:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(token)
				tok = T_CURLY_OPEN
				lex.call(477, 121)
				goto _out
			}
		case 148:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(477, 512)
				goto _out
			}
		case 150:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(token)
				tok = T_ENCAPSED_AND_WHITESPACE

				if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
					lex.cs = 495
				}
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr692:
		lex.cs = 477
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:410
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 495
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr696:
		lex.cs = 477
// line scanner/scanner.rl:410
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 495
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st477:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof477
		}
	st_case_477:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:16324
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
			goto st478
		case 1403:
			goto st100
		case 1546:
			goto st480
		case 1572:
			goto st481
		case 1659:
			goto st482
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr685
		}
		goto st0
	st478:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof478
		}
	st_case_478:
		if lex.data[(lex.p)] == 123 {
			goto tr690
		}
		goto tr689
	st100:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		if lex.data[(lex.p)] == 36 {
			goto tr143
		}
		goto st0
	tr685:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:410
		lex.act = 150
		goto st479
	tr693:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:410
		lex.act = 150
		goto st479
	tr695:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:408
		lex.act = 148
		goto st479
	tr697:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:407
		lex.act = 147
		goto st479
	st479:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof479
		}
	st_case_479:
// line scanner/scanner.go:16432
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
			goto st480
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr685
		}
		goto tr691
	tr694:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st480
	st480:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof480
		}
	st_case_480:
// line scanner/scanner.go:16487
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
			goto tr694
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr693
		}
		goto tr692
	st481:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof481
		}
	st_case_481:
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
			goto tr690
		case 1546:
			goto st480
		case 1659:
			goto tr695
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr685
		}
		goto tr689
	st482:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof482
		}
	st_case_482:
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
			goto tr143
		case 1546:
			goto st480
		case 1572:
			goto tr697
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr685
		}
		goto tr696
	tr145:
// line scanner/scanner.rl:422
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_CURLY_OPEN
			lex.call(483, 121)
			goto _out
		}
		goto st483
	tr699:
		lex.cs = 483
// line scanner/scanner.rl:425
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('`'))
			lex.cs = 121
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr706:
// line scanner/scanner.rl:424
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 483
					lex.top++
					goto st497
				}
			}
		}
		goto st483
	tr707:
// line scanner/scanner.rl:423
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(483, 512)
			goto _out
		}
		goto st483
	tr708:
		lex.cs = 483
// line NONE:1
		switch lex.act {
		case 151:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(token)
				tok = T_CURLY_OPEN
				lex.call(483, 121)
				goto _out
			}
		case 152:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(483, 512)
				goto _out
			}
		case 154:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = TokenID(int('`'))
				lex.cs = 121
				{
					(lex.p)++
					goto _out
				}
			}
		case 155:
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
	tr709:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:426
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 483
				goto _out
			}
		}
		goto st483
	tr713:
// line scanner/scanner.rl:426
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 483
				goto _out
			}
		}
		goto st483
	st483:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof483
		}
	st_case_483:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:16721
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
			goto st484
		case 1888:
			goto tr699
		case 1915:
			goto st101
		case 2058:
			goto st486
		case 2084:
			goto st487
		case 2144:
			goto tr704
		case 2171:
			goto st488
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr701
		}
		goto st0
	st484:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof484
		}
	st_case_484:
		if lex.data[(lex.p)] == 123 {
			goto tr707
		}
		goto tr706
	st101:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if lex.data[(lex.p)] == 36 {
			goto tr145
		}
		goto st0
	tr701:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:426
		lex.act = 155
		goto st485
	tr704:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:425
		lex.act = 154
		goto st485
	tr710:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:426
		lex.act = 155
		goto st485
	tr712:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:423
		lex.act = 152
		goto st485
	tr714:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:422
		lex.act = 151
		goto st485
	st485:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof485
		}
	st_case_485:
// line scanner/scanner.go:16840
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
			goto st486
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr701
		}
		goto tr708
	tr711:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st486
	st486:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof486
		}
	st_case_486:
// line scanner/scanner.go:16895
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
			goto tr711
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr710
		}
		goto tr709
	st487:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof487
		}
	st_case_487:
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
			goto tr707
		case 2058:
			goto st486
		case 2171:
			goto tr712
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr701
		}
		goto tr706
	st488:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof488
		}
	st_case_488:
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
			goto tr145
		case 2058:
			goto st486
		case 2084:
			goto tr714
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr701
		}
		goto tr713
	tr146:
// line scanner/scanner.rl:434
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_CURLY_OPEN
			lex.call(489, 121)
			goto _out
		}
		goto st489
	tr715:
		lex.cs = 489
// line scanner/scanner.rl:437
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('"'))
			lex.cs = 121
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr723:
// line scanner/scanner.rl:436
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 489
					lex.top++
					goto st497
				}
			}
		}
		goto st489
	tr724:
// line scanner/scanner.rl:435
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(489, 512)
			goto _out
		}
		goto st489
	tr725:
		lex.cs = 489
// line NONE:1
		switch lex.act {
		case 156:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(token)
				tok = T_CURLY_OPEN
				lex.call(489, 121)
				goto _out
			}
		case 157:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(489, 512)
				goto _out
			}
		case 159:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(token)
				tok = TokenID(int('"'))
				lex.cs = 121
				{
					(lex.p)++
					goto _out
				}
			}
		case 160:
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
	tr726:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:438
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 489
				goto _out
			}
		}
		goto st489
	tr730:
// line scanner/scanner.rl:438
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 489
				goto _out
			}
		}
		goto st489
	st489:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof489
		}
	st_case_489:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:17129
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
			goto tr715
		case 2340:
			goto st490
		case 2427:
			goto st102
		case 2570:
			goto st492
		case 2594:
			goto tr720
		case 2596:
			goto st493
		case 2683:
			goto st494
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr718
		}
		goto st0
	st490:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof490
		}
	st_case_490:
		if lex.data[(lex.p)] == 123 {
			goto tr724
		}
		goto tr723
	st102:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		if lex.data[(lex.p)] == 36 {
			goto tr146
		}
		goto st0
	tr718:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:438
		lex.act = 160
		goto st491
	tr720:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:437
		lex.act = 159
		goto st491
	tr727:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:438
		lex.act = 160
		goto st491
	tr729:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:435
		lex.act = 157
		goto st491
	tr731:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:434
		lex.act = 156
		goto st491
	st491:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof491
		}
	st_case_491:
// line scanner/scanner.go:17248
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
			goto st492
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr718
		}
		goto tr725
	tr728:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st492
	st492:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof492
		}
	st_case_492:
// line scanner/scanner.go:17303
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
			goto tr728
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr727
		}
		goto tr726
	st493:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof493
		}
	st_case_493:
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
			goto tr724
		case 2570:
			goto st492
		case 2683:
			goto tr729
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr718
		}
		goto tr723
	st494:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof494
		}
	st_case_494:
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
			goto tr146
		case 2570:
			goto st492
		case 2596:
			goto tr731
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr718
		}
		goto tr730
	tr733:
		lex.cs = 495
// line scanner/scanner.rl:446
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_END_HEREDOC
			lex.cs = 121
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st495:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof495
		}
	st_case_495:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:17482
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
		goto st496
	st496:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof496
		}
	st_case_496:
		if lex.data[(lex.p)] == 96 {
			goto tr733
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr733
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr733
				}
			case lex.data[(lex.p)] >= 91:
				goto tr733
			}
		default:
			goto tr733
		}
		goto st496
	tr147:
// line scanner/scanner.rl:465
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st497
	tr148:
// line scanner/scanner.rl:462
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 497
				goto _out
			}
		}
		goto st497
	tr734:
// line scanner/scanner.rl:465
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st497
	tr738:
// line scanner/scanner.rl:464
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('['))
			lex.call(497, 502)
			goto _out
		}
		goto st497
	tr739:
// line scanner/scanner.rl:465
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
		goto st497
	tr741:
// line scanner/scanner.rl:461
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 497
				goto _out
			}
		}
		goto st497
	tr743:
// line scanner/scanner.rl:463
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 497
				goto _out
			}
		}
		goto st497
	st497:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof497
		}
	st_case_497:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:17574
		switch lex.data[(lex.p)] {
		case 36:
			goto st498
		case 45:
			goto tr736
		case 91:
			goto tr738
		case 96:
			goto tr734
		}
		switch {
		case lex.data[(lex.p)] < 92:
			if lex.data[(lex.p)] <= 64 {
				goto tr734
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr734
			}
		default:
			goto tr734
		}
		goto st501
	st498:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof498
		}
	st_case_498:
		if lex.data[(lex.p)] == 96 {
			goto tr739
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr739
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr739
			}
		default:
			goto tr739
		}
		goto st499
	st499:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof499
		}
	st_case_499:
		if lex.data[(lex.p)] == 96 {
			goto tr741
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr741
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr741
				}
			case lex.data[(lex.p)] >= 91:
				goto tr741
			}
		default:
			goto tr741
		}
		goto st499
	tr736:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st500
	st500:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof500
		}
	st_case_500:
// line scanner/scanner.go:17655
		if lex.data[(lex.p)] == 62 {
			goto st103
		}
		goto tr739
	st103:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		if lex.data[(lex.p)] == 96 {
			goto tr147
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr147
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr147
			}
		default:
			goto tr147
		}
		goto tr148
	st501:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof501
		}
	st_case_501:
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
		goto st501
	tr149:
// line scanner/scanner.rl:469
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(token)
			tok = T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 502
				goto _out
			}
		}
		goto st502
	tr744:
// line scanner/scanner.rl:475
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st502
	tr745:
// line scanner/scanner.rl:472
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st502
	tr748:
// line scanner/scanner.rl:473
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 502
				goto _out
			}
		}
		goto st502
	tr752:
// line scanner/scanner.rl:474
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(']'))
			lex.ret(2)
			goto _out
		}
		goto st502
	tr753:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:472
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st502
	tr754:
// line scanner/scanner.rl:475
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st502
	tr755:
// line scanner/scanner.rl:473
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 502
				goto _out
			}
		}
		goto st502
	tr757:
// line scanner/scanner.rl:470
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_VARIABLE
			{
				(lex.p)++
				lex.cs = 502
				goto _out
			}
		}
		goto st502
	tr758:
// line scanner/scanner.rl:469
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 502
				goto _out
			}
		}
		goto st502
	tr762:
// line scanner/scanner.rl:471
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(token)
			tok = T_STRING
			{
				(lex.p)++
				lex.cs = 502
				goto _out
			}
		}
		goto st502
	st502:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof502
		}
	st_case_502:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:17787
		switch lex.data[(lex.p)] {
		case 10:
			goto st503
		case 13:
			goto st504
		case 32:
			goto tr745
		case 33:
			goto tr748
		case 35:
			goto tr745
		case 36:
			goto st505
		case 39:
			goto tr745
		case 48:
			goto tr750
		case 92:
			goto tr745
		case 93:
			goto tr752
		case 96:
			goto tr744
		case 124:
			goto tr748
		case 126:
			goto tr748
		}
		switch {
		case lex.data[(lex.p)] < 37:
			switch {
			case lex.data[(lex.p)] < 9:
				if lex.data[(lex.p)] <= 8 {
					goto tr744
				}
			case lex.data[(lex.p)] > 12:
				if 14 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 34 {
					goto tr744
				}
			default:
				goto tr745
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 58:
				if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
					goto tr150
				}
			case lex.data[(lex.p)] > 64:
				switch {
				case lex.data[(lex.p)] > 94:
					if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
						goto tr744
					}
				case lex.data[(lex.p)] >= 91:
					goto tr748
				}
			default:
				goto tr748
			}
		default:
			goto tr748
		}
		goto st511
	st503:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof503
		}
	st_case_503:
		goto tr753
	st504:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof504
		}
	st_case_504:
		if lex.data[(lex.p)] == 10 {
			goto st503
		}
		goto tr754
	st505:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof505
		}
	st_case_505:
		if lex.data[(lex.p)] == 96 {
			goto tr755
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr755
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr755
			}
		default:
			goto tr755
		}
		goto st506
	st506:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof506
		}
	st_case_506:
		if lex.data[(lex.p)] == 96 {
			goto tr757
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr757
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr757
				}
			case lex.data[(lex.p)] >= 91:
				goto tr757
			}
		default:
			goto tr757
		}
		goto st506
	tr750:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st507
	st507:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof507
		}
	st_case_507:
// line scanner/scanner.go:17924
		switch lex.data[(lex.p)] {
		case 95:
			goto st104
		case 98:
			goto st105
		case 120:
			goto st106
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr150
		}
		goto tr758
	tr150:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st508
	st508:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof508
		}
	st_case_508:
// line scanner/scanner.go:17947
		if lex.data[(lex.p)] == 95 {
			goto st104
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr150
		}
		goto tr758
	st104:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr150
		}
		goto tr149
	st105:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr151
		}
		goto tr149
	tr151:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st509
	st509:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof509
		}
	st_case_509:
// line scanner/scanner.go:17983
		if lex.data[(lex.p)] == 95 {
			goto st105
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr151
		}
		goto tr758
	st106:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr152
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr152
			}
		default:
			goto tr152
		}
		goto tr149
	tr152:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st510
	st510:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof510
		}
	st_case_510:
// line scanner/scanner.go:18019
		if lex.data[(lex.p)] == 95 {
			goto st106
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr152
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr152
			}
		default:
			goto tr152
		}
		goto tr758
	st511:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof511
		}
	st_case_511:
		if lex.data[(lex.p)] == 96 {
			goto tr762
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr762
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr762
				}
			case lex.data[(lex.p)] >= 91:
				goto tr762
			}
		default:
			goto tr762
		}
		goto st511
	tr153:
		lex.cs = 512
// line scanner/scanner.rl:483
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	tr155:
		lex.cs = 512
// line scanner/scanner.rl:482
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(token)
			tok = T_STRING_VARNAME
			lex.cs = 121
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr763:
		lex.cs = 512
// line scanner/scanner.rl:483
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	tr765:
		lex.cs = 512
// line scanner/scanner.rl:483
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	st512:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof512
		}
	st_case_512:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:18098
		if lex.data[(lex.p)] == 96 {
			goto tr763
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr763
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr763
			}
		default:
			goto tr763
		}
		goto tr764
	tr764:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st513
	st513:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof513
		}
	st_case_513:
// line scanner/scanner.go:18125
		switch lex.data[(lex.p)] {
		case 91:
			goto tr155
		case 96:
			goto tr765
		case 125:
			goto tr155
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr765
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr765
				}
			case lex.data[(lex.p)] >= 92:
				goto tr765
			}
		default:
			goto tr765
		}
		goto st107
	st107:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch lex.data[(lex.p)] {
		case 91:
			goto tr155
		case 96:
			goto tr153
		case 125:
			goto tr155
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr153
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr153
				}
			case lex.data[(lex.p)] >= 92:
				goto tr153
			}
		default:
			goto tr153
		}
		goto st107
	tr156:
// line scanner/scanner.rl:487
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st514
	tr766:
		lex.cs = 514
// line scanner/scanner.rl:489
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	tr769:
		lex.cs = 514
// line scanner/scanner.rl:488
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int('('))
			lex.cs = 518
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr770:
// line scanner/scanner.rl:487
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st514
	tr772:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:487
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st514
	tr776:
		lex.cs = 514
// line scanner/scanner.rl:489
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	st514:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof514
		}
	st_case_514:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:18232
		switch lex.data[(lex.p)] {
		case 10:
			goto tr157
		case 13:
			goto st517
		case 32:
			goto tr767
		case 40:
			goto tr769
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr767
		}
		goto tr766
	tr767:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st515
	tr773:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st515
	st515:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof515
		}
	st_case_515:
// line scanner/scanner.go:18264
		switch lex.data[(lex.p)] {
		case 10:
			goto tr157
		case 13:
			goto st108
		case 32:
			goto tr767
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr767
		}
		goto tr770
	tr157:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st516
	tr774:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st516
	st516:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof516
		}
	st_case_516:
// line scanner/scanner.go:18294
		switch lex.data[(lex.p)] {
		case 10:
			goto tr774
		case 13:
			goto tr775
		case 32:
			goto tr773
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr773
		}
		goto tr772
	tr775:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st108
	st108:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof108
		}
	st_case_108:
// line scanner/scanner.go:18316
		if lex.data[(lex.p)] == 10 {
			goto tr157
		}
		goto tr156
	st517:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof517
		}
	st_case_517:
		if lex.data[(lex.p)] == 10 {
			goto tr157
		}
		goto tr776
	tr158:
// line scanner/scanner.rl:493
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st518
	tr777:
		lex.cs = 518
// line scanner/scanner.rl:495
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	tr780:
		lex.cs = 518
// line scanner/scanner.rl:494
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(')'))
			lex.cs = 522
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr781:
// line scanner/scanner.rl:493
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st518
	tr783:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:493
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st518
	tr787:
		lex.cs = 518
// line scanner/scanner.rl:495
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	st518:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof518
		}
	st_case_518:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:18379
		switch lex.data[(lex.p)] {
		case 10:
			goto tr159
		case 13:
			goto st521
		case 32:
			goto tr778
		case 41:
			goto tr780
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr778
		}
		goto tr777
	tr778:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st519
	tr784:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st519
	st519:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof519
		}
	st_case_519:
// line scanner/scanner.go:18411
		switch lex.data[(lex.p)] {
		case 10:
			goto tr159
		case 13:
			goto st109
		case 32:
			goto tr778
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr778
		}
		goto tr781
	tr159:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st520
	tr785:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st520
	st520:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof520
		}
	st_case_520:
// line scanner/scanner.go:18441
		switch lex.data[(lex.p)] {
		case 10:
			goto tr785
		case 13:
			goto tr786
		case 32:
			goto tr784
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr784
		}
		goto tr783
	tr786:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st109
	st109:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof109
		}
	st_case_109:
// line scanner/scanner.go:18463
		if lex.data[(lex.p)] == 10 {
			goto tr159
		}
		goto tr158
	st521:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof521
		}
	st_case_521:
		if lex.data[(lex.p)] == 10 {
			goto tr159
		}
		goto tr787
	tr160:
// line scanner/scanner.rl:499
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st522
	tr788:
		lex.cs = 522
// line scanner/scanner.rl:501
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	tr791:
		lex.cs = 522
// line scanner/scanner.rl:500
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(token)
			tok = TokenID(int(';'))
			lex.cs = 526
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr792:
// line scanner/scanner.rl:499
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st522
	tr794:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:499
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)
		}
		goto st522
	tr798:
		lex.cs = 522
// line scanner/scanner.rl:501
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 121
		}
		goto _again
	st522:
// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof522
		}
	st_case_522:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:18526
		switch lex.data[(lex.p)] {
		case 10:
			goto tr161
		case 13:
			goto st525
		case 32:
			goto tr789
		case 59:
			goto tr791
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr789
		}
		goto tr788
	tr789:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st523
	tr795:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st523
	st523:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof523
		}
	st_case_523:
// line scanner/scanner.go:18558
		switch lex.data[(lex.p)] {
		case 10:
			goto tr161
		case 13:
			goto st110
		case 32:
			goto tr789
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr789
		}
		goto tr792
	tr161:
// line NONE:1
		lex.te = (lex.p) + 1

		goto st524
	tr796:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st524
	st524:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof524
		}
	st_case_524:
// line scanner/scanner.go:18588
		switch lex.data[(lex.p)] {
		case 10:
			goto tr796
		case 13:
			goto tr797
		case 32:
			goto tr795
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr795
		}
		goto tr794
	tr797:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st110
	st110:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof110
		}
	st_case_110:
// line scanner/scanner.go:18610
		if lex.data[(lex.p)] == 10 {
			goto tr161
		}
		goto tr160
	st525:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof525
		}
	st_case_525:
		if lex.data[(lex.p)] == 10 {
			goto tr161
		}
		goto tr798
	tr801:
// line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 186:
			{
				(lex.p) = (lex.te) - 1
				lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
			}
		}

		goto st526
	tr802:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:505
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
		}
		goto st526
	st526:
// line NONE:1
		lex.ts = 0

// line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof526
		}
	st_case_526:
// line NONE:1
		lex.ts = (lex.p)

// line scanner/scanner.go:18657
		if lex.data[(lex.p)] == 10 {
			goto st528
		}
		goto tr799
	tr799:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:505
		lex.act = 186
		goto st527
	tr803:
// line NONE:1
		lex.te = (lex.p) + 1

// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
// line scanner/scanner.rl:505
		lex.act = 186
		goto st527
	st527:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof527
		}
	st_case_527:
// line scanner/scanner.go:18683
		if lex.data[(lex.p)] == 10 {
			goto st528
		}
		goto tr799
	tr804:
// line scanner/scanner.rl:66
		lex.NewLines.Append(lex.p)
		goto st528
	st528:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof528
		}
	st_case_528:
// line scanner/scanner.go:18697
		if lex.data[(lex.p)] == 10 {
			goto tr804
		}
		goto tr803
	st_out:
	_test_eof111:
		lex.cs = 111
		goto _test_eof
	_test_eof112:
		lex.cs = 112
		goto _test_eof
	_test_eof1:
		lex.cs = 1
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
	_test_eof117:
		lex.cs = 117
		goto _test_eof
	_test_eof118:
		lex.cs = 118
		goto _test_eof
	_test_eof119:
		lex.cs = 119
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
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
	_test_eof6:
		lex.cs = 6
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
	_test_eof132:
		lex.cs = 132
		goto _test_eof
	_test_eof133:
		lex.cs = 133
		goto _test_eof
	_test_eof134:
		lex.cs = 134
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof135:
		lex.cs = 135
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
	_test_eof66:
		lex.cs = 66
		goto _test_eof
	_test_eof136:
		lex.cs = 136
		goto _test_eof
	_test_eof137:
		lex.cs = 137
		goto _test_eof
	_test_eof138:
		lex.cs = 138
		goto _test_eof
	_test_eof139:
		lex.cs = 139
		goto _test_eof
	_test_eof140:
		lex.cs = 140
		goto _test_eof
	_test_eof67:
		lex.cs = 67
		goto _test_eof
	_test_eof141:
		lex.cs = 141
		goto _test_eof
	_test_eof68:
		lex.cs = 68
		goto _test_eof
	_test_eof69:
		lex.cs = 69
		goto _test_eof
	_test_eof142:
		lex.cs = 142
		goto _test_eof
	_test_eof70:
		lex.cs = 70
		goto _test_eof
	_test_eof143:
		lex.cs = 143
		goto _test_eof
	_test_eof71:
		lex.cs = 71
		goto _test_eof
	_test_eof72:
		lex.cs = 72
		goto _test_eof
	_test_eof73:
		lex.cs = 73
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
	_test_eof74:
		lex.cs = 74
		goto _test_eof
	_test_eof75:
		lex.cs = 75
		goto _test_eof
	_test_eof147:
		lex.cs = 147
		goto _test_eof
	_test_eof76:
		lex.cs = 76
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
	_test_eof77:
		lex.cs = 77
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
	_test_eof151:
		lex.cs = 151
		goto _test_eof
	_test_eof152:
		lex.cs = 152
		goto _test_eof
	_test_eof81:
		lex.cs = 81
		goto _test_eof
	_test_eof153:
		lex.cs = 153
		goto _test_eof
	_test_eof154:
		lex.cs = 154
		goto _test_eof
	_test_eof82:
		lex.cs = 82
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
	_test_eof155:
		lex.cs = 155
		goto _test_eof
	_test_eof86:
		lex.cs = 86
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
	_test_eof90:
		lex.cs = 90
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
	_test_eof176:
		lex.cs = 176
		goto _test_eof
	_test_eof177:
		lex.cs = 177
		goto _test_eof
	_test_eof91:
		lex.cs = 91
		goto _test_eof
	_test_eof92:
		lex.cs = 92
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
	_test_eof396:
		lex.cs = 396
		goto _test_eof
	_test_eof397:
		lex.cs = 397
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
	_test_eof96:
		lex.cs = 96
		goto _test_eof
	_test_eof97:
		lex.cs = 97
		goto _test_eof
	_test_eof98:
		lex.cs = 98
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
	_test_eof469:
		lex.cs = 469
		goto _test_eof
	_test_eof470:
		lex.cs = 470
		goto _test_eof
	_test_eof99:
		lex.cs = 99
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
	_test_eof477:
		lex.cs = 477
		goto _test_eof
	_test_eof478:
		lex.cs = 478
		goto _test_eof
	_test_eof100:
		lex.cs = 100
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
	_test_eof483:
		lex.cs = 483
		goto _test_eof
	_test_eof484:
		lex.cs = 484
		goto _test_eof
	_test_eof101:
		lex.cs = 101
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
	_test_eof102:
		lex.cs = 102
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
	_test_eof499:
		lex.cs = 499
		goto _test_eof
	_test_eof500:
		lex.cs = 500
		goto _test_eof
	_test_eof103:
		lex.cs = 103
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
	_test_eof507:
		lex.cs = 507
		goto _test_eof
	_test_eof508:
		lex.cs = 508
		goto _test_eof
	_test_eof104:
		lex.cs = 104
		goto _test_eof
	_test_eof105:
		lex.cs = 105
		goto _test_eof
	_test_eof509:
		lex.cs = 509
		goto _test_eof
	_test_eof106:
		lex.cs = 106
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
	_test_eof513:
		lex.cs = 513
		goto _test_eof
	_test_eof107:
		lex.cs = 107
		goto _test_eof
	_test_eof514:
		lex.cs = 514
		goto _test_eof
	_test_eof515:
		lex.cs = 515
		goto _test_eof
	_test_eof516:
		lex.cs = 516
		goto _test_eof
	_test_eof108:
		lex.cs = 108
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
	_test_eof520:
		lex.cs = 520
		goto _test_eof
	_test_eof109:
		lex.cs = 109
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
	_test_eof524:
		lex.cs = 524
		goto _test_eof
	_test_eof110:
		lex.cs = 110
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
	_test_eof528:
		lex.cs = 528
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 112:
				goto tr164
			case 1:
				goto tr0
			case 113:
				goto tr165
			case 115:
				goto tr169
			case 116:
				goto tr171
			case 117:
				goto tr169
			case 118:
				goto tr169
			case 119:
				goto tr176
			case 2:
				goto tr3
			case 3:
				goto tr3
			case 4:
				goto tr3
			case 120:
				goto tr179
			case 5:
				goto tr3
			case 122:
				goto tr232
			case 123:
				goto tr234
			case 6:
				goto tr9
			case 124:
				goto tr238
			case 125:
				goto tr239
			case 126:
				goto tr241
			case 127:
				goto tr243
			case 7:
				goto tr11
			case 8:
				goto tr11
			case 9:
				goto tr11
			case 10:
				goto tr11
			case 128:
				goto tr244
			case 129:
				goto tr246
			case 130:
				goto tr239
			case 131:
				goto tr250
			case 132:
				goto tr239
			case 133:
				goto tr239
			case 134:
				goto tr238
			case 11:
				goto tr18
			case 12:
				goto tr18
			case 135:
				goto tr239
			case 13:
				goto tr22
			case 14:
				goto tr22
			case 15:
				goto tr22
			case 16:
				goto tr22
			case 17:
				goto tr22
			case 18:
				goto tr22
			case 19:
				goto tr22
			case 20:
				goto tr22
			case 21:
				goto tr22
			case 22:
				goto tr22
			case 23:
				goto tr22
			case 24:
				goto tr22
			case 25:
				goto tr22
			case 26:
				goto tr22
			case 27:
				goto tr22
			case 28:
				goto tr22
			case 29:
				goto tr22
			case 30:
				goto tr22
			case 31:
				goto tr22
			case 32:
				goto tr22
			case 33:
				goto tr22
			case 34:
				goto tr22
			case 35:
				goto tr22
			case 36:
				goto tr22
			case 37:
				goto tr22
			case 38:
				goto tr22
			case 39:
				goto tr22
			case 40:
				goto tr22
			case 41:
				goto tr22
			case 42:
				goto tr22
			case 43:
				goto tr22
			case 44:
				goto tr22
			case 45:
				goto tr22
			case 46:
				goto tr22
			case 47:
				goto tr22
			case 48:
				goto tr22
			case 49:
				goto tr22
			case 50:
				goto tr22
			case 51:
				goto tr22
			case 52:
				goto tr22
			case 53:
				goto tr22
			case 54:
				goto tr22
			case 55:
				goto tr22
			case 56:
				goto tr22
			case 57:
				goto tr22
			case 58:
				goto tr22
			case 59:
				goto tr22
			case 60:
				goto tr22
			case 61:
				goto tr22
			case 62:
				goto tr22
			case 63:
				goto tr22
			case 64:
				goto tr22
			case 65:
				goto tr22
			case 66:
				goto tr22
			case 136:
				goto tr239
			case 137:
				goto tr256
			case 138:
				goto tr239
			case 139:
				goto tr239
			case 140:
				goto tr239
			case 67:
				goto tr22
			case 141:
				goto tr265
			case 68:
				goto tr11
			case 69:
				goto tr11
			case 142:
				goto tr265
			case 70:
				goto tr87
			case 143:
				goto tr239
			case 71:
				goto tr22
			case 72:
				goto tr22
			case 73:
				goto tr22
			case 144:
				goto tr269
			case 145:
				goto tr265
			case 146:
				goto tr269
			case 74:
				goto tr96
			case 75:
				goto tr11
			case 147:
				goto tr274
			case 76:
				goto tr11
			case 148:
				goto tr275
			case 149:
				goto tr239
			case 150:
				goto tr239
			case 77:
				goto tr22
			case 78:
				goto tr22
			case 79:
				goto tr22
			case 80:
				goto tr22
			case 151:
				goto tr277
			case 152:
				goto tr279
			case 81:
				goto tr109
			case 153:
				goto tr239
			case 154:
				goto tr283
			case 82:
				goto tr11
			case 83:
				goto tr11
			case 84:
				goto tr11
			case 85:
				goto tr11
			case 155:
				goto tr285
			case 86:
				goto tr11
			case 87:
				goto tr11
			case 88:
				goto tr11
			case 89:
				goto tr11
			case 156:
				goto tr286
			case 157:
				goto tr239
			case 158:
				goto tr290
			case 159:
				goto tr239
			case 160:
				goto tr294
			case 161:
				goto tr239
			case 162:
				goto tr298
			case 163:
				goto tr300
			case 90:
				goto tr125
			case 164:
				goto tr301
			case 165:
				goto tr303
			case 166:
				goto tr11
			case 167:
				goto tr303
			case 168:
				goto tr303
			case 169:
				goto tr303
			case 170:
				goto tr303
			case 171:
				goto tr303
			case 172:
				goto tr303
			case 173:
				goto tr303
			case 174:
				goto tr303
			case 175:
				goto tr303
			case 176:
				goto tr303
			case 177:
				goto tr303
			case 91:
				goto tr127
			case 92:
				goto tr127
			case 178:
				goto tr303
			case 179:
				goto tr303
			case 180:
				goto tr303
			case 181:
				goto tr303
			case 182:
				goto tr303
			case 183:
				goto tr303
			case 184:
				goto tr303
			case 185:
				goto tr303
			case 186:
				goto tr303
			case 187:
				goto tr303
			case 188:
				goto tr303
			case 189:
				goto tr303
			case 190:
				goto tr303
			case 191:
				goto tr303
			case 192:
				goto tr303
			case 193:
				goto tr303
			case 194:
				goto tr303
			case 195:
				goto tr303
			case 196:
				goto tr303
			case 197:
				goto tr303
			case 198:
				goto tr303
			case 199:
				goto tr303
			case 200:
				goto tr303
			case 201:
				goto tr303
			case 202:
				goto tr303
			case 203:
				goto tr303
			case 204:
				goto tr303
			case 205:
				goto tr303
			case 206:
				goto tr303
			case 207:
				goto tr303
			case 208:
				goto tr303
			case 209:
				goto tr303
			case 210:
				goto tr303
			case 211:
				goto tr303
			case 212:
				goto tr303
			case 213:
				goto tr303
			case 214:
				goto tr303
			case 215:
				goto tr303
			case 216:
				goto tr303
			case 217:
				goto tr303
			case 218:
				goto tr303
			case 219:
				goto tr303
			case 220:
				goto tr303
			case 221:
				goto tr303
			case 222:
				goto tr303
			case 223:
				goto tr303
			case 224:
				goto tr303
			case 225:
				goto tr303
			case 226:
				goto tr383
			case 227:
				goto tr303
			case 228:
				goto tr303
			case 229:
				goto tr303
			case 230:
				goto tr303
			case 231:
				goto tr303
			case 232:
				goto tr303
			case 233:
				goto tr303
			case 234:
				goto tr303
			case 235:
				goto tr303
			case 236:
				goto tr303
			case 237:
				goto tr303
			case 238:
				goto tr303
			case 239:
				goto tr303
			case 240:
				goto tr303
			case 241:
				goto tr403
			case 242:
				goto tr303
			case 243:
				goto tr303
			case 244:
				goto tr303
			case 245:
				goto tr303
			case 246:
				goto tr303
			case 247:
				goto tr303
			case 248:
				goto tr303
			case 249:
				goto tr303
			case 250:
				goto tr303
			case 251:
				goto tr303
			case 252:
				goto tr303
			case 253:
				goto tr303
			case 254:
				goto tr303
			case 255:
				goto tr303
			case 256:
				goto tr303
			case 257:
				goto tr303
			case 258:
				goto tr303
			case 259:
				goto tr303
			case 260:
				goto tr303
			case 261:
				goto tr303
			case 262:
				goto tr303
			case 263:
				goto tr303
			case 264:
				goto tr303
			case 265:
				goto tr303
			case 266:
				goto tr303
			case 267:
				goto tr432
			case 268:
				goto tr303
			case 269:
				goto tr303
			case 270:
				goto tr436
			case 271:
				goto tr303
			case 272:
				goto tr303
			case 273:
				goto tr303
			case 274:
				goto tr303
			case 275:
				goto tr303
			case 276:
				goto tr303
			case 277:
				goto tr303
			case 278:
				goto tr303
			case 279:
				goto tr303
			case 280:
				goto tr303
			case 281:
				goto tr303
			case 282:
				goto tr303
			case 283:
				goto tr303
			case 284:
				goto tr303
			case 285:
				goto tr303
			case 286:
				goto tr303
			case 287:
				goto tr303
			case 288:
				goto tr303
			case 289:
				goto tr303
			case 290:
				goto tr303
			case 291:
				goto tr303
			case 292:
				goto tr303
			case 293:
				goto tr303
			case 294:
				goto tr303
			case 295:
				goto tr468
			case 296:
				goto tr303
			case 297:
				goto tr303
			case 298:
				goto tr303
			case 299:
				goto tr303
			case 300:
				goto tr303
			case 301:
				goto tr303
			case 302:
				goto tr303
			case 303:
				goto tr303
			case 304:
				goto tr303
			case 305:
				goto tr303
			case 306:
				goto tr303
			case 307:
				goto tr303
			case 308:
				goto tr303
			case 309:
				goto tr303
			case 310:
				goto tr303
			case 311:
				goto tr303
			case 312:
				goto tr303
			case 313:
				goto tr303
			case 314:
				goto tr303
			case 315:
				goto tr303
			case 316:
				goto tr303
			case 317:
				goto tr303
			case 318:
				goto tr303
			case 319:
				goto tr303
			case 320:
				goto tr303
			case 321:
				goto tr303
			case 322:
				goto tr303
			case 323:
				goto tr303
			case 324:
				goto tr303
			case 325:
				goto tr303
			case 326:
				goto tr303
			case 327:
				goto tr303
			case 328:
				goto tr303
			case 329:
				goto tr303
			case 330:
				goto tr303
			case 331:
				goto tr303
			case 332:
				goto tr303
			case 333:
				goto tr303
			case 334:
				goto tr303
			case 335:
				goto tr303
			case 336:
				goto tr303
			case 337:
				goto tr303
			case 338:
				goto tr303
			case 339:
				goto tr303
			case 340:
				goto tr303
			case 341:
				goto tr303
			case 342:
				goto tr303
			case 343:
				goto tr303
			case 344:
				goto tr303
			case 345:
				goto tr303
			case 346:
				goto tr303
			case 347:
				goto tr303
			case 348:
				goto tr303
			case 349:
				goto tr303
			case 350:
				goto tr303
			case 351:
				goto tr303
			case 352:
				goto tr303
			case 353:
				goto tr303
			case 354:
				goto tr303
			case 355:
				goto tr303
			case 356:
				goto tr536
			case 357:
				goto tr303
			case 358:
				goto tr303
			case 359:
				goto tr303
			case 360:
				goto tr303
			case 361:
				goto tr303
			case 362:
				goto tr303
			case 363:
				goto tr303
			case 364:
				goto tr303
			case 365:
				goto tr303
			case 366:
				goto tr303
			case 367:
				goto tr303
			case 368:
				goto tr303
			case 369:
				goto tr303
			case 370:
				goto tr303
			case 371:
				goto tr303
			case 372:
				goto tr303
			case 373:
				goto tr303
			case 374:
				goto tr303
			case 375:
				goto tr303
			case 376:
				goto tr303
			case 377:
				goto tr303
			case 378:
				goto tr303
			case 379:
				goto tr303
			case 380:
				goto tr303
			case 381:
				goto tr303
			case 382:
				goto tr303
			case 383:
				goto tr303
			case 384:
				goto tr303
			case 385:
				goto tr303
			case 386:
				goto tr303
			case 387:
				goto tr303
			case 388:
				goto tr303
			case 389:
				goto tr303
			case 390:
				goto tr303
			case 391:
				goto tr303
			case 392:
				goto tr303
			case 393:
				goto tr303
			case 394:
				goto tr303
			case 395:
				goto tr303
			case 396:
				goto tr303
			case 397:
				goto tr582
			case 93:
				goto tr129
			case 94:
				goto tr129
			case 95:
				goto tr129
			case 96:
				goto tr129
			case 97:
				goto tr129
			case 98:
				goto tr129
			case 398:
				goto tr303
			case 399:
				goto tr303
			case 400:
				goto tr303
			case 401:
				goto tr239
			case 402:
				goto tr303
			case 403:
				goto tr303
			case 404:
				goto tr303
			case 405:
				goto tr303
			case 406:
				goto tr303
			case 407:
				goto tr303
			case 408:
				goto tr303
			case 409:
				goto tr303
			case 410:
				goto tr303
			case 411:
				goto tr303
			case 412:
				goto tr303
			case 413:
				goto tr303
			case 414:
				goto tr303
			case 415:
				goto tr303
			case 416:
				goto tr303
			case 417:
				goto tr303
			case 418:
				goto tr303
			case 419:
				goto tr303
			case 420:
				goto tr303
			case 421:
				goto tr303
			case 422:
				goto tr303
			case 423:
				goto tr303
			case 424:
				goto tr303
			case 425:
				goto tr303
			case 426:
				goto tr303
			case 427:
				goto tr303
			case 428:
				goto tr303
			case 429:
				goto tr303
			case 430:
				goto tr303
			case 431:
				goto tr303
			case 432:
				goto tr303
			case 433:
				goto tr303
			case 434:
				goto tr303
			case 435:
				goto tr303
			case 436:
				goto tr303
			case 437:
				goto tr303
			case 438:
				goto tr303
			case 439:
				goto tr303
			case 440:
				goto tr303
			case 441:
				goto tr303
			case 442:
				goto tr303
			case 443:
				goto tr303
			case 444:
				goto tr303
			case 445:
				goto tr303
			case 446:
				goto tr303
			case 447:
				goto tr303
			case 448:
				goto tr303
			case 449:
				goto tr303
			case 450:
				goto tr303
			case 451:
				goto tr303
			case 452:
				goto tr303
			case 453:
				goto tr303
			case 454:
				goto tr303
			case 455:
				goto tr303
			case 456:
				goto tr303
			case 457:
				goto tr303
			case 458:
				goto tr303
			case 459:
				goto tr303
			case 460:
				goto tr303
			case 461:
				goto tr303
			case 462:
				goto tr303
			case 463:
				goto tr303
			case 464:
				goto tr303
			case 465:
				goto tr303
			case 466:
				goto tr303
			case 467:
				goto tr239
			case 469:
				goto tr668
			case 470:
				goto tr670
			case 99:
				goto tr141
			case 471:
				goto tr674
			case 472:
				goto tr674
			case 473:
				goto tr676
			case 475:
				goto tr679
			case 476:
				goto tr680
			case 478:
				goto tr689
			case 479:
				goto tr691
			case 480:
				goto tr692
			case 481:
				goto tr689
			case 482:
				goto tr696
			case 484:
				goto tr706
			case 485:
				goto tr708
			case 486:
				goto tr709
			case 487:
				goto tr706
			case 488:
				goto tr713
			case 490:
				goto tr723
			case 491:
				goto tr725
			case 492:
				goto tr726
			case 493:
				goto tr723
			case 494:
				goto tr730
			case 496:
				goto tr733
			case 498:
				goto tr739
			case 499:
				goto tr741
			case 500:
				goto tr739
			case 103:
				goto tr147
			case 501:
				goto tr743
			case 503:
				goto tr753
			case 504:
				goto tr754
			case 505:
				goto tr755
			case 506:
				goto tr757
			case 507:
				goto tr758
			case 508:
				goto tr758
			case 104:
				goto tr149
			case 105:
				goto tr149
			case 509:
				goto tr758
			case 106:
				goto tr149
			case 510:
				goto tr758
			case 511:
				goto tr762
			case 513:
				goto tr765
			case 107:
				goto tr153
			case 515:
				goto tr770
			case 516:
				goto tr772
			case 108:
				goto tr156
			case 517:
				goto tr776
			case 519:
				goto tr781
			case 520:
				goto tr783
			case 109:
				goto tr158
			case 521:
				goto tr787
			case 523:
				goto tr792
			case 524:
				goto tr794
			case 110:
				goto tr160
			case 525:
				goto tr798
			case 527:
				goto tr801
			case 528:
				goto tr802
			}
		}

	_out:
		{
		}
	}

// line scanner/scanner.rl:509

	token.FreeFloating = lex.FreeFloating
	token.Value = string(lex.data[lex.ts:lex.te])

	lval.Token(token)

	return int(tok)
}
