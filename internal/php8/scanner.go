
// line internal/php8/scanner.rl:1
package php8

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/z7zmey/php-parser/pkg/token"
)


// line internal/php8/scanner.go:15
const lexer_start int = 117
const lexer_first_final int = 117
const lexer_error int = 0

const lexer_en_main int = 117
const lexer_en_html int = 120
const lexer_en_php int = 127
const lexer_en_property int = 482
const lexer_en_nowdoc int = 489
const lexer_en_heredoc int = 493
const lexer_en_backqote int = 500
const lexer_en_template_string int = 506
const lexer_en_heredoc_end int = 512
const lexer_en_string_var int = 514
const lexer_en_string_var_index int = 520
const lexer_en_string_var_name int = 530
const lexer_en_halt_compiller_open_parenthesis int = 532
const lexer_en_halt_compiller_close_parenthesis int = 536
const lexer_en_halt_compiller_close_semicolon int = 540
const lexer_en_halt_compiller_end int = 544


// line internal/php8/scanner.rl:17


func initLexer(lex *Lexer)  {
    
// line internal/php8/scanner.go:43
	{
	 lex.cs = lexer_start
	 lex.top = 0
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

// line internal/php8/scanner.rl:21
}

func (lex *Lexer) Lex() *token.Token {
    eof := lex.pe
    var tok token.ID

    tkn := lex.tokenPool.Get()

    lblStart := 0
    lblEnd   := 0

    _, _ = lblStart, lblEnd

    
// line internal/php8/scanner.go:67
	{
	var _widec int16
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	goto _resume

_again:
	switch  lex.cs {
	case 117:
		goto st117
	case 118:
		goto st118
	case 1:
		goto st1
	case 119:
		goto st119
	case 120:
		goto st120
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
	case 2:
		goto st2
	case 3:
		goto st3
	case 4:
		goto st4
	case 126:
		goto st126
	case 5:
		goto st5
	case 127:
		goto st127
	case 128:
		goto st128
	case 129:
		goto st129
	case 6:
		goto st6
	case 130:
		goto st130
	case 131:
		goto st131
	case 132:
		goto st132
	case 133:
		goto st133
	case 7:
		goto st7
	case 8:
		goto st8
	case 9:
		goto st9
	case 10:
		goto st10
	case 134:
		goto st134
	case 135:
		goto st135
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
	case 141:
		goto st141
	case 142:
		goto st142
	case 11:
		goto st11
	case 12:
		goto st12
	case 143:
		goto st143
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
	case 64:
		goto st64
	case 149:
		goto st149
	case 65:
		goto st65
	case 66:
		goto st66
	case 150:
		goto st150
	case 67:
		goto st67
	case 151:
		goto st151
	case 68:
		goto st68
	case 69:
		goto st69
	case 70:
		goto st70
	case 152:
		goto st152
	case 153:
		goto st153
	case 154:
		goto st154
	case 71:
		goto st71
	case 72:
		goto st72
	case 155:
		goto st155
	case 73:
		goto st73
	case 156:
		goto st156
	case 157:
		goto st157
	case 158:
		goto st158
	case 74:
		goto st74
	case 75:
		goto st75
	case 76:
		goto st76
	case 77:
		goto st77
	case 159:
		goto st159
	case 160:
		goto st160
	case 78:
		goto st78
	case 161:
		goto st161
	case 162:
		goto st162
	case 79:
		goto st79
	case 80:
		goto st80
	case 81:
		goto st81
	case 82:
		goto st82
	case 163:
		goto st163
	case 83:
		goto st83
	case 84:
		goto st84
	case 85:
		goto st85
	case 86:
		goto st86
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
	case 87:
		goto st87
	case 170:
		goto st170
	case 171:
		goto st171
	case 88:
		goto st88
	case 172:
		goto st172
	case 173:
		goto st173
	case 174:
		goto st174
	case 89:
		goto st89
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
	case 90:
		goto st90
	case 91:
		goto st91
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
	case 92:
		goto st92
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
	case 413:
		goto st413
	case 414:
		goto st414
	case 99:
		goto st99
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
	case 482:
		goto st482
	case 483:
		goto st483
	case 484:
		goto st484
	case 100:
		goto st100
	case 485:
		goto st485
	case 486:
		goto st486
	case 487:
		goto st487
	case 101:
		goto st101
	case 488:
		goto st488
	case 489:
		goto st489
	case 0:
		goto st0
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
	case 102:
		goto st102
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
	case 104:
		goto st104
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
	case 106:
		goto st106
	case 507:
		goto st507
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
	case 513:
		goto st513
	case 514:
		goto st514
	case 515:
		goto st515
	case 516:
		goto st516
	case 517:
		goto st517
	case 107:
		goto st107
	case 518:
		goto st518
	case 108:
		goto st108
	case 109:
		goto st109
	case 519:
		goto st519
	case 520:
		goto st520
	case 521:
		goto st521
	case 522:
		goto st522
	case 523:
		goto st523
	case 524:
		goto st524
	case 525:
		goto st525
	case 526:
		goto st526
	case 110:
		goto st110
	case 111:
		goto st111
	case 527:
		goto st527
	case 112:
		goto st112
	case 528:
		goto st528
	case 529:
		goto st529
	case 530:
		goto st530
	case 531:
		goto st531
	case 113:
		goto st113
	case 532:
		goto st532
	case 533:
		goto st533
	case 534:
		goto st534
	case 114:
		goto st114
	case 535:
		goto st535
	case 536:
		goto st536
	case 537:
		goto st537
	case 538:
		goto st538
	case 115:
		goto st115
	case 539:
		goto st539
	case 540:
		goto st540
	case 541:
		goto st541
	case 542:
		goto st542
	case 116:
		goto st116
	case 543:
		goto st543
	case 544:
		goto st544
	case 545:
		goto st545
	case 546:
		goto st546
	case 547:
		goto st547
	}

	if ( lex.p)++; ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
_resume:
	switch  lex.cs {
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 1:
		goto st_case_1
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
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
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 126:
		goto st_case_126
	case 5:
		goto st_case_5
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 6:
		goto st_case_6
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
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
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 143:
		goto st_case_143
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
	case 64:
		goto st_case_64
	case 149:
		goto st_case_149
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 150:
		goto st_case_150
	case 67:
		goto st_case_67
	case 151:
		goto st_case_151
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 155:
		goto st_case_155
	case 73:
		goto st_case_73
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 78:
		goto st_case_78
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 163:
		goto st_case_163
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
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
	case 87:
		goto st_case_87
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 88:
		goto st_case_88
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 89:
		goto st_case_89
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
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
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
	case 92:
		goto st_case_92
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
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 99:
		goto st_case_99
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
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 484:
		goto st_case_484
	case 100:
		goto st_case_100
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 101:
		goto st_case_101
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 0:
		goto st_case_0
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
	case 102:
		goto st_case_102
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
	case 104:
		goto st_case_104
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
	case 106:
		goto st_case_106
	case 507:
		goto st_case_507
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
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 517:
		goto st_case_517
	case 107:
		goto st_case_107
	case 518:
		goto st_case_518
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 519:
		goto st_case_519
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 527:
		goto st_case_527
	case 112:
		goto st_case_112
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 113:
		goto st_case_113
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 114:
		goto st_case_114
	case 535:
		goto st_case_535
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
	case 538:
		goto st_case_538
	case 115:
		goto st_case_115
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 116:
		goto st_case_116
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 545:
		goto st_case_545
	case 546:
		goto st_case_546
	case 547:
		goto st_case_547
	}
	goto st_out
tr0:
	 lex.cs = 117
// line internal/php8/scanner.rl:130
( lex.p) = ( lex.te) - 1
{
                 lex.cs = 120;
                lex.ungetCnt(1)
            }
	goto _again
tr175:
	 lex.cs = 117
// line internal/php8/scanner.rl:130
 lex.te = ( lex.p)+1
{
                 lex.cs = 120;
                lex.ungetCnt(1)
            }
	goto _again
tr177:
	 lex.cs = 117
// line internal/php8/scanner.rl:130
 lex.te = ( lex.p)
( lex.p)--
{
                 lex.cs = 120;
                lex.ungetCnt(1)
            }
	goto _again
tr178:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:127
 lex.te = ( lex.p)
( lex.p)--
{
                lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
            }
	goto st117
	st117:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof117
		}
	st_case_117:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:2327
		if  lex.data[( lex.p)] == 35 {
			goto tr176
		}
		goto tr175
tr176:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st118
	st118:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof118
		}
	st_case_118:
// line internal/php8/scanner.go:2342
		if  lex.data[( lex.p)] == 33 {
			goto st1
		}
		goto tr177
tr3:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st1
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
// line internal/php8/scanner.go:2364
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr2
		case 13:
			goto tr3
		}
		goto st1
tr2:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st119
	st119:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof119
		}
	st_case_119:
// line internal/php8/scanner.go:2389
		goto tr178
tr4:
	 lex.cs = 120
// line internal/php8/scanner.rl:143
( lex.p) = ( lex.te) - 1
{
                lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.te)
                 lex.cs = 127;
            }
	goto _again
tr7:
	 lex.cs = 120
// line internal/php8/scanner.rl:147
 lex.te = ( lex.p)+1
{
                lex.ungetCnt(lex.te - lex.ts - 5)
                lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
                 lex.cs = 127;
            }
	goto _again
tr183:
// line internal/php8/scanner.rl:137
 lex.te = ( lex.p)
( lex.p)--
{
                lex.ungetStr("<")
                lex.setTokenPosition(tkn)
                tok = token.T_INLINE_HTML;
                {( lex.p)++;  lex.cs = 120; goto _out }
            }
	goto st120
tr185:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:137
 lex.te = ( lex.p)
( lex.p)--
{
                lex.ungetStr("<")
                lex.setTokenPosition(tkn)
                tok = token.T_INLINE_HTML;
                {( lex.p)++;  lex.cs = 120; goto _out }
            }
	goto st120
tr191:
	 lex.cs = 120
// line internal/php8/scanner.rl:143
 lex.te = ( lex.p)
( lex.p)--
{
                lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.te)
                 lex.cs = 127;
            }
	goto _again
tr192:
	 lex.cs = 120
// line internal/php8/scanner.rl:152
 lex.te = ( lex.p)+1
{
                lex.setTokenPosition(tkn);
                tok = token.T_ECHO;
                 lex.cs = 127;
                {( lex.p)++; goto _out }
            }
	goto _again
tr194:
	 lex.cs = 120
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:147
 lex.te = ( lex.p)
( lex.p)--
{
                lex.ungetCnt(lex.te - lex.ts - 5)
                lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
                 lex.cs = 127;
            }
	goto _again
	st120:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof120
		}
	st_case_120:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:2479
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr180
		case 13:
			goto tr181
		case 60:
			goto st124
		}
		goto st121
tr181:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st121
tr186:
// line internal/php8/scanner.rl:54

	goto st121
tr188:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st121
	st121:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof121
		}
	st_case_121:
// line internal/php8/scanner.go:2524
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr180
		case 13:
			goto tr181
		case 60:
			goto st123
		}
		goto st121
tr180:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st122
tr187:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st122
	st122:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof122
		}
	st_case_122:
// line internal/php8/scanner.go:2565
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr187
		case 13:
			goto tr188
		case 60:
			goto tr189
		}
		goto tr186
tr189:
// line internal/php8/scanner.rl:54

	goto st123
	st123:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof123
		}
	st_case_123:
// line internal/php8/scanner.go:2584
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr180
		case 13:
			goto tr181
		case 60:
			goto st123
		case 63:
			goto tr183
		}
		goto st121
	st124:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr180
		case 13:
			goto tr181
		case 60:
			goto st123
		case 63:
			goto tr190
		}
		goto st121
tr190:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st125
	st125:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof125
		}
	st_case_125:
// line internal/php8/scanner.go:2622
		switch  lex.data[( lex.p)] {
		case 61:
			goto tr192
		case 80:
			goto st2
		case 112:
			goto st2
		}
		goto tr191
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch  lex.data[( lex.p)] {
		case 72:
			goto st3
		case 104:
			goto st3
		}
		goto tr4
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch  lex.data[( lex.p)] {
		case 80:
			goto st4
		case 112:
			goto st4
		}
		goto tr4
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 9:
			goto tr7
		case 10:
			goto tr8
		case 13:
			goto tr9
		case 32:
			goto tr7
		}
		goto tr4
tr8:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st126
	st126:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof126
		}
	st_case_126:
// line internal/php8/scanner.go:2689
		goto tr194
tr9:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st5
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
// line internal/php8/scanner.go:2708
		if  lex.data[( lex.p)] == 10 {
			goto tr8
		}
		goto tr4
tr10:
// line internal/php8/scanner.rl:161
( lex.p) = ( lex.te) - 1
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st127
tr12:
	 lex.cs = 127
// line NONE:1
	switch  lex.act {
	case 10:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++; goto _out }}
	case 11:
	{( lex.p) = ( lex.te) - 1

                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 2, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; {( lex.p)++; goto _out }
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++; goto _out }
            }
	case 12:
	{( lex.p) = ( lex.te) - 1

                base := 10
                if lex.data[lex.ts] == '0' {
                    base = 8
                }

                s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, base, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; {( lex.p)++; goto _out }
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++; goto _out }
            }
	case 13:
	{( lex.p) = ( lex.te) - 1

                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 16, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; {( lex.p)++; goto _out }
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++; goto _out }
            }
	case 14:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_NAME_RELATIVE; {( lex.p)++; goto _out }}
	case 15:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_NAME_QUALIFIED; {( lex.p)++; goto _out }}
	case 18:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ABSTRACT; {( lex.p)++; goto _out }}
	case 19:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ARRAY; {( lex.p)++; goto _out }}
	case 20:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_AS; {( lex.p)++; goto _out }}
	case 21:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_BREAK; {( lex.p)++; goto _out }}
	case 22:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CALLABLE; {( lex.p)++; goto _out }}
	case 23:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CASE; {( lex.p)++; goto _out }}
	case 24:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CATCH; {( lex.p)++; goto _out }}
	case 25:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CLASS; {( lex.p)++; goto _out }}
	case 26:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CLONE; {( lex.p)++; goto _out }}
	case 27:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CONST; {( lex.p)++; goto _out }}
	case 28:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CONTINUE; {( lex.p)++; goto _out }}
	case 29:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DECLARE; {( lex.p)++; goto _out }}
	case 30:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DEFAULT; {( lex.p)++; goto _out }}
	case 31:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DO; {( lex.p)++; goto _out }}
	case 32:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ECHO; {( lex.p)++; goto _out }}
	case 33:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ELSE; {( lex.p)++; goto _out }}
	case 34:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ELSEIF; {( lex.p)++; goto _out }}
	case 35:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_EMPTY; {( lex.p)++; goto _out }}
	case 36:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ENDDECLARE; {( lex.p)++; goto _out }}
	case 37:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ENDFOR; {( lex.p)++; goto _out }}
	case 38:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ENDFOREACH; {( lex.p)++; goto _out }}
	case 39:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ENDIF; {( lex.p)++; goto _out }}
	case 40:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ENDSWITCH; {( lex.p)++; goto _out }}
	case 41:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ENDWHILE; {( lex.p)++; goto _out }}
	case 42:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_EVAL; {( lex.p)++; goto _out }}
	case 43:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_EXIT; {( lex.p)++; goto _out }}
	case 44:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_EXTENDS; {( lex.p)++; goto _out }}
	case 45:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FINAL; {( lex.p)++; goto _out }}
	case 46:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FINALLY; {( lex.p)++; goto _out }}
	case 47:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FOR; {( lex.p)++; goto _out }}
	case 48:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FOREACH; {( lex.p)++; goto _out }}
	case 49:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FUNCTION; {( lex.p)++; goto _out }}
	case 50:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FN; {( lex.p)++; goto _out }}
	case 51:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_GLOBAL; {( lex.p)++; goto _out }}
	case 52:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_GOTO; {( lex.p)++; goto _out }}
	case 53:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_IF; {( lex.p)++; goto _out }}
	case 54:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ISSET; {( lex.p)++; goto _out }}
	case 55:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_IMPLEMENTS; {( lex.p)++; goto _out }}
	case 56:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_INSTANCEOF; {( lex.p)++; goto _out }}
	case 57:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_INSTEADOF; {( lex.p)++; goto _out }}
	case 58:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_INTERFACE; {( lex.p)++; goto _out }}
	case 59:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_LIST; {( lex.p)++; goto _out }}
	case 60:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_NAMESPACE; {( lex.p)++; goto _out }}
	case 61:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_PRIVATE; {( lex.p)++; goto _out }}
	case 62:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_PUBLIC; {( lex.p)++; goto _out }}
	case 63:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_PRINT; {( lex.p)++; goto _out }}
	case 64:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_PROTECTED; {( lex.p)++; goto _out }}
	case 65:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_RETURN; {( lex.p)++; goto _out }}
	case 66:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_STATIC; {( lex.p)++; goto _out }}
	case 67:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_SWITCH; {( lex.p)++; goto _out }}
	case 68:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_MATCH; {( lex.p)++; goto _out }}
	case 69:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_THROW; {( lex.p)++; goto _out }}
	case 70:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_TRAIT; {( lex.p)++; goto _out }}
	case 71:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_TRY; {( lex.p)++; goto _out }}
	case 72:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_UNSET; {( lex.p)++; goto _out }}
	case 73:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_USE; {( lex.p)++; goto _out }}
	case 74:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_VAR; {( lex.p)++; goto _out }}
	case 75:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_WHILE; {( lex.p)++; goto _out }}
	case 77:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_YIELD; {( lex.p)++; goto _out }}
	case 78:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_INCLUDE; {( lex.p)++; goto _out }}
	case 79:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_INCLUDE_ONCE; {( lex.p)++; goto _out }}
	case 80:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_REQUIRE; {( lex.p)++; goto _out }}
	case 81:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_REQUIRE_ONCE; {( lex.p)++; goto _out }}
	case 82:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_CLASS_C; {( lex.p)++; goto _out }}
	case 83:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DIR; {( lex.p)++; goto _out }}
	case 84:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FILE; {( lex.p)++; goto _out }}
	case 85:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_FUNC_C; {( lex.p)++; goto _out }}
	case 86:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_LINE; {( lex.p)++; goto _out }}
	case 87:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_NS_C; {( lex.p)++; goto _out }}
	case 88:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_METHOD_C; {( lex.p)++; goto _out }}
	case 89:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_TRAIT_C; {( lex.p)++; goto _out }}
	case 90:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_HALT_COMPILER;  lex.cs = 532; {( lex.p)++; goto _out }}
	case 91:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_NEW; {( lex.p)++; goto _out }}
	case 92:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_LOGICAL_AND; {( lex.p)++; goto _out }}
	case 93:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_LOGICAL_OR; {( lex.p)++; goto _out }}
	case 94:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_LOGICAL_XOR; {( lex.p)++; goto _out }}
	case 95:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_ATTRIBUTE; {( lex.p)++; goto _out }}
	case 123:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_SL; {( lex.p)++; goto _out }}
	case 134:
	{( lex.p) = ( lex.te) - 1

                lex.ungetStr("?>")
                lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
            }
	case 140:
	{( lex.p) = ( lex.te) - 1
 lex.setTokenPosition(tkn); tok = token.T_STRING;   {( lex.p)++; goto _out } }
	case 146:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.ID(int('"'));  lex.cs = 506; {( lex.p)++; goto _out }}
	}
	
	goto _again
