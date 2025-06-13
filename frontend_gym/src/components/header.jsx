import Login from "../pages/login.jsx";
import "../Styles/header.css";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

const Header = () => {
    const navigate = useNavigate();
    const { isAuthenticated, logout } = useAuth();

    const handleLogout = () => {
        if (window.confirm('¿Estás seguro que deseas cerrar sesión?')) {
            logout();
        }
    };

    return (
        <header>
            <h1>GYM</h1>
            <nav>
                <a href="/">Home</a>
                {isAuthenticated ? (
                    <>
                        <a href="/activities">Actividades</a>
                        <button onClick={handleLogout}>Cerrar Sesión</button>
                    </>
                ) : (
                    <>
                        <a href="/login">Login</a>
                        <a href="/register">Registro</a>
                        <a href="/activities">Ver Actividades</a>
                    </>
                )}
            </nav>
        </header>
    );
}

export default Header;