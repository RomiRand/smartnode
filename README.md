# Rocket Pool - Smart Node Package for ARM64

Please refer to https://github.com/rocket-pool/smartnode for official information.

## Installation

NOTE: Currently only supports lighthouse client!

```
cd $HOME
wget https://raw.githubusercontent.com/RomiRand/smartnode/master/downloads/install.sh
chmod +x install.sh  
./install.sh
rocketpool service config
rocketpool service start
```

Please don't run `rocketpool service install`, it may causes trouble with the docker-compose installation. 

If you have any questions or feedback, text me on Discord: `RamiRond#1730`