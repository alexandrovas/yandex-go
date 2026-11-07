package main

type Config struct {
	Version string
	Plugins []string
	Stat    map[string]int
}

func (cfg *Config) Clone() *Config {
	clone := &Config{
		Version: cfg.Version,
		Plugins: make([]string, len(cfg.Plugins)),
		Stat:    make(map[string]int, len(cfg.Stat)),
	}
	copy(clone.Plugins, cfg.Plugins)
	for k, v := range cfg.Stat {
		clone.Stat[k] = v
	}
	return clone
}

func main() {

}
