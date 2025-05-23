package mg

type Context struct {
	AcceptNewline bool
	InsideForm    bool
}

func makeContext() Context {
	return Context{AcceptNewline: true}
}

func (c Context) setInsideForm(insideForm bool) Context {
	c.AcceptNewline = insideForm
	c.InsideForm = insideForm
	return c
}

func (c Context) setAcceptNewline(acceptNewline bool) Context {
	c.AcceptNewline = acceptNewline
	return c
}

func (c Context) setInsideDelimiters(insideBraces bool) Context {
	c.AcceptNewline = insideBraces
	return c.setInsideForm(false).setAcceptNewline(insideBraces)
}
