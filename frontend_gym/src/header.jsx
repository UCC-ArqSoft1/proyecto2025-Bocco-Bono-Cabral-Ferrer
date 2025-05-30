import Login from "./login";
import "./Header.css";
import { useNavigate } from "react-router-dom";
const Header = () => {
    const navigate = useNavigate();
    const isloggedin = localStorage.getItem("islogin") === "true";

    const logout = () => {
        localStorage.removeItem("islogin");
        navigate("/");
    }
    return (
        <header>
            <h1>GYM</h1>
            <nav>
                <a href="/">Home</a>
                {isloggedin ? (<button onClick={logout}>
                    Cerrar Sesión
                </button>) :
                    (
                        <a href="/login">Login</a>
                    )}

                <a href="/activities">Actividades</a>

            </nav>
        </header>
    );
}

export default Header;