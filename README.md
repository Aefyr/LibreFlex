# LibreFlex
**LibreFlex** is an opensource implementation of Sova Lite's FlexApi and FlexAds written in go.

## Running LibreFlex
Just provide a valid config.json file in the same directory as the LibreFlex binary and start the binary. FlexApi will run on `http://yourIp:flexApiPortFromConfig/FlexMusic` and FlexAds will run on `http://yourIp:flexAdsPortFromConfig/FlexAds`.  

It's recommended to wrap these servers with a reverse proxy such as nginx since there is no SSL support in LibreFlex and FlexApi and FlexAds servers run on different ports.  

Config file structure is described in [docs/config_setup](docs/config_setup.md).