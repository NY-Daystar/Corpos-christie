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
func FormatMail(user *user.User, cfg *config.Config, params settings.Settings, popup *widgets.MailPopup) string {
	var result = tax.CalculateTax(user, cfg)
	var tranches = getTaxTrancheResult(result, cfg.Tax.Year)

	// TODO mettre dans une langue spécifique les labels
	// TODO voir pour mettre du CSS surtout pour la partie tax et remainder a mettre en vert
	// TODO css pour les bordures du tableau du détail
	// TODO mettre des h2 ou h3 entre les parties

	// TODO mettre le .Prelude en placeholder

	// TODO faire une méthode pour ajouter une signature

	var emailTmpl = `{{.Prelude}}
	<br/>
	Hi {{.Name}}, Here's your result for <em>{{.Year}}</em>
	<hr>
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
		<td>{{.Tax}} {{.Currency}}</td>
		</tr>
		<tr>
		<td>Remainder</td>
		<td>{{.Remainder}} {{.Currency}}</td>
		</tr>
	</tbody>
	</table>
	<hr>
	<ul>
		{{ range .Tranches }}
		 	<tr>
				{{ range . }}
					<td>{{.}}</td>
				{{end}}
			</tr>
		{{ end }}
	</ul>`

	// TODO language
	var isInCouple = "No"
	if user.IsInCouple {
		isInCouple = "Yes"
	}

	data := map[string]interface{}{
		"Prelude":   popup.BodyEntry.Text,
		"Year":      cfg.Tax.Year,
		"Name":      popup.Username,
		"Currency":  params.Currency,
		"Income":    user.Income,
		"InCouple":  isInCouple,
		"Children":  user.Children,
		"Shares":    user.Shares,
		"Tax":       user.Tax,
		"Remainder": user.Remainder,
		"Tranches":  tranches,
	}

	t := template.Must(template.New("email").Parse(emailTmpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}

// getTaxTrancheResult get details of each tranche
func getTaxTrancheResult(result tax.Result, year int) [][]string {
	// Setting header
	var header = []string{"Tranche", "Min", "Max", "Rate", "Tax"} // TODO LANGUAGE

	// Create data to append on the table
	var data [][]string
	data = append(data, header)
	for i, val := range result.TaxTranches {
		index := i + 1

		var trancheNumber = fmt.Sprintf("Tranche %d", index)
		var min = fmt.Sprintf("%d €", val.Tranche.Min)
		var max = fmt.Sprintf("%d €", val.Tranche.Max)
		var rate = fmt.Sprintf("%d %%", val.Tranche.Rate)
		var tax = fmt.Sprintf("%d €", int(val.Tax))

		// handle max number
		if val.Tranche.Max > 1e10 {
			max = fmt.Sprint("∞ €")
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
		"Result",
		"Remainder",
		fmt.Sprintf("%s €", strconv.Itoa(int(result.Remainder))),
		"Total Tax",
		fmt.Sprintf("%s €", strconv.Itoa(int(result.Tax))),
	}

	data = append(data, footer)
	return data
}
