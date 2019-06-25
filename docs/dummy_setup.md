# Dummy ads repo provider
Dummy ads repo provider will just return the same ad from its config all the time. Useful for testing or when you don't wanna implement actual ad stuff.

### Here's how to set it up:
Use "dummy" as your **adsRepoProvider** and provide a valid config, config fields are explained after the example  

Example:
```json
"flexApiServerConfig":{
    ...
   "adsRepoProvider":"dummy",
   "adsRepoConfig":{
      "dummyBanner":"https://flexvk.es3n.in/wtf.png",
      "dummyTarget":"https://google.com"
   }
   ...
}
```

### adsRepoConfig fields explanation:
**dummyBanner** is a banner image link that will be returned as ad all the time  
**dummyTarget** is the link that will be opened when user clicks on the ad