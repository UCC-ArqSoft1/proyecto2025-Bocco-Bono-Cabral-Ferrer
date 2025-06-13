import { createContext, useContext, useState, useEffect } from 'react';
import axios from 'axios';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    // Limpiar localStorage al inicio
    useEffect(() => {
        // Mantener solo el token y userInfo si existen
        const token = localStorage.getItem('token');
        const userInfo = localStorage.getItem('userInfo');

        // Limpiar todo el localStorage
        localStorage.clear();

        // Restaurar solo token y userInfo si existen
        if (token) localStorage.setItem('token', token);
        if (userInfo) localStorage.setItem('userInfo', userInfo);
    }, []);

    const isTokenExpired = (token) => {
        try {
            const decoded = jwtDecode(token);
            return decoded.exp * 1000 < Date.now();
        } catch {
            return true;
        }
    };

    const setAuthToken = (token) => {
        if (token) {
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
        } else {
            delete axios.defaults.headers.common['Authorization'];
        }
    };

    const logout = () => {
        localStorage.clear();
        setAuthToken(null);
        setUser(null);
        navigate('/login');
    };

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            if (isTokenExpired(token)) {
                // Token expirado, limpiar y redirigir al login
                logout();
                toast.error('Tu sesión ha expirado. Por favor, inicia sesión nuevamente.');
            } else {
                const userInfo = JSON.parse(localStorage.getItem('userInfo') || '{}');
                setUser(userInfo);
                setAuthToken(token);
            }
        }
        setLoading(false);
    }, [navigate]);

    // Interceptor para manejar errores de token
    useEffect(() => {
        const interceptor = axios.interceptors.response.use(
            (response) => response,
            (error) => {
                // Solo manejar expiración de sesión si el usuario está autenticado
                if (error.response?.status === 401 && user) {
                    logout();
                    toast.error('Tu sesión ha expirado. Por favor, inicia sesión nuevamente.');
                }
                return Promise.reject(error);
            }
        );

        return () => {
            axios.interceptors.response.eject(interceptor);
        };
    }, [user]); // Agregamos user como dependencia

    const login = async (email, password) => {
        try {
            const response = await axios.post('http://localhost:8080/users/login', {
                email,
                password,
            });

            const { token, user_id, user_type_id } = response.data;
            const userInfo = { id: user_id, typeId: user_type_id };

            // Limpiar localStorage antes de guardar nuevos datos
            localStorage.clear();

            localStorage.setItem('token', token);
            localStorage.setItem('userInfo', JSON.stringify(userInfo));
            setAuthToken(token);
            setUser(userInfo);
            return true;
        } catch (error) {
            toast.error(error.response?.data?.error || 'Error al iniciar sesión');
            return false;
        }
    };

    const register = async (userData) => {
        try {
            const response = await axios.post('http://localhost:8080/users/register', userData);
            toast.success('Registro exitoso! Por favor inicia sesión.');
            return true;
        } catch (error) {
            toast.error(error.response?.data?.error || 'Error al registrarse');
            return false;
        }
    };

    const value = {
        user,
        isAuthenticated: !!user && !!localStorage.getItem('token'),
        login,
        logout,
        register,
        loading
    };

    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
}; 