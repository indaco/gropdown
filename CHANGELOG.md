# Changelog

All notable changes to this project will be documented in this file.

The format adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html),
and is generated by [changelogen](https://github.com/unjs/changelogen) and managed with [Changie](https://github.com/miniscruff/changie).

## v0.1.2 - 2025-04-02

[compare changes](https://github.com/indaco/gropdown/compare/v0.1.1...v0.1.2)

**No functional or API-level changes are introduced in this release.**

The main improvement in this release is a packaging optimization:

- The demo.gif file previously embedded in the main branch has been moved to a separate repos.
- This change ensures that users installing the module no longer download the demo asset, resulting in reduced disk usage and faster installs.

### 📦 Build

- Bump go version and update deps ([8c18ac1](https://github.com/indaco/gropdown/commit/8c18ac1))

### 🏡 Chore

- Regenerate _templ.go files with templ v0.3.857 ([a402ae9](https://github.com/indaco/gropdown/commit/a402ae9))

### 📖 Documentation

- **README.md:** Use same key in the configuration section ([3bd4226](https://github.com/indaco/gropdown/commit/3bd4226))
- **README:** Move preview assets to gh-assets and update image links ([af4a9ba](https://github.com/indaco/gropdown/commit/af4a9ba))
- **README:** Improve feature list and refine project description ([20cc464](https://github.com/indaco/gropdown/commit/20cc464))

### ❤️ Contributors

- Indaco ([@indaco](https://github.com/indaco))

## v0.1.1 - 2024-05-07

[compare changes](https://github.com/indaco/gropdown/compare/v0.1.0...v0.1.1)

### 🩹 Fixes

- **gropdown-content.templ:** Id and aria prop (460c438)

### 🏡 Chore

- Bump templ from v0.2.663 to v0.2.680 (739111b)
- Regenerate templ files with v0.2.680 (a5a975c)

### ❤️ Contributors

- Indaco <github@mircoveltri.me>

## v0.1.0 - 2024-05-06

### 🚀 Enhancements

Initial Release
