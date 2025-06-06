import Login from "../pages/login.jsx";
import "../Styles/header.css";
import { useNavigate } from "react-router-dom";

const Header = () => {
    const navigate = useNavigate();
    const isloggedin = localStorage.getItem("isLogin") === "true";

    const logout = () => {
        localStorage.removeItem("isLogin");
        localStorage.removeItem("token");
        localStorage.removeItem("userId");
        localStorage.removeItem("userTypeId");
        navigate("/");
    }

    return (
        <header>
            <h1>GYM</h1>
            <nav>
                <a href="/">Home</a>
                {isloggedin ? (
                    <>
                        <a href="/activities">Mis Actividades</a>
                        <button onClick={logout}>Cerrar Sesión</button>
                    </>
                ) : (
                    <>
                        <a href="/login">Login</a>
                        <a href="/activities">Ver Actividades</a>
                    </>
                )}
            </nav>
        </header>
    );
}

export default Header;