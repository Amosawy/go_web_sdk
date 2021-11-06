package router

import (
	"encoding/json"
	"github.com/Amosawy/go_web_sdk/seelog"
	"github.com/Amosawy/go_web_sdk/tools"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type TomlConfig struct {
	Route  routeConfig
	Routes map[string]Route
}

type routeConfig struct {
	RoutePath string `toml:"route_path"`
}

type Route struct {
	Host string `json:"host"`
	Scheme string `json:"scheme"`
	Uri  string `json:"uri"`
}

func LoadRouteConfig() (map[string]Route, error) {
	abspath, path, err := tools.GetRouteConfigPath()
	if err != nil {
		seelog.Errorf("get config path error %s", err.Error())
		return nil, err
	}
	if _, err := os.Stat(abspath); err != nil {
		return nil, err
	}
	Conf := new(TomlConfig)
	if _, err := toml.DecodeFile(abspath, &Conf); err != nil {
		seelog.Errorf("toml.DecodeFile error %s", err.Error())
		return nil, err
	}
	Conf.Routes = make(map[string]Route)
	routes, err := ioutil.ReadFile(path + Conf.Route.RoutePath)
	if err != nil {
		seelog.Errorf("read routePath error %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(routes, &Conf.Routes)
	if err != nil {
		seelog.Errorf("json.Unmarshal routePath error %s", err.Error())
		return nil, err
	}
	return Conf.Routes, nil
}