tr15:
// line internal/php8/scanner.rl:356
 lex.te = ( lex.p)+1
{
                lex.setTokenPosition(tkn);
                tok = token.T_CONSTANT_ENCAPSED_STRING;
                {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr19:
// line internal/php8/scanner.rl:379
( lex.p) = ( lex.te) - 1
{
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            }
	goto st127
tr23:
// line internal/php8/scanner.rl:342
( lex.p) = ( lex.te) - 1
{
                lex.setTokenPosition(tkn);
                tok = token.ID(int(lex.data[lex.ts]));
                {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr37:
// line internal/php8/scanner.rl:317
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_ARRAY_CAST; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr44:
// line internal/php8/scanner.rl:322
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_STRING_CAST; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr48:
// line internal/php8/scanner.rl:318
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_BOOL_CAST; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr56:
// line internal/php8/scanner.rl:319
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_DOUBLE_CAST; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr63:
// line internal/php8/scanner.rl:320
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_INT_CAST; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr72:
// line internal/php8/scanner.rl:321
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_OBJECT_CAST; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr81:
// line internal/php8/scanner.rl:323
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_UNSET_CAST; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr82:
// line internal/php8/scanner.rl:285
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_ELLIPSIS; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr85:
// line internal/php8/scanner.rl:165
( lex.p) = ( lex.te) - 1
{lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr95:
// line internal/php8/scanner.rl:329
 lex.te = ( lex.p)+1
{
                isDocComment := false;
                if lex.te - lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
                    isDocComment = true;
                }

                if isDocComment {
                    lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
                } else {
                    lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
                }
            }
	goto st127
tr96:
// line internal/php8/scanner.rl:176
( lex.p) = ( lex.te) - 1
{
                base := 10
                if lex.data[lex.ts] == '0' {
                    base = 8
                }

                s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, base, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr109:
	 lex.cs = 127
// line internal/php8/scanner.rl:163
( lex.p) = ( lex.te) - 1
{lex.setTokenPosition(tkn); tok = token.ID(int(';'));  lex.cs = 120; {( lex.p)++; goto _out }}
	goto _again
tr125:
	 lex.cs = 127
// line internal/php8/scanner.rl:354
 lex.te = ( lex.p)+1
{ lex.setTokenPosition(tkn); tok = token.T_NULLSAFE_OBJECT_OPERATOR;  lex.cs = 482; {( lex.p)++; goto _out } }
	goto _again
tr126:
	 lex.cs = 127
// line internal/php8/scanner.rl:162
( lex.p) = ( lex.te) - 1
{lex.setTokenPosition(tkn); tok = token.ID(int(';'));  lex.cs = 120; {( lex.p)++; goto _out }}
	goto _again
tr129:
// line internal/php8/scanner.rl:351
( lex.p) = ( lex.te) - 1
{ lex.setTokenPosition(tkn); tok = token.T_STRING;   {( lex.p)++;  lex.cs = 127; goto _out } }
	goto st127
tr132:
// line internal/php8/scanner.rl:266
( lex.p) = ( lex.te) - 1
{lex.setTokenPosition(tkn); tok = token.T_YIELD; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr143:
// line internal/php8/scanner.rl:265
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_YIELD_FROM; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr144:
// line internal/php8/scanner.rl:204
( lex.p) = ( lex.te) - 1
{lex.setTokenPosition(tkn); tok = token.T_NAME_FULLY_QUALIFIED; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr195:
// line internal/php8/scanner.rl:379
 lex.te = ( lex.p)+1
{
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            }
	goto st127
tr206:
// line internal/php8/scanner.rl:342
 lex.te = ( lex.p)+1
{
                lex.setTokenPosition(tkn);
                tok = token.ID(int(lex.data[lex.ts]));
                {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr244:
	 lex.cs = 127
// line internal/php8/scanner.rl:376
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int('`'));  lex.cs = 500; {( lex.p)++; goto _out }}
	goto _again
tr245:
// line internal/php8/scanner.rl:348
 lex.te = ( lex.p)+1
{ lex.setTokenPosition(tkn); tok = token.ID(int('{')); lex.call(127, 127); goto _out; }
	goto st127
tr247:
// line internal/php8/scanner.rl:349
 lex.te = ( lex.p)+1
{ lex.setTokenPosition(tkn); tok = token.ID(int('}')); lex.ret(1); goto _out;}
	goto st127
tr248:
// line internal/php8/scanner.rl:161
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st127
tr250:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:161
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st127
tr254:
// line internal/php8/scanner.rl:379
 lex.te = ( lex.p)
( lex.p)--
{
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            }
	goto st127
tr255:
// line internal/php8/scanner.rl:342
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.ID(int(lex.data[lex.ts]));
                {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr257:
// line internal/php8/scanner.rl:303
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_IS_NOT_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr258:
// line internal/php8/scanner.rl:304
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_IS_NOT_IDENTICAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr259:
	 lex.cs = 127
// line internal/php8/scanner.rl:377
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.ID(int('"'));  lex.cs = 506; {( lex.p)++; goto _out }}
	goto _again
tr260:
// line internal/php8/scanner.rl:325
 lex.te = ( lex.p)
( lex.p)--
{
                lex.ungetStr("?>")
                lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
            }
	goto st127
tr261:
// line internal/php8/scanner.rl:284
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_ATTRIBUTE; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr266:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:325
 lex.te = ( lex.p)
( lex.p)--
{
                lex.ungetStr("?>")
                lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
            }
	goto st127
tr271:
// line internal/php8/scanner.rl:350
 lex.te = ( lex.p)
( lex.p)--
{ lex.setTokenPosition(tkn); tok = token.T_VARIABLE; {( lex.p)++;  lex.cs = 127; goto _out } }
	goto st127
tr272:
// line internal/php8/scanner.rl:298
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_MOD_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr273:
// line internal/php8/scanner.rl:287
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_BOOLEAN_AND; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr274:
// line internal/php8/scanner.rl:289
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_AND_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr276:
// line internal/php8/scanner.rl:292
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_MUL_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr277:
// line internal/php8/scanner.rl:311
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_POW; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr278:
// line internal/php8/scanner.rl:293
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_POW_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr279:
// line internal/php8/scanner.rl:300
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_INC; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr280:
// line internal/php8/scanner.rl:295
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_PLUS_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr281:
// line internal/php8/scanner.rl:299
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_DEC; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr282:
// line internal/php8/scanner.rl:296
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_MINUS_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr283:
	 lex.cs = 127
// line internal/php8/scanner.rl:353
 lex.te = ( lex.p)+1
{ lex.setTokenPosition(tkn); tok = token.T_OBJECT_OPERATOR;  lex.cs = 482; {( lex.p)++; goto _out } }
	goto _again
tr285:
// line internal/php8/scanner.rl:291
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_CONCAT_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr286:
// line internal/php8/scanner.rl:165
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr289:
// line internal/php8/scanner.rl:294
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_DIV_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr290:
// line internal/php8/scanner.rl:176
 lex.te = ( lex.p)
( lex.p)--
{
                base := 10
                if lex.data[lex.ts] == '0' {
                    base = 8
                }

                s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, base, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr295:
// line internal/php8/scanner.rl:166
 lex.te = ( lex.p)
( lex.p)--
{
                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 2, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr296:
// line internal/php8/scanner.rl:191
 lex.te = ( lex.p)
( lex.p)--
{
                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 16, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; {( lex.p)++;  lex.cs = 127; goto _out }
            }
	goto st127
tr297:
// line internal/php8/scanner.rl:286
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_PAAMAYIM_NEKUDOTAYIM; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr298:
	 lex.cs = 127
// line internal/php8/scanner.rl:163
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.ID(int(';'));  lex.cs = 120; {( lex.p)++; goto _out }}
	goto _again
tr300:
	 lex.cs = 127
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:163
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.ID(int(';'));  lex.cs = 120; {( lex.p)++; goto _out }}
	goto _again
tr303:
// line internal/php8/scanner.rl:303
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_IS_NOT_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr304:
// line internal/php8/scanner.rl:312
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_SL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr305:
// line internal/php8/scanner.rl:307
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_SL_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr306:
	 lex.cs = 127
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:362
 lex.te = ( lex.p)
( lex.p)--
{
                lex.heredocLabel = lex.data[lblStart:lblEnd]
                lex.setTokenPosition(tkn);
                tok = token.T_START_HEREDOC;

                if lex.isHeredocEnd(lex.p+1) {
                     lex.cs = 512;
                } else if lex.data[lblStart-1] == '\'' {
                     lex.cs = 489;
                } else {
                     lex.cs = 493;
                }
                {( lex.p)++; goto _out }
            }
	goto _again
tr307:
// line internal/php8/scanner.rl:310
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_IS_SMALLER_OR_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr308:
// line internal/php8/scanner.rl:302
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_SPACESHIP; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr310:
// line internal/php8/scanner.rl:301
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_DOUBLE_ARROW; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr311:
// line internal/php8/scanner.rl:305
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_IS_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr312:
// line internal/php8/scanner.rl:306
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_IS_IDENTICAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr313:
// line internal/php8/scanner.rl:309
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_IS_GREATER_OR_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr315:
// line internal/php8/scanner.rl:313
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_SR; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr316:
// line internal/php8/scanner.rl:308
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_SR_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr320:
	 lex.cs = 127
// line internal/php8/scanner.rl:162
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.ID(int(';'));  lex.cs = 120; {( lex.p)++; goto _out }}
	goto _again
tr322:
	 lex.cs = 127
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:162
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.ID(int(';'));  lex.cs = 120; {( lex.p)++; goto _out }}
	goto _again
tr323:
// line internal/php8/scanner.rl:314
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_COALESCE; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr324:
// line internal/php8/scanner.rl:315
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_COALESCE_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr325:
// line internal/php8/scanner.rl:351
 lex.te = ( lex.p)
( lex.p)--
{ lex.setTokenPosition(tkn); tok = token.T_STRING;   {( lex.p)++;  lex.cs = 127; goto _out } }
	goto st127
tr331:
// line internal/php8/scanner.rl:203
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_NAME_QUALIFIED; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr407:
// line internal/php8/scanner.rl:222
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_ELSE; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr427:
// line internal/php8/scanner.rl:226
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_ENDFOR; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr456:
// line internal/php8/scanner.rl:234
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_FINAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr460:
// line internal/php8/scanner.rl:236
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_FOR; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr492:
// line internal/php8/scanner.rl:267
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_INCLUDE; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr535:
// line internal/php8/scanner.rl:249
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_NAMESPACE; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr537:
// line internal/php8/scanner.rl:202
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_NAME_RELATIVE; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr567:
// line internal/php8/scanner.rl:269
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_REQUIRE; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr613:
// line internal/php8/scanner.rl:266
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_YIELD; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr614:
// line internal/php8/scanner.rl:205
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_NS_SEPARATOR; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr615:
// line internal/php8/scanner.rl:204
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_NAME_FULLY_QUALIFIED; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr617:
// line internal/php8/scanner.rl:297
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_XOR_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr691:
// line internal/php8/scanner.rl:290
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_OR_EQUAL; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
tr692:
// line internal/php8/scanner.rl:288
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_BOOLEAN_OR; {( lex.p)++;  lex.cs = 127; goto _out }}
	goto st127
	st127:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof127
		}
	st_case_127:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:3636
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr197
		case 32:
			goto tr196
		case 33:
			goto st131
		case 34:
			goto tr199
		case 35:
			goto st134
		case 36:
			goto st138
		case 37:
			goto st140
		case 38:
			goto st141
		case 39:
			goto tr204
		case 40:
			goto tr205
		case 42:
			goto st144
		case 43:
			goto st146
		case 45:
			goto st147
		case 46:
			goto tr210
		case 47:
			goto tr211
		case 48:
			goto tr212
		case 58:
			goto st157
		case 59:
			goto tr214
		case 60:
			goto st161
		case 61:
			goto st165
		case 62:
			goto st167
		case 63:
			goto tr218
		case 64:
			goto tr206
		case 65:
			goto tr219
		case 66:
			goto tr220
		case 67:
			goto tr221
		case 68:
			goto tr222
		case 69:
			goto tr223
		case 70:
			goto tr224
		case 71:
			goto tr225
		case 73:
			goto tr227
		case 76:
			goto tr228
		case 77:
			goto tr229
		case 78:
			goto tr230
		case 79:
			goto tr231
		case 80:
			goto tr232
		case 82:
			goto tr233
		case 83:
			goto tr234
		case 84:
			goto tr235
		case 85:
			goto tr236
		case 86:
			goto tr237
		case 87:
			goto tr238
		case 88:
			goto tr239
		case 89:
			goto tr240
		case 92:
			goto st413
		case 94:
			goto st415
		case 95:
			goto tr243
		case 96:
			goto tr244
		case 97:
			goto tr219
		case 98:
			goto tr220
		case 99:
			goto tr221
		case 100:
			goto tr222
		case 101:
			goto tr223
		case 102:
			goto tr224
		case 103:
			goto tr225
		case 105:
			goto tr227
		case 108:
			goto tr228
		case 109:
			goto tr229
		case 110:
			goto tr230
		case 111:
			goto tr231
		case 112:
			goto tr232
		case 114:
			goto tr233
		case 115:
			goto tr234
		case 116:
			goto tr235
		case 117:
			goto tr236
		case 118:
			goto tr237
		case 119:
			goto tr238
		case 120:
			goto tr239
		case 121:
			goto tr240
		case 123:
			goto tr245
		case 124:
			goto st481
		case 125:
			goto tr247
		case 126:
			goto tr206
		case 127:
			goto tr195
		}
		switch {
		case  lex.data[( lex.p)] < 14:
			switch {
			case  lex.data[( lex.p)] > 8:
				if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
					goto tr196
				}
			default:
				goto tr195
			}
		case  lex.data[( lex.p)] > 31:
			switch {
			case  lex.data[( lex.p)] < 49:
				if 41 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 44 {
					goto tr206
				}
			case  lex.data[( lex.p)] > 57:
				if 91 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 93 {
					goto tr206
				}
			default:
				goto tr97
			}
		default:
			goto tr195
		}
		goto tr226
tr196:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st128
tr251:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

	goto st128
	st128:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof128
		}
	st_case_128:
// line internal/php8/scanner.go:3833
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr249
		case 32:
			goto tr196
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr196
		}
		goto tr248
tr11:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st129
tr252:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st129
	st129:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof129
		}
	st_case_129:
// line internal/php8/scanner.go:3883
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr252
		case 13:
			goto tr253
		case 32:
			goto tr251
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr251
		}
		goto tr250
tr249:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st6
tr253:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st6
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
// line internal/php8/scanner.go:3927
		if  lex.data[( lex.p)] == 10 {
			goto tr11
		}
		goto tr10
tr197:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st130
	st130:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof130
		}
	st_case_130:
// line internal/php8/scanner.go:3949
		if  lex.data[( lex.p)] == 10 {
			goto tr11
		}
		goto tr254
	st131:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof131
		}
	st_case_131:
		if  lex.data[( lex.p)] == 61 {
			goto st132
		}
		goto tr255
	st132:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		if  lex.data[( lex.p)] == 61 {
			goto tr258
		}
		goto tr257
tr199:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:377
 lex.act = 146;
	goto st133
	st133:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof133
		}
	st_case_133:
// line internal/php8/scanner.go:3984
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
tr14:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st7
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
// line internal/php8/scanner.go:4017
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 96:
			goto st7
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto st7
			}
		case  lex.data[( lex.p)] > 94:
			if 124 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto st7
			}
		default:
			goto st7
		}
		goto tr12
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		}
		goto st7
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto tr12
		case 92:
			goto st9
		}
		goto st7
	st134:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof134
		}
	st_case_134:
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
				_widec += 256
			}
		}
		switch _widec {
		case 347:
			goto tr261
		case 522:
			goto tr263
		case 525:
			goto tr264
		case 603:
			goto tr265
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr262
		}
		goto tr260
tr262:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:325
 lex.act = 134;
	goto st135
tr265:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:284
 lex.act = 95;
	goto st135
tr267:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:325
 lex.act = 134;
	goto st135
	st135:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof135
		}
	st_case_135:
// line internal/php8/scanner.go:4180
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr263
		case 525:
			goto tr264
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr262
		}
		goto tr12
tr263:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st136
tr268:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st136
	st136:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof136
		}
	st_case_136:
// line internal/php8/scanner.go:4260
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr268
		case 525:
			goto tr269
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr267
		}
		goto tr266
tr264:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st137
tr269:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st137
	st137:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof137
		}
	st_case_137:
// line internal/php8/scanner.go:4340
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotPhpCloseToken() && lex.isNotNewLine()   {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr263
		case 525:
			goto tr264
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr262
		}
		goto tr260
	st138:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof138
		}
	st_case_138:
		if  lex.data[( lex.p)] == 96 {
			goto tr255
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr255
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr255
			}
		default:
			goto tr255
		}
		goto st139
	st139:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof139
		}
	st_case_139:
		if  lex.data[( lex.p)] == 96 {
			goto tr271
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr271
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr271
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr271
			}
		default:
			goto tr271
		}
		goto st139
	st140:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof140
		}
	st_case_140:
		if  lex.data[( lex.p)] == 61 {
			goto tr272
		}
		goto tr255
	st141:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof141
		}
	st_case_141:
		switch  lex.data[( lex.p)] {
		case 38:
			goto tr273
		case 61:
			goto tr274
		}
		goto tr255
tr204:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st142
	st142:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof142
		}
	st_case_142:
// line internal/php8/scanner.go:4467
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr21
		case 13:
			goto tr21
		case 39:
			goto tr15
		case 92:
			goto st12
		}
		goto st11
tr21:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st11
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
// line internal/php8/scanner.go:4496
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr21
		case 13:
			goto tr21
		case 39:
			goto tr15
		case 92:
			goto st12
		}
		goto st11
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr21
		case 13:
			goto tr21
		}
		goto st11
tr205:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st143
	st143:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof143
		}
	st_case_143:
// line internal/php8/scanner.go:4530
		switch  lex.data[( lex.p)] {
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
		case 83:
			goto st54
		case 85:
			goto st59
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
		case 115:
			goto st54
		case 117:
			goto st59
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st13
		}
		goto tr255
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
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
		case 83:
			goto st54
		case 85:
			goto st59
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
		case 115:
			goto st54
		case 117:
			goto st59
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st13
		}
		goto tr23
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 82:
			goto st15
		case 114:
			goto st15
		}
		goto tr23
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 82:
			goto st16
		case 114:
			goto st16
		}
		goto tr23
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 65:
			goto st17
		case 97:
			goto st17
		}
		goto tr23
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 89:
			goto st18
		case 121:
			goto st18
		}
		goto tr23
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st18
		case 32:
			goto st18
		case 41:
			goto tr37
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st18
		}
		goto tr23
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 73:
			goto st20
		case 79:
			goto st25
		case 105:
			goto st20
		case 111:
			goto st25
		}
		goto tr23
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 78:
			goto st21
		case 110:
			goto st21
		}
		goto tr23
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 65:
			goto st22
		case 97:
			goto st22
		}
		goto tr23
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 82:
			goto st23
		case 114:
			goto st23
		}
		goto tr23
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 89:
			goto st24
		case 121:
			goto st24
		}
		goto tr23
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st24
		case 32:
			goto st24
		case 41:
			goto tr44
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st24
		}
		goto tr23
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 79:
			goto st26
		case 111:
			goto st26
		}
		goto tr23
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 76:
			goto st27
		case 108:
			goto st27
		}
		goto tr23
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
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
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st28
		}
		goto tr23
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 41:
			goto tr48
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st28
		}
		goto tr23
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  lex.data[( lex.p)] {
		case 65:
			goto st30
		case 97:
			goto st30
		}
		goto tr23
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  lex.data[( lex.p)] {
		case 78:
			goto st28
		case 110:
			goto st28
		}
		goto tr23
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch  lex.data[( lex.p)] {
		case 79:
			goto st32
		case 111:
			goto st32
		}
		goto tr23
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch  lex.data[( lex.p)] {
		case 85:
			goto st33
		case 117:
			goto st33
		}
		goto tr23
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch  lex.data[( lex.p)] {
		case 66:
			goto st34
		case 98:
			goto st34
		}
		goto tr23
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch  lex.data[( lex.p)] {
		case 76:
			goto st35
		case 108:
			goto st35
		}
		goto tr23
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch  lex.data[( lex.p)] {
		case 69:
			goto st36
		case 101:
			goto st36
		}
		goto tr23
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st36
		case 32:
			goto st36
		case 41:
			goto tr56
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st36
		}
		goto tr23
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 76:
			goto st38
		case 108:
			goto st38
		}
		goto tr23
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch  lex.data[( lex.p)] {
		case 79:
			goto st39
		case 111:
			goto st39
		}
		goto tr23
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  lex.data[( lex.p)] {
		case 65:
			goto st40
		case 97:
			goto st40
		}
		goto tr23
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch  lex.data[( lex.p)] {
		case 84:
			goto st36
		case 116:
			goto st36
		}
		goto tr23
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch  lex.data[( lex.p)] {
		case 78:
			goto st42
		case 110:
			goto st42
		}
		goto tr23
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch  lex.data[( lex.p)] {
		case 84:
			goto st43
		case 116:
			goto st43
		}
		goto tr23
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch  lex.data[( lex.p)] {
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
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st44
		}
		goto tr23
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st44
		case 32:
			goto st44
		case 41:
			goto tr63
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st44
		}
		goto tr23
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch  lex.data[( lex.p)] {
		case 71:
			goto st46
		case 103:
			goto st46
		}
		goto tr23
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch  lex.data[( lex.p)] {
		case 69:
			goto st47
		case 101:
			goto st47
		}
		goto tr23
	st47:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch  lex.data[( lex.p)] {
		case 82:
			goto st44
		case 114:
			goto st44
		}
		goto tr23
	st48:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch  lex.data[( lex.p)] {
		case 66:
			goto st49
		case 98:
			goto st49
		}
		goto tr23
	st49:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch  lex.data[( lex.p)] {
		case 74:
			goto st50
		case 106:
			goto st50
		}
		goto tr23
	st50:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch  lex.data[( lex.p)] {
		case 69:
			goto st51
		case 101:
			goto st51
		}
		goto tr23
	st51:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch  lex.data[( lex.p)] {
		case 67:
			goto st52
		case 99:
			goto st52
		}
		goto tr23
	st52:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch  lex.data[( lex.p)] {
		case 84:
			goto st53
		case 116:
			goto st53
		}
		goto tr23
	st53:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st53
		case 32:
			goto st53
		case 41:
			goto tr72
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st53
		}
		goto tr23
	st54:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch  lex.data[( lex.p)] {
		case 84:
			goto st55
		case 116:
			goto st55
		}
		goto tr23
	st55:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch  lex.data[( lex.p)] {
		case 82:
			goto st56
		case 114:
			goto st56
		}
		goto tr23
	st56:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch  lex.data[( lex.p)] {
		case 73:
			goto st57
		case 105:
			goto st57
		}
		goto tr23
	st57:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch  lex.data[( lex.p)] {
		case 78:
			goto st58
		case 110:
			goto st58
		}
		goto tr23
	st58:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch  lex.data[( lex.p)] {
		case 71:
			goto st24
		case 103:
			goto st24
		}
		goto tr23
	st59:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch  lex.data[( lex.p)] {
		case 78:
			goto st60
		case 110:
			goto st60
		}
		goto tr23
	st60:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch  lex.data[( lex.p)] {
		case 83:
			goto st61
		case 115:
			goto st61
		}
		goto tr23
	st61:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch  lex.data[( lex.p)] {
		case 69:
			goto st62
		case 101:
			goto st62
		}
		goto tr23
	st62:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch  lex.data[( lex.p)] {
		case 84:
			goto st63
		case 116:
			goto st63
		}
		goto tr23
	st63:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st63
		case 32:
			goto st63
		case 41:
			goto tr81
		}
		if 11 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st63
		}
		goto tr23
	st144:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof144
		}
	st_case_144:
		switch  lex.data[( lex.p)] {
		case 42:
			goto st145
		case 61:
			goto tr276
		}
		goto tr255
	st145:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof145
		}
	st_case_145:
		if  lex.data[( lex.p)] == 61 {
			goto tr278
		}
		goto tr277
	st146:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof146
		}
	st_case_146:
		switch  lex.data[( lex.p)] {
		case 43:
			goto tr279
		case 61:
			goto tr280
		}
		goto tr255
	st147:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof147
		}
	st_case_147:
		switch  lex.data[( lex.p)] {
		case 45:
			goto tr281
		case 61:
			goto tr282
		case 62:
			goto tr283
		}
		goto tr255
tr210:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st148
	st148:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof148
		}
	st_case_148:
// line internal/php8/scanner.go:5334
		switch  lex.data[( lex.p)] {
		case 46:
			goto st64
		case 61:
			goto tr285
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr86
		}
		goto tr255
	st64:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		if  lex.data[( lex.p)] == 46 {
			goto tr82
		}
		goto tr23
tr86:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:165
 lex.act = 10;
	goto st149
	st149:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof149
		}
	st_case_149:
// line internal/php8/scanner.go:5366
		switch  lex.data[( lex.p)] {
		case 69:
			goto st65
		case 95:
			goto st67
		case 101:
			goto st65
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr86
		}
		goto tr286
	st65:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch  lex.data[( lex.p)] {
		case 43:
			goto st66
		case 45:
			goto st66
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr84
		}
		goto tr12
	st66:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr84
		}
		goto tr12
tr84:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:165
 lex.act = 10;
	goto st150
	st150:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof150
		}
	st_case_150:
// line internal/php8/scanner.go:5415
		if  lex.data[( lex.p)] == 95 {
			goto st66
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr84
		}
		goto tr286
	st67:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr86
		}
		goto tr85
tr211:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st151
	st151:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof151
		}
	st_case_151:
// line internal/php8/scanner.go:5442
		switch  lex.data[( lex.p)] {
		case 42:
			goto st68
		case 47:
			goto tr262
		case 61:
			goto tr289
		}
		goto tr255
tr89:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st68
tr91:
// line internal/php8/scanner.rl:54

	goto st68
tr93:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st68
	st68:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof68
		}
	st_case_68:
// line internal/php8/scanner.go:5487
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr88
		case 13:
			goto tr89
		case 42:
			goto st70
		}
		goto st68
tr88:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st69
tr92:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st69
	st69:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof69
		}
	st_case_69:
// line internal/php8/scanner.go:5528
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr92
		case 13:
			goto tr93
		case 42:
			goto tr94
		}
		goto tr91
tr94:
// line internal/php8/scanner.rl:54

	goto st70
	st70:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof70
		}
	st_case_70:
// line internal/php8/scanner.go:5547
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr88
		case 13:
			goto tr89
		case 42:
			goto st70
		case 47:
			goto tr95
		}
		goto st68
tr212:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:176
 lex.act = 12;
	goto st152
	st152:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof152
		}
	st_case_152:
// line internal/php8/scanner.go:5571
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr291
		case 69:
			goto st65
		case 95:
			goto st71
		case 98:
			goto st72
		case 101:
			goto st65
		case 120:
			goto st73
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr97
		}
		goto tr290
tr291:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:165
 lex.act = 10;
	goto st153
	st153:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof153
		}
	st_case_153:
// line internal/php8/scanner.go:5602
		switch  lex.data[( lex.p)] {
		case 69:
			goto st65
		case 101:
			goto st65
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr86
		}
		goto tr286
tr97:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:176
 lex.act = 12;
	goto st154
	st154:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof154
		}
	st_case_154:
// line internal/php8/scanner.go:5625
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr291
		case 69:
			goto st65
		case 95:
			goto st71
		case 101:
			goto st65
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr97
		}
		goto tr290
	st71:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr97
		}
		goto tr96
	st72:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 49 {
			goto tr98
		}
		goto tr12
tr98:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:166
 lex.act = 11;
	goto st155
	st155:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof155
		}
	st_case_155:
// line internal/php8/scanner.go:5670
		if  lex.data[( lex.p)] == 95 {
			goto st72
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 49 {
			goto tr98
		}
		goto tr295
	st73:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr99
			}
		case  lex.data[( lex.p)] > 70:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 102 {
				goto tr99
			}
		default:
			goto tr99
		}
		goto tr12
tr99:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:191
 lex.act = 13;
	goto st156
	st156:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof156
		}
	st_case_156:
// line internal/php8/scanner.go:5708
		if  lex.data[( lex.p)] == 95 {
			goto st73
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr99
			}
		case  lex.data[( lex.p)] > 70:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 102 {
				goto tr99
			}
		default:
			goto tr99
		}
		goto tr296
	st157:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof157
		}
	st_case_157:
		if  lex.data[( lex.p)] == 58 {
			goto tr297
		}
		goto tr255
tr214:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st158
	st158:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof158
		}
	st_case_158:
// line internal/php8/scanner.go:5744
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr101
		case 13:
			goto tr102
		case 32:
			goto st74
		case 63:
			goto st77
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st74
		}
		goto tr255
tr104:
// line internal/php8/scanner.rl:54

	goto st74
	st74:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof74
		}
	st_case_74:
// line internal/php8/scanner.go:5768
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr101
		case 13:
			goto tr102
		case 32:
			goto st74
		case 63:
			goto st77
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st74
		}
		goto tr23
tr101:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st75
tr105:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st75
	st75:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof75
		}
	st_case_75:
// line internal/php8/scanner.go:5814
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr105
		case 13:
			goto tr106
		case 32:
			goto tr104
		case 63:
			goto tr107
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr104
		}
		goto tr23
tr102:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st76
tr106:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st76
	st76:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof76
		}
	st_case_76:
