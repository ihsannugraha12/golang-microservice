package utils

import (
	"os"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func BindFromJson(dest any, filename, path string) error {
	viper := viper.New()
	viper.SetConfigType("json")
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("Failed to unmarshal config: %v", err)
		return err
	}

	return nil
}

func setEnvFromConsulKV(v *viper.Viper) error {
	env := make(map[string]any)
	err := v.Unmarshal(&env)
	if err != nil {
		logrus.Errorf("Failed to unmarshal config: %v", err)
		return err
	}

	for key, value := range env {
		var (
			valOf = reflect.ValueOf(value)
			val   string
		)

		switch valOf.Kind() {
		case reflect.String:
			val = valOf.String()
		case reflect.Int, reflect.Int64:
			val = strconv.FormatInt(valOf.Int(), 10)
		case reflect.Uint, reflect.Uint64:
			val = strconv.FormatUint(valOf.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			val = strconv.FormatFloat(valOf.Float(), 'f', -1, 64)
		case reflect.Bool:
			val = strconv.FormatBool(valOf.Bool())
		default:
			logrus.Warnf("Unsupported type for key %s: %s", key, valOf.Kind())
			panic("Unsupported type for environment variable")
		}

		err = os.Setenv(key, val)
		if err != nil {
			logrus.Errorf("Failed to set environment variable %s: %v", key, err)
			return err
		}
	}

	return nil
}

func BindFromConsul(dest any, endpoint, path string) error {
	viper := viper.New()
	viper.SetConfigType("json")

	err := viper.AddRemoteProvider("consul", endpoint, path)
	if err != nil {
		logrus.Errorf("Failed to add remote provider: %v", err)
		return err
	}

	err = viper.ReadRemoteConfig()
	if err != nil {
		logrus.Errorf("Failed to read config: %v", err)
		return err
	}

	err = viper.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("Failed to unmarshal config: %v", err)
		return err
	}

	err = setEnvFromConsulKV(viper)
	if err != nil {
		logrus.Errorf("Failed to set environment variables from Consul KV: %v", err)
		return err
	}

	return nil
}
