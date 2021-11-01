import React from "react";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import NavBar from "./components/navbar";
import Map from "./components/map";
// import MapHCH from "./components/mapHCH";
import './App.css';

function App() {
  return (
      <React.Fragment>
          <BrowserRouter>
              <div className="App">
                  <NavBar/>
                  {/*<MapHCH/>*/}
                  <main>
                      <Switch>
                          <Route path="/" component={Map}/>
                      </Switch>
                  </main>
              </div>
          </BrowserRouter>
      </React.Fragment>
  );
}

export default App;
