package gscan

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"time"

	enums "github.com/MitchellWT/gscan/internal/enums"
)

func RawExportToJSON(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	collectedMap := CollectRaw(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".json"
	jsonData := ExportRaw{
		UnixStartTime: intervalStart,
		UnixEndTime:   intervalEnd,
		RootDir:       rootDir,
		ScanFiles:     collectedMap,
	}
	jsonBytes, err := json.Marshal(jsonData)
	ErrorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	ErrorCheck(err)

	return fileName
}

func TotalExportToJSON(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotal(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".json"
	jsonData := ExportCollectedRaw{
		UnixStartTime: intervalStart,
		UnixEndTime:   intervalEnd,
		RootDir:       rootDir,
		TotalDiff:     totalDiff,
	}
	jsonBytes, err := json.Marshal(jsonData)
	ErrorCheck(err)

	err = os.MkdirAll(outDir, 0755)
	ErrorCheck(err)

	err = os.WriteFile(fileName, jsonBytes, 0766)
	ErrorCheck(err)

	return fileName
}

func generateRandomLineColour() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func TotalExportToHTML(rootDir string, outDir string, interval enums.Interval) string {
	currentTime := time.Now().Unix()
	intervalStart := interval.GetStart()
	intervalEnd := interval.GetEnd()
	totalDiff := CollectTotal(rootDir, intervalStart, intervalEnd)
	// Builds file name to save data
	fileName := outDir + "export-" + fmt.Sprint(currentTime) + ".html"
	labelSlice := make([]string, 0)
	dataSlice := make([]float32, 0)

	for unixTime, totalSize := range totalDiff {
		outputTime := time.Unix(unixTime, 0).Format("15:04 02/01/06")
		outputSize := float32(totalSize) / 1024
		labelSlice = append(labelSlice, string(outputTime))
		dataSlice = append(dataSlice, outputSize)
	}

	templateData := TotalHTMLTemplateData{
		Title:       rootDir + " Export Graph",
		GraphLabels: labelSlice,
		DataSets: []HTMLTemplateDataSet{
			HTMLTemplateDataSet{
				Label:      rootDir,
				LineColour: generateRandomLineColour(),
				Data:       dataSlice,
			},
		},
	}

	template, err := template.New("html-export").Parse(`
		<!DOCTYPE html>
		<head>
    	<meta charset="utf-8"/>

    	<link rel="icon" href="data:,"/>
    	<title>{{ .Title}}</title>

    	<style media="screen">
        	body {
            	margin: 0;
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

    	<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.7.0/chart.min.js" integrity="sha512-TW5s0IT/IppJtu76UbysrBH9Hy/5X41OTAbQuffZFU6lQ1rdcLHzpU5BzVvr/YFykoiMYZVWlr/PX1mDcfM9Qg==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
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
	ErrorCheck(err)

	exportFile, err := os.Create(fileName)
	ErrorCheck(err)

	err = template.Execute(exportFile, templateData)
	ErrorCheck(err)

	return fileName
}
