package report

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type ReportData struct {
	Results   [][]string
	Historico []map[string]interface{}
}

func GenerateUltimateReport(results [][]string, historico []map[string]interface{}, filePath string) error {
	if len(results) > 0 {
		results = results[1:] // remove header
	}

	data := ReportData{
		Results:   results,
		Historico: historico,
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating HTML file: %v", err)
	}
	defer file.Close()

	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>API Test Dashboard</title>
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.1/dist/css/bootstrap.min.css" rel="stylesheet">
<link href="https://cdn.datatables.net/1.13.6/css/dataTables.bootstrap5.min.css" rel="stylesheet">
<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
<script src="https://code.jquery.com/jquery-3.7.1.min.js"></script>
<script src="https://cdn.datatables.net/1.13.6/js/jquery.dataTables.min.js"></script>
<script src="https://cdn.datatables.net/1.13.6/js/dataTables.bootstrap5.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/xlsx/dist/xlsx.full.min.js"></script>
<style>
body { font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif; background:#f8f9fa; color:#343a40; }
h1,h2{ text-align:center; margin:20px 0; }
.card { margin-bottom:15px; }
.table td, .table th { vertical-align: middle; }
.status-pass { color:#28a745; font-weight:bold; }
.status-fail { color:#dc3545; font-weight:bold; }
.message-cell { max-width:350px; overflow:hidden; text-overflow:ellipsis; white-space:nowrap; }
.message-cell.expanded { white-space:normal; }
.expand-btn { cursor:pointer; color:#0d6efd; text-decoration:underline; font-size:12px; }
.filter-container { display:flex; justify-content:center; gap:10px; flex-wrap:wrap; margin-bottom:20px; }
</style>
</head>
<body>
<div class="container">
<h1>📊 API Test Dashboard</h1>

<!-- KPIs -->
<div class="row text-center">
  <div class="col-md-3"><div class="card"><div class="card-body"><h5>Total Tests</h5><p class="fs-3">{{len .Results}}</p></div></div></div>
  <div class="col-md-3"><div class="card"><div class="card-body"><h5>Passed</h5><p class="fs-3">{{passCount .Results}}</p></div></div></div>
  <div class="col-md-3"><div class="card"><div class="card-body"><h5>Failed</h5><p class="fs-3">{{failCount .Results}}</p></div></div></div>
  <div class="col-md-3"><div class="card"><div class="card-body"><h5>Overall Status</h5><p class="fs-3">{{generalStatus .Results}}</p></div></div></div>
</div>

<!-- Filters -->
<div class="filter-container">
  <select id="filterTestCase" class="form-select w-auto"><option value="">All TestCases</option></select>
  <select id="filterResult" class="form-select w-auto"><option value="">All Results</option><option value="true">Passed</option><option value="false">Failed</option></select>
  <button class="btn btn-primary" onclick="exportTableToExcel('resultsTable','API_Test_Report')">Export Excel</button>
</div>

<!-- Charts -->
<div class="row">
  <div class="col-md-6"><h3 class="text-center">Tests per TestCase</h3><canvas id="barChart" height="200"></canvas></div>
  <div class="col-md-6"><h3 class="text-center">Result Distribution</h3><canvas id="pieChart" height="200"></canvas></div>
</div>

<!-- Results Table -->
<div class="table-responsive mt-4">
<table id="resultsTable" class="table table-striped table-bordered table-hover">
<thead class="table-dark"><tr><th>TestId</th><th>TestCase</th><th>Result</th><th>Message</th></tr></thead>
<tbody>
{{range $index, $row := .Results}}
<tr>
<td>{{index $row 0}}</td>
<td>{{index $row 1}}</td>
<td class="{{if eq (index $row 2) "true"}}status-pass{{else}}status-fail{{end}}">{{if eq (index $row 2) "true"}}Passed{{else}}Failed{{end}}</td>
<td><div id="msg-{{$index}}" class="message-cell">{{index $row 3}}</div>{{if gt (len (index $row 3)) 50}} <span class="expand-btn" onclick="toggleMessage('msg-{{$index}}', this)">See More</span>{{end}}</td>
</tr>
{{end}}
</tbody>
</table>
</div>
</div>

<script>
function toggleMessage(id, btn){
  const cell = document.getElementById(id);
  cell.classList.toggle('expanded');
  btn.innerText = cell.classList.contains('expanded') ? 'See Less' : 'See More';
}
function exportTableToExcel(tableID, filename=''){
  var table = document.getElementById(tableID);
  var wb = XLSX.utils.table_to_book(table,{sheet:"Sheet1"});
  XLSX.writeFile(wb, filename+'.xlsx');
}

$(document).ready(function(){
  var table = $('#resultsTable').DataTable({pageLength:10});
  var testCaseSet = new Set();
  table.column(1).data().each(function(value){ testCaseSet.add(value); });
  testCaseSet.forEach(function(tc){ $('#filterTestCase').append('<option value="'+tc+'">'+tc+'</option>'); });
  $('#filterTestCase,#filterResult').on('change', function(){
    var tc = $('#filterTestCase').val();
    var res = $('#filterResult').val();
    table.rows().every(function(){
      var show=true;
      if(tc && this.data()[1]!=tc) show=false;
      if(res && this.data()[2]!=res) show=false;
      $(this.node()).toggle(show);
    });
  });
});

// Chart Data
const historico = {{marshal .Historico}};
let grouped={}, totalPassed=0, totalFailed=0;
historico.forEach(h=>{
  const tc = h.test_case || h.TestCase;
  if(!grouped[tc]) grouped[tc]={passed:0, failed:0};
  if(h.result || h.Result){ grouped[tc].passed++; totalPassed++; }else{ grouped[tc].failed++; totalFailed++; }
});
const labels = Object.keys(grouped);
const passedData = labels.map(l=>grouped[l].passed);
const failedData = labels.map(l=>grouped[l].failed);

// Bar Chart
new Chart(document.getElementById('barChart').getContext('2d'),{
  type:'bar',
  data:{labels:labels,datasets:[{label:'Passed',data:passedData,backgroundColor:'#28a745'},{label:'Failed',data:failedData,backgroundColor:'#dc3545'}]},
  options:{responsive:true,plugins:{legend:{position:'top'}},scales:{y:{beginAtZero:true,stepSize:1}}}
});

// Pie Chart
new Chart(document.getElementById('pieChart').getContext('2d'),{
  type:'pie',
  data:{labels:['Passed','Failed'],datasets:[{data:[totalPassed,totalFailed],backgroundColor:['#28a745','#dc3545']}]},
  options:{responsive:true,plugins:{legend:{position:'top'}}}
});
</script>
</body>
</html>
`

	funcMap := template.FuncMap{
		"passCount": func(results [][]string) int {
			count := 0
			for _, r := range results {
				if r[2] == "true" {
					count++
				}
			}
			return count
		},
		"failCount": func(results [][]string) int {
			count := 0
			for _, r := range results {
				if r[2] != "true" {
					count++
				}
			}
			return count
		},
		"generalStatus": func(results [][]string) string {
			total := len(results)
			fails := 0
			for _, r := range results {
				if r[2] != "true" {
					fails++
				}
			}
			ratio := float64(fails) / float64(total)
			switch {
			case ratio == 0:
				return "✔️ Excellent"
			case ratio < 0.2:
				return "⚠️ Warning"
			default:
				return "❌ Fail"
			}
		},
		"marshal": func(v interface{}) template.JS {
			b, _ := json.Marshal(v)
			return template.JS(b)
		},
	}

	t, err := template.New("report").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("template parse error: %v", err)
	}

	return t.Execute(file, data)
}
