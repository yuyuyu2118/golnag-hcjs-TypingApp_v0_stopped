const temperatureChart = document.getElementById('temperatureChart')
const label_values = document.querySelectorAll('#label_value')
const data_values = document.querySelectorAll('#data_value')
const datasets = []

for (let i = 0; i < label_values.length; i++){
    data = data_values[i].value.replace(/ /g, ',')
    datasets[i] = {
        "label": label_values[i].value,
        "data": JSON.parse(data),
        "borderColor": `rgb(${(i + 1) * 50}, ${(i + 1) * 100}, ${(i + 1) * 100})`,
        "fill": false,
        "lineTension": 0.1
    }
}

new Chart(temperatureChart, {
    "type": "line",
    "data": {
        "labels": ["1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"],
        "datasets": datasets
    },
    "options": {
        scales: {
            yAxes: [{
                ticks: {
                    suggestedMax: 40,
                    suggestedMin: -20,
                }
            }]
        },
    }
})