package gscan

import "os"

var LibDir = "/var/lib/gscan/"

func preSetupCheck() bool {
	_, err := os.Stat(LibDir + "data/")
	if err != nil && os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(LibDir + "templates/")
	if err != nil && os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(LibDir + "templates/template.html")
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func createDirs() {
	err := os.MkdirAll(LibDir+"data/", 0755)
	ErrorCheck(err)

	err = os.MkdirAll(LibDir+"templates/", 0755)
	ErrorCheck(err)
}

func writeTemplates() {
	templateBytes := []byte(`
<!DOCTYPE html>
<head>
    <meta charset="utf-8"/>

    <link rel="icon" href="data:,"/>
    <title>{{ .Title}}</title>

    <style media="screen">
    	body {
            margin: 0;
            font-family: sans-serif;
    	}

    	h1 {
        	text-align: center;
        	margin: 0;
        	padding: 25px 0;
    	}

    	#chart {
        	padding: 25px;
        	margin: 0 auto;
        	width: 90vw !important;
    	}

    	.body-container {
        	height: 100vh;
    	}
    </style>

    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.7.0/chart.min.js"
        integrity="sha512-TW5s0IT/IppJtu76UbysrBH9Hy/5X41OTAbQuffZFU6lQ1rdcLHzpU5BzVvr/YFykoiMYZVWlr/PX1mDcfM9Qg=="
        crossorigin="anonymous" referrerpolicy="no-referrer"></script>
</head>
<body>
    <div class="body-container">
    	<h1>{{ .Title}}</h1>
    	<canvas id="chart"></canvas>
    </div>

    <script>
    	const fontSize = 18;
    	const ctx = document.getElementById('chart').getContext('2d');
    	const labels = {{ .GraphLabels}};
    	const options = {
        	scales: {
            	x: {
                	title: {
                    	display: true,
                    	text: "Time (hh:ss dd/mm/yy)",
                    	font: {
                        	size: fontSize
                    	}
                	},
                	ticks: {
                    	font: {
                        	size: fontSize
                    	}
                	}
            	},
            	y: {
                	title: {
                    	display: true,
                    	text: "Size (KB)",
                    	font: {
                        	size: fontSize
                    	}

                	},
                	ticks: {
                    	font: {
                        	size: fontSize
                    	}
                	}
            	}
        	},
        	plugins: {
            	legend: {
                	labels: {
                    	font: {
                        	size: fontSize
                    	}
                	}
            	}
        	}
    	}
    	const data = {
        	labels: labels,
    		datasets: [
    		{{ range .DataSets}}
    		{
            	label: {{ .Label}},
            	data: {{ .Data}},
            	fill: false,
            	backgroundColor: {{ .LineColour}},
            	borderColor: {{ .LineColour}},
            	borderWidth: 2.5,
            	tension: 0.1
        	},
    		{{ end}}
    		]
    	};
    	const chart = new Chart(ctx, {
        	type: 'line',
        	data: data,
        	options: options
    	});
    </script>
</body>
`)

	err := os.WriteFile(LibDir+"templates/template.html", templateBytes, 0644)
	ErrorCheck(err)
}

func Setup() {
	if preSetupCheck() {
		return
	}

	createDirs()
	writeTemplates()
}
