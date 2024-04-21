<h1 align="center" style="font-size: 2.5rem;">
  gropdown
</h1>
<h2 align="center" style="font-size: 1.5rem;">
A fully accessible, configurable and themeable server-rendered dropdown component for Go web applications.
</h2>
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
    &nbsp;
    <a href="https://jetpack.io/devbox/docs/contributor-quickstart/" target="_blank">
      <img
          src="https://jetpack.io/img/devbox/shield_galaxy.svg"
          alt="Built with Devbox"
      />
    </a>
</p>

Built with [templ](https://github.com/a-h/templ) library for seamless integration with Go-based web frontends.

## Features

- **Accessible**: Fully compliant with the [WAI-ARIA Menu Button Design Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/menu-button/), to ensure accessibility for all users.
- **No External Dependencies**: Built with native Go and the `templ` library, requiring no external dependencies.
- **Configurable**: The component offers various configuration options to customize its behavior (e.g. positioning, open by default...)
- **Themeable**: Supports theming via CSS variables, allowing easy customization of appearance. Comes with built-in support for light and dark modes, as well as the ability to define custom themes using the `data-theme` attribute.
- **Versatile**: Items can be buttons or links (`<a>`). When a link item is marked as _external_, a visual icon will be added to indicate it.

<div style="display: flex; justify-content: center;">
   <img src="statics/demo.gif" alt="Image" >
</div>

## Installation

To install the Dropdown module, use the `go get` command:

```sh
go get github.com/indaco/gropdown
```

Ensure your project is using Go Modules (it will have a go.mod file in its root if it already does).

## Usage

Import the Dropdown module into your project:

```go
import "github.com/indaco/gropdown"
```

### Creating a Dropdown

```go
// Set the button label.
button := gropdown.DropdownButton{Label: "Menu"}

// Set the items for the Dropdown menu.
items := []gropdown.DropdownItem{
    {Label: "Settings", Href: "/settings"},
    {Label: "GitHub", Href: "https://github.com", External: true},
    {Divider: true},
    {Label: "Button", Attrs: templ.Attributes{"onclick": "alert('Hello gropdown');"}},
}

// Build the Dropdown component.
dropdown := gropdown.NewDropdownBuilder().WithButton(button).WithItems(items)
```

or customize the dropdown with options:

```go
// Here we set the Dropdown menu opened as default and the content positioned as absolute instead of relative.
dropdown := gropdown.NewDropdownBuilder().SetOpen(true).SetPositionAbsolute(true)
// Here we set the Dropdown menu opened as default and the content positioned on top.
dropdown := gropdown.NewDropdownBuilder().SetOpen(true).SetPosition(gropdown.Top)
dropdown.WithButton(button).WithItems(items)
```

### Add gropdown CSS and Javascript

`gropdown` leverages the `templ` library's features, including CSS Components and JavaScript Templates, to encapsulate all necessary styling and functionality without relying on external dependencies.

- `GropdownCSS`: it supplies the required CSS, encapsulating the visual design and layout specifics of the component.
- `GropdownJS`: it provides the JavaScript logic essential for dynamic behaviors such as displaying, keyboard navigation and interaction with the component.

To facilitate integration with Go's `template/html` standard library, `gropdown` includes a dedicated HTMLGenerator type to seamlessly integrate the component into web applications built with Go's `html/template standard` library.

There are methods acting as wrappers to the templ's `templ.ToGoHTML`, generate the necessary HTML to be embedded them into server-rendered pages:

- `GropdownCSSToGoHTML`: render the `GropdownCSS` component into a `template.HTML` value.
- `GropdownJSToGoHTML`: render the `GropdownJS` component into a `template.HTML` value.

> Note: refer to the [Examples](#examples) section to see how to use `gropdown` with templ and `html/template`.

## A11Y

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

Here below is the list of all CSS variables defined and their default values:

### Dropdown Button

| CSS Variable                              | Default Value                                                          | Description                                                      |
|-------------------------------------------|------------------------------------------------------------------------|------------------------------------------------------------------|
| `--gdd-button-min-w`                      | 4.5em                                                                  | Minimum width of the dropdown button                             |
| `--gdd-button-py`                         | 1ch                                                                    | Padding on the y-axis of the dropdown button                     |
| `--gdd-button-px`                         | 2ch                                                                    | Padding on the x-axis of the dropdown button                     |
| `--gdd-button-icon-space`                 | 0.5ch                                                                  | Space between the dropdown button label and icon                 |
| `--gdd-button-color`                      | ![Color Preview](https://via.placeholder.com/20/1f2937?text=+) #1f2937 | Text color of the dropdown button                                |
| `--gdd-button-color-hover`                | ![Color Preview](https://via.placeholder.com/20/1f2937?text=+) #1f2937 | Text color of the dropdown button on hover                       |
| `--gdd-button-font-size`                  | 1rem                                                                   | Font size of the dropdown button label                           |
| `--gdd-button-font-family`                | inherit                                                                | Font family of the dropdown button label                         |
| `--gdd-button-font-weight`                | 500                                                                    | Font weight of the dropdown button label                         |
| `--gdd-button-line-height`                | 1.25                                                                   | Line height of the dropdown button label                         |
| `--gdd-button-letter-spacing`             | 0.025em                                                                | Letter spacing of the dropdown button label                      |
| `--gdd-button-bg-color`                   | ![Color Preview](https://via.placeholder.com/20/f9fafb?text=+) #f9fafb | Background color of the dropdown button                          |
| `--gdd-button-bg-color-hover`             | ![Color Preview](https://via.placeholder.com/20/f3f4f6?text=+) #f3f4f6 | Background color of the dropdown button on hover                 |
| `--gdd-button-border-width`               | 1px                                                                    | Border width of the dropdown button                              |
| `--gdd-button-border-style`               | solid                                                                  | Border style of the dropdown button                              |
| `--gdd-button-border-color`               | transparent                                                            | Border color of the dropdown button                              |
| `--gdd-button-ring-color`                 | ![Color Preview](https://via.placeholder.com/20/e5e7eb?text=+) #e5e7eb | Color of the focus ring around the dropdown button               |
| `--gdd-button-border-radius`              | 0.25rem                                                                | Border radius of the dropdown button                             |
| `--gdd-button-transition-property`        | background                                                             | CSS property to transition for the dropdown button               |
| `--gdd-button-transition-duration`        | 300ms                                                                  | Duration of the transition for the dropdown button               |
| `--gdd-button-transition-timing-function` | cubic-bezier(0.4, 0, 0.2, 1)                                           | Timing function of the transition for the dropdown button        |
| `--gdd-button-ring-width`                 | 1px                                                                    | Width of the focus ring around the dropdown button               |
| `--gdd-button-ring-style`                 | solid                                                                  | Style of the focus ring around the dropdown button               |
| `--gdd-button-ring-offset`                | 1px                                                                    | Offset of the focus ring around the dropdown button              |
| `--gdd-button-animation-open-name`        | flipOutX                                                               | Animation property (name) for button icon when dropdown is open  |
| `--gdd-button-animation-close-name`       | flipInX                                                                | Animation property (name) for button icon when dropdown is close |

### Dropdown Content

| CSS Variable                                       | Default Value                                                          | Description                                                        |
|----------------------------------------------------|------------------------------------------------------------------------|--------------------------------------------------------------------|
| `--gdd-content-w`                                  | 13rem                                                                  | Width of the dropdown content                                      |
| `--gdd-content-max-w`                              | 16rem                                                                  | Maximum width of the dropdown content                              |
| `--gdd-content-mx`                                 | 0                                                                      | Margin on the x-axis of the dropdown content                       |
| `--gdd-content-my`                                 | 0.25rem                                                                | Margin on the y-axis of the dropdown content                       |
| `--gdd-content-px`                                 | 0.375rem                                                               | Padding on the x-axis of the dropdown content                      |
| `--gdd-content-py`                                 | 0.5rem                                                                 | Padding on the y-axis of the dropdown content                      |
| `--gdd-content-bg-color`                           | ![Color Preview](https://via.placeholder.com/20/ffffff?text=+) #ffffff | Background color of the dropdown content                           |
| `--gdd-content-border-width`                       | 1px                                                                    | Border width of the dropdown content                               |
| `--gdd-content-border-style`                       | solid                                                                  | Border style of the dropdown content                               |
| `--gdd-content-border-color`                       | ![Color Preview](https://via.placeholder.com/20/030712?text=+) #030712 | Border color of the dropdown content                               |
| `--gdd-content-border-radius`                      | 0.25rem                                                                | Border radius of the dropdown content                              |
| `--gdd-content-animation-entrance-duration`        | 0.3s                                                                   | Duration of the entrance animation for the dropdown content        |
| `--gdd-content-animation-entrance-timing-function` | ease-in-out                                                            | Timing function of the entrance animation for the dropdown content |

### Dropdown Item

| CSS Variable                | Default Value                                                          | Description                                                       |
|-----------------------------|------------------------------------------------------------------------|-------------------------------------------------------------------|
| `--gdd-item-px`             | 0.375rem                                                               | Padding on the x-axis of the dropdown item                        |
| `--gdd-item-py`             | 0.375rem                                                               | Padding on the y-axis of the dropdown item                        |
| `--gdd-item-icon-space`     | 1ch                                                                    | Space between the dropdown item label and icon                    |
| `--gdd-item-color`          | ![Color Preview](https://via.placeholder.com/20/f3f4f6?text=+) #f3f4f6 | Color of the item text                                            |
| `--gdd-item-color-hover`    | ![Color Preview](https://via.placeholder.com/20/f3f4f6?text=+) #f3f4f6 | Color of the item text on hover                                   |
| `--gdd-item-font-family`    | inherit                                                                | Font family of the dropdown item label                            |
| `--gdd-item-font-size`      | 1rem                                                                   | Font size of the dropdown item label                              |
| `--gdd-item-font-weight`    | 500                                                                    | Font weight of the dropdown item label                            |
| `--gdd-item-line-height`    | 1.25                                                                   | Line height of the dropdown item label                            |
| `--gdd-item-letter-spacing` | 0.025em                                                                | Letter spacing of the dropdown item label                         |
| `--gdd-item-bg-color`       | transparent                                                            | Background color of the item                                      |
| `--gdd-item-bg-color-hover` | ![Color Preview](https://via.placeholder.com/20/030712?text=+) #030712 | Background color of the item on hover                             |
| `--gdd-item-border-width`   | 1px                                                                    | Border width of the dropdown item                                 |
| `--gdd-item-border-style`   | solid                                                                  | Border style of the dropdown item                                 |
| `--gdd-item-border-color`   | transparent                                                            | Border color of the dropdown item                                 |
| `--gdd-item-border-radius`  | 0.25rem                                                                | Border radius of the dropdown item                                |
| `--gdd-item-ring-width`     | 1px                                                                    | Width of the focus ring around the dropdown item                  |
| `--gdd-item-ring-style`     | solid                                                                  | Style of the focus ring around the dropdown item                  |
| `--gdd-item-ring-offset`    | 0                                                                      | Offset of the focus ring around the dropdown item                 |
| `--gdd-item-ring-color`     | transparent                                                            | Color of the focus ring around the dropdown item                  |
| `--gdd-item-divider-width`  | 1px                                                                    | Width of the divider between dropdown items                       |
| `--gdd-item-divider-style`  | solid                                                                  | Style of the divider between dropdown items (e.g., solid, dashed) |
| `--gdd-item-divider-color`  | ![Color Preview](https://via.placeholder.com/20/030712?text=+) #4b5563 | Color of the item divider                                         |


## Examples

- [use with `a-h/templ`](_examples/a-h-templ)
- [icon only button](_examples/icon-only-button)
- [theming](_examples/theming)
- [positioning](_examples/positioning/)
- [custom animations](_examples/custom-animations)
- [custom-button-icon](_examples/custom-button-icon)
- [use with `template/html`](_examples/go-html-template)

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
