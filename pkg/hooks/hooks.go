package hooks

import "fmt"

type Hooks interface {
	Init(map[string]string) error
	Name() string
	Version() string
	Success(*SuccessHookConfig) error
	NoRelease(*NoReleaseConfig) error
	PreRelease(*SuccessHookConfig) error
}

type ChainedHooksExecutor struct {
	HooksChain []Hooks
}

func (c *ChainedHooksExecutor) Success(config *SuccessHookConfig) error {
	for _, h := range c.HooksChain {
		name := h.Name()
		err := h.Success(config)
		if err != nil {
			return fmt.Errorf("%s hook has failed: %w", name, err)
		}
	}
	return nil
}

func (c *ChainedHooksExecutor) NoRelease(config *NoReleaseConfig) error {
	for _, h := range c.HooksChain {
		name := h.Name()
		err := h.NoRelease(config)
		if err != nil {
			return fmt.Errorf("%s hook has failed: %w", name, err)
		}
	}
	return nil
}

func (c *ChainedHooksExecutor) PreRelease(config *SuccessHookConfig) error {
	for _, h := range c.HooksChain {
		name := h.Name()
		err := h.PreRelease(config)
		if err != nil {
			return fmt.Errorf("%s hook has failed: %w", name, err)
		}
	}
	return nil
}

func (c *ChainedHooksExecutor) Init(conf map[string]string) error {
	for _, h := range c.HooksChain {
		err := h.Init(conf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ChainedHooksExecutor) GetNameVersionPairs() []string {
	ret := make([]string, len(c.HooksChain))
	for i, h := range c.HooksChain {
		ret[i] = fmt.Sprintf("%s@%s", h.Name(), h.Version())
	}
	return ret
}
