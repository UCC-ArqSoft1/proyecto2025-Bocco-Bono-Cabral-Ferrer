import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import '../Styles/index.css';
import Login from '../pages/login.jsx';
import Register from '../pages/register.jsx';
import Layout from '../components/Layout.jsx';
import Home from '../pages/home.jsx';
import Activities from '../pages/Activities.jsx';
import ActivitiesAdmin from '../pages/activitiesAdmin.jsx';
import ActivityDetail from '../pages/ActivityDetail.jsx';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path='/activities' element={<Activities />} />
          <Route path='/activity/:id' element={<ActivityDetail />} />
          <Route path='/admin/activities' element={<ActivitiesAdmin />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