// line internal/php8/scanner.go:5860
		if  lex.data[( lex.p)] == 10 {
			goto tr101
		}
		goto tr23
tr107:
// line internal/php8/scanner.rl:54

	goto st77
	st77:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof77
		}
	st_case_77:
// line internal/php8/scanner.go:5874
		if  lex.data[( lex.p)] == 62 {
			goto tr108
		}
		goto tr23
tr108:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st159
	st159:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof159
		}
	st_case_159:
// line internal/php8/scanner.go:5889
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr110
		case 13:
			goto tr299
		}
		goto tr298
tr110:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st160
	st160:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof160
		}
	st_case_160:
// line internal/php8/scanner.go:5914
		goto tr300
tr299:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st78
	st78:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof78
		}
	st_case_78:
// line internal/php8/scanner.go:5933
		if  lex.data[( lex.p)] == 10 {
			goto tr110
		}
		goto tr109
	st161:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch  lex.data[( lex.p)] {
		case 60:
			goto tr301
		case 61:
			goto st164
		case 62:
			goto tr303
		}
		goto tr255
tr301:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:312
 lex.act = 123;
	goto st162
	st162:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof162
		}
	st_case_162:
// line internal/php8/scanner.go:5964
		switch  lex.data[( lex.p)] {
		case 60:
			goto st79
		case 61:
			goto tr305
		}
		goto tr304
	st79:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch  lex.data[( lex.p)] {
		case 9:
			goto st79
		case 32:
			goto st79
		case 34:
			goto st80
		case 39:
			goto st84
		case 96:
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr114
	st80:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		if  lex.data[( lex.p)] == 96 {
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr115
tr115:
// line internal/php8/scanner.rl:35
lblStart = lex.p
	goto st81
	st81:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof81
		}
	st_case_81:
// line internal/php8/scanner.go:6032
		switch  lex.data[( lex.p)] {
		case 34:
			goto tr116
		case 96:
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr12
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st81
tr116:
// line internal/php8/scanner.rl:36
lblEnd = lex.p
	goto st82
	st82:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof82
		}
	st_case_82:
// line internal/php8/scanner.go:6066
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr118
		case 13:
			goto tr119
		}
		goto tr12
tr118:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st163
tr122:
// line internal/php8/scanner.rl:36
lblEnd = lex.p
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st163
	st163:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof163
		}
	st_case_163:
// line internal/php8/scanner.go:6105
		goto tr306
tr119:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st83
tr123:
// line internal/php8/scanner.rl:36
lblEnd = lex.p
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st83
	st83:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof83
		}
	st_case_83:
// line internal/php8/scanner.go:6138
		if  lex.data[( lex.p)] == 10 {
			goto tr118
		}
		goto tr12
	st84:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		if  lex.data[( lex.p)] == 96 {
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr120
tr120:
// line internal/php8/scanner.rl:35
lblStart = lex.p
	goto st85
	st85:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof85
		}
	st_case_85:
// line internal/php8/scanner.go:6173
		switch  lex.data[( lex.p)] {
		case 39:
			goto tr116
		case 96:
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr12
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st85
tr114:
// line internal/php8/scanner.rl:35
lblStart = lex.p
	goto st86
	st86:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof86
		}
	st_case_86:
// line internal/php8/scanner.go:6207
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr122
		case 13:
			goto tr123
		case 96:
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr12
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st86
	st164:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		if  lex.data[( lex.p)] == 62 {
			goto tr308
		}
		goto tr307
	st165:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof165
		}
	st_case_165:
		switch  lex.data[( lex.p)] {
		case 61:
			goto st166
		case 62:
			goto tr310
		}
		goto tr255
	st166:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof166
		}
	st_case_166:
		if  lex.data[( lex.p)] == 61 {
			goto tr312
		}
		goto tr311
	st167:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof167
		}
	st_case_167:
		switch  lex.data[( lex.p)] {
		case 61:
			goto tr313
		case 62:
			goto st168
		}
		goto tr255
	st168:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof168
		}
	st_case_168:
		if  lex.data[( lex.p)] == 61 {
			goto tr316
		}
		goto tr315
tr218:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st169
	st169:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof169
		}
	st_case_169:
// line internal/php8/scanner.go:6295
		switch  lex.data[( lex.p)] {
		case 45:
			goto st87
		case 62:
			goto tr318
		case 63:
			goto st172
		}
		goto tr255
	st87:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		if  lex.data[( lex.p)] == 62 {
			goto tr125
		}
		goto tr23
tr318:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st170
	st170:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof170
		}
	st_case_170:
// line internal/php8/scanner.go:6324
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr127
		case 13:
			goto tr321
		}
		goto tr320
tr127:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st171
	st171:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof171
		}
	st_case_171:
// line internal/php8/scanner.go:6349
		goto tr322
tr321:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st88
	st88:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof88
		}
	st_case_88:
// line internal/php8/scanner.go:6368
		if  lex.data[( lex.p)] == 10 {
			goto tr127
		}
		goto tr126
	st172:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof172
		}
	st_case_172:
		if  lex.data[( lex.p)] == 61 {
			goto tr324
		}
		goto tr323
tr219:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st173
	st173:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof173
		}
	st_case_173:
// line internal/php8/scanner.go:6394
		switch  lex.data[( lex.p)] {
		case 66:
			goto tr326
		case 78:
			goto tr327
		case 82:
			goto tr328
		case 83:
			goto tr329
		case 92:
			goto st89
		case 96:
			goto tr325
		case 98:
			goto tr326
		case 110:
			goto tr327
		case 114:
			goto tr328
		case 115:
			goto tr329
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr226:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st174
tr329:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:209
 lex.act = 20;
	goto st174
tr337:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:207
 lex.act = 18;
	goto st174
tr338:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:281
 lex.act = 92;
	goto st174
tr341:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:208
 lex.act = 19;
	goto st174
tr346:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:210
 lex.act = 21;
	goto st174
tr358:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:211
 lex.act = 22;
	goto st174
tr359:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:212
 lex.act = 23;
	goto st174
tr361:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:213
 lex.act = 24;
	goto st174
tr368:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:238
 lex.act = 49;
	goto st174
tr372:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:214
 lex.act = 25;
	goto st174
tr374:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:215
 lex.act = 26;
	goto st174
tr378:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:216
 lex.act = 27;
	goto st174
tr382:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:217
 lex.act = 28;
	goto st174
tr385:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:220
 lex.act = 31;
	goto st174
tr391:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:218
 lex.act = 29;
	goto st174
tr395:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:219
 lex.act = 30;
	goto st174
tr396:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:232
 lex.act = 43;
	goto st174
tr404:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:221
 lex.act = 32;
	goto st174
tr409:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:223
 lex.act = 34;
	goto st174
tr412:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:224
 lex.act = 35;
	goto st174
tr424:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:225
 lex.act = 36;
	goto st174
tr431:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:227
 lex.act = 38;
	goto st174
tr432:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:228
 lex.act = 39;
	goto st174
tr437:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:229
 lex.act = 40;
	goto st174
tr441:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:230
 lex.act = 41;
	goto st174
tr443:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:231
 lex.act = 42;
	goto st174
tr449:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:233
 lex.act = 44;
	goto st174
tr451:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:239
 lex.act = 50;
	goto st174
tr458:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:235
 lex.act = 46;
	goto st174
tr464:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:237
 lex.act = 48;
	goto st174
tr470:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:240
 lex.act = 51;
	goto st174
tr472:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:241
 lex.act = 52;
	goto st174
tr473:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:242
 lex.act = 53;
	goto st174
tr484:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:244
 lex.act = 55;
	goto st174
tr497:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:268
 lex.act = 79;
	goto st174
tr505:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:245
 lex.act = 56;
	goto st174
tr509:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:246
 lex.act = 57;
	goto st174
tr515:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:247
 lex.act = 58;
	goto st174
tr518:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:243
 lex.act = 54;
	goto st174
tr521:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:248
 lex.act = 59;
	goto st174
tr525:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:257
 lex.act = 68;
	goto st174
tr538:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:280
 lex.act = 91;
	goto st174
tr539:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:282
 lex.act = 93;
	goto st174
tr546:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:252
 lex.act = 63;
	goto st174
tr549:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:250
 lex.act = 61;
	goto st174
tr555:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:253
 lex.act = 64;
	goto st174
tr559:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:251
 lex.act = 62;
	goto st174
tr572:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:270
 lex.act = 81;
	goto st174
tr575:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:254
 lex.act = 65;
	goto st174
tr581:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:255
 lex.act = 66;
	goto st174
tr585:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:256
 lex.act = 67;
	goto st174
tr590:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:258
 lex.act = 69;
	goto st174
tr592:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:260
 lex.act = 71;
	goto st174
tr594:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:259
 lex.act = 70;
	goto st174
tr599:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:261
 lex.act = 72;
	goto st174
tr600:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:262
 lex.act = 73;
	goto st174
tr602:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:263
 lex.act = 74;
	goto st174
tr606:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:264
 lex.act = 75;
	goto st174
tr608:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:283
 lex.act = 94;
	goto st174
tr632:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:271
 lex.act = 82;
	goto st174
tr636:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:272
 lex.act = 83;
	goto st174
tr642:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:273
 lex.act = 84;
	goto st174
tr650:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:274
 lex.act = 85;
	goto st174
tr662:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:279
 lex.act = 90;
	goto st174
tr667:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:275
 lex.act = 86;
	goto st174
tr674:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:277
 lex.act = 88;
	goto st174
tr684:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:276
 lex.act = 87;
	goto st174
tr690:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:278
 lex.act = 89;
	goto st174
	st174:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof174
		}
	st_case_174:
// line internal/php8/scanner.go:6923
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 96:
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr12
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr226
	st89:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		if  lex.data[( lex.p)] == 96 {
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr128
tr128:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:203
 lex.act = 15;
	goto st175
	st175:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof175
		}
	st_case_175:
// line internal/php8/scanner.go:6981
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 96:
			goto tr331
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr331
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr331
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr331
			}
		default:
			goto tr331
		}
		goto tr128
tr326:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st176
	st176:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof176
		}
	st_case_176:
// line internal/php8/scanner.go:7018
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr332
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr332
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr332:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st177
	st177:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof177
		}
	st_case_177:
// line internal/php8/scanner.go:7059
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr333
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr333
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr333:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st178
	st178:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof178
		}
	st_case_178:
// line internal/php8/scanner.go:7100
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr334
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr334
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr334:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st179
	st179:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof179
		}
	st_case_179:
// line internal/php8/scanner.go:7141
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr335
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr335
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr335:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st180
	st180:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof180
		}
	st_case_180:
// line internal/php8/scanner.go:7182
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr336
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr336
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr336:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st181
	st181:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof181
		}
	st_case_181:
// line internal/php8/scanner.go:7223
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr337
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr337
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr327:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st182
	st182:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof182
		}
	st_case_182:
// line internal/php8/scanner.go:7264
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr338
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr338
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr328:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st183
	st183:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof183
		}
	st_case_183:
// line internal/php8/scanner.go:7305
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr339
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr339
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr339:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st184
	st184:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof184
		}
	st_case_184:
// line internal/php8/scanner.go:7346
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr340
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr340
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr340:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st185
	st185:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof185
		}
	st_case_185:
// line internal/php8/scanner.go:7387
		switch  lex.data[( lex.p)] {
		case 89:
			goto tr341
		case 92:
			goto st89
		case 96:
			goto tr325
		case 121:
			goto tr341
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr220:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st186
	st186:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof186
		}
	st_case_186:
// line internal/php8/scanner.go:7428
		switch  lex.data[( lex.p)] {
		case 34:
			goto st7
		case 60:
			goto st90
		case 82:
			goto tr343
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr343
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
	st90:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		if  lex.data[( lex.p)] == 60 {
			goto st91
		}
		goto tr129
	st91:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		if  lex.data[( lex.p)] == 60 {
			goto st79
		}
		goto tr129
tr343:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st187
	st187:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof187
		}
	st_case_187:
// line internal/php8/scanner.go:7491
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr344
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr344
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr344:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st188
	st188:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof188
		}
	st_case_188:
// line internal/php8/scanner.go:7532
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr345
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr345
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr345:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st189
	st189:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof189
		}
	st_case_189:
// line internal/php8/scanner.go:7573
		switch  lex.data[( lex.p)] {
		case 75:
			goto tr346
		case 92:
			goto st89
		case 96:
			goto tr325
		case 107:
			goto tr346
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr221:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st190
	st190:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof190
		}
	st_case_190:
// line internal/php8/scanner.go:7614
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr347
		case 70:
			goto tr348
		case 76:
			goto tr349
		case 79:
			goto tr350
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr347
		case 102:
			goto tr348
		case 108:
			goto tr349
		case 111:
			goto tr350
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr347:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st191
	st191:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof191
		}
	st_case_191:
// line internal/php8/scanner.go:7667
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr351
		case 83:
			goto tr352
		case 84:
			goto tr353
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr351
		case 115:
			goto tr352
		case 116:
			goto tr353
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr351:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st192
	st192:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof192
		}
	st_case_192:
// line internal/php8/scanner.go:7716
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr354
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr354
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr354:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st193
	st193:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof193
		}
	st_case_193:
// line internal/php8/scanner.go:7757
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr355
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr355
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr355:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st194
	st194:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof194
		}
	st_case_194:
// line internal/php8/scanner.go:7798
		switch  lex.data[( lex.p)] {
		case 66:
			goto tr356
		case 92:
			goto st89
		case 96:
			goto tr325
		case 98:
			goto tr356
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr356:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st195
	st195:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof195
		}
	st_case_195:
// line internal/php8/scanner.go:7839
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr357
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr357
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr357:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st196
	st196:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof196
		}
	st_case_196:
// line internal/php8/scanner.go:7880
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr358
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr358
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr352:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st197
	st197:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof197
		}
	st_case_197:
// line internal/php8/scanner.go:7921
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr359
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr359
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr353:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st198
	st198:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof198
		}
	st_case_198:
// line internal/php8/scanner.go:7962
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr360
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr360
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr360:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st199
	st199:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof199
		}
	st_case_199:
// line internal/php8/scanner.go:8003
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr361
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr361
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr348:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st200
	st200:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof200
		}
	st_case_200:
// line internal/php8/scanner.go:8044
		switch  lex.data[( lex.p)] {
		case 85:
			goto tr362
		case 92:
			goto st89
		case 96:
			goto tr325
		case 117:
			goto tr362
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr362:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st201
	st201:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof201
		}
	st_case_201:
// line internal/php8/scanner.go:8085
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr363
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr363
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr363:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st202
	st202:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof202
		}
	st_case_202:
// line internal/php8/scanner.go:8126
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr364
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr364
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr364:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st203
	st203:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof203
		}
	st_case_203:
// line internal/php8/scanner.go:8167
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr365
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr365
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr365:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st204
	st204:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof204
		}
	st_case_204:
// line internal/php8/scanner.go:8208
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr366
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr366
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr366:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st205
	st205:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof205
		}
	st_case_205:
// line internal/php8/scanner.go:8249
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr367
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr367
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr367:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st206
	st206:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof206
		}
	st_case_206:
// line internal/php8/scanner.go:8290
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr368
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr368
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr349:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st207
	st207:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof207
		}
	st_case_207:
// line internal/php8/scanner.go:8331
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr369
		case 79:
			goto tr370
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr369
		case 111:
			goto tr370
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr369:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st208
	st208:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof208
		}
	st_case_208:
// line internal/php8/scanner.go:8376
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr371
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr371
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr371:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st209
	st209:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof209
		}
	st_case_209:
// line internal/php8/scanner.go:8417
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr372
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr372
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr370:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st210
	st210:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof210
		}
	st_case_210:
// line internal/php8/scanner.go:8458
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr373
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr373
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr373:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st211
	st211:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof211
		}
	st_case_211:
// line internal/php8/scanner.go:8499
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr374
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr374
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr350:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st212
	st212:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof212
		}
	st_case_212:
// line internal/php8/scanner.go:8540
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr375
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr375
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr375:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st213
	st213:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof213
		}
	st_case_213:
// line internal/php8/scanner.go:8581
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr376
		case 84:
			goto tr377
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr376
		case 116:
			goto tr377
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr376:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st214
	st214:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof214
		}
	st_case_214:
// line internal/php8/scanner.go:8626
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr378
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr378
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr377:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st215
	st215:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof215
		}
	st_case_215:
// line internal/php8/scanner.go:8667
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr379
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr379
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr379:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st216
	st216:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof216
		}
	st_case_216:
// line internal/php8/scanner.go:8708
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr380
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr380
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr380:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st217
	st217:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof217
		}
	st_case_217:
// line internal/php8/scanner.go:8749
		switch  lex.data[( lex.p)] {
		case 85:
			goto tr381
		case 92:
			goto st89
		case 96:
			goto tr325
		case 117:
			goto tr381
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr381:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st218
	st218:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof218
		}
	st_case_218:
// line internal/php8/scanner.go:8790
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr382
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr382
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr222:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st219
	st219:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof219
		}
	st_case_219:
// line internal/php8/scanner.go:8831
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr383
		case 73:
			goto tr384
		case 79:
			goto tr385
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr383
		case 105:
			goto tr384
		case 111:
			goto tr385
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr383:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st220
	st220:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof220
		}
	st_case_220:
// line internal/php8/scanner.go:8880
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr386
		case 70:
			goto tr387
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr386
		case 102:
			goto tr387
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr386:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st221
	st221:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof221
		}
	st_case_221:
// line internal/php8/scanner.go:8925
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr388
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr388
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr388:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st222
	st222:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof222
		}
	st_case_222:
// line internal/php8/scanner.go:8966
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr389
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr389
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr389:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st223
	st223:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof223
		}
	st_case_223:
// line internal/php8/scanner.go:9007
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr390
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr390
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr390:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st224
	st224:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof224
		}
	st_case_224:
// line internal/php8/scanner.go:9048
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr391
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr391
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr387:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st225
	st225:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof225
		}
	st_case_225:
// line internal/php8/scanner.go:9089
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr392
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr392
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr392:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st226
	st226:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof226
		}
	st_case_226:
// line internal/php8/scanner.go:9130
		switch  lex.data[( lex.p)] {
		case 85:
			goto tr393
		case 92:
			goto st89
		case 96:
			goto tr325
		case 117:
			goto tr393
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr393:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st227
	st227:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof227
		}
	st_case_227:
// line internal/php8/scanner.go:9171
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr394
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr394
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr394:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st228
	st228:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof228
		}
	st_case_228:
// line internal/php8/scanner.go:9212
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr395
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr395
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr384:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st229
	st229:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof229
		}
	st_case_229:
// line internal/php8/scanner.go:9253
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr396
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr396
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr223:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st230
	st230:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof230
		}
	st_case_230:
// line internal/php8/scanner.go:9294
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr397
		case 76:
			goto tr398
		case 77:
			goto tr399
		case 78:
			goto tr400
		case 86:
			goto tr401
		case 88:
			goto tr402
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr397
		case 108:
			goto tr398
		case 109:
			goto tr399
		case 110:
			goto tr400
		case 118:
			goto tr401
		case 120:
			goto tr402
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr397:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st231
	st231:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof231
		}
	st_case_231:
// line internal/php8/scanner.go:9355
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr403
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr403
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr403:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st232
	st232:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof232
		}
	st_case_232:
// line internal/php8/scanner.go:9396
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr404
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr404
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr398:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st233
	st233:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof233
		}
	st_case_233:
// line internal/php8/scanner.go:9437
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr405
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr405
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr405:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st234
	st234:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof234
		}
	st_case_234:
// line internal/php8/scanner.go:9478
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr406
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr406
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr406:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:222
 lex.act = 33;
	goto st235
	st235:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof235
		}
	st_case_235:
// line internal/php8/scanner.go:9519
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr408
		case 92:
			goto st89
		case 96:
			goto tr407
		case 105:
			goto tr408
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr407
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr407
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr407
			}
		default:
			goto tr407
		}
		goto tr226
tr408:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st236
	st236:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof236
		}
	st_case_236:
// line internal/php8/scanner.go:9560
		switch  lex.data[( lex.p)] {
		case 70:
			goto tr409
		case 92:
			goto st89
		case 96:
			goto tr325
		case 102:
			goto tr409
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr399:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st237
	st237:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof237
		}
	st_case_237:
// line internal/php8/scanner.go:9601
		switch  lex.data[( lex.p)] {
		case 80:
			goto tr410
		case 92:
			goto st89
		case 96:
			goto tr325
		case 112:
			goto tr410
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr410:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st238
	st238:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof238
		}
	st_case_238:
// line internal/php8/scanner.go:9642
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr411
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr411
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr411:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st239
	st239:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof239
		}
	st_case_239:
// line internal/php8/scanner.go:9683
		switch  lex.data[( lex.p)] {
		case 89:
			goto tr412
		case 92:
			goto st89
		case 96:
			goto tr325
		case 121:
			goto tr412
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr400:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st240
	st240:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof240
		}
	st_case_240:
// line internal/php8/scanner.go:9724
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr413
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr413
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr413:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st241
	st241:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof241
		}
	st_case_241:
// line internal/php8/scanner.go:9765
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr414
		case 70:
			goto tr415
		case 73:
			goto tr416
		case 83:
			goto tr417
		case 87:
			goto tr418
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr414
		case 102:
			goto tr415
		case 105:
			goto tr416
		case 115:
			goto tr417
		case 119:
			goto tr418
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr414:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st242
	st242:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof242
		}
	st_case_242:
// line internal/php8/scanner.go:9822
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr419
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr419
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr419:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st243
	st243:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof243
		}
	st_case_243:
// line internal/php8/scanner.go:9863
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr420
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr420
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr420:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st244
	st244:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof244
		}
	st_case_244:
// line internal/php8/scanner.go:9904
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr421
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr421
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr421:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st245
	st245:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof245
		}
	st_case_245:
// line internal/php8/scanner.go:9945
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr422
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr422
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr422:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st246
	st246:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof246
		}
	st_case_246:
// line internal/php8/scanner.go:9986
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr423
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr423
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr423:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st247
	st247:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof247
		}
	st_case_247:
// line internal/php8/scanner.go:10027
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr424
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr424
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr415:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st248
	st248:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof248
		}
	st_case_248:
// line internal/php8/scanner.go:10068
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr425
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr425
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr425:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st249
	st249:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof249
		}
	st_case_249:
// line internal/php8/scanner.go:10109
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr426
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr426
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr426:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:226
 lex.act = 37;
	goto st250
	st250:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof250
		}
	st_case_250:
// line internal/php8/scanner.go:10150
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr428
		case 92:
			goto st89
		case 96:
			goto tr427
		case 101:
			goto tr428
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr427
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr427
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr427
			}
		default:
			goto tr427
		}
		goto tr226
tr428:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st251
	st251:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof251
		}
	st_case_251:
// line internal/php8/scanner.go:10191
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr429
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr429
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr429:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st252
	st252:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof252
		}
	st_case_252:
// line internal/php8/scanner.go:10232
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr430
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr430
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr430:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st253
	st253:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof253
		}
	st_case_253:
// line internal/php8/scanner.go:10273
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr431
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr431
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr416:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st254
	st254:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof254
		}
	st_case_254:
// line internal/php8/scanner.go:10314
		switch  lex.data[( lex.p)] {
		case 70:
			goto tr432
		case 92:
			goto st89
		case 96:
			goto tr325
		case 102:
			goto tr432
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr417:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st255
	st255:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof255
		}
	st_case_255:
// line internal/php8/scanner.go:10355
		switch  lex.data[( lex.p)] {
		case 87:
			goto tr433
		case 92:
			goto st89
		case 96:
			goto tr325
		case 119:
			goto tr433
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr433:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st256
	st256:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof256
		}
	st_case_256:
// line internal/php8/scanner.go:10396
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr434
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr434
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr434:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st257
	st257:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof257
		}
	st_case_257:
// line internal/php8/scanner.go:10437
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr435
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr435
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr435:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st258
	st258:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof258
		}
	st_case_258:
// line internal/php8/scanner.go:10478
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr436
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr436
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr436:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st259
	st259:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof259
		}
	st_case_259:
// line internal/php8/scanner.go:10519
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr437
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr437
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr418:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st260
	st260:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof260
		}
	st_case_260:
// line internal/php8/scanner.go:10560
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr438
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr438
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr438:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st261
	st261:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof261
		}
	st_case_261:
// line internal/php8/scanner.go:10601
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr439
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr439
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr439:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st262
	st262:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof262
		}
	st_case_262:
