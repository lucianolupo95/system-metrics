import { useEffect, useState } from "react";

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
    <div>
      <h1>System Metrics Dashboard</h1>

      {!metrics && <p>Waiting for metrics...</p>}

      {metrics && <pre>{JSON.stringify(metrics, null, 2)}</pre>}
    </div>
  );
}

export default App;
