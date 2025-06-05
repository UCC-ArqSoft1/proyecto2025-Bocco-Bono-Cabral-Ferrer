import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import '../Styles/index.css';
import Login from '../pages/login.jsx';
import Layout from '../components/Layout.jsx';
import Home from '../pages/home.jsx';
import Activities from '../pages/Activities.jsx';

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