// line internal/php8/scanner.go:10642
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr440
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr440
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr440:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st263
	st263:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof263
		}
	st_case_263:
// line internal/php8/scanner.go:10683
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr441
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr441
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr401:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st264
	st264:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof264
		}
	st_case_264:
// line internal/php8/scanner.go:10724
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr442
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr442
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr442:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st265
	st265:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof265
		}
	st_case_265:
// line internal/php8/scanner.go:10765
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr443
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr443
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr402:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st266
	st266:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof266
		}
	st_case_266:
// line internal/php8/scanner.go:10806
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr444
		case 84:
			goto tr445
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr444
		case 116:
			goto tr445
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr444:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st267
	st267:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof267
		}
	st_case_267:
// line internal/php8/scanner.go:10851
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr396
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr396
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr445:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st268
	st268:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof268
		}
	st_case_268:
// line internal/php8/scanner.go:10892
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr446
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr446
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr446:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st269
	st269:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof269
		}
	st_case_269:
// line internal/php8/scanner.go:10933
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr447
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr447
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr447:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st270
	st270:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof270
		}
	st_case_270:
// line internal/php8/scanner.go:10974
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr448
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr448
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr448:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st271
	st271:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof271
		}
	st_case_271:
// line internal/php8/scanner.go:11015
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr449
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr449
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr224:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st272
	st272:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof272
		}
	st_case_272:
// line internal/php8/scanner.go:11056
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr450
		case 78:
			goto tr451
		case 79:
			goto tr452
		case 85:
			goto tr362
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr450
		case 110:
			goto tr451
		case 111:
			goto tr452
		case 117:
			goto tr362
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr450:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st273
	st273:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof273
		}
	st_case_273:
// line internal/php8/scanner.go:11109
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr453
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr453
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr453:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st274
	st274:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof274
		}
	st_case_274:
// line internal/php8/scanner.go:11150
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr454
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr454
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr454:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st275
	st275:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof275
		}
	st_case_275:
// line internal/php8/scanner.go:11191
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr455
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr455
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr455:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:234
 lex.act = 45;
	goto st276
	st276:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof276
		}
	st_case_276:
// line internal/php8/scanner.go:11232
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr457
		case 92:
			goto st89
		case 96:
			goto tr456
		case 108:
			goto tr457
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr456
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr456
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr456
			}
		default:
			goto tr456
		}
		goto tr226
tr457:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st277
	st277:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof277
		}
	st_case_277:
// line internal/php8/scanner.go:11273
		switch  lex.data[( lex.p)] {
		case 89:
			goto tr458
		case 92:
			goto st89
		case 96:
			goto tr325
		case 121:
			goto tr458
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr452:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st278
	st278:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof278
		}
	st_case_278:
// line internal/php8/scanner.go:11314
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr459
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr459
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr459:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:236
 lex.act = 47;
	goto st279
	st279:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof279
		}
	st_case_279:
// line internal/php8/scanner.go:11355
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr461
		case 92:
			goto st89
		case 96:
			goto tr460
		case 101:
			goto tr461
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr460
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr460
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr460
			}
		default:
			goto tr460
		}
		goto tr226
tr461:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st280
	st280:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof280
		}
	st_case_280:
// line internal/php8/scanner.go:11396
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr462
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr462
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr462:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st281
	st281:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof281
		}
	st_case_281:
// line internal/php8/scanner.go:11437
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr463
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr463
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr463:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st282
	st282:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof282
		}
	st_case_282:
// line internal/php8/scanner.go:11478
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr464
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr464
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr225:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st283
	st283:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof283
		}
	st_case_283:
// line internal/php8/scanner.go:11519
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr465
		case 79:
			goto tr466
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr465
		case 111:
			goto tr466
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr465:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st284
	st284:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof284
		}
	st_case_284:
// line internal/php8/scanner.go:11564
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr467
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr467
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr467:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st285
	st285:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof285
		}
	st_case_285:
// line internal/php8/scanner.go:11605
		switch  lex.data[( lex.p)] {
		case 66:
			goto tr468
		case 92:
			goto st89
		case 96:
			goto tr325
		case 98:
			goto tr468
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr468:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st286
	st286:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof286
		}
	st_case_286:
// line internal/php8/scanner.go:11646
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr469
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr469
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr469:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st287
	st287:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof287
		}
	st_case_287:
// line internal/php8/scanner.go:11687
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr470
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr470
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr466:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st288
	st288:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof288
		}
	st_case_288:
// line internal/php8/scanner.go:11728
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr471
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr471
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr471:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st289
	st289:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof289
		}
	st_case_289:
// line internal/php8/scanner.go:11769
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr472
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr472
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr227:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st290
	st290:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof290
		}
	st_case_290:
// line internal/php8/scanner.go:11810
		switch  lex.data[( lex.p)] {
		case 70:
			goto tr473
		case 77:
			goto tr474
		case 78:
			goto tr475
		case 83:
			goto tr476
		case 92:
			goto st89
		case 96:
			goto tr325
		case 102:
			goto tr473
		case 109:
			goto tr474
		case 110:
			goto tr475
		case 115:
			goto tr476
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr474:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st291
	st291:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof291
		}
	st_case_291:
// line internal/php8/scanner.go:11863
		switch  lex.data[( lex.p)] {
		case 80:
			goto tr477
		case 92:
			goto st89
		case 96:
			goto tr325
		case 112:
			goto tr477
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr477:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st292
	st292:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof292
		}
	st_case_292:
// line internal/php8/scanner.go:11904
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr478
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr478
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr478:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st293
	st293:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof293
		}
	st_case_293:
// line internal/php8/scanner.go:11945
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr479
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr479
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr479:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st294
	st294:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof294
		}
	st_case_294:
// line internal/php8/scanner.go:11986
		switch  lex.data[( lex.p)] {
		case 77:
			goto tr480
		case 92:
			goto st89
		case 96:
			goto tr325
		case 109:
			goto tr480
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr480:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st295
	st295:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof295
		}
	st_case_295:
// line internal/php8/scanner.go:12027
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr481
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr481
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr481:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st296
	st296:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof296
		}
	st_case_296:
// line internal/php8/scanner.go:12068
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr482
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr482
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr482:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st297
	st297:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof297
		}
	st_case_297:
// line internal/php8/scanner.go:12109
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr483
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr483
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr483:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st298
	st298:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof298
		}
	st_case_298:
// line internal/php8/scanner.go:12150
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr484
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr484
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr475:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st299
	st299:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof299
		}
	st_case_299:
// line internal/php8/scanner.go:12191
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr485
		case 83:
			goto tr486
		case 84:
			goto tr487
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr485
		case 115:
			goto tr486
		case 116:
			goto tr487
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr485:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st300
	st300:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof300
		}
	st_case_300:
// line internal/php8/scanner.go:12240
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr488
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr488
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr488:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st301
	st301:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof301
		}
	st_case_301:
// line internal/php8/scanner.go:12281
		switch  lex.data[( lex.p)] {
		case 85:
			goto tr489
		case 92:
			goto st89
		case 96:
			goto tr325
		case 117:
			goto tr489
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr489:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st302
	st302:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof302
		}
	st_case_302:
// line internal/php8/scanner.go:12322
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr490
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr490
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr490:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st303
	st303:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof303
		}
	st_case_303:
// line internal/php8/scanner.go:12363
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr491
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr491
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr491:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:267
 lex.act = 78;
	goto st304
	st304:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof304
		}
	st_case_304:
// line internal/php8/scanner.go:12404
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr493
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr492
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr492
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr492
			}
		default:
			goto tr492
		}
		goto tr226
tr493:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st305
	st305:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof305
		}
	st_case_305:
// line internal/php8/scanner.go:12441
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr494
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr494
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr494:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st306
	st306:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof306
		}
	st_case_306:
// line internal/php8/scanner.go:12482
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr495
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr495
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr495:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st307
	st307:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof307
		}
	st_case_307:
// line internal/php8/scanner.go:12523
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr496
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr496
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr496:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st308
	st308:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof308
		}
	st_case_308:
// line internal/php8/scanner.go:12564
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr497
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr497
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr486:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st309
	st309:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof309
		}
	st_case_309:
// line internal/php8/scanner.go:12605
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr498
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr498
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr498:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st310
	st310:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof310
		}
	st_case_310:
// line internal/php8/scanner.go:12646
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr499
		case 69:
			goto tr500
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr499
		case 101:
			goto tr500
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr499:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st311
	st311:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof311
		}
	st_case_311:
// line internal/php8/scanner.go:12691
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr501
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr501
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr501:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st312
	st312:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof312
		}
	st_case_312:
// line internal/php8/scanner.go:12732
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr502
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr502
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr502:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st313
	st313:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof313
		}
	st_case_313:
// line internal/php8/scanner.go:12773
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr503
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr503
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr503:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st314
	st314:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof314
		}
	st_case_314:
// line internal/php8/scanner.go:12814
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr504
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr504
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr504:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st315
	st315:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof315
		}
	st_case_315:
// line internal/php8/scanner.go:12855
		switch  lex.data[( lex.p)] {
		case 70:
			goto tr505
		case 92:
			goto st89
		case 96:
			goto tr325
		case 102:
			goto tr505
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr500:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st316
	st316:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof316
		}
	st_case_316:
// line internal/php8/scanner.go:12896
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr506
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr506
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr506:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st317
	st317:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof317
		}
	st_case_317:
// line internal/php8/scanner.go:12937
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr507
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr507
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr507:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st318
	st318:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof318
		}
	st_case_318:
// line internal/php8/scanner.go:12978
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr508
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr508
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr508:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st319
	st319:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof319
		}
	st_case_319:
// line internal/php8/scanner.go:13019
		switch  lex.data[( lex.p)] {
		case 70:
			goto tr509
		case 92:
			goto st89
		case 96:
			goto tr325
		case 102:
			goto tr509
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr487:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st320
	st320:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof320
		}
	st_case_320:
// line internal/php8/scanner.go:13060
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr510
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr510
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr510:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st321
	st321:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof321
		}
	st_case_321:
// line internal/php8/scanner.go:13101
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr511
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr511
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr511:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st322
	st322:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof322
		}
	st_case_322:
// line internal/php8/scanner.go:13142
		switch  lex.data[( lex.p)] {
		case 70:
			goto tr512
		case 92:
			goto st89
		case 96:
			goto tr325
		case 102:
			goto tr512
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr512:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st323
	st323:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof323
		}
	st_case_323:
// line internal/php8/scanner.go:13183
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr513
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr513
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr513:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st324
	st324:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof324
		}
	st_case_324:
// line internal/php8/scanner.go:13224
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr514
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr514
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr514:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st325
	st325:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof325
		}
	st_case_325:
// line internal/php8/scanner.go:13265
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr515
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr515
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr476:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st326
	st326:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof326
		}
	st_case_326:
// line internal/php8/scanner.go:13306
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr516
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr516
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr516:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st327
	st327:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof327
		}
	st_case_327:
// line internal/php8/scanner.go:13347
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr517
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr517
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr517:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st328
	st328:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof328
		}
	st_case_328:
// line internal/php8/scanner.go:13388
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr518
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr518
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr228:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st329
	st329:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof329
		}
	st_case_329:
// line internal/php8/scanner.go:13429
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr519
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr519
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr519:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st330
	st330:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof330
		}
	st_case_330:
// line internal/php8/scanner.go:13470
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr520
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr520
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr520:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st331
	st331:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof331
		}
	st_case_331:
// line internal/php8/scanner.go:13511
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr521
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr521
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr229:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st332
	st332:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof332
		}
	st_case_332:
// line internal/php8/scanner.go:13552
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr522
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr522
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr522:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st333
	st333:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof333
		}
	st_case_333:
// line internal/php8/scanner.go:13593
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr523
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr523
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr523:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st334
	st334:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof334
		}
	st_case_334:
// line internal/php8/scanner.go:13634
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr524
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr524
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr524:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st335
	st335:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof335
		}
	st_case_335:
// line internal/php8/scanner.go:13675
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr525
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr525
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr230:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st336
	st336:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof336
		}
	st_case_336:
// line internal/php8/scanner.go:13716
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr526
		case 69:
			goto tr527
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr526
		case 101:
			goto tr527
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr526:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st337
	st337:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof337
		}
	st_case_337:
// line internal/php8/scanner.go:13761
		switch  lex.data[( lex.p)] {
		case 77:
			goto tr528
		case 92:
			goto st89
		case 96:
			goto tr325
		case 109:
			goto tr528
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr528:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st338
	st338:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof338
		}
	st_case_338:
// line internal/php8/scanner.go:13802
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr529
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr529
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr529:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st339
	st339:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof339
		}
	st_case_339:
// line internal/php8/scanner.go:13843
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr530
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr530
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr530:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st340
	st340:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof340
		}
	st_case_340:
// line internal/php8/scanner.go:13884
		switch  lex.data[( lex.p)] {
		case 80:
			goto tr531
		case 92:
			goto st89
		case 96:
			goto tr325
		case 112:
			goto tr531
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr531:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st341
	st341:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof341
		}
	st_case_341:
// line internal/php8/scanner.go:13925
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr532
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr532
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr532:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st342
	st342:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof342
		}
	st_case_342:
// line internal/php8/scanner.go:13966
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr533
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr533
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr533:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st343
	st343:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof343
		}
	st_case_343:
// line internal/php8/scanner.go:14007
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr534
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr534
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr534:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:249
 lex.act = 60;
	goto st344
	st344:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof344
		}
	st_case_344:
// line internal/php8/scanner.go:14048
		switch  lex.data[( lex.p)] {
		case 92:
			goto st92
		case 96:
			goto tr535
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr535
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr535
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr226
	st92:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		if  lex.data[( lex.p)] == 96 {
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr12
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr131
tr131:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:202
 lex.act = 14;
	goto st345
	st345:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof345
		}
	st_case_345:
// line internal/php8/scanner.go:14106
		switch  lex.data[( lex.p)] {
		case 92:
			goto st92
		case 96:
			goto tr537
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr537
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr537
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr537
			}
		default:
			goto tr537
		}
		goto tr131
tr527:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st346
	st346:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof346
		}
	st_case_346:
// line internal/php8/scanner.go:14143
		switch  lex.data[( lex.p)] {
		case 87:
			goto tr538
		case 92:
			goto st89
		case 96:
			goto tr325
		case 119:
			goto tr538
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr231:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st347
	st347:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof347
		}
	st_case_347:
// line internal/php8/scanner.go:14184
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr539
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr539
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr232:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st348
	st348:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof348
		}
	st_case_348:
// line internal/php8/scanner.go:14225
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr540
		case 85:
			goto tr541
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr540
		case 117:
			goto tr541
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr540:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st349
	st349:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof349
		}
	st_case_349:
// line internal/php8/scanner.go:14270
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr542
		case 79:
			goto tr543
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr542
		case 111:
			goto tr543
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr542:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st350
	st350:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof350
		}
	st_case_350:
// line internal/php8/scanner.go:14315
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr544
		case 86:
			goto tr545
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr544
		case 118:
			goto tr545
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr544:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st351
	st351:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof351
		}
	st_case_351:
// line internal/php8/scanner.go:14360
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr546
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr546
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr545:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st352
	st352:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof352
		}
	st_case_352:
// line internal/php8/scanner.go:14401
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr547
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr547
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr547:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st353
	st353:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof353
		}
	st_case_353:
// line internal/php8/scanner.go:14442
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr548
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr548
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr548:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st354
	st354:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof354
		}
	st_case_354:
// line internal/php8/scanner.go:14483
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr549
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr549
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr543:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st355
	st355:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof355
		}
	st_case_355:
// line internal/php8/scanner.go:14524
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr550
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr550
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr550:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st356
	st356:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof356
		}
	st_case_356:
// line internal/php8/scanner.go:14565
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr551
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr551
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr551:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st357
	st357:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof357
		}
	st_case_357:
// line internal/php8/scanner.go:14606
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr552
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr552
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr552:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st358
	st358:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof358
		}
	st_case_358:
// line internal/php8/scanner.go:14647
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr553
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr553
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr553:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st359
	st359:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof359
		}
	st_case_359:
// line internal/php8/scanner.go:14688
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr554
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr554
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr554:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st360
	st360:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof360
		}
	st_case_360:
// line internal/php8/scanner.go:14729
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr555
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr555
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr541:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st361
	st361:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof361
		}
	st_case_361:
// line internal/php8/scanner.go:14770
		switch  lex.data[( lex.p)] {
		case 66:
			goto tr556
		case 92:
			goto st89
		case 96:
			goto tr325
		case 98:
			goto tr556
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr556:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st362
	st362:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof362
		}
	st_case_362:
// line internal/php8/scanner.go:14811
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr557
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr557
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr557:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st363
	st363:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof363
		}
	st_case_363:
// line internal/php8/scanner.go:14852
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr558
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr558
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr558:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st364
	st364:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof364
		}
	st_case_364:
// line internal/php8/scanner.go:14893
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr559
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr559
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr233:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st365
	st365:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof365
		}
	st_case_365:
// line internal/php8/scanner.go:14934
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr560
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr560
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr560:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st366
	st366:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof366
		}
	st_case_366:
// line internal/php8/scanner.go:14975
		switch  lex.data[( lex.p)] {
		case 81:
			goto tr561
		case 84:
			goto tr562
		case 92:
			goto st89
		case 96:
			goto tr325
		case 113:
			goto tr561
		case 116:
			goto tr562
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr561:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st367
	st367:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof367
		}
	st_case_367:
// line internal/php8/scanner.go:15020
		switch  lex.data[( lex.p)] {
		case 85:
			goto tr563
		case 92:
			goto st89
		case 96:
			goto tr325
		case 117:
			goto tr563
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr563:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st368
	st368:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof368
		}
	st_case_368:
// line internal/php8/scanner.go:15061
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr564
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr564
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr564:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st369
	st369:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof369
		}
	st_case_369:
// line internal/php8/scanner.go:15102
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr565
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr565
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr565:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st370
	st370:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof370
		}
	st_case_370:
// line internal/php8/scanner.go:15143
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr566
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr566
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr566:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:269
 lex.act = 80;
	goto st371
	st371:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof371
		}
	st_case_371:
// line internal/php8/scanner.go:15184
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr568
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr567
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr567
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr567
			}
		default:
			goto tr567
		}
		goto tr226
tr568:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st372
	st372:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof372
		}
	st_case_372:
// line internal/php8/scanner.go:15221
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr569
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr569
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr569:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st373
	st373:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof373
		}
	st_case_373:
// line internal/php8/scanner.go:15262
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr570
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr570
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr570:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st374
	st374:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof374
		}
	st_case_374:
// line internal/php8/scanner.go:15303
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr571
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr571
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr571:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st375
	st375:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof375
		}
	st_case_375:
// line internal/php8/scanner.go:15344
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr572
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr572
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr562:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st376
	st376:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof376
		}
	st_case_376:
// line internal/php8/scanner.go:15385
		switch  lex.data[( lex.p)] {
		case 85:
			goto tr573
		case 92:
			goto st89
		case 96:
			goto tr325
		case 117:
			goto tr573
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr573:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st377
	st377:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof377
		}
	st_case_377:
// line internal/php8/scanner.go:15426
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr574
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr574
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr574:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st378
	st378:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof378
		}
	st_case_378:
// line internal/php8/scanner.go:15467
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr575
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr575
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr234:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st379
	st379:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof379
		}
	st_case_379:
// line internal/php8/scanner.go:15508
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr576
		case 87:
			goto tr577
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr576
		case 119:
			goto tr577
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr576:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st380
	st380:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof380
		}
	st_case_380:
// line internal/php8/scanner.go:15553
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr578
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr578
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr578:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st381
	st381:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof381
		}
	st_case_381:
// line internal/php8/scanner.go:15594
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr579
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr579
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr579:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st382
	st382:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof382
		}
	st_case_382:
// line internal/php8/scanner.go:15635
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr580
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr580
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr580:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st383
	st383:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof383
		}
	st_case_383:
// line internal/php8/scanner.go:15676
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr581
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr581
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr577:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st384
	st384:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof384
		}
	st_case_384:
// line internal/php8/scanner.go:15717
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr582
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr582
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr582:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st385
	st385:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof385
		}
	st_case_385:
// line internal/php8/scanner.go:15758
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr583
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr583
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr583:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st386
	st386:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof386
		}
	st_case_386:
// line internal/php8/scanner.go:15799
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr584
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr584
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr584:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st387
	st387:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof387
		}
	st_case_387:
// line internal/php8/scanner.go:15840
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr585
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr585
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr235:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st388
	st388:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof388
		}
	st_case_388:
// line internal/php8/scanner.go:15881
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr586
		case 82:
			goto tr587
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr586
		case 114:
			goto tr587
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr586:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st389
	st389:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof389
		}
	st_case_389:
// line internal/php8/scanner.go:15926
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr588
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr588
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr588:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st390
	st390:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof390
		}
	st_case_390:
// line internal/php8/scanner.go:15967
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr589
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr589
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr589:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st391
	st391:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof391
		}
	st_case_391:
// line internal/php8/scanner.go:16008
		switch  lex.data[( lex.p)] {
		case 87:
			goto tr590
		case 92:
			goto st89
		case 96:
			goto tr325
		case 119:
			goto tr590
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr587:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st392
	st392:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof392
		}
	st_case_392:
// line internal/php8/scanner.go:16049
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr591
		case 89:
			goto tr592
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr591
		case 121:
			goto tr592
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr591:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st393
	st393:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof393
		}
	st_case_393:
// line internal/php8/scanner.go:16094
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr593
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr593
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr593:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st394
	st394:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof394
		}
	st_case_394:
// line internal/php8/scanner.go:16135
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr594
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr594
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr236:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st395
	st395:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof395
		}
	st_case_395:
// line internal/php8/scanner.go:16176
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr595
		case 83:
			goto tr596
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr595
		case 115:
			goto tr596
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr595:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st396
	st396:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof396
		}
	st_case_396:
// line internal/php8/scanner.go:16221
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr597
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr597
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr597:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st397
	st397:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof397
		}
	st_case_397:
// line internal/php8/scanner.go:16262
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr598
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr598
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr598:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st398
	st398:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof398
		}
	st_case_398:
// line internal/php8/scanner.go:16303
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr599
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr599
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr596:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st399
	st399:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof399
		}
	st_case_399:
// line internal/php8/scanner.go:16344
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr600
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr600
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr237:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st400
	st400:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof400
		}
	st_case_400:
// line internal/php8/scanner.go:16385
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr601
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr601
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr601:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st401
	st401:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof401
		}
	st_case_401:
// line internal/php8/scanner.go:16426
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr602
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr602
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr238:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st402
	st402:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof402
		}
	st_case_402:
// line internal/php8/scanner.go:16467
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr603
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr603
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr603:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st403
	st403:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof403
		}
	st_case_403:
// line internal/php8/scanner.go:16508
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr604
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr604
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr604:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st404
	st404:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof404
		}
	st_case_404:
// line internal/php8/scanner.go:16549
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr605
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr605
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr605:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st405
	st405:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof405
		}
	st_case_405:
// line internal/php8/scanner.go:16590
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr606
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr606
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr239:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st406
	st406:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof406
		}
	st_case_406:
// line internal/php8/scanner.go:16631
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr607
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr607
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr607:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st407
	st407:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof407
		}
	st_case_407:
// line internal/php8/scanner.go:16672
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr608
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr608
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr240:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st408
	st408:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof408
		}
	st_case_408:
// line internal/php8/scanner.go:16713
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr609
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr609
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr609:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st409
	st409:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof409
		}
	st_case_409:
// line internal/php8/scanner.go:16754
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr610
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr610
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr610:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st410
	st410:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof410
		}
	st_case_410:
// line internal/php8/scanner.go:16795
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr611
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr611
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr611:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st411
	st411:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof411
		}
	st_case_411:
// line internal/php8/scanner.go:16836
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr612
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr612
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr612:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:266
 lex.act = 77;
	goto st412
	st412:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof412
		}
	st_case_412:
// line internal/php8/scanner.go:16877
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr134
		case 13:
			goto tr135
		case 32:
			goto st93
		case 92:
			goto st89
		case 96:
			goto tr613
		}
		switch {
		case  lex.data[( lex.p)] < 14:
			switch {
			case  lex.data[( lex.p)] > 8:
				if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
					goto st93
				}
			default:
				goto tr613
			}
		case  lex.data[( lex.p)] > 47:
			switch {
			case  lex.data[( lex.p)] < 91:
				if 58 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 64 {
					goto tr613
				}
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr613
				}
			default:
				goto tr613
			}
		default:
			goto tr613
		}
		goto tr226
