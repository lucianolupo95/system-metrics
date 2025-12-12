import { useEffect, useState } from "react";
import { Card, CardContent, Typography, Container, Box } from "@mui/material";

function formatUptime(seconds) {
  const total = Math.floor(Number(seconds));
  const hrs = Math.floor(total / 3600);
  const mins = Math.floor((total % 3600) / 60);
  const secs = total % 60;

  return `${hrs}h ${mins}m ${secs}s`;
}
function cpuColor(cpu) {
  if (cpu < 30) return "green";
  if (cpu < 70) return "orange";
  return "red";
}

function App() {
  const [metrics, setMetrics] = useState(null);

  useEffect(() => {
    const fetchMetrics = () => {
      fetch("http://localhost:8080/metrics")
        .then((res) => {
          if (res.status === 204) return null;
          return res.json();
        })
        .then((data) => {
          setMetrics(data);
        })
        .catch((err) => {
          console.error("Error fetching metrics:", err);
        });
    };

    fetchMetrics(); // primera llamada inmediata
    const interval = setInterval(fetchMetrics, 3000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <div style={{ textAlign: "center" }}>
        <h1>System Metrics Dashboard</h1>

        {!metrics && <p>Waiting for metrics...</p>}

        {metrics && (
          <div
            style={{
              display: "flex",
              gap: "16px",
              justifyContent: "center",
              flexWrap: "wrap",
              marginTop: "24px",
            }}
          >
            <Card>
              <CardContent>
                <Typography variant="h6">Uptime</Typography>
                <Typography>{formatUptime(metrics.uptime_seconds)}</Typography>
              </CardContent>
            </Card>

            <Card>
              <CardContent>
                <Typography variant="h6">Memory Used</Typography>
                <Typography>{metrics.memory_used_kb} KB</Typography>
              </CardContent>
            </Card>

            <Card>
              <CardContent>
                <Typography variant="h6">CPU Usage</Typography>
                <Typography style={{ color: cpuColor(metrics.cpu_usage_pct) }}>
                  {metrics.cpu_usage_pct.toFixed(2)} %
                </Typography>
              </CardContent>
            </Card>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
