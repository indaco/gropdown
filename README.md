<h1 align="center" style="font-size: 2.5rem;">
  gropdown
</h1>
<p align="center">
  <a href="https://github.com/indaco/gropdown/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square&logo=none" alt="license" />
  </a>
  &nbsp;
  <a href="https://goreportcard.com/report/github.com/indaco/gropdown/" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/gropdown" alt="go report card" />
  </a>
  &nbsp;
  <a href="https://pkg.go.dev/github.com/indaco/gropdown/" target="_blank">
    <img src="https://pkg.go.dev/badge/github.com/indaco/gropdown/.svg" alt="go reference" />
  </a>
  <a href="https://www.jetify.com/devbox/docs/contributor-quickstart/">
    <img  src="https://www.jetify.com/img/devbox/shield_moon.svg" alt="Built with Devbox" />
  </a>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#api-reference">API Reference</a> •
  <a href="#accessibility-a11y">Accessibility (A11Y)</a> •
  <a href="#theming">Theming</a> •
  <a href="#examples">Examples</a>
</p>

A fully accessible, configurable and themeable server-rendered dropdown component for Go web applications. Built with [templ](https://github.com/a-h/templ) library for seamless integration with Go-based web frontends.

![gropdown](https://raw.githubusercontent.com/indaco/gh-assets/main/gropdown/demo.gif)

## Features

- **Accessible**: Fully compliant with the [WAI-ARIA Menu Button Design Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/menu-button/), to ensure accessibility for all users.
- **No External Dependencies**: Built with native Go and the `templ` library, requiring no external dependencies.
- **Configurable**: The component offers various configuration options to customize its behavior (e.g. positioning, open by default...)
- **Multiple dropdown**: multiple `gropdown` components on the same page, each with its own style.
- **Themeable**: Supports theming via CSS variables, allowing easy customization of appearance. Comes with built-in support for light and dark modes, as well as the ability to define custom themes using the `data-theme` attribute.
- **Versatile**: Items can be buttons or links (`<a>`). When a link item is marked as _external_, a visual icon will be added to indicate it.

## Installation

To install the Dropdown module, use the `go get` command:

```sh
go get github.com/indaco/gropdown
```

Ensure your project is using Go Modules (it will have a `go.mod` file in its root if it already does).

## Usage

> [!TIP]
> refer to the [Examples](#examples) section.

Import the Dropdown module into your project:

```go
import "github.com/indaco/gropdown"
```

### Configuration & Context

```go
// Default options
dropdownConfig := gropdown.NewConfigBuilder().Build()
```

#### Available Options

Users can access each configuration option using the corresponding `With` method, such as `gropdown.WithOpen(true)` or `gropdown.WithPosition(gropdown.Left)`.

| Option                | Type          | Default  | Description                                                                                                           |
|:----------------------|:--------------|:---------|:----------------------------------------------------------------------------------------------------------------------|
| `Open`                | _bool_        | `false`  | indicates whether the dropdown menu is currently open.                                                                |
| `Placement`           | [_Placement_] | `Bottom` | indicates the position of the dropdown content relative to the button. Options: `Top`, `Bottom`, `Left`, and `Right`. |
| `Animation`           | _bool_        | `true`   | indicates whether the dropdown button should use animations on open and close.                                        |
| `CloseOnOutsideClick` | _bool_        | `true`   | indicates whether the dropdown should be closed when a click occurs outside of it.                                    |

To allow configurations to be accessible at any level in the dropdown hierarchy and make each component on the same page function independently, `gropdown` makes use of the templ component context and the implicit `ctx` variable. You can read more about templ component context [here](https://templ.guide/syntax-and-usage/context#using-context).

In your function handler, create a configuration for `gropdown` and attach it to the request context passed into the handler function.

```go
func HandleHome(w http.ResponseWriter, r *http.Request) {
  dropdownConfig := gropdown.NewConfigBuilder().WithPlacement(gropdown.Top).Build()

  configMap := gropdown.NewConfigMap()
  configMap.Add("demo", dropdownConfig)
  ctx := context.WithValue(r.Context(), gropdown.ConfigContextKey, configMap)
  err := Page().Render(ctx, w)
  if err != nil {
    return
  }
}
```

### Dropdown Component Structure - Markup

In your `templ` file, defines the structure for the dropdown component.

> [!IMPORTANT]
> It is crucial to ensure that the value passed to `gropdown.Root` **matches** the one used when adding the `gropdownConfig` to the `configMap` as per step above. This ensures that multiple dropdowns on the same page function independently.

```go
// Set the button label.
@gropdown.Root("demo") {
  @gropdown.Button("Menu")
  @gropdown.Content() {
    @gropdown.Item("Profile",
      gropdown.ItemOptions{
        Href: "/profile",
        Icon: profileIcon,
      },
    )
    @gropdown.Item("Settings",
      gropdown.ItemOptions{
        Href: "/settings",
        Icon: settingsIcon,
      },
    )
    @gropdown.Divider()
    @gropdown.Item("GitHub",
      gropdown.ItemOptions{
        Href:     "https://github.com",
        External: true,
        Icon:     globeIcon,
      },
    )
    @gropdown.Divider()
    @gropdown.Item("Button",
      gropdown.ItemOptions{
        Icon:  clickIcon,
        Attrs: templ.Attributes{"onclick": "alert('Hello gropdown');"},
      },
    )
  }
}
```

### CSS and Javascript

`gropdown` leverages the `templ` library's features, including CSS Components and JavaScript Templates, to encapsulate all necessary styling and functionality without relying on external dependencies.

- `GropdownCSS`: it supplies the required CSS, encapsulating the visual design and layout specifics of the component.
- `GropdownJS`: it provides the JavaScript logic essential for dynamic behaviors such as displaying, keyboard navigation and interaction with the component.

To facilitate integration with Go's `template/html` standard library, `gropdown` includes a dedicated `HTMLGenerator` type to seamlessly integrate the component into web applications built with Go's `html/template standard` library.

There are methods acting as wrappers to the templ's `templ.ToGoHTML`, generate the necessary HTML to be embedded them into server-rendered pages:

- `GropdownCSSToGoHTML`: render the `GropdownCSS` component into a `template.HTML` value.
- `GropdownJSToGoHTML`: render the `GropdownJS` component into a `template.HTML` value.

## API Reference

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">gropdown.<span style="font-weight: 600; color: #1f2937;">Root</span></h3>

| Property | Type     | Description                                            |
|:---------|:---------|:-------------------------------------------------------|
| `id`     | _string_ | The unique identifier for the dropdown menu component. |

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">gropdown.<span style="font-weight: 600; color: #1f2937;">Button</span></h3>

| Property | Type           | Description                                          |
|:---------|:---------------|:-----------------------------------------------------|
| `label`  | _string_       | The text displayed for the dropdown menu button.     |
| `icon`   | [_ButtonIcon_] | The icon displayed next to the dropdown menu button. |

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">gropdown.<span style="font-weight: 600; color: #1f2937;">Content</span></h3>

None

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">gropdown.<span style="font-weight: 600; color: #1f2937;">Item</span></h3>

| Property | Type            | Description                                                          |
|:---------|:----------------|:---------------------------------------------------------------------|
| `label`  | _string_        | The text displayed for the dropdown menu item.                       |
| `opts`   | [_ItemOptions_] | The options for configuring the behavior and appearance of the item. |

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">gropdown.<span style="font-weight: 600; color: #1f2937;">Divider</span></h3>

None

## Accessibility (A11Y)

The dropdown component is designed to be accessible to screen readers and supports keyboard navigation according to the [WAI-ARIA pattern for menu buttons](https://www.w3.org/WAI/ARIA/apg/patterns/menu-button/examples/menu-button-actions/#kbd_label).

### Screen Reader Support

Ensure that proper ARIA attributes are used to convey the state and role of the dropdown elements.

### Keyboard Interaction

- **Focusing the Dropdown**:
  - Use the `Tab` key to navigate to the dropdown button. Pressing `Enter` or `Space` will open the dropdown menu.
- **Navigating within the Dropdown Menu**:
  - Use the `Arrow` keys to move focus between items within the dropdown menu.
  - Pressing `Home` or `End` keys will move focus to the first or last item respectively.
  - Use `A-Z` or `a-z` keys to move focus to the next menu item with a label that starts with the typed character if such a menu item exists. Otherwise, focus does not move.
- **Selecting an Item**:
  - Press `Enter` to select the currently focused item in the dropdown menu.
- **Closing the Dropdown**:
  - Press `Escape` to close the dropdown and sets focus to the menu button.

## Theming

Dropdown is themeable using CSS variables (prefix `gdd`) to customize the appearance according to your design.

By default, it supports both light and dark modes. In addition to the built-in modes, you can define your
own custom themes using the `data-theme` attribute. Simply add a `data-theme` attribute to the root element of your application
and define the corresponding CSS variables for your custom theme.

For a comprehensive list of CSS custom properties, along with their default values and descriptions, please consult the `gropdown` [CSS custom Props](./docs/css-props.md) document.

## Examples

- [with `a-h/templ`](_examples/a-h-templ)
- [custom animations](_examples/custom-animations)
- [custom button icon](_examples/custom-button-icon)
- [with `template/html`](_examples/go-html-template)
- [icon only button](_examples/icon-only-button)
- [multiple dropdowns](_examples/multiple-dropdowns/)
- [positioning](_examples/positioning/)
- [theming](_examples/theming)

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

### Development Environment Setup

To set up a development environment for this repository, you can use [devbox](https://www.jetify.com/devbox) along with the provided `devbox.json` configuration file.

1. Install devbox by following the instructions in the [devbox documentation](https://www.jetify.com/devbox/docs/installing_devbox/).
2. Clone this repository to your local machine.
3. Navigate to the root directory of the cloned repository.
4. Run `devbox install` to install all packages mentioned in the `devbox.json` file.
5. Run `devbox shell` to start a new shell with access to the environment.
6. Once the devbox environment is set up, you can start developing, testing, and contributing to the repository.

### Using the Makefile

Additionally, you can make use of the provided `Makefile` to run various tasks:

```bash
make build       # The main build target
make examples    # Process templ files in the _examples folder
make templ       # Process TEMPL files
make test        # Run go tests
make help        # Print this help message
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

<!-- Resources -->
[_ItemOptions_]: ./types.go#L22-L27
[_ButtonIcon_]: ./types.go#L30-L33
[_Placement_]: ./constants.go#L8-L11
