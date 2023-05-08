package template

import "text/template"

var GetWeather = template.Must(template.New("get_weather").Parse(`
<b>{{ .city_name }}, {{ .country }}</b>

{{ .icon }} <b>{{ .temp }} °C</b>

<b>По ощущениям {{ .feels_like }} °C. {{ .description }}</b>
{{ .date }}
`))
