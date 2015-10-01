package req
// Generated ggen -- do not change

const (
	PUT = "PUT"
	POST = "POST"
	PATCH = "PATCH"
	INS = "INS"
	DEL = "DEL"
	GET = "GET"
	HEAD = "HEAD"
)

// Methods for the Rest state
func (r *Rest) Put() *Rest {
	return r.Method(PUT)
}
func (r *Rest) Post() *Rest {
	return r.Method(POST)
}
func (r *Rest) Patch() *Rest {
	return r.Method(PATCH)
}
func (r *Rest) Ins() *Rest {
	return r.Method(INS)
}
func (r *Rest) Del() *Rest {
	return r.Method(DEL)
}
func (r *Rest) Get() *Rest {
	return r.Method(GET)
}
func (r *Rest) Head() *Rest {
	return r.Method(HEAD)
}

// Methods for Req state
func (r *Req) Put() *Req {
	return r.Method(PUT)
}
func (r *Req) Post() *Req {
	return r.Method(POST)
}
func (r *Req) Patch() *Req {
	return r.Method(PATCH)
}
func (r *Req) Ins() *Req {
	return r.Method(INS)
}
func (r *Req) Del() *Req {
	return r.Method(DEL)
}
func (r *Req) Get() *Req {
	return r.Method(GET)
}
func (r *Req) Head() *Req {
	return r.Method(HEAD)
}
