// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func welcome() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Linux Course For Builders</title><link href=\"/static/styles.css\" rel=\"stylesheet\"></head><body class=\"min-h-screen bg-gray-900 flex items-center justify-center p-4\"><div class=\"max-w-4xl w-full\"><div class=\"bg-gray-800 rounded-lg shadow-2xl border-2 border-orange-600/20 p-8\"><div class=\"flex items-center justify-center space-x-4 mb-8\"><!-- Terminal Icon using SVG --><svg class=\"w-12 h-12 text-orange-500\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\"><rect x=\"2\" y=\"4\" width=\"20\" height=\"16\" rx=\"2\"></rect> <path d=\"M6 8l4 4-4 4\"></path> <path d=\"M12 16h6\"></path></svg><h1 class=\"text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-orange-500 to-red-600\">Linux Course For Builders</h1></div><div class=\"grid grid-cols-1 md:grid-cols-3 gap-6 mb-8\"><div class=\"bg-gray-700/50 p-4 rounded-lg border border-orange-500/10 hover:border-orange-500/30 transition-colors\"><h3 class=\"text-orange-400 font-semibold mb-2\">Cross-Platform</h3><p class=\"text-gray-300 text-sm\">Build once, run anywhere</p></div><div class=\"bg-gray-700/50 p-4 rounded-lg border border-orange-500/10 hover:border-orange-500/30 transition-colors\"><h3 class=\"text-orange-400 font-semibold mb-2\">Performance</h3><p class=\"text-gray-300 text-sm\">Optimized for real-world use</p></div><div class=\"bg-gray-700/50 p-4 rounded-lg border border-orange-500/10 hover:border-orange-500/30 transition-colors\"><h3 class=\"text-orange-400 font-semibold mb-2\">Modern Tools</h3><p class=\"text-gray-300 text-sm\">Using cutting-edge technology</p></div></div><div class=\"text-center\"><div class=\"inline-block\"><code class=\"bg-gray-950 text-gray-300 p-4 rounded-lg font-mono text-sm block\">$ go run main.go <span class=\"text-orange-500 ml-2\"></span></code></div></div></div><footer class=\"mt-6 text-center text-gray-500 text-sm\"><p>Built with Go and Templ</p></footer></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
