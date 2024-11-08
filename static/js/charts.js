debugger;

const ctx = document.getElementById("temperatureChart").getContext("2d");
const temperatureData = JSON.parse(
  document.getElementById("temperatureData").textContent
);
const feelsLikeData = JSON.parse(
  document.getElementById("feelsLikeData").textContent
);
const labels = JSON.parse(document.getElementById("chartLabels").textContent);

// Creating gradients
const temperatureGradient = ctx.createLinearGradient(0, 0, 0, 400);
temperatureGradient.addColorStop(0, "rgba(75, 192, 192, 0.5)");
temperatureGradient.addColorStop(1, "rgba(75, 192, 192, 0)");

const feelsLikeGradient = ctx.createLinearGradient(0, 0, 0, 400);
feelsLikeGradient.addColorStop(0, "rgba(255, 99, 132, 0.5)");
feelsLikeGradient.addColorStop(1, "rgba(255, 99, 132, 0)");

const temperatureChart = new Chart(ctx, {
  type: "line",
  data: {
    // Labels basically what the label will look like.
    labels: labels,
    // datasets with what it would look like with various customizations.
    datasets: [
      {
        label: "Feels Like",
        data: feelsLikeData,
        borderColor: "rgba(255, 99, 132, 1)",
        backgroundColor: feelsLikeGradient,
        fill: true,
        borderWidth: 2,
        tension: 0.3,
        pointBackgroundColor: "rgba(255, 99, 132, 1)",
        pointRadius: 4,
        pointHoverRadius: 8,
        hoverBorderWidth: 2,
        hoverBorderColor: "rgba(255, 99, 132, 0.8)",
      },
      {
        label: "Temperature",
        data: temperatureData,
        borderColor: "rgba(75, 192, 192, 1)",
        backgroundColor: temperatureGradient,
        fill: true,
        borderWidth: 2,
        tension: 0.3,
        pointBackgroundColor: "rgba(75, 192, 192, 1)",
        pointRadius: 4,
        pointHoverRadius: 8,
        hoverBorderWidth: 2,
        hoverBorderColor: "rgba(75, 192, 192, 0.8)",
      },
    ],
  },
  options: {
    responsive: true,
    // Axis customization.
    scales: {
      x: {
        grid: { display: false },
        ticks: {
          color: "rgb(102, 102, 102)",
          font: { size: 12 },
        },
      },
      y: {
        beginAtZero: true,
        grid: { color: "rgba(200, 200, 200, 0.2)" },
        ticks: {
          color: "rgb(102, 102, 102)",
          font: { size: 12 },
          callback: function (value) {
            return `${value}°C`;
          },
        },
      },
    },
    plugins: {
      // For Legend customization and tooltip customization.
      legend: {
        display: true,
        labels: {
          color: "rgb(102, 102, 102)",
          font: { size: 14 },
          padding: 20,
        },
      },
      tooltip: {
        enabled: true,
        mode: "nearest",
        backgroundColor: "rgba(0,0,0,0.7)",
        titleFont: { size: 14, weight: "bold" },
        bodyFont: { size: 12 },
        footerFont: { style: "italic" },
        callbacks: {
          label: function (context) {
            return context.dataset.label + ": " + context.raw.toFixed(2) + "°C";
          },
        },
      },
    },
    animations: {
      // Animation effects.
      tension: {
        duration: 1000,
        easing: "easeInOutQuad",
        from: 0.4,
        to: 0.1,
        loop: true,
      },
    },
    hover: {
      // Hover customizations.
      mode: "nearest",
      intersect: true,
      onHover: function (e) {
        const point = temperatureChart.getElementsAtEventForMode(
          e,
          "nearest",
          { intersect: true },
          false
        );
        if (point.length) e.native.target.style.cursor = "pointer";
        else e.native.target.style.cursor = "default";
      },
    },
  },
});
