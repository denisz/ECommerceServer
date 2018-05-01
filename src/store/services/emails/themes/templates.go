package themes

//https://templates.mailchimp.com/resources/inline-css/
const t = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" style="background-color: #fff;">
<head style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
  <meta name="viewport" content="width=device-width, initial-scale=1.0" style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
  <style type="text/css" rel="stylesheet" media="all" style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
    /* Base ------------------------------ */
    *:not(br):not(tr):not(html) {
      font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;
      -webkit-box-sizing: border-box;
      box-sizing: border-box;
    }
	html {
		background-color: #fff;
	}
    body {
      width: 100% !important;
      height: 100%;
      margin: 0;
      padding: 0;
      border:0; 
      outline:0;
      line-height: 1.4;
      background-color: #fff;
      color: #000;
      -webkit-text-size-adjust: none;
    }
    a {
      color: #3869D4;
    }

    /* Layout ------------------------------ */
    .email-wrapper {
      width: 100%;
      margin: 0;
      padding: 0; 
    }
    .email-content {
      width: 100%;
      margin: 0;
      padding: 0;
    }
    /* Masthead ----------------------- */
    .email-masthead {
      padding: 15px 10px;
      text-align: left;
      background-color: #0d253a;
      height: 220px;
    }
    .email-masthead_logo {
      max-width: 400px;
      border: 0;
    }
    .email-masthead_name {
      font-size: 16px;
      font-weight: bold;
      color: #2F3133;
      text-decoration: none; 
      text-shadow: 0 1px 0 white;
    }
    .email-logo {
      max-height: 130px;
    }
    /* Footer ------------------------------ */
	.email-footer {
		background-color: #0d253a;
	}
    /* Body ------------------------------ */
    .email-body {
      width: 100%;
      margin: 0;
      padding: 0;
      border-top: 1px solid #EDEFF2;
      border-bottom: 1px solid #EDEFF2;
      background-color: #FFF;
    }
    .email-body_inner {
      //max-width: 640px;
      //margin: 0 auto;
      padding: 0;
    }
    .email-footer_inner {
      width: 570px;
      margin: 0 auto;
      padding: 0;
      background-color: rgb(13, 37, 58);
      text-align: center;
    }
    .email-footer_inner p {
      color: #eaeaea;
    }
    .body-action {
      width: 100%;
      margin: 30px auto;
      padding: 0;
      text-align: center;
    }

    .body-dictionary {
      width: 100%;
      overflow: hidden;
      margin: 20px auto 20px;
      padding: 0;
    }
    .body-dictionary dt {
      clear: both;
      color: #000;
      font-weight: bold;
      float: left;
      width: 40%;
      padding: 0;
      margin: 0;
      margin-bottom: 0.3em;
    }
    .body-dictionary dd {
      float: left;
      width: 60%;
      padding: 0;
      margin: 0;
    }
    .body-sub {
      margin-top: 25px;
      padding-top: 25px;
      border-top: 1px solid #EDEFF2;
      table-layout: fixed;
    }
    .body-sub a {
      word-break: break-all;
    }
    .content-cell {
      padding: 10px;
    }
    .align-right {
      text-align: right;
    }
    /* Type ------------------------------ */
    h1 {
      margin-top: 0;
      color: #2F3133;
      font-size: 19px;
      font-weight: bold;
    }
    h2 {
      margin-top: 0;
      color: #2F3133;
      font-size: 16px;
      font-weight: bold;
    }
    h3 {
      margin-top: 0;
      color: #2F3133;
      font-size: 14px;
      font-weight: bold;
    }
    blockquote {
      margin: 1.7rem 0;
      padding-left: 0.85rem;
      border-left: 10px solid #F0F2F4;
    }
    blockquote p {
        font-size: 1.1rem;
        color: #999;
    }
    blockquote cite {
        display: block;
        text-align: right;
        color: #666;
        font-size: 1.2rem;
    }
    cite {
      display: block;
      font-size: 0.925rem; 
    }
    cite:before {
      content: "\2014 \0020";
    }
    p {
      margin-top: 0;
      color: #000;
      font-size: 16px;
      line-height: 1.5em;
    }
    p.sub {
      font-size: 12px;
    }
    p.center {
      text-align: center;
    }
    table {
      width: 100%;
    }
    th {
      padding: 0px 5px;
      padding-bottom: 8px;
      border-bottom: 1px solid #EDEFF2;
    }
    th p {
      margin: 0;
      color: #9BA2AB;
      font-size: 12px;
    }
    td {
      padding: 10px 5px;
      color: #000;
      font-size: 15px;
      line-height: 18px;
    }
    .content {
      align: center;
      padding: 0;
    }
    /* Data table ------------------------------ */
    .data-wrapper {
      width: 100%;
      margin: 0;
      padding: 5px 0;
    }
    .data-table {
      width: 100%;
      margin: 0;
    }
    .data-table th {
      text-align: left;
      padding: 0px 5px;
      padding-bottom: 8px;
      border-bottom: 1px solid #EDEFF2;
    }
    .data-table th p {
      margin: 0;
      color: #9BA2AB;
      font-size: 12px;
    }
    .data-table td {
      padding: 10px 5px;
      color: #000;
      font-size: 15px;
      line-height: 18px;
    }
    /* Buttons ------------------------------ */
    .button {
      display: inline-block;
      width: 100%;
      background-color: #00948d;
      color: #ffffff;
      font-size: 15px;
      line-height: 45px;
      text-align: center;
      text-decoration: none;
      -webkit-text-size-adjust: none;
      mso-hide: all;
    }
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
                  <img src="{{.Hermes.Product.Logo | url }}" class="email-logo" style="max-height: 130px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
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
                                    <a href="{{ $action.Button.Link }}" class="button" style="background-color: {{ $action.Button.Color }};color: #ffffff;display: inline-block;width: 100%;font-size: 15px;line-height: 45px;text-align: center;text-decoration: none;-webkit-text-size-adjust: none;mso-hide: all;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;" target="_blank">
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

                    {{ if (eq .Email.Body.FreeMarkdown "") }}
                      {{ with .Email.Body.Actions }} 
                        <table class="body-sub" style="width: 100%;margin-top: 25px;padding-top: 25px;border-top: 1px solid #EDEFF2;table-layout: fixed;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                          <tbody style="font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                              {{ range $action := . }}
                                <tr>
                                  <td style="padding: 10px 5px;color: #000;font-size: 15px;line-height: 18px;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">
                                    <p class="sub" style="margin-top: 0;color: #000;font-size: 12px;line-height: 1.5em;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{$.Hermes.Product.TroubleText | replace "{ACTION}" $action.Button.Text}}</p>
                                    <p class="sub" style="margin-top: 0;color: #000;font-size: 12px;line-height: 1.5em;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;"><a href="{{ $action.Button.Link }}" style="color: #3869D4;word-break: break-all;font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;-webkit-box-sizing: border-box;box-sizing: border-box;">{{ $action.Button.Link }}</a></p>
                                  </td>
                                </tr>
                              {{ end }}
                          </tbody>
                        </table>
                      {{ end }}
                    {{ end }}
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
</html>`
