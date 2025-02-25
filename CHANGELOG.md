
<a name="Onvif Camera Device Service (found in device-onvif-camera) Changelog"></a>
## EdgeX ONVIF Camera Device Service
[Github repository](https://github.com/edgexfoundry/device-onvif-camera)

### Change Logs for EdgeX Dependencies
- [device-sdk-go](https://github.com/edgexfoundry/device-sdk-go/blob/main/CHANGELOG.md)
- [go-mod-core-contracts](https://github.com/edgexfoundry/go-mod-core-contracts/blob/main/CHANGELOG.md)
- [go-mod-bootstrap](https://github.com/edgexfoundry/go-mod-bootstrap/blob/main/CHANGELOG.md)
- [go-mod-messaging](https://github.com/edgexfoundry/go-mod-messaging/blob/main/CHANGELOG.md) (indirect dependency)
- [go-mod-registry](https://github.com/edgexfoundry/go-mod-registry/blob/main/CHANGELOG.md)  (indirect dependency)
- [go-mod-secrets](https://github.com/edgexfoundry/go-mod-secrets/blob/main/CHANGELOG.md) (indirect dependency)
- [go-mod-configuration](https://github.com/edgexfoundry/go-mod-configuration/blob/main/CHANGELOG.md) (indirect dependency)

## [4.0.0] Odessa - 2025-03-12 (Only compatible with the 4.x releases)

### ✨  Features

- Enable PIE support for ASLR and full RELRO ([a83730c…](https://github.com/edgexfoundry/device-onvif-camera/commit/a83730c2ca7660bd580c8c11ae1217f2c0590c44))


### ♻ Code Refactoring

- Update go module to v4 ([9683ead…](https://github.com/edgexfoundry/device-onvif-camera/commit/9683eadad6a6839e290b802346720f2e8068c870))
```text

BREAKING CHANGE: update go module to v4

```

### 🐛 Bug Fixes

- Add Clients and MessageBus Host in config to allow env overrides ([b15dd37…](https://github.com/edgexfoundry/device-onvif-camera/commit/b15dd379a12335186161310b4f4cd640302a8412))
- Ensures that the parsed value falls within the int32 range ([22d23d7…](https://github.com/edgexfoundry/device-onvif-camera/commit/22d23d7ae3353e0154e35ebad73c72e7f4fdf7f8))
- Fix incorrect conversion between integer types ([528a90e…](https://github.com/edgexfoundry/device-onvif-camera/commit/528a90ea40b70368105a33b93baeba7a4762d77e))
- Only one ldflags flag is allowed ([1059907…](https://github.com/edgexfoundry/device-onvif-camera/commit/10599072c9655e29548c116613113411af81b17c))

### 👷 Build

- Upgrade to go-1.23, Linter1.61.0 and Alpine 3.20 ([2342540…](https://github.com/edgexfoundry/device-onvif-camera/commit/2342540ef65f1397ecb34c477184ee441b1d0a60))


## [v3.1.0] Napa - 2023-11-15 (Only compatible with the 3.x releases)


### <!-- 0 -->✨  Features

- Remove snap packaging ([#396](https://github.com/edgexfoundry/device-onvif-camera/issues/396)) ([f9aa1e7…](https://github.com/edgexfoundry/device-onvif-camera/commit/f9aa1e7837344c75d2a4d485c1648989cda54cfb))
```text

BREAKING CHANGE: Remove snap packaging ([#396](https://github.com/edgexfoundry/device-onvif-camera/issues/396))

```
- Replace gorilla/mux with labstack/echo ([9a9a093…](https://github.com/edgexfoundry/device-onvif-camera/commit/9a9a093bba13e018aff1d59529fadb53acfad024))


### ♻ Code Refactoring

- Remove obsolete comments from config file ([#398](https://github.com/edgexfoundry/device-onvif-camera/issues/398)) ([bf74ebf…](https://github.com/edgexfoundry/device-onvif-camera/commit/bf74ebf12077028c42a990cb9f5eb3b4e8875dd7))


### 📖 Documentation

- Fix license link ([#373](https://github.com/edgexfoundry/device-onvif-camera/issues/373)) ([f8d6efe…](https://github.com/edgexfoundry/device-onvif-camera/commit/f8d6efeffbcc72b13ace8b41daac95c88037c92b))
- Add readme and rename collections to clarify collection usage ([f9a51d9…](https://github.com/edgexfoundry/device-onvif-camera/commit/f9a51d90d45793b537885a974cdf3c636c247a85))


### 👷 Build

- Upgrade to go-1.21, Linter1.54.2 and Alpine 3.18 ([#384](https://github.com/edgexfoundry/device-onvif-camera/issues/384)) ([9825d35…](https://github.com/edgexfoundry/device-onvif-camera/commit/9825d353c0ae3edd45665d4480ee2bc38ab097aa))


### 🤖 Continuous Integration

- Add automated release workflow on tag creation ([6bc619b…](https://github.com/edgexfoundry/device-onvif-camera/commit/6bc619b75e5ee3284532e8add0ff78116f112ad2))


## [3.0.0] Minnesota - 2023-05-31 (Only compatible with the 3.x releases)

### Features ✨
- Implement driver Start command ([#343](https://github.com/edgexfoundry/device-onvif-camera/issues/343)) ([#b5d61d0](https://github.com/edgexfoundry/device-onvif-camera/commits/b5d61d0))
- Implement secret change support, refactor cred mgmt ([#333](https://github.com/edgexfoundry/device-onvif-camera/issues/333)) ([#a510015](https://github.com/edgexfoundry/device-onvif-camera/commits/a510015))
- Consume SDK interface changes ([#7c14cf8](https://github.com/edgexfoundry/device-onvif-camera/commits/7c14cf8))
  ```text
  BREAKING CHANGE: Consume SDK interface changes by adding Discover and ValidateDevice func on driver
  ``` 
- Log raw response when xml unmarshal fails ([#ff76725](https://github.com/edgexfoundry/device-onvif-camera/commits/ff76725))
- Updates for common config ([#ce7fe07](https://github.com/edgexfoundry/device-onvif-camera/commits/ce7fe07))
  ```text
  BREAKING CHANGE: Configuration file changed to remove common config settings
  ``` 
- Use latest SDK for MessageBus Request API ([#4452f94](https://github.com/edgexfoundry/device-onvif-camera/commits/4452f94))
  ```text
  BREAKING CHANGE: Commands via MessageBus topic configuration changed
  ``` 
- Remove ZeroMQ messagebus capability ([#848d5d6](https://github.com/edgexfoundry/device-onvif-camera/commits/848d5d6))
  ```text
  BREAKING CHANGE: Remove ZeroMQ messagebus capability
  ```  
  
### Bug Fixes 🐛
- Get command should return server error instead of timeout error ([#7f95bcf](https://github.com/edgexfoundry/device-onvif-camera/commits/7f95bcf))
- **snap:** Refactor to avoid conflicts with readonly config provider directory ([#279](https://github.com/edgexfoundry/device-onvif-camera/issues/279)) ([#e03c6e9](https://github.com/edgexfoundry/device-onvif-camera/commits/e03c6e9))

### Code Refactoring ♻
- Prefix some commands with onvif service name and also add SnapshotUri command ([#7c6b29b](https://github.com/edgexfoundry/device-onvif-camera/commits/7c6b29b))
  ```text
  BREAKING CHANGE: Changed the name of many device commands to include a prefix of the onvif service they belong to. For example Status becomes PTZStatus, Profiles becomes MediaProfiles, etc.
  ``` 
- Modify UpdateDevice calls to use PatchDevice calls ([#355](https://github.com/edgexfoundry/device-onvif-camera/issues/355)) ([#78f21e3](https://github.com/edgexfoundry/device-onvif-camera/commits/78f21e3))
- Fix inconsistent naming of Edgex commands ([#afec685](https://github.com/edgexfoundry/device-onvif-camera/commits/afec685))
  ```text
  BREAKING CHANGE: Making Edgex command names consistent by removing GET and SET keywords
  ```
- Consume Provision Watcher changes for running multiple instances ([#349](https://github.com/edgexfoundry/device-onvif-camera/issues/349)) ([#1961be3](https://github.com/edgexfoundry/device-onvif-camera/commits/1961be3))
- Changed configuration and provision watcher file format to yaml ([#163d81e](https://github.com/edgexfoundry/device-onvif-camera/commits/163d81e))
  ```text
  BREAKING CHANGE: Configuration and provision watcher files are now in YAML format
  ``` 
- Consume device-sdk-go breaking changes ([#35e0b63](https://github.com/edgexfoundry/device-onvif-camera/commits/35e0b63))
  ```text
  BREAKING CHANGE: Update ProtocolDriver implementation for the new ProtocolDriver interface changes
  ```
- Rename path to secret name ([#33d7f92](https://github.com/edgexfoundry/device-onvif-camera/commits/33d7f92))
  ```text
  BREAKING CHANGE: Rename path to secret name
  ``` 
- Use device sdk for adding provision watchers and remove manual code ([#650149e](https://github.com/edgexfoundry/device-onvif-camera/commits/650149e))
  ```text
  BREAKING CHANGE: Remove manual code to add provision watchers and instead use device-sdk to add them
  ```
- Updated sdk versions and secret path to SecretName & Secrets to SecretData ([#5eec883](https://github.com/edgexfoundry/device-onvif-camera/commits/5eec883))
  ```text
  BREAKING CHANGE: Update version to support secret `path` updated to `SecretName` and `Secrets` renamed to `SecretData`
  ```
- Replace internal topics from config with new constants ([#03e3920](https://github.com/edgexfoundry/device-onvif-camera/commits/03e3920))
  ```text
  BREAKING CHANGE: Internal topics no longer configurable, except the base topic.
  ```
- Rework code for refactored MessageBus Configuration ([#916f6b3](https://github.com/edgexfoundry/device-onvif-camera/commits/916f6b3))
  ```text
  BREAKING CHANGE: MessageQueue renamed to MessageBus and fields changed.
  ```
- Rename command line flags for the sake of consistency ([#e283d27](https://github.com/edgexfoundry/device-onvif-camera/commits/e283d27))
  ```text
  BREAKING CHANGE: Renamed -c/--confdir to -cd/--configDir and -f/--file to -cf/--configFile
  ```
- Use latest SDK for flattened config stem ([#4eac92d](https://github.com/edgexfoundry/device-onvif-camera/commits/4eac92d))
  ```text
  BREAKING CHANGE: Location of service configuration in Consul changed
  ```
- **snap:** Update command and metadata sourcing ([#270](https://github.com/edgexfoundry/device-onvif-camera/issues/270)) ([#77ccddf](https://github.com/edgexfoundry/device-onvif-camera/commits/77ccddf))
- **snap:** Refactor and upgrade to edgex-snap-hooks v3 ([#217](https://github.com/edgexfoundry/device-onvif-camera/issues/217)) ([#6741b26](https://github.com/edgexfoundry/device-onvif-camera/commits/6741b26))

### Documentation 📖
- Updated remaining docs from v2 to v3 ([#366](https://github.com/edgexfoundry/device-onvif-camera/issues/366)) ([#4c3943a](https://github.com/edgexfoundry/device-onvif-camera/commits/4c3943a))
- Add operationId, support for deviceCommands, and do not remove schemas for Get commands ([#361](https://github.com/edgexfoundry/device-onvif-camera/issues/361)) ([#0b14d5e](https://github.com/edgexfoundry/device-onvif-camera/commits/0b14d5e))
- Move openapi files to v3 folder ([#359](https://github.com/edgexfoundry/device-onvif-camera/issues/359)) ([#c4faddb](https://github.com/edgexfoundry/device-onvif-camera/commits/c4faddb))
- Update main branch warning with standard text agreed to by TSC ([#354](https://github.com/edgexfoundry/device-onvif-camera/issues/354)) ([#bc8694f](https://github.com/edgexfoundry/device-onvif-camera/commits/bc8694f))
- Remove docs ([#305](https://github.com/edgexfoundry/device-onvif-camera/issues/305)) ([#ca50ff0](https://github.com/edgexfoundry/device-onvif-camera/commits/ca50ff0))
- Remove demo-app docs ([#273](https://github.com/edgexfoundry/device-onvif-camera/issues/273)) ([#30f142b](https://github.com/edgexfoundry/device-onvif-camera/commits/30f142b))
- Add warning to main branch and link to levski ([#271](https://github.com/edgexfoundry/device-onvif-camera/issues/271)) ([#dcb5ff7](https://github.com/edgexfoundry/device-onvif-camera/commits/dcb5ff7))
- Change location of nats documentation ([#237](https://github.com/edgexfoundry/device-onvif-camera/issues/237)) ([#d65f0f9](https://github.com/edgexfoundry/device-onvif-camera/commits/d65f0f9))
- Update docker compose download instructions to the latest version ([#223](https://github.com/edgexfoundry/device-onvif-camera/issues/223)) ([#81966fa](https://github.com/edgexfoundry/device-onvif-camera/commits/81966fa))
- Updating validation metrics with  Hikvision camera ([#218](https://github.com/edgexfoundry/device-onvif-camera/issues/218)) ([#35c6aae](https://github.com/edgexfoundry/device-onvif-camera/commits/35c6aae))

### Build 👷
- Ignore all go-mods except device-sdk-go ([#56b1444](https://github.com/edgexfoundry/device-onvif-camera/commits/56b1444))
- Fixed small issue in makefile ([#209](https://github.com/edgexfoundry/device-onvif-camera/issues/209)) ([#ac64796](https://github.com/edgexfoundry/device-onvif-camera/commits/ac64796))
- Update to Go 1.20, Alpine 3.17 and linter v1.51.2 ([#dd9876e](https://github.com/edgexfoundry/device-onvif-camera/commits/dd9876e))
- Ignore all go-mods except device-sdk-go ([#a4daa83](https://github.com/edgexfoundry/device-onvif-camera/commits/a4daa83))
- **snap:** Upgrade snap base to core22, remove deprecated plug ([#323](https://github.com/edgexfoundry/device-onvif-camera/issues/323)) ([#e2d27ba](https://github.com/edgexfoundry/device-onvif-camera/commits/e2d27ba))

### Continuous Integration 🔄
- Change dependabot gomod schedule to daily ([#264](https://github.com/edgexfoundry/device-onvif-camera/issues/264)) ([#b19be15](https://github.com/edgexfoundry/device-onvif-camera/commits/b19be15))


## [v2.3.1] - 2023-03-22

### Bug Fixes 🐛
- Device-sdk levski patch stable tag added ([#282](https://github.com/edgexfoundry/device-onvif-camera/issues/282)) ([#a19c49b](https://github.com/edgexfoundry/device-onvif-camera/commits/a19c49b))
- Upgrade sdk to fix device cache issue ([#261](https://github.com/edgexfoundry/device-onvif-camera/issues/261)) ([#0c6690d](https://github.com/edgexfoundry/device-onvif-camera/commits/0c6690d))

### Documentation 📖
- Change log updated with levski patch ([#283](https://github.com/edgexfoundry/device-onvif-camera/issues/283)) ([#4d60271](https://github.com/edgexfoundry/device-onvif-camera/commits/4d60271))
- Add note about stable levski branch ([#272](https://github.com/edgexfoundry/device-onvif-camera/issues/272)) ([#d06d196](https://github.com/edgexfoundry/device-onvif-camera/commits/d06d196))
- Remove demo-app docs from levski ([#274](https://github.com/edgexfoundry/device-onvif-camera/issues/274)) ([#c87f573](https://github.com/edgexfoundry/device-onvif-camera/commits/c87f573))


## [v2.3.0] Levski - 2022-11-09 (Not Compatible with 1.x releases)

### Features ✨

- Add Service Metrics configuration ([#187](https://github.com/edgexfoundry/device-onvif-camera/issues/187)) ([#c606ad2](https://github.com/edgexfoundry/device-onvif-camera/commits/c606ad2))
- Add NATS configuration and build option ([#164](https://github.com/edgexfoundry/device-onvif-camera/issues/164)) ([#0803938](https://github.com/edgexfoundry/device-onvif-camera/commits/0803938))
- Add commanding via message configuration ([#4fecb75](https://github.com/edgexfoundry/device-onvif-camera/commits/4fecb75))
- Add GetCapabilities and GetEndpointReference ([#189](https://github.com/edgexfoundry/device-onvif-camera/issues/189)) ([#8e26798](https://github.com/edgexfoundry/device-onvif-camera/commits/8e26798))
- **snap:** add config interface with unique identifier ([#179](https://github.com/edgexfoundry/device-onvif-camera/issues/179)) ([#b5d5716](https://github.com/edgexfoundry/device-onvif-camera/commits/b5d5716))

### Code Refactoring ♻

- Remove control-plane-device ([#153](https://github.com/edgexfoundry/device-onvif-camera/issues/153)) ([#93a2116](https://github.com/edgexfoundry/device-onvif-camera/commits/93a2116))
- Remove custom service interface and use sdk interface ([#150](https://github.com/edgexfoundry/device-onvif-camera/issues/150)) ([#87f15e2](https://github.com/edgexfoundry/device-onvif-camera/commits/87f15e2))

### Test

- Add nestscan discovery tests ([#152](https://github.com/edgexfoundry/device-onvif-camera/issues/152)) ([#d315776](https://github.com/edgexfoundry/device-onvif-camera/commits/d315776))
- Unit test additions ([#149](https://github.com/edgexfoundry/device-onvif-camera/issues/149)) ([#06d0f2a](https://github.com/edgexfoundry/device-onvif-camera/commits/06d0f2a))

### Bug Fixes 🐛

- Status delay [#121](https://github.com/edgexfoundry/device-onvif-camera/issues/121) ([#182](https://github.com/edgexfoundry/device-onvif-camera/issues/182)) ([#55762a2](https://github.com/edgexfoundry/device-onvif-camera/commits/55762a2))
- Discovery mode netscan when using docker-compose ([#176](https://github.com/edgexfoundry/device-onvif-camera/issues/176)) ([#6644798](https://github.com/edgexfoundry/device-onvif-camera/commits/6644798))
- Check whether the onvif response is valid ([#146](https://github.com/edgexfoundry/device-onvif-camera/issues/146)) ([#dc3ccaf](https://github.com/edgexfoundry/device-onvif-camera/commits/dc3ccaf))
- Netscan search only for NetworkVideoTransmitter ([#136](https://github.com/edgexfoundry/device-onvif-camera/issues/136)) ([#a7f4bbd](https://github.com/edgexfoundry/device-onvif-camera/commits/a7f4bbd))

### Documentation 📖

- Integrate onvif schemas and validation matrix to swagger ([#180](https://github.com/edgexfoundry/device-onvif-camera/issues/180)) ([#68a981f](https://github.com/edgexfoundry/device-onvif-camera/commits/68a981f))
- Updated date in onvif function info ([#177](https://github.com/edgexfoundry/device-onvif-camera/issues/177)) ([#c32d12c](https://github.com/edgexfoundry/device-onvif-camera/commits/c32d12c))
- Changes based on validation feeback ([#171](https://github.com/edgexfoundry/device-onvif-camera/issues/171)) ([#bca8455](https://github.com/edgexfoundry/device-onvif-camera/commits/bca8455))
- Add video reference and text ([#166](https://github.com/edgexfoundry/device-onvif-camera/issues/166)) ([#eff61ec](https://github.com/edgexfoundry/device-onvif-camera/commits/eff61ec))
- Add Go 1.18+ required for native install ([#60be3eb](https://github.com/edgexfoundry/device-onvif-camera/commits/60be3eb))
- Fix yaml tag for multicast ([#143](https://github.com/edgexfoundry/device-onvif-camera/issues/143)) ([#51d17d6](https://github.com/edgexfoundry/device-onvif-camera/commits/51d17d6))
- Add missing param to postman call ([#2f0c3ce](https://github.com/edgexfoundry/device-onvif-camera/commits/2f0c3ce))
- Add openapi spec from updated postman collection ([#123](https://github.com/edgexfoundry/device-onvif-camera/issues/123)) ([#9690a38](https://github.com/edgexfoundry/device-onvif-camera/commits/9690a38))
- Publish swagger files, update swagger readme ([#130](https://github.com/edgexfoundry/device-onvif-camera/issues/130)) ([#59a7980](https://github.com/edgexfoundry/device-onvif-camera/commits/59a7980))
- In-depth discovery and credential guides ([#126](https://github.com/edgexfoundry/device-onvif-camera/issues/126)) ([#803b45c](https://github.com/edgexfoundry/device-onvif-camera/commits/803b45c))

### Build 👷

- Upgrade to Go 1.18, Alpine 3.16, linter version and latest SDK/go-mod versions ([#127](https://github.com/edgexfoundry/device-onvif-camera/issues/127)) ([#616f9b7](https://github.com/edgexfoundry/device-onvif-camera/commits/616f9b7))


## [v2.2.0] Kamakura - 2022-07-26 (Not Compatible with 1.x releases)

This is the initial release of this ONVIF camera device service. Refer to the [README](https://github.com/edgexfoundry/device-onvif-camera/blob/v2.2.0/README.md) for details about this service.

### Known Issues 

The following issues are known at the time of this initial release and will be addressed in a future release:

-  [#121](https://github.com/edgexfoundry/device-onvif-camera/issues/121) Delay in obtaining DeviceStatus for pre-defined and manually added devices
