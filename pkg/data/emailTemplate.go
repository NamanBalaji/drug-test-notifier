package data

const emailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Notification about drug test</title>
    <style>
        body { font-family: Arial, sans-serif; }
        .container { width: 80%; margin: 0 auto; }
        .header { background-color: #f2f2f2; padding: 20px; text-align: center; }
        .content { padding: 20px; }
        .footer { background-color: #f2f2f2; padding: 10px; text-align: center; }
        .selected { color: red; font-weight: bold; }
        .not-selected { color: green; font-weight: bold; }
        .confirmation { font-size: 24px; font-weight: bold; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1 class="{{if .Selected}}selected{{else}}not-selected{{end}}">
				{{if .Selected}}IMPORTANT{{else}}HOORAY{{end}}<br>
                You have {{if .Selected}}been selected{{else}}not been selected{{end}} today ({{.Date}})
            </h1>
        </div>
        <div class="content">
            {{if .Selected}}
                <p>Please check your Recovery Trek account for more details.</p>
            {{end}}
            <p class="confirmation">
                Your confirmation number is #{{.ConfirmationNumber}}
            </p>
            <p>Please note you have {{.BillsDue}} bills due.</p>
            <p>{{.Message}}</p>
        </div>
        <div class="footer">
            <p>I love you <span style="color:red;">&hearts;</span></p>
        </div>
    </div>
</body>
</html>
`
