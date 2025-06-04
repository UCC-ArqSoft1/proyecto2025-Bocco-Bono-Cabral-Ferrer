import { createContext, useContext, useState, useEffect } from 'react';
import axios from 'axios';
import { toast } from 'react-toastify';

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            const userInfo = JSON.parse(localStorage.getItem('userInfo') || '{}');
            setUser(userInfo);
        }
        setLoading(false);
    }, []);

    const login = async (email, password) => {
        try {
            const response = await axios.post('http://localhost:8080/login', {
                email,
                password,
            });

            const { token, user_id, user_type_id } = response.data;
            const userInfo = { id: user_id, typeId: user_type_id };

            localStorage.setItem('token', token);
            localStorage.setItem('userInfo', JSON.stringify(userInfo));

            // Set default authorization header for all future requests
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

            setUser(userInfo);
            return true;
        } catch (error) {
            toast.error(error.response?.data?.error || 'Error al iniciar sesión');
            return false;
        }
    };

    const register = async (userData) => {
        try {
            const response = await axios.post('http://localhost:8080/register', userData);
            toast.success('Registro exitoso! Por favor inicia sesión.');
            return true;
        } catch (error) {
            toast.error(error.response?.data?.error || 'Error al registrarse');
            return false;
        }
    };

    const logout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('userInfo');
        delete axios.defaults.headers.common['Authorization'];
        setUser(null);
    };

    const value = {
        user,
        isAuthenticated: !!user,
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