package command

type GetFullLinkCommand struct {
	ShortenedLink string
}

func (c *GetFullLinkCommand) Execute() (string, error) {
	// business logic
	//fullLink, err := c.Database.GetFullLink(context.Background(), c.ShortenedLink)
	//if err != nil {
	//	log.Println(err)
	//	return ""
	//}
	//return fullLink
	return "", nil
}
