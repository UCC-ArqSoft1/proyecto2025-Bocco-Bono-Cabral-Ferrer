import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './index.css';
import Login from './Login.jsx';
import Layout from './layout.jsx';
import Home from './home.jsx';
import Activities from './Activities.jsx';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="/Login" element={<Login />} />
          <Route path='/Activities' element={<Activities />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