tr137:
// line internal/php8/scanner.rl:54

	goto st93
	st93:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof93
		}
	st_case_93:
// line internal/php8/scanner.go:16926
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr134
		case 13:
			goto tr135
		case 32:
			goto st93
		case 70:
			goto st96
		case 102:
			goto st96
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto st93
		}
		goto tr132
tr134:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st94
tr138:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st94
	st94:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof94
		}
	st_case_94:
// line internal/php8/scanner.go:16974
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr138
		case 13:
			goto tr139
		case 32:
			goto tr137
		case 70:
			goto tr140
		case 102:
			goto tr140
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr137
		}
		goto tr132
tr135:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st95
tr139:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st95
	st95:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof95
		}
	st_case_95:
// line internal/php8/scanner.go:17022
		if  lex.data[( lex.p)] == 10 {
			goto tr134
		}
		goto tr132
tr140:
// line internal/php8/scanner.rl:54

	goto st96
	st96:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof96
		}
	st_case_96:
// line internal/php8/scanner.go:17036
		switch  lex.data[( lex.p)] {
		case 82:
			goto st97
		case 114:
			goto st97
		}
		goto tr132
	st97:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch  lex.data[( lex.p)] {
		case 79:
			goto st98
		case 111:
			goto st98
		}
		goto tr132
	st98:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch  lex.data[( lex.p)] {
		case 77:
			goto tr143
		case 109:
			goto tr143
		}
		goto tr132
	st413:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof413
		}
	st_case_413:
		if  lex.data[( lex.p)] == 96 {
			goto tr614
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr614
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto tr145
tr145:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st414
	st414:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof414
		}
	st_case_414:
// line internal/php8/scanner.go:17099
		switch  lex.data[( lex.p)] {
		case 92:
			goto st99
		case 96:
			goto tr615
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr615
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr615
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr615
			}
		default:
			goto tr615
		}
		goto tr145
	st99:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		if  lex.data[( lex.p)] == 96 {
			goto tr144
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr144
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr144
			}
		default:
			goto tr144
		}
		goto tr145
	st415:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof415
		}
	st_case_415:
		if  lex.data[( lex.p)] == 61 {
			goto tr617
		}
		goto tr255
tr243:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st416
	st416:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof416
		}
	st_case_416:
// line internal/php8/scanner.go:17166
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr618
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr618:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st417
	st417:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof417
		}
	st_case_417:
// line internal/php8/scanner.go:17203
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr619
		case 68:
			goto tr620
		case 70:
			goto tr621
		case 72:
			goto tr622
		case 76:
			goto tr623
		case 77:
			goto tr624
		case 78:
			goto tr625
		case 84:
			goto tr626
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr619
		case 100:
			goto tr620
		case 102:
			goto tr621
		case 104:
			goto tr622
		case 108:
			goto tr623
		case 109:
			goto tr624
		case 110:
			goto tr625
		case 116:
			goto tr626
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr619:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st418
	st418:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof418
		}
	st_case_418:
// line internal/php8/scanner.go:17272
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr627
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr627
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr627:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st419
	st419:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof419
		}
	st_case_419:
// line internal/php8/scanner.go:17313
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr628
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr628
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr628:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st420
	st420:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof420
		}
	st_case_420:
// line internal/php8/scanner.go:17354
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr629
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr629
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr629:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st421
	st421:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof421
		}
	st_case_421:
// line internal/php8/scanner.go:17395
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr630
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr630
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr630:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st422
	st422:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof422
		}
	st_case_422:
// line internal/php8/scanner.go:17436
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr631
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr631:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st423
	st423:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof423
		}
	st_case_423:
// line internal/php8/scanner.go:17473
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr632
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr620:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st424
	st424:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof424
		}
	st_case_424:
// line internal/php8/scanner.go:17510
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr633
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr633
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr633:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st425
	st425:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof425
		}
	st_case_425:
// line internal/php8/scanner.go:17551
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr634
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr634
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr634:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st426
	st426:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof426
		}
	st_case_426:
// line internal/php8/scanner.go:17592
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr635
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr635:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st427
	st427:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof427
		}
	st_case_427:
// line internal/php8/scanner.go:17629
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr636
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr621:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st428
	st428:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof428
		}
	st_case_428:
// line internal/php8/scanner.go:17666
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr637
		case 85:
			goto tr638
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr637
		case 117:
			goto tr638
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr637:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st429
	st429:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof429
		}
	st_case_429:
// line internal/php8/scanner.go:17711
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr639
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr639
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr639:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st430
	st430:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof430
		}
	st_case_430:
// line internal/php8/scanner.go:17752
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr640
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr640
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr640:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st431
	st431:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof431
		}
	st_case_431:
// line internal/php8/scanner.go:17793
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr641
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr641:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st432
	st432:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof432
		}
	st_case_432:
// line internal/php8/scanner.go:17830
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr642
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr638:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st433
	st433:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof433
		}
	st_case_433:
// line internal/php8/scanner.go:17867
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr643
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr643
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr643:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st434
	st434:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof434
		}
	st_case_434:
// line internal/php8/scanner.go:17908
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr644
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr644
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr644:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st435
	st435:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof435
		}
	st_case_435:
// line internal/php8/scanner.go:17949
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr645
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr645
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr645:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st436
	st436:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof436
		}
	st_case_436:
// line internal/php8/scanner.go:17990
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr646
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr646
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr646:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st437
	st437:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof437
		}
	st_case_437:
// line internal/php8/scanner.go:18031
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr647
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr647
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr647:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st438
	st438:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof438
		}
	st_case_438:
// line internal/php8/scanner.go:18072
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr648
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr648
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr648:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st439
	st439:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof439
		}
	st_case_439:
// line internal/php8/scanner.go:18113
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr649
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr649:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st440
	st440:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof440
		}
	st_case_440:
// line internal/php8/scanner.go:18150
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr650
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr622:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st441
	st441:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof441
		}
	st_case_441:
// line internal/php8/scanner.go:18187
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr651
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr651
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr651:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st442
	st442:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof442
		}
	st_case_442:
// line internal/php8/scanner.go:18228
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr652
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr652
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr652:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st443
	st443:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof443
		}
	st_case_443:
// line internal/php8/scanner.go:18269
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr653
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr653
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr653:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st444
	st444:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof444
		}
	st_case_444:
// line internal/php8/scanner.go:18310
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr654
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr654:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st445
	st445:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof445
		}
	st_case_445:
// line internal/php8/scanner.go:18347
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr655
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr655
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr655:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st446
	st446:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof446
		}
	st_case_446:
// line internal/php8/scanner.go:18388
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr656
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr656
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr656:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st447
	st447:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof447
		}
	st_case_447:
// line internal/php8/scanner.go:18429
		switch  lex.data[( lex.p)] {
		case 77:
			goto tr657
		case 92:
			goto st89
		case 96:
			goto tr325
		case 109:
			goto tr657
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr657:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st448
	st448:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof448
		}
	st_case_448:
// line internal/php8/scanner.go:18470
		switch  lex.data[( lex.p)] {
		case 80:
			goto tr658
		case 92:
			goto st89
		case 96:
			goto tr325
		case 112:
			goto tr658
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr658:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st449
	st449:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof449
		}
	st_case_449:
// line internal/php8/scanner.go:18511
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr659
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr659
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr659:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st450
	st450:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof450
		}
	st_case_450:
// line internal/php8/scanner.go:18552
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr660
		case 92:
			goto st89
		case 96:
			goto tr325
		case 108:
			goto tr660
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr660:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st451
	st451:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof451
		}
	st_case_451:
// line internal/php8/scanner.go:18593
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr661
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr661
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr661:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st452
	st452:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof452
		}
	st_case_452:
// line internal/php8/scanner.go:18634
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr662
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr662
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr623:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st453
	st453:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof453
		}
	st_case_453:
// line internal/php8/scanner.go:18675
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr663
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr663
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr663:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st454
	st454:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof454
		}
	st_case_454:
// line internal/php8/scanner.go:18716
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr664
		case 92:
			goto st89
		case 96:
			goto tr325
		case 110:
			goto tr664
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr664:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st455
	st455:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof455
		}
	st_case_455:
// line internal/php8/scanner.go:18757
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr665
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr665
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr665:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st456
	st456:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof456
		}
	st_case_456:
// line internal/php8/scanner.go:18798
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr666
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr666:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st457
	st457:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof457
		}
	st_case_457:
// line internal/php8/scanner.go:18835
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr667
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr624:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st458
	st458:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof458
		}
	st_case_458:
// line internal/php8/scanner.go:18872
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr668
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr668
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr668:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st459
	st459:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof459
		}
	st_case_459:
// line internal/php8/scanner.go:18913
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr669
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr669
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr669:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st460
	st460:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof460
		}
	st_case_460:
// line internal/php8/scanner.go:18954
		switch  lex.data[( lex.p)] {
		case 72:
			goto tr670
		case 92:
			goto st89
		case 96:
			goto tr325
		case 104:
			goto tr670
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr670:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st461
	st461:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof461
		}
	st_case_461:
// line internal/php8/scanner.go:18995
		switch  lex.data[( lex.p)] {
		case 79:
			goto tr671
		case 92:
			goto st89
		case 96:
			goto tr325
		case 111:
			goto tr671
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr671:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st462
	st462:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof462
		}
	st_case_462:
// line internal/php8/scanner.go:19036
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr672
		case 92:
			goto st89
		case 96:
			goto tr325
		case 100:
			goto tr672
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr672:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st463
	st463:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof463
		}
	st_case_463:
// line internal/php8/scanner.go:19077
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr673
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr673:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st464
	st464:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof464
		}
	st_case_464:
// line internal/php8/scanner.go:19114
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr674
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr625:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st465
	st465:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof465
		}
	st_case_465:
// line internal/php8/scanner.go:19151
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr675
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr675
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr675:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st466
	st466:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof466
		}
	st_case_466:
// line internal/php8/scanner.go:19192
		switch  lex.data[( lex.p)] {
		case 77:
			goto tr676
		case 92:
			goto st89
		case 96:
			goto tr325
		case 109:
			goto tr676
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr676:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st467
	st467:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof467
		}
	st_case_467:
// line internal/php8/scanner.go:19233
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr677
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr677
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr677:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st468
	st468:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof468
		}
	st_case_468:
// line internal/php8/scanner.go:19274
		switch  lex.data[( lex.p)] {
		case 83:
			goto tr678
		case 92:
			goto st89
		case 96:
			goto tr325
		case 115:
			goto tr678
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr678:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st469
	st469:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof469
		}
	st_case_469:
// line internal/php8/scanner.go:19315
		switch  lex.data[( lex.p)] {
		case 80:
			goto tr679
		case 92:
			goto st89
		case 96:
			goto tr325
		case 112:
			goto tr679
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr679:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st470
	st470:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof470
		}
	st_case_470:
// line internal/php8/scanner.go:19356
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr680
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr680
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr680:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st471
	st471:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof471
		}
	st_case_471:
// line internal/php8/scanner.go:19397
		switch  lex.data[( lex.p)] {
		case 67:
			goto tr681
		case 92:
			goto st89
		case 96:
			goto tr325
		case 99:
			goto tr681
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr681:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st472
	st472:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof472
		}
	st_case_472:
// line internal/php8/scanner.go:19438
		switch  lex.data[( lex.p)] {
		case 69:
			goto tr682
		case 92:
			goto st89
		case 96:
			goto tr325
		case 101:
			goto tr682
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr682:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st473
	st473:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof473
		}
	st_case_473:
// line internal/php8/scanner.go:19479
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr683
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr683:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st474
	st474:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof474
		}
	st_case_474:
// line internal/php8/scanner.go:19516
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr684
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr626:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st475
	st475:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof475
		}
	st_case_475:
// line internal/php8/scanner.go:19553
		switch  lex.data[( lex.p)] {
		case 82:
			goto tr685
		case 92:
			goto st89
		case 96:
			goto tr325
		case 114:
			goto tr685
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr685:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st476
	st476:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof476
		}
	st_case_476:
// line internal/php8/scanner.go:19594
		switch  lex.data[( lex.p)] {
		case 65:
			goto tr686
		case 92:
			goto st89
		case 96:
			goto tr325
		case 97:
			goto tr686
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr686:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st477
	st477:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof477
		}
	st_case_477:
// line internal/php8/scanner.go:19635
		switch  lex.data[( lex.p)] {
		case 73:
			goto tr687
		case 92:
			goto st89
		case 96:
			goto tr325
		case 105:
			goto tr687
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr687:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st478
	st478:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof478
		}
	st_case_478:
// line internal/php8/scanner.go:19676
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr688
		case 92:
			goto st89
		case 96:
			goto tr325
		case 116:
			goto tr688
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr688:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st479
	st479:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof479
		}
	st_case_479:
// line internal/php8/scanner.go:19717
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr689
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
tr689:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:351
 lex.act = 140;
	goto st480
	st480:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof480
		}
	st_case_480:
// line internal/php8/scanner.go:19754
		switch  lex.data[( lex.p)] {
		case 92:
			goto st89
		case 95:
			goto tr690
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr325
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 96:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr325
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr325
			}
		default:
			goto tr325
		}
		goto tr226
	st481:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof481
		}
	st_case_481:
		switch  lex.data[( lex.p)] {
		case 61:
			goto tr691
		case 124:
			goto tr692
		}
		goto tr255
tr146:
// line internal/php8/scanner.rl:386
( lex.p) = ( lex.te) - 1
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st482
tr148:
// line internal/php8/scanner.rl:390
( lex.p) = ( lex.te) - 1
{lex.ungetCnt(1); {goto st127 }}
	goto st482
tr149:
// line internal/php8/scanner.rl:388
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_NULLSAFE_OBJECT_OPERATOR; {( lex.p)++;  lex.cs = 482; goto _out }}
	goto st482
tr693:
// line internal/php8/scanner.rl:390
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); {goto st127 }}
	goto st482
tr699:
// line internal/php8/scanner.rl:386
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st482
tr701:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:386
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st482
tr705:
// line internal/php8/scanner.rl:390
 lex.te = ( lex.p)
( lex.p)--
{lex.ungetCnt(1); {goto st127 }}
	goto st482
tr706:
// line internal/php8/scanner.rl:387
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_OBJECT_OPERATOR; {( lex.p)++;  lex.cs = 482; goto _out }}
	goto st482
tr708:
	 lex.cs = 482
// line internal/php8/scanner.rl:389
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_STRING;  lex.cs = 127; {( lex.p)++; goto _out }}
	goto _again
	st482:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof482
		}
	st_case_482:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:19854
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr147
		case 13:
			goto tr695
		case 32:
			goto tr694
		case 45:
			goto st486
		case 63:
			goto tr697
		case 96:
			goto tr693
		}
		switch {
		case  lex.data[( lex.p)] < 14:
			switch {
			case  lex.data[( lex.p)] > 8:
				if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
					goto tr694
				}
			default:
				goto tr693
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr693
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr693
			}
		default:
			goto tr693
		}
		goto st488
tr694:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st483
tr702:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

	goto st483
	st483:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof483
		}
	st_case_483:
// line internal/php8/scanner.go:19909
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr147
		case 13:
			goto tr700
		case 32:
			goto tr694
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr694
		}
		goto tr699
tr147:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st484
tr703:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st484
	st484:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof484
		}
	st_case_484:
// line internal/php8/scanner.go:19959
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr703
		case 13:
			goto tr704
		case 32:
			goto tr702
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr702
		}
		goto tr701
tr700:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st100
tr704:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st100
	st100:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof100
		}
	st_case_100:
// line internal/php8/scanner.go:20003
		if  lex.data[( lex.p)] == 10 {
			goto tr147
		}
		goto tr146
tr695:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st485
	st485:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof485
		}
	st_case_485:
// line internal/php8/scanner.go:20025
		if  lex.data[( lex.p)] == 10 {
			goto tr147
		}
		goto tr705
	st486:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof486
		}
	st_case_486:
		if  lex.data[( lex.p)] == 62 {
			goto tr706
		}
		goto tr705
tr697:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st487
	st487:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof487
		}
	st_case_487:
// line internal/php8/scanner.go:20049
		if  lex.data[( lex.p)] == 45 {
			goto st101
		}
		goto tr705
	st101:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if  lex.data[( lex.p)] == 62 {
			goto tr149
		}
		goto tr148
	st488:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof488
		}
	st_case_488:
		if  lex.data[( lex.p)] == 96 {
			goto tr708
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr708
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr708
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr708
			}
		default:
			goto tr708
		}
		goto st488
tr712:
	 lex.cs = 489
// line NONE:1
	switch  lex.act {
	case 0:
	{{goto st0 }}
	case 153:
	{( lex.p) = ( lex.te) - 1

                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                 lex.cs = 512;
                {( lex.p)++; goto _out }
            }
	}
	
	goto _again
tr713:
	 lex.cs = 489
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:394
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                 lex.cs = 512;
                {( lex.p)++; goto _out }
            }
	goto _again
tr717:
	 lex.cs = 489
// line internal/php8/scanner.rl:394
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                 lex.cs = 512;
                {( lex.p)++; goto _out }
            }
	goto _again
	st489:
// line NONE:1
 lex.ts = 0

// line NONE:1
 lex.act = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof489
		}
	st_case_489:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:20146
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p)  {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr710
		case 1037:
			goto tr711
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr709
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
tr709:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:394
 lex.act = 153;
	goto st490
tr714:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:394
 lex.act = 153;
	goto st490
	st490:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof490
		}
	st_case_490:
// line internal/php8/scanner.go:20220
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p)  {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr710
		case 1037:
			goto tr711
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr709
		}
		goto tr712
tr710:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st491
tr715:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st491
	st491:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof491
		}
	st_case_491:
// line internal/php8/scanner.go:20300
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p)  {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr715
		case 1037:
			goto tr716
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr714
		}
		goto tr713
tr711:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st492
tr716:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st492
	st492:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof492
		}
	st_case_492:
// line internal/php8/scanner.go:20380
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p)  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p)  {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p)  {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr710
		case 1037:
			goto tr711
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr709
		}
		goto tr717
tr150:
// line internal/php8/scanner.rl:403
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(493, 127); goto _out;}
	goto st493
tr725:
// line internal/php8/scanner.rl:405
 lex.te = ( lex.p)
( lex.p)--
{lex.ungetCnt(1); { lex.growCallStack(); { lex.stack[ lex.top] = 493;  lex.top++; goto st514 }}}
	goto st493
tr726:
// line internal/php8/scanner.rl:404
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(493, 530); goto _out;}
	goto st493
tr727:
	 lex.cs = 493
// line NONE:1
	switch  lex.act {
	case 154:
	{( lex.p) = ( lex.te) - 1
lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(493, 127); goto _out;}
	case 155:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(493, 530); goto _out;}
	case 157:
	{( lex.p) = ( lex.te) - 1

                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;

                if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
                     lex.cs = 512;
                }
                {( lex.p)++; goto _out }
            }
	}
	
	goto _again
tr728:
	 lex.cs = 493
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:406
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;

                if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
                     lex.cs = 512;
                }
                {( lex.p)++; goto _out }
            }
	goto _again
tr732:
	 lex.cs = 493
// line internal/php8/scanner.rl:406
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;

                if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
                     lex.cs = 512;
                }
                {( lex.p)++; goto _out }
            }
	goto _again
	st493:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof493
		}
	st_case_493:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:20512
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto st494
		case 1403:
			goto st102
		case 1546:
			goto tr721
		case 1549:
			goto tr722
		case 1572:
			goto st498
		case 1659:
			goto st499
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr720
		}
		goto st0
	st494:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof494
		}
	st_case_494:
		if  lex.data[( lex.p)] == 123 {
			goto tr726
		}
		goto tr725
	st102:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		if  lex.data[( lex.p)] == 36 {
			goto tr150
		}
		goto st0
tr720:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:406
 lex.act = 157;
	goto st495
tr729:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:406
 lex.act = 157;
	goto st495
tr733:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:404
 lex.act = 155;
	goto st495
tr734:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:403
 lex.act = 154;
	goto st495
	st495:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof495
		}
	st_case_495:
// line internal/php8/scanner.go:20622
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr721
		case 1549:
			goto tr722
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr720
		}
		goto tr727
tr721:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st496
tr730:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st496
	st496:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof496
		}
	st_case_496:
// line internal/php8/scanner.go:20702
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr730
		case 1549:
			goto tr731
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr729
		}
		goto tr728
tr722:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st497
tr731:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st497
	st497:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof497
		}
	st_case_497:
// line internal/php8/scanner.go:20782
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr721
		case 1549:
			goto tr722
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr720
		}
		goto tr732
	st498:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof498
		}
	st_case_498:
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1403:
			goto tr726
		case 1546:
			goto tr721
		case 1549:
			goto tr722
		case 1659:
			goto tr733
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr720
		}
		goto tr725
	st499:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof499
		}
	st_case_499:
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto tr150
		case 1546:
			goto tr721
		case 1549:
			goto tr722
		case 1572:
			goto tr734
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr720
		}
		goto tr732
tr152:
// line internal/php8/scanner.rl:420
 lex.te = ( lex.p)+1
{lex.ungetCnt(2); { lex.growCallStack(); { lex.stack[ lex.top] = 500;  lex.top++; goto st514 }}}
	goto st500
tr153:
// line internal/php8/scanner.rl:419
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(500, 530); goto _out;}
	goto st500
tr154:
// line internal/php8/scanner.rl:418
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(500, 127); goto _out;}
	goto st500
tr736:
	 lex.cs = 500
// line internal/php8/scanner.rl:421
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int('`'));  lex.cs = 127; {( lex.p)++; goto _out }}
	goto _again
tr744:
	 lex.cs = 500
// line NONE:1
	switch  lex.act {
	case 158:
	{( lex.p) = ( lex.te) - 1
lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(500, 127); goto _out;}
	case 159:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(500, 530); goto _out;}
	case 160:
	{( lex.p) = ( lex.te) - 1
lex.ungetCnt(2); { lex.growCallStack(); { lex.stack[ lex.top] = 500;  lex.top++; goto st514 }}}
	case 161:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.ID(int('`'));  lex.cs = 127; {( lex.p)++; goto _out }}
	case 162:
	{( lex.p) = ( lex.te) - 1

                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                {( lex.p)++; goto _out }
            }
	}
	
	goto _again
tr745:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:422
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                {( lex.p)++;  lex.cs = 500; goto _out }
            }
	goto st500
tr749:
// line internal/php8/scanner.rl:422
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                {( lex.p)++;  lex.cs = 500; goto _out }
            }
	goto st500
	st500:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof500
		}
	st_case_500:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:21025
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto st103
		case 1888:
			goto tr736
		case 1915:
			goto st104
		case 2058:
			goto tr739
		case 2061:
			goto tr740
		case 2084:
			goto st504
		case 2144:
			goto tr742
		case 2171:
			goto st505
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr738
		}
		goto st0
	st103:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch  lex.data[( lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr153
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto st0
			}
		case  lex.data[( lex.p)] > 94:
			if 124 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr152
	st104:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		if  lex.data[( lex.p)] == 36 {
			goto tr154
		}
		goto st0
tr738:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:422
 lex.act = 162;
	goto st501
tr742:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:421
 lex.act = 161;
	goto st501
tr746:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:422
 lex.act = 162;
	goto st501
tr750:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:420
 lex.act = 160;
	goto st501
tr751:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:419
 lex.act = 159;
	goto st501
tr752:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:418
 lex.act = 158;
	goto st501
	st501:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof501
		}
	st_case_501:
// line internal/php8/scanner.go:21168
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr739
		case 2061:
			goto tr740
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr738
		}
		goto tr744
tr739:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st502
tr747:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st502
	st502:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof502
		}
	st_case_502:
// line internal/php8/scanner.go:21248
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr747
		case 2061:
			goto tr748
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr746
		}
		goto tr745
tr740:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st503
tr748:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st503
	st503:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof503
		}
	st_case_503:
