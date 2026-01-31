# Changelog

## 0.2.0 (2026-01-31)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/beeper/desktop-api-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** add `description` field to chats, make `title` optional ([fbf0470](https://github.com/beeper/desktop-api-go/commit/fbf047029f11c57542963ed3446ab191ae738643))
* **api:** add upload asset and edit message endpoints ([6d37bf4](https://github.com/beeper/desktop-api-go/commit/6d37bf45aeaeef173628d7bd16ecb5fcfe49ca1e))
* **api:** manual updates ([d52c6cb](https://github.com/beeper/desktop-api-go/commit/d52c6cbcfa8217c6fde16793dac7ed98d2c61348))
* **api:** manual updates ([782c77d](https://github.com/beeper/desktop-api-go/commit/782c77d0de3e076765031d3b311c81f55a82b02b))
* **api:** remove mcp for now ([0cf432f](https://github.com/beeper/desktop-api-go/commit/0cf432f5b3aa26727a571becd287501776a58726))
* **client:** add a convenient param.SetJSON helper ([9ea70da](https://github.com/beeper/desktop-api-go/commit/9ea70da2a4b22a60440ca738ba6367037de2b63b))
* **encoder:** support bracket encoding form-data object members ([a338742](https://github.com/beeper/desktop-api-go/commit/a33874294b78dc962f1acf74668ab59697d61711))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([2bcd6c1](https://github.com/beeper/desktop-api-go/commit/2bcd6c1aaaf2b7e6e554d8a87743d2593856f678))
* **docs:** add missing pointer prefix to api.md return types ([9fd334b](https://github.com/beeper/desktop-api-go/commit/9fd334be8ea4e83d5414c7d044cd326643137685))
* fix for namespace collisions with client and resource test methods ([7454269](https://github.com/beeper/desktop-api-go/commit/7454269a9eb54a359ce5efdfe31c4f036251707a))
* **mcp:** correct code tool API endpoint ([568939e](https://github.com/beeper/desktop-api-go/commit/568939eff6ad5bcdcd42d7bae3439af3389e4b8a))
* rename param to avoid collision ([d304cbe](https://github.com/beeper/desktop-api-go/commit/d304cbe9e74f9486567156cb28a60cf7cee5b053))


### Chores

* add float64 to valid types for RegisterFieldValidator ([4658d5c](https://github.com/beeper/desktop-api-go/commit/4658d5c7d11ff703a1b7e5ccfb594b22f2ee5144))
* bump gjson version ([1194594](https://github.com/beeper/desktop-api-go/commit/1194594b4e0744f59f66e0cda1929e75eb8973a2))
* configure new SDK language ([ebb892c](https://github.com/beeper/desktop-api-go/commit/ebb892cd8a5b4bad877c0b88724a9a55f43dac94))
* elide duplicate aliases ([fa3d456](https://github.com/beeper/desktop-api-go/commit/fa3d4564c46606bcb63fd98906eb0687758ac059))
* **internal:** codegen related update ([f2d3a30](https://github.com/beeper/desktop-api-go/commit/f2d3a30447ffaae00b456f610aba7afee08d32f0))
* **internal:** codegen related update ([77a6d60](https://github.com/beeper/desktop-api-go/commit/77a6d605f55019d31fba8ed59ddc5232f527f9ff))
* **internal:** grammar fix (it's -&gt; its) ([60171f3](https://github.com/beeper/desktop-api-go/commit/60171f3c6eed029a5689bc8ac58656baf63d88f3))
* **internal:** update `actions/checkout` version ([1f84ff0](https://github.com/beeper/desktop-api-go/commit/1f84ff0e78592fb0b97fb1185dbe049a298af53d))


### Documentation

* prominently feature MCP server setup in root SDK readmes ([fcad6f3](https://github.com/beeper/desktop-api-go/commit/fcad6f33198927cfd1e144648a38fb43b1d43f06))

## 0.1.0 (2025-10-16)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/beeper/desktop-api-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** bump for new endpoints ([26fc4bd](https://github.com/beeper/desktop-api-go/commit/26fc4bd1ef7b297fc07cb7b10a63ff33713cb8c1))
* **api:** manual updates ([78129af](https://github.com/beeper/desktop-api-go/commit/78129af51edc7883b0e41aefd57b6adc4e605f5a))
* **api:** manual updates ([4408f20](https://github.com/beeper/desktop-api-go/commit/4408f2032cc58ab59a465bddf91ebe5e8d248cad))
* **api:** manual updates ([0c8017f](https://github.com/beeper/desktop-api-go/commit/0c8017fad7c347bcf437cf685aafdb4270c44a84))
* **api:** manual updates ([db71efa](https://github.com/beeper/desktop-api-go/commit/db71efac468bc6f4b608997bbf0721082733c36b))
* **api:** manual updates ([589a68b](https://github.com/beeper/desktop-api-go/commit/589a68bbf8a2d60a1281db464df43ad9eab67668))
* **api:** manual updates ([39aa081](https://github.com/beeper/desktop-api-go/commit/39aa081c0cf963d834dc7b2ae2e9a6fe38c4cec5))
* **api:** manual updates ([4c76c3b](https://github.com/beeper/desktop-api-go/commit/4c76c3b0b3a134888001bce5e150db1971ac8f14))
* **api:** manual updates ([5c2b75b](https://github.com/beeper/desktop-api-go/commit/5c2b75bdbc04d7548db3b8f5a202ade4590f7143))
* **api:** manual updates ([658f1d0](https://github.com/beeper/desktop-api-go/commit/658f1d09e48180e35248adc7da6e27c395d1fcca))
* **api:** manual updates ([d681504](https://github.com/beeper/desktop-api-go/commit/d681504bf0ff37084a75d5f9a64c4bc48733fcff))
* **api:** remove limit from list routes ([92c49b3](https://github.com/beeper/desktop-api-go/commit/92c49b369b5ace61e67a1f65c4f44c32297bbb5a))
* **api:** update via SDK Studio ([e6707f7](https://github.com/beeper/desktop-api-go/commit/e6707f7a74894867f206865d4d2a36ae00372917))
* **api:** update via SDK Studio ([36be378](https://github.com/beeper/desktop-api-go/commit/36be3780fb98ff95cecde1c4ecdd7d45cdb4aec1))
* **api:** update via SDK Studio ([f32d341](https://github.com/beeper/desktop-api-go/commit/f32d3418ec8b806313a58dac62e4de26c2894a17))
* **api:** update via SDK Studio ([f390e62](https://github.com/beeper/desktop-api-go/commit/f390e629ec0230972d8f84fbffe27ea6a56e92ee))
* **api:** update via SDK Studio ([278cdb3](https://github.com/beeper/desktop-api-go/commit/278cdb321d3459c4019254ea75e5705bd70597b1))
* **api:** update via SDK Studio ([25b5788](https://github.com/beeper/desktop-api-go/commit/25b578877564808fdd3eddb5a4a7238c6b249bc7))
* **api:** update via SDK Studio ([b5a1ff4](https://github.com/beeper/desktop-api-go/commit/b5a1ff452e59274b58586d826e5ebd34efe37a0a))
* **api:** update via SDK Studio ([3b0da67](https://github.com/beeper/desktop-api-go/commit/3b0da676abda606cbb2d58f5dbc558dfbebf16a6))
* **api:** update via SDK Studio ([95da01f](https://github.com/beeper/desktop-api-go/commit/95da01f346d6483cce93fafe9b0ba0a59c50c0a3))
* **api:** update via SDK Studio ([8efb1dd](https://github.com/beeper/desktop-api-go/commit/8efb1ddb9a3c01b51c69e04d16f3e40fcbfebbdb))
* **api:** update via SDK Studio ([cdee3ee](https://github.com/beeper/desktop-api-go/commit/cdee3ee8a2f49c9cbd6a3fd1948b7576b6c5d645))
* **api:** update via SDK Studio ([7cafbfd](https://github.com/beeper/desktop-api-go/commit/7cafbfdcbd5d5aae479a6b665ec5578586dbc215))
* **api:** update via SDK Studio ([9f84441](https://github.com/beeper/desktop-api-go/commit/9f84441fa443bd5f02323eca7676de4d960924e0))
* **api:** update via SDK Studio ([0a97aff](https://github.com/beeper/desktop-api-go/commit/0a97aff2981e3f611cfd829353868f3fb40633d9))


### Bug Fixes

* close body before retrying ([c6145c4](https://github.com/beeper/desktop-api-go/commit/c6145c4b7e1880a13319137608951f0049d94c25))


### Chores

* configure new SDK language ([7f956b3](https://github.com/beeper/desktop-api-go/commit/7f956b3cf69545aed9b0147efa467636fe2ee43b))
* configure new SDK language ([310166b](https://github.com/beeper/desktop-api-go/commit/310166b0b45746db244d1a9e34de1c8dcc04d316))
* **internal:** codegen related update ([b891408](https://github.com/beeper/desktop-api-go/commit/b891408e549e98abd2106329a3c362830639690b))
* **internal:** codegen related update ([3c735fe](https://github.com/beeper/desktop-api-go/commit/3c735fe37a02273620e047afa8e7a28f272fdbd2))
