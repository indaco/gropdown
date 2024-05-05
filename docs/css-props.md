# gropdown CSS custom props

Here below is the list of all CSS variables defined and their default values:

## Colors

[Source Code](../gropdown-css.templ#L12-L72).

| CSS Property             | Default Value                                                                                                                                                       | Description                                          |
|:-------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------|
| `--gdd-color-space`      | _oklab_                                                                                                                                                             | The color space used for color mixing.               |
| `--gdd-primary`          | <div style="display: flex; gap: 0.25rem;"><span style="width: 20px; height: 20px; background-color: hsl(215deg 16.3% 46.9%)"></span>_hsl(215deg 16.3% 46.9%)_</div> | Defines the primary color of the dropdown component. |
| `--gdd-primary-10`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 10% brightness.        |
| `--gdd-primary-20`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 20% brightness.        |
| `--gdd-primary-30`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 30% brightness.        |
| `--gdd-primary-40`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 40% brightness.        |
| `--gdd-primary-50`       | _var(--gdd-primary)_                                                                                                                                                | The primary color itself.                            |
| `--gdd-primary-60`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 60% brightness.        |
| `--gdd-primary-70`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 70% brightness.        |
| `--gdd-primary-80`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 80% brightness.        |
| `--gdd-primary-90`       | Result of color-mix function                                                                                                                                        | Shade of the primary color at 90% brightness.        |
| `--gdd-primary-lightest` | _var(--gdd-primary-10)_                                                                                                                                             | The lightest shade of the primary color.             |
| `--gdd-primary-lighter`  | _var(--gdd-primary-20)_                                                                                                                                             | A lighter shade of the primary color.                |
| `--gdd-primary-light`    | _var(--gdd-primary-30)_                                                                                                                                             | A light shade of the primary color.                  |
| `--gdd-primary-dark`     | _var(--gdd-primary-70)_                                                                                                                                             | A dark shade of the primary color.                   |
| `--gdd-primary-darker`   | _var(--gdd-primary-80)_                                                                                                                                             | A darker shade of the primary color.                 |
| `--gdd-primary-darkest`  | _var(--gdd-primary-90)_                                                                                                                                             | The darkest shade of the primary color.              |

## Dropdown Button

| CSS Property                              | Default Value                  | Description                                                       |
|:------------------------------------------|:-------------------------------|:------------------------------------------------------------------|
| `--gdd-button-min-w`                      | _4.5em_                        | Minimum width of the dropdown button.                             |
| `--gdd-button-py`                         | _1ch_                          | Padding on the y-axis of the dropdown button.                     |
| `--gdd-button-px`                         | _2ch_                          | Padding on the x-axis of the dropdown button.                     |
| `--gdd-button-icon-space`                 | _0.5ch_                        | Space between the dropdown button label and icon.                 |
| `--gdd-button-font-size`                  | _1rem_                         | Font size of the dropdown button label.                           |
| `--gdd-button-font-family`                | _inherit_                      | Font family of the dropdown button label.                         |
| `--gdd-button-font-weight`                | _500_                          | Font weight of the dropdown button label.                         |
| `--gdd-button-line-height`                | _1.25_                         | Line height of the dropdown button label.                         |
| `--gdd-button-letter-spacing`             | _0.025em_                      | Letter spacing of the dropdown button label.                      |
| `--gdd-button-border-width`               | _1px_                          | Border width of the dropdown button.                              |
| `--gdd-button-border-style`               | _solid_                        | Border style of the dropdown button.                              |
| `--gdd-button-border-color`               | _transparent_                  | Border color of the dropdown button.                              |
| `--gdd-button-border-radius`              | _0.25rem_                      | Border radius of the dropdown button.                             |
| `--gdd-button-transition-property`        | _background_                   | CSS property to transition for the dropdown button.               |
| `--gdd-button-transition-duration`        | _300ms_                        | Duration of the transition for the dropdown button.               |
| `--gdd-button-transition-timing-function` | _cubic-bezier(0.4, 0, 0.2, 1)_ | Timing function of the transition for the dropdown button.        |
| `--gdd-button-ring-width`                 | _2px_                          | Width of the focus ring around the dropdown button.               |
| `--gdd-button-ring-style`                 | _solid_                        | Style of the focus ring around the dropdown button.               |
| `--gdd-button-ring-offset`                | _1px_                          | Offset of the focus ring around the dropdown button.              |
| `--gdd-button-animation-open-name`        | _flipOutX_                     | Animation property (name) for button icon when dropdown is open.  |
| `--gdd-button-animation-close-name`       | _flipInX_                      | Animation property (name) for button icon when dropdown is close. |

### Light

| CSS Property               | Default Value                 | Description                                         |
|:---------------------------|:------------------------------|:----------------------------------------------------|
| `--gdd-button-color`       | _var(--gdd-primary-darker)_   | Text color of the dropdown button.                  |
| `--gdd-button-color-hover` | _var(--gdd-primary-darker)_   | Text color of the dropdown button on hover.         |
| `--gdd-button-color-focus` | _var(--gdd-primary-darker)_   | Text color of the dropdown button on focus.         |
| `--gdd-button-bg`          | _var(--gdd-primary-lightest)_ | Background color of the dropdown button.            |
| `--gdd-button-bg-hover`    | _var(--gdd-primary-lighter)_  | Background color of the dropdown button on hover.   |
| `--gdd-button-bg-focus`    | _var(--gdd-primary-lighter)_  | Background color of the dropdown button on focus.   |
| `--gdd-button-ring-color`  | _var(--gdd-primary-light)_    | Color of the focus ring around the dropdown button. |

### Dark

| CSS Property               | Default Value                | Description                                         |
|:---------------------------|:-----------------------------|:----------------------------------------------------|
| `--gdd-button-color`       | _var(--gdd-primary-lighter)_ | Text color of the dropdown button.                  |
| `--gdd-button-color-hover` | _var(--gdd-primary-lighter)_ | Text color of the dropdown button on hover.         |
| `--gdd-button-color-focus` | _var(--gdd-primary-lighter)_ | Text color of the dropdown button on focus.         |
| `--gdd-button-bg`          | _var(--gdd-primary-dark)_    | Background color of the dropdown button.            |
| `--gdd-button-bg-hover`    | _var(--gdd-primary-darker)_  | Background color of the dropdown button on hover.   |
| `--gdd-button-bg-focus`    | _var(--gdd-primary-darker)_  | Background color of the dropdown button on focus.   |
| `--gdd-button-ring-color`  | _var(--gdd-primary)_         | Color of the focus ring around the dropdown button. |

## Dropdown Content

| CSS Property                                       | Default Value | Description                                                         |
|:---------------------------------------------------|:--------------|:--------------------------------------------------------------------|
| `--gdd-content-w`                                  | _13rem_       | Width of the dropdown content.                                      |
| `--gdd-content-max-w`                              | _16rem_       | Maximum width of the dropdown content.                              |
| `--gdd-content-mx`                                 | _0_           | Margin on the x-axis of the dropdown content.                       |
| `--gdd-content-my`                                 | _0.25em_      | Margin on the y-axis of the dropdown content.                       |
| `--gdd-content-px`                                 | _0.375em_     | Padding on the x-axis of the dropdown content.                      |
| `--gdd-content-py`                                 | _0.5em_       | Padding on the y-axis of the dropdown content.                      |
| `--gdd-content-border-width`                       | _1px_         | Border width of the dropdown content.                               |
| `--gdd-content-border-style`                       | _solid_       | Border style of the dropdown content.                               |
| `--gdd-content-border-radius`                      | _0.25rem_     | Border radius of the dropdown content.                              |
| `--gdd-content-animation-entrance-duration`        | _0.3s_        | Duration of the entrance animation for the dropdown content.        |
| `--gdd-content-animation-entrance-timing-function` | _ease-in-out_ | Timing function of the entrance animation for the dropdown content. |

### Light

| CSS Property                 | Default Value                                                                                                                       | Description                               |
|:-----------------------------|:------------------------------------------------------------------------------------------------------------------------------------|:------------------------------------------|
| `--gdd-content-bg-color`     | <div style="display: flex; gap: 0.25rem;"><span style="width: 20px; height: 20px; background-color: #ffffff"></span>_#ffffff_</div> | Background color of the dropdown content. |
| `--gdd-content-border-color` | _var(--gdd-primary-lighter)_                                                                                                        | Border color of the dropdown content.     |

### Dark

| CSS Property                 | Default Value                | Description                               |
|:-----------------------------|:-----------------------------|:------------------------------------------|
| `--gdd-content-bg-color`     | _var(--gdd-primary-darker)_  | Background color of the dropdown content. |
| `--gdd-content-border-color` | _var(--gdd-primary-darkest)_ | Border color of the dropdown content.     |

### Dropdown Item

| CSS Property                | Default Value | Description                                        |
|:----------------------------|:--------------|:---------------------------------------------------|
| `--gdd-item-px`             | _0.375em_     | Padding on the x-axis of the dropdown item.        |
| `--gdd-item-py`             | _0.375em_     | Padding on the y-axis of the dropdown item.        |
| `--gdd-item-icon-space`     | _1ch_         | Space between the dropdown item label and icon.    |
| `--gdd-item-font-family`    | _inherit_     | Font family of the dropdown item label.            |
| `--gdd-item-font-size`      | _1rem_        | Font size of the dropdown item label.              |
| `--gdd-item-font-weight`    | _500_         | Font weight of the dropdown item label.            |
| `--gdd-item-line-height`    | _1.25_        | Line height of the dropdown item label.            |
| `--gdd-item-letter-spacing` | _0.025em_     | Letter spacing of the dropdown item label.         |
| `--gdd-item-border-width`   | _1px_         | Border width of the dropdown item.                 |
| `--gdd-item-border-style`   | _solid_       | Border style of the dropdown item.                 |
| `--gdd-item-border-color`   | _transparent_ | Border color of the dropdown item.                 |
| `--gdd-item-border-radius`  | _0.25rem_     | Border radius of the dropdown item.                |
| `--gdd-item-ring-width`     | _1px_         | Width of the focus ring around the dropdown item.  |
| `--gdd-item-ring-style`     | _solid_       | Style of the focus ring around the dropdown item.  |
| `--gdd-item-ring-offset`    | _0_           | Offset of the focus ring around the dropdown item. |
| `--gdd-item-ring-color`     | _transparent_ | Color of the focus ring around the dropdown item.  |

### Light

| CSS Property             | Default Value                 | Description                            |
|:-------------------------|:------------------------------|:---------------------------------------|
| `--gdd-item-color`       | _var(--gdd-primary-dark)_     | Color of the item text.                |
| `--gdd-item-color-hover` | _var(--gdd-primary-darker)_   | Color of the item text on hover.       |
| `--gdd-item-color-focus` | _var(--gdd-primary-darker)_   | Color of the item text on focus.       |
| `--gdd-item-bg`          | _transparent_                 | Background color of the item.          |
| `--gdd-item-bg-hover`    | _var(--gdd-primary-lightest)_ | Background color of the item on hover. |
| `--gdd-item-bg-focus`    | _var(--gdd-primary-lightest)_ | Background color of the item on focus. |

### Dark

| CSS Property             | Default Value                 | Description                            |
|:-------------------------|:------------------------------|:---------------------------------------|
| `--gdd-item-color`       | _var(--gdd-primary-light)_    | Color of the item text.                |
| `--gdd-item-color-hover` | _var(--gdd-primary-lightest)_ | Color of the item text on hover.       |
| `--gdd-item-color-focus` | _var(--gdd-primary-lightest)_ | Color of the item text on focus.       |
| `--gdd-item-bg`          | _transparent_                 | Background color of the item.          |
| `--gdd-item-bg-hover`    | _var(--gdd-primary-60)_       | Background color of the item on hover. |
| `--gdd-item-bg-focus`    | _var(--gdd-primary-60)_       | Background color of the item on focus. |

### Dropdown Divider

| CSS Property               | Default Value        | Description                                                        |
|:---------------------------|:---------------------|:-------------------------------------------------------------------|
| `--gdd-item-divider-width` | _1px_                | Width of the divider between dropdown items.                       |
| `--gdd-item-divider-style` | _solid_              | Style of the divider between dropdown items (e.g., solid, dashed). |
| `--gdd-item-divider-color` | _var(--gdd-primary)_ | Color of the item divider.                                         |