// line internal/php8/scanner.go:21328
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr739
		case 2061:
			goto tr740
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr738
		}
		goto tr749
	st504:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof504
		}
	st_case_504:
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1887:
			goto tr152
		case 1915:
			goto tr153
		case 2058:
			goto tr739
		case 2061:
			goto tr740
		case 2143:
			goto tr750
		case 2171:
			goto tr751
		}
		switch {
		case _widec < 2113:
			switch {
			case _widec < 1889:
				if 1857 <= _widec && _widec <= 1882 {
					goto tr152
				}
			case _widec > 1914:
				switch {
				case _widec > 2047:
					if 2048 <= _widec && _widec <= 2112 {
						goto tr738
					}
				case _widec >= 1920:
					goto tr152
				}
			default:
				goto tr152
			}
		case _widec > 2138:
			switch {
			case _widec < 2145:
				if 2139 <= _widec && _widec <= 2144 {
					goto tr738
				}
			case _widec > 2170:
				switch {
				case _widec > 2175:
					if 2176 <= _widec && _widec <= 2303 {
						goto tr750
					}
				case _widec >= 2172:
					goto tr738
				}
			default:
				goto tr750
			}
		default:
			goto tr750
		}
		goto tr749
	st505:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof505
		}
	st_case_505:
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('`') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto tr154
		case 2058:
			goto tr739
		case 2061:
			goto tr740
		case 2084:
			goto tr752
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr738
		}
		goto tr749
tr155:
// line internal/php8/scanner.rl:432
 lex.te = ( lex.p)+1
{lex.ungetCnt(2); { lex.growCallStack(); { lex.stack[ lex.top] = 506;  lex.top++; goto st514 }}}
	goto st506
tr156:
// line internal/php8/scanner.rl:431
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(506, 530); goto _out;}
	goto st506
tr157:
// line internal/php8/scanner.rl:430
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(506, 127); goto _out;}
	goto st506
tr753:
	 lex.cs = 506
// line internal/php8/scanner.rl:433
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int('"'));  lex.cs = 127; {( lex.p)++; goto _out }}
	goto _again
tr762:
	 lex.cs = 506
// line NONE:1
	switch  lex.act {
	case 163:
	{( lex.p) = ( lex.te) - 1
lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(506, 127); goto _out;}
	case 164:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(506, 530); goto _out;}
	case 165:
	{( lex.p) = ( lex.te) - 1
lex.ungetCnt(2); { lex.growCallStack(); { lex.stack[ lex.top] = 506;  lex.top++; goto st514 }}}
	case 166:
	{( lex.p) = ( lex.te) - 1
lex.setTokenPosition(tkn); tok = token.ID(int('"'));  lex.cs = 127; {( lex.p)++; goto _out }}
	case 167:
	{( lex.p) = ( lex.te) - 1

                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                {( lex.p)++; goto _out }
            }
	}
	
	goto _again
tr763:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:434
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                {( lex.p)++;  lex.cs = 506; goto _out }
            }
	goto st506
tr767:
// line internal/php8/scanner.rl:434
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                {( lex.p)++;  lex.cs = 506; goto _out }
            }
	goto st506
	st506:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof506
		}
	st_case_506:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:21612
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2338:
			goto tr753
		case 2340:
			goto st105
		case 2427:
			goto st106
		case 2570:
			goto tr757
		case 2573:
			goto tr758
		case 2594:
			goto tr759
		case 2596:
			goto st510
		case 2683:
			goto st511
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr756
		}
		goto st0
	st105:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		switch  lex.data[( lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr156
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto st0
			}
		case  lex.data[( lex.p)] > 94:
			if 124 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr155
	st106:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		if  lex.data[( lex.p)] == 36 {
			goto tr157
		}
		goto st0
tr756:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:434
 lex.act = 167;
	goto st507
tr759:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:433
 lex.act = 166;
	goto st507
tr764:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:434
 lex.act = 167;
	goto st507
tr768:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:432
 lex.act = 165;
	goto st507
tr769:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:431
 lex.act = 164;
	goto st507
tr770:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:430
 lex.act = 163;
	goto st507
	st507:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof507
		}
	st_case_507:
// line internal/php8/scanner.go:21755
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr757
		case 2573:
			goto tr758
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr756
		}
		goto tr762
tr757:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st508
tr765:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st508
	st508:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof508
		}
	st_case_508:
// line internal/php8/scanner.go:21835
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr765
		case 2573:
			goto tr766
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr764
		}
		goto tr763
tr758:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st509
tr766:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st509
	st509:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof509
		}
	st_case_509:
// line internal/php8/scanner.go:21915
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr757
		case 2573:
			goto tr758
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr756
		}
		goto tr767
	st510:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof510
		}
	st_case_510:
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2399:
			goto tr155
		case 2427:
			goto tr156
		case 2570:
			goto tr757
		case 2573:
			goto tr758
		case 2655:
			goto tr768
		case 2683:
			goto tr769
		}
		switch {
		case _widec < 2625:
			switch {
			case _widec < 2401:
				if 2369 <= _widec && _widec <= 2394 {
					goto tr155
				}
			case _widec > 2426:
				switch {
				case _widec > 2559:
					if 2560 <= _widec && _widec <= 2624 {
						goto tr756
					}
				case _widec >= 2432:
					goto tr155
				}
			default:
				goto tr155
			}
		case _widec > 2650:
			switch {
			case _widec < 2657:
				if 2651 <= _widec && _widec <= 2656 {
					goto tr756
				}
			case _widec > 2682:
				switch {
				case _widec > 2687:
					if 2688 <= _widec && _widec <= 2815 {
						goto tr768
					}
				case _widec >= 2684:
					goto tr756
				}
			default:
				goto tr768
			}
		default:
			goto tr768
		}
		goto tr767
	st511:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof511
		}
	st_case_511:
		_widec = int16( lex.data[( lex.p)])
		switch {
		case  lex.data[( lex.p)] < 11:
			switch {
			case  lex.data[( lex.p)] > 9:
				if 10 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 10 {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		case  lex.data[( lex.p)] > 12:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 14 <=  lex.data[( lex.p)] {
					_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
						_widec += 256
					}
				}
			case  lex.data[( lex.p)] >= 13:
				_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
				if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16( lex.data[( lex.p)]) - 0)
			if  lex.isNotStringEnd('"') && lex.isNotStringVar()  {
				_widec += 256
			}
		}
		switch _widec {
		case 2340:
			goto tr157
		case 2570:
			goto tr757
		case 2573:
			goto tr758
		case 2596:
			goto tr770
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr756
		}
		goto tr767
tr772:
	 lex.cs = 512
// line internal/php8/scanner.rl:442
 lex.te = ( lex.p)
( lex.p)--
{
                lex.setTokenPosition(tkn);
                tok = token.T_END_HEREDOC;
                 lex.cs = 127;
                {( lex.p)++; goto _out }
            }
	goto _again
	st512:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof512
		}
	st_case_512:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:22142
		if  lex.data[( lex.p)] == 96 {
			goto st0
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto st0
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto st513
	st513:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof513
		}
	st_case_513:
		if  lex.data[( lex.p)] == 96 {
			goto tr772
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr772
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr772
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr772
			}
		default:
			goto tr772
		}
		goto st513
tr158:
// line internal/php8/scanner.rl:462
( lex.p) = ( lex.te) - 1
{lex.ungetCnt(1); { lex.top--;  lex.cs =  lex.stack[ lex.top];goto _again }}
	goto st514
tr159:
// line internal/php8/scanner.rl:458
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_OBJECT_OPERATOR; {( lex.p)++;  lex.cs = 514; goto _out }}
	goto st514
tr161:
// line internal/php8/scanner.rl:459
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_NULLSAFE_OBJECT_OPERATOR; {( lex.p)++;  lex.cs = 514; goto _out }}
	goto st514
tr773:
// line internal/php8/scanner.rl:462
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); { lex.top--;  lex.cs =  lex.stack[ lex.top];goto _again }}
	goto st514
tr778:
// line internal/php8/scanner.rl:461
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int('[')); lex.call(514, 520); goto _out;}
	goto st514
tr779:
// line internal/php8/scanner.rl:462
 lex.te = ( lex.p)
( lex.p)--
{lex.ungetCnt(1); { lex.top--;  lex.cs =  lex.stack[ lex.top];goto _again }}
	goto st514
tr781:
// line internal/php8/scanner.rl:457
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_VARIABLE; {( lex.p)++;  lex.cs = 514; goto _out }}
	goto st514
tr784:
// line internal/php8/scanner.rl:460
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_STRING; {( lex.p)++;  lex.cs = 514; goto _out }}
	goto st514
	st514:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof514
		}
	st_case_514:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:22239
		switch  lex.data[( lex.p)] {
		case 36:
			goto st515
		case 45:
			goto tr775
		case 63:
			goto tr776
		case 91:
			goto tr778
		case 96:
			goto tr773
		}
		switch {
		case  lex.data[( lex.p)] < 92:
			if  lex.data[( lex.p)] <= 64 {
				goto tr773
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr773
			}
		default:
			goto tr773
		}
		goto st519
	st515:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof515
		}
	st_case_515:
		if  lex.data[( lex.p)] == 96 {
			goto tr779
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr779
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr779
			}
		default:
			goto tr779
		}
		goto st516
	st516:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof516
		}
	st_case_516:
		if  lex.data[( lex.p)] == 96 {
			goto tr781
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr781
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr781
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr781
			}
		default:
			goto tr781
		}
		goto st516
tr775:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st517
	st517:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof517
		}
	st_case_517:
// line internal/php8/scanner.go:22322
		if  lex.data[( lex.p)] == 62 {
			goto st107
		}
		goto tr779
	st107:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		if  lex.data[( lex.p)] == 96 {
			goto tr158
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr158
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr158
			}
		default:
			goto tr158
		}
		goto tr159
tr776:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st518
	st518:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof518
		}
	st_case_518:
// line internal/php8/scanner.go:22358
		if  lex.data[( lex.p)] == 45 {
			goto st108
		}
		goto tr779
	st108:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		if  lex.data[( lex.p)] == 62 {
			goto st109
		}
		goto tr158
	st109:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		if  lex.data[( lex.p)] == 96 {
			goto tr158
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr158
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr158
			}
		default:
			goto tr158
		}
		goto tr161
	st519:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof519
		}
	st_case_519:
		if  lex.data[( lex.p)] == 96 {
			goto tr784
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr784
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr784
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr784
			}
		default:
			goto tr784
		}
		goto st519
tr162:
// line internal/php8/scanner.rl:466
( lex.p) = ( lex.te) - 1
{lex.setTokenPosition(tkn); tok = token.T_NUM_STRING; {( lex.p)++;  lex.cs = 520; goto _out }}
	goto st520
tr785:
// line internal/php8/scanner.rl:472
 lex.te = ( lex.p)+1
{
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            }
	goto st520
tr786:
// line internal/php8/scanner.rl:469
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.T_ENCAPSED_AND_WHITESPACE; lex.ret(2); goto _out;}
	goto st520
tr789:
// line internal/php8/scanner.rl:470
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int(lex.data[lex.ts])); {( lex.p)++;  lex.cs = 520; goto _out }}
	goto st520
tr793:
// line internal/php8/scanner.rl:471
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int(']')); lex.ret(2); goto _out;}
	goto st520
tr794:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:469
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_ENCAPSED_AND_WHITESPACE; lex.ret(2); goto _out;}
	goto st520
tr795:
// line internal/php8/scanner.rl:472
 lex.te = ( lex.p)
( lex.p)--
{
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            }
	goto st520
tr796:
// line internal/php8/scanner.rl:470
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.ID(int(lex.data[lex.ts])); {( lex.p)++;  lex.cs = 520; goto _out }}
	goto st520
tr798:
// line internal/php8/scanner.rl:467
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_VARIABLE; {( lex.p)++;  lex.cs = 520; goto _out }}
	goto st520
tr799:
// line internal/php8/scanner.rl:466
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_NUM_STRING; {( lex.p)++;  lex.cs = 520; goto _out }}
	goto st520
tr803:
// line internal/php8/scanner.rl:468
 lex.te = ( lex.p)
( lex.p)--
{lex.setTokenPosition(tkn); tok = token.T_STRING; {( lex.p)++;  lex.cs = 520; goto _out }}
	goto st520
	st520:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof520
		}
	st_case_520:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:22499
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr787
		case 13:
			goto tr788
		case 32:
			goto tr786
		case 33:
			goto tr789
		case 35:
			goto tr786
		case 36:
			goto st523
		case 39:
			goto tr786
		case 48:
			goto tr791
		case 92:
			goto tr786
		case 93:
			goto tr793
		case 96:
			goto tr785
		case 124:
			goto tr789
		case 126:
			goto tr789
		}
		switch {
		case  lex.data[( lex.p)] < 37:
			switch {
			case  lex.data[( lex.p)] < 9:
				if  lex.data[( lex.p)] <= 8 {
					goto tr785
				}
			case  lex.data[( lex.p)] > 12:
				if 14 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 34 {
					goto tr785
				}
			default:
				goto tr786
			}
		case  lex.data[( lex.p)] > 47:
			switch {
			case  lex.data[( lex.p)] < 58:
				if 49 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
					goto tr163
				}
			case  lex.data[( lex.p)] > 64:
				switch {
				case  lex.data[( lex.p)] > 94:
					if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
						goto tr785
					}
				case  lex.data[( lex.p)] >= 91:
					goto tr789
				}
			default:
				goto tr789
			}
		default:
			goto tr789
		}
		goto st529
tr787:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st521
	st521:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof521
		}
	st_case_521:
// line internal/php8/scanner.go:22581
		goto tr794
tr788:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st522
	st522:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof522
		}
	st_case_522:
// line internal/php8/scanner.go:22600
		if  lex.data[( lex.p)] == 10 {
			goto tr787
		}
		goto tr795
	st523:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof523
		}
	st_case_523:
		if  lex.data[( lex.p)] == 96 {
			goto tr796
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr796
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr796
			}
		default:
			goto tr796
		}
		goto st524
	st524:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof524
		}
	st_case_524:
		if  lex.data[( lex.p)] == 96 {
			goto tr798
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr798
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr798
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr798
			}
		default:
			goto tr798
		}
		goto st524
tr791:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st525
	st525:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof525
		}
	st_case_525:
// line internal/php8/scanner.go:22662
		switch  lex.data[( lex.p)] {
		case 95:
			goto st110
		case 98:
			goto st111
		case 120:
			goto st112
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr163
		}
		goto tr799
tr163:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st526
	st526:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof526
		}
	st_case_526:
// line internal/php8/scanner.go:22685
		if  lex.data[( lex.p)] == 95 {
			goto st110
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr163
		}
		goto tr799
	st110:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr163
		}
		goto tr162
	st111:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof111
		}
	st_case_111:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 49 {
			goto tr164
		}
		goto tr162
tr164:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st527
	st527:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof527
		}
	st_case_527:
// line internal/php8/scanner.go:22721
		if  lex.data[( lex.p)] == 95 {
			goto st111
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 49 {
			goto tr164
		}
		goto tr799
	st112:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr165
			}
		case  lex.data[( lex.p)] > 70:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 102 {
				goto tr165
			}
		default:
			goto tr165
		}
		goto tr162
tr165:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st528
	st528:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof528
		}
	st_case_528:
// line internal/php8/scanner.go:22757
		if  lex.data[( lex.p)] == 95 {
			goto st112
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr165
			}
		case  lex.data[( lex.p)] > 70:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 102 {
				goto tr165
			}
		default:
			goto tr165
		}
		goto tr799
	st529:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof529
		}
	st_case_529:
		if  lex.data[( lex.p)] == 96 {
			goto tr803
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr803
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr803
				}
			case  lex.data[( lex.p)] >= 91:
				goto tr803
			}
		default:
			goto tr803
		}
		goto st529
tr166:
	 lex.cs = 530
// line internal/php8/scanner.rl:480
( lex.p) = ( lex.te) - 1
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
tr168:
	 lex.cs = 530
// line internal/php8/scanner.rl:479
 lex.te = ( lex.p)+1
{lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_STRING_VARNAME;  lex.cs = 127; {( lex.p)++; goto _out }}
	goto _again
tr804:
	 lex.cs = 530
// line internal/php8/scanner.rl:480
 lex.te = ( lex.p)+1
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
tr806:
	 lex.cs = 530
// line internal/php8/scanner.rl:480
 lex.te = ( lex.p)
( lex.p)--
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
	st530:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof530
		}
	st_case_530:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:22836
		if  lex.data[( lex.p)] == 96 {
			goto tr804
		}
		switch {
		case  lex.data[( lex.p)] < 91:
			if  lex.data[( lex.p)] <= 64 {
				goto tr804
			}
		case  lex.data[( lex.p)] > 94:
			if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
				goto tr804
			}
		default:
			goto tr804
		}
		goto tr805
tr805:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st531
	st531:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof531
		}
	st_case_531:
// line internal/php8/scanner.go:22863
		switch  lex.data[( lex.p)] {
		case 91:
			goto tr168
		case 96:
			goto tr806
		case 125:
			goto tr168
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr806
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr806
				}
			case  lex.data[( lex.p)] >= 92:
				goto tr806
			}
		default:
			goto tr806
		}
		goto st113
	st113:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		switch  lex.data[( lex.p)] {
		case 91:
			goto tr168
		case 96:
			goto tr166
		case 125:
			goto tr168
		}
		switch {
		case  lex.data[( lex.p)] < 58:
			if  lex.data[( lex.p)] <= 47 {
				goto tr166
			}
		case  lex.data[( lex.p)] > 64:
			switch {
			case  lex.data[( lex.p)] > 94:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 127 {
					goto tr166
				}
			case  lex.data[( lex.p)] >= 92:
				goto tr166
			}
		default:
			goto tr166
		}
		goto st113
tr169:
// line internal/php8/scanner.rl:484
( lex.p) = ( lex.te) - 1
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st532
tr807:
	 lex.cs = 532
// line internal/php8/scanner.rl:486
 lex.te = ( lex.p)+1
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
tr810:
	 lex.cs = 532
// line internal/php8/scanner.rl:485
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int('('));  lex.cs = 536; {( lex.p)++; goto _out }}
	goto _again
tr811:
// line internal/php8/scanner.rl:484
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st532
tr813:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:484
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st532
tr817:
	 lex.cs = 532
// line internal/php8/scanner.rl:486
 lex.te = ( lex.p)
( lex.p)--
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
	st532:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof532
		}
	st_case_532:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:22970
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr170
		case 13:
			goto tr809
		case 32:
			goto tr808
		case 40:
			goto tr810
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr808
		}
		goto tr807
tr808:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st533
tr814:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

	goto st533
	st533:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof533
		}
	st_case_533:
// line internal/php8/scanner.go:23002
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr170
		case 13:
			goto tr812
		case 32:
			goto tr808
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr808
		}
		goto tr811
tr170:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st534
tr815:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st534
	st534:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof534
		}
	st_case_534:
// line internal/php8/scanner.go:23052
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr815
		case 13:
			goto tr816
		case 32:
			goto tr814
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr814
		}
		goto tr813
tr812:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st114
tr816:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st114
	st114:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof114
		}
	st_case_114:
// line internal/php8/scanner.go:23096
		if  lex.data[( lex.p)] == 10 {
			goto tr170
		}
		goto tr169
tr809:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st535
	st535:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof535
		}
	st_case_535:
// line internal/php8/scanner.go:23118
		if  lex.data[( lex.p)] == 10 {
			goto tr170
		}
		goto tr817
tr171:
// line internal/php8/scanner.rl:490
( lex.p) = ( lex.te) - 1
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st536
tr818:
	 lex.cs = 536
// line internal/php8/scanner.rl:492
 lex.te = ( lex.p)+1
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
tr821:
	 lex.cs = 536
// line internal/php8/scanner.rl:491
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int(')'));  lex.cs = 540; {( lex.p)++; goto _out }}
	goto _again
tr822:
// line internal/php8/scanner.rl:490
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st536
tr824:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:490
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st536
tr828:
	 lex.cs = 536
// line internal/php8/scanner.rl:492
 lex.te = ( lex.p)
( lex.p)--
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
	st536:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof536
		}
	st_case_536:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:23172
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr172
		case 13:
			goto tr820
		case 32:
			goto tr819
		case 41:
			goto tr821
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr819
		}
		goto tr818
tr819:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st537
tr825:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

	goto st537
	st537:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof537
		}
	st_case_537:
// line internal/php8/scanner.go:23204
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr172
		case 13:
			goto tr823
		case 32:
			goto tr819
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr819
		}
		goto tr822
tr172:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st538
tr826:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st538
	st538:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof538
		}
	st_case_538:
// line internal/php8/scanner.go:23254
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr826
		case 13:
			goto tr827
		case 32:
			goto tr825
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr825
		}
		goto tr824
tr823:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st115
tr827:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st115
	st115:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof115
		}
	st_case_115:
// line internal/php8/scanner.go:23298
		if  lex.data[( lex.p)] == 10 {
			goto tr172
		}
		goto tr171
tr820:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st539
	st539:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof539
		}
	st_case_539:
// line internal/php8/scanner.go:23320
		if  lex.data[( lex.p)] == 10 {
			goto tr172
		}
		goto tr828
tr173:
// line internal/php8/scanner.rl:496
( lex.p) = ( lex.te) - 1
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st540
tr829:
	 lex.cs = 540
// line internal/php8/scanner.rl:498
 lex.te = ( lex.p)+1
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
tr832:
	 lex.cs = 540
// line internal/php8/scanner.rl:497
 lex.te = ( lex.p)+1
{lex.setTokenPosition(tkn); tok = token.ID(int(';'));  lex.cs = 544; {( lex.p)++; goto _out }}
	goto _again
tr833:
// line internal/php8/scanner.rl:496
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st540
tr835:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:496
 lex.te = ( lex.p)
( lex.p)--
{lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)}
	goto st540
tr839:
	 lex.cs = 540
// line internal/php8/scanner.rl:498
 lex.te = ( lex.p)
( lex.p)--
{lex.ungetCnt(1);  lex.cs = 127;}
	goto _again
	st540:
// line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof540
		}
	st_case_540:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:23374
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr174
		case 13:
			goto tr831
		case 32:
			goto tr830
		case 59:
			goto tr832
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr830
		}
		goto tr829
tr830:
// line NONE:1
 lex.te = ( lex.p)+1

	goto st541
tr836:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

	goto st541
	st541:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof541
		}
	st_case_541:
// line internal/php8/scanner.go:23406
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr174
		case 13:
			goto tr834
		case 32:
			goto tr830
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr830
		}
		goto tr833
tr174:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st542
tr837:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st542
	st542:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof542
		}
	st_case_542:
// line internal/php8/scanner.go:23456
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr837
		case 13:
			goto tr838
		case 32:
			goto tr836
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 12 {
			goto tr836
		}
		goto tr835
tr834:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st116
tr838:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st116
	st116:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof116
		}
	st_case_116:
// line internal/php8/scanner.go:23500
		if  lex.data[( lex.p)] == 10 {
			goto tr174
		}
		goto tr173
tr831:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st543
	st543:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof543
		}
	st_case_543:
// line internal/php8/scanner.go:23522
		if  lex.data[( lex.p)] == 10 {
			goto tr174
		}
		goto tr839
tr843:
// line NONE:1
	switch  lex.act {
	case 0:
	{{goto st0 }}
	case 194:
	{( lex.p) = ( lex.te) - 1
 lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te); }
	}
	
	goto st544
tr844:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:502
 lex.te = ( lex.p)
( lex.p)--
{ lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te); }
	goto st544
tr848:
// line internal/php8/scanner.rl:502
 lex.te = ( lex.p)
( lex.p)--
{ lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te); }
	goto st544
	st544:
// line NONE:1
 lex.ts = 0

// line NONE:1
 lex.act = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof544
		}
	st_case_544:
// line NONE:1
 lex.ts = ( lex.p)

// line internal/php8/scanner.go:23566
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr841
		case 13:
			goto tr842
		}
		goto tr840
tr840:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:502
 lex.act = 194;
	goto st545
tr845:
// line NONE:1
 lex.te = ( lex.p)+1

// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:502
 lex.act = 194;
	goto st545
	st545:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof545
		}
	st_case_545:
// line internal/php8/scanner.go:23595
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr841
		case 13:
			goto tr842
		}
		goto tr840
tr841:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st546
tr846:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st546
	st546:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof546
		}
	st_case_546:
// line internal/php8/scanner.go:23634
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr846
		case 13:
			goto tr847
		}
		goto tr845
tr842:
// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st547
tr847:
// line internal/php8/scanner.rl:54

// line internal/php8/scanner.rl:38

            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        
	goto st547
	st547:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof547
		}
	st_case_547:
