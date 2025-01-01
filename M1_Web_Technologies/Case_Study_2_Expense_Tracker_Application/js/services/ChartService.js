import { categories } from "../config/categories.js";

export class ChartService {
  constructor() {
    this.pieChart = null;
    this.barChart = null;
    this.initializeCharts();
  }

  initializeCharts() {
    const pieCtx = document.getElementById("pieChart").getContext("2d");
    const barCtx = document.getElementById("barChart").getContext("2d");

    this.pieChart = new Chart(pieCtx, {
      type: "pie",
      data: {
        labels: [],
        datasets: [
          {
            data: [],
            backgroundColor: [],
          },
        ],
      },
      options: {
        responsive: true,
        plugins: {
          legend: { position: "bottom" },
        },
      },
    });

    this.barChart = new Chart(barCtx, {
      type: "bar",
      data: {
        labels: [],
        datasets: [
          {
            label: "Expenses by Category",
            data: [],
            backgroundColor: [],
          },
        ],
      },
      options: {
        responsive: true,
        scales: {
          y: { beginAtZero: true },
        },
      },
    });
  }

  updateCharts(categoryTotals) {
    const chartData = this.prepareChartData(categoryTotals);
    this.updateChartData(this.pieChart, chartData);
    this.updateChartData(this.barChart, chartData);
  }

  prepareChartData(categoryTotals) {
    return {
      labels: Object.keys(categoryTotals),
      data: Object.values(categoryTotals),
      colors: Object.keys(categoryTotals).map(
        (category) => categories[category]
      ),
    };
  }

  updateChartData(chart, { labels, data, colors }) {
    chart.data.labels = labels;
    chart.data.datasets[0].data = data;
    chart.data.datasets[0].backgroundColor = colors;
    chart.update();
  }
}
