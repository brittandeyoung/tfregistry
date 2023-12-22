import {BrowserRouter, Routes, Route} from "react-router-dom";
import Header from './components/Header/Header';
import ModuleList from './components/ModuleList/ModuleList';
import NamespaceList from './components/NamespaceList/NamespaceList';

import './App.css';

function App() {
  return (
    <div className="App">
      <Header />
      <div className="ui container">
        <div className="wrapper">
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<NamespaceList />}></Route>
              <Route path="/:namespace" element={<ModuleList />}></Route>
            </Routes>
          </BrowserRouter>
        </div>
      </div>
    </div>
  );
}

export default App;
