import React, { useState, useEffect } from 'react';
import SensorChart from './components/SensorChart';
import SensorMetric from "./components/SensorMetric.jsx";

function App() {
  const [sensors, setSensors] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [selectedValue, setSelectedValue] = useState("");
  const handleChange = (e) => {
    setSelectedValue(e.target.value)
  };

  const endpointUrl = `/api/sensors`;
  useEffect(() => {
    if (sensors.length > 0 && !selectedValue) {
      setSelectedValue(sensors[0]);
    }

    fetch(endpointUrl)
      .then((res) => {
        if (!res.ok) throw new Error(`Failed to fetch from ${endpointUrl}`);
        return res.json();
      })
      .then((json) => {
        setSensors(json.data || []);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [sensors, selectedValue]);


  return (
    <div style={{ textAlign: 'center', display: 'flex', flexDirection: 'column', alignItems: 'middle', justifyContent: 'center', adding: '20px', fontFamily: 'sans-serif', maxWidth: '800px', margin: '0 auto' }}>
      <h1>Sensor Dashboard</h1>
      <label htmlFor="sensor-select">Choose a sensor:</label>

      <select id="sensor-select" value={selectedValue} onChange={handleChange}>
        {sensors.map((sensor) => (
          <option key={sensor} value={sensor}>
            {sensor}
          </option>
        ))}
      </select>

      {selectedValue && 
        <div>
          <h4>{selectedValue}</h4>
          <div style={{ display: 'inline-flex', justifyContent: 'center', width: '100%' }}>
            <SensorMetric
              sensor={selectedValue}
              metric="sum"/>

            <SensorMetric
              sensor={selectedValue}
              metric="avg"/>
          </div>

          <SensorChart
            endpointUrl={`/api/sensors/${selectedValue}`}
            title={selectedValue}/>
        </div>}
    </div>
  );
}

export default App;
