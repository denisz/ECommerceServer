package themes

// Flat is a theme
type Default struct{}

// Name returns the name of the flat theme
func (dt *Default) Name() string {
	return "flat"
}


// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Default) HTMLTemplate() string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" style="background-color: #fff;">
<head style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
  <meta name="viewport" content="width=device-width, initial-scale=1.0" style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
  <style type="text/css" rel="stylesheet" media="all" style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;"> 
    /*Media Queries ------------------------------ */
    @media only screen and (max-width: 600px) {
      .email-body_inner,
      .email-footer_inner {
        width: 100% !important;
      } 
    }
  </style>
</head>
<body dir="{{.Hermes.TextDirection}}" style="height: 100%;margin: 0;padding: 0;border: 0;outline: 0;line-height: 1.4;background-color: #fff;color: #000;-webkit-text-size-adjust: none;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;width: 100% !important;">
  <table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0" style="width: 100%;margin: 0;padding: 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
    <tr>
      <td class="content" style="padding: 0;color: #000;font-size: 15px;line-height: 18px;align: center;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
        <table class="email-content" width="100%" cellpadding="0" cellspacing="0" style="width: 100%;margin: 0;padding: 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
          <!-- Logo -->
          <tr>
            <td class="email-masthead" style="padding: 15px 10px;color: #000;font-size: 15px;line-height: 18px;text-align: left;background-color: #0d253a;height: 220px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
              <a class="email-masthead_name" href="{{.Hermes.Product.Link}}" target="_blank" style="color: #2F3133;font-size: 16px;font-weight: bold;text-decoration: none;text-shadow: 0 1px 0 white;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                {{ if .Hermes.Product.Logo }}
                  <img src="{{.Hermes.Product.Logo | url }}" alt="Logo" title="Logo" class="email-logo" style="max-height: 130px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                {{ else }}
                  {{ .Hermes.Product.Name }}
                {{ end }}
                </a>
            </td>
          </tr>

          <!-- Email Body -->
          <tr>
            <td class="email-body" width="100%" style="padding: 0;color: #000;font-size: 15px;line-height: 18px;width: 100%;margin: 0;border-top: 1px solid #EDEFF2;border-bottom: 1px solid #EDEFF2;background-color: #FFF;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
              <table class="email-body_inner" width="640" cellpadding="0" cellspacing="0" style="width: 100%;//max-width: 640px;//margin: 0 auto;padding: 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                <!-- Body content -->
                <tr>
                  <td class="content-cell" style="padding: 10px;color: #000;font-size: 15px;line-height: 18px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                    <h1 style="margin-top: 0;color: #2F3133;font-size: 19px;font-weight: bold;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }}, {{ .Email.Body.Name }}! {{ end }}</h1>
                    {{ with .Email.Body.Intros }}
                        {{ if gt (len .) 0 }}
                          {{ range $line := . }}
                            <p style="margin-top: 0;color: #000;font-size: 16px;line-height: 1.5em;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{ $line }}</p>
                          {{ end }}
                        {{ end }}
                    {{ end }}

                    {{ if (ne .Email.Body.FreeMarkdown "") }}
                      {{ .Email.Body.FreeMarkdown.ToHTML }}
                    {{ end }}

                      {{ with .Email.Body.Dictionary }} 
                        {{ if gt (len .) 0 }}
                          <dl class="body-dictionary" style="width: 100%;overflow: hidden;margin: 20px auto 20px;padding: 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                            {{ range $entry := . }}
                              <dt style="clear: both;color: #000;font-weight: bold;float: left;width: 40%;padding: 0;margin: 0;margin-bottom: 0.3em;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{ $entry.Key }}:</dt>
                              <dd style="float: left;width: 60%;padding: 0;margin: 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{ $entry.Value }}</dd>
                            {{ end }}
                          </dl>
                        {{ end }}
                      {{ end }}

                      <!-- Table -->
                      {{ with .Email.Body.Table }}
                        {{ $data := .Data }}
                        {{ $columns := .Columns }}
                        {{ if gt (len $data) 0 }}
                          <table class="data-wrapper" width="100%" cellpadding="0" cellspacing="0" style="width: 100%;margin: 0;padding: 5px 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                            <tr>
                              <td colspan="2" style="padding: 10px 5px;color: #000;font-size: 15px;line-height: 18px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                                <table class="data-table" width="100%" cellpadding="0" cellspacing="0" style="width: 100%;margin: 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                                  <tr>
                                    {{ $col := index $data 0 }}
                                    {{ range $entry := $col }}
                                      <th 
                                        {{ with $columns }}
                                          {{ $width := index .CustomWidth $entry.Key }}
                                          {{ with $width }}
                                            width="{{ . }}"
                                          {{ end }}
                                          {{ $align := index .CustomAlignment $entry.Key }}
                                          {{ with $align }}                                            
                                            style="text-align:{{ . }}; padding: 0px 5px;padding-bottom: 8px;border-bottom: 1px solid #EDEFF2;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;"
                                          {{ end }}
                                        {{ end }}>
                                        <p style="margin-top: 0;color: #9BA2AB;font-size: 12px;line-height: 1.5em;margin: 0;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{ $entry.Key }}
                                        </p>
                                      </th>
                                    {{ end }}
                                  </tr>
                                  {{ range $row := $data }}
                                    <tr>
                                      {{ range $cell := $row }}
                                        <td 
                                          {{ with $columns }}
                                            {{ $align := index .CustomAlignment $cell.Key }}
                                            {{ with $align }}
                                              style="text-align:{{ . }}; padding: 10px 5px;color: #000;font-size: 15px;line-height: 18px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;"
                                            {{ end }}
                                          {{ end }}>                                          
                                          {{ $cell.Value }}
                                        </td>
                                      {{ end }}
                                    </tr>
                                  {{ end }}
                                </table>
                              </td>
                            </tr>
                          </table>
                        {{ end }}
                      {{ end }}

                      <!-- Action -->
                      {{ with .Email.Body.Actions }}
                        {{ if gt (len .) 0 }}
                          {{ range $action := . }}
                            <p style="margin-top: 0;color: #000;font-size: 16px;line-height: 1.5em;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{ $action.Instructions }}</p>
                            <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0" style="width: 100%;margin: 30px auto;padding: 0;text-align: center;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                              <tr>
                                <td align="center" style="padding: 10px 5px;color: #000;font-size: 15px;line-height: 18px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                                  <div style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                                    <a href="{{ $action.Button.Link }}" class="button" style="padding: 10px 5px;background-color: {{ $action.Button.Color }};color: #ffffff;display: block;width: 100%;font-size: 15px;line-height: 45px;text-align: center;text-decoration: none;-webkit-text-size-adjust: none;mso-hide: all;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;" target="_blank">
                                      {{ $action.Button.Text }}
                                    </a>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          {{ end }}
                        {{ end }}
                      {{ end }}
 
                    {{ with .Email.Body.Outros }} 
                        {{ if gt (len .) 0 }}
                          {{ range $line := . }}
                            <p style="margin-top: 0;color: #000;font-size: 16px;line-height: 1.5em;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{ $line }}</p>
                          {{ end }}
                        {{ end }}
                      {{ end }}

                    <p style="margin-top: 0;color: #000;font-size: 16px;line-height: 1.5em;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                      {{.Email.Body.Signature}},
                      <br>
                      {{.Hermes.Product.Name}}
                    </p> 
                  </td>
                </tr>
              </table>
            </td>
          </tr>
          <tr>
            <td class="email-footer" style="padding: 10px 5px;color: #000;font-size: 15px;line-height: 18px;background-color: #0d253a;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
              <table class="email-footer_inner" align="center" width="570" cellpadding="0" cellspacing="0" style="width: 570px;margin: 0 auto;padding: 0;background-color: rgb(13, 37, 58);text-align: center;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                <tr>
                  <td class="content-cell" style="padding: 10px;color: #000;font-size: 15px;line-height: 18px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                    <p class="sub center" style="margin-top: 0;color: #eaeaea;font-size: 12px;line-height: 1.5em;text-align: center;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                      {{.Hermes.Product.Copyright}}	
                    </p>
                  </td>
                </tr>
              </table>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Default) PlainTextTemplate() string {
	return `
<h2>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }}{{ end }},</h2>
{{ with .Email.Body.Intros }}
  {{ range $line := . }}
    <p>{{ $line }}</p>
  {{ end }}
{{ end }}
{{ if (ne .Email.Body.FreeMarkdown "") }}
  {{ .Email.Body.FreeMarkdown.ToHTML }}
{{ else }}
  {{ with .Email.Body.Dictionary }}
    <ul>
    {{ range $entry := . }}
      <li>{{ $entry.Key }}: {{ $entry.Value }}</li>
    {{ end }}
    </ul>
  {{ end }}
  {{ with .Email.Body.Table }}
    {{ $data := .Data }}
    {{ $columns := .Columns }}
    {{ if gt (len $data) 0 }}
      <table class="data-table" width="100%" cellpadding="0" cellspacing="0">
        <tr>
          {{ $col := index $data 0 }}
          {{ range $entry := $col }}
            <th>{{ $entry.Key }} </th>
          {{ end }}
        </tr>
        {{ range $row := $data }}
          <tr>
            {{ range $cell := $row }}
              <td>
                {{ $cell.Value }}
              </td>
            {{ end }}
          </tr>
        {{ end }}
      </table>
    {{ end }}
  {{ end }}
  {{ with .Email.Body.Actions }} 
    {{ range $action := . }}
      <p>{{ $action.Instructions }} {{ $action.Button.Link }}</p> 
    {{ end }}
  {{ end }}
{{ end }}
{{ with .Email.Body.Outros }} 
  {{ range $line := . }}
    <p>{{ $line }}<p>
  {{ end }}
{{ end }}
<p>{{.Email.Body.Signature}},<br>{{.Hermes.Product.Name}} - {{.Hermes.Product.Link}}</p>

<p>{{.Hermes.Product.Copyright}}</p>
`
}