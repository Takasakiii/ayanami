package cuid

func (c *Cuid) Generate() string {
	return c.cuid()
}
