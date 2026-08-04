import React, { useState, useEffect } from 'react';
import SensorChart from './components/SensorChart';
import SensorMetric from "./components/SensorMetric.jsx";

function App() {
  const [sensorData, setSensorData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const sensor = 'sensor-01';

  return (
    <div style={{ textAlign: 'center', display: 'flex', flexDirection: 'column', alignItems: 'middle', justifyContent: 'center', adding: '20px', fontFamily: 'sans-serif', maxWidth: '800px', margin: '0 auto' }}>
      <h1>Sensor Dashboard</h1>
      <h4>{sensor}</h4>

      <div style={{ display: 'inline-flex', justifyContent: 'center', width: '100%' }}>
        <SensorMetric
          sensor={sensor}
          metric="sum"/>

        <SensorMetric
          sensor={sensor}
          metric="avg"/>
      </div>

      <SensorChart
        endpointUrl={`/api/sensors/${sensor}`}
        title={sensor}/>  
    </div>
  );
}

export default App;
