package coco

type Context struct {
	handlers []Handler
}

// next calls the next handler in the chain if there is one.
// If there is no next handler, the request is terminated.
func (c *Context) next(rw Response, req *Request) {

	if len(c.handlers) == 0 {
		return
	}

	// Take the first handler off the list and call it.
	h := c.handlers[0]
	c.handlers = c.handlers[1:]

	h(rw, req, c.next)

}
