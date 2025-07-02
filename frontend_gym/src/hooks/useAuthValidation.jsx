import { useNavigate } from 'react-router-dom';
import { useAuth } from './useAuth';

export const useAuthValidation = () => {
    const navigate = useNavigate();
    const { isAuthenticated, user } = useAuth();
    const isAdmin = user?.typeId === 1;

    const validateAuth = () => {
        if (!isAuthenticated) {
            navigate('/login');
            return false;
        }
        return true;
    };

    const validateAdmin = () => {
        if (!validateAuth()) {
            return false;
        }

        if (!isAdmin) {
            alert('No tienes permisos de administrador para acceder a esta página');
            navigate('/activities');
            return false;
        }

        return true;
    };

    const handleAuthError = (error) => {
        if (error.message.includes('sesión ha expirado') ||
            error.message.includes('token de autenticación') ||
            error.message.includes('inicia sesión')) {
            navigate('/login');
            return true; // Error manejado
        }
        return false; // Error no manejado
    };

    return {
        isAuthenticated,
        isAdmin,
        validateAuth,
        validateAdmin,
        handleAuthError
    };
}; 