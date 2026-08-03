import React,{ useState, useEffect } from 'react';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS, CategoryScale, LinearScale, PointElement,
  LineElement, Title, Tooltip, Legend,
} from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

export default function SensorChart({ endpointUrl, title }) {
  const [sensorData, setSensorData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(endpointUrl)
      .then((res) => {
        if (!res.ok) throw new Error(`Failed to fetch from ${endpointUrl}`);
        return res.json();
      })
      .then((json) => {
        setSensorData(json.data || []);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [endpointUrl]);

  const chartData = {
    labels: sensorData.map((_, index) => `${index + 1}`),
    datasets: [
      {
        label: title,
        data: sensorData,
        borderColor: 'rgb(75, 192, 192)',
        backgroundColor: 'rgba(75, 192, 192, 0.2)',
        tension: 0.5,
      },
    ],
  };

  if (loading) return <p>Loading {title}...</p>;
  if (error) return <p style={{ color: 'red' }}>Error: {error}</p>;

  return (
    <div style={{ background: '#fff', padding: '15px', borderRadius: '8px', boxShadow: '0 2px 4px rgba(0,0,0,0.1)', marginBottom: '20px' }}>
      <Line data={chartData} options={{ responsive: true }} />
    </div>
  );
}
