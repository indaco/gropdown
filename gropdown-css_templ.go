// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package gropdown

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GropdownCSS() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\r\n    * {\r\n      box-sizing: border-box;\r\n      border-width: 0;\r\n      border-style: solid;\r\n    }\r\n\r\n    :root {\r\n      --gdd-color-space: oklab;\r\n\r\n      /* - Color Shades - */\r\n      --gdd-primary: hsl(215deg 16.3% 46.9%);\r\n\r\n      --gdd-primary-10: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 10%,\r\n          white\r\n      );\r\n\r\n      --gdd-primary-20: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 30%,\r\n          white\r\n      );\r\n\r\n      --gdd-primary-30: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 50%,\r\n          white\r\n      );\r\n\r\n      --gdd-primary-40: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 70%,\r\n          white\r\n      );\r\n\r\n      --gdd-primary-50: var(--gdd-primary);\r\n\r\n      --gdd-primary-60: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 70%,\r\n          black\r\n      );\r\n\r\n      --gdd-primary-70: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 50%,\r\n          black\r\n      );\r\n\r\n      --gdd-primary-80: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 30%,\r\n          black\r\n      );\r\n\r\n      --gdd-primary-90: color-mix(\r\n          in var(--gdd-color-space),\r\n          var(--gdd-primary) 10%,\r\n          black\r\n      );\r\n\r\n      --gdd-primary-lightest: var(--gdd-primary-10);\r\n      --gdd-primary-lighter: var(--gdd-primary-20);\r\n      --gdd-primary-light: var(--gdd-primary-30);\r\n      --gdd-primary-dark: var(--gdd-primary-70);\r\n      --gdd-primary-darker: var(--gdd-primary-80);\r\n      --gdd-primary-darkest: var(--gdd-primary-90);\r\n\r\n      /* Dropdown Button */\r\n      --gdd-button-min-w: 4.5em;\r\n      --gdd-button-py: 1ch;\r\n      --gdd-button-px: 2ch;\r\n      --gdd-button-icon-space: 0.5ch;\r\n\r\n      --gdd-button-font-size: 1rem;\r\n      --gdd-button-font-family: inherit;\r\n      --gdd-button-font-weight: 500;\r\n      --gdd-button-line-height: 1.25;\r\n      --gdd-button-letter-spacing: 0.025em;\r\n\r\n      --gdd-button-border-width: 1px;\r\n      --gdd-button-border-style: solid;\r\n      --gdd-button-border-radius: 0.25rem;\r\n\r\n      --gdd-button-transition-property: background;\r\n      --gdd-button-transition-duration: 300ms;\r\n      --gdd-button-transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\r\n\r\n      --gdd-button-ring-width: 2px;\r\n      --gdd-button-ring-style: solid;\r\n      --gdd-button-ring-offset: 1px;\r\n\r\n      --gdd-button-animation-open-name: flipOutX;\r\n      --gdd-button-animation-close-name: flipInX;\r\n\r\n      /* Dropdown Content */\r\n      --gdd-content-w: 13rem;\r\n      --gdd-content-max-w: 16rem;\r\n      --gdd-content-mx: 0;\r\n      --gdd-content-my: 0.25em;\r\n      --gdd-content-px: 0.375em;\r\n      --gdd-content-py: 0.5em;\r\n\r\n      --gdd-content-border-width: 1px;\r\n      --gdd-content-border-style: solid;\r\n      --gdd-content-border-radius: 0.25rem;\r\n\r\n      --gdd-content-animation-name: fadeIn;\r\n      --gdd-content-animation-duration: 0.3s;\r\n      --gdd-content-animation-timing-function: ease-in-out;\r\n      --gdd-content-animation-direction: forwards;\r\n\r\n      /* Dropdown Item */\r\n      --gdd-item-px: 0.375em;\r\n      --gdd-item-py: 0.375em;\r\n      --gdd-item-icon-space: 1ch;\r\n\r\n      --gdd-item-font-family: inherit;\r\n      --gdd-item-font-size: 1rem;\r\n      --gdd-item-font-weight: 500;\r\n      --gdd-item-line-height: 1.25;\r\n      --gdd-item-letter-spacing: 0.025em;\r\n\r\n      --gdd-item-border-width: 1px;\r\n      --gdd-item-border-style: solid;\r\n      --gdd-item-border-color: transparent;\r\n      --gdd-item-border-radius: 0.25rem;\r\n\r\n      --gdd-item-ring-width: 1px;\r\n      --gdd-item-ring-style: solid;\r\n      --gdd-item-ring-offset: 0;\r\n      --gdd-item-ring-color: transparent;\r\n\r\n      /* Divider */\r\n      --gdd-item-divider-width: 1px;\r\n      --gdd-item-divider-style: solid;\r\n    }\r\n\r\n    @media (prefers-color-scheme: light) {\r\n      :root {\r\n        /* button */\r\n        --gdd-button-color: var(--gdd-primary-darker);\r\n        --gdd-button-color-hover: var(--gdd-primary-darker);\r\n        --gdd-button-color-focus: var(--gdd-primary-darker);\r\n        --gdd-button-bg: var(--gdd-primary-lightest);\r\n        --gdd-button-bg-hover: var(--gdd-primary-lighter);\r\n        --gdd-button-bg-focus: var(--gdd-primary-lighter);\r\n        --gdd-button-border-color: transparent;\r\n        --gdd-button-ring-color: var(--gdd-primary-light);\r\n        /* content */\r\n        --gdd-content-bg: #FFFFFF;\r\n        --gdd-content-border-color: var(--gdd-primary-lighter);\r\n        /* item */\r\n        --gdd-item-color: var(--gdd-primary-dark);\r\n        --gdd-item-color-hover: var(--gdd-primary-darker);\r\n        --gdd-item-color-focus: var(--gdd-primary-darker);\r\n        --gdd-item-bg: transparent;\r\n        --gdd-item-bg-hover: var(--gdd-primary-lightest);\r\n        --gdd-item-bg-focus: var(--gdd-primary-lightest);\r\n        /* divider */\r\n        --gdd-divider-color: var(--gdd-primary-lighter);\r\n      }\r\n    }\r\n\r\n    @media (prefers-color-scheme: dark){\r\n      :root {\r\n        /* button */\r\n        --gdd-button-color: var(--gdd-primary-lighter);\r\n        --gdd-button-color-hover: var(--gdd-primary-lighter);\r\n        --gdd-button-color-focus: var(--gdd-primary-lighter);\r\n        --gdd-button-bg: var(--gdd-primary-dark);\r\n        --gdd-button-bg-hover: var(--gdd-primary-darker);\r\n        --gdd-button-bg-focus: var(--gdd-primary-darker);\r\n        --gdd-button-border-color: transparent;\r\n        --gdd-button-ring-color: var(--gdd-primary);\r\n        /* content */\r\n        --gdd-content-bg: var(--gdd-primary-darker);\r\n        --gdd-content-border-color: var(--gdd-primary-darkest);\r\n        /* item */\r\n        --gdd-item-color: var(--gdd-primary-light);\r\n        --gdd-item-color-hover: var(--gdd-primary-lightest);\r\n        --gdd-item-color-focus: var(--gdd-primary-lightest);\r\n        --gdd-item-bg: transparent;\r\n        --gdd-item-bg-hover:  var(--gdd-primary-60);\r\n        --gdd-item-bg-focus:  var(--gdd-primary-60);\r\n        /* divider */\r\n        --gdd-divider-color: var(--gdd-primary);\r\n      }\r\n    }\r\n\r\n    :root[data-theme=\"light\"] {\r\n      /* button */\r\n      --gdd-button-color: var(--gdd-primary-darker);\r\n      --gdd-button-color-hover: var(--gdd-primary-darker);\r\n      --gdd-button-color-focus: var(--gdd-primary-darker);\r\n      --gdd-button-bg: var(--gdd-primary-lightest);\r\n      --gdd-button-bg-hover: var(--gdd-primary-lighter);\r\n      --gdd-button-bg-focus: var(--gdd-primary-lighter);\r\n      --gdd-button-border-color: transparent;\r\n      --gdd-button-ring-color: var(--gdd-primary-light);\r\n      /* content */\r\n      --gdd-content-bg: #FFFFFF;\r\n      --gdd-content-border-color: var(--gdd-primary-lighter);\r\n      /* item */\r\n      --gdd-item-color: var(--gdd-primary-dark);\r\n      --gdd-item-color-hover: var(--gdd-primary-darker);\r\n      --gdd-item-color-focus: var(--gdd-primary-darker);\r\n      --gdd-item-bg: transparent;\r\n      --gdd-item-bg-hover: var(--gdd-primary-lightest);\r\n      --gdd-item-bg-focus: var(--gdd-primary-lightest);\r\n      /* divider */\r\n      --gdd-divider-color: var(--gdd-primary-lighter);\r\n    }\r\n\r\n    :root[data-theme=\"dark\"] {\r\n      /* button */\r\n      --gdd-button-color: var(--gdd-primary-lighter);\r\n      --gdd-button-color-hover: var(--gdd-primary-lighter);\r\n      --gdd-button-color-focus: var(--gdd-primary-lighter);\r\n      --gdd-button-bg: var(--gdd-primary-dark);\r\n      --gdd-button-bg-hover: var(--gdd-primary-darker);\r\n      --gdd-button-bg-focus: var(--gdd-primary-darker);\r\n      --gdd-button-border-color: transparent;\r\n      --gdd-button-ring-color: var(--gdd-primary);\r\n      /* content */\r\n      --gdd-content-bg: var(--gdd-primary-darker);\r\n      --gdd-content-border-color: var(--gdd-primary-darkest);\r\n      /* item */\r\n      --gdd-item-color: var(--gdd-primary-light);\r\n      --gdd-item-color-hover: var(--gdd-primary-lightest);\r\n      --gdd-item-color-focus: var(--gdd-primary-lightest);\r\n      --gdd-item-bg: transparent;\r\n      --gdd-item-bg-hover:  var(--gdd-primary-60);\r\n      --gdd-item-bg-focus:  var(--gdd-primary-60);\r\n      /* divider */\r\n      --gdd-divider-color: var(--gdd-primary);\r\n    }\r\n\r\n    /* ********************************************* *\r\n     * Dropdown Container\r\n     * ********************************************* */\r\n    .gdd_gropdown_container {\r\n      -webkit-font-smoothing: antialiased;\r\n      -moz-osx-font-smoothing: grayscale;\r\n      font-smooth: auto;\r\n      position: relative;\r\n      display: inline-block;\r\n    }\r\n\r\n    /* ********************************************* *\r\n     * Dropdown Button\r\n     * ********************************************* */\r\n    .gdd_button {\r\n      cursor: pointer;\r\n      min-width: var(--gdd-button-min-w);\r\n      display: inline-flex;\r\n      align-items: center;\r\n      justify-content: center;\r\n      gap: var(--gdd-button-icon-space);\r\n      margin: 0;\r\n      padding: var(--gdd-button-py) var(--gdd-button-px);\r\n      color: var(--gdd-button-color);\r\n      border-width: var(--gdd-button-border-width);\r\n      border-style: var(--gdd-button-border-style);\r\n      border-color: var(--gdd-button-border-color);\r\n      border-radius: var(--gdd-button-border-radius);\r\n      background-color: var(--gdd-button-bg);\r\n      font-family: var(--gdd-button-font-family);\r\n      font-size: var(--gdd-button-font-size);\r\n      font-weight: var(--gdd-button-font-weight);\r\n      line-height: var(--gdd-button-line-height);\r\n      letter-spacing: var(--gdd-button-letter-spacing);\r\n      text-transform: none;\r\n      text-decoration-line: var(--_text-decoration-line);\r\n      transition: var(--gdd-button-transition-property) var(--gdd-button-transition-duration) var(--gdd-button-transition-timing-function);\r\n    }\r\n\r\n    .gdd_button:hover {\r\n      --gdd-button-color: var(--gdd-button-color-hover);\r\n      --gdd-button-bg: var(--gdd-button-bg-hover);\r\n    }\r\n\r\n    .gdd_button:focus,\r\n    .gdd_button:focus-visible {\r\n      --gdd-button-color: var(--gdd-button-color-focus);\r\n      --gdd-button-bg: var(--gdd-button-bg-focus);\r\n      outline: var(--gdd-button-ring-color) var(--gdd-button-ring-style) var(--gdd-button-ring-width);\r\n      outline-offset: var(--gdd-button-ring-offset);\r\n    }\r\n\r\n    .gdd_button_icon {\r\n      display: inline-flex;\r\n      align-items: center;\r\n      flex-shrink: 0;\r\n      width: 1em;\r\n      height: 1em;\r\n    }\r\n\r\n    .gdd_button_icon_only {\r\n      min-width: 2.5rem;\r\n      padding: 0;\r\n      width: 2.5rem;\r\n      height: 2.5rem;\r\n      border-radius: 1e5px;\r\n      aspect-ratio: 1;\r\n    }\r\n\r\n    /* Apply animations based on classes */\r\n    .iconToOpen {\r\n      animation: var(--gdd-button-animation-open-name) 0.3s ease forwards;\r\n    }\r\n\r\n    .iconToClose {\r\n      animation: var(--gdd-button-animation-close-name) 0.3s ease forwards;\r\n    }\r\n\r\n    /* ********************************************* *\r\n     * Dropdown Content\r\n     * ********************************************* */\r\n    .gdd_content {\r\n      position: absolute;\r\n      z-index: 10;\r\n      overflow: hidden;\r\n      list-style: none;\r\n      width: var(--gdd-content-w);\r\n      max-width: var(--gdd-content-max-w);\r\n      margin: var(--gdd-content-my) var(--gdd-content-mx);\r\n      padding: var(--gdd-content-py) var(--gdd-content-px);\r\n      background-color: var(--gdd-content-bg);\r\n      border-width: var(--gdd-content-border-width);\r\n      border-style: var(--gdd-content-border-style);\r\n      border-color: var(--gdd-content-border-color);\r\n      border-radius: var(--gdd-content-border-radius);\r\n      transition: opacity var(--gdd-content-animation-duration) var(--gdd-content-animation-timing-function);\r\n      animation-duration: var(--gdd-content-animation-duration);\r\n      animation-direction: var(--gdd-content-animation-direction);\r\n      animation-timing-function: var(--gdd-content-animation-timing-function);\r\n    }\r\n\r\n    [data-position=\"top\"] .gdd_content {\r\n      top: auto;\r\n      bottom: 100%;\r\n    }\r\n\r\n    [data-position=\"bottom\"] .gdd_content {\r\n      top: 100%;\r\n      bottom: auto;\r\n    }\r\n\r\n    [data-position=\"left\"] .gdd_content {\r\n      --gdd-content-my: 0;\r\n      --gdd-content-mx: 0.25rem;\r\n      inset-inline-end: 100%;\r\n      top: 0;\r\n      bottom: auto;\r\n    }\r\n\r\n    [data-position=\"right\"] .gdd_content {\r\n      --gdd-content-my: 0;\r\n      --gdd-content-mx: 0.25rem;\r\n      inset-inline-start: 100%;\r\n      top: 0;\r\n      bottom: auto;\r\n    }\r\n\r\n    .gdd_content[data-state=\"open\"] {\r\n      visibility: visible;\r\n      opacity: 1;\r\n      animation-name: var(--gdd-content-animation-name);\r\n    }\r\n\r\n    .gdd_content[data-state=\"close\"] {\r\n      visibility: hidden;\r\n      opacity: 0;\r\n    }\r\n\r\n    /* ********************************************* *\r\n     * Dropdown Content Item\r\n     * ********************************************* */\r\n    .gdd_content_item {\r\n      position: relative;\r\n      display: flex;\r\n      width: 100%;\r\n      text-decoration: none;\r\n      gap: var(--gdd-item-icon-space);\r\n      padding: var(--gdd-item-py) var(--gdd-item-px);\r\n      color: var(--gdd-item-color);\r\n      font-family: var(--gdd-item-font-family);\r\n      font-size: var(--gdd-item-font-size);\r\n      line-height: var(--gdd-item-line-height);\r\n      letter-spacing: var(--gdd-item-letter-spacing);\r\n      background-color: var(--gdd-item-bg);\r\n      border-width: var(--gdd-item-border-width);\r\n      border-style: var(--gdd-item-border-style);\r\n      border-color: var(--gdd-item-border-color);\r\n      border-radius: var(--gdd-item-border-radius);\r\n    }\r\n\r\n    .gdd_content_item[href^=\"http\"] {\r\n      padding-right: 1.25em;\r\n    }\r\n\r\n    .gdd_content_item:hover {\r\n      --gdd-item-color: var(--gdd-item-color-hover);\r\n      --gdd-item-bg: var(--gdd-item-bg-hover);\r\n    }\r\n\r\n    .gdd_content_item:focus,\r\n    .gdd_content_item:focus-visible {\r\n      --gdd-item-color: var(--gdd-item-color-focus);\r\n      --gdd-item-bg: var(--gdd-item-bg-focus);\r\n      outline: var(--gdd-item-ring-color) var(--gdd-item-ring-style) var(--gdd-item-ring-width);\r\n      outline-offset: var(--gdd-item-ring-offset);\r\n      opacity: 1;\r\n    }\r\n\r\n    .gdd_content_item_icon {\r\n      display: inline-flex;\r\n      align-items: center;\r\n      flex-shrink: 0;\r\n      width: 1.25em;\r\n      height: 1.25em;\r\n    }\r\n\r\n    .gdd_content_divider {\r\n      height: 0;\r\n      margin: 0.125rem 0;\r\n      overflow: hidden;\r\n      border-top: var(--gdd-item-divider-width) var(--gdd-item-divider-style) var(--gdd-divider-color);\r\n    }\r\n\r\n    .gdd_sr_only {\r\n      position: absolute;\r\n      width: 1px;\r\n      height: 1px;\r\n      padding: 0;\r\n      margin: -1px;\r\n      overflow: hidden;\r\n      clip: rect(0, 0, 0, 0);\r\n      white-space: nowrap;\r\n      border-width: 0;\r\n    }\r\n\r\n    /* ********************************************* *\r\n     * Animations\r\n     * ********************************************* */\r\n    @keyframes fadeIn {\r\n      from {\r\n        opacity: 0;\r\n      }\r\n\r\n      to {\r\n        opacity: 1;\r\n      }\r\n    }\r\n\r\n    @keyframes flipInX {\r\n      from {\r\n        transform: scaleY(-1);\r\n      }\r\n\r\n      to {\r\n        transform: scaleY(1);\r\n      }\r\n    }\r\n\r\n    @keyframes flipOutX {\r\n      from {\r\n        transform: scaleY(1);\r\n      }\r\n\r\n      to {\r\n        transform: scaleY(-1);\r\n      }\r\n    }\r\n</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
