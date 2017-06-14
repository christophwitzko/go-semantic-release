			Author:    Signature{Name: "Foo", Email: "foo@example.local", When: ts},
			Committer: Signature{Name: "Bar", Email: "bar@example.local", When: ts},
			Message:   "Message\n\nFoo\nBar\nWith trailing blank lines\n\n",
			tree:      plumbing.NewHash("f000000000000000000000000000000000000001"),
			parents:   []plumbing.Hash{plumbing.NewHash("f000000000000000000000000000000000000002")},
			tree:      plumbing.NewHash("0000000000000000000000000000000000000003"),
			parents: []plumbing.Hash{
func (s *SuiteCommit) TestHistory(c *C) {
	commits, err := s.Commit.History()
	c.Assert(err, IsNil)
	c.Assert(commits, HasLen, 5)
	c.Assert(commits[0].Hash.String(), Equals, s.Commit.Hash.String())
	c.Assert(commits[len(commits)-1].Hash.String(), Equals, "b029517f6300c2da0f4b651b8642506cd6aaf45d")
}
