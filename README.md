# Rocket Pool - Smart Node Package for ARM64

Please refer to https://github.com/rocket-pool/smartnode for official information.

## Installation

```
cd $HOME
wget https://raw.githubusercontent.com/RomiRand/smartnode/master/downloads/install.sh
chmod +x install.sh  
./install.sh
rocketpool service config
rocketpool service start
```

Please don't run `rocketpool service install`, it may causes trouble with the docker-compose installation. 

### Updating:
First, stop the rocketpool stack: `rocketpool service pause`.
Then you can just run the steps above to update. Note that your config files will be overwritten, so if you've made any config changes, better take a backup: `sudo cp -r ~/.rocketpool ~/.rocketpool.bak`

If you have any questions or feedback, text me on Discord: `RamiRond#1730`