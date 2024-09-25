package helper

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/widgets"
	"github.com/NY-Daystar/corpos-christie/settings"
	"github.com/NY-Daystar/corpos-christie/tax"
	"github.com/NY-Daystar/corpos-christie/user"
	"gopkg.in/gomail.v2"
)

// Format and configure mail to send
func NewMail(from string, to string, subject string, body string) *gomail.Message {
	var mail = gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	return mail
}

// Configure and return SMTP client to send mail
func NewSMTP(config *settings.Smtp) *gomail.Dialer {
	return gomail.NewDialer(config.Host, config.Port, config.User, config.Password)
}

// FormatMail - Get taxes data and format it into html content
func FormatMail(user *user.User, cfg *config.Config, params settings.Settings, language settings.Yaml, popup *widgets.MailPopup) string {
	var result = tax.CalculateTax(user, cfg)
	var tranches = getTaxTrancheResult(result, *params.Currency, language)

	var emailTmpl = `{{.Prelude}}
	<br/>
	<hr>
	<h3 style="width: 70%;margin: auto;">{{.TaxResults}}</h3>
	<table style="width: 70%;margin: auto;border: 1px solid black;border-collapse: collapse;">
		<tbody>
			<tr>
				<td>Income</td>
				<td>{{.Income}} {{.Currency}}</td>
			</tr>
			<tr>
				<td>In couple</td>
				<td>{{.InCouple}} </td>
			</tr>
			<tr>
				<td>Children</td>
				<td>{{.Children}}</td>
			</tr>
			<tr>
				<td>Shares</td>
				<td>{{.Shares}}</td>
			</tr>
			<tr>
				<td>Tax</td>
				<td><span style="color:green";>{{.Tax}}</span> {{.Currency}}</td>
			</tr>
			<tr>
				<td>Remainder</td>
				<td><span style="color:green";>{{.Remainder}}</span> {{.Currency}}</td>
			</tr>
		</tbody>
	</table>

	<hr>
	
	<h3 style="width: 70%;margin: auto;">{{.TaxDetails}}</h3>
	<table style="width: 70%;margin: auto;border: 1px solid black;border-collapse: collapse;">
		{{ range .Tranches }}
		 	<tr>
				{{ range . }}
					<td style="border: 1px solid black;">{{.}}</td>
				{{end}}
			</tr>
		{{ end }}
	</table>
	
	<br/>
	<br/>

	<span style="font-size:1.2em;">Mail sent by : <a href="https://github.com/NY-Daystar">NY-Daystar</a></span>
	<img width="50" height="50" src="https://avatars.githubusercontent.com/u/123415822"/> 
	`

	var isInCouple = language.No
	if user.IsInCouple {
		isInCouple = language.Yes
	}

	data := map[string]interface{}{
		"Prelude":    popup.BodyEntry.Text,
		"Year":       cfg.Tax.Year,
		"Currency":   params.Currency,
		"Income":     user.Income,
		"InCouple":   isInCouple,
		"Children":   user.Children,
		"Shares":     user.Shares,
		"Tax":        user.Tax,
		"Remainder":  user.Remainder,
		"Tranches":   tranches,
		"TaxResults": fmt.Sprintf("%s %s", language.Tax, language.Result),
		"TaxDetails": "Details",
	}

	t := template.Must(template.New("email").Parse(emailTmpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}

// getTaxTrancheResult get details of each tranche
func getTaxTrancheResult(result tax.Result, currency string, language settings.Yaml) [][]string {
	var header = []string{
		language.TaxHeaders.Header1,
		language.TaxHeaders.Header2,
		language.TaxHeaders.Header3,
		language.TaxHeaders.Header4,
		language.TaxHeaders.Header5,
	}

	// Create data to append on the table
	var data [][]string
	data = append(data, header)
	for i, val := range result.TaxTranches {
		index := i + 1

		var trancheNumber = fmt.Sprintf("Tranche %d", index)
		var min = fmt.Sprintf("%d %s", val.Tranche.Min, currency)
		var max = fmt.Sprintf("%d %s", val.Tranche.Max, currency)
		var rate = fmt.Sprintf("%d %%", val.Tranche.Rate)
		var tax = fmt.Sprintf("%d %s", int(val.Tax), currency)

		// handle max number
		if val.Tranche.Max > 1e10 {
			max = fmt.Sprintf("∞ %s", currency)
		}

		var line = make([]string, 5)
		line[0] = trancheNumber
		line[1] = min
		line[2] = max
		line[3] = rate
		line[4] = tax
		data = append(data, line)
	}

	// Add footer
	var footer = []string{
		language.Result,
		language.Remainder,
		fmt.Sprintf("%s %s", strconv.Itoa(int(result.Remainder)), currency),
		language.TotalTax,
		fmt.Sprintf("%s %s", strconv.Itoa(int(result.Tax)), currency),
	}

	data = append(data, footer)
	return data
}
