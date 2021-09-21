import logo from './logo.svg';
import './App.css';
import NavBar from "./components/navbar";
import React from "react";
import Map from "./components/map";
// import MapHCH from "./components/mapHCH";

function App() {
  return (
    <div className="App">
      <NavBar/>
      {/*<MapHCH/>*/}
      <Map/>
    </div>
  );
}

export default App;
