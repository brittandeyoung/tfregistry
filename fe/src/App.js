import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Header from "./Components/Header"; 
import NamespaceList from "./Components/NamespaceList"; 
import ModuleList from "./Components/ModuleList"; 

import './App.css';

function App() {
  return (

    <div>
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
