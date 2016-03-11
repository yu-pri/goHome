package home

type Config struct {
  emailTo string
  emailFrom string
}

func LoadConfiguration() Config {
    var cfg Config
    cfg.emailTo = "alPrihodko@gmail.com"
    cfg.emailFrom = "alPrihodko@gmail.com"
    return cfg
}
