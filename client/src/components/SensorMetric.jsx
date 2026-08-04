import React,{ useState, useEffect } from 'react';

export default function SensorChart({ sensor, metric }) {
  const [sensorData, setSensorData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const endpointUrl = `/api/sensors/${sensor}/${metric}`;
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

  if (loading) return <p>Loading {metric}...</p>;
  if (error) return <p style={{ color: 'red' }}>Error: {error}</p>;

  return (
    <div style={{ background: '#fff', padding: '15px', borderRadius: '8px', boxShadow: '0 2px 4px rgba(0,0,0,0.1)', marginBottom: '20px' }}>
      <span>{metric}: {sensorData}</span>
    </div>
  );
}
