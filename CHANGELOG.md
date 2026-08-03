# Changelog

## [1.1.1](https://github.com/s0cks/sp-interview-prep-01/compare/1.1.0...v1.1.1) (2026-08-03)


### Bug Fixes

* **release-please:** 🐛 lets try fixing the release-please docker step ([7ead643](https://github.com/s0cks/sp-interview-prep-01/commit/7ead643e9e6492251570aafb2cb037bf8c59efca))

## [1.1.0](https://github.com/s0cks/sp-interview-prep-01/compare/v1.0.0...v1.1.0) (2026-08-03)


### Features

* ✨ first commit ([55744ce](https://github.com/s0cks/sp-interview-prep-01/commit/55744ce56ad1dff9a777bb94e6f99fb74a88bdaa)), closes [#1](https://github.com/s0cks/sp-interview-prep-01/issues/1)
* **buffer:** ✨ add a function to SensorBuffer to calc metrics ([303da90](https://github.com/s0cks/sp-interview-prep-01/commit/303da901fcea7709c8030692d7850b008dd91beb)), closes [#12](https://github.com/s0cks/sp-interview-prep-01/issues/12) [#13](https://github.com/s0cks/sp-interview-prep-01/issues/13)
* **buffer:** ✨ add initial SensorBuffer implementation ([12cc20d](https://github.com/s0cks/sp-interview-prep-01/commit/12cc20dbb83821817104c3c488b8b686e8386d94))
* **buffer:** ✨ add sync.RWMutex to SensorBuffer struct ([99af46e](https://github.com/s0cks/sp-interview-prep-01/commit/99af46e18847c0d07ec971aaceae1901f23a8cce))
* **sensor-mock:** ✨ add a mock sensor ([f977c25](https://github.com/s0cks/sp-interview-prep-01/commit/f977c25edc3c1791a6dcd6775991094268c0439b))
* **sensor-mock:** ✨ create docker image for mock sensor ([e3f7305](https://github.com/s0cks/sp-interview-prep-01/commit/e3f7305911f5aeaf48778fb5b42c308fa3b63178))
* **sensor-web:** ✨ add a basic sensor dashboard  ([917af09](https://github.com/s0cks/sp-interview-prep-01/commit/917af090bf7b9921c6f0c6a2ad42bee4a2089bf6))
* **sensor-web:** ✨ add docker image for web dashboard ([f7246cd](https://github.com/s0cks/sp-interview-prep-01/commit/f7246cd625f9f03bc65cdd975ea7d9ce134af1ed))
* **sensor-web:** ✨ create a basic web client structure ([e091237](https://github.com/s0cks/sp-interview-prep-01/commit/e091237e1ff1c3e9f51c2280d36cf43562b6fb2a))
* **service:** ✨ add GET /sensors/:id endpoint ([39b9a5b](https://github.com/s0cks/sp-interview-prep-01/commit/39b9a5b00d733119454487b5ad765539981d821a))
* **service:** ✨ add GET /sensors/:id/:metric endpoint ([28e204f](https://github.com/s0cks/sp-interview-prep-01/commit/28e204f6565dfca7f8137c9b3a1c4e3c87b6878c))
* **service:** ✨ add POST /sensors/:id endpoint ([085c842](https://github.com/s0cks/sp-interview-prep-01/commit/085c842dc4db6f17bd4ffdd59a17b77de0c6ec5c)), closes [#5](https://github.com/s0cks/sp-interview-prep-01/issues/5)
* **service:** ✨ Create initial SensorService implementation ([dd59ba2](https://github.com/s0cks/sp-interview-prep-01/commit/dd59ba240db95ec333da854b7c0a03948fee4355))


### Bug Fixes

* more release-please fixes? ([a628670](https://github.com/s0cks/sp-interview-prep-01/commit/a628670c79b2f32911c881b9c9aa12c0a8dce177))
* **release-please:** working on fixing ([481580c](https://github.com/s0cks/sp-interview-prep-01/commit/481580c1c9af024c6357798b62fcded7152acb89))
* trying to get release-please working still ([7a851c7](https://github.com/s0cks/sp-interview-prep-01/commit/7a851c724d0a8847850717ae6908039c36042f44))


### Reverts

* **release-please:** ⏪ reset release-please version for now ([72bd724](https://github.com/s0cks/sp-interview-prep-01/commit/72bd724faaf1b7d85f9203e3071bbd9be5fdd44f))

## [1.2.0](https://github.com/s0cks/sp-interview-prep-01/compare/v1.1.2...v1.2.0) (2026-08-01)


### Features

* ✨ first commit ([55744ce](https://github.com/s0cks/sp-interview-prep-01/commit/55744ce56ad1dff9a777bb94e6f99fb74a88bdaa)), closes [#1](https://github.com/s0cks/sp-interview-prep-01/issues/1)
* **buffer:** ✨ add a function to SensorBuffer to calc metrics ([303da90](https://github.com/s0cks/sp-interview-prep-01/commit/303da901fcea7709c8030692d7850b008dd91beb)), closes [#12](https://github.com/s0cks/sp-interview-prep-01/issues/12) [#13](https://github.com/s0cks/sp-interview-prep-01/issues/13)
* **buffer:** ✨ add initial SensorBuffer implementation ([12cc20d](https://github.com/s0cks/sp-interview-prep-01/commit/12cc20dbb83821817104c3c488b8b686e8386d94))
* **buffer:** ✨ add sync.RWMutex to SensorBuffer struct ([99af46e](https://github.com/s0cks/sp-interview-prep-01/commit/99af46e18847c0d07ec971aaceae1901f23a8cce))
* **service:** ✨ add GET /sensors/:id endpoint ([39b9a5b](https://github.com/s0cks/sp-interview-prep-01/commit/39b9a5b00d733119454487b5ad765539981d821a))
* **service:** ✨ add GET /sensors/:id/:metric endpoint ([28e204f](https://github.com/s0cks/sp-interview-prep-01/commit/28e204f6565dfca7f8137c9b3a1c4e3c87b6878c))
* **service:** ✨ add POST /sensors/:id endpoint ([085c842](https://github.com/s0cks/sp-interview-prep-01/commit/085c842dc4db6f17bd4ffdd59a17b77de0c6ec5c)), closes [#5](https://github.com/s0cks/sp-interview-prep-01/issues/5)
* **service:** ✨ Create initial SensorService implementation ([dd59ba2](https://github.com/s0cks/sp-interview-prep-01/commit/dd59ba240db95ec333da854b7c0a03948fee4355))


### Bug Fixes

* more release-please fixes? ([a628670](https://github.com/s0cks/sp-interview-prep-01/commit/a628670c79b2f32911c881b9c9aa12c0a8dce177))
* **release-please:** working on fixing ([481580c](https://github.com/s0cks/sp-interview-prep-01/commit/481580c1c9af024c6357798b62fcded7152acb89))
* trying to get release-please working still ([7a851c7](https://github.com/s0cks/sp-interview-prep-01/commit/7a851c724d0a8847850717ae6908039c36042f44))

## [1.1.2](https://github.com/s0cks/sp-interview-prep-01/compare/sensor-service-v1.1.1...sensor-service-v1.1.2) (2026-08-01)


### Bug Fixes

* trying to get release-please working still ([7a851c7](https://github.com/s0cks/sp-interview-prep-01/commit/7a851c724d0a8847850717ae6908039c36042f44))

## [1.1.1](https://github.com/s0cks/sp-interview-prep-01/compare/sensor-service-v1.1.0...sensor-service-v1.1.1) (2026-08-01)


### Bug Fixes

* **release-please:** working on fixing ([481580c](https://github.com/s0cks/sp-interview-prep-01/commit/481580c1c9af024c6357798b62fcded7152acb89))

## [1.1.0](https://github.com/s0cks/sp-interview-prep-01/compare/sensor-service-v1.0.0...sensor-service-v1.1.0) (2026-08-01)


### Features

* ✨ first commit ([55744ce](https://github.com/s0cks/sp-interview-prep-01/commit/55744ce56ad1dff9a777bb94e6f99fb74a88bdaa)), closes [#1](https://github.com/s0cks/sp-interview-prep-01/issues/1)
* **buffer:** ✨ add a function to SensorBuffer to calc metrics ([303da90](https://github.com/s0cks/sp-interview-prep-01/commit/303da901fcea7709c8030692d7850b008dd91beb)), closes [#12](https://github.com/s0cks/sp-interview-prep-01/issues/12) [#13](https://github.com/s0cks/sp-interview-prep-01/issues/13)
* **buffer:** ✨ add initial SensorBuffer implementation ([12cc20d](https://github.com/s0cks/sp-interview-prep-01/commit/12cc20dbb83821817104c3c488b8b686e8386d94))
* **buffer:** ✨ add sync.RWMutex to SensorBuffer struct ([99af46e](https://github.com/s0cks/sp-interview-prep-01/commit/99af46e18847c0d07ec971aaceae1901f23a8cce))
* **service:** ✨ add GET /sensors/:id endpoint ([39b9a5b](https://github.com/s0cks/sp-interview-prep-01/commit/39b9a5b00d733119454487b5ad765539981d821a))
* **service:** ✨ add GET /sensors/:id/:metric endpoint ([28e204f](https://github.com/s0cks/sp-interview-prep-01/commit/28e204f6565dfca7f8137c9b3a1c4e3c87b6878c))
* **service:** ✨ add POST /sensors/:id endpoint ([085c842](https://github.com/s0cks/sp-interview-prep-01/commit/085c842dc4db6f17bd4ffdd59a17b77de0c6ec5c)), closes [#5](https://github.com/s0cks/sp-interview-prep-01/issues/5)
* **service:** ✨ Create initial SensorService implementation ([dd59ba2](https://github.com/s0cks/sp-interview-prep-01/commit/dd59ba240db95ec333da854b7c0a03948fee4355))