// line internal/php8/scanner.go:23673
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr841
		case 13:
			goto tr842
		}
		goto tr840
	st_out:
	_test_eof117:  lex.cs = 117; goto _test_eof
	_test_eof118:  lex.cs = 118; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof119:  lex.cs = 119; goto _test_eof
	_test_eof120:  lex.cs = 120; goto _test_eof
	_test_eof121:  lex.cs = 121; goto _test_eof
	_test_eof122:  lex.cs = 122; goto _test_eof
	_test_eof123:  lex.cs = 123; goto _test_eof
	_test_eof124:  lex.cs = 124; goto _test_eof
	_test_eof125:  lex.cs = 125; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof126:  lex.cs = 126; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof127:  lex.cs = 127; goto _test_eof
	_test_eof128:  lex.cs = 128; goto _test_eof
	_test_eof129:  lex.cs = 129; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof130:  lex.cs = 130; goto _test_eof
	_test_eof131:  lex.cs = 131; goto _test_eof
	_test_eof132:  lex.cs = 132; goto _test_eof
	_test_eof133:  lex.cs = 133; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof134:  lex.cs = 134; goto _test_eof
	_test_eof135:  lex.cs = 135; goto _test_eof
	_test_eof136:  lex.cs = 136; goto _test_eof
	_test_eof137:  lex.cs = 137; goto _test_eof
	_test_eof138:  lex.cs = 138; goto _test_eof
	_test_eof139:  lex.cs = 139; goto _test_eof
	_test_eof140:  lex.cs = 140; goto _test_eof
	_test_eof141:  lex.cs = 141; goto _test_eof
	_test_eof142:  lex.cs = 142; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof143:  lex.cs = 143; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof
	_test_eof27:  lex.cs = 27; goto _test_eof
	_test_eof28:  lex.cs = 28; goto _test_eof
	_test_eof29:  lex.cs = 29; goto _test_eof
	_test_eof30:  lex.cs = 30; goto _test_eof
	_test_eof31:  lex.cs = 31; goto _test_eof
	_test_eof32:  lex.cs = 32; goto _test_eof
	_test_eof33:  lex.cs = 33; goto _test_eof
	_test_eof34:  lex.cs = 34; goto _test_eof
	_test_eof35:  lex.cs = 35; goto _test_eof
	_test_eof36:  lex.cs = 36; goto _test_eof
	_test_eof37:  lex.cs = 37; goto _test_eof
	_test_eof38:  lex.cs = 38; goto _test_eof
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof42:  lex.cs = 42; goto _test_eof
	_test_eof43:  lex.cs = 43; goto _test_eof
	_test_eof44:  lex.cs = 44; goto _test_eof
	_test_eof45:  lex.cs = 45; goto _test_eof
	_test_eof46:  lex.cs = 46; goto _test_eof
	_test_eof47:  lex.cs = 47; goto _test_eof
	_test_eof48:  lex.cs = 48; goto _test_eof
	_test_eof49:  lex.cs = 49; goto _test_eof
	_test_eof50:  lex.cs = 50; goto _test_eof
	_test_eof51:  lex.cs = 51; goto _test_eof
	_test_eof52:  lex.cs = 52; goto _test_eof
	_test_eof53:  lex.cs = 53; goto _test_eof
	_test_eof54:  lex.cs = 54; goto _test_eof
	_test_eof55:  lex.cs = 55; goto _test_eof
	_test_eof56:  lex.cs = 56; goto _test_eof
	_test_eof57:  lex.cs = 57; goto _test_eof
	_test_eof58:  lex.cs = 58; goto _test_eof
	_test_eof59:  lex.cs = 59; goto _test_eof
	_test_eof60:  lex.cs = 60; goto _test_eof
	_test_eof61:  lex.cs = 61; goto _test_eof
	_test_eof62:  lex.cs = 62; goto _test_eof
	_test_eof63:  lex.cs = 63; goto _test_eof
	_test_eof144:  lex.cs = 144; goto _test_eof
	_test_eof145:  lex.cs = 145; goto _test_eof
	_test_eof146:  lex.cs = 146; goto _test_eof
	_test_eof147:  lex.cs = 147; goto _test_eof
	_test_eof148:  lex.cs = 148; goto _test_eof
	_test_eof64:  lex.cs = 64; goto _test_eof
	_test_eof149:  lex.cs = 149; goto _test_eof
	_test_eof65:  lex.cs = 65; goto _test_eof
	_test_eof66:  lex.cs = 66; goto _test_eof
	_test_eof150:  lex.cs = 150; goto _test_eof
	_test_eof67:  lex.cs = 67; goto _test_eof
	_test_eof151:  lex.cs = 151; goto _test_eof
	_test_eof68:  lex.cs = 68; goto _test_eof
	_test_eof69:  lex.cs = 69; goto _test_eof
	_test_eof70:  lex.cs = 70; goto _test_eof
	_test_eof152:  lex.cs = 152; goto _test_eof
	_test_eof153:  lex.cs = 153; goto _test_eof
	_test_eof154:  lex.cs = 154; goto _test_eof
	_test_eof71:  lex.cs = 71; goto _test_eof
	_test_eof72:  lex.cs = 72; goto _test_eof
	_test_eof155:  lex.cs = 155; goto _test_eof
	_test_eof73:  lex.cs = 73; goto _test_eof
	_test_eof156:  lex.cs = 156; goto _test_eof
	_test_eof157:  lex.cs = 157; goto _test_eof
	_test_eof158:  lex.cs = 158; goto _test_eof
	_test_eof74:  lex.cs = 74; goto _test_eof
	_test_eof75:  lex.cs = 75; goto _test_eof
	_test_eof76:  lex.cs = 76; goto _test_eof
	_test_eof77:  lex.cs = 77; goto _test_eof
	_test_eof159:  lex.cs = 159; goto _test_eof
	_test_eof160:  lex.cs = 160; goto _test_eof
	_test_eof78:  lex.cs = 78; goto _test_eof
	_test_eof161:  lex.cs = 161; goto _test_eof
	_test_eof162:  lex.cs = 162; goto _test_eof
	_test_eof79:  lex.cs = 79; goto _test_eof
	_test_eof80:  lex.cs = 80; goto _test_eof
	_test_eof81:  lex.cs = 81; goto _test_eof
	_test_eof82:  lex.cs = 82; goto _test_eof
	_test_eof163:  lex.cs = 163; goto _test_eof
	_test_eof83:  lex.cs = 83; goto _test_eof
	_test_eof84:  lex.cs = 84; goto _test_eof
	_test_eof85:  lex.cs = 85; goto _test_eof
	_test_eof86:  lex.cs = 86; goto _test_eof
	_test_eof164:  lex.cs = 164; goto _test_eof
	_test_eof165:  lex.cs = 165; goto _test_eof
	_test_eof166:  lex.cs = 166; goto _test_eof
	_test_eof167:  lex.cs = 167; goto _test_eof
	_test_eof168:  lex.cs = 168; goto _test_eof
	_test_eof169:  lex.cs = 169; goto _test_eof
	_test_eof87:  lex.cs = 87; goto _test_eof
	_test_eof170:  lex.cs = 170; goto _test_eof
	_test_eof171:  lex.cs = 171; goto _test_eof
	_test_eof88:  lex.cs = 88; goto _test_eof
	_test_eof172:  lex.cs = 172; goto _test_eof
	_test_eof173:  lex.cs = 173; goto _test_eof
	_test_eof174:  lex.cs = 174; goto _test_eof
	_test_eof89:  lex.cs = 89; goto _test_eof
	_test_eof175:  lex.cs = 175; goto _test_eof
	_test_eof176:  lex.cs = 176; goto _test_eof
	_test_eof177:  lex.cs = 177; goto _test_eof
	_test_eof178:  lex.cs = 178; goto _test_eof
	_test_eof179:  lex.cs = 179; goto _test_eof
	_test_eof180:  lex.cs = 180; goto _test_eof
	_test_eof181:  lex.cs = 181; goto _test_eof
	_test_eof182:  lex.cs = 182; goto _test_eof
	_test_eof183:  lex.cs = 183; goto _test_eof
	_test_eof184:  lex.cs = 184; goto _test_eof
	_test_eof185:  lex.cs = 185; goto _test_eof
	_test_eof186:  lex.cs = 186; goto _test_eof
	_test_eof90:  lex.cs = 90; goto _test_eof
	_test_eof91:  lex.cs = 91; goto _test_eof
	_test_eof187:  lex.cs = 187; goto _test_eof
	_test_eof188:  lex.cs = 188; goto _test_eof
	_test_eof189:  lex.cs = 189; goto _test_eof
	_test_eof190:  lex.cs = 190; goto _test_eof
	_test_eof191:  lex.cs = 191; goto _test_eof
	_test_eof192:  lex.cs = 192; goto _test_eof
	_test_eof193:  lex.cs = 193; goto _test_eof
	_test_eof194:  lex.cs = 194; goto _test_eof
	_test_eof195:  lex.cs = 195; goto _test_eof
	_test_eof196:  lex.cs = 196; goto _test_eof
	_test_eof197:  lex.cs = 197; goto _test_eof
	_test_eof198:  lex.cs = 198; goto _test_eof
	_test_eof199:  lex.cs = 199; goto _test_eof
	_test_eof200:  lex.cs = 200; goto _test_eof
	_test_eof201:  lex.cs = 201; goto _test_eof
	_test_eof202:  lex.cs = 202; goto _test_eof
	_test_eof203:  lex.cs = 203; goto _test_eof
	_test_eof204:  lex.cs = 204; goto _test_eof
	_test_eof205:  lex.cs = 205; goto _test_eof
	_test_eof206:  lex.cs = 206; goto _test_eof
	_test_eof207:  lex.cs = 207; goto _test_eof
	_test_eof208:  lex.cs = 208; goto _test_eof
	_test_eof209:  lex.cs = 209; goto _test_eof
	_test_eof210:  lex.cs = 210; goto _test_eof
	_test_eof211:  lex.cs = 211; goto _test_eof
	_test_eof212:  lex.cs = 212; goto _test_eof
	_test_eof213:  lex.cs = 213; goto _test_eof
	_test_eof214:  lex.cs = 214; goto _test_eof
	_test_eof215:  lex.cs = 215; goto _test_eof
	_test_eof216:  lex.cs = 216; goto _test_eof
	_test_eof217:  lex.cs = 217; goto _test_eof
	_test_eof218:  lex.cs = 218; goto _test_eof
	_test_eof219:  lex.cs = 219; goto _test_eof
	_test_eof220:  lex.cs = 220; goto _test_eof
	_test_eof221:  lex.cs = 221; goto _test_eof
	_test_eof222:  lex.cs = 222; goto _test_eof
	_test_eof223:  lex.cs = 223; goto _test_eof
	_test_eof224:  lex.cs = 224; goto _test_eof
	_test_eof225:  lex.cs = 225; goto _test_eof
	_test_eof226:  lex.cs = 226; goto _test_eof
	_test_eof227:  lex.cs = 227; goto _test_eof
	_test_eof228:  lex.cs = 228; goto _test_eof
	_test_eof229:  lex.cs = 229; goto _test_eof
	_test_eof230:  lex.cs = 230; goto _test_eof
	_test_eof231:  lex.cs = 231; goto _test_eof
	_test_eof232:  lex.cs = 232; goto _test_eof
	_test_eof233:  lex.cs = 233; goto _test_eof
	_test_eof234:  lex.cs = 234; goto _test_eof
	_test_eof235:  lex.cs = 235; goto _test_eof
	_test_eof236:  lex.cs = 236; goto _test_eof
	_test_eof237:  lex.cs = 237; goto _test_eof
	_test_eof238:  lex.cs = 238; goto _test_eof
	_test_eof239:  lex.cs = 239; goto _test_eof
	_test_eof240:  lex.cs = 240; goto _test_eof
	_test_eof241:  lex.cs = 241; goto _test_eof
	_test_eof242:  lex.cs = 242; goto _test_eof
	_test_eof243:  lex.cs = 243; goto _test_eof
	_test_eof244:  lex.cs = 244; goto _test_eof
	_test_eof245:  lex.cs = 245; goto _test_eof
	_test_eof246:  lex.cs = 246; goto _test_eof
	_test_eof247:  lex.cs = 247; goto _test_eof
	_test_eof248:  lex.cs = 248; goto _test_eof
	_test_eof249:  lex.cs = 249; goto _test_eof
	_test_eof250:  lex.cs = 250; goto _test_eof
	_test_eof251:  lex.cs = 251; goto _test_eof
	_test_eof252:  lex.cs = 252; goto _test_eof
	_test_eof253:  lex.cs = 253; goto _test_eof
	_test_eof254:  lex.cs = 254; goto _test_eof
	_test_eof255:  lex.cs = 255; goto _test_eof
	_test_eof256:  lex.cs = 256; goto _test_eof
	_test_eof257:  lex.cs = 257; goto _test_eof
	_test_eof258:  lex.cs = 258; goto _test_eof
	_test_eof259:  lex.cs = 259; goto _test_eof
	_test_eof260:  lex.cs = 260; goto _test_eof
	_test_eof261:  lex.cs = 261; goto _test_eof
	_test_eof262:  lex.cs = 262; goto _test_eof
	_test_eof263:  lex.cs = 263; goto _test_eof
	_test_eof264:  lex.cs = 264; goto _test_eof
	_test_eof265:  lex.cs = 265; goto _test_eof
	_test_eof266:  lex.cs = 266; goto _test_eof
	_test_eof267:  lex.cs = 267; goto _test_eof
	_test_eof268:  lex.cs = 268; goto _test_eof
	_test_eof269:  lex.cs = 269; goto _test_eof
	_test_eof270:  lex.cs = 270; goto _test_eof
	_test_eof271:  lex.cs = 271; goto _test_eof
	_test_eof272:  lex.cs = 272; goto _test_eof
	_test_eof273:  lex.cs = 273; goto _test_eof
	_test_eof274:  lex.cs = 274; goto _test_eof
	_test_eof275:  lex.cs = 275; goto _test_eof
	_test_eof276:  lex.cs = 276; goto _test_eof
	_test_eof277:  lex.cs = 277; goto _test_eof
	_test_eof278:  lex.cs = 278; goto _test_eof
	_test_eof279:  lex.cs = 279; goto _test_eof
	_test_eof280:  lex.cs = 280; goto _test_eof
	_test_eof281:  lex.cs = 281; goto _test_eof
	_test_eof282:  lex.cs = 282; goto _test_eof
	_test_eof283:  lex.cs = 283; goto _test_eof
	_test_eof284:  lex.cs = 284; goto _test_eof
	_test_eof285:  lex.cs = 285; goto _test_eof
	_test_eof286:  lex.cs = 286; goto _test_eof
	_test_eof287:  lex.cs = 287; goto _test_eof
	_test_eof288:  lex.cs = 288; goto _test_eof
	_test_eof289:  lex.cs = 289; goto _test_eof
	_test_eof290:  lex.cs = 290; goto _test_eof
	_test_eof291:  lex.cs = 291; goto _test_eof
	_test_eof292:  lex.cs = 292; goto _test_eof
	_test_eof293:  lex.cs = 293; goto _test_eof
	_test_eof294:  lex.cs = 294; goto _test_eof
	_test_eof295:  lex.cs = 295; goto _test_eof
	_test_eof296:  lex.cs = 296; goto _test_eof
	_test_eof297:  lex.cs = 297; goto _test_eof
	_test_eof298:  lex.cs = 298; goto _test_eof
	_test_eof299:  lex.cs = 299; goto _test_eof
	_test_eof300:  lex.cs = 300; goto _test_eof
	_test_eof301:  lex.cs = 301; goto _test_eof
	_test_eof302:  lex.cs = 302; goto _test_eof
	_test_eof303:  lex.cs = 303; goto _test_eof
	_test_eof304:  lex.cs = 304; goto _test_eof
	_test_eof305:  lex.cs = 305; goto _test_eof
	_test_eof306:  lex.cs = 306; goto _test_eof
	_test_eof307:  lex.cs = 307; goto _test_eof
	_test_eof308:  lex.cs = 308; goto _test_eof
	_test_eof309:  lex.cs = 309; goto _test_eof
	_test_eof310:  lex.cs = 310; goto _test_eof
	_test_eof311:  lex.cs = 311; goto _test_eof
	_test_eof312:  lex.cs = 312; goto _test_eof
	_test_eof313:  lex.cs = 313; goto _test_eof
	_test_eof314:  lex.cs = 314; goto _test_eof
	_test_eof315:  lex.cs = 315; goto _test_eof
	_test_eof316:  lex.cs = 316; goto _test_eof
	_test_eof317:  lex.cs = 317; goto _test_eof
	_test_eof318:  lex.cs = 318; goto _test_eof
	_test_eof319:  lex.cs = 319; goto _test_eof
	_test_eof320:  lex.cs = 320; goto _test_eof
	_test_eof321:  lex.cs = 321; goto _test_eof
	_test_eof322:  lex.cs = 322; goto _test_eof
	_test_eof323:  lex.cs = 323; goto _test_eof
	_test_eof324:  lex.cs = 324; goto _test_eof
	_test_eof325:  lex.cs = 325; goto _test_eof
	_test_eof326:  lex.cs = 326; goto _test_eof
	_test_eof327:  lex.cs = 327; goto _test_eof
	_test_eof328:  lex.cs = 328; goto _test_eof
	_test_eof329:  lex.cs = 329; goto _test_eof
	_test_eof330:  lex.cs = 330; goto _test_eof
	_test_eof331:  lex.cs = 331; goto _test_eof
	_test_eof332:  lex.cs = 332; goto _test_eof
	_test_eof333:  lex.cs = 333; goto _test_eof
	_test_eof334:  lex.cs = 334; goto _test_eof
	_test_eof335:  lex.cs = 335; goto _test_eof
	_test_eof336:  lex.cs = 336; goto _test_eof
	_test_eof337:  lex.cs = 337; goto _test_eof
	_test_eof338:  lex.cs = 338; goto _test_eof
	_test_eof339:  lex.cs = 339; goto _test_eof
	_test_eof340:  lex.cs = 340; goto _test_eof
	_test_eof341:  lex.cs = 341; goto _test_eof
	_test_eof342:  lex.cs = 342; goto _test_eof
	_test_eof343:  lex.cs = 343; goto _test_eof
	_test_eof344:  lex.cs = 344; goto _test_eof
	_test_eof92:  lex.cs = 92; goto _test_eof
	_test_eof345:  lex.cs = 345; goto _test_eof
	_test_eof346:  lex.cs = 346; goto _test_eof
	_test_eof347:  lex.cs = 347; goto _test_eof
	_test_eof348:  lex.cs = 348; goto _test_eof
	_test_eof349:  lex.cs = 349; goto _test_eof
	_test_eof350:  lex.cs = 350; goto _test_eof
	_test_eof351:  lex.cs = 351; goto _test_eof
	_test_eof352:  lex.cs = 352; goto _test_eof
	_test_eof353:  lex.cs = 353; goto _test_eof
	_test_eof354:  lex.cs = 354; goto _test_eof
	_test_eof355:  lex.cs = 355; goto _test_eof
	_test_eof356:  lex.cs = 356; goto _test_eof
	_test_eof357:  lex.cs = 357; goto _test_eof
	_test_eof358:  lex.cs = 358; goto _test_eof
	_test_eof359:  lex.cs = 359; goto _test_eof
	_test_eof360:  lex.cs = 360; goto _test_eof
	_test_eof361:  lex.cs = 361; goto _test_eof
	_test_eof362:  lex.cs = 362; goto _test_eof
	_test_eof363:  lex.cs = 363; goto _test_eof
	_test_eof364:  lex.cs = 364; goto _test_eof
	_test_eof365:  lex.cs = 365; goto _test_eof
	_test_eof366:  lex.cs = 366; goto _test_eof
	_test_eof367:  lex.cs = 367; goto _test_eof
	_test_eof368:  lex.cs = 368; goto _test_eof
	_test_eof369:  lex.cs = 369; goto _test_eof
	_test_eof370:  lex.cs = 370; goto _test_eof
	_test_eof371:  lex.cs = 371; goto _test_eof
	_test_eof372:  lex.cs = 372; goto _test_eof
	_test_eof373:  lex.cs = 373; goto _test_eof
	_test_eof374:  lex.cs = 374; goto _test_eof
	_test_eof375:  lex.cs = 375; goto _test_eof
	_test_eof376:  lex.cs = 376; goto _test_eof
	_test_eof377:  lex.cs = 377; goto _test_eof
	_test_eof378:  lex.cs = 378; goto _test_eof
	_test_eof379:  lex.cs = 379; goto _test_eof
	_test_eof380:  lex.cs = 380; goto _test_eof
	_test_eof381:  lex.cs = 381; goto _test_eof
	_test_eof382:  lex.cs = 382; goto _test_eof
	_test_eof383:  lex.cs = 383; goto _test_eof
	_test_eof384:  lex.cs = 384; goto _test_eof
	_test_eof385:  lex.cs = 385; goto _test_eof
	_test_eof386:  lex.cs = 386; goto _test_eof
	_test_eof387:  lex.cs = 387; goto _test_eof
	_test_eof388:  lex.cs = 388; goto _test_eof
	_test_eof389:  lex.cs = 389; goto _test_eof
	_test_eof390:  lex.cs = 390; goto _test_eof
	_test_eof391:  lex.cs = 391; goto _test_eof
	_test_eof392:  lex.cs = 392; goto _test_eof
	_test_eof393:  lex.cs = 393; goto _test_eof
	_test_eof394:  lex.cs = 394; goto _test_eof
	_test_eof395:  lex.cs = 395; goto _test_eof
	_test_eof396:  lex.cs = 396; goto _test_eof
	_test_eof397:  lex.cs = 397; goto _test_eof
	_test_eof398:  lex.cs = 398; goto _test_eof
	_test_eof399:  lex.cs = 399; goto _test_eof
	_test_eof400:  lex.cs = 400; goto _test_eof
	_test_eof401:  lex.cs = 401; goto _test_eof
	_test_eof402:  lex.cs = 402; goto _test_eof
	_test_eof403:  lex.cs = 403; goto _test_eof
	_test_eof404:  lex.cs = 404; goto _test_eof
	_test_eof405:  lex.cs = 405; goto _test_eof
	_test_eof406:  lex.cs = 406; goto _test_eof
	_test_eof407:  lex.cs = 407; goto _test_eof
	_test_eof408:  lex.cs = 408; goto _test_eof
	_test_eof409:  lex.cs = 409; goto _test_eof
	_test_eof410:  lex.cs = 410; goto _test_eof
	_test_eof411:  lex.cs = 411; goto _test_eof
	_test_eof412:  lex.cs = 412; goto _test_eof
	_test_eof93:  lex.cs = 93; goto _test_eof
	_test_eof94:  lex.cs = 94; goto _test_eof
	_test_eof95:  lex.cs = 95; goto _test_eof
	_test_eof96:  lex.cs = 96; goto _test_eof
	_test_eof97:  lex.cs = 97; goto _test_eof
	_test_eof98:  lex.cs = 98; goto _test_eof
	_test_eof413:  lex.cs = 413; goto _test_eof
	_test_eof414:  lex.cs = 414; goto _test_eof
	_test_eof99:  lex.cs = 99; goto _test_eof
	_test_eof415:  lex.cs = 415; goto _test_eof
	_test_eof416:  lex.cs = 416; goto _test_eof
	_test_eof417:  lex.cs = 417; goto _test_eof
	_test_eof418:  lex.cs = 418; goto _test_eof
	_test_eof419:  lex.cs = 419; goto _test_eof
	_test_eof420:  lex.cs = 420; goto _test_eof
	_test_eof421:  lex.cs = 421; goto _test_eof
	_test_eof422:  lex.cs = 422; goto _test_eof
	_test_eof423:  lex.cs = 423; goto _test_eof
	_test_eof424:  lex.cs = 424; goto _test_eof
	_test_eof425:  lex.cs = 425; goto _test_eof
	_test_eof426:  lex.cs = 426; goto _test_eof
	_test_eof427:  lex.cs = 427; goto _test_eof
	_test_eof428:  lex.cs = 428; goto _test_eof
	_test_eof429:  lex.cs = 429; goto _test_eof
	_test_eof430:  lex.cs = 430; goto _test_eof
	_test_eof431:  lex.cs = 431; goto _test_eof
	_test_eof432:  lex.cs = 432; goto _test_eof
	_test_eof433:  lex.cs = 433; goto _test_eof
	_test_eof434:  lex.cs = 434; goto _test_eof
	_test_eof435:  lex.cs = 435; goto _test_eof
	_test_eof436:  lex.cs = 436; goto _test_eof
	_test_eof437:  lex.cs = 437; goto _test_eof
	_test_eof438:  lex.cs = 438; goto _test_eof
	_test_eof439:  lex.cs = 439; goto _test_eof
	_test_eof440:  lex.cs = 440; goto _test_eof
	_test_eof441:  lex.cs = 441; goto _test_eof
	_test_eof442:  lex.cs = 442; goto _test_eof
	_test_eof443:  lex.cs = 443; goto _test_eof
	_test_eof444:  lex.cs = 444; goto _test_eof
	_test_eof445:  lex.cs = 445; goto _test_eof
	_test_eof446:  lex.cs = 446; goto _test_eof
	_test_eof447:  lex.cs = 447; goto _test_eof
	_test_eof448:  lex.cs = 448; goto _test_eof
	_test_eof449:  lex.cs = 449; goto _test_eof
	_test_eof450:  lex.cs = 450; goto _test_eof
	_test_eof451:  lex.cs = 451; goto _test_eof
	_test_eof452:  lex.cs = 452; goto _test_eof
	_test_eof453:  lex.cs = 453; goto _test_eof
	_test_eof454:  lex.cs = 454; goto _test_eof
	_test_eof455:  lex.cs = 455; goto _test_eof
	_test_eof456:  lex.cs = 456; goto _test_eof
	_test_eof457:  lex.cs = 457; goto _test_eof
	_test_eof458:  lex.cs = 458; goto _test_eof
	_test_eof459:  lex.cs = 459; goto _test_eof
	_test_eof460:  lex.cs = 460; goto _test_eof
	_test_eof461:  lex.cs = 461; goto _test_eof
	_test_eof462:  lex.cs = 462; goto _test_eof
	_test_eof463:  lex.cs = 463; goto _test_eof
	_test_eof464:  lex.cs = 464; goto _test_eof
	_test_eof465:  lex.cs = 465; goto _test_eof
	_test_eof466:  lex.cs = 466; goto _test_eof
	_test_eof467:  lex.cs = 467; goto _test_eof
	_test_eof468:  lex.cs = 468; goto _test_eof
	_test_eof469:  lex.cs = 469; goto _test_eof
	_test_eof470:  lex.cs = 470; goto _test_eof
	_test_eof471:  lex.cs = 471; goto _test_eof
	_test_eof472:  lex.cs = 472; goto _test_eof
	_test_eof473:  lex.cs = 473; goto _test_eof
	_test_eof474:  lex.cs = 474; goto _test_eof
	_test_eof475:  lex.cs = 475; goto _test_eof
	_test_eof476:  lex.cs = 476; goto _test_eof
	_test_eof477:  lex.cs = 477; goto _test_eof
	_test_eof478:  lex.cs = 478; goto _test_eof
	_test_eof479:  lex.cs = 479; goto _test_eof
	_test_eof480:  lex.cs = 480; goto _test_eof
	_test_eof481:  lex.cs = 481; goto _test_eof
	_test_eof482:  lex.cs = 482; goto _test_eof
	_test_eof483:  lex.cs = 483; goto _test_eof
	_test_eof484:  lex.cs = 484; goto _test_eof
	_test_eof100:  lex.cs = 100; goto _test_eof
	_test_eof485:  lex.cs = 485; goto _test_eof
	_test_eof486:  lex.cs = 486; goto _test_eof
	_test_eof487:  lex.cs = 487; goto _test_eof
	_test_eof101:  lex.cs = 101; goto _test_eof
	_test_eof488:  lex.cs = 488; goto _test_eof
	_test_eof489:  lex.cs = 489; goto _test_eof
	_test_eof490:  lex.cs = 490; goto _test_eof
	_test_eof491:  lex.cs = 491; goto _test_eof
	_test_eof492:  lex.cs = 492; goto _test_eof
	_test_eof493:  lex.cs = 493; goto _test_eof
	_test_eof494:  lex.cs = 494; goto _test_eof
	_test_eof102:  lex.cs = 102; goto _test_eof
	_test_eof495:  lex.cs = 495; goto _test_eof
	_test_eof496:  lex.cs = 496; goto _test_eof
	_test_eof497:  lex.cs = 497; goto _test_eof
	_test_eof498:  lex.cs = 498; goto _test_eof
	_test_eof499:  lex.cs = 499; goto _test_eof
	_test_eof500:  lex.cs = 500; goto _test_eof
	_test_eof103:  lex.cs = 103; goto _test_eof
	_test_eof104:  lex.cs = 104; goto _test_eof
	_test_eof501:  lex.cs = 501; goto _test_eof
	_test_eof502:  lex.cs = 502; goto _test_eof
	_test_eof503:  lex.cs = 503; goto _test_eof
	_test_eof504:  lex.cs = 504; goto _test_eof
	_test_eof505:  lex.cs = 505; goto _test_eof
	_test_eof506:  lex.cs = 506; goto _test_eof
	_test_eof105:  lex.cs = 105; goto _test_eof
	_test_eof106:  lex.cs = 106; goto _test_eof
	_test_eof507:  lex.cs = 507; goto _test_eof
	_test_eof508:  lex.cs = 508; goto _test_eof
	_test_eof509:  lex.cs = 509; goto _test_eof
	_test_eof510:  lex.cs = 510; goto _test_eof
	_test_eof511:  lex.cs = 511; goto _test_eof
	_test_eof512:  lex.cs = 512; goto _test_eof
	_test_eof513:  lex.cs = 513; goto _test_eof
	_test_eof514:  lex.cs = 514; goto _test_eof
	_test_eof515:  lex.cs = 515; goto _test_eof
	_test_eof516:  lex.cs = 516; goto _test_eof
	_test_eof517:  lex.cs = 517; goto _test_eof
	_test_eof107:  lex.cs = 107; goto _test_eof
	_test_eof518:  lex.cs = 518; goto _test_eof
	_test_eof108:  lex.cs = 108; goto _test_eof
	_test_eof109:  lex.cs = 109; goto _test_eof
	_test_eof519:  lex.cs = 519; goto _test_eof
	_test_eof520:  lex.cs = 520; goto _test_eof
	_test_eof521:  lex.cs = 521; goto _test_eof
	_test_eof522:  lex.cs = 522; goto _test_eof
	_test_eof523:  lex.cs = 523; goto _test_eof
	_test_eof524:  lex.cs = 524; goto _test_eof
	_test_eof525:  lex.cs = 525; goto _test_eof
	_test_eof526:  lex.cs = 526; goto _test_eof
	_test_eof110:  lex.cs = 110; goto _test_eof
	_test_eof111:  lex.cs = 111; goto _test_eof
	_test_eof527:  lex.cs = 527; goto _test_eof
	_test_eof112:  lex.cs = 112; goto _test_eof
	_test_eof528:  lex.cs = 528; goto _test_eof
	_test_eof529:  lex.cs = 529; goto _test_eof
	_test_eof530:  lex.cs = 530; goto _test_eof
	_test_eof531:  lex.cs = 531; goto _test_eof
	_test_eof113:  lex.cs = 113; goto _test_eof
	_test_eof532:  lex.cs = 532; goto _test_eof
	_test_eof533:  lex.cs = 533; goto _test_eof
	_test_eof534:  lex.cs = 534; goto _test_eof
	_test_eof114:  lex.cs = 114; goto _test_eof
	_test_eof535:  lex.cs = 535; goto _test_eof
	_test_eof536:  lex.cs = 536; goto _test_eof
	_test_eof537:  lex.cs = 537; goto _test_eof
	_test_eof538:  lex.cs = 538; goto _test_eof
	_test_eof115:  lex.cs = 115; goto _test_eof
	_test_eof539:  lex.cs = 539; goto _test_eof
	_test_eof540:  lex.cs = 540; goto _test_eof
	_test_eof541:  lex.cs = 541; goto _test_eof
	_test_eof542:  lex.cs = 542; goto _test_eof
	_test_eof116:  lex.cs = 116; goto _test_eof
	_test_eof543:  lex.cs = 543; goto _test_eof
	_test_eof544:  lex.cs = 544; goto _test_eof
	_test_eof545:  lex.cs = 545; goto _test_eof
	_test_eof546:  lex.cs = 546; goto _test_eof
	_test_eof547:  lex.cs = 547; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 118:
			goto tr177
		case 1:
			goto tr0
		case 119:
			goto tr178
		case 121:
			goto tr183
		case 122:
			goto tr185
		case 123:
			goto tr183
		case 124:
			goto tr183
		case 125:
			goto tr191
		case 2:
			goto tr4
		case 3:
			goto tr4
		case 4:
			goto tr4
		case 126:
			goto tr194
		case 5:
			goto tr4
		case 128:
			goto tr248
		case 129:
			goto tr250
		case 6:
			goto tr10
		case 130:
			goto tr254
		case 131:
			goto tr255
		case 132:
			goto tr257
		case 133:
			goto tr259
		case 7:
			goto tr12
		case 8:
			goto tr12
		case 9:
			goto tr12
		case 10:
			goto tr12
		case 134:
			goto tr260
		case 135:
			goto tr12
		case 136:
			goto tr266
		case 137:
			goto tr260
		case 138:
			goto tr255
		case 139:
			goto tr271
		case 140:
			goto tr255
		case 141:
			goto tr255
		case 142:
			goto tr254
		case 11:
			goto tr19
		case 12:
			goto tr19
		case 143:
			goto tr255
		case 13:
			goto tr23
		case 14:
			goto tr23
		case 15:
			goto tr23
		case 16:
			goto tr23
		case 17:
			goto tr23
		case 18:
			goto tr23
		case 19:
			goto tr23
		case 20:
			goto tr23
		case 21:
			goto tr23
		case 22:
			goto tr23
		case 23:
			goto tr23
		case 24:
			goto tr23
		case 25:
			goto tr23
		case 26:
			goto tr23
		case 27:
			goto tr23
		case 28:
			goto tr23
		case 29:
			goto tr23
		case 30:
			goto tr23
		case 31:
			goto tr23
		case 32:
			goto tr23
		case 33:
			goto tr23
		case 34:
			goto tr23
		case 35:
			goto tr23
		case 36:
			goto tr23
		case 37:
			goto tr23
		case 38:
			goto tr23
		case 39:
			goto tr23
		case 40:
			goto tr23
		case 41:
			goto tr23
		case 42:
			goto tr23
		case 43:
			goto tr23
		case 44:
			goto tr23
		case 45:
			goto tr23
		case 46:
			goto tr23
		case 47:
			goto tr23
		case 48:
			goto tr23
		case 49:
			goto tr23
		case 50:
			goto tr23
		case 51:
			goto tr23
		case 52:
			goto tr23
		case 53:
			goto tr23
		case 54:
			goto tr23
		case 55:
			goto tr23
		case 56:
			goto tr23
		case 57:
			goto tr23
		case 58:
			goto tr23
		case 59:
			goto tr23
		case 60:
			goto tr23
		case 61:
			goto tr23
		case 62:
			goto tr23
		case 63:
			goto tr23
		case 144:
			goto tr255
		case 145:
			goto tr277
		case 146:
			goto tr255
		case 147:
			goto tr255
		case 148:
			goto tr255
		case 64:
			goto tr23
		case 149:
			goto tr286
		case 65:
			goto tr12
		case 66:
			goto tr12
		case 150:
			goto tr286
		case 67:
			goto tr85
		case 151:
			goto tr255
		case 68:
			goto tr23
		case 69:
			goto tr23
		case 70:
			goto tr23
		case 152:
			goto tr290
		case 153:
			goto tr286
		case 154:
			goto tr290
		case 71:
			goto tr96
		case 72:
			goto tr12
		case 155:
			goto tr295
		case 73:
			goto tr12
		case 156:
			goto tr296
		case 157:
			goto tr255
		case 158:
			goto tr255
		case 74:
			goto tr23
		case 75:
			goto tr23
		case 76:
			goto tr23
		case 77:
			goto tr23
		case 159:
			goto tr298
		case 160:
			goto tr300
		case 78:
			goto tr109
		case 161:
			goto tr255
		case 162:
			goto tr304
		case 79:
			goto tr12
		case 80:
			goto tr12
		case 81:
			goto tr12
		case 82:
			goto tr12
		case 163:
			goto tr306
		case 83:
			goto tr12
		case 84:
			goto tr12
		case 85:
			goto tr12
		case 86:
			goto tr12
		case 164:
			goto tr307
		case 165:
			goto tr255
		case 166:
			goto tr311
		case 167:
			goto tr255
		case 168:
			goto tr315
		case 169:
			goto tr255
		case 87:
			goto tr23
		case 170:
			goto tr320
		case 171:
			goto tr322
		case 88:
			goto tr126
		case 172:
			goto tr323
		case 173:
			goto tr325
		case 174:
			goto tr12
		case 89:
			goto tr12
		case 175:
			goto tr331
		case 176:
			goto tr325
		case 177:
			goto tr325
		case 178:
			goto tr325
		case 179:
			goto tr325
		case 180:
			goto tr325
		case 181:
			goto tr325
		case 182:
			goto tr325
		case 183:
			goto tr325
		case 184:
			goto tr325
		case 185:
			goto tr325
		case 186:
			goto tr325
		case 90:
			goto tr129
		case 91:
			goto tr129
		case 187:
			goto tr325
		case 188:
			goto tr325
		case 189:
			goto tr325
		case 190:
			goto tr325
		case 191:
			goto tr325
		case 192:
			goto tr325
		case 193:
			goto tr325
		case 194:
			goto tr325
		case 195:
			goto tr325
		case 196:
			goto tr325
		case 197:
			goto tr325
		case 198:
			goto tr325
		case 199:
			goto tr325
		case 200:
			goto tr325
		case 201:
			goto tr325
		case 202:
			goto tr325
		case 203:
			goto tr325
		case 204:
			goto tr325
		case 205:
			goto tr325
		case 206:
			goto tr325
		case 207:
			goto tr325
		case 208:
			goto tr325
		case 209:
			goto tr325
		case 210:
			goto tr325
		case 211:
			goto tr325
		case 212:
			goto tr325
		case 213:
			goto tr325
		case 214:
			goto tr325
		case 215:
			goto tr325
		case 216:
			goto tr325
		case 217:
			goto tr325
		case 218:
			goto tr325
		case 219:
			goto tr325
		case 220:
			goto tr325
		case 221:
			goto tr325
		case 222:
			goto tr325
		case 223:
			goto tr325
		case 224:
			goto tr325
		case 225:
			goto tr325
		case 226:
			goto tr325
		case 227:
			goto tr325
		case 228:
			goto tr325
		case 229:
			goto tr325
		case 230:
			goto tr325
		case 231:
			goto tr325
		case 232:
			goto tr325
		case 233:
			goto tr325
		case 234:
			goto tr325
		case 235:
			goto tr407
		case 236:
			goto tr325
		case 237:
			goto tr325
		case 238:
			goto tr325
		case 239:
			goto tr325
		case 240:
			goto tr325
		case 241:
			goto tr325
		case 242:
			goto tr325
		case 243:
			goto tr325
		case 244:
			goto tr325
		case 245:
			goto tr325
		case 246:
			goto tr325
		case 247:
			goto tr325
		case 248:
			goto tr325
		case 249:
			goto tr325
		case 250:
			goto tr427
		case 251:
			goto tr325
		case 252:
			goto tr325
		case 253:
			goto tr325
		case 254:
			goto tr325
		case 255:
			goto tr325
		case 256:
			goto tr325
		case 257:
			goto tr325
		case 258:
			goto tr325
		case 259:
			goto tr325
		case 260:
			goto tr325
		case 261:
			goto tr325
		case 262:
			goto tr325
		case 263:
			goto tr325
		case 264:
			goto tr325
		case 265:
			goto tr325
		case 266:
			goto tr325
		case 267:
			goto tr325
		case 268:
			goto tr325
		case 269:
			goto tr325
		case 270:
			goto tr325
		case 271:
			goto tr325
		case 272:
			goto tr325
		case 273:
			goto tr325
		case 274:
			goto tr325
		case 275:
			goto tr325
		case 276:
			goto tr456
		case 277:
			goto tr325
		case 278:
			goto tr325
		case 279:
			goto tr460
		case 280:
			goto tr325
		case 281:
			goto tr325
		case 282:
			goto tr325
		case 283:
			goto tr325
		case 284:
			goto tr325
		case 285:
			goto tr325
		case 286:
			goto tr325
		case 287:
			goto tr325
		case 288:
			goto tr325
		case 289:
			goto tr325
		case 290:
			goto tr325
		case 291:
			goto tr325
		case 292:
			goto tr325
		case 293:
			goto tr325
		case 294:
			goto tr325
		case 295:
			goto tr325
		case 296:
			goto tr325
		case 297:
			goto tr325
		case 298:
			goto tr325
		case 299:
			goto tr325
		case 300:
			goto tr325
		case 301:
			goto tr325
		case 302:
			goto tr325
		case 303:
			goto tr325
		case 304:
			goto tr492
		case 305:
			goto tr325
		case 306:
			goto tr325
		case 307:
			goto tr325
		case 308:
			goto tr325
		case 309:
			goto tr325
		case 310:
			goto tr325
		case 311:
			goto tr325
		case 312:
			goto tr325
		case 313:
			goto tr325
		case 314:
			goto tr325
		case 315:
			goto tr325
		case 316:
			goto tr325
		case 317:
			goto tr325
		case 318:
			goto tr325
		case 319:
			goto tr325
		case 320:
			goto tr325
		case 321:
			goto tr325
		case 322:
			goto tr325
		case 323:
			goto tr325
		case 324:
			goto tr325
		case 325:
			goto tr325
		case 326:
			goto tr325
		case 327:
			goto tr325
		case 328:
			goto tr325
		case 329:
			goto tr325
		case 330:
			goto tr325
		case 331:
			goto tr325
		case 332:
			goto tr325
		case 333:
			goto tr325
		case 334:
			goto tr325
		case 335:
			goto tr325
		case 336:
			goto tr325
		case 337:
			goto tr325
		case 338:
			goto tr325
		case 339:
			goto tr325
		case 340:
			goto tr325
		case 341:
			goto tr325
		case 342:
			goto tr325
		case 343:
			goto tr325
		case 344:
			goto tr535
		case 92:
			goto tr12
		case 345:
			goto tr537
		case 346:
			goto tr325
		case 347:
			goto tr325
		case 348:
			goto tr325
		case 349:
			goto tr325
		case 350:
			goto tr325
		case 351:
			goto tr325
		case 352:
			goto tr325
		case 353:
			goto tr325
		case 354:
			goto tr325
		case 355:
			goto tr325
		case 356:
			goto tr325
		case 357:
			goto tr325
		case 358:
			goto tr325
		case 359:
			goto tr325
		case 360:
			goto tr325
		case 361:
			goto tr325
		case 362:
			goto tr325
		case 363:
			goto tr325
		case 364:
			goto tr325
		case 365:
			goto tr325
		case 366:
			goto tr325
		case 367:
			goto tr325
		case 368:
			goto tr325
		case 369:
			goto tr325
		case 370:
			goto tr325
		case 371:
			goto tr567
		case 372:
			goto tr325
		case 373:
			goto tr325
		case 374:
			goto tr325
		case 375:
			goto tr325
		case 376:
			goto tr325
		case 377:
			goto tr325
		case 378:
			goto tr325
		case 379:
			goto tr325
		case 380:
			goto tr325
		case 381:
			goto tr325
		case 382:
			goto tr325
		case 383:
			goto tr325
		case 384:
			goto tr325
		case 385:
			goto tr325
		case 386:
			goto tr325
		case 387:
			goto tr325
		case 388:
			goto tr325
		case 389:
			goto tr325
		case 390:
			goto tr325
		case 391:
			goto tr325
		case 392:
			goto tr325
		case 393:
			goto tr325
		case 394:
			goto tr325
		case 395:
			goto tr325
		case 396:
			goto tr325
		case 397:
			goto tr325
		case 398:
			goto tr325
		case 399:
			goto tr325
		case 400:
			goto tr325
		case 401:
			goto tr325
		case 402:
			goto tr325
		case 403:
			goto tr325
		case 404:
			goto tr325
		case 405:
			goto tr325
		case 406:
			goto tr325
		case 407:
			goto tr325
		case 408:
			goto tr325
		case 409:
			goto tr325
		case 410:
			goto tr325
		case 411:
			goto tr325
		case 412:
			goto tr613
		case 93:
			goto tr132
		case 94:
			goto tr132
		case 95:
			goto tr132
		case 96:
			goto tr132
		case 97:
			goto tr132
		case 98:
			goto tr132
		case 413:
			goto tr614
		case 414:
			goto tr615
		case 99:
			goto tr144
		case 415:
			goto tr255
		case 416:
			goto tr325
		case 417:
			goto tr325
		case 418:
			goto tr325
		case 419:
			goto tr325
		case 420:
			goto tr325
		case 421:
			goto tr325
		case 422:
			goto tr325
		case 423:
			goto tr325
		case 424:
			goto tr325
		case 425:
			goto tr325
		case 426:
			goto tr325
		case 427:
			goto tr325
		case 428:
			goto tr325
		case 429:
			goto tr325
		case 430:
			goto tr325
		case 431:
			goto tr325
		case 432:
			goto tr325
		case 433:
			goto tr325
		case 434:
			goto tr325
		case 435:
			goto tr325
		case 436:
			goto tr325
		case 437:
			goto tr325
		case 438:
			goto tr325
		case 439:
			goto tr325
		case 440:
			goto tr325
		case 441:
			goto tr325
		case 442:
			goto tr325
		case 443:
			goto tr325
		case 444:
			goto tr325
		case 445:
			goto tr325
		case 446:
			goto tr325
		case 447:
			goto tr325
		case 448:
			goto tr325
		case 449:
			goto tr325
		case 450:
			goto tr325
		case 451:
			goto tr325
		case 452:
			goto tr325
		case 453:
			goto tr325
		case 454:
			goto tr325
		case 455:
			goto tr325
		case 456:
			goto tr325
		case 457:
			goto tr325
		case 458:
			goto tr325
		case 459:
			goto tr325
		case 460:
			goto tr325
		case 461:
			goto tr325
		case 462:
			goto tr325
		case 463:
			goto tr325
		case 464:
			goto tr325
		case 465:
			goto tr325
		case 466:
			goto tr325
		case 467:
			goto tr325
		case 468:
			goto tr325
		case 469:
			goto tr325
		case 470:
			goto tr325
		case 471:
			goto tr325
		case 472:
			goto tr325
		case 473:
			goto tr325
		case 474:
			goto tr325
		case 475:
			goto tr325
		case 476:
			goto tr325
		case 477:
			goto tr325
		case 478:
			goto tr325
		case 479:
			goto tr325
		case 480:
			goto tr325
		case 481:
			goto tr255
		case 483:
			goto tr699
		case 484:
			goto tr701
		case 100:
			goto tr146
		case 485:
			goto tr705
		case 486:
			goto tr705
		case 487:
			goto tr705
		case 101:
			goto tr148
		case 488:
			goto tr708
		case 490:
			goto tr712
		case 491:
			goto tr713
		case 492:
			goto tr717
		case 494:
			goto tr725
		case 495:
			goto tr727
		case 496:
			goto tr728
		case 497:
			goto tr732
		case 498:
			goto tr725
		case 499:
			goto tr732
		case 501:
			goto tr744
		case 502:
			goto tr745
		case 503:
			goto tr749
		case 504:
			goto tr749
		case 505:
			goto tr749
		case 507:
			goto tr762
		case 508:
			goto tr763
		case 509:
			goto tr767
		case 510:
			goto tr767
		case 511:
			goto tr767
		case 513:
			goto tr772
		case 515:
			goto tr779
		case 516:
			goto tr781
		case 517:
			goto tr779
		case 107:
			goto tr158
		case 518:
			goto tr779
		case 108:
			goto tr158
		case 109:
			goto tr158
		case 519:
			goto tr784
		case 521:
			goto tr794
		case 522:
			goto tr795
		case 523:
			goto tr796
		case 524:
			goto tr798
		case 525:
			goto tr799
		case 526:
			goto tr799
		case 110:
			goto tr162
		case 111:
			goto tr162
		case 527:
			goto tr799
		case 112:
			goto tr162
		case 528:
			goto tr799
		case 529:
			goto tr803
		case 531:
			goto tr806
		case 113:
			goto tr166
		case 533:
			goto tr811
		case 534:
			goto tr813
		case 114:
			goto tr169
		case 535:
			goto tr817
		case 537:
			goto tr822
		case 538:
			goto tr824
		case 115:
			goto tr171
		case 539:
			goto tr828
		case 541:
			goto tr833
		case 542:
			goto tr835
		case 116:
			goto tr173
		case 543:
			goto tr839
		case 545:
			goto tr843
		case 546:
			goto tr844
		case 547:
			goto tr848
		}
	}

	_out: {}
	}

// line internal/php8/scanner.rl:506


    tkn.Value = lex.data[lex.ts:lex.te]
    tkn.ID = token.ID(tok)

    return tkn
}