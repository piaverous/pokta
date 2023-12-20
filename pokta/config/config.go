package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Okta   OktaConfig `mapstructure:"okta"`
	DryRun bool       `mapstructure:"dry_run"`
}

// Configuration of how pokta should interact with Jira.
type OktaConfig struct {
	PrivateKey   string `mapstructure:"private_key"`
	Kid          string `mapstructure:"kid"`
	Aud          string `mapstructure:"aud"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	ApiKey       string `mapstructure:"api_key"`
	Scope        string `mapstructure:"scope"`
}

// Load pokta coniguration.
func (c *Config) Load(flags *pflag.FlagSet) error {
	v := viper.New()

	// pokta looks for configuration files called config.yaml, config.json,
	// config.toml, config.hcl, etc.
	v.SetConfigName("config")

	// pokta looks for configuration files in the common configuration
	// directories.
	v.AddConfigPath("/etc/pokta")
	v.AddConfigPath("$HOME/.local/.pokta")
	v.AddConfigPath("$HOME/.pokta")

	// Viper logs the configuration file it uses, if any.
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}

	// configFile := v.ConfigFileUsed()
	// if configFile != "" {
	// 	fmt.Fprintf(os.Stderr, "Found config file at %s\n", configFile)
	// }

	// pokta can be configured with environment variables that start with
	// pokta_.
	v.SetEnvPrefix("pokta")
	v.AutomaticEnv()

	// Options with dashes in flag names have underscores when set inside a
	// configuration file or with environment variables.
	flags.SetNormalizeFunc(func(fs *pflag.FlagSet, name string) pflag.NormalizedName {
		name = strings.ReplaceAll(name, "-", "_")
		return pflag.NormalizedName(name)
	})
	v.BindPFlags(flags)

	// Nested configuration options set with environment variables use an
	// underscore as a separator.
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvironmentVariables(v, *c)

	return v.Unmarshal(c)
}

// bindEnvironmentVariables inspects iface's structure and recursively binds its
// fields to environment variables. This is a workaround to a limitation of
// Viper, found here:
// https://github.com/spf13/viper/issues/188#issuecomment-399884438
func bindEnvironmentVariables(v *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		val := ifv.Field(i)
		typ := ift.Field(i)
		tv, ok := typ.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch val.Kind() {
		case reflect.Struct:
			bindEnvironmentVariables(v, val.Interface(), append(parts, tv)...)
		default:
			v.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
