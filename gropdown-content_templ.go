// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package gropdown

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func content(dropdown *Dropdown) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{gddContent()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(menuId(dropdown.Button.Label)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" role=\"menu\" aria-labelledby=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(buttonId(dropdown.Button.Label)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if dropdown.Open {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" data-state=\"open\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" data-state=\"close\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range dropdown.Items {
			templ_7745c5c3_Err = contentItem(item).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func gddContent() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`position:absolute;`)
	templ_7745c5c3_CSSBuilder.WriteString(`z-index:10;`)
	templ_7745c5c3_CSSBuilder.WriteString(`overflow:hidden;`)
	templ_7745c5c3_CSSBuilder.WriteString(`list-style:none;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:var(--gdd-content-w);`)
	templ_7745c5c3_CSSBuilder.WriteString(`max-width:var(--gdd-content-max-w);`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:var(--gdd-content-my) var(--gdd-content-mx);`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:var(--gdd-content-py) var(--gdd-content-px);`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--gdd-content-bg-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-width:var(--gdd-content-border-width);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-style:var(--gdd-content-border-style);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-color:var(--gdd-content-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--gdd-content-border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`transition:opacity var(--gdd-content-animation-duration) var(--gdd-content-animation-timing-function);`)
	templ_7745c5c3_CSSBuilder.WriteString(`animation-duration:var(--gdd-content-animation-duration);`)
	templ_7745c5c3_CSSBuilder.WriteString(`animation-direction:var(--gdd-content-animation-direction);`)
	templ_7745c5c3_CSSBuilder.WriteString(`animation-timing-function:var(--gdd-content-animation-timing-function);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gddContent`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}
