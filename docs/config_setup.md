# LibreFlex config.json setup

### Here's an example config.json file:
```json
{
    "flexApiServerConfig":{
        "enabled": true,
        "port": 9090,
        "donateUrl": "https://google.com",
        "adsApiUrl": "http://192.168.1.207:9091/FlexAds",
        "vkApiToken": "abcdef123",

        "adsRepoProvider": "dummy",
        "adsRepoConfig":{
            "dummyBanner": "https://flexvk.es3n.in/wtf.png",
            "dummyTarget": "https://google.com"
        }
    },
    
    "flexAdsServerConfig":{
        "enabled": true,
        "port": 9091, 

        "adsRepoProvider": "dummy",
        "adsRepoConfig":{
            "dummyBanner": "https://flexvk.es3n.in/wtf.png",
            "dummyTarget": "https://google.com"
        }
    }
}
```

As you can see, it consists of 2 top-level entities - **flexApiServerConfig** and **flexAdsServerConfig**, we will cover their fields in this doc.


## flexApiServerConfig
This is a config for the FlexApi server
### Fields: 
**enabled** - whether FlexApi server is enabled  
**port** - the port on which FlexApi server will run  
**donateUrl** - donation link for this server. Will be returned in getInfo method and such  
**adsApiUrl** - FlexAds server url for this FlexApi server  
**vkApiToken** - VK API token with audio access and subscription
**adsRepoProvider** - Ads repo provider type, we'll cover this later  
**adsRepoConfig** - Ads repo provider config, we'll cover this later  


## flexAdsServerConfig
This is a config for the FlexAds server
### Fields: 
**enabled** - whether FlexApi server is enabled  
**port** - the port on which FlexApi server will run  
**adsRepoProvider** - Ads repo provider type, we'll cover this later  
**adsRepoConfig** - Ads repo provider config, we'll cover this later  


## Ads repo provider
Some of the entities above have **adsRepoProvider** and **adsRepoConfig** fields, those are responsible for providing ads. At the moment, there are 2 available providers:
* [dummy](dummy_setup)
* [postgres](postgres_setup)

Check out their respective articles to find out, how to set them up.