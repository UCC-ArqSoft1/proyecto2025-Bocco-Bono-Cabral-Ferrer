import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import '../Styles/index.css';
import Login from '../pages/login.jsx';
import Register from '../pages/register.jsx';
import Layout from '../components/Layout.jsx';
import Home from '../pages/home.jsx';
import Activities from '../pages/Activities.jsx';
import ActivitiesAdmin from '../pages/activitiesAdmin.jsx';
import ActivityDetail from '../pages/ActivityDetail.jsx';
import { AuthProvider } from '../hooks/useAuth';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter>
      <AuthProvider>
        <ToastContainer position="top-right" autoClose={3000} />
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
      </AuthProvider>
    </BrowserRouter>
  </StrictMode>,
)
