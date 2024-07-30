package main

type reqmethod string
type status int

const (
	// Request method
	METHOD_GET     reqmethod = "GET"
	METHOD_PUT     reqmethod = "PUT"
	METHOD_POST    reqmethod = "POST"
	METHOD_DELETE  reqmethod = "DELETE"
	METHOD_PATCH   reqmethod = "PATCH"
	METHOD_HEAD    reqmethod = "HEAD"
	METHOD_OPTIONS reqmethod = "OPTIONS"
	METHOD_TRACE   reqmethod = "TRACE"
	METHOD_CONNECT reqmethod = "CONNECT"

	// Status codes
	CODE_100 status = 100
	CODE_101 status = 101
	CODE_102 status = 102
	CODE_103 status = 103
	CODE_122 status = 122
	CODE_200 status = 200
	CODE_201 status = 201
	CODE_202 status = 202
	CODE_203 status = 203
	CODE_204 status = 204
	CODE_205 status = 205
	CODE_206 status = 206
	CODE_207 status = 207
	CODE_208 status = 208
	CODE_226 status = 226
	CODE_300 status = 300
	CODE_301 status = 301
	CODE_302 status = 302
	CODE_303 status = 303
	CODE_304 status = 304
	CODE_305 status = 305
	CODE_306 status = 306
	CODE_307 status = 307
	CODE_308 status = 308
	CODE_400 status = 400
	CODE_401 status = 401
	CODE_402 status = 402
	CODE_403 status = 403
	CODE_404 status = 404
	CODE_405 status = 405
	CODE_406 status = 406
	CODE_407 status = 407
	CODE_408 status = 408
	CODE_409 status = 409
	CODE_410 status = 410
	CODE_411 status = 411
	CODE_412 status = 412
	CODE_413 status = 413
	CODE_414 status = 414
	CODE_415 status = 415
	CODE_416 status = 416
	CODE_417 status = 417
	CODE_418 status = 418
	CODE_421 status = 421
	CODE_422 status = 422
	CODE_423 status = 423
	CODE_424 status = 424
	CODE_425 status = 425
	CODE_426 status = 426
	CODE_428 status = 428
	CODE_429 status = 429
	CODE_431 status = 431
	CODE_444 status = 444
	CODE_449 status = 449
	CODE_450 status = 450
	CODE_451 status = 451
	CODE_499 status = 499
	CODE_500 status = 500
	CODE_501 status = 501
	CODE_502 status = 502
	CODE_503 status = 503
	CODE_504 status = 504
	CODE_505 status = 505
	CODE_506 status = 506
	CODE_507 status = 507
	CODE_509 status = 509
	CODE_510 status = 510
	CODE_511 status = 511
)
